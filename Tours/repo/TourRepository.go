package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) FindById(id string) (model.Tour, error) {
	var tour model.Tour
	dbResult := repo.DatabaseConnection.First(&tour, "id = ?", id)
	if dbResult.Error != nil {
		return tour, dbResult.Error
	}
	var keypoints []model.TourKeypoint
	dbResult = repo.DatabaseConnection.Where("tour_id = ?", tour.ID).Find(&keypoints)
	if dbResult.Error != nil {
		return tour, dbResult.Error
	}
	tour.KeyPoints = keypoints

	return tour, nil
}

func (repo *TourRepository) Create(tour *model.Tour) (*model.Tour, error) {
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return tour, nil
}

func (repo *TourRepository) FindAll() ([]model.Tour, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Find(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	for i := range tours {
		var keypoints []model.TourKeypoint
		dbResult := repo.DatabaseConnection.Where("tour_id = ?", tours[i].ID).Find(&keypoints)
		if dbResult.Error != nil {
			return nil, dbResult.Error
		}
		tours[i].KeyPoints = keypoints
	}
	return tours, nil
}

func (repo *TourRepository) Update(tour *model.Tour) (*model.Tour, error) {
	dbResult := repo.DatabaseConnection.Updates(tour)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return tour, nil
}
