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
	datatypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.DATATYPES
	for _, datatype := range datatypes.DATATYPE_DEFINITION_ENUMERATION {

		sb.WriteString(fmt.Sprintf(`type %s string`, GenerateGoIdentifier(datatype.Name)))
		sb.WriteString("\n")
		sb.WriteString("\n")

		sb.WriteString("const (\n")

		for _, enum := range datatype.SPECIFIED_VALUES.ENUM_VALUE {

			sb.WriteString(fmt.Sprintf("\t%s_%s %s = \"%s\"\n",
				GenerateGoIdentifier(datatype.Name),
				GenerateGoIdentifier(enum.LONG_NAME),
				GenerateGoIdentifier(datatype.Name),
				enum.LONG_NAME))

		}

		sb.WriteString("\n)\n")
	}

	// generates the struct
	for specObjectType := range *m.GetGongstructInstancesSet[m.SPEC_OBJECT_TYPE](stager.GetStage()) {

		sb.WriteString(fmt.Sprintf("type %s struct {", GenerateGoIdentifier(specObjectType.Name)))
		sb.WriteString("\tName string\n")
		sb.WriteString("\n")

		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			sb.WriteString(fmt.Sprintf("\t%s string\n", GenerateGoIdentifier(attribute.Name)))
		}

		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {

			var attributeType string
			if datatype, ok := stager.Map_id_datatypes_enumeration[attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF", attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF,
					"unknown ref")
			}

			sb.WriteString(fmt.Sprintf("\t%s %s\n",
				GenerateGoIdentifier(attribute.Name),
				GenerateGoIdentifier(attributeType),
			))
		}

		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			sb.WriteString(fmt.Sprintf("\t%s bool\n", GenerateGoIdentifier(attribute.Name)))
		}
		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			sb.WriteString(fmt.Sprintf("\t%s string\n", GenerateGoIdentifier(attribute.Name)))
		}
		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			sb.WriteString(fmt.Sprintf("\t%s int\n", GenerateGoIdentifier(attribute.Name)))
		}
		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			sb.WriteString(fmt.Sprintf("\t%s float64\n", GenerateGoIdentifier(attribute.Name)))
		}
		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			sb.WriteString(fmt.Sprintf("\t%s string\n", GenerateGoIdentifier(attribute.Name)))
		}

		sb.WriteString("\n}\n")
	}

	result := sb.String()
	err := os.WriteFile(modelGenerator.PathToGoModelFile, []byte(result), 0644)
	if err != nil {
		log.Panicln("Generate go models", err.Error())
	}
}
