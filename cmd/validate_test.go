package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func validate_test(t *testing.T) {

	err := validate("testdata/petstore.json")

	assert.Nilf(t, err, "spec was valid but returned error")

	err = validate("testdata/petstore.yaml")
	assert.Nilf(t, err, "spec was valid but returned error")

	err = validate("testdata/invalid.yaml")
	assert.NotNilf(t, err, "expected an error due to invalid spec, but parsed succesfully")

}
