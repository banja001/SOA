package service

import (
	"encgo/model"
	"encgo/repo"
	"fmt"
)

type ChallengeService struct {
	ChallengeRepository *repo.ChallengeRepository
}

func (service *ChallengeService) GetAll() ([]model.Challenge, error) {
	challenges, err := service.ChallengeRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("Items not found")
	}
	return challenges, nil
}

func (service *ChallengeService) Delete(id int) error {
	err := service.ChallengeRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
