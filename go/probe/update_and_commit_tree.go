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
	sidebar := &tree.Tree{Name: "Sidebar"}

	nodeRefreshButton := &tree.Node{Name: fmt.Sprintf("Stage %s, # %d, %s",
		probe.stageOfInterest.GetName(),
		probe.stageOfInterest.GetCommitId(),
		probe.stageOfInterest.GetCommitTS().Local().Format(time.Kitchen))}
	nodeRefreshButton.Name +=
		fmt.Sprintf(" (%d/%d/%d)", 
			len(probe.stageOfInterest.GetNew()), 
			len(probe.stageOfInterest.GetModified()), 
			len(probe.stageOfInterest.GetDeleted()),
		)
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
	setOfGongStructs := *gong_models.GetGongstructInstancesSetFromPointerType[*gong_models.GongStruct](probe.gongStage)

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
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ALTERNATIVE_ID](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _alternative_id := range set {
				nodeInstance := &tree.Node{Name: _alternative_id.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_alternative_id, "ALTERNATIVE_ID", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_alternative_id]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_alternative_id]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_alternative_id]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_BOOLEAN](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_definition_boolean := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_boolean.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_boolean, "ATTRIBUTE_DEFINITION_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_definition_boolean]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_definition_boolean]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_definition_boolean]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_DEFINITION_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_DATE](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_definition_date := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_date.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_date, "ATTRIBUTE_DEFINITION_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_definition_date]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_definition_date]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_definition_date]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_ENUMERATION](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_definition_enumeration := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_enumeration.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_enumeration, "ATTRIBUTE_DEFINITION_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_definition_enumeration]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_definition_enumeration]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_definition_enumeration]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_DEFINITION_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_INTEGER](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_definition_integer := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_integer.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_integer, "ATTRIBUTE_DEFINITION_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_definition_integer]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_definition_integer]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_definition_integer]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_DEFINITION_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_REAL](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_definition_real := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_real.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_real, "ATTRIBUTE_DEFINITION_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_definition_real]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_definition_real]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_definition_real]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_DEFINITION_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_STRING](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_definition_string := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_string.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_string, "ATTRIBUTE_DEFINITION_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_definition_string]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_definition_string]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_definition_string]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_DEFINITION_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_XHTML](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_definition_xhtml := range set {
				nodeInstance := &tree.Node{Name: _attribute_definition_xhtml.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_xhtml, "ATTRIBUTE_DEFINITION_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_definition_xhtml]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_definition_xhtml]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_definition_xhtml]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_VALUE_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_BOOLEAN](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_value_boolean := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_boolean.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_boolean, "ATTRIBUTE_VALUE_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_value_boolean]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_value_boolean]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_value_boolean]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_VALUE_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_DATE](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_value_date := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_date.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_date, "ATTRIBUTE_VALUE_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_value_date]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_value_date]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_value_date]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_VALUE_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_ENUMERATION](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_value_enumeration := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_enumeration.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_enumeration, "ATTRIBUTE_VALUE_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_value_enumeration]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_value_enumeration]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_value_enumeration]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_VALUE_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_INTEGER](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_value_integer := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_integer.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_integer, "ATTRIBUTE_VALUE_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_value_integer]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_value_integer]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_value_integer]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_VALUE_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_REAL](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_value_real := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_real.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_real, "ATTRIBUTE_VALUE_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_value_real]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_value_real]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_value_real]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_VALUE_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_STRING](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_value_string := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_string.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_string, "ATTRIBUTE_VALUE_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_value_string]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_value_string]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_value_string]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ATTRIBUTE_VALUE_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_XHTML](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attribute_value_xhtml := range set {
				nodeInstance := &tree.Node{Name: _attribute_value_xhtml.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_xhtml, "ATTRIBUTE_VALUE_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attribute_value_xhtml]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attribute_value_xhtml]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attribute_value_xhtml]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ALTERNATIVE_ID":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ALTERNATIVE_ID](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_alternative_id := range set {
				nodeInstance := &tree.Node{Name: _a_alternative_id.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_alternative_id, "A_ALTERNATIVE_ID", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_alternative_id]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_alternative_id]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_alternative_id]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_definition_boolean_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_boolean_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_boolean_ref, "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_definition_boolean_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_definition_boolean_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_definition_boolean_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_DEFINITION_DATE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_DATE_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_definition_date_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_date_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_date_ref, "A_ATTRIBUTE_DEFINITION_DATE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_definition_date_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_definition_date_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_definition_date_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_definition_enumeration_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_enumeration_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_enumeration_ref, "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_definition_enumeration_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_definition_enumeration_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_definition_enumeration_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_definition_integer_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_integer_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_integer_ref, "A_ATTRIBUTE_DEFINITION_INTEGER_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_definition_integer_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_definition_integer_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_definition_integer_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_DEFINITION_REAL_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_REAL_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_definition_real_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_real_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_real_ref, "A_ATTRIBUTE_DEFINITION_REAL_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_definition_real_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_definition_real_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_definition_real_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_DEFINITION_STRING_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_STRING_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_definition_string_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_string_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_string_ref, "A_ATTRIBUTE_DEFINITION_STRING_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_definition_string_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_definition_string_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_definition_string_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_definition_xhtml_ref := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_definition_xhtml_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_definition_xhtml_ref, "A_ATTRIBUTE_DEFINITION_XHTML_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_definition_xhtml_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_definition_xhtml_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_definition_xhtml_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_VALUE_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_BOOLEAN](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_value_boolean := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_boolean.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_boolean, "A_ATTRIBUTE_VALUE_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_value_boolean]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_value_boolean]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_value_boolean]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_VALUE_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_DATE](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_value_date := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_date.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_date, "A_ATTRIBUTE_VALUE_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_value_date]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_value_date]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_value_date]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_VALUE_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_ENUMERATION](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_value_enumeration := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_enumeration.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_enumeration, "A_ATTRIBUTE_VALUE_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_value_enumeration]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_value_enumeration]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_value_enumeration]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_VALUE_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_INTEGER](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_value_integer := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_integer.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_integer, "A_ATTRIBUTE_VALUE_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_value_integer]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_value_integer]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_value_integer]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_VALUE_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_REAL](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_value_real := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_real.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_real, "A_ATTRIBUTE_VALUE_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_value_real]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_value_real]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_value_real]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_VALUE_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_STRING](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_value_string := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_string.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_string, "A_ATTRIBUTE_VALUE_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_value_string]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_value_string]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_value_string]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_VALUE_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_XHTML](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_value_xhtml := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_xhtml.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_xhtml, "A_ATTRIBUTE_VALUE_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_value_xhtml]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_value_xhtml]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_value_xhtml]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_XHTML_1](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_attribute_value_xhtml_1 := range set {
				nodeInstance := &tree.Node{Name: _a_attribute_value_xhtml_1.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_attribute_value_xhtml_1, "A_ATTRIBUTE_VALUE_XHTML_1", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_attribute_value_xhtml_1]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_attribute_value_xhtml_1]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_attribute_value_xhtml_1]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_CHILDREN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_CHILDREN](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_children := range set {
				nodeInstance := &tree.Node{Name: _a_children.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_children, "A_CHILDREN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_children]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_children]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_children]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_CORE_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_CORE_CONTENT](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_core_content := range set {
				nodeInstance := &tree.Node{Name: _a_core_content.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_core_content, "A_CORE_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_core_content]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_core_content]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_core_content]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_DATATYPES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPES](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_datatypes := range set {
				nodeInstance := &tree.Node{Name: _a_datatypes.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatypes, "A_DATATYPES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_datatypes]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_datatypes]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_datatypes]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_datatype_definition_boolean_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_boolean_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_boolean_ref, "A_DATATYPE_DEFINITION_BOOLEAN_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_datatype_definition_boolean_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_datatype_definition_boolean_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_datatype_definition_boolean_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_DATATYPE_DEFINITION_DATE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_DATE_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_datatype_definition_date_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_date_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_date_ref, "A_DATATYPE_DEFINITION_DATE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_datatype_definition_date_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_datatype_definition_date_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_datatype_definition_date_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_datatype_definition_enumeration_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_enumeration_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_enumeration_ref, "A_DATATYPE_DEFINITION_ENUMERATION_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_datatype_definition_enumeration_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_datatype_definition_enumeration_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_datatype_definition_enumeration_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_DATATYPE_DEFINITION_INTEGER_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_INTEGER_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_datatype_definition_integer_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_integer_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_integer_ref, "A_DATATYPE_DEFINITION_INTEGER_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_datatype_definition_integer_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_datatype_definition_integer_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_datatype_definition_integer_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_DATATYPE_DEFINITION_REAL_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_REAL_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_datatype_definition_real_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_real_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_real_ref, "A_DATATYPE_DEFINITION_REAL_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_datatype_definition_real_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_datatype_definition_real_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_datatype_definition_real_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_DATATYPE_DEFINITION_STRING_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_STRING_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_datatype_definition_string_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_string_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_string_ref, "A_DATATYPE_DEFINITION_STRING_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_datatype_definition_string_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_datatype_definition_string_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_datatype_definition_string_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_DATATYPE_DEFINITION_XHTML_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_XHTML_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_datatype_definition_xhtml_ref := range set {
				nodeInstance := &tree.Node{Name: _a_datatype_definition_xhtml_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatype_definition_xhtml_ref, "A_DATATYPE_DEFINITION_XHTML_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_datatype_definition_xhtml_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_datatype_definition_xhtml_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_datatype_definition_xhtml_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_EDITABLE_ATTS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_EDITABLE_ATTS](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_editable_atts := range set {
				nodeInstance := &tree.Node{Name: _a_editable_atts.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_editable_atts, "A_EDITABLE_ATTS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_editable_atts]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_editable_atts]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_editable_atts]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_ENUM_VALUE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ENUM_VALUE_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_enum_value_ref := range set {
				nodeInstance := &tree.Node{Name: _a_enum_value_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_enum_value_ref, "A_ENUM_VALUE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_enum_value_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_enum_value_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_enum_value_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_OBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_OBJECT](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_object := range set {
				nodeInstance := &tree.Node{Name: _a_object.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_object, "A_OBJECT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_object]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_object]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_object]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_PROPERTIES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_PROPERTIES](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_properties := range set {
				nodeInstance := &tree.Node{Name: _a_properties.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_properties, "A_PROPERTIES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_properties]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_properties]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_properties]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_RELATION_GROUP_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_RELATION_GROUP_TYPE_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_relation_group_type_ref := range set {
				nodeInstance := &tree.Node{Name: _a_relation_group_type_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_relation_group_type_ref, "A_RELATION_GROUP_TYPE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_relation_group_type_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_relation_group_type_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_relation_group_type_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SOURCE_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SOURCE_1](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_source_1 := range set {
				nodeInstance := &tree.Node{Name: _a_source_1.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_source_1, "A_SOURCE_1", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_source_1]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_source_1]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_source_1]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SOURCE_SPECIFICATION_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SOURCE_SPECIFICATION_1](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_source_specification_1 := range set {
				nodeInstance := &tree.Node{Name: _a_source_specification_1.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_source_specification_1, "A_SOURCE_SPECIFICATION_1", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_source_specification_1]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_source_specification_1]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_source_specification_1]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPECIFICATIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPECIFICATIONS](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_specifications := range set {
				nodeInstance := &tree.Node{Name: _a_specifications.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_specifications, "A_SPECIFICATIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_specifications]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_specifications]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_specifications]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPECIFICATION_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPECIFICATION_TYPE_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_specification_type_ref := range set {
				nodeInstance := &tree.Node{Name: _a_specification_type_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_specification_type_ref, "A_SPECIFICATION_TYPE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_specification_type_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_specification_type_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_specification_type_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPECIFIED_VALUES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPECIFIED_VALUES](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_specified_values := range set {
				nodeInstance := &tree.Node{Name: _a_specified_values.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_specified_values, "A_SPECIFIED_VALUES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_specified_values]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_specified_values]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_specified_values]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPEC_ATTRIBUTES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_ATTRIBUTES](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_spec_attributes := range set {
				nodeInstance := &tree.Node{Name: _a_spec_attributes.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_attributes, "A_SPEC_ATTRIBUTES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_spec_attributes]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_spec_attributes]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_spec_attributes]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPEC_OBJECTS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_OBJECTS](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_spec_objects := range set {
				nodeInstance := &tree.Node{Name: _a_spec_objects.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_objects, "A_SPEC_OBJECTS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_spec_objects]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_spec_objects]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_spec_objects]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPEC_OBJECT_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_OBJECT_TYPE_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_spec_object_type_ref := range set {
				nodeInstance := &tree.Node{Name: _a_spec_object_type_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_object_type_ref, "A_SPEC_OBJECT_TYPE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_spec_object_type_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_spec_object_type_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_spec_object_type_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPEC_RELATIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_RELATIONS](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_spec_relations := range set {
				nodeInstance := &tree.Node{Name: _a_spec_relations.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relations, "A_SPEC_RELATIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_spec_relations]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_spec_relations]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_spec_relations]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPEC_RELATION_GROUPS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_RELATION_GROUPS](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_spec_relation_groups := range set {
				nodeInstance := &tree.Node{Name: _a_spec_relation_groups.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relation_groups, "A_SPEC_RELATION_GROUPS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_spec_relation_groups]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_spec_relation_groups]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_spec_relation_groups]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPEC_RELATION_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_RELATION_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_spec_relation_ref := range set {
				nodeInstance := &tree.Node{Name: _a_spec_relation_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relation_ref, "A_SPEC_RELATION_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_spec_relation_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_spec_relation_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_spec_relation_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPEC_RELATION_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_RELATION_TYPE_REF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_spec_relation_type_ref := range set {
				nodeInstance := &tree.Node{Name: _a_spec_relation_type_ref.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relation_type_ref, "A_SPEC_RELATION_TYPE_REF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_spec_relation_type_ref]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_spec_relation_type_ref]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_spec_relation_type_ref]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_SPEC_TYPES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_TYPES](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_spec_types := range set {
				nodeInstance := &tree.Node{Name: _a_spec_types.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_types, "A_SPEC_TYPES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_spec_types]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_spec_types]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_spec_types]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_THE_HEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_THE_HEADER](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_the_header := range set {
				nodeInstance := &tree.Node{Name: _a_the_header.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_the_header, "A_THE_HEADER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_the_header]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_the_header]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_the_header]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "A_TOOL_EXTENSIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_TOOL_EXTENSIONS](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _a_tool_extensions := range set {
				nodeInstance := &tree.Node{Name: _a_tool_extensions.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_tool_extensions, "A_TOOL_EXTENSIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_a_tool_extensions]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_a_tool_extensions]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_a_tool_extensions]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "DATATYPE_DEFINITION_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_BOOLEAN](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _datatype_definition_boolean := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_boolean.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_boolean, "DATATYPE_DEFINITION_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_datatype_definition_boolean]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_datatype_definition_boolean]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_datatype_definition_boolean]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "DATATYPE_DEFINITION_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_DATE](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _datatype_definition_date := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_date.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_date, "DATATYPE_DEFINITION_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_datatype_definition_date]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_datatype_definition_date]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_datatype_definition_date]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "DATATYPE_DEFINITION_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_ENUMERATION](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _datatype_definition_enumeration := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_enumeration.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_enumeration, "DATATYPE_DEFINITION_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_datatype_definition_enumeration]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_datatype_definition_enumeration]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_datatype_definition_enumeration]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "DATATYPE_DEFINITION_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_INTEGER](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _datatype_definition_integer := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_integer.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_integer, "DATATYPE_DEFINITION_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_datatype_definition_integer]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_datatype_definition_integer]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_datatype_definition_integer]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "DATATYPE_DEFINITION_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_REAL](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _datatype_definition_real := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_real.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_real, "DATATYPE_DEFINITION_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_datatype_definition_real]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_datatype_definition_real]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_datatype_definition_real]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "DATATYPE_DEFINITION_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_STRING](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _datatype_definition_string := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_string.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_string, "DATATYPE_DEFINITION_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_datatype_definition_string]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_datatype_definition_string]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_datatype_definition_string]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "DATATYPE_DEFINITION_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_XHTML](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _datatype_definition_xhtml := range set {
				nodeInstance := &tree.Node{Name: _datatype_definition_xhtml.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_xhtml, "DATATYPE_DEFINITION_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_datatype_definition_xhtml]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_datatype_definition_xhtml]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_datatype_definition_xhtml]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "EMBEDDED_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EMBEDDED_VALUE](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _embedded_value := range set {
				nodeInstance := &tree.Node{Name: _embedded_value.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_embedded_value, "EMBEDDED_VALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_embedded_value]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_embedded_value]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_embedded_value]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "ENUM_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ENUM_VALUE](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _enum_value := range set {
				nodeInstance := &tree.Node{Name: _enum_value.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_enum_value, "ENUM_VALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_enum_value]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_enum_value]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_enum_value]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "EmbeddedJpgImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EmbeddedJpgImage](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _embeddedjpgimage := range set {
				nodeInstance := &tree.Node{Name: _embeddedjpgimage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_embeddedjpgimage, "EmbeddedJpgImage", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_embeddedjpgimage]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_embeddedjpgimage]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_embeddedjpgimage]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "EmbeddedPngImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EmbeddedPngImage](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _embeddedpngimage := range set {
				nodeInstance := &tree.Node{Name: _embeddedpngimage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_embeddedpngimage, "EmbeddedPngImage", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_embeddedpngimage]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_embeddedpngimage]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_embeddedpngimage]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "EmbeddedSvgImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EmbeddedSvgImage](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _embeddedsvgimage := range set {
				nodeInstance := &tree.Node{Name: _embeddedsvgimage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_embeddedsvgimage, "EmbeddedSvgImage", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_embeddedsvgimage]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_embeddedsvgimage]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_embeddedsvgimage]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Kill":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Kill](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _kill := range set {
				nodeInstance := &tree.Node{Name: _kill.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_kill, "Kill", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_kill]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_kill]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_kill]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_boolean_showinsubjectentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_boolean_showinsubjectentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_boolean_showinsubjectentry, "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_boolean_showinsubjectentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_boolean_showinsubjectentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_boolean_showinsubjectentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_boolean_showintableentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_boolean_showintableentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_boolean_showintableentry, "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_boolean_showintableentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_boolean_showintableentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_boolean_showintableentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_boolean_showintitleentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_boolean_showintitleentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_boolean_showintitleentry, "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_boolean_showintitleentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_boolean_showintitleentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_boolean_showintitleentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_date_showinsubjectentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_date_showinsubjectentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_date_showinsubjectentry, "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_date_showinsubjectentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_date_showinsubjectentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_date_showinsubjectentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_date_showintableentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_date_showintableentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_date_showintableentry, "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_date_showintableentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_date_showintableentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_date_showintableentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_date_showintitleentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_date_showintitleentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_date_showintitleentry, "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_date_showintitleentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_date_showintitleentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_date_showintitleentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_enumeration_showinsubjectentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_enumeration_showinsubjectentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_enumeration_showinsubjectentry, "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_enumeration_showinsubjectentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_enumeration_showinsubjectentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_enumeration_showinsubjectentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_enumeration_showintableentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_enumeration_showintableentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_enumeration_showintableentry, "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_enumeration_showintableentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_enumeration_showintableentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_enumeration_showintableentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_enumeration_showintitleentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_enumeration_showintitleentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_enumeration_showintitleentry, "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_enumeration_showintitleentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_enumeration_showintitleentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_enumeration_showintitleentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_integer_showinsubjectentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_integer_showinsubjectentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_integer_showinsubjectentry, "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_integer_showinsubjectentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_integer_showinsubjectentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_integer_showinsubjectentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_integer_showintableentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_integer_showintableentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_integer_showintableentry, "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_integer_showintableentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_integer_showintableentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_integer_showintableentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_integer_showintitleentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_integer_showintitleentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_integer_showintitleentry, "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_integer_showintitleentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_integer_showintitleentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_integer_showintitleentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_real_showinsubjectentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_real_showinsubjectentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_real_showinsubjectentry, "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_real_showinsubjectentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_real_showinsubjectentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_real_showinsubjectentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_real_showintableentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_real_showintableentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_real_showintableentry, "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_real_showintableentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_real_showintableentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_real_showintableentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_real_showintitleentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_real_showintitleentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_real_showintitleentry, "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_real_showintitleentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_real_showintitleentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_real_showintitleentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_string_showinsubjectentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_string_showinsubjectentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_string_showinsubjectentry, "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_string_showinsubjectentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_string_showinsubjectentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_string_showinsubjectentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_string_showintableentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_string_showintableentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_string_showintableentry, "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_string_showintableentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_string_showintableentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_string_showintableentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_string_showintitleentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_string_showintitleentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_string_showintitleentry, "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_string_showintitleentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_string_showintitleentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_string_showintitleentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_xhtml_showinsubjectentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_xhtml_showinsubjectentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_xhtml_showinsubjectentry, "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_xhtml_showinsubjectentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_xhtml_showinsubjectentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_xhtml_showinsubjectentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_xhtml_showintableentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_xhtml_showintableentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_xhtml_showintableentry, "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_xhtml_showintableentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_xhtml_showintableentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_xhtml_showintableentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_attribute_definition_xhtml_showintitleentry := range set {
				nodeInstance := &tree.Node{Name: _map_attribute_definition_xhtml_showintitleentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_attribute_definition_xhtml_showintitleentry, "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_attribute_definition_xhtml_showintitleentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_attribute_definition_xhtml_showintitleentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_attribute_definition_xhtml_showintitleentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_SPECIFICATION_Nodes_expandedEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_SPECIFICATION_Nodes_expandedEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_specification_nodes_expandedentry := range set {
				nodeInstance := &tree.Node{Name: _map_specification_nodes_expandedentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_specification_nodes_expandedentry, "Map_SPECIFICATION_Nodes_expandedEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_specification_nodes_expandedentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_specification_nodes_expandedentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_specification_nodes_expandedentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_spec_object_type_isnodeexpandedentry := range set {
				nodeInstance := &tree.Node{Name: _map_spec_object_type_isnodeexpandedentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_spec_object_type_isnodeexpandedentry, "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_spec_object_type_isnodeexpandedentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_spec_object_type_isnodeexpandedentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_spec_object_type_isnodeexpandedentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_SPEC_OBJECT_TYPE_showIdentifierEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_spec_object_type_showidentifierentry := range set {
				nodeInstance := &tree.Node{Name: _map_spec_object_type_showidentifierentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_spec_object_type_showidentifierentry, "Map_SPEC_OBJECT_TYPE_showIdentifierEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_spec_object_type_showidentifierentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_spec_object_type_showidentifierentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_spec_object_type_showidentifierentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Map_SPEC_OBJECT_TYPE_showNameEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_SPEC_OBJECT_TYPE_showNameEntry](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _map_spec_object_type_shownameentry := range set {
				nodeInstance := &tree.Node{Name: _map_spec_object_type_shownameentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_map_spec_object_type_shownameentry, "Map_SPEC_OBJECT_TYPE_showNameEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_map_spec_object_type_shownameentry]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_map_spec_object_type_shownameentry]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_map_spec_object_type_shownameentry]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "RELATION_GROUP":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RELATION_GROUP](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _relation_group := range set {
				nodeInstance := &tree.Node{Name: _relation_group.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_relation_group, "RELATION_GROUP", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_relation_group]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_relation_group]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_relation_group]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "RELATION_GROUP_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RELATION_GROUP_TYPE](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _relation_group_type := range set {
				nodeInstance := &tree.Node{Name: _relation_group_type.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_relation_group_type, "RELATION_GROUP_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_relation_group_type]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_relation_group_type]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_relation_group_type]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "REQ_IF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.REQ_IF](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _req_if := range set {
				nodeInstance := &tree.Node{Name: _req_if.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if, "REQ_IF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_req_if]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_req_if]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_req_if]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "REQ_IF_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.REQ_IF_CONTENT](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _req_if_content := range set {
				nodeInstance := &tree.Node{Name: _req_if_content.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_content, "REQ_IF_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_req_if_content]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_req_if_content]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_req_if_content]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "REQ_IF_HEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.REQ_IF_HEADER](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _req_if_header := range set {
				nodeInstance := &tree.Node{Name: _req_if_header.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_header, "REQ_IF_HEADER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_req_if_header]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_req_if_header]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_req_if_header]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "REQ_IF_TOOL_EXTENSION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.REQ_IF_TOOL_EXTENSION](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _req_if_tool_extension := range set {
				nodeInstance := &tree.Node{Name: _req_if_tool_extension.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_tool_extension, "REQ_IF_TOOL_EXTENSION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_req_if_tool_extension]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_req_if_tool_extension]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_req_if_tool_extension]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "RenderingConfiguration":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RenderingConfiguration](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _renderingconfiguration := range set {
				nodeInstance := &tree.Node{Name: _renderingconfiguration.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_renderingconfiguration, "RenderingConfiguration", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_renderingconfiguration]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_renderingconfiguration]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_renderingconfiguration]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "SPECIFICATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _specification := range set {
				nodeInstance := &tree.Node{Name: _specification.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specification, "SPECIFICATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_specification]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_specification]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_specification]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "SPECIFICATION_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION_TYPE](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _specification_type := range set {
				nodeInstance := &tree.Node{Name: _specification_type.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specification_type, "SPECIFICATION_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_specification_type]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_specification_type]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_specification_type]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "SPEC_HIERARCHY":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_HIERARCHY](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _spec_hierarchy := range set {
				nodeInstance := &tree.Node{Name: _spec_hierarchy.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_hierarchy, "SPEC_HIERARCHY", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_spec_hierarchy]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_spec_hierarchy]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_spec_hierarchy]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "SPEC_OBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_OBJECT](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _spec_object := range set {
				nodeInstance := &tree.Node{Name: _spec_object.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_object, "SPEC_OBJECT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_spec_object]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_spec_object]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_spec_object]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "SPEC_OBJECT_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_OBJECT_TYPE](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _spec_object_type := range set {
				nodeInstance := &tree.Node{Name: _spec_object_type.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_object_type, "SPEC_OBJECT_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_spec_object_type]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_spec_object_type]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_spec_object_type]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "SPEC_RELATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_RELATION](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _spec_relation := range set {
				nodeInstance := &tree.Node{Name: _spec_relation.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_relation, "SPEC_RELATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_spec_relation]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_spec_relation]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_spec_relation]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "SPEC_RELATION_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_RELATION_TYPE](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _spec_relation_type := range set {
				nodeInstance := &tree.Node{Name: _spec_relation_type.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_relation_type, "SPEC_RELATION_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_spec_relation_type]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_spec_relation_type]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_spec_relation_type]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "StaticWebSite":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StaticWebSite](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _staticwebsite := range set {
				nodeInstance := &tree.Node{Name: _staticwebsite.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staticwebsite, "StaticWebSite", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_staticwebsite]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_staticwebsite]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_staticwebsite]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "StaticWebSiteChapter":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StaticWebSiteChapter](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _staticwebsitechapter := range set {
				nodeInstance := &tree.Node{Name: _staticwebsitechapter.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staticwebsitechapter, "StaticWebSiteChapter", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_staticwebsitechapter]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_staticwebsitechapter]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_staticwebsitechapter]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "StaticWebSiteGeneratedImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StaticWebSiteGeneratedImage](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _staticwebsitegeneratedimage := range set {
				nodeInstance := &tree.Node{Name: _staticwebsitegeneratedimage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staticwebsitegeneratedimage, "StaticWebSiteGeneratedImage", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_staticwebsitegeneratedimage]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_staticwebsitegeneratedimage]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_staticwebsitegeneratedimage]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "StaticWebSiteImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StaticWebSiteImage](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _staticwebsiteimage := range set {
				nodeInstance := &tree.Node{Name: _staticwebsiteimage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staticwebsiteimage, "StaticWebSiteImage", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_staticwebsiteimage]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_staticwebsiteimage]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_staticwebsiteimage]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "StaticWebSiteParagraph":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StaticWebSiteParagraph](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _staticwebsiteparagraph := range set {
				nodeInstance := &tree.Node{Name: _staticwebsiteparagraph.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_staticwebsiteparagraph, "StaticWebSiteParagraph", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_staticwebsiteparagraph]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_staticwebsiteparagraph]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_staticwebsiteparagraph]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "XHTML_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.XHTML_CONTENT](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _xhtml_content := range set {
				nodeInstance := &tree.Node{Name: _xhtml_content.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_content, "XHTML_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_xhtml_content]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_xhtml_content]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_xhtml_content]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
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

type InstanceNodeCallback[T models.PointerToGongstruct] struct {
	Instance       T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.PointerToGongstruct](
	instance T,
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
