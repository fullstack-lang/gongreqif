package spectypes

import (
	"log"

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

		for _, spectype := range spectypes.SPEC_OBJECT_TYPE {
			node := &tree.Node{
				Name: spectype.Name,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, node)

			addAttibutesNotes(
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
		for _, spectype := range spectypes.SPEC_RELATION_TYPE {
			node := &tree.Node{
				Name: spectype.Name,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, node)

			addAttibutesNotes(
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
		for _, spectype := range spectypes.SPECIFICATION_TYPE {
			node := &tree.Node{
				Name: spectype.Name,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, node)

			if spectype.SPEC_ATTRIBUTES == nil {
				continue
			}
			addAttibutesNotes(
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

// addAttibutesNotes creates category nodes for attribute definitions
// only if there is more than one attribute definition for that category.
func addAttibutesNotes(
	stager *m.Stager,
	nodeSpecType *tree.Node,
	specAttributes *m.A_SPEC_ATTRIBUTES) {

	if specAttributes == nil {
		return
	}

	// XHTML Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_XHTML) > 1 {
		nodeSpecTypeAttributeCategory := &tree.Node{
			Name:       "XHTML",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_XHTML {
			var attributeType string
			if datatype, ok := stager.Map_id_datatypes_xhtml[attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF", attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttributeCategory.Children = append(nodeSpecTypeAttributeCategory.Children, nodeAttribute)
		}
	}

	// String Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_STRING) > 1 {
		nodeSpecTypeAttributeCategory := &tree.Node{
			Name:       "String",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_STRING {
			var attributeType string
			if datatype, ok := stager.Map_id_datatypes_string[attribute.TYPE.DATATYPE_DEFINITION_STRING_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_STRING_REF", attribute.TYPE.DATATYPE_DEFINITION_STRING_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttributeCategory.Children = append(nodeSpecTypeAttributeCategory.Children, nodeAttribute)
		}
	}

	// Boolean Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_BOOLEAN) > 1 {
		nodeSpecTypeAttributeCategory := &tree.Node{
			Name:       "Boolean",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			var attributeType string
			if datatype, ok := stager.Map_id_datatypes_boolean[attribute.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF", attribute.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttributeCategory.Children = append(nodeSpecTypeAttributeCategory.Children, nodeAttribute)
		}
	}

	// Integer Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_INTEGER) > 1 {
		nodeSpecTypeAttributeCategory := &tree.Node{
			Name:       "Integer",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_INTEGER {
			var attributeType string
			if datatype, ok := stager.Map_id_datatypes_integer[attribute.TYPE.DATATYPE_DEFINITION_INTEGER_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_INTEGER_REF", attribute.TYPE.DATATYPE_DEFINITION_INTEGER_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttributeCategory.Children = append(nodeSpecTypeAttributeCategory.Children, nodeAttribute)
		}
	}

	// Real Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_REAL) > 1 {
		nodeSpecTypeAttributeCategory := &tree.Node{
			Name:       "Real",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_REAL {
			var attributeType string
			if datatype, ok := stager.Map_id_datatypes_real[attribute.TYPE.DATATYPE_DEFINITION_REAL_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_REAL_REF", attribute.TYPE.DATATYPE_DEFINITION_REAL_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttributeCategory.Children = append(nodeSpecTypeAttributeCategory.Children, nodeAttribute)
		}
	}

	// Date Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_DATE) > 1 {
		nodeSpecTypeAttributeCategory := &tree.Node{
			Name:       "Date",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_DATE {
			var attributeType string
			if datatype, ok := stager.Map_id_datatypes_date[attribute.TYPE.DATATYPE_DEFINITION_DATE_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_DATE_REF", attribute.TYPE.DATATYPE_DEFINITION_DATE_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttributeCategory.Children = append(nodeSpecTypeAttributeCategory.Children, nodeAttribute)
		}
	}

	// ENUMERATION Attributes
	if len(specAttributes.ATTRIBUTE_DEFINITION_ENUMERATION) > 1 {
		nodeSpecTypeAttributeCategory := &tree.Node{
			Name:       "ENUMERATION",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttributeCategory)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			var attributeType string
			if datatype, ok := stager.Map_id_datatypes_enumeration[attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF", attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + " : " + attributeType, // Note: Original had " : " here, kept for consistency
			}
			nodeSpecTypeAttributeCategory.Children = append(nodeSpecTypeAttributeCategory.Children, nodeAttribute)
		}
	}
}
