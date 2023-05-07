package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	// redirect output to buffer
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	// set path to example spec
	path = "../testdata/petstore.yaml"

	// execute command
	args := []string{"--path", path}
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()

	assert.NoError(t, err)

	// verify output
	assert.Contains(t, buf.String(), rootCmd.Long)

	Execute()
}

func TestRootPreRun(t *testing.T) {
	t.Skipf("t.b.d. how to test stdinput")
	// Hack stdin to use prompt_server.txt as input
	input, err := os.Open("testdata/root.txt")
	assert.NoError(t, err)
	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
		err := input.Close()
		if err != nil {
			return
		}
	}()
	os.Stdin = input

	rootPreRun(nil, nil)
}
