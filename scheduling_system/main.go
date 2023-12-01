package main

import (
	"github.com/danielbahrami/se07-atsa/scheduling_system/api"
	"github.com/danielbahrami/se07-atsa/scheduling_system/broker"
	"github.com/danielbahrami/se07-atsa/scheduling_system/database"
	"github.com/danielbahrami/se07-atsa/scheduling_system/env"
	"github.com/danielbahrami/se07-atsa/scheduling_system/scheduler"
)

func main() {
	env.Load(".env")

	pg := database.NewPG(
		env.Get("PG_USER"),
		env.Get("PG_PASSWORD"),
		env.Get("PG_HOST"),
		env.Get("PG_PORT"),
		env.Get("PG_DATABASE"))

	b := api.NewBuilder()
	b.Database(&pg)
	b.Scheduler(scheduler.DummyScheduler{})
	b.Broker(broker.NewMQTT())
	sys := b.Build()
	sys.Start()

}
