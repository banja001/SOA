package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tour struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name"`
	Description   string             `bson:"description" json:"description"`
	Difficulty    TourDifficulty     `bson:"difficulty" json:"difficulty"`
	Tags          []string           `bson:"tags" json:"tags"`
	Status        TourStatus         `bson:"status" json:"status"`
	Price         float64            `bson:"price" json:"price"`
	AuthorId      int                `bson:"authorId" json:"authorId"`
	Equipment     []string           `bson:"equipment" json:"equipment"`
	DistanceInKm  float64            `bson:"distanceInKm" json:"distanceInKm"`
	ArchivedDate  *time.Time         `bson:"archivedDate,omitempty" json:"archivedDate"`
	PublishedDate *time.Time         `bson:"publishedDate,omitempty" json:"publishedDate"`
	Durations     TourDurations      `bson:"durations" json:"durations"`
	KeyPoints     []TourKeypoint     `bson:"keyPoints" json:"keyPoints"`
	Image         string             `bson:"image" json:"image"`
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
