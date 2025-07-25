package spectypes

import (
	"log"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

type SpecObjectTypeNodeProxy struct {
	stager         *m.Stager
	specObjectType *m.SPEC_OBJECT_TYPE
}

func (proxy *SpecObjectTypeNodeProxy) OnAfterUpdate(
	treeStage *tree.Stage,
	staged, front *tree.Node) {

	if staged.IsExpanded != front.IsExpanded {
		isNodeExpanded, ok := proxy.stager.Map_SPEC_OBJECT_TYPE_isNodeExpanded[proxy.specObjectType]

		if !ok {
			log.Fatalln("Unknown specificiation in map", proxy.specObjectType.Name)
		}

		proxy.stager.Map_SPEC_OBJECT_TYPE_isNodeExpanded[proxy.specObjectType] = !isNodeExpanded
	}
}
