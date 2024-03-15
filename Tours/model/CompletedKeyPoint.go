package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type CompletedKeyPoint struct {
	KeypointId     int        `json:"KeyPointId"`
	CompletionTime *time.Time `json:"CompletionTime"`
}

type CompletedKeyPoints []CompletedKeyPoint

func (keypoints *CompletedKeyPoints) Scan(value interface{}) error {
	if value == nil {
		*keypoints = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}


	if err := json.Unmarshal(bytes, keypoints); err != nil {
		return err
	}

	return nil
}

func (keypoints CompletedKeyPoints) Value() (driver.Value, error) {
	if keypoints == nil {
		return nil, nil
	}

	bytes, err := json.Marshal(keypoints)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
