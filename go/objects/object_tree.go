package objects

import (
	"strings"

	m "github.com/fullstack-lang/gongreqif/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ObjectTreeStage struct {
}

// UpdateAndCommitObjectTreeStage implements models.ObjectTreeStageInterface.
func (o *ObjectTreeStage) UpdateAndCommitObjectTreeStage(stager *m.Stager) {
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

	attributeDefinitionXHTML := *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_XHTML](stager.GetStage())
	attributeDefinitionXHTML_id_map := make(map[string]*m.ATTRIBUTE_DEFINITION_XHTML)
	for attributeDefinition := range attributeDefinitionXHTML {
		attributeDefinitionXHTML_id_map[attributeDefinition.IDENTIFIER] = attributeDefinition
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
