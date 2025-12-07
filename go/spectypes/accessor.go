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
