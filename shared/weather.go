package shared

import "time"

type Weather struct {
	DateTime      time.Time
	Temperature   int
	FeelsLike     int
	AvgWindSpeed  int
	MaxWindSpeed  int
	WindDirection string
	CloudCover    int
	WMOCode       int // Above 50 is precipitation
}

type FlyabilityRating struct {
	Weather Weather
	Rating  int // 1-10
}
