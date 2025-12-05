package specobjects

import (
	m "github.com/fullstack-lang/gongreqif/go/models"
)

func AddIdentifierAndNameToTitle(stager *m.Stager, specObjectType *m.SPEC_OBJECT_TYPE, markDownContent *string, specObject *m.SPEC_OBJECT) {
	if stager.Map_SPEC_OBJECT_TYPE_showIdentifier[specObjectType] {
		*markDownContent += specObject.IDENTIFIER
	}

	if stager.Map_SPEC_OBJECT_TYPE_showIdentifier[specObjectType] &&
		stager.Map_SPEC_OBJECT_TYPE_showName[specObjectType] {
		*markDownContent += " - "
	}

	if stager.Map_SPEC_OBJECT_TYPE_showName[specObjectType] {
		*markDownContent += specObject.Name
	}

	if stager.Map_SPEC_OBJECT_TYPE_showIdentifier[specObjectType] ||
		stager.Map_SPEC_OBJECT_TYPE_showName[specObjectType] {
		*markDownContent += " "
	}
}
