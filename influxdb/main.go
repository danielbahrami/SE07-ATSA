package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/danielbahrami/se07-atsa/influxdb/broker"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {

	// Initialize MQTT broker
	mqttBroker := broker.NewMQTT()
	if mqttBroker.Connect() {
		mqttBroker.Subscribe("robot", func(message string) {
			if message == "RobotRunning" {
				logRobotStatus(true)
			}
		})
	}

	// Log 'robot' status into InfluxDB
	logRobotStatus(false) // Initially assuming the robot is not running

	// Keep the main function running to maintain MQTT subscription
	select {}
}

// Log 'robot' status into InfluxDB
func logRobotStatus(running bool) {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)

	org := "SE07-ATSA"
	bucket := "my-bucket"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	tags := map[string]string{
		"status": "robot",
	}
	fields := map[string]interface{}{
		"running": running,
	}

	point := write.NewPoint("robot_status", tags, fields, time.Now())

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}
}
