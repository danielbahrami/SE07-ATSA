package scheduler

import "log"

type Process struct {
	Id string
}

type IScheduler interface {
	ScheduleProcess(process Process)
}

type DummyScheduler struct {
}

func (s DummyScheduler) ScheduleProcess(process Process) {
	log.Printf("Scheduling process {%s}\n", process.Id)
}
