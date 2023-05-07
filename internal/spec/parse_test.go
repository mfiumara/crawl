package spec

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
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

	options, serverMap := GetOptionsAndServers(doc)

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
		"GET     /foo | Get Foo":    {Path: "/foo", Item: *path1},
		"POST    /bar | Create Foo": {Path: "/bar", Item: *path2},
	}

	options, pathMap := GetOptionsAndPaths(doc)

	assert.Equal(t, expectedOptions, options)
	assert.Equal(t, expectedPathMap, pathMap)
}

func TestValidate(t *testing.T) {

	_, err := Validate("../testdata/petstore.json")
	assert.Nilf(t, err, "spec was valid but returned error")

	_, err = Validate("../testdata/petstore.yaml")
	assert.Nilf(t, err, "spec was valid but returned error")

	_, err = Validate("nonexistent.yaml")
	assert.NotNilf(t, err, "expected an error due to invalid path, but parsed succesfully")

	_, err = Validate("../testdata/invalid.yaml")
	assert.NotNilf(t, err, "expected an error due to invalid spec, but parsed succesfully")
}

func TestListMethods(t *testing.T) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	doc, err := loader.LoadFromFile("../testdata/petstore.json")
	err = ListMethods(*doc)
	assert.Nilf(t, err, "spec was valid but returned error")
}

func TestListSecurity(t *testing.T) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	doc, err := loader.LoadFromFile("../testdata/petstore.json")
	err = ListSecurity(*doc)
	assert.Nilf(t, err, "spec was valid but returned error")
}
