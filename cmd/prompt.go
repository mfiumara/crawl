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
	Short: "Open an API spec and perform a single request",
	Long: `Opens an API spec provided and goes through it to list servers and paths.
Prompts the user to choose a server and path and performs the request`,
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

func Prompt(doc openapi3.T) error {
	err, server := promptServer(doc)
	if err != nil {
		return err
	}

	err, path := promptPath(doc)
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

func getOptionsAndPaths(doc openapi3.T) ([]string, map[string]Path) {
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
	return options, ops
}

func promptPath(doc openapi3.T) (error, Path) {
	var input string

	options, ops := getOptionsAndPaths(doc)
	methodPrompt := &survey.Select{
		Message: "Select method to call",
		Options: options,
	}
	err := survey.AskOne(methodPrompt, &input, survey.WithValidator(survey.Required))
	if err != nil {
		return err, Path{}
	}

	return nil, ops[input]
}

func promptServer(doc openapi3.T) (error, openapi3.Server) {
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

func init() {
	rootCmd.AddCommand(promptCmd)
}
