package model

import (
	"encoding/json"
	"io"
)

type Tour struct {
	ID            int				 `json:"id,omitempty"`
	Name          string             `json:"name,omitempty"`
	Description   string             `json:"description,omitempty"`
	//Difficulty    TourDifficulty     `json:"difficulty,omitempty"`
	Tags          []string           `json:"tags,omitempty"`
	//Status        TourStatus         `json:"status,omitempty"`
	Price         float64            `json:"price,omitempty"`
	AuthorId      int                `json:"authorId,omitempty"`
	Equipment     []string           `json:"equipment,omitempty"`
	DistanceInKm  float64            `json:"distanceInKm,omitempty"`
	//ArchivedDate  *time.Time         `json:"archivedDate,omitempty"`
	//PublishedDate *time.Time         `json:"publishedDate,omitempty"`
	Durations     TourDurations      `json:"durations,omitempty"`
	KeyPoints     []TourKeypoint     `json:"keyPoints,omitempty"`
	Image         string             `json:"image,omitempty"`
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
