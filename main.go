package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hacktiv8-assignment3/controller"
	"hacktiv8-assignment3/dto"
	"hacktiv8-assignment3/model"
	"hacktiv8-assignment3/repository"
	"math/rand"
	"net/http"
	"time"
)

const PORT = "8080"

type ApiWeatherResponse struct {
	Message string `json:"message"`
	Data    struct {
		Wind        int    `json:"wind"`
		Water       int    `json:"water"`
		WindStatus  string `json:"wind_status"`
		WaterStatus string `json:"water_status"`
	}
}

func UpdateWeatherPer15Seconds() {

	apiUrl := "http://localhost:" + PORT + "/weather"
	for {

		weatherRequest := dto.WeatherRequest{
			Wind:  generateRandomData(),
			Water: generateRandomData(),
		}
		data, _ := json.Marshal(weatherRequest)
		response := ApiCall(apiUrl, http.MethodPut, data)

		var apiWeatherResponse ApiWeatherResponse
		err := json.Unmarshal([]byte(response), &apiWeatherResponse)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("")
		fmt.Println("Wind:", apiWeatherResponse.Data.Wind)
		fmt.Println("Water:", apiWeatherResponse.Data.Water)
		fmt.Println("Water Status:", apiWeatherResponse.Data.WaterStatus)
		fmt.Println("Wind Status:", apiWeatherResponse.Data.WindStatus)
		fmt.Println("")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	initDB()

	router := gin.Default()

	go UpdateWeatherPer15Seconds()

	weatherRepository := repository.NewWeatherRepository(GetDatabaseConnection())
	weatherController := controller.NewWeatherController(weatherRepository)

	router.PUT("/weather", weatherController.UpdateWeather)
	router.GET("/weather", weatherController.GetWeather)
	router.Run(":" + PORT)

}

func generateRandomData() int {
	return rand.Intn(100) + 1
}

func initDB() {
	db := GetDatabaseConnection()
	err := db.AutoMigrate(&model.Weather{})
	if err != nil {
		panic(err)
	}
}
