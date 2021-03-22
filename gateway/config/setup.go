package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var cfg = readConfigs()
var Clients = cfg.Clients
var Devices = cfg.Devices

type sensorCfg struct {
	Topic string `yaml:"topic"`
}

type temperatureSensorCfg struct {
	Topic    string `yaml:"topic"`
	Settings struct {
		Humidity struct {
			From        float64 `yaml:"from"`
			To          float64 `yaml:"to"`
			NormalDelta float64 `yaml:"normal_delta"`
		} `yaml:"humidity"`
		Temperature struct {
			NormalDelta float64 `yaml:"normal_delta"`
			ColdsMonths struct {
				From   float64 `yaml:"from"`
				To     float64 `yaml:"to"`
				Months []int64 `yaml:"months"`
			} `yaml:"cold_months"`
			HotMonths struct {
				From   float64 `yaml:"from"`
				To     float64 `yaml:"to"`
				Months []int64 `yaml:"months"`
			} `yaml:"hot_months"`
		} `yaml:"temperature"`
	} `yaml:"settings"`
}

type devicesCfg struct {
	Door        sensorCfg            `yaml:"door"`
	Temperature temperatureSensorCfg `yaml:"temperature"`
	Light       sensorCfg            `yaml:"light"`
	Siren       sensorCfg            `yaml:"siren"`
	Switch      sensorCfg            `yaml:"switch"`
}

type clientsCfg struct {
	Mqtt struct {
		BrokerUrl string `yaml:"broker_url"`
		ClientId  string `yaml:"client_id"`
	} `yaml:"mqtt"`
}

type gatewayCfg struct {
	Clients clientsCfg `yaml:"clients"`
	Devices devicesCfg `yaml:"devices"`
}

func readConfigs() *gatewayCfg {
	config := &gatewayCfg{}
	file, err := os.Open("configs/gateway.yaml")
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
