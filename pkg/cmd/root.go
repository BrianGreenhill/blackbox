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
