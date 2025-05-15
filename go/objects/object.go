package objects

import (
	"log"
	"strings"

	m "github.com/fullstack-lang/gongreqif/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ObjectTreeStageUpdater struct {
}

// UpdateAndCommitObjectTreeStage implements models.ObjectTreeStageInterface.
func (o *ObjectTreeStageUpdater) UpdateAndCommitObjectTreeStage(stager *m.Stager) {
	treeStage := stager.GetObjectTreeStage()

	treeStage.Reset()

	rootNode := &tree.Node{
		Name:       "Objects",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}

	objects := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_OBJECTS

	specObjectTypes := *m.GetGongstructInstancesSet[m.SPEC_OBJECT_TYPE](stager.GetStage())
	specobjectTypes_id_map := make(map[string]*m.SPEC_OBJECT_TYPE)
	for specObjectType := range specObjectTypes {
		specobjectTypes_id_map[specObjectType.IDENTIFIER] = specObjectType
	}

	attributeDefinitionXHTML_id_map := make(map[string]*m.ATTRIBUTE_DEFINITION_XHTML)
	{
		attributeDefinitionXHTMLs := *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_XHTML](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			attributeDefinitionXHTML_id_map[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	attributeDefinitionENUM_id_map := make(map[string]*m.ATTRIBUTE_DEFINITION_ENUMERATION)
	{
		attributeDefinitionENUMs := *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_ENUMERATION](stager.GetStage())
		for attributeDefinition := range attributeDefinitionENUMs {
			attributeDefinitionENUM_id_map[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	enumValuess := *m.GetGongstructInstancesSet[m.ENUM_VALUE](stager.GetStage())
	enumValues_id_map := make(map[string]*m.ENUM_VALUE)
	for enumValue := range enumValuess {
		enumValues_id_map[enumValue.IDENTIFIER] = enumValue
	}

	for _, specObject := range objects.SPEC_OBJECT {

		// try to match the type
		typeId_ := specObject.TYPE.SPEC_OBJECT_TYPE_REF
		specObjectType := specobjectTypes_id_map[typeId_]

		objectNode := &tree.Node{
			Name: specObject.Name + " : " + specObjectType.Name,
		}
		rootNode.Children = append(rootNode.Children, objectNode)

		{
			objectNodeAttributeCategoryXHTML := &tree.Node{
				Name:       "XHTML",
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryXHTML)
			for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
				// provide the type
				var attributeType string
				if datatype, ok := attributeDefinitionXHTML_id_map[attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF]; ok {
					attributeType = datatype.LONG_NAME
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
					Name: attributeType + " : " + enclosedText,
				}
				objectNodeAttributeCategoryXHTML.Children = append(objectNodeAttributeCategoryXHTML.Children, nodeXHTMLAttribute)
			}
		}
		{
			objectNodeAttributeCategory := &tree.Node{
				Name:       "Enums",
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			objectNode.Children = append(objectNode.Children, objectNodeAttributeCategory)
			for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
				// provide the type
				var enumTypeString string
				if enumType, ok := attributeDefinitionENUM_id_map[attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF]; ok {
					enumTypeString = enumType.LONG_NAME
				} else {
					log.Panic("ATTRIBUTE_DEFINITION_ENUMERATION_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF,
						"unkonwn ref")
				}

				valueIdentifier := attribute.VALUES.Name
				var enumValueString string
				if enumValue, ok := enumValues_id_map[valueIdentifier]; ok {
					enumValueString = enumValue.Name
				}

				nodeXHTMLAttribute := &tree.Node{
					Name: enumTypeString + " : " + enumValueString,
				}
				objectNodeAttributeCategory.Children = append(objectNodeAttributeCategory.Children, nodeXHTMLAttribute)
			}
		}
	}

	tree.StageBranch(treeStage,
		&tree.Tree{
			Name: stager.GetObjectTreeName(),
			RootNodes: []*tree.Node{
				rootNode,
			},
		},
	)

	treeStage.Commit()
}
