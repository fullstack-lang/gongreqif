package specrelations

import (
	m "github.com/fullstack-lang/gongreqif/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type SpecRelationsTreeStageUpdater struct {
}

func (o *SpecRelationsTreeStageUpdater) UpdateAndCommitSpecRelationsTreeStage(stager *m.Stager) {
	treeStage := stager.GetSpecRelationTreeStage()

	tree.StageBranch(treeStage,
		&tree.Tree{
			Name:      stager.GetSpecRelationsTreeName(),
			RootNodes: nil,
		},
	)

	treeStage.Commit()
}
