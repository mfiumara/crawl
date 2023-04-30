package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Prompt(doc)
		if err != nil {
			os.Exit(1)
		}
	},
}

type Path struct {
	path string
	item openapi3.PathItem
}

func PromptServer(doc openapi3.T) (error, openapi3.Server) {
	// Set the printing format
	opLen := 0
	for _, value := range doc.Servers {
		if len(value.Description) > opLen {
			opLen = len(value.Description)
		}
	}
	format := fmt.Sprintf("%%-%ds | %%s", opLen)

	var servers []string
	serverMap := make(map[string]openapi3.Server)
	for _, value := range doc.Servers {
		choice := fmt.Sprintf(format, value.Description, value.URL)
		servers = append(servers, choice)
		serverMap[choice] = *value
	}
	serverPrompt := &survey.Select{
		Message: "Select which server to use",
		Options: servers,
	}
	var server string
	err := survey.AskOne(serverPrompt, &server, survey.WithValidator(survey.Required))
	if err != nil {
		return err, openapi3.Server{}
	}
	return nil, serverMap[server]
}

func PromptPath(doc openapi3.T) (error, Path) {
	// Set the printing format
	opLen := 0
	for key := range doc.Paths {
		if len(key) > opLen {
			opLen = len(key)
		}
	}
	format := fmt.Sprintf("%%-%ds %%-%ds | %%s", 7, opLen)

	var options []string
	ops := make(map[string]Path)
	for key, value := range doc.Paths {
		for method, operation := range value.Operations() {
			path := fmt.Sprintf(format, method, key, operation.Summary)
			options = append(options, path)

			ops[path] = Path{
				path: key,
				item: *value,
			}
		}
	}
	methodPrompt := &survey.Select{
		Message: "Select method to call",
		Options: options,
	}
	var answer string
	err := survey.AskOne(methodPrompt, &answer, survey.WithValidator(survey.Required))
	if err != nil {
		return err, Path{}
	}
	return nil, ops[answer]
}

func Prompt(doc openapi3.T) error {
	err, server := PromptServer(doc)
	if err != nil {
		return err
	}

	err, path := PromptPath(doc)
	if err != nil {
		return err
	}

	resp, err := http.Get(server.URL + path.path)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func init() {
	rootCmd.AddCommand(promptCmd)
}
