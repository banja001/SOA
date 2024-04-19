package model

import (
	"time"
)

type Session struct {
	ID                     int                `json:"Id" bson:"_id,omitempty"`
	TourId                 int                `json:"TourId" bson:"tourId,omitempty"`
	TouristId              int                `json:"TouristId" bson:"touristId,omitempty"`
	LocationId             int                `json:"LocationId" bson:"locationId,omitempty"`
	SessionStatus          SessionStatus      `json:"SessionStatus" bson:"status,omitempty"`
	Transportation         int                `json:"Transportation" bson:"transportation,omitempty"`
	DistanceCrossedPercent int                `json:"DistanceCrossedPercent" bson:"distanceCrossed,omitempty"`
	LastActivity           *time.Time         `json:"LastActivity" bson:"lastActivity,omitempty"`
	CompletedKeyPoints     CompletedKeyPoints `json:"CompletedKeyPoints" bson:"completedKeyPoints,omitempty"`
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
