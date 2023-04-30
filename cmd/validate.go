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
	Long:  `Validates whether an openAPI specification is valid or not. Can give a URL or a local file path as input.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Validates an openapi spec
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
