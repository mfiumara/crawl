/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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
		err := Prompt()
		if err != nil {
			os.Exit(1)
		}
	},
}

type Path struct {
	path string
	item openapi3.PathItem
}

func Prompt() error {

	// Set the printing format
	opLen := 0
	for _, value := range doc.Servers {
		if len(value.Description) > opLen {
			opLen = len(value.Description)
		}
	}
	format := fmt.Sprintf("%%-%ds | %%s", opLen)

	var servers []string
	svrs := make(map[string]openapi3.Server)
	for _, value := range doc.Servers {
		choice := fmt.Sprintf(format, value.Description, value.URL)
		servers = append(servers, choice)
		svrs[choice] = *value
	}
	serverPrompt := &survey.Select{
		Message: "Select which server to use",
		Options: servers,
	}
	var server string
	err := survey.AskOne(serverPrompt, &server, survey.WithValidator(survey.Required))
	if err != nil {
		println(err.Error())
		return err
	}
	baseUrl := svrs[server].URL
	// Set the printing format
	opLen = 0
	for key, _ := range doc.Paths {
		if len(key) > opLen {
			opLen = len(key)
		}
	}
	format = fmt.Sprintf("%%-%ds %%-%ds | %%s", 7, opLen)

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
	err = survey.AskOne(methodPrompt, &answer, survey.WithValidator(survey.Required))
	if err != nil {
		println(err.Error())
		return err
	}
	resp, err := http.Get(baseUrl + ops[answer].path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
	return nil
}

func init() {
	rootCmd.AddCommand(promptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// promptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// promptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
