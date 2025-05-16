package specobjects

import (
	"fmt"
	"log"
	"strings"

	m "github.com/fullstack-lang/gongreqif/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type SpecObjectsTreeStageUpdater struct {
}

func (o *SpecObjectsTreeStageUpdater) UpdateAndCommitSpecObjectsTreeStage(stager *m.Stager) {
	treeStage := stager.GetSpecObjectsTreeStage()

	treeStage.Reset()

	sliceOfSpecObjectNodes := make([]*tree.Node, 0)

	objects := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_OBJECTS

	map_id_specobjectTypes := make(map[string]*m.SPEC_OBJECT_TYPE)
	{
		specObjectTypes := *m.GetGongstructInstancesSet[m.SPEC_OBJECT_TYPE](stager.GetStage())
		for specObjectType := range specObjectTypes {
			map_id_specobjectTypes[specObjectType.IDENTIFIER] = specObjectType
		}
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

	// prepare one node per spec object type
	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES
	map_specObjectType_node := make(map[*m.SPEC_OBJECT_TYPE]*tree.Node)
	map_specObjectType_nbInstances := make(map[*m.SPEC_OBJECT_TYPE]int)
	for _, specObjectType := range spectypes.SPEC_OBJECT_TYPE {
		nodeSpecObjectType := &tree.Node{
			Name: specObjectType.Name,
		}
		sliceOfSpecObjectNodes = append(sliceOfSpecObjectNodes, nodeSpecObjectType)
		map_specObjectType_node[specObjectType] = nodeSpecObjectType
	}

	for _, specObject := range objects.SPEC_OBJECT {

		specObjectType, ok := map_id_specobjectTypes[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
		if !ok {
			log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
				"unknown object type")
		}

		objectNode := &tree.Node{
			Name: specObject.Name,
		}
		map_specObjectType_node[specObjectType].Children =
			append(map_specObjectType_node[specObjectType].Children, objectNode)
		map_specObjectType_nbInstances[specObjectType] = map_specObjectType_nbInstances[specObjectType] + 1

		{
			objectNodeAttributeCategoryXHTML := &tree.Node{
				Name:       "XHTML",
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			objectNode.Children = append(objectNode.Children, objectNodeAttributeCategoryXHTML)
			for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
				// provide the type
				var attributeDefinition string
				if datatype, ok := attributeDefinitionXHTML_id_map[attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF]; ok {
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
		nbInstances := map_specObjectType_nbInstances[specObjectType]
		nodeSpecType := map_specObjectType_node[specObjectType]
		nodeSpecType.Name = nodeSpecType.Name + fmt.Sprintf(" (%d)", nbInstances)
	}

	tree.StageBranch(treeStage,
		&tree.Tree{
			Name:      stager.GetSpecObjectsTreeName(),
			RootNodes: sliceOfSpecObjectNodes,
		},
	)

	treeStage.Commit()
}
