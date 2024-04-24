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

func (fs *FollowerService) GetAllPersonsNodes() (*model.Persons, error) {
	return fs.repo.GetAllPersonsNodes()
}

func (fs *FollowerService) GetAllFollowed(id int, uid int) (*model.Persons, error) {
	return fs.repo.GetAllFollowed(id, uid)
}

func (fs *FollowerService) RewriteFollower(updatedFollower *model.Follower) error {
	err := fs.repo.RewriteFollower(updatedFollower)
	if err != nil {
		fs.logger.Println("Error updating follower in service:", err)
		return err
	}
	return nil
}
