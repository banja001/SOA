package repo

import (
	"encgo/model"

	"gorm.io/gorm"
)

type ChallengeRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *ChallengeRepository) GetAll() ([]model.Challenge, error) {
	challenges := []model.Challenge{}
	dbResult := repo.DatabaseConnection.Find(&challenges)
	if dbResult != nil {
		return challenges, dbResult.Error
	}
	return challenges, nil
}
