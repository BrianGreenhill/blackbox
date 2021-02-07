package web

import (
	"fmt"
	"net"
	"net/http"

	"github.com/briangreenhill/blackbox/internal"
)

// Checker performs simple site availability checks
type Checker struct {
	Client   Doer
	Notifier internal.Notifier
}

// Doer defines an interface for a client that can prform an availability check
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// DoCheck is used to check the availability of a target and uses a notifier to communicate the result
func (w *Checker) DoCheck(t internal.Target) error {
	var msg = &internal.CheckResult{Message: ""}
	req, err := http.NewRequest(http.MethodGet, t.URL, nil)
	if err != nil {
		return err
	}
	resp, err := w.Client.Do(req)
	if err != nil {
		_, err := net.LookupHost(t.URL)
		if err != nil {
			msg.Message = "That host did not exist"
			return w.Notifier.Notify(msg)
		}
		return err
	}
	defer resp.Body.Close()

	statuscode := resp.StatusCode
	var statuscodemsg string
	if statuscode == 301 {
		statuscodemsg = "We were redirected.\n"
	}

	msg.Message = fmt.Sprintf("%sresponse was: %d", statuscodemsg, statuscode)
	return w.Notifier.Notify(msg)
}
