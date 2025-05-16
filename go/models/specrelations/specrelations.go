package specrelations

import (
	"fmt"
	"log"

	m "github.com/fullstack-lang/gongreqif/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type SpecRelationsTreeStageUpdater struct {
}

func (o *SpecRelationsTreeStageUpdater) UpdateAndCommitSpecRelationsTreeStage(stager *m.Stager) {
	treeStage := stager.GetSpecRelationsTreeStage()

	sliceOfSpecRelationNodes := make([]*tree.Node, 0)

	relations := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_RELATIONS

	map_id_specobjectTypes := make(map[string]*m.SPEC_OBJECT_TYPE)
	{
		specObjectTypes := *m.GetGongstructInstancesSet[m.SPEC_OBJECT_TYPE](stager.GetStage())
		for specObjectType := range specObjectTypes {
			map_id_specobjectTypes[specObjectType.IDENTIFIER] = specObjectType
		}
	}

	map_id_specRelationType := make(map[string]*m.SPEC_RELATION_TYPE)
	{
		specRelationTypes := *m.GetGongstructInstancesSet[m.SPEC_RELATION_TYPE](stager.GetStage())
		for specRelationType := range specRelationTypes {
			map_id_specRelationType[specRelationType.IDENTIFIER] = specRelationType
		}
	}

	map_id_specObject := make(map[string]*m.SPEC_OBJECT)
	{
		specObjects := *m.GetGongstructInstancesSet[m.SPEC_OBJECT](stager.GetStage())
		for specObject := range specObjects {
			map_id_specObject[specObject.IDENTIFIER] = specObject
		}
	}

	// prepare one node per spec relation type
	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES
	map_specRelationType_node := make(map[*m.SPEC_RELATION_TYPE]*tree.Node)
	map_specRelationType_nbInstances := make(map[*m.SPEC_RELATION_TYPE]int)
	for _, specRelationType := range spectypes.SPEC_RELATION_TYPE {
		nodeSpecRelationType := &tree.Node{
			Name: specRelationType.Name,
		}
		sliceOfSpecRelationNodes = append(sliceOfSpecRelationNodes, nodeSpecRelationType)
		map_specRelationType_node[specRelationType] = nodeSpecRelationType
	}

	for _, specRelation := range relations.SPEC_RELATION {

		specRelationType, ok := map_id_specRelationType[specRelation.TYPE.SPEC_RELATION_TYPE_REF]
		if !ok {
			log.Panic("specRelation.TYPE.SPEC_RELATION_TYPE_REF", specRelation.TYPE.SPEC_RELATION_TYPE_REF,
				"unknown relation type")
		}

		relationNode := &tree.Node{
			Name: specRelation.Name,
		}
		map_specRelationType_node[specRelationType].Children =
			append(map_specRelationType_node[specRelationType].Children, relationNode)
		map_specRelationType_nbInstances[specRelationType] = map_specRelationType_nbInstances[specRelationType] + 1

		{
			specObject, ok := map_id_specObject[specRelation.SOURCE.SPEC_OBJECT_REF]
			if !ok {
				log.Panic("specRelation.SOURCE.SPEC_OBJECT_REF", specRelation.SOURCE.SPEC_OBJECT_REF,
					"unknown ref")
			}

			specObjectType, ok := map_id_specobjectTypes[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
			if !ok {
				log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
					"unknown object type")
			}
			sourceNode := &tree.Node{
				Name:       "Source: " + specObject.Name + " : " + specObjectType.Name,
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			relationNode.Children = append(relationNode.Children, sourceNode)
		}

		{
			specObject, ok := map_id_specObject[specRelation.TARGET.SPEC_OBJECT_REF]
			if !ok {
				log.Panic("specRelation.TARGET.SPEC_OBJECT_REF", specRelation.TARGET.SPEC_OBJECT_REF,
					"unknown ref")
			}

			specObjectType, ok := map_id_specobjectTypes[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
			if !ok {
				log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
					"unknown object type")
			}
			targetNode := &tree.Node{
				Name:       "Target: " + specObject.Name + " : " + specObjectType.Name,
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			relationNode.Children = append(relationNode.Children, targetNode)
		}

	}

	// update the node with the number of instances
	for _, specRelationType := range spectypes.SPEC_RELATION_TYPE {
		nbInstances := map_specRelationType_nbInstances[specRelationType]
		nodeSpecType := map_specRelationType_node[specRelationType]
		nodeSpecType.Name = nodeSpecType.Name + fmt.Sprintf(" (%d)", nbInstances)
	}

	tree.StageBranch(treeStage,
		&tree.Tree{
			Name:      stager.GetSpecRelationsTreeName(),
			RootNodes: sliceOfSpecRelationNodes,
		},
	)

	treeStage.Commit()
}
