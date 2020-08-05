package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var Cfg = NewConfig()

type SensorConfig struct {
	Topic string 						  `yaml:"topic"`
}

type Config struct {
	MqttClient struct {
		BrokerUrl string                  `yaml:"broker_url"`
		ClientId  string                  `yaml:"client_id"`
		Sensors struct {
			Door SensorConfig 			  `yaml:"door"`
			Temperature SensorConfig 	  `yaml:"temperature"`
		} 								  `yaml:"sensors"`
	} 									  `yaml:"mqtt_client"`
	TgBot struct {
		ApiToken         string   		  `yaml:"api_token"`
	} 									  `yaml:"tg_bot"`
}

func NewConfig() *Config {
	config := &Config{}
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal("Config file was not found")
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		log.Fatal("Can not decode config file")
	}
	return config
}
