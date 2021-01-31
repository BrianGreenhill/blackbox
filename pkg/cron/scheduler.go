package cron

import (
	"github.com/briangreenhill/blackbox/pkg/internal"
)

type CronScheduler struct {
	checker internal.Checker
}

func NewCronScheduler(checker internal.Checker) *CronScheduler {
	return &CronScheduler{
		checker: checker,
	}
}

func (s *CronScheduler) Schedule(schedule string, target internal.Target) error {
	return s.checker.DoCheck(target)
}
