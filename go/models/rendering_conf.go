package models

// ATTRIBUTE_DEFINITION_Rendering store the rendering choices
type ATTRIBUTE_DEFINITION_Rendering struct {
	Name          string // the identifier
	ShowInTable   bool
	ShowInTitle   bool
	ShowInSubject bool
}

func (r *ATTRIBUTE_DEFINITION_Rendering) GetShowInTitlePtr() *bool {
	return &r.ShowInTitle
}

func (r *ATTRIBUTE_DEFINITION_Rendering) GetShowInTablePtr() *bool {
	return &r.ShowInTable
}

func (r *ATTRIBUTE_DEFINITION_Rendering) GetShowInSubjectPtr() *bool {
	return &r.ShowInSubject
}

type ATTRIBUTE_DEFINITION_XHTML_Rendering struct {
	ATTRIBUTE_DEFINITION_Rendering
}

type ATTRIBUTE_DEFINITION_STRING_Rendering struct {
	ATTRIBUTE_DEFINITION_Rendering
}

type ATTRIBUTE_DEFINITION_BOOLEAN_Rendering struct {
	ATTRIBUTE_DEFINITION_Rendering
}

type ATTRIBUTE_DEFINITION_INTEGER_Rendering struct {
	ATTRIBUTE_DEFINITION_Rendering
}

type ATTRIBUTE_DEFINITION_DATE_Rendering struct {
	ATTRIBUTE_DEFINITION_Rendering
}

type ATTRIBUTE_DEFINITION_REAL_Rendering struct {
	ATTRIBUTE_DEFINITION_Rendering
}

type ATTRIBUTE_DEFINITION_ENUMERATION_Rendering struct {
	ATTRIBUTE_DEFINITION_Rendering
}

// SPEC_OBJECT_TYPE_Rendering store the rendering choices
type SPEC_OBJECT_TYPE_Rendering struct {
	Name           string // the identifier
	IsNodeExpanded bool
	ShowIdentifier bool
	ShowName       bool
	ShowRelations  bool
	IsHeading      bool
}

// SPECIFICATION_Rendering store the rendering choices
type SPECIFICATION_Rendering struct {
	Name           string // the identifier
	IsNodeExpanded bool
	IsSelected     bool
}

type Map_identifier_bool struct {
	Name  string // the ReqIF unique identifier
	Value bool
}
