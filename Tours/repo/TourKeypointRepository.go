package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourKeypointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourKeypointRepository) FindById(id string) (model.TourKeypoint, error) {
	tourKeypoint := model.TourKeypoint{}
	dbResult := repo.DatabaseConnection.First(&tourKeypoint, "id = ?", id)
	if dbResult != nil {
		return tourKeypoint, dbResult.Error
	}
	return tourKeypoint, nil
}

func (repo *TourKeypointRepository) Create(tourKeypoint *model.TourKeypoint) error {
	dbResult := repo.DatabaseConnection.Create(tourKeypoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourKeypointRepository) Update(tourKeypoint *model.TourKeypoint) error {
	dbResult := repo.DatabaseConnection.Updates(tourKeypoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourKeypointRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Unscoped().Delete(&model.TourKeypoint{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
