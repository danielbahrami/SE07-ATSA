package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/danielbahrami/se07-atsa/influxdb/broker"
	"github.com/danielbahrami/se07-atsa/influxdb/env"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func logRobotStatus(running bool, gpuProduced bool) {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://" + env.Get("INFLUXDB")
	client := influxdb2.NewClient(url, token)

	org := "my-org"
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

func logTestStatus(testName string, passed bool) {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://" + env.Get("INFLUXDB")
	client := influxdb2.NewClient(url, token)
	org := "my-org"
	bucket := "my-bucket"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	tags := map[string]string{
		"status": "test",
		"test":   testName,
	}
	fields := map[string]interface{}{
		"passed": passed,
	}

	point := write.NewPoint("test_status", tags, fields, time.Now())

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}
}

func main() {
	env.Load(".env")
	mqttBroker := broker.NewMQTT()
	if mqttBroker.Connect() {
		mqttBroker.Subscribe("topic/robot/status", func(message string) {
			switch message {
			case "RUNNING":
				logRobotStatus(true, false)
			case "GPUProduced":
				logRobotStatus(true, true)
			}
		})

		mqttBroker.Subscribe("topic/testing/gpu/completed", func(message string) {
			parts := strings.Split(message, ":")
			if len(parts) == 2 {
				testName := parts[0]
				result := parts[1]
				passed := result == "true"
				logTestStatus(testName, passed)
			}
		})
	}
	logRobotStatus(false, false)
	select {}
}
