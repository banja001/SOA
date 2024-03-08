package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourStatus string

const (
	Draft       TourStatus = "Draft"
	Published   TourStatus = "Published"
	Archived    TourStatus = "Archived"
	TouristMade TourStatus = "TouristMade"
)

type TourDifficulty string

const (
	Beginner     TourDifficulty = "Beginner"
	Intermediate TourDifficulty = "Intermediate"
	Advanced     TourDifficulty = "Advanced"
	Pro          TourDifficulty = "Pro"
)

type Tour struct {
	ID          uuid.UUID      `json:"Id" gorm:"primary_key"`
	Name        string         `json:"Name" gorm:"not null;type:string"`
	Description string         `json:"Description" gorm:"not null;type:string"`
	Difficulty  TourDifficulty `json:"Difficulty" gorm:"type:string"`
	//Tags []string `json:"Tags" gorm:"not null;type:text"`
	Status   TourStatus `json:"Status" gorm:"type:string"`
	Price    float64    `json:"Price"`
	AuthorId int        `json:"AuthorId"`
	//Equipment     []int      `json:"Equipment" gorm:"type:integer[]"`
	DistanceInKm  float64    `json:"DistanceInKm"`
	ArchivedDate  *time.Time `json:"ArchivedDate"`  //?
	PublishedDate *time.Time `json:"PublishedDate"` //?
	//Durations     []TourDuration `gorm:"foreignKey:TourID"`
	//KeyPoints     []TourKeyPoint `gorm:"foreignKey:TourID"`
	Image string `json:"Image" gorm:"type:string"` //?
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	tour.ID = uuid.New()
	return nil
}
