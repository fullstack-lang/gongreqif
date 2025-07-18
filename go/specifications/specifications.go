package specifications

import (
	"fmt"
	"log"
	"strings"

	// Corrected path
	m "github.com/fullstack-lang/gongreqif/go/models"

	"github.com/fullstack-lang/gongreqif/go/specobjects"

	markdown "github.com/fullstack-lang/gong/lib/markdown/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type SpecificationsTreeStageUpdater struct {
}

// UpdateAndCommitSpecificationsMarkdownStage implements models.SpecificationsTreeUpdaterInterface.
func (o *SpecificationsTreeStageUpdater) UpdateAndCommitSpecificationsMarkdownStage(stager *m.Stager) {

	markdownStage := stager.GetMarkdownStage()
	markdownStage.Reset()

	specifications := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPECIFICATIONS.SPECIFICATION
	for _, specification := range specifications {

		if stager.GetSelectedSpecification() == nil {
			continue
		}

		if stager.GetSelectedSpecification() != specification {
			continue
		}

		// --- updated logic to generate and assign markdown content ---

		// 1. Initialize markdown content string
		markDownContent := "# " + specification.Name + "\n\n"

		// 2. A dummy parent node is created because processSpecHierarchy expects a parent
		// to append children to. This node is temporary and will be discarded.
		hierarchyParentNode := &tree.Node{}
		depth := 2 // initial depth for chapters

		// 3. Recursively process spec hierarchies to build the markdown string
		for _, specHierarchy := range specification.CHILDREN.SPEC_HIERARCHY {
			processSpecHierarchy(
				stager,
				specHierarchy,
				hierarchyParentNode,
				depth,
				&markDownContent)
		}
		// --- end of update ---

		content := &markdown.Content{
			Name:    specification.GetName(),
			Content: markDownContent, // Assign the generated markdown
		}

		markdown.StageBranch(markdownStage, content)
	}

	markdownStage.Commit()
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
			Name:       specificationType.Name,
			IsExpanded: true,
		}
		sliceOfSpecificationNodes = append(sliceOfSpecificationNodes, nodeSpecificationType)
		map_specificationType_node[specificationType] = nodeSpecificationType
	}

	for _, specification := range specifications.SPECIFICATION {

		specificationType, ok := stager.Map_id_SPECIFICATION_TYPE[specification.TYPE.SPECIFICATION_TYPE_REF]
		if !ok {
			log.Panic("specRelation.TYPE.SPECIFICATION_TYPE_REF", specification.TYPE.SPECIFICATION_TYPE_REF,
				"unknown relation type")
		}

		isSelectedSpecification := stager.GetSelectedSpecification() == specification

		specificationNode := &tree.Node{
			Name:              specification.Name,
			HasCheckboxButton: true,
			IsChecked:         isSelectedSpecification,
			IsExpanded:        stager.Map_SPECIFICATION_Nodes_expanded[specification],
			Impl: &ProxySpecification{
				stager:        stager,
				specification: specification,
			},
		}

		markDownContent := "# " + specification.Name
		map_specificationType_node[specificationType].Children =
			append(map_specificationType_node[specificationType].Children, specificationNode)
		map_specificationType_nbInstances[specificationType] = map_specificationType_nbInstances[specificationType] + 1

		{

			// specificationAttributeCategoryXHTML := &tree.Node{
			// 	Name:       "XHTML",
			// 	IsExpanded: true,
			// 	FontStyle:  tree.ITALIC,
			// }
			// specificationNode.Children = append(specificationNode.Children, specificationAttributeCategoryXHTML)
			if specification.VALUES != nil {
				for _, attribute := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
					// provide the type
					var attributeDefinition string
					if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF]; ok {
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
					specificationNode.Children = append(specificationNode.Children, nodeXHTMLAttribute)
				}
			}
		}
		{
			hierarchyParentNode := &tree.Node{
				Name:       "Hierarchy",
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			specificationNode.Children = append(specificationNode.Children, hierarchyParentNode)

			depth := 2 // depth of chapters

			for _, specHierarchy := range specification.CHILDREN.SPEC_HIERARCHY {

				processSpecHierarchy(
					stager,
					specHierarchy,
					hierarchyParentNode,
					depth,
					&markDownContent)
			}

			// log.Println(markDownContent)
		}

		specobjects.AddAttributeNodes(stager, specificationNode, specification)

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
	hierarchyParentNode *tree.Node,
	depth int,
	markDownContent *string) {

	specObject, ok := stager.Map_id_SPEC_OBJECT[specHierarchy.OBJECT.SPEC_OBJECT_REF]
	if !ok {
		log.Panic("specHierarchy.OBJECT.SPEC_OBJECT_REF", specHierarchy.OBJECT.SPEC_OBJECT_REF,
			"unknown ref")
	}

	specObjectType, ok := stager.Map_id_SPEC_OBJECT_TYPE[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
	if !ok {
		log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
			"unknown ref")
	}

	if specHierarchy.CHILDREN == nil || len(specHierarchy.CHILDREN.SPEC_HIERARCHY) == 0 {
		*markDownContent += "#"
	}
	if true {
		for range depth {
			*markDownContent += "#"
		}
		*markDownContent += " "
	}

	if stager.Map_SPEC_OBJECT_TYPE_showIdentifier[specObjectType] {
		*markDownContent += fmt.Sprintf("%s - %s\n\n", specObject.IDENTIFIER, specObject.Name)
	} else {
		*markDownContent += fmt.Sprintf("%s\n\n", specObject.Name)
	}

	specObjectNode := &tree.Node{
		Name: specObject.Name + " : " + specObjectType.Name,
	}
	hierarchyParentNode.Children = append(hierarchyParentNode.Children, specObjectNode)

	specobjects.AddAttributeNodes(stager, specObjectNode, specObject)
	specobjects.AppendAttributesToMarkdown(stager, specObject, markDownContent)

	m.AddIconForEditabilityOfAttribute(specHierarchy.IS_EDITABLE, specObject.Name, specObjectNode)

	if specHierarchy.CHILDREN != nil {
		for _, specHierarchyChildren := range specHierarchy.CHILDREN.SPEC_HIERARCHY {
			processSpecHierarchy(
				stager,
				specHierarchyChildren,
				specObjectNode,
				depth+1,
				markDownContent)
		}
	}
}

