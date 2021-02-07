package stdout

import (
	"fmt"
	"io"

	"github.com/briangreenhill/blackbox/internal"
)

type Notifier struct {
	stdout io.Writer
}

func NewNotifier(w io.Writer) *Notifier {
	return &Notifier{
		stdout: w,
	}
}

func (n *Notifier) Notify(result *internal.CheckResult) error {
	_, err := fmt.Fprintln(n.stdout, result.Message)
	return err
}
