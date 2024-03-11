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
	ID            uuid.UUID      `json:"id" gorm:"primary_key"`
	Name          string         `json:"name" gorm:"not null;type:string"`
	Description   string         `json:"description" gorm:"not null;type:string"`
	//Difficulty    TourDifficulty `json:"difficulty" gorm:"type:enum('Beginner', 'Intermediate', 'Advanced', 'Pro')"`
	Tags          []string       `json:"tags" gorm:"not null;type:string"`
	//Status        TourStatus     `json:"status" gorm:"type:enum('Draft', 'Published', 'Archived', 'TouristMade')"`
	Price         float64        `json:"price"`
	AuthorId      int            `json:"authorId"`
	Equipment     []int          `json:"equipment" gorm:"type:string"`
	DistanceInKm  float64        `json:"distanceInKm"`
	ArchivedDate  *time.Time     `json:"archivedDate"`
	PublishedDate *time.Time     `json:"publishedDate"`
	//Durations     []TourDuration `gorm:"foreignKey:TourID"`
	//KeyPoints     []TourKeyPoint `gorm:"foreignKey:TourID"`
	Image string `json:"image" gorm:"not null;type:string"`
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	tour.ID = uuid.New()
	return nil
}
