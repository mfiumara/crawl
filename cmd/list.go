package cmd

import (
	"context"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
		doc, err := loader.LoadFromFile(path)
		if err != nil {
			println("Could not load spec from path: ", path)
			println(err.Error())
			os.Exit(1)
		}
		err = ListMethods(*doc)
		if err != nil {
			println("Could not load spec from path: ", path)
			println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func ListMethods(doc openapi3.T) error {

	// Set the printing format
	opLen := 0
	for key := range doc.Paths {
		if len(key) > opLen {
			opLen = len(key)
		}
	}
	format := fmt.Sprintf("%%-%ds %%-%ds | %%s\n", 7, opLen)

	for key, value := range doc.Paths {
		for method, operation := range value.Operations() {
			fmt.Printf(format, method, key, operation.Summary)
		}
	}
	return nil
}

func ListSecurity(doc openapi3.T) error {
	for key, value := range doc.Components.SecuritySchemes {
		fmt.Printf("name: %s | type: %s\n", key, value.Value.Type)
	}
	return nil
}
