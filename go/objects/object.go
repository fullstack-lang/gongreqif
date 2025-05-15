package objects

import (
	"fmt"
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

	sliceOfSpecObjectNodes := make([]*tree.Node, 0)

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

	// prepare one node per spec obect type
	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES
	map_specType_node := make(map[*m.SPEC_OBJECT_TYPE]*tree.Node)
	map_specType_nbInstances := make(map[*m.SPEC_OBJECT_TYPE]int)
	for _, specObjectType := range spectypes.SPEC_OBJECT_TYPE {
		nodeSpecType := &tree.Node{
			Name: specObjectType.Name,
		}
		sliceOfSpecObjectNodes = append(sliceOfSpecObjectNodes, nodeSpecType)
		map_specType_node[specObjectType] = nodeSpecType

	}

	for _, specObject := range objects.SPEC_OBJECT {

		// try to match the type
		typeId_ := specObject.TYPE.SPEC_OBJECT_TYPE_REF
		specObjectType := specobjectTypes_id_map[typeId_]

		objectNode := &tree.Node{
			Name: specObject.Name + " : " + specObjectType.Name,
		}
		map_specType_node[specObjectType].Children =
			append(map_specType_node[specObjectType].Children, objectNode)
		map_specType_nbInstances[specObjectType] = map_specType_nbInstances[specObjectType] + 1

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

	// update the node with the number of instances
	for _, specObjectType := range spectypes.SPEC_OBJECT_TYPE {
		nbInstances := map_specType_nbInstances[specObjectType]
		nodeSpecType := map_specType_node[specObjectType]
		nodeSpecType.Name = nodeSpecType.Name + fmt.Sprintf(" (%d)", nbInstances)
	}

	tree.StageBranch(treeStage,
		&tree.Tree{
			Name:      stager.GetObjectTreeName(),
			RootNodes: sliceOfSpecObjectNodes,
		},
	)

	treeStage.Commit()
}
