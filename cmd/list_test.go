package cmd

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListMethods(t *testing.T) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	doc, err := loader.LoadFromFile("testdata/petstore.json")
	err = ListMethods(*doc)
	assert.Nilf(t, err, "spec was valid but returned error")

	doc, err = loader.LoadFromFile("testdata/invalid.yaml")
	assert.NotNilf(t, err, "expected an error due to invalid spec, but parsed succesfully")
}

func TestListSecurity(t *testing.T) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	doc, err := loader.LoadFromFile("testdata/petstore.json")
	err = ListSecurity(*doc)
	assert.Nilf(t, err, "spec was valid but returned error")

}
