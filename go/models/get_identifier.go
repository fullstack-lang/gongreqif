package models

func (d *DATATYPE_DEFINITION_XHTML) GetIdentifier() string        { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_STRING) GetIdentifier() string       { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_BOOLEAN) GetIdentifier() string      { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_INTEGER) GetIdentifier() string      { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_REAL) GetIdentifier() string         { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_DATE) GetIdentifier() string         { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_ENUMERATION) GetIdentifier() string  { return d.IDENTIFIER }
func (s *SPEC_OBJECT_TYPE) GetIdentifier() string                 { return s.IDENTIFIER }
func (s *SPECIFICATION_TYPE) GetIdentifier() string               { return s.IDENTIFIER }
func (s *SPEC_OBJECT) GetIdentifier() string                      { return s.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_XHTML) GetIdentifier() string       { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_STRING) GetIdentifier() string      { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_BOOLEAN) GetIdentifier() string     { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_INTEGER) GetIdentifier() string     { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_DATE) GetIdentifier() string        { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_REAL) GetIdentifier() string        { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_ENUMERATION) GetIdentifier() string { return a.IDENTIFIER }
func (e *ENUM_VALUE) GetIdentifier() string                       { return e.IDENTIFIER }
func (s *SPEC_RELATION_TYPE) GetIdentifier() string               { return s.IDENTIFIER }
func (s *SPECIFICATION) GetIdentifier() string                    { return s.IDENTIFIER }

func (a *ATTRIBUTE_DEFINITION_XHTML) GetDatatypeDefinitionRef() string {
	return a.TYPE.DATATYPE_DEFINITION_XHTML_REF
}
func (a *ATTRIBUTE_DEFINITION_STRING) GetDatatypeDefinitionRef() string {
	return a.TYPE.DATATYPE_DEFINITION_STRING_REF
}
func (a *ATTRIBUTE_DEFINITION_BOOLEAN) GetDatatypeDefinitionRef() string {
	return a.TYPE.DATATYPE_DEFINITION_BOOLEAN_REF
}
func (a *ATTRIBUTE_DEFINITION_INTEGER) GetDatatypeDefinitionRef() string {
	return a.TYPE.DATATYPE_DEFINITION_INTEGER_REF
}
func (a *ATTRIBUTE_DEFINITION_DATE) GetDatatypeDefinitionRef() string {
	return a.TYPE.DATATYPE_DEFINITION_DATE_REF
}
func (a *ATTRIBUTE_DEFINITION_REAL) GetDatatypeDefinitionRef() string {
	return a.TYPE.DATATYPE_DEFINITION_REAL_REF
}
func (a *ATTRIBUTE_DEFINITION_ENUMERATION) GetDatatypeDefinitionRef() string {
	return a.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF
}

func (a *ATTRIBUTE_VALUE_XHTML) GetDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF
}
func (a *ATTRIBUTE_VALUE_STRING) GetDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF
}
func (a *ATTRIBUTE_VALUE_BOOLEAN) GetDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF
}
func (a *ATTRIBUTE_VALUE_INTEGER) GetDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF
}
func (a *ATTRIBUTE_VALUE_DATE) GetDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF
}
func (a *ATTRIBUTE_VALUE_REAL) GetDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF
}
func (a *ATTRIBUTE_VALUE_ENUMERATION) GetDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF
}

func (a *ATTRIBUTE_DEFINITION_XHTML) GetIsEditable() bool {
	return a.IS_EDITABLE
}
func (a *ATTRIBUTE_DEFINITION_STRING) GetIsEditable() bool {
	return a.IS_EDITABLE
}
func (a *ATTRIBUTE_DEFINITION_BOOLEAN) GetIsEditable() bool {
	return a.IS_EDITABLE
}
func (a *ATTRIBUTE_DEFINITION_INTEGER) GetIsEditable() bool {
	return a.IS_EDITABLE
}
func (a *ATTRIBUTE_DEFINITION_DATE) GetIsEditable() bool {
	return a.IS_EDITABLE
}
func (a *ATTRIBUTE_DEFINITION_REAL) GetIsEditable() bool {
	return a.IS_EDITABLE
}
func (a *ATTRIBUTE_DEFINITION_ENUMERATION) GetIsEditable() bool {
	return a.IS_EDITABLE
}

func (a *ATTRIBUTE_DEFINITION_XHTML) GetLongName() string {
	return a.LONG_NAME
}
func (a *ATTRIBUTE_DEFINITION_STRING) GetLongName() string {
	return a.LONG_NAME
}
func (a *ATTRIBUTE_DEFINITION_BOOLEAN) GetLongName() string {
	return a.LONG_NAME
}
func (a *ATTRIBUTE_DEFINITION_INTEGER) GetLongName() string {
	return a.LONG_NAME
}
func (a *ATTRIBUTE_DEFINITION_DATE) GetLongName() string {
	return a.LONG_NAME
}
func (a *ATTRIBUTE_DEFINITION_REAL) GetLongName() string {
	return a.LONG_NAME
}
func (a *ATTRIBUTE_DEFINITION_ENUMERATION) GetLongName() string {
	return a.LONG_NAME
}

func (a *A_ATTRIBUTE_DEFINITION_XHTML_REF) GetRef() string {
	return a.ATTRIBUTE_DEFINITION_XHTML_REF
}

func (a *A_ATTRIBUTE_DEFINITION_STRING_REF) GetRef() string {
	return a.ATTRIBUTE_DEFINITION_STRING_REF
}

func (a *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GetRef() string {
	return a.ATTRIBUTE_DEFINITION_BOOLEAN_REF
}

func (a *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GetRef() string {
	return a.ATTRIBUTE_DEFINITION_INTEGER_REF
}

func (a *A_ATTRIBUTE_DEFINITION_DATE_REF) GetRef() string {
	return a.ATTRIBUTE_DEFINITION_DATE_REF
}

func (a *A_ATTRIBUTE_DEFINITION_REAL_REF) GetRef() string {
	return a.ATTRIBUTE_DEFINITION_REAL_REF
}

func (a *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GetRef() string {
	return a.ATTRIBUTE_DEFINITION_ENUMERATION_REF
}

func (a *DATATYPE_DEFINITION_XHTML) GetLongName() string {
	return a.LONG_NAME
}
func (a *DATATYPE_DEFINITION_STRING) GetLongName() string {
	return a.LONG_NAME
}
func (a *DATATYPE_DEFINITION_BOOLEAN) GetLongName() string {
	return a.LONG_NAME
}
func (a *DATATYPE_DEFINITION_INTEGER) GetLongName() string {
	return a.LONG_NAME
}
func (a *DATATYPE_DEFINITION_DATE) GetLongName() string {
	return a.LONG_NAME
}
func (a *DATATYPE_DEFINITION_REAL) GetLongName() string {
	return a.LONG_NAME
}
func (a *DATATYPE_DEFINITION_ENUMERATION) GetLongName() string {
	return a.LONG_NAME
}

func (o *SPEC_OBJECT) GetValues() *A_ATTRIBUTE_VALUE_XHTML_1 {
	return o.VALUES
}

func (o *SPEC_RELATION) GetValues() *A_ATTRIBUTE_VALUE_XHTML_1 {
	return o.VALUES
}

func (o *SPECIFICATION) GetValues() *A_ATTRIBUTE_VALUE_XHTML_1 {
	return o.VALUES
}
