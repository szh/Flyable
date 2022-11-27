package weatherfetcher

import (
	"env"
	"errors"
)

type Config struct {
	Latitude      float64
	Longitude     float64
	LocationName  string
	TimeZone      string
	WeatherApiUrl string
}

func (c Config) Validate() error {
	if c.Latitude == 0 {
		return errors.New("Latitude is required")
	}

	if c.Longitude == 0 {
		return errors.New("Longitude is required")
	}

	return nil
}

func LoadConfigFromEnv() (Config, error) {
	var config Config
	config.Latitude = env.Float64("LATITUDE", 0)
	config.Longitude = env.Float64("LONGITUDE", 0)
	config.LocationName = env.String("LOCATION_NAME", "")
	config.TimeZone = env.String("TIME_ZONE", "UTC")
	config.WeatherApiUrl = env.String("WEATHER_API_URL", "https://api.open-meteo.com/v1/forecast")

	err := config.Validate()
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
