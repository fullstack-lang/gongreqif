// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_BOOLEAN),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_BOOLEAN, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_DATE),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_DATE, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_ENUMERATION),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_ENUMERATION, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_INTEGER),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_INTEGER, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_REAL),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_REAL, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_STRING),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_STRING, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_XHTML),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_XHTML, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_BOOLEAN"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_BOOLEAN),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_BOOLEAN, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_DATE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_DATE),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_DATE, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_ENUMERATION"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_ENUMERATION),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_ENUMERATION, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_INTEGER"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_INTEGER),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_INTEGER, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_REAL"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_REAL),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_REAL, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_STRING"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_STRING),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_STRING, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_XHTML"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_XHTML),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_XHTML, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ENUM_VALUE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ENUM_VALUE),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ENUM_VALUE, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP_TYPE),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP_TYPE, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION_TYPE),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION_TYPE, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_HIERARCHY),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_HIERARCHY, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT_TYPE),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT_TYPE, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION_TYPE),
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION_TYPE, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID.ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP_TYPE, *models.ATTRIBUTE_DEFINITION_BOOLEAN](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION_TYPE, *models.ATTRIBUTE_DEFINITION_BOOLEAN](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT_TYPE, *models.ATTRIBUTE_DEFINITION_BOOLEAN](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION_TYPE, *models.ATTRIBUTE_DEFINITION_BOOLEAN](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP_TYPE, *models.ATTRIBUTE_DEFINITION_DATE](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION_TYPE, *models.ATTRIBUTE_DEFINITION_DATE](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT_TYPE, *models.ATTRIBUTE_DEFINITION_DATE](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION_TYPE, *models.ATTRIBUTE_DEFINITION_DATE](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MULTI_VALUED", instanceWithInferedType.MULTI_VALUED, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP_TYPE, *models.ATTRIBUTE_DEFINITION_ENUMERATION](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION_TYPE, *models.ATTRIBUTE_DEFINITION_ENUMERATION](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT_TYPE, *models.ATTRIBUTE_DEFINITION_ENUMERATION](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION_TYPE, *models.ATTRIBUTE_DEFINITION_ENUMERATION](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP_TYPE, *models.ATTRIBUTE_DEFINITION_INTEGER](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION_TYPE, *models.ATTRIBUTE_DEFINITION_INTEGER](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT_TYPE, *models.ATTRIBUTE_DEFINITION_INTEGER](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION_TYPE, *models.ATTRIBUTE_DEFINITION_INTEGER](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP_TYPE, *models.ATTRIBUTE_DEFINITION_REAL](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION_TYPE, *models.ATTRIBUTE_DEFINITION_REAL](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT_TYPE, *models.ATTRIBUTE_DEFINITION_REAL](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION_TYPE, *models.ATTRIBUTE_DEFINITION_REAL](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP_TYPE, *models.ATTRIBUTE_DEFINITION_STRING](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION_TYPE, *models.ATTRIBUTE_DEFINITION_STRING](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT_TYPE, *models.ATTRIBUTE_DEFINITION_STRING](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION_TYPE, *models.ATTRIBUTE_DEFINITION_STRING](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP_TYPE, *models.ATTRIBUTE_DEFINITION_XHTML](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION_TYPE, *models.ATTRIBUTE_DEFINITION_XHTML](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT_TYPE, *models.ATTRIBUTE_DEFINITION_XHTML](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION_TYPE),
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION_TYPE, *models.ATTRIBUTE_DEFINITION_XHTML](
					nil,
					"SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_BOOLEAN),
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_BOOLEAN, *models.ATTRIBUTE_VALUE_BOOLEAN](
					nil,
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"VALUES.ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.ATTRIBUTE_VALUE_BOOLEAN](
					nil,
					"VALUES.ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"VALUES.ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.ATTRIBUTE_VALUE_BOOLEAN](
					nil,
					"VALUES.ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"VALUES.ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.ATTRIBUTE_VALUE_BOOLEAN](
					nil,
					"VALUES.ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_DATE),
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_DATE, *models.ATTRIBUTE_VALUE_DATE](
					nil,
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"VALUES.ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.ATTRIBUTE_VALUE_DATE](
					nil,
					"VALUES.ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"VALUES.ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.ATTRIBUTE_VALUE_DATE](
					nil,
					"VALUES.ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"VALUES.ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.ATTRIBUTE_VALUE_DATE](
					nil,
					"VALUES.ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_ENUMERATION),
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_ENUMERATION, *models.ATTRIBUTE_VALUE_ENUMERATION](
					nil,
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"VALUES.ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.ATTRIBUTE_VALUE_ENUMERATION](
					nil,
					"VALUES.ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"VALUES.ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.ATTRIBUTE_VALUE_ENUMERATION](
					nil,
					"VALUES.ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"VALUES.ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.ATTRIBUTE_VALUE_ENUMERATION](
					nil,
					"VALUES.ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_INTEGER),
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_INTEGER, *models.ATTRIBUTE_VALUE_INTEGER](
					nil,
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"VALUES.ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.ATTRIBUTE_VALUE_INTEGER](
					nil,
					"VALUES.ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"VALUES.ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.ATTRIBUTE_VALUE_INTEGER](
					nil,
					"VALUES.ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"VALUES.ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.ATTRIBUTE_VALUE_INTEGER](
					nil,
					"VALUES.ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_REAL),
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_REAL, *models.ATTRIBUTE_VALUE_REAL](
					nil,
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"VALUES.ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.ATTRIBUTE_VALUE_REAL](
					nil,
					"VALUES.ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"VALUES.ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.ATTRIBUTE_VALUE_REAL](
					nil,
					"VALUES.ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"VALUES.ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.ATTRIBUTE_VALUE_REAL](
					nil,
					"VALUES.ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_STRING),
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_STRING, *models.ATTRIBUTE_VALUE_STRING](
					nil,
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"VALUES.ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.ATTRIBUTE_VALUE_STRING](
					nil,
					"VALUES.ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"VALUES.ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.ATTRIBUTE_VALUE_STRING](
					nil,
					"VALUES.ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"VALUES.ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.ATTRIBUTE_VALUE_STRING](
					nil,
					"VALUES.ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("THE_VALUE", instanceWithInferedType, &instanceWithInferedType.THE_VALUE, formGroup, probe)
		AssociationSliceToForm("THE_ORIGINAL_VALUE", instanceWithInferedType, &instanceWithInferedType.THE_ORIGINAL_VALUE, formGroup, probe)
		BasicFieldtoForm("IS_SIMPLIFIED", instanceWithInferedType.IS_SIMPLIFIED, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_XHTML),
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_XHTML, *models.ATTRIBUTE_VALUE_XHTML](
					nil,
					"DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"VALUES.ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.ATTRIBUTE_VALUE_XHTML](
					nil,
					"VALUES.ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"VALUES.ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.ATTRIBUTE_VALUE_XHTML](
					nil,
					"VALUES.ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"VALUES.ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.ATTRIBUTE_VALUE_XHTML](
					nil,
					"VALUES.ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.AnyType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InnerXML", instanceWithInferedType.InnerXML, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.DATATYPE_DEFINITION_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_BOOLEAN"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"DATATYPES.DATATYPE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.DATATYPE_DEFINITION_BOOLEAN](
					nil,
					"DATATYPES.DATATYPE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_DATE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"DATATYPES.DATATYPE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.DATATYPE_DEFINITION_DATE](
					nil,
					"DATATYPES.DATATYPE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SPECIFIED_VALUES.ENUM_VALUE", instanceWithInferedType, &instanceWithInferedType.SPECIFIED_VALUES.ENUM_VALUE, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_ENUMERATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"DATATYPES.DATATYPE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.DATATYPE_DEFINITION_ENUMERATION](
					nil,
					"DATATYPES.DATATYPE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_INTEGER"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"DATATYPES.DATATYPE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.DATATYPE_DEFINITION_INTEGER](
					nil,
					"DATATYPES.DATATYPE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MAX", instanceWithInferedType.MAX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MIN", instanceWithInferedType.MIN, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_REAL"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"DATATYPES.DATATYPE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.DATATYPE_DEFINITION_REAL](
					nil,
					"DATATYPES.DATATYPE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_STRING"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"DATATYPES.DATATYPE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.DATATYPE_DEFINITION_STRING](
					nil,
					"DATATYPES.DATATYPE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_XHTML"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"DATATYPES.DATATYPE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.DATATYPE_DEFINITION_XHTML](
					nil,
					"DATATYPES.DATATYPE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.EMBEDDED_VALUE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OTHER_CONTENT", instanceWithInferedType.OTHER_CONTENT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ENUM_VALUE"
			rf.Fieldname = "PROPERTIES.EMBEDDED_VALUE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ENUM_VALUE),
					"PROPERTIES.EMBEDDED_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ENUM_VALUE, *models.EMBEDDED_VALUE](
					nil,
					"PROPERTIES.EMBEDDED_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ENUM_VALUE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("PROPERTIES.EMBEDDED_VALUE", instanceWithInferedType, &instanceWithInferedType.PROPERTIES.EMBEDDED_VALUE, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_ENUMERATION"
			rf.Fieldname = "SPECIFIED_VALUES.ENUM_VALUE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_ENUMERATION),
					"SPECIFIED_VALUES.ENUM_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_ENUMERATION, *models.ENUM_VALUE](
					nil,
					"SPECIFIED_VALUES.ENUM_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.RELATION_GROUP:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_RELATION_GROUPS.RELATION_GROUP"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_RELATION_GROUPS.RELATION_GROUP",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.RELATION_GROUP](
					nil,
					"SPEC_RELATION_GROUPS.RELATION_GROUP",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.RELATION_GROUP_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_TYPES.RELATION_GROUP_TYPE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_TYPES.RELATION_GROUP_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.RELATION_GROUP_TYPE](
					nil,
					"SPEC_TYPES.RELATION_GROUP_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.REQ_IF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("THE_HEADER.REQ_IF_HEADER", instanceWithInferedType.THE_HEADER.REQ_IF_HEADER, formGroup, probe)
		AssociationFieldToForm("CORE_CONTENT.REQ_IF_CONTENT", instanceWithInferedType.CORE_CONTENT.REQ_IF_CONTENT, formGroup, probe)
		AssociationSliceToForm("TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION", instanceWithInferedType, &instanceWithInferedType.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION, formGroup, probe)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.REQ_IF_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DATATYPES.DATATYPE_DEFINITION_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.DATATYPES.DATATYPE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("DATATYPES.DATATYPE_DEFINITION_DATE", instanceWithInferedType, &instanceWithInferedType.DATATYPES.DATATYPE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("DATATYPES.DATATYPE_DEFINITION_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.DATATYPES.DATATYPE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("DATATYPES.DATATYPE_DEFINITION_INTEGER", instanceWithInferedType, &instanceWithInferedType.DATATYPES.DATATYPE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("DATATYPES.DATATYPE_DEFINITION_REAL", instanceWithInferedType, &instanceWithInferedType.DATATYPES.DATATYPE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("DATATYPES.DATATYPE_DEFINITION_STRING", instanceWithInferedType, &instanceWithInferedType.DATATYPES.DATATYPE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("DATATYPES.DATATYPE_DEFINITION_XHTML", instanceWithInferedType, &instanceWithInferedType.DATATYPES.DATATYPE_DEFINITION_XHTML, formGroup, probe)
		AssociationSliceToForm("SPEC_TYPES.RELATION_GROUP_TYPE", instanceWithInferedType, &instanceWithInferedType.SPEC_TYPES.RELATION_GROUP_TYPE, formGroup, probe)
		AssociationSliceToForm("SPEC_TYPES.SPEC_OBJECT_TYPE", instanceWithInferedType, &instanceWithInferedType.SPEC_TYPES.SPEC_OBJECT_TYPE, formGroup, probe)
		AssociationSliceToForm("SPEC_TYPES.SPEC_RELATION_TYPE", instanceWithInferedType, &instanceWithInferedType.SPEC_TYPES.SPEC_RELATION_TYPE, formGroup, probe)
		AssociationSliceToForm("SPEC_TYPES.SPECIFICATION_TYPE", instanceWithInferedType, &instanceWithInferedType.SPEC_TYPES.SPECIFICATION_TYPE, formGroup, probe)
		AssociationSliceToForm("SPEC_OBJECTS.SPEC_OBJECT", instanceWithInferedType, &instanceWithInferedType.SPEC_OBJECTS.SPEC_OBJECT, formGroup, probe)
		AssociationSliceToForm("SPEC_RELATIONS.SPEC_RELATION", instanceWithInferedType, &instanceWithInferedType.SPEC_RELATIONS.SPEC_RELATION, formGroup, probe)
		AssociationSliceToForm("SPECIFICATIONS.SPECIFICATION", instanceWithInferedType, &instanceWithInferedType.SPECIFICATIONS.SPECIFICATION, formGroup, probe)
		AssociationSliceToForm("SPEC_RELATION_GROUPS.RELATION_GROUP", instanceWithInferedType, &instanceWithInferedType.SPEC_RELATION_GROUPS.RELATION_GROUP, formGroup, probe)

	case *models.REQ_IF_HEADER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("COMMENT", instanceWithInferedType.COMMENT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CREATION_TIME", instanceWithInferedType.CREATION_TIME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REPOSITORY_ID", instanceWithInferedType.REPOSITORY_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REQ_IF_TOOL_ID", instanceWithInferedType.REQ_IF_TOOL_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REQ_IF_VERSION", instanceWithInferedType.REQ_IF_VERSION, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SOURCE_TOOL_ID", instanceWithInferedType.SOURCE_TOOL_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TITLE", instanceWithInferedType.TITLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.REQ_IF_TOOL_EXTENSION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF"
			rf.Fieldname = "TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF),
					"TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF, *models.REQ_IF_TOOL_EXTENSION](
					nil,
					"TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECIFICATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_DATE", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_DATE, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_INTEGER", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_REAL", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_REAL, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_STRING", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_STRING, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_XHTML", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_XHTML, formGroup, probe)
		AssociationSliceToForm("CHILDREN.SPEC_HIERARCHY", instanceWithInferedType, &instanceWithInferedType.CHILDREN.SPEC_HIERARCHY, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPECIFICATIONS.SPECIFICATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPECIFICATIONS.SPECIFICATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.SPECIFICATION](
					nil,
					"SPECIFICATIONS.SPECIFICATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECIFICATION_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_TYPES.SPECIFICATION_TYPE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_TYPES.SPECIFICATION_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.SPECIFICATION_TYPE](
					nil,
					"SPEC_TYPES.SPECIFICATION_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPEC_HIERARCHY:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("CHILDREN.SPEC_HIERARCHY", instanceWithInferedType, &instanceWithInferedType.CHILDREN.SPEC_HIERARCHY, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_TABLE_INTERNAL", instanceWithInferedType.IS_TABLE_INTERNAL, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "CHILDREN.SPEC_HIERARCHY"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"CHILDREN.SPEC_HIERARCHY",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.SPEC_HIERARCHY](
					nil,
					"CHILDREN.SPEC_HIERARCHY",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "CHILDREN.SPEC_HIERARCHY"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_HIERARCHY),
					"CHILDREN.SPEC_HIERARCHY",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_HIERARCHY, *models.SPEC_HIERARCHY](
					nil,
					"CHILDREN.SPEC_HIERARCHY",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPEC_OBJECT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_DATE", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_DATE, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_INTEGER", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_REAL", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_REAL, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_STRING", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_STRING, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_XHTML", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_XHTML, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_OBJECTS.SPEC_OBJECT"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_OBJECTS.SPEC_OBJECT",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.SPEC_OBJECT](
					nil,
					"SPEC_OBJECTS.SPEC_OBJECT",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPEC_OBJECT_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_TYPES.SPEC_OBJECT_TYPE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_TYPES.SPEC_OBJECT_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.SPEC_OBJECT_TYPE](
					nil,
					"SPEC_TYPES.SPEC_OBJECT_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPEC_RELATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_DATE", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_DATE, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_INTEGER", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_REAL", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_REAL, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_STRING", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_STRING, formGroup, probe)
		AssociationSliceToForm("VALUES.ATTRIBUTE_VALUE_XHTML", instanceWithInferedType, &instanceWithInferedType.VALUES.ATTRIBUTE_VALUE_XHTML, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_RELATIONS.SPEC_RELATION"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_RELATIONS.SPEC_RELATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.SPEC_RELATION](
					nil,
					"SPEC_RELATIONS.SPEC_RELATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPEC_RELATION_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID.ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_TYPES.SPEC_RELATION_TYPE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_TYPES.SPEC_RELATION_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.SPEC_RELATION_TYPE](
					nil,
					"SPEC_TYPES.SPEC_RELATION_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.XHTML_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "THE_VALUE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_XHTML),
					"THE_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_XHTML, *models.XHTML_CONTENT](
					nil,
					"THE_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "THE_ORIGINAL_VALUE"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_XHTML),
					"THE_ORIGINAL_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_XHTML, *models.XHTML_CONTENT](
					nil,
					"THE_ORIGINAL_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Xhtml_InlPres_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_a_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_abbr_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_acronym_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_address_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_blockquote_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_br_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_caption_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_cite_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_code_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_col_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_colgroup_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_dd_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_dfn_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_div_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_dl_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_dt_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_edit_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_em_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_h1_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_h2_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_h3_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_h4_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_h5_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_h6_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_heading_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_hr_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_kbd_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_li_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_object_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_ol_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_p_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_param_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_pre_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_q_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_samp_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_span_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_strong_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_table_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_tbody_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_td_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_tfoot_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_th_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_thead_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_tr_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_ul_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Xhtml_var_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	default:
		_ = instanceWithInferedType
	}
}
