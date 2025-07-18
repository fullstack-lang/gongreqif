package models

// Identifiable is a constraint for types that can provide a string identifier.
type Identifiable interface {
	PointerToGongstruct
	GetIdentifier() string
}

type AttributeDefinition interface {
	Identifiable
	GetDatatypeDefinitionRef() string
	GetIsEditable() bool
	GetLongName() string
}

// Things like A_ATTRIBUTE_DEFINITION_XHTML_REF
type AttributeDefinitionRef interface {
	PointerToGongstruct

	GetRef() string
}

type Attribute interface {
	Identifiable
	GetLongName() string
}

type DatatypeDefinition interface {
	PointerToGongstruct

	GetLongName() string
}

// for SPEC OBJECTS and SPEC RELATIONS
type ObjectWithValues interface {
	GetValues() *A_ATTRIBUTE_VALUE_XHTML_1
}
