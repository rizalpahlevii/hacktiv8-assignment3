package model

import "time"

type Weather struct {
	ID        uint      `gorm:"primary_key" json:"-"`
	Water     int       `json:"water"`
	Wind      int       `json:"wind"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (weather Weather) GetWindStatus() string {
	windStatus := "Aman"
	if weather.Wind < 6 {
		windStatus = "Aman"
	} else if weather.Wind >= 7 && weather.Wind <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}
	return windStatus
}

func (weather Weather) GetWaterStatus() string {
	waterStatus := "Aman"
	if weather.Water < 5 {
		waterStatus = "Aman"
	} else if weather.Water >= 6 && weather.Water <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}
	return waterStatus
}
