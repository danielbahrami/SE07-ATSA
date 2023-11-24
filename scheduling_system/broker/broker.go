package broker

import (
	"atse/scheduler_system/env"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type IBroker interface {
	Connect() bool
	Message(topic string, message string)
}

type Broker struct {
	client mqtt.Client
}

func (b *Broker) Connect() bool {
	log.Println("Connecting to message broker ...")
	brokerAddr := env.Get("BROKER")
	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcp://%s", brokerAddr))
	options.SetClientID("atse_scheduling_system")
	options.OnConnect = onConnect
	options.OnConnectionLost = onConnectionLost

	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Printf("Could not connect to message broker {%s}\n", token.Error())
		return false
	}
	return true
}

func (b *Broker) Message(topic string, message string) {
	if b.client != nil {
		token := b.client.Publish(topic, 1, false, message)
		token.Wait()
	}
}

func onConnect(client mqtt.Client) {
	log.Println("Connection established")
}

func onConnectionLost(client mqtt.Client, err error) {
	log.Printf("Connection lost {%v}\n", err)
}
