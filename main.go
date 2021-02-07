package main

import (
	"flag"
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

var (
	url      string
	schedule string
	name     string
)

func init() {
	flag.StringVar(&name, "name", os.Getenv("TARGET_NAME"), "The name of the check")
	flag.StringVar(&url, "url", os.Getenv("TARGET_URL"), "The URL to check")
	flag.StringVar(&schedule, "schedule", os.Getenv("SCHEDULE"), "The schedule in six star format")
}

func main() {
	flag.Parse()
	if url == "" {
		log.Fatal("url is required")
	}

	checker := &web.Checker{
		Client:   http.DefaultClient,
		Notifier: stdout.NewNotifier(os.Stdout),
	}

	target := internal.Target{
		Name: name,
		URL:  url,
	}

	scheduleCommand := cmd.ScheduleCommand{
		Scheduler: cron.NewCronScheduler(checker, target),
	}

	if err := scheduleCommand.RunSchedule(schedule); err != nil {
		log.Fatal(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	s := <-signals
	log.Printf("got %s signal, shutting down", s.String())
	os.Exit(0)
}
