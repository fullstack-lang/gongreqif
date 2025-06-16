package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

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
	sidebar := &tree.Tree{Name: SideBarTreeName}

	nodeRefreshButton := &tree.Node{Name: fmt.Sprintf("Stage %s, # %d, %s",
		probe.stageOfInterest.GetName(),
		probe.stageOfInterest.GetCommitId(),
		probe.stageOfInterest.GetCommitTS().Local().Format(time.Kitchen))}
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := &tree.Button{
		Name:            "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon:            string(gongtree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree.Left,
	}

	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

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

		nodeGongstruct := &tree.Node{Name: name}
		nodeGongstruct.HasToolTip = true
		nodeGongstruct.ToolTipText = "Display table of all " + name + " instances"
		nodeGongstruct.ToolTipPosition = tree.Right

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
				nodeInstance := &tree.Node{Name: _alternative_id.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_alternative_id, "ALTERNATIVE_ID", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_BOOLEAN](probe.stageOfInterest)
			for _attribute_definition_boolean := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_boolean.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_boolean, "ATTRIBUTE_DEFINITION_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_DATE](probe.stageOfInterest)
			for _attribute_definition_date := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_date.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_date, "ATTRIBUTE_DEFINITION_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_ENUMERATION](probe.stageOfInterest)
			for _attribute_definition_enumeration := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_enumeration.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_enumeration, "ATTRIBUTE_DEFINITION_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_INTEGER](probe.stageOfInterest)
			for _attribute_definition_integer := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_integer.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_integer, "ATTRIBUTE_DEFINITION_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_REAL](probe.stageOfInterest)
			for _attribute_definition_real := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_real.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_real, "ATTRIBUTE_DEFINITION_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_STRING](probe.stageOfInterest)
			for _attribute_definition_string := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_string.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_string, "ATTRIBUTE_DEFINITION_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_XHTML](probe.stageOfInterest)
			for _attribute_definition_xhtml := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_xhtml.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_xhtml, "ATTRIBUTE_DEFINITION_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_BOOLEAN](probe.stageOfInterest)
			for _attribute_value_boolean := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_boolean.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_boolean, "ATTRIBUTE_VALUE_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_DATE](probe.stageOfInterest)
			for _attribute_value_date := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_date.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_date, "ATTRIBUTE_VALUE_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_ENUMERATION](probe.stageOfInterest)
			for _attribute_value_enumeration := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_enumeration.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_enumeration, "ATTRIBUTE_VALUE_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_INTEGER](probe.stageOfInterest)
			for _attribute_value_integer := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_integer.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_integer, "ATTRIBUTE_VALUE_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_REAL](probe.stageOfInterest)
			for _attribute_value_real := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_real.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_real, "ATTRIBUTE_VALUE_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_STRING](probe.stageOfInterest)
			for _attribute_value_string := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_string.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_string, "ATTRIBUTE_VALUE_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_XHTML](probe.stageOfInterest)
			for _attribute_value_xhtml := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_xhtml.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_xhtml, "ATTRIBUTE_VALUE_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ALTERNATIVE_ID":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ALTERNATIVE_ID](probe.stageOfInterest)
			for _a_alternative_id := range set {
				nodeInstance := &tree.Node{Name: _a_alternative_id.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_alternative_id, "A_ALTERNATIVE_ID", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](probe.stageOfInterest)
			for _a_attribute_definition_boolean_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_boolean_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_boolean_ref, "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_DEFINITION_DATE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_DEFINITION_DATE_REF](probe.stageOfInterest)
			for _a_attribute_definition_date_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_date_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_date_ref, "A_ATTRIBUTE_DEFINITION_DATE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](probe.stageOfInterest)
			for _a_attribute_definition_enumeration_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_enumeration_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_enumeration_ref, "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](probe.stageOfInterest)
			for _a_attribute_definition_integer_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_integer_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_integer_ref, "A_ATTRIBUTE_DEFINITION_INTEGER_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_DEFINITION_REAL_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_DEFINITION_REAL_REF](probe.stageOfInterest)
			for _a_attribute_definition_real_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_real_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_real_ref, "A_ATTRIBUTE_DEFINITION_REAL_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_DEFINITION_STRING_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_DEFINITION_STRING_REF](probe.stageOfInterest)
			for _a_attribute_definition_string_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_string_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_string_ref, "A_ATTRIBUTE_DEFINITION_STRING_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_DEFINITION_XHTML_REF](probe.stageOfInterest)
			for _a_attribute_definition_xhtml_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_xhtml_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_xhtml_ref, "A_ATTRIBUTE_DEFINITION_XHTML_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_VALUE_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_BOOLEAN](probe.stageOfInterest)
			for _a_attribute_value_boolean := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_boolean.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_boolean, "A_ATTRIBUTE_VALUE_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_VALUE_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_DATE](probe.stageOfInterest)
			for _a_attribute_value_date := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_date.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_date, "A_ATTRIBUTE_VALUE_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_VALUE_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_ENUMERATION](probe.stageOfInterest)
			for _a_attribute_value_enumeration := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_enumeration.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_enumeration, "A_ATTRIBUTE_VALUE_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_VALUE_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_INTEGER](probe.stageOfInterest)
			for _a_attribute_value_integer := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_integer.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_integer, "A_ATTRIBUTE_VALUE_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_VALUE_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_REAL](probe.stageOfInterest)
			for _a_attribute_value_real := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_real.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_real, "A_ATTRIBUTE_VALUE_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_VALUE_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_STRING](probe.stageOfInterest)
			for _a_attribute_value_string := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_string.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_string, "A_ATTRIBUTE_VALUE_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_VALUE_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML](probe.stageOfInterest)
			for _a_attribute_value_xhtml := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_xhtml.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_xhtml, "A_ATTRIBUTE_VALUE_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML_1](probe.stageOfInterest)
			for _a_attribute_value_xhtml_1 := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_xhtml_1.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_xhtml_1, "A_ATTRIBUTE_VALUE_XHTML_1", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_CHILDREN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_CHILDREN](probe.stageOfInterest)
			for _a_children := range set {
				nodeInstance := &tree.Node{Name: _a_children.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_children, "A_CHILDREN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_CORE_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_CORE_CONTENT](probe.stageOfInterest)
			for _a_core_content := range set {
				nodeInstance := &tree.Node{Name: _a_core_content.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_core_content, "A_CORE_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DATATYPES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DATATYPES](probe.stageOfInterest)
			for _a_datatypes := range set {
				nodeInstance := &tree.Node{Name: _a_datatypes.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatypes, "A_DATATYPES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DATATYPE_DEFINITION_BOOLEAN_REF](probe.stageOfInterest)
			for _a_datatype_definition_boolean_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_boolean_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_boolean_ref, "A_DATATYPE_DEFINITION_BOOLEAN_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DATATYPE_DEFINITION_DATE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DATATYPE_DEFINITION_DATE_REF](probe.stageOfInterest)
			for _a_datatype_definition_date_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_date_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_date_ref, "A_DATATYPE_DEFINITION_DATE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DATATYPE_DEFINITION_ENUMERATION_REF](probe.stageOfInterest)
			for _a_datatype_definition_enumeration_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_enumeration_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_enumeration_ref, "A_DATATYPE_DEFINITION_ENUMERATION_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DATATYPE_DEFINITION_INTEGER_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DATATYPE_DEFINITION_INTEGER_REF](probe.stageOfInterest)
			for _a_datatype_definition_integer_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_integer_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_integer_ref, "A_DATATYPE_DEFINITION_INTEGER_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DATATYPE_DEFINITION_REAL_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DATATYPE_DEFINITION_REAL_REF](probe.stageOfInterest)
			for _a_datatype_definition_real_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_real_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_real_ref, "A_DATATYPE_DEFINITION_REAL_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DATATYPE_DEFINITION_STRING_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DATATYPE_DEFINITION_STRING_REF](probe.stageOfInterest)
			for _a_datatype_definition_string_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_string_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_string_ref, "A_DATATYPE_DEFINITION_STRING_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DATATYPE_DEFINITION_XHTML_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DATATYPE_DEFINITION_XHTML_REF](probe.stageOfInterest)
			for _a_datatype_definition_xhtml_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_xhtml_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_xhtml_ref, "A_DATATYPE_DEFINITION_XHTML_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_EDITABLE_ATTS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_EDITABLE_ATTS](probe.stageOfInterest)
			for _a_editable_atts := range set {
				nodeInstance := &tree.Node{Name: _a_editable_atts.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_editable_atts, "A_EDITABLE_ATTS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_ENUM_VALUE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ENUM_VALUE_REF](probe.stageOfInterest)
			for _a_enum_value_ref := range set {
				nodeInstance := &tree.Node{Name: _a_enum_value_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_enum_value_ref, "A_ENUM_VALUE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_OBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_OBJECT](probe.stageOfInterest)
			for _a_object := range set {
				nodeInstance := &tree.Node{Name: _a_object.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_object, "A_OBJECT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_PROPERTIES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_PROPERTIES](probe.stageOfInterest)
			for _a_properties := range set {
				nodeInstance := &tree.Node{Name: _a_properties.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_properties, "A_PROPERTIES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_RELATION_GROUP_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_RELATION_GROUP_TYPE_REF](probe.stageOfInterest)
			for _a_relation_group_type_ref := range set {
				nodeInstance := &tree.Node{Name: _a_relation_group_type_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_relation_group_type_ref, "A_RELATION_GROUP_TYPE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SOURCE_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SOURCE_1](probe.stageOfInterest)
			for _a_source_1 := range set {
				nodeInstance := &tree.Node{Name: _a_source_1.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_source_1, "A_SOURCE_1", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SOURCE_SPECIFICATION_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SOURCE_SPECIFICATION_1](probe.stageOfInterest)
			for _a_source_specification_1 := range set {
				nodeInstance := &tree.Node{Name: _a_source_specification_1.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_source_specification_1, "A_SOURCE_SPECIFICATION_1", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPECIFICATIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPECIFICATIONS](probe.stageOfInterest)
			for _a_specifications := range set {
				nodeInstance := &tree.Node{Name: _a_specifications.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_specifications, "A_SPECIFICATIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPECIFICATION_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPECIFICATION_TYPE_REF](probe.stageOfInterest)
			for _a_specification_type_ref := range set {
				nodeInstance := &tree.Node{Name: _a_specification_type_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_specification_type_ref, "A_SPECIFICATION_TYPE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPECIFIED_VALUES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPECIFIED_VALUES](probe.stageOfInterest)
			for _a_specified_values := range set {
				nodeInstance := &tree.Node{Name: _a_specified_values.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_specified_values, "A_SPECIFIED_VALUES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_ATTRIBUTES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](probe.stageOfInterest)
			for _a_spec_attributes := range set {
				nodeInstance := &tree.Node{Name: _a_spec_attributes.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_attributes, "A_SPEC_ATTRIBUTES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_OBJECTS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_OBJECTS](probe.stageOfInterest)
			for _a_spec_objects := range set {
				nodeInstance := &tree.Node{Name: _a_spec_objects.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_objects, "A_SPEC_OBJECTS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_OBJECT_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_OBJECT_TYPE_REF](probe.stageOfInterest)
			for _a_spec_object_type_ref := range set {
				nodeInstance := &tree.Node{Name: _a_spec_object_type_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_object_type_ref, "A_SPEC_OBJECT_TYPE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_RELATIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_RELATIONS](probe.stageOfInterest)
			for _a_spec_relations := range set {
				nodeInstance := &tree.Node{Name: _a_spec_relations.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relations, "A_SPEC_RELATIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_RELATION_GROUPS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_RELATION_GROUPS](probe.stageOfInterest)
			for _a_spec_relation_groups := range set {
				nodeInstance := &tree.Node{Name: _a_spec_relation_groups.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relation_groups, "A_SPEC_RELATION_GROUPS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_RELATION_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_RELATION_REF](probe.stageOfInterest)
			for _a_spec_relation_ref := range set {
				nodeInstance := &tree.Node{Name: _a_spec_relation_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relation_ref, "A_SPEC_RELATION_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_RELATION_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_RELATION_TYPE_REF](probe.stageOfInterest)
			for _a_spec_relation_type_ref := range set {
				nodeInstance := &tree.Node{Name: _a_spec_relation_type_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relation_type_ref, "A_SPEC_RELATION_TYPE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_TYPES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_TYPES](probe.stageOfInterest)
			for _a_spec_types := range set {
				nodeInstance := &tree.Node{Name: _a_spec_types.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_types, "A_SPEC_TYPES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_THE_HEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_THE_HEADER](probe.stageOfInterest)
			for _a_the_header := range set {
				nodeInstance := &tree.Node{Name: _a_the_header.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_the_header, "A_THE_HEADER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TOOL_EXTENSIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TOOL_EXTENSIONS](probe.stageOfInterest)
			for _a_tool_extensions := range set {
				nodeInstance := &tree.Node{Name: _a_tool_extensions.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_tool_extensions, "A_TOOL_EXTENSIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_BOOLEAN](probe.stageOfInterest)
			for _datatype_definition_boolean := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_boolean.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_boolean, "DATATYPE_DEFINITION_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_DATE](probe.stageOfInterest)
			for _datatype_definition_date := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_date.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_date, "DATATYPE_DEFINITION_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_ENUMERATION](probe.stageOfInterest)
			for _datatype_definition_enumeration := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_enumeration.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_enumeration, "DATATYPE_DEFINITION_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_INTEGER](probe.stageOfInterest)
			for _datatype_definition_integer := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_integer.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_integer, "DATATYPE_DEFINITION_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_REAL](probe.stageOfInterest)
			for _datatype_definition_real := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_real.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_real, "DATATYPE_DEFINITION_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_STRING](probe.stageOfInterest)
			for _datatype_definition_string := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_string.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_string, "DATATYPE_DEFINITION_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_XHTML](probe.stageOfInterest)
			for _datatype_definition_xhtml := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_xhtml.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_xhtml, "DATATYPE_DEFINITION_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "EMBEDDED_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.EMBEDDED_VALUE](probe.stageOfInterest)
			for _embedded_value := range set {
				nodeInstance := &tree.Node{Name: _embedded_value.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_embedded_value, "EMBEDDED_VALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ENUM_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ENUM_VALUE](probe.stageOfInterest)
			for _enum_value := range set {
				nodeInstance := &tree.Node{Name: _enum_value.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_enum_value, "ENUM_VALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GeneratedImageMetamodel":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GeneratedImageMetamodel](probe.stageOfInterest)
			for _generatedimagemetamodel := range set {
				nodeInstance := &tree.Node{Name: _generatedimagemetamodel.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_generatedimagemetamodel, "GeneratedImageMetamodel", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Image":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Image](probe.stageOfInterest)
			for _image := range set {
				nodeInstance := &tree.Node{Name: _image.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_image, "Image", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Paragraph":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Paragraph](probe.stageOfInterest)
			for _paragraph := range set {
				nodeInstance := &tree.Node{Name: _paragraph.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_paragraph, "Paragraph", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RELATION_GROUP":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RELATION_GROUP](probe.stageOfInterest)
			for _relation_group := range set {
				nodeInstance := &tree.Node{Name: _relation_group.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_relation_group, "RELATION_GROUP", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RELATION_GROUP_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](probe.stageOfInterest)
			for _relation_group_type := range set {
				nodeInstance := &tree.Node{Name: _relation_group_type.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_relation_group_type, "RELATION_GROUP_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF](probe.stageOfInterest)
			for _req_if := range set {
				nodeInstance := &tree.Node{Name: _req_if.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if, "REQ_IF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](probe.stageOfInterest)
			for _req_if_content := range set {
				nodeInstance := &tree.Node{Name: _req_if_content.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_content, "REQ_IF_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF_HEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF_HEADER](probe.stageOfInterest)
			for _req_if_header := range set {
				nodeInstance := &tree.Node{Name: _req_if_header.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_header, "REQ_IF_HEADER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF_TOOL_EXTENSION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF_TOOL_EXTENSION](probe.stageOfInterest)
			for _req_if_tool_extension := range set {
				nodeInstance := &tree.Node{Name: _req_if_tool_extension.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_tool_extension, "REQ_IF_TOOL_EXTENSION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECIFICATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECIFICATION](probe.stageOfInterest)
			for _specification := range set {
				nodeInstance := &tree.Node{Name: _specification.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specification, "SPECIFICATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECIFICATION_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](probe.stageOfInterest)
			for _specification_type := range set {
				nodeInstance := &tree.Node{Name: _specification_type.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specification_type, "SPECIFICATION_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_HIERARCHY":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_HIERARCHY](probe.stageOfInterest)
			for _spec_hierarchy := range set {
				nodeInstance := &tree.Node{Name: _spec_hierarchy.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_hierarchy, "SPEC_HIERARCHY", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_OBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_OBJECT](probe.stageOfInterest)
			for _spec_object := range set {
				nodeInstance := &tree.Node{Name: _spec_object.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_object, "SPEC_OBJECT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_OBJECT_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](probe.stageOfInterest)
			for _spec_object_type := range set {
				nodeInstance := &tree.Node{Name: _spec_object_type.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_object_type, "SPEC_OBJECT_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_RELATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_RELATION](probe.stageOfInterest)
			for _spec_relation := range set {
				nodeInstance := &tree.Node{Name: _spec_relation.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_relation, "SPEC_RELATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_RELATION_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](probe.stageOfInterest)
			for _spec_relation_type := range set {
				nodeInstance := &tree.Node{Name: _spec_relation_type.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_relation_type, "SPEC_RELATION_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "StaticWebSite":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.StaticWebSite](probe.stageOfInterest)
			for _staticwebsite := range set {
				nodeInstance := &tree.Node{Name: _staticwebsite.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staticwebsite, "StaticWebSite", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "StaticWebSiteChapter":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.StaticWebSiteChapter](probe.stageOfInterest)
			for _staticwebsitechapter := range set {
				nodeInstance := &tree.Node{Name: _staticwebsitechapter.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staticwebsitechapter, "StaticWebSiteChapter", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "XHTML_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.XHTML_CONTENT](probe.stageOfInterest)
			for _xhtml_content := range set {
				nodeInstance := &tree.Node{Name: _xhtml_content.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_content, "XHTML_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := &tree.Button{
			Name:            gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon:            string(gongtree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of " + gongStruct.GetName(),
			ToolTipPosition: tree.Right,
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	tree.StageBranch(probe.treeStage, sidebar)

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
