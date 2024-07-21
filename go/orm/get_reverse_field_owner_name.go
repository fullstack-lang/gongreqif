// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongreqif/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _attribute_definition_boolean, ok := stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_boolean.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _attribute_definition_date, ok := stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_date.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _attribute_definition_enumeration, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_enumeration.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _attribute_definition_integer, ok := stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_integer.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _attribute_definition_real, ok := stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_real.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _attribute_definition_string, ok := stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_string.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _attribute_definition_xhtml, ok := stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_xhtml.Name
				}
			}
		case "DATATYPE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _datatype_definition_boolean, ok := stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_boolean.Name
				}
			}
		case "DATATYPE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _datatype_definition_date, ok := stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_date.Name
				}
			}
		case "DATATYPE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _datatype_definition_enumeration, ok := stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_enumeration.Name
				}
			}
		case "DATATYPE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _datatype_definition_integer, ok := stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_integer.Name
				}
			}
		case "DATATYPE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _datatype_definition_real, ok := stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_real.Name
				}
			}
		case "DATATYPE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _datatype_definition_string, ok := stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_string.Name
				}
			}
		case "DATATYPE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _datatype_definition_xhtml, ok := stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_xhtml.Name
				}
			}
		case "ENUM_VALUE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _enum_value, ok := stage.ENUM_VALUE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _enum_value.Name
				}
			}
		case "RELATION_GROUP":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _relation_group, ok := stage.RELATION_GROUP_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _relation_group.Name
				}
			}
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _relation_group_type, ok := stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _relation_group_type.Name
				}
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _specification, ok := stage.SPECIFICATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _specification_type, ok := stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _specification_type.Name
				}
			}
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _spec_hierarchy, ok := stage.SPEC_HIERARCHY_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _spec_hierarchy.Name
				}
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _spec_object, ok := stage.SPEC_OBJECT_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _spec_object_type, ok := stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _spec_object_type.Name
				}
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _spec_relation, ok := stage.SPEC_RELATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				if _spec_relation_type, ok := stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _spec_relation_type.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN":
				if _relation_group_type, ok := stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]; ok {
					res = _relation_group_type.Name
				}
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN":
				if _specification_type, ok := stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]; ok {
					res = _specification_type.Name
				}
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN":
				if _spec_object_type, ok := stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]; ok {
					res = _spec_object_type.Name
				}
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN":
				if _spec_relation_type, ok := stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]; ok {
					res = _spec_relation_type.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE":
				if _relation_group_type, ok := stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]; ok {
					res = _relation_group_type.Name
				}
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE":
				if _specification_type, ok := stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]; ok {
					res = _specification_type.Name
				}
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE":
				if _spec_object_type, ok := stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]; ok {
					res = _spec_object_type.Name
				}
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE":
				if _spec_relation_type, ok := stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]; ok {
					res = _spec_relation_type.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION":
				if _relation_group_type, ok := stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]; ok {
					res = _relation_group_type.Name
				}
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION":
				if _specification_type, ok := stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]; ok {
					res = _specification_type.Name
				}
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION":
				if _spec_object_type, ok := stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]; ok {
					res = _spec_object_type.Name
				}
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION":
				if _spec_relation_type, ok := stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]; ok {
					res = _spec_relation_type.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER":
				if _relation_group_type, ok := stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]; ok {
					res = _relation_group_type.Name
				}
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER":
				if _specification_type, ok := stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]; ok {
					res = _specification_type.Name
				}
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER":
				if _spec_object_type, ok := stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]; ok {
					res = _spec_object_type.Name
				}
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER":
				if _spec_relation_type, ok := stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]; ok {
					res = _spec_relation_type.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL":
				if _relation_group_type, ok := stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]; ok {
					res = _relation_group_type.Name
				}
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL":
				if _specification_type, ok := stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]; ok {
					res = _specification_type.Name
				}
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL":
				if _spec_object_type, ok := stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]; ok {
					res = _spec_object_type.Name
				}
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL":
				if _spec_relation_type, ok := stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]; ok {
					res = _spec_relation_type.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING":
				if _relation_group_type, ok := stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]; ok {
					res = _relation_group_type.Name
				}
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING":
				if _specification_type, ok := stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]; ok {
					res = _specification_type.Name
				}
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING":
				if _spec_object_type, ok := stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]; ok {
					res = _spec_object_type.Name
				}
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING":
				if _spec_relation_type, ok := stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]; ok {
					res = _spec_relation_type.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML":
				if _relation_group_type, ok := stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]; ok {
					res = _relation_group_type.Name
				}
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML":
				if _specification_type, ok := stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]; ok {
					res = _specification_type.Name
				}
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML":
				if _spec_object_type, ok := stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]; ok {
					res = _spec_object_type.Name
				}
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML":
				if _spec_relation_type, ok := stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]; ok {
					res = _spec_relation_type.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN":
				if _attribute_definition_boolean, ok := stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]; ok {
					res = _attribute_definition_boolean.Name
				}
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_BOOLEAN":
				if _specification, ok := stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_BOOLEAN":
				if _spec_object, ok := stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_BOOLEAN":
				if _spec_relation, ok := stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE":
				if _attribute_definition_date, ok := stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE_reverseMap[inst]; ok {
					res = _attribute_definition_date.Name
				}
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_DATE":
				if _specification, ok := stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_DATE":
				if _spec_object, ok := stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_DATE":
				if _spec_relation, ok := stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION":
				if _attribute_definition_enumeration, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]; ok {
					res = _attribute_definition_enumeration.Name
				}
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_ENUMERATION":
				if _specification, ok := stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_ENUMERATION":
				if _spec_object, ok := stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_ENUMERATION":
				if _spec_relation, ok := stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER":
				if _attribute_definition_integer, ok := stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]; ok {
					res = _attribute_definition_integer.Name
				}
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_INTEGER":
				if _specification, ok := stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_INTEGER":
				if _spec_object, ok := stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_INTEGER":
				if _spec_relation, ok := stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL":
				if _attribute_definition_real, ok := stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL_reverseMap[inst]; ok {
					res = _attribute_definition_real.Name
				}
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_REAL":
				if _specification, ok := stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_REAL":
				if _spec_object, ok := stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_REAL":
				if _spec_relation, ok := stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING":
				if _attribute_definition_string, ok := stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING_reverseMap[inst]; ok {
					res = _attribute_definition_string.Name
				}
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_STRING":
				if _specification, ok := stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_STRING":
				if _spec_object, ok := stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_STRING":
				if _spec_relation, ok := stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML":
				if _attribute_definition_xhtml, ok := stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]; ok {
					res = _attribute_definition_xhtml.Name
				}
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_XHTML":
				if _specification, ok := stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_XHTML":
				if _spec_object, ok := stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_XHTML":
				if _spec_relation, ok := stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
		}

	case *models.AnyType:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_BOOLEAN":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_DATE":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_ENUMERATION":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_INTEGER":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_REAL":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_STRING":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_XHTML":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.EMBEDDED_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		case "ENUM_VALUE":
			switch reverseField.Fieldname {
			case "PROPERTIES_EMBEDDED_VALUE":
				if _enum_value, ok := stage.ENUM_VALUE_PROPERTIES_EMBEDDED_VALUE_reverseMap[inst]; ok {
					res = _enum_value.Name
				}
			}
		}

	case *models.ENUM_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "SPECIFIED_VALUES_ENUM_VALUE":
				if _datatype_definition_enumeration, ok := stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[inst]; ok {
					res = _datatype_definition_enumeration.Name
				}
			}
		}

	case *models.RELATION_GROUP:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_RELATION_GROUPS_RELATION_GROUP":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.RELATION_GROUP_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_TYPES_RELATION_GROUP_TYPE":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.REQ_IF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_TOOL_EXTENSION:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF":
			switch reverseField.Fieldname {
			case "TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION":
				if _req_if, ok := stage.REQ_IF_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[inst]; ok {
					res = _req_if.Name
				}
			}
		}

	case *models.SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPECIFICATIONS_SPECIFICATION":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPECIFICATIONS_SPECIFICATION_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.SPECIFICATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_TYPES_SPECIFICATION_TYPE":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.SPEC_HIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "CHILDREN_SPEC_HIERARCHY":
				if _specification, ok := stage.SPECIFICATION_CHILDREN_SPEC_HIERARCHY_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "CHILDREN_SPEC_HIERARCHY":
				if _spec_hierarchy, ok := stage.SPEC_HIERARCHY_CHILDREN_SPEC_HIERARCHY_reverseMap[inst]; ok {
					res = _spec_hierarchy.Name
				}
			}
		}

	case *models.SPEC_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_OBJECTS_SPEC_OBJECT":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.SPEC_OBJECT_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_TYPES_SPEC_OBJECT_TYPE":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.SPEC_RELATION:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_RELATIONS_SPEC_RELATION":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_RELATIONS_SPEC_RELATION_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.SPEC_RELATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_TYPES_SPEC_RELATION_TYPE":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.XHTML_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_XHTML":
			switch reverseField.Fieldname {
			case "THE_VALUE":
				if _attribute_value_xhtml, ok := stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap[inst]; ok {
					res = _attribute_value_xhtml.Name
				}
			case "THE_ORIGINAL_VALUE":
				if _attribute_value_xhtml, ok := stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap[inst]; ok {
					res = _attribute_value_xhtml.Name
				}
			}
		}

	case *models.Xhtml_InlPres_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_a_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_abbr_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_acronym_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_address_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_blockquote_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_br_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_caption_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_cite_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_code_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_col_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_colgroup_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_dd_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_dfn_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_div_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_dl_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_dt_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_edit_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_em_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h1_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h2_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h3_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h4_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h5_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h6_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_heading_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_hr_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_kbd_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_li_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_object_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_ol_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_p_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_param_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_pre_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_q_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_samp_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_span_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_strong_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_table_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_tbody_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_td_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_tfoot_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_th_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_thead_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_tr_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_ul_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_var_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ENUM_VALUE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.ENUM_VALUE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "RELATION_GROUP":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.RELATION_GROUP_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.SPECIFICATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.SPEC_HIERARCHY_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.SPEC_OBJECT_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.SPEC_RELATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID_ALTERNATIVE_ID":
				res = stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN":
				res = stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN":
				res = stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN":
				res = stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN":
				res = stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE":
				res = stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE":
				res = stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE":
				res = stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE":
				res = stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION":
				res = stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION":
				res = stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION":
				res = stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION":
				res = stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER":
				res = stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER":
				res = stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER":
				res = stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER":
				res = stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL":
				res = stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL":
				res = stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL":
				res = stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL":
				res = stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING":
				res = stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING":
				res = stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING":
				res = stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING":
				res = stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML":
				res = stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML":
				res = stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML":
				res = stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML":
				res = stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN":
				res = stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_BOOLEAN":
				res = stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_BOOLEAN":
				res = stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_BOOLEAN":
				res = stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE":
				res = stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE_reverseMap[inst]
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_DATE":
				res = stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[inst]
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_DATE":
				res = stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[inst]
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_DATE":
				res = stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION":
				res = stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_ENUMERATION":
				res = stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_ENUMERATION":
				res = stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_ENUMERATION":
				res = stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER":
				res = stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_INTEGER":
				res = stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_INTEGER":
				res = stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_INTEGER":
				res = stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL":
				res = stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL_reverseMap[inst]
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_REAL":
				res = stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[inst]
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_REAL":
				res = stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[inst]
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_REAL":
				res = stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING":
				res = stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING_reverseMap[inst]
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_STRING":
				res = stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[inst]
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_STRING":
				res = stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[inst]
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_STRING":
				res = stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML":
				res = stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_XHTML":
				res = stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_XHTML":
				res = stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES_ATTRIBUTE_VALUE_XHTML":
				res = stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]
			}
		}

	case *models.AnyType:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_BOOLEAN":
				res = stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_DATE":
				res = stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_ENUMERATION":
				res = stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_INTEGER":
				res = stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_REAL":
				res = stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_STRING":
				res = stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES_DATATYPE_DEFINITION_XHTML":
				res = stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[inst]
			}
		}

	case *models.EMBEDDED_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		case "ENUM_VALUE":
			switch reverseField.Fieldname {
			case "PROPERTIES_EMBEDDED_VALUE":
				res = stage.ENUM_VALUE_PROPERTIES_EMBEDDED_VALUE_reverseMap[inst]
			}
		}

	case *models.ENUM_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "SPECIFIED_VALUES_ENUM_VALUE":
				res = stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[inst]
			}
		}

	case *models.RELATION_GROUP:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_RELATION_GROUPS_RELATION_GROUP":
				res = stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[inst]
			}
		}

	case *models.RELATION_GROUP_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_TYPES_RELATION_GROUP_TYPE":
				res = stage.REQ_IF_CONTENT_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[inst]
			}
		}

	case *models.REQ_IF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_TOOL_EXTENSION:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF":
			switch reverseField.Fieldname {
			case "TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION":
				res = stage.REQ_IF_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[inst]
			}
		}

	case *models.SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPECIFICATIONS_SPECIFICATION":
				res = stage.REQ_IF_CONTENT_SPECIFICATIONS_SPECIFICATION_reverseMap[inst]
			}
		}

	case *models.SPECIFICATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_TYPES_SPECIFICATION_TYPE":
				res = stage.REQ_IF_CONTENT_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[inst]
			}
		}

	case *models.SPEC_HIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "CHILDREN_SPEC_HIERARCHY":
				res = stage.SPECIFICATION_CHILDREN_SPEC_HIERARCHY_reverseMap[inst]
			}
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "CHILDREN_SPEC_HIERARCHY":
				res = stage.SPEC_HIERARCHY_CHILDREN_SPEC_HIERARCHY_reverseMap[inst]
			}
		}

	case *models.SPEC_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_OBJECTS_SPEC_OBJECT":
				res = stage.REQ_IF_CONTENT_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[inst]
			}
		}

	case *models.SPEC_OBJECT_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_TYPES_SPEC_OBJECT_TYPE":
				res = stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[inst]
			}
		}

	case *models.SPEC_RELATION:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_RELATIONS_SPEC_RELATION":
				res = stage.REQ_IF_CONTENT_SPEC_RELATIONS_SPEC_RELATION_reverseMap[inst]
			}
		}

	case *models.SPEC_RELATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_TYPES_SPEC_RELATION_TYPE":
				res = stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[inst]
			}
		}

	case *models.XHTML_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_XHTML":
			switch reverseField.Fieldname {
			case "THE_VALUE":
				res = stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap[inst]
			case "THE_ORIGINAL_VALUE":
				res = stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap[inst]
			}
		}

	case *models.Xhtml_InlPres_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_a_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_abbr_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_acronym_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_address_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_blockquote_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_br_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_caption_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_cite_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_code_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_col_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_colgroup_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_dd_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_dfn_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_div_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_dl_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_dt_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_edit_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_em_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h1_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h2_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h3_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h4_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h5_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_h6_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_heading_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_hr_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_kbd_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_li_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_object_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_ol_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_p_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_param_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_pre_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_q_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_samp_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_span_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_strong_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_table_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_tbody_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_td_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_tfoot_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_th_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_thead_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_tr_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_ul_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Xhtml_var_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
