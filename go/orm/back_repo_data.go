// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ALTERNATIVE_IDAPIs []*ALTERNATIVE_IDAPI

	ATTRIBUTE_DEFINITION_BOOLEANAPIs []*ATTRIBUTE_DEFINITION_BOOLEANAPI

	ATTRIBUTE_DEFINITION_DATEAPIs []*ATTRIBUTE_DEFINITION_DATEAPI

	ATTRIBUTE_DEFINITION_ENUMERATIONAPIs []*ATTRIBUTE_DEFINITION_ENUMERATIONAPI

	ATTRIBUTE_DEFINITION_INTEGERAPIs []*ATTRIBUTE_DEFINITION_INTEGERAPI

	ATTRIBUTE_DEFINITION_REALAPIs []*ATTRIBUTE_DEFINITION_REALAPI

	ATTRIBUTE_DEFINITION_STRINGAPIs []*ATTRIBUTE_DEFINITION_STRINGAPI

	ATTRIBUTE_DEFINITION_XHTMLAPIs []*ATTRIBUTE_DEFINITION_XHTMLAPI

	ATTRIBUTE_VALUE_BOOLEANAPIs []*ATTRIBUTE_VALUE_BOOLEANAPI

	ATTRIBUTE_VALUE_DATEAPIs []*ATTRIBUTE_VALUE_DATEAPI

	ATTRIBUTE_VALUE_ENUMERATIONAPIs []*ATTRIBUTE_VALUE_ENUMERATIONAPI

	ATTRIBUTE_VALUE_INTEGERAPIs []*ATTRIBUTE_VALUE_INTEGERAPI

	ATTRIBUTE_VALUE_REALAPIs []*ATTRIBUTE_VALUE_REALAPI

	ATTRIBUTE_VALUE_STRINGAPIs []*ATTRIBUTE_VALUE_STRINGAPI

	ATTRIBUTE_VALUE_XHTMLAPIs []*ATTRIBUTE_VALUE_XHTMLAPI

	AnyTypeAPIs []*AnyTypeAPI

	DATATYPE_DEFINITION_BOOLEANAPIs []*DATATYPE_DEFINITION_BOOLEANAPI

	DATATYPE_DEFINITION_DATEAPIs []*DATATYPE_DEFINITION_DATEAPI

	DATATYPE_DEFINITION_ENUMERATIONAPIs []*DATATYPE_DEFINITION_ENUMERATIONAPI

	DATATYPE_DEFINITION_INTEGERAPIs []*DATATYPE_DEFINITION_INTEGERAPI

	DATATYPE_DEFINITION_REALAPIs []*DATATYPE_DEFINITION_REALAPI

	DATATYPE_DEFINITION_STRINGAPIs []*DATATYPE_DEFINITION_STRINGAPI

	DATATYPE_DEFINITION_XHTMLAPIs []*DATATYPE_DEFINITION_XHTMLAPI

	EMBEDDED_VALUEAPIs []*EMBEDDED_VALUEAPI

	ENUM_VALUEAPIs []*ENUM_VALUEAPI

	RELATION_GROUPAPIs []*RELATION_GROUPAPI

	RELATION_GROUP_TYPEAPIs []*RELATION_GROUP_TYPEAPI

	REQ_IFAPIs []*REQ_IFAPI

	REQ_IF_CONTENTAPIs []*REQ_IF_CONTENTAPI

	REQ_IF_HEADERAPIs []*REQ_IF_HEADERAPI

	REQ_IF_TOOL_EXTENSIONAPIs []*REQ_IF_TOOL_EXTENSIONAPI

	SPECIFICATIONAPIs []*SPECIFICATIONAPI

	SPECIFICATION_TYPEAPIs []*SPECIFICATION_TYPEAPI

	SPEC_HIERARCHYAPIs []*SPEC_HIERARCHYAPI

	SPEC_OBJECTAPIs []*SPEC_OBJECTAPI

	SPEC_OBJECT_TYPEAPIs []*SPEC_OBJECT_TYPEAPI

	SPEC_RELATIONAPIs []*SPEC_RELATIONAPI

	SPEC_RELATION_TYPEAPIs []*SPEC_RELATION_TYPEAPI

	XHTML_CONTENTAPIs []*XHTML_CONTENTAPI

	Xhtml_InlPres_typeAPIs []*Xhtml_InlPres_typeAPI

	Xhtml_a_typeAPIs []*Xhtml_a_typeAPI

	Xhtml_abbr_typeAPIs []*Xhtml_abbr_typeAPI

	Xhtml_acronym_typeAPIs []*Xhtml_acronym_typeAPI

	Xhtml_address_typeAPIs []*Xhtml_address_typeAPI

	Xhtml_blockquote_typeAPIs []*Xhtml_blockquote_typeAPI

	Xhtml_br_typeAPIs []*Xhtml_br_typeAPI

	Xhtml_caption_typeAPIs []*Xhtml_caption_typeAPI

	Xhtml_cite_typeAPIs []*Xhtml_cite_typeAPI

	Xhtml_code_typeAPIs []*Xhtml_code_typeAPI

	Xhtml_col_typeAPIs []*Xhtml_col_typeAPI

	Xhtml_colgroup_typeAPIs []*Xhtml_colgroup_typeAPI

	Xhtml_dd_typeAPIs []*Xhtml_dd_typeAPI

	Xhtml_dfn_typeAPIs []*Xhtml_dfn_typeAPI

	Xhtml_div_typeAPIs []*Xhtml_div_typeAPI

	Xhtml_dl_typeAPIs []*Xhtml_dl_typeAPI

	Xhtml_dt_typeAPIs []*Xhtml_dt_typeAPI

	Xhtml_edit_typeAPIs []*Xhtml_edit_typeAPI

	Xhtml_em_typeAPIs []*Xhtml_em_typeAPI

	Xhtml_h1_typeAPIs []*Xhtml_h1_typeAPI

	Xhtml_h2_typeAPIs []*Xhtml_h2_typeAPI

	Xhtml_h3_typeAPIs []*Xhtml_h3_typeAPI

	Xhtml_h4_typeAPIs []*Xhtml_h4_typeAPI

	Xhtml_h5_typeAPIs []*Xhtml_h5_typeAPI

	Xhtml_h6_typeAPIs []*Xhtml_h6_typeAPI

	Xhtml_heading_typeAPIs []*Xhtml_heading_typeAPI

	Xhtml_hr_typeAPIs []*Xhtml_hr_typeAPI

	Xhtml_kbd_typeAPIs []*Xhtml_kbd_typeAPI

	Xhtml_li_typeAPIs []*Xhtml_li_typeAPI

	Xhtml_object_typeAPIs []*Xhtml_object_typeAPI

	Xhtml_ol_typeAPIs []*Xhtml_ol_typeAPI

	Xhtml_p_typeAPIs []*Xhtml_p_typeAPI

	Xhtml_param_typeAPIs []*Xhtml_param_typeAPI

	Xhtml_pre_typeAPIs []*Xhtml_pre_typeAPI

	Xhtml_q_typeAPIs []*Xhtml_q_typeAPI

	Xhtml_samp_typeAPIs []*Xhtml_samp_typeAPI

	Xhtml_span_typeAPIs []*Xhtml_span_typeAPI

	Xhtml_strong_typeAPIs []*Xhtml_strong_typeAPI

	Xhtml_table_typeAPIs []*Xhtml_table_typeAPI

	Xhtml_tbody_typeAPIs []*Xhtml_tbody_typeAPI

	Xhtml_td_typeAPIs []*Xhtml_td_typeAPI

	Xhtml_tfoot_typeAPIs []*Xhtml_tfoot_typeAPI

	Xhtml_th_typeAPIs []*Xhtml_th_typeAPI

	Xhtml_thead_typeAPIs []*Xhtml_thead_typeAPI

	Xhtml_tr_typeAPIs []*Xhtml_tr_typeAPI

	Xhtml_ul_typeAPIs []*Xhtml_ul_typeAPI

	Xhtml_var_typeAPIs []*Xhtml_var_typeAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, alternative_idDB := range backRepo.BackRepoALTERNATIVE_ID.Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDDB {

		var alternative_idAPI ALTERNATIVE_IDAPI
		alternative_idAPI.ID = alternative_idDB.ID
		alternative_idAPI.ALTERNATIVE_IDPointersEncoding = alternative_idDB.ALTERNATIVE_IDPointersEncoding
		alternative_idDB.CopyBasicFieldsToALTERNATIVE_ID_WOP(&alternative_idAPI.ALTERNATIVE_ID_WOP)

		backRepoData.ALTERNATIVE_IDAPIs = append(backRepoData.ALTERNATIVE_IDAPIs, &alternative_idAPI)
	}

	for _, attribute_definition_booleanDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB {

		var attribute_definition_booleanAPI ATTRIBUTE_DEFINITION_BOOLEANAPI
		attribute_definition_booleanAPI.ID = attribute_definition_booleanDB.ID
		attribute_definition_booleanAPI.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding = attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding
		attribute_definition_booleanDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN_WOP(&attribute_definition_booleanAPI.ATTRIBUTE_DEFINITION_BOOLEAN_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_BOOLEANAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_BOOLEANAPIs, &attribute_definition_booleanAPI)
	}

	for _, attribute_definition_dateDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.Map_ATTRIBUTE_DEFINITION_DATEDBID_ATTRIBUTE_DEFINITION_DATEDB {

		var attribute_definition_dateAPI ATTRIBUTE_DEFINITION_DATEAPI
		attribute_definition_dateAPI.ID = attribute_definition_dateDB.ID
		attribute_definition_dateAPI.ATTRIBUTE_DEFINITION_DATEPointersEncoding = attribute_definition_dateDB.ATTRIBUTE_DEFINITION_DATEPointersEncoding
		attribute_definition_dateDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_DATE_WOP(&attribute_definition_dateAPI.ATTRIBUTE_DEFINITION_DATE_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_DATEAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_DATEAPIs, &attribute_definition_dateAPI)
	}

	for _, attribute_definition_enumerationDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.Map_ATTRIBUTE_DEFINITION_ENUMERATIONDBID_ATTRIBUTE_DEFINITION_ENUMERATIONDB {

		var attribute_definition_enumerationAPI ATTRIBUTE_DEFINITION_ENUMERATIONAPI
		attribute_definition_enumerationAPI.ID = attribute_definition_enumerationDB.ID
		attribute_definition_enumerationAPI.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding = attribute_definition_enumerationDB.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding
		attribute_definition_enumerationDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_ENUMERATION_WOP(&attribute_definition_enumerationAPI.ATTRIBUTE_DEFINITION_ENUMERATION_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_ENUMERATIONAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_ENUMERATIONAPIs, &attribute_definition_enumerationAPI)
	}

	for _, attribute_definition_integerDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB {

		var attribute_definition_integerAPI ATTRIBUTE_DEFINITION_INTEGERAPI
		attribute_definition_integerAPI.ID = attribute_definition_integerDB.ID
		attribute_definition_integerAPI.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding = attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding
		attribute_definition_integerDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER_WOP(&attribute_definition_integerAPI.ATTRIBUTE_DEFINITION_INTEGER_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_INTEGERAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_INTEGERAPIs, &attribute_definition_integerAPI)
	}

	for _, attribute_definition_realDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.Map_ATTRIBUTE_DEFINITION_REALDBID_ATTRIBUTE_DEFINITION_REALDB {

		var attribute_definition_realAPI ATTRIBUTE_DEFINITION_REALAPI
		attribute_definition_realAPI.ID = attribute_definition_realDB.ID
		attribute_definition_realAPI.ATTRIBUTE_DEFINITION_REALPointersEncoding = attribute_definition_realDB.ATTRIBUTE_DEFINITION_REALPointersEncoding
		attribute_definition_realDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_REAL_WOP(&attribute_definition_realAPI.ATTRIBUTE_DEFINITION_REAL_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_REALAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_REALAPIs, &attribute_definition_realAPI)
	}

	for _, attribute_definition_stringDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB {

		var attribute_definition_stringAPI ATTRIBUTE_DEFINITION_STRINGAPI
		attribute_definition_stringAPI.ID = attribute_definition_stringDB.ID
		attribute_definition_stringAPI.ATTRIBUTE_DEFINITION_STRINGPointersEncoding = attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding
		attribute_definition_stringDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING_WOP(&attribute_definition_stringAPI.ATTRIBUTE_DEFINITION_STRING_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_STRINGAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_STRINGAPIs, &attribute_definition_stringAPI)
	}

	for _, attribute_definition_xhtmlDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.Map_ATTRIBUTE_DEFINITION_XHTMLDBID_ATTRIBUTE_DEFINITION_XHTMLDB {

		var attribute_definition_xhtmlAPI ATTRIBUTE_DEFINITION_XHTMLAPI
		attribute_definition_xhtmlAPI.ID = attribute_definition_xhtmlDB.ID
		attribute_definition_xhtmlAPI.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding = attribute_definition_xhtmlDB.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding
		attribute_definition_xhtmlDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_XHTML_WOP(&attribute_definition_xhtmlAPI.ATTRIBUTE_DEFINITION_XHTML_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_XHTMLAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_XHTMLAPIs, &attribute_definition_xhtmlAPI)
	}

	for _, attribute_value_booleanDB := range backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.Map_ATTRIBUTE_VALUE_BOOLEANDBID_ATTRIBUTE_VALUE_BOOLEANDB {

		var attribute_value_booleanAPI ATTRIBUTE_VALUE_BOOLEANAPI
		attribute_value_booleanAPI.ID = attribute_value_booleanDB.ID
		attribute_value_booleanAPI.ATTRIBUTE_VALUE_BOOLEANPointersEncoding = attribute_value_booleanDB.ATTRIBUTE_VALUE_BOOLEANPointersEncoding
		attribute_value_booleanDB.CopyBasicFieldsToATTRIBUTE_VALUE_BOOLEAN_WOP(&attribute_value_booleanAPI.ATTRIBUTE_VALUE_BOOLEAN_WOP)

		backRepoData.ATTRIBUTE_VALUE_BOOLEANAPIs = append(backRepoData.ATTRIBUTE_VALUE_BOOLEANAPIs, &attribute_value_booleanAPI)
	}

	for _, attribute_value_dateDB := range backRepo.BackRepoATTRIBUTE_VALUE_DATE.Map_ATTRIBUTE_VALUE_DATEDBID_ATTRIBUTE_VALUE_DATEDB {

		var attribute_value_dateAPI ATTRIBUTE_VALUE_DATEAPI
		attribute_value_dateAPI.ID = attribute_value_dateDB.ID
		attribute_value_dateAPI.ATTRIBUTE_VALUE_DATEPointersEncoding = attribute_value_dateDB.ATTRIBUTE_VALUE_DATEPointersEncoding
		attribute_value_dateDB.CopyBasicFieldsToATTRIBUTE_VALUE_DATE_WOP(&attribute_value_dateAPI.ATTRIBUTE_VALUE_DATE_WOP)

		backRepoData.ATTRIBUTE_VALUE_DATEAPIs = append(backRepoData.ATTRIBUTE_VALUE_DATEAPIs, &attribute_value_dateAPI)
	}

	for _, attribute_value_enumerationDB := range backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.Map_ATTRIBUTE_VALUE_ENUMERATIONDBID_ATTRIBUTE_VALUE_ENUMERATIONDB {

		var attribute_value_enumerationAPI ATTRIBUTE_VALUE_ENUMERATIONAPI
		attribute_value_enumerationAPI.ID = attribute_value_enumerationDB.ID
		attribute_value_enumerationAPI.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding = attribute_value_enumerationDB.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding
		attribute_value_enumerationDB.CopyBasicFieldsToATTRIBUTE_VALUE_ENUMERATION_WOP(&attribute_value_enumerationAPI.ATTRIBUTE_VALUE_ENUMERATION_WOP)

		backRepoData.ATTRIBUTE_VALUE_ENUMERATIONAPIs = append(backRepoData.ATTRIBUTE_VALUE_ENUMERATIONAPIs, &attribute_value_enumerationAPI)
	}

	for _, attribute_value_integerDB := range backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.Map_ATTRIBUTE_VALUE_INTEGERDBID_ATTRIBUTE_VALUE_INTEGERDB {

		var attribute_value_integerAPI ATTRIBUTE_VALUE_INTEGERAPI
		attribute_value_integerAPI.ID = attribute_value_integerDB.ID
		attribute_value_integerAPI.ATTRIBUTE_VALUE_INTEGERPointersEncoding = attribute_value_integerDB.ATTRIBUTE_VALUE_INTEGERPointersEncoding
		attribute_value_integerDB.CopyBasicFieldsToATTRIBUTE_VALUE_INTEGER_WOP(&attribute_value_integerAPI.ATTRIBUTE_VALUE_INTEGER_WOP)

		backRepoData.ATTRIBUTE_VALUE_INTEGERAPIs = append(backRepoData.ATTRIBUTE_VALUE_INTEGERAPIs, &attribute_value_integerAPI)
	}

	for _, attribute_value_realDB := range backRepo.BackRepoATTRIBUTE_VALUE_REAL.Map_ATTRIBUTE_VALUE_REALDBID_ATTRIBUTE_VALUE_REALDB {

		var attribute_value_realAPI ATTRIBUTE_VALUE_REALAPI
		attribute_value_realAPI.ID = attribute_value_realDB.ID
		attribute_value_realAPI.ATTRIBUTE_VALUE_REALPointersEncoding = attribute_value_realDB.ATTRIBUTE_VALUE_REALPointersEncoding
		attribute_value_realDB.CopyBasicFieldsToATTRIBUTE_VALUE_REAL_WOP(&attribute_value_realAPI.ATTRIBUTE_VALUE_REAL_WOP)

		backRepoData.ATTRIBUTE_VALUE_REALAPIs = append(backRepoData.ATTRIBUTE_VALUE_REALAPIs, &attribute_value_realAPI)
	}

	for _, attribute_value_stringDB := range backRepo.BackRepoATTRIBUTE_VALUE_STRING.Map_ATTRIBUTE_VALUE_STRINGDBID_ATTRIBUTE_VALUE_STRINGDB {

		var attribute_value_stringAPI ATTRIBUTE_VALUE_STRINGAPI
		attribute_value_stringAPI.ID = attribute_value_stringDB.ID
		attribute_value_stringAPI.ATTRIBUTE_VALUE_STRINGPointersEncoding = attribute_value_stringDB.ATTRIBUTE_VALUE_STRINGPointersEncoding
		attribute_value_stringDB.CopyBasicFieldsToATTRIBUTE_VALUE_STRING_WOP(&attribute_value_stringAPI.ATTRIBUTE_VALUE_STRING_WOP)

		backRepoData.ATTRIBUTE_VALUE_STRINGAPIs = append(backRepoData.ATTRIBUTE_VALUE_STRINGAPIs, &attribute_value_stringAPI)
	}

	for _, attribute_value_xhtmlDB := range backRepo.BackRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB {

		var attribute_value_xhtmlAPI ATTRIBUTE_VALUE_XHTMLAPI
		attribute_value_xhtmlAPI.ID = attribute_value_xhtmlDB.ID
		attribute_value_xhtmlAPI.ATTRIBUTE_VALUE_XHTMLPointersEncoding = attribute_value_xhtmlDB.ATTRIBUTE_VALUE_XHTMLPointersEncoding
		attribute_value_xhtmlDB.CopyBasicFieldsToATTRIBUTE_VALUE_XHTML_WOP(&attribute_value_xhtmlAPI.ATTRIBUTE_VALUE_XHTML_WOP)

		backRepoData.ATTRIBUTE_VALUE_XHTMLAPIs = append(backRepoData.ATTRIBUTE_VALUE_XHTMLAPIs, &attribute_value_xhtmlAPI)
	}

	for _, anytypeDB := range backRepo.BackRepoAnyType.Map_AnyTypeDBID_AnyTypeDB {

		var anytypeAPI AnyTypeAPI
		anytypeAPI.ID = anytypeDB.ID
		anytypeAPI.AnyTypePointersEncoding = anytypeDB.AnyTypePointersEncoding
		anytypeDB.CopyBasicFieldsToAnyType_WOP(&anytypeAPI.AnyType_WOP)

		backRepoData.AnyTypeAPIs = append(backRepoData.AnyTypeAPIs, &anytypeAPI)
	}

	for _, datatype_definition_booleanDB := range backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.Map_DATATYPE_DEFINITION_BOOLEANDBID_DATATYPE_DEFINITION_BOOLEANDB {

		var datatype_definition_booleanAPI DATATYPE_DEFINITION_BOOLEANAPI
		datatype_definition_booleanAPI.ID = datatype_definition_booleanDB.ID
		datatype_definition_booleanAPI.DATATYPE_DEFINITION_BOOLEANPointersEncoding = datatype_definition_booleanDB.DATATYPE_DEFINITION_BOOLEANPointersEncoding
		datatype_definition_booleanDB.CopyBasicFieldsToDATATYPE_DEFINITION_BOOLEAN_WOP(&datatype_definition_booleanAPI.DATATYPE_DEFINITION_BOOLEAN_WOP)

		backRepoData.DATATYPE_DEFINITION_BOOLEANAPIs = append(backRepoData.DATATYPE_DEFINITION_BOOLEANAPIs, &datatype_definition_booleanAPI)
	}

	for _, datatype_definition_dateDB := range backRepo.BackRepoDATATYPE_DEFINITION_DATE.Map_DATATYPE_DEFINITION_DATEDBID_DATATYPE_DEFINITION_DATEDB {

		var datatype_definition_dateAPI DATATYPE_DEFINITION_DATEAPI
		datatype_definition_dateAPI.ID = datatype_definition_dateDB.ID
		datatype_definition_dateAPI.DATATYPE_DEFINITION_DATEPointersEncoding = datatype_definition_dateDB.DATATYPE_DEFINITION_DATEPointersEncoding
		datatype_definition_dateDB.CopyBasicFieldsToDATATYPE_DEFINITION_DATE_WOP(&datatype_definition_dateAPI.DATATYPE_DEFINITION_DATE_WOP)

		backRepoData.DATATYPE_DEFINITION_DATEAPIs = append(backRepoData.DATATYPE_DEFINITION_DATEAPIs, &datatype_definition_dateAPI)
	}

	for _, datatype_definition_enumerationDB := range backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.Map_DATATYPE_DEFINITION_ENUMERATIONDBID_DATATYPE_DEFINITION_ENUMERATIONDB {

		var datatype_definition_enumerationAPI DATATYPE_DEFINITION_ENUMERATIONAPI
		datatype_definition_enumerationAPI.ID = datatype_definition_enumerationDB.ID
		datatype_definition_enumerationAPI.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding = datatype_definition_enumerationDB.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding
		datatype_definition_enumerationDB.CopyBasicFieldsToDATATYPE_DEFINITION_ENUMERATION_WOP(&datatype_definition_enumerationAPI.DATATYPE_DEFINITION_ENUMERATION_WOP)

		backRepoData.DATATYPE_DEFINITION_ENUMERATIONAPIs = append(backRepoData.DATATYPE_DEFINITION_ENUMERATIONAPIs, &datatype_definition_enumerationAPI)
	}

	for _, datatype_definition_integerDB := range backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB {

		var datatype_definition_integerAPI DATATYPE_DEFINITION_INTEGERAPI
		datatype_definition_integerAPI.ID = datatype_definition_integerDB.ID
		datatype_definition_integerAPI.DATATYPE_DEFINITION_INTEGERPointersEncoding = datatype_definition_integerDB.DATATYPE_DEFINITION_INTEGERPointersEncoding
		datatype_definition_integerDB.CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER_WOP(&datatype_definition_integerAPI.DATATYPE_DEFINITION_INTEGER_WOP)

		backRepoData.DATATYPE_DEFINITION_INTEGERAPIs = append(backRepoData.DATATYPE_DEFINITION_INTEGERAPIs, &datatype_definition_integerAPI)
	}

	for _, datatype_definition_realDB := range backRepo.BackRepoDATATYPE_DEFINITION_REAL.Map_DATATYPE_DEFINITION_REALDBID_DATATYPE_DEFINITION_REALDB {

		var datatype_definition_realAPI DATATYPE_DEFINITION_REALAPI
		datatype_definition_realAPI.ID = datatype_definition_realDB.ID
		datatype_definition_realAPI.DATATYPE_DEFINITION_REALPointersEncoding = datatype_definition_realDB.DATATYPE_DEFINITION_REALPointersEncoding
		datatype_definition_realDB.CopyBasicFieldsToDATATYPE_DEFINITION_REAL_WOP(&datatype_definition_realAPI.DATATYPE_DEFINITION_REAL_WOP)

		backRepoData.DATATYPE_DEFINITION_REALAPIs = append(backRepoData.DATATYPE_DEFINITION_REALAPIs, &datatype_definition_realAPI)
	}

	for _, datatype_definition_stringDB := range backRepo.BackRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB {

		var datatype_definition_stringAPI DATATYPE_DEFINITION_STRINGAPI
		datatype_definition_stringAPI.ID = datatype_definition_stringDB.ID
		datatype_definition_stringAPI.DATATYPE_DEFINITION_STRINGPointersEncoding = datatype_definition_stringDB.DATATYPE_DEFINITION_STRINGPointersEncoding
		datatype_definition_stringDB.CopyBasicFieldsToDATATYPE_DEFINITION_STRING_WOP(&datatype_definition_stringAPI.DATATYPE_DEFINITION_STRING_WOP)

		backRepoData.DATATYPE_DEFINITION_STRINGAPIs = append(backRepoData.DATATYPE_DEFINITION_STRINGAPIs, &datatype_definition_stringAPI)
	}

	for _, datatype_definition_xhtmlDB := range backRepo.BackRepoDATATYPE_DEFINITION_XHTML.Map_DATATYPE_DEFINITION_XHTMLDBID_DATATYPE_DEFINITION_XHTMLDB {

		var datatype_definition_xhtmlAPI DATATYPE_DEFINITION_XHTMLAPI
		datatype_definition_xhtmlAPI.ID = datatype_definition_xhtmlDB.ID
		datatype_definition_xhtmlAPI.DATATYPE_DEFINITION_XHTMLPointersEncoding = datatype_definition_xhtmlDB.DATATYPE_DEFINITION_XHTMLPointersEncoding
		datatype_definition_xhtmlDB.CopyBasicFieldsToDATATYPE_DEFINITION_XHTML_WOP(&datatype_definition_xhtmlAPI.DATATYPE_DEFINITION_XHTML_WOP)

		backRepoData.DATATYPE_DEFINITION_XHTMLAPIs = append(backRepoData.DATATYPE_DEFINITION_XHTMLAPIs, &datatype_definition_xhtmlAPI)
	}

	for _, embedded_valueDB := range backRepo.BackRepoEMBEDDED_VALUE.Map_EMBEDDED_VALUEDBID_EMBEDDED_VALUEDB {

		var embedded_valueAPI EMBEDDED_VALUEAPI
		embedded_valueAPI.ID = embedded_valueDB.ID
		embedded_valueAPI.EMBEDDED_VALUEPointersEncoding = embedded_valueDB.EMBEDDED_VALUEPointersEncoding
		embedded_valueDB.CopyBasicFieldsToEMBEDDED_VALUE_WOP(&embedded_valueAPI.EMBEDDED_VALUE_WOP)

		backRepoData.EMBEDDED_VALUEAPIs = append(backRepoData.EMBEDDED_VALUEAPIs, &embedded_valueAPI)
	}

	for _, enum_valueDB := range backRepo.BackRepoENUM_VALUE.Map_ENUM_VALUEDBID_ENUM_VALUEDB {

		var enum_valueAPI ENUM_VALUEAPI
		enum_valueAPI.ID = enum_valueDB.ID
		enum_valueAPI.ENUM_VALUEPointersEncoding = enum_valueDB.ENUM_VALUEPointersEncoding
		enum_valueDB.CopyBasicFieldsToENUM_VALUE_WOP(&enum_valueAPI.ENUM_VALUE_WOP)

		backRepoData.ENUM_VALUEAPIs = append(backRepoData.ENUM_VALUEAPIs, &enum_valueAPI)
	}

	for _, relation_groupDB := range backRepo.BackRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB {

		var relation_groupAPI RELATION_GROUPAPI
		relation_groupAPI.ID = relation_groupDB.ID
		relation_groupAPI.RELATION_GROUPPointersEncoding = relation_groupDB.RELATION_GROUPPointersEncoding
		relation_groupDB.CopyBasicFieldsToRELATION_GROUP_WOP(&relation_groupAPI.RELATION_GROUP_WOP)

		backRepoData.RELATION_GROUPAPIs = append(backRepoData.RELATION_GROUPAPIs, &relation_groupAPI)
	}

	for _, relation_group_typeDB := range backRepo.BackRepoRELATION_GROUP_TYPE.Map_RELATION_GROUP_TYPEDBID_RELATION_GROUP_TYPEDB {

		var relation_group_typeAPI RELATION_GROUP_TYPEAPI
		relation_group_typeAPI.ID = relation_group_typeDB.ID
		relation_group_typeAPI.RELATION_GROUP_TYPEPointersEncoding = relation_group_typeDB.RELATION_GROUP_TYPEPointersEncoding
		relation_group_typeDB.CopyBasicFieldsToRELATION_GROUP_TYPE_WOP(&relation_group_typeAPI.RELATION_GROUP_TYPE_WOP)

		backRepoData.RELATION_GROUP_TYPEAPIs = append(backRepoData.RELATION_GROUP_TYPEAPIs, &relation_group_typeAPI)
	}

	for _, req_ifDB := range backRepo.BackRepoREQ_IF.Map_REQ_IFDBID_REQ_IFDB {

		var req_ifAPI REQ_IFAPI
		req_ifAPI.ID = req_ifDB.ID
		req_ifAPI.REQ_IFPointersEncoding = req_ifDB.REQ_IFPointersEncoding
		req_ifDB.CopyBasicFieldsToREQ_IF_WOP(&req_ifAPI.REQ_IF_WOP)

		backRepoData.REQ_IFAPIs = append(backRepoData.REQ_IFAPIs, &req_ifAPI)
	}

	for _, req_if_contentDB := range backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB {

		var req_if_contentAPI REQ_IF_CONTENTAPI
		req_if_contentAPI.ID = req_if_contentDB.ID
		req_if_contentAPI.REQ_IF_CONTENTPointersEncoding = req_if_contentDB.REQ_IF_CONTENTPointersEncoding
		req_if_contentDB.CopyBasicFieldsToREQ_IF_CONTENT_WOP(&req_if_contentAPI.REQ_IF_CONTENT_WOP)

		backRepoData.REQ_IF_CONTENTAPIs = append(backRepoData.REQ_IF_CONTENTAPIs, &req_if_contentAPI)
	}

	for _, req_if_headerDB := range backRepo.BackRepoREQ_IF_HEADER.Map_REQ_IF_HEADERDBID_REQ_IF_HEADERDB {

		var req_if_headerAPI REQ_IF_HEADERAPI
		req_if_headerAPI.ID = req_if_headerDB.ID
		req_if_headerAPI.REQ_IF_HEADERPointersEncoding = req_if_headerDB.REQ_IF_HEADERPointersEncoding
		req_if_headerDB.CopyBasicFieldsToREQ_IF_HEADER_WOP(&req_if_headerAPI.REQ_IF_HEADER_WOP)

		backRepoData.REQ_IF_HEADERAPIs = append(backRepoData.REQ_IF_HEADERAPIs, &req_if_headerAPI)
	}

	for _, req_if_tool_extensionDB := range backRepo.BackRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB {

		var req_if_tool_extensionAPI REQ_IF_TOOL_EXTENSIONAPI
		req_if_tool_extensionAPI.ID = req_if_tool_extensionDB.ID
		req_if_tool_extensionAPI.REQ_IF_TOOL_EXTENSIONPointersEncoding = req_if_tool_extensionDB.REQ_IF_TOOL_EXTENSIONPointersEncoding
		req_if_tool_extensionDB.CopyBasicFieldsToREQ_IF_TOOL_EXTENSION_WOP(&req_if_tool_extensionAPI.REQ_IF_TOOL_EXTENSION_WOP)

		backRepoData.REQ_IF_TOOL_EXTENSIONAPIs = append(backRepoData.REQ_IF_TOOL_EXTENSIONAPIs, &req_if_tool_extensionAPI)
	}

	for _, specificationDB := range backRepo.BackRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB {

		var specificationAPI SPECIFICATIONAPI
		specificationAPI.ID = specificationDB.ID
		specificationAPI.SPECIFICATIONPointersEncoding = specificationDB.SPECIFICATIONPointersEncoding
		specificationDB.CopyBasicFieldsToSPECIFICATION_WOP(&specificationAPI.SPECIFICATION_WOP)

		backRepoData.SPECIFICATIONAPIs = append(backRepoData.SPECIFICATIONAPIs, &specificationAPI)
	}

	for _, specification_typeDB := range backRepo.BackRepoSPECIFICATION_TYPE.Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEDB {

		var specification_typeAPI SPECIFICATION_TYPEAPI
		specification_typeAPI.ID = specification_typeDB.ID
		specification_typeAPI.SPECIFICATION_TYPEPointersEncoding = specification_typeDB.SPECIFICATION_TYPEPointersEncoding
		specification_typeDB.CopyBasicFieldsToSPECIFICATION_TYPE_WOP(&specification_typeAPI.SPECIFICATION_TYPE_WOP)

		backRepoData.SPECIFICATION_TYPEAPIs = append(backRepoData.SPECIFICATION_TYPEAPIs, &specification_typeAPI)
	}

	for _, spec_hierarchyDB := range backRepo.BackRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB {

		var spec_hierarchyAPI SPEC_HIERARCHYAPI
		spec_hierarchyAPI.ID = spec_hierarchyDB.ID
		spec_hierarchyAPI.SPEC_HIERARCHYPointersEncoding = spec_hierarchyDB.SPEC_HIERARCHYPointersEncoding
		spec_hierarchyDB.CopyBasicFieldsToSPEC_HIERARCHY_WOP(&spec_hierarchyAPI.SPEC_HIERARCHY_WOP)

		backRepoData.SPEC_HIERARCHYAPIs = append(backRepoData.SPEC_HIERARCHYAPIs, &spec_hierarchyAPI)
	}

	for _, spec_objectDB := range backRepo.BackRepoSPEC_OBJECT.Map_SPEC_OBJECTDBID_SPEC_OBJECTDB {

		var spec_objectAPI SPEC_OBJECTAPI
		spec_objectAPI.ID = spec_objectDB.ID
		spec_objectAPI.SPEC_OBJECTPointersEncoding = spec_objectDB.SPEC_OBJECTPointersEncoding
		spec_objectDB.CopyBasicFieldsToSPEC_OBJECT_WOP(&spec_objectAPI.SPEC_OBJECT_WOP)

		backRepoData.SPEC_OBJECTAPIs = append(backRepoData.SPEC_OBJECTAPIs, &spec_objectAPI)
	}

	for _, spec_object_typeDB := range backRepo.BackRepoSPEC_OBJECT_TYPE.Map_SPEC_OBJECT_TYPEDBID_SPEC_OBJECT_TYPEDB {

		var spec_object_typeAPI SPEC_OBJECT_TYPEAPI
		spec_object_typeAPI.ID = spec_object_typeDB.ID
		spec_object_typeAPI.SPEC_OBJECT_TYPEPointersEncoding = spec_object_typeDB.SPEC_OBJECT_TYPEPointersEncoding
		spec_object_typeDB.CopyBasicFieldsToSPEC_OBJECT_TYPE_WOP(&spec_object_typeAPI.SPEC_OBJECT_TYPE_WOP)

		backRepoData.SPEC_OBJECT_TYPEAPIs = append(backRepoData.SPEC_OBJECT_TYPEAPIs, &spec_object_typeAPI)
	}

	for _, spec_relationDB := range backRepo.BackRepoSPEC_RELATION.Map_SPEC_RELATIONDBID_SPEC_RELATIONDB {

		var spec_relationAPI SPEC_RELATIONAPI
		spec_relationAPI.ID = spec_relationDB.ID
		spec_relationAPI.SPEC_RELATIONPointersEncoding = spec_relationDB.SPEC_RELATIONPointersEncoding
		spec_relationDB.CopyBasicFieldsToSPEC_RELATION_WOP(&spec_relationAPI.SPEC_RELATION_WOP)

		backRepoData.SPEC_RELATIONAPIs = append(backRepoData.SPEC_RELATIONAPIs, &spec_relationAPI)
	}

	for _, spec_relation_typeDB := range backRepo.BackRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB {

		var spec_relation_typeAPI SPEC_RELATION_TYPEAPI
		spec_relation_typeAPI.ID = spec_relation_typeDB.ID
		spec_relation_typeAPI.SPEC_RELATION_TYPEPointersEncoding = spec_relation_typeDB.SPEC_RELATION_TYPEPointersEncoding
		spec_relation_typeDB.CopyBasicFieldsToSPEC_RELATION_TYPE_WOP(&spec_relation_typeAPI.SPEC_RELATION_TYPE_WOP)

		backRepoData.SPEC_RELATION_TYPEAPIs = append(backRepoData.SPEC_RELATION_TYPEAPIs, &spec_relation_typeAPI)
	}

	for _, xhtml_contentDB := range backRepo.BackRepoXHTML_CONTENT.Map_XHTML_CONTENTDBID_XHTML_CONTENTDB {

		var xhtml_contentAPI XHTML_CONTENTAPI
		xhtml_contentAPI.ID = xhtml_contentDB.ID
		xhtml_contentAPI.XHTML_CONTENTPointersEncoding = xhtml_contentDB.XHTML_CONTENTPointersEncoding
		xhtml_contentDB.CopyBasicFieldsToXHTML_CONTENT_WOP(&xhtml_contentAPI.XHTML_CONTENT_WOP)

		backRepoData.XHTML_CONTENTAPIs = append(backRepoData.XHTML_CONTENTAPIs, &xhtml_contentAPI)
	}

	for _, xhtml_inlpres_typeDB := range backRepo.BackRepoXhtml_InlPres_type.Map_Xhtml_InlPres_typeDBID_Xhtml_InlPres_typeDB {

		var xhtml_inlpres_typeAPI Xhtml_InlPres_typeAPI
		xhtml_inlpres_typeAPI.ID = xhtml_inlpres_typeDB.ID
		xhtml_inlpres_typeAPI.Xhtml_InlPres_typePointersEncoding = xhtml_inlpres_typeDB.Xhtml_InlPres_typePointersEncoding
		xhtml_inlpres_typeDB.CopyBasicFieldsToXhtml_InlPres_type_WOP(&xhtml_inlpres_typeAPI.Xhtml_InlPres_type_WOP)

		backRepoData.Xhtml_InlPres_typeAPIs = append(backRepoData.Xhtml_InlPres_typeAPIs, &xhtml_inlpres_typeAPI)
	}

	for _, xhtml_a_typeDB := range backRepo.BackRepoXhtml_a_type.Map_Xhtml_a_typeDBID_Xhtml_a_typeDB {

		var xhtml_a_typeAPI Xhtml_a_typeAPI
		xhtml_a_typeAPI.ID = xhtml_a_typeDB.ID
		xhtml_a_typeAPI.Xhtml_a_typePointersEncoding = xhtml_a_typeDB.Xhtml_a_typePointersEncoding
		xhtml_a_typeDB.CopyBasicFieldsToXhtml_a_type_WOP(&xhtml_a_typeAPI.Xhtml_a_type_WOP)

		backRepoData.Xhtml_a_typeAPIs = append(backRepoData.Xhtml_a_typeAPIs, &xhtml_a_typeAPI)
	}

	for _, xhtml_abbr_typeDB := range backRepo.BackRepoXhtml_abbr_type.Map_Xhtml_abbr_typeDBID_Xhtml_abbr_typeDB {

		var xhtml_abbr_typeAPI Xhtml_abbr_typeAPI
		xhtml_abbr_typeAPI.ID = xhtml_abbr_typeDB.ID
		xhtml_abbr_typeAPI.Xhtml_abbr_typePointersEncoding = xhtml_abbr_typeDB.Xhtml_abbr_typePointersEncoding
		xhtml_abbr_typeDB.CopyBasicFieldsToXhtml_abbr_type_WOP(&xhtml_abbr_typeAPI.Xhtml_abbr_type_WOP)

		backRepoData.Xhtml_abbr_typeAPIs = append(backRepoData.Xhtml_abbr_typeAPIs, &xhtml_abbr_typeAPI)
	}

	for _, xhtml_acronym_typeDB := range backRepo.BackRepoXhtml_acronym_type.Map_Xhtml_acronym_typeDBID_Xhtml_acronym_typeDB {

		var xhtml_acronym_typeAPI Xhtml_acronym_typeAPI
		xhtml_acronym_typeAPI.ID = xhtml_acronym_typeDB.ID
		xhtml_acronym_typeAPI.Xhtml_acronym_typePointersEncoding = xhtml_acronym_typeDB.Xhtml_acronym_typePointersEncoding
		xhtml_acronym_typeDB.CopyBasicFieldsToXhtml_acronym_type_WOP(&xhtml_acronym_typeAPI.Xhtml_acronym_type_WOP)

		backRepoData.Xhtml_acronym_typeAPIs = append(backRepoData.Xhtml_acronym_typeAPIs, &xhtml_acronym_typeAPI)
	}

	for _, xhtml_address_typeDB := range backRepo.BackRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB {

		var xhtml_address_typeAPI Xhtml_address_typeAPI
		xhtml_address_typeAPI.ID = xhtml_address_typeDB.ID
		xhtml_address_typeAPI.Xhtml_address_typePointersEncoding = xhtml_address_typeDB.Xhtml_address_typePointersEncoding
		xhtml_address_typeDB.CopyBasicFieldsToXhtml_address_type_WOP(&xhtml_address_typeAPI.Xhtml_address_type_WOP)

		backRepoData.Xhtml_address_typeAPIs = append(backRepoData.Xhtml_address_typeAPIs, &xhtml_address_typeAPI)
	}

	for _, xhtml_blockquote_typeDB := range backRepo.BackRepoXhtml_blockquote_type.Map_Xhtml_blockquote_typeDBID_Xhtml_blockquote_typeDB {

		var xhtml_blockquote_typeAPI Xhtml_blockquote_typeAPI
		xhtml_blockquote_typeAPI.ID = xhtml_blockquote_typeDB.ID
		xhtml_blockquote_typeAPI.Xhtml_blockquote_typePointersEncoding = xhtml_blockquote_typeDB.Xhtml_blockquote_typePointersEncoding
		xhtml_blockquote_typeDB.CopyBasicFieldsToXhtml_blockquote_type_WOP(&xhtml_blockquote_typeAPI.Xhtml_blockquote_type_WOP)

		backRepoData.Xhtml_blockquote_typeAPIs = append(backRepoData.Xhtml_blockquote_typeAPIs, &xhtml_blockquote_typeAPI)
	}

	for _, xhtml_br_typeDB := range backRepo.BackRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB {

		var xhtml_br_typeAPI Xhtml_br_typeAPI
		xhtml_br_typeAPI.ID = xhtml_br_typeDB.ID
		xhtml_br_typeAPI.Xhtml_br_typePointersEncoding = xhtml_br_typeDB.Xhtml_br_typePointersEncoding
		xhtml_br_typeDB.CopyBasicFieldsToXhtml_br_type_WOP(&xhtml_br_typeAPI.Xhtml_br_type_WOP)

		backRepoData.Xhtml_br_typeAPIs = append(backRepoData.Xhtml_br_typeAPIs, &xhtml_br_typeAPI)
	}

	for _, xhtml_caption_typeDB := range backRepo.BackRepoXhtml_caption_type.Map_Xhtml_caption_typeDBID_Xhtml_caption_typeDB {

		var xhtml_caption_typeAPI Xhtml_caption_typeAPI
		xhtml_caption_typeAPI.ID = xhtml_caption_typeDB.ID
		xhtml_caption_typeAPI.Xhtml_caption_typePointersEncoding = xhtml_caption_typeDB.Xhtml_caption_typePointersEncoding
		xhtml_caption_typeDB.CopyBasicFieldsToXhtml_caption_type_WOP(&xhtml_caption_typeAPI.Xhtml_caption_type_WOP)

		backRepoData.Xhtml_caption_typeAPIs = append(backRepoData.Xhtml_caption_typeAPIs, &xhtml_caption_typeAPI)
	}

	for _, xhtml_cite_typeDB := range backRepo.BackRepoXhtml_cite_type.Map_Xhtml_cite_typeDBID_Xhtml_cite_typeDB {

		var xhtml_cite_typeAPI Xhtml_cite_typeAPI
		xhtml_cite_typeAPI.ID = xhtml_cite_typeDB.ID
		xhtml_cite_typeAPI.Xhtml_cite_typePointersEncoding = xhtml_cite_typeDB.Xhtml_cite_typePointersEncoding
		xhtml_cite_typeDB.CopyBasicFieldsToXhtml_cite_type_WOP(&xhtml_cite_typeAPI.Xhtml_cite_type_WOP)

		backRepoData.Xhtml_cite_typeAPIs = append(backRepoData.Xhtml_cite_typeAPIs, &xhtml_cite_typeAPI)
	}

	for _, xhtml_code_typeDB := range backRepo.BackRepoXhtml_code_type.Map_Xhtml_code_typeDBID_Xhtml_code_typeDB {

		var xhtml_code_typeAPI Xhtml_code_typeAPI
		xhtml_code_typeAPI.ID = xhtml_code_typeDB.ID
		xhtml_code_typeAPI.Xhtml_code_typePointersEncoding = xhtml_code_typeDB.Xhtml_code_typePointersEncoding
		xhtml_code_typeDB.CopyBasicFieldsToXhtml_code_type_WOP(&xhtml_code_typeAPI.Xhtml_code_type_WOP)

		backRepoData.Xhtml_code_typeAPIs = append(backRepoData.Xhtml_code_typeAPIs, &xhtml_code_typeAPI)
	}

	for _, xhtml_col_typeDB := range backRepo.BackRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB {

		var xhtml_col_typeAPI Xhtml_col_typeAPI
		xhtml_col_typeAPI.ID = xhtml_col_typeDB.ID
		xhtml_col_typeAPI.Xhtml_col_typePointersEncoding = xhtml_col_typeDB.Xhtml_col_typePointersEncoding
		xhtml_col_typeDB.CopyBasicFieldsToXhtml_col_type_WOP(&xhtml_col_typeAPI.Xhtml_col_type_WOP)

		backRepoData.Xhtml_col_typeAPIs = append(backRepoData.Xhtml_col_typeAPIs, &xhtml_col_typeAPI)
	}

	for _, xhtml_colgroup_typeDB := range backRepo.BackRepoXhtml_colgroup_type.Map_Xhtml_colgroup_typeDBID_Xhtml_colgroup_typeDB {

		var xhtml_colgroup_typeAPI Xhtml_colgroup_typeAPI
		xhtml_colgroup_typeAPI.ID = xhtml_colgroup_typeDB.ID
		xhtml_colgroup_typeAPI.Xhtml_colgroup_typePointersEncoding = xhtml_colgroup_typeDB.Xhtml_colgroup_typePointersEncoding
		xhtml_colgroup_typeDB.CopyBasicFieldsToXhtml_colgroup_type_WOP(&xhtml_colgroup_typeAPI.Xhtml_colgroup_type_WOP)

		backRepoData.Xhtml_colgroup_typeAPIs = append(backRepoData.Xhtml_colgroup_typeAPIs, &xhtml_colgroup_typeAPI)
	}

	for _, xhtml_dd_typeDB := range backRepo.BackRepoXhtml_dd_type.Map_Xhtml_dd_typeDBID_Xhtml_dd_typeDB {

		var xhtml_dd_typeAPI Xhtml_dd_typeAPI
		xhtml_dd_typeAPI.ID = xhtml_dd_typeDB.ID
		xhtml_dd_typeAPI.Xhtml_dd_typePointersEncoding = xhtml_dd_typeDB.Xhtml_dd_typePointersEncoding
		xhtml_dd_typeDB.CopyBasicFieldsToXhtml_dd_type_WOP(&xhtml_dd_typeAPI.Xhtml_dd_type_WOP)

		backRepoData.Xhtml_dd_typeAPIs = append(backRepoData.Xhtml_dd_typeAPIs, &xhtml_dd_typeAPI)
	}

	for _, xhtml_dfn_typeDB := range backRepo.BackRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB {

		var xhtml_dfn_typeAPI Xhtml_dfn_typeAPI
		xhtml_dfn_typeAPI.ID = xhtml_dfn_typeDB.ID
		xhtml_dfn_typeAPI.Xhtml_dfn_typePointersEncoding = xhtml_dfn_typeDB.Xhtml_dfn_typePointersEncoding
		xhtml_dfn_typeDB.CopyBasicFieldsToXhtml_dfn_type_WOP(&xhtml_dfn_typeAPI.Xhtml_dfn_type_WOP)

		backRepoData.Xhtml_dfn_typeAPIs = append(backRepoData.Xhtml_dfn_typeAPIs, &xhtml_dfn_typeAPI)
	}

	for _, xhtml_div_typeDB := range backRepo.BackRepoXhtml_div_type.Map_Xhtml_div_typeDBID_Xhtml_div_typeDB {

		var xhtml_div_typeAPI Xhtml_div_typeAPI
		xhtml_div_typeAPI.ID = xhtml_div_typeDB.ID
		xhtml_div_typeAPI.Xhtml_div_typePointersEncoding = xhtml_div_typeDB.Xhtml_div_typePointersEncoding
		xhtml_div_typeDB.CopyBasicFieldsToXhtml_div_type_WOP(&xhtml_div_typeAPI.Xhtml_div_type_WOP)

		backRepoData.Xhtml_div_typeAPIs = append(backRepoData.Xhtml_div_typeAPIs, &xhtml_div_typeAPI)
	}

	for _, xhtml_dl_typeDB := range backRepo.BackRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB {

		var xhtml_dl_typeAPI Xhtml_dl_typeAPI
		xhtml_dl_typeAPI.ID = xhtml_dl_typeDB.ID
		xhtml_dl_typeAPI.Xhtml_dl_typePointersEncoding = xhtml_dl_typeDB.Xhtml_dl_typePointersEncoding
		xhtml_dl_typeDB.CopyBasicFieldsToXhtml_dl_type_WOP(&xhtml_dl_typeAPI.Xhtml_dl_type_WOP)

		backRepoData.Xhtml_dl_typeAPIs = append(backRepoData.Xhtml_dl_typeAPIs, &xhtml_dl_typeAPI)
	}

	for _, xhtml_dt_typeDB := range backRepo.BackRepoXhtml_dt_type.Map_Xhtml_dt_typeDBID_Xhtml_dt_typeDB {

		var xhtml_dt_typeAPI Xhtml_dt_typeAPI
		xhtml_dt_typeAPI.ID = xhtml_dt_typeDB.ID
		xhtml_dt_typeAPI.Xhtml_dt_typePointersEncoding = xhtml_dt_typeDB.Xhtml_dt_typePointersEncoding
		xhtml_dt_typeDB.CopyBasicFieldsToXhtml_dt_type_WOP(&xhtml_dt_typeAPI.Xhtml_dt_type_WOP)

		backRepoData.Xhtml_dt_typeAPIs = append(backRepoData.Xhtml_dt_typeAPIs, &xhtml_dt_typeAPI)
	}

	for _, xhtml_edit_typeDB := range backRepo.BackRepoXhtml_edit_type.Map_Xhtml_edit_typeDBID_Xhtml_edit_typeDB {

		var xhtml_edit_typeAPI Xhtml_edit_typeAPI
		xhtml_edit_typeAPI.ID = xhtml_edit_typeDB.ID
		xhtml_edit_typeAPI.Xhtml_edit_typePointersEncoding = xhtml_edit_typeDB.Xhtml_edit_typePointersEncoding
		xhtml_edit_typeDB.CopyBasicFieldsToXhtml_edit_type_WOP(&xhtml_edit_typeAPI.Xhtml_edit_type_WOP)

		backRepoData.Xhtml_edit_typeAPIs = append(backRepoData.Xhtml_edit_typeAPIs, &xhtml_edit_typeAPI)
	}

	for _, xhtml_em_typeDB := range backRepo.BackRepoXhtml_em_type.Map_Xhtml_em_typeDBID_Xhtml_em_typeDB {

		var xhtml_em_typeAPI Xhtml_em_typeAPI
		xhtml_em_typeAPI.ID = xhtml_em_typeDB.ID
		xhtml_em_typeAPI.Xhtml_em_typePointersEncoding = xhtml_em_typeDB.Xhtml_em_typePointersEncoding
		xhtml_em_typeDB.CopyBasicFieldsToXhtml_em_type_WOP(&xhtml_em_typeAPI.Xhtml_em_type_WOP)

		backRepoData.Xhtml_em_typeAPIs = append(backRepoData.Xhtml_em_typeAPIs, &xhtml_em_typeAPI)
	}

	for _, xhtml_h1_typeDB := range backRepo.BackRepoXhtml_h1_type.Map_Xhtml_h1_typeDBID_Xhtml_h1_typeDB {

		var xhtml_h1_typeAPI Xhtml_h1_typeAPI
		xhtml_h1_typeAPI.ID = xhtml_h1_typeDB.ID
		xhtml_h1_typeAPI.Xhtml_h1_typePointersEncoding = xhtml_h1_typeDB.Xhtml_h1_typePointersEncoding
		xhtml_h1_typeDB.CopyBasicFieldsToXhtml_h1_type_WOP(&xhtml_h1_typeAPI.Xhtml_h1_type_WOP)

		backRepoData.Xhtml_h1_typeAPIs = append(backRepoData.Xhtml_h1_typeAPIs, &xhtml_h1_typeAPI)
	}

	for _, xhtml_h2_typeDB := range backRepo.BackRepoXhtml_h2_type.Map_Xhtml_h2_typeDBID_Xhtml_h2_typeDB {

		var xhtml_h2_typeAPI Xhtml_h2_typeAPI
		xhtml_h2_typeAPI.ID = xhtml_h2_typeDB.ID
		xhtml_h2_typeAPI.Xhtml_h2_typePointersEncoding = xhtml_h2_typeDB.Xhtml_h2_typePointersEncoding
		xhtml_h2_typeDB.CopyBasicFieldsToXhtml_h2_type_WOP(&xhtml_h2_typeAPI.Xhtml_h2_type_WOP)

		backRepoData.Xhtml_h2_typeAPIs = append(backRepoData.Xhtml_h2_typeAPIs, &xhtml_h2_typeAPI)
	}

	for _, xhtml_h3_typeDB := range backRepo.BackRepoXhtml_h3_type.Map_Xhtml_h3_typeDBID_Xhtml_h3_typeDB {

		var xhtml_h3_typeAPI Xhtml_h3_typeAPI
		xhtml_h3_typeAPI.ID = xhtml_h3_typeDB.ID
		xhtml_h3_typeAPI.Xhtml_h3_typePointersEncoding = xhtml_h3_typeDB.Xhtml_h3_typePointersEncoding
		xhtml_h3_typeDB.CopyBasicFieldsToXhtml_h3_type_WOP(&xhtml_h3_typeAPI.Xhtml_h3_type_WOP)

		backRepoData.Xhtml_h3_typeAPIs = append(backRepoData.Xhtml_h3_typeAPIs, &xhtml_h3_typeAPI)
	}

	for _, xhtml_h4_typeDB := range backRepo.BackRepoXhtml_h4_type.Map_Xhtml_h4_typeDBID_Xhtml_h4_typeDB {

		var xhtml_h4_typeAPI Xhtml_h4_typeAPI
		xhtml_h4_typeAPI.ID = xhtml_h4_typeDB.ID
		xhtml_h4_typeAPI.Xhtml_h4_typePointersEncoding = xhtml_h4_typeDB.Xhtml_h4_typePointersEncoding
		xhtml_h4_typeDB.CopyBasicFieldsToXhtml_h4_type_WOP(&xhtml_h4_typeAPI.Xhtml_h4_type_WOP)

		backRepoData.Xhtml_h4_typeAPIs = append(backRepoData.Xhtml_h4_typeAPIs, &xhtml_h4_typeAPI)
	}

	for _, xhtml_h5_typeDB := range backRepo.BackRepoXhtml_h5_type.Map_Xhtml_h5_typeDBID_Xhtml_h5_typeDB {

		var xhtml_h5_typeAPI Xhtml_h5_typeAPI
		xhtml_h5_typeAPI.ID = xhtml_h5_typeDB.ID
		xhtml_h5_typeAPI.Xhtml_h5_typePointersEncoding = xhtml_h5_typeDB.Xhtml_h5_typePointersEncoding
		xhtml_h5_typeDB.CopyBasicFieldsToXhtml_h5_type_WOP(&xhtml_h5_typeAPI.Xhtml_h5_type_WOP)

		backRepoData.Xhtml_h5_typeAPIs = append(backRepoData.Xhtml_h5_typeAPIs, &xhtml_h5_typeAPI)
	}

	for _, xhtml_h6_typeDB := range backRepo.BackRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB {

		var xhtml_h6_typeAPI Xhtml_h6_typeAPI
		xhtml_h6_typeAPI.ID = xhtml_h6_typeDB.ID
		xhtml_h6_typeAPI.Xhtml_h6_typePointersEncoding = xhtml_h6_typeDB.Xhtml_h6_typePointersEncoding
		xhtml_h6_typeDB.CopyBasicFieldsToXhtml_h6_type_WOP(&xhtml_h6_typeAPI.Xhtml_h6_type_WOP)

		backRepoData.Xhtml_h6_typeAPIs = append(backRepoData.Xhtml_h6_typeAPIs, &xhtml_h6_typeAPI)
	}

	for _, xhtml_heading_typeDB := range backRepo.BackRepoXhtml_heading_type.Map_Xhtml_heading_typeDBID_Xhtml_heading_typeDB {

		var xhtml_heading_typeAPI Xhtml_heading_typeAPI
		xhtml_heading_typeAPI.ID = xhtml_heading_typeDB.ID
		xhtml_heading_typeAPI.Xhtml_heading_typePointersEncoding = xhtml_heading_typeDB.Xhtml_heading_typePointersEncoding
		xhtml_heading_typeDB.CopyBasicFieldsToXhtml_heading_type_WOP(&xhtml_heading_typeAPI.Xhtml_heading_type_WOP)

		backRepoData.Xhtml_heading_typeAPIs = append(backRepoData.Xhtml_heading_typeAPIs, &xhtml_heading_typeAPI)
	}

	for _, xhtml_hr_typeDB := range backRepo.BackRepoXhtml_hr_type.Map_Xhtml_hr_typeDBID_Xhtml_hr_typeDB {

		var xhtml_hr_typeAPI Xhtml_hr_typeAPI
		xhtml_hr_typeAPI.ID = xhtml_hr_typeDB.ID
		xhtml_hr_typeAPI.Xhtml_hr_typePointersEncoding = xhtml_hr_typeDB.Xhtml_hr_typePointersEncoding
		xhtml_hr_typeDB.CopyBasicFieldsToXhtml_hr_type_WOP(&xhtml_hr_typeAPI.Xhtml_hr_type_WOP)

		backRepoData.Xhtml_hr_typeAPIs = append(backRepoData.Xhtml_hr_typeAPIs, &xhtml_hr_typeAPI)
	}

	for _, xhtml_kbd_typeDB := range backRepo.BackRepoXhtml_kbd_type.Map_Xhtml_kbd_typeDBID_Xhtml_kbd_typeDB {

		var xhtml_kbd_typeAPI Xhtml_kbd_typeAPI
		xhtml_kbd_typeAPI.ID = xhtml_kbd_typeDB.ID
		xhtml_kbd_typeAPI.Xhtml_kbd_typePointersEncoding = xhtml_kbd_typeDB.Xhtml_kbd_typePointersEncoding
		xhtml_kbd_typeDB.CopyBasicFieldsToXhtml_kbd_type_WOP(&xhtml_kbd_typeAPI.Xhtml_kbd_type_WOP)

		backRepoData.Xhtml_kbd_typeAPIs = append(backRepoData.Xhtml_kbd_typeAPIs, &xhtml_kbd_typeAPI)
	}

	for _, xhtml_li_typeDB := range backRepo.BackRepoXhtml_li_type.Map_Xhtml_li_typeDBID_Xhtml_li_typeDB {

		var xhtml_li_typeAPI Xhtml_li_typeAPI
		xhtml_li_typeAPI.ID = xhtml_li_typeDB.ID
		xhtml_li_typeAPI.Xhtml_li_typePointersEncoding = xhtml_li_typeDB.Xhtml_li_typePointersEncoding
		xhtml_li_typeDB.CopyBasicFieldsToXhtml_li_type_WOP(&xhtml_li_typeAPI.Xhtml_li_type_WOP)

		backRepoData.Xhtml_li_typeAPIs = append(backRepoData.Xhtml_li_typeAPIs, &xhtml_li_typeAPI)
	}

	for _, xhtml_object_typeDB := range backRepo.BackRepoXhtml_object_type.Map_Xhtml_object_typeDBID_Xhtml_object_typeDB {

		var xhtml_object_typeAPI Xhtml_object_typeAPI
		xhtml_object_typeAPI.ID = xhtml_object_typeDB.ID
		xhtml_object_typeAPI.Xhtml_object_typePointersEncoding = xhtml_object_typeDB.Xhtml_object_typePointersEncoding
		xhtml_object_typeDB.CopyBasicFieldsToXhtml_object_type_WOP(&xhtml_object_typeAPI.Xhtml_object_type_WOP)

		backRepoData.Xhtml_object_typeAPIs = append(backRepoData.Xhtml_object_typeAPIs, &xhtml_object_typeAPI)
	}

	for _, xhtml_ol_typeDB := range backRepo.BackRepoXhtml_ol_type.Map_Xhtml_ol_typeDBID_Xhtml_ol_typeDB {

		var xhtml_ol_typeAPI Xhtml_ol_typeAPI
		xhtml_ol_typeAPI.ID = xhtml_ol_typeDB.ID
		xhtml_ol_typeAPI.Xhtml_ol_typePointersEncoding = xhtml_ol_typeDB.Xhtml_ol_typePointersEncoding
		xhtml_ol_typeDB.CopyBasicFieldsToXhtml_ol_type_WOP(&xhtml_ol_typeAPI.Xhtml_ol_type_WOP)

		backRepoData.Xhtml_ol_typeAPIs = append(backRepoData.Xhtml_ol_typeAPIs, &xhtml_ol_typeAPI)
	}

	for _, xhtml_p_typeDB := range backRepo.BackRepoXhtml_p_type.Map_Xhtml_p_typeDBID_Xhtml_p_typeDB {

		var xhtml_p_typeAPI Xhtml_p_typeAPI
		xhtml_p_typeAPI.ID = xhtml_p_typeDB.ID
		xhtml_p_typeAPI.Xhtml_p_typePointersEncoding = xhtml_p_typeDB.Xhtml_p_typePointersEncoding
		xhtml_p_typeDB.CopyBasicFieldsToXhtml_p_type_WOP(&xhtml_p_typeAPI.Xhtml_p_type_WOP)

		backRepoData.Xhtml_p_typeAPIs = append(backRepoData.Xhtml_p_typeAPIs, &xhtml_p_typeAPI)
	}

	for _, xhtml_param_typeDB := range backRepo.BackRepoXhtml_param_type.Map_Xhtml_param_typeDBID_Xhtml_param_typeDB {

		var xhtml_param_typeAPI Xhtml_param_typeAPI
		xhtml_param_typeAPI.ID = xhtml_param_typeDB.ID
		xhtml_param_typeAPI.Xhtml_param_typePointersEncoding = xhtml_param_typeDB.Xhtml_param_typePointersEncoding
		xhtml_param_typeDB.CopyBasicFieldsToXhtml_param_type_WOP(&xhtml_param_typeAPI.Xhtml_param_type_WOP)

		backRepoData.Xhtml_param_typeAPIs = append(backRepoData.Xhtml_param_typeAPIs, &xhtml_param_typeAPI)
	}

	for _, xhtml_pre_typeDB := range backRepo.BackRepoXhtml_pre_type.Map_Xhtml_pre_typeDBID_Xhtml_pre_typeDB {

		var xhtml_pre_typeAPI Xhtml_pre_typeAPI
		xhtml_pre_typeAPI.ID = xhtml_pre_typeDB.ID
		xhtml_pre_typeAPI.Xhtml_pre_typePointersEncoding = xhtml_pre_typeDB.Xhtml_pre_typePointersEncoding
		xhtml_pre_typeDB.CopyBasicFieldsToXhtml_pre_type_WOP(&xhtml_pre_typeAPI.Xhtml_pre_type_WOP)

		backRepoData.Xhtml_pre_typeAPIs = append(backRepoData.Xhtml_pre_typeAPIs, &xhtml_pre_typeAPI)
	}

	for _, xhtml_q_typeDB := range backRepo.BackRepoXhtml_q_type.Map_Xhtml_q_typeDBID_Xhtml_q_typeDB {

		var xhtml_q_typeAPI Xhtml_q_typeAPI
		xhtml_q_typeAPI.ID = xhtml_q_typeDB.ID
		xhtml_q_typeAPI.Xhtml_q_typePointersEncoding = xhtml_q_typeDB.Xhtml_q_typePointersEncoding
		xhtml_q_typeDB.CopyBasicFieldsToXhtml_q_type_WOP(&xhtml_q_typeAPI.Xhtml_q_type_WOP)

		backRepoData.Xhtml_q_typeAPIs = append(backRepoData.Xhtml_q_typeAPIs, &xhtml_q_typeAPI)
	}

	for _, xhtml_samp_typeDB := range backRepo.BackRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB {

		var xhtml_samp_typeAPI Xhtml_samp_typeAPI
		xhtml_samp_typeAPI.ID = xhtml_samp_typeDB.ID
		xhtml_samp_typeAPI.Xhtml_samp_typePointersEncoding = xhtml_samp_typeDB.Xhtml_samp_typePointersEncoding
		xhtml_samp_typeDB.CopyBasicFieldsToXhtml_samp_type_WOP(&xhtml_samp_typeAPI.Xhtml_samp_type_WOP)

		backRepoData.Xhtml_samp_typeAPIs = append(backRepoData.Xhtml_samp_typeAPIs, &xhtml_samp_typeAPI)
	}

	for _, xhtml_span_typeDB := range backRepo.BackRepoXhtml_span_type.Map_Xhtml_span_typeDBID_Xhtml_span_typeDB {

		var xhtml_span_typeAPI Xhtml_span_typeAPI
		xhtml_span_typeAPI.ID = xhtml_span_typeDB.ID
		xhtml_span_typeAPI.Xhtml_span_typePointersEncoding = xhtml_span_typeDB.Xhtml_span_typePointersEncoding
		xhtml_span_typeDB.CopyBasicFieldsToXhtml_span_type_WOP(&xhtml_span_typeAPI.Xhtml_span_type_WOP)

		backRepoData.Xhtml_span_typeAPIs = append(backRepoData.Xhtml_span_typeAPIs, &xhtml_span_typeAPI)
	}

	for _, xhtml_strong_typeDB := range backRepo.BackRepoXhtml_strong_type.Map_Xhtml_strong_typeDBID_Xhtml_strong_typeDB {

		var xhtml_strong_typeAPI Xhtml_strong_typeAPI
		xhtml_strong_typeAPI.ID = xhtml_strong_typeDB.ID
		xhtml_strong_typeAPI.Xhtml_strong_typePointersEncoding = xhtml_strong_typeDB.Xhtml_strong_typePointersEncoding
		xhtml_strong_typeDB.CopyBasicFieldsToXhtml_strong_type_WOP(&xhtml_strong_typeAPI.Xhtml_strong_type_WOP)

		backRepoData.Xhtml_strong_typeAPIs = append(backRepoData.Xhtml_strong_typeAPIs, &xhtml_strong_typeAPI)
	}

	for _, xhtml_table_typeDB := range backRepo.BackRepoXhtml_table_type.Map_Xhtml_table_typeDBID_Xhtml_table_typeDB {

		var xhtml_table_typeAPI Xhtml_table_typeAPI
		xhtml_table_typeAPI.ID = xhtml_table_typeDB.ID
		xhtml_table_typeAPI.Xhtml_table_typePointersEncoding = xhtml_table_typeDB.Xhtml_table_typePointersEncoding
		xhtml_table_typeDB.CopyBasicFieldsToXhtml_table_type_WOP(&xhtml_table_typeAPI.Xhtml_table_type_WOP)

		backRepoData.Xhtml_table_typeAPIs = append(backRepoData.Xhtml_table_typeAPIs, &xhtml_table_typeAPI)
	}

	for _, xhtml_tbody_typeDB := range backRepo.BackRepoXhtml_tbody_type.Map_Xhtml_tbody_typeDBID_Xhtml_tbody_typeDB {

		var xhtml_tbody_typeAPI Xhtml_tbody_typeAPI
		xhtml_tbody_typeAPI.ID = xhtml_tbody_typeDB.ID
		xhtml_tbody_typeAPI.Xhtml_tbody_typePointersEncoding = xhtml_tbody_typeDB.Xhtml_tbody_typePointersEncoding
		xhtml_tbody_typeDB.CopyBasicFieldsToXhtml_tbody_type_WOP(&xhtml_tbody_typeAPI.Xhtml_tbody_type_WOP)

		backRepoData.Xhtml_tbody_typeAPIs = append(backRepoData.Xhtml_tbody_typeAPIs, &xhtml_tbody_typeAPI)
	}

	for _, xhtml_td_typeDB := range backRepo.BackRepoXhtml_td_type.Map_Xhtml_td_typeDBID_Xhtml_td_typeDB {

		var xhtml_td_typeAPI Xhtml_td_typeAPI
		xhtml_td_typeAPI.ID = xhtml_td_typeDB.ID
		xhtml_td_typeAPI.Xhtml_td_typePointersEncoding = xhtml_td_typeDB.Xhtml_td_typePointersEncoding
		xhtml_td_typeDB.CopyBasicFieldsToXhtml_td_type_WOP(&xhtml_td_typeAPI.Xhtml_td_type_WOP)

		backRepoData.Xhtml_td_typeAPIs = append(backRepoData.Xhtml_td_typeAPIs, &xhtml_td_typeAPI)
	}

	for _, xhtml_tfoot_typeDB := range backRepo.BackRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB {

		var xhtml_tfoot_typeAPI Xhtml_tfoot_typeAPI
		xhtml_tfoot_typeAPI.ID = xhtml_tfoot_typeDB.ID
		xhtml_tfoot_typeAPI.Xhtml_tfoot_typePointersEncoding = xhtml_tfoot_typeDB.Xhtml_tfoot_typePointersEncoding
		xhtml_tfoot_typeDB.CopyBasicFieldsToXhtml_tfoot_type_WOP(&xhtml_tfoot_typeAPI.Xhtml_tfoot_type_WOP)

		backRepoData.Xhtml_tfoot_typeAPIs = append(backRepoData.Xhtml_tfoot_typeAPIs, &xhtml_tfoot_typeAPI)
	}

	for _, xhtml_th_typeDB := range backRepo.BackRepoXhtml_th_type.Map_Xhtml_th_typeDBID_Xhtml_th_typeDB {

		var xhtml_th_typeAPI Xhtml_th_typeAPI
		xhtml_th_typeAPI.ID = xhtml_th_typeDB.ID
		xhtml_th_typeAPI.Xhtml_th_typePointersEncoding = xhtml_th_typeDB.Xhtml_th_typePointersEncoding
		xhtml_th_typeDB.CopyBasicFieldsToXhtml_th_type_WOP(&xhtml_th_typeAPI.Xhtml_th_type_WOP)

		backRepoData.Xhtml_th_typeAPIs = append(backRepoData.Xhtml_th_typeAPIs, &xhtml_th_typeAPI)
	}

	for _, xhtml_thead_typeDB := range backRepo.BackRepoXhtml_thead_type.Map_Xhtml_thead_typeDBID_Xhtml_thead_typeDB {

		var xhtml_thead_typeAPI Xhtml_thead_typeAPI
		xhtml_thead_typeAPI.ID = xhtml_thead_typeDB.ID
		xhtml_thead_typeAPI.Xhtml_thead_typePointersEncoding = xhtml_thead_typeDB.Xhtml_thead_typePointersEncoding
		xhtml_thead_typeDB.CopyBasicFieldsToXhtml_thead_type_WOP(&xhtml_thead_typeAPI.Xhtml_thead_type_WOP)

		backRepoData.Xhtml_thead_typeAPIs = append(backRepoData.Xhtml_thead_typeAPIs, &xhtml_thead_typeAPI)
	}

	for _, xhtml_tr_typeDB := range backRepo.BackRepoXhtml_tr_type.Map_Xhtml_tr_typeDBID_Xhtml_tr_typeDB {

		var xhtml_tr_typeAPI Xhtml_tr_typeAPI
		xhtml_tr_typeAPI.ID = xhtml_tr_typeDB.ID
		xhtml_tr_typeAPI.Xhtml_tr_typePointersEncoding = xhtml_tr_typeDB.Xhtml_tr_typePointersEncoding
		xhtml_tr_typeDB.CopyBasicFieldsToXhtml_tr_type_WOP(&xhtml_tr_typeAPI.Xhtml_tr_type_WOP)

		backRepoData.Xhtml_tr_typeAPIs = append(backRepoData.Xhtml_tr_typeAPIs, &xhtml_tr_typeAPI)
	}

	for _, xhtml_ul_typeDB := range backRepo.BackRepoXhtml_ul_type.Map_Xhtml_ul_typeDBID_Xhtml_ul_typeDB {

		var xhtml_ul_typeAPI Xhtml_ul_typeAPI
		xhtml_ul_typeAPI.ID = xhtml_ul_typeDB.ID
		xhtml_ul_typeAPI.Xhtml_ul_typePointersEncoding = xhtml_ul_typeDB.Xhtml_ul_typePointersEncoding
		xhtml_ul_typeDB.CopyBasicFieldsToXhtml_ul_type_WOP(&xhtml_ul_typeAPI.Xhtml_ul_type_WOP)

		backRepoData.Xhtml_ul_typeAPIs = append(backRepoData.Xhtml_ul_typeAPIs, &xhtml_ul_typeAPI)
	}

	for _, xhtml_var_typeDB := range backRepo.BackRepoXhtml_var_type.Map_Xhtml_var_typeDBID_Xhtml_var_typeDB {

		var xhtml_var_typeAPI Xhtml_var_typeAPI
		xhtml_var_typeAPI.ID = xhtml_var_typeDB.ID
		xhtml_var_typeAPI.Xhtml_var_typePointersEncoding = xhtml_var_typeDB.Xhtml_var_typePointersEncoding
		xhtml_var_typeDB.CopyBasicFieldsToXhtml_var_type_WOP(&xhtml_var_typeAPI.Xhtml_var_type_WOP)

		backRepoData.Xhtml_var_typeAPIs = append(backRepoData.Xhtml_var_typeAPIs, &xhtml_var_typeAPI)
	}

}
