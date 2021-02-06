package cron

import (
	"github.com/robfig/cron/v3"

	"github.com/briangreenhill/blackbox/internal"
)

type CronScheduler struct {
	checker internal.Checker
	runner  *cron.Cron
	target  internal.Target
}

func NewCronScheduler(checker internal.Checker, target internal.Target) *CronScheduler {
	return &CronScheduler{
		checker: checker,
		runner:  cron.New(cron.WithSeconds()),
		target:  target,
	}
}

func (s *CronScheduler) Schedule(schedule string) error {
	_, err := s.runner.AddFunc(schedule, func() {
		_ = s.checker.DoCheck(s.target)
	})

	s.runner.Start()
	return err
}
