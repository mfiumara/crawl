package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListMethods(t *testing.T) {

	err := ListMethods("testdata/petstore.json")
	assert.Nilf(t, err, "spec was valid but returned error")

	err = ListMethods("testdata/invalid.yaml")
	assert.NotNilf(t, err, "expected an error due to invalid spec, but parsed succesfully")
}
