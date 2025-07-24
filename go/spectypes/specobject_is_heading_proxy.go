package spectypes

import (
	"log"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

type ButtonToggleShowSpecObjectTypeIsHeadingProxy struct {
	stager         *m.Stager
	specObjectType *m.SPEC_OBJECT_TYPE
}

func (proxy *ButtonToggleShowSpecObjectTypeIsHeadingProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	shoeIsHeading, ok := proxy.stager.Map_SPEC_OBJECT_TYPE_isHeading[proxy.specObjectType]

	if !ok {
		log.Fatalln("Unknown specificiation in map", proxy.specObjectType.Name)
	}

	proxy.stager.Map_SPEC_OBJECT_TYPE_isHeading[proxy.specObjectType] = !shoeIsHeading

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
}
