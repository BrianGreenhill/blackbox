package cron

import (
	"github.com/robfig/cron/v3"

	"github.com/briangreenhill/blackbox/pkg/internal"
)

type CronScheduler struct {
	checker internal.Checker
	runner  *cron.Cron
}

func NewCronScheduler(checker internal.Checker) *CronScheduler {
	return &CronScheduler{
		checker: checker,
		runner:  cron.New(cron.WithSeconds()),
	}
}

func (s *CronScheduler) Schedule(schedule string, target internal.Target) error {
	_, err := s.runner.AddFunc(schedule, func() {
		_ = s.checker.DoCheck(target)
	})

	s.runner.Start()
	return err
}
