package models

// Identifiable is a constraint for types that can provide a string identifier.
type Identifiable interface {
	PointerToGongstruct
	GetIdentifier() string
}

type Attribute interface {
	Identifiable
	GetTypeRef() string
	GetIsEditable() bool
	GetLongName() string
}
