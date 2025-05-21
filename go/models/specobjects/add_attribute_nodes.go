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

func AddAttributeXHTMLNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {

	objectNodeAttributeCategoryXHTML := &tree.Node{
		Name:       "XHTML",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}
	objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryXHTML)
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
		// provide the type
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

func AddAttributeStringNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	objectNodeAttributeCategoryString := &tree.Node{
		Name:       "STRING",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}
	objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryString)
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_STRING {
		// provide the type
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

func AddAttributeBooleanNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	objectNodeAttributeCategoryBoolean := &tree.Node{
		Name:       "BOOLEAN",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}
	objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryBoolean)
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		// provide the type
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

func AddAttributeIntegerNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	objectNodeAttributeCategoryInteger := &tree.Node{
		Name:       "Integer",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}
	objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryInteger)
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_INTEGER {
		// provide the type
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

func AddAttributeDateNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	objectNodeAttributeCategoryDate := &tree.Node{
		Name:       "Date",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}
	objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryDate)
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_DATE {
		// provide the type
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

func AddAttributeRealNodes(stager *m.Stager, objectNode *tree.Node, specObject *m.SPEC_OBJECT) {
	objectNodeAttributeCategoryReal := &tree.Node{
		Name:       "Real",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}
	objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryReal)
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_REAL {
		// provide the type
		var attributeDefinition string
		if datatype, ok := stager.Map_id_attributeDefinitionReal[attribute.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF]; ok {
			attributeDefinition = datatype.LONG_NAME
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_REAL_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF,
				"unknown ref")
		}

		nodeAttribute := &tree.Node{
			Name: attributeDefinition + " : " + fmt.Sprintf("%d", attribute.THE_VALUE),
		}

		objectNodeAttributeCategoryReal.Children = append(objectNodeAttributeCategoryReal.Children, nodeAttribute)
	}
}

func AddAttributeEnumNodes(
	stager *m.Stager,
	objectNode *tree.Node,
	specObject *m.SPEC_OBJECT,
) {
	objectNodeAttributeCategory := &tree.Node{
		Name:       "Enums",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}
	objectNode.Children = append(objectNode.Children, objectNodeAttributeCategory)
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		// provide the type
		var enumTypeString string
		if enumType, ok := stager.Map_id_attributeDefinitionEnum[attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF]; ok {
			enumTypeString = enumType.LONG_NAME
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_ENUMERATION_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF,
				"unkonwn ref")
		}

		valueIdentifier := attribute.VALUES.Name
		var enumValueString string
		if enumValue, ok := stager.Map_id_enumValues[valueIdentifier]; ok {
			enumValueString = enumValue.Name
		}

		nodeXHTMLAttribute := &tree.Node{
			Name: enumTypeString + " : " + enumValueString,
		}
		objectNodeAttributeCategory.Children = append(objectNodeAttributeCategory.Children, nodeXHTMLAttribute)
	}
}
