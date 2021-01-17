package stdout

import (
	"bytes"
	"greenhill/backend/pkg/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriterWritesToStdout(t *testing.T) {
	buf := new(bytes.Buffer)
	usecase := NewNotifier(buf)
	expected := &internal.CheckResult{
		Message: "test message",
	}

	err := usecase.Notify(expected)
	assert.NoError(t, err)
	actual := string(buf.Bytes())
	assert.Equal(t, "test message", actual)
}
