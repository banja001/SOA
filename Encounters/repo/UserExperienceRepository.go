package repo

import (
	"encgo/model"

	"gorm.io/gorm"
)

type UserExperienceRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *UserExperienceRepository) FindByUserId(userId int) (model.UserExperience, error) {
	userExperience := model.UserExperience{}
	dbResult := repo.DatabaseConnection.First(&userExperience, "user_id = ?", userId)
	if dbResult != nil {
		return userExperience, dbResult.Error
	}
	return userExperience, nil
}