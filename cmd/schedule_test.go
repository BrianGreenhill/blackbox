package cmd_test

import (
	"errors"
	"testing"

	"github.com/briangreenhill/blackbox/cmd"
	"github.com/briangreenhill/blackbox/mocks"
	"github.com/stretchr/testify/assert"
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
			var err error
			if tt.errorExpected {
				err = errors.New("you suck")
			}
			scheduler.On("Schedule", tt.schedule).Return(err)

			scheduleCommand := cmd.ScheduleCommand{
				Scheduler: scheduler,
			}
			err = scheduleCommand.RunSchedule(tt.schedule)
			if tt.errorExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			scheduler.AssertExpectations(t)
		})
	}
}
