package weatherstation 

import (
    "encoding/json"
    "net/http"
)

type OpenWeatherService struct{}

func (ows *OpenWeatherService) GetWeather(city string) (WeatherData, error) {
    apiKey := "e7d615b7307b9865a8c6b386cfb77a9e"
    url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey
    resp, err := http.Get(url)
    if err != nil {
        return WeatherData{}, err
    }
    defer resp.Body.Close()

    var data struct {
        Main struct {
            Temp     float64 `json:"temp"`
            Humidity int     `json:"humidity"`
        }
        Weather []struct {
            Description string `json:"description"`
        }
    }
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return WeatherData{}, err
    }

    return WeatherData{
        Temperature: data.Main.Temp,
        Humidity:    data.Main.Humidity,
        Description: data.Weather[0].Description,
    }, nil
}
