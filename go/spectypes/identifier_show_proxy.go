package spectypes

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

type ButtonToggleShowSpecObjectTypeIdentifierProxy struct {
	stager         *m.Stager
	specObjectType *m.SPEC_OBJECT_TYPE
}

func (proxy *ButtonToggleShowSpecObjectTypeIdentifierProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	showIdentifier := proxy.stager.RenderingConf.Get_SPEC_OBJECT_TYPE_showIdentifier(proxy.specObjectType)

	proxy.stager.RenderingConf.Set_SPEC_OBJECT_TYPE_showIdentifier(proxy.specObjectType, !showIdentifier)

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
}
