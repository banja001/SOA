package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TourService struct {
	TourRepo *repo.TourRepository
}

func (service *TourService) Find(id string) (*model.Tour, error) {
	tour, err := service.TourRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tour, nil
}

func (service *TourService) Create(tour *model.Tour) (*model.Tour, error) {
	createdTour, err := service.TourRepo.Create(tour)
	if err != nil {
		return nil, err
	}
	return createdTour, nil
}

func (service *TourService) GetAll() ([]*model.Tour, error) {
	tours, err := service.TourRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var tourPointers []*model.Tour
	for i := range tours {
		tourPointers = append(tourPointers, &tours[i])
	}

	return tourPointers, nil
}

func (service *TourService) Update(tour *model.Tour) (*model.Tour, error) {
	updatedTour, err := service.TourRepo.Update(tour)
	if err != nil {
		return nil, err
	}
	return updatedTour, nil
}
