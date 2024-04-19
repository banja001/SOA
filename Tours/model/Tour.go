package model

import (
	"encoding/json"
	"io"
	"time"
)

type Tour struct {
	ID            int				 `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	Difficulty    TourDifficulty     `json:"difficulty,omitempty" bson:"difficulty,omitempty"`
	Tags          []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Status        TourStatus         `json:"status,omitempty" bson:"status,omitempty"`
	Price         float64            `json:"price,omitempty" bson:"price,omitempty"`
	AuthorId      int                `json:"authorId,omitempty" bson:"authorId,omitempty"`
	Equipment     []string           `json:"equipment,omitempty" bson:"equipment,omitempty"`
	DistanceInKm  float64            `json:"distanceInKm,omitempty" bson:"distance,omitempty"`
	ArchivedDate  *time.Time         `json:"archivedDate,omitempty" bson:"archived,omitempty"`
	PublishedDate *time.Time         `json:"publishedDate,omitempty" bson:"published,omitempty"`
	Durations     TourDurations      `json:"durations,omitempty" bson:"durations,omitempty"`
	KeyPoints     []TourKeypoint     `json:"keyPoints,omitempty" bson:"keypoints,omitempty"`
	Image         string             `json:"image,omitempty" bson:"image,omitempty"`
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
