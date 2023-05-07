package prompt

import (
	"crawl/internal/spec"
	"github.com/Netflix/go-expect"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPromptExpect(t *testing.T) {
	c, _ := expect.NewConsole()
	defer c.Close()
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		c.SendLine("https")
	}()
	<-donec

}

func TestPromptServer(t *testing.T) {
	t.Skipf("t.b.d. how to test stdinput")
	doc, err := spec.Validate("../testdata/petstore.json")
	assert.NoError(t, err)

	c, _ := expect.NewConsole()

	err, server := promptServer(*doc)
	c.SendLine("https")
	assert.NoError(t, err)
	assert.Equal(t, server.URL, "https://petstore.swagger.io/v2")
}

func TestPromptPath(t *testing.T) {
	doc, err := spec.Validate("../testdata/petstore.json")
	assert.NoError(t, err)

	// Hack stdin to use prompt_server.txt as input
	input, err := os.Open("testdata/prompt_path.txt")
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

	err, path := promptPath(*doc)
	assert.NoError(t, err)
	assert.Equal(t, path.Path, "/pets/")
}
