package model

import "time"

type ChallengeExecution struct {
	ID             int        `json:"id" gorm:"primary_key"`
	TouristId      int        `json:"touristId"`
	Challenge      Challenge  `json:"-"`
	ChallengeId    int        `json:"challenge"`
	Latitude       float64    `json:"latitude"`
	Longitude      float64    `json:"longitude"`
	ActivationTime *time.Time `json:"activationTime"`
	CompletionTime *time.Time `json:"completionTime"`
	IsCompleted    bool       `json:"isCompleted"`
}
