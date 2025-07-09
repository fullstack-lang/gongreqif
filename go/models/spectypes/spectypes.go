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

		for _, spectype := range spectypes.SPEC_OBJECT_TYPE {
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

	// XHTML Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_XHTML) > 0 {
		// nodeSpecTypeAttributeCategory := &tree.Node{
		// 	Name:       "XHTML",
		// 	IsExpanded: true,
		// 	FontStyle:  tree.ITALIC,
		// }
		// nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)

		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_XHTML]int)
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_XHTML_REF](stager.GetStage()) {

			attributeDefinition, ok := stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[x.ATTRIBUTE_DEFINITION_XHTML_REF]
			if !ok {
				log.Panic("x.ATTRIBUTE_DEFINITION_XHTML_REF", x.ATTRIBUTE_DEFINITION_XHTML_REF,
					"unknown ref")
			} else {
				map_attributeDefinition_nbInstance[attributeDefinition]++
			}
		}

		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_XHTML {
			var attributeType string
			if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_XHTML[attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF", attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType + fmt.Sprintf(" (%d)", map_attributeDefinition_nbInstance[attribute]),
			}
			nodeSpecType.Children = append(nodeSpecType.Children, nodeAttribute)
			m.AddIconForEditabilityOfAttribute(attribute.IS_EDITABLE, attribute.LONG_NAME, nodeAttribute)
		}
	}

	// String Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_STRING) > 0 {
		// nodeSpecTypeAttributeCategory := &tree.Node{
		// 	Name:       "String",
		// 	IsExpanded: true,
		// 	FontStyle:  tree.ITALIC,
		// }
		// nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)

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
					fmt.Sprintf(" (%d/%d)", nbInstances,
						map_attributeDefinition_nbInstance[attributeDefinition]),
			}
			if nbInstances > 0 {
				nodeAttribute.IsWithPreceedingIcon = true
				nodeAttribute.PreceedingIcon = string(buttons.BUTTON_check_circle)
			}
			nodeSpecType.Children = append(nodeSpecType.Children, nodeAttribute)
			m.AddIconForEditabilityOfAttribute(attributeDefinition.IS_EDITABLE, attributeDefinition.LONG_NAME, nodeAttribute)
		}
	}

	// Boolean Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_BOOLEAN) > 0 {
		// nodeSpecTypeAttributeCategory := &tree.Node{
		// 	Name:       "Boolean",
		// 	IsExpanded: true,
		// 	FontStyle:  tree.ITALIC,
		// }
		// nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)

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

		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			var attributeType string
			if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_BOOLEAN[attribute.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF", attribute.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType + fmt.Sprintf(" (%d)", map_attributeDefinition_nbInstance[attribute]),
			}
			nodeSpecType.Children = append(nodeSpecType.Children, nodeAttribute)
			m.AddIconForEditabilityOfAttribute(attribute.IS_EDITABLE, attribute.LONG_NAME, nodeAttribute)
		}
	}

	// Integer Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_INTEGER) > 0 {
		// nodeSpecTypeAttributeCategory := &tree.Node{
		// 	Name:       "Integer",
		// 	IsExpanded: true,
		// 	FontStyle:  tree.ITALIC,
		// }
		// nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)

		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_INTEGER]int)
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_INTEGER_REF](stager.GetStage()) {

			attributeDefinition, ok := stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER[x.ATTRIBUTE_DEFINITION_INTEGER_REF]
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

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType + fmt.Sprintf(" (%d)", map_attributeDefinition_nbInstance[attribute]),
			}
			nodeSpecType.Children = append(nodeSpecType.Children, nodeAttribute)
			m.AddIconForEditabilityOfAttribute(attribute.IS_EDITABLE, attribute.LONG_NAME, nodeAttribute)
		}
	}

	// Real Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_REAL) > 0 {
		// nodeSpecTypeAttributeCategory := &tree.Node{
		// 	Name:       "Real",
		// 	IsExpanded: true,
		// 	FontStyle:  tree.ITALIC,
		// }
		// nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)

		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_REAL]int)
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_REAL_REF](stager.GetStage()) {

			attributeDefinition, ok := stager.Map_id_ATTRIBUTE_DEFINITION_REAL[x.ATTRIBUTE_DEFINITION_REAL_REF]
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

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType + fmt.Sprintf(" (%d)", map_attributeDefinition_nbInstance[attribute]),
			}
			nodeSpecType.Children = append(nodeSpecType.Children, nodeAttribute)
			m.AddIconForEditabilityOfAttribute(attribute.IS_EDITABLE, attribute.LONG_NAME, nodeAttribute)
		}
	}

	// Date Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_DATE) > 0 {
		// nodeSpecTypeAttributeCategory := &tree.Node{
		// 	Name:       "Date",
		// 	IsExpanded: true,
		// 	FontStyle:  tree.ITALIC,
		// }
		// nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)

		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_DATE]int)
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_DATE_REF](stager.GetStage()) {

			attributeDefinition, ok := stager.Map_id_ATTRIBUTE_DEFINITION_DATE[x.ATTRIBUTE_DEFINITION_DATE_REF]
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

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType + fmt.Sprintf(" (%d)", map_attributeDefinition_nbInstance[attribute]),
			}
			nodeSpecType.Children = append(nodeSpecType.Children, nodeAttribute)
			m.AddIconForEditabilityOfAttribute(attribute.IS_EDITABLE, attribute.LONG_NAME, nodeAttribute)
		}
	}

	// ENUMERATION Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_ENUMERATION) > 0 {
		// nodeSpecTypeAttributeCategory := &tree.Node{
		// 	Name:       "ENUMERATION",
		// 	IsExpanded: true,
		// 	FontStyle:  tree.ITALIC,
		// }
		// nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)

		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[*m.ATTRIBUTE_DEFINITION_ENUMERATION]int)
		for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](stager.GetStage()) {

			attributeDefinition, ok := stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION[x.ATTRIBUTE_DEFINITION_ENUMERATION_REF]
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

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + " : " + attributeType +
					fmt.Sprintf(" (%d)", map_attributeDefinition_nbInstance[attribute]),
			}

			nodeSpecType.Children = append(nodeSpecType.Children, nodeAttribute)
			m.AddIconForEditabilityOfAttribute(attribute.IS_EDITABLE, attribute.LONG_NAME, nodeAttribute)
		}
	}
}
