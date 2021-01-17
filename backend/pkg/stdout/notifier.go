package stdout

import (
	"fmt"
	"io"

	"greenhill/backend/pkg/internal"
)

type notifier struct {
	stdout io.Writer
}

func NewNotifier(w io.Writer) notifier {
	return notifier{
		stdout: w,
	}
}

func (n *notifier) Notify(result *internal.CheckResult) error {
	_, err := fmt.Fprintf(n.stdout, result.Message)
	return err
}
