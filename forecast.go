package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func getForecast() WeatherForecast {

	key := viper.GetString("darksky.key")
	coordinates := viper.GetString("darksky.coordinates")
	queryString := viper.GetString("darksky.query_string")

	forecast := WeatherForecast{}

	response, err := http.Get("https://api.darksky.net/forecast/" + key + "/" + coordinates + queryString)
	if err != nil {
		log.Fatal(err)
		return forecast
	}

	jsonBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return forecast
	}

	unmarshalErr := json.Unmarshal(jsonBody, &forecast)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
		return forecast
	}

	return forecast
}
