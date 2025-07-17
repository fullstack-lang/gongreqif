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
				Name: spectype.Name + fmt.Sprintf(" (%d)", map_specType_nbInstance[spectype]),
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
				Name: spectype.Name + fmt.Sprintf(" (%d)", map_specType_nbInstance[spectype]),
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

	addAttributeNodeExperimental(
		specAttributes.ATTRIBUTE_DEFINITION_XHTML,
		stager.Map_id_ATTRIBUTE_DEFINITION_XHTML,
		stager.Map_ATTRIBUTE_DEFINITION_XHTML_Spec_nbInstance,
		stager,
		nodeSpecType)

	// String Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_STRING) > 0 {
		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_STRING]int)
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_STRING_REF](stager.GetStage()) {

			attributeDefinition, ok := stager.Map_id_ATTRIBUTE_DEFINITION_STRING[x.ATTRIBUTE_DEFINITION_STRING_REF]
			if !ok {
				log.Panic("x.ATTRIBUTE_DEFINITION_STRING_REF", x.ATTRIBUTE_DEFINITION_STRING_REF,
					"unknown ref")
			} else {
				map_attributeDefinition_nbInstance[attributeDefinition]++
			}
		}

		for _, attributeDefinition := range specAttributes.ATTRIBUTE_DEFINITION_STRING {
			var attributeType string
			if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_STRING[attributeDefinition.TYPE.DATATYPE_DEFINITION_STRING_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_STRING_REF", attributeDefinition.TYPE.DATATYPE_DEFINITION_STRING_REF,
					"unknown ref")
			}

			nbInstances := stager.Map_ATTRIBUTE_DEFINITION_STRING_Spec_nbInstance[attributeDefinition]
			nodeAttribute := &tree.Node{
				Name: attributeDefinition.LONG_NAME + ":" + attributeType +
					fmt.Sprintf(" (%d/%d)", nbInstances, map_attributeDefinition_nbInstance[attributeDefinition]),
			}
			configureAndAddAttributeNode(nodeSpecType, nodeAttribute, nbInstances, attributeDefinition.IS_EDITABLE, attributeDefinition.LONG_NAME)
		}
	}

	// Boolean Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_BOOLEAN) > 0 {
		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_BOOLEAN]int)
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](stager.GetStage()) {

			attributeDefinition, ok := stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN[x.ATTRIBUTE_DEFINITION_BOOLEAN_REF]
			if !ok {
				log.Panic("x.ATTRIBUTE_DEFINITION_BOOLEAN_REF", x.ATTRIBUTE_DEFINITION_BOOLEAN_REF,
					"unknown ref")
			} else {
				map_attributeDefinition_nbInstance[attributeDefinition]++
			}
		}

		for _, attributeDefinition := range specAttributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			var attributeType string
			if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_BOOLEAN[attributeDefinition.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF", attributeDefinition.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF,
					"unknown ref")
			}

			nbInstances := stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_Spec_nbInstance[attributeDefinition]
			nodeAttribute := &tree.Node{
				Name: attributeDefinition.LONG_NAME + ":" + attributeType +
					fmt.Sprintf(" (%d/%d)", nbInstances, map_attributeDefinition_nbInstance[attributeDefinition]),
			}
			configureAndAddAttributeNode(nodeSpecType, nodeAttribute, nbInstances, attributeDefinition.IS_EDITABLE, attributeDefinition.LONG_NAME)
		}
	}

	// Integer Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_INTEGER) > 0 {
		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_INTEGER]int)
		var attributeDefinition *m.ATTRIBUTE_DEFINITION_INTEGER
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_INTEGER_REF](stager.GetStage()) {

			var ok bool
			attributeDefinition, ok = stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER[x.ATTRIBUTE_DEFINITION_INTEGER_REF]
			if !ok {
				log.Panic("x.ATTRIBUTE_DEFINITION_INTEGER_REF", x.ATTRIBUTE_DEFINITION_INTEGER_REF,
					"unknown ref")
			} else {
				map_attributeDefinition_nbInstance[attributeDefinition]++
			}
		}

		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_INTEGER {
			var attributeType string
			if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_INTEGER[attribute.TYPE.DATATYPE_DEFINITION_INTEGER_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_INTEGER_REF", attribute.TYPE.DATATYPE_DEFINITION_INTEGER_REF,
					"unknown ref")
			}

			nbInstances := stager.Map_ATTRIBUTE_DEFINITION_INTEGER_Spec_nbInstance[attribute]
			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType +
					fmt.Sprintf(" (%d/%d)", nbInstances, map_attributeDefinition_nbInstance[attribute]),
			}
			configureAndAddAttributeNode(nodeSpecType, nodeAttribute, nbInstances, attribute.IS_EDITABLE, attribute.LONG_NAME)
		}
	}

	// Real Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_REAL) > 0 {
		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_REAL]int)
		var attributeDefinition *m.ATTRIBUTE_DEFINITION_REAL
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_REAL_REF](stager.GetStage()) {

			var ok bool
			attributeDefinition, ok = stager.Map_id_ATTRIBUTE_DEFINITION_REAL[x.ATTRIBUTE_DEFINITION_REAL_REF]
			if !ok {
				log.Panic("x.ATTRIBUTE_DEFINITION_REAL_REF", x.ATTRIBUTE_DEFINITION_REAL_REF,
					"unknown ref")
			} else {
				map_attributeDefinition_nbInstance[attributeDefinition]++
			}
		}

		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_REAL {
			var attributeType string
			if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_REAL[attribute.TYPE.DATATYPE_DEFINITION_REAL_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_REAL_REF", attribute.TYPE.DATATYPE_DEFINITION_REAL_REF,
					"unknown ref")
			}

			nbInstances := stager.Map_ATTRIBUTE_DEFINITION_REAL_Spec_nbInstance[attribute]
			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType +
					fmt.Sprintf(" (%d/%d)", nbInstances, map_attributeDefinition_nbInstance[attribute]),
			}
			configureAndAddAttributeNode(nodeSpecType, nodeAttribute, nbInstances, attribute.IS_EDITABLE, attribute.LONG_NAME)
		}
	}

	// Date Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_DATE) > 0 {
		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_DATE]int)
		var attributeDefinition *m.ATTRIBUTE_DEFINITION_DATE
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_DATE_REF](stager.GetStage()) {

			var ok bool
			attributeDefinition, ok = stager.Map_id_ATTRIBUTE_DEFINITION_DATE[x.ATTRIBUTE_DEFINITION_DATE_REF]
			if !ok {
				log.Panic("x.ATTRIBUTE_DEFINITION_DATE_REF", x.ATTRIBUTE_DEFINITION_DATE_REF,
					"unknown ref")
			} else {
				map_attributeDefinition_nbInstance[attributeDefinition]++
			}
		}

		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_DATE {
			var attributeType string
			if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_DATE[attribute.TYPE.DATATYPE_DEFINITION_DATE_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_DATE_REF", attribute.TYPE.DATATYPE_DEFINITION_DATE_REF,
					"unknown ref")
			}

			nbInstances := stager.Map_ATTRIBUTE_DEFINITION_DATE_Spec_nbInstance[attribute]
			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType +
					fmt.Sprintf(" (%d/%d)", nbInstances, map_attributeDefinition_nbInstance[attribute]),
			}
			configureAndAddAttributeNode(nodeSpecType, nodeAttribute, nbInstances, attribute.IS_EDITABLE, attribute.LONG_NAME)
		}
	}

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
			configureAndAddAttributeNode(nodeSpecType, nodeAttribute, nbInstances, attribute.IS_EDITABLE, attribute.LONG_NAME)
		}
	}
}

