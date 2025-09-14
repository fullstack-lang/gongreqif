package models

type RenderingConfiguration struct {
	Name string

	Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries                []*Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry
	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries         []*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
	Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries        []*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries       []*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry
	Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries       []*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry
	Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries          []*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry
	Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries          []*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries   []*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries         []*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
	Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries        []*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries       []*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry
	Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries       []*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry
	Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries          []*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry
	Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries          []*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries   []*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry
	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries       []*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry
	Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries      []*Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries     []*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry
	Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries     []*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry
	Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries        []*Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry
	Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries        []*Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries []*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
	Map_SPECIFICATION_Nodes_expandedEntries                   []*Map_SPECIFICATION_Nodes_expandedEntry
	Map_SPEC_OBJECT_TYPE_showIdentifierEntries                []*Map_SPEC_OBJECT_TYPE_showIdentifierEntry
	Map_SPEC_OBJECT_TYPE_showNameEntries                      []*Map_SPEC_OBJECT_TYPE_showNameEntry
}

type Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry struct {
	Name  string
	Value bool
}

type Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry struct {
	Name  string
	Value bool
}

type Map_SPECIFICATION_Nodes_expandedEntry struct {
	Name  string
	Value bool
}

type Map_SPEC_OBJECT_TYPE_showIdentifierEntry struct {
	Name  string
	Value bool
}

type Map_SPEC_OBJECT_TYPE_showNameEntry struct {
	Name  string
	Value bool
}
