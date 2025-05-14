package objects

import (
	m "github.com/fullstack-lang/gongreqif/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ObjectTreeStage struct {
}

// UpdateAndCommitObjectTreeStage implements models.ObjectTreeStageInterface.
func (o *ObjectTreeStage) UpdateAndCommitObjectTreeStage(stager *m.Stager) {
	treeStage := stager.GetObjectTreeStage()

	treeStage.Reset()

	rootNode := &tree.Node{
		Name:       "Objects",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}

	objects := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_OBJECTS

	for _, object := range objects.SPEC_OBJECT {
		objectNode := &tree.Node{
			Name: object.Name,
		}
		rootNode.Children = append(rootNode.Children, objectNode)

	}

	tree.StageBranch(treeStage,
		&tree.Tree{
			Name: stager.GetObjectTreeName(),
			RootNodes: []*tree.Node{
				rootNode,
			},
		},
	)

	treeStage.Commit()
}
