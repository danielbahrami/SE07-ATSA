package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"../scheduling_system/env"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type IBroker interface {
	Connect() bool
	Message(topic string, message string)
}

type Broker struct {
	client mqtt.Client
}

func (b Broker) Connect() bool {
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

func (b Broker) Message(topic string, message string) {
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

func main() {
	// Initialize client
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)

	// Write data
	org := "SE07-ATSA"
	bucket := "my-bucket"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 5; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
	}

	// Execute a flux query
	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "my-bucket")
			  |> range(start: -10m)
			  |> filter(fn: (r) => r._measurement == "measurement1")`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}

	// Execute an Aggregate Query
	query = `from(bucket: "my-bucket")
              |> range(start: -10m)
              |> filter(fn: (r) => r._measurement == "measurement1")
              |> mean()`
	results, err = queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
}
