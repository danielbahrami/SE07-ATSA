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

func logRobotStatus(running bool, gpuProduced bool) {
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
		"running":      running,
		"gpu_produced": gpuProduced,
	}

	point := write.NewPoint("robot_status", tags, fields, time.Now())

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Initialize MQTT broker
	mqttBroker := broker.NewMQTT()
	if mqttBroker.Connect() {
		mqttBroker.Subscribe("robot", func(message string) {
			switch message {
			case "RobotRunning":
				logRobotStatus(true, false) // Robot is running, GPU not produced yet
			case "GPUProduced":
				logRobotStatus(true, true) // Robot produced a GPU
			}
		})
	}

	// Log 'robot' status into InfluxDB
	logRobotStatus(false, false) // Initially assuming the robot hasn't produced GPU

	// Keep the main function running to maintain MQTT subscription
	select {}
}
