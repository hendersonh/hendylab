package main

import (
	"fmt"
	"github.com/hendersonh/hendylab/pkg/weatherstation"
	"log"
)

func main() {
	weatherService := weatherstation.OpenWeatherService{}
	city := "Paris"
	weatherData, err := weatherService.GetWeather(city)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Weather in %v: %v°C, Humidity: %v%%, Description: %s\n",
		city,
		weatherData.Temperature-273.15, // Convert Kelvin to Celsius
		weatherData.Humidity,
		weatherData.Description)
}
