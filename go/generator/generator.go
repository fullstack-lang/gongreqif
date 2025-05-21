package generator

import (
	"fmt"
	"log"
	"os"
	"strings"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

type ModelGenerator struct {
	PathToGoModelFile string
}

// GenerateModels implements m.ModelGeneratorInterface.
func (modelGenerator *ModelGenerator) GenerateModels(stager *m.Stager) {
	var sb strings.Builder

	sb.WriteString(`// Generated code, do not edit
package models

`)

	// generates the enums

	// generates the struct
	for specObjectType := range *m.GetGongstructInstancesSet[m.SPEC_OBJECT_TYPE](stager.GetStage()) {

		sb.WriteString(fmt.Sprintf(`type %s struct {
	Name string`, GenerateGoIdentifier(specObjectType.Name)))

		sb.WriteString(`	
	
	// Attributes XHTML
`)

		for _, attributeXHTML := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			sb.WriteString(fmt.Sprintf(`	%s string
`, GenerateGoIdentifier(attributeXHTML.Name)))
		}

		for _, attributeString := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			sb.WriteString(fmt.Sprintf(`	%s string
`, GenerateGoIdentifier(attributeString.Name)))
		}

		sb.WriteString(`
}`)
	}

	result := sb.String()
	err := os.WriteFile(modelGenerator.PathToGoModelFile, []byte(result), 0644)
	if err != nil {
		log.Panicln("Generate go models", err.Error())
	}
}
