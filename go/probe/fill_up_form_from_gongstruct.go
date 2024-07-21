// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ALTERNATIVE_ID Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ALTERNATIVE_IDFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_DATE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_REAL:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_STRING:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_XHTML:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AnyType:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "AnyType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnyTypeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_DATE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_INTEGER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_REAL:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_REALFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_STRING:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_XHTML:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EMBEDDED_VALUE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "EMBEDDED_VALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EMBEDDED_VALUEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ENUM_VALUE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ENUM_VALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ENUM_VALUEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RELATION_GROUP:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RELATION_GROUP Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATION_GROUPFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RELATION_GROUP_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RELATION_GROUP_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATION_GROUP_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQ_IF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQ_IF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQ_IF_CONTENT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQ_IF_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_CONTENTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQ_IF_HEADER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQ_IF_HEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_HEADERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQ_IF_TOOL_EXTENSION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQ_IF_TOOL_EXTENSION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECIFICATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECIFICATION_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECIFICATION_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATION_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_HIERARCHY:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_HIERARCHY Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_HIERARCHYFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_OBJECT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_OBJECT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_OBJECTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_OBJECT_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_OBJECT_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_OBJECT_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_RELATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_RELATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_RELATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_RELATION_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_RELATION_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_RELATION_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XHTML_CONTENT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "XHTML_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XHTML_CONTENTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_InlPres_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_InlPres_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_InlPres_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_a_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_a_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_a_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_abbr_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_abbr_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_abbr_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_acronym_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_acronym_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_acronym_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_address_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_address_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_address_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_blockquote_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_blockquote_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_blockquote_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_br_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_br_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_br_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_caption_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_caption_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_caption_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_cite_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_cite_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_cite_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_code_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_code_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_code_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_col_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_col_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_col_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_colgroup_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_colgroup_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_colgroup_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_dd_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_dd_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_dd_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_dfn_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_dfn_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_dfn_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_div_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_div_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_div_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_dl_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_dl_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_dl_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_dt_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_dt_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_dt_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_edit_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_edit_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_edit_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_em_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_em_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_em_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_h1_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_h1_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h1_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_h2_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_h2_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h2_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_h3_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_h3_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h3_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_h4_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_h4_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h4_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_h5_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_h5_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h5_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_h6_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_h6_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h6_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_heading_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_heading_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_heading_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_hr_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_hr_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_hr_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_kbd_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_kbd_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_kbd_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_li_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_li_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_li_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_object_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_object_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_object_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_ol_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_ol_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_ol_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_p_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_p_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_p_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_param_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_param_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_param_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_pre_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_pre_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_pre_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_q_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_q_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_q_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_samp_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_samp_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_samp_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_span_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_span_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_span_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_strong_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_strong_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_strong_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_table_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_table_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_table_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_tbody_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_tbody_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_tbody_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_td_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_td_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_td_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_tfoot_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_tfoot_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_tfoot_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_th_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_th_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_th_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_thead_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_thead_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_thead_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_tr_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_tr_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_tr_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_ul_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_ul_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_ul_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Xhtml_var_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Xhtml_var_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_var_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
