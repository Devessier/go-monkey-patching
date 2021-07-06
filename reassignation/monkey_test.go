package reassignation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	assert := assert.New(t)

	const (
		input          = "hi"
		expectedOutput = input + "!"
	)

	result := HelloWorld(input)
	assert.Equal(expectedOutput, result)
}

func TestHelloWorldWithMonkeyPatchedFormatter(t *testing.T) {
	assert := assert.New(t)

	const (
		input          = "hi"
		expectedOutput = input + "?"
	)

	originalImplementation := PrinterFn
	PrinterFn = func(s string) string {
		return s + "?"
	}
	defer func() {
		PrinterFn = originalImplementation
	}()

	result := HelloWorld(input)
	assert.Equal(expectedOutput, result)
}
