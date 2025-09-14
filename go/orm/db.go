// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongreqif/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	alternative_idDBs map[uint]*ALTERNATIVE_IDDB

	nextIDALTERNATIVE_IDDB uint

	attribute_definition_booleanDBs map[uint]*ATTRIBUTE_DEFINITION_BOOLEANDB

	nextIDATTRIBUTE_DEFINITION_BOOLEANDB uint

	attribute_definition_dateDBs map[uint]*ATTRIBUTE_DEFINITION_DATEDB

	nextIDATTRIBUTE_DEFINITION_DATEDB uint

	attribute_definition_enumerationDBs map[uint]*ATTRIBUTE_DEFINITION_ENUMERATIONDB

	nextIDATTRIBUTE_DEFINITION_ENUMERATIONDB uint

	attribute_definition_integerDBs map[uint]*ATTRIBUTE_DEFINITION_INTEGERDB

	nextIDATTRIBUTE_DEFINITION_INTEGERDB uint

	attribute_definition_realDBs map[uint]*ATTRIBUTE_DEFINITION_REALDB

	nextIDATTRIBUTE_DEFINITION_REALDB uint

	attribute_definition_stringDBs map[uint]*ATTRIBUTE_DEFINITION_STRINGDB

	nextIDATTRIBUTE_DEFINITION_STRINGDB uint

	attribute_definition_xhtmlDBs map[uint]*ATTRIBUTE_DEFINITION_XHTMLDB

	nextIDATTRIBUTE_DEFINITION_XHTMLDB uint

	attribute_value_booleanDBs map[uint]*ATTRIBUTE_VALUE_BOOLEANDB

	nextIDATTRIBUTE_VALUE_BOOLEANDB uint

	attribute_value_dateDBs map[uint]*ATTRIBUTE_VALUE_DATEDB

	nextIDATTRIBUTE_VALUE_DATEDB uint

	attribute_value_enumerationDBs map[uint]*ATTRIBUTE_VALUE_ENUMERATIONDB

	nextIDATTRIBUTE_VALUE_ENUMERATIONDB uint

	attribute_value_integerDBs map[uint]*ATTRIBUTE_VALUE_INTEGERDB

	nextIDATTRIBUTE_VALUE_INTEGERDB uint

	attribute_value_realDBs map[uint]*ATTRIBUTE_VALUE_REALDB

	nextIDATTRIBUTE_VALUE_REALDB uint

	attribute_value_stringDBs map[uint]*ATTRIBUTE_VALUE_STRINGDB

	nextIDATTRIBUTE_VALUE_STRINGDB uint

	attribute_value_xhtmlDBs map[uint]*ATTRIBUTE_VALUE_XHTMLDB

	nextIDATTRIBUTE_VALUE_XHTMLDB uint

	a_alternative_idDBs map[uint]*A_ALTERNATIVE_IDDB

	nextIDA_ALTERNATIVE_IDDB uint

	a_attribute_definition_boolean_refDBs map[uint]*A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB

	nextIDA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB uint

	a_attribute_definition_date_refDBs map[uint]*A_ATTRIBUTE_DEFINITION_DATE_REFDB

	nextIDA_ATTRIBUTE_DEFINITION_DATE_REFDB uint

	a_attribute_definition_enumeration_refDBs map[uint]*A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB

	nextIDA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB uint

	a_attribute_definition_integer_refDBs map[uint]*A_ATTRIBUTE_DEFINITION_INTEGER_REFDB

	nextIDA_ATTRIBUTE_DEFINITION_INTEGER_REFDB uint

	a_attribute_definition_real_refDBs map[uint]*A_ATTRIBUTE_DEFINITION_REAL_REFDB

	nextIDA_ATTRIBUTE_DEFINITION_REAL_REFDB uint

	a_attribute_definition_string_refDBs map[uint]*A_ATTRIBUTE_DEFINITION_STRING_REFDB

	nextIDA_ATTRIBUTE_DEFINITION_STRING_REFDB uint

	a_attribute_definition_xhtml_refDBs map[uint]*A_ATTRIBUTE_DEFINITION_XHTML_REFDB

	nextIDA_ATTRIBUTE_DEFINITION_XHTML_REFDB uint

	a_attribute_value_booleanDBs map[uint]*A_ATTRIBUTE_VALUE_BOOLEANDB

	nextIDA_ATTRIBUTE_VALUE_BOOLEANDB uint

	a_attribute_value_dateDBs map[uint]*A_ATTRIBUTE_VALUE_DATEDB

	nextIDA_ATTRIBUTE_VALUE_DATEDB uint

	a_attribute_value_enumerationDBs map[uint]*A_ATTRIBUTE_VALUE_ENUMERATIONDB

	nextIDA_ATTRIBUTE_VALUE_ENUMERATIONDB uint

	a_attribute_value_integerDBs map[uint]*A_ATTRIBUTE_VALUE_INTEGERDB

	nextIDA_ATTRIBUTE_VALUE_INTEGERDB uint

	a_attribute_value_realDBs map[uint]*A_ATTRIBUTE_VALUE_REALDB

	nextIDA_ATTRIBUTE_VALUE_REALDB uint

	a_attribute_value_stringDBs map[uint]*A_ATTRIBUTE_VALUE_STRINGDB

	nextIDA_ATTRIBUTE_VALUE_STRINGDB uint

	a_attribute_value_xhtmlDBs map[uint]*A_ATTRIBUTE_VALUE_XHTMLDB

	nextIDA_ATTRIBUTE_VALUE_XHTMLDB uint

	a_attribute_value_xhtml_1DBs map[uint]*A_ATTRIBUTE_VALUE_XHTML_1DB

	nextIDA_ATTRIBUTE_VALUE_XHTML_1DB uint

	a_childrenDBs map[uint]*A_CHILDRENDB

	nextIDA_CHILDRENDB uint

	a_core_contentDBs map[uint]*A_CORE_CONTENTDB

	nextIDA_CORE_CONTENTDB uint

	a_datatypesDBs map[uint]*A_DATATYPESDB

	nextIDA_DATATYPESDB uint

	a_datatype_definition_boolean_refDBs map[uint]*A_DATATYPE_DEFINITION_BOOLEAN_REFDB

	nextIDA_DATATYPE_DEFINITION_BOOLEAN_REFDB uint

	a_datatype_definition_date_refDBs map[uint]*A_DATATYPE_DEFINITION_DATE_REFDB

	nextIDA_DATATYPE_DEFINITION_DATE_REFDB uint

	a_datatype_definition_enumeration_refDBs map[uint]*A_DATATYPE_DEFINITION_ENUMERATION_REFDB

	nextIDA_DATATYPE_DEFINITION_ENUMERATION_REFDB uint

	a_datatype_definition_integer_refDBs map[uint]*A_DATATYPE_DEFINITION_INTEGER_REFDB

	nextIDA_DATATYPE_DEFINITION_INTEGER_REFDB uint

	a_datatype_definition_real_refDBs map[uint]*A_DATATYPE_DEFINITION_REAL_REFDB

	nextIDA_DATATYPE_DEFINITION_REAL_REFDB uint

	a_datatype_definition_string_refDBs map[uint]*A_DATATYPE_DEFINITION_STRING_REFDB

	nextIDA_DATATYPE_DEFINITION_STRING_REFDB uint

	a_datatype_definition_xhtml_refDBs map[uint]*A_DATATYPE_DEFINITION_XHTML_REFDB

	nextIDA_DATATYPE_DEFINITION_XHTML_REFDB uint

	a_editable_attsDBs map[uint]*A_EDITABLE_ATTSDB

	nextIDA_EDITABLE_ATTSDB uint

	a_enum_value_refDBs map[uint]*A_ENUM_VALUE_REFDB

	nextIDA_ENUM_VALUE_REFDB uint

	a_objectDBs map[uint]*A_OBJECTDB

	nextIDA_OBJECTDB uint

	a_propertiesDBs map[uint]*A_PROPERTIESDB

	nextIDA_PROPERTIESDB uint

	a_relation_group_type_refDBs map[uint]*A_RELATION_GROUP_TYPE_REFDB

	nextIDA_RELATION_GROUP_TYPE_REFDB uint

	a_source_1DBs map[uint]*A_SOURCE_1DB

	nextIDA_SOURCE_1DB uint

	a_source_specification_1DBs map[uint]*A_SOURCE_SPECIFICATION_1DB

	nextIDA_SOURCE_SPECIFICATION_1DB uint

	a_specificationsDBs map[uint]*A_SPECIFICATIONSDB

	nextIDA_SPECIFICATIONSDB uint

	a_specification_type_refDBs map[uint]*A_SPECIFICATION_TYPE_REFDB

	nextIDA_SPECIFICATION_TYPE_REFDB uint

	a_specified_valuesDBs map[uint]*A_SPECIFIED_VALUESDB

	nextIDA_SPECIFIED_VALUESDB uint

	a_spec_attributesDBs map[uint]*A_SPEC_ATTRIBUTESDB

	nextIDA_SPEC_ATTRIBUTESDB uint

	a_spec_objectsDBs map[uint]*A_SPEC_OBJECTSDB

	nextIDA_SPEC_OBJECTSDB uint

	a_spec_object_type_refDBs map[uint]*A_SPEC_OBJECT_TYPE_REFDB

	nextIDA_SPEC_OBJECT_TYPE_REFDB uint

	a_spec_relationsDBs map[uint]*A_SPEC_RELATIONSDB

	nextIDA_SPEC_RELATIONSDB uint

	a_spec_relation_groupsDBs map[uint]*A_SPEC_RELATION_GROUPSDB

	nextIDA_SPEC_RELATION_GROUPSDB uint

	a_spec_relation_refDBs map[uint]*A_SPEC_RELATION_REFDB

	nextIDA_SPEC_RELATION_REFDB uint

	a_spec_relation_type_refDBs map[uint]*A_SPEC_RELATION_TYPE_REFDB

	nextIDA_SPEC_RELATION_TYPE_REFDB uint

	a_spec_typesDBs map[uint]*A_SPEC_TYPESDB

	nextIDA_SPEC_TYPESDB uint

	a_the_headerDBs map[uint]*A_THE_HEADERDB

	nextIDA_THE_HEADERDB uint

	a_tool_extensionsDBs map[uint]*A_TOOL_EXTENSIONSDB

	nextIDA_TOOL_EXTENSIONSDB uint

	datatype_definition_booleanDBs map[uint]*DATATYPE_DEFINITION_BOOLEANDB

	nextIDDATATYPE_DEFINITION_BOOLEANDB uint

	datatype_definition_dateDBs map[uint]*DATATYPE_DEFINITION_DATEDB

	nextIDDATATYPE_DEFINITION_DATEDB uint

	datatype_definition_enumerationDBs map[uint]*DATATYPE_DEFINITION_ENUMERATIONDB

	nextIDDATATYPE_DEFINITION_ENUMERATIONDB uint

	datatype_definition_integerDBs map[uint]*DATATYPE_DEFINITION_INTEGERDB

	nextIDDATATYPE_DEFINITION_INTEGERDB uint

	datatype_definition_realDBs map[uint]*DATATYPE_DEFINITION_REALDB

	nextIDDATATYPE_DEFINITION_REALDB uint

	datatype_definition_stringDBs map[uint]*DATATYPE_DEFINITION_STRINGDB

	nextIDDATATYPE_DEFINITION_STRINGDB uint

	datatype_definition_xhtmlDBs map[uint]*DATATYPE_DEFINITION_XHTMLDB

	nextIDDATATYPE_DEFINITION_XHTMLDB uint

	embedded_valueDBs map[uint]*EMBEDDED_VALUEDB

	nextIDEMBEDDED_VALUEDB uint

	enum_valueDBs map[uint]*ENUM_VALUEDB

	nextIDENUM_VALUEDB uint

	map_attribute_definition_boolean_showinsubjectentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB uint

	map_attribute_definition_boolean_showintableentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB uint

	map_attribute_definition_boolean_showintitleentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB uint

	map_attribute_definition_date_showinsubjectentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB uint

	map_attribute_definition_date_showintableentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB uint

	map_attribute_definition_date_showintitleentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB uint

	map_attribute_definition_enumeration_showinsubjectentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB uint

	map_attribute_definition_enumeration_showintableentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB uint

	map_attribute_definition_enumeration_showintitleentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB uint

	map_attribute_definition_integer_showinsubjectentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB uint

	map_attribute_definition_integer_showintableentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB uint

	map_attribute_definition_integer_showintitleentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB uint

	map_attribute_definition_real_showinsubjectentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB uint

	map_attribute_definition_real_showintableentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB uint

	map_attribute_definition_real_showintitleentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB uint

	map_attribute_definition_string_showinsubjectentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB uint

	map_attribute_definition_string_showintableentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB uint

	map_attribute_definition_string_showintitleentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB uint

	map_attribute_definition_xhtml_showinsubjectentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB uint

	map_attribute_definition_xhtml_showintableentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB uint

	map_attribute_definition_xhtml_showintitleentryDBs map[uint]*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB

	nextIDMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB uint

	map_specification_nodes_expandedentryDBs map[uint]*Map_SPECIFICATION_Nodes_expandedEntryDB

	nextIDMap_SPECIFICATION_Nodes_expandedEntryDB uint

	map_spec_object_type_isnodeexpandedentryDBs map[uint]*Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB

	nextIDMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB uint

	map_spec_object_type_showidentifierentryDBs map[uint]*Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB

	nextIDMap_SPEC_OBJECT_TYPE_showIdentifierEntryDB uint

	map_spec_object_type_shownameentryDBs map[uint]*Map_SPEC_OBJECT_TYPE_showNameEntryDB

	nextIDMap_SPEC_OBJECT_TYPE_showNameEntryDB uint

	relation_groupDBs map[uint]*RELATION_GROUPDB

	nextIDRELATION_GROUPDB uint

	relation_group_typeDBs map[uint]*RELATION_GROUP_TYPEDB

	nextIDRELATION_GROUP_TYPEDB uint

	req_ifDBs map[uint]*REQ_IFDB

	nextIDREQ_IFDB uint

	req_if_contentDBs map[uint]*REQ_IF_CONTENTDB

	nextIDREQ_IF_CONTENTDB uint

	req_if_headerDBs map[uint]*REQ_IF_HEADERDB

	nextIDREQ_IF_HEADERDB uint

	req_if_tool_extensionDBs map[uint]*REQ_IF_TOOL_EXTENSIONDB

	nextIDREQ_IF_TOOL_EXTENSIONDB uint

	renderingconfigurationDBs map[uint]*RenderingConfigurationDB

	nextIDRenderingConfigurationDB uint

	specificationDBs map[uint]*SPECIFICATIONDB

	nextIDSPECIFICATIONDB uint

	specification_typeDBs map[uint]*SPECIFICATION_TYPEDB

	nextIDSPECIFICATION_TYPEDB uint

	spec_hierarchyDBs map[uint]*SPEC_HIERARCHYDB

	nextIDSPEC_HIERARCHYDB uint

	spec_objectDBs map[uint]*SPEC_OBJECTDB

	nextIDSPEC_OBJECTDB uint

	spec_object_typeDBs map[uint]*SPEC_OBJECT_TYPEDB

	nextIDSPEC_OBJECT_TYPEDB uint

	spec_relationDBs map[uint]*SPEC_RELATIONDB

	nextIDSPEC_RELATIONDB uint

	spec_relation_typeDBs map[uint]*SPEC_RELATION_TYPEDB

	nextIDSPEC_RELATION_TYPEDB uint

	staticwebsiteDBs map[uint]*StaticWebSiteDB

	nextIDStaticWebSiteDB uint

	staticwebsitechapterDBs map[uint]*StaticWebSiteChapterDB

	nextIDStaticWebSiteChapterDB uint

	staticwebsitegeneratedimageDBs map[uint]*StaticWebSiteGeneratedImageDB

	nextIDStaticWebSiteGeneratedImageDB uint

	staticwebsiteimageDBs map[uint]*StaticWebSiteImageDB

	nextIDStaticWebSiteImageDB uint

	staticwebsiteparagraphDBs map[uint]*StaticWebSiteParagraphDB

	nextIDStaticWebSiteParagraphDB uint

	xhtml_contentDBs map[uint]*XHTML_CONTENTDB

	nextIDXHTML_CONTENTDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		alternative_idDBs: make(map[uint]*ALTERNATIVE_IDDB),

		attribute_definition_booleanDBs: make(map[uint]*ATTRIBUTE_DEFINITION_BOOLEANDB),

		attribute_definition_dateDBs: make(map[uint]*ATTRIBUTE_DEFINITION_DATEDB),

		attribute_definition_enumerationDBs: make(map[uint]*ATTRIBUTE_DEFINITION_ENUMERATIONDB),

		attribute_definition_integerDBs: make(map[uint]*ATTRIBUTE_DEFINITION_INTEGERDB),

		attribute_definition_realDBs: make(map[uint]*ATTRIBUTE_DEFINITION_REALDB),

		attribute_definition_stringDBs: make(map[uint]*ATTRIBUTE_DEFINITION_STRINGDB),

		attribute_definition_xhtmlDBs: make(map[uint]*ATTRIBUTE_DEFINITION_XHTMLDB),

		attribute_value_booleanDBs: make(map[uint]*ATTRIBUTE_VALUE_BOOLEANDB),

		attribute_value_dateDBs: make(map[uint]*ATTRIBUTE_VALUE_DATEDB),

		attribute_value_enumerationDBs: make(map[uint]*ATTRIBUTE_VALUE_ENUMERATIONDB),

		attribute_value_integerDBs: make(map[uint]*ATTRIBUTE_VALUE_INTEGERDB),

		attribute_value_realDBs: make(map[uint]*ATTRIBUTE_VALUE_REALDB),

		attribute_value_stringDBs: make(map[uint]*ATTRIBUTE_VALUE_STRINGDB),

		attribute_value_xhtmlDBs: make(map[uint]*ATTRIBUTE_VALUE_XHTMLDB),

		a_alternative_idDBs: make(map[uint]*A_ALTERNATIVE_IDDB),

		a_attribute_definition_boolean_refDBs: make(map[uint]*A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB),

		a_attribute_definition_date_refDBs: make(map[uint]*A_ATTRIBUTE_DEFINITION_DATE_REFDB),

		a_attribute_definition_enumeration_refDBs: make(map[uint]*A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB),

		a_attribute_definition_integer_refDBs: make(map[uint]*A_ATTRIBUTE_DEFINITION_INTEGER_REFDB),

		a_attribute_definition_real_refDBs: make(map[uint]*A_ATTRIBUTE_DEFINITION_REAL_REFDB),

		a_attribute_definition_string_refDBs: make(map[uint]*A_ATTRIBUTE_DEFINITION_STRING_REFDB),

		a_attribute_definition_xhtml_refDBs: make(map[uint]*A_ATTRIBUTE_DEFINITION_XHTML_REFDB),

		a_attribute_value_booleanDBs: make(map[uint]*A_ATTRIBUTE_VALUE_BOOLEANDB),

		a_attribute_value_dateDBs: make(map[uint]*A_ATTRIBUTE_VALUE_DATEDB),

		a_attribute_value_enumerationDBs: make(map[uint]*A_ATTRIBUTE_VALUE_ENUMERATIONDB),

		a_attribute_value_integerDBs: make(map[uint]*A_ATTRIBUTE_VALUE_INTEGERDB),

		a_attribute_value_realDBs: make(map[uint]*A_ATTRIBUTE_VALUE_REALDB),

		a_attribute_value_stringDBs: make(map[uint]*A_ATTRIBUTE_VALUE_STRINGDB),

		a_attribute_value_xhtmlDBs: make(map[uint]*A_ATTRIBUTE_VALUE_XHTMLDB),

		a_attribute_value_xhtml_1DBs: make(map[uint]*A_ATTRIBUTE_VALUE_XHTML_1DB),

		a_childrenDBs: make(map[uint]*A_CHILDRENDB),

		a_core_contentDBs: make(map[uint]*A_CORE_CONTENTDB),

		a_datatypesDBs: make(map[uint]*A_DATATYPESDB),

		a_datatype_definition_boolean_refDBs: make(map[uint]*A_DATATYPE_DEFINITION_BOOLEAN_REFDB),

		a_datatype_definition_date_refDBs: make(map[uint]*A_DATATYPE_DEFINITION_DATE_REFDB),

		a_datatype_definition_enumeration_refDBs: make(map[uint]*A_DATATYPE_DEFINITION_ENUMERATION_REFDB),

		a_datatype_definition_integer_refDBs: make(map[uint]*A_DATATYPE_DEFINITION_INTEGER_REFDB),

		a_datatype_definition_real_refDBs: make(map[uint]*A_DATATYPE_DEFINITION_REAL_REFDB),

		a_datatype_definition_string_refDBs: make(map[uint]*A_DATATYPE_DEFINITION_STRING_REFDB),

		a_datatype_definition_xhtml_refDBs: make(map[uint]*A_DATATYPE_DEFINITION_XHTML_REFDB),

		a_editable_attsDBs: make(map[uint]*A_EDITABLE_ATTSDB),

		a_enum_value_refDBs: make(map[uint]*A_ENUM_VALUE_REFDB),

		a_objectDBs: make(map[uint]*A_OBJECTDB),

		a_propertiesDBs: make(map[uint]*A_PROPERTIESDB),

		a_relation_group_type_refDBs: make(map[uint]*A_RELATION_GROUP_TYPE_REFDB),

		a_source_1DBs: make(map[uint]*A_SOURCE_1DB),

		a_source_specification_1DBs: make(map[uint]*A_SOURCE_SPECIFICATION_1DB),

		a_specificationsDBs: make(map[uint]*A_SPECIFICATIONSDB),

		a_specification_type_refDBs: make(map[uint]*A_SPECIFICATION_TYPE_REFDB),

		a_specified_valuesDBs: make(map[uint]*A_SPECIFIED_VALUESDB),

		a_spec_attributesDBs: make(map[uint]*A_SPEC_ATTRIBUTESDB),

		a_spec_objectsDBs: make(map[uint]*A_SPEC_OBJECTSDB),

		a_spec_object_type_refDBs: make(map[uint]*A_SPEC_OBJECT_TYPE_REFDB),

		a_spec_relationsDBs: make(map[uint]*A_SPEC_RELATIONSDB),

		a_spec_relation_groupsDBs: make(map[uint]*A_SPEC_RELATION_GROUPSDB),

		a_spec_relation_refDBs: make(map[uint]*A_SPEC_RELATION_REFDB),

		a_spec_relation_type_refDBs: make(map[uint]*A_SPEC_RELATION_TYPE_REFDB),

		a_spec_typesDBs: make(map[uint]*A_SPEC_TYPESDB),

		a_the_headerDBs: make(map[uint]*A_THE_HEADERDB),

		a_tool_extensionsDBs: make(map[uint]*A_TOOL_EXTENSIONSDB),

		datatype_definition_booleanDBs: make(map[uint]*DATATYPE_DEFINITION_BOOLEANDB),

		datatype_definition_dateDBs: make(map[uint]*DATATYPE_DEFINITION_DATEDB),

		datatype_definition_enumerationDBs: make(map[uint]*DATATYPE_DEFINITION_ENUMERATIONDB),

		datatype_definition_integerDBs: make(map[uint]*DATATYPE_DEFINITION_INTEGERDB),

		datatype_definition_realDBs: make(map[uint]*DATATYPE_DEFINITION_REALDB),

		datatype_definition_stringDBs: make(map[uint]*DATATYPE_DEFINITION_STRINGDB),

		datatype_definition_xhtmlDBs: make(map[uint]*DATATYPE_DEFINITION_XHTMLDB),

		embedded_valueDBs: make(map[uint]*EMBEDDED_VALUEDB),

		enum_valueDBs: make(map[uint]*ENUM_VALUEDB),

		map_attribute_definition_boolean_showinsubjectentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB),

		map_attribute_definition_boolean_showintableentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB),

		map_attribute_definition_boolean_showintitleentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB),

		map_attribute_definition_date_showinsubjectentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB),

		map_attribute_definition_date_showintableentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB),

		map_attribute_definition_date_showintitleentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB),

		map_attribute_definition_enumeration_showinsubjectentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB),

		map_attribute_definition_enumeration_showintableentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB),

		map_attribute_definition_enumeration_showintitleentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB),

		map_attribute_definition_integer_showinsubjectentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB),

		map_attribute_definition_integer_showintableentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB),

		map_attribute_definition_integer_showintitleentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB),

		map_attribute_definition_real_showinsubjectentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB),

		map_attribute_definition_real_showintableentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB),

		map_attribute_definition_real_showintitleentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB),

		map_attribute_definition_string_showinsubjectentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB),

		map_attribute_definition_string_showintableentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB),

		map_attribute_definition_string_showintitleentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB),

		map_attribute_definition_xhtml_showinsubjectentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB),

		map_attribute_definition_xhtml_showintableentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB),

		map_attribute_definition_xhtml_showintitleentryDBs: make(map[uint]*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB),

		map_specification_nodes_expandedentryDBs: make(map[uint]*Map_SPECIFICATION_Nodes_expandedEntryDB),

		map_spec_object_type_isnodeexpandedentryDBs: make(map[uint]*Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB),

		map_spec_object_type_showidentifierentryDBs: make(map[uint]*Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB),

		map_spec_object_type_shownameentryDBs: make(map[uint]*Map_SPEC_OBJECT_TYPE_showNameEntryDB),

		relation_groupDBs: make(map[uint]*RELATION_GROUPDB),

		relation_group_typeDBs: make(map[uint]*RELATION_GROUP_TYPEDB),

		req_ifDBs: make(map[uint]*REQ_IFDB),

		req_if_contentDBs: make(map[uint]*REQ_IF_CONTENTDB),

		req_if_headerDBs: make(map[uint]*REQ_IF_HEADERDB),

		req_if_tool_extensionDBs: make(map[uint]*REQ_IF_TOOL_EXTENSIONDB),

		renderingconfigurationDBs: make(map[uint]*RenderingConfigurationDB),

		specificationDBs: make(map[uint]*SPECIFICATIONDB),

		specification_typeDBs: make(map[uint]*SPECIFICATION_TYPEDB),

		spec_hierarchyDBs: make(map[uint]*SPEC_HIERARCHYDB),

		spec_objectDBs: make(map[uint]*SPEC_OBJECTDB),

		spec_object_typeDBs: make(map[uint]*SPEC_OBJECT_TYPEDB),

		spec_relationDBs: make(map[uint]*SPEC_RELATIONDB),

		spec_relation_typeDBs: make(map[uint]*SPEC_RELATION_TYPEDB),

		staticwebsiteDBs: make(map[uint]*StaticWebSiteDB),

		staticwebsitechapterDBs: make(map[uint]*StaticWebSiteChapterDB),

		staticwebsitegeneratedimageDBs: make(map[uint]*StaticWebSiteGeneratedImageDB),

		staticwebsiteimageDBs: make(map[uint]*StaticWebSiteImageDB),

		staticwebsiteparagraphDBs: make(map[uint]*StaticWebSiteParagraphDB),

		xhtml_contentDBs: make(map[uint]*XHTML_CONTENTDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *ALTERNATIVE_IDDB:
		db.nextIDALTERNATIVE_IDDB++
		v.ID = db.nextIDALTERNATIVE_IDDB
		db.alternative_idDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_BOOLEANDB:
		db.nextIDATTRIBUTE_DEFINITION_BOOLEANDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_BOOLEANDB
		db.attribute_definition_booleanDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_DATEDB:
		db.nextIDATTRIBUTE_DEFINITION_DATEDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_DATEDB
		db.attribute_definition_dateDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		db.nextIDATTRIBUTE_DEFINITION_ENUMERATIONDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_ENUMERATIONDB
		db.attribute_definition_enumerationDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_INTEGERDB:
		db.nextIDATTRIBUTE_DEFINITION_INTEGERDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_INTEGERDB
		db.attribute_definition_integerDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_REALDB:
		db.nextIDATTRIBUTE_DEFINITION_REALDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_REALDB
		db.attribute_definition_realDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_STRINGDB:
		db.nextIDATTRIBUTE_DEFINITION_STRINGDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_STRINGDB
		db.attribute_definition_stringDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_XHTMLDB:
		db.nextIDATTRIBUTE_DEFINITION_XHTMLDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_XHTMLDB
		db.attribute_definition_xhtmlDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_BOOLEANDB:
		db.nextIDATTRIBUTE_VALUE_BOOLEANDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_BOOLEANDB
		db.attribute_value_booleanDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_DATEDB:
		db.nextIDATTRIBUTE_VALUE_DATEDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_DATEDB
		db.attribute_value_dateDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_ENUMERATIONDB:
		db.nextIDATTRIBUTE_VALUE_ENUMERATIONDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_ENUMERATIONDB
		db.attribute_value_enumerationDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_INTEGERDB:
		db.nextIDATTRIBUTE_VALUE_INTEGERDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_INTEGERDB
		db.attribute_value_integerDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_REALDB:
		db.nextIDATTRIBUTE_VALUE_REALDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_REALDB
		db.attribute_value_realDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_STRINGDB:
		db.nextIDATTRIBUTE_VALUE_STRINGDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_STRINGDB
		db.attribute_value_stringDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_XHTMLDB:
		db.nextIDATTRIBUTE_VALUE_XHTMLDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_XHTMLDB
		db.attribute_value_xhtmlDBs[v.ID] = v
	case *A_ALTERNATIVE_IDDB:
		db.nextIDA_ALTERNATIVE_IDDB++
		v.ID = db.nextIDA_ALTERNATIVE_IDDB
		db.a_alternative_idDBs[v.ID] = v
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB:
		db.nextIDA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB++
		v.ID = db.nextIDA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB
		db.a_attribute_definition_boolean_refDBs[v.ID] = v
	case *A_ATTRIBUTE_DEFINITION_DATE_REFDB:
		db.nextIDA_ATTRIBUTE_DEFINITION_DATE_REFDB++
		v.ID = db.nextIDA_ATTRIBUTE_DEFINITION_DATE_REFDB
		db.a_attribute_definition_date_refDBs[v.ID] = v
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB:
		db.nextIDA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB++
		v.ID = db.nextIDA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB
		db.a_attribute_definition_enumeration_refDBs[v.ID] = v
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REFDB:
		db.nextIDA_ATTRIBUTE_DEFINITION_INTEGER_REFDB++
		v.ID = db.nextIDA_ATTRIBUTE_DEFINITION_INTEGER_REFDB
		db.a_attribute_definition_integer_refDBs[v.ID] = v
	case *A_ATTRIBUTE_DEFINITION_REAL_REFDB:
		db.nextIDA_ATTRIBUTE_DEFINITION_REAL_REFDB++
		v.ID = db.nextIDA_ATTRIBUTE_DEFINITION_REAL_REFDB
		db.a_attribute_definition_real_refDBs[v.ID] = v
	case *A_ATTRIBUTE_DEFINITION_STRING_REFDB:
		db.nextIDA_ATTRIBUTE_DEFINITION_STRING_REFDB++
		v.ID = db.nextIDA_ATTRIBUTE_DEFINITION_STRING_REFDB
		db.a_attribute_definition_string_refDBs[v.ID] = v
	case *A_ATTRIBUTE_DEFINITION_XHTML_REFDB:
		db.nextIDA_ATTRIBUTE_DEFINITION_XHTML_REFDB++
		v.ID = db.nextIDA_ATTRIBUTE_DEFINITION_XHTML_REFDB
		db.a_attribute_definition_xhtml_refDBs[v.ID] = v
	case *A_ATTRIBUTE_VALUE_BOOLEANDB:
		db.nextIDA_ATTRIBUTE_VALUE_BOOLEANDB++
		v.ID = db.nextIDA_ATTRIBUTE_VALUE_BOOLEANDB
		db.a_attribute_value_booleanDBs[v.ID] = v
	case *A_ATTRIBUTE_VALUE_DATEDB:
		db.nextIDA_ATTRIBUTE_VALUE_DATEDB++
		v.ID = db.nextIDA_ATTRIBUTE_VALUE_DATEDB
		db.a_attribute_value_dateDBs[v.ID] = v
	case *A_ATTRIBUTE_VALUE_ENUMERATIONDB:
		db.nextIDA_ATTRIBUTE_VALUE_ENUMERATIONDB++
		v.ID = db.nextIDA_ATTRIBUTE_VALUE_ENUMERATIONDB
		db.a_attribute_value_enumerationDBs[v.ID] = v
	case *A_ATTRIBUTE_VALUE_INTEGERDB:
		db.nextIDA_ATTRIBUTE_VALUE_INTEGERDB++
		v.ID = db.nextIDA_ATTRIBUTE_VALUE_INTEGERDB
		db.a_attribute_value_integerDBs[v.ID] = v
	case *A_ATTRIBUTE_VALUE_REALDB:
		db.nextIDA_ATTRIBUTE_VALUE_REALDB++
		v.ID = db.nextIDA_ATTRIBUTE_VALUE_REALDB
		db.a_attribute_value_realDBs[v.ID] = v
	case *A_ATTRIBUTE_VALUE_STRINGDB:
		db.nextIDA_ATTRIBUTE_VALUE_STRINGDB++
		v.ID = db.nextIDA_ATTRIBUTE_VALUE_STRINGDB
		db.a_attribute_value_stringDBs[v.ID] = v
	case *A_ATTRIBUTE_VALUE_XHTMLDB:
		db.nextIDA_ATTRIBUTE_VALUE_XHTMLDB++
		v.ID = db.nextIDA_ATTRIBUTE_VALUE_XHTMLDB
		db.a_attribute_value_xhtmlDBs[v.ID] = v
	case *A_ATTRIBUTE_VALUE_XHTML_1DB:
		db.nextIDA_ATTRIBUTE_VALUE_XHTML_1DB++
		v.ID = db.nextIDA_ATTRIBUTE_VALUE_XHTML_1DB
		db.a_attribute_value_xhtml_1DBs[v.ID] = v
	case *A_CHILDRENDB:
		db.nextIDA_CHILDRENDB++
		v.ID = db.nextIDA_CHILDRENDB
		db.a_childrenDBs[v.ID] = v
	case *A_CORE_CONTENTDB:
		db.nextIDA_CORE_CONTENTDB++
		v.ID = db.nextIDA_CORE_CONTENTDB
		db.a_core_contentDBs[v.ID] = v
	case *A_DATATYPESDB:
		db.nextIDA_DATATYPESDB++
		v.ID = db.nextIDA_DATATYPESDB
		db.a_datatypesDBs[v.ID] = v
	case *A_DATATYPE_DEFINITION_BOOLEAN_REFDB:
		db.nextIDA_DATATYPE_DEFINITION_BOOLEAN_REFDB++
		v.ID = db.nextIDA_DATATYPE_DEFINITION_BOOLEAN_REFDB
		db.a_datatype_definition_boolean_refDBs[v.ID] = v
	case *A_DATATYPE_DEFINITION_DATE_REFDB:
		db.nextIDA_DATATYPE_DEFINITION_DATE_REFDB++
		v.ID = db.nextIDA_DATATYPE_DEFINITION_DATE_REFDB
		db.a_datatype_definition_date_refDBs[v.ID] = v
	case *A_DATATYPE_DEFINITION_ENUMERATION_REFDB:
		db.nextIDA_DATATYPE_DEFINITION_ENUMERATION_REFDB++
		v.ID = db.nextIDA_DATATYPE_DEFINITION_ENUMERATION_REFDB
		db.a_datatype_definition_enumeration_refDBs[v.ID] = v
	case *A_DATATYPE_DEFINITION_INTEGER_REFDB:
		db.nextIDA_DATATYPE_DEFINITION_INTEGER_REFDB++
		v.ID = db.nextIDA_DATATYPE_DEFINITION_INTEGER_REFDB
		db.a_datatype_definition_integer_refDBs[v.ID] = v
	case *A_DATATYPE_DEFINITION_REAL_REFDB:
		db.nextIDA_DATATYPE_DEFINITION_REAL_REFDB++
		v.ID = db.nextIDA_DATATYPE_DEFINITION_REAL_REFDB
		db.a_datatype_definition_real_refDBs[v.ID] = v
	case *A_DATATYPE_DEFINITION_STRING_REFDB:
		db.nextIDA_DATATYPE_DEFINITION_STRING_REFDB++
		v.ID = db.nextIDA_DATATYPE_DEFINITION_STRING_REFDB
		db.a_datatype_definition_string_refDBs[v.ID] = v
	case *A_DATATYPE_DEFINITION_XHTML_REFDB:
		db.nextIDA_DATATYPE_DEFINITION_XHTML_REFDB++
		v.ID = db.nextIDA_DATATYPE_DEFINITION_XHTML_REFDB
		db.a_datatype_definition_xhtml_refDBs[v.ID] = v
	case *A_EDITABLE_ATTSDB:
		db.nextIDA_EDITABLE_ATTSDB++
		v.ID = db.nextIDA_EDITABLE_ATTSDB
		db.a_editable_attsDBs[v.ID] = v
	case *A_ENUM_VALUE_REFDB:
		db.nextIDA_ENUM_VALUE_REFDB++
		v.ID = db.nextIDA_ENUM_VALUE_REFDB
		db.a_enum_value_refDBs[v.ID] = v
	case *A_OBJECTDB:
		db.nextIDA_OBJECTDB++
		v.ID = db.nextIDA_OBJECTDB
		db.a_objectDBs[v.ID] = v
	case *A_PROPERTIESDB:
		db.nextIDA_PROPERTIESDB++
		v.ID = db.nextIDA_PROPERTIESDB
		db.a_propertiesDBs[v.ID] = v
	case *A_RELATION_GROUP_TYPE_REFDB:
		db.nextIDA_RELATION_GROUP_TYPE_REFDB++
		v.ID = db.nextIDA_RELATION_GROUP_TYPE_REFDB
		db.a_relation_group_type_refDBs[v.ID] = v
	case *A_SOURCE_1DB:
		db.nextIDA_SOURCE_1DB++
		v.ID = db.nextIDA_SOURCE_1DB
		db.a_source_1DBs[v.ID] = v
	case *A_SOURCE_SPECIFICATION_1DB:
		db.nextIDA_SOURCE_SPECIFICATION_1DB++
		v.ID = db.nextIDA_SOURCE_SPECIFICATION_1DB
		db.a_source_specification_1DBs[v.ID] = v
	case *A_SPECIFICATIONSDB:
		db.nextIDA_SPECIFICATIONSDB++
		v.ID = db.nextIDA_SPECIFICATIONSDB
		db.a_specificationsDBs[v.ID] = v
	case *A_SPECIFICATION_TYPE_REFDB:
		db.nextIDA_SPECIFICATION_TYPE_REFDB++
		v.ID = db.nextIDA_SPECIFICATION_TYPE_REFDB
		db.a_specification_type_refDBs[v.ID] = v
	case *A_SPECIFIED_VALUESDB:
		db.nextIDA_SPECIFIED_VALUESDB++
		v.ID = db.nextIDA_SPECIFIED_VALUESDB
		db.a_specified_valuesDBs[v.ID] = v
	case *A_SPEC_ATTRIBUTESDB:
		db.nextIDA_SPEC_ATTRIBUTESDB++
		v.ID = db.nextIDA_SPEC_ATTRIBUTESDB
		db.a_spec_attributesDBs[v.ID] = v
	case *A_SPEC_OBJECTSDB:
		db.nextIDA_SPEC_OBJECTSDB++
		v.ID = db.nextIDA_SPEC_OBJECTSDB
		db.a_spec_objectsDBs[v.ID] = v
	case *A_SPEC_OBJECT_TYPE_REFDB:
		db.nextIDA_SPEC_OBJECT_TYPE_REFDB++
		v.ID = db.nextIDA_SPEC_OBJECT_TYPE_REFDB
		db.a_spec_object_type_refDBs[v.ID] = v
	case *A_SPEC_RELATIONSDB:
		db.nextIDA_SPEC_RELATIONSDB++
		v.ID = db.nextIDA_SPEC_RELATIONSDB
		db.a_spec_relationsDBs[v.ID] = v
	case *A_SPEC_RELATION_GROUPSDB:
		db.nextIDA_SPEC_RELATION_GROUPSDB++
		v.ID = db.nextIDA_SPEC_RELATION_GROUPSDB
		db.a_spec_relation_groupsDBs[v.ID] = v
	case *A_SPEC_RELATION_REFDB:
		db.nextIDA_SPEC_RELATION_REFDB++
		v.ID = db.nextIDA_SPEC_RELATION_REFDB
		db.a_spec_relation_refDBs[v.ID] = v
	case *A_SPEC_RELATION_TYPE_REFDB:
		db.nextIDA_SPEC_RELATION_TYPE_REFDB++
		v.ID = db.nextIDA_SPEC_RELATION_TYPE_REFDB
		db.a_spec_relation_type_refDBs[v.ID] = v
	case *A_SPEC_TYPESDB:
		db.nextIDA_SPEC_TYPESDB++
		v.ID = db.nextIDA_SPEC_TYPESDB
		db.a_spec_typesDBs[v.ID] = v
	case *A_THE_HEADERDB:
		db.nextIDA_THE_HEADERDB++
		v.ID = db.nextIDA_THE_HEADERDB
		db.a_the_headerDBs[v.ID] = v
	case *A_TOOL_EXTENSIONSDB:
		db.nextIDA_TOOL_EXTENSIONSDB++
		v.ID = db.nextIDA_TOOL_EXTENSIONSDB
		db.a_tool_extensionsDBs[v.ID] = v
	case *DATATYPE_DEFINITION_BOOLEANDB:
		db.nextIDDATATYPE_DEFINITION_BOOLEANDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_BOOLEANDB
		db.datatype_definition_booleanDBs[v.ID] = v
	case *DATATYPE_DEFINITION_DATEDB:
		db.nextIDDATATYPE_DEFINITION_DATEDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_DATEDB
		db.datatype_definition_dateDBs[v.ID] = v
	case *DATATYPE_DEFINITION_ENUMERATIONDB:
		db.nextIDDATATYPE_DEFINITION_ENUMERATIONDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_ENUMERATIONDB
		db.datatype_definition_enumerationDBs[v.ID] = v
	case *DATATYPE_DEFINITION_INTEGERDB:
		db.nextIDDATATYPE_DEFINITION_INTEGERDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_INTEGERDB
		db.datatype_definition_integerDBs[v.ID] = v
	case *DATATYPE_DEFINITION_REALDB:
		db.nextIDDATATYPE_DEFINITION_REALDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_REALDB
		db.datatype_definition_realDBs[v.ID] = v
	case *DATATYPE_DEFINITION_STRINGDB:
		db.nextIDDATATYPE_DEFINITION_STRINGDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_STRINGDB
		db.datatype_definition_stringDBs[v.ID] = v
	case *DATATYPE_DEFINITION_XHTMLDB:
		db.nextIDDATATYPE_DEFINITION_XHTMLDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_XHTMLDB
		db.datatype_definition_xhtmlDBs[v.ID] = v
	case *EMBEDDED_VALUEDB:
		db.nextIDEMBEDDED_VALUEDB++
		v.ID = db.nextIDEMBEDDED_VALUEDB
		db.embedded_valueDBs[v.ID] = v
	case *ENUM_VALUEDB:
		db.nextIDENUM_VALUEDB++
		v.ID = db.nextIDENUM_VALUEDB
		db.enum_valueDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB
		db.map_attribute_definition_boolean_showinsubjectentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB
		db.map_attribute_definition_boolean_showintableentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB
		db.map_attribute_definition_boolean_showintitleentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB
		db.map_attribute_definition_date_showinsubjectentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB
		db.map_attribute_definition_date_showintableentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB
		db.map_attribute_definition_date_showintitleentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB
		db.map_attribute_definition_enumeration_showinsubjectentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB
		db.map_attribute_definition_enumeration_showintableentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB
		db.map_attribute_definition_enumeration_showintitleentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB
		db.map_attribute_definition_integer_showinsubjectentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB
		db.map_attribute_definition_integer_showintableentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB
		db.map_attribute_definition_integer_showintitleentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB
		db.map_attribute_definition_real_showinsubjectentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB
		db.map_attribute_definition_real_showintableentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB
		db.map_attribute_definition_real_showintitleentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB
		db.map_attribute_definition_string_showinsubjectentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB
		db.map_attribute_definition_string_showintableentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB
		db.map_attribute_definition_string_showintitleentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB
		db.map_attribute_definition_xhtml_showinsubjectentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB
		db.map_attribute_definition_xhtml_showintableentryDBs[v.ID] = v
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB:
		db.nextIDMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB++
		v.ID = db.nextIDMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB
		db.map_attribute_definition_xhtml_showintitleentryDBs[v.ID] = v
	case *Map_SPECIFICATION_Nodes_expandedEntryDB:
		db.nextIDMap_SPECIFICATION_Nodes_expandedEntryDB++
		v.ID = db.nextIDMap_SPECIFICATION_Nodes_expandedEntryDB
		db.map_specification_nodes_expandedentryDBs[v.ID] = v
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB:
		db.nextIDMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB++
		v.ID = db.nextIDMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB
		db.map_spec_object_type_isnodeexpandedentryDBs[v.ID] = v
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB:
		db.nextIDMap_SPEC_OBJECT_TYPE_showIdentifierEntryDB++
		v.ID = db.nextIDMap_SPEC_OBJECT_TYPE_showIdentifierEntryDB
		db.map_spec_object_type_showidentifierentryDBs[v.ID] = v
	case *Map_SPEC_OBJECT_TYPE_showNameEntryDB:
		db.nextIDMap_SPEC_OBJECT_TYPE_showNameEntryDB++
		v.ID = db.nextIDMap_SPEC_OBJECT_TYPE_showNameEntryDB
		db.map_spec_object_type_shownameentryDBs[v.ID] = v
	case *RELATION_GROUPDB:
		db.nextIDRELATION_GROUPDB++
		v.ID = db.nextIDRELATION_GROUPDB
		db.relation_groupDBs[v.ID] = v
	case *RELATION_GROUP_TYPEDB:
		db.nextIDRELATION_GROUP_TYPEDB++
		v.ID = db.nextIDRELATION_GROUP_TYPEDB
		db.relation_group_typeDBs[v.ID] = v
	case *REQ_IFDB:
		db.nextIDREQ_IFDB++
		v.ID = db.nextIDREQ_IFDB
		db.req_ifDBs[v.ID] = v
	case *REQ_IF_CONTENTDB:
		db.nextIDREQ_IF_CONTENTDB++
		v.ID = db.nextIDREQ_IF_CONTENTDB
		db.req_if_contentDBs[v.ID] = v
	case *REQ_IF_HEADERDB:
		db.nextIDREQ_IF_HEADERDB++
		v.ID = db.nextIDREQ_IF_HEADERDB
		db.req_if_headerDBs[v.ID] = v
	case *REQ_IF_TOOL_EXTENSIONDB:
		db.nextIDREQ_IF_TOOL_EXTENSIONDB++
		v.ID = db.nextIDREQ_IF_TOOL_EXTENSIONDB
		db.req_if_tool_extensionDBs[v.ID] = v
	case *RenderingConfigurationDB:
		db.nextIDRenderingConfigurationDB++
		v.ID = db.nextIDRenderingConfigurationDB
		db.renderingconfigurationDBs[v.ID] = v
	case *SPECIFICATIONDB:
		db.nextIDSPECIFICATIONDB++
		v.ID = db.nextIDSPECIFICATIONDB
		db.specificationDBs[v.ID] = v
	case *SPECIFICATION_TYPEDB:
		db.nextIDSPECIFICATION_TYPEDB++
		v.ID = db.nextIDSPECIFICATION_TYPEDB
		db.specification_typeDBs[v.ID] = v
	case *SPEC_HIERARCHYDB:
		db.nextIDSPEC_HIERARCHYDB++
		v.ID = db.nextIDSPEC_HIERARCHYDB
		db.spec_hierarchyDBs[v.ID] = v
	case *SPEC_OBJECTDB:
		db.nextIDSPEC_OBJECTDB++
		v.ID = db.nextIDSPEC_OBJECTDB
		db.spec_objectDBs[v.ID] = v
	case *SPEC_OBJECT_TYPEDB:
		db.nextIDSPEC_OBJECT_TYPEDB++
		v.ID = db.nextIDSPEC_OBJECT_TYPEDB
		db.spec_object_typeDBs[v.ID] = v
	case *SPEC_RELATIONDB:
		db.nextIDSPEC_RELATIONDB++
		v.ID = db.nextIDSPEC_RELATIONDB
		db.spec_relationDBs[v.ID] = v
	case *SPEC_RELATION_TYPEDB:
		db.nextIDSPEC_RELATION_TYPEDB++
		v.ID = db.nextIDSPEC_RELATION_TYPEDB
		db.spec_relation_typeDBs[v.ID] = v
	case *StaticWebSiteDB:
		db.nextIDStaticWebSiteDB++
		v.ID = db.nextIDStaticWebSiteDB
		db.staticwebsiteDBs[v.ID] = v
	case *StaticWebSiteChapterDB:
		db.nextIDStaticWebSiteChapterDB++
		v.ID = db.nextIDStaticWebSiteChapterDB
		db.staticwebsitechapterDBs[v.ID] = v
	case *StaticWebSiteGeneratedImageDB:
		db.nextIDStaticWebSiteGeneratedImageDB++
		v.ID = db.nextIDStaticWebSiteGeneratedImageDB
		db.staticwebsitegeneratedimageDBs[v.ID] = v
	case *StaticWebSiteImageDB:
		db.nextIDStaticWebSiteImageDB++
		v.ID = db.nextIDStaticWebSiteImageDB
		db.staticwebsiteimageDBs[v.ID] = v
	case *StaticWebSiteParagraphDB:
		db.nextIDStaticWebSiteParagraphDB++
		v.ID = db.nextIDStaticWebSiteParagraphDB
		db.staticwebsiteparagraphDBs[v.ID] = v
	case *XHTML_CONTENTDB:
		db.nextIDXHTML_CONTENTDB++
		v.ID = db.nextIDXHTML_CONTENTDB
		db.xhtml_contentDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ALTERNATIVE_IDDB:
		delete(db.alternative_idDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_BOOLEANDB:
		delete(db.attribute_definition_booleanDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_DATEDB:
		delete(db.attribute_definition_dateDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		delete(db.attribute_definition_enumerationDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_INTEGERDB:
		delete(db.attribute_definition_integerDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_REALDB:
		delete(db.attribute_definition_realDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_STRINGDB:
		delete(db.attribute_definition_stringDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_XHTMLDB:
		delete(db.attribute_definition_xhtmlDBs, v.ID)
	case *ATTRIBUTE_VALUE_BOOLEANDB:
		delete(db.attribute_value_booleanDBs, v.ID)
	case *ATTRIBUTE_VALUE_DATEDB:
		delete(db.attribute_value_dateDBs, v.ID)
	case *ATTRIBUTE_VALUE_ENUMERATIONDB:
		delete(db.attribute_value_enumerationDBs, v.ID)
	case *ATTRIBUTE_VALUE_INTEGERDB:
		delete(db.attribute_value_integerDBs, v.ID)
	case *ATTRIBUTE_VALUE_REALDB:
		delete(db.attribute_value_realDBs, v.ID)
	case *ATTRIBUTE_VALUE_STRINGDB:
		delete(db.attribute_value_stringDBs, v.ID)
	case *ATTRIBUTE_VALUE_XHTMLDB:
		delete(db.attribute_value_xhtmlDBs, v.ID)
	case *A_ALTERNATIVE_IDDB:
		delete(db.a_alternative_idDBs, v.ID)
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB:
		delete(db.a_attribute_definition_boolean_refDBs, v.ID)
	case *A_ATTRIBUTE_DEFINITION_DATE_REFDB:
		delete(db.a_attribute_definition_date_refDBs, v.ID)
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB:
		delete(db.a_attribute_definition_enumeration_refDBs, v.ID)
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REFDB:
		delete(db.a_attribute_definition_integer_refDBs, v.ID)
	case *A_ATTRIBUTE_DEFINITION_REAL_REFDB:
		delete(db.a_attribute_definition_real_refDBs, v.ID)
	case *A_ATTRIBUTE_DEFINITION_STRING_REFDB:
		delete(db.a_attribute_definition_string_refDBs, v.ID)
	case *A_ATTRIBUTE_DEFINITION_XHTML_REFDB:
		delete(db.a_attribute_definition_xhtml_refDBs, v.ID)
	case *A_ATTRIBUTE_VALUE_BOOLEANDB:
		delete(db.a_attribute_value_booleanDBs, v.ID)
	case *A_ATTRIBUTE_VALUE_DATEDB:
		delete(db.a_attribute_value_dateDBs, v.ID)
	case *A_ATTRIBUTE_VALUE_ENUMERATIONDB:
		delete(db.a_attribute_value_enumerationDBs, v.ID)
	case *A_ATTRIBUTE_VALUE_INTEGERDB:
		delete(db.a_attribute_value_integerDBs, v.ID)
	case *A_ATTRIBUTE_VALUE_REALDB:
		delete(db.a_attribute_value_realDBs, v.ID)
	case *A_ATTRIBUTE_VALUE_STRINGDB:
		delete(db.a_attribute_value_stringDBs, v.ID)
	case *A_ATTRIBUTE_VALUE_XHTMLDB:
		delete(db.a_attribute_value_xhtmlDBs, v.ID)
	case *A_ATTRIBUTE_VALUE_XHTML_1DB:
		delete(db.a_attribute_value_xhtml_1DBs, v.ID)
	case *A_CHILDRENDB:
		delete(db.a_childrenDBs, v.ID)
	case *A_CORE_CONTENTDB:
		delete(db.a_core_contentDBs, v.ID)
	case *A_DATATYPESDB:
		delete(db.a_datatypesDBs, v.ID)
	case *A_DATATYPE_DEFINITION_BOOLEAN_REFDB:
		delete(db.a_datatype_definition_boolean_refDBs, v.ID)
	case *A_DATATYPE_DEFINITION_DATE_REFDB:
		delete(db.a_datatype_definition_date_refDBs, v.ID)
	case *A_DATATYPE_DEFINITION_ENUMERATION_REFDB:
		delete(db.a_datatype_definition_enumeration_refDBs, v.ID)
	case *A_DATATYPE_DEFINITION_INTEGER_REFDB:
		delete(db.a_datatype_definition_integer_refDBs, v.ID)
	case *A_DATATYPE_DEFINITION_REAL_REFDB:
		delete(db.a_datatype_definition_real_refDBs, v.ID)
	case *A_DATATYPE_DEFINITION_STRING_REFDB:
		delete(db.a_datatype_definition_string_refDBs, v.ID)
	case *A_DATATYPE_DEFINITION_XHTML_REFDB:
		delete(db.a_datatype_definition_xhtml_refDBs, v.ID)
	case *A_EDITABLE_ATTSDB:
		delete(db.a_editable_attsDBs, v.ID)
	case *A_ENUM_VALUE_REFDB:
		delete(db.a_enum_value_refDBs, v.ID)
	case *A_OBJECTDB:
		delete(db.a_objectDBs, v.ID)
	case *A_PROPERTIESDB:
		delete(db.a_propertiesDBs, v.ID)
	case *A_RELATION_GROUP_TYPE_REFDB:
		delete(db.a_relation_group_type_refDBs, v.ID)
	case *A_SOURCE_1DB:
		delete(db.a_source_1DBs, v.ID)
	case *A_SOURCE_SPECIFICATION_1DB:
		delete(db.a_source_specification_1DBs, v.ID)
	case *A_SPECIFICATIONSDB:
		delete(db.a_specificationsDBs, v.ID)
	case *A_SPECIFICATION_TYPE_REFDB:
		delete(db.a_specification_type_refDBs, v.ID)
	case *A_SPECIFIED_VALUESDB:
		delete(db.a_specified_valuesDBs, v.ID)
	case *A_SPEC_ATTRIBUTESDB:
		delete(db.a_spec_attributesDBs, v.ID)
	case *A_SPEC_OBJECTSDB:
		delete(db.a_spec_objectsDBs, v.ID)
	case *A_SPEC_OBJECT_TYPE_REFDB:
		delete(db.a_spec_object_type_refDBs, v.ID)
	case *A_SPEC_RELATIONSDB:
		delete(db.a_spec_relationsDBs, v.ID)
	case *A_SPEC_RELATION_GROUPSDB:
		delete(db.a_spec_relation_groupsDBs, v.ID)
	case *A_SPEC_RELATION_REFDB:
		delete(db.a_spec_relation_refDBs, v.ID)
	case *A_SPEC_RELATION_TYPE_REFDB:
		delete(db.a_spec_relation_type_refDBs, v.ID)
	case *A_SPEC_TYPESDB:
		delete(db.a_spec_typesDBs, v.ID)
	case *A_THE_HEADERDB:
		delete(db.a_the_headerDBs, v.ID)
	case *A_TOOL_EXTENSIONSDB:
		delete(db.a_tool_extensionsDBs, v.ID)
	case *DATATYPE_DEFINITION_BOOLEANDB:
		delete(db.datatype_definition_booleanDBs, v.ID)
	case *DATATYPE_DEFINITION_DATEDB:
		delete(db.datatype_definition_dateDBs, v.ID)
	case *DATATYPE_DEFINITION_ENUMERATIONDB:
		delete(db.datatype_definition_enumerationDBs, v.ID)
	case *DATATYPE_DEFINITION_INTEGERDB:
		delete(db.datatype_definition_integerDBs, v.ID)
	case *DATATYPE_DEFINITION_REALDB:
		delete(db.datatype_definition_realDBs, v.ID)
	case *DATATYPE_DEFINITION_STRINGDB:
		delete(db.datatype_definition_stringDBs, v.ID)
	case *DATATYPE_DEFINITION_XHTMLDB:
		delete(db.datatype_definition_xhtmlDBs, v.ID)
	case *EMBEDDED_VALUEDB:
		delete(db.embedded_valueDBs, v.ID)
	case *ENUM_VALUEDB:
		delete(db.enum_valueDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB:
		delete(db.map_attribute_definition_boolean_showinsubjectentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB:
		delete(db.map_attribute_definition_boolean_showintableentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB:
		delete(db.map_attribute_definition_boolean_showintitleentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB:
		delete(db.map_attribute_definition_date_showinsubjectentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB:
		delete(db.map_attribute_definition_date_showintableentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB:
		delete(db.map_attribute_definition_date_showintitleentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB:
		delete(db.map_attribute_definition_enumeration_showinsubjectentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB:
		delete(db.map_attribute_definition_enumeration_showintableentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB:
		delete(db.map_attribute_definition_enumeration_showintitleentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB:
		delete(db.map_attribute_definition_integer_showinsubjectentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB:
		delete(db.map_attribute_definition_integer_showintableentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB:
		delete(db.map_attribute_definition_integer_showintitleentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB:
		delete(db.map_attribute_definition_real_showinsubjectentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB:
		delete(db.map_attribute_definition_real_showintableentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB:
		delete(db.map_attribute_definition_real_showintitleentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB:
		delete(db.map_attribute_definition_string_showinsubjectentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB:
		delete(db.map_attribute_definition_string_showintableentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB:
		delete(db.map_attribute_definition_string_showintitleentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB:
		delete(db.map_attribute_definition_xhtml_showinsubjectentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB:
		delete(db.map_attribute_definition_xhtml_showintableentryDBs, v.ID)
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB:
		delete(db.map_attribute_definition_xhtml_showintitleentryDBs, v.ID)
	case *Map_SPECIFICATION_Nodes_expandedEntryDB:
		delete(db.map_specification_nodes_expandedentryDBs, v.ID)
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB:
		delete(db.map_spec_object_type_isnodeexpandedentryDBs, v.ID)
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB:
		delete(db.map_spec_object_type_showidentifierentryDBs, v.ID)
	case *Map_SPEC_OBJECT_TYPE_showNameEntryDB:
		delete(db.map_spec_object_type_shownameentryDBs, v.ID)
	case *RELATION_GROUPDB:
		delete(db.relation_groupDBs, v.ID)
	case *RELATION_GROUP_TYPEDB:
		delete(db.relation_group_typeDBs, v.ID)
	case *REQ_IFDB:
		delete(db.req_ifDBs, v.ID)
	case *REQ_IF_CONTENTDB:
		delete(db.req_if_contentDBs, v.ID)
	case *REQ_IF_HEADERDB:
		delete(db.req_if_headerDBs, v.ID)
	case *REQ_IF_TOOL_EXTENSIONDB:
		delete(db.req_if_tool_extensionDBs, v.ID)
	case *RenderingConfigurationDB:
		delete(db.renderingconfigurationDBs, v.ID)
	case *SPECIFICATIONDB:
		delete(db.specificationDBs, v.ID)
	case *SPECIFICATION_TYPEDB:
		delete(db.specification_typeDBs, v.ID)
	case *SPEC_HIERARCHYDB:
		delete(db.spec_hierarchyDBs, v.ID)
	case *SPEC_OBJECTDB:
		delete(db.spec_objectDBs, v.ID)
	case *SPEC_OBJECT_TYPEDB:
		delete(db.spec_object_typeDBs, v.ID)
	case *SPEC_RELATIONDB:
		delete(db.spec_relationDBs, v.ID)
	case *SPEC_RELATION_TYPEDB:
		delete(db.spec_relation_typeDBs, v.ID)
	case *StaticWebSiteDB:
		delete(db.staticwebsiteDBs, v.ID)
	case *StaticWebSiteChapterDB:
		delete(db.staticwebsitechapterDBs, v.ID)
	case *StaticWebSiteGeneratedImageDB:
		delete(db.staticwebsitegeneratedimageDBs, v.ID)
	case *StaticWebSiteImageDB:
		delete(db.staticwebsiteimageDBs, v.ID)
	case *StaticWebSiteParagraphDB:
		delete(db.staticwebsiteparagraphDBs, v.ID)
	case *XHTML_CONTENTDB:
		delete(db.xhtml_contentDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ALTERNATIVE_IDDB:
		db.alternative_idDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_BOOLEANDB:
		db.attribute_definition_booleanDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_DATEDB:
		db.attribute_definition_dateDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		db.attribute_definition_enumerationDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_INTEGERDB:
		db.attribute_definition_integerDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_REALDB:
		db.attribute_definition_realDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_STRINGDB:
		db.attribute_definition_stringDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_XHTMLDB:
		db.attribute_definition_xhtmlDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_BOOLEANDB:
		db.attribute_value_booleanDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_DATEDB:
		db.attribute_value_dateDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_ENUMERATIONDB:
		db.attribute_value_enumerationDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_INTEGERDB:
		db.attribute_value_integerDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_REALDB:
		db.attribute_value_realDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_STRINGDB:
		db.attribute_value_stringDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_XHTMLDB:
		db.attribute_value_xhtmlDBs[v.ID] = v
		return db, nil
	case *A_ALTERNATIVE_IDDB:
		db.a_alternative_idDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB:
		db.a_attribute_definition_boolean_refDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_DEFINITION_DATE_REFDB:
		db.a_attribute_definition_date_refDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB:
		db.a_attribute_definition_enumeration_refDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REFDB:
		db.a_attribute_definition_integer_refDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_DEFINITION_REAL_REFDB:
		db.a_attribute_definition_real_refDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_DEFINITION_STRING_REFDB:
		db.a_attribute_definition_string_refDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_DEFINITION_XHTML_REFDB:
		db.a_attribute_definition_xhtml_refDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_VALUE_BOOLEANDB:
		db.a_attribute_value_booleanDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_VALUE_DATEDB:
		db.a_attribute_value_dateDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_VALUE_ENUMERATIONDB:
		db.a_attribute_value_enumerationDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_VALUE_INTEGERDB:
		db.a_attribute_value_integerDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_VALUE_REALDB:
		db.a_attribute_value_realDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_VALUE_STRINGDB:
		db.a_attribute_value_stringDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_VALUE_XHTMLDB:
		db.a_attribute_value_xhtmlDBs[v.ID] = v
		return db, nil
	case *A_ATTRIBUTE_VALUE_XHTML_1DB:
		db.a_attribute_value_xhtml_1DBs[v.ID] = v
		return db, nil
	case *A_CHILDRENDB:
		db.a_childrenDBs[v.ID] = v
		return db, nil
	case *A_CORE_CONTENTDB:
		db.a_core_contentDBs[v.ID] = v
		return db, nil
	case *A_DATATYPESDB:
		db.a_datatypesDBs[v.ID] = v
		return db, nil
	case *A_DATATYPE_DEFINITION_BOOLEAN_REFDB:
		db.a_datatype_definition_boolean_refDBs[v.ID] = v
		return db, nil
	case *A_DATATYPE_DEFINITION_DATE_REFDB:
		db.a_datatype_definition_date_refDBs[v.ID] = v
		return db, nil
	case *A_DATATYPE_DEFINITION_ENUMERATION_REFDB:
		db.a_datatype_definition_enumeration_refDBs[v.ID] = v
		return db, nil
	case *A_DATATYPE_DEFINITION_INTEGER_REFDB:
		db.a_datatype_definition_integer_refDBs[v.ID] = v
		return db, nil
	case *A_DATATYPE_DEFINITION_REAL_REFDB:
		db.a_datatype_definition_real_refDBs[v.ID] = v
		return db, nil
	case *A_DATATYPE_DEFINITION_STRING_REFDB:
		db.a_datatype_definition_string_refDBs[v.ID] = v
		return db, nil
	case *A_DATATYPE_DEFINITION_XHTML_REFDB:
		db.a_datatype_definition_xhtml_refDBs[v.ID] = v
		return db, nil
	case *A_EDITABLE_ATTSDB:
		db.a_editable_attsDBs[v.ID] = v
		return db, nil
	case *A_ENUM_VALUE_REFDB:
		db.a_enum_value_refDBs[v.ID] = v
		return db, nil
	case *A_OBJECTDB:
		db.a_objectDBs[v.ID] = v
		return db, nil
	case *A_PROPERTIESDB:
		db.a_propertiesDBs[v.ID] = v
		return db, nil
	case *A_RELATION_GROUP_TYPE_REFDB:
		db.a_relation_group_type_refDBs[v.ID] = v
		return db, nil
	case *A_SOURCE_1DB:
		db.a_source_1DBs[v.ID] = v
		return db, nil
	case *A_SOURCE_SPECIFICATION_1DB:
		db.a_source_specification_1DBs[v.ID] = v
		return db, nil
	case *A_SPECIFICATIONSDB:
		db.a_specificationsDBs[v.ID] = v
		return db, nil
	case *A_SPECIFICATION_TYPE_REFDB:
		db.a_specification_type_refDBs[v.ID] = v
		return db, nil
	case *A_SPECIFIED_VALUESDB:
		db.a_specified_valuesDBs[v.ID] = v
		return db, nil
	case *A_SPEC_ATTRIBUTESDB:
		db.a_spec_attributesDBs[v.ID] = v
		return db, nil
	case *A_SPEC_OBJECTSDB:
		db.a_spec_objectsDBs[v.ID] = v
		return db, nil
	case *A_SPEC_OBJECT_TYPE_REFDB:
		db.a_spec_object_type_refDBs[v.ID] = v
		return db, nil
	case *A_SPEC_RELATIONSDB:
		db.a_spec_relationsDBs[v.ID] = v
		return db, nil
	case *A_SPEC_RELATION_GROUPSDB:
		db.a_spec_relation_groupsDBs[v.ID] = v
		return db, nil
	case *A_SPEC_RELATION_REFDB:
		db.a_spec_relation_refDBs[v.ID] = v
		return db, nil
	case *A_SPEC_RELATION_TYPE_REFDB:
		db.a_spec_relation_type_refDBs[v.ID] = v
		return db, nil
	case *A_SPEC_TYPESDB:
		db.a_spec_typesDBs[v.ID] = v
		return db, nil
	case *A_THE_HEADERDB:
		db.a_the_headerDBs[v.ID] = v
		return db, nil
	case *A_TOOL_EXTENSIONSDB:
		db.a_tool_extensionsDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_BOOLEANDB:
		db.datatype_definition_booleanDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_DATEDB:
		db.datatype_definition_dateDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_ENUMERATIONDB:
		db.datatype_definition_enumerationDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_INTEGERDB:
		db.datatype_definition_integerDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_REALDB:
		db.datatype_definition_realDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_STRINGDB:
		db.datatype_definition_stringDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_XHTMLDB:
		db.datatype_definition_xhtmlDBs[v.ID] = v
		return db, nil
	case *EMBEDDED_VALUEDB:
		db.embedded_valueDBs[v.ID] = v
		return db, nil
	case *ENUM_VALUEDB:
		db.enum_valueDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB:
		db.map_attribute_definition_boolean_showinsubjectentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB:
		db.map_attribute_definition_boolean_showintableentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB:
		db.map_attribute_definition_boolean_showintitleentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB:
		db.map_attribute_definition_date_showinsubjectentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB:
		db.map_attribute_definition_date_showintableentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB:
		db.map_attribute_definition_date_showintitleentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB:
		db.map_attribute_definition_enumeration_showinsubjectentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB:
		db.map_attribute_definition_enumeration_showintableentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB:
		db.map_attribute_definition_enumeration_showintitleentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB:
		db.map_attribute_definition_integer_showinsubjectentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB:
		db.map_attribute_definition_integer_showintableentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB:
		db.map_attribute_definition_integer_showintitleentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB:
		db.map_attribute_definition_real_showinsubjectentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB:
		db.map_attribute_definition_real_showintableentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB:
		db.map_attribute_definition_real_showintitleentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB:
		db.map_attribute_definition_string_showinsubjectentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB:
		db.map_attribute_definition_string_showintableentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB:
		db.map_attribute_definition_string_showintitleentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB:
		db.map_attribute_definition_xhtml_showinsubjectentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB:
		db.map_attribute_definition_xhtml_showintableentryDBs[v.ID] = v
		return db, nil
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB:
		db.map_attribute_definition_xhtml_showintitleentryDBs[v.ID] = v
		return db, nil
	case *Map_SPECIFICATION_Nodes_expandedEntryDB:
		db.map_specification_nodes_expandedentryDBs[v.ID] = v
		return db, nil
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB:
		db.map_spec_object_type_isnodeexpandedentryDBs[v.ID] = v
		return db, nil
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB:
		db.map_spec_object_type_showidentifierentryDBs[v.ID] = v
		return db, nil
	case *Map_SPEC_OBJECT_TYPE_showNameEntryDB:
		db.map_spec_object_type_shownameentryDBs[v.ID] = v
		return db, nil
	case *RELATION_GROUPDB:
		db.relation_groupDBs[v.ID] = v
		return db, nil
	case *RELATION_GROUP_TYPEDB:
		db.relation_group_typeDBs[v.ID] = v
		return db, nil
	case *REQ_IFDB:
		db.req_ifDBs[v.ID] = v
		return db, nil
	case *REQ_IF_CONTENTDB:
		db.req_if_contentDBs[v.ID] = v
		return db, nil
	case *REQ_IF_HEADERDB:
		db.req_if_headerDBs[v.ID] = v
		return db, nil
	case *REQ_IF_TOOL_EXTENSIONDB:
		db.req_if_tool_extensionDBs[v.ID] = v
		return db, nil
	case *RenderingConfigurationDB:
		db.renderingconfigurationDBs[v.ID] = v
		return db, nil
	case *SPECIFICATIONDB:
		db.specificationDBs[v.ID] = v
		return db, nil
	case *SPECIFICATION_TYPEDB:
		db.specification_typeDBs[v.ID] = v
		return db, nil
	case *SPEC_HIERARCHYDB:
		db.spec_hierarchyDBs[v.ID] = v
		return db, nil
	case *SPEC_OBJECTDB:
		db.spec_objectDBs[v.ID] = v
		return db, nil
	case *SPEC_OBJECT_TYPEDB:
		db.spec_object_typeDBs[v.ID] = v
		return db, nil
	case *SPEC_RELATIONDB:
		db.spec_relationDBs[v.ID] = v
		return db, nil
	case *SPEC_RELATION_TYPEDB:
		db.spec_relation_typeDBs[v.ID] = v
		return db, nil
	case *StaticWebSiteDB:
		db.staticwebsiteDBs[v.ID] = v
		return db, nil
	case *StaticWebSiteChapterDB:
		db.staticwebsitechapterDBs[v.ID] = v
		return db, nil
	case *StaticWebSiteGeneratedImageDB:
		db.staticwebsitegeneratedimageDBs[v.ID] = v
		return db, nil
	case *StaticWebSiteImageDB:
		db.staticwebsiteimageDBs[v.ID] = v
		return db, nil
	case *StaticWebSiteParagraphDB:
		db.staticwebsiteparagraphDBs[v.ID] = v
		return db, nil
	case *XHTML_CONTENTDB:
		db.xhtml_contentDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ALTERNATIVE_IDDB:
		if existing, ok := db.alternative_idDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ALTERNATIVE_ID github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_BOOLEANDB:
		if existing, ok := db.attribute_definition_booleanDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_BOOLEAN github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_DATEDB:
		if existing, ok := db.attribute_definition_dateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_DATE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		if existing, ok := db.attribute_definition_enumerationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_ENUMERATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_INTEGERDB:
		if existing, ok := db.attribute_definition_integerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_INTEGER github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_REALDB:
		if existing, ok := db.attribute_definition_realDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_REAL github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_STRINGDB:
		if existing, ok := db.attribute_definition_stringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_STRING github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_XHTMLDB:
		if existing, ok := db.attribute_definition_xhtmlDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_XHTML github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_BOOLEANDB:
		if existing, ok := db.attribute_value_booleanDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_BOOLEAN github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_DATEDB:
		if existing, ok := db.attribute_value_dateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_DATE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_ENUMERATIONDB:
		if existing, ok := db.attribute_value_enumerationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_ENUMERATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_INTEGERDB:
		if existing, ok := db.attribute_value_integerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_INTEGER github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_REALDB:
		if existing, ok := db.attribute_value_realDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_REAL github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_STRINGDB:
		if existing, ok := db.attribute_value_stringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_STRING github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_XHTMLDB:
		if existing, ok := db.attribute_value_xhtmlDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_XHTML github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ALTERNATIVE_IDDB:
		if existing, ok := db.a_alternative_idDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ALTERNATIVE_ID github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB:
		if existing, ok := db.a_attribute_definition_boolean_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_DEFINITION_BOOLEAN_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_DEFINITION_DATE_REFDB:
		if existing, ok := db.a_attribute_definition_date_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_DEFINITION_DATE_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB:
		if existing, ok := db.a_attribute_definition_enumeration_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_DEFINITION_ENUMERATION_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REFDB:
		if existing, ok := db.a_attribute_definition_integer_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_DEFINITION_INTEGER_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_DEFINITION_REAL_REFDB:
		if existing, ok := db.a_attribute_definition_real_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_DEFINITION_REAL_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_DEFINITION_STRING_REFDB:
		if existing, ok := db.a_attribute_definition_string_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_DEFINITION_STRING_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REFDB:
		if existing, ok := db.a_attribute_definition_xhtml_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_DEFINITION_XHTML_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_VALUE_BOOLEANDB:
		if existing, ok := db.a_attribute_value_booleanDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_VALUE_BOOLEAN github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_VALUE_DATEDB:
		if existing, ok := db.a_attribute_value_dateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_VALUE_DATE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_VALUE_ENUMERATIONDB:
		if existing, ok := db.a_attribute_value_enumerationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_VALUE_ENUMERATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_VALUE_INTEGERDB:
		if existing, ok := db.a_attribute_value_integerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_VALUE_INTEGER github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_VALUE_REALDB:
		if existing, ok := db.a_attribute_value_realDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_VALUE_REAL github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_VALUE_STRINGDB:
		if existing, ok := db.a_attribute_value_stringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_VALUE_STRING github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_VALUE_XHTMLDB:
		if existing, ok := db.a_attribute_value_xhtmlDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_VALUE_XHTML github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ATTRIBUTE_VALUE_XHTML_1DB:
		if existing, ok := db.a_attribute_value_xhtml_1DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ATTRIBUTE_VALUE_XHTML_1 github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_CHILDRENDB:
		if existing, ok := db.a_childrenDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_CHILDREN github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_CORE_CONTENTDB:
		if existing, ok := db.a_core_contentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_CORE_CONTENT github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_DATATYPESDB:
		if existing, ok := db.a_datatypesDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_DATATYPES github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REFDB:
		if existing, ok := db.a_datatype_definition_boolean_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_DATATYPE_DEFINITION_BOOLEAN_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_DATATYPE_DEFINITION_DATE_REFDB:
		if existing, ok := db.a_datatype_definition_date_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_DATATYPE_DEFINITION_DATE_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REFDB:
		if existing, ok := db.a_datatype_definition_enumeration_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_DATATYPE_DEFINITION_ENUMERATION_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_DATATYPE_DEFINITION_INTEGER_REFDB:
		if existing, ok := db.a_datatype_definition_integer_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_DATATYPE_DEFINITION_INTEGER_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_DATATYPE_DEFINITION_REAL_REFDB:
		if existing, ok := db.a_datatype_definition_real_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_DATATYPE_DEFINITION_REAL_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_DATATYPE_DEFINITION_STRING_REFDB:
		if existing, ok := db.a_datatype_definition_string_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_DATATYPE_DEFINITION_STRING_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_DATATYPE_DEFINITION_XHTML_REFDB:
		if existing, ok := db.a_datatype_definition_xhtml_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_DATATYPE_DEFINITION_XHTML_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_EDITABLE_ATTSDB:
		if existing, ok := db.a_editable_attsDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_EDITABLE_ATTS github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_ENUM_VALUE_REFDB:
		if existing, ok := db.a_enum_value_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_ENUM_VALUE_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_OBJECTDB:
		if existing, ok := db.a_objectDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_OBJECT github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_PROPERTIESDB:
		if existing, ok := db.a_propertiesDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_PROPERTIES github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_RELATION_GROUP_TYPE_REFDB:
		if existing, ok := db.a_relation_group_type_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_RELATION_GROUP_TYPE_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SOURCE_1DB:
		if existing, ok := db.a_source_1DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SOURCE_1 github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SOURCE_SPECIFICATION_1DB:
		if existing, ok := db.a_source_specification_1DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SOURCE_SPECIFICATION_1 github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPECIFICATIONSDB:
		if existing, ok := db.a_specificationsDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPECIFICATIONS github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPECIFICATION_TYPE_REFDB:
		if existing, ok := db.a_specification_type_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPECIFICATION_TYPE_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPECIFIED_VALUESDB:
		if existing, ok := db.a_specified_valuesDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPECIFIED_VALUES github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPEC_ATTRIBUTESDB:
		if existing, ok := db.a_spec_attributesDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPEC_ATTRIBUTES github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPEC_OBJECTSDB:
		if existing, ok := db.a_spec_objectsDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPEC_OBJECTS github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPEC_OBJECT_TYPE_REFDB:
		if existing, ok := db.a_spec_object_type_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPEC_OBJECT_TYPE_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPEC_RELATIONSDB:
		if existing, ok := db.a_spec_relationsDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPEC_RELATIONS github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPEC_RELATION_GROUPSDB:
		if existing, ok := db.a_spec_relation_groupsDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPEC_RELATION_GROUPS github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPEC_RELATION_REFDB:
		if existing, ok := db.a_spec_relation_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPEC_RELATION_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPEC_RELATION_TYPE_REFDB:
		if existing, ok := db.a_spec_relation_type_refDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPEC_RELATION_TYPE_REF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_SPEC_TYPESDB:
		if existing, ok := db.a_spec_typesDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_SPEC_TYPES github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_THE_HEADERDB:
		if existing, ok := db.a_the_headerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_THE_HEADER github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *A_TOOL_EXTENSIONSDB:
		if existing, ok := db.a_tool_extensionsDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db A_TOOL_EXTENSIONS github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_BOOLEANDB:
		if existing, ok := db.datatype_definition_booleanDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_BOOLEAN github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_DATEDB:
		if existing, ok := db.datatype_definition_dateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_DATE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_ENUMERATIONDB:
		if existing, ok := db.datatype_definition_enumerationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_ENUMERATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_INTEGERDB:
		if existing, ok := db.datatype_definition_integerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_INTEGER github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_REALDB:
		if existing, ok := db.datatype_definition_realDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_REAL github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_STRINGDB:
		if existing, ok := db.datatype_definition_stringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_STRING github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_XHTMLDB:
		if existing, ok := db.datatype_definition_xhtmlDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_XHTML github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *EMBEDDED_VALUEDB:
		if existing, ok := db.embedded_valueDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db EMBEDDED_VALUE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ENUM_VALUEDB:
		if existing, ok := db.enum_valueDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ENUM_VALUE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB:
		if existing, ok := db.map_attribute_definition_boolean_showinsubjectentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB:
		if existing, ok := db.map_attribute_definition_boolean_showintableentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB:
		if existing, ok := db.map_attribute_definition_boolean_showintitleentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB:
		if existing, ok := db.map_attribute_definition_date_showinsubjectentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB:
		if existing, ok := db.map_attribute_definition_date_showintableentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB:
		if existing, ok := db.map_attribute_definition_date_showintitleentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB:
		if existing, ok := db.map_attribute_definition_enumeration_showinsubjectentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB:
		if existing, ok := db.map_attribute_definition_enumeration_showintableentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB:
		if existing, ok := db.map_attribute_definition_enumeration_showintitleentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB:
		if existing, ok := db.map_attribute_definition_integer_showinsubjectentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB:
		if existing, ok := db.map_attribute_definition_integer_showintableentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB:
		if existing, ok := db.map_attribute_definition_integer_showintitleentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB:
		if existing, ok := db.map_attribute_definition_real_showinsubjectentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB:
		if existing, ok := db.map_attribute_definition_real_showintableentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB:
		if existing, ok := db.map_attribute_definition_real_showintitleentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB:
		if existing, ok := db.map_attribute_definition_string_showinsubjectentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB:
		if existing, ok := db.map_attribute_definition_string_showintableentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB:
		if existing, ok := db.map_attribute_definition_string_showintitleentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB:
		if existing, ok := db.map_attribute_definition_xhtml_showinsubjectentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB:
		if existing, ok := db.map_attribute_definition_xhtml_showintableentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB:
		if existing, ok := db.map_attribute_definition_xhtml_showintitleentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_SPECIFICATION_Nodes_expandedEntryDB:
		if existing, ok := db.map_specification_nodes_expandedentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_SPECIFICATION_Nodes_expandedEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB:
		if existing, ok := db.map_spec_object_type_isnodeexpandedentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB:
		if existing, ok := db.map_spec_object_type_showidentifierentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_SPEC_OBJECT_TYPE_showIdentifierEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Map_SPEC_OBJECT_TYPE_showNameEntryDB:
		if existing, ok := db.map_spec_object_type_shownameentryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Map_SPEC_OBJECT_TYPE_showNameEntry github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *RELATION_GROUPDB:
		if existing, ok := db.relation_groupDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RELATION_GROUP github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *RELATION_GROUP_TYPEDB:
		if existing, ok := db.relation_group_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RELATION_GROUP_TYPE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *REQ_IFDB:
		if existing, ok := db.req_ifDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db REQ_IF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *REQ_IF_CONTENTDB:
		if existing, ok := db.req_if_contentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db REQ_IF_CONTENT github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *REQ_IF_HEADERDB:
		if existing, ok := db.req_if_headerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db REQ_IF_HEADER github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *REQ_IF_TOOL_EXTENSIONDB:
		if existing, ok := db.req_if_tool_extensionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db REQ_IF_TOOL_EXTENSION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *RenderingConfigurationDB:
		if existing, ok := db.renderingconfigurationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RenderingConfiguration github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPECIFICATIONDB:
		if existing, ok := db.specificationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPECIFICATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPECIFICATION_TYPEDB:
		if existing, ok := db.specification_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPECIFICATION_TYPE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPEC_HIERARCHYDB:
		if existing, ok := db.spec_hierarchyDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPEC_HIERARCHY github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPEC_OBJECTDB:
		if existing, ok := db.spec_objectDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPEC_OBJECT github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPEC_OBJECT_TYPEDB:
		if existing, ok := db.spec_object_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPEC_OBJECT_TYPE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPEC_RELATIONDB:
		if existing, ok := db.spec_relationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPEC_RELATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPEC_RELATION_TYPEDB:
		if existing, ok := db.spec_relation_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPEC_RELATION_TYPE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *StaticWebSiteDB:
		if existing, ok := db.staticwebsiteDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db StaticWebSite github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *StaticWebSiteChapterDB:
		if existing, ok := db.staticwebsitechapterDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db StaticWebSiteChapter github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *StaticWebSiteGeneratedImageDB:
		if existing, ok := db.staticwebsitegeneratedimageDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db StaticWebSiteGeneratedImage github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *StaticWebSiteImageDB:
		if existing, ok := db.staticwebsiteimageDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db StaticWebSiteImage github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *StaticWebSiteParagraphDB:
		if existing, ok := db.staticwebsiteparagraphDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db StaticWebSiteParagraph github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *XHTML_CONTENTDB:
		if existing, ok := db.xhtml_contentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db XHTML_CONTENT github.com/fullstack-lang/gongreqif/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]ALTERNATIVE_IDDB:
		*ptr = make([]ALTERNATIVE_IDDB, 0, len(db.alternative_idDBs))
		for _, v := range db.alternative_idDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_BOOLEANDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_BOOLEANDB, 0, len(db.attribute_definition_booleanDBs))
		for _, v := range db.attribute_definition_booleanDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_DATEDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_DATEDB, 0, len(db.attribute_definition_dateDBs))
		for _, v := range db.attribute_definition_dateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_ENUMERATIONDB, 0, len(db.attribute_definition_enumerationDBs))
		for _, v := range db.attribute_definition_enumerationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_INTEGERDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_INTEGERDB, 0, len(db.attribute_definition_integerDBs))
		for _, v := range db.attribute_definition_integerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_REALDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_REALDB, 0, len(db.attribute_definition_realDBs))
		for _, v := range db.attribute_definition_realDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_STRINGDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_STRINGDB, 0, len(db.attribute_definition_stringDBs))
		for _, v := range db.attribute_definition_stringDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_XHTMLDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_XHTMLDB, 0, len(db.attribute_definition_xhtmlDBs))
		for _, v := range db.attribute_definition_xhtmlDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_BOOLEANDB:
		*ptr = make([]ATTRIBUTE_VALUE_BOOLEANDB, 0, len(db.attribute_value_booleanDBs))
		for _, v := range db.attribute_value_booleanDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_DATEDB:
		*ptr = make([]ATTRIBUTE_VALUE_DATEDB, 0, len(db.attribute_value_dateDBs))
		for _, v := range db.attribute_value_dateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_ENUMERATIONDB:
		*ptr = make([]ATTRIBUTE_VALUE_ENUMERATIONDB, 0, len(db.attribute_value_enumerationDBs))
		for _, v := range db.attribute_value_enumerationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_INTEGERDB:
		*ptr = make([]ATTRIBUTE_VALUE_INTEGERDB, 0, len(db.attribute_value_integerDBs))
		for _, v := range db.attribute_value_integerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_REALDB:
		*ptr = make([]ATTRIBUTE_VALUE_REALDB, 0, len(db.attribute_value_realDBs))
		for _, v := range db.attribute_value_realDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_STRINGDB:
		*ptr = make([]ATTRIBUTE_VALUE_STRINGDB, 0, len(db.attribute_value_stringDBs))
		for _, v := range db.attribute_value_stringDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_XHTMLDB:
		*ptr = make([]ATTRIBUTE_VALUE_XHTMLDB, 0, len(db.attribute_value_xhtmlDBs))
		for _, v := range db.attribute_value_xhtmlDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ALTERNATIVE_IDDB:
		*ptr = make([]A_ALTERNATIVE_IDDB, 0, len(db.a_alternative_idDBs))
		for _, v := range db.a_alternative_idDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB:
		*ptr = make([]A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB, 0, len(db.a_attribute_definition_boolean_refDBs))
		for _, v := range db.a_attribute_definition_boolean_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_DEFINITION_DATE_REFDB:
		*ptr = make([]A_ATTRIBUTE_DEFINITION_DATE_REFDB, 0, len(db.a_attribute_definition_date_refDBs))
		for _, v := range db.a_attribute_definition_date_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB:
		*ptr = make([]A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB, 0, len(db.a_attribute_definition_enumeration_refDBs))
		for _, v := range db.a_attribute_definition_enumeration_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_DEFINITION_INTEGER_REFDB:
		*ptr = make([]A_ATTRIBUTE_DEFINITION_INTEGER_REFDB, 0, len(db.a_attribute_definition_integer_refDBs))
		for _, v := range db.a_attribute_definition_integer_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_DEFINITION_REAL_REFDB:
		*ptr = make([]A_ATTRIBUTE_DEFINITION_REAL_REFDB, 0, len(db.a_attribute_definition_real_refDBs))
		for _, v := range db.a_attribute_definition_real_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_DEFINITION_STRING_REFDB:
		*ptr = make([]A_ATTRIBUTE_DEFINITION_STRING_REFDB, 0, len(db.a_attribute_definition_string_refDBs))
		for _, v := range db.a_attribute_definition_string_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_DEFINITION_XHTML_REFDB:
		*ptr = make([]A_ATTRIBUTE_DEFINITION_XHTML_REFDB, 0, len(db.a_attribute_definition_xhtml_refDBs))
		for _, v := range db.a_attribute_definition_xhtml_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_VALUE_BOOLEANDB:
		*ptr = make([]A_ATTRIBUTE_VALUE_BOOLEANDB, 0, len(db.a_attribute_value_booleanDBs))
		for _, v := range db.a_attribute_value_booleanDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_VALUE_DATEDB:
		*ptr = make([]A_ATTRIBUTE_VALUE_DATEDB, 0, len(db.a_attribute_value_dateDBs))
		for _, v := range db.a_attribute_value_dateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_VALUE_ENUMERATIONDB:
		*ptr = make([]A_ATTRIBUTE_VALUE_ENUMERATIONDB, 0, len(db.a_attribute_value_enumerationDBs))
		for _, v := range db.a_attribute_value_enumerationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_VALUE_INTEGERDB:
		*ptr = make([]A_ATTRIBUTE_VALUE_INTEGERDB, 0, len(db.a_attribute_value_integerDBs))
		for _, v := range db.a_attribute_value_integerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_VALUE_REALDB:
		*ptr = make([]A_ATTRIBUTE_VALUE_REALDB, 0, len(db.a_attribute_value_realDBs))
		for _, v := range db.a_attribute_value_realDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_VALUE_STRINGDB:
		*ptr = make([]A_ATTRIBUTE_VALUE_STRINGDB, 0, len(db.a_attribute_value_stringDBs))
		for _, v := range db.a_attribute_value_stringDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_VALUE_XHTMLDB:
		*ptr = make([]A_ATTRIBUTE_VALUE_XHTMLDB, 0, len(db.a_attribute_value_xhtmlDBs))
		for _, v := range db.a_attribute_value_xhtmlDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ATTRIBUTE_VALUE_XHTML_1DB:
		*ptr = make([]A_ATTRIBUTE_VALUE_XHTML_1DB, 0, len(db.a_attribute_value_xhtml_1DBs))
		for _, v := range db.a_attribute_value_xhtml_1DBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_CHILDRENDB:
		*ptr = make([]A_CHILDRENDB, 0, len(db.a_childrenDBs))
		for _, v := range db.a_childrenDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_CORE_CONTENTDB:
		*ptr = make([]A_CORE_CONTENTDB, 0, len(db.a_core_contentDBs))
		for _, v := range db.a_core_contentDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_DATATYPESDB:
		*ptr = make([]A_DATATYPESDB, 0, len(db.a_datatypesDBs))
		for _, v := range db.a_datatypesDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_DATATYPE_DEFINITION_BOOLEAN_REFDB:
		*ptr = make([]A_DATATYPE_DEFINITION_BOOLEAN_REFDB, 0, len(db.a_datatype_definition_boolean_refDBs))
		for _, v := range db.a_datatype_definition_boolean_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_DATATYPE_DEFINITION_DATE_REFDB:
		*ptr = make([]A_DATATYPE_DEFINITION_DATE_REFDB, 0, len(db.a_datatype_definition_date_refDBs))
		for _, v := range db.a_datatype_definition_date_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_DATATYPE_DEFINITION_ENUMERATION_REFDB:
		*ptr = make([]A_DATATYPE_DEFINITION_ENUMERATION_REFDB, 0, len(db.a_datatype_definition_enumeration_refDBs))
		for _, v := range db.a_datatype_definition_enumeration_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_DATATYPE_DEFINITION_INTEGER_REFDB:
		*ptr = make([]A_DATATYPE_DEFINITION_INTEGER_REFDB, 0, len(db.a_datatype_definition_integer_refDBs))
		for _, v := range db.a_datatype_definition_integer_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_DATATYPE_DEFINITION_REAL_REFDB:
		*ptr = make([]A_DATATYPE_DEFINITION_REAL_REFDB, 0, len(db.a_datatype_definition_real_refDBs))
		for _, v := range db.a_datatype_definition_real_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_DATATYPE_DEFINITION_STRING_REFDB:
		*ptr = make([]A_DATATYPE_DEFINITION_STRING_REFDB, 0, len(db.a_datatype_definition_string_refDBs))
		for _, v := range db.a_datatype_definition_string_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_DATATYPE_DEFINITION_XHTML_REFDB:
		*ptr = make([]A_DATATYPE_DEFINITION_XHTML_REFDB, 0, len(db.a_datatype_definition_xhtml_refDBs))
		for _, v := range db.a_datatype_definition_xhtml_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_EDITABLE_ATTSDB:
		*ptr = make([]A_EDITABLE_ATTSDB, 0, len(db.a_editable_attsDBs))
		for _, v := range db.a_editable_attsDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_ENUM_VALUE_REFDB:
		*ptr = make([]A_ENUM_VALUE_REFDB, 0, len(db.a_enum_value_refDBs))
		for _, v := range db.a_enum_value_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_OBJECTDB:
		*ptr = make([]A_OBJECTDB, 0, len(db.a_objectDBs))
		for _, v := range db.a_objectDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_PROPERTIESDB:
		*ptr = make([]A_PROPERTIESDB, 0, len(db.a_propertiesDBs))
		for _, v := range db.a_propertiesDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_RELATION_GROUP_TYPE_REFDB:
		*ptr = make([]A_RELATION_GROUP_TYPE_REFDB, 0, len(db.a_relation_group_type_refDBs))
		for _, v := range db.a_relation_group_type_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SOURCE_1DB:
		*ptr = make([]A_SOURCE_1DB, 0, len(db.a_source_1DBs))
		for _, v := range db.a_source_1DBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SOURCE_SPECIFICATION_1DB:
		*ptr = make([]A_SOURCE_SPECIFICATION_1DB, 0, len(db.a_source_specification_1DBs))
		for _, v := range db.a_source_specification_1DBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPECIFICATIONSDB:
		*ptr = make([]A_SPECIFICATIONSDB, 0, len(db.a_specificationsDBs))
		for _, v := range db.a_specificationsDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPECIFICATION_TYPE_REFDB:
		*ptr = make([]A_SPECIFICATION_TYPE_REFDB, 0, len(db.a_specification_type_refDBs))
		for _, v := range db.a_specification_type_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPECIFIED_VALUESDB:
		*ptr = make([]A_SPECIFIED_VALUESDB, 0, len(db.a_specified_valuesDBs))
		for _, v := range db.a_specified_valuesDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPEC_ATTRIBUTESDB:
		*ptr = make([]A_SPEC_ATTRIBUTESDB, 0, len(db.a_spec_attributesDBs))
		for _, v := range db.a_spec_attributesDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPEC_OBJECTSDB:
		*ptr = make([]A_SPEC_OBJECTSDB, 0, len(db.a_spec_objectsDBs))
		for _, v := range db.a_spec_objectsDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPEC_OBJECT_TYPE_REFDB:
		*ptr = make([]A_SPEC_OBJECT_TYPE_REFDB, 0, len(db.a_spec_object_type_refDBs))
		for _, v := range db.a_spec_object_type_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPEC_RELATIONSDB:
		*ptr = make([]A_SPEC_RELATIONSDB, 0, len(db.a_spec_relationsDBs))
		for _, v := range db.a_spec_relationsDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPEC_RELATION_GROUPSDB:
		*ptr = make([]A_SPEC_RELATION_GROUPSDB, 0, len(db.a_spec_relation_groupsDBs))
		for _, v := range db.a_spec_relation_groupsDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPEC_RELATION_REFDB:
		*ptr = make([]A_SPEC_RELATION_REFDB, 0, len(db.a_spec_relation_refDBs))
		for _, v := range db.a_spec_relation_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPEC_RELATION_TYPE_REFDB:
		*ptr = make([]A_SPEC_RELATION_TYPE_REFDB, 0, len(db.a_spec_relation_type_refDBs))
		for _, v := range db.a_spec_relation_type_refDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_SPEC_TYPESDB:
		*ptr = make([]A_SPEC_TYPESDB, 0, len(db.a_spec_typesDBs))
		for _, v := range db.a_spec_typesDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_THE_HEADERDB:
		*ptr = make([]A_THE_HEADERDB, 0, len(db.a_the_headerDBs))
		for _, v := range db.a_the_headerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]A_TOOL_EXTENSIONSDB:
		*ptr = make([]A_TOOL_EXTENSIONSDB, 0, len(db.a_tool_extensionsDBs))
		for _, v := range db.a_tool_extensionsDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_BOOLEANDB:
		*ptr = make([]DATATYPE_DEFINITION_BOOLEANDB, 0, len(db.datatype_definition_booleanDBs))
		for _, v := range db.datatype_definition_booleanDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_DATEDB:
		*ptr = make([]DATATYPE_DEFINITION_DATEDB, 0, len(db.datatype_definition_dateDBs))
		for _, v := range db.datatype_definition_dateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_ENUMERATIONDB:
		*ptr = make([]DATATYPE_DEFINITION_ENUMERATIONDB, 0, len(db.datatype_definition_enumerationDBs))
		for _, v := range db.datatype_definition_enumerationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_INTEGERDB:
		*ptr = make([]DATATYPE_DEFINITION_INTEGERDB, 0, len(db.datatype_definition_integerDBs))
		for _, v := range db.datatype_definition_integerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_REALDB:
		*ptr = make([]DATATYPE_DEFINITION_REALDB, 0, len(db.datatype_definition_realDBs))
		for _, v := range db.datatype_definition_realDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_STRINGDB:
		*ptr = make([]DATATYPE_DEFINITION_STRINGDB, 0, len(db.datatype_definition_stringDBs))
		for _, v := range db.datatype_definition_stringDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_XHTMLDB:
		*ptr = make([]DATATYPE_DEFINITION_XHTMLDB, 0, len(db.datatype_definition_xhtmlDBs))
		for _, v := range db.datatype_definition_xhtmlDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]EMBEDDED_VALUEDB:
		*ptr = make([]EMBEDDED_VALUEDB, 0, len(db.embedded_valueDBs))
		for _, v := range db.embedded_valueDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ENUM_VALUEDB:
		*ptr = make([]ENUM_VALUEDB, 0, len(db.enum_valueDBs))
		for _, v := range db.enum_valueDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB, 0, len(db.map_attribute_definition_boolean_showinsubjectentryDBs))
		for _, v := range db.map_attribute_definition_boolean_showinsubjectentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB, 0, len(db.map_attribute_definition_boolean_showintableentryDBs))
		for _, v := range db.map_attribute_definition_boolean_showintableentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB, 0, len(db.map_attribute_definition_boolean_showintitleentryDBs))
		for _, v := range db.map_attribute_definition_boolean_showintitleentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB, 0, len(db.map_attribute_definition_date_showinsubjectentryDBs))
		for _, v := range db.map_attribute_definition_date_showinsubjectentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB, 0, len(db.map_attribute_definition_date_showintableentryDBs))
		for _, v := range db.map_attribute_definition_date_showintableentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB, 0, len(db.map_attribute_definition_date_showintitleentryDBs))
		for _, v := range db.map_attribute_definition_date_showintitleentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB, 0, len(db.map_attribute_definition_enumeration_showinsubjectentryDBs))
		for _, v := range db.map_attribute_definition_enumeration_showinsubjectentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB, 0, len(db.map_attribute_definition_enumeration_showintableentryDBs))
		for _, v := range db.map_attribute_definition_enumeration_showintableentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB, 0, len(db.map_attribute_definition_enumeration_showintitleentryDBs))
		for _, v := range db.map_attribute_definition_enumeration_showintitleentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB, 0, len(db.map_attribute_definition_integer_showinsubjectentryDBs))
		for _, v := range db.map_attribute_definition_integer_showinsubjectentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB, 0, len(db.map_attribute_definition_integer_showintableentryDBs))
		for _, v := range db.map_attribute_definition_integer_showintableentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB, 0, len(db.map_attribute_definition_integer_showintitleentryDBs))
		for _, v := range db.map_attribute_definition_integer_showintitleentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB, 0, len(db.map_attribute_definition_real_showinsubjectentryDBs))
		for _, v := range db.map_attribute_definition_real_showinsubjectentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB, 0, len(db.map_attribute_definition_real_showintableentryDBs))
		for _, v := range db.map_attribute_definition_real_showintableentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB, 0, len(db.map_attribute_definition_real_showintitleentryDBs))
		for _, v := range db.map_attribute_definition_real_showintitleentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB, 0, len(db.map_attribute_definition_string_showinsubjectentryDBs))
		for _, v := range db.map_attribute_definition_string_showinsubjectentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB, 0, len(db.map_attribute_definition_string_showintableentryDBs))
		for _, v := range db.map_attribute_definition_string_showintableentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB, 0, len(db.map_attribute_definition_string_showintitleentryDBs))
		for _, v := range db.map_attribute_definition_string_showintitleentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB, 0, len(db.map_attribute_definition_xhtml_showinsubjectentryDBs))
		for _, v := range db.map_attribute_definition_xhtml_showinsubjectentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB, 0, len(db.map_attribute_definition_xhtml_showintableentryDBs))
		for _, v := range db.map_attribute_definition_xhtml_showintableentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB:
		*ptr = make([]Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB, 0, len(db.map_attribute_definition_xhtml_showintitleentryDBs))
		for _, v := range db.map_attribute_definition_xhtml_showintitleentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_SPECIFICATION_Nodes_expandedEntryDB:
		*ptr = make([]Map_SPECIFICATION_Nodes_expandedEntryDB, 0, len(db.map_specification_nodes_expandedentryDBs))
		for _, v := range db.map_specification_nodes_expandedentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB:
		*ptr = make([]Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB, 0, len(db.map_spec_object_type_isnodeexpandedentryDBs))
		for _, v := range db.map_spec_object_type_isnodeexpandedentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB:
		*ptr = make([]Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB, 0, len(db.map_spec_object_type_showidentifierentryDBs))
		for _, v := range db.map_spec_object_type_showidentifierentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Map_SPEC_OBJECT_TYPE_showNameEntryDB:
		*ptr = make([]Map_SPEC_OBJECT_TYPE_showNameEntryDB, 0, len(db.map_spec_object_type_shownameentryDBs))
		for _, v := range db.map_spec_object_type_shownameentryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RELATION_GROUPDB:
		*ptr = make([]RELATION_GROUPDB, 0, len(db.relation_groupDBs))
		for _, v := range db.relation_groupDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RELATION_GROUP_TYPEDB:
		*ptr = make([]RELATION_GROUP_TYPEDB, 0, len(db.relation_group_typeDBs))
		for _, v := range db.relation_group_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]REQ_IFDB:
		*ptr = make([]REQ_IFDB, 0, len(db.req_ifDBs))
		for _, v := range db.req_ifDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]REQ_IF_CONTENTDB:
		*ptr = make([]REQ_IF_CONTENTDB, 0, len(db.req_if_contentDBs))
		for _, v := range db.req_if_contentDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]REQ_IF_HEADERDB:
		*ptr = make([]REQ_IF_HEADERDB, 0, len(db.req_if_headerDBs))
		for _, v := range db.req_if_headerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]REQ_IF_TOOL_EXTENSIONDB:
		*ptr = make([]REQ_IF_TOOL_EXTENSIONDB, 0, len(db.req_if_tool_extensionDBs))
		for _, v := range db.req_if_tool_extensionDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RenderingConfigurationDB:
		*ptr = make([]RenderingConfigurationDB, 0, len(db.renderingconfigurationDBs))
		for _, v := range db.renderingconfigurationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPECIFICATIONDB:
		*ptr = make([]SPECIFICATIONDB, 0, len(db.specificationDBs))
		for _, v := range db.specificationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPECIFICATION_TYPEDB:
		*ptr = make([]SPECIFICATION_TYPEDB, 0, len(db.specification_typeDBs))
		for _, v := range db.specification_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPEC_HIERARCHYDB:
		*ptr = make([]SPEC_HIERARCHYDB, 0, len(db.spec_hierarchyDBs))
		for _, v := range db.spec_hierarchyDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPEC_OBJECTDB:
		*ptr = make([]SPEC_OBJECTDB, 0, len(db.spec_objectDBs))
		for _, v := range db.spec_objectDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPEC_OBJECT_TYPEDB:
		*ptr = make([]SPEC_OBJECT_TYPEDB, 0, len(db.spec_object_typeDBs))
		for _, v := range db.spec_object_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPEC_RELATIONDB:
		*ptr = make([]SPEC_RELATIONDB, 0, len(db.spec_relationDBs))
		for _, v := range db.spec_relationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPEC_RELATION_TYPEDB:
		*ptr = make([]SPEC_RELATION_TYPEDB, 0, len(db.spec_relation_typeDBs))
		for _, v := range db.spec_relation_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]StaticWebSiteDB:
		*ptr = make([]StaticWebSiteDB, 0, len(db.staticwebsiteDBs))
		for _, v := range db.staticwebsiteDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]StaticWebSiteChapterDB:
		*ptr = make([]StaticWebSiteChapterDB, 0, len(db.staticwebsitechapterDBs))
		for _, v := range db.staticwebsitechapterDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]StaticWebSiteGeneratedImageDB:
		*ptr = make([]StaticWebSiteGeneratedImageDB, 0, len(db.staticwebsitegeneratedimageDBs))
		for _, v := range db.staticwebsitegeneratedimageDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]StaticWebSiteImageDB:
		*ptr = make([]StaticWebSiteImageDB, 0, len(db.staticwebsiteimageDBs))
		for _, v := range db.staticwebsiteimageDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]StaticWebSiteParagraphDB:
		*ptr = make([]StaticWebSiteParagraphDB, 0, len(db.staticwebsiteparagraphDBs))
		for _, v := range db.staticwebsiteparagraphDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]XHTML_CONTENTDB:
		*ptr = make([]XHTML_CONTENTDB, 0, len(db.xhtml_contentDBs))
		for _, v := range db.xhtml_contentDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gongreqif/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *ALTERNATIVE_IDDB:
		tmp, ok := db.alternative_idDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ALTERNATIVE_ID Unkown entry %d", i))
		}

		alternative_idDB, _ := instanceDB.(*ALTERNATIVE_IDDB)
		*alternative_idDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_BOOLEANDB:
		tmp, ok := db.attribute_definition_booleanDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_BOOLEAN Unkown entry %d", i))
		}

		attribute_definition_booleanDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_BOOLEANDB)
		*attribute_definition_booleanDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_DATEDB:
		tmp, ok := db.attribute_definition_dateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_DATE Unkown entry %d", i))
		}

		attribute_definition_dateDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_DATEDB)
		*attribute_definition_dateDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		tmp, ok := db.attribute_definition_enumerationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_ENUMERATION Unkown entry %d", i))
		}

		attribute_definition_enumerationDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_ENUMERATIONDB)
		*attribute_definition_enumerationDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_INTEGERDB:
		tmp, ok := db.attribute_definition_integerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_INTEGER Unkown entry %d", i))
		}

		attribute_definition_integerDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_INTEGERDB)
		*attribute_definition_integerDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_REALDB:
		tmp, ok := db.attribute_definition_realDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_REAL Unkown entry %d", i))
		}

		attribute_definition_realDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_REALDB)
		*attribute_definition_realDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_STRINGDB:
		tmp, ok := db.attribute_definition_stringDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_STRING Unkown entry %d", i))
		}

		attribute_definition_stringDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_STRINGDB)
		*attribute_definition_stringDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_XHTMLDB:
		tmp, ok := db.attribute_definition_xhtmlDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_XHTML Unkown entry %d", i))
		}

		attribute_definition_xhtmlDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_XHTMLDB)
		*attribute_definition_xhtmlDB = *tmp
		
	case *ATTRIBUTE_VALUE_BOOLEANDB:
		tmp, ok := db.attribute_value_booleanDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_BOOLEAN Unkown entry %d", i))
		}

		attribute_value_booleanDB, _ := instanceDB.(*ATTRIBUTE_VALUE_BOOLEANDB)
		*attribute_value_booleanDB = *tmp
		
	case *ATTRIBUTE_VALUE_DATEDB:
		tmp, ok := db.attribute_value_dateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_DATE Unkown entry %d", i))
		}

		attribute_value_dateDB, _ := instanceDB.(*ATTRIBUTE_VALUE_DATEDB)
		*attribute_value_dateDB = *tmp
		
	case *ATTRIBUTE_VALUE_ENUMERATIONDB:
		tmp, ok := db.attribute_value_enumerationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_ENUMERATION Unkown entry %d", i))
		}

		attribute_value_enumerationDB, _ := instanceDB.(*ATTRIBUTE_VALUE_ENUMERATIONDB)
		*attribute_value_enumerationDB = *tmp
		
	case *ATTRIBUTE_VALUE_INTEGERDB:
		tmp, ok := db.attribute_value_integerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_INTEGER Unkown entry %d", i))
		}

		attribute_value_integerDB, _ := instanceDB.(*ATTRIBUTE_VALUE_INTEGERDB)
		*attribute_value_integerDB = *tmp
		
	case *ATTRIBUTE_VALUE_REALDB:
		tmp, ok := db.attribute_value_realDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_REAL Unkown entry %d", i))
		}

		attribute_value_realDB, _ := instanceDB.(*ATTRIBUTE_VALUE_REALDB)
		*attribute_value_realDB = *tmp
		
	case *ATTRIBUTE_VALUE_STRINGDB:
		tmp, ok := db.attribute_value_stringDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_STRING Unkown entry %d", i))
		}

		attribute_value_stringDB, _ := instanceDB.(*ATTRIBUTE_VALUE_STRINGDB)
		*attribute_value_stringDB = *tmp
		
	case *ATTRIBUTE_VALUE_XHTMLDB:
		tmp, ok := db.attribute_value_xhtmlDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_XHTML Unkown entry %d", i))
		}

		attribute_value_xhtmlDB, _ := instanceDB.(*ATTRIBUTE_VALUE_XHTMLDB)
		*attribute_value_xhtmlDB = *tmp
		
	case *A_ALTERNATIVE_IDDB:
		tmp, ok := db.a_alternative_idDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ALTERNATIVE_ID Unkown entry %d", i))
		}

		a_alternative_idDB, _ := instanceDB.(*A_ALTERNATIVE_IDDB)
		*a_alternative_idDB = *tmp
		
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB:
		tmp, ok := db.a_attribute_definition_boolean_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_DEFINITION_BOOLEAN_REF Unkown entry %d", i))
		}

		a_attribute_definition_boolean_refDB, _ := instanceDB.(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB)
		*a_attribute_definition_boolean_refDB = *tmp
		
	case *A_ATTRIBUTE_DEFINITION_DATE_REFDB:
		tmp, ok := db.a_attribute_definition_date_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_DEFINITION_DATE_REF Unkown entry %d", i))
		}

		a_attribute_definition_date_refDB, _ := instanceDB.(*A_ATTRIBUTE_DEFINITION_DATE_REFDB)
		*a_attribute_definition_date_refDB = *tmp
		
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB:
		tmp, ok := db.a_attribute_definition_enumeration_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_DEFINITION_ENUMERATION_REF Unkown entry %d", i))
		}

		a_attribute_definition_enumeration_refDB, _ := instanceDB.(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB)
		*a_attribute_definition_enumeration_refDB = *tmp
		
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REFDB:
		tmp, ok := db.a_attribute_definition_integer_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_DEFINITION_INTEGER_REF Unkown entry %d", i))
		}

		a_attribute_definition_integer_refDB, _ := instanceDB.(*A_ATTRIBUTE_DEFINITION_INTEGER_REFDB)
		*a_attribute_definition_integer_refDB = *tmp
		
	case *A_ATTRIBUTE_DEFINITION_REAL_REFDB:
		tmp, ok := db.a_attribute_definition_real_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_DEFINITION_REAL_REF Unkown entry %d", i))
		}

		a_attribute_definition_real_refDB, _ := instanceDB.(*A_ATTRIBUTE_DEFINITION_REAL_REFDB)
		*a_attribute_definition_real_refDB = *tmp
		
	case *A_ATTRIBUTE_DEFINITION_STRING_REFDB:
		tmp, ok := db.a_attribute_definition_string_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_DEFINITION_STRING_REF Unkown entry %d", i))
		}

		a_attribute_definition_string_refDB, _ := instanceDB.(*A_ATTRIBUTE_DEFINITION_STRING_REFDB)
		*a_attribute_definition_string_refDB = *tmp
		
	case *A_ATTRIBUTE_DEFINITION_XHTML_REFDB:
		tmp, ok := db.a_attribute_definition_xhtml_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_DEFINITION_XHTML_REF Unkown entry %d", i))
		}

		a_attribute_definition_xhtml_refDB, _ := instanceDB.(*A_ATTRIBUTE_DEFINITION_XHTML_REFDB)
		*a_attribute_definition_xhtml_refDB = *tmp
		
	case *A_ATTRIBUTE_VALUE_BOOLEANDB:
		tmp, ok := db.a_attribute_value_booleanDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_VALUE_BOOLEAN Unkown entry %d", i))
		}

		a_attribute_value_booleanDB, _ := instanceDB.(*A_ATTRIBUTE_VALUE_BOOLEANDB)
		*a_attribute_value_booleanDB = *tmp
		
	case *A_ATTRIBUTE_VALUE_DATEDB:
		tmp, ok := db.a_attribute_value_dateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_VALUE_DATE Unkown entry %d", i))
		}

		a_attribute_value_dateDB, _ := instanceDB.(*A_ATTRIBUTE_VALUE_DATEDB)
		*a_attribute_value_dateDB = *tmp
		
	case *A_ATTRIBUTE_VALUE_ENUMERATIONDB:
		tmp, ok := db.a_attribute_value_enumerationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_VALUE_ENUMERATION Unkown entry %d", i))
		}

		a_attribute_value_enumerationDB, _ := instanceDB.(*A_ATTRIBUTE_VALUE_ENUMERATIONDB)
		*a_attribute_value_enumerationDB = *tmp
		
	case *A_ATTRIBUTE_VALUE_INTEGERDB:
		tmp, ok := db.a_attribute_value_integerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_VALUE_INTEGER Unkown entry %d", i))
		}

		a_attribute_value_integerDB, _ := instanceDB.(*A_ATTRIBUTE_VALUE_INTEGERDB)
		*a_attribute_value_integerDB = *tmp
		
	case *A_ATTRIBUTE_VALUE_REALDB:
		tmp, ok := db.a_attribute_value_realDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_VALUE_REAL Unkown entry %d", i))
		}

		a_attribute_value_realDB, _ := instanceDB.(*A_ATTRIBUTE_VALUE_REALDB)
		*a_attribute_value_realDB = *tmp
		
	case *A_ATTRIBUTE_VALUE_STRINGDB:
		tmp, ok := db.a_attribute_value_stringDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_VALUE_STRING Unkown entry %d", i))
		}

		a_attribute_value_stringDB, _ := instanceDB.(*A_ATTRIBUTE_VALUE_STRINGDB)
		*a_attribute_value_stringDB = *tmp
		
	case *A_ATTRIBUTE_VALUE_XHTMLDB:
		tmp, ok := db.a_attribute_value_xhtmlDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_VALUE_XHTML Unkown entry %d", i))
		}

		a_attribute_value_xhtmlDB, _ := instanceDB.(*A_ATTRIBUTE_VALUE_XHTMLDB)
		*a_attribute_value_xhtmlDB = *tmp
		
	case *A_ATTRIBUTE_VALUE_XHTML_1DB:
		tmp, ok := db.a_attribute_value_xhtml_1DBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ATTRIBUTE_VALUE_XHTML_1 Unkown entry %d", i))
		}

		a_attribute_value_xhtml_1DB, _ := instanceDB.(*A_ATTRIBUTE_VALUE_XHTML_1DB)
		*a_attribute_value_xhtml_1DB = *tmp
		
	case *A_CHILDRENDB:
		tmp, ok := db.a_childrenDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_CHILDREN Unkown entry %d", i))
		}

		a_childrenDB, _ := instanceDB.(*A_CHILDRENDB)
		*a_childrenDB = *tmp
		
	case *A_CORE_CONTENTDB:
		tmp, ok := db.a_core_contentDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_CORE_CONTENT Unkown entry %d", i))
		}

		a_core_contentDB, _ := instanceDB.(*A_CORE_CONTENTDB)
		*a_core_contentDB = *tmp
		
	case *A_DATATYPESDB:
		tmp, ok := db.a_datatypesDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_DATATYPES Unkown entry %d", i))
		}

		a_datatypesDB, _ := instanceDB.(*A_DATATYPESDB)
		*a_datatypesDB = *tmp
		
	case *A_DATATYPE_DEFINITION_BOOLEAN_REFDB:
		tmp, ok := db.a_datatype_definition_boolean_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_DATATYPE_DEFINITION_BOOLEAN_REF Unkown entry %d", i))
		}

		a_datatype_definition_boolean_refDB, _ := instanceDB.(*A_DATATYPE_DEFINITION_BOOLEAN_REFDB)
		*a_datatype_definition_boolean_refDB = *tmp
		
	case *A_DATATYPE_DEFINITION_DATE_REFDB:
		tmp, ok := db.a_datatype_definition_date_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_DATATYPE_DEFINITION_DATE_REF Unkown entry %d", i))
		}

		a_datatype_definition_date_refDB, _ := instanceDB.(*A_DATATYPE_DEFINITION_DATE_REFDB)
		*a_datatype_definition_date_refDB = *tmp
		
	case *A_DATATYPE_DEFINITION_ENUMERATION_REFDB:
		tmp, ok := db.a_datatype_definition_enumeration_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_DATATYPE_DEFINITION_ENUMERATION_REF Unkown entry %d", i))
		}

		a_datatype_definition_enumeration_refDB, _ := instanceDB.(*A_DATATYPE_DEFINITION_ENUMERATION_REFDB)
		*a_datatype_definition_enumeration_refDB = *tmp
		
	case *A_DATATYPE_DEFINITION_INTEGER_REFDB:
		tmp, ok := db.a_datatype_definition_integer_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_DATATYPE_DEFINITION_INTEGER_REF Unkown entry %d", i))
		}

		a_datatype_definition_integer_refDB, _ := instanceDB.(*A_DATATYPE_DEFINITION_INTEGER_REFDB)
		*a_datatype_definition_integer_refDB = *tmp
		
	case *A_DATATYPE_DEFINITION_REAL_REFDB:
		tmp, ok := db.a_datatype_definition_real_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_DATATYPE_DEFINITION_REAL_REF Unkown entry %d", i))
		}

		a_datatype_definition_real_refDB, _ := instanceDB.(*A_DATATYPE_DEFINITION_REAL_REFDB)
		*a_datatype_definition_real_refDB = *tmp
		
	case *A_DATATYPE_DEFINITION_STRING_REFDB:
		tmp, ok := db.a_datatype_definition_string_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_DATATYPE_DEFINITION_STRING_REF Unkown entry %d", i))
		}

		a_datatype_definition_string_refDB, _ := instanceDB.(*A_DATATYPE_DEFINITION_STRING_REFDB)
		*a_datatype_definition_string_refDB = *tmp
		
	case *A_DATATYPE_DEFINITION_XHTML_REFDB:
		tmp, ok := db.a_datatype_definition_xhtml_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_DATATYPE_DEFINITION_XHTML_REF Unkown entry %d", i))
		}

		a_datatype_definition_xhtml_refDB, _ := instanceDB.(*A_DATATYPE_DEFINITION_XHTML_REFDB)
		*a_datatype_definition_xhtml_refDB = *tmp
		
	case *A_EDITABLE_ATTSDB:
		tmp, ok := db.a_editable_attsDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_EDITABLE_ATTS Unkown entry %d", i))
		}

		a_editable_attsDB, _ := instanceDB.(*A_EDITABLE_ATTSDB)
		*a_editable_attsDB = *tmp
		
	case *A_ENUM_VALUE_REFDB:
		tmp, ok := db.a_enum_value_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_ENUM_VALUE_REF Unkown entry %d", i))
		}

		a_enum_value_refDB, _ := instanceDB.(*A_ENUM_VALUE_REFDB)
		*a_enum_value_refDB = *tmp
		
	case *A_OBJECTDB:
		tmp, ok := db.a_objectDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_OBJECT Unkown entry %d", i))
		}

		a_objectDB, _ := instanceDB.(*A_OBJECTDB)
		*a_objectDB = *tmp
		
	case *A_PROPERTIESDB:
		tmp, ok := db.a_propertiesDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_PROPERTIES Unkown entry %d", i))
		}

		a_propertiesDB, _ := instanceDB.(*A_PROPERTIESDB)
		*a_propertiesDB = *tmp
		
	case *A_RELATION_GROUP_TYPE_REFDB:
		tmp, ok := db.a_relation_group_type_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_RELATION_GROUP_TYPE_REF Unkown entry %d", i))
		}

		a_relation_group_type_refDB, _ := instanceDB.(*A_RELATION_GROUP_TYPE_REFDB)
		*a_relation_group_type_refDB = *tmp
		
	case *A_SOURCE_1DB:
		tmp, ok := db.a_source_1DBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SOURCE_1 Unkown entry %d", i))
		}

		a_source_1DB, _ := instanceDB.(*A_SOURCE_1DB)
		*a_source_1DB = *tmp
		
	case *A_SOURCE_SPECIFICATION_1DB:
		tmp, ok := db.a_source_specification_1DBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SOURCE_SPECIFICATION_1 Unkown entry %d", i))
		}

		a_source_specification_1DB, _ := instanceDB.(*A_SOURCE_SPECIFICATION_1DB)
		*a_source_specification_1DB = *tmp
		
	case *A_SPECIFICATIONSDB:
		tmp, ok := db.a_specificationsDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPECIFICATIONS Unkown entry %d", i))
		}

		a_specificationsDB, _ := instanceDB.(*A_SPECIFICATIONSDB)
		*a_specificationsDB = *tmp
		
	case *A_SPECIFICATION_TYPE_REFDB:
		tmp, ok := db.a_specification_type_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPECIFICATION_TYPE_REF Unkown entry %d", i))
		}

		a_specification_type_refDB, _ := instanceDB.(*A_SPECIFICATION_TYPE_REFDB)
		*a_specification_type_refDB = *tmp
		
	case *A_SPECIFIED_VALUESDB:
		tmp, ok := db.a_specified_valuesDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPECIFIED_VALUES Unkown entry %d", i))
		}

		a_specified_valuesDB, _ := instanceDB.(*A_SPECIFIED_VALUESDB)
		*a_specified_valuesDB = *tmp
		
	case *A_SPEC_ATTRIBUTESDB:
		tmp, ok := db.a_spec_attributesDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPEC_ATTRIBUTES Unkown entry %d", i))
		}

		a_spec_attributesDB, _ := instanceDB.(*A_SPEC_ATTRIBUTESDB)
		*a_spec_attributesDB = *tmp
		
	case *A_SPEC_OBJECTSDB:
		tmp, ok := db.a_spec_objectsDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPEC_OBJECTS Unkown entry %d", i))
		}

		a_spec_objectsDB, _ := instanceDB.(*A_SPEC_OBJECTSDB)
		*a_spec_objectsDB = *tmp
		
	case *A_SPEC_OBJECT_TYPE_REFDB:
		tmp, ok := db.a_spec_object_type_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPEC_OBJECT_TYPE_REF Unkown entry %d", i))
		}

		a_spec_object_type_refDB, _ := instanceDB.(*A_SPEC_OBJECT_TYPE_REFDB)
		*a_spec_object_type_refDB = *tmp
		
	case *A_SPEC_RELATIONSDB:
		tmp, ok := db.a_spec_relationsDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPEC_RELATIONS Unkown entry %d", i))
		}

		a_spec_relationsDB, _ := instanceDB.(*A_SPEC_RELATIONSDB)
		*a_spec_relationsDB = *tmp
		
	case *A_SPEC_RELATION_GROUPSDB:
		tmp, ok := db.a_spec_relation_groupsDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPEC_RELATION_GROUPS Unkown entry %d", i))
		}

		a_spec_relation_groupsDB, _ := instanceDB.(*A_SPEC_RELATION_GROUPSDB)
		*a_spec_relation_groupsDB = *tmp
		
	case *A_SPEC_RELATION_REFDB:
		tmp, ok := db.a_spec_relation_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPEC_RELATION_REF Unkown entry %d", i))
		}

		a_spec_relation_refDB, _ := instanceDB.(*A_SPEC_RELATION_REFDB)
		*a_spec_relation_refDB = *tmp
		
	case *A_SPEC_RELATION_TYPE_REFDB:
		tmp, ok := db.a_spec_relation_type_refDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPEC_RELATION_TYPE_REF Unkown entry %d", i))
		}

		a_spec_relation_type_refDB, _ := instanceDB.(*A_SPEC_RELATION_TYPE_REFDB)
		*a_spec_relation_type_refDB = *tmp
		
	case *A_SPEC_TYPESDB:
		tmp, ok := db.a_spec_typesDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_SPEC_TYPES Unkown entry %d", i))
		}

		a_spec_typesDB, _ := instanceDB.(*A_SPEC_TYPESDB)
		*a_spec_typesDB = *tmp
		
	case *A_THE_HEADERDB:
		tmp, ok := db.a_the_headerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_THE_HEADER Unkown entry %d", i))
		}

		a_the_headerDB, _ := instanceDB.(*A_THE_HEADERDB)
		*a_the_headerDB = *tmp
		
	case *A_TOOL_EXTENSIONSDB:
		tmp, ok := db.a_tool_extensionsDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First A_TOOL_EXTENSIONS Unkown entry %d", i))
		}

		a_tool_extensionsDB, _ := instanceDB.(*A_TOOL_EXTENSIONSDB)
		*a_tool_extensionsDB = *tmp
		
	case *DATATYPE_DEFINITION_BOOLEANDB:
		tmp, ok := db.datatype_definition_booleanDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_BOOLEAN Unkown entry %d", i))
		}

		datatype_definition_booleanDB, _ := instanceDB.(*DATATYPE_DEFINITION_BOOLEANDB)
		*datatype_definition_booleanDB = *tmp
		
	case *DATATYPE_DEFINITION_DATEDB:
		tmp, ok := db.datatype_definition_dateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_DATE Unkown entry %d", i))
		}

		datatype_definition_dateDB, _ := instanceDB.(*DATATYPE_DEFINITION_DATEDB)
		*datatype_definition_dateDB = *tmp
		
	case *DATATYPE_DEFINITION_ENUMERATIONDB:
		tmp, ok := db.datatype_definition_enumerationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_ENUMERATION Unkown entry %d", i))
		}

		datatype_definition_enumerationDB, _ := instanceDB.(*DATATYPE_DEFINITION_ENUMERATIONDB)
		*datatype_definition_enumerationDB = *tmp
		
	case *DATATYPE_DEFINITION_INTEGERDB:
		tmp, ok := db.datatype_definition_integerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_INTEGER Unkown entry %d", i))
		}

		datatype_definition_integerDB, _ := instanceDB.(*DATATYPE_DEFINITION_INTEGERDB)
		*datatype_definition_integerDB = *tmp
		
	case *DATATYPE_DEFINITION_REALDB:
		tmp, ok := db.datatype_definition_realDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_REAL Unkown entry %d", i))
		}

		datatype_definition_realDB, _ := instanceDB.(*DATATYPE_DEFINITION_REALDB)
		*datatype_definition_realDB = *tmp
		
	case *DATATYPE_DEFINITION_STRINGDB:
		tmp, ok := db.datatype_definition_stringDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_STRING Unkown entry %d", i))
		}

		datatype_definition_stringDB, _ := instanceDB.(*DATATYPE_DEFINITION_STRINGDB)
		*datatype_definition_stringDB = *tmp
		
	case *DATATYPE_DEFINITION_XHTMLDB:
		tmp, ok := db.datatype_definition_xhtmlDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_XHTML Unkown entry %d", i))
		}

		datatype_definition_xhtmlDB, _ := instanceDB.(*DATATYPE_DEFINITION_XHTMLDB)
		*datatype_definition_xhtmlDB = *tmp
		
	case *EMBEDDED_VALUEDB:
		tmp, ok := db.embedded_valueDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First EMBEDDED_VALUE Unkown entry %d", i))
		}

		embedded_valueDB, _ := instanceDB.(*EMBEDDED_VALUEDB)
		*embedded_valueDB = *tmp
		
	case *ENUM_VALUEDB:
		tmp, ok := db.enum_valueDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ENUM_VALUE Unkown entry %d", i))
		}

		enum_valueDB, _ := instanceDB.(*ENUM_VALUEDB)
		*enum_valueDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB:
		tmp, ok := db.map_attribute_definition_boolean_showinsubjectentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry Unkown entry %d", i))
		}

		map_attribute_definition_boolean_showinsubjectentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB)
		*map_attribute_definition_boolean_showinsubjectentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB:
		tmp, ok := db.map_attribute_definition_boolean_showintableentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry Unkown entry %d", i))
		}

		map_attribute_definition_boolean_showintableentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB)
		*map_attribute_definition_boolean_showintableentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB:
		tmp, ok := db.map_attribute_definition_boolean_showintitleentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry Unkown entry %d", i))
		}

		map_attribute_definition_boolean_showintitleentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB)
		*map_attribute_definition_boolean_showintitleentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB:
		tmp, ok := db.map_attribute_definition_date_showinsubjectentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry Unkown entry %d", i))
		}

		map_attribute_definition_date_showinsubjectentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB)
		*map_attribute_definition_date_showinsubjectentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB:
		tmp, ok := db.map_attribute_definition_date_showintableentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry Unkown entry %d", i))
		}

		map_attribute_definition_date_showintableentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB)
		*map_attribute_definition_date_showintableentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB:
		tmp, ok := db.map_attribute_definition_date_showintitleentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry Unkown entry %d", i))
		}

		map_attribute_definition_date_showintitleentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB)
		*map_attribute_definition_date_showintitleentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB:
		tmp, ok := db.map_attribute_definition_enumeration_showinsubjectentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry Unkown entry %d", i))
		}

		map_attribute_definition_enumeration_showinsubjectentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB)
		*map_attribute_definition_enumeration_showinsubjectentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB:
		tmp, ok := db.map_attribute_definition_enumeration_showintableentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry Unkown entry %d", i))
		}

		map_attribute_definition_enumeration_showintableentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB)
		*map_attribute_definition_enumeration_showintableentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB:
		tmp, ok := db.map_attribute_definition_enumeration_showintitleentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry Unkown entry %d", i))
		}

		map_attribute_definition_enumeration_showintitleentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB)
		*map_attribute_definition_enumeration_showintitleentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB:
		tmp, ok := db.map_attribute_definition_integer_showinsubjectentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry Unkown entry %d", i))
		}

		map_attribute_definition_integer_showinsubjectentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB)
		*map_attribute_definition_integer_showinsubjectentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB:
		tmp, ok := db.map_attribute_definition_integer_showintableentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry Unkown entry %d", i))
		}

		map_attribute_definition_integer_showintableentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB)
		*map_attribute_definition_integer_showintableentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB:
		tmp, ok := db.map_attribute_definition_integer_showintitleentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry Unkown entry %d", i))
		}

		map_attribute_definition_integer_showintitleentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB)
		*map_attribute_definition_integer_showintitleentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB:
		tmp, ok := db.map_attribute_definition_real_showinsubjectentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry Unkown entry %d", i))
		}

		map_attribute_definition_real_showinsubjectentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB)
		*map_attribute_definition_real_showinsubjectentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB:
		tmp, ok := db.map_attribute_definition_real_showintableentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry Unkown entry %d", i))
		}

		map_attribute_definition_real_showintableentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB)
		*map_attribute_definition_real_showintableentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB:
		tmp, ok := db.map_attribute_definition_real_showintitleentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry Unkown entry %d", i))
		}

		map_attribute_definition_real_showintitleentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB)
		*map_attribute_definition_real_showintitleentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB:
		tmp, ok := db.map_attribute_definition_string_showinsubjectentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry Unkown entry %d", i))
		}

		map_attribute_definition_string_showinsubjectentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB)
		*map_attribute_definition_string_showinsubjectentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB:
		tmp, ok := db.map_attribute_definition_string_showintableentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry Unkown entry %d", i))
		}

		map_attribute_definition_string_showintableentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB)
		*map_attribute_definition_string_showintableentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB:
		tmp, ok := db.map_attribute_definition_string_showintitleentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry Unkown entry %d", i))
		}

		map_attribute_definition_string_showintitleentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB)
		*map_attribute_definition_string_showintitleentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB:
		tmp, ok := db.map_attribute_definition_xhtml_showinsubjectentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry Unkown entry %d", i))
		}

		map_attribute_definition_xhtml_showinsubjectentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB)
		*map_attribute_definition_xhtml_showinsubjectentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB:
		tmp, ok := db.map_attribute_definition_xhtml_showintableentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry Unkown entry %d", i))
		}

		map_attribute_definition_xhtml_showintableentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB)
		*map_attribute_definition_xhtml_showintableentryDB = *tmp
		
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB:
		tmp, ok := db.map_attribute_definition_xhtml_showintitleentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry Unkown entry %d", i))
		}

		map_attribute_definition_xhtml_showintitleentryDB, _ := instanceDB.(*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB)
		*map_attribute_definition_xhtml_showintitleentryDB = *tmp
		
	case *Map_SPECIFICATION_Nodes_expandedEntryDB:
		tmp, ok := db.map_specification_nodes_expandedentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_SPECIFICATION_Nodes_expandedEntry Unkown entry %d", i))
		}

		map_specification_nodes_expandedentryDB, _ := instanceDB.(*Map_SPECIFICATION_Nodes_expandedEntryDB)
		*map_specification_nodes_expandedentryDB = *tmp
		
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB:
		tmp, ok := db.map_spec_object_type_isnodeexpandedentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry Unkown entry %d", i))
		}

		map_spec_object_type_isnodeexpandedentryDB, _ := instanceDB.(*Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB)
		*map_spec_object_type_isnodeexpandedentryDB = *tmp
		
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB:
		tmp, ok := db.map_spec_object_type_showidentifierentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_SPEC_OBJECT_TYPE_showIdentifierEntry Unkown entry %d", i))
		}

		map_spec_object_type_showidentifierentryDB, _ := instanceDB.(*Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB)
		*map_spec_object_type_showidentifierentryDB = *tmp
		
	case *Map_SPEC_OBJECT_TYPE_showNameEntryDB:
		tmp, ok := db.map_spec_object_type_shownameentryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Map_SPEC_OBJECT_TYPE_showNameEntry Unkown entry %d", i))
		}

		map_spec_object_type_shownameentryDB, _ := instanceDB.(*Map_SPEC_OBJECT_TYPE_showNameEntryDB)
		*map_spec_object_type_shownameentryDB = *tmp
		
	case *RELATION_GROUPDB:
		tmp, ok := db.relation_groupDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RELATION_GROUP Unkown entry %d", i))
		}

		relation_groupDB, _ := instanceDB.(*RELATION_GROUPDB)
		*relation_groupDB = *tmp
		
	case *RELATION_GROUP_TYPEDB:
		tmp, ok := db.relation_group_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RELATION_GROUP_TYPE Unkown entry %d", i))
		}

		relation_group_typeDB, _ := instanceDB.(*RELATION_GROUP_TYPEDB)
		*relation_group_typeDB = *tmp
		
	case *REQ_IFDB:
		tmp, ok := db.req_ifDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First REQ_IF Unkown entry %d", i))
		}

		req_ifDB, _ := instanceDB.(*REQ_IFDB)
		*req_ifDB = *tmp
		
	case *REQ_IF_CONTENTDB:
		tmp, ok := db.req_if_contentDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First REQ_IF_CONTENT Unkown entry %d", i))
		}

		req_if_contentDB, _ := instanceDB.(*REQ_IF_CONTENTDB)
		*req_if_contentDB = *tmp
		
	case *REQ_IF_HEADERDB:
		tmp, ok := db.req_if_headerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First REQ_IF_HEADER Unkown entry %d", i))
		}

		req_if_headerDB, _ := instanceDB.(*REQ_IF_HEADERDB)
		*req_if_headerDB = *tmp
		
	case *REQ_IF_TOOL_EXTENSIONDB:
		tmp, ok := db.req_if_tool_extensionDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First REQ_IF_TOOL_EXTENSION Unkown entry %d", i))
		}

		req_if_tool_extensionDB, _ := instanceDB.(*REQ_IF_TOOL_EXTENSIONDB)
		*req_if_tool_extensionDB = *tmp
		
	case *RenderingConfigurationDB:
		tmp, ok := db.renderingconfigurationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RenderingConfiguration Unkown entry %d", i))
		}

		renderingconfigurationDB, _ := instanceDB.(*RenderingConfigurationDB)
		*renderingconfigurationDB = *tmp
		
	case *SPECIFICATIONDB:
		tmp, ok := db.specificationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPECIFICATION Unkown entry %d", i))
		}

		specificationDB, _ := instanceDB.(*SPECIFICATIONDB)
		*specificationDB = *tmp
		
	case *SPECIFICATION_TYPEDB:
		tmp, ok := db.specification_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPECIFICATION_TYPE Unkown entry %d", i))
		}

		specification_typeDB, _ := instanceDB.(*SPECIFICATION_TYPEDB)
		*specification_typeDB = *tmp
		
	case *SPEC_HIERARCHYDB:
		tmp, ok := db.spec_hierarchyDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPEC_HIERARCHY Unkown entry %d", i))
		}

		spec_hierarchyDB, _ := instanceDB.(*SPEC_HIERARCHYDB)
		*spec_hierarchyDB = *tmp
		
	case *SPEC_OBJECTDB:
		tmp, ok := db.spec_objectDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPEC_OBJECT Unkown entry %d", i))
		}

		spec_objectDB, _ := instanceDB.(*SPEC_OBJECTDB)
		*spec_objectDB = *tmp
		
	case *SPEC_OBJECT_TYPEDB:
		tmp, ok := db.spec_object_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPEC_OBJECT_TYPE Unkown entry %d", i))
		}

		spec_object_typeDB, _ := instanceDB.(*SPEC_OBJECT_TYPEDB)
		*spec_object_typeDB = *tmp
		
	case *SPEC_RELATIONDB:
		tmp, ok := db.spec_relationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPEC_RELATION Unkown entry %d", i))
		}

		spec_relationDB, _ := instanceDB.(*SPEC_RELATIONDB)
		*spec_relationDB = *tmp
		
	case *SPEC_RELATION_TYPEDB:
		tmp, ok := db.spec_relation_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPEC_RELATION_TYPE Unkown entry %d", i))
		}

		spec_relation_typeDB, _ := instanceDB.(*SPEC_RELATION_TYPEDB)
		*spec_relation_typeDB = *tmp
		
	case *StaticWebSiteDB:
		tmp, ok := db.staticwebsiteDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First StaticWebSite Unkown entry %d", i))
		}

		staticwebsiteDB, _ := instanceDB.(*StaticWebSiteDB)
		*staticwebsiteDB = *tmp
		
	case *StaticWebSiteChapterDB:
		tmp, ok := db.staticwebsitechapterDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First StaticWebSiteChapter Unkown entry %d", i))
		}

		staticwebsitechapterDB, _ := instanceDB.(*StaticWebSiteChapterDB)
		*staticwebsitechapterDB = *tmp
		
	case *StaticWebSiteGeneratedImageDB:
		tmp, ok := db.staticwebsitegeneratedimageDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First StaticWebSiteGeneratedImage Unkown entry %d", i))
		}

		staticwebsitegeneratedimageDB, _ := instanceDB.(*StaticWebSiteGeneratedImageDB)
		*staticwebsitegeneratedimageDB = *tmp
		
	case *StaticWebSiteImageDB:
		tmp, ok := db.staticwebsiteimageDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First StaticWebSiteImage Unkown entry %d", i))
		}

		staticwebsiteimageDB, _ := instanceDB.(*StaticWebSiteImageDB)
		*staticwebsiteimageDB = *tmp
		
	case *StaticWebSiteParagraphDB:
		tmp, ok := db.staticwebsiteparagraphDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First StaticWebSiteParagraph Unkown entry %d", i))
		}

		staticwebsiteparagraphDB, _ := instanceDB.(*StaticWebSiteParagraphDB)
		*staticwebsiteparagraphDB = *tmp
		
	case *XHTML_CONTENTDB:
		tmp, ok := db.xhtml_contentDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First XHTML_CONTENT Unkown entry %d", i))
		}

		xhtml_contentDB, _ := instanceDB.(*XHTML_CONTENTDB)
		*xhtml_contentDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, Unkown type")
	}
	
	return db, nil
}

