package models

import (
	"ticket-booking/utils" // Import the utils package where DB is initialized
	"github.com/jinzhu/gorm"
)

type Ticket struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	IdCardNo string `json:"id_card_no"`
	AgencyName  string `json:"agency_name"`
	SeatNo   string `json:"seat_no"`
	TravelDate string `json:"travel_date"`
}

func (ticket *Ticket) Save() error {
	return utils.DB.Create(&ticket).Error
}
