package models

// GetBoolValueFromMap retrieves a boolean value from a slice of Map_identifier_bool
// based on the provided identifier. It returns the value and true if found,
// otherwise it returns false and false.
func (renderingConf *RenderingConfiguration) getBoolValueFromMap(entries []*Map_identifier_bool, identifier string) (bool, bool) {
	for _, entry := range entries {
		if entry.Name == identifier {
			return entry.Value, true
		}
	}
	return false, false
}

// SetBoolValueInMap sets a boolean value in a slice of Map_identifier_bool
// based on the provided identifier. If an entry with the identifier is found,
// it updates the value. Otherwise, it creates a new entry and appends it to the slice.
func (renderingConf *RenderingConfiguration) setBoolValueInMap(entries *[]*Map_identifier_bool, identifier string, value bool) {
	for _, entry := range *entries {
		if entry.Name == identifier {
			entry.Value = value
			return
		}
	}
	newEntry := &Map_identifier_bool{
		Name:  identifier,
		Value: value,
	}
	*entries = append(*entries, newEntry)
}

// Get_SPEC_OBJECT_TYPE_isNodeExpanded is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_SPEC_OBJECT_TYPE_isNodeExpanded(instance *SPEC_OBJECT_TYPE) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries, instance.GetIdentifier())
	return val
}

// Set_SPEC_OBJECT_TYPE_isNodeExpanded is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_SPEC_OBJECT_TYPE_isNodeExpanded(instance *SPEC_OBJECT_TYPE, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle(instance *ATTRIBUTE_DEFINITION_XHTML) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle(instance *ATTRIBUTE_DEFINITION_XHTML, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_STRING_ShowInTitle is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_STRING_ShowInTitle(instance *ATTRIBUTE_DEFINITION_STRING) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_STRING_ShowInTitle is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_STRING_ShowInTitle(instance *ATTRIBUTE_DEFINITION_STRING, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle(instance *ATTRIBUTE_DEFINITION_BOOLEAN) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle(instance *ATTRIBUTE_DEFINITION_BOOLEAN, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle(instance *ATTRIBUTE_DEFINITION_INTEGER) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle(instance *ATTRIBUTE_DEFINITION_INTEGER, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_DATE_ShowInTitle is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_DATE_ShowInTitle(instance *ATTRIBUTE_DEFINITION_DATE) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_DATE_ShowInTitle is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_DATE_ShowInTitle(instance *ATTRIBUTE_DEFINITION_DATE, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_REAL_ShowInTitle is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_REAL_ShowInTitle(instance *ATTRIBUTE_DEFINITION_REAL) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_REAL_ShowInTitle is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_REAL_ShowInTitle(instance *ATTRIBUTE_DEFINITION_REAL, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle(instance *ATTRIBUTE_DEFINITION_ENUMERATION) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle(instance *ATTRIBUTE_DEFINITION_ENUMERATION, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_XHTML_ShowInTable is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_XHTML_ShowInTable(instance *ATTRIBUTE_DEFINITION_XHTML) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_XHTML_ShowInTable is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_XHTML_ShowInTable(instance *ATTRIBUTE_DEFINITION_XHTML, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_STRING_ShowInTable is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_STRING_ShowInTable(instance *ATTRIBUTE_DEFINITION_STRING) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_STRING_ShowInTable is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_STRING_ShowInTable(instance *ATTRIBUTE_DEFINITION_STRING, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable(instance *ATTRIBUTE_DEFINITION_BOOLEAN) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable(instance *ATTRIBUTE_DEFINITION_BOOLEAN, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable(instance *ATTRIBUTE_DEFINITION_INTEGER) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable(instance *ATTRIBUTE_DEFINITION_INTEGER, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_DATE_ShowInTable is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_DATE_ShowInTable(instance *ATTRIBUTE_DEFINITION_DATE) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_DATE_ShowInTable is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_DATE_ShowInTable(instance *ATTRIBUTE_DEFINITION_DATE, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_REAL_ShowInTable is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_REAL_ShowInTable(instance *ATTRIBUTE_DEFINITION_REAL) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_REAL_ShowInTable is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_REAL_ShowInTable(instance *ATTRIBUTE_DEFINITION_REAL, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable(instance *ATTRIBUTE_DEFINITION_ENUMERATION) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable(instance *ATTRIBUTE_DEFINITION_ENUMERATION, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject(instance *ATTRIBUTE_DEFINITION_XHTML) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject(instance *ATTRIBUTE_DEFINITION_XHTML, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_STRING_ShowInSubject is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_STRING_ShowInSubject(instance *ATTRIBUTE_DEFINITION_STRING) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_STRING_ShowInSubject is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_STRING_ShowInSubject(instance *ATTRIBUTE_DEFINITION_STRING, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject(instance *ATTRIBUTE_DEFINITION_BOOLEAN) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject(instance *ATTRIBUTE_DEFINITION_BOOLEAN, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject(instance *ATTRIBUTE_DEFINITION_INTEGER) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject(instance *ATTRIBUTE_DEFINITION_INTEGER, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_DATE_ShowInSubject is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_DATE_ShowInSubject(instance *ATTRIBUTE_DEFINITION_DATE) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_DATE_ShowInSubject is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_DATE_ShowInSubject(instance *ATTRIBUTE_DEFINITION_DATE, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_REAL_ShowInSubject is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_REAL_ShowInSubject(instance *ATTRIBUTE_DEFINITION_REAL) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_REAL_ShowInSubject is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_REAL_ShowInSubject(instance *ATTRIBUTE_DEFINITION_REAL, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries, instance.GetIdentifier(), value)
}

// Get_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject(instance *ATTRIBUTE_DEFINITION_ENUMERATION) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries, instance.GetIdentifier())
	return val
}

// Set_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject(instance *ATTRIBUTE_DEFINITION_ENUMERATION, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries, instance.GetIdentifier(), value)
}

// Get_SPECIFICATION_Nodes_expanded is a wrapper for getBoolValueFromMap
func (renderingConf *RenderingConfiguration) Get_SPECIFICATION_Nodes_expanded(instance *SPECIFICATION) bool {
	val, _ := renderingConf.getBoolValueFromMap(renderingConf.Map_SPECIFICATION_Nodes_expandedEntries, instance.GetIdentifier())
	return val
}

// Set_SPECIFICATION_Nodes_expanded is a wrapper for setBoolValueInMap
func (renderingConf *RenderingConfiguration) Set_SPECIFICATION_Nodes_expanded(instance *SPECIFICATION, value bool) {
	renderingConf.setBoolValueInMap(&renderingConf.Map_SPECIFICATION_Nodes_expandedEntries, instance.GetIdentifier(), value)
}
