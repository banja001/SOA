package model

import (
	"encoding/json"
	"io"
)

type UserRole int

const (
	Administrator UserRole = iota
	Author
	Tourist
)

type User struct {
	ID                     int      `json:"Id" gorm:"primary_key"`
	PersonId               int      `json:"PersonId" orm:"not null;`
	Username               string   `json: "Username" gorm:"not null;`
	Password               string   `json: "Password" gorm:"not null;`
	Role                   UserRole `json: "Role" gorm:"not null;`
	IsActive               bool     `json: "IsActive" gorm:"not null;`
	ResetPasswordToken     string   `json: "ResetPasswordToken" gorm:"not null;`
	EmailVerificationToken string   `json: "EmailVerificationToken" gorm:"not null;`
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

func (ur UserRole) String() string {
	switch ur {
	case Administrator:
		return "administrator"
	case Author:
		return "author"
	case Tourist:
		return "tourist"
	default:
		return "unknown"
	}
}
