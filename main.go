package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"github.com/yosssi/gmq/mqtt"
	mqttClient "github.com/yosssi/gmq/mqtt/client"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	configErr := viper.ReadInConfig()
	if configErr != nil {
		panic(fmt.Errorf("fatal error config file: %s", configErr))
	}

	mqttHost := viper.GetString("mqtt.host")
	mqttClientID := []byte(viper.GetString("mqtt.client_id"))
	mqttUser := []byte(viper.GetString("mqtt.user"))
	mqttPassword := []byte(viper.GetString("mqtt.password"))
	mqttTopic := []byte(viper.GetString("mqtt.topic"))
	mqttStatus := []byte(viper.GetString("mqtt.status"))

	mqttCli := mqttClient.New(&mqttClient.Options{
		ErrorHandler: func(err error) {
			log.Fatal(err)
		},
	})

	defer mqttCli.Terminate()
	mqttErr := mqttCli.Connect(&mqttClient.ConnectOptions{
		Network:  "tcp",
		Address:  mqttHost,
		ClientID: mqttClientID,
		UserName: mqttUser,
		Password: mqttPassword,
	})

	if mqttErr != nil {
		log.Fatal(mqttErr)
	}

	mqttErr = mqttCli.Publish(&mqttClient.PublishOptions{
		QoS:       mqtt.QoS0,
		Retain:    true,
		TopicName: []byte(mqttStatus),
		Message:   []byte("Ready!"),
	})

	if mqttErr != nil {
		log.Fatal(mqttErr)
	}

	for {
		forecastData := getForecast()

		earliestForecast := DailyForecast{}
		earliestTimestamp := 0
		for _, forecast := range forecastData.Daily.Data {
			if earliestTimestamp == 0 || forecast.Time < earliestTimestamp {
				earliestTimestamp = forecast.Time
				earliestForecast = forecast
			}
		}

		message := MQTTForecast{
			MinTemp: earliestForecast.ApparentTemperatureMin,
			MaxTemp: earliestForecast.ApparentTemperatureMax,
			Icon:    earliestForecast.Icon,
		}
		log.Println(message)
		weatherString, error := json.Marshal(message)
		if error != nil {
			fmt.Println(error)
			return
		}

		mqttCli.Publish(&mqttClient.PublishOptions{
			QoS:       mqtt.QoS0,
			Retain:    true,
			TopicName: []byte(mqttTopic),
			Message:   []byte(weatherString),
		})

		time.Sleep(5 * time.Minute)
	}
}
