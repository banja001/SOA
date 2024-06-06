package service

import (
	"encgo/model"
	"encgo/repo"
	"fmt"
)

type UserExperienceService struct {
	UserExperienceRepo *repo.UserExperienceRepository
	model              model.UserExperience
}

func NewProductService(store model.UserExperience) *UserExperienceService {
	return &UserExperienceService{
		model: store,
	}
}

func (service *UserExperienceService) AddXP(id int, xp int) (*model.UserExperience, error) {
	userExperience, err := service.UserExperienceRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %d not found", id))
	}
	userExperience.AddXP(xp)
	userExperience.CalculateLevel()
	service.UserExperienceRepo.Update(&userExperience)

	return &userExperience, nil
}

func (service *UserExperienceService) FindByUserId(userId int) (*model.UserExperience, error) {
	userExperience, err := service.UserExperienceRepo.FindByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with user id %d not found", userId))
	}
	return &userExperience, nil
}

func (service *UserExperienceService) Create(userExperience *model.UserExperience) (*model.UserExperience, error) {
	createdUserExperience, err := service.UserExperienceRepo.Create(userExperience)
	if err != nil {
		return nil, err
	}
	return createdUserExperience, nil
}

func (service *UserExperienceService) Delete(id string) error {
	err := service.UserExperienceRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserExperienceService) Update(userExperience *model.UserExperience) (*model.UserExperience, error) {
	updatedUserExperience, err := service.UserExperienceRepo.Update(userExperience)
	if err != nil {
		return nil, err
	}
	return updatedUserExperience, nil
}
