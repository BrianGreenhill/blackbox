package web_test

import (
	"greenhill/backend/pkg/internal"
	"greenhill/backend/pkg/mocks"
	"greenhill/backend/pkg/web"
	"net/http"
	"testing"

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
