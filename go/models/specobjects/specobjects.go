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
	nameNode := &tree.Node{
		Name:      "Spec Objects",
		FontStyle: tree.ITALIC,
	}
	sliceOfSpecObjectNodes = append(sliceOfSpecObjectNodes, nameNode)

	objects := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_OBJECTS

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

		specObjectType, ok := stager.Map_id_specobjectTypes[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
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

		AddAttributeXHTMLNodes(stager, objectNode, specObject)
		AddAttributeENUMNodes(stager, objectNode, specObject)

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

func AddAttributeENUMNodes(
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
		if enumType, ok := stager.Map_id_attributeDefinitionENUM[attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF]; ok {
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
