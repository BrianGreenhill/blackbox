package web_test

import (
	"net/http"
	"testing"

	"github.com/briangreenhill/blackbox/internal"
	"github.com/briangreenhill/blackbox/mocks"
	"github.com/briangreenhill/blackbox/web"

	"github.com/stretchr/testify/assert"
)

type ClientMock struct {
	ExpectedStatusCode int
}

func newClientMock(status int) *ClientMock {
	return &ClientMock{
		ExpectedStatusCode: status,
	}
}

func (c *ClientMock) Do(r *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: c.ExpectedStatusCode,
	}, nil
}

func TestCheckReturnsCheckResultOnFailure(t *testing.T) {
	n := new(mocks.Notifier)
	w := web.Checker{
		Notifier: n,
		Client:   newClientMock(500),
	}
	checkResult := &internal.CheckResult{
		Message: "response was: 500",
	}
	n.On("Notify", checkResult).Return(nil)

	err := w.DoCheck(internal.Target{})
	assert.NoError(t, err)
}
