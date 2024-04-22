package service

import (
	"log"

	"Followers/model"
	"Followers/repo"
)

type FollowerService struct {
    repo   *repo.FollowerRepo
    logger *log.Logger
}

func NewFollowerService(repo *repo.FollowerRepo, logger *log.Logger) *FollowerService {
    return &FollowerService{
        repo:   repo,
        logger: logger,
    }
}

func (fs *FollowerService) GetAllFollowers() (model.Followers, error) {
    return fs.repo.GetAllFollowerNodes()
}
