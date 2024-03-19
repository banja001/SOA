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

func (repo *ChallengeRepository) Delete(id int) error {
	dbResult := repo.DatabaseConnection.Unscoped().Delete(&model.Challenge{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Challenge deleted, with id: ", id)
	return nil
}

func (repo *ChallengeRepository) Update(challenge *model.Challenge) (model.Challenge, error) {
	dbResult := repo.DatabaseConnection.Updates(challenge)
	if dbResult.Error != nil {
		return *challenge, dbResult.Error
	}
	println("Challenge updated: ", challenge.ID)
	return *challenge, nil
}
