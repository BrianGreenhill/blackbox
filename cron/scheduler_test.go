package cron_test

import (
	"testing"
	"time"

	"github.com/briangreenhill/blackbox/cron"
	"github.com/briangreenhill/blackbox/internal"
	"github.com/briangreenhill/blackbox/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestScheduler(t *testing.T) {
	scheduleEverySecond := "* * * * * *"
	target := internal.Target{
		Name: "test",
		URL:  "test.com",
	}

	checkerMock := new(mocks.ConcurrentChecker)
	checkerMock.Wg.Add(2)
	start := time.Now()
	checkerMock.On("DoCheck", mock.Anything, mock.Anything).Return(nil)
	usecase := cron.NewCronScheduler(checkerMock, target)
	err := usecase.Schedule(scheduleEverySecond)
	assert.Nil(t, err)
	checkerMock.Wg.Wait()
	end := time.Now()
	diff := end.Sub(start)

	assert.GreaterOrEqual(t, diff, 1*time.Second)

	checkerMock.AssertExpectations(t)
}
