package repository

import (
	"gorm.io/gorm"
	"hacktiv8-assignment3/dto"
	"hacktiv8-assignment3/model"
)

type WeatherRepositoryInterface interface {
	UpdateWeather(payload dto.WeatherRequest)
	GetWeather() model.Weather
}

type WeatherRepository struct {
	db *gorm.DB
}

func (repository *WeatherRepository) UpdateWeather(payload dto.WeatherRequest) model.Weather {
	db := repository.db

	// Get the first record, if no records found, create one . If record found, update it
	var weather model.Weather

	if db.First(&weather).RowsAffected == 0 {
		db.Create(&payload)
	} else {
		db.Model(&weather).Updates(&payload)
	}

	return weather
}

func (repository *WeatherRepository) GetWeather() model.Weather {
	db := repository.db

	var weather model.Weather

	db.First(&weather)

	return weather
}

func NewWeatherRepository(db *gorm.DB) WeatherRepository {
	return WeatherRepository{db: db}
}
