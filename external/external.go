package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

type dataValidator struct {
	data string
	err  error
}

func main() {

	schemaLoader := gojsonschema.NewReferenceLoader("file://schema/Main.json")
	documentLoader := gojsonschema.NewReferenceLoader("file://data/valid.json")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
