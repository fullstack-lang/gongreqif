package spectypes

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

type ButtonToggleShowSpecObjectTypeRelationsProxy struct {
	stager         *m.Stager
	specObjectType *m.SPEC_OBJECT_TYPE
}

func (proxy *ButtonToggleShowSpecObjectTypeRelationsProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	show := proxy.stager.RenderingConf.Get_SPEC_OBJECT_TYPE_showRelations(proxy.specObjectType)

	proxy.stager.RenderingConf.Set_SPEC_OBJECT_TYPE_showRelations(proxy.specObjectType, !show)

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
}
