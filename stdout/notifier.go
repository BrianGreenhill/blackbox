package stdout

import (
	"fmt"
	"io"

	"github.com/briangreenhill/blackbox/internal"
)

type notifier struct {
	stdout io.Writer
}

func NewNotifier(w io.Writer) *notifier {
	return &notifier{
		stdout: w,
	}
}

func (n *notifier) Notify(result *internal.CheckResult) error {
	_, err := fmt.Fprintln(n.stdout, result.Message)
	return err
}
