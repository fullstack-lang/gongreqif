package specifications

import (
	"fmt"
	"log"
	"strings"

	m "github.com/fullstack-lang/gongreqif/go/models"
	"github.com/fullstack-lang/gongreqif/go/models/specobjects"

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

	attributeDefinitionXHTML_id_map := make(map[string]*m.ATTRIBUTE_DEFINITION_XHTML)
	{
		attributeDefinitionXHTMLs := *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_XHTML](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			attributeDefinitionXHTML_id_map[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	attributeDefinitionENUM_id_map := make(map[string]*m.ATTRIBUTE_DEFINITION_ENUMERATION)
	{
		attributeDefinitionENUMs := *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_ENUMERATION](stager.GetStage())
		for attributeDefinition := range attributeDefinitionENUMs {
			attributeDefinitionENUM_id_map[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	enumValues_id_map := make(map[string]*m.ENUM_VALUE)
	{
		enumValuess := *m.GetGongstructInstancesSet[m.ENUM_VALUE](stager.GetStage())
		for enumValue := range enumValuess {
			enumValues_id_map[enumValue.IDENTIFIER] = enumValue
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

		{

			specificationAttributeCategoryXHTML := &tree.Node{
				Name:       "XHTML",
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			relationNode.Children = append(relationNode.Children, specificationAttributeCategoryXHTML)
			for _, attribute := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
				// provide the type
				var attributeDefinition string
				if datatype, ok := attributeDefinitionXHTML_id_map[attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF]; ok {
					attributeDefinition = datatype.LONG_NAME
				} else {
					log.Panic("ATTRIBUTE_DEFINITION_XHTML_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF,
						"unknown ref")
				}

				enclosedText := attribute.THE_VALUE.EnclosedText

				enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:div>", " ")
				enclosedText = strings.ReplaceAll(enclosedText, "</reqif-xhtml:div>", "\n")
				enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:br >", "-")
				enclosedText = strings.ReplaceAll(enclosedText, "</reqif-xhtml:br >", "\n")
				enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:br />", "\n")
				nodeXHTMLAttribute := &tree.Node{
					Name: attributeDefinition + " : " + enclosedText,
				}
				specificationAttributeCategoryXHTML.Children = append(specificationAttributeCategoryXHTML.Children, nodeXHTMLAttribute)
			}
		}
		{
			hierarchyParentNode := &tree.Node{
				Name:       "Hierarchy",
				IsExpanded: true,
				FontStyle:  tree.ITALIC,
			}
			relationNode.Children = append(relationNode.Children, hierarchyParentNode)

			for _, specHierarchy := range specification.CHILDREN.SPEC_HIERARCHY {

				processSpecHierarchy(
					map_id_specObject,
					attributeDefinitionXHTML_id_map,
					attributeDefinitionENUM_id_map,
					enumValues_id_map,
					specHierarchy,
					hierarchyParentNode)
			}
		}

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

func processSpecHierarchy(
	map_id_specObject map[string]*m.SPEC_OBJECT,
	attributeDefinitionXHTML_id_map map[string]*m.ATTRIBUTE_DEFINITION_XHTML,
	attributeDefinitionENUM_id_map map[string]*m.ATTRIBUTE_DEFINITION_ENUMERATION,
	enumValues_id_map map[string]*m.ENUM_VALUE,
	specHierarchy *m.SPEC_HIERARCHY,
	hierarchyParentNode *tree.Node) {
	specObject, ok := map_id_specObject[specHierarchy.OBJECT.SPEC_OBJECT_REF]
	if !ok {
		log.Panic("specHierarchy.OBJECT.SPEC_OBJECT_REF", specHierarchy.OBJECT.SPEC_OBJECT_REF,
			"unknown ref")
	}
	hierarchyNode := &tree.Node{
		Name: specObject.Name,
	}
	hierarchyParentNode.Children = append(hierarchyParentNode.Children, hierarchyNode)

	specobjects.AddAttributeXHTMLNodes(hierarchyNode, specObject, attributeDefinitionXHTML_id_map)
	specobjects.AddAttributeENUMNodes(hierarchyNode, specObject, attributeDefinitionENUM_id_map, enumValues_id_map)

	if specHierarchy.CHILDREN != nil {
		for _, specHierarchyChildren := range specHierarchy.CHILDREN.SPEC_HIERARCHY {
			processSpecHierarchy(
				map_id_specObject,
				attributeDefinitionXHTML_id_map,
				attributeDefinitionENUM_id_map,
				enumValues_id_map,
				specHierarchyChildren,
				hierarchyNode)
		}
	}
}
