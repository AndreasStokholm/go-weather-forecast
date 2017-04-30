package main

// WeatherForecast - Darksky forecast root object
type WeatherForecast struct {
	Latitude  float64
	Longitude float64
	Timezone  string
	TZOffset  int `json:"offset"`
	Daily     DailyForecastList
}

// DailyForecastList - Darksky forecast list structure
type DailyForecastList struct {
	Summary string
	Icon    string
	Data    []DailyForecast
}

// DailyForecast - Darksky forecast structure
type DailyForecast struct {
	Time                       int
	Summary                    string
	Icon                       string
	SunriseTime                int
	SunsetTime                 int
	MoonPhase                  float64
	PrecipIntensity            float64
	PrecipIntensityMax         float64
	PrecitIntensityMaxTime     int
	PrecipProbability          float64
	PrecipType                 string
	TemperatureMin             float64
	TemperatureMinTime         int
	TemperatureMax             float64
	TemperatureMaxTime         int
	ApparentTemperatureMin     float64
	ApparentTemperatureMinTime int
	ApparentTemperatureMax     float64
	ApparentTemperatureMaxTime int
	DewPoint                   float64
	Humidity                   float64
	WindSpeed                  float64
	WindBearing                int
	Visibility                 float64
	CloudCover                 float64
	Pressure                   float64
	Ozone                      float64
}

// MQTTForecast - Darksky data exposed to MQTT
type MQTTForecast struct {
	MinTemp float64 `json:"min_temp"`
	MaxTemp float64 `json:"max_temp"`
	Icon    string  `json:"icon"`
}
