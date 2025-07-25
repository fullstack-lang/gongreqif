package specifications

import (
	"fmt"
	"log"
	"strings"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	m "github.com/fullstack-lang/gongreqif/go/models"
	"github.com/fullstack-lang/gongreqif/go/specobjects"
)

func processSpecHierarchy(
	stager *m.Stager,
	specHierarchy *m.SPEC_HIERARCHY,
	hierarchyParentNode *tree.Node,
	outerDepth int,
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

	// starting makr bold

	markdownBoldStartingMark := `
**`
	markdownBoldEndingMark := `**
`

	if stager.Map_SPEC_OBJECT_TYPE_isHeading[specObjectType] {

		*markDownContent += `
`

		if specHierarchy.CHILDREN == nil || len(specHierarchy.CHILDREN.SPEC_HIERARCHY) == 0 {
			*markDownContent += "#"
		}

		if true {
			for range outerDepth {
				*markDownContent += "#"
			}
			*markDownContent += " "
		}
	} else {
		*markDownContent += markdownBoldStartingMark
	}

	if stager.Map_SPEC_OBJECT_TYPE_showIdentifier[specObjectType] {
		*markDownContent += fmt.Sprintf("%s", specObject.IDENTIFIER)
	}

	if stager.Map_SPEC_OBJECT_TYPE_showIdentifier[specObjectType] &&
		stager.Map_SPEC_OBJECT_TYPE_showName[specObjectType] {
		*markDownContent += " - "
	}

	if stager.Map_SPEC_OBJECT_TYPE_showName[specObjectType] {
		*markDownContent += fmt.Sprintf("%s", specObject.Name)
	}

	titleComplement := fillUpStringWithAttributes(stager, specObject, Title)

	subjectComplement := fillUpStringWithAttributes(stager, specObject, Subject)

	if !stager.Map_SPEC_OBJECT_TYPE_isHeading[specObjectType] && !strings.HasSuffix(*markDownContent, "**") && titleComplement != "" {
		*markDownContent += " - "
	}

	*markDownContent += titleComplement

	if stager.Map_SPEC_OBJECT_TYPE_isHeading[specObjectType] {
		*markDownContent += `
`
	} else {
		// ending mark for bold
		*markDownContent += markdownBoldEndingMark
	}

	// remove "****" if no title is present
	if strings.HasSuffix(*markDownContent, markdownBoldStartingMark+markdownBoldEndingMark) {
		*markDownContent = strings.TrimSuffix(*markDownContent, markdownBoldStartingMark+markdownBoldEndingMark)
	}

	// add the subject after the title
	if subjectComplement != "" {
		*markDownContent += `
` + subjectComplement + `
`
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
				outerDepth+1,
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
