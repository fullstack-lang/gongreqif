package specobjects

import (
	m "github.com/fullstack-lang/gongreqif/go/models"
)

func AddIdentifierAndNameToTitle(stager *m.Stager, specObjectType *m.SPEC_OBJECT_TYPE, markDownContent *string, specObject *m.SPEC_OBJECT) {
	if stager.RenderingConf.Get_SPEC_OBJECT_TYPE_showIdentifier(specObjectType) {
		*markDownContent += specObject.IDENTIFIER
	}

	if stager.RenderingConf.Get_SPEC_OBJECT_TYPE_showIdentifier(specObjectType) &&
		stager.RenderingConf.Get_SPEC_OBJECT_TYPE_showName(specObjectType) {
		*markDownContent += " - "
	}

	if stager.RenderingConf.Get_SPEC_OBJECT_TYPE_showName(specObjectType) {
		*markDownContent += specObject.Name
	}

	if stager.RenderingConf.Get_SPEC_OBJECT_TYPE_showIdentifier(specObjectType) ||
		stager.RenderingConf.Get_SPEC_OBJECT_TYPE_showName(specObjectType) {
		*markDownContent += " "
	}
}
