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
		valid, err := spec.IsValid(doc)
		if err != nil || !valid {
			println("Spec not valid: ", err.Error())
			os.Exit(1)
		}
		println("✅ Spec valid")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
