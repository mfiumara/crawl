package command

import (
	"crawl/internal/spec"
	"os"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate an openAPI spec definition",
	Long:  `Validates whether the openAPI specification is valid or not.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := spec.Validate(path)
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
