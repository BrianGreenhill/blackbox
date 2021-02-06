package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/briangreenhill/blackbox/cmd"
	"github.com/briangreenhill/blackbox/cron"
	"github.com/briangreenhill/blackbox/internal"
	"github.com/briangreenhill/blackbox/stdout"
	"github.com/briangreenhill/blackbox/web"
)

const URL = "https://www.google.com"
const ScheduleFormat = "* * * * * *"

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	checker := &web.Checker{
		Client:   http.DefaultClient,
		Notifier: stdout.NewNotifier(os.Stdout),
	}

	target := internal.Target{
		Name: "Google",
		URL:  URL,
	}

	scheduleCommand := cmd.ScheduleCommand{
		Scheduler: cron.NewCronScheduler(checker, target),
	}

	if err := scheduleCommand.RunSchedule(ScheduleFormat); err != nil {
		log.Fatal(err)
	}

	s := <-signals
	log.Printf("got %s signal, shutting down", s.String())
	os.Exit(0)
}
