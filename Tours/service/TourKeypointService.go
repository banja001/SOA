package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TourKeypointService struct {
	TourKeypointRepo *repo.TourKeypointRepository
}

func (service *TourKeypointService) Find(id string) (*model.TourKeypoint, error) {
	tourKeypoint, err := service.TourKeypointRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tourKeypoint, nil
}

func (service *TourKeypointService) Create(tourKeypoint *model.TourKeypoint) error {
	err := service.TourKeypointRepo.CreateTourKeypoint(tourKeypoint)
	if err != nil {
		return err
	}
	return nil
}

