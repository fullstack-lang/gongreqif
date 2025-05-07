package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
)

const SideBarTreeName = "gong"

func updateAndCommitTree(
	probe *Probe,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range probe.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	probe.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: SideBarTreeName}).Stage(probe.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](probe.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&tree.Node{Name: name}).Stage(probe.treeStage)

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "ALTERNATIVE_ID":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ALTERNATIVE_ID](probe.stageOfInterest)
			for _alternative_id := range set {
				nodeInstance := (&tree.Node{Name: _alternative_id.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_alternative_id, "ALTERNATIVE_ID", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_BOOLEAN](probe.stageOfInterest)
			for _attribute_definition_boolean := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_boolean.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_boolean, "ATTRIBUTE_DEFINITION_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_DATE](probe.stageOfInterest)
			for _attribute_definition_date := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_date.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_date, "ATTRIBUTE_DEFINITION_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_ENUMERATION](probe.stageOfInterest)
			for _attribute_definition_enumeration := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_enumeration.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_enumeration, "ATTRIBUTE_DEFINITION_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_INTEGER](probe.stageOfInterest)
			for _attribute_definition_integer := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_integer.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_integer, "ATTRIBUTE_DEFINITION_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_REAL](probe.stageOfInterest)
			for _attribute_definition_real := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_real.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_real, "ATTRIBUTE_DEFINITION_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_STRING](probe.stageOfInterest)
			for _attribute_definition_string := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_string.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_string, "ATTRIBUTE_DEFINITION_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_XHTML](probe.stageOfInterest)
			for _attribute_definition_xhtml := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_xhtml.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_xhtml, "ATTRIBUTE_DEFINITION_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_BOOLEAN](probe.stageOfInterest)
			for _attribute_value_boolean := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_boolean.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_boolean, "ATTRIBUTE_VALUE_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_DATE](probe.stageOfInterest)
			for _attribute_value_date := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_date.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_date, "ATTRIBUTE_VALUE_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_ENUMERATION](probe.stageOfInterest)
			for _attribute_value_enumeration := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_enumeration.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_enumeration, "ATTRIBUTE_VALUE_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_INTEGER](probe.stageOfInterest)
			for _attribute_value_integer := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_integer.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_integer, "ATTRIBUTE_VALUE_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_REAL](probe.stageOfInterest)
			for _attribute_value_real := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_real.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_real, "ATTRIBUTE_VALUE_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_STRING](probe.stageOfInterest)
			for _attribute_value_string := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_string.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_string, "ATTRIBUTE_VALUE_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_XHTML](probe.stageOfInterest)
			for _attribute_value_xhtml := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_xhtml.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_xhtml, "ATTRIBUTE_VALUE_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "AnyType":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.AnyType](probe.stageOfInterest)
			for _anytype := range set {
				nodeInstance := (&tree.Node{Name: _anytype.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_anytype, "AnyType", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_BOOLEAN](probe.stageOfInterest)
			for _datatype_definition_boolean := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_boolean.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_boolean, "DATATYPE_DEFINITION_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_DATE](probe.stageOfInterest)
			for _datatype_definition_date := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_date.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_date, "DATATYPE_DEFINITION_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_ENUMERATION](probe.stageOfInterest)
			for _datatype_definition_enumeration := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_enumeration.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_enumeration, "DATATYPE_DEFINITION_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_INTEGER](probe.stageOfInterest)
			for _datatype_definition_integer := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_integer.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_integer, "DATATYPE_DEFINITION_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_REAL](probe.stageOfInterest)
			for _datatype_definition_real := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_real.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_real, "DATATYPE_DEFINITION_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_STRING](probe.stageOfInterest)
			for _datatype_definition_string := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_string.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_string, "DATATYPE_DEFINITION_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_XHTML](probe.stageOfInterest)
			for _datatype_definition_xhtml := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_xhtml.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_xhtml, "DATATYPE_DEFINITION_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "EMBEDDED_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.EMBEDDED_VALUE](probe.stageOfInterest)
			for _embedded_value := range set {
				nodeInstance := (&tree.Node{Name: _embedded_value.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_embedded_value, "EMBEDDED_VALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ENUM_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ENUM_VALUE](probe.stageOfInterest)
			for _enum_value := range set {
				nodeInstance := (&tree.Node{Name: _enum_value.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_enum_value, "ENUM_VALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RELATION_GROUP":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RELATION_GROUP](probe.stageOfInterest)
			for _relation_group := range set {
				nodeInstance := (&tree.Node{Name: _relation_group.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_relation_group, "RELATION_GROUP", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RELATION_GROUP_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](probe.stageOfInterest)
			for _relation_group_type := range set {
				nodeInstance := (&tree.Node{Name: _relation_group_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_relation_group_type, "RELATION_GROUP_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF](probe.stageOfInterest)
			for _req_if := range set {
				nodeInstance := (&tree.Node{Name: _req_if.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if, "REQ_IF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](probe.stageOfInterest)
			for _req_if_content := range set {
				nodeInstance := (&tree.Node{Name: _req_if_content.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_content, "REQ_IF_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF_HEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF_HEADER](probe.stageOfInterest)
			for _req_if_header := range set {
				nodeInstance := (&tree.Node{Name: _req_if_header.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_header, "REQ_IF_HEADER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF_TOOL_EXTENSION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF_TOOL_EXTENSION](probe.stageOfInterest)
			for _req_if_tool_extension := range set {
				nodeInstance := (&tree.Node{Name: _req_if_tool_extension.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_tool_extension, "REQ_IF_TOOL_EXTENSION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECIFICATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECIFICATION](probe.stageOfInterest)
			for _specification := range set {
				nodeInstance := (&tree.Node{Name: _specification.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specification, "SPECIFICATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECIFICATION_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](probe.stageOfInterest)
			for _specification_type := range set {
				nodeInstance := (&tree.Node{Name: _specification_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specification_type, "SPECIFICATION_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_HIERARCHY":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_HIERARCHY](probe.stageOfInterest)
			for _spec_hierarchy := range set {
				nodeInstance := (&tree.Node{Name: _spec_hierarchy.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_hierarchy, "SPEC_HIERARCHY", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_OBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_OBJECT](probe.stageOfInterest)
			for _spec_object := range set {
				nodeInstance := (&tree.Node{Name: _spec_object.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_object, "SPEC_OBJECT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_OBJECT_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](probe.stageOfInterest)
			for _spec_object_type := range set {
				nodeInstance := (&tree.Node{Name: _spec_object_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_object_type, "SPEC_OBJECT_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_RELATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_RELATION](probe.stageOfInterest)
			for _spec_relation := range set {
				nodeInstance := (&tree.Node{Name: _spec_relation.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_relation, "SPEC_RELATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_RELATION_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](probe.stageOfInterest)
			for _spec_relation_type := range set {
				nodeInstance := (&tree.Node{Name: _spec_relation_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_relation_type, "SPEC_RELATION_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "XHTML_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.XHTML_CONTENT](probe.stageOfInterest)
			for _xhtml_content := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_content.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_content, "XHTML_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_InlPres_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_InlPres_type](probe.stageOfInterest)
			for _xhtml_inlpres_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_inlpres_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_inlpres_type, "Xhtml_InlPres_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_a_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_a_type](probe.stageOfInterest)
			for _xhtml_a_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_a_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_a_type, "Xhtml_a_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_abbr_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_abbr_type](probe.stageOfInterest)
			for _xhtml_abbr_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_abbr_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_abbr_type, "Xhtml_abbr_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_acronym_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_acronym_type](probe.stageOfInterest)
			for _xhtml_acronym_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_acronym_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_acronym_type, "Xhtml_acronym_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_address_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_address_type](probe.stageOfInterest)
			for _xhtml_address_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_address_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_address_type, "Xhtml_address_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_blockquote_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_blockquote_type](probe.stageOfInterest)
			for _xhtml_blockquote_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_blockquote_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_blockquote_type, "Xhtml_blockquote_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_br_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_br_type](probe.stageOfInterest)
			for _xhtml_br_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_br_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_br_type, "Xhtml_br_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_caption_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_caption_type](probe.stageOfInterest)
			for _xhtml_caption_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_caption_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_caption_type, "Xhtml_caption_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_cite_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_cite_type](probe.stageOfInterest)
			for _xhtml_cite_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_cite_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_cite_type, "Xhtml_cite_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_code_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_code_type](probe.stageOfInterest)
			for _xhtml_code_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_code_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_code_type, "Xhtml_code_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_col_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_col_type](probe.stageOfInterest)
			for _xhtml_col_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_col_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_col_type, "Xhtml_col_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_colgroup_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_colgroup_type](probe.stageOfInterest)
			for _xhtml_colgroup_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_colgroup_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_colgroup_type, "Xhtml_colgroup_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_dd_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_dd_type](probe.stageOfInterest)
			for _xhtml_dd_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_dd_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_dd_type, "Xhtml_dd_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_dfn_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_dfn_type](probe.stageOfInterest)
			for _xhtml_dfn_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_dfn_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_dfn_type, "Xhtml_dfn_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_div_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_div_type](probe.stageOfInterest)
			for _xhtml_div_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_div_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_div_type, "Xhtml_div_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_dl_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_dl_type](probe.stageOfInterest)
			for _xhtml_dl_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_dl_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_dl_type, "Xhtml_dl_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_dt_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_dt_type](probe.stageOfInterest)
			for _xhtml_dt_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_dt_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_dt_type, "Xhtml_dt_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_edit_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_edit_type](probe.stageOfInterest)
			for _xhtml_edit_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_edit_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_edit_type, "Xhtml_edit_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_em_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_em_type](probe.stageOfInterest)
			for _xhtml_em_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_em_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_em_type, "Xhtml_em_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_h1_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_h1_type](probe.stageOfInterest)
			for _xhtml_h1_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_h1_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_h1_type, "Xhtml_h1_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_h2_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_h2_type](probe.stageOfInterest)
			for _xhtml_h2_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_h2_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_h2_type, "Xhtml_h2_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_h3_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_h3_type](probe.stageOfInterest)
			for _xhtml_h3_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_h3_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_h3_type, "Xhtml_h3_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_h4_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_h4_type](probe.stageOfInterest)
			for _xhtml_h4_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_h4_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_h4_type, "Xhtml_h4_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_h5_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_h5_type](probe.stageOfInterest)
			for _xhtml_h5_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_h5_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_h5_type, "Xhtml_h5_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_h6_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_h6_type](probe.stageOfInterest)
			for _xhtml_h6_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_h6_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_h6_type, "Xhtml_h6_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_heading_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_heading_type](probe.stageOfInterest)
			for _xhtml_heading_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_heading_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_heading_type, "Xhtml_heading_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_hr_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_hr_type](probe.stageOfInterest)
			for _xhtml_hr_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_hr_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_hr_type, "Xhtml_hr_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_kbd_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_kbd_type](probe.stageOfInterest)
			for _xhtml_kbd_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_kbd_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_kbd_type, "Xhtml_kbd_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_li_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_li_type](probe.stageOfInterest)
			for _xhtml_li_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_li_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_li_type, "Xhtml_li_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_object_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_object_type](probe.stageOfInterest)
			for _xhtml_object_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_object_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_object_type, "Xhtml_object_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_ol_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_ol_type](probe.stageOfInterest)
			for _xhtml_ol_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_ol_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_ol_type, "Xhtml_ol_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_p_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_p_type](probe.stageOfInterest)
			for _xhtml_p_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_p_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_p_type, "Xhtml_p_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_param_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_param_type](probe.stageOfInterest)
			for _xhtml_param_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_param_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_param_type, "Xhtml_param_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_pre_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_pre_type](probe.stageOfInterest)
			for _xhtml_pre_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_pre_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_pre_type, "Xhtml_pre_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_q_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_q_type](probe.stageOfInterest)
			for _xhtml_q_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_q_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_q_type, "Xhtml_q_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_samp_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_samp_type](probe.stageOfInterest)
			for _xhtml_samp_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_samp_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_samp_type, "Xhtml_samp_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_span_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_span_type](probe.stageOfInterest)
			for _xhtml_span_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_span_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_span_type, "Xhtml_span_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_strong_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_strong_type](probe.stageOfInterest)
			for _xhtml_strong_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_strong_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_strong_type, "Xhtml_strong_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_table_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_table_type](probe.stageOfInterest)
			for _xhtml_table_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_table_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_table_type, "Xhtml_table_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_tbody_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_tbody_type](probe.stageOfInterest)
			for _xhtml_tbody_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_tbody_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_tbody_type, "Xhtml_tbody_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_td_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_td_type](probe.stageOfInterest)
			for _xhtml_td_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_td_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_td_type, "Xhtml_td_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_tfoot_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_tfoot_type](probe.stageOfInterest)
			for _xhtml_tfoot_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_tfoot_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_tfoot_type, "Xhtml_tfoot_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_th_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_th_type](probe.stageOfInterest)
			for _xhtml_th_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_th_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_th_type, "Xhtml_th_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_thead_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_thead_type](probe.stageOfInterest)
			for _xhtml_thead_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_thead_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_thead_type, "Xhtml_thead_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_tr_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_tr_type](probe.stageOfInterest)
			for _xhtml_tr_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_tr_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_tr_type, "Xhtml_tr_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_ul_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_ul_type](probe.stageOfInterest)
			for _xhtml_ul_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_ul_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_ul_type, "Xhtml_ul_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xhtml_var_type":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xhtml_var_type](probe.stageOfInterest)
			for _xhtml_var_type := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_var_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_var_type, "Xhtml_var_type", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(probe.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	// Add a refresh button
	nodeRefreshButton := (&tree.Node{Name: ""}).Stage(probe.treeStage)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := (&tree.Button{
		Name: "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon: string(gongtree_buttons.BUTTON_refresh)}).Stage(probe.treeStage)
	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	probe.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	probe *Probe) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.probe = probe
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.Stage,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
