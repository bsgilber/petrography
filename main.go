package main

import (
  "os"
  "fmt"

  "github.com/getkin/kin-openapi/openapi3"
  "gopkg.in/alecthomas/kingpin.v2"
)

func validateSchemaFile(path string) {
	loader := openapi3.NewSwaggerLoader()
	swagger, _ := loader.LoadSwaggerFromFile(path)
	err := swagger.Validate(loader.Context)

	if err != nil {
		fmt.Printf("File located at path [%s] is [INVALID]\n  -  %s", path, err.Error())
	} else {
		fmt.Printf("File located at path [%s] is [VALID]", path)
	}

	return
}

func validateManifestFile(path string) {
	println("not implemented yet")
}

var (
  app      			= kingpin.New("petro", "A command-line application for inhouse file validations.")
  validate      	= app.Command("validate", "Validate:\n  *    OpenApi Schema\n  *    Manifest File") 
  
  schema			= validate.Command("schema", "The Schema to validate against OpenAPI standards.")
  schemaFile	  	= schema.Arg("schema-file", "Schema file to validate.").Required().String()

  manifest			= validate.Command("manifest", "The Manifest file to validate against fetchetta-stone standards.")
  manifestFile	  	= manifest.Arg("manifest-file", "The Manifest file to validate.").Required().String()
)

func main() {
  switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case schema.FullCommand():
		validateSchemaFile(*schemaFile)
	case manifest.FullCommand():
		validateManifestFile(*manifestFile)
  }
}