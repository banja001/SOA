package model

import (
	"encoding/json"
	"io"
)

type Person struct {
	ID      int64  `json:"Id" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"Name,omitempty"`
	Surname string `json:"Surname,omitempty"`
	Email   string `json:"Email,omitempty"`
}

type Persons []*Person

func (persons Persons) ToJSON(w io.Writer) error {
	// Write the opening square bracket '['
	_, err := w.Write([]byte{'['})
	if err != nil {
		return err
	}

	// Create a JSON encoder
	encoder := json.NewEncoder(w)
	// Iterate over each person in the slice
	for i, person := range persons {
		// Encode each person to JSON and write it to the writer
		if err := encoder.Encode(person); err != nil {
			return err
		}
		// If it's not the last person, write a comma separator
		if i < len(persons)-1 {
			_, err := w.Write([]byte{','})
			if err != nil {
				return err
			}
		}
	}

	// Write the closing square bracket ']'
	_, err = w.Write([]byte{']'})
	if err != nil {
		return err
	}

	return nil
}

func (person *Person) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(person)
}

func (person *Person) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(person)
}
