package cmd

import (
	"github.com/briangreenhill/blackbox/internal"
)

type ScheduleCommand struct {
	Scheduler internal.Scheduler
	Target    internal.Target
}

func (s *ScheduleCommand) RunSchedule(schedule string) error {
	return s.Scheduler.Schedule(schedule)
}
