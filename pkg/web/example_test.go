package web_test

import (
	"log"
	"net/http"
	"os"

	"github.com/briangreenhill/blackbox/pkg/internal"
	"github.com/briangreenhill/blackbox/pkg/stdout"
	"github.com/briangreenhill/blackbox/pkg/web"
)

func Example() {
	notifier := stdout.NewNotifier(os.Stdout)

	target := internal.Target{
		Name: "GitHub",
		URL:  "https://www.github.com",
	}

	checker := web.Checker{
		Client: http.DefaultClient,
	}

	err := checker.DoCheck(target, &notifier)
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	//  response was: 200
}
