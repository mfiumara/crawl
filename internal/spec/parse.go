package spec

import (
	"context"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
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
func Validate(path string) (*openapi3.T, error) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc, err := loader.LoadFromFile(path)
	if err != nil {
		return nil, err
	}

	// Validate document
	err = doc.Validate(ctx)
	if err != nil {
		return nil, err
	}

	return doc, nil
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
