package tests

import (
	"testing"

	"github.com/danielbahrami/se07-atsa/scheduling_system/scheduler"
)

func TestScheduler(t *testing.T) {
	t.Run("process scheduling", func(t *testing.T) {
		sch := scheduler.DummyScheduler{}
		prc := scheduler.Process{Id: "TESTER_ID"}
		sch.ScheduleProcess(prc)
	})
}
