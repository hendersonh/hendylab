package weatherstation 

type WeatherService interface {
	GetWeather(city string) (WeatherData, error)
}

type WeatherData struct {
	Temperature float64
	Humidity    int
	Description string
}
