package model

import (
	"encoding/json"
	"io"
	"time"
)

type Tour struct {
	ID            int				 `json:"Id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"Name,omitempty" bson:"name,omitempty"`
	Description   string             `json:"Description,omitempty" bson:"description,omitempty"`
	Difficulty    TourDifficulty     `json:"Difficulty,omitempty" bson:"difficulty,omitempty"`
	Tags          []string           `json:"Tags,omitempty" bson:"tags,omitempty"`
	Status        TourStatus         `json:"Status,omitempty" bson:"status,omitempty"`
	Price         float64            `json:"Price,omitempty" bson:"price,omitempty"`
	AuthorId      int                `json:"AuthorId,omitempty" bson:"authorId,omitempty"`
	Equipment     []string           `json:"Equipment,omitempty" bson:"equipment,omitempty"`
	DistanceInKm  float64            `json:"DistanceInKm,omitempty" bson:"distance,omitempty"`
	ArchivedDate  *time.Time         `json:"ArchivedDate,omitempty" bson:"archived,omitempty"`
	PublishedDate *time.Time         `json:"PublishedDate,omitempty" bson:"published,omitempty"`
	Durations     TourDurations      `json:"Durations,omitempty" bson:"durations,omitempty"`
	KeyPoints     []TourKeypoint     `json:"KeyPoints,omitempty" bson:"keypoints,omitempty"`
	Image         string             `json:"Image,omitempty" bson:"image,omitempty"`
}

type TourDifficulty int

const (
	Beginner TourDifficulty = iota
	Intermediate
	Advanced
	Pro
)

type TourStatus int

const (
	Draft TourStatus = iota
	Published
	Archived
	TouristMade
)

func (p *Tour) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tour) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
