package service

import (
	"database-example/model"
	"database-example/repo"
)

type SessionService struct {
	SessionRepo *repo.SessionRepository
}

func (service *SessionService) Create(session *model.Session) (*model.Session, error) {
	createdSession, err := service.SessionRepo.Create(session)
	if err != nil {
		return nil, err
	}
	return createdSession, nil
}
