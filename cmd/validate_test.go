package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {

	err := Validate("testdata/petstore.json")

	assert.Nilf(t, err, "spec was valid but returned error")

	err = Validate("testdata/petstore.yaml")
	assert.Nilf(t, err, "spec was valid but returned error")

	err = Validate("testdata/invalid.yaml")
	assert.NotNilf(t, err, "expected an error due to invalid spec, but parsed succesfully")

}
