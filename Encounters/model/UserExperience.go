package model

import (
	"gorm.io/gorm"
)

type UserExperience struct {
	gorm.Model
	ID            int            `json:"id" gorm:"primary_key"`
	UserID        int            `json:"userId"`
	XP            int            `json:"xp"`
	Level         int            `json:"level"`
}