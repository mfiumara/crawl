package cmd

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetOptionsAndServers(t *testing.T) {
	doc := openapi3.T{
		Servers: []*openapi3.Server{
			{URL: "http://localhost:8000", Description: "Local Server"},
			{URL: "https://api.example.com", Description: "Production Server"},
		},
	}

	expectedOptions := []string{
		"Local Server      | http://localhost:8000",
		"Production Server | https://api.example.com",
	}

	expectedServerMap := map[string]openapi3.Server{
		"Local Server      | http://localhost:8000":   {URL: "http://localhost:8000", Description: "Local Server"},
		"Production Server | https://api.example.com": {URL: "https://api.example.com", Description: "Production Server"},
	}

	options, serverMap := getOptionsAndServers(doc)

	assert.Equal(t, expectedOptions, options)
	assert.Equal(t, expectedServerMap, serverMap)
}

func TestGetOptionsAndPaths(t *testing.T) {
	doc := openapi3.T{
		Paths: make(map[string]*openapi3.PathItem),
	}

	path1 := &openapi3.PathItem{
		Get: &openapi3.Operation{Summary: "Get Foo"},
	}
	path2 := &openapi3.PathItem{
		Post: &openapi3.Operation{Summary: "Create Foo"},
	}
	doc.Paths["/foo"] = path1
	doc.Paths["/bar"] = path2

	expectedOptions := []string{
		"GET     /foo | Get Foo",
		"POST    /bar | Create Foo",
	}

	expectedPathMap := map[string]Path{
		"GET     /foo | Get Foo":    {path: "/foo", item: *path1},
		"POST    /bar | Create Foo": {path: "/bar", item: *path2},
	}

	options, pathMap := getOptionsAndPaths(doc)

	assert.Equal(t, expectedOptions, options)
	assert.Equal(t, expectedPathMap, pathMap)
}

func TestPromptServer(t *testing.T) {
	doc, err := Validate("testdata/petstore.json")
	assert.NoError(t, err)

	// Hack stdin to use prompt_server.txt as input
	input, err := os.Open("testdata/prompt_server.txt")
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

	err, server := promptServer(*doc)
	assert.NoError(t, err)
	assert.Equal(t, server.URL, "https://petstore.swagger.io/v2")
}

func TestPromptPath(t *testing.T) {
	doc, err := Validate("testdata/petstore.json")
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
	assert.Equal(t, path.path, "/pets/")
}
