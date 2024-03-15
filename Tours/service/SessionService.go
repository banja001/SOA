package service

import (
	"database-example/model"
	"database-example/repo"
	"time"
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

func (service *SessionService) Update(session *model.Session) (*model.Session, error) {
	updatedSession, err := service.SessionRepo.Update(session)
	if err != nil {
		return nil, err
	}
	return updatedSession, nil
}

func (service *SessionService) CompleteKeypoint(sessionId string, keypointId int) (*model.Session, error){
	session, err := service.SessionRepo.FindById(sessionId)
	if err != nil {
		return nil, err
	}

	session = *service.AddKeypoint(&session, keypointId)
	updatedSession, err := service.Update(&session)
	if err != nil {
		return nil, err
	}

	return updatedSession, nil
}

func (service *SessionService) AddKeypoint(session *model.Session, keypointId int) (*model.Session){
	currentTime := time.Now()
	completedKeypoint := model.CompletedKeyPoint{
		KeypointId: keypointId,
		CompletionTime: &currentTime,
	}

	completeKeypointCheck := false
	for _, keypoint := range session.CompletedKeyPoints{
		if keypoint.KeypointId == keypointId{
			completeKeypointCheck = true
			break
		}
	}

	if !completeKeypointCheck {
		session.CompletedKeyPoints = append(session.CompletedKeyPoints, completedKeypoint)
	}
	
	return session
}
