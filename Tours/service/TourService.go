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

func (service *TourService) GetByAuthorId(authorID string) ([]*model.Tour, error) {
	tours, err := service.TourRepo.FindByAuthorId(authorID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", authorID))
	}
	var tourPointers []*model.Tour
	for i := range tours {
		tourPointers = append(tourPointers, &tours[i])
	}

	return tourPointers, nil
}

// func (service *TourService) ChangeStatus(id string, authorID int, tourStatus model.TourStatus) (*model.Tour, error) {
// 	tour, err := service.TourRepo.FindById(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("tour with ID %s not found", id)
// 	}

// 	if tour.AuthorId != authorID {
// 		return nil, fmt.Errorf("not authorized to change the status of this tour")
// 	}
// 	currentTime := time.Now()
// 	if tourStatus == model.Published {
// 		tour.PublishedDate = &currentTime
// 	}
// 	if tourStatus == model.Archived {
// 		if tour.Status != model.Published {
// 			return nil, fmt.Errorf("tour must be published in order to be archived")
// 		}
// 		tour.ArchivedDate = &currentTime
// 	}
// 	tour.Status = tourStatus
// 	updatedTour, err := service.TourRepo.Update(&tour)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to update tour: %v", err)
// 	}
// 	return updatedTour, nil
// }
