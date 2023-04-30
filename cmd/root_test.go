package cmd

import (
	"bytes"
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
