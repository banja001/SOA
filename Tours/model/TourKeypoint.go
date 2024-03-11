package model

type TourKeypoint struct {
	ID             int     `json:"Id"`
	Name           string  `json:"Name" gorm:"not null;type:string"`
	Description    string  `json:"Description"`
	Image          string  `json:"Image"`
	Latitude       float64 `json:"Latitude"`
	Longitude      float64 `json:"Longitude"`
	TourID         int     `json:"TourId"`
	Secret         string  `json:"Secret"`
	PositionInTour int     `json:"PositionInTour"`
	PublicPointID  int     `json:"PublicPointId"`
}

// func (tourKeypoint *TourKeypoint) BeforeCreate(scope *gorm.DB) error {
// 	tourKeypoint.ID = uuid.New()
// 	return nil
// }
