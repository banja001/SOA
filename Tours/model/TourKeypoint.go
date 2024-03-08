package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourKeypoint struct {
	ID   uuid.UUID `json:"Id"`
	Name string    `json:"Name" gorm:"not null;type:string"`
	Description string `json:"Description"`
	Image string `json:"Image"`
	Latitude float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	TourID uuid.UUID `json:"TourId"`
	Secret string `json:"Secret"`
	PositionInTour int32 `json:"PositionInTour"`
	PublicPointID uuid.UUID `json:"PublicPointId"`
}

func (tourKeypoint *TourKeypoint) BeforeCreate(scope *gorm.DB) error {
	tourKeypoint.ID = uuid.New()
	return nil
}