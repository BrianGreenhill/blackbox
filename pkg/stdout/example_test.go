package stdout_test

import (
	"log"
	"os"

	"github.com/briangreenhill/blackbox/pkg/internal"
	"github.com/briangreenhill/blackbox/pkg/stdout"
)

func Example() {
	result := &internal.CheckResult{
		Message: "response was: 500",
	}

	notifier := stdout.NewNotifier(os.Stdout)

	err := notifier.Notify(result)
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// response was: 500
}