func addAttributeNodeExperimental[T m.Attribute](
	attributes []T,
	map_Id_Attribute map[string]T,
	map_Atttribute_nbInstance map[T]int,
	stager *m.Stager, nodeSpecType *tree.Node) {
	if len(attributes) > 0 {
		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[T]int)
		var attributeDefinition T
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_XHTML_REF](stager.GetStage()) {

			var ok bool
			attributeDefinition, ok = map_Id_Attribute[x.ATTRIBUTE_DEFINITION_XHTML_REF]
			if !ok {
				log.Panic("x.ATTRIBUTE_DEFINITION_XHTML_REF", x.ATTRIBUTE_DEFINITION_XHTML_REF,
					"unknown ref")
			} else {
				map_attributeDefinition_nbInstance[attributeDefinition]++
			}
		}

		for _, attribute := range attributes {
			var attributeType string
			if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_XHTML[attribute.GetTypeRef()]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.GetTypeRef()", attribute.GetTypeRef(),
					"unknown ref")
			}

			nbInstances := map_Atttribute_nbInstance[attribute]
			nodeAttribute := &tree.Node{
				Name: attribute.GetLongName() + ":" + attributeType +
					fmt.Sprintf(" (%d/%d)", nbInstances, map_attributeDefinition_nbInstance[attribute]),
			}
			configureAndAddAttributeNode(nodeSpecType, nodeAttribute, nbInstances, attribute.GetIsEditable(), attribute.GetLongName())
		}
	}
}

// configureAndAddAttributeNode configures a new attribute node, sets its icon,
// and adds it to the parent node in the tree.
func configureAndAddAttributeNode(
	nodeSpecType *tree.Node,
	nodeAttribute *tree.Node,
	nbInstances int,
	isEditable bool,
	longName string,
) {
	if nbInstances > 0 {
		nodeAttribute.IsWithPreceedingIcon = true
		nodeAttribute.PreceedingIcon = string(buttons.BUTTON_check_circle)
	}
	nodeSpecType.Children = append(nodeSpecType.Children, nodeAttribute)
	m.AddIconForEditabilityOfAttribute(isEditable, longName, nodeAttribute)

	nodeAttribute.Buttons = append(nodeAttribute.Buttons,
		&tree.Button{
			SVGIcon: svgIconTitle,
		},
	)
}
