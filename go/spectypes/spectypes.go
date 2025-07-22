package spectypes

import (
	"fmt"
	"log"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

type SpecTypesTreeStageUpdater struct {
}

func (updater *SpecTypesTreeStageUpdater) UpdateAndCommitSpecTypesTreeStage(stager *m.Stager) {

	stager.GetSpecTypesTreeStage().Reset()

	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES

	rootNode := &tree.Node{
		Name:       "Spec types",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}

	{
		spectypeCategory := &tree.Node{
			Name:       "Spec Object Types",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, spectypeCategory)

		// compute the number of time this spec object type is used
		map_specType_nbInstance := make(map[*m.SPEC_OBJECT_TYPE]int)
		for x := range *m.GetGongstructInstancesSet[m.A_SPEC_OBJECT_TYPE_REF](stager.GetStage()) {

			specObjecType, ok := stager.Map_id_SPEC_OBJECT_TYPE[x.SPEC_OBJECT_TYPE_REF]
			if !ok {
				log.Panicf("Object %s has a x.SPEC_OBJECT_TYPE_REF %s which is not known", x.Name, x.SPEC_OBJECT_TYPE_REF)
			} else {
				map_specType_nbInstance[specObjecType]++
			}
		}

		for _, specObjectType := range spectypes.SPEC_OBJECT_TYPE {
			specObjectTypeNode := &tree.Node{
				Name: specObjectType.Name + fmt.Sprintf(" (%d/%d)",
					stager.Map_SPEC_OBJECT_TYPE_Spec_nbInstance[specObjectType],
					map_specType_nbInstance[specObjectType]),
				IsExpanded: true,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, specObjectTypeNode)

			{
				button := &tree.Button{
					Name: "Show/Unshow identifier",
					Impl: &ButtonToggleShowSpecObjectTypeIdentifierProxy{
						stager:         stager,
						specObjectType: specObjectType,
					},
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if !stager.Map_SPEC_OBJECT_TYPE_showIdentifier[specObjectType] {
					button.ToolTipText = "Show identifier in title"
					button.SVGIcon = svgIconBadge
				} else {
					button.ToolTipText = "Hide identifier in title"
					button.SVGIcon = svgIconBadgeOff
				}
				specObjectTypeNode.Buttons = append(specObjectTypeNode.Buttons, button)
			}

			{
				button := &tree.Button{
					Name: "Show/Unshow long name",
					Impl: &ButtonToggleShowSpecObjectTypeLongNameProxy{
						stager:         stager,
						specObjectType: specObjectType,
					},
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if !stager.Map_SPEC_OBJECT_TYPE_showName[specObjectType] {
					button.ToolTipText = "Show name in title"
					button.SVGIcon = svgIconBadge
				} else {
					button.ToolTipText = "Hide name in title"
					button.SVGIcon = svgIconBadgeOff
				}
				specObjectTypeNode.Buttons = append(specObjectTypeNode.Buttons, button)
			}

			addAttibutesNodes(
				stager,
				specObjectTypeNode,
				specObjectType.SPEC_ATTRIBUTES)

		}
	}

	{
		spectypeCategory := &tree.Node{
			Name:       "Relation Types",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, spectypeCategory)

		// compute the number of time this spec object type is used
		map_specType_nbInstance := make(map[*m.SPEC_RELATION_TYPE]int)
		for x := range *m.GetGongstructInstancesSet[m.A_SPEC_RELATION_TYPE_REF](stager.GetStage()) {

			datatypeDefinition, ok := stager.Map_id_SPEC_RELATION_TYPE[x.SPEC_RELATION_TYPE_REF]
			if !ok {
				log.Panic("x.SPEC_RELATION_TYPE_REF", x.SPEC_RELATION_TYPE_REF,
					"unknown ref")
			} else {
				map_specType_nbInstance[datatypeDefinition]++
			}
		}

		for _, spectype := range spectypes.SPEC_RELATION_TYPE {
			node := &tree.Node{
				Name:       spectype.Name + fmt.Sprintf(" (%d)", map_specType_nbInstance[spectype]),
				IsExpanded: true,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, node)

			addAttibutesNodes(
				stager,
				node,
				spectype.SPEC_ATTRIBUTES)
		}
	}

	{
		spectypeCategory := &tree.Node{
			Name:       "Specification Types",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, spectypeCategory)

		// compute the number of time this spec object type is used
		map_specType_nbInstance := make(map[*m.SPECIFICATION_TYPE]int)
		for x := range *m.GetGongstructInstancesSet[m.A_SPECIFICATION_TYPE_REF](stager.GetStage()) {

			specificationType, ok := stager.Map_id_SPECIFICATION_TYPE[x.SPECIFICATION_TYPE_REF]
			if !ok {
				log.Panic("x.SPECIFICATION_TYPE_REF: x: ", x.Name, " ,Type Ref: ", x.SPECIFICATION_TYPE_REF,
					", unknown ref")
			} else {
				map_specType_nbInstance[specificationType]++
			}
		}

		for _, spectype := range spectypes.SPECIFICATION_TYPE {
			node := &tree.Node{
				Name: spectype.Name + fmt.Sprintf(" (%d/%d)",
					stager.Map_SPECIFICATION_TYPE_Spec_nbInstance[spectype],
					map_specType_nbInstance[spectype]),
				IsExpanded: true,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, node)

			if spectype.SPEC_ATTRIBUTES == nil {
				continue
			}
			addAttibutesNodes(
				stager,
				node,
				spectype.SPEC_ATTRIBUTES)
		}
	}

	tree.StageBranch(stager.GetSpecTypesTreeStage(),
		&tree.Tree{
			Name: stager.GetSpecTypesTreeName(),
			RootNodes: []*tree.Node{
				rootNode,
			},
		},
	)

	stager.GetSpecTypesTreeStage().Commit()
}

// addAttibutesNodes creates category nodes for attribute definitions
// only if there is more than one attribute definition for that category.
func addAttibutesNodes(
	stager *m.Stager,
	nodeSpecType *tree.Node,
	specAttributes *m.A_SPEC_ATTRIBUTES) {

	if specAttributes == nil {
		return
	}

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_XHTML,
		stager.Map_id_ATTRIBUTE_DEFINITION_XHTML,
		stager.Map_ATTRIBUTE_DEFINITION_XHTML_Spec_nbInstance,
		stager.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTile,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_XHTML_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_XHTML,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_STRING,
		stager.Map_id_ATTRIBUTE_DEFINITION_STRING,
		stager.Map_ATTRIBUTE_DEFINITION_STRING_Spec_nbInstance,
		stager.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitle,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_STRING_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_STRING,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_BOOLEAN,
		stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN,
		stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_Spec_nbInstance,
		stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_BOOLEAN,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_INTEGER,
		stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER,
		stager.Map_ATTRIBUTE_DEFINITION_INTEGER_Spec_nbInstance,
		stager.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_INTEGER_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_INTEGER,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_REAL,
		stager.Map_id_ATTRIBUTE_DEFINITION_REAL,
		stager.Map_ATTRIBUTE_DEFINITION_REAL_Spec_nbInstance,
		stager.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitle,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_REAL_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_REAL,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_DATE,
		stager.Map_id_ATTRIBUTE_DEFINITION_DATE,
		stager.Map_ATTRIBUTE_DEFINITION_DATE_Spec_nbInstance,
		stager.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitle,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_DATE_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_DATE,
		nodeSpecType)

	// ENUMERATION Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_ENUMERATION) > 0 {
		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_ENUMERATION]int)
		var attributeDefinition *m.ATTRIBUTE_DEFINITION_ENUMERATION
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](stager.GetStage()) {

			var ok bool
			attributeDefinition, ok = stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION[x.ATTRIBUTE_DEFINITION_ENUMERATION_REF]
			if !ok {
				log.Panic("x.ATTRIBUTE_DEFINITION_ENUMERATION_REF", x.ATTRIBUTE_DEFINITION_ENUMERATION_REF,
					"unknown ref")
			} else {
				map_attributeDefinition_nbInstance[attributeDefinition]++
			}
		}

		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			var attributeType string
			if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_ENUMERATION[attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF", attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF,
					"unknown ref")
			}

			nbInstances := stager.Map_ATTRIBUTE_DEFINITION_ENUMERATION_Spec_nbInstance[attribute]
			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + " : " + attributeType +
					fmt.Sprintf(" (%d/%d)", nbInstances, map_attributeDefinition_nbInstance[attribute]),
			}
			configureAndAddAttributeNode(
				stager,
				nodeSpecType,
				nodeAttribute,
				nbInstances,
				attribute.IS_EDITABLE,
				attribute.LONG_NAME,
				attributeDefinition,
				stager.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle)
		}
	}
}

