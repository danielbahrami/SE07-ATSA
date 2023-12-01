package main

import (
	"atse/scheduler_system/api"
	"atse/scheduler_system/broker"
	"atse/scheduler_system/database"
	"atse/scheduler_system/env"
	"atse/scheduler_system/scheduler"
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
