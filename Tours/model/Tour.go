package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ArrayValue []interface{}

func (a *ArrayValue) Scan(value interface{}) error {
	if value == nil {
		*a = []interface{}{}
		return nil
	}

	switch v := value.(type) {
	case string:
		str := strings.TrimPrefix(v, "{")
		str = strings.TrimSuffix(str, "}")
		strValues := strings.Split(str, ",")
		var valueSlice []interface{}
		for _, strValue := range strValues {
			strValue = strings.TrimSpace(strValue)
			if strValue == "" {
				continue
			}
			intValue, err := strconv.Atoi(strValue)
			if err == nil {
				valueSlice = append(valueSlice, intValue)
			} else {
				valueSlice = append(valueSlice, strValue)
			}
		}
		*a = valueSlice
	case []interface{}:
		*a = v
	default:
		return errors.New("failed to scan ArrayValue: unexpected type")
	}

	return nil
}

func (a ArrayValue) Value() (driver.Value, error) {
	var strArray []string
	for _, v := range a {
		strArray = append(strArray, fmt.Sprintf("%v", v))
	}
	return "{" + strings.Join(strArray, ",") + "}", nil
}

type Tour struct {
	ID            int            `json:"Id" gorm:"primary_key"`
	Name          string         `json:"Name" gorm:"not null;type:text"`
	Description   string         `json:"Description" gorm:"not null;type:text"`
	Difficulty    TourDifficulty `json:"Difficulty" gorm:"type:integer"`
	Tags          ArrayValue     `json:"Tags" gorm:"not null;type:text[]"`
	Status        TourStatus     `json:"Status" gorm:"type:integer"`
	Price         float64        `json:"Price"`
	AuthorId      int            `json:"AuthorId" gorm:"type:integer"`
	Equipment     ArrayValue     `json:"Equipment" gorm:"type:integer[]"`
	DistanceInKm  float64        `json:"DistanceInKm"`
	ArchivedDate  *time.Time     `json:"ArchivedDate"`
	PublishedDate *time.Time     `json:"PublishedDate"`
	Durations     TourDurations  `json:"Durations" gorm:"type:jsonb"`
	KeyPoints     []TourKeypoint `json:"KeyPoints" gorm:"-"`
	Image         string         `json:"Image" gorm:"type:text"`
}

type TourStatus int

const (
	Draft TourStatus = iota
	Published
	Archived
	TouristMade
)

type TourDifficulty int

const (
	Beginner TourDifficulty = iota
	Intermediate
	Advanced
	Pro
)

func (Tour) TableName() string {
	return "Tour"
}
