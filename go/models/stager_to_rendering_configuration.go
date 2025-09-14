package models

// Do not edit, use GEMINI

// ToRenderingConfiguration creates a RenderingConfiguration instance from the Stager's fields.
// This method iterates through each map in the Stager, converts its key-value pairs
// into a slice of the corresponding Entry structs, and populates a new RenderingConfiguration.
//
// The code assumes that all key types (e.g., *SPEC_OBJECT_TYPE, *ATTRIBUTE_DEFINITION_STRING)
// have an exported string field named 'IDENTIFIER' which is used to populate the 'Name'
// field of each entry.
//
// This version omits explicit length checks on the maps, as a 'for...range' loop
// over a nil or empty map is a safe no-op in Go. This results in cleaner, more concise code.
func (s *Stager) ToRenderingConfiguration(name string) *RenderingConfiguration {
	config := &RenderingConfiguration{
		Name: name,
	}

	for key, value := range s.Map_SPEC_OBJECT_TYPE_isNodeExpanded {
		config.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries = append(config.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries, &Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle {
		config.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries = append(config.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries, &Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitle {
		config.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries = append(config.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries, &Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle {
		config.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries = append(config.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries, &Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle {
		config.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries = append(config.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries, &Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitle {
		config.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries = append(config.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries, &Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitle {
		config.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries = append(config.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries, &Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle {
		config.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries = append(config.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries, &Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTable {
		config.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries = append(config.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries, &Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTable {
		config.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries = append(config.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries, &Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable {
		config.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries = append(config.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries, &Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable {
		config.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries = append(config.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries, &Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTable {
		config.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries = append(config.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries, &Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTable {
		config.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries = append(config.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries, &Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable {
		config.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries = append(config.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries, &Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject {
		config.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries = append(config.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries, &Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubject {
		config.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries = append(config.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries, &Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject {
		config.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries = append(config.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries, &Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject {
		config.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries = append(config.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries, &Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubject {
		config.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries = append(config.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries, &Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubject {
		config.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries = append(config.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries, &Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject {
		config.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries = append(config.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries, &Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_SPECIFICATION_Nodes_expanded {
		config.Map_SPECIFICATION_Nodes_expandedEntries = append(config.Map_SPECIFICATION_Nodes_expandedEntries, &Map_SPECIFICATION_Nodes_expandedEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_SPEC_OBJECT_TYPE_showIdentifier {
		config.Map_SPEC_OBJECT_TYPE_showIdentifierEntries = append(config.Map_SPEC_OBJECT_TYPE_showIdentifierEntries, &Map_SPEC_OBJECT_TYPE_showIdentifierEntry{Name: key.IDENTIFIER, Value: value})
	}

	for key, value := range s.Map_SPEC_OBJECT_TYPE_showName {
		config.Map_SPEC_OBJECT_TYPE_showNameEntries = append(config.Map_SPEC_OBJECT_TYPE_showNameEntries, &Map_SPEC_OBJECT_TYPE_showNameEntry{Name: key.IDENTIFIER, Value: value})
	}

	return config
}

// FromRenderingConfiguration populates the Stager's maps based on a RenderingConfiguration.
//
// This method iterates through the entries in the configuration and updates the
// corresponding values in the Stager's maps.
//
// CRITICAL ASSUMPTION: This method assumes that the Stager's maps (e.g., s.Map_SPEC_OBJECT_TYPE_isNodeExpanded)
// are already populated with all possible pointer keys. It finds the matching key by comparing
// IDENTIFIER fields. It does NOT create new pointer keys.
// NOTE: This implementation uses nested loops for simplicity, which may be inefficient for very large maps.
func (s *Stager) FromRenderingConfiguration(conf *RenderingConfiguration) {
	for _, entry := range conf.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries {
		for key := range s.Map_SPEC_OBJECT_TYPE_isNodeExpanded {
			if key.IDENTIFIER == entry.Name {
				s.Map_SPEC_OBJECT_TYPE_isNodeExpanded[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitle {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitle[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitle {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitle[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitle {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitle[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTable {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTable[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTable {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTable[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTable {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTable[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTable {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTable[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubject {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubject[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubject {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubject[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubject {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubject[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries {
		for key := range s.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject {
			if key.IDENTIFIER == entry.Name {
				s.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_SPECIFICATION_Nodes_expandedEntries {
		for key := range s.Map_SPECIFICATION_Nodes_expanded {
			if key.IDENTIFIER == entry.Name {
				s.Map_SPECIFICATION_Nodes_expanded[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_SPEC_OBJECT_TYPE_showIdentifierEntries {
		for key := range s.Map_SPEC_OBJECT_TYPE_showIdentifier {
			if key.IDENTIFIER == entry.Name {
				s.Map_SPEC_OBJECT_TYPE_showIdentifier[key] = entry.Value
				break
			}
		}
	}

	for _, entry := range conf.Map_SPEC_OBJECT_TYPE_showNameEntries {
		for key := range s.Map_SPEC_OBJECT_TYPE_showName {
			if key.IDENTIFIER == entry.Name {
				s.Map_SPEC_OBJECT_TYPE_showName[key] = entry.Value
				break
			}
		}
	}
}
