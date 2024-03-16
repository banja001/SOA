package model

type TourKeypoint struct {
	ID             int     `json:"Id"`
	Name           string  `json:"Name" gorm:"not null;type:string"`
	Description    string  `json:"Description" gorm:"not null"`
	Image          string  `json:"Image"`
	Latitude       float64 `json:"Latitude"`
	Longitude      float64 `json:"Longitude"`
	TourID         int     `json:"TourId"`
	Secret         string  `json:"Secret"`
	PositionInTour int     `json:"PositionInTour"`
	PublicPointID  int     `json:"PublicPointId"`
}

func (TourKeypoint) TableName() string {
	return "TourKeypoints"
}


// func (tourKeypoint TourKeypoint) Validate(db *gorm.DB) {
// 	if tourKeypoint.Latitude > 90{
// 		db.AddError(errors.New("latitude needs to be between -90 and 90"))
// 	}
// }

// func (tourKeypoint *TourKeypoint) BeforeCreate(db *gorm.DB) (err error) {
//     tourKeypoint.Validate(db)

// 	if db.Error != nil {
//         return db.Error
//     }

//     return nil
// }

// func (tourKeypoint *TourKeypoint) BeforeUpdate(db *gorm.DB) (err error) {
//     tourKeypoint.Validate(db)

// 	if db.Error != nil {
//         return db.Error
//     }

//     return nil
// }
