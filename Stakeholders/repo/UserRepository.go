package repo

import (
	"stakeholders/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *UserRepository) FindActiveByUsername(username string) (model.User, error) {
	user := model.User{}
	dbResult := repo.DatabaseConnection.First(&user, "username = ? AND is_active = ?", username, true)
	if dbResult.Error != nil {
		return model.User{}, dbResult.Error
	}
	return user, nil
}
