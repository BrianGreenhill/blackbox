package cron

import (
	"log"

	"github.com/robfig/cron/v3"

	"github.com/briangreenhill/blackbox/internal"
)

type Scheduler struct {
	checker internal.Checker
	runner  *cron.Cron
	target  internal.Target
}

func NewCronScheduler(checker internal.Checker, target internal.Target) *Scheduler {
	return &Scheduler{
		checker: checker,
		runner:  cron.New(cron.WithSeconds()),
		target:  target,
	}
}

func (s *Scheduler) Schedule(schedule string) error {
	_, err := s.runner.AddFunc(schedule, func() {
		err := s.checker.DoCheck(s.target)
		if err != nil {
			log.Println(err)
		}
	})

	s.runner.Start()
	return err
}
