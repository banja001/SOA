package repo

import (
	"encgo/model"

	"gorm.io/gorm"
)

type ChallengeExecutionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *ChallengeExecutionRepository) Delete(id int) error {
	dbResult := repo.DatabaseConnection.Unscoped().Delete(&model.ChallengeExecution{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("ChallengeExecution deleted, with id: ", id)
	return nil
}
