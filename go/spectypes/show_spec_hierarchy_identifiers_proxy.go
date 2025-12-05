package spectypes

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

type ButtonToggleShowSpecHierachyIdentifiersProxy struct {
	stager *m.Stager
}

func (proxy *ButtonToggleShowSpecHierachyIdentifiersProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	proxy.stager.RenderingConf.ShowSpecHierachyIdentifiers = !proxy.stager.RenderingConf.ShowSpecHierachyIdentifiers

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
}
