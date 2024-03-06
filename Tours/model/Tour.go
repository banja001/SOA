package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tour struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name" gorm:"not null;type:string"`
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	tour.ID = uuid.New()
	return nil
}
