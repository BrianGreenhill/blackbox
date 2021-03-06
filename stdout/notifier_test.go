package stdout_test

import (
	"bytes"
	"testing"

	"github.com/briangreenhill/blackbox/internal"
	"github.com/briangreenhill/blackbox/stdout"

	"github.com/stretchr/testify/assert"
)

func TestWriterWritesToStdout(t *testing.T) {
	buf := new(bytes.Buffer)
	usecase := stdout.NewNotifier(buf)
	expected := &internal.CheckResult{
		Message: "test message",
	}

	err := usecase.Notify(expected)
	assert.NoError(t, err)
	actual := buf.String()
	assert.Equal(t, "test message\n", actual)
}
