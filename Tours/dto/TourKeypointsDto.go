package dto

import (
	"github.com/google/uuid"
)

type TourKeypoint struct {
	ID   uuid.UUID `json:"Id"`
	Name string    `json:"Name"`
	Description string `json:"Description"`
	Image string `json:"Image"`
	Latitude float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	TourID uuid.UUID `json:"TourId"`
	Secret string `json:"Secret"`
	PositionInTour int32 `json:"PositionInTour"`
	PublicPointID uuid.UUID `json:"PublicPointId"`
}
