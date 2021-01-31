package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/briangreenhill/blackbox/pkg/internal"
	"github.com/briangreenhill/blackbox/pkg/stdout"
	"github.com/briangreenhill/blackbox/pkg/web"
)

// Run is the entrypoint to the application
func Run(name, url string) {
	// create notifier
	notifier := stdout.NewNotifier(os.Stdout)

	// create checker
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
}
