/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate an openAPI spec definition",
	Long:  `Validates whether the openAPI specification is valid or not.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Validate(path)
		if err != nil {
			println("Could not load spec: ", err.Error())
			os.Exit(1)
		}
		println("✅ Spec valid")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

// Validate Validates an openapi spec
func Validate(path string) (*openapi3.T, error) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc, err := loader.LoadFromFile(path)
	if err != nil {
		return nil, err
	}

	// Validate document
	err = doc.Validate(ctx)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
