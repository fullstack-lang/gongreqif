package spectypes

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

type ButtonToggleShowSpecObjectTypeLongNameProxy struct {
	stager         *m.Stager
	specObjectType *m.SPEC_OBJECT_TYPE
}

func (proxy *ButtonToggleShowSpecObjectTypeLongNameProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	showLongName := proxy.stager.RenderingConf.Get_SPEC_OBJECT_TYPE_showName(proxy.specObjectType)

	proxy.stager.RenderingConf.Set_SPEC_OBJECT_TYPE_showName(proxy.specObjectType, !showLongName)

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
}
