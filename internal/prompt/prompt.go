package prompt

import (
	"crawl/internal/spec"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/getkin/kin-openapi/openapi3"
	"io"
	"net/http"
)

func Prompt(doc openapi3.T) error {
	err, server := promptServer(doc)
	if err != nil {
		return err
	}
	err, p := promptPath(doc)
	if err != nil {
		return err
	}

	resp, err := http.Get(server.URL + p.Path)
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

func promptPath(doc openapi3.T) (error, spec.Path) {
	var input string

	options, paths := spec.GetOptionsAndPaths(doc)
	methodPrompt := &survey.Select{
		Message: "Select method to call",
		Options: options,
	}
	err := survey.AskOne(methodPrompt, &input, survey.WithValidator(survey.Required))
	if err != nil {
		return err, spec.Path{}
	}

	return nil, paths[input]
}

func promptServer(doc openapi3.T) (error, openapi3.Server) {
	options, servers := spec.GetOptionsAndServers(doc)
	serverPrompt := &survey.Select{
		Message: "Select which server to use",
		Options: options,
	}
	var input string
	err := survey.AskOne(
		serverPrompt,
		&input,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return err, openapi3.Server{}
	}
	return nil, servers[input]
}
