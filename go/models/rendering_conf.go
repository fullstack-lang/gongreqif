package models

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

	Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries []*Map_identifier_bool

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

	Map_SPECIFICATION_Nodes_expandedEntries []*Map_identifier_bool

	Map_SPEC_OBJECT_TYPE_showIdentifierEntries []*Map_identifier_bool
	Map_SPEC_OBJECT_TYPE_showNameEntries       []*Map_identifier_bool
	Map_SPEC_OBJECT_TYPE_showRelations         []*Map_identifier_bool

	ShowSpecHierachyIdentifiers bool
}

type Map_identifier_bool struct {
	Name  string
	Value bool
}
