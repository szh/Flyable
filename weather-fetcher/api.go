package weatherfetcher

import (
	"context"
	"math"
	"time"

	"github.com/hectormalot/omgo"
	"github.com/szh/Flyable/shared"
)

func FetchWeather(config Config, startDate time.Time, endDate time.Time) ([]shared.Weather, error) {
	forecast, err := getForecast(config, startDate, endDate)
	if err != nil {
		return nil, err
	}

	weathers, err := parseForecast(forecast)
	if err != nil {
		return nil, err
	}
	return weathers, nil
}

func getForecast(config Config, startDate time.Time, endDate time.Time) (*omgo.Forecast, error) {
	c, err := omgo.NewClient()

	if err != nil {
		return nil, err
	}

	loc, err := omgo.NewLocation(config.Latitude, config.Longitude)
	if err != nil {
		return nil, err
	}

	opts := omgo.Options{
		TemperatureUnit:   "fahrenheit",
		WindspeedUnit:     "mph",
		PrecipitationUnit: "inch",
		Timezone:          config.TimeZone,
		HourlyMetrics: []string{
			"temperature_2m",
			"apparent_temperature",
			"wind_speed_10m",
			"windgusts_10m",
			"winddirection_10m",
			"cloudcover",
			"weathercode",
		},
	}

	forecast, err := c.Forecast(context.Background(), loc, &opts)
	if err != nil {
		return nil, err
	}

	return forecast, nil
}

func parseForecast(forecast *omgo.Forecast) ([]shared.Weather, error) {
	var weathers []shared.Weather

	for i, time := range forecast.HourlyTimes {
		weather := shared.Weather{
			DateTime:      time,
			Temperature:   parseNumeric(forecast.HourlyMetrics["temperature_2m"][i]),
			FeelsLike:     parseNumeric(forecast.HourlyMetrics["apparent_temperature"][i]),
			AvgWindSpeed:  parseNumeric(forecast.HourlyMetrics["wind_speed_10m"][i]),
			MaxWindSpeed:  parseNumeric(forecast.HourlyMetrics["windgusts_10m"][i]),
			WindDirection: parseWindDirection(forecast.HourlyMetrics["winddirection_10m"][i]),
			CloudCover:    parseNumeric(forecast.HourlyMetrics["cloudcover"][i]),
			WMOCode:       parseNumeric(forecast.HourlyMetrics["weathercode"][i]),
		}

		weathers = append(weathers, weather)
	}

	return weathers, nil
}

func parseNumeric(f float64) int {
	return int(math.Round(f))
}

func parseWindDirection(f float64) string {
	val := int((f / 22.5) + .5)
	dirs := []string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}
	return dirs[(val % 16)]
}
