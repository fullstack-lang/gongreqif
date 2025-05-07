package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) updateAndCommit_data_type_tree_stage() {

	stager.dataTypeTreeStage.Reset()

	datatypes := stager.rootReqif.CORE_CONTENT.REQ_IF_CONTENT.DATATYPES

	rootNode := &tree.Node{
		Name:       "Data types",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Data types XHTML",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)
		for _, datatype := range datatypes.DATATYPE_DEFINITION_XHTML {
			datatypeCategory.Children = append(datatypeCategory.Children,
				&tree.Node{
					Name: datatype.LONG_NAME,
				},
			)
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Data types String",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)
		for _, datatype := range datatypes.DATATYPE_DEFINITION_STRING {
			datatypeCategory.Children = append(datatypeCategory.Children,
				&tree.Node{
					Name: datatype.LONG_NAME,
				},
			)
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Data types Enumeration",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)
		for _, datatype := range datatypes.DATATYPE_DEFINITION_ENUMERATION {
			node := &tree.Node{
				Name: datatype.LONG_NAME,
			}
			datatypeCategory.Children = append(datatypeCategory.Children, node)
			for _, enum := range datatype.SPECIFIED_VALUES.ENUM_VALUE {
				nodeEnum := &tree.Node{
					Name: enum.LONG_NAME,
				}
				node.Children = append(node.Children, nodeEnum)
			}
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Data types Boolean",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)
		for _, datatype := range datatypes.DATATYPE_DEFINITION_BOOLEAN {
			node := &tree.Node{
				Name: datatype.LONG_NAME,
			}
			datatypeCategory.Children = append(datatypeCategory.Children, node)
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Data types Integer",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)
		for _, datatype := range datatypes.DATATYPE_DEFINITION_INTEGER {
			node := &tree.Node{
				Name: datatype.LONG_NAME,
			}
			datatypeCategory.Children = append(datatypeCategory.Children, node)
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Data types Date",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)
		for _, datatype := range datatypes.DATATYPE_DEFINITION_DATE {
			node := &tree.Node{
				Name: datatype.LONG_NAME,
			}
			datatypeCategory.Children = append(datatypeCategory.Children, node)
		}
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "Data types Real",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, datatypeCategory)
		for _, datatype := range datatypes.DATATYPE_DEFINITION_REAL {
			node := &tree.Node{
				Name: datatype.LONG_NAME,
			}
			datatypeCategory.Children = append(datatypeCategory.Children, node)
		}
	}

	tree.StageBranch(stager.dataTypeTreeStage,
		&tree.Tree{
			Name: stager.dataTypeTreeName,
			RootNodes: []*tree.Node{
				rootNode,
			},
		},
	)

	stager.dataTypeTreeStage.Commit()
}
