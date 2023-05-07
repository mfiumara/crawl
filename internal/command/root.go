package command

import (
	"crawl/internal/spec"
	"github.com/AlecAivazis/survey/v2"
	"github.com/getkin/kin-openapi/openapi3"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	path string
	doc  openapi3.T
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crawl",
	Short: "A tool for interacting with openAPI specifications.",
	Long: `Crawl is a CLI for interacting with openAPI specifications. It can validate, crawl through
and perform requests based on a given openAPI specification.

Running the CLI in interactive mode will always first validate the spec, then proceed as an interactive prompt.
`,
	PersistentPreRun: rootPreRun,
}

func rootPreRun(*cobra.Command, []string) {
	specPrompt := &survey.Input{
		Message: "Point to a filepath of an openAPI spec:",
		Suggest: func(toComplete string) []string {
			files, _ := filepath.Glob(toComplete + "*")
			return files
		},
	}

	var err error
	// Use the cli path variable if provided
	if path == "" {
		err = survey.AskOne(specPrompt, &path)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
	d, err := spec.Validate(path)
	if err != nil {
		os.Exit(1)
	}
	doc = *d
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Available for all subcommands
	rootCmd.PersistentFlags().StringVar(&path, "path", "", "Path to an openAPI specification")
}
