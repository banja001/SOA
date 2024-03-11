package service

import (
	"encgo/model"
	"encgo/repo"
	"fmt"
)

type UserExperienceService struct {
	UserExperienceRepo *repo.UserExperienceRepository
}

func (service *UserExperienceService) FindByUserId(userId int) (*model.UserExperience, error) {
	userExperience, err := service.UserExperienceRepo.FindByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with user id %d not found", userId))
	}
	return &userExperience, nil
}