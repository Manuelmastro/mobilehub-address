package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID     string `json:"user_id" validate:"required"`
	Country    string `json:"country" validate:"required"`
	State      string `json:"state" validate:"required"`
	District   string `json:"district" validate:"required"`
	StreetName string `json:"street_name" validate:"required"`
	PinCode    string `json:"pin_code" validate:"required,numeric"`
	Phone      string `json:"phone" validate:"required,numeric,len=10"`
}
