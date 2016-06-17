package containerreaper

import (
	"errors"
	"time"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
	"github.com/pivotal-golang/lager"
)

type ContainerReaper interface {
	Run() error
}

//go:generate counterfeiter . ContainerReaperDB

type ContainerReaperDB interface {
	FindJobIDForBuild(buildID int) (int, bool, error)
	GetContainersWithInfiniteTTL() ([]db.SavedContainer, error)
	FindContainersFromSuccessfulBuildsWithInfiniteTTL() ([]db.SavedContainer, error)
	FindContainersFromUnsuccessfulBuildsWithInfiniteTTL() ([]db.SavedContainer, error)
	UpdateExpiresAtOnContainer(handle string, ttl time.Duration) error
}

type containerReaper struct {
	logger            lager.Logger
	workerClient      worker.Client
	db                ContainerReaperDB
	pipelineDBFactory db.PipelineDBFactory
}

func NewContainerReaper(
	logger lager.Logger,
	workerClient worker.Client,
	db ContainerReaperDB,
	pipelineDBFactory db.PipelineDBFactory,
) ContainerReaper {
	return &containerReaper{
		logger:            logger,
		workerClient:      workerClient,
		db:                db,
		pipelineDBFactory: pipelineDBFactory,
	}
}

func (cr *containerReaper) updateWorkerContainerTTL(handle string) error {
	workerContainer, found, err := cr.workerClient.LookupContainer(cr.logger, handle)
	if err != nil {
		cr.logger.Error("error-finding-worker-container", err)
		return err
	}

	if !found {
		cr.logger.Error("worker-containerr-not-found", nil)
		return errors.New("worker-container-not-found")
	}

	workerContainer.Release(worker.FinalTTL(worker.ContainerTTL))
	return nil
}

func (cr *containerReaper) Run() error {
	successfulContainers, err := cr.db.FindContainersFromSuccessfulBuildsWithInfiniteTTL()
	if err != nil {
		cr.logger.Error("failed-to-find-successful-containers", err)
	}

	for _, container := range successfulContainers {
		err := cr.updateWorkerContainerTTL(container.Handle)
		if err != nil {
			continue
		}
		err = cr.db.UpdateExpiresAtOnContainer(container.Handle, worker.ContainerTTL)
		if err != nil {
			cr.logger.Error("error-updating-db-container-ttl", err)
		}
	}

	failedContainers, err := cr.db.FindContainersFromUnsuccessfulBuildsWithInfiniteTTL()
	if err != nil {
		return err
	}

	var jobContainerMap map[int][]db.SavedContainer
	jobContainerMap = make(map[int][]db.SavedContainer)

	for _, container := range failedContainers {
		buildID := container.BuildID

		// TODO: if the container's job no longer exists in the configuration, expire it
		jobID, found, err := cr.db.FindJobIDForBuild(buildID)
		if err != nil {
			cr.logger.Error("find-job-id-for-build", err, lager.Data{"build-id": buildID})
		}
		if !found {
			cr.logger.Error("unable-to-find-job-id-for-build", nil, lager.Data{"build-id": buildID})
		}

		jobContainers := jobContainerMap[jobID]
		if jobContainers == nil {
			jobContainerMap[jobID] = []db.SavedContainer{container}
		} else {
			jobContainers = append(jobContainers, container)
			jobContainerMap[jobID] = jobContainers
		}
	}

	for _, jobContainers := range jobContainerMap {
		maxBuildID := -1
		for _, jobContainer := range jobContainers {
			if jobContainer.BuildID > maxBuildID {
				maxBuildID = jobContainer.BuildID
			}
		}

		for _, jobContainer := range jobContainers {
			if jobContainer.BuildID < maxBuildID {
				handle := jobContainer.Container.Handle
				err := cr.updateWorkerContainerTTL(handle)
				if err != nil {
					continue
				}
				err = cr.db.UpdateExpiresAtOnContainer(handle, worker.ContainerTTL)
				if err != nil {
					cr.logger.Error("error-updating-db-container-ttl", err)
				}
			}
		}
	}

	return nil
}