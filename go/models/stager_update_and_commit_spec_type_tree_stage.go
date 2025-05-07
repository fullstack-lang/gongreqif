package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) updateAndCommit_spec_type_tree_stage() {

	stager.specTypeTreeStage.Reset()

	spectypes := stager.rootReqif.CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES

	datatypes_xhtml := *GetGongstructInstancesSet[DATATYPE_DEFINITION_XHTML](stager.stage)
	datatypes_xhtml_id_map := make(map[string]*DATATYPE_DEFINITION_XHTML)
	for datatype_xhtml := range datatypes_xhtml {
		datatypes_xhtml_id_map[datatype_xhtml.IDENTIFIER] = datatype_xhtml
	}

	datatypes_enumeration := *GetGongstructInstancesSet[DATATYPE_DEFINITION_ENUMERATION](stager.stage)
	datatypes_enumeration_id_map := make(map[string]*DATATYPE_DEFINITION_ENUMERATION)
	for datatype_enumeration := range datatypes_enumeration {
		datatypes_enumeration_id_map[datatype_enumeration.IDENTIFIER] = datatype_enumeration
	}

	rootNode := &tree.Node{
		Name:       "Spec types",
		IsExpanded: true,
	}

	{
		spectypeCategory := &tree.Node{
			Name:       "Spec Object Types",
			IsExpanded: true,
		}
		rootNode.Children = append(rootNode.Children, spectypeCategory)
		for _, spectype := range spectypes.SPEC_OBJECT_TYPE {
			node := &tree.Node{
				Name: spectype.LONG_NAME,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, node)

			if spectype.SPEC_ATTRIBUTES == nil {
				continue
			}
			for _, attribute := range spectype.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				// provide the type
				var attributeType string
				if datatype, ok := datatypes_xhtml_id_map[attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF]; ok {
					attributeType = datatype.LONG_NAME
				}

				nodeAttribute := &tree.Node{
					Name: attribute.LONG_NAME + ":" + attributeType,
				}
				node.Children = append(node.Children, nodeAttribute)
			}
		}
	}

	{
		spectypeCategory := &tree.Node{
			Name:       "Relation Types",
			IsExpanded: true,
		}
		rootNode.Children = append(rootNode.Children, spectypeCategory)
		for _, spectype := range spectypes.SPEC_RELATION_TYPE {
			node := &tree.Node{
				Name: spectype.LONG_NAME,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, node)

			if spectype.SPEC_ATTRIBUTES == nil {
				continue
			}
			for _, attribute := range spectype.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				// provide the type
				var attributeType string
				if datatype, ok := datatypes_xhtml_id_map[attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF]; ok {
					attributeType = datatype.LONG_NAME
				}

				nodeAttribute := &tree.Node{
					Name: attribute.LONG_NAME + ":" + attributeType,
				}
				node.Children = append(node.Children, nodeAttribute)
			}
		}
	}

	{
		spectypeCategory := &tree.Node{
			Name:       "Specification Types",
			IsExpanded: true,
		}
		rootNode.Children = append(rootNode.Children, spectypeCategory)
		for _, spectype := range spectypes.SPECIFICATION_TYPE {
			node := &tree.Node{
				Name: spectype.LONG_NAME,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, node)

			if spectype.SPEC_ATTRIBUTES == nil {
				continue
			}
			for _, attribute := range spectype.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				// provide the type
				var attributeType string
				if datatype, ok := datatypes_xhtml_id_map[attribute.TYPE.DATATYPE_DEFINITION_XHTML_REF]; ok {
					attributeType = datatype.LONG_NAME
				}

				nodeAttribute := &tree.Node{
					Name: attribute.LONG_NAME + ":" + attributeType,
				}
				node.Children = append(node.Children, nodeAttribute)
			}
			for _, attribute := range spectype.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
				// provide the type
				var attributeType string
				if datatype, ok := datatypes_enumeration_id_map[attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF]; ok {
					attributeType = datatype.LONG_NAME
				}

				nodeAttribute := &tree.Node{
					Name: attribute.LONG_NAME + ":" + attributeType,
				}
				node.Children = append(node.Children, nodeAttribute)
			}
		}
	}

	tree.StageBranch(stager.specTypeTreeStage,
		&tree.Tree{
			Name: stager.specTypeTreeName,
			RootNodes: []*tree.Node{
				rootNode,
			},
		},
	)

	stager.specTypeTreeStage.Commit()
}
