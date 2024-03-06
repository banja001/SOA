package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourKeypoint struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name" gorm:"not null;type:string"`
	Description string 
	Image string
	Latitude float64
	Longitude float64
	TourID uuid.UUID
	Secret string
	PositionInTour int
	PublicPointID uuid.UUID
}

func (tourKeypoint *TourKeypoint) BeforeCreate(scope *gorm.DB) error {
	tourKeypoint.ID = uuid.New()
	return nil
}