package spectypes

import (
	"fmt"
	"log"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	m "github.com/fullstack-lang/gongreqif/go/models"
)

func addAttributeNode[
	AttrDef m.AttributeDefinition,
	AttrDefRef m.AttributeDefinitionRef,
	DatatypeDef m.DatatypeDefinition](

	stager *m.Stager,

	attributeDefinitions []AttrDef, // the attributes definitions to display
	map_Id_AttributeDefinition map[string]AttrDef, // map to find attribute def from the ref
	map_AtttributeDefinition_Spec_nbInstance map[AttrDef]int, // number of use in the selected specification
	map_AtttributeDefinition_showInTitle map[AttrDef]bool,
	map_AtttributeDefinition_showInTable map[AttrDef]bool,
	map_AtttributeDefinition_showInSubject map[AttrDef]bool,

	attributesDefinitionRefs map[AttrDefRef]any, // the set of all reference to this kind of attribute definition
	map_Id_DatatypeDefinition map[string]DatatypeDef,

	nodeSpecType *tree.Node) {
	if len(attributeDefinitions) > 0 {

		// compute the number of time this attribute is used
		map_attributeDefinition_nbInstance := make(map[AttrDef]int)
		var attributeDefinition AttrDef
		for x := range attributesDefinitionRefs {

			var ok bool
			attributeDefinition, ok = map_Id_AttributeDefinition[x.GetRef()]
			if !ok {
				log.Panic("x.GetRef()", x.GetRef(),
					"unknown ref")
			} else {
				map_attributeDefinition_nbInstance[attributeDefinition]++
			}
		}

		for _, attributeDefinition := range attributeDefinitions {
			var attributeType string
			if datatype, ok := map_Id_DatatypeDefinition[attributeDefinition.GetDatatypeDefinitionRef()]; ok {
				attributeType = datatype.GetLongName()
			} else {
				log.Panic("attributeDefinition: ",
					attributeDefinition.GetLongName(), " ref:",
					attributeDefinition.GetDatatypeDefinitionRef(), ": ",
					"unknown ref")
			}

			nbInstances := map_AtttributeDefinition_Spec_nbInstance[attributeDefinition]
			nodeAttribute := &tree.Node{
				Name: attributeDefinition.GetLongName() + ":" + attributeType +
					fmt.Sprintf(" (%d/%d)", nbInstances, map_attributeDefinition_nbInstance[attributeDefinition]),
			}
			configureAndAddAttributeNode(
				stager,
				nodeSpecType,
				nodeAttribute,
				nbInstances,
				attributeDefinition.GetIsEditable(),
				attributeDefinition.GetLongName(),
				attributeDefinition,
				map_AtttributeDefinition_showInTitle,
				map_AtttributeDefinition_showInTable,
				map_AtttributeDefinition_showInSubject,
			)
		}
	}
}
