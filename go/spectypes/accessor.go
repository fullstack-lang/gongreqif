package spectypes

import (
	"log"

	"github.com/fullstack-lang/gongreqif/go/models"
)

func GetSpecObjectTypeRendering(
	stage *models.Stage,
	spectObjectType *models.SPEC_OBJECT_TYPE,
) (
	specObjectTypeRendering *models.SPEC_OBJECT_TYPE_Rendering) {

	// SPEC_OBJECT_TYPE_Rendering instances Names are the identifiers of the
	// SPEC_OBJECT_TYPE instance. Since, by ReqIF design, those identifiers are unique,
	// we can use the gong map of those instances
	map_ := models.GongGetMap[*models.SPEC_OBJECT_TYPE_Rendering](stage)

	var ok bool
	if specObjectTypeRendering, ok = map_[spectObjectType.GetIdentifier()]; !ok {
		log.Panic("GetSpecObjectTypeRendering: Unknown specObjectType", spectObjectType.GetName())
	}

	return
}

func GetSpecAttributeDefinitionRendering[
	AttrDef models.AttributeDefinition, AttrDefRendering models.AttributeDefinitionRendering](
	stage *models.Stage,
	specAttributeDefinition AttrDef,
) (
	specAttributeDefinitionRendering AttrDefRendering) {

	// ATTRIBUTE_DEFINITION_Rendering instances Names are the identifiers of the
	// ATTRIBUTE_DEFINITION instance. Since, by ReqIF design, those identifiers are unique,
	// we can use the gong map of those instances
	map_ := models.GongGetMap[AttrDefRendering](stage)

	var ok bool
	if specAttributeDefinitionRendering, ok = map_[specAttributeDefinition.GetIdentifier()]; !ok {
		log.Panic("GetSpecAttributeDefinitionRendering: Unknown specAttributeDefinition", specAttributeDefinition.GetName())
	}

	return
}
