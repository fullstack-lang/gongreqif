package models

func (stager *Stager) initMaps() {

	stager.Map_id_DATATYPE_DEFINITION_XHTML = populateIdMap[*DATATYPE_DEFINITION_XHTML](stager)
	stager.Map_id_DATATYPE_DEFINITION_STRING = populateIdMap[*DATATYPE_DEFINITION_STRING](stager)
	stager.Map_id_DATATYPE_DEFINITION_BOOLEAN = populateIdMap[*DATATYPE_DEFINITION_BOOLEAN](stager)
	stager.Map_id_DATATYPE_DEFINITION_INTEGER = populateIdMap[*DATATYPE_DEFINITION_INTEGER](stager)
	stager.Map_id_DATATYPE_DEFINITION_REAL = populateIdMap[*DATATYPE_DEFINITION_REAL](stager)
	stager.Map_id_DATATYPE_DEFINITION_DATE = populateIdMap[*DATATYPE_DEFINITION_DATE](stager)
	stager.Map_id_DATATYPE_DEFINITION_ENUMERATION = populateIdMap[*DATATYPE_DEFINITION_ENUMERATION](stager)

	stager.Map_id_SPEC_OBJECT_TYPE = populateIdMap[*SPEC_OBJECT_TYPE](stager)
	stager.Map_id_SPECIFICATION_TYPE = populateIdMap[*SPECIFICATION_TYPE](stager)
	stager.Map_id_SPEC_OBJECT = populateIdMap[*SPEC_OBJECT](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_XHTML = populateIdMap[*ATTRIBUTE_DEFINITION_XHTML](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_STRING = populateIdMap[*ATTRIBUTE_DEFINITION_STRING](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN = populateIdMap[*ATTRIBUTE_DEFINITION_BOOLEAN](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER = populateIdMap[*ATTRIBUTE_DEFINITION_INTEGER](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_DATE = populateIdMap[*ATTRIBUTE_DEFINITION_DATE](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_REAL = populateIdMap[*ATTRIBUTE_DEFINITION_REAL](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION = populateIdMap[*ATTRIBUTE_DEFINITION_ENUMERATION](stager)

	stager.Map_id_ENUM_VALUE = populateIdMap[*ENUM_VALUE](stager)
	stager.Map_id_SPEC_RELATION_TYPE = populateIdMap[*SPEC_RELATION_TYPE](stager)

	stager.Map_SPECIFICATION_Nodes_expanded = populateBoolMap[*SPECIFICATION](stager)
	stager.Map_SPEC_OBJECT_TYPE_showIdentifier = populateBoolMap[*SPEC_OBJECT_TYPE](stager)
	stager.Map_SPEC_OBJECT_TYPE_showName = populateBoolMap[*SPEC_OBJECT_TYPE](stager)

	stager.Map_SPEC_OBJECT_TYPE_isHeading = initializePointerToGongstructMap[*SPEC_OBJECT_TYPE](stager)
	stager.Map_SPEC_OBJECT_TYPE_isNodeExpanded = initializePointerToGongstructMap[*SPEC_OBJECT_TYPE](stager)

	stager.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_XHTML](stager)
	stager.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitle = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_STRING](stager)
	stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_BOOLEAN](stager)
	stager.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_INTEGER](stager)
	stager.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitle = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_DATE](stager)
	stager.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitle = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_REAL](stager)
	stager.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_ENUMERATION](stager)

	stager.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTable = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_XHTML](stager)
	stager.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTable = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_STRING](stager)
	stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_BOOLEAN](stager)
	stager.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_INTEGER](stager)
	stager.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTable = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_DATE](stager)
	stager.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTable = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_REAL](stager)
	stager.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_ENUMERATION](stager)

	stager.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_XHTML](stager)
	stager.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubject = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_STRING](stager)
	stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_BOOLEAN](stager)
	stager.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_INTEGER](stager)
	stager.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubject = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_DATE](stager)
	stager.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubject = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_REAL](stager)
	stager.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject = initializePointerToGongstructMap[*ATTRIBUTE_DEFINITION_ENUMERATION](stager)
}

// Generic function to initialize a map with *T as key and bool as value
func initializePointerToGongstructMap[T PointerToGongstruct](stager *Stager) map[T]bool {
	resultMap := make(map[T]bool)

	instances := *GetGongstructInstancesSetFromPointerType[T](stager.GetStage())
	for instance := range instances {
		resultMap[instance] = false
	}

	return resultMap
}

// populateIdMap fetches instances of a Gongstruct type T and populates a map
// where keys are string identifiers and values are pointers to the instances.
// T must satisfy the Identifiable interface.
func populateIdMap[T Identifiable](stager *Stager) map[string]T {
	instances := *GetGongstructInstancesSetFromPointerType[T](stager.GetStage())
	resultMap := make(map[string]T)
	for instance := range instances {
		resultMap[instance.GetIdentifier()] = instance // Correctly calls the method
	}
	return resultMap
}

// populateBoolMap fetches instances of a Gongstruct type T and populates a map
// where keys are pointers to instances and values are booleans (initialized to false).
func populateBoolMap[T Identifiable](stager *Stager) map[T]bool {
	resultMap := make(map[T]bool)
	instances := *GetGongstructInstancesSetFromPointerType[T](stager.GetStage())
	for instance := range instances {
		resultMap[instance] = false // Initialize with false
	}
	return resultMap
}

// populateBoolMap fetches instances of a Gongstruct type T and populates a map
// where keys are pointers to instances and values are booleans (initialized to false).
func populateBoolMapTrue[T Identifiable](stager *Stager) map[T]bool {
	resultMap := make(map[T]bool)
	instances := *GetGongstructInstancesSetFromPointerType[T](stager.GetStage())
	for instance := range instances {
		resultMap[instance] = true // Initialize with false
	}
	return resultMap
}
