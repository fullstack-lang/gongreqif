package spectypes

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/gongreqif/go/models"
)

type toggleButtonProxy struct {
	stager      *models.Stager
	toggleValue *bool
}

func (p *toggleButtonProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	*p.toggleValue = !*p.toggleValue

	p.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(p.stager)
	p.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(p.stager)
}
