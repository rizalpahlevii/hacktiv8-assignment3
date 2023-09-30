package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-assignment3/dto"
	"hacktiv8-assignment3/repository"
	"hacktiv8-assignment3/resource"
)

type WeatherControllerInterface interface {
	UpdateWeather(context *gin.Context)
	GetWeather(ctx *gin.Context)
}

type WeatherController struct {
	weatherRepository repository.WeatherRepository
}

func NewWeatherController(weatherRepository repository.WeatherRepository) WeatherController {
	return WeatherController{weatherRepository: weatherRepository}
}

func (controller WeatherController) UpdateWeather(context *gin.Context) {
	var weatherRequest dto.WeatherRequest
	if err := context.ShouldBindJSON(&weatherRequest); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	weather := controller.weatherRepository.UpdateWeather(weatherRequest)

	response := resource.WeatherResource{
		Wind:        weather.Wind,
		Water:       weather.Water,
		WindStatus:  weather.GetWindStatus(),
		WaterStatus: weather.GetWaterStatus(),
	}

	context.JSON(200, gin.H{
		"message": "Weather updated successfully",
		"data":    response,
	})
}

func (controller WeatherController) GetWeather(context *gin.Context) {
	weather := controller.weatherRepository.GetWeather()
	context.JSON(200, weather)
}
