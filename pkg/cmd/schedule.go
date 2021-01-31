package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/briangreenhill/blackbox/pkg/cron"
	"github.com/briangreenhill/blackbox/pkg/internal"
	"github.com/briangreenhill/blackbox/pkg/stdout"
	"github.com/briangreenhill/blackbox/pkg/web"
)

func RunSchedule(name, url, schedule string) {
	notifier := stdout.NewNotifier(os.Stdout)
	checker := web.Checker{
		Client:   http.DefaultClient,
		Notifier: &notifier,
	}

	target := internal.Target{
		Name: name,
		URL:  url,
	}

	err := checker.DoCheck(target)
	if err != nil {
		log.Fatal(err)
	}
	scheduler := cron.NewCronScheduler(&checker)
	scheduler.Schedule(schedule, target)
}
