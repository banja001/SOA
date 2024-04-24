package model

import (
	"encoding/json"
	"io"
	"time"
)

type Follower struct {
	ID           int64                `json:"Id" gorm:"primaryKey;autoIncrement"`
	FollowerId   int64                `json:"FollowerId,omitempty"`
	FollowedId   int64                `json:"FollowedId,omitempty"`
	Notification FollowerNotification `json:"Notification,omitempty"`
}

type FollowerNotification struct {
	Content       string     `json:"Content,omitempty"`
	TimeOfArrival *time.Time `json:"TimeOfArrival,omitempty"`
	Read          bool       `json:"Read,omitempty"`
}

type Followers []*Follower
type FollowerNotifications []*FollowerNotification

func (o *Followers) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *FollowerNotifications) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Follower) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *FollowerNotification) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}
