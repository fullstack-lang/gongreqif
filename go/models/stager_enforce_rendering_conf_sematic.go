package models

// ensure that for every object that need a rendering configuration
// there is a rendering configuration
//
// In principle, the rendering configuration should be fields
// for OBJECT_TYPE, ATTRIBUTE_DEFINITIONxxx, ...
// however, since those objects are created from the parsing of the ReqIF
// file, there can be persisted in another file for the rendering conf
// only purpose
func (stager *Stager) enforceRenderingConfigurationSemantic() {

	stage := stager.stage

	for specObjectType := range *GetGongstructInstancesSetFromPointerType[*SPEC_OBJECT_TYPE](stage) {
		if _, ok := stage.SPEC_OBJECT_TYPE_Renderings_mapString[specObjectType.GetIdentifier()]; !ok {
			o := (&SPEC_OBJECT_TYPE_Rendering{}).Stage(stage)
			o.Name = specObjectType.GetIdentifier()
			stage.SPEC_OBJECT_TYPE_Renderings_mapString[specObjectType.GetIdentifier()] = o
		}
	}

	for specObjectTypeRendering := range *GetGongstructInstancesSetFromPointerType[*SPEC_OBJECT_TYPE_Rendering](stage) {
		if _, ok := stager.Map_id_SPEC_OBJECT_TYPE[specObjectTypeRendering.GetName()]; !ok {
			specObjectTypeRendering.Unstage(stage)
		}
	}

	for specification := range *GetGongstructInstancesSetFromPointerType[*SPECIFICATION](stage) {
		if _, ok := stage.SPECIFICATION_Renderings_mapString[specification.GetIdentifier()]; !ok {
			o := (&SPECIFICATION_Rendering{}).Stage(stage)
			o.Name = specification.GetIdentifier()
			stage.SPECIFICATION_Renderings_mapString[specification.GetIdentifier()] = o
		}
	}

	for specificationRendering := range *GetGongstructInstancesSetFromPointerType[*SPECIFICATION_Rendering](stage) {
		if _, ok := stager.Map_id_SPECIFICATION_TYPE[specificationRendering.GetName()]; !ok {
			specificationRendering.Unstage(stage)
		}
	}
}
