package model

type ChallengeStatus int

const (
	Draft ChallengeStatus = iota
	Active
	Archived
)

type ChallengeType int

const (
	Social ChallengeType = iota
	Location
	Misc
)

type Challenge struct {
	ID                 int             `json:"id" gorm:"primary_key"`
	CreatorId          int             `json: "creatorId" gorm:"not null;`
	Description        string          `json: "description" gorm:"not null;`
	Name               string          `json: "name" gorm:"not null;`
	Status             ChallengeStatus `json: "status" gorm:"not null;`
	Type               ChallengeType   `json: "type" gorm:"not null;`
	Latitude           float64         `json: "latitude" gorm:"not null;`
	Longitude          float64         `json: "longitude" gorm:"not null;`
	ExperiencePoints   int             `json: "experiencePoints" gorm:"not null;`
	KeyPointId         int             `json: "keyPointId"`
	Image              string          `json: "image"`
	LatitudeImage      float64         `json: "latitudeImage"`
	LongitudeImage     float64         `json: "longitudeImage"`
	Range              float64         `json: "range" gorm:"not null;`
	RequiredAttendance int             `json: "requiredAttendance"`
}
