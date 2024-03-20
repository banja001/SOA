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

func (repo *UserExperienceRepository) Create(userExperience *model.UserExperience) (*model.UserExperience, error) {
	dbResult := repo.DatabaseConnection.Create(userExperience)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("User experience created: ", dbResult.RowsAffected)
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

func (repo *UserExperienceRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Unscoped().Delete(&model.UserExperience{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("User experience deleted: ", dbResult.RowsAffected)
	return nil
}