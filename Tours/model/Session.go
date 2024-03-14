package model

import (
	"time"
)

type Session struct {
	ID                     int                `json:"Id" gorm:"primary_key"`
	TourId                 int                `json:"TourId" gorm:"type:integer"`
	TouristId              int                `json:"TouristId" gorm:"type:integer"`
	LocationId             int                `json:"LocationId" gorm:"type:integer"`
	SessionStatus          SessionStatus      `json:"SessionStatus" gorm:"type:integer"`
	Transportation         int                `json:"Transportation" gorm:"type:integer"`
	DistanceCrossedPercent int                `json:"DistanceCrossedPercent" gorm:"type:integer"`
	LastActivity           *time.Time         `json:"LastActivity"`
	CompletedKeyPoints     CompletedKeyPoints `json:"CompletedKeyPoints" gorm:"type:jsonb"`
}

type SessionStatus int

const (
	Active SessionStatus = iota
	Completed
	Abandoned
)

func (Session) TableName() string {
	return "Session"
}
