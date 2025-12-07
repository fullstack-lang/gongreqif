package specifications

import "github.com/fullstack-lang/gongreqif/go/models"

func GetSelectedSpecification(stage *models.Stage) (specification *models.SPECIFICATION) {
	for specificationRendering := range *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION_Rendering](stage) {
		if specificationRendering.IsSelected {
			for specification_ := range *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION](stage) {
				if specification_.GetIdentifier() == specificationRendering.GetName() {
					specification = specification_
				}
			}
		}
	}

	return
}

func SetSelectedSpecification(stage *models.Stage, specification *models.SPECIFICATION) {

	for specificationRendering := range *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION_Rendering](stage) {
		specificationRendering.IsSelected = false
		if specificationRendering.GetName() == specification.GetIdentifier() {
			specificationRendering.IsSelected = true
		}
	}
}
