package specobjects

import (
	"log"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

type SpecObjectsTreeStageUpdater struct {
}

func (spectTreeStageUpdater *SpecObjectsTreeStageUpdater) UpdateAndCommitSpecTreeStage(stager *m.Stager) {

	stager.GetSpecTreeStage().Reset()

	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES

	datatypes_xhtml := *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_XHTML](stager.GetStage())
	datatypes_xhtml_id_map := make(map[string]*m.DATATYPE_DEFINITION_XHTML)
	for datatype_xhtml := range datatypes_xhtml {
		datatypes_xhtml_id_map[datatype_xhtml.IDENTIFIER] = datatype_xhtml
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
			nodeSpecType := &tree.Node{
				Name: spectype.Name,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, nodeSpecType)

			if spectype.SPEC_ATTRIBUTES == nil {
				continue
			}

			{
				nodeSpecTypeAttribute := &tree.Node{
					Name:       "XHTML",
					IsExpanded: true,
					FontStyle:  tree.ITALIC,
				}
				nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttribute)
				for _, attribute := range spectype.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
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
					Name:       "ENUMERATION",
					IsExpanded: true,
					FontStyle:  tree.ITALIC,
				}
				nodeSpecType.Children = append(nodeSpecType.Children, nodeSpecTypeAttribute)
				for _, attribute := range spectype.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
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

	tree.StageBranch(stager.GetSpecTreeStage(),
		&tree.Tree{
			Name: stager.GetSpecTreeName(),
			RootNodes: []*tree.Node{
				rootNode,
			},
		},
	)

	stager.GetSpecTreeStage().Commit()
}
