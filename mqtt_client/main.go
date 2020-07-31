package mqtt_client

import (
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

var (
	client mqtt.Client
	broker = os.Getenv("BROKER")
	clientId = os.Getenv("CLIENT_ID")
)

var defaultHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {}

func SetUpClient() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientId)

	opts.SetKeepAlive(120 * time.Second)
	opts.SetDefaultPublishHandler(defaultHandler)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	setUpSubscriptions()
}

func Disconnect() {
	client.Disconnect(250)
}

func setUpSubscriptions() {}
