package tests

import (
	"atse/scheduler_system/scheduler"
	"testing"
)

func TestScheduler(t *testing.T) {
	t.Run("process scheduling", func(t *testing.T) {
		sch := scheduler.DummyScheduler{}
		prc := scheduler.Process{Id: "TESTER_ID"}
		sch.ScheduleProcess(prc)
	})
}
