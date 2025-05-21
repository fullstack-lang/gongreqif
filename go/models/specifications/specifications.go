package specifications

import (
	"fmt"
	"log"
	"strings"

	m "github.com/fullstack-lang/gongreqif/go/models"
	"github.com/fullstack-lang/gongreqif/go/models/specobjects"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type SpecificationsTreeStageUpdater struct {
}

func (o *SpecificationsTreeStageUpdater) UpdateAndCommitSpecificationsTreeStage(stager *m.Stager) {
	treeStage := stager.GetSpecificationsTreeStage()

	sliceOfSpecificationNodes := make([]*tree.Node, 0)
	nameNode := &tree.Node{
		Name:      "Specifications",
		FontStyle: tree.ITALIC,
	}
	sliceOfSpecificationNodes = append(sliceOfSpecificationNodes, nameNode)

	specifications := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPECIFICATIONS

	// prepare one node per specification type
	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES
	map_specificationType_node := make(map[*m.SPECIFICATION_TYPE]*tree.Node)
	map_specificationType_nbInstances := make(map[*m.SPECIFICATION_TYPE]int)
	for _, specificationType := range spectypes.SPECIFICATION_TYPE {
		nodeSpecificationType := &tree.Node{
			Name: specificationType.Name,
		}
		sliceOfSpecificationNodes = append(sliceOfSpecificationNodes, nodeSpecificationType)
		map_specificationType_node[specificationType] = nodeSpecificationType
	}

	for _, specification := range specifications.SPECIFICATION {

		specificationType, ok := stager.Map_id_specificationType[specification.TYPE.SPECIFICATION_TYPE_REF]
		if !ok {
			log.Panic("specRelation.TYPE.SPECIFICATION_TYPE_REF", specification.TYPE.SPECIFICATION_TYPE_REF,
				"unknown relation type")
		}

		relationNode := &tree.Node{
			Name: specification.Name,
		}
		map_specificationType_node[specificationType].Children =
			append(map_specificationType_node[specificationType].Children, relationNode)
		map_specificationType_nbInstances[specificationType] = map_specificationType_nbInstances[specificationType] + 1

		{

			specificationAttributeCategoryXHTML := &tree.Node{
				Name:       "XHTML",
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			relationNode.Children = append(relationNode.Children, specificationAttributeCategoryXHTML)
			if specification.VALUES != nil {
				for _, attribute := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
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
					specificationAttributeCategoryXHTML.Children = append(specificationAttributeCategoryXHTML.Children, nodeXHTMLAttribute)
				}
			}
		}
		{
			hierarchyParentNode := &tree.Node{
				Name:       "Hierarchy",
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			relationNode.Children = append(relationNode.Children, hierarchyParentNode)

			for _, specHierarchy := range specification.CHILDREN.SPEC_HIERARCHY {

				processSpecHierarchy(
					stager,
					specHierarchy,
					hierarchyParentNode)
			}
		}

	}

	// update the node with the number of instances
	for _, type_ := range spectypes.SPECIFICATION_TYPE {
		nbInstances := map_specificationType_nbInstances[type_]
		nodeSpecType := map_specificationType_node[type_]
		nodeSpecType.Name = nodeSpecType.Name + fmt.Sprintf(" (%d)", nbInstances)
	}

	tree.StageBranch(treeStage,
		&tree.Tree{
			Name:      stager.GetSpecificationsTreeName(),
			RootNodes: sliceOfSpecificationNodes,
		},
	)

	treeStage.Commit()
}

func processSpecHierarchy(
	stager *m.Stager,
	specHierarchy *m.SPEC_HIERARCHY,
	hierarchyParentNode *tree.Node) {

	specObject, ok := stager.Map_id_specObject[specHierarchy.OBJECT.SPEC_OBJECT_REF]
	if !ok {
		log.Panic("specHierarchy.OBJECT.SPEC_OBJECT_REF", specHierarchy.OBJECT.SPEC_OBJECT_REF,
			"unknown ref")
	}
	hierarchyNode := &tree.Node{
		Name: specObject.Name,
	}
	hierarchyParentNode.Children = append(hierarchyParentNode.Children, hierarchyNode)

	specobjects.AddAttributeXHTMLNodes(stager, hierarchyNode, specObject)
	specobjects.AddAttributeENUMNodes(stager, hierarchyNode, specObject)

	if specHierarchy.CHILDREN != nil {
		for _, specHierarchyChildren := range specHierarchy.CHILDREN.SPEC_HIERARCHY {
			processSpecHierarchy(
				stager,
				specHierarchyChildren,
				hierarchyNode)
		}
	}
}
