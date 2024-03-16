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

func (userExperience *UserExperience) AddXP(xp int) {
	userExperience.XP += xp
}

func (userExperience *UserExperience) CalculateLevel() (int) {
	userExperience.Level += userExperience.XP / 20 + 1
	return userExperience.Level
}