package model

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"
	"time"
)

type TagsArray []string

func (tags *TagsArray) Scan(value interface{}) error {
	if value == nil {
		*tags = []string{}
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan TagsArray: unexpected type")
	}

	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")

	tagsArray := strings.Split(str, ",")

	for i := range tagsArray {
		tagsArray[i] = strings.TrimSpace(tagsArray[i])
	}

	*tags = tagsArray
	return nil
}

type IntArray []int

func (ints *IntArray) Scan(value interface{}) error {
	if value == nil {
		*ints = []int{}
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan IntArray: unexpected type")
	}

	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")

	strInts := strings.Split(str, ",")
	var intSlice []int
	for _, strInt := range strInts {
		i, err := strconv.Atoi(strInt)
		if err != nil {
			return err
		}
		intSlice = append(intSlice, i)
	}

	*ints = intSlice
	return nil
}

func (r IntArray) Value() (driver.Value, error) {
	var strArray []string
	for _, v := range r {
		strArray = append(strArray, strconv.Itoa(v))
	}
	return "{" + strings.Join(strArray, ",") + "}", nil
}

func (r TagsArray) Value() (driver.Value, error) {
	return "{" + strings.Join(r, ",") + "}", nil
}

type Tour struct {
	ID            int            `json:"Id" gorm:"primary_key"`
	Name          string         `json:"Name" gorm:"not null;type:text"`
	Description   string         `json:"Description" gorm:"not null;type:text"`
	Difficulty    TourDifficulty `json:"Difficulty" gorm:"type:integer"`
	Tags          TagsArray      `json:"Tags" gorm:"not null;type:text[]"`
	Status        TourStatus     `json:"Status" gorm:"type:integer"`
	Price         float64        `json:"Price"`
	AuthorId      int            `json:"AuthorId"`
	Equipment     IntArray       `json:"Equipment" gorm:"type:integer[]"`
	DistanceInKm  float64        `json:"DistanceInKm"`
	ArchivedDate  *time.Time     `json:"ArchivedDate"`  //?
	PublishedDate *time.Time     `json:"PublishedDate"` //?
	Durations     []TourDuration `json:"Durations" gorm:"-"`
	KeyPoints     []TourKeypoint `json:"KeyPoints" gorm:"-"`
	Image         string         `json:"Image" gorm:"type:text"` //?
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
