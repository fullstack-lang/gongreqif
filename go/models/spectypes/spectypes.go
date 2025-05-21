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

	datatypes_xhtml := *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_XHTML](stager.GetStage())
	datatypes_xhtml_id_map := make(map[string]*m.DATATYPE_DEFINITION_XHTML)
	for datatype_xhtml := range datatypes_xhtml {
		datatypes_xhtml_id_map[datatype_xhtml.IDENTIFIER] = datatype_xhtml
	}

	datatypes_string := *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_STRING](stager.GetStage())
	datatypes_string_id_map := make(map[string]*m.DATATYPE_DEFINITION_STRING)
	for datatype_string := range datatypes_string {
		datatypes_string_id_map[datatype_string.IDENTIFIER] = datatype_string
	}

	datatypes_boolean := *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_BOOLEAN](stager.GetStage())
	datatypes_boolean_id_map := make(map[string]*m.DATATYPE_DEFINITION_BOOLEAN)
	for datatype_boolean := range datatypes_boolean {
		datatypes_boolean_id_map[datatype_boolean.IDENTIFIER] = datatype_boolean
	}

	datatypes_integer := *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_INTEGER](stager.GetStage())
	datatypes_integer_id_map := make(map[string]*m.DATATYPE_DEFINITION_INTEGER)
	for datatype_integer := range datatypes_integer {
		datatypes_integer_id_map[datatype_integer.IDENTIFIER] = datatype_integer
	}

	datatypes_real := *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_REAL](stager.GetStage())
	datatypes_real_id_map := make(map[string]*m.DATATYPE_DEFINITION_REAL)
	for datatype_real := range datatypes_real {
		datatypes_real_id_map[datatype_real.IDENTIFIER] = datatype_real
	}

	datatypes_date := *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_DATE](stager.GetStage())
	datatypes_date_id_map := make(map[string]*m.DATATYPE_DEFINITION_DATE)
	for datatype_date := range datatypes_date {
		datatypes_date_id_map[datatype_date.IDENTIFIER] = datatype_date
	}

	datatypes_enumeration := *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_ENUMERATION](stager.GetStage())
	datatypes_enumeration_id_map := make(map[string]*m.DATATYPE_DEFINITION_ENUMERATION)
	for datatype_enumeration := range datatypes_enumeration {
		datatypes_enumeration_id_map[datatype_enumeration.IDENTIFIER] = datatype_enumeration
	}

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
				node,
				spectype.SPEC_ATTRIBUTES,
				datatypes_xhtml_id_map,
				datatypes_string_id_map,
				datatypes_boolean_id_map,
				datatypes_integer_id_map,
				datatypes_real_id_map,
				datatypes_date_id_map,
				datatypes_enumeration_id_map)

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
				node,
				spectype.SPEC_ATTRIBUTES,
				datatypes_xhtml_id_map,
				datatypes_string_id_map,
				datatypes_boolean_id_map,
				datatypes_integer_id_map,
				datatypes_real_id_map,
				datatypes_date_id_map,
				datatypes_enumeration_id_map)
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
				node,
				spectype.SPEC_ATTRIBUTES,
				datatypes_xhtml_id_map,
				datatypes_string_id_map,
				datatypes_boolean_id_map,
				datatypes_integer_id_map,
				datatypes_real_id_map,
				datatypes_date_id_map,
				datatypes_enumeration_id_map)
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

func addAttibutesNotes(
	nodeSpecType *tree.Node,
	specAttributes *m.A_SPEC_ATTRIBUTES,
	datatypes_xhtml_id_map map[string]*m.DATATYPE_DEFINITION_XHTML,
	datatypes_string_id_map map[string]*m.DATATYPE_DEFINITION_STRING,
	datatypes_boolean_id_map map[string]*m.DATATYPE_DEFINITION_BOOLEAN,
	datatypes_integer_id_map map[string]*m.DATATYPE_DEFINITION_INTEGER,
	datatypes_real_id_map map[string]*m.DATATYPE_DEFINITION_REAL,
	datatypes_date_id_map map[string]*m.DATATYPE_DEFINITION_DATE,
	datatypes_enumeration_id_map map[string]*m.DATATYPE_DEFINITION_ENUMERATION) {
	{

		if specAttributes == nil {
			return
		}

		nodeSpecTypeAttribute := &tree.Node{
			Name:       "XHTML",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttribute)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_XHTML {
			// provide the type
			var attributeType string
			if datatype, ok := datatypes_xhtml_id_map[attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF", attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttribute.Children = append(nodeSpecTypeAttribute.Children, nodeAttribute)
		}
	}
	{
		nodeSpecTypeAttribute := &tree.Node{
			Name:       "String",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttribute)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_STRING {
			// provide the type
			var attributeType string
			if datatype, ok := datatypes_string_id_map[attribute.TYPE.DATATYPE_DEFINITION_STRING_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_STRING_REF", attribute.TYPE.DATATYPE_DEFINITION_STRING_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttribute.Children = append(nodeSpecTypeAttribute.Children, nodeAttribute)
		}
	}
	{
		nodeSpecTypeAttribute := &tree.Node{
			Name:       "Boolean",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttribute)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			// provide the type
			var attributeType string
			if datatype, ok := datatypes_boolean_id_map[attribute.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF", attribute.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttribute.Children = append(nodeSpecTypeAttribute.Children, nodeAttribute)
		}
	}
	{
		nodeSpecTypeAttribute := &tree.Node{
			Name:       "Integer",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttribute)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_INTEGER {
			// provide the type
			var attributeType string
			if datatype, ok := datatypes_integer_id_map[attribute.TYPE.DATATYPE_DEFINITION_INTEGER_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_INTEGER_REF", attribute.TYPE.DATATYPE_DEFINITION_INTEGER_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttribute.Children = append(nodeSpecTypeAttribute.Children, nodeAttribute)
		}
	}
	{
		nodeSpecTypeAttribute := &tree.Node{
			Name:       "Real",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttribute)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_REAL {
			// provide the type
			var attributeType string
			if datatype, ok := datatypes_real_id_map[attribute.TYPE.DATATYPE_DEFINITION_REAL_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_REAL_REF", attribute.TYPE.DATATYPE_DEFINITION_REAL_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttribute.Children = append(nodeSpecTypeAttribute.Children, nodeAttribute)
		}
	}
	{
		nodeSpecTypeAttribute := &tree.Node{
			Name:       "Date",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttribute)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_DATE {
			// provide the type
			var attributeType string
			if datatype, ok := datatypes_date_id_map[attribute.TYPE.DATATYPE_DEFINITION_DATE_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_DATE_REF", attribute.TYPE.DATATYPE_DEFINITION_DATE_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + ":" + attributeType,
			}
			nodeSpecTypeAttribute.Children = append(nodeSpecTypeAttribute.Children, nodeAttribute)
		}
	}
	{
		nodeSpecTypeAttribute := &tree.Node{
			Name:       "ENUMERATION",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttribute)
		for _, attribute := range specAttributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			// provide the type
			var attributeType string
			if datatype, ok := datatypes_enumeration_id_map[attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF", attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF,
					"unknown ref")
			}

			nodeAttribute := &tree.Node{
				Name: attribute.LONG_NAME + " : " + attributeType,
			}
			nodeSpecTypeAttribute.Children = append(nodeSpecTypeAttribute.Children, nodeAttribute)
		}
	}
}
