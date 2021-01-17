package web

import (
	"fmt"
	"net"
	"net/http"

	"greenhill/backend/pkg/internal"
)

type Checker struct {
	Client Doer
}

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

func (w *Checker) DoCheck(t internal.Target, n internal.Notifier) error {
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
			return n.Notify(msg)
		}
		return err
	}

	statuscode := resp.StatusCode
	var statuscodemsg string
	switch statuscode {
	case 301:
		statuscodemsg = "We were redirected.\n"
		break
	}

	msg.Message = fmt.Sprintf("%sresponse was: %d", statuscodemsg, statuscode)
	return n.Notify(msg)
}
