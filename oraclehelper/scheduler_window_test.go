package oraclehelper

import (
	"testing"
)

func TestSchedulerWindowService(t *testing.T) {

	tstWindow := ResourceSchedulerWindow{
		WindowName:     "TEST01",
		ResourcePlan:   "INTERNAL_PLAN",
		Duration:       "numtodsinterval(60, 'minute')",
		RepeatInterval: "freq=daily;byday=SAT;byhour=6;byminute=0; bysecond=0",
		WindowPriority: "LOW",
		Comments:       "test01 commments",
	}
	c.SchedulerWindowService.CreateSchedulerWindow(tstWindow)
	windows, err := c.SchedulerWindowService.ReadSchedulerWindow(ResourceSchedulerWindow{Owner: "SYS", WindowName: "TEST01"})
	if err != nil {
		t.Error("Failed to read window")
	}
	if windows.WindowName != tstWindow.WindowName {
		t.Error("window name not equal")

	}
	c.SchedulerWindowService.DropSchedulerWindow(ResourceSchedulerWindow{WindowName: "TEST01"})
}
