package cmd

import (
	"greenhill/backend/pkg/internal"
	"greenhill/backend/pkg/stdout"
	"greenhill/backend/pkg/web"
	"log"
	"net/http"
	"os"
)

func Run(name, url string) {
	// create checker
	checker := web.Checker{
		Client: http.DefaultClient,
	}

	// create notifier
	notifier := stdout.NewNotifier(os.Stdout)

	target := internal.Target{
		Name: name,
		URL:  url,
	}

	err := checker.DoCheck(target, &notifier)
	if err != nil {
		log.Fatal(err)
	}
}
