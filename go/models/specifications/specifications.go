package specifications

import (
	"fmt"
	"log"

	m "github.com/fullstack-lang/gongreqif/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type SpecificationsTreeStageUpdater struct {
}

func (o *SpecificationsTreeStageUpdater) UpdateAndCommitSpecificationsTreeStage(stager *m.Stager) {
	treeStage := stager.GetSpecificationsTreeStage()

	sliceOfSpecificationNodes := make([]*tree.Node, 0)

	specifications := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPECIFICATIONS

	map_id_specobjectTypes := make(map[string]*m.SPEC_OBJECT_TYPE)
	{
		specObjectTypes := *m.GetGongstructInstancesSet[m.SPEC_OBJECT_TYPE](stager.GetStage())
		for specObjectType := range specObjectTypes {
			map_id_specobjectTypes[specObjectType.IDENTIFIER] = specObjectType
		}
	}

	map_id_specificationType := make(map[string]*m.SPECIFICATION_TYPE)
	{
		types := *m.GetGongstructInstancesSet[m.SPECIFICATION_TYPE](stager.GetStage())
		for type_ := range types {
			map_id_specificationType[type_.IDENTIFIER] = type_
		}
	}

	map_id_specObject := make(map[string]*m.SPEC_OBJECT)
	{
		specObjects := *m.GetGongstructInstancesSet[m.SPEC_OBJECT](stager.GetStage())
		for specObject := range specObjects {
			map_id_specObject[specObject.IDENTIFIER] = specObject
		}
	}

	// prepare one node per specification type
	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES
	map_specificationType_node := make(map[*m.SPECIFICATION_TYPE]*tree.Node)
	map_specificationType_nbInstances := make(map[*m.SPECIFICATION_TYPE]int)
	for _, specificationType := range spectypes.SPECIFICATION_TYPE {
		nodeSpecificationType := &tree.Node{
			Name: specificationType.Name,
		}
		sliceOfSpecificationNodes = append(sliceOfSpecificationNodes, nodeSpecificationType)
		map_specificationType_node[specificationType] = nodeSpecificationType
	}

	for _, specification := range specifications.SPECIFICATION {

		specificationType, ok := map_id_specificationType[specification.TYPE.SPECIFICATION_TYPE_REF]
		if !ok {
			log.Panic("specRelation.TYPE.SPECIFICATION_TYPE_REF", specification.TYPE.SPECIFICATION_TYPE_REF,
				"unknown relation type")
		}

		relationNode := &tree.Node{
			Name: specification.Name,
		}
		map_specificationType_node[specificationType].Children =
			append(map_specificationType_node[specificationType].Children, relationNode)
		map_specificationType_nbInstances[specificationType] = map_specificationType_nbInstances[specificationType] + 1
	}

	// update the node with the number of instances
	for _, type_ := range spectypes.SPECIFICATION_TYPE {
		nbInstances := map_specificationType_nbInstances[type_]
		nodeSpecType := map_specificationType_node[type_]
		nodeSpecType.Name = nodeSpecType.Name + fmt.Sprintf(" (%d)", nbInstances)
	}

	tree.StageBranch(treeStage,
		&tree.Tree{
			Name:      stager.GetSpecificationsTreeName(),
			RootNodes: sliceOfSpecificationNodes,
		},
	)

	treeStage.Commit()
}
