package spec

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/getkin/kin-openapi/openapi3"
	"net/http"
)

type Path struct {
	Path string
	Item openapi3.PathItem
}

func GetOptionsAndPaths(doc openapi3.T) ([]string, map[string]Path) {
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
				Path: key,
				Item: *value,
			}
		}
	}
	return options, ops
}

func GetOptionsAndServers(doc openapi3.T) ([]string, map[string]openapi3.Server) {
	// Set the printing format
	opLen := 0
	for _, value := range doc.Servers {
		if len(value.Description) > opLen {
			opLen = len(value.Description)
		}
	}
	format := fmt.Sprintf("%%-%ds | %%s", opLen)

	var options []string
	serverMap := make(map[string]openapi3.Server)
	for _, value := range doc.Servers {
		choice := fmt.Sprintf(format, value.Description, value.URL)
		options = append(options, choice)
		serverMap[choice] = *value
	}
	return options, serverMap
}

// Validate Validates an openapi spec
func IsValid(doc openapi3.T) (bool, error) {
	ctx := context.Background()

	// Validate document
	err := doc.Validate(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ListMethods(doc openapi3.T) error {
	// Set the printing format
	opLen := 0
	for key := range doc.Paths {
		if len(key) > opLen {
			opLen = len(key)
		}
	}
	format := fmt.Sprintf("%%-%ds %%-%ds | %%s\n", 12, opLen)

	for key, value := range doc.Paths {
		for method, operation := range value.Operations() {
			fmt.Printf(format, methodColorized(method), key, operation.Summary)
		}
	}
	return nil
}

func methodColorized(method string) string {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	hiGreen := color.New(color.FgHiGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	switch method {
	case http.MethodGet:
		return blue(method)
	case http.MethodPut:
		return yellow(method)
	case http.MethodPost:
		return green(method)
	case http.MethodPatch:
		return hiGreen(method)
	case http.MethodDelete:
		return red(method)

	case http.MethodConnect:
	case http.MethodHead:
	case http.MethodOptions:
	case http.MethodTrace:
	default:
		return method
	}
	return method
}

func ListSecurity(doc openapi3.T) error {
	for key, value := range doc.Components.SecuritySchemes {
		fmt.Printf("name: %s | type: %s\n", key, value.Value.Type)
	}
	return nil
}
