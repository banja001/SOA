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

func (service *TourKeypointService) Create(tourKeypoint *model.TourKeypoint) (*model.TourKeypoint, error) {
	createdTourKeypoint, err := service.TourKeypointRepo.Create(tourKeypoint)
	if err != nil {
		return nil, err
	}
	return &createdTourKeypoint, nil
}

func (service *TourKeypointService) Update(tourKeypoint *model.TourKeypoint) (*model.TourKeypoint, error) {
	updatedTourKeypoint, err := service.TourKeypointRepo.Update(tourKeypoint)
	if err != nil {
		return nil, err
	}
	return &updatedTourKeypoint, nil
}

func (service *TourKeypointService) Delete(id string) error {
	err := service.TourKeypointRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

