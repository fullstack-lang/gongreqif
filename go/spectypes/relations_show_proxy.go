package spectypes

import (
	"log"

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

	show, ok := proxy.stager.Map_SPEC_OBJECT_TYPE_showRelations[proxy.specObjectType]

	if !ok {
		log.Fatalln("Unknown specificiation in map", proxy.specObjectType.Name)
	}

	proxy.stager.Map_SPEC_OBJECT_TYPE_showRelations[proxy.specObjectType] = !show

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
}
