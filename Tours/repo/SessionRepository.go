package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type SessionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *SessionRepository) FindById(id string) (model.Session, error) {
	session := model.Session{}
	dbResult := repo.DatabaseConnection.First(&session, "id = ?", id)
	if dbResult != nil {
		return session, dbResult.Error
	}
	return session, nil
}

func (repo *SessionRepository) Create(session *model.Session) (*model.Session, error) {
	dbResult := repo.DatabaseConnection.Create(session)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("Sessions created: ", dbResult.RowsAffected)
	return session, nil
}

func (repo *SessionRepository) Update(session *model.Session) (*model.Session, error) {
	dbResult := repo.DatabaseConnection.Updates(session)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("Sessions updated: ", dbResult.RowsAffected)
	return session, nil
}

