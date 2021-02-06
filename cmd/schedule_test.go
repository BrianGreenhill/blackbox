package cmd_test

import (
	"errors"
	"testing"

	"github.com/briangreenhill/blackbox/cmd"
	"github.com/briangreenhill/blackbox/mocks"
)

func TestScheduleCommand(t *testing.T) {
	var tests = []struct {
		name          string
		schedule      string
		errorExpected bool
	}{
		{
			name:          "it schedules a check correctly",
			schedule:      "* * * * * *",
			errorExpected: false,
		},
		{
			name:          "it fails if the schedule is wrong",
			schedule:      "this is not a schedule",
			errorExpected: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			scheduler := new(mocks.Scheduler)
			if tt.errorExpected == false {
				scheduler.On("Schedule", tt.schedule).Return(nil)
			} else {
				scheduler.On("Schedule", tt.schedule).Return(errors.New("you suck"))
			}

			scheduleCommand := cmd.ScheduleCommand{
				Scheduler: scheduler,
			}
			scheduleCommand.RunSchedule(tt.schedule)
			scheduler.AssertExpectations(t)
		})
	}
}
