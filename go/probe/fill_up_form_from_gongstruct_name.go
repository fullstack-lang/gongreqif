// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "ALTERNATIVE_ID":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ALTERNATIVE_ID Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ALTERNATIVE_IDFormCallback(
			nil,
			probe,
			formGroup,
		)
		alternative_id := new(models.ALTERNATIVE_ID)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(alternative_id, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_BOOLEAN":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_DEFINITION_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_boolean := new(models.ATTRIBUTE_DEFINITION_BOOLEAN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_boolean, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_DATE":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_DEFINITION_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_date := new(models.ATTRIBUTE_DEFINITION_DATE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_date, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_ENUMERATION":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_DEFINITION_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_enumeration := new(models.ATTRIBUTE_DEFINITION_ENUMERATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_enumeration, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_INTEGER":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_DEFINITION_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_integer := new(models.ATTRIBUTE_DEFINITION_INTEGER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_integer, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_REAL":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_DEFINITION_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_real := new(models.ATTRIBUTE_DEFINITION_REAL)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_real, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_STRING":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_DEFINITION_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_string := new(models.ATTRIBUTE_DEFINITION_STRING)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_string, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_XHTML":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_DEFINITION_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_xhtml := new(models.ATTRIBUTE_DEFINITION_XHTML)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_xhtml, formGroup, probe)
	case "ATTRIBUTE_VALUE_BOOLEAN":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_VALUE_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_boolean := new(models.ATTRIBUTE_VALUE_BOOLEAN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_boolean, formGroup, probe)
	case "ATTRIBUTE_VALUE_DATE":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_VALUE_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_date := new(models.ATTRIBUTE_VALUE_DATE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_date, formGroup, probe)
	case "ATTRIBUTE_VALUE_ENUMERATION":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_VALUE_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_enumeration := new(models.ATTRIBUTE_VALUE_ENUMERATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_enumeration, formGroup, probe)
	case "ATTRIBUTE_VALUE_INTEGER":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_VALUE_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_integer := new(models.ATTRIBUTE_VALUE_INTEGER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_integer, formGroup, probe)
	case "ATTRIBUTE_VALUE_REAL":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_VALUE_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_real := new(models.ATTRIBUTE_VALUE_REAL)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_real, formGroup, probe)
	case "ATTRIBUTE_VALUE_STRING":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_VALUE_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_string := new(models.ATTRIBUTE_VALUE_STRING)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_string, formGroup, probe)
	case "ATTRIBUTE_VALUE_XHTML":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ATTRIBUTE_VALUE_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_xhtml := new(models.ATTRIBUTE_VALUE_XHTML)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_xhtml, formGroup, probe)
	case "AnyType":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AnyType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnyTypeFormCallback(
			nil,
			probe,
			formGroup,
		)
		anytype := new(models.AnyType)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(anytype, formGroup, probe)
	case "DATATYPE_DEFINITION_BOOLEAN":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DATATYPE_DEFINITION_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_boolean := new(models.DATATYPE_DEFINITION_BOOLEAN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_boolean, formGroup, probe)
	case "DATATYPE_DEFINITION_DATE":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DATATYPE_DEFINITION_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_date := new(models.DATATYPE_DEFINITION_DATE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_date, formGroup, probe)
	case "DATATYPE_DEFINITION_ENUMERATION":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DATATYPE_DEFINITION_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_enumeration := new(models.DATATYPE_DEFINITION_ENUMERATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_enumeration, formGroup, probe)
	case "DATATYPE_DEFINITION_INTEGER":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DATATYPE_DEFINITION_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_integer := new(models.DATATYPE_DEFINITION_INTEGER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_integer, formGroup, probe)
	case "DATATYPE_DEFINITION_REAL":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DATATYPE_DEFINITION_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_REALFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_real := new(models.DATATYPE_DEFINITION_REAL)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_real, formGroup, probe)
	case "DATATYPE_DEFINITION_STRING":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DATATYPE_DEFINITION_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_string := new(models.DATATYPE_DEFINITION_STRING)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_string, formGroup, probe)
	case "DATATYPE_DEFINITION_XHTML":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DATATYPE_DEFINITION_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_xhtml := new(models.DATATYPE_DEFINITION_XHTML)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_xhtml, formGroup, probe)
	case "EMBEDDED_VALUE":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EMBEDDED_VALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EMBEDDED_VALUEFormCallback(
			nil,
			probe,
			formGroup,
		)
		embedded_value := new(models.EMBEDDED_VALUE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(embedded_value, formGroup, probe)
	case "ENUM_VALUE":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ENUM_VALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ENUM_VALUEFormCallback(
			nil,
			probe,
			formGroup,
		)
		enum_value := new(models.ENUM_VALUE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(enum_value, formGroup, probe)
	case "RELATION_GROUP":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RELATION_GROUP Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATION_GROUPFormCallback(
			nil,
			probe,
			formGroup,
		)
		relation_group := new(models.RELATION_GROUP)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(relation_group, formGroup, probe)
	case "RELATION_GROUP_TYPE":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RELATION_GROUP_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATION_GROUP_TYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		relation_group_type := new(models.RELATION_GROUP_TYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(relation_group_type, formGroup, probe)
	case "REQ_IF":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "REQ_IF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IFFormCallback(
			nil,
			probe,
			formGroup,
		)
		req_if := new(models.REQ_IF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(req_if, formGroup, probe)
	case "REQ_IF_CONTENT":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "REQ_IF_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_CONTENTFormCallback(
			nil,
			probe,
			formGroup,
		)
		req_if_content := new(models.REQ_IF_CONTENT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(req_if_content, formGroup, probe)
	case "REQ_IF_HEADER":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "REQ_IF_HEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_HEADERFormCallback(
			nil,
			probe,
			formGroup,
		)
		req_if_header := new(models.REQ_IF_HEADER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(req_if_header, formGroup, probe)
	case "REQ_IF_TOOL_EXTENSION":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "REQ_IF_TOOL_EXTENSION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		req_if_tool_extension := new(models.REQ_IF_TOOL_EXTENSION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(req_if_tool_extension, formGroup, probe)
	case "SPECIFICATION":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		specification := new(models.SPECIFICATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specification, formGroup, probe)
	case "SPECIFICATION_TYPE":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SPECIFICATION_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATION_TYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		specification_type := new(models.SPECIFICATION_TYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specification_type, formGroup, probe)
	case "SPEC_HIERARCHY":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SPEC_HIERARCHY Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_HIERARCHYFormCallback(
			nil,
			probe,
			formGroup,
		)
		spec_hierarchy := new(models.SPEC_HIERARCHY)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spec_hierarchy, formGroup, probe)
	case "SPEC_OBJECT":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SPEC_OBJECT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_OBJECTFormCallback(
			nil,
			probe,
			formGroup,
		)
		spec_object := new(models.SPEC_OBJECT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spec_object, formGroup, probe)
	case "SPEC_OBJECT_TYPE":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SPEC_OBJECT_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_OBJECT_TYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		spec_object_type := new(models.SPEC_OBJECT_TYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spec_object_type, formGroup, probe)
	case "SPEC_RELATION":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SPEC_RELATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_RELATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		spec_relation := new(models.SPEC_RELATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spec_relation, formGroup, probe)
	case "SPEC_RELATION_TYPE":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SPEC_RELATION_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_RELATION_TYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		spec_relation_type := new(models.SPEC_RELATION_TYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spec_relation_type, formGroup, probe)
	case "XHTML_CONTENT":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "XHTML_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XHTML_CONTENTFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_content := new(models.XHTML_CONTENT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_content, formGroup, probe)
	case "Xhtml_InlPres_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_InlPres_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_InlPres_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_inlpres_type := new(models.Xhtml_InlPres_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_inlpres_type, formGroup, probe)
	case "Xhtml_a_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_a_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_a_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_a_type := new(models.Xhtml_a_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_a_type, formGroup, probe)
	case "Xhtml_abbr_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_abbr_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_abbr_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_abbr_type := new(models.Xhtml_abbr_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_abbr_type, formGroup, probe)
	case "Xhtml_acronym_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_acronym_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_acronym_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_acronym_type := new(models.Xhtml_acronym_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_acronym_type, formGroup, probe)
	case "Xhtml_address_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_address_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_address_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_address_type := new(models.Xhtml_address_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_address_type, formGroup, probe)
	case "Xhtml_blockquote_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_blockquote_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_blockquote_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_blockquote_type := new(models.Xhtml_blockquote_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_blockquote_type, formGroup, probe)
	case "Xhtml_br_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_br_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_br_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_br_type := new(models.Xhtml_br_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_br_type, formGroup, probe)
	case "Xhtml_caption_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_caption_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_caption_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_caption_type := new(models.Xhtml_caption_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_caption_type, formGroup, probe)
	case "Xhtml_cite_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_cite_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_cite_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_cite_type := new(models.Xhtml_cite_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_cite_type, formGroup, probe)
	case "Xhtml_code_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_code_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_code_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_code_type := new(models.Xhtml_code_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_code_type, formGroup, probe)
	case "Xhtml_col_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_col_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_col_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_col_type := new(models.Xhtml_col_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_col_type, formGroup, probe)
	case "Xhtml_colgroup_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_colgroup_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_colgroup_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_colgroup_type := new(models.Xhtml_colgroup_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_colgroup_type, formGroup, probe)
	case "Xhtml_dd_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_dd_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_dd_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_dd_type := new(models.Xhtml_dd_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_dd_type, formGroup, probe)
	case "Xhtml_dfn_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_dfn_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_dfn_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_dfn_type := new(models.Xhtml_dfn_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_dfn_type, formGroup, probe)
	case "Xhtml_div_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_div_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_div_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_div_type := new(models.Xhtml_div_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_div_type, formGroup, probe)
	case "Xhtml_dl_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_dl_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_dl_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_dl_type := new(models.Xhtml_dl_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_dl_type, formGroup, probe)
	case "Xhtml_dt_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_dt_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_dt_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_dt_type := new(models.Xhtml_dt_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_dt_type, formGroup, probe)
	case "Xhtml_edit_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_edit_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_edit_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_edit_type := new(models.Xhtml_edit_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_edit_type, formGroup, probe)
	case "Xhtml_em_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_em_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_em_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_em_type := new(models.Xhtml_em_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_em_type, formGroup, probe)
	case "Xhtml_h1_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_h1_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h1_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_h1_type := new(models.Xhtml_h1_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_h1_type, formGroup, probe)
	case "Xhtml_h2_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_h2_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h2_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_h2_type := new(models.Xhtml_h2_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_h2_type, formGroup, probe)
	case "Xhtml_h3_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_h3_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h3_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_h3_type := new(models.Xhtml_h3_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_h3_type, formGroup, probe)
	case "Xhtml_h4_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_h4_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h4_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_h4_type := new(models.Xhtml_h4_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_h4_type, formGroup, probe)
	case "Xhtml_h5_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_h5_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h5_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_h5_type := new(models.Xhtml_h5_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_h5_type, formGroup, probe)
	case "Xhtml_h6_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_h6_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_h6_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_h6_type := new(models.Xhtml_h6_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_h6_type, formGroup, probe)
	case "Xhtml_heading_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_heading_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_heading_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_heading_type := new(models.Xhtml_heading_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_heading_type, formGroup, probe)
	case "Xhtml_hr_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_hr_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_hr_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_hr_type := new(models.Xhtml_hr_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_hr_type, formGroup, probe)
	case "Xhtml_kbd_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_kbd_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_kbd_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_kbd_type := new(models.Xhtml_kbd_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_kbd_type, formGroup, probe)
	case "Xhtml_li_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_li_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_li_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_li_type := new(models.Xhtml_li_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_li_type, formGroup, probe)
	case "Xhtml_object_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_object_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_object_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_object_type := new(models.Xhtml_object_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_object_type, formGroup, probe)
	case "Xhtml_ol_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_ol_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_ol_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_ol_type := new(models.Xhtml_ol_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_ol_type, formGroup, probe)
	case "Xhtml_p_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_p_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_p_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_p_type := new(models.Xhtml_p_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_p_type, formGroup, probe)
	case "Xhtml_param_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_param_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_param_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_param_type := new(models.Xhtml_param_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_param_type, formGroup, probe)
	case "Xhtml_pre_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_pre_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_pre_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_pre_type := new(models.Xhtml_pre_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_pre_type, formGroup, probe)
	case "Xhtml_q_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_q_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_q_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_q_type := new(models.Xhtml_q_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_q_type, formGroup, probe)
	case "Xhtml_samp_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_samp_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_samp_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_samp_type := new(models.Xhtml_samp_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_samp_type, formGroup, probe)
	case "Xhtml_span_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_span_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_span_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_span_type := new(models.Xhtml_span_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_span_type, formGroup, probe)
	case "Xhtml_strong_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_strong_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_strong_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_strong_type := new(models.Xhtml_strong_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_strong_type, formGroup, probe)
	case "Xhtml_table_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_table_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_table_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_table_type := new(models.Xhtml_table_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_table_type, formGroup, probe)
	case "Xhtml_tbody_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_tbody_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_tbody_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_tbody_type := new(models.Xhtml_tbody_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_tbody_type, formGroup, probe)
	case "Xhtml_td_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_td_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_td_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_td_type := new(models.Xhtml_td_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_td_type, formGroup, probe)
	case "Xhtml_tfoot_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_tfoot_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_tfoot_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_tfoot_type := new(models.Xhtml_tfoot_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_tfoot_type, formGroup, probe)
	case "Xhtml_th_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_th_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_th_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_th_type := new(models.Xhtml_th_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_th_type, formGroup, probe)
	case "Xhtml_thead_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_thead_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_thead_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_thead_type := new(models.Xhtml_thead_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_thead_type, formGroup, probe)
	case "Xhtml_tr_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_tr_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_tr_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_tr_type := new(models.Xhtml_tr_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_tr_type, formGroup, probe)
	case "Xhtml_ul_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_ul_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_ul_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_ul_type := new(models.Xhtml_ul_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_ul_type, formGroup, probe)
	case "Xhtml_var_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xhtml_var_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Xhtml_var_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_var_type := new(models.Xhtml_var_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_var_type, formGroup, probe)
	}
	formStage.Commit()
}
