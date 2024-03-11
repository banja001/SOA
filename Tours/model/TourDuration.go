package model

type TourDuration struct {
	//ID             int                `json:"Id"`
	TimeInSeconds  uint               `json:"TimeInSeconds"`
	Transportation TransportationType `json:"TransportationType"`
	//TourID         int                `json:"TourId"`
}

type TransportationType int

const (
	Walking TransportationType = iota
	Bicycle
	Car
)
