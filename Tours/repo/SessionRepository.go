package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type SessionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *SessionRepository) Create(session *model.Session) (*model.Session, error) {
	dbResult := repo.DatabaseConnection.Create(session)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return session, nil
}
