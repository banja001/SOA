package model

import (
	"encoding/json"
	"io"
)

type TourKeypoint struct {
	ID             int     `json:"Id" bson:"_id,omitempty"`
	Name           string  `json:"Name" bson:"Name,omitempty"`
	Description    string  `json:"Description" bson:"Description,omitempty"`
	Image          string  `json:"Image" bson:"Image,omitempty"`
	Latitude       float64 `json:"Latitude" bson:"Latitude,omitempty"`
	Longitude      float64 `json:"Longitude" bson:"Longitude,omitempty"`
	TourID         int     `json:"TourId" bson:"TourId,omitempty"`
	Secret         string  `json:"Secret" bson:"Secret,omitempty"`
	PositionInTour int     `json:"PositionInTour" bson:"PositionInTour,omitempty"`
	PublicPointID  int     `json:"PublicPointId" bson:"PublicPointId,omitempty"`
}

func (tkp *TourKeypoint) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(tkp)
}

func (tkp *TourKeypoint) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(tkp)
}
