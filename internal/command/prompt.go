package command

import (
	"crawl/internal/prompt"
	"github.com/spf13/cobra"
	"os"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Open an API spec and perform a single request",
	Long: `Opens an API spec provided and goes through it to list servers and paths.
Prompts the user to choose a server and path and performs the request`,
	Run: func(cmd *cobra.Command, args []string) {
		err := prompt.Prompt(doc)
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}
