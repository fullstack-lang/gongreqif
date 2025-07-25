// generated code - do not edit
package probe

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe      *Probe
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	probe *Probe,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.probe = probe
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.Stage,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	// log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "ALTERNATIVE_ID" {
		updateAndCommitTable[models.ALTERNATIVE_ID](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_BOOLEAN" {
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_DATE" {
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_ENUMERATION" {
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_INTEGER" {
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_REAL" {
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_STRING" {
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_XHTML" {
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_BOOLEAN" {
		updateAndCommitTable[models.ATTRIBUTE_VALUE_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_DATE" {
		updateAndCommitTable[models.ATTRIBUTE_VALUE_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_ENUMERATION" {
		updateAndCommitTable[models.ATTRIBUTE_VALUE_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_INTEGER" {
		updateAndCommitTable[models.ATTRIBUTE_VALUE_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_REAL" {
		updateAndCommitTable[models.ATTRIBUTE_VALUE_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_STRING" {
		updateAndCommitTable[models.ATTRIBUTE_VALUE_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_XHTML" {
		updateAndCommitTable[models.ATTRIBUTE_VALUE_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ALTERNATIVE_ID" {
		updateAndCommitTable[models.A_ALTERNATIVE_ID](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF" {
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_DATE_REF" {
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_DATE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF" {
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_INTEGER_REF" {
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_REAL_REF" {
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_REAL_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_STRING_REF" {
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_STRING_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_XHTML_REF" {
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_XHTML_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_BOOLEAN" {
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_DATE" {
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_ENUMERATION" {
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_INTEGER" {
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_REAL" {
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_STRING" {
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_XHTML" {
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_XHTML_1" {
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_XHTML_1](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_CHILDREN" {
		updateAndCommitTable[models.A_CHILDREN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_CORE_CONTENT" {
		updateAndCommitTable[models.A_CORE_CONTENT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPES" {
		updateAndCommitTable[models.A_DATATYPES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_BOOLEAN_REF" {
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_BOOLEAN_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_DATE_REF" {
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_DATE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_ENUMERATION_REF" {
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_ENUMERATION_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_INTEGER_REF" {
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_INTEGER_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_REAL_REF" {
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_REAL_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_STRING_REF" {
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_STRING_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_XHTML_REF" {
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_XHTML_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_EDITABLE_ATTS" {
		updateAndCommitTable[models.A_EDITABLE_ATTS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ENUM_VALUE_REF" {
		updateAndCommitTable[models.A_ENUM_VALUE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_OBJECT" {
		updateAndCommitTable[models.A_OBJECT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_PROPERTIES" {
		updateAndCommitTable[models.A_PROPERTIES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_RELATION_GROUP_TYPE_REF" {
		updateAndCommitTable[models.A_RELATION_GROUP_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SOURCE_1" {
		updateAndCommitTable[models.A_SOURCE_1](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SOURCE_SPECIFICATION_1" {
		updateAndCommitTable[models.A_SOURCE_SPECIFICATION_1](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPECIFICATIONS" {
		updateAndCommitTable[models.A_SPECIFICATIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPECIFICATION_TYPE_REF" {
		updateAndCommitTable[models.A_SPECIFICATION_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPECIFIED_VALUES" {
		updateAndCommitTable[models.A_SPECIFIED_VALUES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_ATTRIBUTES" {
		updateAndCommitTable[models.A_SPEC_ATTRIBUTES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_OBJECTS" {
		updateAndCommitTable[models.A_SPEC_OBJECTS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_OBJECT_TYPE_REF" {
		updateAndCommitTable[models.A_SPEC_OBJECT_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATIONS" {
		updateAndCommitTable[models.A_SPEC_RELATIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATION_GROUPS" {
		updateAndCommitTable[models.A_SPEC_RELATION_GROUPS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATION_REF" {
		updateAndCommitTable[models.A_SPEC_RELATION_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATION_TYPE_REF" {
		updateAndCommitTable[models.A_SPEC_RELATION_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_TYPES" {
		updateAndCommitTable[models.A_SPEC_TYPES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_THE_HEADER" {
		updateAndCommitTable[models.A_THE_HEADER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_TOOL_EXTENSIONS" {
		updateAndCommitTable[models.A_TOOL_EXTENSIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_BOOLEAN" {
		updateAndCommitTable[models.DATATYPE_DEFINITION_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_DATE" {
		updateAndCommitTable[models.DATATYPE_DEFINITION_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_ENUMERATION" {
		updateAndCommitTable[models.DATATYPE_DEFINITION_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_INTEGER" {
		updateAndCommitTable[models.DATATYPE_DEFINITION_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_REAL" {
		updateAndCommitTable[models.DATATYPE_DEFINITION_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_STRING" {
		updateAndCommitTable[models.DATATYPE_DEFINITION_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_XHTML" {
		updateAndCommitTable[models.DATATYPE_DEFINITION_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "EMBEDDED_VALUE" {
		updateAndCommitTable[models.EMBEDDED_VALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ENUM_VALUE" {
		updateAndCommitTable[models.ENUM_VALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RELATION_GROUP" {
		updateAndCommitTable[models.RELATION_GROUP](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RELATION_GROUP_TYPE" {
		updateAndCommitTable[models.RELATION_GROUP_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF" {
		updateAndCommitTable[models.REQ_IF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_CONTENT" {
		updateAndCommitTable[models.REQ_IF_CONTENT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_HEADER" {
		updateAndCommitTable[models.REQ_IF_HEADER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_TOOL_EXTENSION" {
		updateAndCommitTable[models.REQ_IF_TOOL_EXTENSION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATION" {
		updateAndCommitTable[models.SPECIFICATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATION_TYPE" {
		updateAndCommitTable[models.SPECIFICATION_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_HIERARCHY" {
		updateAndCommitTable[models.SPEC_HIERARCHY](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_OBJECT" {
		updateAndCommitTable[models.SPEC_OBJECT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_OBJECT_TYPE" {
		updateAndCommitTable[models.SPEC_OBJECT_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_RELATION" {
		updateAndCommitTable[models.SPEC_RELATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_RELATION_TYPE" {
		updateAndCommitTable[models.SPEC_RELATION_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "StaticWebSite" {
		updateAndCommitTable[models.StaticWebSite](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "StaticWebSiteChapter" {
		updateAndCommitTable[models.StaticWebSiteChapter](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "StaticWebSiteGeneratedImage" {
		updateAndCommitTable[models.StaticWebSiteGeneratedImage](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "StaticWebSiteImage" {
		updateAndCommitTable[models.StaticWebSiteImage](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "StaticWebSiteParagraph" {
		updateAndCommitTable[models.StaticWebSiteParagraph](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "XHTML_CONTENT" {
		updateAndCommitTable[models.XHTML_CONTENT](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
