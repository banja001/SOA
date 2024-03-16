package repo

import (
	"encgo/model"

	"gorm.io/gorm"
)

type UserExperienceRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *UserExperienceRepository) FindById(id int) (model.UserExperience, error) {
	userExperience := model.UserExperience{}
	dbResult := repo.DatabaseConnection.First(&userExperience, "id = ?", id)
	if dbResult != nil {
		return userExperience, dbResult.Error
	}
	return userExperience, nil
}

func (repo *UserExperienceRepository) Update(userExperience *model.UserExperience) error {
	dbResult := repo.DatabaseConnection.Updates(userExperience)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *UserExperienceRepository) FindByUserId(userId int) (model.UserExperience, error) {
	userExperience := model.UserExperience{}
	dbResult := repo.DatabaseConnection.First(&userExperience, "user_id = ?", userId)
	if dbResult != nil {
		return userExperience, dbResult.Error
	}
	return userExperience, nil
}