type ProxySpecification struct {
	stager        *m.Stager
	specification *m.SPECIFICATION
}

func (proxy *ProxySpecification) OnAfterUpdate(treeStage *tree.Stage, stageNode, frontNode *tree.Node) {

	if frontNode.IsChecked && !stageNode.IsChecked {
		frontNode.IsChecked = stageNode.IsChecked

		// log.Println("Specification", proxy.specification.Name, "selected")
		proxy.stager.SetSelectedSpecification(proxy.specification)

		proxy.stager.Map_SPECIFICATION_TYPE_Spec_nbInstance = initializeNbInstanceMap[m.SPECIFICATION_TYPE]()
		proxy.stager.Map_SPEC_RELATION_TYPE_Spec_nbInstance = initializeNbInstanceMap[m.SPEC_RELATION_TYPE]()

		proxy.stager.GetSpecificationsTreeUpdater().UpdateAttributeDefinitionNb(proxy.stager)
		proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
		proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsTreeStage(proxy.stager)
		proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
	}

	if !frontNode.IsChecked && stageNode.IsChecked {
		frontNode.IsChecked = stageNode.IsChecked
	}

	if frontNode.IsExpanded && !stageNode.IsExpanded {
		proxy.stager.Map_SPECIFICATION_Nodes_expanded[proxy.specification] = true
	}

	if !frontNode.IsExpanded && stageNode.IsExpanded {
		proxy.stager.Map_SPECIFICATION_Nodes_expanded[proxy.specification] = false
	}
}
