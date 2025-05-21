package specobjects

import (
	"fmt"
	"log"
	"strings"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	m "github.com/fullstack-lang/gongreqif/go/models"
)

func AddAttributeNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	AddAttributeXHTMLNodes(stager, objectNode, specObject)
	AddAttributeStringNodes(stager, objectNode, specObject)
	AddAttributeBooleanNodes(stager, objectNode, specObject)
	AddAttributeIntegerNodes(stager, objectNode, specObject)
	AddAttributeRealNodes(stager, objectNode, specObject)
	AddAttributeDateNodes(stager, objectNode, specObject)
	AddAttributeEnumNodes(stager, objectNode, specObject)
}

// AddAttributeXHTMLNodes - This is the function you provided as an example for the change.
// For completeness and to ensure the pattern is clear, here's how it would look fully:
func AddAttributeXHTMLNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_XHTML) > 0 {
		objectNodeAttributeCategoryXHTML := &tree.Node{
			Name:       "XHTML",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryXHTML)

		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_attributeDefinitionXHTML[attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_XHTML_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF,
					"unknown ref")
			}

			enclosedText := attribute.THE_VALUE.EnclosedText
			enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:div>", " ")
			enclosedText = strings.ReplaceAll(enclosedText, "</reqif-xhtml:div>", "\n")
			enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:br >", "-")
			enclosedText = strings.ReplaceAll(enclosedText, "</reqif-xhtml:br >", "\n")
			enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:br />", "\n")

			nodeXHTMLAttribute := &tree.Node{
				Name: attributeDefinition + " : " + enclosedText,
			}
			objectNodeAttributeCategoryXHTML.Children = append(objectNodeAttributeCategoryXHTML.Children, nodeXHTMLAttribute)
		}
	}
}

func AddAttributeStringNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_STRING) > 0 {
		objectNodeAttributeCategoryString := &tree.Node{
			Name:       "STRING",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryString)
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_STRING {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_attributeDefinitionString[attribute.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_STRING_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attributeDefinition + " : " + attribute.THE_VALUE,
			}
			objectNodeAttributeCategoryString.Children = append(objectNodeAttributeCategoryString.Children, nodeAttribute)
		}
	}
}

func AddAttributeBooleanNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_BOOLEAN) > 0 {
		objectNodeAttributeCategoryBoolean := &tree.Node{
			Name:       "BOOLEAN",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryBoolean)
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_attributeDefinitionBoolean[attribute.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_BOOLEAN_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF,
					"unknown ref")
			}

			value := "false"
			if attribute.THE_VALUE {
				value = "true"
			}
			nodeAttribute := &tree.Node{
				Name: attributeDefinition + " : " + value,
			}
			objectNodeAttributeCategoryBoolean.Children = append(objectNodeAttributeCategoryBoolean.Children, nodeAttribute)
		}
	}
}

func AddAttributeIntegerNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_INTEGER) > 0 {
		objectNodeAttributeCategoryInteger := &tree.Node{
			Name:       "Integer",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryInteger)
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_INTEGER {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_attributeDefinitionInteger[attribute.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_INTEGER_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attributeDefinition + " : " + fmt.Sprintf("%d", attribute.THE_VALUE),
			}
			objectNodeAttributeCategoryInteger.Children = append(objectNodeAttributeCategoryInteger.Children, nodeAttribute)
		}
	}
}

func AddAttributeDateNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_DATE) > 0 {
		objectNodeAttributeCategoryDate := &tree.Node{
			Name:       "Date",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryDate)
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_DATE {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_attributeDefinitionDate[attribute.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_DATE_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attributeDefinition + " : " + attribute.THE_VALUE,
			}
			objectNodeAttributeCategoryDate.Children = append(objectNodeAttributeCategoryDate.Children, nodeAttribute)
		}
	}
}

func AddAttributeRealNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_REAL) > 0 {
		objectNodeAttributeCategoryReal := &tree.Node{
			Name:       "Real",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryReal)
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_REAL {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_attributeDefinitionReal[attribute.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_REAL_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attributeDefinition + " : " + fmt.Sprintf("%f", attribute.THE_VALUE),
			}
			objectNodeAttributeCategoryReal.Children = append(objectNodeAttributeCategoryReal.Children, nodeAttribute)
		}
	}
}

func AddAttributeEnumNodes(
	stager *m.Stager,
	objectNode *tree.Node,
	specObject *m.SPEC_OBJECT,
) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION) > 0 {
		objectNodeAttributeCategory := &tree.Node{
			Name:       "Enums",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		objectNode.Children = append(objectNode.Children, objectNodeAttributeCategory)
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			var enumTypeString string
			if enumType, ok := stager.Map_id_attributeDefinitionEnum[attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF]; ok {
				enumTypeString = enumType.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_ENUMERATION_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF,
					"unknown ref") // Corrected "unkonwn" to "unknown"
			}

			valueIdentifier := attribute.VALUES.Name // Assuming VALUES here refers to the ENUM_VALUE_REF or similar structure
			var enumValueString string
			// The original code for attribute.VALUES.Name might need to be attribute.VALUES.ENUM_VALUE_REF or similar
			// depending on the actual structure of ATTRIBUTE_VALUE_ENUMERATION.
			// For now, assuming attribute.VALUES.Name correctly gives the identifier for Map_id_enumValues.
			// If attribute.VALUES itself is a slice of references, the logic might be attribute.VALUES[0].ENUM_VALUE_REF for single-select enums.
			// The provided snippet uses `attribute.VALUES.Name`, I'll stick to that.
			if enumValue, ok := stager.Map_id_enumValues[valueIdentifier]; ok {
				enumValueString = enumValue.Name
			} else {
				// It might be valid for an enum value to not be found if the reference is optional or cleared
				// log.Printf("Warning: ENUM_VALUE_REF %s not found in Map_id_enumValues", valueIdentifier)
				// For now, keeping it strict as per other panic behavior, though for enums, a missing value might not always be a panic.
				// If an empty string is acceptable for missing values:
				// enumValueString = "" // Or some placeholder
				// Or if it's an error:
				log.Panic("ENUM_VALUE_REF", valueIdentifier, "unknown ref in Map_id_enumValues")
			}

			// The original code named this `nodeXHTMLAttribute`, renaming for clarity if preferred, but kept for consistency here.
			nodeEnumAttribute := &tree.Node{
				Name: enumTypeString + " : " + enumValueString,
			}
			objectNodeAttributeCategory.Children = append(objectNodeAttributeCategory.Children, nodeEnumAttribute)
		}
	}
}
