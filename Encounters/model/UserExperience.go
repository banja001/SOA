package model

type UserExperience struct {
	ID     int `json:"Id" gorm:"primaryKey;autoIncrement"`
	UserID int `json:"UserId"`
	XP     int `json:"XP"`
	Level  int `json:"Level"`
}

func (userExperience *UserExperience) AddXP(xp int) {
	userExperience.XP += xp
}

func (userExperience *UserExperience) CalculateLevel() int {
	userExperience.Level += userExperience.XP/20 + 1
	return userExperience.Level
}