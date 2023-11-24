package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

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
