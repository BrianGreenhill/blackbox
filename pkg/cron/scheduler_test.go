package cron_test

import (
	"testing"
	"time"

	"github.com/briangreenhill/blackbox/pkg/cron"
	"github.com/briangreenhill/blackbox/pkg/internal"
	"github.com/briangreenhill/blackbox/pkg/mocks"
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
	checkerMock.Wg.Add(1)
	start := time.Now()
	checkerMock.On("DoCheck", mock.Anything, mock.Anything).Return(nil)
	usecase := cron.NewCronScheduler(checkerMock)
	usecase.Schedule(scheduleEverySecond, target)
	checkerMock.Wg.Wait()
	end := time.Now()
	diff := end.Sub(start)
	assert.GreaterOrEqual(t, diff, time.Second)

	checkerMock.AssertExpectations(t)
}
