package service

import "encgo/repo"

type ChallengeExecutionService struct {
	ChallengeExecutionRepository *repo.ChallengeExecutionRepository
}

func (service *ChallengeExecutionService) Delete(id int) error {
	err := service.ChallengeExecutionRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
