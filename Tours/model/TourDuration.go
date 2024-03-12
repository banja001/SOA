package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type TourDuration struct {
	TimeInSeconds  uint               `json:"TimeInSeconds"`
	Transportation TransportationType `json:"TransportationType"`
}

type TransportationType int

const (
	Walking TransportationType = iota
	Bicycle
	Car
)

type TourDurations []TourDuration

func (durations *TourDurations) Scan(value interface{}) error {
	if value == nil {
		*durations = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}

	if err := json.Unmarshal(bytes, durations); err != nil {
		return err
	}

	return nil
}

func (durations TourDurations) Value() (driver.Value, error) {
	if durations == nil {
		return nil, nil
	}

	bytes, err := json.Marshal(durations)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
