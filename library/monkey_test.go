package library

import (
	"fmt"
	"testing"

	"bou.ke/monkey"
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

	monkey.Patch(Formatter, func(s string) string {
		return s + "?"
	})
	defer monkey.UnpatchAll()

	fmt.Printf("result = %s\n", Formatter("lol"))

	result := HelloWorld(input)
	assert.Equal(expectedOutput, result)
}