// configureAndAddAttributeNode configures a new attribute node, sets its icon,
// and adds it to the parent node in the tree.
func configureAndAddAttributeNode[AttrDef m.AttributeDefinition](
	stager *m.Stager,
	nodeSpecType *tree.Node,
	nodeAttribute *tree.Node,
	nbInstances int,
	isEditable bool,
	longName string,
	attributeDefinition AttrDef,
	map_AtttributeDefinition_showInTitle map[AttrDef]bool,
) {
	if nbInstances > 0 {
		nodeAttribute.IsWithPreceedingIcon = true
		nodeAttribute.PreceedingIcon = string(buttons.BUTTON_check_circle)
	}
	nodeSpecType.Children = append(nodeSpecType.Children, nodeAttribute)
	m.AddIconForEditabilityOfAttribute(isEditable, longName, nodeAttribute)

	button := &tree.Button{
		Name: "button for showing off in title",
		Impl: &ButtonToggleShowAttributeFieldInTitleProxy[AttrDef]{
			stager:                               stager,
			attributeDefinition:                  attributeDefinition,
			map_AtttributeDefinition_showInTitle: map_AtttributeDefinition_showInTitle,
		},
	}

	if map_AtttributeDefinition_showInTitle[attributeDefinition] {
		button.SVGIcon = svgIconTitleOff
	} else {
		button.SVGIcon = svgIconTitle
	}
	nodeAttribute.Buttons = append(nodeAttribute.Buttons,
		button,
	)
}

type ButtonToggleShowAttributeFieldInTitleProxy[AttrDef m.AttributeDefinition] struct {
	stager                               *m.Stager
	attributeDefinition                  AttrDef
	map_AtttributeDefinition_showInTitle map[AttrDef]bool
}

func (proxy *ButtonToggleShowAttributeFieldInTitleProxy[AttrDef]) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	showInTitle, ok := proxy.map_AtttributeDefinition_showInTitle[proxy.attributeDefinition]

	if !ok {
		log.Fatalln("Unknown specificiation in map", proxy.attributeDefinition.GetName())
	}

	proxy.map_AtttributeDefinition_showInTitle[proxy.attributeDefinition] = !showInTitle

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
}
