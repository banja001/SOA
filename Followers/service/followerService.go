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

func (fs *FollowerService) RewriteFollower(updatedFollower *model.Follower) error {
	err := fs.repo.RewriteFollower(updatedFollower)
    if err != nil {
        fs.logger.Println("Error updating follower in service:", err)
        return err
    }
    return nil
}

