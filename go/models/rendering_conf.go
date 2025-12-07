package models

// ATTRIBUTE_DEFINITION_Rendering store the rendering choices
type ATTRIBUTE_DEFINITION_Rendering struct {
	Name                 string // the identifier
	ShowInTableEntries   bool
	ShowInTitleEntries   bool
	ShowInSubjectEntries bool
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

// RenderingConfiguration is the persistance in gong
// of the rendering configuration
//
// the need. The rendering configuration is a set of boolean values related to
// object types or attributes type.
// For instance, for the object of type X, sho the identifier in the title
//
// problem, in gong, it is no field of type map[<gong instance>]bool
//
// therefore, one need to trick the model
type RenderingConfiguration struct {
	Name string

	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries       []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries      []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries     []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries     []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries        []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries        []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries []*Map_identifier_bool

	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries       []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries      []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries     []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries     []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries        []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries        []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries []*Map_identifier_bool

	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries       []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries      []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries     []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries     []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries        []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries        []*Map_identifier_bool
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries []*Map_identifier_bool

	ShowSpecHierachyIdentifiers bool
}

type Map_identifier_bool struct {
	Name  string // the ReqIF unique identifier
	Value bool
}
