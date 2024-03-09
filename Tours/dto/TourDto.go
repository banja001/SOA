package dto

import (
	"time"

	"github.com/google/uuid"
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
	ID            uuid.UUID
	Name          string
	Description   string
	Difficulty    TourDifficulty
	Tags          []string
	Status        TourStatus
	Price         float64
	AuthorId      int
	Equipment     []int
	DistanceInKm  float64
	ArchivedDate  *time.Time
	PublishedDate *time.Time
	//Durations     []TourDurationDto
	//KeyPoints     []TourKeyPointDto
	Image string
}
