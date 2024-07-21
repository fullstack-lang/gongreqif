// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		// insertion point per field

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_BOOLEAN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_BOOLEAN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_BOOLEAN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_BOOLEAN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN = _inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN =
								append(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_VALUE_BOOLEAN))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_DATE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_DATE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_DATE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_DATE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_DATE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE = _inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE =
								append(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE, any(fieldInstance).(*ATTRIBUTE_VALUE_DATE))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		// insertion point per field
		if fieldName == "DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION = _inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION =
								append(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_VALUE_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_INTEGER:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_INTEGER) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_INTEGER)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_INTEGER) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_INTEGER)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER = _inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER =
								append(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, any(fieldInstance).(*ATTRIBUTE_VALUE_INTEGER))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_REAL:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_REAL) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_REAL)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_REAL) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_REAL)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL = _inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL =
								append(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL, any(fieldInstance).(*ATTRIBUTE_VALUE_REAL))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_STRING:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_STRING) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_STRING)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_STRING) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_STRING)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING = _inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING =
								append(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, any(fieldInstance).(*ATTRIBUTE_VALUE_STRING))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_XHTML:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML = _inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML =
								append(_inferedTypeInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML, any(fieldInstance).(*ATTRIBUTE_VALUE_XHTML))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point per field

	case *ATTRIBUTE_VALUE_DATE:
		// insertion point per field

	case *ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point per field

	case *ATTRIBUTE_VALUE_INTEGER:
		// insertion point per field

	case *ATTRIBUTE_VALUE_REAL:
		// insertion point per field

	case *ATTRIBUTE_VALUE_STRING:
		// insertion point per field

	case *ATTRIBUTE_VALUE_XHTML:
		// insertion point per field
		if fieldName == "THE_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.THE_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.THE_VALUE = _inferedTypeInstance.THE_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.THE_VALUE =
								append(_inferedTypeInstance.THE_VALUE, any(fieldInstance).(*XHTML_CONTENT))
						}
					}
				}
			}
		}
		if fieldName == "THE_ORIGINAL_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.THE_ORIGINAL_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.THE_ORIGINAL_VALUE = _inferedTypeInstance.THE_ORIGINAL_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.THE_ORIGINAL_VALUE =
								append(_inferedTypeInstance.THE_ORIGINAL_VALUE, any(fieldInstance).(*XHTML_CONTENT))
						}
					}
				}
			}
		}

	case *AnyType:
		// insertion point per field

	case *DATATYPE_DEFINITION_BOOLEAN:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_BOOLEAN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_BOOLEAN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_DATE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_DATE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_DATE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_ENUMERATION:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SPECIFIED_VALUES.ENUM_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECIFIED_VALUES.ENUM_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECIFIED_VALUES.ENUM_VALUE = _inferedTypeInstance.SPECIFIED_VALUES.ENUM_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECIFIED_VALUES.ENUM_VALUE =
								append(_inferedTypeInstance.SPECIFIED_VALUES.ENUM_VALUE, any(fieldInstance).(*ENUM_VALUE))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_INTEGER:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_INTEGER) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_INTEGER)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_REAL:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_REAL) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_REAL)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_STRING:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_STRING) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_STRING)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_XHTML:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *EMBEDDED_VALUE:
		// insertion point per field

	case *ENUM_VALUE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ENUM_VALUE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ENUM_VALUE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "PROPERTIES.EMBEDDED_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ENUM_VALUE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ENUM_VALUE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.PROPERTIES.EMBEDDED_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.PROPERTIES.EMBEDDED_VALUE = _inferedTypeInstance.PROPERTIES.EMBEDDED_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.PROPERTIES.EMBEDDED_VALUE =
								append(_inferedTypeInstance.PROPERTIES.EMBEDDED_VALUE, any(fieldInstance).(*EMBEDDED_VALUE))
						}
					}
				}
			}
		}

	case *RELATION_GROUP:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *RELATION_GROUP_TYPE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_DEFINITION_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, any(fieldInstance).(*ATTRIBUTE_DEFINITION_DATE))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_DEFINITION_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, any(fieldInstance).(*ATTRIBUTE_DEFINITION_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, any(fieldInstance).(*ATTRIBUTE_DEFINITION_REAL))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, any(fieldInstance).(*ATTRIBUTE_DEFINITION_STRING))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, any(fieldInstance).(*ATTRIBUTE_DEFINITION_XHTML))
						}
					}
				}
			}
		}

	case *REQ_IF:
		// insertion point per field
		if fieldName == "TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION = _inferedTypeInstance.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION =
								append(_inferedTypeInstance.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION, any(fieldInstance).(*REQ_IF_TOOL_EXTENSION))
						}
					}
				}
			}
		}

	case *REQ_IF_CONTENT:
		// insertion point per field
		if fieldName == "DATATYPES.DATATYPE_DEFINITION_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_BOOLEAN = _inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_BOOLEAN =
								append(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_BOOLEAN, any(fieldInstance).(*DATATYPE_DEFINITION_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPES.DATATYPE_DEFINITION_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_DATE = _inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_DATE =
								append(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_DATE, any(fieldInstance).(*DATATYPE_DEFINITION_DATE))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPES.DATATYPE_DEFINITION_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_ENUMERATION = _inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_ENUMERATION =
								append(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_ENUMERATION, any(fieldInstance).(*DATATYPE_DEFINITION_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPES.DATATYPE_DEFINITION_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_INTEGER = _inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_INTEGER =
								append(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_INTEGER, any(fieldInstance).(*DATATYPE_DEFINITION_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPES.DATATYPE_DEFINITION_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_REAL = _inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_REAL =
								append(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_REAL, any(fieldInstance).(*DATATYPE_DEFINITION_REAL))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPES.DATATYPE_DEFINITION_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_STRING = _inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_STRING =
								append(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_STRING, any(fieldInstance).(*DATATYPE_DEFINITION_STRING))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPES.DATATYPE_DEFINITION_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_XHTML = _inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_XHTML =
								append(_inferedTypeInstance.DATATYPES.DATATYPE_DEFINITION_XHTML, any(fieldInstance).(*DATATYPE_DEFINITION_XHTML))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_TYPES.RELATION_GROUP_TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_TYPES.RELATION_GROUP_TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_TYPES.RELATION_GROUP_TYPE = _inferedTypeInstance.SPEC_TYPES.RELATION_GROUP_TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_TYPES.RELATION_GROUP_TYPE =
								append(_inferedTypeInstance.SPEC_TYPES.RELATION_GROUP_TYPE, any(fieldInstance).(*RELATION_GROUP_TYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_TYPES.SPEC_OBJECT_TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_TYPES.SPEC_OBJECT_TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_TYPES.SPEC_OBJECT_TYPE = _inferedTypeInstance.SPEC_TYPES.SPEC_OBJECT_TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_TYPES.SPEC_OBJECT_TYPE =
								append(_inferedTypeInstance.SPEC_TYPES.SPEC_OBJECT_TYPE, any(fieldInstance).(*SPEC_OBJECT_TYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_TYPES.SPEC_RELATION_TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_TYPES.SPEC_RELATION_TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_TYPES.SPEC_RELATION_TYPE = _inferedTypeInstance.SPEC_TYPES.SPEC_RELATION_TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_TYPES.SPEC_RELATION_TYPE =
								append(_inferedTypeInstance.SPEC_TYPES.SPEC_RELATION_TYPE, any(fieldInstance).(*SPEC_RELATION_TYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_TYPES.SPECIFICATION_TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_TYPES.SPECIFICATION_TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_TYPES.SPECIFICATION_TYPE = _inferedTypeInstance.SPEC_TYPES.SPECIFICATION_TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_TYPES.SPECIFICATION_TYPE =
								append(_inferedTypeInstance.SPEC_TYPES.SPECIFICATION_TYPE, any(fieldInstance).(*SPECIFICATION_TYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_OBJECTS.SPEC_OBJECT" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_OBJECTS.SPEC_OBJECT).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_OBJECTS.SPEC_OBJECT = _inferedTypeInstance.SPEC_OBJECTS.SPEC_OBJECT[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_OBJECTS.SPEC_OBJECT =
								append(_inferedTypeInstance.SPEC_OBJECTS.SPEC_OBJECT, any(fieldInstance).(*SPEC_OBJECT))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_RELATIONS.SPEC_RELATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_RELATIONS.SPEC_RELATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_RELATIONS.SPEC_RELATION = _inferedTypeInstance.SPEC_RELATIONS.SPEC_RELATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_RELATIONS.SPEC_RELATION =
								append(_inferedTypeInstance.SPEC_RELATIONS.SPEC_RELATION, any(fieldInstance).(*SPEC_RELATION))
						}
					}
				}
			}
		}
		if fieldName == "SPECIFICATIONS.SPECIFICATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECIFICATIONS.SPECIFICATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECIFICATIONS.SPECIFICATION = _inferedTypeInstance.SPECIFICATIONS.SPECIFICATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECIFICATIONS.SPECIFICATION =
								append(_inferedTypeInstance.SPECIFICATIONS.SPECIFICATION, any(fieldInstance).(*SPECIFICATION))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_RELATION_GROUPS.RELATION_GROUP" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_RELATION_GROUPS.RELATION_GROUP).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_RELATION_GROUPS.RELATION_GROUP = _inferedTypeInstance.SPEC_RELATION_GROUPS.RELATION_GROUP[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_RELATION_GROUPS.RELATION_GROUP =
								append(_inferedTypeInstance.SPEC_RELATION_GROUPS.RELATION_GROUP, any(fieldInstance).(*RELATION_GROUP))
						}
					}
				}
			}
		}

	case *REQ_IF_HEADER:
		// insertion point per field

	case *REQ_IF_TOOL_EXTENSION:
		// insertion point per field

	case *SPECIFICATION:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_VALUE_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE, any(fieldInstance).(*ATTRIBUTE_VALUE_DATE))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_VALUE_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER, any(fieldInstance).(*ATTRIBUTE_VALUE_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL, any(fieldInstance).(*ATTRIBUTE_VALUE_REAL))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING, any(fieldInstance).(*ATTRIBUTE_VALUE_STRING))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML, any(fieldInstance).(*ATTRIBUTE_VALUE_XHTML))
						}
					}
				}
			}
		}
		if fieldName == "CHILDREN.SPEC_HIERARCHY" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.CHILDREN.SPEC_HIERARCHY).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.CHILDREN.SPEC_HIERARCHY = _inferedTypeInstance.CHILDREN.SPEC_HIERARCHY[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.CHILDREN.SPEC_HIERARCHY =
								append(_inferedTypeInstance.CHILDREN.SPEC_HIERARCHY, any(fieldInstance).(*SPEC_HIERARCHY))
						}
					}
				}
			}
		}

	case *SPECIFICATION_TYPE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_DEFINITION_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, any(fieldInstance).(*ATTRIBUTE_DEFINITION_DATE))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_DEFINITION_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, any(fieldInstance).(*ATTRIBUTE_DEFINITION_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, any(fieldInstance).(*ATTRIBUTE_DEFINITION_REAL))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, any(fieldInstance).(*ATTRIBUTE_DEFINITION_STRING))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, any(fieldInstance).(*ATTRIBUTE_DEFINITION_XHTML))
						}
					}
				}
			}
		}

	case *SPEC_HIERARCHY:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_HIERARCHY) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_HIERARCHY)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "CHILDREN.SPEC_HIERARCHY" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_HIERARCHY) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_HIERARCHY)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.CHILDREN.SPEC_HIERARCHY).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.CHILDREN.SPEC_HIERARCHY = _inferedTypeInstance.CHILDREN.SPEC_HIERARCHY[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.CHILDREN.SPEC_HIERARCHY =
								append(_inferedTypeInstance.CHILDREN.SPEC_HIERARCHY, any(fieldInstance).(*SPEC_HIERARCHY))
						}
					}
				}
			}
		}

	case *SPEC_OBJECT:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_VALUE_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE, any(fieldInstance).(*ATTRIBUTE_VALUE_DATE))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_VALUE_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER, any(fieldInstance).(*ATTRIBUTE_VALUE_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL, any(fieldInstance).(*ATTRIBUTE_VALUE_REAL))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING, any(fieldInstance).(*ATTRIBUTE_VALUE_STRING))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML, any(fieldInstance).(*ATTRIBUTE_VALUE_XHTML))
						}
					}
				}
			}
		}

	case *SPEC_OBJECT_TYPE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_DEFINITION_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, any(fieldInstance).(*ATTRIBUTE_DEFINITION_DATE))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_DEFINITION_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, any(fieldInstance).(*ATTRIBUTE_DEFINITION_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, any(fieldInstance).(*ATTRIBUTE_DEFINITION_REAL))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, any(fieldInstance).(*ATTRIBUTE_DEFINITION_STRING))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, any(fieldInstance).(*ATTRIBUTE_DEFINITION_XHTML))
						}
					}
				}
			}
		}

	case *SPEC_RELATION:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_VALUE_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_DATE, any(fieldInstance).(*ATTRIBUTE_VALUE_DATE))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_VALUE_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_INTEGER, any(fieldInstance).(*ATTRIBUTE_VALUE_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_REAL, any(fieldInstance).(*ATTRIBUTE_VALUE_REAL))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_STRING, any(fieldInstance).(*ATTRIBUTE_VALUE_STRING))
						}
					}
				}
			}
		}
		if fieldName == "VALUES.ATTRIBUTE_VALUE_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML = _inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML =
								append(_inferedTypeInstance.VALUES.ATTRIBUTE_VALUE_XHTML, any(fieldInstance).(*ATTRIBUTE_VALUE_XHTML))
						}
					}
				}
			}
		}

	case *SPEC_RELATION_TYPE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID.ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_DEFINITION_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, any(fieldInstance).(*ATTRIBUTE_DEFINITION_DATE))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_DEFINITION_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, any(fieldInstance).(*ATTRIBUTE_DEFINITION_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, any(fieldInstance).(*ATTRIBUTE_DEFINITION_REAL))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, any(fieldInstance).(*ATTRIBUTE_DEFINITION_STRING))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = _inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, any(fieldInstance).(*ATTRIBUTE_DEFINITION_XHTML))
						}
					}
				}
			}
		}

	case *XHTML_CONTENT:
		// insertion point per field

	case *Xhtml_InlPres_type:
		// insertion point per field

	case *Xhtml_a_type:
		// insertion point per field

	case *Xhtml_abbr_type:
		// insertion point per field

	case *Xhtml_acronym_type:
		// insertion point per field

	case *Xhtml_address_type:
		// insertion point per field

	case *Xhtml_blockquote_type:
		// insertion point per field

	case *Xhtml_br_type:
		// insertion point per field

	case *Xhtml_caption_type:
		// insertion point per field

	case *Xhtml_cite_type:
		// insertion point per field

	case *Xhtml_code_type:
		// insertion point per field

	case *Xhtml_col_type:
		// insertion point per field

	case *Xhtml_colgroup_type:
		// insertion point per field

	case *Xhtml_dd_type:
		// insertion point per field

	case *Xhtml_dfn_type:
		// insertion point per field

	case *Xhtml_div_type:
		// insertion point per field

	case *Xhtml_dl_type:
		// insertion point per field

	case *Xhtml_dt_type:
		// insertion point per field

	case *Xhtml_edit_type:
		// insertion point per field

	case *Xhtml_em_type:
		// insertion point per field

	case *Xhtml_h1_type:
		// insertion point per field

	case *Xhtml_h2_type:
		// insertion point per field

	case *Xhtml_h3_type:
		// insertion point per field

	case *Xhtml_h4_type:
		// insertion point per field

	case *Xhtml_h5_type:
		// insertion point per field

	case *Xhtml_h6_type:
		// insertion point per field

	case *Xhtml_heading_type:
		// insertion point per field

	case *Xhtml_hr_type:
		// insertion point per field

	case *Xhtml_kbd_type:
		// insertion point per field

	case *Xhtml_li_type:
		// insertion point per field

	case *Xhtml_object_type:
		// insertion point per field

	case *Xhtml_ol_type:
		// insertion point per field

	case *Xhtml_p_type:
		// insertion point per field

	case *Xhtml_param_type:
		// insertion point per field

	case *Xhtml_pre_type:
		// insertion point per field

	case *Xhtml_q_type:
		// insertion point per field

	case *Xhtml_samp_type:
		// insertion point per field

	case *Xhtml_span_type:
		// insertion point per field

	case *Xhtml_strong_type:
		// insertion point per field

	case *Xhtml_table_type:
		// insertion point per field

	case *Xhtml_tbody_type:
		// insertion point per field

	case *Xhtml_td_type:
		// insertion point per field

	case *Xhtml_tfoot_type:
		// insertion point per field

	case *Xhtml_th_type:
		// insertion point per field

	case *Xhtml_thead_type:
		// insertion point per field

	case *Xhtml_tr_type:
		// insertion point per field

	case *Xhtml_ul_type:
		// insertion point per field

	case *Xhtml_var_type:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct ALTERNATIVE_ID
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_BOOLEAN
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_BOOLEAN)
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		_ = attribute_definition_boolean
		for _, _alternative_id := range attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_boolean
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*ATTRIBUTE_DEFINITION_BOOLEAN)
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		_ = attribute_definition_boolean
		for _, _attribute_value_boolean := range attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
			stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = attribute_definition_boolean
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_DATE
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_DATE)
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		_ = attribute_definition_date
		for _, _alternative_id := range attribute_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_date
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*ATTRIBUTE_DEFINITION_DATE)
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		_ = attribute_definition_date
		for _, _attribute_value_date := range attribute_definition_date.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE {
			stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = attribute_definition_date
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_ENUMERATION
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*ATTRIBUTE_DEFINITION_ENUMERATION)
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		_ = attribute_definition_enumeration
		for _, _attribute_value_enumeration := range attribute_definition_enumeration.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION {
			stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = attribute_definition_enumeration
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_ENUMERATION)
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		_ = attribute_definition_enumeration
		for _, _alternative_id := range attribute_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_enumeration
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_INTEGER
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_INTEGER)
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		_ = attribute_definition_integer
		for _, _alternative_id := range attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_integer
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*ATTRIBUTE_DEFINITION_INTEGER)
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		_ = attribute_definition_integer
		for _, _attribute_value_integer := range attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
			stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = attribute_definition_integer
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_REAL
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_REAL)
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		_ = attribute_definition_real
		for _, _alternative_id := range attribute_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_real
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*ATTRIBUTE_DEFINITION_REAL)
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		_ = attribute_definition_real
		for _, _attribute_value_real := range attribute_definition_real.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL {
			stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = attribute_definition_real
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_STRING
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_STRING)
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		_ = attribute_definition_string
		for _, _alternative_id := range attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_string
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*ATTRIBUTE_DEFINITION_STRING)
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		_ = attribute_definition_string
		for _, _attribute_value_string := range attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
			stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = attribute_definition_string
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_XHTML
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_XHTML)
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		_ = attribute_definition_xhtml
		for _, _alternative_id := range attribute_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_xhtml
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*ATTRIBUTE_DEFINITION_XHTML)
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		_ = attribute_definition_xhtml
		for _, _attribute_value_xhtml := range attribute_definition_xhtml.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML {
			stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = attribute_definition_xhtml
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_VALUE_BOOLEAN
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_DATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_INTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_REAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_STRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_XHTML
	// insertion point per field
	clear(stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap)
	stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap = make(map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML)
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		_ = attribute_value_xhtml
		for _, _xhtml_content := range attribute_value_xhtml.THE_VALUE {
			stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap[_xhtml_content] = attribute_value_xhtml
		}
	}
	clear(stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap)
	stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap = make(map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML)
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		_ = attribute_value_xhtml
		for _, _xhtml_content := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
			stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap[_xhtml_content] = attribute_value_xhtml
		}
	}

	// Compute reverse map for named struct AnyType
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_BOOLEAN
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_BOOLEAN)
	for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
		_ = datatype_definition_boolean
		for _, _alternative_id := range datatype_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_boolean
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_DATE
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_DATE)
	for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
		_ = datatype_definition_date
		for _, _alternative_id := range datatype_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_date
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_ENUMERATION
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_ENUMERATION)
	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		_ = datatype_definition_enumeration
		for _, _alternative_id := range datatype_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_enumeration
		}
	}
	clear(stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_ENUM_VALUE_reverseMap)
	stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_ENUM_VALUE_reverseMap = make(map[*ENUM_VALUE]*DATATYPE_DEFINITION_ENUMERATION)
	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		_ = datatype_definition_enumeration
		for _, _enum_value := range datatype_definition_enumeration.SPECIFIED_VALUES.ENUM_VALUE {
			stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[_enum_value] = datatype_definition_enumeration
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_INTEGER
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_INTEGER)
	for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
		_ = datatype_definition_integer
		for _, _alternative_id := range datatype_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_integer
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_REAL
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_REAL)
	for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
		_ = datatype_definition_real
		for _, _alternative_id := range datatype_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_real
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_STRING
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_STRING)
	for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
		_ = datatype_definition_string
		for _, _alternative_id := range datatype_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_string
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_XHTML
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_XHTML)
	for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
		_ = datatype_definition_xhtml
		for _, _alternative_id := range datatype_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_xhtml
		}
	}

	// Compute reverse map for named struct EMBEDDED_VALUE
	// insertion point per field

	// Compute reverse map for named struct ENUM_VALUE
	// insertion point per field
	clear(stage.ENUM_VALUE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ENUM_VALUE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ENUM_VALUE)
	for enum_value := range stage.ENUM_VALUEs {
		_ = enum_value
		for _, _alternative_id := range enum_value.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ENUM_VALUE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = enum_value
		}
	}
	clear(stage.ENUM_VALUE_PROPERTIES_EMBEDDED_VALUE_reverseMap)
	stage.ENUM_VALUE_PROPERTIES_EMBEDDED_VALUE_reverseMap = make(map[*EMBEDDED_VALUE]*ENUM_VALUE)
	for enum_value := range stage.ENUM_VALUEs {
		_ = enum_value
		for _, _embedded_value := range enum_value.PROPERTIES.EMBEDDED_VALUE {
			stage.ENUM_VALUE_PROPERTIES_EMBEDDED_VALUE_reverseMap[_embedded_value] = enum_value
		}
	}

	// Compute reverse map for named struct RELATION_GROUP
	// insertion point per field
	clear(stage.RELATION_GROUP_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.RELATION_GROUP_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*RELATION_GROUP)
	for relation_group := range stage.RELATION_GROUPs {
		_ = relation_group
		for _, _alternative_id := range relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.RELATION_GROUP_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = relation_group
		}
	}

	// Compute reverse map for named struct RELATION_GROUP_TYPE
	// insertion point per field
	clear(stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _alternative_id := range relation_group_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_boolean := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_date := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_enumeration := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_integer := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_real := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_string := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_xhtml := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = relation_group_type
		}
	}

	// Compute reverse map for named struct REQ_IF
	// insertion point per field
	clear(stage.REQ_IF_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap)
	stage.REQ_IF_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap = make(map[*REQ_IF_TOOL_EXTENSION]*REQ_IF)
	for req_if := range stage.REQ_IFs {
		_ = req_if
		for _, _req_if_tool_extension := range req_if.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION {
			stage.REQ_IF_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[_req_if_tool_extension] = req_if
		}
	}

	// Compute reverse map for named struct REQ_IF_CONTENT
	// insertion point per field
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap = make(map[*DATATYPE_DEFINITION_BOOLEAN]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_boolean := range req_if_content.DATATYPES.DATATYPE_DEFINITION_BOOLEAN {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[_datatype_definition_boolean] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap = make(map[*DATATYPE_DEFINITION_DATE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_date := range req_if_content.DATATYPES.DATATYPE_DEFINITION_DATE {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[_datatype_definition_date] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap = make(map[*DATATYPE_DEFINITION_ENUMERATION]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_enumeration := range req_if_content.DATATYPES.DATATYPE_DEFINITION_ENUMERATION {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[_datatype_definition_enumeration] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap = make(map[*DATATYPE_DEFINITION_INTEGER]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_integer := range req_if_content.DATATYPES.DATATYPE_DEFINITION_INTEGER {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[_datatype_definition_integer] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap = make(map[*DATATYPE_DEFINITION_REAL]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_real := range req_if_content.DATATYPES.DATATYPE_DEFINITION_REAL {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[_datatype_definition_real] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap = make(map[*DATATYPE_DEFINITION_STRING]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_string := range req_if_content.DATATYPES.DATATYPE_DEFINITION_STRING {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[_datatype_definition_string] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap = make(map[*DATATYPE_DEFINITION_XHTML]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_xhtml := range req_if_content.DATATYPES.DATATYPE_DEFINITION_XHTML {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[_datatype_definition_xhtml] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap = make(map[*RELATION_GROUP_TYPE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _relation_group_type := range req_if_content.SPEC_TYPES.RELATION_GROUP_TYPE {
			stage.REQ_IF_CONTENT_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[_relation_group_type] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap = make(map[*SPEC_OBJECT_TYPE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _spec_object_type := range req_if_content.SPEC_TYPES.SPEC_OBJECT_TYPE {
			stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[_spec_object_type] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap = make(map[*SPEC_RELATION_TYPE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _spec_relation_type := range req_if_content.SPEC_TYPES.SPEC_RELATION_TYPE {
			stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[_spec_relation_type] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap = make(map[*SPECIFICATION_TYPE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _specification_type := range req_if_content.SPEC_TYPES.SPECIFICATION_TYPE {
			stage.REQ_IF_CONTENT_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[_specification_type] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_OBJECTS_SPEC_OBJECT_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_OBJECTS_SPEC_OBJECT_reverseMap = make(map[*SPEC_OBJECT]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _spec_object := range req_if_content.SPEC_OBJECTS.SPEC_OBJECT {
			stage.REQ_IF_CONTENT_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[_spec_object] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_RELATIONS_SPEC_RELATION_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_RELATIONS_SPEC_RELATION_reverseMap = make(map[*SPEC_RELATION]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _spec_relation := range req_if_content.SPEC_RELATIONS.SPEC_RELATION {
			stage.REQ_IF_CONTENT_SPEC_RELATIONS_SPEC_RELATION_reverseMap[_spec_relation] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPECIFICATIONS_SPECIFICATION_reverseMap)
	stage.REQ_IF_CONTENT_SPECIFICATIONS_SPECIFICATION_reverseMap = make(map[*SPECIFICATION]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _specification := range req_if_content.SPECIFICATIONS.SPECIFICATION {
			stage.REQ_IF_CONTENT_SPECIFICATIONS_SPECIFICATION_reverseMap[_specification] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap = make(map[*RELATION_GROUP]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _relation_group := range req_if_content.SPEC_RELATION_GROUPS.RELATION_GROUP {
			stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[_relation_group] = req_if_content
		}
	}

	// Compute reverse map for named struct REQ_IF_HEADER
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_TOOL_EXTENSION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION
	// insertion point per field
	clear(stage.SPECIFICATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPECIFICATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _alternative_id := range specification.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPECIFICATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_boolean := range specification.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_date := range specification.VALUES.ATTRIBUTE_VALUE_DATE {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_enumeration := range specification.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_integer := range specification.VALUES.ATTRIBUTE_VALUE_INTEGER {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_real := range specification.VALUES.ATTRIBUTE_VALUE_REAL {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_string := range specification.VALUES.ATTRIBUTE_VALUE_STRING {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_xhtml := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = specification
		}
	}
	clear(stage.SPECIFICATION_CHILDREN_SPEC_HIERARCHY_reverseMap)
	stage.SPECIFICATION_CHILDREN_SPEC_HIERARCHY_reverseMap = make(map[*SPEC_HIERARCHY]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _spec_hierarchy := range specification.CHILDREN.SPEC_HIERARCHY {
			stage.SPECIFICATION_CHILDREN_SPEC_HIERARCHY_reverseMap[_spec_hierarchy] = specification
		}
	}

	// Compute reverse map for named struct SPECIFICATION_TYPE
	// insertion point per field
	clear(stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _alternative_id := range specification_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_boolean := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_date := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_enumeration := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_integer := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_real := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_string := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_xhtml := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = specification_type
		}
	}

	// Compute reverse map for named struct SPEC_HIERARCHY
	// insertion point per field
	clear(stage.SPEC_HIERARCHY_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_HIERARCHY_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPEC_HIERARCHY)
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		_ = spec_hierarchy
		for _, _alternative_id := range spec_hierarchy.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPEC_HIERARCHY_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = spec_hierarchy
		}
	}
	clear(stage.SPEC_HIERARCHY_CHILDREN_SPEC_HIERARCHY_reverseMap)
	stage.SPEC_HIERARCHY_CHILDREN_SPEC_HIERARCHY_reverseMap = make(map[*SPEC_HIERARCHY]*SPEC_HIERARCHY)
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		_ = spec_hierarchy
		for _, _spec_hierarchy := range spec_hierarchy.CHILDREN.SPEC_HIERARCHY {
			stage.SPEC_HIERARCHY_CHILDREN_SPEC_HIERARCHY_reverseMap[_spec_hierarchy] = spec_hierarchy
		}
	}

	// Compute reverse map for named struct SPEC_OBJECT
	// insertion point per field
	clear(stage.SPEC_OBJECT_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_OBJECT_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _alternative_id := range spec_object.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPEC_OBJECT_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_boolean := range spec_object.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_date := range spec_object.VALUES.ATTRIBUTE_VALUE_DATE {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_enumeration := range spec_object.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_integer := range spec_object.VALUES.ATTRIBUTE_VALUE_INTEGER {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_real := range spec_object.VALUES.ATTRIBUTE_VALUE_REAL {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_string := range spec_object.VALUES.ATTRIBUTE_VALUE_STRING {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_xhtml := range spec_object.VALUES.ATTRIBUTE_VALUE_XHTML {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = spec_object
		}
	}

	// Compute reverse map for named struct SPEC_OBJECT_TYPE
	// insertion point per field
	clear(stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _alternative_id := range spec_object_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_boolean := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_date := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_enumeration := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_integer := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_real := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_string := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_xhtml := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = spec_object_type
		}
	}

	// Compute reverse map for named struct SPEC_RELATION
	// insertion point per field
	clear(stage.SPEC_RELATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_RELATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _alternative_id := range spec_relation.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPEC_RELATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_boolean := range spec_relation.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_date := range spec_relation.VALUES.ATTRIBUTE_VALUE_DATE {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_enumeration := range spec_relation.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_integer := range spec_relation.VALUES.ATTRIBUTE_VALUE_INTEGER {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_real := range spec_relation.VALUES.ATTRIBUTE_VALUE_REAL {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_string := range spec_relation.VALUES.ATTRIBUTE_VALUE_STRING {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_xhtml := range spec_relation.VALUES.ATTRIBUTE_VALUE_XHTML {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = spec_relation
		}
	}

	// Compute reverse map for named struct SPEC_RELATION_TYPE
	// insertion point per field
	clear(stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _alternative_id := range spec_relation_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_boolean := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_date := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_enumeration := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_integer := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_real := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_string := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_xhtml := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = spec_relation_type
		}
	}

	// Compute reverse map for named struct XHTML_CONTENT
	// insertion point per field

	// Compute reverse map for named struct Xhtml_InlPres_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_a_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_abbr_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_acronym_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_address_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_blockquote_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_br_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_caption_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_cite_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_code_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_col_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_colgroup_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_dd_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_dfn_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_div_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_dl_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_dt_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_edit_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_em_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h1_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h2_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h3_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h4_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h5_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h6_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_heading_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_hr_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_kbd_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_li_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_object_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_ol_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_p_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_param_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_pre_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_q_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_samp_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_span_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_strong_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_table_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_tbody_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_td_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_tfoot_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_th_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_thead_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_tr_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_ul_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_var_type
	// insertion point per field

}
