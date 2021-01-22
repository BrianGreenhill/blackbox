package web_test

import (
	"net/http"
	"testing"

	"github.com/briangreenhill/blackbox/pkg/internal"
	"github.com/briangreenhill/blackbox/pkg/mocks"
	"github.com/briangreenhill/blackbox/pkg/web"

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
	w := web.Checker{}
	w.Client = newClientMock(500)
	n := new(mocks.Notifier)
	checkResult := &internal.CheckResult{
		Message: "response was: 500",
	}
	n.On("Notify", checkResult).Return(nil)

	err := w.DoCheck(internal.Target{}, n)
	assert.NoError(t, err)
}
