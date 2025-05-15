package datatype

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

type DataTypeTreeStageUpdater struct {
}

func (dataTypeTreeStageUpdater *DataTypeTreeStageUpdater) UpdateAndCommitDataTypeTreeStage(stager *m.Stager) {

	stager.GetDataTypeTreeStage().Reset()

	datatypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.DATATYPES

	rootNode := &tree.Node{
		Name:       "Data types",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}

	{
		datatypeCategory := &tree.Node{
			Name:       "XHTML",
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
			Name:       "String",
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
			Name:       "Enumeration",
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
			Name:       "Boolean",
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
			Name:       "Integer",
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
			Name:       "Date",
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
			Name:       "Real",
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

	tree.StageBranch(stager.GetDataTypeTreeStage(),
		&tree.Tree{
			Name: stager.GetDataTypeTreeName(),
			RootNodes: []*tree.Node{
				rootNode,
			},
		},
	)

	stager.GetDataTypeTreeStage().Commit()
}
