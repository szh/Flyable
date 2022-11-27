package weatherfetcher

import (
	"fmt"
	"time"
)

const dateFormat = "2006-01-02"

func FetchWeather(config Config, startDate time.Time, endDate time.Time) ([]Weather, error) {
}

func buildApiUrl(config Config, startDate time.Time, endDate time.Time) string {
	// return config.WeatherApiUrl + "?latitude=" + config.Latitude + "&longitude=" + config.Longitude

	baseUrl := config.WeatherApiUrl
	queryStr := fmt.Sprintf("?latitude=%f&longitude=%f", config.Latitude, config.Longitude)
	queryStr += "&hourly=temperature_2m,apparent_temperature,wind_speed_10m,winddirection_10m,windgusts_10m,visibility,precipitation,cloudcover"
	queryStr += "&temperature_unit=fahrenheit&windspeed_unit=mph&precipitation_unit=inch"
	queryStr += fmt.Sprintf("&timezone=%s&start_date=%s&end_date=%s", config.TimeZone, startDate.Format(dateFormat), endDate.Format(dateFormat))

	return baseUrl + queryStr
}
