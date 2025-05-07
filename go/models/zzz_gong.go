// generated code - do not edit
package models

import (
	"cmp"
	"embed"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"time"

	gongreqif_go "github.com/fullstack-lang/gongreqif/go"
)

// can be used for
//     days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var _ = __Gong__Abs

const ProbeTreeSidebarSuffix = "-sidebar"
const ProbeTableSuffix = "-table"
const ProbeFormSuffix = "-form"
const ProbeSplitSuffix = "-probe"

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// Stage enables storage of staged instances
// swagger:ignore
type Stage struct {
	name string

	// insertion point for definition of arrays registering instances
	ALTERNATIVE_IDs           map[*ALTERNATIVE_ID]any
	ALTERNATIVE_IDs_mapString map[string]*ALTERNATIVE_ID

	// insertion point for slice of pointers maps
	OnAfterALTERNATIVE_IDCreateCallback OnAfterCreateInterface[ALTERNATIVE_ID]
	OnAfterALTERNATIVE_IDUpdateCallback OnAfterUpdateInterface[ALTERNATIVE_ID]
	OnAfterALTERNATIVE_IDDeleteCallback OnAfterDeleteInterface[ALTERNATIVE_ID]
	OnAfterALTERNATIVE_IDReadCallback   OnAfterReadInterface[ALTERNATIVE_ID]

	ATTRIBUTE_DEFINITION_BOOLEANs           map[*ATTRIBUTE_DEFINITION_BOOLEAN]any
	ATTRIBUTE_DEFINITION_BOOLEANs_mapString map[string]*ATTRIBUTE_DEFINITION_BOOLEAN

	// insertion point for slice of pointers maps
	ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_BOOLEAN

	ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN_reverseMap map[*ATTRIBUTE_VALUE_BOOLEAN]*ATTRIBUTE_DEFINITION_BOOLEAN

	OnAfterATTRIBUTE_DEFINITION_BOOLEANCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_BOOLEAN]
	OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_BOOLEAN]
	OnAfterATTRIBUTE_DEFINITION_BOOLEANDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_BOOLEAN]
	OnAfterATTRIBUTE_DEFINITION_BOOLEANReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_BOOLEAN]

	ATTRIBUTE_DEFINITION_DATEs           map[*ATTRIBUTE_DEFINITION_DATE]any
	ATTRIBUTE_DEFINITION_DATEs_mapString map[string]*ATTRIBUTE_DEFINITION_DATE

	// insertion point for slice of pointers maps
	ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_DATE

	ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE_reverseMap map[*ATTRIBUTE_VALUE_DATE]*ATTRIBUTE_DEFINITION_DATE

	OnAfterATTRIBUTE_DEFINITION_DATECreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_DATE]
	OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_DATE]
	OnAfterATTRIBUTE_DEFINITION_DATEDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_DATE]
	OnAfterATTRIBUTE_DEFINITION_DATEReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_DATE]

	ATTRIBUTE_DEFINITION_ENUMERATIONs           map[*ATTRIBUTE_DEFINITION_ENUMERATION]any
	ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString map[string]*ATTRIBUTE_DEFINITION_ENUMERATION

	// insertion point for slice of pointers maps
	ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION_reverseMap map[*ATTRIBUTE_VALUE_ENUMERATION]*ATTRIBUTE_DEFINITION_ENUMERATION

	ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_ENUMERATION

	OnAfterATTRIBUTE_DEFINITION_ENUMERATIONCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_ENUMERATION]
	OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_ENUMERATION]
	OnAfterATTRIBUTE_DEFINITION_ENUMERATIONDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_ENUMERATION]
	OnAfterATTRIBUTE_DEFINITION_ENUMERATIONReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_ENUMERATION]

	ATTRIBUTE_DEFINITION_INTEGERs           map[*ATTRIBUTE_DEFINITION_INTEGER]any
	ATTRIBUTE_DEFINITION_INTEGERs_mapString map[string]*ATTRIBUTE_DEFINITION_INTEGER

	// insertion point for slice of pointers maps
	ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_INTEGER

	ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER_reverseMap map[*ATTRIBUTE_VALUE_INTEGER]*ATTRIBUTE_DEFINITION_INTEGER

	OnAfterATTRIBUTE_DEFINITION_INTEGERCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_INTEGER]
	OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_INTEGER]
	OnAfterATTRIBUTE_DEFINITION_INTEGERDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_INTEGER]
	OnAfterATTRIBUTE_DEFINITION_INTEGERReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_INTEGER]

	ATTRIBUTE_DEFINITION_REALs           map[*ATTRIBUTE_DEFINITION_REAL]any
	ATTRIBUTE_DEFINITION_REALs_mapString map[string]*ATTRIBUTE_DEFINITION_REAL

	// insertion point for slice of pointers maps
	ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_REAL

	ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL_reverseMap map[*ATTRIBUTE_VALUE_REAL]*ATTRIBUTE_DEFINITION_REAL

	OnAfterATTRIBUTE_DEFINITION_REALCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_REAL]
	OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_REAL]
	OnAfterATTRIBUTE_DEFINITION_REALDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_REAL]
	OnAfterATTRIBUTE_DEFINITION_REALReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_REAL]

	ATTRIBUTE_DEFINITION_STRINGs           map[*ATTRIBUTE_DEFINITION_STRING]any
	ATTRIBUTE_DEFINITION_STRINGs_mapString map[string]*ATTRIBUTE_DEFINITION_STRING

	// insertion point for slice of pointers maps
	ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_STRING

	ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING_reverseMap map[*ATTRIBUTE_VALUE_STRING]*ATTRIBUTE_DEFINITION_STRING

	OnAfterATTRIBUTE_DEFINITION_STRINGCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_STRING]
	OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_STRING]
	OnAfterATTRIBUTE_DEFINITION_STRINGDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_STRING]
	OnAfterATTRIBUTE_DEFINITION_STRINGReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_STRING]

	ATTRIBUTE_DEFINITION_XHTMLs           map[*ATTRIBUTE_DEFINITION_XHTML]any
	ATTRIBUTE_DEFINITION_XHTMLs_mapString map[string]*ATTRIBUTE_DEFINITION_XHTML

	// insertion point for slice of pointers maps
	ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_XHTML

	ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML_reverseMap map[*ATTRIBUTE_VALUE_XHTML]*ATTRIBUTE_DEFINITION_XHTML

	OnAfterATTRIBUTE_DEFINITION_XHTMLCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_XHTML]
	OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_XHTML]
	OnAfterATTRIBUTE_DEFINITION_XHTMLDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_XHTML]
	OnAfterATTRIBUTE_DEFINITION_XHTMLReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_XHTML]

	ATTRIBUTE_VALUE_BOOLEANs           map[*ATTRIBUTE_VALUE_BOOLEAN]any
	ATTRIBUTE_VALUE_BOOLEANs_mapString map[string]*ATTRIBUTE_VALUE_BOOLEAN

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_BOOLEANCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_BOOLEAN]
	OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_BOOLEAN]
	OnAfterATTRIBUTE_VALUE_BOOLEANDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_BOOLEAN]
	OnAfterATTRIBUTE_VALUE_BOOLEANReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_BOOLEAN]

	ATTRIBUTE_VALUE_DATEs           map[*ATTRIBUTE_VALUE_DATE]any
	ATTRIBUTE_VALUE_DATEs_mapString map[string]*ATTRIBUTE_VALUE_DATE

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_DATECreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_DATE]
	OnAfterATTRIBUTE_VALUE_DATEUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_DATE]
	OnAfterATTRIBUTE_VALUE_DATEDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_DATE]
	OnAfterATTRIBUTE_VALUE_DATEReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_DATE]

	ATTRIBUTE_VALUE_ENUMERATIONs           map[*ATTRIBUTE_VALUE_ENUMERATION]any
	ATTRIBUTE_VALUE_ENUMERATIONs_mapString map[string]*ATTRIBUTE_VALUE_ENUMERATION

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_ENUMERATIONCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_ENUMERATION]
	OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_ENUMERATION]
	OnAfterATTRIBUTE_VALUE_ENUMERATIONDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_ENUMERATION]
	OnAfterATTRIBUTE_VALUE_ENUMERATIONReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_ENUMERATION]

	ATTRIBUTE_VALUE_INTEGERs           map[*ATTRIBUTE_VALUE_INTEGER]any
	ATTRIBUTE_VALUE_INTEGERs_mapString map[string]*ATTRIBUTE_VALUE_INTEGER

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_INTEGERCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_INTEGER]
	OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_INTEGER]
	OnAfterATTRIBUTE_VALUE_INTEGERDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_INTEGER]
	OnAfterATTRIBUTE_VALUE_INTEGERReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_INTEGER]

	ATTRIBUTE_VALUE_REALs           map[*ATTRIBUTE_VALUE_REAL]any
	ATTRIBUTE_VALUE_REALs_mapString map[string]*ATTRIBUTE_VALUE_REAL

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_REALCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_REAL]
	OnAfterATTRIBUTE_VALUE_REALUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_REAL]
	OnAfterATTRIBUTE_VALUE_REALDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_REAL]
	OnAfterATTRIBUTE_VALUE_REALReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_REAL]

	ATTRIBUTE_VALUE_STRINGs           map[*ATTRIBUTE_VALUE_STRING]any
	ATTRIBUTE_VALUE_STRINGs_mapString map[string]*ATTRIBUTE_VALUE_STRING

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_STRINGCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_STRING]
	OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_STRING]
	OnAfterATTRIBUTE_VALUE_STRINGDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_STRING]
	OnAfterATTRIBUTE_VALUE_STRINGReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_STRING]

	ATTRIBUTE_VALUE_XHTMLs           map[*ATTRIBUTE_VALUE_XHTML]any
	ATTRIBUTE_VALUE_XHTMLs_mapString map[string]*ATTRIBUTE_VALUE_XHTML

	// insertion point for slice of pointers maps
	ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML

	ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML

	OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_XHTML]
	OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_XHTML]
	OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_XHTML]
	OnAfterATTRIBUTE_VALUE_XHTMLReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_XHTML]

	AnyTypes           map[*AnyType]any
	AnyTypes_mapString map[string]*AnyType

	// insertion point for slice of pointers maps
	OnAfterAnyTypeCreateCallback OnAfterCreateInterface[AnyType]
	OnAfterAnyTypeUpdateCallback OnAfterUpdateInterface[AnyType]
	OnAfterAnyTypeDeleteCallback OnAfterDeleteInterface[AnyType]
	OnAfterAnyTypeReadCallback   OnAfterReadInterface[AnyType]

	DATATYPE_DEFINITION_BOOLEANs           map[*DATATYPE_DEFINITION_BOOLEAN]any
	DATATYPE_DEFINITION_BOOLEANs_mapString map[string]*DATATYPE_DEFINITION_BOOLEAN

	// insertion point for slice of pointers maps
	DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_BOOLEAN

	OnAfterDATATYPE_DEFINITION_BOOLEANCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_BOOLEAN]
	OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_BOOLEAN]
	OnAfterDATATYPE_DEFINITION_BOOLEANDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_BOOLEAN]
	OnAfterDATATYPE_DEFINITION_BOOLEANReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_BOOLEAN]

	DATATYPE_DEFINITION_DATEs           map[*DATATYPE_DEFINITION_DATE]any
	DATATYPE_DEFINITION_DATEs_mapString map[string]*DATATYPE_DEFINITION_DATE

	// insertion point for slice of pointers maps
	DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_DATE

	OnAfterDATATYPE_DEFINITION_DATECreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_DATE]
	OnAfterDATATYPE_DEFINITION_DATEUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_DATE]
	OnAfterDATATYPE_DEFINITION_DATEDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_DATE]
	OnAfterDATATYPE_DEFINITION_DATEReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_DATE]

	DATATYPE_DEFINITION_ENUMERATIONs           map[*DATATYPE_DEFINITION_ENUMERATION]any
	DATATYPE_DEFINITION_ENUMERATIONs_mapString map[string]*DATATYPE_DEFINITION_ENUMERATION

	// insertion point for slice of pointers maps
	DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_ENUMERATION

	DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_ENUM_VALUE_reverseMap map[*ENUM_VALUE]*DATATYPE_DEFINITION_ENUMERATION

	OnAfterDATATYPE_DEFINITION_ENUMERATIONCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_ENUMERATION]
	OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_ENUMERATION]
	OnAfterDATATYPE_DEFINITION_ENUMERATIONDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_ENUMERATION]
	OnAfterDATATYPE_DEFINITION_ENUMERATIONReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_ENUMERATION]

	DATATYPE_DEFINITION_INTEGERs           map[*DATATYPE_DEFINITION_INTEGER]any
	DATATYPE_DEFINITION_INTEGERs_mapString map[string]*DATATYPE_DEFINITION_INTEGER

	// insertion point for slice of pointers maps
	DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_INTEGER

	OnAfterDATATYPE_DEFINITION_INTEGERCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_INTEGER]
	OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_INTEGER]
	OnAfterDATATYPE_DEFINITION_INTEGERDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_INTEGER]
	OnAfterDATATYPE_DEFINITION_INTEGERReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_INTEGER]

	DATATYPE_DEFINITION_REALs           map[*DATATYPE_DEFINITION_REAL]any
	DATATYPE_DEFINITION_REALs_mapString map[string]*DATATYPE_DEFINITION_REAL

	// insertion point for slice of pointers maps
	DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_REAL

	OnAfterDATATYPE_DEFINITION_REALCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_REAL]
	OnAfterDATATYPE_DEFINITION_REALUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_REAL]
	OnAfterDATATYPE_DEFINITION_REALDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_REAL]
	OnAfterDATATYPE_DEFINITION_REALReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_REAL]

	DATATYPE_DEFINITION_STRINGs           map[*DATATYPE_DEFINITION_STRING]any
	DATATYPE_DEFINITION_STRINGs_mapString map[string]*DATATYPE_DEFINITION_STRING

	// insertion point for slice of pointers maps
	DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_STRING

	OnAfterDATATYPE_DEFINITION_STRINGCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_STRING]
	OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_STRING]
	OnAfterDATATYPE_DEFINITION_STRINGDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_STRING]
	OnAfterDATATYPE_DEFINITION_STRINGReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_STRING]

	DATATYPE_DEFINITION_XHTMLs           map[*DATATYPE_DEFINITION_XHTML]any
	DATATYPE_DEFINITION_XHTMLs_mapString map[string]*DATATYPE_DEFINITION_XHTML

	// insertion point for slice of pointers maps
	DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_XHTML

	OnAfterDATATYPE_DEFINITION_XHTMLCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_XHTML]
	OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_XHTML]
	OnAfterDATATYPE_DEFINITION_XHTMLDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_XHTML]
	OnAfterDATATYPE_DEFINITION_XHTMLReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_XHTML]

	EMBEDDED_VALUEs           map[*EMBEDDED_VALUE]any
	EMBEDDED_VALUEs_mapString map[string]*EMBEDDED_VALUE

	// insertion point for slice of pointers maps
	OnAfterEMBEDDED_VALUECreateCallback OnAfterCreateInterface[EMBEDDED_VALUE]
	OnAfterEMBEDDED_VALUEUpdateCallback OnAfterUpdateInterface[EMBEDDED_VALUE]
	OnAfterEMBEDDED_VALUEDeleteCallback OnAfterDeleteInterface[EMBEDDED_VALUE]
	OnAfterEMBEDDED_VALUEReadCallback   OnAfterReadInterface[EMBEDDED_VALUE]

	ENUM_VALUEs           map[*ENUM_VALUE]any
	ENUM_VALUEs_mapString map[string]*ENUM_VALUE

	// insertion point for slice of pointers maps
	ENUM_VALUE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*ENUM_VALUE

	ENUM_VALUE_PROPERTIES_EMBEDDED_VALUE_reverseMap map[*EMBEDDED_VALUE]*ENUM_VALUE

	OnAfterENUM_VALUECreateCallback OnAfterCreateInterface[ENUM_VALUE]
	OnAfterENUM_VALUEUpdateCallback OnAfterUpdateInterface[ENUM_VALUE]
	OnAfterENUM_VALUEDeleteCallback OnAfterDeleteInterface[ENUM_VALUE]
	OnAfterENUM_VALUEReadCallback   OnAfterReadInterface[ENUM_VALUE]

	RELATION_GROUPs           map[*RELATION_GROUP]any
	RELATION_GROUPs_mapString map[string]*RELATION_GROUP

	// insertion point for slice of pointers maps
	RELATION_GROUP_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*RELATION_GROUP

	OnAfterRELATION_GROUPCreateCallback OnAfterCreateInterface[RELATION_GROUP]
	OnAfterRELATION_GROUPUpdateCallback OnAfterUpdateInterface[RELATION_GROUP]
	OnAfterRELATION_GROUPDeleteCallback OnAfterDeleteInterface[RELATION_GROUP]
	OnAfterRELATION_GROUPReadCallback   OnAfterReadInterface[RELATION_GROUP]

	RELATION_GROUP_TYPEs           map[*RELATION_GROUP_TYPE]any
	RELATION_GROUP_TYPEs_mapString map[string]*RELATION_GROUP_TYPE

	// insertion point for slice of pointers maps
	RELATION_GROUP_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*RELATION_GROUP_TYPE

	RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap map[*ATTRIBUTE_DEFINITION_BOOLEAN]*RELATION_GROUP_TYPE

	RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap map[*ATTRIBUTE_DEFINITION_DATE]*RELATION_GROUP_TYPE

	RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap map[*ATTRIBUTE_DEFINITION_ENUMERATION]*RELATION_GROUP_TYPE

	RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap map[*ATTRIBUTE_DEFINITION_INTEGER]*RELATION_GROUP_TYPE

	RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap map[*ATTRIBUTE_DEFINITION_REAL]*RELATION_GROUP_TYPE

	RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap map[*ATTRIBUTE_DEFINITION_STRING]*RELATION_GROUP_TYPE

	RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap map[*ATTRIBUTE_DEFINITION_XHTML]*RELATION_GROUP_TYPE

	OnAfterRELATION_GROUP_TYPECreateCallback OnAfterCreateInterface[RELATION_GROUP_TYPE]
	OnAfterRELATION_GROUP_TYPEUpdateCallback OnAfterUpdateInterface[RELATION_GROUP_TYPE]
	OnAfterRELATION_GROUP_TYPEDeleteCallback OnAfterDeleteInterface[RELATION_GROUP_TYPE]
	OnAfterRELATION_GROUP_TYPEReadCallback   OnAfterReadInterface[RELATION_GROUP_TYPE]

	REQ_IFs           map[*REQ_IF]any
	REQ_IFs_mapString map[string]*REQ_IF

	// insertion point for slice of pointers maps
	REQ_IF_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap map[*REQ_IF_TOOL_EXTENSION]*REQ_IF

	OnAfterREQ_IFCreateCallback OnAfterCreateInterface[REQ_IF]
	OnAfterREQ_IFUpdateCallback OnAfterUpdateInterface[REQ_IF]
	OnAfterREQ_IFDeleteCallback OnAfterDeleteInterface[REQ_IF]
	OnAfterREQ_IFReadCallback   OnAfterReadInterface[REQ_IF]

	REQ_IF_CONTENTs           map[*REQ_IF_CONTENT]any
	REQ_IF_CONTENTs_mapString map[string]*REQ_IF_CONTENT

	// insertion point for slice of pointers maps
	REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap map[*DATATYPE_DEFINITION_BOOLEAN]*REQ_IF_CONTENT

	REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap map[*DATATYPE_DEFINITION_DATE]*REQ_IF_CONTENT

	REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap map[*DATATYPE_DEFINITION_ENUMERATION]*REQ_IF_CONTENT

	REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap map[*DATATYPE_DEFINITION_INTEGER]*REQ_IF_CONTENT

	REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap map[*DATATYPE_DEFINITION_REAL]*REQ_IF_CONTENT

	REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap map[*DATATYPE_DEFINITION_STRING]*REQ_IF_CONTENT

	REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap map[*DATATYPE_DEFINITION_XHTML]*REQ_IF_CONTENT

	REQ_IF_CONTENT_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap map[*RELATION_GROUP_TYPE]*REQ_IF_CONTENT

	REQ_IF_CONTENT_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap map[*SPEC_OBJECT_TYPE]*REQ_IF_CONTENT

	REQ_IF_CONTENT_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap map[*SPEC_RELATION_TYPE]*REQ_IF_CONTENT

	REQ_IF_CONTENT_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap map[*SPECIFICATION_TYPE]*REQ_IF_CONTENT

	REQ_IF_CONTENT_SPEC_OBJECTS_SPEC_OBJECT_reverseMap map[*SPEC_OBJECT]*REQ_IF_CONTENT

	REQ_IF_CONTENT_SPEC_RELATIONS_SPEC_RELATION_reverseMap map[*SPEC_RELATION]*REQ_IF_CONTENT

	REQ_IF_CONTENT_SPECIFICATIONS_SPECIFICATION_reverseMap map[*SPECIFICATION]*REQ_IF_CONTENT

	REQ_IF_CONTENT_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap map[*RELATION_GROUP]*REQ_IF_CONTENT

	OnAfterREQ_IF_CONTENTCreateCallback OnAfterCreateInterface[REQ_IF_CONTENT]
	OnAfterREQ_IF_CONTENTUpdateCallback OnAfterUpdateInterface[REQ_IF_CONTENT]
	OnAfterREQ_IF_CONTENTDeleteCallback OnAfterDeleteInterface[REQ_IF_CONTENT]
	OnAfterREQ_IF_CONTENTReadCallback   OnAfterReadInterface[REQ_IF_CONTENT]

	REQ_IF_HEADERs           map[*REQ_IF_HEADER]any
	REQ_IF_HEADERs_mapString map[string]*REQ_IF_HEADER

	// insertion point for slice of pointers maps
	OnAfterREQ_IF_HEADERCreateCallback OnAfterCreateInterface[REQ_IF_HEADER]
	OnAfterREQ_IF_HEADERUpdateCallback OnAfterUpdateInterface[REQ_IF_HEADER]
	OnAfterREQ_IF_HEADERDeleteCallback OnAfterDeleteInterface[REQ_IF_HEADER]
	OnAfterREQ_IF_HEADERReadCallback   OnAfterReadInterface[REQ_IF_HEADER]

	REQ_IF_TOOL_EXTENSIONs           map[*REQ_IF_TOOL_EXTENSION]any
	REQ_IF_TOOL_EXTENSIONs_mapString map[string]*REQ_IF_TOOL_EXTENSION

	// insertion point for slice of pointers maps
	OnAfterREQ_IF_TOOL_EXTENSIONCreateCallback OnAfterCreateInterface[REQ_IF_TOOL_EXTENSION]
	OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback OnAfterUpdateInterface[REQ_IF_TOOL_EXTENSION]
	OnAfterREQ_IF_TOOL_EXTENSIONDeleteCallback OnAfterDeleteInterface[REQ_IF_TOOL_EXTENSION]
	OnAfterREQ_IF_TOOL_EXTENSIONReadCallback   OnAfterReadInterface[REQ_IF_TOOL_EXTENSION]

	SPECIFICATIONs           map[*SPECIFICATION]any
	SPECIFICATIONs_mapString map[string]*SPECIFICATION

	// insertion point for slice of pointers maps
	SPECIFICATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*SPECIFICATION

	SPECIFICATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap map[*ATTRIBUTE_VALUE_BOOLEAN]*SPECIFICATION

	SPECIFICATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap map[*ATTRIBUTE_VALUE_DATE]*SPECIFICATION

	SPECIFICATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap map[*ATTRIBUTE_VALUE_ENUMERATION]*SPECIFICATION

	SPECIFICATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap map[*ATTRIBUTE_VALUE_INTEGER]*SPECIFICATION

	SPECIFICATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap map[*ATTRIBUTE_VALUE_REAL]*SPECIFICATION

	SPECIFICATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap map[*ATTRIBUTE_VALUE_STRING]*SPECIFICATION

	SPECIFICATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap map[*ATTRIBUTE_VALUE_XHTML]*SPECIFICATION

	SPECIFICATION_CHILDREN_SPEC_HIERARCHY_reverseMap map[*SPEC_HIERARCHY]*SPECIFICATION

	OnAfterSPECIFICATIONCreateCallback OnAfterCreateInterface[SPECIFICATION]
	OnAfterSPECIFICATIONUpdateCallback OnAfterUpdateInterface[SPECIFICATION]
	OnAfterSPECIFICATIONDeleteCallback OnAfterDeleteInterface[SPECIFICATION]
	OnAfterSPECIFICATIONReadCallback   OnAfterReadInterface[SPECIFICATION]

	SPECIFICATION_TYPEs           map[*SPECIFICATION_TYPE]any
	SPECIFICATION_TYPEs_mapString map[string]*SPECIFICATION_TYPE

	// insertion point for slice of pointers maps
	SPECIFICATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*SPECIFICATION_TYPE

	SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPECIFICATION_TYPE

	SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap map[*ATTRIBUTE_DEFINITION_DATE]*SPECIFICATION_TYPE

	SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPECIFICATION_TYPE

	SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap map[*ATTRIBUTE_DEFINITION_INTEGER]*SPECIFICATION_TYPE

	SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap map[*ATTRIBUTE_DEFINITION_REAL]*SPECIFICATION_TYPE

	SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap map[*ATTRIBUTE_DEFINITION_STRING]*SPECIFICATION_TYPE

	SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap map[*ATTRIBUTE_DEFINITION_XHTML]*SPECIFICATION_TYPE

	OnAfterSPECIFICATION_TYPECreateCallback OnAfterCreateInterface[SPECIFICATION_TYPE]
	OnAfterSPECIFICATION_TYPEUpdateCallback OnAfterUpdateInterface[SPECIFICATION_TYPE]
	OnAfterSPECIFICATION_TYPEDeleteCallback OnAfterDeleteInterface[SPECIFICATION_TYPE]
	OnAfterSPECIFICATION_TYPEReadCallback   OnAfterReadInterface[SPECIFICATION_TYPE]

	SPEC_HIERARCHYs           map[*SPEC_HIERARCHY]any
	SPEC_HIERARCHYs_mapString map[string]*SPEC_HIERARCHY

	// insertion point for slice of pointers maps
	SPEC_HIERARCHY_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*SPEC_HIERARCHY

	SPEC_HIERARCHY_CHILDREN_SPEC_HIERARCHY_reverseMap map[*SPEC_HIERARCHY]*SPEC_HIERARCHY

	OnAfterSPEC_HIERARCHYCreateCallback OnAfterCreateInterface[SPEC_HIERARCHY]
	OnAfterSPEC_HIERARCHYUpdateCallback OnAfterUpdateInterface[SPEC_HIERARCHY]
	OnAfterSPEC_HIERARCHYDeleteCallback OnAfterDeleteInterface[SPEC_HIERARCHY]
	OnAfterSPEC_HIERARCHYReadCallback   OnAfterReadInterface[SPEC_HIERARCHY]

	SPEC_OBJECTs           map[*SPEC_OBJECT]any
	SPEC_OBJECTs_mapString map[string]*SPEC_OBJECT

	// insertion point for slice of pointers maps
	SPEC_OBJECT_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*SPEC_OBJECT

	SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap map[*ATTRIBUTE_VALUE_BOOLEAN]*SPEC_OBJECT

	SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap map[*ATTRIBUTE_VALUE_DATE]*SPEC_OBJECT

	SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap map[*ATTRIBUTE_VALUE_ENUMERATION]*SPEC_OBJECT

	SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap map[*ATTRIBUTE_VALUE_INTEGER]*SPEC_OBJECT

	SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap map[*ATTRIBUTE_VALUE_REAL]*SPEC_OBJECT

	SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap map[*ATTRIBUTE_VALUE_STRING]*SPEC_OBJECT

	SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap map[*ATTRIBUTE_VALUE_XHTML]*SPEC_OBJECT

	OnAfterSPEC_OBJECTCreateCallback OnAfterCreateInterface[SPEC_OBJECT]
	OnAfterSPEC_OBJECTUpdateCallback OnAfterUpdateInterface[SPEC_OBJECT]
	OnAfterSPEC_OBJECTDeleteCallback OnAfterDeleteInterface[SPEC_OBJECT]
	OnAfterSPEC_OBJECTReadCallback   OnAfterReadInterface[SPEC_OBJECT]

	SPEC_OBJECT_TYPEs           map[*SPEC_OBJECT_TYPE]any
	SPEC_OBJECT_TYPEs_mapString map[string]*SPEC_OBJECT_TYPE

	// insertion point for slice of pointers maps
	SPEC_OBJECT_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*SPEC_OBJECT_TYPE

	SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPEC_OBJECT_TYPE

	SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap map[*ATTRIBUTE_DEFINITION_DATE]*SPEC_OBJECT_TYPE

	SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPEC_OBJECT_TYPE

	SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap map[*ATTRIBUTE_DEFINITION_INTEGER]*SPEC_OBJECT_TYPE

	SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap map[*ATTRIBUTE_DEFINITION_REAL]*SPEC_OBJECT_TYPE

	SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap map[*ATTRIBUTE_DEFINITION_STRING]*SPEC_OBJECT_TYPE

	SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap map[*ATTRIBUTE_DEFINITION_XHTML]*SPEC_OBJECT_TYPE

	OnAfterSPEC_OBJECT_TYPECreateCallback OnAfterCreateInterface[SPEC_OBJECT_TYPE]
	OnAfterSPEC_OBJECT_TYPEUpdateCallback OnAfterUpdateInterface[SPEC_OBJECT_TYPE]
	OnAfterSPEC_OBJECT_TYPEDeleteCallback OnAfterDeleteInterface[SPEC_OBJECT_TYPE]
	OnAfterSPEC_OBJECT_TYPEReadCallback   OnAfterReadInterface[SPEC_OBJECT_TYPE]

	SPEC_RELATIONs           map[*SPEC_RELATION]any
	SPEC_RELATIONs_mapString map[string]*SPEC_RELATION

	// insertion point for slice of pointers maps
	SPEC_RELATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*SPEC_RELATION

	SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap map[*ATTRIBUTE_VALUE_BOOLEAN]*SPEC_RELATION

	SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap map[*ATTRIBUTE_VALUE_DATE]*SPEC_RELATION

	SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap map[*ATTRIBUTE_VALUE_ENUMERATION]*SPEC_RELATION

	SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap map[*ATTRIBUTE_VALUE_INTEGER]*SPEC_RELATION

	SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap map[*ATTRIBUTE_VALUE_REAL]*SPEC_RELATION

	SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap map[*ATTRIBUTE_VALUE_STRING]*SPEC_RELATION

	SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap map[*ATTRIBUTE_VALUE_XHTML]*SPEC_RELATION

	OnAfterSPEC_RELATIONCreateCallback OnAfterCreateInterface[SPEC_RELATION]
	OnAfterSPEC_RELATIONUpdateCallback OnAfterUpdateInterface[SPEC_RELATION]
	OnAfterSPEC_RELATIONDeleteCallback OnAfterDeleteInterface[SPEC_RELATION]
	OnAfterSPEC_RELATIONReadCallback   OnAfterReadInterface[SPEC_RELATION]

	SPEC_RELATION_TYPEs           map[*SPEC_RELATION_TYPE]any
	SPEC_RELATION_TYPEs_mapString map[string]*SPEC_RELATION_TYPE

	// insertion point for slice of pointers maps
	SPEC_RELATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap map[*ALTERNATIVE_ID]*SPEC_RELATION_TYPE

	SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPEC_RELATION_TYPE

	SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap map[*ATTRIBUTE_DEFINITION_DATE]*SPEC_RELATION_TYPE

	SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPEC_RELATION_TYPE

	SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap map[*ATTRIBUTE_DEFINITION_INTEGER]*SPEC_RELATION_TYPE

	SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap map[*ATTRIBUTE_DEFINITION_REAL]*SPEC_RELATION_TYPE

	SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap map[*ATTRIBUTE_DEFINITION_STRING]*SPEC_RELATION_TYPE

	SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap map[*ATTRIBUTE_DEFINITION_XHTML]*SPEC_RELATION_TYPE

	OnAfterSPEC_RELATION_TYPECreateCallback OnAfterCreateInterface[SPEC_RELATION_TYPE]
	OnAfterSPEC_RELATION_TYPEUpdateCallback OnAfterUpdateInterface[SPEC_RELATION_TYPE]
	OnAfterSPEC_RELATION_TYPEDeleteCallback OnAfterDeleteInterface[SPEC_RELATION_TYPE]
	OnAfterSPEC_RELATION_TYPEReadCallback   OnAfterReadInterface[SPEC_RELATION_TYPE]

	XHTML_CONTENTs           map[*XHTML_CONTENT]any
	XHTML_CONTENTs_mapString map[string]*XHTML_CONTENT

	// insertion point for slice of pointers maps
	OnAfterXHTML_CONTENTCreateCallback OnAfterCreateInterface[XHTML_CONTENT]
	OnAfterXHTML_CONTENTUpdateCallback OnAfterUpdateInterface[XHTML_CONTENT]
	OnAfterXHTML_CONTENTDeleteCallback OnAfterDeleteInterface[XHTML_CONTENT]
	OnAfterXHTML_CONTENTReadCallback   OnAfterReadInterface[XHTML_CONTENT]

	Xhtml_InlPres_types           map[*Xhtml_InlPres_type]any
	Xhtml_InlPres_types_mapString map[string]*Xhtml_InlPres_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_InlPres_typeCreateCallback OnAfterCreateInterface[Xhtml_InlPres_type]
	OnAfterXhtml_InlPres_typeUpdateCallback OnAfterUpdateInterface[Xhtml_InlPres_type]
	OnAfterXhtml_InlPres_typeDeleteCallback OnAfterDeleteInterface[Xhtml_InlPres_type]
	OnAfterXhtml_InlPres_typeReadCallback   OnAfterReadInterface[Xhtml_InlPres_type]

	Xhtml_a_types           map[*Xhtml_a_type]any
	Xhtml_a_types_mapString map[string]*Xhtml_a_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_a_typeCreateCallback OnAfterCreateInterface[Xhtml_a_type]
	OnAfterXhtml_a_typeUpdateCallback OnAfterUpdateInterface[Xhtml_a_type]
	OnAfterXhtml_a_typeDeleteCallback OnAfterDeleteInterface[Xhtml_a_type]
	OnAfterXhtml_a_typeReadCallback   OnAfterReadInterface[Xhtml_a_type]

	Xhtml_abbr_types           map[*Xhtml_abbr_type]any
	Xhtml_abbr_types_mapString map[string]*Xhtml_abbr_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_abbr_typeCreateCallback OnAfterCreateInterface[Xhtml_abbr_type]
	OnAfterXhtml_abbr_typeUpdateCallback OnAfterUpdateInterface[Xhtml_abbr_type]
	OnAfterXhtml_abbr_typeDeleteCallback OnAfterDeleteInterface[Xhtml_abbr_type]
	OnAfterXhtml_abbr_typeReadCallback   OnAfterReadInterface[Xhtml_abbr_type]

	Xhtml_acronym_types           map[*Xhtml_acronym_type]any
	Xhtml_acronym_types_mapString map[string]*Xhtml_acronym_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_acronym_typeCreateCallback OnAfterCreateInterface[Xhtml_acronym_type]
	OnAfterXhtml_acronym_typeUpdateCallback OnAfterUpdateInterface[Xhtml_acronym_type]
	OnAfterXhtml_acronym_typeDeleteCallback OnAfterDeleteInterface[Xhtml_acronym_type]
	OnAfterXhtml_acronym_typeReadCallback   OnAfterReadInterface[Xhtml_acronym_type]

	Xhtml_address_types           map[*Xhtml_address_type]any
	Xhtml_address_types_mapString map[string]*Xhtml_address_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_address_typeCreateCallback OnAfterCreateInterface[Xhtml_address_type]
	OnAfterXhtml_address_typeUpdateCallback OnAfterUpdateInterface[Xhtml_address_type]
	OnAfterXhtml_address_typeDeleteCallback OnAfterDeleteInterface[Xhtml_address_type]
	OnAfterXhtml_address_typeReadCallback   OnAfterReadInterface[Xhtml_address_type]

	Xhtml_blockquote_types           map[*Xhtml_blockquote_type]any
	Xhtml_blockquote_types_mapString map[string]*Xhtml_blockquote_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_blockquote_typeCreateCallback OnAfterCreateInterface[Xhtml_blockquote_type]
	OnAfterXhtml_blockquote_typeUpdateCallback OnAfterUpdateInterface[Xhtml_blockquote_type]
	OnAfterXhtml_blockquote_typeDeleteCallback OnAfterDeleteInterface[Xhtml_blockquote_type]
	OnAfterXhtml_blockquote_typeReadCallback   OnAfterReadInterface[Xhtml_blockquote_type]

	Xhtml_br_types           map[*Xhtml_br_type]any
	Xhtml_br_types_mapString map[string]*Xhtml_br_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_br_typeCreateCallback OnAfterCreateInterface[Xhtml_br_type]
	OnAfterXhtml_br_typeUpdateCallback OnAfterUpdateInterface[Xhtml_br_type]
	OnAfterXhtml_br_typeDeleteCallback OnAfterDeleteInterface[Xhtml_br_type]
	OnAfterXhtml_br_typeReadCallback   OnAfterReadInterface[Xhtml_br_type]

	Xhtml_caption_types           map[*Xhtml_caption_type]any
	Xhtml_caption_types_mapString map[string]*Xhtml_caption_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_caption_typeCreateCallback OnAfterCreateInterface[Xhtml_caption_type]
	OnAfterXhtml_caption_typeUpdateCallback OnAfterUpdateInterface[Xhtml_caption_type]
	OnAfterXhtml_caption_typeDeleteCallback OnAfterDeleteInterface[Xhtml_caption_type]
	OnAfterXhtml_caption_typeReadCallback   OnAfterReadInterface[Xhtml_caption_type]

	Xhtml_cite_types           map[*Xhtml_cite_type]any
	Xhtml_cite_types_mapString map[string]*Xhtml_cite_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_cite_typeCreateCallback OnAfterCreateInterface[Xhtml_cite_type]
	OnAfterXhtml_cite_typeUpdateCallback OnAfterUpdateInterface[Xhtml_cite_type]
	OnAfterXhtml_cite_typeDeleteCallback OnAfterDeleteInterface[Xhtml_cite_type]
	OnAfterXhtml_cite_typeReadCallback   OnAfterReadInterface[Xhtml_cite_type]

	Xhtml_code_types           map[*Xhtml_code_type]any
	Xhtml_code_types_mapString map[string]*Xhtml_code_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_code_typeCreateCallback OnAfterCreateInterface[Xhtml_code_type]
	OnAfterXhtml_code_typeUpdateCallback OnAfterUpdateInterface[Xhtml_code_type]
	OnAfterXhtml_code_typeDeleteCallback OnAfterDeleteInterface[Xhtml_code_type]
	OnAfterXhtml_code_typeReadCallback   OnAfterReadInterface[Xhtml_code_type]

	Xhtml_col_types           map[*Xhtml_col_type]any
	Xhtml_col_types_mapString map[string]*Xhtml_col_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_col_typeCreateCallback OnAfterCreateInterface[Xhtml_col_type]
	OnAfterXhtml_col_typeUpdateCallback OnAfterUpdateInterface[Xhtml_col_type]
	OnAfterXhtml_col_typeDeleteCallback OnAfterDeleteInterface[Xhtml_col_type]
	OnAfterXhtml_col_typeReadCallback   OnAfterReadInterface[Xhtml_col_type]

	Xhtml_colgroup_types           map[*Xhtml_colgroup_type]any
	Xhtml_colgroup_types_mapString map[string]*Xhtml_colgroup_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_colgroup_typeCreateCallback OnAfterCreateInterface[Xhtml_colgroup_type]
	OnAfterXhtml_colgroup_typeUpdateCallback OnAfterUpdateInterface[Xhtml_colgroup_type]
	OnAfterXhtml_colgroup_typeDeleteCallback OnAfterDeleteInterface[Xhtml_colgroup_type]
	OnAfterXhtml_colgroup_typeReadCallback   OnAfterReadInterface[Xhtml_colgroup_type]

	Xhtml_dd_types           map[*Xhtml_dd_type]any
	Xhtml_dd_types_mapString map[string]*Xhtml_dd_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_dd_typeCreateCallback OnAfterCreateInterface[Xhtml_dd_type]
	OnAfterXhtml_dd_typeUpdateCallback OnAfterUpdateInterface[Xhtml_dd_type]
	OnAfterXhtml_dd_typeDeleteCallback OnAfterDeleteInterface[Xhtml_dd_type]
	OnAfterXhtml_dd_typeReadCallback   OnAfterReadInterface[Xhtml_dd_type]

	Xhtml_dfn_types           map[*Xhtml_dfn_type]any
	Xhtml_dfn_types_mapString map[string]*Xhtml_dfn_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_dfn_typeCreateCallback OnAfterCreateInterface[Xhtml_dfn_type]
	OnAfterXhtml_dfn_typeUpdateCallback OnAfterUpdateInterface[Xhtml_dfn_type]
	OnAfterXhtml_dfn_typeDeleteCallback OnAfterDeleteInterface[Xhtml_dfn_type]
	OnAfterXhtml_dfn_typeReadCallback   OnAfterReadInterface[Xhtml_dfn_type]

	Xhtml_div_types           map[*Xhtml_div_type]any
	Xhtml_div_types_mapString map[string]*Xhtml_div_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_div_typeCreateCallback OnAfterCreateInterface[Xhtml_div_type]
	OnAfterXhtml_div_typeUpdateCallback OnAfterUpdateInterface[Xhtml_div_type]
	OnAfterXhtml_div_typeDeleteCallback OnAfterDeleteInterface[Xhtml_div_type]
	OnAfterXhtml_div_typeReadCallback   OnAfterReadInterface[Xhtml_div_type]

	Xhtml_dl_types           map[*Xhtml_dl_type]any
	Xhtml_dl_types_mapString map[string]*Xhtml_dl_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_dl_typeCreateCallback OnAfterCreateInterface[Xhtml_dl_type]
	OnAfterXhtml_dl_typeUpdateCallback OnAfterUpdateInterface[Xhtml_dl_type]
	OnAfterXhtml_dl_typeDeleteCallback OnAfterDeleteInterface[Xhtml_dl_type]
	OnAfterXhtml_dl_typeReadCallback   OnAfterReadInterface[Xhtml_dl_type]

	Xhtml_dt_types           map[*Xhtml_dt_type]any
	Xhtml_dt_types_mapString map[string]*Xhtml_dt_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_dt_typeCreateCallback OnAfterCreateInterface[Xhtml_dt_type]
	OnAfterXhtml_dt_typeUpdateCallback OnAfterUpdateInterface[Xhtml_dt_type]
	OnAfterXhtml_dt_typeDeleteCallback OnAfterDeleteInterface[Xhtml_dt_type]
	OnAfterXhtml_dt_typeReadCallback   OnAfterReadInterface[Xhtml_dt_type]

	Xhtml_edit_types           map[*Xhtml_edit_type]any
	Xhtml_edit_types_mapString map[string]*Xhtml_edit_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_edit_typeCreateCallback OnAfterCreateInterface[Xhtml_edit_type]
	OnAfterXhtml_edit_typeUpdateCallback OnAfterUpdateInterface[Xhtml_edit_type]
	OnAfterXhtml_edit_typeDeleteCallback OnAfterDeleteInterface[Xhtml_edit_type]
	OnAfterXhtml_edit_typeReadCallback   OnAfterReadInterface[Xhtml_edit_type]

	Xhtml_em_types           map[*Xhtml_em_type]any
	Xhtml_em_types_mapString map[string]*Xhtml_em_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_em_typeCreateCallback OnAfterCreateInterface[Xhtml_em_type]
	OnAfterXhtml_em_typeUpdateCallback OnAfterUpdateInterface[Xhtml_em_type]
	OnAfterXhtml_em_typeDeleteCallback OnAfterDeleteInterface[Xhtml_em_type]
	OnAfterXhtml_em_typeReadCallback   OnAfterReadInterface[Xhtml_em_type]

	Xhtml_h1_types           map[*Xhtml_h1_type]any
	Xhtml_h1_types_mapString map[string]*Xhtml_h1_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_h1_typeCreateCallback OnAfterCreateInterface[Xhtml_h1_type]
	OnAfterXhtml_h1_typeUpdateCallback OnAfterUpdateInterface[Xhtml_h1_type]
	OnAfterXhtml_h1_typeDeleteCallback OnAfterDeleteInterface[Xhtml_h1_type]
	OnAfterXhtml_h1_typeReadCallback   OnAfterReadInterface[Xhtml_h1_type]

	Xhtml_h2_types           map[*Xhtml_h2_type]any
	Xhtml_h2_types_mapString map[string]*Xhtml_h2_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_h2_typeCreateCallback OnAfterCreateInterface[Xhtml_h2_type]
	OnAfterXhtml_h2_typeUpdateCallback OnAfterUpdateInterface[Xhtml_h2_type]
	OnAfterXhtml_h2_typeDeleteCallback OnAfterDeleteInterface[Xhtml_h2_type]
	OnAfterXhtml_h2_typeReadCallback   OnAfterReadInterface[Xhtml_h2_type]

	Xhtml_h3_types           map[*Xhtml_h3_type]any
	Xhtml_h3_types_mapString map[string]*Xhtml_h3_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_h3_typeCreateCallback OnAfterCreateInterface[Xhtml_h3_type]
	OnAfterXhtml_h3_typeUpdateCallback OnAfterUpdateInterface[Xhtml_h3_type]
	OnAfterXhtml_h3_typeDeleteCallback OnAfterDeleteInterface[Xhtml_h3_type]
	OnAfterXhtml_h3_typeReadCallback   OnAfterReadInterface[Xhtml_h3_type]

	Xhtml_h4_types           map[*Xhtml_h4_type]any
	Xhtml_h4_types_mapString map[string]*Xhtml_h4_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_h4_typeCreateCallback OnAfterCreateInterface[Xhtml_h4_type]
	OnAfterXhtml_h4_typeUpdateCallback OnAfterUpdateInterface[Xhtml_h4_type]
	OnAfterXhtml_h4_typeDeleteCallback OnAfterDeleteInterface[Xhtml_h4_type]
	OnAfterXhtml_h4_typeReadCallback   OnAfterReadInterface[Xhtml_h4_type]

	Xhtml_h5_types           map[*Xhtml_h5_type]any
	Xhtml_h5_types_mapString map[string]*Xhtml_h5_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_h5_typeCreateCallback OnAfterCreateInterface[Xhtml_h5_type]
	OnAfterXhtml_h5_typeUpdateCallback OnAfterUpdateInterface[Xhtml_h5_type]
	OnAfterXhtml_h5_typeDeleteCallback OnAfterDeleteInterface[Xhtml_h5_type]
	OnAfterXhtml_h5_typeReadCallback   OnAfterReadInterface[Xhtml_h5_type]

	Xhtml_h6_types           map[*Xhtml_h6_type]any
	Xhtml_h6_types_mapString map[string]*Xhtml_h6_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_h6_typeCreateCallback OnAfterCreateInterface[Xhtml_h6_type]
	OnAfterXhtml_h6_typeUpdateCallback OnAfterUpdateInterface[Xhtml_h6_type]
	OnAfterXhtml_h6_typeDeleteCallback OnAfterDeleteInterface[Xhtml_h6_type]
	OnAfterXhtml_h6_typeReadCallback   OnAfterReadInterface[Xhtml_h6_type]

	Xhtml_heading_types           map[*Xhtml_heading_type]any
	Xhtml_heading_types_mapString map[string]*Xhtml_heading_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_heading_typeCreateCallback OnAfterCreateInterface[Xhtml_heading_type]
	OnAfterXhtml_heading_typeUpdateCallback OnAfterUpdateInterface[Xhtml_heading_type]
	OnAfterXhtml_heading_typeDeleteCallback OnAfterDeleteInterface[Xhtml_heading_type]
	OnAfterXhtml_heading_typeReadCallback   OnAfterReadInterface[Xhtml_heading_type]

	Xhtml_hr_types           map[*Xhtml_hr_type]any
	Xhtml_hr_types_mapString map[string]*Xhtml_hr_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_hr_typeCreateCallback OnAfterCreateInterface[Xhtml_hr_type]
	OnAfterXhtml_hr_typeUpdateCallback OnAfterUpdateInterface[Xhtml_hr_type]
	OnAfterXhtml_hr_typeDeleteCallback OnAfterDeleteInterface[Xhtml_hr_type]
	OnAfterXhtml_hr_typeReadCallback   OnAfterReadInterface[Xhtml_hr_type]

	Xhtml_kbd_types           map[*Xhtml_kbd_type]any
	Xhtml_kbd_types_mapString map[string]*Xhtml_kbd_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_kbd_typeCreateCallback OnAfterCreateInterface[Xhtml_kbd_type]
	OnAfterXhtml_kbd_typeUpdateCallback OnAfterUpdateInterface[Xhtml_kbd_type]
	OnAfterXhtml_kbd_typeDeleteCallback OnAfterDeleteInterface[Xhtml_kbd_type]
	OnAfterXhtml_kbd_typeReadCallback   OnAfterReadInterface[Xhtml_kbd_type]

	Xhtml_li_types           map[*Xhtml_li_type]any
	Xhtml_li_types_mapString map[string]*Xhtml_li_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_li_typeCreateCallback OnAfterCreateInterface[Xhtml_li_type]
	OnAfterXhtml_li_typeUpdateCallback OnAfterUpdateInterface[Xhtml_li_type]
	OnAfterXhtml_li_typeDeleteCallback OnAfterDeleteInterface[Xhtml_li_type]
	OnAfterXhtml_li_typeReadCallback   OnAfterReadInterface[Xhtml_li_type]

	Xhtml_object_types           map[*Xhtml_object_type]any
	Xhtml_object_types_mapString map[string]*Xhtml_object_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_object_typeCreateCallback OnAfterCreateInterface[Xhtml_object_type]
	OnAfterXhtml_object_typeUpdateCallback OnAfterUpdateInterface[Xhtml_object_type]
	OnAfterXhtml_object_typeDeleteCallback OnAfterDeleteInterface[Xhtml_object_type]
	OnAfterXhtml_object_typeReadCallback   OnAfterReadInterface[Xhtml_object_type]

	Xhtml_ol_types           map[*Xhtml_ol_type]any
	Xhtml_ol_types_mapString map[string]*Xhtml_ol_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_ol_typeCreateCallback OnAfterCreateInterface[Xhtml_ol_type]
	OnAfterXhtml_ol_typeUpdateCallback OnAfterUpdateInterface[Xhtml_ol_type]
	OnAfterXhtml_ol_typeDeleteCallback OnAfterDeleteInterface[Xhtml_ol_type]
	OnAfterXhtml_ol_typeReadCallback   OnAfterReadInterface[Xhtml_ol_type]

	Xhtml_p_types           map[*Xhtml_p_type]any
	Xhtml_p_types_mapString map[string]*Xhtml_p_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_p_typeCreateCallback OnAfterCreateInterface[Xhtml_p_type]
	OnAfterXhtml_p_typeUpdateCallback OnAfterUpdateInterface[Xhtml_p_type]
	OnAfterXhtml_p_typeDeleteCallback OnAfterDeleteInterface[Xhtml_p_type]
	OnAfterXhtml_p_typeReadCallback   OnAfterReadInterface[Xhtml_p_type]

	Xhtml_param_types           map[*Xhtml_param_type]any
	Xhtml_param_types_mapString map[string]*Xhtml_param_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_param_typeCreateCallback OnAfterCreateInterface[Xhtml_param_type]
	OnAfterXhtml_param_typeUpdateCallback OnAfterUpdateInterface[Xhtml_param_type]
	OnAfterXhtml_param_typeDeleteCallback OnAfterDeleteInterface[Xhtml_param_type]
	OnAfterXhtml_param_typeReadCallback   OnAfterReadInterface[Xhtml_param_type]

	Xhtml_pre_types           map[*Xhtml_pre_type]any
	Xhtml_pre_types_mapString map[string]*Xhtml_pre_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_pre_typeCreateCallback OnAfterCreateInterface[Xhtml_pre_type]
	OnAfterXhtml_pre_typeUpdateCallback OnAfterUpdateInterface[Xhtml_pre_type]
	OnAfterXhtml_pre_typeDeleteCallback OnAfterDeleteInterface[Xhtml_pre_type]
	OnAfterXhtml_pre_typeReadCallback   OnAfterReadInterface[Xhtml_pre_type]

	Xhtml_q_types           map[*Xhtml_q_type]any
	Xhtml_q_types_mapString map[string]*Xhtml_q_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_q_typeCreateCallback OnAfterCreateInterface[Xhtml_q_type]
	OnAfterXhtml_q_typeUpdateCallback OnAfterUpdateInterface[Xhtml_q_type]
	OnAfterXhtml_q_typeDeleteCallback OnAfterDeleteInterface[Xhtml_q_type]
	OnAfterXhtml_q_typeReadCallback   OnAfterReadInterface[Xhtml_q_type]

	Xhtml_samp_types           map[*Xhtml_samp_type]any
	Xhtml_samp_types_mapString map[string]*Xhtml_samp_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_samp_typeCreateCallback OnAfterCreateInterface[Xhtml_samp_type]
	OnAfterXhtml_samp_typeUpdateCallback OnAfterUpdateInterface[Xhtml_samp_type]
	OnAfterXhtml_samp_typeDeleteCallback OnAfterDeleteInterface[Xhtml_samp_type]
	OnAfterXhtml_samp_typeReadCallback   OnAfterReadInterface[Xhtml_samp_type]

	Xhtml_span_types           map[*Xhtml_span_type]any
	Xhtml_span_types_mapString map[string]*Xhtml_span_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_span_typeCreateCallback OnAfterCreateInterface[Xhtml_span_type]
	OnAfterXhtml_span_typeUpdateCallback OnAfterUpdateInterface[Xhtml_span_type]
	OnAfterXhtml_span_typeDeleteCallback OnAfterDeleteInterface[Xhtml_span_type]
	OnAfterXhtml_span_typeReadCallback   OnAfterReadInterface[Xhtml_span_type]

	Xhtml_strong_types           map[*Xhtml_strong_type]any
	Xhtml_strong_types_mapString map[string]*Xhtml_strong_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_strong_typeCreateCallback OnAfterCreateInterface[Xhtml_strong_type]
	OnAfterXhtml_strong_typeUpdateCallback OnAfterUpdateInterface[Xhtml_strong_type]
	OnAfterXhtml_strong_typeDeleteCallback OnAfterDeleteInterface[Xhtml_strong_type]
	OnAfterXhtml_strong_typeReadCallback   OnAfterReadInterface[Xhtml_strong_type]

	Xhtml_table_types           map[*Xhtml_table_type]any
	Xhtml_table_types_mapString map[string]*Xhtml_table_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_table_typeCreateCallback OnAfterCreateInterface[Xhtml_table_type]
	OnAfterXhtml_table_typeUpdateCallback OnAfterUpdateInterface[Xhtml_table_type]
	OnAfterXhtml_table_typeDeleteCallback OnAfterDeleteInterface[Xhtml_table_type]
	OnAfterXhtml_table_typeReadCallback   OnAfterReadInterface[Xhtml_table_type]

	Xhtml_tbody_types           map[*Xhtml_tbody_type]any
	Xhtml_tbody_types_mapString map[string]*Xhtml_tbody_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_tbody_typeCreateCallback OnAfterCreateInterface[Xhtml_tbody_type]
	OnAfterXhtml_tbody_typeUpdateCallback OnAfterUpdateInterface[Xhtml_tbody_type]
	OnAfterXhtml_tbody_typeDeleteCallback OnAfterDeleteInterface[Xhtml_tbody_type]
	OnAfterXhtml_tbody_typeReadCallback   OnAfterReadInterface[Xhtml_tbody_type]

	Xhtml_td_types           map[*Xhtml_td_type]any
	Xhtml_td_types_mapString map[string]*Xhtml_td_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_td_typeCreateCallback OnAfterCreateInterface[Xhtml_td_type]
	OnAfterXhtml_td_typeUpdateCallback OnAfterUpdateInterface[Xhtml_td_type]
	OnAfterXhtml_td_typeDeleteCallback OnAfterDeleteInterface[Xhtml_td_type]
	OnAfterXhtml_td_typeReadCallback   OnAfterReadInterface[Xhtml_td_type]

	Xhtml_tfoot_types           map[*Xhtml_tfoot_type]any
	Xhtml_tfoot_types_mapString map[string]*Xhtml_tfoot_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_tfoot_typeCreateCallback OnAfterCreateInterface[Xhtml_tfoot_type]
	OnAfterXhtml_tfoot_typeUpdateCallback OnAfterUpdateInterface[Xhtml_tfoot_type]
	OnAfterXhtml_tfoot_typeDeleteCallback OnAfterDeleteInterface[Xhtml_tfoot_type]
	OnAfterXhtml_tfoot_typeReadCallback   OnAfterReadInterface[Xhtml_tfoot_type]

	Xhtml_th_types           map[*Xhtml_th_type]any
	Xhtml_th_types_mapString map[string]*Xhtml_th_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_th_typeCreateCallback OnAfterCreateInterface[Xhtml_th_type]
	OnAfterXhtml_th_typeUpdateCallback OnAfterUpdateInterface[Xhtml_th_type]
	OnAfterXhtml_th_typeDeleteCallback OnAfterDeleteInterface[Xhtml_th_type]
	OnAfterXhtml_th_typeReadCallback   OnAfterReadInterface[Xhtml_th_type]

	Xhtml_thead_types           map[*Xhtml_thead_type]any
	Xhtml_thead_types_mapString map[string]*Xhtml_thead_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_thead_typeCreateCallback OnAfterCreateInterface[Xhtml_thead_type]
	OnAfterXhtml_thead_typeUpdateCallback OnAfterUpdateInterface[Xhtml_thead_type]
	OnAfterXhtml_thead_typeDeleteCallback OnAfterDeleteInterface[Xhtml_thead_type]
	OnAfterXhtml_thead_typeReadCallback   OnAfterReadInterface[Xhtml_thead_type]

	Xhtml_tr_types           map[*Xhtml_tr_type]any
	Xhtml_tr_types_mapString map[string]*Xhtml_tr_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_tr_typeCreateCallback OnAfterCreateInterface[Xhtml_tr_type]
	OnAfterXhtml_tr_typeUpdateCallback OnAfterUpdateInterface[Xhtml_tr_type]
	OnAfterXhtml_tr_typeDeleteCallback OnAfterDeleteInterface[Xhtml_tr_type]
	OnAfterXhtml_tr_typeReadCallback   OnAfterReadInterface[Xhtml_tr_type]

	Xhtml_ul_types           map[*Xhtml_ul_type]any
	Xhtml_ul_types_mapString map[string]*Xhtml_ul_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_ul_typeCreateCallback OnAfterCreateInterface[Xhtml_ul_type]
	OnAfterXhtml_ul_typeUpdateCallback OnAfterUpdateInterface[Xhtml_ul_type]
	OnAfterXhtml_ul_typeDeleteCallback OnAfterDeleteInterface[Xhtml_ul_type]
	OnAfterXhtml_ul_typeReadCallback   OnAfterReadInterface[Xhtml_ul_type]

	Xhtml_var_types           map[*Xhtml_var_type]any
	Xhtml_var_types_mapString map[string]*Xhtml_var_type

	// insertion point for slice of pointers maps
	OnAfterXhtml_var_typeCreateCallback OnAfterCreateInterface[Xhtml_var_type]
	OnAfterXhtml_var_typeUpdateCallback OnAfterUpdateInterface[Xhtml_var_type]
	OnAfterXhtml_var_typeDeleteCallback OnAfterDeleteInterface[Xhtml_var_type]
	OnAfterXhtml_var_typeReadCallback   OnAfterReadInterface[Xhtml_var_type]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	ALTERNATIVE_IDOrder            uint
	ALTERNATIVE_IDMap_Staged_Order map[*ALTERNATIVE_ID]uint

	ATTRIBUTE_DEFINITION_BOOLEANOrder            uint
	ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order map[*ATTRIBUTE_DEFINITION_BOOLEAN]uint

	ATTRIBUTE_DEFINITION_DATEOrder            uint
	ATTRIBUTE_DEFINITION_DATEMap_Staged_Order map[*ATTRIBUTE_DEFINITION_DATE]uint

	ATTRIBUTE_DEFINITION_ENUMERATIONOrder            uint
	ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order map[*ATTRIBUTE_DEFINITION_ENUMERATION]uint

	ATTRIBUTE_DEFINITION_INTEGEROrder            uint
	ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order map[*ATTRIBUTE_DEFINITION_INTEGER]uint

	ATTRIBUTE_DEFINITION_REALOrder            uint
	ATTRIBUTE_DEFINITION_REALMap_Staged_Order map[*ATTRIBUTE_DEFINITION_REAL]uint

	ATTRIBUTE_DEFINITION_STRINGOrder            uint
	ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order map[*ATTRIBUTE_DEFINITION_STRING]uint

	ATTRIBUTE_DEFINITION_XHTMLOrder            uint
	ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order map[*ATTRIBUTE_DEFINITION_XHTML]uint

	ATTRIBUTE_VALUE_BOOLEANOrder            uint
	ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order map[*ATTRIBUTE_VALUE_BOOLEAN]uint

	ATTRIBUTE_VALUE_DATEOrder            uint
	ATTRIBUTE_VALUE_DATEMap_Staged_Order map[*ATTRIBUTE_VALUE_DATE]uint

	ATTRIBUTE_VALUE_ENUMERATIONOrder            uint
	ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order map[*ATTRIBUTE_VALUE_ENUMERATION]uint

	ATTRIBUTE_VALUE_INTEGEROrder            uint
	ATTRIBUTE_VALUE_INTEGERMap_Staged_Order map[*ATTRIBUTE_VALUE_INTEGER]uint

	ATTRIBUTE_VALUE_REALOrder            uint
	ATTRIBUTE_VALUE_REALMap_Staged_Order map[*ATTRIBUTE_VALUE_REAL]uint

	ATTRIBUTE_VALUE_STRINGOrder            uint
	ATTRIBUTE_VALUE_STRINGMap_Staged_Order map[*ATTRIBUTE_VALUE_STRING]uint

	ATTRIBUTE_VALUE_XHTMLOrder            uint
	ATTRIBUTE_VALUE_XHTMLMap_Staged_Order map[*ATTRIBUTE_VALUE_XHTML]uint

	AnyTypeOrder            uint
	AnyTypeMap_Staged_Order map[*AnyType]uint

	DATATYPE_DEFINITION_BOOLEANOrder            uint
	DATATYPE_DEFINITION_BOOLEANMap_Staged_Order map[*DATATYPE_DEFINITION_BOOLEAN]uint

	DATATYPE_DEFINITION_DATEOrder            uint
	DATATYPE_DEFINITION_DATEMap_Staged_Order map[*DATATYPE_DEFINITION_DATE]uint

	DATATYPE_DEFINITION_ENUMERATIONOrder            uint
	DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order map[*DATATYPE_DEFINITION_ENUMERATION]uint

	DATATYPE_DEFINITION_INTEGEROrder            uint
	DATATYPE_DEFINITION_INTEGERMap_Staged_Order map[*DATATYPE_DEFINITION_INTEGER]uint

	DATATYPE_DEFINITION_REALOrder            uint
	DATATYPE_DEFINITION_REALMap_Staged_Order map[*DATATYPE_DEFINITION_REAL]uint

	DATATYPE_DEFINITION_STRINGOrder            uint
	DATATYPE_DEFINITION_STRINGMap_Staged_Order map[*DATATYPE_DEFINITION_STRING]uint

	DATATYPE_DEFINITION_XHTMLOrder            uint
	DATATYPE_DEFINITION_XHTMLMap_Staged_Order map[*DATATYPE_DEFINITION_XHTML]uint

	EMBEDDED_VALUEOrder            uint
	EMBEDDED_VALUEMap_Staged_Order map[*EMBEDDED_VALUE]uint

	ENUM_VALUEOrder            uint
	ENUM_VALUEMap_Staged_Order map[*ENUM_VALUE]uint

	RELATION_GROUPOrder            uint
	RELATION_GROUPMap_Staged_Order map[*RELATION_GROUP]uint

	RELATION_GROUP_TYPEOrder            uint
	RELATION_GROUP_TYPEMap_Staged_Order map[*RELATION_GROUP_TYPE]uint

	REQ_IFOrder            uint
	REQ_IFMap_Staged_Order map[*REQ_IF]uint

	REQ_IF_CONTENTOrder            uint
	REQ_IF_CONTENTMap_Staged_Order map[*REQ_IF_CONTENT]uint

	REQ_IF_HEADEROrder            uint
	REQ_IF_HEADERMap_Staged_Order map[*REQ_IF_HEADER]uint

	REQ_IF_TOOL_EXTENSIONOrder            uint
	REQ_IF_TOOL_EXTENSIONMap_Staged_Order map[*REQ_IF_TOOL_EXTENSION]uint

	SPECIFICATIONOrder            uint
	SPECIFICATIONMap_Staged_Order map[*SPECIFICATION]uint

	SPECIFICATION_TYPEOrder            uint
	SPECIFICATION_TYPEMap_Staged_Order map[*SPECIFICATION_TYPE]uint

	SPEC_HIERARCHYOrder            uint
	SPEC_HIERARCHYMap_Staged_Order map[*SPEC_HIERARCHY]uint

	SPEC_OBJECTOrder            uint
	SPEC_OBJECTMap_Staged_Order map[*SPEC_OBJECT]uint

	SPEC_OBJECT_TYPEOrder            uint
	SPEC_OBJECT_TYPEMap_Staged_Order map[*SPEC_OBJECT_TYPE]uint

	SPEC_RELATIONOrder            uint
	SPEC_RELATIONMap_Staged_Order map[*SPEC_RELATION]uint

	SPEC_RELATION_TYPEOrder            uint
	SPEC_RELATION_TYPEMap_Staged_Order map[*SPEC_RELATION_TYPE]uint

	XHTML_CONTENTOrder            uint
	XHTML_CONTENTMap_Staged_Order map[*XHTML_CONTENT]uint

	Xhtml_InlPres_typeOrder            uint
	Xhtml_InlPres_typeMap_Staged_Order map[*Xhtml_InlPres_type]uint

	Xhtml_a_typeOrder            uint
	Xhtml_a_typeMap_Staged_Order map[*Xhtml_a_type]uint

	Xhtml_abbr_typeOrder            uint
	Xhtml_abbr_typeMap_Staged_Order map[*Xhtml_abbr_type]uint

	Xhtml_acronym_typeOrder            uint
	Xhtml_acronym_typeMap_Staged_Order map[*Xhtml_acronym_type]uint

	Xhtml_address_typeOrder            uint
	Xhtml_address_typeMap_Staged_Order map[*Xhtml_address_type]uint

	Xhtml_blockquote_typeOrder            uint
	Xhtml_blockquote_typeMap_Staged_Order map[*Xhtml_blockquote_type]uint

	Xhtml_br_typeOrder            uint
	Xhtml_br_typeMap_Staged_Order map[*Xhtml_br_type]uint

	Xhtml_caption_typeOrder            uint
	Xhtml_caption_typeMap_Staged_Order map[*Xhtml_caption_type]uint

	Xhtml_cite_typeOrder            uint
	Xhtml_cite_typeMap_Staged_Order map[*Xhtml_cite_type]uint

	Xhtml_code_typeOrder            uint
	Xhtml_code_typeMap_Staged_Order map[*Xhtml_code_type]uint

	Xhtml_col_typeOrder            uint
	Xhtml_col_typeMap_Staged_Order map[*Xhtml_col_type]uint

	Xhtml_colgroup_typeOrder            uint
	Xhtml_colgroup_typeMap_Staged_Order map[*Xhtml_colgroup_type]uint

	Xhtml_dd_typeOrder            uint
	Xhtml_dd_typeMap_Staged_Order map[*Xhtml_dd_type]uint

	Xhtml_dfn_typeOrder            uint
	Xhtml_dfn_typeMap_Staged_Order map[*Xhtml_dfn_type]uint

	Xhtml_div_typeOrder            uint
	Xhtml_div_typeMap_Staged_Order map[*Xhtml_div_type]uint

	Xhtml_dl_typeOrder            uint
	Xhtml_dl_typeMap_Staged_Order map[*Xhtml_dl_type]uint

	Xhtml_dt_typeOrder            uint
	Xhtml_dt_typeMap_Staged_Order map[*Xhtml_dt_type]uint

	Xhtml_edit_typeOrder            uint
	Xhtml_edit_typeMap_Staged_Order map[*Xhtml_edit_type]uint

	Xhtml_em_typeOrder            uint
	Xhtml_em_typeMap_Staged_Order map[*Xhtml_em_type]uint

	Xhtml_h1_typeOrder            uint
	Xhtml_h1_typeMap_Staged_Order map[*Xhtml_h1_type]uint

	Xhtml_h2_typeOrder            uint
	Xhtml_h2_typeMap_Staged_Order map[*Xhtml_h2_type]uint

	Xhtml_h3_typeOrder            uint
	Xhtml_h3_typeMap_Staged_Order map[*Xhtml_h3_type]uint

	Xhtml_h4_typeOrder            uint
	Xhtml_h4_typeMap_Staged_Order map[*Xhtml_h4_type]uint

	Xhtml_h5_typeOrder            uint
	Xhtml_h5_typeMap_Staged_Order map[*Xhtml_h5_type]uint

	Xhtml_h6_typeOrder            uint
	Xhtml_h6_typeMap_Staged_Order map[*Xhtml_h6_type]uint

	Xhtml_heading_typeOrder            uint
	Xhtml_heading_typeMap_Staged_Order map[*Xhtml_heading_type]uint

	Xhtml_hr_typeOrder            uint
	Xhtml_hr_typeMap_Staged_Order map[*Xhtml_hr_type]uint

	Xhtml_kbd_typeOrder            uint
	Xhtml_kbd_typeMap_Staged_Order map[*Xhtml_kbd_type]uint

	Xhtml_li_typeOrder            uint
	Xhtml_li_typeMap_Staged_Order map[*Xhtml_li_type]uint

	Xhtml_object_typeOrder            uint
	Xhtml_object_typeMap_Staged_Order map[*Xhtml_object_type]uint

	Xhtml_ol_typeOrder            uint
	Xhtml_ol_typeMap_Staged_Order map[*Xhtml_ol_type]uint

	Xhtml_p_typeOrder            uint
	Xhtml_p_typeMap_Staged_Order map[*Xhtml_p_type]uint

	Xhtml_param_typeOrder            uint
	Xhtml_param_typeMap_Staged_Order map[*Xhtml_param_type]uint

	Xhtml_pre_typeOrder            uint
	Xhtml_pre_typeMap_Staged_Order map[*Xhtml_pre_type]uint

	Xhtml_q_typeOrder            uint
	Xhtml_q_typeMap_Staged_Order map[*Xhtml_q_type]uint

	Xhtml_samp_typeOrder            uint
	Xhtml_samp_typeMap_Staged_Order map[*Xhtml_samp_type]uint

	Xhtml_span_typeOrder            uint
	Xhtml_span_typeMap_Staged_Order map[*Xhtml_span_type]uint

	Xhtml_strong_typeOrder            uint
	Xhtml_strong_typeMap_Staged_Order map[*Xhtml_strong_type]uint

	Xhtml_table_typeOrder            uint
	Xhtml_table_typeMap_Staged_Order map[*Xhtml_table_type]uint

	Xhtml_tbody_typeOrder            uint
	Xhtml_tbody_typeMap_Staged_Order map[*Xhtml_tbody_type]uint

	Xhtml_td_typeOrder            uint
	Xhtml_td_typeMap_Staged_Order map[*Xhtml_td_type]uint

	Xhtml_tfoot_typeOrder            uint
	Xhtml_tfoot_typeMap_Staged_Order map[*Xhtml_tfoot_type]uint

	Xhtml_th_typeOrder            uint
	Xhtml_th_typeMap_Staged_Order map[*Xhtml_th_type]uint

	Xhtml_thead_typeOrder            uint
	Xhtml_thead_typeMap_Staged_Order map[*Xhtml_thead_type]uint

	Xhtml_tr_typeOrder            uint
	Xhtml_tr_typeMap_Staged_Order map[*Xhtml_tr_type]uint

	Xhtml_ul_typeOrder            uint
	Xhtml_ul_typeMap_Staged_Order map[*Xhtml_ul_type]uint

	Xhtml_var_typeOrder            uint
	Xhtml_var_typeMap_Staged_Order map[*Xhtml_var_type]uint

	// end of insertion point

	NamedStructs []*NamedStruct
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]any, order map[T]uint) (res []string) {

	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	for _, instance := range orderedSet {
		res = append(res, instance.GetName())
	}

	return
}

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {

	switch namedStructName {
	// insertion point for case 
		case "ALTERNATIVE_ID":
			res = GetNamedStructInstances(stage.ALTERNATIVE_IDs, stage.ALTERNATIVE_IDMap_Staged_Order)
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			res = GetNamedStructInstances(stage.ATTRIBUTE_DEFINITION_BOOLEANs, stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order)
		case "ATTRIBUTE_DEFINITION_DATE":
			res = GetNamedStructInstances(stage.ATTRIBUTE_DEFINITION_DATEs, stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order)
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			res = GetNamedStructInstances(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs, stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order)
		case "ATTRIBUTE_DEFINITION_INTEGER":
			res = GetNamedStructInstances(stage.ATTRIBUTE_DEFINITION_INTEGERs, stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order)
		case "ATTRIBUTE_DEFINITION_REAL":
			res = GetNamedStructInstances(stage.ATTRIBUTE_DEFINITION_REALs, stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order)
		case "ATTRIBUTE_DEFINITION_STRING":
			res = GetNamedStructInstances(stage.ATTRIBUTE_DEFINITION_STRINGs, stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order)
		case "ATTRIBUTE_DEFINITION_XHTML":
			res = GetNamedStructInstances(stage.ATTRIBUTE_DEFINITION_XHTMLs, stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order)
		case "ATTRIBUTE_VALUE_BOOLEAN":
			res = GetNamedStructInstances(stage.ATTRIBUTE_VALUE_BOOLEANs, stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order)
		case "ATTRIBUTE_VALUE_DATE":
			res = GetNamedStructInstances(stage.ATTRIBUTE_VALUE_DATEs, stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order)
		case "ATTRIBUTE_VALUE_ENUMERATION":
			res = GetNamedStructInstances(stage.ATTRIBUTE_VALUE_ENUMERATIONs, stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order)
		case "ATTRIBUTE_VALUE_INTEGER":
			res = GetNamedStructInstances(stage.ATTRIBUTE_VALUE_INTEGERs, stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order)
		case "ATTRIBUTE_VALUE_REAL":
			res = GetNamedStructInstances(stage.ATTRIBUTE_VALUE_REALs, stage.ATTRIBUTE_VALUE_REALMap_Staged_Order)
		case "ATTRIBUTE_VALUE_STRING":
			res = GetNamedStructInstances(stage.ATTRIBUTE_VALUE_STRINGs, stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order)
		case "ATTRIBUTE_VALUE_XHTML":
			res = GetNamedStructInstances(stage.ATTRIBUTE_VALUE_XHTMLs, stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order)
		case "AnyType":
			res = GetNamedStructInstances(stage.AnyTypes, stage.AnyTypeMap_Staged_Order)
		case "DATATYPE_DEFINITION_BOOLEAN":
			res = GetNamedStructInstances(stage.DATATYPE_DEFINITION_BOOLEANs, stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order)
		case "DATATYPE_DEFINITION_DATE":
			res = GetNamedStructInstances(stage.DATATYPE_DEFINITION_DATEs, stage.DATATYPE_DEFINITION_DATEMap_Staged_Order)
		case "DATATYPE_DEFINITION_ENUMERATION":
			res = GetNamedStructInstances(stage.DATATYPE_DEFINITION_ENUMERATIONs, stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order)
		case "DATATYPE_DEFINITION_INTEGER":
			res = GetNamedStructInstances(stage.DATATYPE_DEFINITION_INTEGERs, stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order)
		case "DATATYPE_DEFINITION_REAL":
			res = GetNamedStructInstances(stage.DATATYPE_DEFINITION_REALs, stage.DATATYPE_DEFINITION_REALMap_Staged_Order)
		case "DATATYPE_DEFINITION_STRING":
			res = GetNamedStructInstances(stage.DATATYPE_DEFINITION_STRINGs, stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order)
		case "DATATYPE_DEFINITION_XHTML":
			res = GetNamedStructInstances(stage.DATATYPE_DEFINITION_XHTMLs, stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order)
		case "EMBEDDED_VALUE":
			res = GetNamedStructInstances(stage.EMBEDDED_VALUEs, stage.EMBEDDED_VALUEMap_Staged_Order)
		case "ENUM_VALUE":
			res = GetNamedStructInstances(stage.ENUM_VALUEs, stage.ENUM_VALUEMap_Staged_Order)
		case "RELATION_GROUP":
			res = GetNamedStructInstances(stage.RELATION_GROUPs, stage.RELATION_GROUPMap_Staged_Order)
		case "RELATION_GROUP_TYPE":
			res = GetNamedStructInstances(stage.RELATION_GROUP_TYPEs, stage.RELATION_GROUP_TYPEMap_Staged_Order)
		case "REQ_IF":
			res = GetNamedStructInstances(stage.REQ_IFs, stage.REQ_IFMap_Staged_Order)
		case "REQ_IF_CONTENT":
			res = GetNamedStructInstances(stage.REQ_IF_CONTENTs, stage.REQ_IF_CONTENTMap_Staged_Order)
		case "REQ_IF_HEADER":
			res = GetNamedStructInstances(stage.REQ_IF_HEADERs, stage.REQ_IF_HEADERMap_Staged_Order)
		case "REQ_IF_TOOL_EXTENSION":
			res = GetNamedStructInstances(stage.REQ_IF_TOOL_EXTENSIONs, stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order)
		case "SPECIFICATION":
			res = GetNamedStructInstances(stage.SPECIFICATIONs, stage.SPECIFICATIONMap_Staged_Order)
		case "SPECIFICATION_TYPE":
			res = GetNamedStructInstances(stage.SPECIFICATION_TYPEs, stage.SPECIFICATION_TYPEMap_Staged_Order)
		case "SPEC_HIERARCHY":
			res = GetNamedStructInstances(stage.SPEC_HIERARCHYs, stage.SPEC_HIERARCHYMap_Staged_Order)
		case "SPEC_OBJECT":
			res = GetNamedStructInstances(stage.SPEC_OBJECTs, stage.SPEC_OBJECTMap_Staged_Order)
		case "SPEC_OBJECT_TYPE":
			res = GetNamedStructInstances(stage.SPEC_OBJECT_TYPEs, stage.SPEC_OBJECT_TYPEMap_Staged_Order)
		case "SPEC_RELATION":
			res = GetNamedStructInstances(stage.SPEC_RELATIONs, stage.SPEC_RELATIONMap_Staged_Order)
		case "SPEC_RELATION_TYPE":
			res = GetNamedStructInstances(stage.SPEC_RELATION_TYPEs, stage.SPEC_RELATION_TYPEMap_Staged_Order)
		case "XHTML_CONTENT":
			res = GetNamedStructInstances(stage.XHTML_CONTENTs, stage.XHTML_CONTENTMap_Staged_Order)
		case "Xhtml_InlPres_type":
			res = GetNamedStructInstances(stage.Xhtml_InlPres_types, stage.Xhtml_InlPres_typeMap_Staged_Order)
		case "Xhtml_a_type":
			res = GetNamedStructInstances(stage.Xhtml_a_types, stage.Xhtml_a_typeMap_Staged_Order)
		case "Xhtml_abbr_type":
			res = GetNamedStructInstances(stage.Xhtml_abbr_types, stage.Xhtml_abbr_typeMap_Staged_Order)
		case "Xhtml_acronym_type":
			res = GetNamedStructInstances(stage.Xhtml_acronym_types, stage.Xhtml_acronym_typeMap_Staged_Order)
		case "Xhtml_address_type":
			res = GetNamedStructInstances(stage.Xhtml_address_types, stage.Xhtml_address_typeMap_Staged_Order)
		case "Xhtml_blockquote_type":
			res = GetNamedStructInstances(stage.Xhtml_blockquote_types, stage.Xhtml_blockquote_typeMap_Staged_Order)
		case "Xhtml_br_type":
			res = GetNamedStructInstances(stage.Xhtml_br_types, stage.Xhtml_br_typeMap_Staged_Order)
		case "Xhtml_caption_type":
			res = GetNamedStructInstances(stage.Xhtml_caption_types, stage.Xhtml_caption_typeMap_Staged_Order)
		case "Xhtml_cite_type":
			res = GetNamedStructInstances(stage.Xhtml_cite_types, stage.Xhtml_cite_typeMap_Staged_Order)
		case "Xhtml_code_type":
			res = GetNamedStructInstances(stage.Xhtml_code_types, stage.Xhtml_code_typeMap_Staged_Order)
		case "Xhtml_col_type":
			res = GetNamedStructInstances(stage.Xhtml_col_types, stage.Xhtml_col_typeMap_Staged_Order)
		case "Xhtml_colgroup_type":
			res = GetNamedStructInstances(stage.Xhtml_colgroup_types, stage.Xhtml_colgroup_typeMap_Staged_Order)
		case "Xhtml_dd_type":
			res = GetNamedStructInstances(stage.Xhtml_dd_types, stage.Xhtml_dd_typeMap_Staged_Order)
		case "Xhtml_dfn_type":
			res = GetNamedStructInstances(stage.Xhtml_dfn_types, stage.Xhtml_dfn_typeMap_Staged_Order)
		case "Xhtml_div_type":
			res = GetNamedStructInstances(stage.Xhtml_div_types, stage.Xhtml_div_typeMap_Staged_Order)
		case "Xhtml_dl_type":
			res = GetNamedStructInstances(stage.Xhtml_dl_types, stage.Xhtml_dl_typeMap_Staged_Order)
		case "Xhtml_dt_type":
			res = GetNamedStructInstances(stage.Xhtml_dt_types, stage.Xhtml_dt_typeMap_Staged_Order)
		case "Xhtml_edit_type":
			res = GetNamedStructInstances(stage.Xhtml_edit_types, stage.Xhtml_edit_typeMap_Staged_Order)
		case "Xhtml_em_type":
			res = GetNamedStructInstances(stage.Xhtml_em_types, stage.Xhtml_em_typeMap_Staged_Order)
		case "Xhtml_h1_type":
			res = GetNamedStructInstances(stage.Xhtml_h1_types, stage.Xhtml_h1_typeMap_Staged_Order)
		case "Xhtml_h2_type":
			res = GetNamedStructInstances(stage.Xhtml_h2_types, stage.Xhtml_h2_typeMap_Staged_Order)
		case "Xhtml_h3_type":
			res = GetNamedStructInstances(stage.Xhtml_h3_types, stage.Xhtml_h3_typeMap_Staged_Order)
		case "Xhtml_h4_type":
			res = GetNamedStructInstances(stage.Xhtml_h4_types, stage.Xhtml_h4_typeMap_Staged_Order)
		case "Xhtml_h5_type":
			res = GetNamedStructInstances(stage.Xhtml_h5_types, stage.Xhtml_h5_typeMap_Staged_Order)
		case "Xhtml_h6_type":
			res = GetNamedStructInstances(stage.Xhtml_h6_types, stage.Xhtml_h6_typeMap_Staged_Order)
		case "Xhtml_heading_type":
			res = GetNamedStructInstances(stage.Xhtml_heading_types, stage.Xhtml_heading_typeMap_Staged_Order)
		case "Xhtml_hr_type":
			res = GetNamedStructInstances(stage.Xhtml_hr_types, stage.Xhtml_hr_typeMap_Staged_Order)
		case "Xhtml_kbd_type":
			res = GetNamedStructInstances(stage.Xhtml_kbd_types, stage.Xhtml_kbd_typeMap_Staged_Order)
		case "Xhtml_li_type":
			res = GetNamedStructInstances(stage.Xhtml_li_types, stage.Xhtml_li_typeMap_Staged_Order)
		case "Xhtml_object_type":
			res = GetNamedStructInstances(stage.Xhtml_object_types, stage.Xhtml_object_typeMap_Staged_Order)
		case "Xhtml_ol_type":
			res = GetNamedStructInstances(stage.Xhtml_ol_types, stage.Xhtml_ol_typeMap_Staged_Order)
		case "Xhtml_p_type":
			res = GetNamedStructInstances(stage.Xhtml_p_types, stage.Xhtml_p_typeMap_Staged_Order)
		case "Xhtml_param_type":
			res = GetNamedStructInstances(stage.Xhtml_param_types, stage.Xhtml_param_typeMap_Staged_Order)
		case "Xhtml_pre_type":
			res = GetNamedStructInstances(stage.Xhtml_pre_types, stage.Xhtml_pre_typeMap_Staged_Order)
		case "Xhtml_q_type":
			res = GetNamedStructInstances(stage.Xhtml_q_types, stage.Xhtml_q_typeMap_Staged_Order)
		case "Xhtml_samp_type":
			res = GetNamedStructInstances(stage.Xhtml_samp_types, stage.Xhtml_samp_typeMap_Staged_Order)
		case "Xhtml_span_type":
			res = GetNamedStructInstances(stage.Xhtml_span_types, stage.Xhtml_span_typeMap_Staged_Order)
		case "Xhtml_strong_type":
			res = GetNamedStructInstances(stage.Xhtml_strong_types, stage.Xhtml_strong_typeMap_Staged_Order)
		case "Xhtml_table_type":
			res = GetNamedStructInstances(stage.Xhtml_table_types, stage.Xhtml_table_typeMap_Staged_Order)
		case "Xhtml_tbody_type":
			res = GetNamedStructInstances(stage.Xhtml_tbody_types, stage.Xhtml_tbody_typeMap_Staged_Order)
		case "Xhtml_td_type":
			res = GetNamedStructInstances(stage.Xhtml_td_types, stage.Xhtml_td_typeMap_Staged_Order)
		case "Xhtml_tfoot_type":
			res = GetNamedStructInstances(stage.Xhtml_tfoot_types, stage.Xhtml_tfoot_typeMap_Staged_Order)
		case "Xhtml_th_type":
			res = GetNamedStructInstances(stage.Xhtml_th_types, stage.Xhtml_th_typeMap_Staged_Order)
		case "Xhtml_thead_type":
			res = GetNamedStructInstances(stage.Xhtml_thead_types, stage.Xhtml_thead_typeMap_Staged_Order)
		case "Xhtml_tr_type":
			res = GetNamedStructInstances(stage.Xhtml_tr_types, stage.Xhtml_tr_typeMap_Staged_Order)
		case "Xhtml_ul_type":
			res = GetNamedStructInstances(stage.Xhtml_ul_types, stage.Xhtml_ul_typeMap_Staged_Order)
		case "Xhtml_var_type":
			res = GetNamedStructInstances(stage.Xhtml_var_types, stage.Xhtml_var_typeMap_Staged_Order)
	}

	return
}


type NamedStruct struct {
	name string
}

func (namedStruct *NamedStruct) GetName() string {
	return namedStruct.name
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gongreqif/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return gongreqif_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return gongreqif_go.GoDiagramsDir
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *Stage,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID)
	CheckoutALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID)
	CommitATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN)
	CheckoutATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN)
	CommitATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE)
	CheckoutATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE)
	CommitATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION)
	CheckoutATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION)
	CommitATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER)
	CheckoutATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER)
	CommitATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL)
	CheckoutATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL)
	CommitATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING)
	CheckoutATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING)
	CommitATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML)
	CheckoutATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML)
	CommitATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN)
	CheckoutATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN)
	CommitATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE)
	CheckoutATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE)
	CommitATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION)
	CheckoutATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION)
	CommitATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER)
	CheckoutATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER)
	CommitATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL)
	CheckoutATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL)
	CommitATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING)
	CheckoutATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING)
	CommitATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML)
	CheckoutATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML)
	CommitAnyType(anytype *AnyType)
	CheckoutAnyType(anytype *AnyType)
	CommitDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN)
	CheckoutDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN)
	CommitDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE)
	CheckoutDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE)
	CommitDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION)
	CheckoutDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION)
	CommitDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER)
	CheckoutDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER)
	CommitDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL)
	CheckoutDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL)
	CommitDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING)
	CheckoutDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING)
	CommitDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML)
	CheckoutDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML)
	CommitEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE)
	CheckoutEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE)
	CommitENUM_VALUE(enum_value *ENUM_VALUE)
	CheckoutENUM_VALUE(enum_value *ENUM_VALUE)
	CommitRELATION_GROUP(relation_group *RELATION_GROUP)
	CheckoutRELATION_GROUP(relation_group *RELATION_GROUP)
	CommitRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE)
	CheckoutRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE)
	CommitREQ_IF(req_if *REQ_IF)
	CheckoutREQ_IF(req_if *REQ_IF)
	CommitREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT)
	CheckoutREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT)
	CommitREQ_IF_HEADER(req_if_header *REQ_IF_HEADER)
	CheckoutREQ_IF_HEADER(req_if_header *REQ_IF_HEADER)
	CommitREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION)
	CheckoutREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION)
	CommitSPECIFICATION(specification *SPECIFICATION)
	CheckoutSPECIFICATION(specification *SPECIFICATION)
	CommitSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE)
	CheckoutSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE)
	CommitSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY)
	CheckoutSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY)
	CommitSPEC_OBJECT(spec_object *SPEC_OBJECT)
	CheckoutSPEC_OBJECT(spec_object *SPEC_OBJECT)
	CommitSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE)
	CheckoutSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE)
	CommitSPEC_RELATION(spec_relation *SPEC_RELATION)
	CheckoutSPEC_RELATION(spec_relation *SPEC_RELATION)
	CommitSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE)
	CheckoutSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE)
	CommitXHTML_CONTENT(xhtml_content *XHTML_CONTENT)
	CheckoutXHTML_CONTENT(xhtml_content *XHTML_CONTENT)
	CommitXhtml_InlPres_type(xhtml_inlpres_type *Xhtml_InlPres_type)
	CheckoutXhtml_InlPres_type(xhtml_inlpres_type *Xhtml_InlPres_type)
	CommitXhtml_a_type(xhtml_a_type *Xhtml_a_type)
	CheckoutXhtml_a_type(xhtml_a_type *Xhtml_a_type)
	CommitXhtml_abbr_type(xhtml_abbr_type *Xhtml_abbr_type)
	CheckoutXhtml_abbr_type(xhtml_abbr_type *Xhtml_abbr_type)
	CommitXhtml_acronym_type(xhtml_acronym_type *Xhtml_acronym_type)
	CheckoutXhtml_acronym_type(xhtml_acronym_type *Xhtml_acronym_type)
	CommitXhtml_address_type(xhtml_address_type *Xhtml_address_type)
	CheckoutXhtml_address_type(xhtml_address_type *Xhtml_address_type)
	CommitXhtml_blockquote_type(xhtml_blockquote_type *Xhtml_blockquote_type)
	CheckoutXhtml_blockquote_type(xhtml_blockquote_type *Xhtml_blockquote_type)
	CommitXhtml_br_type(xhtml_br_type *Xhtml_br_type)
	CheckoutXhtml_br_type(xhtml_br_type *Xhtml_br_type)
	CommitXhtml_caption_type(xhtml_caption_type *Xhtml_caption_type)
	CheckoutXhtml_caption_type(xhtml_caption_type *Xhtml_caption_type)
	CommitXhtml_cite_type(xhtml_cite_type *Xhtml_cite_type)
	CheckoutXhtml_cite_type(xhtml_cite_type *Xhtml_cite_type)
	CommitXhtml_code_type(xhtml_code_type *Xhtml_code_type)
	CheckoutXhtml_code_type(xhtml_code_type *Xhtml_code_type)
	CommitXhtml_col_type(xhtml_col_type *Xhtml_col_type)
	CheckoutXhtml_col_type(xhtml_col_type *Xhtml_col_type)
	CommitXhtml_colgroup_type(xhtml_colgroup_type *Xhtml_colgroup_type)
	CheckoutXhtml_colgroup_type(xhtml_colgroup_type *Xhtml_colgroup_type)
	CommitXhtml_dd_type(xhtml_dd_type *Xhtml_dd_type)
	CheckoutXhtml_dd_type(xhtml_dd_type *Xhtml_dd_type)
	CommitXhtml_dfn_type(xhtml_dfn_type *Xhtml_dfn_type)
	CheckoutXhtml_dfn_type(xhtml_dfn_type *Xhtml_dfn_type)
	CommitXhtml_div_type(xhtml_div_type *Xhtml_div_type)
	CheckoutXhtml_div_type(xhtml_div_type *Xhtml_div_type)
	CommitXhtml_dl_type(xhtml_dl_type *Xhtml_dl_type)
	CheckoutXhtml_dl_type(xhtml_dl_type *Xhtml_dl_type)
	CommitXhtml_dt_type(xhtml_dt_type *Xhtml_dt_type)
	CheckoutXhtml_dt_type(xhtml_dt_type *Xhtml_dt_type)
	CommitXhtml_edit_type(xhtml_edit_type *Xhtml_edit_type)
	CheckoutXhtml_edit_type(xhtml_edit_type *Xhtml_edit_type)
	CommitXhtml_em_type(xhtml_em_type *Xhtml_em_type)
	CheckoutXhtml_em_type(xhtml_em_type *Xhtml_em_type)
	CommitXhtml_h1_type(xhtml_h1_type *Xhtml_h1_type)
	CheckoutXhtml_h1_type(xhtml_h1_type *Xhtml_h1_type)
	CommitXhtml_h2_type(xhtml_h2_type *Xhtml_h2_type)
	CheckoutXhtml_h2_type(xhtml_h2_type *Xhtml_h2_type)
	CommitXhtml_h3_type(xhtml_h3_type *Xhtml_h3_type)
	CheckoutXhtml_h3_type(xhtml_h3_type *Xhtml_h3_type)
	CommitXhtml_h4_type(xhtml_h4_type *Xhtml_h4_type)
	CheckoutXhtml_h4_type(xhtml_h4_type *Xhtml_h4_type)
	CommitXhtml_h5_type(xhtml_h5_type *Xhtml_h5_type)
	CheckoutXhtml_h5_type(xhtml_h5_type *Xhtml_h5_type)
	CommitXhtml_h6_type(xhtml_h6_type *Xhtml_h6_type)
	CheckoutXhtml_h6_type(xhtml_h6_type *Xhtml_h6_type)
	CommitXhtml_heading_type(xhtml_heading_type *Xhtml_heading_type)
	CheckoutXhtml_heading_type(xhtml_heading_type *Xhtml_heading_type)
	CommitXhtml_hr_type(xhtml_hr_type *Xhtml_hr_type)
	CheckoutXhtml_hr_type(xhtml_hr_type *Xhtml_hr_type)
	CommitXhtml_kbd_type(xhtml_kbd_type *Xhtml_kbd_type)
	CheckoutXhtml_kbd_type(xhtml_kbd_type *Xhtml_kbd_type)
	CommitXhtml_li_type(xhtml_li_type *Xhtml_li_type)
	CheckoutXhtml_li_type(xhtml_li_type *Xhtml_li_type)
	CommitXhtml_object_type(xhtml_object_type *Xhtml_object_type)
	CheckoutXhtml_object_type(xhtml_object_type *Xhtml_object_type)
	CommitXhtml_ol_type(xhtml_ol_type *Xhtml_ol_type)
	CheckoutXhtml_ol_type(xhtml_ol_type *Xhtml_ol_type)
	CommitXhtml_p_type(xhtml_p_type *Xhtml_p_type)
	CheckoutXhtml_p_type(xhtml_p_type *Xhtml_p_type)
	CommitXhtml_param_type(xhtml_param_type *Xhtml_param_type)
	CheckoutXhtml_param_type(xhtml_param_type *Xhtml_param_type)
	CommitXhtml_pre_type(xhtml_pre_type *Xhtml_pre_type)
	CheckoutXhtml_pre_type(xhtml_pre_type *Xhtml_pre_type)
	CommitXhtml_q_type(xhtml_q_type *Xhtml_q_type)
	CheckoutXhtml_q_type(xhtml_q_type *Xhtml_q_type)
	CommitXhtml_samp_type(xhtml_samp_type *Xhtml_samp_type)
	CheckoutXhtml_samp_type(xhtml_samp_type *Xhtml_samp_type)
	CommitXhtml_span_type(xhtml_span_type *Xhtml_span_type)
	CheckoutXhtml_span_type(xhtml_span_type *Xhtml_span_type)
	CommitXhtml_strong_type(xhtml_strong_type *Xhtml_strong_type)
	CheckoutXhtml_strong_type(xhtml_strong_type *Xhtml_strong_type)
	CommitXhtml_table_type(xhtml_table_type *Xhtml_table_type)
	CheckoutXhtml_table_type(xhtml_table_type *Xhtml_table_type)
	CommitXhtml_tbody_type(xhtml_tbody_type *Xhtml_tbody_type)
	CheckoutXhtml_tbody_type(xhtml_tbody_type *Xhtml_tbody_type)
	CommitXhtml_td_type(xhtml_td_type *Xhtml_td_type)
	CheckoutXhtml_td_type(xhtml_td_type *Xhtml_td_type)
	CommitXhtml_tfoot_type(xhtml_tfoot_type *Xhtml_tfoot_type)
	CheckoutXhtml_tfoot_type(xhtml_tfoot_type *Xhtml_tfoot_type)
	CommitXhtml_th_type(xhtml_th_type *Xhtml_th_type)
	CheckoutXhtml_th_type(xhtml_th_type *Xhtml_th_type)
	CommitXhtml_thead_type(xhtml_thead_type *Xhtml_thead_type)
	CheckoutXhtml_thead_type(xhtml_thead_type *Xhtml_thead_type)
	CommitXhtml_tr_type(xhtml_tr_type *Xhtml_tr_type)
	CheckoutXhtml_tr_type(xhtml_tr_type *Xhtml_tr_type)
	CommitXhtml_ul_type(xhtml_ul_type *Xhtml_ul_type)
	CheckoutXhtml_ul_type(xhtml_ul_type *Xhtml_ul_type)
	CommitXhtml_var_type(xhtml_var_type *Xhtml_var_type)
	CheckoutXhtml_var_type(xhtml_var_type *Xhtml_var_type)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		ALTERNATIVE_IDs:           make(map[*ALTERNATIVE_ID]any),
		ALTERNATIVE_IDs_mapString: make(map[string]*ALTERNATIVE_ID),

		ATTRIBUTE_DEFINITION_BOOLEANs:           make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]any),
		ATTRIBUTE_DEFINITION_BOOLEANs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_BOOLEAN),

		ATTRIBUTE_DEFINITION_DATEs:           make(map[*ATTRIBUTE_DEFINITION_DATE]any),
		ATTRIBUTE_DEFINITION_DATEs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_DATE),

		ATTRIBUTE_DEFINITION_ENUMERATIONs:           make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]any),
		ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_ENUMERATION),

		ATTRIBUTE_DEFINITION_INTEGERs:           make(map[*ATTRIBUTE_DEFINITION_INTEGER]any),
		ATTRIBUTE_DEFINITION_INTEGERs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_INTEGER),

		ATTRIBUTE_DEFINITION_REALs:           make(map[*ATTRIBUTE_DEFINITION_REAL]any),
		ATTRIBUTE_DEFINITION_REALs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_REAL),

		ATTRIBUTE_DEFINITION_STRINGs:           make(map[*ATTRIBUTE_DEFINITION_STRING]any),
		ATTRIBUTE_DEFINITION_STRINGs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_STRING),

		ATTRIBUTE_DEFINITION_XHTMLs:           make(map[*ATTRIBUTE_DEFINITION_XHTML]any),
		ATTRIBUTE_DEFINITION_XHTMLs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_XHTML),

		ATTRIBUTE_VALUE_BOOLEANs:           make(map[*ATTRIBUTE_VALUE_BOOLEAN]any),
		ATTRIBUTE_VALUE_BOOLEANs_mapString: make(map[string]*ATTRIBUTE_VALUE_BOOLEAN),

		ATTRIBUTE_VALUE_DATEs:           make(map[*ATTRIBUTE_VALUE_DATE]any),
		ATTRIBUTE_VALUE_DATEs_mapString: make(map[string]*ATTRIBUTE_VALUE_DATE),

		ATTRIBUTE_VALUE_ENUMERATIONs:           make(map[*ATTRIBUTE_VALUE_ENUMERATION]any),
		ATTRIBUTE_VALUE_ENUMERATIONs_mapString: make(map[string]*ATTRIBUTE_VALUE_ENUMERATION),

		ATTRIBUTE_VALUE_INTEGERs:           make(map[*ATTRIBUTE_VALUE_INTEGER]any),
		ATTRIBUTE_VALUE_INTEGERs_mapString: make(map[string]*ATTRIBUTE_VALUE_INTEGER),

		ATTRIBUTE_VALUE_REALs:           make(map[*ATTRIBUTE_VALUE_REAL]any),
		ATTRIBUTE_VALUE_REALs_mapString: make(map[string]*ATTRIBUTE_VALUE_REAL),

		ATTRIBUTE_VALUE_STRINGs:           make(map[*ATTRIBUTE_VALUE_STRING]any),
		ATTRIBUTE_VALUE_STRINGs_mapString: make(map[string]*ATTRIBUTE_VALUE_STRING),

		ATTRIBUTE_VALUE_XHTMLs:           make(map[*ATTRIBUTE_VALUE_XHTML]any),
		ATTRIBUTE_VALUE_XHTMLs_mapString: make(map[string]*ATTRIBUTE_VALUE_XHTML),

		AnyTypes:           make(map[*AnyType]any),
		AnyTypes_mapString: make(map[string]*AnyType),

		DATATYPE_DEFINITION_BOOLEANs:           make(map[*DATATYPE_DEFINITION_BOOLEAN]any),
		DATATYPE_DEFINITION_BOOLEANs_mapString: make(map[string]*DATATYPE_DEFINITION_BOOLEAN),

		DATATYPE_DEFINITION_DATEs:           make(map[*DATATYPE_DEFINITION_DATE]any),
		DATATYPE_DEFINITION_DATEs_mapString: make(map[string]*DATATYPE_DEFINITION_DATE),

		DATATYPE_DEFINITION_ENUMERATIONs:           make(map[*DATATYPE_DEFINITION_ENUMERATION]any),
		DATATYPE_DEFINITION_ENUMERATIONs_mapString: make(map[string]*DATATYPE_DEFINITION_ENUMERATION),

		DATATYPE_DEFINITION_INTEGERs:           make(map[*DATATYPE_DEFINITION_INTEGER]any),
		DATATYPE_DEFINITION_INTEGERs_mapString: make(map[string]*DATATYPE_DEFINITION_INTEGER),

		DATATYPE_DEFINITION_REALs:           make(map[*DATATYPE_DEFINITION_REAL]any),
		DATATYPE_DEFINITION_REALs_mapString: make(map[string]*DATATYPE_DEFINITION_REAL),

		DATATYPE_DEFINITION_STRINGs:           make(map[*DATATYPE_DEFINITION_STRING]any),
		DATATYPE_DEFINITION_STRINGs_mapString: make(map[string]*DATATYPE_DEFINITION_STRING),

		DATATYPE_DEFINITION_XHTMLs:           make(map[*DATATYPE_DEFINITION_XHTML]any),
		DATATYPE_DEFINITION_XHTMLs_mapString: make(map[string]*DATATYPE_DEFINITION_XHTML),

		EMBEDDED_VALUEs:           make(map[*EMBEDDED_VALUE]any),
		EMBEDDED_VALUEs_mapString: make(map[string]*EMBEDDED_VALUE),

		ENUM_VALUEs:           make(map[*ENUM_VALUE]any),
		ENUM_VALUEs_mapString: make(map[string]*ENUM_VALUE),

		RELATION_GROUPs:           make(map[*RELATION_GROUP]any),
		RELATION_GROUPs_mapString: make(map[string]*RELATION_GROUP),

		RELATION_GROUP_TYPEs:           make(map[*RELATION_GROUP_TYPE]any),
		RELATION_GROUP_TYPEs_mapString: make(map[string]*RELATION_GROUP_TYPE),

		REQ_IFs:           make(map[*REQ_IF]any),
		REQ_IFs_mapString: make(map[string]*REQ_IF),

		REQ_IF_CONTENTs:           make(map[*REQ_IF_CONTENT]any),
		REQ_IF_CONTENTs_mapString: make(map[string]*REQ_IF_CONTENT),

		REQ_IF_HEADERs:           make(map[*REQ_IF_HEADER]any),
		REQ_IF_HEADERs_mapString: make(map[string]*REQ_IF_HEADER),

		REQ_IF_TOOL_EXTENSIONs:           make(map[*REQ_IF_TOOL_EXTENSION]any),
		REQ_IF_TOOL_EXTENSIONs_mapString: make(map[string]*REQ_IF_TOOL_EXTENSION),

		SPECIFICATIONs:           make(map[*SPECIFICATION]any),
		SPECIFICATIONs_mapString: make(map[string]*SPECIFICATION),

		SPECIFICATION_TYPEs:           make(map[*SPECIFICATION_TYPE]any),
		SPECIFICATION_TYPEs_mapString: make(map[string]*SPECIFICATION_TYPE),

		SPEC_HIERARCHYs:           make(map[*SPEC_HIERARCHY]any),
		SPEC_HIERARCHYs_mapString: make(map[string]*SPEC_HIERARCHY),

		SPEC_OBJECTs:           make(map[*SPEC_OBJECT]any),
		SPEC_OBJECTs_mapString: make(map[string]*SPEC_OBJECT),

		SPEC_OBJECT_TYPEs:           make(map[*SPEC_OBJECT_TYPE]any),
		SPEC_OBJECT_TYPEs_mapString: make(map[string]*SPEC_OBJECT_TYPE),

		SPEC_RELATIONs:           make(map[*SPEC_RELATION]any),
		SPEC_RELATIONs_mapString: make(map[string]*SPEC_RELATION),

		SPEC_RELATION_TYPEs:           make(map[*SPEC_RELATION_TYPE]any),
		SPEC_RELATION_TYPEs_mapString: make(map[string]*SPEC_RELATION_TYPE),

		XHTML_CONTENTs:           make(map[*XHTML_CONTENT]any),
		XHTML_CONTENTs_mapString: make(map[string]*XHTML_CONTENT),

		Xhtml_InlPres_types:           make(map[*Xhtml_InlPres_type]any),
		Xhtml_InlPres_types_mapString: make(map[string]*Xhtml_InlPres_type),

		Xhtml_a_types:           make(map[*Xhtml_a_type]any),
		Xhtml_a_types_mapString: make(map[string]*Xhtml_a_type),

		Xhtml_abbr_types:           make(map[*Xhtml_abbr_type]any),
		Xhtml_abbr_types_mapString: make(map[string]*Xhtml_abbr_type),

		Xhtml_acronym_types:           make(map[*Xhtml_acronym_type]any),
		Xhtml_acronym_types_mapString: make(map[string]*Xhtml_acronym_type),

		Xhtml_address_types:           make(map[*Xhtml_address_type]any),
		Xhtml_address_types_mapString: make(map[string]*Xhtml_address_type),

		Xhtml_blockquote_types:           make(map[*Xhtml_blockquote_type]any),
		Xhtml_blockquote_types_mapString: make(map[string]*Xhtml_blockquote_type),

		Xhtml_br_types:           make(map[*Xhtml_br_type]any),
		Xhtml_br_types_mapString: make(map[string]*Xhtml_br_type),

		Xhtml_caption_types:           make(map[*Xhtml_caption_type]any),
		Xhtml_caption_types_mapString: make(map[string]*Xhtml_caption_type),

		Xhtml_cite_types:           make(map[*Xhtml_cite_type]any),
		Xhtml_cite_types_mapString: make(map[string]*Xhtml_cite_type),

		Xhtml_code_types:           make(map[*Xhtml_code_type]any),
		Xhtml_code_types_mapString: make(map[string]*Xhtml_code_type),

		Xhtml_col_types:           make(map[*Xhtml_col_type]any),
		Xhtml_col_types_mapString: make(map[string]*Xhtml_col_type),

		Xhtml_colgroup_types:           make(map[*Xhtml_colgroup_type]any),
		Xhtml_colgroup_types_mapString: make(map[string]*Xhtml_colgroup_type),

		Xhtml_dd_types:           make(map[*Xhtml_dd_type]any),
		Xhtml_dd_types_mapString: make(map[string]*Xhtml_dd_type),

		Xhtml_dfn_types:           make(map[*Xhtml_dfn_type]any),
		Xhtml_dfn_types_mapString: make(map[string]*Xhtml_dfn_type),

		Xhtml_div_types:           make(map[*Xhtml_div_type]any),
		Xhtml_div_types_mapString: make(map[string]*Xhtml_div_type),

		Xhtml_dl_types:           make(map[*Xhtml_dl_type]any),
		Xhtml_dl_types_mapString: make(map[string]*Xhtml_dl_type),

		Xhtml_dt_types:           make(map[*Xhtml_dt_type]any),
		Xhtml_dt_types_mapString: make(map[string]*Xhtml_dt_type),

		Xhtml_edit_types:           make(map[*Xhtml_edit_type]any),
		Xhtml_edit_types_mapString: make(map[string]*Xhtml_edit_type),

		Xhtml_em_types:           make(map[*Xhtml_em_type]any),
		Xhtml_em_types_mapString: make(map[string]*Xhtml_em_type),

		Xhtml_h1_types:           make(map[*Xhtml_h1_type]any),
		Xhtml_h1_types_mapString: make(map[string]*Xhtml_h1_type),

		Xhtml_h2_types:           make(map[*Xhtml_h2_type]any),
		Xhtml_h2_types_mapString: make(map[string]*Xhtml_h2_type),

		Xhtml_h3_types:           make(map[*Xhtml_h3_type]any),
		Xhtml_h3_types_mapString: make(map[string]*Xhtml_h3_type),

		Xhtml_h4_types:           make(map[*Xhtml_h4_type]any),
		Xhtml_h4_types_mapString: make(map[string]*Xhtml_h4_type),

		Xhtml_h5_types:           make(map[*Xhtml_h5_type]any),
		Xhtml_h5_types_mapString: make(map[string]*Xhtml_h5_type),

		Xhtml_h6_types:           make(map[*Xhtml_h6_type]any),
		Xhtml_h6_types_mapString: make(map[string]*Xhtml_h6_type),

		Xhtml_heading_types:           make(map[*Xhtml_heading_type]any),
		Xhtml_heading_types_mapString: make(map[string]*Xhtml_heading_type),

		Xhtml_hr_types:           make(map[*Xhtml_hr_type]any),
		Xhtml_hr_types_mapString: make(map[string]*Xhtml_hr_type),

		Xhtml_kbd_types:           make(map[*Xhtml_kbd_type]any),
		Xhtml_kbd_types_mapString: make(map[string]*Xhtml_kbd_type),

		Xhtml_li_types:           make(map[*Xhtml_li_type]any),
		Xhtml_li_types_mapString: make(map[string]*Xhtml_li_type),

		Xhtml_object_types:           make(map[*Xhtml_object_type]any),
		Xhtml_object_types_mapString: make(map[string]*Xhtml_object_type),

		Xhtml_ol_types:           make(map[*Xhtml_ol_type]any),
		Xhtml_ol_types_mapString: make(map[string]*Xhtml_ol_type),

		Xhtml_p_types:           make(map[*Xhtml_p_type]any),
		Xhtml_p_types_mapString: make(map[string]*Xhtml_p_type),

		Xhtml_param_types:           make(map[*Xhtml_param_type]any),
		Xhtml_param_types_mapString: make(map[string]*Xhtml_param_type),

		Xhtml_pre_types:           make(map[*Xhtml_pre_type]any),
		Xhtml_pre_types_mapString: make(map[string]*Xhtml_pre_type),

		Xhtml_q_types:           make(map[*Xhtml_q_type]any),
		Xhtml_q_types_mapString: make(map[string]*Xhtml_q_type),

		Xhtml_samp_types:           make(map[*Xhtml_samp_type]any),
		Xhtml_samp_types_mapString: make(map[string]*Xhtml_samp_type),

		Xhtml_span_types:           make(map[*Xhtml_span_type]any),
		Xhtml_span_types_mapString: make(map[string]*Xhtml_span_type),

		Xhtml_strong_types:           make(map[*Xhtml_strong_type]any),
		Xhtml_strong_types_mapString: make(map[string]*Xhtml_strong_type),

		Xhtml_table_types:           make(map[*Xhtml_table_type]any),
		Xhtml_table_types_mapString: make(map[string]*Xhtml_table_type),

		Xhtml_tbody_types:           make(map[*Xhtml_tbody_type]any),
		Xhtml_tbody_types_mapString: make(map[string]*Xhtml_tbody_type),

		Xhtml_td_types:           make(map[*Xhtml_td_type]any),
		Xhtml_td_types_mapString: make(map[string]*Xhtml_td_type),

		Xhtml_tfoot_types:           make(map[*Xhtml_tfoot_type]any),
		Xhtml_tfoot_types_mapString: make(map[string]*Xhtml_tfoot_type),

		Xhtml_th_types:           make(map[*Xhtml_th_type]any),
		Xhtml_th_types_mapString: make(map[string]*Xhtml_th_type),

		Xhtml_thead_types:           make(map[*Xhtml_thead_type]any),
		Xhtml_thead_types_mapString: make(map[string]*Xhtml_thead_type),

		Xhtml_tr_types:           make(map[*Xhtml_tr_type]any),
		Xhtml_tr_types_mapString: make(map[string]*Xhtml_tr_type),

		Xhtml_ul_types:           make(map[*Xhtml_ul_type]any),
		Xhtml_ul_types_mapString: make(map[string]*Xhtml_ul_type),

		Xhtml_var_types:           make(map[*Xhtml_var_type]any),
		Xhtml_var_types_mapString: make(map[string]*Xhtml_var_type),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ALTERNATIVE_IDMap_Staged_Order: make(map[*ALTERNATIVE_ID]uint),

		ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order: make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]uint),

		ATTRIBUTE_DEFINITION_DATEMap_Staged_Order: make(map[*ATTRIBUTE_DEFINITION_DATE]uint),

		ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order: make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]uint),

		ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order: make(map[*ATTRIBUTE_DEFINITION_INTEGER]uint),

		ATTRIBUTE_DEFINITION_REALMap_Staged_Order: make(map[*ATTRIBUTE_DEFINITION_REAL]uint),

		ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order: make(map[*ATTRIBUTE_DEFINITION_STRING]uint),

		ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order: make(map[*ATTRIBUTE_DEFINITION_XHTML]uint),

		ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order: make(map[*ATTRIBUTE_VALUE_BOOLEAN]uint),

		ATTRIBUTE_VALUE_DATEMap_Staged_Order: make(map[*ATTRIBUTE_VALUE_DATE]uint),

		ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order: make(map[*ATTRIBUTE_VALUE_ENUMERATION]uint),

		ATTRIBUTE_VALUE_INTEGERMap_Staged_Order: make(map[*ATTRIBUTE_VALUE_INTEGER]uint),

		ATTRIBUTE_VALUE_REALMap_Staged_Order: make(map[*ATTRIBUTE_VALUE_REAL]uint),

		ATTRIBUTE_VALUE_STRINGMap_Staged_Order: make(map[*ATTRIBUTE_VALUE_STRING]uint),

		ATTRIBUTE_VALUE_XHTMLMap_Staged_Order: make(map[*ATTRIBUTE_VALUE_XHTML]uint),

		AnyTypeMap_Staged_Order: make(map[*AnyType]uint),

		DATATYPE_DEFINITION_BOOLEANMap_Staged_Order: make(map[*DATATYPE_DEFINITION_BOOLEAN]uint),

		DATATYPE_DEFINITION_DATEMap_Staged_Order: make(map[*DATATYPE_DEFINITION_DATE]uint),

		DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order: make(map[*DATATYPE_DEFINITION_ENUMERATION]uint),

		DATATYPE_DEFINITION_INTEGERMap_Staged_Order: make(map[*DATATYPE_DEFINITION_INTEGER]uint),

		DATATYPE_DEFINITION_REALMap_Staged_Order: make(map[*DATATYPE_DEFINITION_REAL]uint),

		DATATYPE_DEFINITION_STRINGMap_Staged_Order: make(map[*DATATYPE_DEFINITION_STRING]uint),

		DATATYPE_DEFINITION_XHTMLMap_Staged_Order: make(map[*DATATYPE_DEFINITION_XHTML]uint),

		EMBEDDED_VALUEMap_Staged_Order: make(map[*EMBEDDED_VALUE]uint),

		ENUM_VALUEMap_Staged_Order: make(map[*ENUM_VALUE]uint),

		RELATION_GROUPMap_Staged_Order: make(map[*RELATION_GROUP]uint),

		RELATION_GROUP_TYPEMap_Staged_Order: make(map[*RELATION_GROUP_TYPE]uint),

		REQ_IFMap_Staged_Order: make(map[*REQ_IF]uint),

		REQ_IF_CONTENTMap_Staged_Order: make(map[*REQ_IF_CONTENT]uint),

		REQ_IF_HEADERMap_Staged_Order: make(map[*REQ_IF_HEADER]uint),

		REQ_IF_TOOL_EXTENSIONMap_Staged_Order: make(map[*REQ_IF_TOOL_EXTENSION]uint),

		SPECIFICATIONMap_Staged_Order: make(map[*SPECIFICATION]uint),

		SPECIFICATION_TYPEMap_Staged_Order: make(map[*SPECIFICATION_TYPE]uint),

		SPEC_HIERARCHYMap_Staged_Order: make(map[*SPEC_HIERARCHY]uint),

		SPEC_OBJECTMap_Staged_Order: make(map[*SPEC_OBJECT]uint),

		SPEC_OBJECT_TYPEMap_Staged_Order: make(map[*SPEC_OBJECT_TYPE]uint),

		SPEC_RELATIONMap_Staged_Order: make(map[*SPEC_RELATION]uint),

		SPEC_RELATION_TYPEMap_Staged_Order: make(map[*SPEC_RELATION_TYPE]uint),

		XHTML_CONTENTMap_Staged_Order: make(map[*XHTML_CONTENT]uint),

		Xhtml_InlPres_typeMap_Staged_Order: make(map[*Xhtml_InlPres_type]uint),

		Xhtml_a_typeMap_Staged_Order: make(map[*Xhtml_a_type]uint),

		Xhtml_abbr_typeMap_Staged_Order: make(map[*Xhtml_abbr_type]uint),

		Xhtml_acronym_typeMap_Staged_Order: make(map[*Xhtml_acronym_type]uint),

		Xhtml_address_typeMap_Staged_Order: make(map[*Xhtml_address_type]uint),

		Xhtml_blockquote_typeMap_Staged_Order: make(map[*Xhtml_blockquote_type]uint),

		Xhtml_br_typeMap_Staged_Order: make(map[*Xhtml_br_type]uint),

		Xhtml_caption_typeMap_Staged_Order: make(map[*Xhtml_caption_type]uint),

		Xhtml_cite_typeMap_Staged_Order: make(map[*Xhtml_cite_type]uint),

		Xhtml_code_typeMap_Staged_Order: make(map[*Xhtml_code_type]uint),

		Xhtml_col_typeMap_Staged_Order: make(map[*Xhtml_col_type]uint),

		Xhtml_colgroup_typeMap_Staged_Order: make(map[*Xhtml_colgroup_type]uint),

		Xhtml_dd_typeMap_Staged_Order: make(map[*Xhtml_dd_type]uint),

		Xhtml_dfn_typeMap_Staged_Order: make(map[*Xhtml_dfn_type]uint),

		Xhtml_div_typeMap_Staged_Order: make(map[*Xhtml_div_type]uint),

		Xhtml_dl_typeMap_Staged_Order: make(map[*Xhtml_dl_type]uint),

		Xhtml_dt_typeMap_Staged_Order: make(map[*Xhtml_dt_type]uint),

		Xhtml_edit_typeMap_Staged_Order: make(map[*Xhtml_edit_type]uint),

		Xhtml_em_typeMap_Staged_Order: make(map[*Xhtml_em_type]uint),

		Xhtml_h1_typeMap_Staged_Order: make(map[*Xhtml_h1_type]uint),

		Xhtml_h2_typeMap_Staged_Order: make(map[*Xhtml_h2_type]uint),

		Xhtml_h3_typeMap_Staged_Order: make(map[*Xhtml_h3_type]uint),

		Xhtml_h4_typeMap_Staged_Order: make(map[*Xhtml_h4_type]uint),

		Xhtml_h5_typeMap_Staged_Order: make(map[*Xhtml_h5_type]uint),

		Xhtml_h6_typeMap_Staged_Order: make(map[*Xhtml_h6_type]uint),

		Xhtml_heading_typeMap_Staged_Order: make(map[*Xhtml_heading_type]uint),

		Xhtml_hr_typeMap_Staged_Order: make(map[*Xhtml_hr_type]uint),

		Xhtml_kbd_typeMap_Staged_Order: make(map[*Xhtml_kbd_type]uint),

		Xhtml_li_typeMap_Staged_Order: make(map[*Xhtml_li_type]uint),

		Xhtml_object_typeMap_Staged_Order: make(map[*Xhtml_object_type]uint),

		Xhtml_ol_typeMap_Staged_Order: make(map[*Xhtml_ol_type]uint),

		Xhtml_p_typeMap_Staged_Order: make(map[*Xhtml_p_type]uint),

		Xhtml_param_typeMap_Staged_Order: make(map[*Xhtml_param_type]uint),

		Xhtml_pre_typeMap_Staged_Order: make(map[*Xhtml_pre_type]uint),

		Xhtml_q_typeMap_Staged_Order: make(map[*Xhtml_q_type]uint),

		Xhtml_samp_typeMap_Staged_Order: make(map[*Xhtml_samp_type]uint),

		Xhtml_span_typeMap_Staged_Order: make(map[*Xhtml_span_type]uint),

		Xhtml_strong_typeMap_Staged_Order: make(map[*Xhtml_strong_type]uint),

		Xhtml_table_typeMap_Staged_Order: make(map[*Xhtml_table_type]uint),

		Xhtml_tbody_typeMap_Staged_Order: make(map[*Xhtml_tbody_type]uint),

		Xhtml_td_typeMap_Staged_Order: make(map[*Xhtml_td_type]uint),

		Xhtml_tfoot_typeMap_Staged_Order: make(map[*Xhtml_tfoot_type]uint),

		Xhtml_th_typeMap_Staged_Order: make(map[*Xhtml_th_type]uint),

		Xhtml_thead_typeMap_Staged_Order: make(map[*Xhtml_thead_type]uint),

		Xhtml_tr_typeMap_Staged_Order: make(map[*Xhtml_tr_type]uint),

		Xhtml_ul_typeMap_Staged_Order: make(map[*Xhtml_ul_type]uint),

		Xhtml_var_typeMap_Staged_Order: make(map[*Xhtml_var_type]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "ALTERNATIVE_ID"},
			{name: "ATTRIBUTE_DEFINITION_BOOLEAN"},
			{name: "ATTRIBUTE_DEFINITION_DATE"},
			{name: "ATTRIBUTE_DEFINITION_ENUMERATION"},
			{name: "ATTRIBUTE_DEFINITION_INTEGER"},
			{name: "ATTRIBUTE_DEFINITION_REAL"},
			{name: "ATTRIBUTE_DEFINITION_STRING"},
			{name: "ATTRIBUTE_DEFINITION_XHTML"},
			{name: "ATTRIBUTE_VALUE_BOOLEAN"},
			{name: "ATTRIBUTE_VALUE_DATE"},
			{name: "ATTRIBUTE_VALUE_ENUMERATION"},
			{name: "ATTRIBUTE_VALUE_INTEGER"},
			{name: "ATTRIBUTE_VALUE_REAL"},
			{name: "ATTRIBUTE_VALUE_STRING"},
			{name: "ATTRIBUTE_VALUE_XHTML"},
			{name: "AnyType"},
			{name: "DATATYPE_DEFINITION_BOOLEAN"},
			{name: "DATATYPE_DEFINITION_DATE"},
			{name: "DATATYPE_DEFINITION_ENUMERATION"},
			{name: "DATATYPE_DEFINITION_INTEGER"},
			{name: "DATATYPE_DEFINITION_REAL"},
			{name: "DATATYPE_DEFINITION_STRING"},
			{name: "DATATYPE_DEFINITION_XHTML"},
			{name: "EMBEDDED_VALUE"},
			{name: "ENUM_VALUE"},
			{name: "RELATION_GROUP"},
			{name: "RELATION_GROUP_TYPE"},
			{name: "REQ_IF"},
			{name: "REQ_IF_CONTENT"},
			{name: "REQ_IF_HEADER"},
			{name: "REQ_IF_TOOL_EXTENSION"},
			{name: "SPECIFICATION"},
			{name: "SPECIFICATION_TYPE"},
			{name: "SPEC_HIERARCHY"},
			{name: "SPEC_OBJECT"},
			{name: "SPEC_OBJECT_TYPE"},
			{name: "SPEC_RELATION"},
			{name: "SPEC_RELATION_TYPE"},
			{name: "XHTML_CONTENT"},
			{name: "Xhtml_InlPres_type"},
			{name: "Xhtml_a_type"},
			{name: "Xhtml_abbr_type"},
			{name: "Xhtml_acronym_type"},
			{name: "Xhtml_address_type"},
			{name: "Xhtml_blockquote_type"},
			{name: "Xhtml_br_type"},
			{name: "Xhtml_caption_type"},
			{name: "Xhtml_cite_type"},
			{name: "Xhtml_code_type"},
			{name: "Xhtml_col_type"},
			{name: "Xhtml_colgroup_type"},
			{name: "Xhtml_dd_type"},
			{name: "Xhtml_dfn_type"},
			{name: "Xhtml_div_type"},
			{name: "Xhtml_dl_type"},
			{name: "Xhtml_dt_type"},
			{name: "Xhtml_edit_type"},
			{name: "Xhtml_em_type"},
			{name: "Xhtml_h1_type"},
			{name: "Xhtml_h2_type"},
			{name: "Xhtml_h3_type"},
			{name: "Xhtml_h4_type"},
			{name: "Xhtml_h5_type"},
			{name: "Xhtml_h6_type"},
			{name: "Xhtml_heading_type"},
			{name: "Xhtml_hr_type"},
			{name: "Xhtml_kbd_type"},
			{name: "Xhtml_li_type"},
			{name: "Xhtml_object_type"},
			{name: "Xhtml_ol_type"},
			{name: "Xhtml_p_type"},
			{name: "Xhtml_param_type"},
			{name: "Xhtml_pre_type"},
			{name: "Xhtml_q_type"},
			{name: "Xhtml_samp_type"},
			{name: "Xhtml_span_type"},
			{name: "Xhtml_strong_type"},
			{name: "Xhtml_table_type"},
			{name: "Xhtml_tbody_type"},
			{name: "Xhtml_td_type"},
			{name: "Xhtml_tfoot_type"},
			{name: "Xhtml_th_type"},
			{name: "Xhtml_thead_type"},
			{name: "Xhtml_tr_type"},
			{name: "Xhtml_ul_type"},
			{name: "Xhtml_var_type"},
		}, // end of insertion point
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *ALTERNATIVE_ID:
		return stage.ALTERNATIVE_IDMap_Staged_Order[instance]
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		return stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[instance]
	case *ATTRIBUTE_DEFINITION_DATE:
		return stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[instance]
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		return stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[instance]
	case *ATTRIBUTE_DEFINITION_INTEGER:
		return stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[instance]
	case *ATTRIBUTE_DEFINITION_REAL:
		return stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[instance]
	case *ATTRIBUTE_DEFINITION_STRING:
		return stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[instance]
	case *ATTRIBUTE_DEFINITION_XHTML:
		return stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[instance]
	case *ATTRIBUTE_VALUE_BOOLEAN:
		return stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[instance]
	case *ATTRIBUTE_VALUE_DATE:
		return stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[instance]
	case *ATTRIBUTE_VALUE_ENUMERATION:
		return stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[instance]
	case *ATTRIBUTE_VALUE_INTEGER:
		return stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[instance]
	case *ATTRIBUTE_VALUE_REAL:
		return stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[instance]
	case *ATTRIBUTE_VALUE_STRING:
		return stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[instance]
	case *ATTRIBUTE_VALUE_XHTML:
		return stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[instance]
	case *AnyType:
		return stage.AnyTypeMap_Staged_Order[instance]
	case *DATATYPE_DEFINITION_BOOLEAN:
		return stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[instance]
	case *DATATYPE_DEFINITION_DATE:
		return stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[instance]
	case *DATATYPE_DEFINITION_ENUMERATION:
		return stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[instance]
	case *DATATYPE_DEFINITION_INTEGER:
		return stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[instance]
	case *DATATYPE_DEFINITION_REAL:
		return stage.DATATYPE_DEFINITION_REALMap_Staged_Order[instance]
	case *DATATYPE_DEFINITION_STRING:
		return stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[instance]
	case *DATATYPE_DEFINITION_XHTML:
		return stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[instance]
	case *EMBEDDED_VALUE:
		return stage.EMBEDDED_VALUEMap_Staged_Order[instance]
	case *ENUM_VALUE:
		return stage.ENUM_VALUEMap_Staged_Order[instance]
	case *RELATION_GROUP:
		return stage.RELATION_GROUPMap_Staged_Order[instance]
	case *RELATION_GROUP_TYPE:
		return stage.RELATION_GROUP_TYPEMap_Staged_Order[instance]
	case *REQ_IF:
		return stage.REQ_IFMap_Staged_Order[instance]
	case *REQ_IF_CONTENT:
		return stage.REQ_IF_CONTENTMap_Staged_Order[instance]
	case *REQ_IF_HEADER:
		return stage.REQ_IF_HEADERMap_Staged_Order[instance]
	case *REQ_IF_TOOL_EXTENSION:
		return stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[instance]
	case *SPECIFICATION:
		return stage.SPECIFICATIONMap_Staged_Order[instance]
	case *SPECIFICATION_TYPE:
		return stage.SPECIFICATION_TYPEMap_Staged_Order[instance]
	case *SPEC_HIERARCHY:
		return stage.SPEC_HIERARCHYMap_Staged_Order[instance]
	case *SPEC_OBJECT:
		return stage.SPEC_OBJECTMap_Staged_Order[instance]
	case *SPEC_OBJECT_TYPE:
		return stage.SPEC_OBJECT_TYPEMap_Staged_Order[instance]
	case *SPEC_RELATION:
		return stage.SPEC_RELATIONMap_Staged_Order[instance]
	case *SPEC_RELATION_TYPE:
		return stage.SPEC_RELATION_TYPEMap_Staged_Order[instance]
	case *XHTML_CONTENT:
		return stage.XHTML_CONTENTMap_Staged_Order[instance]
	case *Xhtml_InlPres_type:
		return stage.Xhtml_InlPres_typeMap_Staged_Order[instance]
	case *Xhtml_a_type:
		return stage.Xhtml_a_typeMap_Staged_Order[instance]
	case *Xhtml_abbr_type:
		return stage.Xhtml_abbr_typeMap_Staged_Order[instance]
	case *Xhtml_acronym_type:
		return stage.Xhtml_acronym_typeMap_Staged_Order[instance]
	case *Xhtml_address_type:
		return stage.Xhtml_address_typeMap_Staged_Order[instance]
	case *Xhtml_blockquote_type:
		return stage.Xhtml_blockquote_typeMap_Staged_Order[instance]
	case *Xhtml_br_type:
		return stage.Xhtml_br_typeMap_Staged_Order[instance]
	case *Xhtml_caption_type:
		return stage.Xhtml_caption_typeMap_Staged_Order[instance]
	case *Xhtml_cite_type:
		return stage.Xhtml_cite_typeMap_Staged_Order[instance]
	case *Xhtml_code_type:
		return stage.Xhtml_code_typeMap_Staged_Order[instance]
	case *Xhtml_col_type:
		return stage.Xhtml_col_typeMap_Staged_Order[instance]
	case *Xhtml_colgroup_type:
		return stage.Xhtml_colgroup_typeMap_Staged_Order[instance]
	case *Xhtml_dd_type:
		return stage.Xhtml_dd_typeMap_Staged_Order[instance]
	case *Xhtml_dfn_type:
		return stage.Xhtml_dfn_typeMap_Staged_Order[instance]
	case *Xhtml_div_type:
		return stage.Xhtml_div_typeMap_Staged_Order[instance]
	case *Xhtml_dl_type:
		return stage.Xhtml_dl_typeMap_Staged_Order[instance]
	case *Xhtml_dt_type:
		return stage.Xhtml_dt_typeMap_Staged_Order[instance]
	case *Xhtml_edit_type:
		return stage.Xhtml_edit_typeMap_Staged_Order[instance]
	case *Xhtml_em_type:
		return stage.Xhtml_em_typeMap_Staged_Order[instance]
	case *Xhtml_h1_type:
		return stage.Xhtml_h1_typeMap_Staged_Order[instance]
	case *Xhtml_h2_type:
		return stage.Xhtml_h2_typeMap_Staged_Order[instance]
	case *Xhtml_h3_type:
		return stage.Xhtml_h3_typeMap_Staged_Order[instance]
	case *Xhtml_h4_type:
		return stage.Xhtml_h4_typeMap_Staged_Order[instance]
	case *Xhtml_h5_type:
		return stage.Xhtml_h5_typeMap_Staged_Order[instance]
	case *Xhtml_h6_type:
		return stage.Xhtml_h6_typeMap_Staged_Order[instance]
	case *Xhtml_heading_type:
		return stage.Xhtml_heading_typeMap_Staged_Order[instance]
	case *Xhtml_hr_type:
		return stage.Xhtml_hr_typeMap_Staged_Order[instance]
	case *Xhtml_kbd_type:
		return stage.Xhtml_kbd_typeMap_Staged_Order[instance]
	case *Xhtml_li_type:
		return stage.Xhtml_li_typeMap_Staged_Order[instance]
	case *Xhtml_object_type:
		return stage.Xhtml_object_typeMap_Staged_Order[instance]
	case *Xhtml_ol_type:
		return stage.Xhtml_ol_typeMap_Staged_Order[instance]
	case *Xhtml_p_type:
		return stage.Xhtml_p_typeMap_Staged_Order[instance]
	case *Xhtml_param_type:
		return stage.Xhtml_param_typeMap_Staged_Order[instance]
	case *Xhtml_pre_type:
		return stage.Xhtml_pre_typeMap_Staged_Order[instance]
	case *Xhtml_q_type:
		return stage.Xhtml_q_typeMap_Staged_Order[instance]
	case *Xhtml_samp_type:
		return stage.Xhtml_samp_typeMap_Staged_Order[instance]
	case *Xhtml_span_type:
		return stage.Xhtml_span_typeMap_Staged_Order[instance]
	case *Xhtml_strong_type:
		return stage.Xhtml_strong_typeMap_Staged_Order[instance]
	case *Xhtml_table_type:
		return stage.Xhtml_table_typeMap_Staged_Order[instance]
	case *Xhtml_tbody_type:
		return stage.Xhtml_tbody_typeMap_Staged_Order[instance]
	case *Xhtml_td_type:
		return stage.Xhtml_td_typeMap_Staged_Order[instance]
	case *Xhtml_tfoot_type:
		return stage.Xhtml_tfoot_typeMap_Staged_Order[instance]
	case *Xhtml_th_type:
		return stage.Xhtml_th_typeMap_Staged_Order[instance]
	case *Xhtml_thead_type:
		return stage.Xhtml_thead_typeMap_Staged_Order[instance]
	case *Xhtml_tr_type:
		return stage.Xhtml_tr_typeMap_Staged_Order[instance]
	case *Xhtml_ul_type:
		return stage.Xhtml_ul_typeMap_Staged_Order[instance]
	case *Xhtml_var_type:
		return stage.Xhtml_var_typeMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func (stage *Stage) GetName() string {
	return stage.name
}

func (stage *Stage) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["ALTERNATIVE_ID"] = len(stage.ALTERNATIVE_IDs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_BOOLEAN"] = len(stage.ATTRIBUTE_DEFINITION_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_DATE"] = len(stage.ATTRIBUTE_DEFINITION_DATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_ENUMERATION"] = len(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_INTEGER"] = len(stage.ATTRIBUTE_DEFINITION_INTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_REAL"] = len(stage.ATTRIBUTE_DEFINITION_REALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_STRING"] = len(stage.ATTRIBUTE_DEFINITION_STRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_XHTML"] = len(stage.ATTRIBUTE_DEFINITION_XHTMLs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_BOOLEAN"] = len(stage.ATTRIBUTE_VALUE_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_DATE"] = len(stage.ATTRIBUTE_VALUE_DATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_ENUMERATION"] = len(stage.ATTRIBUTE_VALUE_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_INTEGER"] = len(stage.ATTRIBUTE_VALUE_INTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_REAL"] = len(stage.ATTRIBUTE_VALUE_REALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_STRING"] = len(stage.ATTRIBUTE_VALUE_STRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_XHTML"] = len(stage.ATTRIBUTE_VALUE_XHTMLs)
	stage.Map_GongStructName_InstancesNb["AnyType"] = len(stage.AnyTypes)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_BOOLEAN"] = len(stage.DATATYPE_DEFINITION_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_DATE"] = len(stage.DATATYPE_DEFINITION_DATEs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_ENUMERATION"] = len(stage.DATATYPE_DEFINITION_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_INTEGER"] = len(stage.DATATYPE_DEFINITION_INTEGERs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_REAL"] = len(stage.DATATYPE_DEFINITION_REALs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_STRING"] = len(stage.DATATYPE_DEFINITION_STRINGs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_XHTML"] = len(stage.DATATYPE_DEFINITION_XHTMLs)
	stage.Map_GongStructName_InstancesNb["EMBEDDED_VALUE"] = len(stage.EMBEDDED_VALUEs)
	stage.Map_GongStructName_InstancesNb["ENUM_VALUE"] = len(stage.ENUM_VALUEs)
	stage.Map_GongStructName_InstancesNb["RELATION_GROUP"] = len(stage.RELATION_GROUPs)
	stage.Map_GongStructName_InstancesNb["RELATION_GROUP_TYPE"] = len(stage.RELATION_GROUP_TYPEs)
	stage.Map_GongStructName_InstancesNb["REQ_IF"] = len(stage.REQ_IFs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_CONTENT"] = len(stage.REQ_IF_CONTENTs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_HEADER"] = len(stage.REQ_IF_HEADERs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_TOOL_EXTENSION"] = len(stage.REQ_IF_TOOL_EXTENSIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION"] = len(stage.SPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION_TYPE"] = len(stage.SPECIFICATION_TYPEs)
	stage.Map_GongStructName_InstancesNb["SPEC_HIERARCHY"] = len(stage.SPEC_HIERARCHYs)
	stage.Map_GongStructName_InstancesNb["SPEC_OBJECT"] = len(stage.SPEC_OBJECTs)
	stage.Map_GongStructName_InstancesNb["SPEC_OBJECT_TYPE"] = len(stage.SPEC_OBJECT_TYPEs)
	stage.Map_GongStructName_InstancesNb["SPEC_RELATION"] = len(stage.SPEC_RELATIONs)
	stage.Map_GongStructName_InstancesNb["SPEC_RELATION_TYPE"] = len(stage.SPEC_RELATION_TYPEs)
	stage.Map_GongStructName_InstancesNb["XHTML_CONTENT"] = len(stage.XHTML_CONTENTs)
	stage.Map_GongStructName_InstancesNb["Xhtml_InlPres_type"] = len(stage.Xhtml_InlPres_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_a_type"] = len(stage.Xhtml_a_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_abbr_type"] = len(stage.Xhtml_abbr_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_acronym_type"] = len(stage.Xhtml_acronym_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_address_type"] = len(stage.Xhtml_address_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_blockquote_type"] = len(stage.Xhtml_blockquote_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_br_type"] = len(stage.Xhtml_br_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_caption_type"] = len(stage.Xhtml_caption_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_cite_type"] = len(stage.Xhtml_cite_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_code_type"] = len(stage.Xhtml_code_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_col_type"] = len(stage.Xhtml_col_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_colgroup_type"] = len(stage.Xhtml_colgroup_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_dd_type"] = len(stage.Xhtml_dd_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_dfn_type"] = len(stage.Xhtml_dfn_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_div_type"] = len(stage.Xhtml_div_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_dl_type"] = len(stage.Xhtml_dl_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_dt_type"] = len(stage.Xhtml_dt_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_edit_type"] = len(stage.Xhtml_edit_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_em_type"] = len(stage.Xhtml_em_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h1_type"] = len(stage.Xhtml_h1_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h2_type"] = len(stage.Xhtml_h2_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h3_type"] = len(stage.Xhtml_h3_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h4_type"] = len(stage.Xhtml_h4_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h5_type"] = len(stage.Xhtml_h5_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h6_type"] = len(stage.Xhtml_h6_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_heading_type"] = len(stage.Xhtml_heading_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_hr_type"] = len(stage.Xhtml_hr_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_kbd_type"] = len(stage.Xhtml_kbd_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_li_type"] = len(stage.Xhtml_li_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_object_type"] = len(stage.Xhtml_object_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_ol_type"] = len(stage.Xhtml_ol_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_p_type"] = len(stage.Xhtml_p_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_param_type"] = len(stage.Xhtml_param_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_pre_type"] = len(stage.Xhtml_pre_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_q_type"] = len(stage.Xhtml_q_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_samp_type"] = len(stage.Xhtml_samp_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_span_type"] = len(stage.Xhtml_span_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_strong_type"] = len(stage.Xhtml_strong_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_table_type"] = len(stage.Xhtml_table_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_tbody_type"] = len(stage.Xhtml_tbody_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_td_type"] = len(stage.Xhtml_td_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_tfoot_type"] = len(stage.Xhtml_tfoot_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_th_type"] = len(stage.Xhtml_th_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_thead_type"] = len(stage.Xhtml_thead_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_tr_type"] = len(stage.Xhtml_tr_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_ul_type"] = len(stage.Xhtml_ul_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_var_type"] = len(stage.Xhtml_var_types)

}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["ALTERNATIVE_ID"] = len(stage.ALTERNATIVE_IDs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_BOOLEAN"] = len(stage.ATTRIBUTE_DEFINITION_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_DATE"] = len(stage.ATTRIBUTE_DEFINITION_DATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_ENUMERATION"] = len(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_INTEGER"] = len(stage.ATTRIBUTE_DEFINITION_INTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_REAL"] = len(stage.ATTRIBUTE_DEFINITION_REALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_STRING"] = len(stage.ATTRIBUTE_DEFINITION_STRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_XHTML"] = len(stage.ATTRIBUTE_DEFINITION_XHTMLs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_BOOLEAN"] = len(stage.ATTRIBUTE_VALUE_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_DATE"] = len(stage.ATTRIBUTE_VALUE_DATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_ENUMERATION"] = len(stage.ATTRIBUTE_VALUE_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_INTEGER"] = len(stage.ATTRIBUTE_VALUE_INTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_REAL"] = len(stage.ATTRIBUTE_VALUE_REALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_STRING"] = len(stage.ATTRIBUTE_VALUE_STRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_XHTML"] = len(stage.ATTRIBUTE_VALUE_XHTMLs)
	stage.Map_GongStructName_InstancesNb["AnyType"] = len(stage.AnyTypes)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_BOOLEAN"] = len(stage.DATATYPE_DEFINITION_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_DATE"] = len(stage.DATATYPE_DEFINITION_DATEs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_ENUMERATION"] = len(stage.DATATYPE_DEFINITION_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_INTEGER"] = len(stage.DATATYPE_DEFINITION_INTEGERs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_REAL"] = len(stage.DATATYPE_DEFINITION_REALs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_STRING"] = len(stage.DATATYPE_DEFINITION_STRINGs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_XHTML"] = len(stage.DATATYPE_DEFINITION_XHTMLs)
	stage.Map_GongStructName_InstancesNb["EMBEDDED_VALUE"] = len(stage.EMBEDDED_VALUEs)
	stage.Map_GongStructName_InstancesNb["ENUM_VALUE"] = len(stage.ENUM_VALUEs)
	stage.Map_GongStructName_InstancesNb["RELATION_GROUP"] = len(stage.RELATION_GROUPs)
	stage.Map_GongStructName_InstancesNb["RELATION_GROUP_TYPE"] = len(stage.RELATION_GROUP_TYPEs)
	stage.Map_GongStructName_InstancesNb["REQ_IF"] = len(stage.REQ_IFs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_CONTENT"] = len(stage.REQ_IF_CONTENTs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_HEADER"] = len(stage.REQ_IF_HEADERs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_TOOL_EXTENSION"] = len(stage.REQ_IF_TOOL_EXTENSIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION"] = len(stage.SPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION_TYPE"] = len(stage.SPECIFICATION_TYPEs)
	stage.Map_GongStructName_InstancesNb["SPEC_HIERARCHY"] = len(stage.SPEC_HIERARCHYs)
	stage.Map_GongStructName_InstancesNb["SPEC_OBJECT"] = len(stage.SPEC_OBJECTs)
	stage.Map_GongStructName_InstancesNb["SPEC_OBJECT_TYPE"] = len(stage.SPEC_OBJECT_TYPEs)
	stage.Map_GongStructName_InstancesNb["SPEC_RELATION"] = len(stage.SPEC_RELATIONs)
	stage.Map_GongStructName_InstancesNb["SPEC_RELATION_TYPE"] = len(stage.SPEC_RELATION_TYPEs)
	stage.Map_GongStructName_InstancesNb["XHTML_CONTENT"] = len(stage.XHTML_CONTENTs)
	stage.Map_GongStructName_InstancesNb["Xhtml_InlPres_type"] = len(stage.Xhtml_InlPres_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_a_type"] = len(stage.Xhtml_a_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_abbr_type"] = len(stage.Xhtml_abbr_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_acronym_type"] = len(stage.Xhtml_acronym_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_address_type"] = len(stage.Xhtml_address_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_blockquote_type"] = len(stage.Xhtml_blockquote_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_br_type"] = len(stage.Xhtml_br_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_caption_type"] = len(stage.Xhtml_caption_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_cite_type"] = len(stage.Xhtml_cite_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_code_type"] = len(stage.Xhtml_code_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_col_type"] = len(stage.Xhtml_col_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_colgroup_type"] = len(stage.Xhtml_colgroup_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_dd_type"] = len(stage.Xhtml_dd_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_dfn_type"] = len(stage.Xhtml_dfn_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_div_type"] = len(stage.Xhtml_div_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_dl_type"] = len(stage.Xhtml_dl_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_dt_type"] = len(stage.Xhtml_dt_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_edit_type"] = len(stage.Xhtml_edit_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_em_type"] = len(stage.Xhtml_em_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h1_type"] = len(stage.Xhtml_h1_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h2_type"] = len(stage.Xhtml_h2_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h3_type"] = len(stage.Xhtml_h3_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h4_type"] = len(stage.Xhtml_h4_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h5_type"] = len(stage.Xhtml_h5_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_h6_type"] = len(stage.Xhtml_h6_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_heading_type"] = len(stage.Xhtml_heading_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_hr_type"] = len(stage.Xhtml_hr_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_kbd_type"] = len(stage.Xhtml_kbd_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_li_type"] = len(stage.Xhtml_li_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_object_type"] = len(stage.Xhtml_object_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_ol_type"] = len(stage.Xhtml_ol_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_p_type"] = len(stage.Xhtml_p_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_param_type"] = len(stage.Xhtml_param_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_pre_type"] = len(stage.Xhtml_pre_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_q_type"] = len(stage.Xhtml_q_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_samp_type"] = len(stage.Xhtml_samp_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_span_type"] = len(stage.Xhtml_span_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_strong_type"] = len(stage.Xhtml_strong_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_table_type"] = len(stage.Xhtml_table_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_tbody_type"] = len(stage.Xhtml_tbody_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_td_type"] = len(stage.Xhtml_td_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_tfoot_type"] = len(stage.Xhtml_tfoot_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_th_type"] = len(stage.Xhtml_th_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_thead_type"] = len(stage.Xhtml_thead_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_tr_type"] = len(stage.Xhtml_tr_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_ul_type"] = len(stage.Xhtml_ul_types)
	stage.Map_GongStructName_InstancesNb["Xhtml_var_type"] = len(stage.Xhtml_var_types)

}

// backup generates backup files in the dirPath
func (stage *Stage) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *Stage) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts alternative_id to the model stage
func (alternative_id *ALTERNATIVE_ID) Stage(stage *Stage) *ALTERNATIVE_ID {

	if _, ok := stage.ALTERNATIVE_IDs[alternative_id]; !ok {
		stage.ALTERNATIVE_IDs[alternative_id] = __member
		stage.ALTERNATIVE_IDMap_Staged_Order[alternative_id] = stage.ALTERNATIVE_IDOrder
		stage.ALTERNATIVE_IDOrder++
	}
	stage.ALTERNATIVE_IDs_mapString[alternative_id.Name] = alternative_id

	return alternative_id
}

// Unstage removes alternative_id off the model stage
func (alternative_id *ALTERNATIVE_ID) Unstage(stage *Stage) *ALTERNATIVE_ID {
	delete(stage.ALTERNATIVE_IDs, alternative_id)
	delete(stage.ALTERNATIVE_IDs_mapString, alternative_id.Name)
	return alternative_id
}

// UnstageVoid removes alternative_id off the model stage
func (alternative_id *ALTERNATIVE_ID) UnstageVoid(stage *Stage) {
	delete(stage.ALTERNATIVE_IDs, alternative_id)
	delete(stage.ALTERNATIVE_IDs_mapString, alternative_id.Name)
}

// commit alternative_id to the back repo (if it is already staged)
func (alternative_id *ALTERNATIVE_ID) Commit(stage *Stage) *ALTERNATIVE_ID {
	if _, ok := stage.ALTERNATIVE_IDs[alternative_id]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitALTERNATIVE_ID(alternative_id)
		}
	}
	return alternative_id
}

func (alternative_id *ALTERNATIVE_ID) CommitVoid(stage *Stage) {
	alternative_id.Commit(stage)
}

// Checkout alternative_id to the back repo (if it is already staged)
func (alternative_id *ALTERNATIVE_ID) Checkout(stage *Stage) *ALTERNATIVE_ID {
	if _, ok := stage.ALTERNATIVE_IDs[alternative_id]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutALTERNATIVE_ID(alternative_id)
		}
	}
	return alternative_id
}

// for satisfaction of GongStruct interface
func (alternative_id *ALTERNATIVE_ID) GetName() (res string) {
	return alternative_id.Name
}

// Stage puts attribute_definition_boolean to the model stage
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) Stage(stage *Stage) *ATTRIBUTE_DEFINITION_BOOLEAN {

	if _, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean]; !ok {
		stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean] = __member
		stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_boolean] = stage.ATTRIBUTE_DEFINITION_BOOLEANOrder
		stage.ATTRIBUTE_DEFINITION_BOOLEANOrder++
	}
	stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString[attribute_definition_boolean.Name] = attribute_definition_boolean

	return attribute_definition_boolean
}

// Unstage removes attribute_definition_boolean off the model stage
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) Unstage(stage *Stage) *ATTRIBUTE_DEFINITION_BOOLEAN {
	delete(stage.ATTRIBUTE_DEFINITION_BOOLEANs, attribute_definition_boolean)
	delete(stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString, attribute_definition_boolean.Name)
	return attribute_definition_boolean
}

// UnstageVoid removes attribute_definition_boolean off the model stage
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_DEFINITION_BOOLEANs, attribute_definition_boolean)
	delete(stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString, attribute_definition_boolean.Name)
}

// commit attribute_definition_boolean to the back repo (if it is already staged)
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) Commit(stage *Stage) *ATTRIBUTE_DEFINITION_BOOLEAN {
	if _, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean)
		}
	}
	return attribute_definition_boolean
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) CommitVoid(stage *Stage) {
	attribute_definition_boolean.Commit(stage)
}

// Checkout attribute_definition_boolean to the back repo (if it is already staged)
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) Checkout(stage *Stage) *ATTRIBUTE_DEFINITION_BOOLEAN {
	if _, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean)
		}
	}
	return attribute_definition_boolean
}

// for satisfaction of GongStruct interface
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GetName() (res string) {
	return attribute_definition_boolean.Name
}

// Stage puts attribute_definition_date to the model stage
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) Stage(stage *Stage) *ATTRIBUTE_DEFINITION_DATE {

	if _, ok := stage.ATTRIBUTE_DEFINITION_DATEs[attribute_definition_date]; !ok {
		stage.ATTRIBUTE_DEFINITION_DATEs[attribute_definition_date] = __member
		stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[attribute_definition_date] = stage.ATTRIBUTE_DEFINITION_DATEOrder
		stage.ATTRIBUTE_DEFINITION_DATEOrder++
	}
	stage.ATTRIBUTE_DEFINITION_DATEs_mapString[attribute_definition_date.Name] = attribute_definition_date

	return attribute_definition_date
}

// Unstage removes attribute_definition_date off the model stage
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) Unstage(stage *Stage) *ATTRIBUTE_DEFINITION_DATE {
	delete(stage.ATTRIBUTE_DEFINITION_DATEs, attribute_definition_date)
	delete(stage.ATTRIBUTE_DEFINITION_DATEs_mapString, attribute_definition_date.Name)
	return attribute_definition_date
}

// UnstageVoid removes attribute_definition_date off the model stage
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_DEFINITION_DATEs, attribute_definition_date)
	delete(stage.ATTRIBUTE_DEFINITION_DATEs_mapString, attribute_definition_date.Name)
}

// commit attribute_definition_date to the back repo (if it is already staged)
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) Commit(stage *Stage) *ATTRIBUTE_DEFINITION_DATE {
	if _, ok := stage.ATTRIBUTE_DEFINITION_DATEs[attribute_definition_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_DATE(attribute_definition_date)
		}
	}
	return attribute_definition_date
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) CommitVoid(stage *Stage) {
	attribute_definition_date.Commit(stage)
}

// Checkout attribute_definition_date to the back repo (if it is already staged)
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) Checkout(stage *Stage) *ATTRIBUTE_DEFINITION_DATE {
	if _, ok := stage.ATTRIBUTE_DEFINITION_DATEs[attribute_definition_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_DATE(attribute_definition_date)
		}
	}
	return attribute_definition_date
}

// for satisfaction of GongStruct interface
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GetName() (res string) {
	return attribute_definition_date.Name
}

// Stage puts attribute_definition_enumeration to the model stage
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) Stage(stage *Stage) *ATTRIBUTE_DEFINITION_ENUMERATION {

	if _, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[attribute_definition_enumeration]; !ok {
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[attribute_definition_enumeration] = __member
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[attribute_definition_enumeration] = stage.ATTRIBUTE_DEFINITION_ENUMERATIONOrder
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONOrder++
	}
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString[attribute_definition_enumeration.Name] = attribute_definition_enumeration

	return attribute_definition_enumeration
}

// Unstage removes attribute_definition_enumeration off the model stage
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) Unstage(stage *Stage) *ATTRIBUTE_DEFINITION_ENUMERATION {
	delete(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs, attribute_definition_enumeration)
	delete(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString, attribute_definition_enumeration.Name)
	return attribute_definition_enumeration
}

// UnstageVoid removes attribute_definition_enumeration off the model stage
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs, attribute_definition_enumeration)
	delete(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString, attribute_definition_enumeration.Name)
}

// commit attribute_definition_enumeration to the back repo (if it is already staged)
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) Commit(stage *Stage) *ATTRIBUTE_DEFINITION_ENUMERATION {
	if _, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[attribute_definition_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration)
		}
	}
	return attribute_definition_enumeration
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) CommitVoid(stage *Stage) {
	attribute_definition_enumeration.Commit(stage)
}

// Checkout attribute_definition_enumeration to the back repo (if it is already staged)
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) Checkout(stage *Stage) *ATTRIBUTE_DEFINITION_ENUMERATION {
	if _, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[attribute_definition_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration)
		}
	}
	return attribute_definition_enumeration
}

// for satisfaction of GongStruct interface
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GetName() (res string) {
	return attribute_definition_enumeration.Name
}

// Stage puts attribute_definition_integer to the model stage
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) Stage(stage *Stage) *ATTRIBUTE_DEFINITION_INTEGER {

	if _, ok := stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer]; !ok {
		stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer] = __member
		stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[attribute_definition_integer] = stage.ATTRIBUTE_DEFINITION_INTEGEROrder
		stage.ATTRIBUTE_DEFINITION_INTEGEROrder++
	}
	stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString[attribute_definition_integer.Name] = attribute_definition_integer

	return attribute_definition_integer
}

// Unstage removes attribute_definition_integer off the model stage
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) Unstage(stage *Stage) *ATTRIBUTE_DEFINITION_INTEGER {
	delete(stage.ATTRIBUTE_DEFINITION_INTEGERs, attribute_definition_integer)
	delete(stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString, attribute_definition_integer.Name)
	return attribute_definition_integer
}

// UnstageVoid removes attribute_definition_integer off the model stage
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_DEFINITION_INTEGERs, attribute_definition_integer)
	delete(stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString, attribute_definition_integer.Name)
}

// commit attribute_definition_integer to the back repo (if it is already staged)
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) Commit(stage *Stage) *ATTRIBUTE_DEFINITION_INTEGER {
	if _, ok := stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer)
		}
	}
	return attribute_definition_integer
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) CommitVoid(stage *Stage) {
	attribute_definition_integer.Commit(stage)
}

// Checkout attribute_definition_integer to the back repo (if it is already staged)
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) Checkout(stage *Stage) *ATTRIBUTE_DEFINITION_INTEGER {
	if _, ok := stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer)
		}
	}
	return attribute_definition_integer
}

// for satisfaction of GongStruct interface
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GetName() (res string) {
	return attribute_definition_integer.Name
}

// Stage puts attribute_definition_real to the model stage
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) Stage(stage *Stage) *ATTRIBUTE_DEFINITION_REAL {

	if _, ok := stage.ATTRIBUTE_DEFINITION_REALs[attribute_definition_real]; !ok {
		stage.ATTRIBUTE_DEFINITION_REALs[attribute_definition_real] = __member
		stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[attribute_definition_real] = stage.ATTRIBUTE_DEFINITION_REALOrder
		stage.ATTRIBUTE_DEFINITION_REALOrder++
	}
	stage.ATTRIBUTE_DEFINITION_REALs_mapString[attribute_definition_real.Name] = attribute_definition_real

	return attribute_definition_real
}

// Unstage removes attribute_definition_real off the model stage
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) Unstage(stage *Stage) *ATTRIBUTE_DEFINITION_REAL {
	delete(stage.ATTRIBUTE_DEFINITION_REALs, attribute_definition_real)
	delete(stage.ATTRIBUTE_DEFINITION_REALs_mapString, attribute_definition_real.Name)
	return attribute_definition_real
}

// UnstageVoid removes attribute_definition_real off the model stage
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_DEFINITION_REALs, attribute_definition_real)
	delete(stage.ATTRIBUTE_DEFINITION_REALs_mapString, attribute_definition_real.Name)
}

// commit attribute_definition_real to the back repo (if it is already staged)
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) Commit(stage *Stage) *ATTRIBUTE_DEFINITION_REAL {
	if _, ok := stage.ATTRIBUTE_DEFINITION_REALs[attribute_definition_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_REAL(attribute_definition_real)
		}
	}
	return attribute_definition_real
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) CommitVoid(stage *Stage) {
	attribute_definition_real.Commit(stage)
}

// Checkout attribute_definition_real to the back repo (if it is already staged)
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) Checkout(stage *Stage) *ATTRIBUTE_DEFINITION_REAL {
	if _, ok := stage.ATTRIBUTE_DEFINITION_REALs[attribute_definition_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_REAL(attribute_definition_real)
		}
	}
	return attribute_definition_real
}

// for satisfaction of GongStruct interface
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GetName() (res string) {
	return attribute_definition_real.Name
}

// Stage puts attribute_definition_string to the model stage
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) Stage(stage *Stage) *ATTRIBUTE_DEFINITION_STRING {

	if _, ok := stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string]; !ok {
		stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string] = __member
		stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_string] = stage.ATTRIBUTE_DEFINITION_STRINGOrder
		stage.ATTRIBUTE_DEFINITION_STRINGOrder++
	}
	stage.ATTRIBUTE_DEFINITION_STRINGs_mapString[attribute_definition_string.Name] = attribute_definition_string

	return attribute_definition_string
}

// Unstage removes attribute_definition_string off the model stage
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) Unstage(stage *Stage) *ATTRIBUTE_DEFINITION_STRING {
	delete(stage.ATTRIBUTE_DEFINITION_STRINGs, attribute_definition_string)
	delete(stage.ATTRIBUTE_DEFINITION_STRINGs_mapString, attribute_definition_string.Name)
	return attribute_definition_string
}

// UnstageVoid removes attribute_definition_string off the model stage
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_DEFINITION_STRINGs, attribute_definition_string)
	delete(stage.ATTRIBUTE_DEFINITION_STRINGs_mapString, attribute_definition_string.Name)
}

// commit attribute_definition_string to the back repo (if it is already staged)
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) Commit(stage *Stage) *ATTRIBUTE_DEFINITION_STRING {
	if _, ok := stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_STRING(attribute_definition_string)
		}
	}
	return attribute_definition_string
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) CommitVoid(stage *Stage) {
	attribute_definition_string.Commit(stage)
}

// Checkout attribute_definition_string to the back repo (if it is already staged)
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) Checkout(stage *Stage) *ATTRIBUTE_DEFINITION_STRING {
	if _, ok := stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_STRING(attribute_definition_string)
		}
	}
	return attribute_definition_string
}

// for satisfaction of GongStruct interface
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GetName() (res string) {
	return attribute_definition_string.Name
}

// Stage puts attribute_definition_xhtml to the model stage
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) Stage(stage *Stage) *ATTRIBUTE_DEFINITION_XHTML {

	if _, ok := stage.ATTRIBUTE_DEFINITION_XHTMLs[attribute_definition_xhtml]; !ok {
		stage.ATTRIBUTE_DEFINITION_XHTMLs[attribute_definition_xhtml] = __member
		stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[attribute_definition_xhtml] = stage.ATTRIBUTE_DEFINITION_XHTMLOrder
		stage.ATTRIBUTE_DEFINITION_XHTMLOrder++
	}
	stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString[attribute_definition_xhtml.Name] = attribute_definition_xhtml

	return attribute_definition_xhtml
}

// Unstage removes attribute_definition_xhtml off the model stage
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) Unstage(stage *Stage) *ATTRIBUTE_DEFINITION_XHTML {
	delete(stage.ATTRIBUTE_DEFINITION_XHTMLs, attribute_definition_xhtml)
	delete(stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString, attribute_definition_xhtml.Name)
	return attribute_definition_xhtml
}

// UnstageVoid removes attribute_definition_xhtml off the model stage
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_DEFINITION_XHTMLs, attribute_definition_xhtml)
	delete(stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString, attribute_definition_xhtml.Name)
}

// commit attribute_definition_xhtml to the back repo (if it is already staged)
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) Commit(stage *Stage) *ATTRIBUTE_DEFINITION_XHTML {
	if _, ok := stage.ATTRIBUTE_DEFINITION_XHTMLs[attribute_definition_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml)
		}
	}
	return attribute_definition_xhtml
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) CommitVoid(stage *Stage) {
	attribute_definition_xhtml.Commit(stage)
}

// Checkout attribute_definition_xhtml to the back repo (if it is already staged)
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) Checkout(stage *Stage) *ATTRIBUTE_DEFINITION_XHTML {
	if _, ok := stage.ATTRIBUTE_DEFINITION_XHTMLs[attribute_definition_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml)
		}
	}
	return attribute_definition_xhtml
}

// for satisfaction of GongStruct interface
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GetName() (res string) {
	return attribute_definition_xhtml.Name
}

// Stage puts attribute_value_boolean to the model stage
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) Stage(stage *Stage) *ATTRIBUTE_VALUE_BOOLEAN {

	if _, ok := stage.ATTRIBUTE_VALUE_BOOLEANs[attribute_value_boolean]; !ok {
		stage.ATTRIBUTE_VALUE_BOOLEANs[attribute_value_boolean] = __member
		stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[attribute_value_boolean] = stage.ATTRIBUTE_VALUE_BOOLEANOrder
		stage.ATTRIBUTE_VALUE_BOOLEANOrder++
	}
	stage.ATTRIBUTE_VALUE_BOOLEANs_mapString[attribute_value_boolean.Name] = attribute_value_boolean

	return attribute_value_boolean
}

// Unstage removes attribute_value_boolean off the model stage
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) Unstage(stage *Stage) *ATTRIBUTE_VALUE_BOOLEAN {
	delete(stage.ATTRIBUTE_VALUE_BOOLEANs, attribute_value_boolean)
	delete(stage.ATTRIBUTE_VALUE_BOOLEANs_mapString, attribute_value_boolean.Name)
	return attribute_value_boolean
}

// UnstageVoid removes attribute_value_boolean off the model stage
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_VALUE_BOOLEANs, attribute_value_boolean)
	delete(stage.ATTRIBUTE_VALUE_BOOLEANs_mapString, attribute_value_boolean.Name)
}

// commit attribute_value_boolean to the back repo (if it is already staged)
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) Commit(stage *Stage) *ATTRIBUTE_VALUE_BOOLEAN {
	if _, ok := stage.ATTRIBUTE_VALUE_BOOLEANs[attribute_value_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean)
		}
	}
	return attribute_value_boolean
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) CommitVoid(stage *Stage) {
	attribute_value_boolean.Commit(stage)
}

// Checkout attribute_value_boolean to the back repo (if it is already staged)
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) Checkout(stage *Stage) *ATTRIBUTE_VALUE_BOOLEAN {
	if _, ok := stage.ATTRIBUTE_VALUE_BOOLEANs[attribute_value_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean)
		}
	}
	return attribute_value_boolean
}

// for satisfaction of GongStruct interface
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GetName() (res string) {
	return attribute_value_boolean.Name
}

// Stage puts attribute_value_date to the model stage
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) Stage(stage *Stage) *ATTRIBUTE_VALUE_DATE {

	if _, ok := stage.ATTRIBUTE_VALUE_DATEs[attribute_value_date]; !ok {
		stage.ATTRIBUTE_VALUE_DATEs[attribute_value_date] = __member
		stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[attribute_value_date] = stage.ATTRIBUTE_VALUE_DATEOrder
		stage.ATTRIBUTE_VALUE_DATEOrder++
	}
	stage.ATTRIBUTE_VALUE_DATEs_mapString[attribute_value_date.Name] = attribute_value_date

	return attribute_value_date
}

// Unstage removes attribute_value_date off the model stage
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) Unstage(stage *Stage) *ATTRIBUTE_VALUE_DATE {
	delete(stage.ATTRIBUTE_VALUE_DATEs, attribute_value_date)
	delete(stage.ATTRIBUTE_VALUE_DATEs_mapString, attribute_value_date.Name)
	return attribute_value_date
}

// UnstageVoid removes attribute_value_date off the model stage
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_VALUE_DATEs, attribute_value_date)
	delete(stage.ATTRIBUTE_VALUE_DATEs_mapString, attribute_value_date.Name)
}

// commit attribute_value_date to the back repo (if it is already staged)
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) Commit(stage *Stage) *ATTRIBUTE_VALUE_DATE {
	if _, ok := stage.ATTRIBUTE_VALUE_DATEs[attribute_value_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_DATE(attribute_value_date)
		}
	}
	return attribute_value_date
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) CommitVoid(stage *Stage) {
	attribute_value_date.Commit(stage)
}

// Checkout attribute_value_date to the back repo (if it is already staged)
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) Checkout(stage *Stage) *ATTRIBUTE_VALUE_DATE {
	if _, ok := stage.ATTRIBUTE_VALUE_DATEs[attribute_value_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_DATE(attribute_value_date)
		}
	}
	return attribute_value_date
}

// for satisfaction of GongStruct interface
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GetName() (res string) {
	return attribute_value_date.Name
}

// Stage puts attribute_value_enumeration to the model stage
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) Stage(stage *Stage) *ATTRIBUTE_VALUE_ENUMERATION {

	if _, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONs[attribute_value_enumeration]; !ok {
		stage.ATTRIBUTE_VALUE_ENUMERATIONs[attribute_value_enumeration] = __member
		stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[attribute_value_enumeration] = stage.ATTRIBUTE_VALUE_ENUMERATIONOrder
		stage.ATTRIBUTE_VALUE_ENUMERATIONOrder++
	}
	stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString[attribute_value_enumeration.Name] = attribute_value_enumeration

	return attribute_value_enumeration
}

// Unstage removes attribute_value_enumeration off the model stage
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) Unstage(stage *Stage) *ATTRIBUTE_VALUE_ENUMERATION {
	delete(stage.ATTRIBUTE_VALUE_ENUMERATIONs, attribute_value_enumeration)
	delete(stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString, attribute_value_enumeration.Name)
	return attribute_value_enumeration
}

// UnstageVoid removes attribute_value_enumeration off the model stage
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_VALUE_ENUMERATIONs, attribute_value_enumeration)
	delete(stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString, attribute_value_enumeration.Name)
}

// commit attribute_value_enumeration to the back repo (if it is already staged)
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) Commit(stage *Stage) *ATTRIBUTE_VALUE_ENUMERATION {
	if _, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONs[attribute_value_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration)
		}
	}
	return attribute_value_enumeration
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) CommitVoid(stage *Stage) {
	attribute_value_enumeration.Commit(stage)
}

// Checkout attribute_value_enumeration to the back repo (if it is already staged)
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) Checkout(stage *Stage) *ATTRIBUTE_VALUE_ENUMERATION {
	if _, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONs[attribute_value_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration)
		}
	}
	return attribute_value_enumeration
}

// for satisfaction of GongStruct interface
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GetName() (res string) {
	return attribute_value_enumeration.Name
}

// Stage puts attribute_value_integer to the model stage
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) Stage(stage *Stage) *ATTRIBUTE_VALUE_INTEGER {

	if _, ok := stage.ATTRIBUTE_VALUE_INTEGERs[attribute_value_integer]; !ok {
		stage.ATTRIBUTE_VALUE_INTEGERs[attribute_value_integer] = __member
		stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[attribute_value_integer] = stage.ATTRIBUTE_VALUE_INTEGEROrder
		stage.ATTRIBUTE_VALUE_INTEGEROrder++
	}
	stage.ATTRIBUTE_VALUE_INTEGERs_mapString[attribute_value_integer.Name] = attribute_value_integer

	return attribute_value_integer
}

// Unstage removes attribute_value_integer off the model stage
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) Unstage(stage *Stage) *ATTRIBUTE_VALUE_INTEGER {
	delete(stage.ATTRIBUTE_VALUE_INTEGERs, attribute_value_integer)
	delete(stage.ATTRIBUTE_VALUE_INTEGERs_mapString, attribute_value_integer.Name)
	return attribute_value_integer
}

// UnstageVoid removes attribute_value_integer off the model stage
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_VALUE_INTEGERs, attribute_value_integer)
	delete(stage.ATTRIBUTE_VALUE_INTEGERs_mapString, attribute_value_integer.Name)
}

// commit attribute_value_integer to the back repo (if it is already staged)
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) Commit(stage *Stage) *ATTRIBUTE_VALUE_INTEGER {
	if _, ok := stage.ATTRIBUTE_VALUE_INTEGERs[attribute_value_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_INTEGER(attribute_value_integer)
		}
	}
	return attribute_value_integer
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) CommitVoid(stage *Stage) {
	attribute_value_integer.Commit(stage)
}

// Checkout attribute_value_integer to the back repo (if it is already staged)
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) Checkout(stage *Stage) *ATTRIBUTE_VALUE_INTEGER {
	if _, ok := stage.ATTRIBUTE_VALUE_INTEGERs[attribute_value_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_INTEGER(attribute_value_integer)
		}
	}
	return attribute_value_integer
}

// for satisfaction of GongStruct interface
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GetName() (res string) {
	return attribute_value_integer.Name
}

// Stage puts attribute_value_real to the model stage
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) Stage(stage *Stage) *ATTRIBUTE_VALUE_REAL {

	if _, ok := stage.ATTRIBUTE_VALUE_REALs[attribute_value_real]; !ok {
		stage.ATTRIBUTE_VALUE_REALs[attribute_value_real] = __member
		stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[attribute_value_real] = stage.ATTRIBUTE_VALUE_REALOrder
		stage.ATTRIBUTE_VALUE_REALOrder++
	}
	stage.ATTRIBUTE_VALUE_REALs_mapString[attribute_value_real.Name] = attribute_value_real

	return attribute_value_real
}

// Unstage removes attribute_value_real off the model stage
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) Unstage(stage *Stage) *ATTRIBUTE_VALUE_REAL {
	delete(stage.ATTRIBUTE_VALUE_REALs, attribute_value_real)
	delete(stage.ATTRIBUTE_VALUE_REALs_mapString, attribute_value_real.Name)
	return attribute_value_real
}

// UnstageVoid removes attribute_value_real off the model stage
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_VALUE_REALs, attribute_value_real)
	delete(stage.ATTRIBUTE_VALUE_REALs_mapString, attribute_value_real.Name)
}

// commit attribute_value_real to the back repo (if it is already staged)
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) Commit(stage *Stage) *ATTRIBUTE_VALUE_REAL {
	if _, ok := stage.ATTRIBUTE_VALUE_REALs[attribute_value_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_REAL(attribute_value_real)
		}
	}
	return attribute_value_real
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) CommitVoid(stage *Stage) {
	attribute_value_real.Commit(stage)
}

// Checkout attribute_value_real to the back repo (if it is already staged)
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) Checkout(stage *Stage) *ATTRIBUTE_VALUE_REAL {
	if _, ok := stage.ATTRIBUTE_VALUE_REALs[attribute_value_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_REAL(attribute_value_real)
		}
	}
	return attribute_value_real
}

// for satisfaction of GongStruct interface
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GetName() (res string) {
	return attribute_value_real.Name
}

// Stage puts attribute_value_string to the model stage
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) Stage(stage *Stage) *ATTRIBUTE_VALUE_STRING {

	if _, ok := stage.ATTRIBUTE_VALUE_STRINGs[attribute_value_string]; !ok {
		stage.ATTRIBUTE_VALUE_STRINGs[attribute_value_string] = __member
		stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[attribute_value_string] = stage.ATTRIBUTE_VALUE_STRINGOrder
		stage.ATTRIBUTE_VALUE_STRINGOrder++
	}
	stage.ATTRIBUTE_VALUE_STRINGs_mapString[attribute_value_string.Name] = attribute_value_string

	return attribute_value_string
}

// Unstage removes attribute_value_string off the model stage
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) Unstage(stage *Stage) *ATTRIBUTE_VALUE_STRING {
	delete(stage.ATTRIBUTE_VALUE_STRINGs, attribute_value_string)
	delete(stage.ATTRIBUTE_VALUE_STRINGs_mapString, attribute_value_string.Name)
	return attribute_value_string
}

// UnstageVoid removes attribute_value_string off the model stage
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_VALUE_STRINGs, attribute_value_string)
	delete(stage.ATTRIBUTE_VALUE_STRINGs_mapString, attribute_value_string.Name)
}

// commit attribute_value_string to the back repo (if it is already staged)
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) Commit(stage *Stage) *ATTRIBUTE_VALUE_STRING {
	if _, ok := stage.ATTRIBUTE_VALUE_STRINGs[attribute_value_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_STRING(attribute_value_string)
		}
	}
	return attribute_value_string
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) CommitVoid(stage *Stage) {
	attribute_value_string.Commit(stage)
}

// Checkout attribute_value_string to the back repo (if it is already staged)
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) Checkout(stage *Stage) *ATTRIBUTE_VALUE_STRING {
	if _, ok := stage.ATTRIBUTE_VALUE_STRINGs[attribute_value_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_STRING(attribute_value_string)
		}
	}
	return attribute_value_string
}

// for satisfaction of GongStruct interface
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GetName() (res string) {
	return attribute_value_string.Name
}

// Stage puts attribute_value_xhtml to the model stage
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) Stage(stage *Stage) *ATTRIBUTE_VALUE_XHTML {

	if _, ok := stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml]; !ok {
		stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml] = __member
		stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtml] = stage.ATTRIBUTE_VALUE_XHTMLOrder
		stage.ATTRIBUTE_VALUE_XHTMLOrder++
	}
	stage.ATTRIBUTE_VALUE_XHTMLs_mapString[attribute_value_xhtml.Name] = attribute_value_xhtml

	return attribute_value_xhtml
}

// Unstage removes attribute_value_xhtml off the model stage
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) Unstage(stage *Stage) *ATTRIBUTE_VALUE_XHTML {
	delete(stage.ATTRIBUTE_VALUE_XHTMLs, attribute_value_xhtml)
	delete(stage.ATTRIBUTE_VALUE_XHTMLs_mapString, attribute_value_xhtml.Name)
	return attribute_value_xhtml
}

// UnstageVoid removes attribute_value_xhtml off the model stage
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) UnstageVoid(stage *Stage) {
	delete(stage.ATTRIBUTE_VALUE_XHTMLs, attribute_value_xhtml)
	delete(stage.ATTRIBUTE_VALUE_XHTMLs_mapString, attribute_value_xhtml.Name)
}

// commit attribute_value_xhtml to the back repo (if it is already staged)
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) Commit(stage *Stage) *ATTRIBUTE_VALUE_XHTML {
	if _, ok := stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_XHTML(attribute_value_xhtml)
		}
	}
	return attribute_value_xhtml
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) CommitVoid(stage *Stage) {
	attribute_value_xhtml.Commit(stage)
}

// Checkout attribute_value_xhtml to the back repo (if it is already staged)
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) Checkout(stage *Stage) *ATTRIBUTE_VALUE_XHTML {
	if _, ok := stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_XHTML(attribute_value_xhtml)
		}
	}
	return attribute_value_xhtml
}

// for satisfaction of GongStruct interface
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GetName() (res string) {
	return attribute_value_xhtml.Name
}

// Stage puts anytype to the model stage
func (anytype *AnyType) Stage(stage *Stage) *AnyType {

	if _, ok := stage.AnyTypes[anytype]; !ok {
		stage.AnyTypes[anytype] = __member
		stage.AnyTypeMap_Staged_Order[anytype] = stage.AnyTypeOrder
		stage.AnyTypeOrder++
	}
	stage.AnyTypes_mapString[anytype.Name] = anytype

	return anytype
}

// Unstage removes anytype off the model stage
func (anytype *AnyType) Unstage(stage *Stage) *AnyType {
	delete(stage.AnyTypes, anytype)
	delete(stage.AnyTypes_mapString, anytype.Name)
	return anytype
}

// UnstageVoid removes anytype off the model stage
func (anytype *AnyType) UnstageVoid(stage *Stage) {
	delete(stage.AnyTypes, anytype)
	delete(stage.AnyTypes_mapString, anytype.Name)
}

// commit anytype to the back repo (if it is already staged)
func (anytype *AnyType) Commit(stage *Stage) *AnyType {
	if _, ok := stage.AnyTypes[anytype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAnyType(anytype)
		}
	}
	return anytype
}

func (anytype *AnyType) CommitVoid(stage *Stage) {
	anytype.Commit(stage)
}

// Checkout anytype to the back repo (if it is already staged)
func (anytype *AnyType) Checkout(stage *Stage) *AnyType {
	if _, ok := stage.AnyTypes[anytype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAnyType(anytype)
		}
	}
	return anytype
}

// for satisfaction of GongStruct interface
func (anytype *AnyType) GetName() (res string) {
	return anytype.Name
}

// Stage puts datatype_definition_boolean to the model stage
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) Stage(stage *Stage) *DATATYPE_DEFINITION_BOOLEAN {

	if _, ok := stage.DATATYPE_DEFINITION_BOOLEANs[datatype_definition_boolean]; !ok {
		stage.DATATYPE_DEFINITION_BOOLEANs[datatype_definition_boolean] = __member
		stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[datatype_definition_boolean] = stage.DATATYPE_DEFINITION_BOOLEANOrder
		stage.DATATYPE_DEFINITION_BOOLEANOrder++
	}
	stage.DATATYPE_DEFINITION_BOOLEANs_mapString[datatype_definition_boolean.Name] = datatype_definition_boolean

	return datatype_definition_boolean
}

// Unstage removes datatype_definition_boolean off the model stage
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) Unstage(stage *Stage) *DATATYPE_DEFINITION_BOOLEAN {
	delete(stage.DATATYPE_DEFINITION_BOOLEANs, datatype_definition_boolean)
	delete(stage.DATATYPE_DEFINITION_BOOLEANs_mapString, datatype_definition_boolean.Name)
	return datatype_definition_boolean
}

// UnstageVoid removes datatype_definition_boolean off the model stage
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) UnstageVoid(stage *Stage) {
	delete(stage.DATATYPE_DEFINITION_BOOLEANs, datatype_definition_boolean)
	delete(stage.DATATYPE_DEFINITION_BOOLEANs_mapString, datatype_definition_boolean.Name)
}

// commit datatype_definition_boolean to the back repo (if it is already staged)
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) Commit(stage *Stage) *DATATYPE_DEFINITION_BOOLEAN {
	if _, ok := stage.DATATYPE_DEFINITION_BOOLEANs[datatype_definition_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean)
		}
	}
	return datatype_definition_boolean
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) CommitVoid(stage *Stage) {
	datatype_definition_boolean.Commit(stage)
}

// Checkout datatype_definition_boolean to the back repo (if it is already staged)
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) Checkout(stage *Stage) *DATATYPE_DEFINITION_BOOLEAN {
	if _, ok := stage.DATATYPE_DEFINITION_BOOLEANs[datatype_definition_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean)
		}
	}
	return datatype_definition_boolean
}

// for satisfaction of GongStruct interface
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GetName() (res string) {
	return datatype_definition_boolean.Name
}

// Stage puts datatype_definition_date to the model stage
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) Stage(stage *Stage) *DATATYPE_DEFINITION_DATE {

	if _, ok := stage.DATATYPE_DEFINITION_DATEs[datatype_definition_date]; !ok {
		stage.DATATYPE_DEFINITION_DATEs[datatype_definition_date] = __member
		stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[datatype_definition_date] = stage.DATATYPE_DEFINITION_DATEOrder
		stage.DATATYPE_DEFINITION_DATEOrder++
	}
	stage.DATATYPE_DEFINITION_DATEs_mapString[datatype_definition_date.Name] = datatype_definition_date

	return datatype_definition_date
}

// Unstage removes datatype_definition_date off the model stage
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) Unstage(stage *Stage) *DATATYPE_DEFINITION_DATE {
	delete(stage.DATATYPE_DEFINITION_DATEs, datatype_definition_date)
	delete(stage.DATATYPE_DEFINITION_DATEs_mapString, datatype_definition_date.Name)
	return datatype_definition_date
}

// UnstageVoid removes datatype_definition_date off the model stage
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) UnstageVoid(stage *Stage) {
	delete(stage.DATATYPE_DEFINITION_DATEs, datatype_definition_date)
	delete(stage.DATATYPE_DEFINITION_DATEs_mapString, datatype_definition_date.Name)
}

// commit datatype_definition_date to the back repo (if it is already staged)
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) Commit(stage *Stage) *DATATYPE_DEFINITION_DATE {
	if _, ok := stage.DATATYPE_DEFINITION_DATEs[datatype_definition_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_DATE(datatype_definition_date)
		}
	}
	return datatype_definition_date
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) CommitVoid(stage *Stage) {
	datatype_definition_date.Commit(stage)
}

// Checkout datatype_definition_date to the back repo (if it is already staged)
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) Checkout(stage *Stage) *DATATYPE_DEFINITION_DATE {
	if _, ok := stage.DATATYPE_DEFINITION_DATEs[datatype_definition_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_DATE(datatype_definition_date)
		}
	}
	return datatype_definition_date
}

// for satisfaction of GongStruct interface
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GetName() (res string) {
	return datatype_definition_date.Name
}

// Stage puts datatype_definition_enumeration to the model stage
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) Stage(stage *Stage) *DATATYPE_DEFINITION_ENUMERATION {

	if _, ok := stage.DATATYPE_DEFINITION_ENUMERATIONs[datatype_definition_enumeration]; !ok {
		stage.DATATYPE_DEFINITION_ENUMERATIONs[datatype_definition_enumeration] = __member
		stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[datatype_definition_enumeration] = stage.DATATYPE_DEFINITION_ENUMERATIONOrder
		stage.DATATYPE_DEFINITION_ENUMERATIONOrder++
	}
	stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString[datatype_definition_enumeration.Name] = datatype_definition_enumeration

	return datatype_definition_enumeration
}

// Unstage removes datatype_definition_enumeration off the model stage
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) Unstage(stage *Stage) *DATATYPE_DEFINITION_ENUMERATION {
	delete(stage.DATATYPE_DEFINITION_ENUMERATIONs, datatype_definition_enumeration)
	delete(stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString, datatype_definition_enumeration.Name)
	return datatype_definition_enumeration
}

// UnstageVoid removes datatype_definition_enumeration off the model stage
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) UnstageVoid(stage *Stage) {
	delete(stage.DATATYPE_DEFINITION_ENUMERATIONs, datatype_definition_enumeration)
	delete(stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString, datatype_definition_enumeration.Name)
}

// commit datatype_definition_enumeration to the back repo (if it is already staged)
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) Commit(stage *Stage) *DATATYPE_DEFINITION_ENUMERATION {
	if _, ok := stage.DATATYPE_DEFINITION_ENUMERATIONs[datatype_definition_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration)
		}
	}
	return datatype_definition_enumeration
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) CommitVoid(stage *Stage) {
	datatype_definition_enumeration.Commit(stage)
}

// Checkout datatype_definition_enumeration to the back repo (if it is already staged)
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) Checkout(stage *Stage) *DATATYPE_DEFINITION_ENUMERATION {
	if _, ok := stage.DATATYPE_DEFINITION_ENUMERATIONs[datatype_definition_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration)
		}
	}
	return datatype_definition_enumeration
}

// for satisfaction of GongStruct interface
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GetName() (res string) {
	return datatype_definition_enumeration.Name
}

// Stage puts datatype_definition_integer to the model stage
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) Stage(stage *Stage) *DATATYPE_DEFINITION_INTEGER {

	if _, ok := stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer]; !ok {
		stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer] = __member
		stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[datatype_definition_integer] = stage.DATATYPE_DEFINITION_INTEGEROrder
		stage.DATATYPE_DEFINITION_INTEGEROrder++
	}
	stage.DATATYPE_DEFINITION_INTEGERs_mapString[datatype_definition_integer.Name] = datatype_definition_integer

	return datatype_definition_integer
}

// Unstage removes datatype_definition_integer off the model stage
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) Unstage(stage *Stage) *DATATYPE_DEFINITION_INTEGER {
	delete(stage.DATATYPE_DEFINITION_INTEGERs, datatype_definition_integer)
	delete(stage.DATATYPE_DEFINITION_INTEGERs_mapString, datatype_definition_integer.Name)
	return datatype_definition_integer
}

// UnstageVoid removes datatype_definition_integer off the model stage
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) UnstageVoid(stage *Stage) {
	delete(stage.DATATYPE_DEFINITION_INTEGERs, datatype_definition_integer)
	delete(stage.DATATYPE_DEFINITION_INTEGERs_mapString, datatype_definition_integer.Name)
}

// commit datatype_definition_integer to the back repo (if it is already staged)
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) Commit(stage *Stage) *DATATYPE_DEFINITION_INTEGER {
	if _, ok := stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_INTEGER(datatype_definition_integer)
		}
	}
	return datatype_definition_integer
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) CommitVoid(stage *Stage) {
	datatype_definition_integer.Commit(stage)
}

// Checkout datatype_definition_integer to the back repo (if it is already staged)
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) Checkout(stage *Stage) *DATATYPE_DEFINITION_INTEGER {
	if _, ok := stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_INTEGER(datatype_definition_integer)
		}
	}
	return datatype_definition_integer
}

// for satisfaction of GongStruct interface
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GetName() (res string) {
	return datatype_definition_integer.Name
}

// Stage puts datatype_definition_real to the model stage
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) Stage(stage *Stage) *DATATYPE_DEFINITION_REAL {

	if _, ok := stage.DATATYPE_DEFINITION_REALs[datatype_definition_real]; !ok {
		stage.DATATYPE_DEFINITION_REALs[datatype_definition_real] = __member
		stage.DATATYPE_DEFINITION_REALMap_Staged_Order[datatype_definition_real] = stage.DATATYPE_DEFINITION_REALOrder
		stage.DATATYPE_DEFINITION_REALOrder++
	}
	stage.DATATYPE_DEFINITION_REALs_mapString[datatype_definition_real.Name] = datatype_definition_real

	return datatype_definition_real
}

// Unstage removes datatype_definition_real off the model stage
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) Unstage(stage *Stage) *DATATYPE_DEFINITION_REAL {
	delete(stage.DATATYPE_DEFINITION_REALs, datatype_definition_real)
	delete(stage.DATATYPE_DEFINITION_REALs_mapString, datatype_definition_real.Name)
	return datatype_definition_real
}

// UnstageVoid removes datatype_definition_real off the model stage
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) UnstageVoid(stage *Stage) {
	delete(stage.DATATYPE_DEFINITION_REALs, datatype_definition_real)
	delete(stage.DATATYPE_DEFINITION_REALs_mapString, datatype_definition_real.Name)
}

// commit datatype_definition_real to the back repo (if it is already staged)
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) Commit(stage *Stage) *DATATYPE_DEFINITION_REAL {
	if _, ok := stage.DATATYPE_DEFINITION_REALs[datatype_definition_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_REAL(datatype_definition_real)
		}
	}
	return datatype_definition_real
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) CommitVoid(stage *Stage) {
	datatype_definition_real.Commit(stage)
}

// Checkout datatype_definition_real to the back repo (if it is already staged)
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) Checkout(stage *Stage) *DATATYPE_DEFINITION_REAL {
	if _, ok := stage.DATATYPE_DEFINITION_REALs[datatype_definition_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_REAL(datatype_definition_real)
		}
	}
	return datatype_definition_real
}

// for satisfaction of GongStruct interface
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GetName() (res string) {
	return datatype_definition_real.Name
}

// Stage puts datatype_definition_string to the model stage
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) Stage(stage *Stage) *DATATYPE_DEFINITION_STRING {

	if _, ok := stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string]; !ok {
		stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string] = __member
		stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[datatype_definition_string] = stage.DATATYPE_DEFINITION_STRINGOrder
		stage.DATATYPE_DEFINITION_STRINGOrder++
	}
	stage.DATATYPE_DEFINITION_STRINGs_mapString[datatype_definition_string.Name] = datatype_definition_string

	return datatype_definition_string
}

// Unstage removes datatype_definition_string off the model stage
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) Unstage(stage *Stage) *DATATYPE_DEFINITION_STRING {
	delete(stage.DATATYPE_DEFINITION_STRINGs, datatype_definition_string)
	delete(stage.DATATYPE_DEFINITION_STRINGs_mapString, datatype_definition_string.Name)
	return datatype_definition_string
}

// UnstageVoid removes datatype_definition_string off the model stage
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) UnstageVoid(stage *Stage) {
	delete(stage.DATATYPE_DEFINITION_STRINGs, datatype_definition_string)
	delete(stage.DATATYPE_DEFINITION_STRINGs_mapString, datatype_definition_string.Name)
}

// commit datatype_definition_string to the back repo (if it is already staged)
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) Commit(stage *Stage) *DATATYPE_DEFINITION_STRING {
	if _, ok := stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_STRING(datatype_definition_string)
		}
	}
	return datatype_definition_string
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) CommitVoid(stage *Stage) {
	datatype_definition_string.Commit(stage)
}

// Checkout datatype_definition_string to the back repo (if it is already staged)
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) Checkout(stage *Stage) *DATATYPE_DEFINITION_STRING {
	if _, ok := stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_STRING(datatype_definition_string)
		}
	}
	return datatype_definition_string
}

// for satisfaction of GongStruct interface
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GetName() (res string) {
	return datatype_definition_string.Name
}

// Stage puts datatype_definition_xhtml to the model stage
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) Stage(stage *Stage) *DATATYPE_DEFINITION_XHTML {

	if _, ok := stage.DATATYPE_DEFINITION_XHTMLs[datatype_definition_xhtml]; !ok {
		stage.DATATYPE_DEFINITION_XHTMLs[datatype_definition_xhtml] = __member
		stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[datatype_definition_xhtml] = stage.DATATYPE_DEFINITION_XHTMLOrder
		stage.DATATYPE_DEFINITION_XHTMLOrder++
	}
	stage.DATATYPE_DEFINITION_XHTMLs_mapString[datatype_definition_xhtml.Name] = datatype_definition_xhtml

	return datatype_definition_xhtml
}

// Unstage removes datatype_definition_xhtml off the model stage
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) Unstage(stage *Stage) *DATATYPE_DEFINITION_XHTML {
	delete(stage.DATATYPE_DEFINITION_XHTMLs, datatype_definition_xhtml)
	delete(stage.DATATYPE_DEFINITION_XHTMLs_mapString, datatype_definition_xhtml.Name)
	return datatype_definition_xhtml
}

// UnstageVoid removes datatype_definition_xhtml off the model stage
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) UnstageVoid(stage *Stage) {
	delete(stage.DATATYPE_DEFINITION_XHTMLs, datatype_definition_xhtml)
	delete(stage.DATATYPE_DEFINITION_XHTMLs_mapString, datatype_definition_xhtml.Name)
}

// commit datatype_definition_xhtml to the back repo (if it is already staged)
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) Commit(stage *Stage) *DATATYPE_DEFINITION_XHTML {
	if _, ok := stage.DATATYPE_DEFINITION_XHTMLs[datatype_definition_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml)
		}
	}
	return datatype_definition_xhtml
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) CommitVoid(stage *Stage) {
	datatype_definition_xhtml.Commit(stage)
}

// Checkout datatype_definition_xhtml to the back repo (if it is already staged)
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) Checkout(stage *Stage) *DATATYPE_DEFINITION_XHTML {
	if _, ok := stage.DATATYPE_DEFINITION_XHTMLs[datatype_definition_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml)
		}
	}
	return datatype_definition_xhtml
}

// for satisfaction of GongStruct interface
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GetName() (res string) {
	return datatype_definition_xhtml.Name
}

// Stage puts embedded_value to the model stage
func (embedded_value *EMBEDDED_VALUE) Stage(stage *Stage) *EMBEDDED_VALUE {

	if _, ok := stage.EMBEDDED_VALUEs[embedded_value]; !ok {
		stage.EMBEDDED_VALUEs[embedded_value] = __member
		stage.EMBEDDED_VALUEMap_Staged_Order[embedded_value] = stage.EMBEDDED_VALUEOrder
		stage.EMBEDDED_VALUEOrder++
	}
	stage.EMBEDDED_VALUEs_mapString[embedded_value.Name] = embedded_value

	return embedded_value
}

// Unstage removes embedded_value off the model stage
func (embedded_value *EMBEDDED_VALUE) Unstage(stage *Stage) *EMBEDDED_VALUE {
	delete(stage.EMBEDDED_VALUEs, embedded_value)
	delete(stage.EMBEDDED_VALUEs_mapString, embedded_value.Name)
	return embedded_value
}

// UnstageVoid removes embedded_value off the model stage
func (embedded_value *EMBEDDED_VALUE) UnstageVoid(stage *Stage) {
	delete(stage.EMBEDDED_VALUEs, embedded_value)
	delete(stage.EMBEDDED_VALUEs_mapString, embedded_value.Name)
}

// commit embedded_value to the back repo (if it is already staged)
func (embedded_value *EMBEDDED_VALUE) Commit(stage *Stage) *EMBEDDED_VALUE {
	if _, ok := stage.EMBEDDED_VALUEs[embedded_value]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEMBEDDED_VALUE(embedded_value)
		}
	}
	return embedded_value
}

func (embedded_value *EMBEDDED_VALUE) CommitVoid(stage *Stage) {
	embedded_value.Commit(stage)
}

// Checkout embedded_value to the back repo (if it is already staged)
func (embedded_value *EMBEDDED_VALUE) Checkout(stage *Stage) *EMBEDDED_VALUE {
	if _, ok := stage.EMBEDDED_VALUEs[embedded_value]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEMBEDDED_VALUE(embedded_value)
		}
	}
	return embedded_value
}

// for satisfaction of GongStruct interface
func (embedded_value *EMBEDDED_VALUE) GetName() (res string) {
	return embedded_value.Name
}

// Stage puts enum_value to the model stage
func (enum_value *ENUM_VALUE) Stage(stage *Stage) *ENUM_VALUE {

	if _, ok := stage.ENUM_VALUEs[enum_value]; !ok {
		stage.ENUM_VALUEs[enum_value] = __member
		stage.ENUM_VALUEMap_Staged_Order[enum_value] = stage.ENUM_VALUEOrder
		stage.ENUM_VALUEOrder++
	}
	stage.ENUM_VALUEs_mapString[enum_value.Name] = enum_value

	return enum_value
}

// Unstage removes enum_value off the model stage
func (enum_value *ENUM_VALUE) Unstage(stage *Stage) *ENUM_VALUE {
	delete(stage.ENUM_VALUEs, enum_value)
	delete(stage.ENUM_VALUEs_mapString, enum_value.Name)
	return enum_value
}

// UnstageVoid removes enum_value off the model stage
func (enum_value *ENUM_VALUE) UnstageVoid(stage *Stage) {
	delete(stage.ENUM_VALUEs, enum_value)
	delete(stage.ENUM_VALUEs_mapString, enum_value.Name)
}

// commit enum_value to the back repo (if it is already staged)
func (enum_value *ENUM_VALUE) Commit(stage *Stage) *ENUM_VALUE {
	if _, ok := stage.ENUM_VALUEs[enum_value]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitENUM_VALUE(enum_value)
		}
	}
	return enum_value
}

func (enum_value *ENUM_VALUE) CommitVoid(stage *Stage) {
	enum_value.Commit(stage)
}

// Checkout enum_value to the back repo (if it is already staged)
func (enum_value *ENUM_VALUE) Checkout(stage *Stage) *ENUM_VALUE {
	if _, ok := stage.ENUM_VALUEs[enum_value]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutENUM_VALUE(enum_value)
		}
	}
	return enum_value
}

// for satisfaction of GongStruct interface
func (enum_value *ENUM_VALUE) GetName() (res string) {
	return enum_value.Name
}

// Stage puts relation_group to the model stage
func (relation_group *RELATION_GROUP) Stage(stage *Stage) *RELATION_GROUP {

	if _, ok := stage.RELATION_GROUPs[relation_group]; !ok {
		stage.RELATION_GROUPs[relation_group] = __member
		stage.RELATION_GROUPMap_Staged_Order[relation_group] = stage.RELATION_GROUPOrder
		stage.RELATION_GROUPOrder++
	}
	stage.RELATION_GROUPs_mapString[relation_group.Name] = relation_group

	return relation_group
}

// Unstage removes relation_group off the model stage
func (relation_group *RELATION_GROUP) Unstage(stage *Stage) *RELATION_GROUP {
	delete(stage.RELATION_GROUPs, relation_group)
	delete(stage.RELATION_GROUPs_mapString, relation_group.Name)
	return relation_group
}

// UnstageVoid removes relation_group off the model stage
func (relation_group *RELATION_GROUP) UnstageVoid(stage *Stage) {
	delete(stage.RELATION_GROUPs, relation_group)
	delete(stage.RELATION_GROUPs_mapString, relation_group.Name)
}

// commit relation_group to the back repo (if it is already staged)
func (relation_group *RELATION_GROUP) Commit(stage *Stage) *RELATION_GROUP {
	if _, ok := stage.RELATION_GROUPs[relation_group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRELATION_GROUP(relation_group)
		}
	}
	return relation_group
}

func (relation_group *RELATION_GROUP) CommitVoid(stage *Stage) {
	relation_group.Commit(stage)
}

// Checkout relation_group to the back repo (if it is already staged)
func (relation_group *RELATION_GROUP) Checkout(stage *Stage) *RELATION_GROUP {
	if _, ok := stage.RELATION_GROUPs[relation_group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRELATION_GROUP(relation_group)
		}
	}
	return relation_group
}

// for satisfaction of GongStruct interface
func (relation_group *RELATION_GROUP) GetName() (res string) {
	return relation_group.Name
}

// Stage puts relation_group_type to the model stage
func (relation_group_type *RELATION_GROUP_TYPE) Stage(stage *Stage) *RELATION_GROUP_TYPE {

	if _, ok := stage.RELATION_GROUP_TYPEs[relation_group_type]; !ok {
		stage.RELATION_GROUP_TYPEs[relation_group_type] = __member
		stage.RELATION_GROUP_TYPEMap_Staged_Order[relation_group_type] = stage.RELATION_GROUP_TYPEOrder
		stage.RELATION_GROUP_TYPEOrder++
	}
	stage.RELATION_GROUP_TYPEs_mapString[relation_group_type.Name] = relation_group_type

	return relation_group_type
}

// Unstage removes relation_group_type off the model stage
func (relation_group_type *RELATION_GROUP_TYPE) Unstage(stage *Stage) *RELATION_GROUP_TYPE {
	delete(stage.RELATION_GROUP_TYPEs, relation_group_type)
	delete(stage.RELATION_GROUP_TYPEs_mapString, relation_group_type.Name)
	return relation_group_type
}

// UnstageVoid removes relation_group_type off the model stage
func (relation_group_type *RELATION_GROUP_TYPE) UnstageVoid(stage *Stage) {
	delete(stage.RELATION_GROUP_TYPEs, relation_group_type)
	delete(stage.RELATION_GROUP_TYPEs_mapString, relation_group_type.Name)
}

// commit relation_group_type to the back repo (if it is already staged)
func (relation_group_type *RELATION_GROUP_TYPE) Commit(stage *Stage) *RELATION_GROUP_TYPE {
	if _, ok := stage.RELATION_GROUP_TYPEs[relation_group_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRELATION_GROUP_TYPE(relation_group_type)
		}
	}
	return relation_group_type
}

func (relation_group_type *RELATION_GROUP_TYPE) CommitVoid(stage *Stage) {
	relation_group_type.Commit(stage)
}

// Checkout relation_group_type to the back repo (if it is already staged)
func (relation_group_type *RELATION_GROUP_TYPE) Checkout(stage *Stage) *RELATION_GROUP_TYPE {
	if _, ok := stage.RELATION_GROUP_TYPEs[relation_group_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRELATION_GROUP_TYPE(relation_group_type)
		}
	}
	return relation_group_type
}

// for satisfaction of GongStruct interface
func (relation_group_type *RELATION_GROUP_TYPE) GetName() (res string) {
	return relation_group_type.Name
}

// Stage puts req_if to the model stage
func (req_if *REQ_IF) Stage(stage *Stage) *REQ_IF {

	if _, ok := stage.REQ_IFs[req_if]; !ok {
		stage.REQ_IFs[req_if] = __member
		stage.REQ_IFMap_Staged_Order[req_if] = stage.REQ_IFOrder
		stage.REQ_IFOrder++
	}
	stage.REQ_IFs_mapString[req_if.Name] = req_if

	return req_if
}

// Unstage removes req_if off the model stage
func (req_if *REQ_IF) Unstage(stage *Stage) *REQ_IF {
	delete(stage.REQ_IFs, req_if)
	delete(stage.REQ_IFs_mapString, req_if.Name)
	return req_if
}

// UnstageVoid removes req_if off the model stage
func (req_if *REQ_IF) UnstageVoid(stage *Stage) {
	delete(stage.REQ_IFs, req_if)
	delete(stage.REQ_IFs_mapString, req_if.Name)
}

// commit req_if to the back repo (if it is already staged)
func (req_if *REQ_IF) Commit(stage *Stage) *REQ_IF {
	if _, ok := stage.REQ_IFs[req_if]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQ_IF(req_if)
		}
	}
	return req_if
}

func (req_if *REQ_IF) CommitVoid(stage *Stage) {
	req_if.Commit(stage)
}

// Checkout req_if to the back repo (if it is already staged)
func (req_if *REQ_IF) Checkout(stage *Stage) *REQ_IF {
	if _, ok := stage.REQ_IFs[req_if]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQ_IF(req_if)
		}
	}
	return req_if
}

// for satisfaction of GongStruct interface
func (req_if *REQ_IF) GetName() (res string) {
	return req_if.Name
}

// Stage puts req_if_content to the model stage
func (req_if_content *REQ_IF_CONTENT) Stage(stage *Stage) *REQ_IF_CONTENT {

	if _, ok := stage.REQ_IF_CONTENTs[req_if_content]; !ok {
		stage.REQ_IF_CONTENTs[req_if_content] = __member
		stage.REQ_IF_CONTENTMap_Staged_Order[req_if_content] = stage.REQ_IF_CONTENTOrder
		stage.REQ_IF_CONTENTOrder++
	}
	stage.REQ_IF_CONTENTs_mapString[req_if_content.Name] = req_if_content

	return req_if_content
}

// Unstage removes req_if_content off the model stage
func (req_if_content *REQ_IF_CONTENT) Unstage(stage *Stage) *REQ_IF_CONTENT {
	delete(stage.REQ_IF_CONTENTs, req_if_content)
	delete(stage.REQ_IF_CONTENTs_mapString, req_if_content.Name)
	return req_if_content
}

// UnstageVoid removes req_if_content off the model stage
func (req_if_content *REQ_IF_CONTENT) UnstageVoid(stage *Stage) {
	delete(stage.REQ_IF_CONTENTs, req_if_content)
	delete(stage.REQ_IF_CONTENTs_mapString, req_if_content.Name)
}

// commit req_if_content to the back repo (if it is already staged)
func (req_if_content *REQ_IF_CONTENT) Commit(stage *Stage) *REQ_IF_CONTENT {
	if _, ok := stage.REQ_IF_CONTENTs[req_if_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQ_IF_CONTENT(req_if_content)
		}
	}
	return req_if_content
}

func (req_if_content *REQ_IF_CONTENT) CommitVoid(stage *Stage) {
	req_if_content.Commit(stage)
}

// Checkout req_if_content to the back repo (if it is already staged)
func (req_if_content *REQ_IF_CONTENT) Checkout(stage *Stage) *REQ_IF_CONTENT {
	if _, ok := stage.REQ_IF_CONTENTs[req_if_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQ_IF_CONTENT(req_if_content)
		}
	}
	return req_if_content
}

// for satisfaction of GongStruct interface
func (req_if_content *REQ_IF_CONTENT) GetName() (res string) {
	return req_if_content.Name
}

// Stage puts req_if_header to the model stage
func (req_if_header *REQ_IF_HEADER) Stage(stage *Stage) *REQ_IF_HEADER {

	if _, ok := stage.REQ_IF_HEADERs[req_if_header]; !ok {
		stage.REQ_IF_HEADERs[req_if_header] = __member
		stage.REQ_IF_HEADERMap_Staged_Order[req_if_header] = stage.REQ_IF_HEADEROrder
		stage.REQ_IF_HEADEROrder++
	}
	stage.REQ_IF_HEADERs_mapString[req_if_header.Name] = req_if_header

	return req_if_header
}

// Unstage removes req_if_header off the model stage
func (req_if_header *REQ_IF_HEADER) Unstage(stage *Stage) *REQ_IF_HEADER {
	delete(stage.REQ_IF_HEADERs, req_if_header)
	delete(stage.REQ_IF_HEADERs_mapString, req_if_header.Name)
	return req_if_header
}

// UnstageVoid removes req_if_header off the model stage
func (req_if_header *REQ_IF_HEADER) UnstageVoid(stage *Stage) {
	delete(stage.REQ_IF_HEADERs, req_if_header)
	delete(stage.REQ_IF_HEADERs_mapString, req_if_header.Name)
}

// commit req_if_header to the back repo (if it is already staged)
func (req_if_header *REQ_IF_HEADER) Commit(stage *Stage) *REQ_IF_HEADER {
	if _, ok := stage.REQ_IF_HEADERs[req_if_header]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQ_IF_HEADER(req_if_header)
		}
	}
	return req_if_header
}

func (req_if_header *REQ_IF_HEADER) CommitVoid(stage *Stage) {
	req_if_header.Commit(stage)
}

// Checkout req_if_header to the back repo (if it is already staged)
func (req_if_header *REQ_IF_HEADER) Checkout(stage *Stage) *REQ_IF_HEADER {
	if _, ok := stage.REQ_IF_HEADERs[req_if_header]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQ_IF_HEADER(req_if_header)
		}
	}
	return req_if_header
}

// for satisfaction of GongStruct interface
func (req_if_header *REQ_IF_HEADER) GetName() (res string) {
	return req_if_header.Name
}

// Stage puts req_if_tool_extension to the model stage
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) Stage(stage *Stage) *REQ_IF_TOOL_EXTENSION {

	if _, ok := stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension]; !ok {
		stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension] = __member
		stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[req_if_tool_extension] = stage.REQ_IF_TOOL_EXTENSIONOrder
		stage.REQ_IF_TOOL_EXTENSIONOrder++
	}
	stage.REQ_IF_TOOL_EXTENSIONs_mapString[req_if_tool_extension.Name] = req_if_tool_extension

	return req_if_tool_extension
}

// Unstage removes req_if_tool_extension off the model stage
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) Unstage(stage *Stage) *REQ_IF_TOOL_EXTENSION {
	delete(stage.REQ_IF_TOOL_EXTENSIONs, req_if_tool_extension)
	delete(stage.REQ_IF_TOOL_EXTENSIONs_mapString, req_if_tool_extension.Name)
	return req_if_tool_extension
}

// UnstageVoid removes req_if_tool_extension off the model stage
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) UnstageVoid(stage *Stage) {
	delete(stage.REQ_IF_TOOL_EXTENSIONs, req_if_tool_extension)
	delete(stage.REQ_IF_TOOL_EXTENSIONs_mapString, req_if_tool_extension.Name)
}

// commit req_if_tool_extension to the back repo (if it is already staged)
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) Commit(stage *Stage) *REQ_IF_TOOL_EXTENSION {
	if _, ok := stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQ_IF_TOOL_EXTENSION(req_if_tool_extension)
		}
	}
	return req_if_tool_extension
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) CommitVoid(stage *Stage) {
	req_if_tool_extension.Commit(stage)
}

// Checkout req_if_tool_extension to the back repo (if it is already staged)
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) Checkout(stage *Stage) *REQ_IF_TOOL_EXTENSION {
	if _, ok := stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQ_IF_TOOL_EXTENSION(req_if_tool_extension)
		}
	}
	return req_if_tool_extension
}

// for satisfaction of GongStruct interface
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GetName() (res string) {
	return req_if_tool_extension.Name
}

// Stage puts specification to the model stage
func (specification *SPECIFICATION) Stage(stage *Stage) *SPECIFICATION {

	if _, ok := stage.SPECIFICATIONs[specification]; !ok {
		stage.SPECIFICATIONs[specification] = __member
		stage.SPECIFICATIONMap_Staged_Order[specification] = stage.SPECIFICATIONOrder
		stage.SPECIFICATIONOrder++
	}
	stage.SPECIFICATIONs_mapString[specification.Name] = specification

	return specification
}

// Unstage removes specification off the model stage
func (specification *SPECIFICATION) Unstage(stage *Stage) *SPECIFICATION {
	delete(stage.SPECIFICATIONs, specification)
	delete(stage.SPECIFICATIONs_mapString, specification.Name)
	return specification
}

// UnstageVoid removes specification off the model stage
func (specification *SPECIFICATION) UnstageVoid(stage *Stage) {
	delete(stage.SPECIFICATIONs, specification)
	delete(stage.SPECIFICATIONs_mapString, specification.Name)
}

// commit specification to the back repo (if it is already staged)
func (specification *SPECIFICATION) Commit(stage *Stage) *SPECIFICATION {
	if _, ok := stage.SPECIFICATIONs[specification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECIFICATION(specification)
		}
	}
	return specification
}

func (specification *SPECIFICATION) CommitVoid(stage *Stage) {
	specification.Commit(stage)
}

// Checkout specification to the back repo (if it is already staged)
func (specification *SPECIFICATION) Checkout(stage *Stage) *SPECIFICATION {
	if _, ok := stage.SPECIFICATIONs[specification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECIFICATION(specification)
		}
	}
	return specification
}

// for satisfaction of GongStruct interface
func (specification *SPECIFICATION) GetName() (res string) {
	return specification.Name
}

// Stage puts specification_type to the model stage
func (specification_type *SPECIFICATION_TYPE) Stage(stage *Stage) *SPECIFICATION_TYPE {

	if _, ok := stage.SPECIFICATION_TYPEs[specification_type]; !ok {
		stage.SPECIFICATION_TYPEs[specification_type] = __member
		stage.SPECIFICATION_TYPEMap_Staged_Order[specification_type] = stage.SPECIFICATION_TYPEOrder
		stage.SPECIFICATION_TYPEOrder++
	}
	stage.SPECIFICATION_TYPEs_mapString[specification_type.Name] = specification_type

	return specification_type
}

// Unstage removes specification_type off the model stage
func (specification_type *SPECIFICATION_TYPE) Unstage(stage *Stage) *SPECIFICATION_TYPE {
	delete(stage.SPECIFICATION_TYPEs, specification_type)
	delete(stage.SPECIFICATION_TYPEs_mapString, specification_type.Name)
	return specification_type
}

// UnstageVoid removes specification_type off the model stage
func (specification_type *SPECIFICATION_TYPE) UnstageVoid(stage *Stage) {
	delete(stage.SPECIFICATION_TYPEs, specification_type)
	delete(stage.SPECIFICATION_TYPEs_mapString, specification_type.Name)
}

// commit specification_type to the back repo (if it is already staged)
func (specification_type *SPECIFICATION_TYPE) Commit(stage *Stage) *SPECIFICATION_TYPE {
	if _, ok := stage.SPECIFICATION_TYPEs[specification_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECIFICATION_TYPE(specification_type)
		}
	}
	return specification_type
}

func (specification_type *SPECIFICATION_TYPE) CommitVoid(stage *Stage) {
	specification_type.Commit(stage)
}

// Checkout specification_type to the back repo (if it is already staged)
func (specification_type *SPECIFICATION_TYPE) Checkout(stage *Stage) *SPECIFICATION_TYPE {
	if _, ok := stage.SPECIFICATION_TYPEs[specification_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECIFICATION_TYPE(specification_type)
		}
	}
	return specification_type
}

// for satisfaction of GongStruct interface
func (specification_type *SPECIFICATION_TYPE) GetName() (res string) {
	return specification_type.Name
}

// Stage puts spec_hierarchy to the model stage
func (spec_hierarchy *SPEC_HIERARCHY) Stage(stage *Stage) *SPEC_HIERARCHY {

	if _, ok := stage.SPEC_HIERARCHYs[spec_hierarchy]; !ok {
		stage.SPEC_HIERARCHYs[spec_hierarchy] = __member
		stage.SPEC_HIERARCHYMap_Staged_Order[spec_hierarchy] = stage.SPEC_HIERARCHYOrder
		stage.SPEC_HIERARCHYOrder++
	}
	stage.SPEC_HIERARCHYs_mapString[spec_hierarchy.Name] = spec_hierarchy

	return spec_hierarchy
}

// Unstage removes spec_hierarchy off the model stage
func (spec_hierarchy *SPEC_HIERARCHY) Unstage(stage *Stage) *SPEC_HIERARCHY {
	delete(stage.SPEC_HIERARCHYs, spec_hierarchy)
	delete(stage.SPEC_HIERARCHYs_mapString, spec_hierarchy.Name)
	return spec_hierarchy
}

// UnstageVoid removes spec_hierarchy off the model stage
func (spec_hierarchy *SPEC_HIERARCHY) UnstageVoid(stage *Stage) {
	delete(stage.SPEC_HIERARCHYs, spec_hierarchy)
	delete(stage.SPEC_HIERARCHYs_mapString, spec_hierarchy.Name)
}

// commit spec_hierarchy to the back repo (if it is already staged)
func (spec_hierarchy *SPEC_HIERARCHY) Commit(stage *Stage) *SPEC_HIERARCHY {
	if _, ok := stage.SPEC_HIERARCHYs[spec_hierarchy]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_HIERARCHY(spec_hierarchy)
		}
	}
	return spec_hierarchy
}

func (spec_hierarchy *SPEC_HIERARCHY) CommitVoid(stage *Stage) {
	spec_hierarchy.Commit(stage)
}

// Checkout spec_hierarchy to the back repo (if it is already staged)
func (spec_hierarchy *SPEC_HIERARCHY) Checkout(stage *Stage) *SPEC_HIERARCHY {
	if _, ok := stage.SPEC_HIERARCHYs[spec_hierarchy]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_HIERARCHY(spec_hierarchy)
		}
	}
	return spec_hierarchy
}

// for satisfaction of GongStruct interface
func (spec_hierarchy *SPEC_HIERARCHY) GetName() (res string) {
	return spec_hierarchy.Name
}

// Stage puts spec_object to the model stage
func (spec_object *SPEC_OBJECT) Stage(stage *Stage) *SPEC_OBJECT {

	if _, ok := stage.SPEC_OBJECTs[spec_object]; !ok {
		stage.SPEC_OBJECTs[spec_object] = __member
		stage.SPEC_OBJECTMap_Staged_Order[spec_object] = stage.SPEC_OBJECTOrder
		stage.SPEC_OBJECTOrder++
	}
	stage.SPEC_OBJECTs_mapString[spec_object.Name] = spec_object

	return spec_object
}

// Unstage removes spec_object off the model stage
func (spec_object *SPEC_OBJECT) Unstage(stage *Stage) *SPEC_OBJECT {
	delete(stage.SPEC_OBJECTs, spec_object)
	delete(stage.SPEC_OBJECTs_mapString, spec_object.Name)
	return spec_object
}

// UnstageVoid removes spec_object off the model stage
func (spec_object *SPEC_OBJECT) UnstageVoid(stage *Stage) {
	delete(stage.SPEC_OBJECTs, spec_object)
	delete(stage.SPEC_OBJECTs_mapString, spec_object.Name)
}

// commit spec_object to the back repo (if it is already staged)
func (spec_object *SPEC_OBJECT) Commit(stage *Stage) *SPEC_OBJECT {
	if _, ok := stage.SPEC_OBJECTs[spec_object]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_OBJECT(spec_object)
		}
	}
	return spec_object
}

func (spec_object *SPEC_OBJECT) CommitVoid(stage *Stage) {
	spec_object.Commit(stage)
}

// Checkout spec_object to the back repo (if it is already staged)
func (spec_object *SPEC_OBJECT) Checkout(stage *Stage) *SPEC_OBJECT {
	if _, ok := stage.SPEC_OBJECTs[spec_object]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_OBJECT(spec_object)
		}
	}
	return spec_object
}

// for satisfaction of GongStruct interface
func (spec_object *SPEC_OBJECT) GetName() (res string) {
	return spec_object.Name
}

// Stage puts spec_object_type to the model stage
func (spec_object_type *SPEC_OBJECT_TYPE) Stage(stage *Stage) *SPEC_OBJECT_TYPE {

	if _, ok := stage.SPEC_OBJECT_TYPEs[spec_object_type]; !ok {
		stage.SPEC_OBJECT_TYPEs[spec_object_type] = __member
		stage.SPEC_OBJECT_TYPEMap_Staged_Order[spec_object_type] = stage.SPEC_OBJECT_TYPEOrder
		stage.SPEC_OBJECT_TYPEOrder++
	}
	stage.SPEC_OBJECT_TYPEs_mapString[spec_object_type.Name] = spec_object_type

	return spec_object_type
}

// Unstage removes spec_object_type off the model stage
func (spec_object_type *SPEC_OBJECT_TYPE) Unstage(stage *Stage) *SPEC_OBJECT_TYPE {
	delete(stage.SPEC_OBJECT_TYPEs, spec_object_type)
	delete(stage.SPEC_OBJECT_TYPEs_mapString, spec_object_type.Name)
	return spec_object_type
}

// UnstageVoid removes spec_object_type off the model stage
func (spec_object_type *SPEC_OBJECT_TYPE) UnstageVoid(stage *Stage) {
	delete(stage.SPEC_OBJECT_TYPEs, spec_object_type)
	delete(stage.SPEC_OBJECT_TYPEs_mapString, spec_object_type.Name)
}

// commit spec_object_type to the back repo (if it is already staged)
func (spec_object_type *SPEC_OBJECT_TYPE) Commit(stage *Stage) *SPEC_OBJECT_TYPE {
	if _, ok := stage.SPEC_OBJECT_TYPEs[spec_object_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_OBJECT_TYPE(spec_object_type)
		}
	}
	return spec_object_type
}

func (spec_object_type *SPEC_OBJECT_TYPE) CommitVoid(stage *Stage) {
	spec_object_type.Commit(stage)
}

// Checkout spec_object_type to the back repo (if it is already staged)
func (spec_object_type *SPEC_OBJECT_TYPE) Checkout(stage *Stage) *SPEC_OBJECT_TYPE {
	if _, ok := stage.SPEC_OBJECT_TYPEs[spec_object_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_OBJECT_TYPE(spec_object_type)
		}
	}
	return spec_object_type
}

// for satisfaction of GongStruct interface
func (spec_object_type *SPEC_OBJECT_TYPE) GetName() (res string) {
	return spec_object_type.Name
}

// Stage puts spec_relation to the model stage
func (spec_relation *SPEC_RELATION) Stage(stage *Stage) *SPEC_RELATION {

	if _, ok := stage.SPEC_RELATIONs[spec_relation]; !ok {
		stage.SPEC_RELATIONs[spec_relation] = __member
		stage.SPEC_RELATIONMap_Staged_Order[spec_relation] = stage.SPEC_RELATIONOrder
		stage.SPEC_RELATIONOrder++
	}
	stage.SPEC_RELATIONs_mapString[spec_relation.Name] = spec_relation

	return spec_relation
}

// Unstage removes spec_relation off the model stage
func (spec_relation *SPEC_RELATION) Unstage(stage *Stage) *SPEC_RELATION {
	delete(stage.SPEC_RELATIONs, spec_relation)
	delete(stage.SPEC_RELATIONs_mapString, spec_relation.Name)
	return spec_relation
}

// UnstageVoid removes spec_relation off the model stage
func (spec_relation *SPEC_RELATION) UnstageVoid(stage *Stage) {
	delete(stage.SPEC_RELATIONs, spec_relation)
	delete(stage.SPEC_RELATIONs_mapString, spec_relation.Name)
}

// commit spec_relation to the back repo (if it is already staged)
func (spec_relation *SPEC_RELATION) Commit(stage *Stage) *SPEC_RELATION {
	if _, ok := stage.SPEC_RELATIONs[spec_relation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_RELATION(spec_relation)
		}
	}
	return spec_relation
}

func (spec_relation *SPEC_RELATION) CommitVoid(stage *Stage) {
	spec_relation.Commit(stage)
}

// Checkout spec_relation to the back repo (if it is already staged)
func (spec_relation *SPEC_RELATION) Checkout(stage *Stage) *SPEC_RELATION {
	if _, ok := stage.SPEC_RELATIONs[spec_relation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_RELATION(spec_relation)
		}
	}
	return spec_relation
}

// for satisfaction of GongStruct interface
func (spec_relation *SPEC_RELATION) GetName() (res string) {
	return spec_relation.Name
}

// Stage puts spec_relation_type to the model stage
func (spec_relation_type *SPEC_RELATION_TYPE) Stage(stage *Stage) *SPEC_RELATION_TYPE {

	if _, ok := stage.SPEC_RELATION_TYPEs[spec_relation_type]; !ok {
		stage.SPEC_RELATION_TYPEs[spec_relation_type] = __member
		stage.SPEC_RELATION_TYPEMap_Staged_Order[spec_relation_type] = stage.SPEC_RELATION_TYPEOrder
		stage.SPEC_RELATION_TYPEOrder++
	}
	stage.SPEC_RELATION_TYPEs_mapString[spec_relation_type.Name] = spec_relation_type

	return spec_relation_type
}

// Unstage removes spec_relation_type off the model stage
func (spec_relation_type *SPEC_RELATION_TYPE) Unstage(stage *Stage) *SPEC_RELATION_TYPE {
	delete(stage.SPEC_RELATION_TYPEs, spec_relation_type)
	delete(stage.SPEC_RELATION_TYPEs_mapString, spec_relation_type.Name)
	return spec_relation_type
}

// UnstageVoid removes spec_relation_type off the model stage
func (spec_relation_type *SPEC_RELATION_TYPE) UnstageVoid(stage *Stage) {
	delete(stage.SPEC_RELATION_TYPEs, spec_relation_type)
	delete(stage.SPEC_RELATION_TYPEs_mapString, spec_relation_type.Name)
}

// commit spec_relation_type to the back repo (if it is already staged)
func (spec_relation_type *SPEC_RELATION_TYPE) Commit(stage *Stage) *SPEC_RELATION_TYPE {
	if _, ok := stage.SPEC_RELATION_TYPEs[spec_relation_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_RELATION_TYPE(spec_relation_type)
		}
	}
	return spec_relation_type
}

func (spec_relation_type *SPEC_RELATION_TYPE) CommitVoid(stage *Stage) {
	spec_relation_type.Commit(stage)
}

// Checkout spec_relation_type to the back repo (if it is already staged)
func (spec_relation_type *SPEC_RELATION_TYPE) Checkout(stage *Stage) *SPEC_RELATION_TYPE {
	if _, ok := stage.SPEC_RELATION_TYPEs[spec_relation_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_RELATION_TYPE(spec_relation_type)
		}
	}
	return spec_relation_type
}

// for satisfaction of GongStruct interface
func (spec_relation_type *SPEC_RELATION_TYPE) GetName() (res string) {
	return spec_relation_type.Name
}

// Stage puts xhtml_content to the model stage
func (xhtml_content *XHTML_CONTENT) Stage(stage *Stage) *XHTML_CONTENT {

	if _, ok := stage.XHTML_CONTENTs[xhtml_content]; !ok {
		stage.XHTML_CONTENTs[xhtml_content] = __member
		stage.XHTML_CONTENTMap_Staged_Order[xhtml_content] = stage.XHTML_CONTENTOrder
		stage.XHTML_CONTENTOrder++
	}
	stage.XHTML_CONTENTs_mapString[xhtml_content.Name] = xhtml_content

	return xhtml_content
}

// Unstage removes xhtml_content off the model stage
func (xhtml_content *XHTML_CONTENT) Unstage(stage *Stage) *XHTML_CONTENT {
	delete(stage.XHTML_CONTENTs, xhtml_content)
	delete(stage.XHTML_CONTENTs_mapString, xhtml_content.Name)
	return xhtml_content
}

// UnstageVoid removes xhtml_content off the model stage
func (xhtml_content *XHTML_CONTENT) UnstageVoid(stage *Stage) {
	delete(stage.XHTML_CONTENTs, xhtml_content)
	delete(stage.XHTML_CONTENTs_mapString, xhtml_content.Name)
}

// commit xhtml_content to the back repo (if it is already staged)
func (xhtml_content *XHTML_CONTENT) Commit(stage *Stage) *XHTML_CONTENT {
	if _, ok := stage.XHTML_CONTENTs[xhtml_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXHTML_CONTENT(xhtml_content)
		}
	}
	return xhtml_content
}

func (xhtml_content *XHTML_CONTENT) CommitVoid(stage *Stage) {
	xhtml_content.Commit(stage)
}

// Checkout xhtml_content to the back repo (if it is already staged)
func (xhtml_content *XHTML_CONTENT) Checkout(stage *Stage) *XHTML_CONTENT {
	if _, ok := stage.XHTML_CONTENTs[xhtml_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXHTML_CONTENT(xhtml_content)
		}
	}
	return xhtml_content
}

// for satisfaction of GongStruct interface
func (xhtml_content *XHTML_CONTENT) GetName() (res string) {
	return xhtml_content.Name
}

// Stage puts xhtml_inlpres_type to the model stage
func (xhtml_inlpres_type *Xhtml_InlPres_type) Stage(stage *Stage) *Xhtml_InlPres_type {

	if _, ok := stage.Xhtml_InlPres_types[xhtml_inlpres_type]; !ok {
		stage.Xhtml_InlPres_types[xhtml_inlpres_type] = __member
		stage.Xhtml_InlPres_typeMap_Staged_Order[xhtml_inlpres_type] = stage.Xhtml_InlPres_typeOrder
		stage.Xhtml_InlPres_typeOrder++
	}
	stage.Xhtml_InlPres_types_mapString[xhtml_inlpres_type.Name] = xhtml_inlpres_type

	return xhtml_inlpres_type
}

// Unstage removes xhtml_inlpres_type off the model stage
func (xhtml_inlpres_type *Xhtml_InlPres_type) Unstage(stage *Stage) *Xhtml_InlPres_type {
	delete(stage.Xhtml_InlPres_types, xhtml_inlpres_type)
	delete(stage.Xhtml_InlPres_types_mapString, xhtml_inlpres_type.Name)
	return xhtml_inlpres_type
}

// UnstageVoid removes xhtml_inlpres_type off the model stage
func (xhtml_inlpres_type *Xhtml_InlPres_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_InlPres_types, xhtml_inlpres_type)
	delete(stage.Xhtml_InlPres_types_mapString, xhtml_inlpres_type.Name)
}

// commit xhtml_inlpres_type to the back repo (if it is already staged)
func (xhtml_inlpres_type *Xhtml_InlPres_type) Commit(stage *Stage) *Xhtml_InlPres_type {
	if _, ok := stage.Xhtml_InlPres_types[xhtml_inlpres_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_InlPres_type(xhtml_inlpres_type)
		}
	}
	return xhtml_inlpres_type
}

func (xhtml_inlpres_type *Xhtml_InlPres_type) CommitVoid(stage *Stage) {
	xhtml_inlpres_type.Commit(stage)
}

// Checkout xhtml_inlpres_type to the back repo (if it is already staged)
func (xhtml_inlpres_type *Xhtml_InlPres_type) Checkout(stage *Stage) *Xhtml_InlPres_type {
	if _, ok := stage.Xhtml_InlPres_types[xhtml_inlpres_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_InlPres_type(xhtml_inlpres_type)
		}
	}
	return xhtml_inlpres_type
}

// for satisfaction of GongStruct interface
func (xhtml_inlpres_type *Xhtml_InlPres_type) GetName() (res string) {
	return xhtml_inlpres_type.Name
}

// Stage puts xhtml_a_type to the model stage
func (xhtml_a_type *Xhtml_a_type) Stage(stage *Stage) *Xhtml_a_type {

	if _, ok := stage.Xhtml_a_types[xhtml_a_type]; !ok {
		stage.Xhtml_a_types[xhtml_a_type] = __member
		stage.Xhtml_a_typeMap_Staged_Order[xhtml_a_type] = stage.Xhtml_a_typeOrder
		stage.Xhtml_a_typeOrder++
	}
	stage.Xhtml_a_types_mapString[xhtml_a_type.Name] = xhtml_a_type

	return xhtml_a_type
}

// Unstage removes xhtml_a_type off the model stage
func (xhtml_a_type *Xhtml_a_type) Unstage(stage *Stage) *Xhtml_a_type {
	delete(stage.Xhtml_a_types, xhtml_a_type)
	delete(stage.Xhtml_a_types_mapString, xhtml_a_type.Name)
	return xhtml_a_type
}

// UnstageVoid removes xhtml_a_type off the model stage
func (xhtml_a_type *Xhtml_a_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_a_types, xhtml_a_type)
	delete(stage.Xhtml_a_types_mapString, xhtml_a_type.Name)
}

// commit xhtml_a_type to the back repo (if it is already staged)
func (xhtml_a_type *Xhtml_a_type) Commit(stage *Stage) *Xhtml_a_type {
	if _, ok := stage.Xhtml_a_types[xhtml_a_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_a_type(xhtml_a_type)
		}
	}
	return xhtml_a_type
}

func (xhtml_a_type *Xhtml_a_type) CommitVoid(stage *Stage) {
	xhtml_a_type.Commit(stage)
}

// Checkout xhtml_a_type to the back repo (if it is already staged)
func (xhtml_a_type *Xhtml_a_type) Checkout(stage *Stage) *Xhtml_a_type {
	if _, ok := stage.Xhtml_a_types[xhtml_a_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_a_type(xhtml_a_type)
		}
	}
	return xhtml_a_type
}

// for satisfaction of GongStruct interface
func (xhtml_a_type *Xhtml_a_type) GetName() (res string) {
	return xhtml_a_type.Name
}

// Stage puts xhtml_abbr_type to the model stage
func (xhtml_abbr_type *Xhtml_abbr_type) Stage(stage *Stage) *Xhtml_abbr_type {

	if _, ok := stage.Xhtml_abbr_types[xhtml_abbr_type]; !ok {
		stage.Xhtml_abbr_types[xhtml_abbr_type] = __member
		stage.Xhtml_abbr_typeMap_Staged_Order[xhtml_abbr_type] = stage.Xhtml_abbr_typeOrder
		stage.Xhtml_abbr_typeOrder++
	}
	stage.Xhtml_abbr_types_mapString[xhtml_abbr_type.Name] = xhtml_abbr_type

	return xhtml_abbr_type
}

// Unstage removes xhtml_abbr_type off the model stage
func (xhtml_abbr_type *Xhtml_abbr_type) Unstage(stage *Stage) *Xhtml_abbr_type {
	delete(stage.Xhtml_abbr_types, xhtml_abbr_type)
	delete(stage.Xhtml_abbr_types_mapString, xhtml_abbr_type.Name)
	return xhtml_abbr_type
}

// UnstageVoid removes xhtml_abbr_type off the model stage
func (xhtml_abbr_type *Xhtml_abbr_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_abbr_types, xhtml_abbr_type)
	delete(stage.Xhtml_abbr_types_mapString, xhtml_abbr_type.Name)
}

// commit xhtml_abbr_type to the back repo (if it is already staged)
func (xhtml_abbr_type *Xhtml_abbr_type) Commit(stage *Stage) *Xhtml_abbr_type {
	if _, ok := stage.Xhtml_abbr_types[xhtml_abbr_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_abbr_type(xhtml_abbr_type)
		}
	}
	return xhtml_abbr_type
}

func (xhtml_abbr_type *Xhtml_abbr_type) CommitVoid(stage *Stage) {
	xhtml_abbr_type.Commit(stage)
}

// Checkout xhtml_abbr_type to the back repo (if it is already staged)
func (xhtml_abbr_type *Xhtml_abbr_type) Checkout(stage *Stage) *Xhtml_abbr_type {
	if _, ok := stage.Xhtml_abbr_types[xhtml_abbr_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_abbr_type(xhtml_abbr_type)
		}
	}
	return xhtml_abbr_type
}

// for satisfaction of GongStruct interface
func (xhtml_abbr_type *Xhtml_abbr_type) GetName() (res string) {
	return xhtml_abbr_type.Name
}

// Stage puts xhtml_acronym_type to the model stage
func (xhtml_acronym_type *Xhtml_acronym_type) Stage(stage *Stage) *Xhtml_acronym_type {

	if _, ok := stage.Xhtml_acronym_types[xhtml_acronym_type]; !ok {
		stage.Xhtml_acronym_types[xhtml_acronym_type] = __member
		stage.Xhtml_acronym_typeMap_Staged_Order[xhtml_acronym_type] = stage.Xhtml_acronym_typeOrder
		stage.Xhtml_acronym_typeOrder++
	}
	stage.Xhtml_acronym_types_mapString[xhtml_acronym_type.Name] = xhtml_acronym_type

	return xhtml_acronym_type
}

// Unstage removes xhtml_acronym_type off the model stage
func (xhtml_acronym_type *Xhtml_acronym_type) Unstage(stage *Stage) *Xhtml_acronym_type {
	delete(stage.Xhtml_acronym_types, xhtml_acronym_type)
	delete(stage.Xhtml_acronym_types_mapString, xhtml_acronym_type.Name)
	return xhtml_acronym_type
}

// UnstageVoid removes xhtml_acronym_type off the model stage
func (xhtml_acronym_type *Xhtml_acronym_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_acronym_types, xhtml_acronym_type)
	delete(stage.Xhtml_acronym_types_mapString, xhtml_acronym_type.Name)
}

// commit xhtml_acronym_type to the back repo (if it is already staged)
func (xhtml_acronym_type *Xhtml_acronym_type) Commit(stage *Stage) *Xhtml_acronym_type {
	if _, ok := stage.Xhtml_acronym_types[xhtml_acronym_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_acronym_type(xhtml_acronym_type)
		}
	}
	return xhtml_acronym_type
}

func (xhtml_acronym_type *Xhtml_acronym_type) CommitVoid(stage *Stage) {
	xhtml_acronym_type.Commit(stage)
}

// Checkout xhtml_acronym_type to the back repo (if it is already staged)
func (xhtml_acronym_type *Xhtml_acronym_type) Checkout(stage *Stage) *Xhtml_acronym_type {
	if _, ok := stage.Xhtml_acronym_types[xhtml_acronym_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_acronym_type(xhtml_acronym_type)
		}
	}
	return xhtml_acronym_type
}

// for satisfaction of GongStruct interface
func (xhtml_acronym_type *Xhtml_acronym_type) GetName() (res string) {
	return xhtml_acronym_type.Name
}

// Stage puts xhtml_address_type to the model stage
func (xhtml_address_type *Xhtml_address_type) Stage(stage *Stage) *Xhtml_address_type {

	if _, ok := stage.Xhtml_address_types[xhtml_address_type]; !ok {
		stage.Xhtml_address_types[xhtml_address_type] = __member
		stage.Xhtml_address_typeMap_Staged_Order[xhtml_address_type] = stage.Xhtml_address_typeOrder
		stage.Xhtml_address_typeOrder++
	}
	stage.Xhtml_address_types_mapString[xhtml_address_type.Name] = xhtml_address_type

	return xhtml_address_type
}

// Unstage removes xhtml_address_type off the model stage
func (xhtml_address_type *Xhtml_address_type) Unstage(stage *Stage) *Xhtml_address_type {
	delete(stage.Xhtml_address_types, xhtml_address_type)
	delete(stage.Xhtml_address_types_mapString, xhtml_address_type.Name)
	return xhtml_address_type
}

// UnstageVoid removes xhtml_address_type off the model stage
func (xhtml_address_type *Xhtml_address_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_address_types, xhtml_address_type)
	delete(stage.Xhtml_address_types_mapString, xhtml_address_type.Name)
}

// commit xhtml_address_type to the back repo (if it is already staged)
func (xhtml_address_type *Xhtml_address_type) Commit(stage *Stage) *Xhtml_address_type {
	if _, ok := stage.Xhtml_address_types[xhtml_address_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_address_type(xhtml_address_type)
		}
	}
	return xhtml_address_type
}

func (xhtml_address_type *Xhtml_address_type) CommitVoid(stage *Stage) {
	xhtml_address_type.Commit(stage)
}

// Checkout xhtml_address_type to the back repo (if it is already staged)
func (xhtml_address_type *Xhtml_address_type) Checkout(stage *Stage) *Xhtml_address_type {
	if _, ok := stage.Xhtml_address_types[xhtml_address_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_address_type(xhtml_address_type)
		}
	}
	return xhtml_address_type
}

// for satisfaction of GongStruct interface
func (xhtml_address_type *Xhtml_address_type) GetName() (res string) {
	return xhtml_address_type.Name
}

// Stage puts xhtml_blockquote_type to the model stage
func (xhtml_blockquote_type *Xhtml_blockquote_type) Stage(stage *Stage) *Xhtml_blockquote_type {

	if _, ok := stage.Xhtml_blockquote_types[xhtml_blockquote_type]; !ok {
		stage.Xhtml_blockquote_types[xhtml_blockquote_type] = __member
		stage.Xhtml_blockquote_typeMap_Staged_Order[xhtml_blockquote_type] = stage.Xhtml_blockquote_typeOrder
		stage.Xhtml_blockquote_typeOrder++
	}
	stage.Xhtml_blockquote_types_mapString[xhtml_blockquote_type.Name] = xhtml_blockquote_type

	return xhtml_blockquote_type
}

// Unstage removes xhtml_blockquote_type off the model stage
func (xhtml_blockquote_type *Xhtml_blockquote_type) Unstage(stage *Stage) *Xhtml_blockquote_type {
	delete(stage.Xhtml_blockquote_types, xhtml_blockquote_type)
	delete(stage.Xhtml_blockquote_types_mapString, xhtml_blockquote_type.Name)
	return xhtml_blockquote_type
}

// UnstageVoid removes xhtml_blockquote_type off the model stage
func (xhtml_blockquote_type *Xhtml_blockquote_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_blockquote_types, xhtml_blockquote_type)
	delete(stage.Xhtml_blockquote_types_mapString, xhtml_blockquote_type.Name)
}

// commit xhtml_blockquote_type to the back repo (if it is already staged)
func (xhtml_blockquote_type *Xhtml_blockquote_type) Commit(stage *Stage) *Xhtml_blockquote_type {
	if _, ok := stage.Xhtml_blockquote_types[xhtml_blockquote_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_blockquote_type(xhtml_blockquote_type)
		}
	}
	return xhtml_blockquote_type
}

func (xhtml_blockquote_type *Xhtml_blockquote_type) CommitVoid(stage *Stage) {
	xhtml_blockquote_type.Commit(stage)
}

// Checkout xhtml_blockquote_type to the back repo (if it is already staged)
func (xhtml_blockquote_type *Xhtml_blockquote_type) Checkout(stage *Stage) *Xhtml_blockquote_type {
	if _, ok := stage.Xhtml_blockquote_types[xhtml_blockquote_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_blockquote_type(xhtml_blockquote_type)
		}
	}
	return xhtml_blockquote_type
}

// for satisfaction of GongStruct interface
func (xhtml_blockquote_type *Xhtml_blockquote_type) GetName() (res string) {
	return xhtml_blockquote_type.Name
}

// Stage puts xhtml_br_type to the model stage
func (xhtml_br_type *Xhtml_br_type) Stage(stage *Stage) *Xhtml_br_type {

	if _, ok := stage.Xhtml_br_types[xhtml_br_type]; !ok {
		stage.Xhtml_br_types[xhtml_br_type] = __member
		stage.Xhtml_br_typeMap_Staged_Order[xhtml_br_type] = stage.Xhtml_br_typeOrder
		stage.Xhtml_br_typeOrder++
	}
	stage.Xhtml_br_types_mapString[xhtml_br_type.Name] = xhtml_br_type

	return xhtml_br_type
}

// Unstage removes xhtml_br_type off the model stage
func (xhtml_br_type *Xhtml_br_type) Unstage(stage *Stage) *Xhtml_br_type {
	delete(stage.Xhtml_br_types, xhtml_br_type)
	delete(stage.Xhtml_br_types_mapString, xhtml_br_type.Name)
	return xhtml_br_type
}

// UnstageVoid removes xhtml_br_type off the model stage
func (xhtml_br_type *Xhtml_br_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_br_types, xhtml_br_type)
	delete(stage.Xhtml_br_types_mapString, xhtml_br_type.Name)
}

// commit xhtml_br_type to the back repo (if it is already staged)
func (xhtml_br_type *Xhtml_br_type) Commit(stage *Stage) *Xhtml_br_type {
	if _, ok := stage.Xhtml_br_types[xhtml_br_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_br_type(xhtml_br_type)
		}
	}
	return xhtml_br_type
}

func (xhtml_br_type *Xhtml_br_type) CommitVoid(stage *Stage) {
	xhtml_br_type.Commit(stage)
}

// Checkout xhtml_br_type to the back repo (if it is already staged)
func (xhtml_br_type *Xhtml_br_type) Checkout(stage *Stage) *Xhtml_br_type {
	if _, ok := stage.Xhtml_br_types[xhtml_br_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_br_type(xhtml_br_type)
		}
	}
	return xhtml_br_type
}

// for satisfaction of GongStruct interface
func (xhtml_br_type *Xhtml_br_type) GetName() (res string) {
	return xhtml_br_type.Name
}

// Stage puts xhtml_caption_type to the model stage
func (xhtml_caption_type *Xhtml_caption_type) Stage(stage *Stage) *Xhtml_caption_type {

	if _, ok := stage.Xhtml_caption_types[xhtml_caption_type]; !ok {
		stage.Xhtml_caption_types[xhtml_caption_type] = __member
		stage.Xhtml_caption_typeMap_Staged_Order[xhtml_caption_type] = stage.Xhtml_caption_typeOrder
		stage.Xhtml_caption_typeOrder++
	}
	stage.Xhtml_caption_types_mapString[xhtml_caption_type.Name] = xhtml_caption_type

	return xhtml_caption_type
}

// Unstage removes xhtml_caption_type off the model stage
func (xhtml_caption_type *Xhtml_caption_type) Unstage(stage *Stage) *Xhtml_caption_type {
	delete(stage.Xhtml_caption_types, xhtml_caption_type)
	delete(stage.Xhtml_caption_types_mapString, xhtml_caption_type.Name)
	return xhtml_caption_type
}

// UnstageVoid removes xhtml_caption_type off the model stage
func (xhtml_caption_type *Xhtml_caption_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_caption_types, xhtml_caption_type)
	delete(stage.Xhtml_caption_types_mapString, xhtml_caption_type.Name)
}

// commit xhtml_caption_type to the back repo (if it is already staged)
func (xhtml_caption_type *Xhtml_caption_type) Commit(stage *Stage) *Xhtml_caption_type {
	if _, ok := stage.Xhtml_caption_types[xhtml_caption_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_caption_type(xhtml_caption_type)
		}
	}
	return xhtml_caption_type
}

func (xhtml_caption_type *Xhtml_caption_type) CommitVoid(stage *Stage) {
	xhtml_caption_type.Commit(stage)
}

// Checkout xhtml_caption_type to the back repo (if it is already staged)
func (xhtml_caption_type *Xhtml_caption_type) Checkout(stage *Stage) *Xhtml_caption_type {
	if _, ok := stage.Xhtml_caption_types[xhtml_caption_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_caption_type(xhtml_caption_type)
		}
	}
	return xhtml_caption_type
}

// for satisfaction of GongStruct interface
func (xhtml_caption_type *Xhtml_caption_type) GetName() (res string) {
	return xhtml_caption_type.Name
}

// Stage puts xhtml_cite_type to the model stage
func (xhtml_cite_type *Xhtml_cite_type) Stage(stage *Stage) *Xhtml_cite_type {

	if _, ok := stage.Xhtml_cite_types[xhtml_cite_type]; !ok {
		stage.Xhtml_cite_types[xhtml_cite_type] = __member
		stage.Xhtml_cite_typeMap_Staged_Order[xhtml_cite_type] = stage.Xhtml_cite_typeOrder
		stage.Xhtml_cite_typeOrder++
	}
	stage.Xhtml_cite_types_mapString[xhtml_cite_type.Name] = xhtml_cite_type

	return xhtml_cite_type
}

// Unstage removes xhtml_cite_type off the model stage
func (xhtml_cite_type *Xhtml_cite_type) Unstage(stage *Stage) *Xhtml_cite_type {
	delete(stage.Xhtml_cite_types, xhtml_cite_type)
	delete(stage.Xhtml_cite_types_mapString, xhtml_cite_type.Name)
	return xhtml_cite_type
}

// UnstageVoid removes xhtml_cite_type off the model stage
func (xhtml_cite_type *Xhtml_cite_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_cite_types, xhtml_cite_type)
	delete(stage.Xhtml_cite_types_mapString, xhtml_cite_type.Name)
}

// commit xhtml_cite_type to the back repo (if it is already staged)
func (xhtml_cite_type *Xhtml_cite_type) Commit(stage *Stage) *Xhtml_cite_type {
	if _, ok := stage.Xhtml_cite_types[xhtml_cite_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_cite_type(xhtml_cite_type)
		}
	}
	return xhtml_cite_type
}

func (xhtml_cite_type *Xhtml_cite_type) CommitVoid(stage *Stage) {
	xhtml_cite_type.Commit(stage)
}

// Checkout xhtml_cite_type to the back repo (if it is already staged)
func (xhtml_cite_type *Xhtml_cite_type) Checkout(stage *Stage) *Xhtml_cite_type {
	if _, ok := stage.Xhtml_cite_types[xhtml_cite_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_cite_type(xhtml_cite_type)
		}
	}
	return xhtml_cite_type
}

// for satisfaction of GongStruct interface
func (xhtml_cite_type *Xhtml_cite_type) GetName() (res string) {
	return xhtml_cite_type.Name
}

// Stage puts xhtml_code_type to the model stage
func (xhtml_code_type *Xhtml_code_type) Stage(stage *Stage) *Xhtml_code_type {

	if _, ok := stage.Xhtml_code_types[xhtml_code_type]; !ok {
		stage.Xhtml_code_types[xhtml_code_type] = __member
		stage.Xhtml_code_typeMap_Staged_Order[xhtml_code_type] = stage.Xhtml_code_typeOrder
		stage.Xhtml_code_typeOrder++
	}
	stage.Xhtml_code_types_mapString[xhtml_code_type.Name] = xhtml_code_type

	return xhtml_code_type
}

// Unstage removes xhtml_code_type off the model stage
func (xhtml_code_type *Xhtml_code_type) Unstage(stage *Stage) *Xhtml_code_type {
	delete(stage.Xhtml_code_types, xhtml_code_type)
	delete(stage.Xhtml_code_types_mapString, xhtml_code_type.Name)
	return xhtml_code_type
}

// UnstageVoid removes xhtml_code_type off the model stage
func (xhtml_code_type *Xhtml_code_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_code_types, xhtml_code_type)
	delete(stage.Xhtml_code_types_mapString, xhtml_code_type.Name)
}

// commit xhtml_code_type to the back repo (if it is already staged)
func (xhtml_code_type *Xhtml_code_type) Commit(stage *Stage) *Xhtml_code_type {
	if _, ok := stage.Xhtml_code_types[xhtml_code_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_code_type(xhtml_code_type)
		}
	}
	return xhtml_code_type
}

func (xhtml_code_type *Xhtml_code_type) CommitVoid(stage *Stage) {
	xhtml_code_type.Commit(stage)
}

// Checkout xhtml_code_type to the back repo (if it is already staged)
func (xhtml_code_type *Xhtml_code_type) Checkout(stage *Stage) *Xhtml_code_type {
	if _, ok := stage.Xhtml_code_types[xhtml_code_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_code_type(xhtml_code_type)
		}
	}
	return xhtml_code_type
}

// for satisfaction of GongStruct interface
func (xhtml_code_type *Xhtml_code_type) GetName() (res string) {
	return xhtml_code_type.Name
}

// Stage puts xhtml_col_type to the model stage
func (xhtml_col_type *Xhtml_col_type) Stage(stage *Stage) *Xhtml_col_type {

	if _, ok := stage.Xhtml_col_types[xhtml_col_type]; !ok {
		stage.Xhtml_col_types[xhtml_col_type] = __member
		stage.Xhtml_col_typeMap_Staged_Order[xhtml_col_type] = stage.Xhtml_col_typeOrder
		stage.Xhtml_col_typeOrder++
	}
	stage.Xhtml_col_types_mapString[xhtml_col_type.Name] = xhtml_col_type

	return xhtml_col_type
}

// Unstage removes xhtml_col_type off the model stage
func (xhtml_col_type *Xhtml_col_type) Unstage(stage *Stage) *Xhtml_col_type {
	delete(stage.Xhtml_col_types, xhtml_col_type)
	delete(stage.Xhtml_col_types_mapString, xhtml_col_type.Name)
	return xhtml_col_type
}

// UnstageVoid removes xhtml_col_type off the model stage
func (xhtml_col_type *Xhtml_col_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_col_types, xhtml_col_type)
	delete(stage.Xhtml_col_types_mapString, xhtml_col_type.Name)
}

// commit xhtml_col_type to the back repo (if it is already staged)
func (xhtml_col_type *Xhtml_col_type) Commit(stage *Stage) *Xhtml_col_type {
	if _, ok := stage.Xhtml_col_types[xhtml_col_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_col_type(xhtml_col_type)
		}
	}
	return xhtml_col_type
}

func (xhtml_col_type *Xhtml_col_type) CommitVoid(stage *Stage) {
	xhtml_col_type.Commit(stage)
}

// Checkout xhtml_col_type to the back repo (if it is already staged)
func (xhtml_col_type *Xhtml_col_type) Checkout(stage *Stage) *Xhtml_col_type {
	if _, ok := stage.Xhtml_col_types[xhtml_col_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_col_type(xhtml_col_type)
		}
	}
	return xhtml_col_type
}

// for satisfaction of GongStruct interface
func (xhtml_col_type *Xhtml_col_type) GetName() (res string) {
	return xhtml_col_type.Name
}

// Stage puts xhtml_colgroup_type to the model stage
func (xhtml_colgroup_type *Xhtml_colgroup_type) Stage(stage *Stage) *Xhtml_colgroup_type {

	if _, ok := stage.Xhtml_colgroup_types[xhtml_colgroup_type]; !ok {
		stage.Xhtml_colgroup_types[xhtml_colgroup_type] = __member
		stage.Xhtml_colgroup_typeMap_Staged_Order[xhtml_colgroup_type] = stage.Xhtml_colgroup_typeOrder
		stage.Xhtml_colgroup_typeOrder++
	}
	stage.Xhtml_colgroup_types_mapString[xhtml_colgroup_type.Name] = xhtml_colgroup_type

	return xhtml_colgroup_type
}

// Unstage removes xhtml_colgroup_type off the model stage
func (xhtml_colgroup_type *Xhtml_colgroup_type) Unstage(stage *Stage) *Xhtml_colgroup_type {
	delete(stage.Xhtml_colgroup_types, xhtml_colgroup_type)
	delete(stage.Xhtml_colgroup_types_mapString, xhtml_colgroup_type.Name)
	return xhtml_colgroup_type
}

// UnstageVoid removes xhtml_colgroup_type off the model stage
func (xhtml_colgroup_type *Xhtml_colgroup_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_colgroup_types, xhtml_colgroup_type)
	delete(stage.Xhtml_colgroup_types_mapString, xhtml_colgroup_type.Name)
}

// commit xhtml_colgroup_type to the back repo (if it is already staged)
func (xhtml_colgroup_type *Xhtml_colgroup_type) Commit(stage *Stage) *Xhtml_colgroup_type {
	if _, ok := stage.Xhtml_colgroup_types[xhtml_colgroup_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_colgroup_type(xhtml_colgroup_type)
		}
	}
	return xhtml_colgroup_type
}

func (xhtml_colgroup_type *Xhtml_colgroup_type) CommitVoid(stage *Stage) {
	xhtml_colgroup_type.Commit(stage)
}

// Checkout xhtml_colgroup_type to the back repo (if it is already staged)
func (xhtml_colgroup_type *Xhtml_colgroup_type) Checkout(stage *Stage) *Xhtml_colgroup_type {
	if _, ok := stage.Xhtml_colgroup_types[xhtml_colgroup_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_colgroup_type(xhtml_colgroup_type)
		}
	}
	return xhtml_colgroup_type
}

// for satisfaction of GongStruct interface
func (xhtml_colgroup_type *Xhtml_colgroup_type) GetName() (res string) {
	return xhtml_colgroup_type.Name
}

// Stage puts xhtml_dd_type to the model stage
func (xhtml_dd_type *Xhtml_dd_type) Stage(stage *Stage) *Xhtml_dd_type {

	if _, ok := stage.Xhtml_dd_types[xhtml_dd_type]; !ok {
		stage.Xhtml_dd_types[xhtml_dd_type] = __member
		stage.Xhtml_dd_typeMap_Staged_Order[xhtml_dd_type] = stage.Xhtml_dd_typeOrder
		stage.Xhtml_dd_typeOrder++
	}
	stage.Xhtml_dd_types_mapString[xhtml_dd_type.Name] = xhtml_dd_type

	return xhtml_dd_type
}

// Unstage removes xhtml_dd_type off the model stage
func (xhtml_dd_type *Xhtml_dd_type) Unstage(stage *Stage) *Xhtml_dd_type {
	delete(stage.Xhtml_dd_types, xhtml_dd_type)
	delete(stage.Xhtml_dd_types_mapString, xhtml_dd_type.Name)
	return xhtml_dd_type
}

// UnstageVoid removes xhtml_dd_type off the model stage
func (xhtml_dd_type *Xhtml_dd_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_dd_types, xhtml_dd_type)
	delete(stage.Xhtml_dd_types_mapString, xhtml_dd_type.Name)
}

// commit xhtml_dd_type to the back repo (if it is already staged)
func (xhtml_dd_type *Xhtml_dd_type) Commit(stage *Stage) *Xhtml_dd_type {
	if _, ok := stage.Xhtml_dd_types[xhtml_dd_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_dd_type(xhtml_dd_type)
		}
	}
	return xhtml_dd_type
}

func (xhtml_dd_type *Xhtml_dd_type) CommitVoid(stage *Stage) {
	xhtml_dd_type.Commit(stage)
}

// Checkout xhtml_dd_type to the back repo (if it is already staged)
func (xhtml_dd_type *Xhtml_dd_type) Checkout(stage *Stage) *Xhtml_dd_type {
	if _, ok := stage.Xhtml_dd_types[xhtml_dd_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_dd_type(xhtml_dd_type)
		}
	}
	return xhtml_dd_type
}

// for satisfaction of GongStruct interface
func (xhtml_dd_type *Xhtml_dd_type) GetName() (res string) {
	return xhtml_dd_type.Name
}

// Stage puts xhtml_dfn_type to the model stage
func (xhtml_dfn_type *Xhtml_dfn_type) Stage(stage *Stage) *Xhtml_dfn_type {

	if _, ok := stage.Xhtml_dfn_types[xhtml_dfn_type]; !ok {
		stage.Xhtml_dfn_types[xhtml_dfn_type] = __member
		stage.Xhtml_dfn_typeMap_Staged_Order[xhtml_dfn_type] = stage.Xhtml_dfn_typeOrder
		stage.Xhtml_dfn_typeOrder++
	}
	stage.Xhtml_dfn_types_mapString[xhtml_dfn_type.Name] = xhtml_dfn_type

	return xhtml_dfn_type
}

// Unstage removes xhtml_dfn_type off the model stage
func (xhtml_dfn_type *Xhtml_dfn_type) Unstage(stage *Stage) *Xhtml_dfn_type {
	delete(stage.Xhtml_dfn_types, xhtml_dfn_type)
	delete(stage.Xhtml_dfn_types_mapString, xhtml_dfn_type.Name)
	return xhtml_dfn_type
}

// UnstageVoid removes xhtml_dfn_type off the model stage
func (xhtml_dfn_type *Xhtml_dfn_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_dfn_types, xhtml_dfn_type)
	delete(stage.Xhtml_dfn_types_mapString, xhtml_dfn_type.Name)
}

// commit xhtml_dfn_type to the back repo (if it is already staged)
func (xhtml_dfn_type *Xhtml_dfn_type) Commit(stage *Stage) *Xhtml_dfn_type {
	if _, ok := stage.Xhtml_dfn_types[xhtml_dfn_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_dfn_type(xhtml_dfn_type)
		}
	}
	return xhtml_dfn_type
}

func (xhtml_dfn_type *Xhtml_dfn_type) CommitVoid(stage *Stage) {
	xhtml_dfn_type.Commit(stage)
}

// Checkout xhtml_dfn_type to the back repo (if it is already staged)
func (xhtml_dfn_type *Xhtml_dfn_type) Checkout(stage *Stage) *Xhtml_dfn_type {
	if _, ok := stage.Xhtml_dfn_types[xhtml_dfn_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_dfn_type(xhtml_dfn_type)
		}
	}
	return xhtml_dfn_type
}

// for satisfaction of GongStruct interface
func (xhtml_dfn_type *Xhtml_dfn_type) GetName() (res string) {
	return xhtml_dfn_type.Name
}

// Stage puts xhtml_div_type to the model stage
func (xhtml_div_type *Xhtml_div_type) Stage(stage *Stage) *Xhtml_div_type {

	if _, ok := stage.Xhtml_div_types[xhtml_div_type]; !ok {
		stage.Xhtml_div_types[xhtml_div_type] = __member
		stage.Xhtml_div_typeMap_Staged_Order[xhtml_div_type] = stage.Xhtml_div_typeOrder
		stage.Xhtml_div_typeOrder++
	}
	stage.Xhtml_div_types_mapString[xhtml_div_type.Name] = xhtml_div_type

	return xhtml_div_type
}

// Unstage removes xhtml_div_type off the model stage
func (xhtml_div_type *Xhtml_div_type) Unstage(stage *Stage) *Xhtml_div_type {
	delete(stage.Xhtml_div_types, xhtml_div_type)
	delete(stage.Xhtml_div_types_mapString, xhtml_div_type.Name)
	return xhtml_div_type
}

// UnstageVoid removes xhtml_div_type off the model stage
func (xhtml_div_type *Xhtml_div_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_div_types, xhtml_div_type)
	delete(stage.Xhtml_div_types_mapString, xhtml_div_type.Name)
}

// commit xhtml_div_type to the back repo (if it is already staged)
func (xhtml_div_type *Xhtml_div_type) Commit(stage *Stage) *Xhtml_div_type {
	if _, ok := stage.Xhtml_div_types[xhtml_div_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_div_type(xhtml_div_type)
		}
	}
	return xhtml_div_type
}

func (xhtml_div_type *Xhtml_div_type) CommitVoid(stage *Stage) {
	xhtml_div_type.Commit(stage)
}

// Checkout xhtml_div_type to the back repo (if it is already staged)
func (xhtml_div_type *Xhtml_div_type) Checkout(stage *Stage) *Xhtml_div_type {
	if _, ok := stage.Xhtml_div_types[xhtml_div_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_div_type(xhtml_div_type)
		}
	}
	return xhtml_div_type
}

// for satisfaction of GongStruct interface
func (xhtml_div_type *Xhtml_div_type) GetName() (res string) {
	return xhtml_div_type.Name
}

// Stage puts xhtml_dl_type to the model stage
func (xhtml_dl_type *Xhtml_dl_type) Stage(stage *Stage) *Xhtml_dl_type {

	if _, ok := stage.Xhtml_dl_types[xhtml_dl_type]; !ok {
		stage.Xhtml_dl_types[xhtml_dl_type] = __member
		stage.Xhtml_dl_typeMap_Staged_Order[xhtml_dl_type] = stage.Xhtml_dl_typeOrder
		stage.Xhtml_dl_typeOrder++
	}
	stage.Xhtml_dl_types_mapString[xhtml_dl_type.Name] = xhtml_dl_type

	return xhtml_dl_type
}

// Unstage removes xhtml_dl_type off the model stage
func (xhtml_dl_type *Xhtml_dl_type) Unstage(stage *Stage) *Xhtml_dl_type {
	delete(stage.Xhtml_dl_types, xhtml_dl_type)
	delete(stage.Xhtml_dl_types_mapString, xhtml_dl_type.Name)
	return xhtml_dl_type
}

// UnstageVoid removes xhtml_dl_type off the model stage
func (xhtml_dl_type *Xhtml_dl_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_dl_types, xhtml_dl_type)
	delete(stage.Xhtml_dl_types_mapString, xhtml_dl_type.Name)
}

// commit xhtml_dl_type to the back repo (if it is already staged)
func (xhtml_dl_type *Xhtml_dl_type) Commit(stage *Stage) *Xhtml_dl_type {
	if _, ok := stage.Xhtml_dl_types[xhtml_dl_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_dl_type(xhtml_dl_type)
		}
	}
	return xhtml_dl_type
}

func (xhtml_dl_type *Xhtml_dl_type) CommitVoid(stage *Stage) {
	xhtml_dl_type.Commit(stage)
}

// Checkout xhtml_dl_type to the back repo (if it is already staged)
func (xhtml_dl_type *Xhtml_dl_type) Checkout(stage *Stage) *Xhtml_dl_type {
	if _, ok := stage.Xhtml_dl_types[xhtml_dl_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_dl_type(xhtml_dl_type)
		}
	}
	return xhtml_dl_type
}

// for satisfaction of GongStruct interface
func (xhtml_dl_type *Xhtml_dl_type) GetName() (res string) {
	return xhtml_dl_type.Name
}

// Stage puts xhtml_dt_type to the model stage
func (xhtml_dt_type *Xhtml_dt_type) Stage(stage *Stage) *Xhtml_dt_type {

	if _, ok := stage.Xhtml_dt_types[xhtml_dt_type]; !ok {
		stage.Xhtml_dt_types[xhtml_dt_type] = __member
		stage.Xhtml_dt_typeMap_Staged_Order[xhtml_dt_type] = stage.Xhtml_dt_typeOrder
		stage.Xhtml_dt_typeOrder++
	}
	stage.Xhtml_dt_types_mapString[xhtml_dt_type.Name] = xhtml_dt_type

	return xhtml_dt_type
}

// Unstage removes xhtml_dt_type off the model stage
func (xhtml_dt_type *Xhtml_dt_type) Unstage(stage *Stage) *Xhtml_dt_type {
	delete(stage.Xhtml_dt_types, xhtml_dt_type)
	delete(stage.Xhtml_dt_types_mapString, xhtml_dt_type.Name)
	return xhtml_dt_type
}

// UnstageVoid removes xhtml_dt_type off the model stage
func (xhtml_dt_type *Xhtml_dt_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_dt_types, xhtml_dt_type)
	delete(stage.Xhtml_dt_types_mapString, xhtml_dt_type.Name)
}

// commit xhtml_dt_type to the back repo (if it is already staged)
func (xhtml_dt_type *Xhtml_dt_type) Commit(stage *Stage) *Xhtml_dt_type {
	if _, ok := stage.Xhtml_dt_types[xhtml_dt_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_dt_type(xhtml_dt_type)
		}
	}
	return xhtml_dt_type
}

func (xhtml_dt_type *Xhtml_dt_type) CommitVoid(stage *Stage) {
	xhtml_dt_type.Commit(stage)
}

// Checkout xhtml_dt_type to the back repo (if it is already staged)
func (xhtml_dt_type *Xhtml_dt_type) Checkout(stage *Stage) *Xhtml_dt_type {
	if _, ok := stage.Xhtml_dt_types[xhtml_dt_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_dt_type(xhtml_dt_type)
		}
	}
	return xhtml_dt_type
}

// for satisfaction of GongStruct interface
func (xhtml_dt_type *Xhtml_dt_type) GetName() (res string) {
	return xhtml_dt_type.Name
}

// Stage puts xhtml_edit_type to the model stage
func (xhtml_edit_type *Xhtml_edit_type) Stage(stage *Stage) *Xhtml_edit_type {

	if _, ok := stage.Xhtml_edit_types[xhtml_edit_type]; !ok {
		stage.Xhtml_edit_types[xhtml_edit_type] = __member
		stage.Xhtml_edit_typeMap_Staged_Order[xhtml_edit_type] = stage.Xhtml_edit_typeOrder
		stage.Xhtml_edit_typeOrder++
	}
	stage.Xhtml_edit_types_mapString[xhtml_edit_type.Name] = xhtml_edit_type

	return xhtml_edit_type
}

// Unstage removes xhtml_edit_type off the model stage
func (xhtml_edit_type *Xhtml_edit_type) Unstage(stage *Stage) *Xhtml_edit_type {
	delete(stage.Xhtml_edit_types, xhtml_edit_type)
	delete(stage.Xhtml_edit_types_mapString, xhtml_edit_type.Name)
	return xhtml_edit_type
}

// UnstageVoid removes xhtml_edit_type off the model stage
func (xhtml_edit_type *Xhtml_edit_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_edit_types, xhtml_edit_type)
	delete(stage.Xhtml_edit_types_mapString, xhtml_edit_type.Name)
}

// commit xhtml_edit_type to the back repo (if it is already staged)
func (xhtml_edit_type *Xhtml_edit_type) Commit(stage *Stage) *Xhtml_edit_type {
	if _, ok := stage.Xhtml_edit_types[xhtml_edit_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_edit_type(xhtml_edit_type)
		}
	}
	return xhtml_edit_type
}

func (xhtml_edit_type *Xhtml_edit_type) CommitVoid(stage *Stage) {
	xhtml_edit_type.Commit(stage)
}

// Checkout xhtml_edit_type to the back repo (if it is already staged)
func (xhtml_edit_type *Xhtml_edit_type) Checkout(stage *Stage) *Xhtml_edit_type {
	if _, ok := stage.Xhtml_edit_types[xhtml_edit_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_edit_type(xhtml_edit_type)
		}
	}
	return xhtml_edit_type
}

// for satisfaction of GongStruct interface
func (xhtml_edit_type *Xhtml_edit_type) GetName() (res string) {
	return xhtml_edit_type.Name
}

// Stage puts xhtml_em_type to the model stage
func (xhtml_em_type *Xhtml_em_type) Stage(stage *Stage) *Xhtml_em_type {

	if _, ok := stage.Xhtml_em_types[xhtml_em_type]; !ok {
		stage.Xhtml_em_types[xhtml_em_type] = __member
		stage.Xhtml_em_typeMap_Staged_Order[xhtml_em_type] = stage.Xhtml_em_typeOrder
		stage.Xhtml_em_typeOrder++
	}
	stage.Xhtml_em_types_mapString[xhtml_em_type.Name] = xhtml_em_type

	return xhtml_em_type
}

// Unstage removes xhtml_em_type off the model stage
func (xhtml_em_type *Xhtml_em_type) Unstage(stage *Stage) *Xhtml_em_type {
	delete(stage.Xhtml_em_types, xhtml_em_type)
	delete(stage.Xhtml_em_types_mapString, xhtml_em_type.Name)
	return xhtml_em_type
}

// UnstageVoid removes xhtml_em_type off the model stage
func (xhtml_em_type *Xhtml_em_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_em_types, xhtml_em_type)
	delete(stage.Xhtml_em_types_mapString, xhtml_em_type.Name)
}

// commit xhtml_em_type to the back repo (if it is already staged)
func (xhtml_em_type *Xhtml_em_type) Commit(stage *Stage) *Xhtml_em_type {
	if _, ok := stage.Xhtml_em_types[xhtml_em_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_em_type(xhtml_em_type)
		}
	}
	return xhtml_em_type
}

func (xhtml_em_type *Xhtml_em_type) CommitVoid(stage *Stage) {
	xhtml_em_type.Commit(stage)
}

// Checkout xhtml_em_type to the back repo (if it is already staged)
func (xhtml_em_type *Xhtml_em_type) Checkout(stage *Stage) *Xhtml_em_type {
	if _, ok := stage.Xhtml_em_types[xhtml_em_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_em_type(xhtml_em_type)
		}
	}
	return xhtml_em_type
}

// for satisfaction of GongStruct interface
func (xhtml_em_type *Xhtml_em_type) GetName() (res string) {
	return xhtml_em_type.Name
}

// Stage puts xhtml_h1_type to the model stage
func (xhtml_h1_type *Xhtml_h1_type) Stage(stage *Stage) *Xhtml_h1_type {

	if _, ok := stage.Xhtml_h1_types[xhtml_h1_type]; !ok {
		stage.Xhtml_h1_types[xhtml_h1_type] = __member
		stage.Xhtml_h1_typeMap_Staged_Order[xhtml_h1_type] = stage.Xhtml_h1_typeOrder
		stage.Xhtml_h1_typeOrder++
	}
	stage.Xhtml_h1_types_mapString[xhtml_h1_type.Name] = xhtml_h1_type

	return xhtml_h1_type
}

// Unstage removes xhtml_h1_type off the model stage
func (xhtml_h1_type *Xhtml_h1_type) Unstage(stage *Stage) *Xhtml_h1_type {
	delete(stage.Xhtml_h1_types, xhtml_h1_type)
	delete(stage.Xhtml_h1_types_mapString, xhtml_h1_type.Name)
	return xhtml_h1_type
}

// UnstageVoid removes xhtml_h1_type off the model stage
func (xhtml_h1_type *Xhtml_h1_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_h1_types, xhtml_h1_type)
	delete(stage.Xhtml_h1_types_mapString, xhtml_h1_type.Name)
}

// commit xhtml_h1_type to the back repo (if it is already staged)
func (xhtml_h1_type *Xhtml_h1_type) Commit(stage *Stage) *Xhtml_h1_type {
	if _, ok := stage.Xhtml_h1_types[xhtml_h1_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_h1_type(xhtml_h1_type)
		}
	}
	return xhtml_h1_type
}

func (xhtml_h1_type *Xhtml_h1_type) CommitVoid(stage *Stage) {
	xhtml_h1_type.Commit(stage)
}

// Checkout xhtml_h1_type to the back repo (if it is already staged)
func (xhtml_h1_type *Xhtml_h1_type) Checkout(stage *Stage) *Xhtml_h1_type {
	if _, ok := stage.Xhtml_h1_types[xhtml_h1_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_h1_type(xhtml_h1_type)
		}
	}
	return xhtml_h1_type
}

// for satisfaction of GongStruct interface
func (xhtml_h1_type *Xhtml_h1_type) GetName() (res string) {
	return xhtml_h1_type.Name
}

// Stage puts xhtml_h2_type to the model stage
func (xhtml_h2_type *Xhtml_h2_type) Stage(stage *Stage) *Xhtml_h2_type {

	if _, ok := stage.Xhtml_h2_types[xhtml_h2_type]; !ok {
		stage.Xhtml_h2_types[xhtml_h2_type] = __member
		stage.Xhtml_h2_typeMap_Staged_Order[xhtml_h2_type] = stage.Xhtml_h2_typeOrder
		stage.Xhtml_h2_typeOrder++
	}
	stage.Xhtml_h2_types_mapString[xhtml_h2_type.Name] = xhtml_h2_type

	return xhtml_h2_type
}

// Unstage removes xhtml_h2_type off the model stage
func (xhtml_h2_type *Xhtml_h2_type) Unstage(stage *Stage) *Xhtml_h2_type {
	delete(stage.Xhtml_h2_types, xhtml_h2_type)
	delete(stage.Xhtml_h2_types_mapString, xhtml_h2_type.Name)
	return xhtml_h2_type
}

// UnstageVoid removes xhtml_h2_type off the model stage
func (xhtml_h2_type *Xhtml_h2_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_h2_types, xhtml_h2_type)
	delete(stage.Xhtml_h2_types_mapString, xhtml_h2_type.Name)
}

// commit xhtml_h2_type to the back repo (if it is already staged)
func (xhtml_h2_type *Xhtml_h2_type) Commit(stage *Stage) *Xhtml_h2_type {
	if _, ok := stage.Xhtml_h2_types[xhtml_h2_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_h2_type(xhtml_h2_type)
		}
	}
	return xhtml_h2_type
}

func (xhtml_h2_type *Xhtml_h2_type) CommitVoid(stage *Stage) {
	xhtml_h2_type.Commit(stage)
}

// Checkout xhtml_h2_type to the back repo (if it is already staged)
func (xhtml_h2_type *Xhtml_h2_type) Checkout(stage *Stage) *Xhtml_h2_type {
	if _, ok := stage.Xhtml_h2_types[xhtml_h2_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_h2_type(xhtml_h2_type)
		}
	}
	return xhtml_h2_type
}

// for satisfaction of GongStruct interface
func (xhtml_h2_type *Xhtml_h2_type) GetName() (res string) {
	return xhtml_h2_type.Name
}

// Stage puts xhtml_h3_type to the model stage
func (xhtml_h3_type *Xhtml_h3_type) Stage(stage *Stage) *Xhtml_h3_type {

	if _, ok := stage.Xhtml_h3_types[xhtml_h3_type]; !ok {
		stage.Xhtml_h3_types[xhtml_h3_type] = __member
		stage.Xhtml_h3_typeMap_Staged_Order[xhtml_h3_type] = stage.Xhtml_h3_typeOrder
		stage.Xhtml_h3_typeOrder++
	}
	stage.Xhtml_h3_types_mapString[xhtml_h3_type.Name] = xhtml_h3_type

	return xhtml_h3_type
}

// Unstage removes xhtml_h3_type off the model stage
func (xhtml_h3_type *Xhtml_h3_type) Unstage(stage *Stage) *Xhtml_h3_type {
	delete(stage.Xhtml_h3_types, xhtml_h3_type)
	delete(stage.Xhtml_h3_types_mapString, xhtml_h3_type.Name)
	return xhtml_h3_type
}

// UnstageVoid removes xhtml_h3_type off the model stage
func (xhtml_h3_type *Xhtml_h3_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_h3_types, xhtml_h3_type)
	delete(stage.Xhtml_h3_types_mapString, xhtml_h3_type.Name)
}

// commit xhtml_h3_type to the back repo (if it is already staged)
func (xhtml_h3_type *Xhtml_h3_type) Commit(stage *Stage) *Xhtml_h3_type {
	if _, ok := stage.Xhtml_h3_types[xhtml_h3_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_h3_type(xhtml_h3_type)
		}
	}
	return xhtml_h3_type
}

func (xhtml_h3_type *Xhtml_h3_type) CommitVoid(stage *Stage) {
	xhtml_h3_type.Commit(stage)
}

// Checkout xhtml_h3_type to the back repo (if it is already staged)
func (xhtml_h3_type *Xhtml_h3_type) Checkout(stage *Stage) *Xhtml_h3_type {
	if _, ok := stage.Xhtml_h3_types[xhtml_h3_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_h3_type(xhtml_h3_type)
		}
	}
	return xhtml_h3_type
}

// for satisfaction of GongStruct interface
func (xhtml_h3_type *Xhtml_h3_type) GetName() (res string) {
	return xhtml_h3_type.Name
}

// Stage puts xhtml_h4_type to the model stage
func (xhtml_h4_type *Xhtml_h4_type) Stage(stage *Stage) *Xhtml_h4_type {

	if _, ok := stage.Xhtml_h4_types[xhtml_h4_type]; !ok {
		stage.Xhtml_h4_types[xhtml_h4_type] = __member
		stage.Xhtml_h4_typeMap_Staged_Order[xhtml_h4_type] = stage.Xhtml_h4_typeOrder
		stage.Xhtml_h4_typeOrder++
	}
	stage.Xhtml_h4_types_mapString[xhtml_h4_type.Name] = xhtml_h4_type

	return xhtml_h4_type
}

// Unstage removes xhtml_h4_type off the model stage
func (xhtml_h4_type *Xhtml_h4_type) Unstage(stage *Stage) *Xhtml_h4_type {
	delete(stage.Xhtml_h4_types, xhtml_h4_type)
	delete(stage.Xhtml_h4_types_mapString, xhtml_h4_type.Name)
	return xhtml_h4_type
}

// UnstageVoid removes xhtml_h4_type off the model stage
func (xhtml_h4_type *Xhtml_h4_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_h4_types, xhtml_h4_type)
	delete(stage.Xhtml_h4_types_mapString, xhtml_h4_type.Name)
}

// commit xhtml_h4_type to the back repo (if it is already staged)
func (xhtml_h4_type *Xhtml_h4_type) Commit(stage *Stage) *Xhtml_h4_type {
	if _, ok := stage.Xhtml_h4_types[xhtml_h4_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_h4_type(xhtml_h4_type)
		}
	}
	return xhtml_h4_type
}

func (xhtml_h4_type *Xhtml_h4_type) CommitVoid(stage *Stage) {
	xhtml_h4_type.Commit(stage)
}

// Checkout xhtml_h4_type to the back repo (if it is already staged)
func (xhtml_h4_type *Xhtml_h4_type) Checkout(stage *Stage) *Xhtml_h4_type {
	if _, ok := stage.Xhtml_h4_types[xhtml_h4_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_h4_type(xhtml_h4_type)
		}
	}
	return xhtml_h4_type
}

// for satisfaction of GongStruct interface
func (xhtml_h4_type *Xhtml_h4_type) GetName() (res string) {
	return xhtml_h4_type.Name
}

// Stage puts xhtml_h5_type to the model stage
func (xhtml_h5_type *Xhtml_h5_type) Stage(stage *Stage) *Xhtml_h5_type {

	if _, ok := stage.Xhtml_h5_types[xhtml_h5_type]; !ok {
		stage.Xhtml_h5_types[xhtml_h5_type] = __member
		stage.Xhtml_h5_typeMap_Staged_Order[xhtml_h5_type] = stage.Xhtml_h5_typeOrder
		stage.Xhtml_h5_typeOrder++
	}
	stage.Xhtml_h5_types_mapString[xhtml_h5_type.Name] = xhtml_h5_type

	return xhtml_h5_type
}

// Unstage removes xhtml_h5_type off the model stage
func (xhtml_h5_type *Xhtml_h5_type) Unstage(stage *Stage) *Xhtml_h5_type {
	delete(stage.Xhtml_h5_types, xhtml_h5_type)
	delete(stage.Xhtml_h5_types_mapString, xhtml_h5_type.Name)
	return xhtml_h5_type
}

// UnstageVoid removes xhtml_h5_type off the model stage
func (xhtml_h5_type *Xhtml_h5_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_h5_types, xhtml_h5_type)
	delete(stage.Xhtml_h5_types_mapString, xhtml_h5_type.Name)
}

// commit xhtml_h5_type to the back repo (if it is already staged)
func (xhtml_h5_type *Xhtml_h5_type) Commit(stage *Stage) *Xhtml_h5_type {
	if _, ok := stage.Xhtml_h5_types[xhtml_h5_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_h5_type(xhtml_h5_type)
		}
	}
	return xhtml_h5_type
}

func (xhtml_h5_type *Xhtml_h5_type) CommitVoid(stage *Stage) {
	xhtml_h5_type.Commit(stage)
}

// Checkout xhtml_h5_type to the back repo (if it is already staged)
func (xhtml_h5_type *Xhtml_h5_type) Checkout(stage *Stage) *Xhtml_h5_type {
	if _, ok := stage.Xhtml_h5_types[xhtml_h5_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_h5_type(xhtml_h5_type)
		}
	}
	return xhtml_h5_type
}

// for satisfaction of GongStruct interface
func (xhtml_h5_type *Xhtml_h5_type) GetName() (res string) {
	return xhtml_h5_type.Name
}

// Stage puts xhtml_h6_type to the model stage
func (xhtml_h6_type *Xhtml_h6_type) Stage(stage *Stage) *Xhtml_h6_type {

	if _, ok := stage.Xhtml_h6_types[xhtml_h6_type]; !ok {
		stage.Xhtml_h6_types[xhtml_h6_type] = __member
		stage.Xhtml_h6_typeMap_Staged_Order[xhtml_h6_type] = stage.Xhtml_h6_typeOrder
		stage.Xhtml_h6_typeOrder++
	}
	stage.Xhtml_h6_types_mapString[xhtml_h6_type.Name] = xhtml_h6_type

	return xhtml_h6_type
}

// Unstage removes xhtml_h6_type off the model stage
func (xhtml_h6_type *Xhtml_h6_type) Unstage(stage *Stage) *Xhtml_h6_type {
	delete(stage.Xhtml_h6_types, xhtml_h6_type)
	delete(stage.Xhtml_h6_types_mapString, xhtml_h6_type.Name)
	return xhtml_h6_type
}

// UnstageVoid removes xhtml_h6_type off the model stage
func (xhtml_h6_type *Xhtml_h6_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_h6_types, xhtml_h6_type)
	delete(stage.Xhtml_h6_types_mapString, xhtml_h6_type.Name)
}

// commit xhtml_h6_type to the back repo (if it is already staged)
func (xhtml_h6_type *Xhtml_h6_type) Commit(stage *Stage) *Xhtml_h6_type {
	if _, ok := stage.Xhtml_h6_types[xhtml_h6_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_h6_type(xhtml_h6_type)
		}
	}
	return xhtml_h6_type
}

func (xhtml_h6_type *Xhtml_h6_type) CommitVoid(stage *Stage) {
	xhtml_h6_type.Commit(stage)
}

// Checkout xhtml_h6_type to the back repo (if it is already staged)
func (xhtml_h6_type *Xhtml_h6_type) Checkout(stage *Stage) *Xhtml_h6_type {
	if _, ok := stage.Xhtml_h6_types[xhtml_h6_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_h6_type(xhtml_h6_type)
		}
	}
	return xhtml_h6_type
}

// for satisfaction of GongStruct interface
func (xhtml_h6_type *Xhtml_h6_type) GetName() (res string) {
	return xhtml_h6_type.Name
}

// Stage puts xhtml_heading_type to the model stage
func (xhtml_heading_type *Xhtml_heading_type) Stage(stage *Stage) *Xhtml_heading_type {

	if _, ok := stage.Xhtml_heading_types[xhtml_heading_type]; !ok {
		stage.Xhtml_heading_types[xhtml_heading_type] = __member
		stage.Xhtml_heading_typeMap_Staged_Order[xhtml_heading_type] = stage.Xhtml_heading_typeOrder
		stage.Xhtml_heading_typeOrder++
	}
	stage.Xhtml_heading_types_mapString[xhtml_heading_type.Name] = xhtml_heading_type

	return xhtml_heading_type
}

// Unstage removes xhtml_heading_type off the model stage
func (xhtml_heading_type *Xhtml_heading_type) Unstage(stage *Stage) *Xhtml_heading_type {
	delete(stage.Xhtml_heading_types, xhtml_heading_type)
	delete(stage.Xhtml_heading_types_mapString, xhtml_heading_type.Name)
	return xhtml_heading_type
}

// UnstageVoid removes xhtml_heading_type off the model stage
func (xhtml_heading_type *Xhtml_heading_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_heading_types, xhtml_heading_type)
	delete(stage.Xhtml_heading_types_mapString, xhtml_heading_type.Name)
}

// commit xhtml_heading_type to the back repo (if it is already staged)
func (xhtml_heading_type *Xhtml_heading_type) Commit(stage *Stage) *Xhtml_heading_type {
	if _, ok := stage.Xhtml_heading_types[xhtml_heading_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_heading_type(xhtml_heading_type)
		}
	}
	return xhtml_heading_type
}

func (xhtml_heading_type *Xhtml_heading_type) CommitVoid(stage *Stage) {
	xhtml_heading_type.Commit(stage)
}

// Checkout xhtml_heading_type to the back repo (if it is already staged)
func (xhtml_heading_type *Xhtml_heading_type) Checkout(stage *Stage) *Xhtml_heading_type {
	if _, ok := stage.Xhtml_heading_types[xhtml_heading_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_heading_type(xhtml_heading_type)
		}
	}
	return xhtml_heading_type
}

// for satisfaction of GongStruct interface
func (xhtml_heading_type *Xhtml_heading_type) GetName() (res string) {
	return xhtml_heading_type.Name
}

// Stage puts xhtml_hr_type to the model stage
func (xhtml_hr_type *Xhtml_hr_type) Stage(stage *Stage) *Xhtml_hr_type {

	if _, ok := stage.Xhtml_hr_types[xhtml_hr_type]; !ok {
		stage.Xhtml_hr_types[xhtml_hr_type] = __member
		stage.Xhtml_hr_typeMap_Staged_Order[xhtml_hr_type] = stage.Xhtml_hr_typeOrder
		stage.Xhtml_hr_typeOrder++
	}
	stage.Xhtml_hr_types_mapString[xhtml_hr_type.Name] = xhtml_hr_type

	return xhtml_hr_type
}

// Unstage removes xhtml_hr_type off the model stage
func (xhtml_hr_type *Xhtml_hr_type) Unstage(stage *Stage) *Xhtml_hr_type {
	delete(stage.Xhtml_hr_types, xhtml_hr_type)
	delete(stage.Xhtml_hr_types_mapString, xhtml_hr_type.Name)
	return xhtml_hr_type
}

// UnstageVoid removes xhtml_hr_type off the model stage
func (xhtml_hr_type *Xhtml_hr_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_hr_types, xhtml_hr_type)
	delete(stage.Xhtml_hr_types_mapString, xhtml_hr_type.Name)
}

// commit xhtml_hr_type to the back repo (if it is already staged)
func (xhtml_hr_type *Xhtml_hr_type) Commit(stage *Stage) *Xhtml_hr_type {
	if _, ok := stage.Xhtml_hr_types[xhtml_hr_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_hr_type(xhtml_hr_type)
		}
	}
	return xhtml_hr_type
}

func (xhtml_hr_type *Xhtml_hr_type) CommitVoid(stage *Stage) {
	xhtml_hr_type.Commit(stage)
}

// Checkout xhtml_hr_type to the back repo (if it is already staged)
func (xhtml_hr_type *Xhtml_hr_type) Checkout(stage *Stage) *Xhtml_hr_type {
	if _, ok := stage.Xhtml_hr_types[xhtml_hr_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_hr_type(xhtml_hr_type)
		}
	}
	return xhtml_hr_type
}

// for satisfaction of GongStruct interface
func (xhtml_hr_type *Xhtml_hr_type) GetName() (res string) {
	return xhtml_hr_type.Name
}

// Stage puts xhtml_kbd_type to the model stage
func (xhtml_kbd_type *Xhtml_kbd_type) Stage(stage *Stage) *Xhtml_kbd_type {

	if _, ok := stage.Xhtml_kbd_types[xhtml_kbd_type]; !ok {
		stage.Xhtml_kbd_types[xhtml_kbd_type] = __member
		stage.Xhtml_kbd_typeMap_Staged_Order[xhtml_kbd_type] = stage.Xhtml_kbd_typeOrder
		stage.Xhtml_kbd_typeOrder++
	}
	stage.Xhtml_kbd_types_mapString[xhtml_kbd_type.Name] = xhtml_kbd_type

	return xhtml_kbd_type
}

// Unstage removes xhtml_kbd_type off the model stage
func (xhtml_kbd_type *Xhtml_kbd_type) Unstage(stage *Stage) *Xhtml_kbd_type {
	delete(stage.Xhtml_kbd_types, xhtml_kbd_type)
	delete(stage.Xhtml_kbd_types_mapString, xhtml_kbd_type.Name)
	return xhtml_kbd_type
}

// UnstageVoid removes xhtml_kbd_type off the model stage
func (xhtml_kbd_type *Xhtml_kbd_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_kbd_types, xhtml_kbd_type)
	delete(stage.Xhtml_kbd_types_mapString, xhtml_kbd_type.Name)
}

// commit xhtml_kbd_type to the back repo (if it is already staged)
func (xhtml_kbd_type *Xhtml_kbd_type) Commit(stage *Stage) *Xhtml_kbd_type {
	if _, ok := stage.Xhtml_kbd_types[xhtml_kbd_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_kbd_type(xhtml_kbd_type)
		}
	}
	return xhtml_kbd_type
}

func (xhtml_kbd_type *Xhtml_kbd_type) CommitVoid(stage *Stage) {
	xhtml_kbd_type.Commit(stage)
}

// Checkout xhtml_kbd_type to the back repo (if it is already staged)
func (xhtml_kbd_type *Xhtml_kbd_type) Checkout(stage *Stage) *Xhtml_kbd_type {
	if _, ok := stage.Xhtml_kbd_types[xhtml_kbd_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_kbd_type(xhtml_kbd_type)
		}
	}
	return xhtml_kbd_type
}

// for satisfaction of GongStruct interface
func (xhtml_kbd_type *Xhtml_kbd_type) GetName() (res string) {
	return xhtml_kbd_type.Name
}

// Stage puts xhtml_li_type to the model stage
func (xhtml_li_type *Xhtml_li_type) Stage(stage *Stage) *Xhtml_li_type {

	if _, ok := stage.Xhtml_li_types[xhtml_li_type]; !ok {
		stage.Xhtml_li_types[xhtml_li_type] = __member
		stage.Xhtml_li_typeMap_Staged_Order[xhtml_li_type] = stage.Xhtml_li_typeOrder
		stage.Xhtml_li_typeOrder++
	}
	stage.Xhtml_li_types_mapString[xhtml_li_type.Name] = xhtml_li_type

	return xhtml_li_type
}

// Unstage removes xhtml_li_type off the model stage
func (xhtml_li_type *Xhtml_li_type) Unstage(stage *Stage) *Xhtml_li_type {
	delete(stage.Xhtml_li_types, xhtml_li_type)
	delete(stage.Xhtml_li_types_mapString, xhtml_li_type.Name)
	return xhtml_li_type
}

// UnstageVoid removes xhtml_li_type off the model stage
func (xhtml_li_type *Xhtml_li_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_li_types, xhtml_li_type)
	delete(stage.Xhtml_li_types_mapString, xhtml_li_type.Name)
}

// commit xhtml_li_type to the back repo (if it is already staged)
func (xhtml_li_type *Xhtml_li_type) Commit(stage *Stage) *Xhtml_li_type {
	if _, ok := stage.Xhtml_li_types[xhtml_li_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_li_type(xhtml_li_type)
		}
	}
	return xhtml_li_type
}

func (xhtml_li_type *Xhtml_li_type) CommitVoid(stage *Stage) {
	xhtml_li_type.Commit(stage)
}

// Checkout xhtml_li_type to the back repo (if it is already staged)
func (xhtml_li_type *Xhtml_li_type) Checkout(stage *Stage) *Xhtml_li_type {
	if _, ok := stage.Xhtml_li_types[xhtml_li_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_li_type(xhtml_li_type)
		}
	}
	return xhtml_li_type
}

// for satisfaction of GongStruct interface
func (xhtml_li_type *Xhtml_li_type) GetName() (res string) {
	return xhtml_li_type.Name
}

// Stage puts xhtml_object_type to the model stage
func (xhtml_object_type *Xhtml_object_type) Stage(stage *Stage) *Xhtml_object_type {

	if _, ok := stage.Xhtml_object_types[xhtml_object_type]; !ok {
		stage.Xhtml_object_types[xhtml_object_type] = __member
		stage.Xhtml_object_typeMap_Staged_Order[xhtml_object_type] = stage.Xhtml_object_typeOrder
		stage.Xhtml_object_typeOrder++
	}
	stage.Xhtml_object_types_mapString[xhtml_object_type.Name] = xhtml_object_type

	return xhtml_object_type
}

// Unstage removes xhtml_object_type off the model stage
func (xhtml_object_type *Xhtml_object_type) Unstage(stage *Stage) *Xhtml_object_type {
	delete(stage.Xhtml_object_types, xhtml_object_type)
	delete(stage.Xhtml_object_types_mapString, xhtml_object_type.Name)
	return xhtml_object_type
}

// UnstageVoid removes xhtml_object_type off the model stage
func (xhtml_object_type *Xhtml_object_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_object_types, xhtml_object_type)
	delete(stage.Xhtml_object_types_mapString, xhtml_object_type.Name)
}

// commit xhtml_object_type to the back repo (if it is already staged)
func (xhtml_object_type *Xhtml_object_type) Commit(stage *Stage) *Xhtml_object_type {
	if _, ok := stage.Xhtml_object_types[xhtml_object_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_object_type(xhtml_object_type)
		}
	}
	return xhtml_object_type
}

func (xhtml_object_type *Xhtml_object_type) CommitVoid(stage *Stage) {
	xhtml_object_type.Commit(stage)
}

// Checkout xhtml_object_type to the back repo (if it is already staged)
func (xhtml_object_type *Xhtml_object_type) Checkout(stage *Stage) *Xhtml_object_type {
	if _, ok := stage.Xhtml_object_types[xhtml_object_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_object_type(xhtml_object_type)
		}
	}
	return xhtml_object_type
}

// for satisfaction of GongStruct interface
func (xhtml_object_type *Xhtml_object_type) GetName() (res string) {
	return xhtml_object_type.Name
}

// Stage puts xhtml_ol_type to the model stage
func (xhtml_ol_type *Xhtml_ol_type) Stage(stage *Stage) *Xhtml_ol_type {

	if _, ok := stage.Xhtml_ol_types[xhtml_ol_type]; !ok {
		stage.Xhtml_ol_types[xhtml_ol_type] = __member
		stage.Xhtml_ol_typeMap_Staged_Order[xhtml_ol_type] = stage.Xhtml_ol_typeOrder
		stage.Xhtml_ol_typeOrder++
	}
	stage.Xhtml_ol_types_mapString[xhtml_ol_type.Name] = xhtml_ol_type

	return xhtml_ol_type
}

// Unstage removes xhtml_ol_type off the model stage
func (xhtml_ol_type *Xhtml_ol_type) Unstage(stage *Stage) *Xhtml_ol_type {
	delete(stage.Xhtml_ol_types, xhtml_ol_type)
	delete(stage.Xhtml_ol_types_mapString, xhtml_ol_type.Name)
	return xhtml_ol_type
}

// UnstageVoid removes xhtml_ol_type off the model stage
func (xhtml_ol_type *Xhtml_ol_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_ol_types, xhtml_ol_type)
	delete(stage.Xhtml_ol_types_mapString, xhtml_ol_type.Name)
}

// commit xhtml_ol_type to the back repo (if it is already staged)
func (xhtml_ol_type *Xhtml_ol_type) Commit(stage *Stage) *Xhtml_ol_type {
	if _, ok := stage.Xhtml_ol_types[xhtml_ol_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_ol_type(xhtml_ol_type)
		}
	}
	return xhtml_ol_type
}

func (xhtml_ol_type *Xhtml_ol_type) CommitVoid(stage *Stage) {
	xhtml_ol_type.Commit(stage)
}

// Checkout xhtml_ol_type to the back repo (if it is already staged)
func (xhtml_ol_type *Xhtml_ol_type) Checkout(stage *Stage) *Xhtml_ol_type {
	if _, ok := stage.Xhtml_ol_types[xhtml_ol_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_ol_type(xhtml_ol_type)
		}
	}
	return xhtml_ol_type
}

// for satisfaction of GongStruct interface
func (xhtml_ol_type *Xhtml_ol_type) GetName() (res string) {
	return xhtml_ol_type.Name
}

// Stage puts xhtml_p_type to the model stage
func (xhtml_p_type *Xhtml_p_type) Stage(stage *Stage) *Xhtml_p_type {

	if _, ok := stage.Xhtml_p_types[xhtml_p_type]; !ok {
		stage.Xhtml_p_types[xhtml_p_type] = __member
		stage.Xhtml_p_typeMap_Staged_Order[xhtml_p_type] = stage.Xhtml_p_typeOrder
		stage.Xhtml_p_typeOrder++
	}
	stage.Xhtml_p_types_mapString[xhtml_p_type.Name] = xhtml_p_type

	return xhtml_p_type
}

// Unstage removes xhtml_p_type off the model stage
func (xhtml_p_type *Xhtml_p_type) Unstage(stage *Stage) *Xhtml_p_type {
	delete(stage.Xhtml_p_types, xhtml_p_type)
	delete(stage.Xhtml_p_types_mapString, xhtml_p_type.Name)
	return xhtml_p_type
}

// UnstageVoid removes xhtml_p_type off the model stage
func (xhtml_p_type *Xhtml_p_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_p_types, xhtml_p_type)
	delete(stage.Xhtml_p_types_mapString, xhtml_p_type.Name)
}

// commit xhtml_p_type to the back repo (if it is already staged)
func (xhtml_p_type *Xhtml_p_type) Commit(stage *Stage) *Xhtml_p_type {
	if _, ok := stage.Xhtml_p_types[xhtml_p_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_p_type(xhtml_p_type)
		}
	}
	return xhtml_p_type
}

func (xhtml_p_type *Xhtml_p_type) CommitVoid(stage *Stage) {
	xhtml_p_type.Commit(stage)
}

// Checkout xhtml_p_type to the back repo (if it is already staged)
func (xhtml_p_type *Xhtml_p_type) Checkout(stage *Stage) *Xhtml_p_type {
	if _, ok := stage.Xhtml_p_types[xhtml_p_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_p_type(xhtml_p_type)
		}
	}
	return xhtml_p_type
}

// for satisfaction of GongStruct interface
func (xhtml_p_type *Xhtml_p_type) GetName() (res string) {
	return xhtml_p_type.Name
}

// Stage puts xhtml_param_type to the model stage
func (xhtml_param_type *Xhtml_param_type) Stage(stage *Stage) *Xhtml_param_type {

	if _, ok := stage.Xhtml_param_types[xhtml_param_type]; !ok {
		stage.Xhtml_param_types[xhtml_param_type] = __member
		stage.Xhtml_param_typeMap_Staged_Order[xhtml_param_type] = stage.Xhtml_param_typeOrder
		stage.Xhtml_param_typeOrder++
	}
	stage.Xhtml_param_types_mapString[xhtml_param_type.Name] = xhtml_param_type

	return xhtml_param_type
}

// Unstage removes xhtml_param_type off the model stage
func (xhtml_param_type *Xhtml_param_type) Unstage(stage *Stage) *Xhtml_param_type {
	delete(stage.Xhtml_param_types, xhtml_param_type)
	delete(stage.Xhtml_param_types_mapString, xhtml_param_type.Name)
	return xhtml_param_type
}

// UnstageVoid removes xhtml_param_type off the model stage
func (xhtml_param_type *Xhtml_param_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_param_types, xhtml_param_type)
	delete(stage.Xhtml_param_types_mapString, xhtml_param_type.Name)
}

// commit xhtml_param_type to the back repo (if it is already staged)
func (xhtml_param_type *Xhtml_param_type) Commit(stage *Stage) *Xhtml_param_type {
	if _, ok := stage.Xhtml_param_types[xhtml_param_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_param_type(xhtml_param_type)
		}
	}
	return xhtml_param_type
}

func (xhtml_param_type *Xhtml_param_type) CommitVoid(stage *Stage) {
	xhtml_param_type.Commit(stage)
}

// Checkout xhtml_param_type to the back repo (if it is already staged)
func (xhtml_param_type *Xhtml_param_type) Checkout(stage *Stage) *Xhtml_param_type {
	if _, ok := stage.Xhtml_param_types[xhtml_param_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_param_type(xhtml_param_type)
		}
	}
	return xhtml_param_type
}

// for satisfaction of GongStruct interface
func (xhtml_param_type *Xhtml_param_type) GetName() (res string) {
	return xhtml_param_type.Name
}

// Stage puts xhtml_pre_type to the model stage
func (xhtml_pre_type *Xhtml_pre_type) Stage(stage *Stage) *Xhtml_pre_type {

	if _, ok := stage.Xhtml_pre_types[xhtml_pre_type]; !ok {
		stage.Xhtml_pre_types[xhtml_pre_type] = __member
		stage.Xhtml_pre_typeMap_Staged_Order[xhtml_pre_type] = stage.Xhtml_pre_typeOrder
		stage.Xhtml_pre_typeOrder++
	}
	stage.Xhtml_pre_types_mapString[xhtml_pre_type.Name] = xhtml_pre_type

	return xhtml_pre_type
}

// Unstage removes xhtml_pre_type off the model stage
func (xhtml_pre_type *Xhtml_pre_type) Unstage(stage *Stage) *Xhtml_pre_type {
	delete(stage.Xhtml_pre_types, xhtml_pre_type)
	delete(stage.Xhtml_pre_types_mapString, xhtml_pre_type.Name)
	return xhtml_pre_type
}

// UnstageVoid removes xhtml_pre_type off the model stage
func (xhtml_pre_type *Xhtml_pre_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_pre_types, xhtml_pre_type)
	delete(stage.Xhtml_pre_types_mapString, xhtml_pre_type.Name)
}

// commit xhtml_pre_type to the back repo (if it is already staged)
func (xhtml_pre_type *Xhtml_pre_type) Commit(stage *Stage) *Xhtml_pre_type {
	if _, ok := stage.Xhtml_pre_types[xhtml_pre_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_pre_type(xhtml_pre_type)
		}
	}
	return xhtml_pre_type
}

func (xhtml_pre_type *Xhtml_pre_type) CommitVoid(stage *Stage) {
	xhtml_pre_type.Commit(stage)
}

// Checkout xhtml_pre_type to the back repo (if it is already staged)
func (xhtml_pre_type *Xhtml_pre_type) Checkout(stage *Stage) *Xhtml_pre_type {
	if _, ok := stage.Xhtml_pre_types[xhtml_pre_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_pre_type(xhtml_pre_type)
		}
	}
	return xhtml_pre_type
}

// for satisfaction of GongStruct interface
func (xhtml_pre_type *Xhtml_pre_type) GetName() (res string) {
	return xhtml_pre_type.Name
}

// Stage puts xhtml_q_type to the model stage
func (xhtml_q_type *Xhtml_q_type) Stage(stage *Stage) *Xhtml_q_type {

	if _, ok := stage.Xhtml_q_types[xhtml_q_type]; !ok {
		stage.Xhtml_q_types[xhtml_q_type] = __member
		stage.Xhtml_q_typeMap_Staged_Order[xhtml_q_type] = stage.Xhtml_q_typeOrder
		stage.Xhtml_q_typeOrder++
	}
	stage.Xhtml_q_types_mapString[xhtml_q_type.Name] = xhtml_q_type

	return xhtml_q_type
}

// Unstage removes xhtml_q_type off the model stage
func (xhtml_q_type *Xhtml_q_type) Unstage(stage *Stage) *Xhtml_q_type {
	delete(stage.Xhtml_q_types, xhtml_q_type)
	delete(stage.Xhtml_q_types_mapString, xhtml_q_type.Name)
	return xhtml_q_type
}

// UnstageVoid removes xhtml_q_type off the model stage
func (xhtml_q_type *Xhtml_q_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_q_types, xhtml_q_type)
	delete(stage.Xhtml_q_types_mapString, xhtml_q_type.Name)
}

// commit xhtml_q_type to the back repo (if it is already staged)
func (xhtml_q_type *Xhtml_q_type) Commit(stage *Stage) *Xhtml_q_type {
	if _, ok := stage.Xhtml_q_types[xhtml_q_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_q_type(xhtml_q_type)
		}
	}
	return xhtml_q_type
}

func (xhtml_q_type *Xhtml_q_type) CommitVoid(stage *Stage) {
	xhtml_q_type.Commit(stage)
}

// Checkout xhtml_q_type to the back repo (if it is already staged)
func (xhtml_q_type *Xhtml_q_type) Checkout(stage *Stage) *Xhtml_q_type {
	if _, ok := stage.Xhtml_q_types[xhtml_q_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_q_type(xhtml_q_type)
		}
	}
	return xhtml_q_type
}

// for satisfaction of GongStruct interface
func (xhtml_q_type *Xhtml_q_type) GetName() (res string) {
	return xhtml_q_type.Name
}

// Stage puts xhtml_samp_type to the model stage
func (xhtml_samp_type *Xhtml_samp_type) Stage(stage *Stage) *Xhtml_samp_type {

	if _, ok := stage.Xhtml_samp_types[xhtml_samp_type]; !ok {
		stage.Xhtml_samp_types[xhtml_samp_type] = __member
		stage.Xhtml_samp_typeMap_Staged_Order[xhtml_samp_type] = stage.Xhtml_samp_typeOrder
		stage.Xhtml_samp_typeOrder++
	}
	stage.Xhtml_samp_types_mapString[xhtml_samp_type.Name] = xhtml_samp_type

	return xhtml_samp_type
}

// Unstage removes xhtml_samp_type off the model stage
func (xhtml_samp_type *Xhtml_samp_type) Unstage(stage *Stage) *Xhtml_samp_type {
	delete(stage.Xhtml_samp_types, xhtml_samp_type)
	delete(stage.Xhtml_samp_types_mapString, xhtml_samp_type.Name)
	return xhtml_samp_type
}

// UnstageVoid removes xhtml_samp_type off the model stage
func (xhtml_samp_type *Xhtml_samp_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_samp_types, xhtml_samp_type)
	delete(stage.Xhtml_samp_types_mapString, xhtml_samp_type.Name)
}

// commit xhtml_samp_type to the back repo (if it is already staged)
func (xhtml_samp_type *Xhtml_samp_type) Commit(stage *Stage) *Xhtml_samp_type {
	if _, ok := stage.Xhtml_samp_types[xhtml_samp_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_samp_type(xhtml_samp_type)
		}
	}
	return xhtml_samp_type
}

func (xhtml_samp_type *Xhtml_samp_type) CommitVoid(stage *Stage) {
	xhtml_samp_type.Commit(stage)
}

// Checkout xhtml_samp_type to the back repo (if it is already staged)
func (xhtml_samp_type *Xhtml_samp_type) Checkout(stage *Stage) *Xhtml_samp_type {
	if _, ok := stage.Xhtml_samp_types[xhtml_samp_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_samp_type(xhtml_samp_type)
		}
	}
	return xhtml_samp_type
}

// for satisfaction of GongStruct interface
func (xhtml_samp_type *Xhtml_samp_type) GetName() (res string) {
	return xhtml_samp_type.Name
}

// Stage puts xhtml_span_type to the model stage
func (xhtml_span_type *Xhtml_span_type) Stage(stage *Stage) *Xhtml_span_type {

	if _, ok := stage.Xhtml_span_types[xhtml_span_type]; !ok {
		stage.Xhtml_span_types[xhtml_span_type] = __member
		stage.Xhtml_span_typeMap_Staged_Order[xhtml_span_type] = stage.Xhtml_span_typeOrder
		stage.Xhtml_span_typeOrder++
	}
	stage.Xhtml_span_types_mapString[xhtml_span_type.Name] = xhtml_span_type

	return xhtml_span_type
}

// Unstage removes xhtml_span_type off the model stage
func (xhtml_span_type *Xhtml_span_type) Unstage(stage *Stage) *Xhtml_span_type {
	delete(stage.Xhtml_span_types, xhtml_span_type)
	delete(stage.Xhtml_span_types_mapString, xhtml_span_type.Name)
	return xhtml_span_type
}

// UnstageVoid removes xhtml_span_type off the model stage
func (xhtml_span_type *Xhtml_span_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_span_types, xhtml_span_type)
	delete(stage.Xhtml_span_types_mapString, xhtml_span_type.Name)
}

// commit xhtml_span_type to the back repo (if it is already staged)
func (xhtml_span_type *Xhtml_span_type) Commit(stage *Stage) *Xhtml_span_type {
	if _, ok := stage.Xhtml_span_types[xhtml_span_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_span_type(xhtml_span_type)
		}
	}
	return xhtml_span_type
}

func (xhtml_span_type *Xhtml_span_type) CommitVoid(stage *Stage) {
	xhtml_span_type.Commit(stage)
}

// Checkout xhtml_span_type to the back repo (if it is already staged)
func (xhtml_span_type *Xhtml_span_type) Checkout(stage *Stage) *Xhtml_span_type {
	if _, ok := stage.Xhtml_span_types[xhtml_span_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_span_type(xhtml_span_type)
		}
	}
	return xhtml_span_type
}

// for satisfaction of GongStruct interface
func (xhtml_span_type *Xhtml_span_type) GetName() (res string) {
	return xhtml_span_type.Name
}

// Stage puts xhtml_strong_type to the model stage
func (xhtml_strong_type *Xhtml_strong_type) Stage(stage *Stage) *Xhtml_strong_type {

	if _, ok := stage.Xhtml_strong_types[xhtml_strong_type]; !ok {
		stage.Xhtml_strong_types[xhtml_strong_type] = __member
		stage.Xhtml_strong_typeMap_Staged_Order[xhtml_strong_type] = stage.Xhtml_strong_typeOrder
		stage.Xhtml_strong_typeOrder++
	}
	stage.Xhtml_strong_types_mapString[xhtml_strong_type.Name] = xhtml_strong_type

	return xhtml_strong_type
}

// Unstage removes xhtml_strong_type off the model stage
func (xhtml_strong_type *Xhtml_strong_type) Unstage(stage *Stage) *Xhtml_strong_type {
	delete(stage.Xhtml_strong_types, xhtml_strong_type)
	delete(stage.Xhtml_strong_types_mapString, xhtml_strong_type.Name)
	return xhtml_strong_type
}

// UnstageVoid removes xhtml_strong_type off the model stage
func (xhtml_strong_type *Xhtml_strong_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_strong_types, xhtml_strong_type)
	delete(stage.Xhtml_strong_types_mapString, xhtml_strong_type.Name)
}

// commit xhtml_strong_type to the back repo (if it is already staged)
func (xhtml_strong_type *Xhtml_strong_type) Commit(stage *Stage) *Xhtml_strong_type {
	if _, ok := stage.Xhtml_strong_types[xhtml_strong_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_strong_type(xhtml_strong_type)
		}
	}
	return xhtml_strong_type
}

func (xhtml_strong_type *Xhtml_strong_type) CommitVoid(stage *Stage) {
	xhtml_strong_type.Commit(stage)
}

// Checkout xhtml_strong_type to the back repo (if it is already staged)
func (xhtml_strong_type *Xhtml_strong_type) Checkout(stage *Stage) *Xhtml_strong_type {
	if _, ok := stage.Xhtml_strong_types[xhtml_strong_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_strong_type(xhtml_strong_type)
		}
	}
	return xhtml_strong_type
}

// for satisfaction of GongStruct interface
func (xhtml_strong_type *Xhtml_strong_type) GetName() (res string) {
	return xhtml_strong_type.Name
}

// Stage puts xhtml_table_type to the model stage
func (xhtml_table_type *Xhtml_table_type) Stage(stage *Stage) *Xhtml_table_type {

	if _, ok := stage.Xhtml_table_types[xhtml_table_type]; !ok {
		stage.Xhtml_table_types[xhtml_table_type] = __member
		stage.Xhtml_table_typeMap_Staged_Order[xhtml_table_type] = stage.Xhtml_table_typeOrder
		stage.Xhtml_table_typeOrder++
	}
	stage.Xhtml_table_types_mapString[xhtml_table_type.Name] = xhtml_table_type

	return xhtml_table_type
}

// Unstage removes xhtml_table_type off the model stage
func (xhtml_table_type *Xhtml_table_type) Unstage(stage *Stage) *Xhtml_table_type {
	delete(stage.Xhtml_table_types, xhtml_table_type)
	delete(stage.Xhtml_table_types_mapString, xhtml_table_type.Name)
	return xhtml_table_type
}

// UnstageVoid removes xhtml_table_type off the model stage
func (xhtml_table_type *Xhtml_table_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_table_types, xhtml_table_type)
	delete(stage.Xhtml_table_types_mapString, xhtml_table_type.Name)
}

// commit xhtml_table_type to the back repo (if it is already staged)
func (xhtml_table_type *Xhtml_table_type) Commit(stage *Stage) *Xhtml_table_type {
	if _, ok := stage.Xhtml_table_types[xhtml_table_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_table_type(xhtml_table_type)
		}
	}
	return xhtml_table_type
}

func (xhtml_table_type *Xhtml_table_type) CommitVoid(stage *Stage) {
	xhtml_table_type.Commit(stage)
}

// Checkout xhtml_table_type to the back repo (if it is already staged)
func (xhtml_table_type *Xhtml_table_type) Checkout(stage *Stage) *Xhtml_table_type {
	if _, ok := stage.Xhtml_table_types[xhtml_table_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_table_type(xhtml_table_type)
		}
	}
	return xhtml_table_type
}

// for satisfaction of GongStruct interface
func (xhtml_table_type *Xhtml_table_type) GetName() (res string) {
	return xhtml_table_type.Name
}

// Stage puts xhtml_tbody_type to the model stage
func (xhtml_tbody_type *Xhtml_tbody_type) Stage(stage *Stage) *Xhtml_tbody_type {

	if _, ok := stage.Xhtml_tbody_types[xhtml_tbody_type]; !ok {
		stage.Xhtml_tbody_types[xhtml_tbody_type] = __member
		stage.Xhtml_tbody_typeMap_Staged_Order[xhtml_tbody_type] = stage.Xhtml_tbody_typeOrder
		stage.Xhtml_tbody_typeOrder++
	}
	stage.Xhtml_tbody_types_mapString[xhtml_tbody_type.Name] = xhtml_tbody_type

	return xhtml_tbody_type
}

// Unstage removes xhtml_tbody_type off the model stage
func (xhtml_tbody_type *Xhtml_tbody_type) Unstage(stage *Stage) *Xhtml_tbody_type {
	delete(stage.Xhtml_tbody_types, xhtml_tbody_type)
	delete(stage.Xhtml_tbody_types_mapString, xhtml_tbody_type.Name)
	return xhtml_tbody_type
}

// UnstageVoid removes xhtml_tbody_type off the model stage
func (xhtml_tbody_type *Xhtml_tbody_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_tbody_types, xhtml_tbody_type)
	delete(stage.Xhtml_tbody_types_mapString, xhtml_tbody_type.Name)
}

// commit xhtml_tbody_type to the back repo (if it is already staged)
func (xhtml_tbody_type *Xhtml_tbody_type) Commit(stage *Stage) *Xhtml_tbody_type {
	if _, ok := stage.Xhtml_tbody_types[xhtml_tbody_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_tbody_type(xhtml_tbody_type)
		}
	}
	return xhtml_tbody_type
}

func (xhtml_tbody_type *Xhtml_tbody_type) CommitVoid(stage *Stage) {
	xhtml_tbody_type.Commit(stage)
}

// Checkout xhtml_tbody_type to the back repo (if it is already staged)
func (xhtml_tbody_type *Xhtml_tbody_type) Checkout(stage *Stage) *Xhtml_tbody_type {
	if _, ok := stage.Xhtml_tbody_types[xhtml_tbody_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_tbody_type(xhtml_tbody_type)
		}
	}
	return xhtml_tbody_type
}

// for satisfaction of GongStruct interface
func (xhtml_tbody_type *Xhtml_tbody_type) GetName() (res string) {
	return xhtml_tbody_type.Name
}

// Stage puts xhtml_td_type to the model stage
func (xhtml_td_type *Xhtml_td_type) Stage(stage *Stage) *Xhtml_td_type {

	if _, ok := stage.Xhtml_td_types[xhtml_td_type]; !ok {
		stage.Xhtml_td_types[xhtml_td_type] = __member
		stage.Xhtml_td_typeMap_Staged_Order[xhtml_td_type] = stage.Xhtml_td_typeOrder
		stage.Xhtml_td_typeOrder++
	}
	stage.Xhtml_td_types_mapString[xhtml_td_type.Name] = xhtml_td_type

	return xhtml_td_type
}

// Unstage removes xhtml_td_type off the model stage
func (xhtml_td_type *Xhtml_td_type) Unstage(stage *Stage) *Xhtml_td_type {
	delete(stage.Xhtml_td_types, xhtml_td_type)
	delete(stage.Xhtml_td_types_mapString, xhtml_td_type.Name)
	return xhtml_td_type
}

// UnstageVoid removes xhtml_td_type off the model stage
func (xhtml_td_type *Xhtml_td_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_td_types, xhtml_td_type)
	delete(stage.Xhtml_td_types_mapString, xhtml_td_type.Name)
}

// commit xhtml_td_type to the back repo (if it is already staged)
func (xhtml_td_type *Xhtml_td_type) Commit(stage *Stage) *Xhtml_td_type {
	if _, ok := stage.Xhtml_td_types[xhtml_td_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_td_type(xhtml_td_type)
		}
	}
	return xhtml_td_type
}

func (xhtml_td_type *Xhtml_td_type) CommitVoid(stage *Stage) {
	xhtml_td_type.Commit(stage)
}

// Checkout xhtml_td_type to the back repo (if it is already staged)
func (xhtml_td_type *Xhtml_td_type) Checkout(stage *Stage) *Xhtml_td_type {
	if _, ok := stage.Xhtml_td_types[xhtml_td_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_td_type(xhtml_td_type)
		}
	}
	return xhtml_td_type
}

// for satisfaction of GongStruct interface
func (xhtml_td_type *Xhtml_td_type) GetName() (res string) {
	return xhtml_td_type.Name
}

// Stage puts xhtml_tfoot_type to the model stage
func (xhtml_tfoot_type *Xhtml_tfoot_type) Stage(stage *Stage) *Xhtml_tfoot_type {

	if _, ok := stage.Xhtml_tfoot_types[xhtml_tfoot_type]; !ok {
		stage.Xhtml_tfoot_types[xhtml_tfoot_type] = __member
		stage.Xhtml_tfoot_typeMap_Staged_Order[xhtml_tfoot_type] = stage.Xhtml_tfoot_typeOrder
		stage.Xhtml_tfoot_typeOrder++
	}
	stage.Xhtml_tfoot_types_mapString[xhtml_tfoot_type.Name] = xhtml_tfoot_type

	return xhtml_tfoot_type
}

// Unstage removes xhtml_tfoot_type off the model stage
func (xhtml_tfoot_type *Xhtml_tfoot_type) Unstage(stage *Stage) *Xhtml_tfoot_type {
	delete(stage.Xhtml_tfoot_types, xhtml_tfoot_type)
	delete(stage.Xhtml_tfoot_types_mapString, xhtml_tfoot_type.Name)
	return xhtml_tfoot_type
}

// UnstageVoid removes xhtml_tfoot_type off the model stage
func (xhtml_tfoot_type *Xhtml_tfoot_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_tfoot_types, xhtml_tfoot_type)
	delete(stage.Xhtml_tfoot_types_mapString, xhtml_tfoot_type.Name)
}

// commit xhtml_tfoot_type to the back repo (if it is already staged)
func (xhtml_tfoot_type *Xhtml_tfoot_type) Commit(stage *Stage) *Xhtml_tfoot_type {
	if _, ok := stage.Xhtml_tfoot_types[xhtml_tfoot_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_tfoot_type(xhtml_tfoot_type)
		}
	}
	return xhtml_tfoot_type
}

func (xhtml_tfoot_type *Xhtml_tfoot_type) CommitVoid(stage *Stage) {
	xhtml_tfoot_type.Commit(stage)
}

// Checkout xhtml_tfoot_type to the back repo (if it is already staged)
func (xhtml_tfoot_type *Xhtml_tfoot_type) Checkout(stage *Stage) *Xhtml_tfoot_type {
	if _, ok := stage.Xhtml_tfoot_types[xhtml_tfoot_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_tfoot_type(xhtml_tfoot_type)
		}
	}
	return xhtml_tfoot_type
}

// for satisfaction of GongStruct interface
func (xhtml_tfoot_type *Xhtml_tfoot_type) GetName() (res string) {
	return xhtml_tfoot_type.Name
}

// Stage puts xhtml_th_type to the model stage
func (xhtml_th_type *Xhtml_th_type) Stage(stage *Stage) *Xhtml_th_type {

	if _, ok := stage.Xhtml_th_types[xhtml_th_type]; !ok {
		stage.Xhtml_th_types[xhtml_th_type] = __member
		stage.Xhtml_th_typeMap_Staged_Order[xhtml_th_type] = stage.Xhtml_th_typeOrder
		stage.Xhtml_th_typeOrder++
	}
	stage.Xhtml_th_types_mapString[xhtml_th_type.Name] = xhtml_th_type

	return xhtml_th_type
}

// Unstage removes xhtml_th_type off the model stage
func (xhtml_th_type *Xhtml_th_type) Unstage(stage *Stage) *Xhtml_th_type {
	delete(stage.Xhtml_th_types, xhtml_th_type)
	delete(stage.Xhtml_th_types_mapString, xhtml_th_type.Name)
	return xhtml_th_type
}

// UnstageVoid removes xhtml_th_type off the model stage
func (xhtml_th_type *Xhtml_th_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_th_types, xhtml_th_type)
	delete(stage.Xhtml_th_types_mapString, xhtml_th_type.Name)
}

// commit xhtml_th_type to the back repo (if it is already staged)
func (xhtml_th_type *Xhtml_th_type) Commit(stage *Stage) *Xhtml_th_type {
	if _, ok := stage.Xhtml_th_types[xhtml_th_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_th_type(xhtml_th_type)
		}
	}
	return xhtml_th_type
}

func (xhtml_th_type *Xhtml_th_type) CommitVoid(stage *Stage) {
	xhtml_th_type.Commit(stage)
}

// Checkout xhtml_th_type to the back repo (if it is already staged)
func (xhtml_th_type *Xhtml_th_type) Checkout(stage *Stage) *Xhtml_th_type {
	if _, ok := stage.Xhtml_th_types[xhtml_th_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_th_type(xhtml_th_type)
		}
	}
	return xhtml_th_type
}

// for satisfaction of GongStruct interface
func (xhtml_th_type *Xhtml_th_type) GetName() (res string) {
	return xhtml_th_type.Name
}

// Stage puts xhtml_thead_type to the model stage
func (xhtml_thead_type *Xhtml_thead_type) Stage(stage *Stage) *Xhtml_thead_type {

	if _, ok := stage.Xhtml_thead_types[xhtml_thead_type]; !ok {
		stage.Xhtml_thead_types[xhtml_thead_type] = __member
		stage.Xhtml_thead_typeMap_Staged_Order[xhtml_thead_type] = stage.Xhtml_thead_typeOrder
		stage.Xhtml_thead_typeOrder++
	}
	stage.Xhtml_thead_types_mapString[xhtml_thead_type.Name] = xhtml_thead_type

	return xhtml_thead_type
}

// Unstage removes xhtml_thead_type off the model stage
func (xhtml_thead_type *Xhtml_thead_type) Unstage(stage *Stage) *Xhtml_thead_type {
	delete(stage.Xhtml_thead_types, xhtml_thead_type)
	delete(stage.Xhtml_thead_types_mapString, xhtml_thead_type.Name)
	return xhtml_thead_type
}

// UnstageVoid removes xhtml_thead_type off the model stage
func (xhtml_thead_type *Xhtml_thead_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_thead_types, xhtml_thead_type)
	delete(stage.Xhtml_thead_types_mapString, xhtml_thead_type.Name)
}

// commit xhtml_thead_type to the back repo (if it is already staged)
func (xhtml_thead_type *Xhtml_thead_type) Commit(stage *Stage) *Xhtml_thead_type {
	if _, ok := stage.Xhtml_thead_types[xhtml_thead_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_thead_type(xhtml_thead_type)
		}
	}
	return xhtml_thead_type
}

func (xhtml_thead_type *Xhtml_thead_type) CommitVoid(stage *Stage) {
	xhtml_thead_type.Commit(stage)
}

// Checkout xhtml_thead_type to the back repo (if it is already staged)
func (xhtml_thead_type *Xhtml_thead_type) Checkout(stage *Stage) *Xhtml_thead_type {
	if _, ok := stage.Xhtml_thead_types[xhtml_thead_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_thead_type(xhtml_thead_type)
		}
	}
	return xhtml_thead_type
}

// for satisfaction of GongStruct interface
func (xhtml_thead_type *Xhtml_thead_type) GetName() (res string) {
	return xhtml_thead_type.Name
}

// Stage puts xhtml_tr_type to the model stage
func (xhtml_tr_type *Xhtml_tr_type) Stage(stage *Stage) *Xhtml_tr_type {

	if _, ok := stage.Xhtml_tr_types[xhtml_tr_type]; !ok {
		stage.Xhtml_tr_types[xhtml_tr_type] = __member
		stage.Xhtml_tr_typeMap_Staged_Order[xhtml_tr_type] = stage.Xhtml_tr_typeOrder
		stage.Xhtml_tr_typeOrder++
	}
	stage.Xhtml_tr_types_mapString[xhtml_tr_type.Name] = xhtml_tr_type

	return xhtml_tr_type
}

// Unstage removes xhtml_tr_type off the model stage
func (xhtml_tr_type *Xhtml_tr_type) Unstage(stage *Stage) *Xhtml_tr_type {
	delete(stage.Xhtml_tr_types, xhtml_tr_type)
	delete(stage.Xhtml_tr_types_mapString, xhtml_tr_type.Name)
	return xhtml_tr_type
}

// UnstageVoid removes xhtml_tr_type off the model stage
func (xhtml_tr_type *Xhtml_tr_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_tr_types, xhtml_tr_type)
	delete(stage.Xhtml_tr_types_mapString, xhtml_tr_type.Name)
}

// commit xhtml_tr_type to the back repo (if it is already staged)
func (xhtml_tr_type *Xhtml_tr_type) Commit(stage *Stage) *Xhtml_tr_type {
	if _, ok := stage.Xhtml_tr_types[xhtml_tr_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_tr_type(xhtml_tr_type)
		}
	}
	return xhtml_tr_type
}

func (xhtml_tr_type *Xhtml_tr_type) CommitVoid(stage *Stage) {
	xhtml_tr_type.Commit(stage)
}

// Checkout xhtml_tr_type to the back repo (if it is already staged)
func (xhtml_tr_type *Xhtml_tr_type) Checkout(stage *Stage) *Xhtml_tr_type {
	if _, ok := stage.Xhtml_tr_types[xhtml_tr_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_tr_type(xhtml_tr_type)
		}
	}
	return xhtml_tr_type
}

// for satisfaction of GongStruct interface
func (xhtml_tr_type *Xhtml_tr_type) GetName() (res string) {
	return xhtml_tr_type.Name
}

// Stage puts xhtml_ul_type to the model stage
func (xhtml_ul_type *Xhtml_ul_type) Stage(stage *Stage) *Xhtml_ul_type {

	if _, ok := stage.Xhtml_ul_types[xhtml_ul_type]; !ok {
		stage.Xhtml_ul_types[xhtml_ul_type] = __member
		stage.Xhtml_ul_typeMap_Staged_Order[xhtml_ul_type] = stage.Xhtml_ul_typeOrder
		stage.Xhtml_ul_typeOrder++
	}
	stage.Xhtml_ul_types_mapString[xhtml_ul_type.Name] = xhtml_ul_type

	return xhtml_ul_type
}

// Unstage removes xhtml_ul_type off the model stage
func (xhtml_ul_type *Xhtml_ul_type) Unstage(stage *Stage) *Xhtml_ul_type {
	delete(stage.Xhtml_ul_types, xhtml_ul_type)
	delete(stage.Xhtml_ul_types_mapString, xhtml_ul_type.Name)
	return xhtml_ul_type
}

// UnstageVoid removes xhtml_ul_type off the model stage
func (xhtml_ul_type *Xhtml_ul_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_ul_types, xhtml_ul_type)
	delete(stage.Xhtml_ul_types_mapString, xhtml_ul_type.Name)
}

// commit xhtml_ul_type to the back repo (if it is already staged)
func (xhtml_ul_type *Xhtml_ul_type) Commit(stage *Stage) *Xhtml_ul_type {
	if _, ok := stage.Xhtml_ul_types[xhtml_ul_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_ul_type(xhtml_ul_type)
		}
	}
	return xhtml_ul_type
}

func (xhtml_ul_type *Xhtml_ul_type) CommitVoid(stage *Stage) {
	xhtml_ul_type.Commit(stage)
}

// Checkout xhtml_ul_type to the back repo (if it is already staged)
func (xhtml_ul_type *Xhtml_ul_type) Checkout(stage *Stage) *Xhtml_ul_type {
	if _, ok := stage.Xhtml_ul_types[xhtml_ul_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_ul_type(xhtml_ul_type)
		}
	}
	return xhtml_ul_type
}

// for satisfaction of GongStruct interface
func (xhtml_ul_type *Xhtml_ul_type) GetName() (res string) {
	return xhtml_ul_type.Name
}

// Stage puts xhtml_var_type to the model stage
func (xhtml_var_type *Xhtml_var_type) Stage(stage *Stage) *Xhtml_var_type {

	if _, ok := stage.Xhtml_var_types[xhtml_var_type]; !ok {
		stage.Xhtml_var_types[xhtml_var_type] = __member
		stage.Xhtml_var_typeMap_Staged_Order[xhtml_var_type] = stage.Xhtml_var_typeOrder
		stage.Xhtml_var_typeOrder++
	}
	stage.Xhtml_var_types_mapString[xhtml_var_type.Name] = xhtml_var_type

	return xhtml_var_type
}

// Unstage removes xhtml_var_type off the model stage
func (xhtml_var_type *Xhtml_var_type) Unstage(stage *Stage) *Xhtml_var_type {
	delete(stage.Xhtml_var_types, xhtml_var_type)
	delete(stage.Xhtml_var_types_mapString, xhtml_var_type.Name)
	return xhtml_var_type
}

// UnstageVoid removes xhtml_var_type off the model stage
func (xhtml_var_type *Xhtml_var_type) UnstageVoid(stage *Stage) {
	delete(stage.Xhtml_var_types, xhtml_var_type)
	delete(stage.Xhtml_var_types_mapString, xhtml_var_type.Name)
}

// commit xhtml_var_type to the back repo (if it is already staged)
func (xhtml_var_type *Xhtml_var_type) Commit(stage *Stage) *Xhtml_var_type {
	if _, ok := stage.Xhtml_var_types[xhtml_var_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXhtml_var_type(xhtml_var_type)
		}
	}
	return xhtml_var_type
}

func (xhtml_var_type *Xhtml_var_type) CommitVoid(stage *Stage) {
	xhtml_var_type.Commit(stage)
}

// Checkout xhtml_var_type to the back repo (if it is already staged)
func (xhtml_var_type *Xhtml_var_type) Checkout(stage *Stage) *Xhtml_var_type {
	if _, ok := stage.Xhtml_var_types[xhtml_var_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXhtml_var_type(xhtml_var_type)
		}
	}
	return xhtml_var_type
}

// for satisfaction of GongStruct interface
func (xhtml_var_type *Xhtml_var_type) GetName() (res string) {
	return xhtml_var_type.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMALTERNATIVE_ID(ALTERNATIVE_ID *ALTERNATIVE_ID)
	CreateORMATTRIBUTE_DEFINITION_BOOLEAN(ATTRIBUTE_DEFINITION_BOOLEAN *ATTRIBUTE_DEFINITION_BOOLEAN)
	CreateORMATTRIBUTE_DEFINITION_DATE(ATTRIBUTE_DEFINITION_DATE *ATTRIBUTE_DEFINITION_DATE)
	CreateORMATTRIBUTE_DEFINITION_ENUMERATION(ATTRIBUTE_DEFINITION_ENUMERATION *ATTRIBUTE_DEFINITION_ENUMERATION)
	CreateORMATTRIBUTE_DEFINITION_INTEGER(ATTRIBUTE_DEFINITION_INTEGER *ATTRIBUTE_DEFINITION_INTEGER)
	CreateORMATTRIBUTE_DEFINITION_REAL(ATTRIBUTE_DEFINITION_REAL *ATTRIBUTE_DEFINITION_REAL)
	CreateORMATTRIBUTE_DEFINITION_STRING(ATTRIBUTE_DEFINITION_STRING *ATTRIBUTE_DEFINITION_STRING)
	CreateORMATTRIBUTE_DEFINITION_XHTML(ATTRIBUTE_DEFINITION_XHTML *ATTRIBUTE_DEFINITION_XHTML)
	CreateORMATTRIBUTE_VALUE_BOOLEAN(ATTRIBUTE_VALUE_BOOLEAN *ATTRIBUTE_VALUE_BOOLEAN)
	CreateORMATTRIBUTE_VALUE_DATE(ATTRIBUTE_VALUE_DATE *ATTRIBUTE_VALUE_DATE)
	CreateORMATTRIBUTE_VALUE_ENUMERATION(ATTRIBUTE_VALUE_ENUMERATION *ATTRIBUTE_VALUE_ENUMERATION)
	CreateORMATTRIBUTE_VALUE_INTEGER(ATTRIBUTE_VALUE_INTEGER *ATTRIBUTE_VALUE_INTEGER)
	CreateORMATTRIBUTE_VALUE_REAL(ATTRIBUTE_VALUE_REAL *ATTRIBUTE_VALUE_REAL)
	CreateORMATTRIBUTE_VALUE_STRING(ATTRIBUTE_VALUE_STRING *ATTRIBUTE_VALUE_STRING)
	CreateORMATTRIBUTE_VALUE_XHTML(ATTRIBUTE_VALUE_XHTML *ATTRIBUTE_VALUE_XHTML)
	CreateORMAnyType(AnyType *AnyType)
	CreateORMDATATYPE_DEFINITION_BOOLEAN(DATATYPE_DEFINITION_BOOLEAN *DATATYPE_DEFINITION_BOOLEAN)
	CreateORMDATATYPE_DEFINITION_DATE(DATATYPE_DEFINITION_DATE *DATATYPE_DEFINITION_DATE)
	CreateORMDATATYPE_DEFINITION_ENUMERATION(DATATYPE_DEFINITION_ENUMERATION *DATATYPE_DEFINITION_ENUMERATION)
	CreateORMDATATYPE_DEFINITION_INTEGER(DATATYPE_DEFINITION_INTEGER *DATATYPE_DEFINITION_INTEGER)
	CreateORMDATATYPE_DEFINITION_REAL(DATATYPE_DEFINITION_REAL *DATATYPE_DEFINITION_REAL)
	CreateORMDATATYPE_DEFINITION_STRING(DATATYPE_DEFINITION_STRING *DATATYPE_DEFINITION_STRING)
	CreateORMDATATYPE_DEFINITION_XHTML(DATATYPE_DEFINITION_XHTML *DATATYPE_DEFINITION_XHTML)
	CreateORMEMBEDDED_VALUE(EMBEDDED_VALUE *EMBEDDED_VALUE)
	CreateORMENUM_VALUE(ENUM_VALUE *ENUM_VALUE)
	CreateORMRELATION_GROUP(RELATION_GROUP *RELATION_GROUP)
	CreateORMRELATION_GROUP_TYPE(RELATION_GROUP_TYPE *RELATION_GROUP_TYPE)
	CreateORMREQ_IF(REQ_IF *REQ_IF)
	CreateORMREQ_IF_CONTENT(REQ_IF_CONTENT *REQ_IF_CONTENT)
	CreateORMREQ_IF_HEADER(REQ_IF_HEADER *REQ_IF_HEADER)
	CreateORMREQ_IF_TOOL_EXTENSION(REQ_IF_TOOL_EXTENSION *REQ_IF_TOOL_EXTENSION)
	CreateORMSPECIFICATION(SPECIFICATION *SPECIFICATION)
	CreateORMSPECIFICATION_TYPE(SPECIFICATION_TYPE *SPECIFICATION_TYPE)
	CreateORMSPEC_HIERARCHY(SPEC_HIERARCHY *SPEC_HIERARCHY)
	CreateORMSPEC_OBJECT(SPEC_OBJECT *SPEC_OBJECT)
	CreateORMSPEC_OBJECT_TYPE(SPEC_OBJECT_TYPE *SPEC_OBJECT_TYPE)
	CreateORMSPEC_RELATION(SPEC_RELATION *SPEC_RELATION)
	CreateORMSPEC_RELATION_TYPE(SPEC_RELATION_TYPE *SPEC_RELATION_TYPE)
	CreateORMXHTML_CONTENT(XHTML_CONTENT *XHTML_CONTENT)
	CreateORMXhtml_InlPres_type(Xhtml_InlPres_type *Xhtml_InlPres_type)
	CreateORMXhtml_a_type(Xhtml_a_type *Xhtml_a_type)
	CreateORMXhtml_abbr_type(Xhtml_abbr_type *Xhtml_abbr_type)
	CreateORMXhtml_acronym_type(Xhtml_acronym_type *Xhtml_acronym_type)
	CreateORMXhtml_address_type(Xhtml_address_type *Xhtml_address_type)
	CreateORMXhtml_blockquote_type(Xhtml_blockquote_type *Xhtml_blockquote_type)
	CreateORMXhtml_br_type(Xhtml_br_type *Xhtml_br_type)
	CreateORMXhtml_caption_type(Xhtml_caption_type *Xhtml_caption_type)
	CreateORMXhtml_cite_type(Xhtml_cite_type *Xhtml_cite_type)
	CreateORMXhtml_code_type(Xhtml_code_type *Xhtml_code_type)
	CreateORMXhtml_col_type(Xhtml_col_type *Xhtml_col_type)
	CreateORMXhtml_colgroup_type(Xhtml_colgroup_type *Xhtml_colgroup_type)
	CreateORMXhtml_dd_type(Xhtml_dd_type *Xhtml_dd_type)
	CreateORMXhtml_dfn_type(Xhtml_dfn_type *Xhtml_dfn_type)
	CreateORMXhtml_div_type(Xhtml_div_type *Xhtml_div_type)
	CreateORMXhtml_dl_type(Xhtml_dl_type *Xhtml_dl_type)
	CreateORMXhtml_dt_type(Xhtml_dt_type *Xhtml_dt_type)
	CreateORMXhtml_edit_type(Xhtml_edit_type *Xhtml_edit_type)
	CreateORMXhtml_em_type(Xhtml_em_type *Xhtml_em_type)
	CreateORMXhtml_h1_type(Xhtml_h1_type *Xhtml_h1_type)
	CreateORMXhtml_h2_type(Xhtml_h2_type *Xhtml_h2_type)
	CreateORMXhtml_h3_type(Xhtml_h3_type *Xhtml_h3_type)
	CreateORMXhtml_h4_type(Xhtml_h4_type *Xhtml_h4_type)
	CreateORMXhtml_h5_type(Xhtml_h5_type *Xhtml_h5_type)
	CreateORMXhtml_h6_type(Xhtml_h6_type *Xhtml_h6_type)
	CreateORMXhtml_heading_type(Xhtml_heading_type *Xhtml_heading_type)
	CreateORMXhtml_hr_type(Xhtml_hr_type *Xhtml_hr_type)
	CreateORMXhtml_kbd_type(Xhtml_kbd_type *Xhtml_kbd_type)
	CreateORMXhtml_li_type(Xhtml_li_type *Xhtml_li_type)
	CreateORMXhtml_object_type(Xhtml_object_type *Xhtml_object_type)
	CreateORMXhtml_ol_type(Xhtml_ol_type *Xhtml_ol_type)
	CreateORMXhtml_p_type(Xhtml_p_type *Xhtml_p_type)
	CreateORMXhtml_param_type(Xhtml_param_type *Xhtml_param_type)
	CreateORMXhtml_pre_type(Xhtml_pre_type *Xhtml_pre_type)
	CreateORMXhtml_q_type(Xhtml_q_type *Xhtml_q_type)
	CreateORMXhtml_samp_type(Xhtml_samp_type *Xhtml_samp_type)
	CreateORMXhtml_span_type(Xhtml_span_type *Xhtml_span_type)
	CreateORMXhtml_strong_type(Xhtml_strong_type *Xhtml_strong_type)
	CreateORMXhtml_table_type(Xhtml_table_type *Xhtml_table_type)
	CreateORMXhtml_tbody_type(Xhtml_tbody_type *Xhtml_tbody_type)
	CreateORMXhtml_td_type(Xhtml_td_type *Xhtml_td_type)
	CreateORMXhtml_tfoot_type(Xhtml_tfoot_type *Xhtml_tfoot_type)
	CreateORMXhtml_th_type(Xhtml_th_type *Xhtml_th_type)
	CreateORMXhtml_thead_type(Xhtml_thead_type *Xhtml_thead_type)
	CreateORMXhtml_tr_type(Xhtml_tr_type *Xhtml_tr_type)
	CreateORMXhtml_ul_type(Xhtml_ul_type *Xhtml_ul_type)
	CreateORMXhtml_var_type(Xhtml_var_type *Xhtml_var_type)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMALTERNATIVE_ID(ALTERNATIVE_ID *ALTERNATIVE_ID)
	DeleteORMATTRIBUTE_DEFINITION_BOOLEAN(ATTRIBUTE_DEFINITION_BOOLEAN *ATTRIBUTE_DEFINITION_BOOLEAN)
	DeleteORMATTRIBUTE_DEFINITION_DATE(ATTRIBUTE_DEFINITION_DATE *ATTRIBUTE_DEFINITION_DATE)
	DeleteORMATTRIBUTE_DEFINITION_ENUMERATION(ATTRIBUTE_DEFINITION_ENUMERATION *ATTRIBUTE_DEFINITION_ENUMERATION)
	DeleteORMATTRIBUTE_DEFINITION_INTEGER(ATTRIBUTE_DEFINITION_INTEGER *ATTRIBUTE_DEFINITION_INTEGER)
	DeleteORMATTRIBUTE_DEFINITION_REAL(ATTRIBUTE_DEFINITION_REAL *ATTRIBUTE_DEFINITION_REAL)
	DeleteORMATTRIBUTE_DEFINITION_STRING(ATTRIBUTE_DEFINITION_STRING *ATTRIBUTE_DEFINITION_STRING)
	DeleteORMATTRIBUTE_DEFINITION_XHTML(ATTRIBUTE_DEFINITION_XHTML *ATTRIBUTE_DEFINITION_XHTML)
	DeleteORMATTRIBUTE_VALUE_BOOLEAN(ATTRIBUTE_VALUE_BOOLEAN *ATTRIBUTE_VALUE_BOOLEAN)
	DeleteORMATTRIBUTE_VALUE_DATE(ATTRIBUTE_VALUE_DATE *ATTRIBUTE_VALUE_DATE)
	DeleteORMATTRIBUTE_VALUE_ENUMERATION(ATTRIBUTE_VALUE_ENUMERATION *ATTRIBUTE_VALUE_ENUMERATION)
	DeleteORMATTRIBUTE_VALUE_INTEGER(ATTRIBUTE_VALUE_INTEGER *ATTRIBUTE_VALUE_INTEGER)
	DeleteORMATTRIBUTE_VALUE_REAL(ATTRIBUTE_VALUE_REAL *ATTRIBUTE_VALUE_REAL)
	DeleteORMATTRIBUTE_VALUE_STRING(ATTRIBUTE_VALUE_STRING *ATTRIBUTE_VALUE_STRING)
	DeleteORMATTRIBUTE_VALUE_XHTML(ATTRIBUTE_VALUE_XHTML *ATTRIBUTE_VALUE_XHTML)
	DeleteORMAnyType(AnyType *AnyType)
	DeleteORMDATATYPE_DEFINITION_BOOLEAN(DATATYPE_DEFINITION_BOOLEAN *DATATYPE_DEFINITION_BOOLEAN)
	DeleteORMDATATYPE_DEFINITION_DATE(DATATYPE_DEFINITION_DATE *DATATYPE_DEFINITION_DATE)
	DeleteORMDATATYPE_DEFINITION_ENUMERATION(DATATYPE_DEFINITION_ENUMERATION *DATATYPE_DEFINITION_ENUMERATION)
	DeleteORMDATATYPE_DEFINITION_INTEGER(DATATYPE_DEFINITION_INTEGER *DATATYPE_DEFINITION_INTEGER)
	DeleteORMDATATYPE_DEFINITION_REAL(DATATYPE_DEFINITION_REAL *DATATYPE_DEFINITION_REAL)
	DeleteORMDATATYPE_DEFINITION_STRING(DATATYPE_DEFINITION_STRING *DATATYPE_DEFINITION_STRING)
	DeleteORMDATATYPE_DEFINITION_XHTML(DATATYPE_DEFINITION_XHTML *DATATYPE_DEFINITION_XHTML)
	DeleteORMEMBEDDED_VALUE(EMBEDDED_VALUE *EMBEDDED_VALUE)
	DeleteORMENUM_VALUE(ENUM_VALUE *ENUM_VALUE)
	DeleteORMRELATION_GROUP(RELATION_GROUP *RELATION_GROUP)
	DeleteORMRELATION_GROUP_TYPE(RELATION_GROUP_TYPE *RELATION_GROUP_TYPE)
	DeleteORMREQ_IF(REQ_IF *REQ_IF)
	DeleteORMREQ_IF_CONTENT(REQ_IF_CONTENT *REQ_IF_CONTENT)
	DeleteORMREQ_IF_HEADER(REQ_IF_HEADER *REQ_IF_HEADER)
	DeleteORMREQ_IF_TOOL_EXTENSION(REQ_IF_TOOL_EXTENSION *REQ_IF_TOOL_EXTENSION)
	DeleteORMSPECIFICATION(SPECIFICATION *SPECIFICATION)
	DeleteORMSPECIFICATION_TYPE(SPECIFICATION_TYPE *SPECIFICATION_TYPE)
	DeleteORMSPEC_HIERARCHY(SPEC_HIERARCHY *SPEC_HIERARCHY)
	DeleteORMSPEC_OBJECT(SPEC_OBJECT *SPEC_OBJECT)
	DeleteORMSPEC_OBJECT_TYPE(SPEC_OBJECT_TYPE *SPEC_OBJECT_TYPE)
	DeleteORMSPEC_RELATION(SPEC_RELATION *SPEC_RELATION)
	DeleteORMSPEC_RELATION_TYPE(SPEC_RELATION_TYPE *SPEC_RELATION_TYPE)
	DeleteORMXHTML_CONTENT(XHTML_CONTENT *XHTML_CONTENT)
	DeleteORMXhtml_InlPres_type(Xhtml_InlPres_type *Xhtml_InlPres_type)
	DeleteORMXhtml_a_type(Xhtml_a_type *Xhtml_a_type)
	DeleteORMXhtml_abbr_type(Xhtml_abbr_type *Xhtml_abbr_type)
	DeleteORMXhtml_acronym_type(Xhtml_acronym_type *Xhtml_acronym_type)
	DeleteORMXhtml_address_type(Xhtml_address_type *Xhtml_address_type)
	DeleteORMXhtml_blockquote_type(Xhtml_blockquote_type *Xhtml_blockquote_type)
	DeleteORMXhtml_br_type(Xhtml_br_type *Xhtml_br_type)
	DeleteORMXhtml_caption_type(Xhtml_caption_type *Xhtml_caption_type)
	DeleteORMXhtml_cite_type(Xhtml_cite_type *Xhtml_cite_type)
	DeleteORMXhtml_code_type(Xhtml_code_type *Xhtml_code_type)
	DeleteORMXhtml_col_type(Xhtml_col_type *Xhtml_col_type)
	DeleteORMXhtml_colgroup_type(Xhtml_colgroup_type *Xhtml_colgroup_type)
	DeleteORMXhtml_dd_type(Xhtml_dd_type *Xhtml_dd_type)
	DeleteORMXhtml_dfn_type(Xhtml_dfn_type *Xhtml_dfn_type)
	DeleteORMXhtml_div_type(Xhtml_div_type *Xhtml_div_type)
	DeleteORMXhtml_dl_type(Xhtml_dl_type *Xhtml_dl_type)
	DeleteORMXhtml_dt_type(Xhtml_dt_type *Xhtml_dt_type)
	DeleteORMXhtml_edit_type(Xhtml_edit_type *Xhtml_edit_type)
	DeleteORMXhtml_em_type(Xhtml_em_type *Xhtml_em_type)
	DeleteORMXhtml_h1_type(Xhtml_h1_type *Xhtml_h1_type)
	DeleteORMXhtml_h2_type(Xhtml_h2_type *Xhtml_h2_type)
	DeleteORMXhtml_h3_type(Xhtml_h3_type *Xhtml_h3_type)
	DeleteORMXhtml_h4_type(Xhtml_h4_type *Xhtml_h4_type)
	DeleteORMXhtml_h5_type(Xhtml_h5_type *Xhtml_h5_type)
	DeleteORMXhtml_h6_type(Xhtml_h6_type *Xhtml_h6_type)
	DeleteORMXhtml_heading_type(Xhtml_heading_type *Xhtml_heading_type)
	DeleteORMXhtml_hr_type(Xhtml_hr_type *Xhtml_hr_type)
	DeleteORMXhtml_kbd_type(Xhtml_kbd_type *Xhtml_kbd_type)
	DeleteORMXhtml_li_type(Xhtml_li_type *Xhtml_li_type)
	DeleteORMXhtml_object_type(Xhtml_object_type *Xhtml_object_type)
	DeleteORMXhtml_ol_type(Xhtml_ol_type *Xhtml_ol_type)
	DeleteORMXhtml_p_type(Xhtml_p_type *Xhtml_p_type)
	DeleteORMXhtml_param_type(Xhtml_param_type *Xhtml_param_type)
	DeleteORMXhtml_pre_type(Xhtml_pre_type *Xhtml_pre_type)
	DeleteORMXhtml_q_type(Xhtml_q_type *Xhtml_q_type)
	DeleteORMXhtml_samp_type(Xhtml_samp_type *Xhtml_samp_type)
	DeleteORMXhtml_span_type(Xhtml_span_type *Xhtml_span_type)
	DeleteORMXhtml_strong_type(Xhtml_strong_type *Xhtml_strong_type)
	DeleteORMXhtml_table_type(Xhtml_table_type *Xhtml_table_type)
	DeleteORMXhtml_tbody_type(Xhtml_tbody_type *Xhtml_tbody_type)
	DeleteORMXhtml_td_type(Xhtml_td_type *Xhtml_td_type)
	DeleteORMXhtml_tfoot_type(Xhtml_tfoot_type *Xhtml_tfoot_type)
	DeleteORMXhtml_th_type(Xhtml_th_type *Xhtml_th_type)
	DeleteORMXhtml_thead_type(Xhtml_thead_type *Xhtml_thead_type)
	DeleteORMXhtml_tr_type(Xhtml_tr_type *Xhtml_tr_type)
	DeleteORMXhtml_ul_type(Xhtml_ul_type *Xhtml_ul_type)
	DeleteORMXhtml_var_type(Xhtml_var_type *Xhtml_var_type)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.ALTERNATIVE_IDs = make(map[*ALTERNATIVE_ID]any)
	stage.ALTERNATIVE_IDs_mapString = make(map[string]*ALTERNATIVE_ID)
	stage.ALTERNATIVE_IDMap_Staged_Order = make(map[*ALTERNATIVE_ID]uint)
	stage.ALTERNATIVE_IDOrder = 0

	stage.ATTRIBUTE_DEFINITION_BOOLEANs = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]any)
	stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_BOOLEAN)
	stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]uint)
	stage.ATTRIBUTE_DEFINITION_BOOLEANOrder = 0

	stage.ATTRIBUTE_DEFINITION_DATEs = make(map[*ATTRIBUTE_DEFINITION_DATE]any)
	stage.ATTRIBUTE_DEFINITION_DATEs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_DATE)
	stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order = make(map[*ATTRIBUTE_DEFINITION_DATE]uint)
	stage.ATTRIBUTE_DEFINITION_DATEOrder = 0

	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]any)
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_ENUMERATION)
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]uint)
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONOrder = 0

	stage.ATTRIBUTE_DEFINITION_INTEGERs = make(map[*ATTRIBUTE_DEFINITION_INTEGER]any)
	stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_INTEGER)
	stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order = make(map[*ATTRIBUTE_DEFINITION_INTEGER]uint)
	stage.ATTRIBUTE_DEFINITION_INTEGEROrder = 0

	stage.ATTRIBUTE_DEFINITION_REALs = make(map[*ATTRIBUTE_DEFINITION_REAL]any)
	stage.ATTRIBUTE_DEFINITION_REALs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_REAL)
	stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order = make(map[*ATTRIBUTE_DEFINITION_REAL]uint)
	stage.ATTRIBUTE_DEFINITION_REALOrder = 0

	stage.ATTRIBUTE_DEFINITION_STRINGs = make(map[*ATTRIBUTE_DEFINITION_STRING]any)
	stage.ATTRIBUTE_DEFINITION_STRINGs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_STRING)
	stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order = make(map[*ATTRIBUTE_DEFINITION_STRING]uint)
	stage.ATTRIBUTE_DEFINITION_STRINGOrder = 0

	stage.ATTRIBUTE_DEFINITION_XHTMLs = make(map[*ATTRIBUTE_DEFINITION_XHTML]any)
	stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_XHTML)
	stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order = make(map[*ATTRIBUTE_DEFINITION_XHTML]uint)
	stage.ATTRIBUTE_DEFINITION_XHTMLOrder = 0

	stage.ATTRIBUTE_VALUE_BOOLEANs = make(map[*ATTRIBUTE_VALUE_BOOLEAN]any)
	stage.ATTRIBUTE_VALUE_BOOLEANs_mapString = make(map[string]*ATTRIBUTE_VALUE_BOOLEAN)
	stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order = make(map[*ATTRIBUTE_VALUE_BOOLEAN]uint)
	stage.ATTRIBUTE_VALUE_BOOLEANOrder = 0

	stage.ATTRIBUTE_VALUE_DATEs = make(map[*ATTRIBUTE_VALUE_DATE]any)
	stage.ATTRIBUTE_VALUE_DATEs_mapString = make(map[string]*ATTRIBUTE_VALUE_DATE)
	stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order = make(map[*ATTRIBUTE_VALUE_DATE]uint)
	stage.ATTRIBUTE_VALUE_DATEOrder = 0

	stage.ATTRIBUTE_VALUE_ENUMERATIONs = make(map[*ATTRIBUTE_VALUE_ENUMERATION]any)
	stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString = make(map[string]*ATTRIBUTE_VALUE_ENUMERATION)
	stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order = make(map[*ATTRIBUTE_VALUE_ENUMERATION]uint)
	stage.ATTRIBUTE_VALUE_ENUMERATIONOrder = 0

	stage.ATTRIBUTE_VALUE_INTEGERs = make(map[*ATTRIBUTE_VALUE_INTEGER]any)
	stage.ATTRIBUTE_VALUE_INTEGERs_mapString = make(map[string]*ATTRIBUTE_VALUE_INTEGER)
	stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order = make(map[*ATTRIBUTE_VALUE_INTEGER]uint)
	stage.ATTRIBUTE_VALUE_INTEGEROrder = 0

	stage.ATTRIBUTE_VALUE_REALs = make(map[*ATTRIBUTE_VALUE_REAL]any)
	stage.ATTRIBUTE_VALUE_REALs_mapString = make(map[string]*ATTRIBUTE_VALUE_REAL)
	stage.ATTRIBUTE_VALUE_REALMap_Staged_Order = make(map[*ATTRIBUTE_VALUE_REAL]uint)
	stage.ATTRIBUTE_VALUE_REALOrder = 0

	stage.ATTRIBUTE_VALUE_STRINGs = make(map[*ATTRIBUTE_VALUE_STRING]any)
	stage.ATTRIBUTE_VALUE_STRINGs_mapString = make(map[string]*ATTRIBUTE_VALUE_STRING)
	stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order = make(map[*ATTRIBUTE_VALUE_STRING]uint)
	stage.ATTRIBUTE_VALUE_STRINGOrder = 0

	stage.ATTRIBUTE_VALUE_XHTMLs = make(map[*ATTRIBUTE_VALUE_XHTML]any)
	stage.ATTRIBUTE_VALUE_XHTMLs_mapString = make(map[string]*ATTRIBUTE_VALUE_XHTML)
	stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order = make(map[*ATTRIBUTE_VALUE_XHTML]uint)
	stage.ATTRIBUTE_VALUE_XHTMLOrder = 0

	stage.AnyTypes = make(map[*AnyType]any)
	stage.AnyTypes_mapString = make(map[string]*AnyType)
	stage.AnyTypeMap_Staged_Order = make(map[*AnyType]uint)
	stage.AnyTypeOrder = 0

	stage.DATATYPE_DEFINITION_BOOLEANs = make(map[*DATATYPE_DEFINITION_BOOLEAN]any)
	stage.DATATYPE_DEFINITION_BOOLEANs_mapString = make(map[string]*DATATYPE_DEFINITION_BOOLEAN)
	stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order = make(map[*DATATYPE_DEFINITION_BOOLEAN]uint)
	stage.DATATYPE_DEFINITION_BOOLEANOrder = 0

	stage.DATATYPE_DEFINITION_DATEs = make(map[*DATATYPE_DEFINITION_DATE]any)
	stage.DATATYPE_DEFINITION_DATEs_mapString = make(map[string]*DATATYPE_DEFINITION_DATE)
	stage.DATATYPE_DEFINITION_DATEMap_Staged_Order = make(map[*DATATYPE_DEFINITION_DATE]uint)
	stage.DATATYPE_DEFINITION_DATEOrder = 0

	stage.DATATYPE_DEFINITION_ENUMERATIONs = make(map[*DATATYPE_DEFINITION_ENUMERATION]any)
	stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString = make(map[string]*DATATYPE_DEFINITION_ENUMERATION)
	stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order = make(map[*DATATYPE_DEFINITION_ENUMERATION]uint)
	stage.DATATYPE_DEFINITION_ENUMERATIONOrder = 0

	stage.DATATYPE_DEFINITION_INTEGERs = make(map[*DATATYPE_DEFINITION_INTEGER]any)
	stage.DATATYPE_DEFINITION_INTEGERs_mapString = make(map[string]*DATATYPE_DEFINITION_INTEGER)
	stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order = make(map[*DATATYPE_DEFINITION_INTEGER]uint)
	stage.DATATYPE_DEFINITION_INTEGEROrder = 0

	stage.DATATYPE_DEFINITION_REALs = make(map[*DATATYPE_DEFINITION_REAL]any)
	stage.DATATYPE_DEFINITION_REALs_mapString = make(map[string]*DATATYPE_DEFINITION_REAL)
	stage.DATATYPE_DEFINITION_REALMap_Staged_Order = make(map[*DATATYPE_DEFINITION_REAL]uint)
	stage.DATATYPE_DEFINITION_REALOrder = 0

	stage.DATATYPE_DEFINITION_STRINGs = make(map[*DATATYPE_DEFINITION_STRING]any)
	stage.DATATYPE_DEFINITION_STRINGs_mapString = make(map[string]*DATATYPE_DEFINITION_STRING)
	stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order = make(map[*DATATYPE_DEFINITION_STRING]uint)
	stage.DATATYPE_DEFINITION_STRINGOrder = 0

	stage.DATATYPE_DEFINITION_XHTMLs = make(map[*DATATYPE_DEFINITION_XHTML]any)
	stage.DATATYPE_DEFINITION_XHTMLs_mapString = make(map[string]*DATATYPE_DEFINITION_XHTML)
	stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order = make(map[*DATATYPE_DEFINITION_XHTML]uint)
	stage.DATATYPE_DEFINITION_XHTMLOrder = 0

	stage.EMBEDDED_VALUEs = make(map[*EMBEDDED_VALUE]any)
	stage.EMBEDDED_VALUEs_mapString = make(map[string]*EMBEDDED_VALUE)
	stage.EMBEDDED_VALUEMap_Staged_Order = make(map[*EMBEDDED_VALUE]uint)
	stage.EMBEDDED_VALUEOrder = 0

	stage.ENUM_VALUEs = make(map[*ENUM_VALUE]any)
	stage.ENUM_VALUEs_mapString = make(map[string]*ENUM_VALUE)
	stage.ENUM_VALUEMap_Staged_Order = make(map[*ENUM_VALUE]uint)
	stage.ENUM_VALUEOrder = 0

	stage.RELATION_GROUPs = make(map[*RELATION_GROUP]any)
	stage.RELATION_GROUPs_mapString = make(map[string]*RELATION_GROUP)
	stage.RELATION_GROUPMap_Staged_Order = make(map[*RELATION_GROUP]uint)
	stage.RELATION_GROUPOrder = 0

	stage.RELATION_GROUP_TYPEs = make(map[*RELATION_GROUP_TYPE]any)
	stage.RELATION_GROUP_TYPEs_mapString = make(map[string]*RELATION_GROUP_TYPE)
	stage.RELATION_GROUP_TYPEMap_Staged_Order = make(map[*RELATION_GROUP_TYPE]uint)
	stage.RELATION_GROUP_TYPEOrder = 0

	stage.REQ_IFs = make(map[*REQ_IF]any)
	stage.REQ_IFs_mapString = make(map[string]*REQ_IF)
	stage.REQ_IFMap_Staged_Order = make(map[*REQ_IF]uint)
	stage.REQ_IFOrder = 0

	stage.REQ_IF_CONTENTs = make(map[*REQ_IF_CONTENT]any)
	stage.REQ_IF_CONTENTs_mapString = make(map[string]*REQ_IF_CONTENT)
	stage.REQ_IF_CONTENTMap_Staged_Order = make(map[*REQ_IF_CONTENT]uint)
	stage.REQ_IF_CONTENTOrder = 0

	stage.REQ_IF_HEADERs = make(map[*REQ_IF_HEADER]any)
	stage.REQ_IF_HEADERs_mapString = make(map[string]*REQ_IF_HEADER)
	stage.REQ_IF_HEADERMap_Staged_Order = make(map[*REQ_IF_HEADER]uint)
	stage.REQ_IF_HEADEROrder = 0

	stage.REQ_IF_TOOL_EXTENSIONs = make(map[*REQ_IF_TOOL_EXTENSION]any)
	stage.REQ_IF_TOOL_EXTENSIONs_mapString = make(map[string]*REQ_IF_TOOL_EXTENSION)
	stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order = make(map[*REQ_IF_TOOL_EXTENSION]uint)
	stage.REQ_IF_TOOL_EXTENSIONOrder = 0

	stage.SPECIFICATIONs = make(map[*SPECIFICATION]any)
	stage.SPECIFICATIONs_mapString = make(map[string]*SPECIFICATION)
	stage.SPECIFICATIONMap_Staged_Order = make(map[*SPECIFICATION]uint)
	stage.SPECIFICATIONOrder = 0

	stage.SPECIFICATION_TYPEs = make(map[*SPECIFICATION_TYPE]any)
	stage.SPECIFICATION_TYPEs_mapString = make(map[string]*SPECIFICATION_TYPE)
	stage.SPECIFICATION_TYPEMap_Staged_Order = make(map[*SPECIFICATION_TYPE]uint)
	stage.SPECIFICATION_TYPEOrder = 0

	stage.SPEC_HIERARCHYs = make(map[*SPEC_HIERARCHY]any)
	stage.SPEC_HIERARCHYs_mapString = make(map[string]*SPEC_HIERARCHY)
	stage.SPEC_HIERARCHYMap_Staged_Order = make(map[*SPEC_HIERARCHY]uint)
	stage.SPEC_HIERARCHYOrder = 0

	stage.SPEC_OBJECTs = make(map[*SPEC_OBJECT]any)
	stage.SPEC_OBJECTs_mapString = make(map[string]*SPEC_OBJECT)
	stage.SPEC_OBJECTMap_Staged_Order = make(map[*SPEC_OBJECT]uint)
	stage.SPEC_OBJECTOrder = 0

	stage.SPEC_OBJECT_TYPEs = make(map[*SPEC_OBJECT_TYPE]any)
	stage.SPEC_OBJECT_TYPEs_mapString = make(map[string]*SPEC_OBJECT_TYPE)
	stage.SPEC_OBJECT_TYPEMap_Staged_Order = make(map[*SPEC_OBJECT_TYPE]uint)
	stage.SPEC_OBJECT_TYPEOrder = 0

	stage.SPEC_RELATIONs = make(map[*SPEC_RELATION]any)
	stage.SPEC_RELATIONs_mapString = make(map[string]*SPEC_RELATION)
	stage.SPEC_RELATIONMap_Staged_Order = make(map[*SPEC_RELATION]uint)
	stage.SPEC_RELATIONOrder = 0

	stage.SPEC_RELATION_TYPEs = make(map[*SPEC_RELATION_TYPE]any)
	stage.SPEC_RELATION_TYPEs_mapString = make(map[string]*SPEC_RELATION_TYPE)
	stage.SPEC_RELATION_TYPEMap_Staged_Order = make(map[*SPEC_RELATION_TYPE]uint)
	stage.SPEC_RELATION_TYPEOrder = 0

	stage.XHTML_CONTENTs = make(map[*XHTML_CONTENT]any)
	stage.XHTML_CONTENTs_mapString = make(map[string]*XHTML_CONTENT)
	stage.XHTML_CONTENTMap_Staged_Order = make(map[*XHTML_CONTENT]uint)
	stage.XHTML_CONTENTOrder = 0

	stage.Xhtml_InlPres_types = make(map[*Xhtml_InlPres_type]any)
	stage.Xhtml_InlPres_types_mapString = make(map[string]*Xhtml_InlPres_type)
	stage.Xhtml_InlPres_typeMap_Staged_Order = make(map[*Xhtml_InlPres_type]uint)
	stage.Xhtml_InlPres_typeOrder = 0

	stage.Xhtml_a_types = make(map[*Xhtml_a_type]any)
	stage.Xhtml_a_types_mapString = make(map[string]*Xhtml_a_type)
	stage.Xhtml_a_typeMap_Staged_Order = make(map[*Xhtml_a_type]uint)
	stage.Xhtml_a_typeOrder = 0

	stage.Xhtml_abbr_types = make(map[*Xhtml_abbr_type]any)
	stage.Xhtml_abbr_types_mapString = make(map[string]*Xhtml_abbr_type)
	stage.Xhtml_abbr_typeMap_Staged_Order = make(map[*Xhtml_abbr_type]uint)
	stage.Xhtml_abbr_typeOrder = 0

	stage.Xhtml_acronym_types = make(map[*Xhtml_acronym_type]any)
	stage.Xhtml_acronym_types_mapString = make(map[string]*Xhtml_acronym_type)
	stage.Xhtml_acronym_typeMap_Staged_Order = make(map[*Xhtml_acronym_type]uint)
	stage.Xhtml_acronym_typeOrder = 0

	stage.Xhtml_address_types = make(map[*Xhtml_address_type]any)
	stage.Xhtml_address_types_mapString = make(map[string]*Xhtml_address_type)
	stage.Xhtml_address_typeMap_Staged_Order = make(map[*Xhtml_address_type]uint)
	stage.Xhtml_address_typeOrder = 0

	stage.Xhtml_blockquote_types = make(map[*Xhtml_blockquote_type]any)
	stage.Xhtml_blockquote_types_mapString = make(map[string]*Xhtml_blockquote_type)
	stage.Xhtml_blockquote_typeMap_Staged_Order = make(map[*Xhtml_blockquote_type]uint)
	stage.Xhtml_blockquote_typeOrder = 0

	stage.Xhtml_br_types = make(map[*Xhtml_br_type]any)
	stage.Xhtml_br_types_mapString = make(map[string]*Xhtml_br_type)
	stage.Xhtml_br_typeMap_Staged_Order = make(map[*Xhtml_br_type]uint)
	stage.Xhtml_br_typeOrder = 0

	stage.Xhtml_caption_types = make(map[*Xhtml_caption_type]any)
	stage.Xhtml_caption_types_mapString = make(map[string]*Xhtml_caption_type)
	stage.Xhtml_caption_typeMap_Staged_Order = make(map[*Xhtml_caption_type]uint)
	stage.Xhtml_caption_typeOrder = 0

	stage.Xhtml_cite_types = make(map[*Xhtml_cite_type]any)
	stage.Xhtml_cite_types_mapString = make(map[string]*Xhtml_cite_type)
	stage.Xhtml_cite_typeMap_Staged_Order = make(map[*Xhtml_cite_type]uint)
	stage.Xhtml_cite_typeOrder = 0

	stage.Xhtml_code_types = make(map[*Xhtml_code_type]any)
	stage.Xhtml_code_types_mapString = make(map[string]*Xhtml_code_type)
	stage.Xhtml_code_typeMap_Staged_Order = make(map[*Xhtml_code_type]uint)
	stage.Xhtml_code_typeOrder = 0

	stage.Xhtml_col_types = make(map[*Xhtml_col_type]any)
	stage.Xhtml_col_types_mapString = make(map[string]*Xhtml_col_type)
	stage.Xhtml_col_typeMap_Staged_Order = make(map[*Xhtml_col_type]uint)
	stage.Xhtml_col_typeOrder = 0

	stage.Xhtml_colgroup_types = make(map[*Xhtml_colgroup_type]any)
	stage.Xhtml_colgroup_types_mapString = make(map[string]*Xhtml_colgroup_type)
	stage.Xhtml_colgroup_typeMap_Staged_Order = make(map[*Xhtml_colgroup_type]uint)
	stage.Xhtml_colgroup_typeOrder = 0

	stage.Xhtml_dd_types = make(map[*Xhtml_dd_type]any)
	stage.Xhtml_dd_types_mapString = make(map[string]*Xhtml_dd_type)
	stage.Xhtml_dd_typeMap_Staged_Order = make(map[*Xhtml_dd_type]uint)
	stage.Xhtml_dd_typeOrder = 0

	stage.Xhtml_dfn_types = make(map[*Xhtml_dfn_type]any)
	stage.Xhtml_dfn_types_mapString = make(map[string]*Xhtml_dfn_type)
	stage.Xhtml_dfn_typeMap_Staged_Order = make(map[*Xhtml_dfn_type]uint)
	stage.Xhtml_dfn_typeOrder = 0

	stage.Xhtml_div_types = make(map[*Xhtml_div_type]any)
	stage.Xhtml_div_types_mapString = make(map[string]*Xhtml_div_type)
	stage.Xhtml_div_typeMap_Staged_Order = make(map[*Xhtml_div_type]uint)
	stage.Xhtml_div_typeOrder = 0

	stage.Xhtml_dl_types = make(map[*Xhtml_dl_type]any)
	stage.Xhtml_dl_types_mapString = make(map[string]*Xhtml_dl_type)
	stage.Xhtml_dl_typeMap_Staged_Order = make(map[*Xhtml_dl_type]uint)
	stage.Xhtml_dl_typeOrder = 0

	stage.Xhtml_dt_types = make(map[*Xhtml_dt_type]any)
	stage.Xhtml_dt_types_mapString = make(map[string]*Xhtml_dt_type)
	stage.Xhtml_dt_typeMap_Staged_Order = make(map[*Xhtml_dt_type]uint)
	stage.Xhtml_dt_typeOrder = 0

	stage.Xhtml_edit_types = make(map[*Xhtml_edit_type]any)
	stage.Xhtml_edit_types_mapString = make(map[string]*Xhtml_edit_type)
	stage.Xhtml_edit_typeMap_Staged_Order = make(map[*Xhtml_edit_type]uint)
	stage.Xhtml_edit_typeOrder = 0

	stage.Xhtml_em_types = make(map[*Xhtml_em_type]any)
	stage.Xhtml_em_types_mapString = make(map[string]*Xhtml_em_type)
	stage.Xhtml_em_typeMap_Staged_Order = make(map[*Xhtml_em_type]uint)
	stage.Xhtml_em_typeOrder = 0

	stage.Xhtml_h1_types = make(map[*Xhtml_h1_type]any)
	stage.Xhtml_h1_types_mapString = make(map[string]*Xhtml_h1_type)
	stage.Xhtml_h1_typeMap_Staged_Order = make(map[*Xhtml_h1_type]uint)
	stage.Xhtml_h1_typeOrder = 0

	stage.Xhtml_h2_types = make(map[*Xhtml_h2_type]any)
	stage.Xhtml_h2_types_mapString = make(map[string]*Xhtml_h2_type)
	stage.Xhtml_h2_typeMap_Staged_Order = make(map[*Xhtml_h2_type]uint)
	stage.Xhtml_h2_typeOrder = 0

	stage.Xhtml_h3_types = make(map[*Xhtml_h3_type]any)
	stage.Xhtml_h3_types_mapString = make(map[string]*Xhtml_h3_type)
	stage.Xhtml_h3_typeMap_Staged_Order = make(map[*Xhtml_h3_type]uint)
	stage.Xhtml_h3_typeOrder = 0

	stage.Xhtml_h4_types = make(map[*Xhtml_h4_type]any)
	stage.Xhtml_h4_types_mapString = make(map[string]*Xhtml_h4_type)
	stage.Xhtml_h4_typeMap_Staged_Order = make(map[*Xhtml_h4_type]uint)
	stage.Xhtml_h4_typeOrder = 0

	stage.Xhtml_h5_types = make(map[*Xhtml_h5_type]any)
	stage.Xhtml_h5_types_mapString = make(map[string]*Xhtml_h5_type)
	stage.Xhtml_h5_typeMap_Staged_Order = make(map[*Xhtml_h5_type]uint)
	stage.Xhtml_h5_typeOrder = 0

	stage.Xhtml_h6_types = make(map[*Xhtml_h6_type]any)
	stage.Xhtml_h6_types_mapString = make(map[string]*Xhtml_h6_type)
	stage.Xhtml_h6_typeMap_Staged_Order = make(map[*Xhtml_h6_type]uint)
	stage.Xhtml_h6_typeOrder = 0

	stage.Xhtml_heading_types = make(map[*Xhtml_heading_type]any)
	stage.Xhtml_heading_types_mapString = make(map[string]*Xhtml_heading_type)
	stage.Xhtml_heading_typeMap_Staged_Order = make(map[*Xhtml_heading_type]uint)
	stage.Xhtml_heading_typeOrder = 0

	stage.Xhtml_hr_types = make(map[*Xhtml_hr_type]any)
	stage.Xhtml_hr_types_mapString = make(map[string]*Xhtml_hr_type)
	stage.Xhtml_hr_typeMap_Staged_Order = make(map[*Xhtml_hr_type]uint)
	stage.Xhtml_hr_typeOrder = 0

	stage.Xhtml_kbd_types = make(map[*Xhtml_kbd_type]any)
	stage.Xhtml_kbd_types_mapString = make(map[string]*Xhtml_kbd_type)
	stage.Xhtml_kbd_typeMap_Staged_Order = make(map[*Xhtml_kbd_type]uint)
	stage.Xhtml_kbd_typeOrder = 0

	stage.Xhtml_li_types = make(map[*Xhtml_li_type]any)
	stage.Xhtml_li_types_mapString = make(map[string]*Xhtml_li_type)
	stage.Xhtml_li_typeMap_Staged_Order = make(map[*Xhtml_li_type]uint)
	stage.Xhtml_li_typeOrder = 0

	stage.Xhtml_object_types = make(map[*Xhtml_object_type]any)
	stage.Xhtml_object_types_mapString = make(map[string]*Xhtml_object_type)
	stage.Xhtml_object_typeMap_Staged_Order = make(map[*Xhtml_object_type]uint)
	stage.Xhtml_object_typeOrder = 0

	stage.Xhtml_ol_types = make(map[*Xhtml_ol_type]any)
	stage.Xhtml_ol_types_mapString = make(map[string]*Xhtml_ol_type)
	stage.Xhtml_ol_typeMap_Staged_Order = make(map[*Xhtml_ol_type]uint)
	stage.Xhtml_ol_typeOrder = 0

	stage.Xhtml_p_types = make(map[*Xhtml_p_type]any)
	stage.Xhtml_p_types_mapString = make(map[string]*Xhtml_p_type)
	stage.Xhtml_p_typeMap_Staged_Order = make(map[*Xhtml_p_type]uint)
	stage.Xhtml_p_typeOrder = 0

	stage.Xhtml_param_types = make(map[*Xhtml_param_type]any)
	stage.Xhtml_param_types_mapString = make(map[string]*Xhtml_param_type)
	stage.Xhtml_param_typeMap_Staged_Order = make(map[*Xhtml_param_type]uint)
	stage.Xhtml_param_typeOrder = 0

	stage.Xhtml_pre_types = make(map[*Xhtml_pre_type]any)
	stage.Xhtml_pre_types_mapString = make(map[string]*Xhtml_pre_type)
	stage.Xhtml_pre_typeMap_Staged_Order = make(map[*Xhtml_pre_type]uint)
	stage.Xhtml_pre_typeOrder = 0

	stage.Xhtml_q_types = make(map[*Xhtml_q_type]any)
	stage.Xhtml_q_types_mapString = make(map[string]*Xhtml_q_type)
	stage.Xhtml_q_typeMap_Staged_Order = make(map[*Xhtml_q_type]uint)
	stage.Xhtml_q_typeOrder = 0

	stage.Xhtml_samp_types = make(map[*Xhtml_samp_type]any)
	stage.Xhtml_samp_types_mapString = make(map[string]*Xhtml_samp_type)
	stage.Xhtml_samp_typeMap_Staged_Order = make(map[*Xhtml_samp_type]uint)
	stage.Xhtml_samp_typeOrder = 0

	stage.Xhtml_span_types = make(map[*Xhtml_span_type]any)
	stage.Xhtml_span_types_mapString = make(map[string]*Xhtml_span_type)
	stage.Xhtml_span_typeMap_Staged_Order = make(map[*Xhtml_span_type]uint)
	stage.Xhtml_span_typeOrder = 0

	stage.Xhtml_strong_types = make(map[*Xhtml_strong_type]any)
	stage.Xhtml_strong_types_mapString = make(map[string]*Xhtml_strong_type)
	stage.Xhtml_strong_typeMap_Staged_Order = make(map[*Xhtml_strong_type]uint)
	stage.Xhtml_strong_typeOrder = 0

	stage.Xhtml_table_types = make(map[*Xhtml_table_type]any)
	stage.Xhtml_table_types_mapString = make(map[string]*Xhtml_table_type)
	stage.Xhtml_table_typeMap_Staged_Order = make(map[*Xhtml_table_type]uint)
	stage.Xhtml_table_typeOrder = 0

	stage.Xhtml_tbody_types = make(map[*Xhtml_tbody_type]any)
	stage.Xhtml_tbody_types_mapString = make(map[string]*Xhtml_tbody_type)
	stage.Xhtml_tbody_typeMap_Staged_Order = make(map[*Xhtml_tbody_type]uint)
	stage.Xhtml_tbody_typeOrder = 0

	stage.Xhtml_td_types = make(map[*Xhtml_td_type]any)
	stage.Xhtml_td_types_mapString = make(map[string]*Xhtml_td_type)
	stage.Xhtml_td_typeMap_Staged_Order = make(map[*Xhtml_td_type]uint)
	stage.Xhtml_td_typeOrder = 0

	stage.Xhtml_tfoot_types = make(map[*Xhtml_tfoot_type]any)
	stage.Xhtml_tfoot_types_mapString = make(map[string]*Xhtml_tfoot_type)
	stage.Xhtml_tfoot_typeMap_Staged_Order = make(map[*Xhtml_tfoot_type]uint)
	stage.Xhtml_tfoot_typeOrder = 0

	stage.Xhtml_th_types = make(map[*Xhtml_th_type]any)
	stage.Xhtml_th_types_mapString = make(map[string]*Xhtml_th_type)
	stage.Xhtml_th_typeMap_Staged_Order = make(map[*Xhtml_th_type]uint)
	stage.Xhtml_th_typeOrder = 0

	stage.Xhtml_thead_types = make(map[*Xhtml_thead_type]any)
	stage.Xhtml_thead_types_mapString = make(map[string]*Xhtml_thead_type)
	stage.Xhtml_thead_typeMap_Staged_Order = make(map[*Xhtml_thead_type]uint)
	stage.Xhtml_thead_typeOrder = 0

	stage.Xhtml_tr_types = make(map[*Xhtml_tr_type]any)
	stage.Xhtml_tr_types_mapString = make(map[string]*Xhtml_tr_type)
	stage.Xhtml_tr_typeMap_Staged_Order = make(map[*Xhtml_tr_type]uint)
	stage.Xhtml_tr_typeOrder = 0

	stage.Xhtml_ul_types = make(map[*Xhtml_ul_type]any)
	stage.Xhtml_ul_types_mapString = make(map[string]*Xhtml_ul_type)
	stage.Xhtml_ul_typeMap_Staged_Order = make(map[*Xhtml_ul_type]uint)
	stage.Xhtml_ul_typeOrder = 0

	stage.Xhtml_var_types = make(map[*Xhtml_var_type]any)
	stage.Xhtml_var_types_mapString = make(map[string]*Xhtml_var_type)
	stage.Xhtml_var_typeMap_Staged_Order = make(map[*Xhtml_var_type]uint)
	stage.Xhtml_var_typeOrder = 0

}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.ALTERNATIVE_IDs = nil
	stage.ALTERNATIVE_IDs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_BOOLEANs = nil
	stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_DATEs = nil
	stage.ATTRIBUTE_DEFINITION_DATEs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs = nil
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_INTEGERs = nil
	stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_REALs = nil
	stage.ATTRIBUTE_DEFINITION_REALs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_STRINGs = nil
	stage.ATTRIBUTE_DEFINITION_STRINGs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_XHTMLs = nil
	stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString = nil

	stage.ATTRIBUTE_VALUE_BOOLEANs = nil
	stage.ATTRIBUTE_VALUE_BOOLEANs_mapString = nil

	stage.ATTRIBUTE_VALUE_DATEs = nil
	stage.ATTRIBUTE_VALUE_DATEs_mapString = nil

	stage.ATTRIBUTE_VALUE_ENUMERATIONs = nil
	stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString = nil

	stage.ATTRIBUTE_VALUE_INTEGERs = nil
	stage.ATTRIBUTE_VALUE_INTEGERs_mapString = nil

	stage.ATTRIBUTE_VALUE_REALs = nil
	stage.ATTRIBUTE_VALUE_REALs_mapString = nil

	stage.ATTRIBUTE_VALUE_STRINGs = nil
	stage.ATTRIBUTE_VALUE_STRINGs_mapString = nil

	stage.ATTRIBUTE_VALUE_XHTMLs = nil
	stage.ATTRIBUTE_VALUE_XHTMLs_mapString = nil

	stage.AnyTypes = nil
	stage.AnyTypes_mapString = nil

	stage.DATATYPE_DEFINITION_BOOLEANs = nil
	stage.DATATYPE_DEFINITION_BOOLEANs_mapString = nil

	stage.DATATYPE_DEFINITION_DATEs = nil
	stage.DATATYPE_DEFINITION_DATEs_mapString = nil

	stage.DATATYPE_DEFINITION_ENUMERATIONs = nil
	stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString = nil

	stage.DATATYPE_DEFINITION_INTEGERs = nil
	stage.DATATYPE_DEFINITION_INTEGERs_mapString = nil

	stage.DATATYPE_DEFINITION_REALs = nil
	stage.DATATYPE_DEFINITION_REALs_mapString = nil

	stage.DATATYPE_DEFINITION_STRINGs = nil
	stage.DATATYPE_DEFINITION_STRINGs_mapString = nil

	stage.DATATYPE_DEFINITION_XHTMLs = nil
	stage.DATATYPE_DEFINITION_XHTMLs_mapString = nil

	stage.EMBEDDED_VALUEs = nil
	stage.EMBEDDED_VALUEs_mapString = nil

	stage.ENUM_VALUEs = nil
	stage.ENUM_VALUEs_mapString = nil

	stage.RELATION_GROUPs = nil
	stage.RELATION_GROUPs_mapString = nil

	stage.RELATION_GROUP_TYPEs = nil
	stage.RELATION_GROUP_TYPEs_mapString = nil

	stage.REQ_IFs = nil
	stage.REQ_IFs_mapString = nil

	stage.REQ_IF_CONTENTs = nil
	stage.REQ_IF_CONTENTs_mapString = nil

	stage.REQ_IF_HEADERs = nil
	stage.REQ_IF_HEADERs_mapString = nil

	stage.REQ_IF_TOOL_EXTENSIONs = nil
	stage.REQ_IF_TOOL_EXTENSIONs_mapString = nil

	stage.SPECIFICATIONs = nil
	stage.SPECIFICATIONs_mapString = nil

	stage.SPECIFICATION_TYPEs = nil
	stage.SPECIFICATION_TYPEs_mapString = nil

	stage.SPEC_HIERARCHYs = nil
	stage.SPEC_HIERARCHYs_mapString = nil

	stage.SPEC_OBJECTs = nil
	stage.SPEC_OBJECTs_mapString = nil

	stage.SPEC_OBJECT_TYPEs = nil
	stage.SPEC_OBJECT_TYPEs_mapString = nil

	stage.SPEC_RELATIONs = nil
	stage.SPEC_RELATIONs_mapString = nil

	stage.SPEC_RELATION_TYPEs = nil
	stage.SPEC_RELATION_TYPEs_mapString = nil

	stage.XHTML_CONTENTs = nil
	stage.XHTML_CONTENTs_mapString = nil

	stage.Xhtml_InlPres_types = nil
	stage.Xhtml_InlPres_types_mapString = nil

	stage.Xhtml_a_types = nil
	stage.Xhtml_a_types_mapString = nil

	stage.Xhtml_abbr_types = nil
	stage.Xhtml_abbr_types_mapString = nil

	stage.Xhtml_acronym_types = nil
	stage.Xhtml_acronym_types_mapString = nil

	stage.Xhtml_address_types = nil
	stage.Xhtml_address_types_mapString = nil

	stage.Xhtml_blockquote_types = nil
	stage.Xhtml_blockquote_types_mapString = nil

	stage.Xhtml_br_types = nil
	stage.Xhtml_br_types_mapString = nil

	stage.Xhtml_caption_types = nil
	stage.Xhtml_caption_types_mapString = nil

	stage.Xhtml_cite_types = nil
	stage.Xhtml_cite_types_mapString = nil

	stage.Xhtml_code_types = nil
	stage.Xhtml_code_types_mapString = nil

	stage.Xhtml_col_types = nil
	stage.Xhtml_col_types_mapString = nil

	stage.Xhtml_colgroup_types = nil
	stage.Xhtml_colgroup_types_mapString = nil

	stage.Xhtml_dd_types = nil
	stage.Xhtml_dd_types_mapString = nil

	stage.Xhtml_dfn_types = nil
	stage.Xhtml_dfn_types_mapString = nil

	stage.Xhtml_div_types = nil
	stage.Xhtml_div_types_mapString = nil

	stage.Xhtml_dl_types = nil
	stage.Xhtml_dl_types_mapString = nil

	stage.Xhtml_dt_types = nil
	stage.Xhtml_dt_types_mapString = nil

	stage.Xhtml_edit_types = nil
	stage.Xhtml_edit_types_mapString = nil

	stage.Xhtml_em_types = nil
	stage.Xhtml_em_types_mapString = nil

	stage.Xhtml_h1_types = nil
	stage.Xhtml_h1_types_mapString = nil

	stage.Xhtml_h2_types = nil
	stage.Xhtml_h2_types_mapString = nil

	stage.Xhtml_h3_types = nil
	stage.Xhtml_h3_types_mapString = nil

	stage.Xhtml_h4_types = nil
	stage.Xhtml_h4_types_mapString = nil

	stage.Xhtml_h5_types = nil
	stage.Xhtml_h5_types_mapString = nil

	stage.Xhtml_h6_types = nil
	stage.Xhtml_h6_types_mapString = nil

	stage.Xhtml_heading_types = nil
	stage.Xhtml_heading_types_mapString = nil

	stage.Xhtml_hr_types = nil
	stage.Xhtml_hr_types_mapString = nil

	stage.Xhtml_kbd_types = nil
	stage.Xhtml_kbd_types_mapString = nil

	stage.Xhtml_li_types = nil
	stage.Xhtml_li_types_mapString = nil

	stage.Xhtml_object_types = nil
	stage.Xhtml_object_types_mapString = nil

	stage.Xhtml_ol_types = nil
	stage.Xhtml_ol_types_mapString = nil

	stage.Xhtml_p_types = nil
	stage.Xhtml_p_types_mapString = nil

	stage.Xhtml_param_types = nil
	stage.Xhtml_param_types_mapString = nil

	stage.Xhtml_pre_types = nil
	stage.Xhtml_pre_types_mapString = nil

	stage.Xhtml_q_types = nil
	stage.Xhtml_q_types_mapString = nil

	stage.Xhtml_samp_types = nil
	stage.Xhtml_samp_types_mapString = nil

	stage.Xhtml_span_types = nil
	stage.Xhtml_span_types_mapString = nil

	stage.Xhtml_strong_types = nil
	stage.Xhtml_strong_types_mapString = nil

	stage.Xhtml_table_types = nil
	stage.Xhtml_table_types_mapString = nil

	stage.Xhtml_tbody_types = nil
	stage.Xhtml_tbody_types_mapString = nil

	stage.Xhtml_td_types = nil
	stage.Xhtml_td_types_mapString = nil

	stage.Xhtml_tfoot_types = nil
	stage.Xhtml_tfoot_types_mapString = nil

	stage.Xhtml_th_types = nil
	stage.Xhtml_th_types_mapString = nil

	stage.Xhtml_thead_types = nil
	stage.Xhtml_thead_types_mapString = nil

	stage.Xhtml_tr_types = nil
	stage.Xhtml_tr_types_mapString = nil

	stage.Xhtml_ul_types = nil
	stage.Xhtml_ul_types_mapString = nil

	stage.Xhtml_var_types = nil
	stage.Xhtml_var_types_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for alternative_id := range stage.ALTERNATIVE_IDs {
		alternative_id.Unstage(stage)
	}

	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		attribute_definition_boolean.Unstage(stage)
	}

	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		attribute_definition_date.Unstage(stage)
	}

	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		attribute_definition_enumeration.Unstage(stage)
	}

	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		attribute_definition_integer.Unstage(stage)
	}

	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		attribute_definition_real.Unstage(stage)
	}

	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		attribute_definition_string.Unstage(stage)
	}

	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		attribute_definition_xhtml.Unstage(stage)
	}

	for attribute_value_boolean := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		attribute_value_boolean.Unstage(stage)
	}

	for attribute_value_date := range stage.ATTRIBUTE_VALUE_DATEs {
		attribute_value_date.Unstage(stage)
	}

	for attribute_value_enumeration := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		attribute_value_enumeration.Unstage(stage)
	}

	for attribute_value_integer := range stage.ATTRIBUTE_VALUE_INTEGERs {
		attribute_value_integer.Unstage(stage)
	}

	for attribute_value_real := range stage.ATTRIBUTE_VALUE_REALs {
		attribute_value_real.Unstage(stage)
	}

	for attribute_value_string := range stage.ATTRIBUTE_VALUE_STRINGs {
		attribute_value_string.Unstage(stage)
	}

	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		attribute_value_xhtml.Unstage(stage)
	}

	for anytype := range stage.AnyTypes {
		anytype.Unstage(stage)
	}

	for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
		datatype_definition_boolean.Unstage(stage)
	}

	for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
		datatype_definition_date.Unstage(stage)
	}

	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		datatype_definition_enumeration.Unstage(stage)
	}

	for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
		datatype_definition_integer.Unstage(stage)
	}

	for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
		datatype_definition_real.Unstage(stage)
	}

	for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
		datatype_definition_string.Unstage(stage)
	}

	for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
		datatype_definition_xhtml.Unstage(stage)
	}

	for embedded_value := range stage.EMBEDDED_VALUEs {
		embedded_value.Unstage(stage)
	}

	for enum_value := range stage.ENUM_VALUEs {
		enum_value.Unstage(stage)
	}

	for relation_group := range stage.RELATION_GROUPs {
		relation_group.Unstage(stage)
	}

	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		relation_group_type.Unstage(stage)
	}

	for req_if := range stage.REQ_IFs {
		req_if.Unstage(stage)
	}

	for req_if_content := range stage.REQ_IF_CONTENTs {
		req_if_content.Unstage(stage)
	}

	for req_if_header := range stage.REQ_IF_HEADERs {
		req_if_header.Unstage(stage)
	}

	for req_if_tool_extension := range stage.REQ_IF_TOOL_EXTENSIONs {
		req_if_tool_extension.Unstage(stage)
	}

	for specification := range stage.SPECIFICATIONs {
		specification.Unstage(stage)
	}

	for specification_type := range stage.SPECIFICATION_TYPEs {
		specification_type.Unstage(stage)
	}

	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		spec_hierarchy.Unstage(stage)
	}

	for spec_object := range stage.SPEC_OBJECTs {
		spec_object.Unstage(stage)
	}

	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		spec_object_type.Unstage(stage)
	}

	for spec_relation := range stage.SPEC_RELATIONs {
		spec_relation.Unstage(stage)
	}

	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		spec_relation_type.Unstage(stage)
	}

	for xhtml_content := range stage.XHTML_CONTENTs {
		xhtml_content.Unstage(stage)
	}

	for xhtml_inlpres_type := range stage.Xhtml_InlPres_types {
		xhtml_inlpres_type.Unstage(stage)
	}

	for xhtml_a_type := range stage.Xhtml_a_types {
		xhtml_a_type.Unstage(stage)
	}

	for xhtml_abbr_type := range stage.Xhtml_abbr_types {
		xhtml_abbr_type.Unstage(stage)
	}

	for xhtml_acronym_type := range stage.Xhtml_acronym_types {
		xhtml_acronym_type.Unstage(stage)
	}

	for xhtml_address_type := range stage.Xhtml_address_types {
		xhtml_address_type.Unstage(stage)
	}

	for xhtml_blockquote_type := range stage.Xhtml_blockquote_types {
		xhtml_blockquote_type.Unstage(stage)
	}

	for xhtml_br_type := range stage.Xhtml_br_types {
		xhtml_br_type.Unstage(stage)
	}

	for xhtml_caption_type := range stage.Xhtml_caption_types {
		xhtml_caption_type.Unstage(stage)
	}

	for xhtml_cite_type := range stage.Xhtml_cite_types {
		xhtml_cite_type.Unstage(stage)
	}

	for xhtml_code_type := range stage.Xhtml_code_types {
		xhtml_code_type.Unstage(stage)
	}

	for xhtml_col_type := range stage.Xhtml_col_types {
		xhtml_col_type.Unstage(stage)
	}

	for xhtml_colgroup_type := range stage.Xhtml_colgroup_types {
		xhtml_colgroup_type.Unstage(stage)
	}

	for xhtml_dd_type := range stage.Xhtml_dd_types {
		xhtml_dd_type.Unstage(stage)
	}

	for xhtml_dfn_type := range stage.Xhtml_dfn_types {
		xhtml_dfn_type.Unstage(stage)
	}

	for xhtml_div_type := range stage.Xhtml_div_types {
		xhtml_div_type.Unstage(stage)
	}

	for xhtml_dl_type := range stage.Xhtml_dl_types {
		xhtml_dl_type.Unstage(stage)
	}

	for xhtml_dt_type := range stage.Xhtml_dt_types {
		xhtml_dt_type.Unstage(stage)
	}

	for xhtml_edit_type := range stage.Xhtml_edit_types {
		xhtml_edit_type.Unstage(stage)
	}

	for xhtml_em_type := range stage.Xhtml_em_types {
		xhtml_em_type.Unstage(stage)
	}

	for xhtml_h1_type := range stage.Xhtml_h1_types {
		xhtml_h1_type.Unstage(stage)
	}

	for xhtml_h2_type := range stage.Xhtml_h2_types {
		xhtml_h2_type.Unstage(stage)
	}

	for xhtml_h3_type := range stage.Xhtml_h3_types {
		xhtml_h3_type.Unstage(stage)
	}

	for xhtml_h4_type := range stage.Xhtml_h4_types {
		xhtml_h4_type.Unstage(stage)
	}

	for xhtml_h5_type := range stage.Xhtml_h5_types {
		xhtml_h5_type.Unstage(stage)
	}

	for xhtml_h6_type := range stage.Xhtml_h6_types {
		xhtml_h6_type.Unstage(stage)
	}

	for xhtml_heading_type := range stage.Xhtml_heading_types {
		xhtml_heading_type.Unstage(stage)
	}

	for xhtml_hr_type := range stage.Xhtml_hr_types {
		xhtml_hr_type.Unstage(stage)
	}

	for xhtml_kbd_type := range stage.Xhtml_kbd_types {
		xhtml_kbd_type.Unstage(stage)
	}

	for xhtml_li_type := range stage.Xhtml_li_types {
		xhtml_li_type.Unstage(stage)
	}

	for xhtml_object_type := range stage.Xhtml_object_types {
		xhtml_object_type.Unstage(stage)
	}

	for xhtml_ol_type := range stage.Xhtml_ol_types {
		xhtml_ol_type.Unstage(stage)
	}

	for xhtml_p_type := range stage.Xhtml_p_types {
		xhtml_p_type.Unstage(stage)
	}

	for xhtml_param_type := range stage.Xhtml_param_types {
		xhtml_param_type.Unstage(stage)
	}

	for xhtml_pre_type := range stage.Xhtml_pre_types {
		xhtml_pre_type.Unstage(stage)
	}

	for xhtml_q_type := range stage.Xhtml_q_types {
		xhtml_q_type.Unstage(stage)
	}

	for xhtml_samp_type := range stage.Xhtml_samp_types {
		xhtml_samp_type.Unstage(stage)
	}

	for xhtml_span_type := range stage.Xhtml_span_types {
		xhtml_span_type.Unstage(stage)
	}

	for xhtml_strong_type := range stage.Xhtml_strong_types {
		xhtml_strong_type.Unstage(stage)
	}

	for xhtml_table_type := range stage.Xhtml_table_types {
		xhtml_table_type.Unstage(stage)
	}

	for xhtml_tbody_type := range stage.Xhtml_tbody_types {
		xhtml_tbody_type.Unstage(stage)
	}

	for xhtml_td_type := range stage.Xhtml_td_types {
		xhtml_td_type.Unstage(stage)
	}

	for xhtml_tfoot_type := range stage.Xhtml_tfoot_types {
		xhtml_tfoot_type.Unstage(stage)
	}

	for xhtml_th_type := range stage.Xhtml_th_types {
		xhtml_th_type.Unstage(stage)
	}

	for xhtml_thead_type := range stage.Xhtml_thead_types {
		xhtml_thead_type.Unstage(stage)
	}

	for xhtml_tr_type := range stage.Xhtml_tr_types {
		xhtml_tr_type.Unstage(stage)
	}

	for xhtml_ul_type := range stage.Xhtml_ul_types {
		xhtml_ul_type.Unstage(stage)
	}

	for xhtml_var_type := range stage.Xhtml_var_types {
		xhtml_var_type.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*Stage)
	UnstageVoid(stage *Stage)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *Stage) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*ALTERNATIVE_ID]any:
		return any(&stage.ALTERNATIVE_IDs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_BOOLEAN]any:
		return any(&stage.ATTRIBUTE_DEFINITION_BOOLEANs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_DATE]any:
		return any(&stage.ATTRIBUTE_DEFINITION_DATEs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_ENUMERATION]any:
		return any(&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_INTEGER]any:
		return any(&stage.ATTRIBUTE_DEFINITION_INTEGERs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_REAL]any:
		return any(&stage.ATTRIBUTE_DEFINITION_REALs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_STRING]any:
		return any(&stage.ATTRIBUTE_DEFINITION_STRINGs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_XHTML]any:
		return any(&stage.ATTRIBUTE_DEFINITION_XHTMLs).(*Type)
	case map[*ATTRIBUTE_VALUE_BOOLEAN]any:
		return any(&stage.ATTRIBUTE_VALUE_BOOLEANs).(*Type)
	case map[*ATTRIBUTE_VALUE_DATE]any:
		return any(&stage.ATTRIBUTE_VALUE_DATEs).(*Type)
	case map[*ATTRIBUTE_VALUE_ENUMERATION]any:
		return any(&stage.ATTRIBUTE_VALUE_ENUMERATIONs).(*Type)
	case map[*ATTRIBUTE_VALUE_INTEGER]any:
		return any(&stage.ATTRIBUTE_VALUE_INTEGERs).(*Type)
	case map[*ATTRIBUTE_VALUE_REAL]any:
		return any(&stage.ATTRIBUTE_VALUE_REALs).(*Type)
	case map[*ATTRIBUTE_VALUE_STRING]any:
		return any(&stage.ATTRIBUTE_VALUE_STRINGs).(*Type)
	case map[*ATTRIBUTE_VALUE_XHTML]any:
		return any(&stage.ATTRIBUTE_VALUE_XHTMLs).(*Type)
	case map[*AnyType]any:
		return any(&stage.AnyTypes).(*Type)
	case map[*DATATYPE_DEFINITION_BOOLEAN]any:
		return any(&stage.DATATYPE_DEFINITION_BOOLEANs).(*Type)
	case map[*DATATYPE_DEFINITION_DATE]any:
		return any(&stage.DATATYPE_DEFINITION_DATEs).(*Type)
	case map[*DATATYPE_DEFINITION_ENUMERATION]any:
		return any(&stage.DATATYPE_DEFINITION_ENUMERATIONs).(*Type)
	case map[*DATATYPE_DEFINITION_INTEGER]any:
		return any(&stage.DATATYPE_DEFINITION_INTEGERs).(*Type)
	case map[*DATATYPE_DEFINITION_REAL]any:
		return any(&stage.DATATYPE_DEFINITION_REALs).(*Type)
	case map[*DATATYPE_DEFINITION_STRING]any:
		return any(&stage.DATATYPE_DEFINITION_STRINGs).(*Type)
	case map[*DATATYPE_DEFINITION_XHTML]any:
		return any(&stage.DATATYPE_DEFINITION_XHTMLs).(*Type)
	case map[*EMBEDDED_VALUE]any:
		return any(&stage.EMBEDDED_VALUEs).(*Type)
	case map[*ENUM_VALUE]any:
		return any(&stage.ENUM_VALUEs).(*Type)
	case map[*RELATION_GROUP]any:
		return any(&stage.RELATION_GROUPs).(*Type)
	case map[*RELATION_GROUP_TYPE]any:
		return any(&stage.RELATION_GROUP_TYPEs).(*Type)
	case map[*REQ_IF]any:
		return any(&stage.REQ_IFs).(*Type)
	case map[*REQ_IF_CONTENT]any:
		return any(&stage.REQ_IF_CONTENTs).(*Type)
	case map[*REQ_IF_HEADER]any:
		return any(&stage.REQ_IF_HEADERs).(*Type)
	case map[*REQ_IF_TOOL_EXTENSION]any:
		return any(&stage.REQ_IF_TOOL_EXTENSIONs).(*Type)
	case map[*SPECIFICATION]any:
		return any(&stage.SPECIFICATIONs).(*Type)
	case map[*SPECIFICATION_TYPE]any:
		return any(&stage.SPECIFICATION_TYPEs).(*Type)
	case map[*SPEC_HIERARCHY]any:
		return any(&stage.SPEC_HIERARCHYs).(*Type)
	case map[*SPEC_OBJECT]any:
		return any(&stage.SPEC_OBJECTs).(*Type)
	case map[*SPEC_OBJECT_TYPE]any:
		return any(&stage.SPEC_OBJECT_TYPEs).(*Type)
	case map[*SPEC_RELATION]any:
		return any(&stage.SPEC_RELATIONs).(*Type)
	case map[*SPEC_RELATION_TYPE]any:
		return any(&stage.SPEC_RELATION_TYPEs).(*Type)
	case map[*XHTML_CONTENT]any:
		return any(&stage.XHTML_CONTENTs).(*Type)
	case map[*Xhtml_InlPres_type]any:
		return any(&stage.Xhtml_InlPres_types).(*Type)
	case map[*Xhtml_a_type]any:
		return any(&stage.Xhtml_a_types).(*Type)
	case map[*Xhtml_abbr_type]any:
		return any(&stage.Xhtml_abbr_types).(*Type)
	case map[*Xhtml_acronym_type]any:
		return any(&stage.Xhtml_acronym_types).(*Type)
	case map[*Xhtml_address_type]any:
		return any(&stage.Xhtml_address_types).(*Type)
	case map[*Xhtml_blockquote_type]any:
		return any(&stage.Xhtml_blockquote_types).(*Type)
	case map[*Xhtml_br_type]any:
		return any(&stage.Xhtml_br_types).(*Type)
	case map[*Xhtml_caption_type]any:
		return any(&stage.Xhtml_caption_types).(*Type)
	case map[*Xhtml_cite_type]any:
		return any(&stage.Xhtml_cite_types).(*Type)
	case map[*Xhtml_code_type]any:
		return any(&stage.Xhtml_code_types).(*Type)
	case map[*Xhtml_col_type]any:
		return any(&stage.Xhtml_col_types).(*Type)
	case map[*Xhtml_colgroup_type]any:
		return any(&stage.Xhtml_colgroup_types).(*Type)
	case map[*Xhtml_dd_type]any:
		return any(&stage.Xhtml_dd_types).(*Type)
	case map[*Xhtml_dfn_type]any:
		return any(&stage.Xhtml_dfn_types).(*Type)
	case map[*Xhtml_div_type]any:
		return any(&stage.Xhtml_div_types).(*Type)
	case map[*Xhtml_dl_type]any:
		return any(&stage.Xhtml_dl_types).(*Type)
	case map[*Xhtml_dt_type]any:
		return any(&stage.Xhtml_dt_types).(*Type)
	case map[*Xhtml_edit_type]any:
		return any(&stage.Xhtml_edit_types).(*Type)
	case map[*Xhtml_em_type]any:
		return any(&stage.Xhtml_em_types).(*Type)
	case map[*Xhtml_h1_type]any:
		return any(&stage.Xhtml_h1_types).(*Type)
	case map[*Xhtml_h2_type]any:
		return any(&stage.Xhtml_h2_types).(*Type)
	case map[*Xhtml_h3_type]any:
		return any(&stage.Xhtml_h3_types).(*Type)
	case map[*Xhtml_h4_type]any:
		return any(&stage.Xhtml_h4_types).(*Type)
	case map[*Xhtml_h5_type]any:
		return any(&stage.Xhtml_h5_types).(*Type)
	case map[*Xhtml_h6_type]any:
		return any(&stage.Xhtml_h6_types).(*Type)
	case map[*Xhtml_heading_type]any:
		return any(&stage.Xhtml_heading_types).(*Type)
	case map[*Xhtml_hr_type]any:
		return any(&stage.Xhtml_hr_types).(*Type)
	case map[*Xhtml_kbd_type]any:
		return any(&stage.Xhtml_kbd_types).(*Type)
	case map[*Xhtml_li_type]any:
		return any(&stage.Xhtml_li_types).(*Type)
	case map[*Xhtml_object_type]any:
		return any(&stage.Xhtml_object_types).(*Type)
	case map[*Xhtml_ol_type]any:
		return any(&stage.Xhtml_ol_types).(*Type)
	case map[*Xhtml_p_type]any:
		return any(&stage.Xhtml_p_types).(*Type)
	case map[*Xhtml_param_type]any:
		return any(&stage.Xhtml_param_types).(*Type)
	case map[*Xhtml_pre_type]any:
		return any(&stage.Xhtml_pre_types).(*Type)
	case map[*Xhtml_q_type]any:
		return any(&stage.Xhtml_q_types).(*Type)
	case map[*Xhtml_samp_type]any:
		return any(&stage.Xhtml_samp_types).(*Type)
	case map[*Xhtml_span_type]any:
		return any(&stage.Xhtml_span_types).(*Type)
	case map[*Xhtml_strong_type]any:
		return any(&stage.Xhtml_strong_types).(*Type)
	case map[*Xhtml_table_type]any:
		return any(&stage.Xhtml_table_types).(*Type)
	case map[*Xhtml_tbody_type]any:
		return any(&stage.Xhtml_tbody_types).(*Type)
	case map[*Xhtml_td_type]any:
		return any(&stage.Xhtml_td_types).(*Type)
	case map[*Xhtml_tfoot_type]any:
		return any(&stage.Xhtml_tfoot_types).(*Type)
	case map[*Xhtml_th_type]any:
		return any(&stage.Xhtml_th_types).(*Type)
	case map[*Xhtml_thead_type]any:
		return any(&stage.Xhtml_thead_types).(*Type)
	case map[*Xhtml_tr_type]any:
		return any(&stage.Xhtml_tr_types).(*Type)
	case map[*Xhtml_ul_type]any:
		return any(&stage.Xhtml_ul_types).(*Type)
	case map[*Xhtml_var_type]any:
		return any(&stage.Xhtml_var_types).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*ALTERNATIVE_ID:
		return any(&stage.ALTERNATIVE_IDs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_BOOLEAN:
		return any(&stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_DATE:
		return any(&stage.ATTRIBUTE_DEFINITION_DATEs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_INTEGER:
		return any(&stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_REAL:
		return any(&stage.ATTRIBUTE_DEFINITION_REALs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_STRING:
		return any(&stage.ATTRIBUTE_DEFINITION_STRINGs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_XHTML:
		return any(&stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.ATTRIBUTE_VALUE_BOOLEANs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_DATE:
		return any(&stage.ATTRIBUTE_VALUE_DATEs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.ATTRIBUTE_VALUE_INTEGERs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_REAL:
		return any(&stage.ATTRIBUTE_VALUE_REALs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_STRING:
		return any(&stage.ATTRIBUTE_VALUE_STRINGs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_XHTML:
		return any(&stage.ATTRIBUTE_VALUE_XHTMLs_mapString).(*Type)
	case map[string]*AnyType:
		return any(&stage.AnyTypes_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_BOOLEAN:
		return any(&stage.DATATYPE_DEFINITION_BOOLEANs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_DATE:
		return any(&stage.DATATYPE_DEFINITION_DATEs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_ENUMERATION:
		return any(&stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_INTEGER:
		return any(&stage.DATATYPE_DEFINITION_INTEGERs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_REAL:
		return any(&stage.DATATYPE_DEFINITION_REALs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_STRING:
		return any(&stage.DATATYPE_DEFINITION_STRINGs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_XHTML:
		return any(&stage.DATATYPE_DEFINITION_XHTMLs_mapString).(*Type)
	case map[string]*EMBEDDED_VALUE:
		return any(&stage.EMBEDDED_VALUEs_mapString).(*Type)
	case map[string]*ENUM_VALUE:
		return any(&stage.ENUM_VALUEs_mapString).(*Type)
	case map[string]*RELATION_GROUP:
		return any(&stage.RELATION_GROUPs_mapString).(*Type)
	case map[string]*RELATION_GROUP_TYPE:
		return any(&stage.RELATION_GROUP_TYPEs_mapString).(*Type)
	case map[string]*REQ_IF:
		return any(&stage.REQ_IFs_mapString).(*Type)
	case map[string]*REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs_mapString).(*Type)
	case map[string]*REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs_mapString).(*Type)
	case map[string]*REQ_IF_TOOL_EXTENSION:
		return any(&stage.REQ_IF_TOOL_EXTENSIONs_mapString).(*Type)
	case map[string]*SPECIFICATION:
		return any(&stage.SPECIFICATIONs_mapString).(*Type)
	case map[string]*SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs_mapString).(*Type)
	case map[string]*SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs_mapString).(*Type)
	case map[string]*SPEC_OBJECT:
		return any(&stage.SPEC_OBJECTs_mapString).(*Type)
	case map[string]*SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs_mapString).(*Type)
	case map[string]*SPEC_RELATION:
		return any(&stage.SPEC_RELATIONs_mapString).(*Type)
	case map[string]*SPEC_RELATION_TYPE:
		return any(&stage.SPEC_RELATION_TYPEs_mapString).(*Type)
	case map[string]*XHTML_CONTENT:
		return any(&stage.XHTML_CONTENTs_mapString).(*Type)
	case map[string]*Xhtml_InlPres_type:
		return any(&stage.Xhtml_InlPres_types_mapString).(*Type)
	case map[string]*Xhtml_a_type:
		return any(&stage.Xhtml_a_types_mapString).(*Type)
	case map[string]*Xhtml_abbr_type:
		return any(&stage.Xhtml_abbr_types_mapString).(*Type)
	case map[string]*Xhtml_acronym_type:
		return any(&stage.Xhtml_acronym_types_mapString).(*Type)
	case map[string]*Xhtml_address_type:
		return any(&stage.Xhtml_address_types_mapString).(*Type)
	case map[string]*Xhtml_blockquote_type:
		return any(&stage.Xhtml_blockquote_types_mapString).(*Type)
	case map[string]*Xhtml_br_type:
		return any(&stage.Xhtml_br_types_mapString).(*Type)
	case map[string]*Xhtml_caption_type:
		return any(&stage.Xhtml_caption_types_mapString).(*Type)
	case map[string]*Xhtml_cite_type:
		return any(&stage.Xhtml_cite_types_mapString).(*Type)
	case map[string]*Xhtml_code_type:
		return any(&stage.Xhtml_code_types_mapString).(*Type)
	case map[string]*Xhtml_col_type:
		return any(&stage.Xhtml_col_types_mapString).(*Type)
	case map[string]*Xhtml_colgroup_type:
		return any(&stage.Xhtml_colgroup_types_mapString).(*Type)
	case map[string]*Xhtml_dd_type:
		return any(&stage.Xhtml_dd_types_mapString).(*Type)
	case map[string]*Xhtml_dfn_type:
		return any(&stage.Xhtml_dfn_types_mapString).(*Type)
	case map[string]*Xhtml_div_type:
		return any(&stage.Xhtml_div_types_mapString).(*Type)
	case map[string]*Xhtml_dl_type:
		return any(&stage.Xhtml_dl_types_mapString).(*Type)
	case map[string]*Xhtml_dt_type:
		return any(&stage.Xhtml_dt_types_mapString).(*Type)
	case map[string]*Xhtml_edit_type:
		return any(&stage.Xhtml_edit_types_mapString).(*Type)
	case map[string]*Xhtml_em_type:
		return any(&stage.Xhtml_em_types_mapString).(*Type)
	case map[string]*Xhtml_h1_type:
		return any(&stage.Xhtml_h1_types_mapString).(*Type)
	case map[string]*Xhtml_h2_type:
		return any(&stage.Xhtml_h2_types_mapString).(*Type)
	case map[string]*Xhtml_h3_type:
		return any(&stage.Xhtml_h3_types_mapString).(*Type)
	case map[string]*Xhtml_h4_type:
		return any(&stage.Xhtml_h4_types_mapString).(*Type)
	case map[string]*Xhtml_h5_type:
		return any(&stage.Xhtml_h5_types_mapString).(*Type)
	case map[string]*Xhtml_h6_type:
		return any(&stage.Xhtml_h6_types_mapString).(*Type)
	case map[string]*Xhtml_heading_type:
		return any(&stage.Xhtml_heading_types_mapString).(*Type)
	case map[string]*Xhtml_hr_type:
		return any(&stage.Xhtml_hr_types_mapString).(*Type)
	case map[string]*Xhtml_kbd_type:
		return any(&stage.Xhtml_kbd_types_mapString).(*Type)
	case map[string]*Xhtml_li_type:
		return any(&stage.Xhtml_li_types_mapString).(*Type)
	case map[string]*Xhtml_object_type:
		return any(&stage.Xhtml_object_types_mapString).(*Type)
	case map[string]*Xhtml_ol_type:
		return any(&stage.Xhtml_ol_types_mapString).(*Type)
	case map[string]*Xhtml_p_type:
		return any(&stage.Xhtml_p_types_mapString).(*Type)
	case map[string]*Xhtml_param_type:
		return any(&stage.Xhtml_param_types_mapString).(*Type)
	case map[string]*Xhtml_pre_type:
		return any(&stage.Xhtml_pre_types_mapString).(*Type)
	case map[string]*Xhtml_q_type:
		return any(&stage.Xhtml_q_types_mapString).(*Type)
	case map[string]*Xhtml_samp_type:
		return any(&stage.Xhtml_samp_types_mapString).(*Type)
	case map[string]*Xhtml_span_type:
		return any(&stage.Xhtml_span_types_mapString).(*Type)
	case map[string]*Xhtml_strong_type:
		return any(&stage.Xhtml_strong_types_mapString).(*Type)
	case map[string]*Xhtml_table_type:
		return any(&stage.Xhtml_table_types_mapString).(*Type)
	case map[string]*Xhtml_tbody_type:
		return any(&stage.Xhtml_tbody_types_mapString).(*Type)
	case map[string]*Xhtml_td_type:
		return any(&stage.Xhtml_td_types_mapString).(*Type)
	case map[string]*Xhtml_tfoot_type:
		return any(&stage.Xhtml_tfoot_types_mapString).(*Type)
	case map[string]*Xhtml_th_type:
		return any(&stage.Xhtml_th_types_mapString).(*Type)
	case map[string]*Xhtml_thead_type:
		return any(&stage.Xhtml_thead_types_mapString).(*Type)
	case map[string]*Xhtml_tr_type:
		return any(&stage.Xhtml_tr_types_mapString).(*Type)
	case map[string]*Xhtml_ul_type:
		return any(&stage.Xhtml_ul_types_mapString).(*Type)
	case map[string]*Xhtml_var_type:
		return any(&stage.Xhtml_var_types_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case ALTERNATIVE_ID:
		return any(&stage.ALTERNATIVE_IDs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		return any(&stage.ATTRIBUTE_DEFINITION_BOOLEANs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_DATE:
		return any(&stage.ATTRIBUTE_DEFINITION_DATEs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_INTEGER:
		return any(&stage.ATTRIBUTE_DEFINITION_INTEGERs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_REAL:
		return any(&stage.ATTRIBUTE_DEFINITION_REALs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_STRING:
		return any(&stage.ATTRIBUTE_DEFINITION_STRINGs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_XHTML:
		return any(&stage.ATTRIBUTE_DEFINITION_XHTMLs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.ATTRIBUTE_VALUE_BOOLEANs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_DATE:
		return any(&stage.ATTRIBUTE_VALUE_DATEs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.ATTRIBUTE_VALUE_ENUMERATIONs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.ATTRIBUTE_VALUE_INTEGERs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_REAL:
		return any(&stage.ATTRIBUTE_VALUE_REALs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_STRING:
		return any(&stage.ATTRIBUTE_VALUE_STRINGs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_XHTML:
		return any(&stage.ATTRIBUTE_VALUE_XHTMLs).(*map[*Type]any)
	case AnyType:
		return any(&stage.AnyTypes).(*map[*Type]any)
	case DATATYPE_DEFINITION_BOOLEAN:
		return any(&stage.DATATYPE_DEFINITION_BOOLEANs).(*map[*Type]any)
	case DATATYPE_DEFINITION_DATE:
		return any(&stage.DATATYPE_DEFINITION_DATEs).(*map[*Type]any)
	case DATATYPE_DEFINITION_ENUMERATION:
		return any(&stage.DATATYPE_DEFINITION_ENUMERATIONs).(*map[*Type]any)
	case DATATYPE_DEFINITION_INTEGER:
		return any(&stage.DATATYPE_DEFINITION_INTEGERs).(*map[*Type]any)
	case DATATYPE_DEFINITION_REAL:
		return any(&stage.DATATYPE_DEFINITION_REALs).(*map[*Type]any)
	case DATATYPE_DEFINITION_STRING:
		return any(&stage.DATATYPE_DEFINITION_STRINGs).(*map[*Type]any)
	case DATATYPE_DEFINITION_XHTML:
		return any(&stage.DATATYPE_DEFINITION_XHTMLs).(*map[*Type]any)
	case EMBEDDED_VALUE:
		return any(&stage.EMBEDDED_VALUEs).(*map[*Type]any)
	case ENUM_VALUE:
		return any(&stage.ENUM_VALUEs).(*map[*Type]any)
	case RELATION_GROUP:
		return any(&stage.RELATION_GROUPs).(*map[*Type]any)
	case RELATION_GROUP_TYPE:
		return any(&stage.RELATION_GROUP_TYPEs).(*map[*Type]any)
	case REQ_IF:
		return any(&stage.REQ_IFs).(*map[*Type]any)
	case REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs).(*map[*Type]any)
	case REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs).(*map[*Type]any)
	case REQ_IF_TOOL_EXTENSION:
		return any(&stage.REQ_IF_TOOL_EXTENSIONs).(*map[*Type]any)
	case SPECIFICATION:
		return any(&stage.SPECIFICATIONs).(*map[*Type]any)
	case SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs).(*map[*Type]any)
	case SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs).(*map[*Type]any)
	case SPEC_OBJECT:
		return any(&stage.SPEC_OBJECTs).(*map[*Type]any)
	case SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs).(*map[*Type]any)
	case SPEC_RELATION:
		return any(&stage.SPEC_RELATIONs).(*map[*Type]any)
	case SPEC_RELATION_TYPE:
		return any(&stage.SPEC_RELATION_TYPEs).(*map[*Type]any)
	case XHTML_CONTENT:
		return any(&stage.XHTML_CONTENTs).(*map[*Type]any)
	case Xhtml_InlPres_type:
		return any(&stage.Xhtml_InlPres_types).(*map[*Type]any)
	case Xhtml_a_type:
		return any(&stage.Xhtml_a_types).(*map[*Type]any)
	case Xhtml_abbr_type:
		return any(&stage.Xhtml_abbr_types).(*map[*Type]any)
	case Xhtml_acronym_type:
		return any(&stage.Xhtml_acronym_types).(*map[*Type]any)
	case Xhtml_address_type:
		return any(&stage.Xhtml_address_types).(*map[*Type]any)
	case Xhtml_blockquote_type:
		return any(&stage.Xhtml_blockquote_types).(*map[*Type]any)
	case Xhtml_br_type:
		return any(&stage.Xhtml_br_types).(*map[*Type]any)
	case Xhtml_caption_type:
		return any(&stage.Xhtml_caption_types).(*map[*Type]any)
	case Xhtml_cite_type:
		return any(&stage.Xhtml_cite_types).(*map[*Type]any)
	case Xhtml_code_type:
		return any(&stage.Xhtml_code_types).(*map[*Type]any)
	case Xhtml_col_type:
		return any(&stage.Xhtml_col_types).(*map[*Type]any)
	case Xhtml_colgroup_type:
		return any(&stage.Xhtml_colgroup_types).(*map[*Type]any)
	case Xhtml_dd_type:
		return any(&stage.Xhtml_dd_types).(*map[*Type]any)
	case Xhtml_dfn_type:
		return any(&stage.Xhtml_dfn_types).(*map[*Type]any)
	case Xhtml_div_type:
		return any(&stage.Xhtml_div_types).(*map[*Type]any)
	case Xhtml_dl_type:
		return any(&stage.Xhtml_dl_types).(*map[*Type]any)
	case Xhtml_dt_type:
		return any(&stage.Xhtml_dt_types).(*map[*Type]any)
	case Xhtml_edit_type:
		return any(&stage.Xhtml_edit_types).(*map[*Type]any)
	case Xhtml_em_type:
		return any(&stage.Xhtml_em_types).(*map[*Type]any)
	case Xhtml_h1_type:
		return any(&stage.Xhtml_h1_types).(*map[*Type]any)
	case Xhtml_h2_type:
		return any(&stage.Xhtml_h2_types).(*map[*Type]any)
	case Xhtml_h3_type:
		return any(&stage.Xhtml_h3_types).(*map[*Type]any)
	case Xhtml_h4_type:
		return any(&stage.Xhtml_h4_types).(*map[*Type]any)
	case Xhtml_h5_type:
		return any(&stage.Xhtml_h5_types).(*map[*Type]any)
	case Xhtml_h6_type:
		return any(&stage.Xhtml_h6_types).(*map[*Type]any)
	case Xhtml_heading_type:
		return any(&stage.Xhtml_heading_types).(*map[*Type]any)
	case Xhtml_hr_type:
		return any(&stage.Xhtml_hr_types).(*map[*Type]any)
	case Xhtml_kbd_type:
		return any(&stage.Xhtml_kbd_types).(*map[*Type]any)
	case Xhtml_li_type:
		return any(&stage.Xhtml_li_types).(*map[*Type]any)
	case Xhtml_object_type:
		return any(&stage.Xhtml_object_types).(*map[*Type]any)
	case Xhtml_ol_type:
		return any(&stage.Xhtml_ol_types).(*map[*Type]any)
	case Xhtml_p_type:
		return any(&stage.Xhtml_p_types).(*map[*Type]any)
	case Xhtml_param_type:
		return any(&stage.Xhtml_param_types).(*map[*Type]any)
	case Xhtml_pre_type:
		return any(&stage.Xhtml_pre_types).(*map[*Type]any)
	case Xhtml_q_type:
		return any(&stage.Xhtml_q_types).(*map[*Type]any)
	case Xhtml_samp_type:
		return any(&stage.Xhtml_samp_types).(*map[*Type]any)
	case Xhtml_span_type:
		return any(&stage.Xhtml_span_types).(*map[*Type]any)
	case Xhtml_strong_type:
		return any(&stage.Xhtml_strong_types).(*map[*Type]any)
	case Xhtml_table_type:
		return any(&stage.Xhtml_table_types).(*map[*Type]any)
	case Xhtml_tbody_type:
		return any(&stage.Xhtml_tbody_types).(*map[*Type]any)
	case Xhtml_td_type:
		return any(&stage.Xhtml_td_types).(*map[*Type]any)
	case Xhtml_tfoot_type:
		return any(&stage.Xhtml_tfoot_types).(*map[*Type]any)
	case Xhtml_th_type:
		return any(&stage.Xhtml_th_types).(*map[*Type]any)
	case Xhtml_thead_type:
		return any(&stage.Xhtml_thead_types).(*map[*Type]any)
	case Xhtml_tr_type:
		return any(&stage.Xhtml_tr_types).(*map[*Type]any)
	case Xhtml_ul_type:
		return any(&stage.Xhtml_ul_types).(*map[*Type]any)
	case Xhtml_var_type:
		return any(&stage.Xhtml_var_types).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *ALTERNATIVE_ID:
		return any(&stage.ALTERNATIVE_IDs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		return any(&stage.ATTRIBUTE_DEFINITION_BOOLEANs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_DATE:
		return any(&stage.ATTRIBUTE_DEFINITION_DATEs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_INTEGER:
		return any(&stage.ATTRIBUTE_DEFINITION_INTEGERs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_REAL:
		return any(&stage.ATTRIBUTE_DEFINITION_REALs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_STRING:
		return any(&stage.ATTRIBUTE_DEFINITION_STRINGs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_XHTML:
		return any(&stage.ATTRIBUTE_DEFINITION_XHTMLs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.ATTRIBUTE_VALUE_BOOLEANs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_DATE:
		return any(&stage.ATTRIBUTE_VALUE_DATEs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.ATTRIBUTE_VALUE_ENUMERATIONs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.ATTRIBUTE_VALUE_INTEGERs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_REAL:
		return any(&stage.ATTRIBUTE_VALUE_REALs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_STRING:
		return any(&stage.ATTRIBUTE_VALUE_STRINGs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_XHTML:
		return any(&stage.ATTRIBUTE_VALUE_XHTMLs).(*map[Type]any)
	case *AnyType:
		return any(&stage.AnyTypes).(*map[Type]any)
	case *DATATYPE_DEFINITION_BOOLEAN:
		return any(&stage.DATATYPE_DEFINITION_BOOLEANs).(*map[Type]any)
	case *DATATYPE_DEFINITION_DATE:
		return any(&stage.DATATYPE_DEFINITION_DATEs).(*map[Type]any)
	case *DATATYPE_DEFINITION_ENUMERATION:
		return any(&stage.DATATYPE_DEFINITION_ENUMERATIONs).(*map[Type]any)
	case *DATATYPE_DEFINITION_INTEGER:
		return any(&stage.DATATYPE_DEFINITION_INTEGERs).(*map[Type]any)
	case *DATATYPE_DEFINITION_REAL:
		return any(&stage.DATATYPE_DEFINITION_REALs).(*map[Type]any)
	case *DATATYPE_DEFINITION_STRING:
		return any(&stage.DATATYPE_DEFINITION_STRINGs).(*map[Type]any)
	case *DATATYPE_DEFINITION_XHTML:
		return any(&stage.DATATYPE_DEFINITION_XHTMLs).(*map[Type]any)
	case *EMBEDDED_VALUE:
		return any(&stage.EMBEDDED_VALUEs).(*map[Type]any)
	case *ENUM_VALUE:
		return any(&stage.ENUM_VALUEs).(*map[Type]any)
	case *RELATION_GROUP:
		return any(&stage.RELATION_GROUPs).(*map[Type]any)
	case *RELATION_GROUP_TYPE:
		return any(&stage.RELATION_GROUP_TYPEs).(*map[Type]any)
	case *REQ_IF:
		return any(&stage.REQ_IFs).(*map[Type]any)
	case *REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs).(*map[Type]any)
	case *REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs).(*map[Type]any)
	case *REQ_IF_TOOL_EXTENSION:
		return any(&stage.REQ_IF_TOOL_EXTENSIONs).(*map[Type]any)
	case *SPECIFICATION:
		return any(&stage.SPECIFICATIONs).(*map[Type]any)
	case *SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs).(*map[Type]any)
	case *SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs).(*map[Type]any)
	case *SPEC_OBJECT:
		return any(&stage.SPEC_OBJECTs).(*map[Type]any)
	case *SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs).(*map[Type]any)
	case *SPEC_RELATION:
		return any(&stage.SPEC_RELATIONs).(*map[Type]any)
	case *SPEC_RELATION_TYPE:
		return any(&stage.SPEC_RELATION_TYPEs).(*map[Type]any)
	case *XHTML_CONTENT:
		return any(&stage.XHTML_CONTENTs).(*map[Type]any)
	case *Xhtml_InlPres_type:
		return any(&stage.Xhtml_InlPres_types).(*map[Type]any)
	case *Xhtml_a_type:
		return any(&stage.Xhtml_a_types).(*map[Type]any)
	case *Xhtml_abbr_type:
		return any(&stage.Xhtml_abbr_types).(*map[Type]any)
	case *Xhtml_acronym_type:
		return any(&stage.Xhtml_acronym_types).(*map[Type]any)
	case *Xhtml_address_type:
		return any(&stage.Xhtml_address_types).(*map[Type]any)
	case *Xhtml_blockquote_type:
		return any(&stage.Xhtml_blockquote_types).(*map[Type]any)
	case *Xhtml_br_type:
		return any(&stage.Xhtml_br_types).(*map[Type]any)
	case *Xhtml_caption_type:
		return any(&stage.Xhtml_caption_types).(*map[Type]any)
	case *Xhtml_cite_type:
		return any(&stage.Xhtml_cite_types).(*map[Type]any)
	case *Xhtml_code_type:
		return any(&stage.Xhtml_code_types).(*map[Type]any)
	case *Xhtml_col_type:
		return any(&stage.Xhtml_col_types).(*map[Type]any)
	case *Xhtml_colgroup_type:
		return any(&stage.Xhtml_colgroup_types).(*map[Type]any)
	case *Xhtml_dd_type:
		return any(&stage.Xhtml_dd_types).(*map[Type]any)
	case *Xhtml_dfn_type:
		return any(&stage.Xhtml_dfn_types).(*map[Type]any)
	case *Xhtml_div_type:
		return any(&stage.Xhtml_div_types).(*map[Type]any)
	case *Xhtml_dl_type:
		return any(&stage.Xhtml_dl_types).(*map[Type]any)
	case *Xhtml_dt_type:
		return any(&stage.Xhtml_dt_types).(*map[Type]any)
	case *Xhtml_edit_type:
		return any(&stage.Xhtml_edit_types).(*map[Type]any)
	case *Xhtml_em_type:
		return any(&stage.Xhtml_em_types).(*map[Type]any)
	case *Xhtml_h1_type:
		return any(&stage.Xhtml_h1_types).(*map[Type]any)
	case *Xhtml_h2_type:
		return any(&stage.Xhtml_h2_types).(*map[Type]any)
	case *Xhtml_h3_type:
		return any(&stage.Xhtml_h3_types).(*map[Type]any)
	case *Xhtml_h4_type:
		return any(&stage.Xhtml_h4_types).(*map[Type]any)
	case *Xhtml_h5_type:
		return any(&stage.Xhtml_h5_types).(*map[Type]any)
	case *Xhtml_h6_type:
		return any(&stage.Xhtml_h6_types).(*map[Type]any)
	case *Xhtml_heading_type:
		return any(&stage.Xhtml_heading_types).(*map[Type]any)
	case *Xhtml_hr_type:
		return any(&stage.Xhtml_hr_types).(*map[Type]any)
	case *Xhtml_kbd_type:
		return any(&stage.Xhtml_kbd_types).(*map[Type]any)
	case *Xhtml_li_type:
		return any(&stage.Xhtml_li_types).(*map[Type]any)
	case *Xhtml_object_type:
		return any(&stage.Xhtml_object_types).(*map[Type]any)
	case *Xhtml_ol_type:
		return any(&stage.Xhtml_ol_types).(*map[Type]any)
	case *Xhtml_p_type:
		return any(&stage.Xhtml_p_types).(*map[Type]any)
	case *Xhtml_param_type:
		return any(&stage.Xhtml_param_types).(*map[Type]any)
	case *Xhtml_pre_type:
		return any(&stage.Xhtml_pre_types).(*map[Type]any)
	case *Xhtml_q_type:
		return any(&stage.Xhtml_q_types).(*map[Type]any)
	case *Xhtml_samp_type:
		return any(&stage.Xhtml_samp_types).(*map[Type]any)
	case *Xhtml_span_type:
		return any(&stage.Xhtml_span_types).(*map[Type]any)
	case *Xhtml_strong_type:
		return any(&stage.Xhtml_strong_types).(*map[Type]any)
	case *Xhtml_table_type:
		return any(&stage.Xhtml_table_types).(*map[Type]any)
	case *Xhtml_tbody_type:
		return any(&stage.Xhtml_tbody_types).(*map[Type]any)
	case *Xhtml_td_type:
		return any(&stage.Xhtml_td_types).(*map[Type]any)
	case *Xhtml_tfoot_type:
		return any(&stage.Xhtml_tfoot_types).(*map[Type]any)
	case *Xhtml_th_type:
		return any(&stage.Xhtml_th_types).(*map[Type]any)
	case *Xhtml_thead_type:
		return any(&stage.Xhtml_thead_types).(*map[Type]any)
	case *Xhtml_tr_type:
		return any(&stage.Xhtml_tr_types).(*map[Type]any)
	case *Xhtml_ul_type:
		return any(&stage.Xhtml_ul_types).(*map[Type]any)
	case *Xhtml_var_type:
		return any(&stage.Xhtml_var_types).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case ALTERNATIVE_ID:
		return any(&stage.ALTERNATIVE_IDs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		return any(&stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_DATE:
		return any(&stage.ATTRIBUTE_DEFINITION_DATEs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_INTEGER:
		return any(&stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_REAL:
		return any(&stage.ATTRIBUTE_DEFINITION_REALs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_STRING:
		return any(&stage.ATTRIBUTE_DEFINITION_STRINGs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_XHTML:
		return any(&stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.ATTRIBUTE_VALUE_BOOLEANs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_DATE:
		return any(&stage.ATTRIBUTE_VALUE_DATEs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.ATTRIBUTE_VALUE_INTEGERs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_REAL:
		return any(&stage.ATTRIBUTE_VALUE_REALs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_STRING:
		return any(&stage.ATTRIBUTE_VALUE_STRINGs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_XHTML:
		return any(&stage.ATTRIBUTE_VALUE_XHTMLs_mapString).(*map[string]*Type)
	case AnyType:
		return any(&stage.AnyTypes_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_BOOLEAN:
		return any(&stage.DATATYPE_DEFINITION_BOOLEANs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_DATE:
		return any(&stage.DATATYPE_DEFINITION_DATEs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_ENUMERATION:
		return any(&stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_INTEGER:
		return any(&stage.DATATYPE_DEFINITION_INTEGERs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_REAL:
		return any(&stage.DATATYPE_DEFINITION_REALs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_STRING:
		return any(&stage.DATATYPE_DEFINITION_STRINGs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_XHTML:
		return any(&stage.DATATYPE_DEFINITION_XHTMLs_mapString).(*map[string]*Type)
	case EMBEDDED_VALUE:
		return any(&stage.EMBEDDED_VALUEs_mapString).(*map[string]*Type)
	case ENUM_VALUE:
		return any(&stage.ENUM_VALUEs_mapString).(*map[string]*Type)
	case RELATION_GROUP:
		return any(&stage.RELATION_GROUPs_mapString).(*map[string]*Type)
	case RELATION_GROUP_TYPE:
		return any(&stage.RELATION_GROUP_TYPEs_mapString).(*map[string]*Type)
	case REQ_IF:
		return any(&stage.REQ_IFs_mapString).(*map[string]*Type)
	case REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs_mapString).(*map[string]*Type)
	case REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs_mapString).(*map[string]*Type)
	case REQ_IF_TOOL_EXTENSION:
		return any(&stage.REQ_IF_TOOL_EXTENSIONs_mapString).(*map[string]*Type)
	case SPECIFICATION:
		return any(&stage.SPECIFICATIONs_mapString).(*map[string]*Type)
	case SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs_mapString).(*map[string]*Type)
	case SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs_mapString).(*map[string]*Type)
	case SPEC_OBJECT:
		return any(&stage.SPEC_OBJECTs_mapString).(*map[string]*Type)
	case SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs_mapString).(*map[string]*Type)
	case SPEC_RELATION:
		return any(&stage.SPEC_RELATIONs_mapString).(*map[string]*Type)
	case SPEC_RELATION_TYPE:
		return any(&stage.SPEC_RELATION_TYPEs_mapString).(*map[string]*Type)
	case XHTML_CONTENT:
		return any(&stage.XHTML_CONTENTs_mapString).(*map[string]*Type)
	case Xhtml_InlPres_type:
		return any(&stage.Xhtml_InlPres_types_mapString).(*map[string]*Type)
	case Xhtml_a_type:
		return any(&stage.Xhtml_a_types_mapString).(*map[string]*Type)
	case Xhtml_abbr_type:
		return any(&stage.Xhtml_abbr_types_mapString).(*map[string]*Type)
	case Xhtml_acronym_type:
		return any(&stage.Xhtml_acronym_types_mapString).(*map[string]*Type)
	case Xhtml_address_type:
		return any(&stage.Xhtml_address_types_mapString).(*map[string]*Type)
	case Xhtml_blockquote_type:
		return any(&stage.Xhtml_blockquote_types_mapString).(*map[string]*Type)
	case Xhtml_br_type:
		return any(&stage.Xhtml_br_types_mapString).(*map[string]*Type)
	case Xhtml_caption_type:
		return any(&stage.Xhtml_caption_types_mapString).(*map[string]*Type)
	case Xhtml_cite_type:
		return any(&stage.Xhtml_cite_types_mapString).(*map[string]*Type)
	case Xhtml_code_type:
		return any(&stage.Xhtml_code_types_mapString).(*map[string]*Type)
	case Xhtml_col_type:
		return any(&stage.Xhtml_col_types_mapString).(*map[string]*Type)
	case Xhtml_colgroup_type:
		return any(&stage.Xhtml_colgroup_types_mapString).(*map[string]*Type)
	case Xhtml_dd_type:
		return any(&stage.Xhtml_dd_types_mapString).(*map[string]*Type)
	case Xhtml_dfn_type:
		return any(&stage.Xhtml_dfn_types_mapString).(*map[string]*Type)
	case Xhtml_div_type:
		return any(&stage.Xhtml_div_types_mapString).(*map[string]*Type)
	case Xhtml_dl_type:
		return any(&stage.Xhtml_dl_types_mapString).(*map[string]*Type)
	case Xhtml_dt_type:
		return any(&stage.Xhtml_dt_types_mapString).(*map[string]*Type)
	case Xhtml_edit_type:
		return any(&stage.Xhtml_edit_types_mapString).(*map[string]*Type)
	case Xhtml_em_type:
		return any(&stage.Xhtml_em_types_mapString).(*map[string]*Type)
	case Xhtml_h1_type:
		return any(&stage.Xhtml_h1_types_mapString).(*map[string]*Type)
	case Xhtml_h2_type:
		return any(&stage.Xhtml_h2_types_mapString).(*map[string]*Type)
	case Xhtml_h3_type:
		return any(&stage.Xhtml_h3_types_mapString).(*map[string]*Type)
	case Xhtml_h4_type:
		return any(&stage.Xhtml_h4_types_mapString).(*map[string]*Type)
	case Xhtml_h5_type:
		return any(&stage.Xhtml_h5_types_mapString).(*map[string]*Type)
	case Xhtml_h6_type:
		return any(&stage.Xhtml_h6_types_mapString).(*map[string]*Type)
	case Xhtml_heading_type:
		return any(&stage.Xhtml_heading_types_mapString).(*map[string]*Type)
	case Xhtml_hr_type:
		return any(&stage.Xhtml_hr_types_mapString).(*map[string]*Type)
	case Xhtml_kbd_type:
		return any(&stage.Xhtml_kbd_types_mapString).(*map[string]*Type)
	case Xhtml_li_type:
		return any(&stage.Xhtml_li_types_mapString).(*map[string]*Type)
	case Xhtml_object_type:
		return any(&stage.Xhtml_object_types_mapString).(*map[string]*Type)
	case Xhtml_ol_type:
		return any(&stage.Xhtml_ol_types_mapString).(*map[string]*Type)
	case Xhtml_p_type:
		return any(&stage.Xhtml_p_types_mapString).(*map[string]*Type)
	case Xhtml_param_type:
		return any(&stage.Xhtml_param_types_mapString).(*map[string]*Type)
	case Xhtml_pre_type:
		return any(&stage.Xhtml_pre_types_mapString).(*map[string]*Type)
	case Xhtml_q_type:
		return any(&stage.Xhtml_q_types_mapString).(*map[string]*Type)
	case Xhtml_samp_type:
		return any(&stage.Xhtml_samp_types_mapString).(*map[string]*Type)
	case Xhtml_span_type:
		return any(&stage.Xhtml_span_types_mapString).(*map[string]*Type)
	case Xhtml_strong_type:
		return any(&stage.Xhtml_strong_types_mapString).(*map[string]*Type)
	case Xhtml_table_type:
		return any(&stage.Xhtml_table_types_mapString).(*map[string]*Type)
	case Xhtml_tbody_type:
		return any(&stage.Xhtml_tbody_types_mapString).(*map[string]*Type)
	case Xhtml_td_type:
		return any(&stage.Xhtml_td_types_mapString).(*map[string]*Type)
	case Xhtml_tfoot_type:
		return any(&stage.Xhtml_tfoot_types_mapString).(*map[string]*Type)
	case Xhtml_th_type:
		return any(&stage.Xhtml_th_types_mapString).(*map[string]*Type)
	case Xhtml_thead_type:
		return any(&stage.Xhtml_thead_types_mapString).(*map[string]*Type)
	case Xhtml_tr_type:
		return any(&stage.Xhtml_tr_types_mapString).(*map[string]*Type)
	case Xhtml_ul_type:
		return any(&stage.Xhtml_ul_types_mapString).(*map[string]*Type)
	case Xhtml_var_type:
		return any(&stage.Xhtml_var_types_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case ALTERNATIVE_ID:
		return any(&ALTERNATIVE_ID{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		return any(&ATTRIBUTE_DEFINITION_BOOLEAN{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_DATE:
		return any(&ATTRIBUTE_DEFINITION_DATE{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&ATTRIBUTE_DEFINITION_ENUMERATION{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_INTEGER:
		return any(&ATTRIBUTE_DEFINITION_INTEGER{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_REAL:
		return any(&ATTRIBUTE_DEFINITION_REAL{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_STRING:
		return any(&ATTRIBUTE_DEFINITION_STRING{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_XHTML:
		return any(&ATTRIBUTE_DEFINITION_XHTML{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_BOOLEAN:
		return any(&ATTRIBUTE_VALUE_BOOLEAN{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_DATE:
		return any(&ATTRIBUTE_VALUE_DATE{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_ENUMERATION:
		return any(&ATTRIBUTE_VALUE_ENUMERATION{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_INTEGER:
		return any(&ATTRIBUTE_VALUE_INTEGER{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_REAL:
		return any(&ATTRIBUTE_VALUE_REAL{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_STRING:
		return any(&ATTRIBUTE_VALUE_STRING{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_XHTML:
		return any(&ATTRIBUTE_VALUE_XHTML{
			// Initialisation of associations
			// field is initialized with an instance of XHTML_CONTENT with the name of the field
			THE_VALUE: []*XHTML_CONTENT{{Name: "THE_VALUE"}},
			// field is initialized with an instance of XHTML_CONTENT with the name of the field
			THE_ORIGINAL_VALUE: []*XHTML_CONTENT{{Name: "THE_ORIGINAL_VALUE"}},
		}).(*Type)
	case AnyType:
		return any(&AnyType{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_BOOLEAN:
		return any(&DATATYPE_DEFINITION_BOOLEAN{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_DATE:
		return any(&DATATYPE_DEFINITION_DATE{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_ENUMERATION:
		return any(&DATATYPE_DEFINITION_ENUMERATION{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_INTEGER:
		return any(&DATATYPE_DEFINITION_INTEGER{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_REAL:
		return any(&DATATYPE_DEFINITION_REAL{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_STRING:
		return any(&DATATYPE_DEFINITION_STRING{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_XHTML:
		return any(&DATATYPE_DEFINITION_XHTML{
			// Initialisation of associations
		}).(*Type)
	case EMBEDDED_VALUE:
		return any(&EMBEDDED_VALUE{
			// Initialisation of associations
		}).(*Type)
	case ENUM_VALUE:
		return any(&ENUM_VALUE{
			// Initialisation of associations
		}).(*Type)
	case RELATION_GROUP:
		return any(&RELATION_GROUP{
			// Initialisation of associations
		}).(*Type)
	case RELATION_GROUP_TYPE:
		return any(&RELATION_GROUP_TYPE{
			// Initialisation of associations
		}).(*Type)
	case REQ_IF:
		return any(&REQ_IF{
			// Initialisation of associations
		}).(*Type)
	case REQ_IF_CONTENT:
		return any(&REQ_IF_CONTENT{
			// Initialisation of associations
		}).(*Type)
	case REQ_IF_HEADER:
		return any(&REQ_IF_HEADER{
			// Initialisation of associations
		}).(*Type)
	case REQ_IF_TOOL_EXTENSION:
		return any(&REQ_IF_TOOL_EXTENSION{
			// Initialisation of associations
		}).(*Type)
	case SPECIFICATION:
		return any(&SPECIFICATION{
			// Initialisation of associations
		}).(*Type)
	case SPECIFICATION_TYPE:
		return any(&SPECIFICATION_TYPE{
			// Initialisation of associations
		}).(*Type)
	case SPEC_HIERARCHY:
		return any(&SPEC_HIERARCHY{
			// Initialisation of associations
		}).(*Type)
	case SPEC_OBJECT:
		return any(&SPEC_OBJECT{
			// Initialisation of associations
		}).(*Type)
	case SPEC_OBJECT_TYPE:
		return any(&SPEC_OBJECT_TYPE{
			// Initialisation of associations
		}).(*Type)
	case SPEC_RELATION:
		return any(&SPEC_RELATION{
			// Initialisation of associations
		}).(*Type)
	case SPEC_RELATION_TYPE:
		return any(&SPEC_RELATION_TYPE{
			// Initialisation of associations
		}).(*Type)
	case XHTML_CONTENT:
		return any(&XHTML_CONTENT{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_InlPres_type:
		return any(&Xhtml_InlPres_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_a_type:
		return any(&Xhtml_a_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_abbr_type:
		return any(&Xhtml_abbr_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_acronym_type:
		return any(&Xhtml_acronym_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_address_type:
		return any(&Xhtml_address_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_blockquote_type:
		return any(&Xhtml_blockquote_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_br_type:
		return any(&Xhtml_br_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_caption_type:
		return any(&Xhtml_caption_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_cite_type:
		return any(&Xhtml_cite_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_code_type:
		return any(&Xhtml_code_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_col_type:
		return any(&Xhtml_col_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_colgroup_type:
		return any(&Xhtml_colgroup_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_dd_type:
		return any(&Xhtml_dd_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_dfn_type:
		return any(&Xhtml_dfn_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_div_type:
		return any(&Xhtml_div_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_dl_type:
		return any(&Xhtml_dl_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_dt_type:
		return any(&Xhtml_dt_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_edit_type:
		return any(&Xhtml_edit_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_em_type:
		return any(&Xhtml_em_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_h1_type:
		return any(&Xhtml_h1_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_h2_type:
		return any(&Xhtml_h2_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_h3_type:
		return any(&Xhtml_h3_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_h4_type:
		return any(&Xhtml_h4_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_h5_type:
		return any(&Xhtml_h5_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_h6_type:
		return any(&Xhtml_h6_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_heading_type:
		return any(&Xhtml_heading_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_hr_type:
		return any(&Xhtml_hr_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_kbd_type:
		return any(&Xhtml_kbd_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_li_type:
		return any(&Xhtml_li_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_object_type:
		return any(&Xhtml_object_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_ol_type:
		return any(&Xhtml_ol_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_p_type:
		return any(&Xhtml_p_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_param_type:
		return any(&Xhtml_param_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_pre_type:
		return any(&Xhtml_pre_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_q_type:
		return any(&Xhtml_q_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_samp_type:
		return any(&Xhtml_samp_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_span_type:
		return any(&Xhtml_span_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_strong_type:
		return any(&Xhtml_strong_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_table_type:
		return any(&Xhtml_table_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_tbody_type:
		return any(&Xhtml_tbody_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_td_type:
		return any(&Xhtml_td_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_tfoot_type:
		return any(&Xhtml_tfoot_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_th_type:
		return any(&Xhtml_th_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_thead_type:
		return any(&Xhtml_thead_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_tr_type:
		return any(&Xhtml_tr_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_ul_type:
		return any(&Xhtml_ul_type{
			// Initialisation of associations
		}).(*Type)
	case Xhtml_var_type:
		return any(&Xhtml_var_type{
			// Initialisation of associations
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of ALTERNATIVE_ID
	case ALTERNATIVE_ID:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_BOOLEAN
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_DATE
	case ATTRIBUTE_DEFINITION_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_ENUMERATION
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_INTEGER
	case ATTRIBUTE_DEFINITION_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_REAL
	case ATTRIBUTE_DEFINITION_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_STRING
	case ATTRIBUTE_DEFINITION_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_XHTML
	case ATTRIBUTE_DEFINITION_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_BOOLEAN
	case ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_DATE
	case ATTRIBUTE_VALUE_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_ENUMERATION
	case ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_INTEGER
	case ATTRIBUTE_VALUE_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_REAL
	case ATTRIBUTE_VALUE_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_STRING
	case ATTRIBUTE_VALUE_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_XHTML
	case ATTRIBUTE_VALUE_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AnyType
	case AnyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_BOOLEAN
	case DATATYPE_DEFINITION_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_DATE
	case DATATYPE_DEFINITION_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_ENUMERATION
	case DATATYPE_DEFINITION_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_INTEGER
	case DATATYPE_DEFINITION_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_REAL
	case DATATYPE_DEFINITION_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_STRING
	case DATATYPE_DEFINITION_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_XHTML
	case DATATYPE_DEFINITION_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EMBEDDED_VALUE
	case EMBEDDED_VALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ENUM_VALUE
	case ENUM_VALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RELATION_GROUP
	case RELATION_GROUP:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RELATION_GROUP_TYPE
	case RELATION_GROUP_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF
	case REQ_IF:
		switch fieldname {
		// insertion point for per direct association field
		case "THE_HEADER.REQ_IF_HEADER":
			res := make(map[*REQ_IF_HEADER][]*REQ_IF)
			for req_if := range stage.REQ_IFs {
				if req_if.THE_HEADER.REQ_IF_HEADER != nil {
					req_if_header_ := req_if.THE_HEADER.REQ_IF_HEADER
					var req_ifs []*REQ_IF
					_, ok := res[req_if_header_]
					if ok {
						req_ifs = res[req_if_header_]
					} else {
						req_ifs = make([]*REQ_IF, 0)
					}
					req_ifs = append(req_ifs, req_if)
					res[req_if_header_] = req_ifs
				}
			}
			return any(res).(map[*End][]*Start)
		case "CORE_CONTENT.REQ_IF_CONTENT":
			res := make(map[*REQ_IF_CONTENT][]*REQ_IF)
			for req_if := range stage.REQ_IFs {
				if req_if.CORE_CONTENT.REQ_IF_CONTENT != nil {
					req_if_content_ := req_if.CORE_CONTENT.REQ_IF_CONTENT
					var req_ifs []*REQ_IF
					_, ok := res[req_if_content_]
					if ok {
						req_ifs = res[req_if_content_]
					} else {
						req_ifs = make([]*REQ_IF, 0)
					}
					req_ifs = append(req_ifs, req_if)
					res[req_if_content_] = req_ifs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of REQ_IF_CONTENT
	case REQ_IF_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_HEADER
	case REQ_IF_HEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_TOOL_EXTENSION
	case REQ_IF_TOOL_EXTENSION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION
	case SPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION_TYPE
	case SPECIFICATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_HIERARCHY
	case SPEC_HIERARCHY:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_OBJECT
	case SPEC_OBJECT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_OBJECT_TYPE
	case SPEC_OBJECT_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_RELATION
	case SPEC_RELATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_RELATION_TYPE
	case SPEC_RELATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XHTML_CONTENT
	case XHTML_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_InlPres_type
	case Xhtml_InlPres_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_a_type
	case Xhtml_a_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_abbr_type
	case Xhtml_abbr_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_acronym_type
	case Xhtml_acronym_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_address_type
	case Xhtml_address_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_blockquote_type
	case Xhtml_blockquote_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_br_type
	case Xhtml_br_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_caption_type
	case Xhtml_caption_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_cite_type
	case Xhtml_cite_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_code_type
	case Xhtml_code_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_col_type
	case Xhtml_col_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_colgroup_type
	case Xhtml_colgroup_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_dd_type
	case Xhtml_dd_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_dfn_type
	case Xhtml_dfn_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_div_type
	case Xhtml_div_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_dl_type
	case Xhtml_dl_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_dt_type
	case Xhtml_dt_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_edit_type
	case Xhtml_edit_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_em_type
	case Xhtml_em_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h1_type
	case Xhtml_h1_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h2_type
	case Xhtml_h2_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h3_type
	case Xhtml_h3_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h4_type
	case Xhtml_h4_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h5_type
	case Xhtml_h5_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h6_type
	case Xhtml_h6_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_heading_type
	case Xhtml_heading_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_hr_type
	case Xhtml_hr_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_kbd_type
	case Xhtml_kbd_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_li_type
	case Xhtml_li_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_object_type
	case Xhtml_object_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_ol_type
	case Xhtml_ol_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_p_type
	case Xhtml_p_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_param_type
	case Xhtml_param_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_pre_type
	case Xhtml_pre_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_q_type
	case Xhtml_q_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_samp_type
	case Xhtml_samp_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_span_type
	case Xhtml_span_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_strong_type
	case Xhtml_strong_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_table_type
	case Xhtml_table_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_tbody_type
	case Xhtml_tbody_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_td_type
	case Xhtml_td_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_tfoot_type
	case Xhtml_tfoot_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_th_type
	case Xhtml_th_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_thead_type
	case Xhtml_thead_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_tr_type
	case Xhtml_tr_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_ul_type
	case Xhtml_ul_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_var_type
	case Xhtml_var_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of ALTERNATIVE_ID
	case ALTERNATIVE_ID:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_BOOLEAN
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_BOOLEAN)
			for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
				for _, alternative_id_ := range attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = attribute_definition_boolean
				}
			}
			return any(res).(map[*End]*Start)
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN":
			res := make(map[*ATTRIBUTE_VALUE_BOOLEAN]*ATTRIBUTE_DEFINITION_BOOLEAN)
			for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
				for _, attribute_value_boolean_ := range attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
					res[attribute_value_boolean_] = attribute_definition_boolean
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_DATE
	case ATTRIBUTE_DEFINITION_DATE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_DATE)
			for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
				for _, alternative_id_ := range attribute_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = attribute_definition_date
				}
			}
			return any(res).(map[*End]*Start)
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE":
			res := make(map[*ATTRIBUTE_VALUE_DATE]*ATTRIBUTE_DEFINITION_DATE)
			for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
				for _, attribute_value_date_ := range attribute_definition_date.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE {
					res[attribute_value_date_] = attribute_definition_date
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_ENUMERATION
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION":
			res := make(map[*ATTRIBUTE_VALUE_ENUMERATION]*ATTRIBUTE_DEFINITION_ENUMERATION)
			for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
				for _, attribute_value_enumeration_ := range attribute_definition_enumeration.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION {
					res[attribute_value_enumeration_] = attribute_definition_enumeration
				}
			}
			return any(res).(map[*End]*Start)
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_ENUMERATION)
			for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
				for _, alternative_id_ := range attribute_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = attribute_definition_enumeration
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_INTEGER
	case ATTRIBUTE_DEFINITION_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_INTEGER)
			for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
				for _, alternative_id_ := range attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = attribute_definition_integer
				}
			}
			return any(res).(map[*End]*Start)
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER":
			res := make(map[*ATTRIBUTE_VALUE_INTEGER]*ATTRIBUTE_DEFINITION_INTEGER)
			for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
				for _, attribute_value_integer_ := range attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
					res[attribute_value_integer_] = attribute_definition_integer
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_REAL
	case ATTRIBUTE_DEFINITION_REAL:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_REAL)
			for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
				for _, alternative_id_ := range attribute_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = attribute_definition_real
				}
			}
			return any(res).(map[*End]*Start)
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL":
			res := make(map[*ATTRIBUTE_VALUE_REAL]*ATTRIBUTE_DEFINITION_REAL)
			for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
				for _, attribute_value_real_ := range attribute_definition_real.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL {
					res[attribute_value_real_] = attribute_definition_real
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_STRING
	case ATTRIBUTE_DEFINITION_STRING:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_STRING)
			for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
				for _, alternative_id_ := range attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = attribute_definition_string
				}
			}
			return any(res).(map[*End]*Start)
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING":
			res := make(map[*ATTRIBUTE_VALUE_STRING]*ATTRIBUTE_DEFINITION_STRING)
			for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
				for _, attribute_value_string_ := range attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
					res[attribute_value_string_] = attribute_definition_string
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_XHTML
	case ATTRIBUTE_DEFINITION_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_XHTML)
			for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
				for _, alternative_id_ := range attribute_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = attribute_definition_xhtml
				}
			}
			return any(res).(map[*End]*Start)
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML":
			res := make(map[*ATTRIBUTE_VALUE_XHTML]*ATTRIBUTE_DEFINITION_XHTML)
			for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
				for _, attribute_value_xhtml_ := range attribute_definition_xhtml.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML {
					res[attribute_value_xhtml_] = attribute_definition_xhtml
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_BOOLEAN
	case ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_DATE
	case ATTRIBUTE_VALUE_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_ENUMERATION
	case ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_INTEGER
	case ATTRIBUTE_VALUE_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_REAL
	case ATTRIBUTE_VALUE_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_STRING
	case ATTRIBUTE_VALUE_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_XHTML
	case ATTRIBUTE_VALUE_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "THE_VALUE":
			res := make(map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML)
			for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
				for _, xhtml_content_ := range attribute_value_xhtml.THE_VALUE {
					res[xhtml_content_] = attribute_value_xhtml
				}
			}
			return any(res).(map[*End]*Start)
		case "THE_ORIGINAL_VALUE":
			res := make(map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML)
			for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
				for _, xhtml_content_ := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
					res[xhtml_content_] = attribute_value_xhtml
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of AnyType
	case AnyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_BOOLEAN
	case DATATYPE_DEFINITION_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_BOOLEAN)
			for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
				for _, alternative_id_ := range datatype_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = datatype_definition_boolean
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_DATE
	case DATATYPE_DEFINITION_DATE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_DATE)
			for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
				for _, alternative_id_ := range datatype_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = datatype_definition_date
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_ENUMERATION
	case DATATYPE_DEFINITION_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_ENUMERATION)
			for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
				for _, alternative_id_ := range datatype_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = datatype_definition_enumeration
				}
			}
			return any(res).(map[*End]*Start)
		case "SPECIFIED_VALUES.ENUM_VALUE":
			res := make(map[*ENUM_VALUE]*DATATYPE_DEFINITION_ENUMERATION)
			for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
				for _, enum_value_ := range datatype_definition_enumeration.SPECIFIED_VALUES.ENUM_VALUE {
					res[enum_value_] = datatype_definition_enumeration
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_INTEGER
	case DATATYPE_DEFINITION_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_INTEGER)
			for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
				for _, alternative_id_ := range datatype_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = datatype_definition_integer
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_REAL
	case DATATYPE_DEFINITION_REAL:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_REAL)
			for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
				for _, alternative_id_ := range datatype_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = datatype_definition_real
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_STRING
	case DATATYPE_DEFINITION_STRING:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_STRING)
			for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
				for _, alternative_id_ := range datatype_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = datatype_definition_string
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_XHTML
	case DATATYPE_DEFINITION_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_XHTML)
			for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
				for _, alternative_id_ := range datatype_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = datatype_definition_xhtml
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of EMBEDDED_VALUE
	case EMBEDDED_VALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ENUM_VALUE
	case ENUM_VALUE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*ENUM_VALUE)
			for enum_value := range stage.ENUM_VALUEs {
				for _, alternative_id_ := range enum_value.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = enum_value
				}
			}
			return any(res).(map[*End]*Start)
		case "PROPERTIES.EMBEDDED_VALUE":
			res := make(map[*EMBEDDED_VALUE]*ENUM_VALUE)
			for enum_value := range stage.ENUM_VALUEs {
				for _, embedded_value_ := range enum_value.PROPERTIES.EMBEDDED_VALUE {
					res[embedded_value_] = enum_value
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of RELATION_GROUP
	case RELATION_GROUP:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*RELATION_GROUP)
			for relation_group := range stage.RELATION_GROUPs {
				for _, alternative_id_ := range relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = relation_group
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of RELATION_GROUP_TYPE
	case RELATION_GROUP_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*RELATION_GROUP_TYPE)
			for relation_group_type := range stage.RELATION_GROUP_TYPEs {
				for _, alternative_id_ := range relation_group_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = relation_group_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			res := make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*RELATION_GROUP_TYPE)
			for relation_group_type := range stage.RELATION_GROUP_TYPEs {
				for _, attribute_definition_boolean_ := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
					res[attribute_definition_boolean_] = relation_group_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			res := make(map[*ATTRIBUTE_DEFINITION_DATE]*RELATION_GROUP_TYPE)
			for relation_group_type := range stage.RELATION_GROUP_TYPEs {
				for _, attribute_definition_date_ := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
					res[attribute_definition_date_] = relation_group_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			res := make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*RELATION_GROUP_TYPE)
			for relation_group_type := range stage.RELATION_GROUP_TYPEs {
				for _, attribute_definition_enumeration_ := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
					res[attribute_definition_enumeration_] = relation_group_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			res := make(map[*ATTRIBUTE_DEFINITION_INTEGER]*RELATION_GROUP_TYPE)
			for relation_group_type := range stage.RELATION_GROUP_TYPEs {
				for _, attribute_definition_integer_ := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
					res[attribute_definition_integer_] = relation_group_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			res := make(map[*ATTRIBUTE_DEFINITION_REAL]*RELATION_GROUP_TYPE)
			for relation_group_type := range stage.RELATION_GROUP_TYPEs {
				for _, attribute_definition_real_ := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
					res[attribute_definition_real_] = relation_group_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			res := make(map[*ATTRIBUTE_DEFINITION_STRING]*RELATION_GROUP_TYPE)
			for relation_group_type := range stage.RELATION_GROUP_TYPEs {
				for _, attribute_definition_string_ := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
					res[attribute_definition_string_] = relation_group_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			res := make(map[*ATTRIBUTE_DEFINITION_XHTML]*RELATION_GROUP_TYPE)
			for relation_group_type := range stage.RELATION_GROUP_TYPEs {
				for _, attribute_definition_xhtml_ := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
					res[attribute_definition_xhtml_] = relation_group_type
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of REQ_IF
	case REQ_IF:
		switch fieldname {
		// insertion point for per direct association field
		case "TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION":
			res := make(map[*REQ_IF_TOOL_EXTENSION]*REQ_IF)
			for req_if := range stage.REQ_IFs {
				for _, req_if_tool_extension_ := range req_if.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION {
					res[req_if_tool_extension_] = req_if
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of REQ_IF_CONTENT
	case REQ_IF_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		case "DATATYPES.DATATYPE_DEFINITION_BOOLEAN":
			res := make(map[*DATATYPE_DEFINITION_BOOLEAN]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, datatype_definition_boolean_ := range req_if_content.DATATYPES.DATATYPE_DEFINITION_BOOLEAN {
					res[datatype_definition_boolean_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPES.DATATYPE_DEFINITION_DATE":
			res := make(map[*DATATYPE_DEFINITION_DATE]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, datatype_definition_date_ := range req_if_content.DATATYPES.DATATYPE_DEFINITION_DATE {
					res[datatype_definition_date_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPES.DATATYPE_DEFINITION_ENUMERATION":
			res := make(map[*DATATYPE_DEFINITION_ENUMERATION]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, datatype_definition_enumeration_ := range req_if_content.DATATYPES.DATATYPE_DEFINITION_ENUMERATION {
					res[datatype_definition_enumeration_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPES.DATATYPE_DEFINITION_INTEGER":
			res := make(map[*DATATYPE_DEFINITION_INTEGER]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, datatype_definition_integer_ := range req_if_content.DATATYPES.DATATYPE_DEFINITION_INTEGER {
					res[datatype_definition_integer_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPES.DATATYPE_DEFINITION_REAL":
			res := make(map[*DATATYPE_DEFINITION_REAL]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, datatype_definition_real_ := range req_if_content.DATATYPES.DATATYPE_DEFINITION_REAL {
					res[datatype_definition_real_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPES.DATATYPE_DEFINITION_STRING":
			res := make(map[*DATATYPE_DEFINITION_STRING]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, datatype_definition_string_ := range req_if_content.DATATYPES.DATATYPE_DEFINITION_STRING {
					res[datatype_definition_string_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPES.DATATYPE_DEFINITION_XHTML":
			res := make(map[*DATATYPE_DEFINITION_XHTML]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, datatype_definition_xhtml_ := range req_if_content.DATATYPES.DATATYPE_DEFINITION_XHTML {
					res[datatype_definition_xhtml_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_TYPES.RELATION_GROUP_TYPE":
			res := make(map[*RELATION_GROUP_TYPE]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, relation_group_type_ := range req_if_content.SPEC_TYPES.RELATION_GROUP_TYPE {
					res[relation_group_type_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_TYPES.SPEC_OBJECT_TYPE":
			res := make(map[*SPEC_OBJECT_TYPE]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, spec_object_type_ := range req_if_content.SPEC_TYPES.SPEC_OBJECT_TYPE {
					res[spec_object_type_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_TYPES.SPEC_RELATION_TYPE":
			res := make(map[*SPEC_RELATION_TYPE]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, spec_relation_type_ := range req_if_content.SPEC_TYPES.SPEC_RELATION_TYPE {
					res[spec_relation_type_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_TYPES.SPECIFICATION_TYPE":
			res := make(map[*SPECIFICATION_TYPE]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, specification_type_ := range req_if_content.SPEC_TYPES.SPECIFICATION_TYPE {
					res[specification_type_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_OBJECTS.SPEC_OBJECT":
			res := make(map[*SPEC_OBJECT]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, spec_object_ := range req_if_content.SPEC_OBJECTS.SPEC_OBJECT {
					res[spec_object_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_RELATIONS.SPEC_RELATION":
			res := make(map[*SPEC_RELATION]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, spec_relation_ := range req_if_content.SPEC_RELATIONS.SPEC_RELATION {
					res[spec_relation_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "SPECIFICATIONS.SPECIFICATION":
			res := make(map[*SPECIFICATION]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, specification_ := range req_if_content.SPECIFICATIONS.SPECIFICATION {
					res[specification_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_RELATION_GROUPS.RELATION_GROUP":
			res := make(map[*RELATION_GROUP]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, relation_group_ := range req_if_content.SPEC_RELATION_GROUPS.RELATION_GROUP {
					res[relation_group_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of REQ_IF_HEADER
	case REQ_IF_HEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_TOOL_EXTENSION
	case REQ_IF_TOOL_EXTENSION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION
	case SPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				for _, alternative_id_ := range specification.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = specification
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			res := make(map[*ATTRIBUTE_VALUE_BOOLEAN]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				for _, attribute_value_boolean_ := range specification.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
					res[attribute_value_boolean_] = specification
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_DATE":
			res := make(map[*ATTRIBUTE_VALUE_DATE]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				for _, attribute_value_date_ := range specification.VALUES.ATTRIBUTE_VALUE_DATE {
					res[attribute_value_date_] = specification
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			res := make(map[*ATTRIBUTE_VALUE_ENUMERATION]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				for _, attribute_value_enumeration_ := range specification.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
					res[attribute_value_enumeration_] = specification
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_INTEGER":
			res := make(map[*ATTRIBUTE_VALUE_INTEGER]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				for _, attribute_value_integer_ := range specification.VALUES.ATTRIBUTE_VALUE_INTEGER {
					res[attribute_value_integer_] = specification
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_REAL":
			res := make(map[*ATTRIBUTE_VALUE_REAL]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				for _, attribute_value_real_ := range specification.VALUES.ATTRIBUTE_VALUE_REAL {
					res[attribute_value_real_] = specification
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_STRING":
			res := make(map[*ATTRIBUTE_VALUE_STRING]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				for _, attribute_value_string_ := range specification.VALUES.ATTRIBUTE_VALUE_STRING {
					res[attribute_value_string_] = specification
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_XHTML":
			res := make(map[*ATTRIBUTE_VALUE_XHTML]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				for _, attribute_value_xhtml_ := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
					res[attribute_value_xhtml_] = specification
				}
			}
			return any(res).(map[*End]*Start)
		case "CHILDREN.SPEC_HIERARCHY":
			res := make(map[*SPEC_HIERARCHY]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				for _, spec_hierarchy_ := range specification.CHILDREN.SPEC_HIERARCHY {
					res[spec_hierarchy_] = specification
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPECIFICATION_TYPE
	case SPECIFICATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*SPECIFICATION_TYPE)
			for specification_type := range stage.SPECIFICATION_TYPEs {
				for _, alternative_id_ := range specification_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = specification_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			res := make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPECIFICATION_TYPE)
			for specification_type := range stage.SPECIFICATION_TYPEs {
				for _, attribute_definition_boolean_ := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
					res[attribute_definition_boolean_] = specification_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			res := make(map[*ATTRIBUTE_DEFINITION_DATE]*SPECIFICATION_TYPE)
			for specification_type := range stage.SPECIFICATION_TYPEs {
				for _, attribute_definition_date_ := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
					res[attribute_definition_date_] = specification_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			res := make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPECIFICATION_TYPE)
			for specification_type := range stage.SPECIFICATION_TYPEs {
				for _, attribute_definition_enumeration_ := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
					res[attribute_definition_enumeration_] = specification_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			res := make(map[*ATTRIBUTE_DEFINITION_INTEGER]*SPECIFICATION_TYPE)
			for specification_type := range stage.SPECIFICATION_TYPEs {
				for _, attribute_definition_integer_ := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
					res[attribute_definition_integer_] = specification_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			res := make(map[*ATTRIBUTE_DEFINITION_REAL]*SPECIFICATION_TYPE)
			for specification_type := range stage.SPECIFICATION_TYPEs {
				for _, attribute_definition_real_ := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
					res[attribute_definition_real_] = specification_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			res := make(map[*ATTRIBUTE_DEFINITION_STRING]*SPECIFICATION_TYPE)
			for specification_type := range stage.SPECIFICATION_TYPEs {
				for _, attribute_definition_string_ := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
					res[attribute_definition_string_] = specification_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			res := make(map[*ATTRIBUTE_DEFINITION_XHTML]*SPECIFICATION_TYPE)
			for specification_type := range stage.SPECIFICATION_TYPEs {
				for _, attribute_definition_xhtml_ := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
					res[attribute_definition_xhtml_] = specification_type
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPEC_HIERARCHY
	case SPEC_HIERARCHY:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*SPEC_HIERARCHY)
			for spec_hierarchy := range stage.SPEC_HIERARCHYs {
				for _, alternative_id_ := range spec_hierarchy.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = spec_hierarchy
				}
			}
			return any(res).(map[*End]*Start)
		case "CHILDREN.SPEC_HIERARCHY":
			res := make(map[*SPEC_HIERARCHY]*SPEC_HIERARCHY)
			for spec_hierarchy := range stage.SPEC_HIERARCHYs {
				for _, spec_hierarchy_ := range spec_hierarchy.CHILDREN.SPEC_HIERARCHY {
					res[spec_hierarchy_] = spec_hierarchy
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPEC_OBJECT
	case SPEC_OBJECT:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				for _, alternative_id_ := range spec_object.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = spec_object
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			res := make(map[*ATTRIBUTE_VALUE_BOOLEAN]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				for _, attribute_value_boolean_ := range spec_object.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
					res[attribute_value_boolean_] = spec_object
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_DATE":
			res := make(map[*ATTRIBUTE_VALUE_DATE]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				for _, attribute_value_date_ := range spec_object.VALUES.ATTRIBUTE_VALUE_DATE {
					res[attribute_value_date_] = spec_object
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			res := make(map[*ATTRIBUTE_VALUE_ENUMERATION]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				for _, attribute_value_enumeration_ := range spec_object.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
					res[attribute_value_enumeration_] = spec_object
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_INTEGER":
			res := make(map[*ATTRIBUTE_VALUE_INTEGER]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				for _, attribute_value_integer_ := range spec_object.VALUES.ATTRIBUTE_VALUE_INTEGER {
					res[attribute_value_integer_] = spec_object
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_REAL":
			res := make(map[*ATTRIBUTE_VALUE_REAL]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				for _, attribute_value_real_ := range spec_object.VALUES.ATTRIBUTE_VALUE_REAL {
					res[attribute_value_real_] = spec_object
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_STRING":
			res := make(map[*ATTRIBUTE_VALUE_STRING]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				for _, attribute_value_string_ := range spec_object.VALUES.ATTRIBUTE_VALUE_STRING {
					res[attribute_value_string_] = spec_object
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_XHTML":
			res := make(map[*ATTRIBUTE_VALUE_XHTML]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				for _, attribute_value_xhtml_ := range spec_object.VALUES.ATTRIBUTE_VALUE_XHTML {
					res[attribute_value_xhtml_] = spec_object
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPEC_OBJECT_TYPE
	case SPEC_OBJECT_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*SPEC_OBJECT_TYPE)
			for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
				for _, alternative_id_ := range spec_object_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = spec_object_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			res := make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPEC_OBJECT_TYPE)
			for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
				for _, attribute_definition_boolean_ := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
					res[attribute_definition_boolean_] = spec_object_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			res := make(map[*ATTRIBUTE_DEFINITION_DATE]*SPEC_OBJECT_TYPE)
			for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
				for _, attribute_definition_date_ := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
					res[attribute_definition_date_] = spec_object_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			res := make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPEC_OBJECT_TYPE)
			for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
				for _, attribute_definition_enumeration_ := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
					res[attribute_definition_enumeration_] = spec_object_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			res := make(map[*ATTRIBUTE_DEFINITION_INTEGER]*SPEC_OBJECT_TYPE)
			for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
				for _, attribute_definition_integer_ := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
					res[attribute_definition_integer_] = spec_object_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			res := make(map[*ATTRIBUTE_DEFINITION_REAL]*SPEC_OBJECT_TYPE)
			for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
				for _, attribute_definition_real_ := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
					res[attribute_definition_real_] = spec_object_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			res := make(map[*ATTRIBUTE_DEFINITION_STRING]*SPEC_OBJECT_TYPE)
			for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
				for _, attribute_definition_string_ := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
					res[attribute_definition_string_] = spec_object_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			res := make(map[*ATTRIBUTE_DEFINITION_XHTML]*SPEC_OBJECT_TYPE)
			for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
				for _, attribute_definition_xhtml_ := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
					res[attribute_definition_xhtml_] = spec_object_type
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPEC_RELATION
	case SPEC_RELATION:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				for _, alternative_id_ := range spec_relation.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = spec_relation
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			res := make(map[*ATTRIBUTE_VALUE_BOOLEAN]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				for _, attribute_value_boolean_ := range spec_relation.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
					res[attribute_value_boolean_] = spec_relation
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_DATE":
			res := make(map[*ATTRIBUTE_VALUE_DATE]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				for _, attribute_value_date_ := range spec_relation.VALUES.ATTRIBUTE_VALUE_DATE {
					res[attribute_value_date_] = spec_relation
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			res := make(map[*ATTRIBUTE_VALUE_ENUMERATION]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				for _, attribute_value_enumeration_ := range spec_relation.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
					res[attribute_value_enumeration_] = spec_relation
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_INTEGER":
			res := make(map[*ATTRIBUTE_VALUE_INTEGER]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				for _, attribute_value_integer_ := range spec_relation.VALUES.ATTRIBUTE_VALUE_INTEGER {
					res[attribute_value_integer_] = spec_relation
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_REAL":
			res := make(map[*ATTRIBUTE_VALUE_REAL]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				for _, attribute_value_real_ := range spec_relation.VALUES.ATTRIBUTE_VALUE_REAL {
					res[attribute_value_real_] = spec_relation
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_STRING":
			res := make(map[*ATTRIBUTE_VALUE_STRING]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				for _, attribute_value_string_ := range spec_relation.VALUES.ATTRIBUTE_VALUE_STRING {
					res[attribute_value_string_] = spec_relation
				}
			}
			return any(res).(map[*End]*Start)
		case "VALUES.ATTRIBUTE_VALUE_XHTML":
			res := make(map[*ATTRIBUTE_VALUE_XHTML]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				for _, attribute_value_xhtml_ := range spec_relation.VALUES.ATTRIBUTE_VALUE_XHTML {
					res[attribute_value_xhtml_] = spec_relation
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPEC_RELATION_TYPE
	case SPEC_RELATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID]*SPEC_RELATION_TYPE)
			for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
				for _, alternative_id_ := range spec_relation_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
					res[alternative_id_] = spec_relation_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			res := make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPEC_RELATION_TYPE)
			for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
				for _, attribute_definition_boolean_ := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
					res[attribute_definition_boolean_] = spec_relation_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			res := make(map[*ATTRIBUTE_DEFINITION_DATE]*SPEC_RELATION_TYPE)
			for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
				for _, attribute_definition_date_ := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
					res[attribute_definition_date_] = spec_relation_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			res := make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPEC_RELATION_TYPE)
			for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
				for _, attribute_definition_enumeration_ := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
					res[attribute_definition_enumeration_] = spec_relation_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			res := make(map[*ATTRIBUTE_DEFINITION_INTEGER]*SPEC_RELATION_TYPE)
			for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
				for _, attribute_definition_integer_ := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
					res[attribute_definition_integer_] = spec_relation_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			res := make(map[*ATTRIBUTE_DEFINITION_REAL]*SPEC_RELATION_TYPE)
			for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
				for _, attribute_definition_real_ := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
					res[attribute_definition_real_] = spec_relation_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			res := make(map[*ATTRIBUTE_DEFINITION_STRING]*SPEC_RELATION_TYPE)
			for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
				for _, attribute_definition_string_ := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
					res[attribute_definition_string_] = spec_relation_type
				}
			}
			return any(res).(map[*End]*Start)
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			res := make(map[*ATTRIBUTE_DEFINITION_XHTML]*SPEC_RELATION_TYPE)
			for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
				for _, attribute_definition_xhtml_ := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
					res[attribute_definition_xhtml_] = spec_relation_type
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of XHTML_CONTENT
	case XHTML_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_InlPres_type
	case Xhtml_InlPres_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_a_type
	case Xhtml_a_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_abbr_type
	case Xhtml_abbr_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_acronym_type
	case Xhtml_acronym_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_address_type
	case Xhtml_address_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_blockquote_type
	case Xhtml_blockquote_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_br_type
	case Xhtml_br_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_caption_type
	case Xhtml_caption_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_cite_type
	case Xhtml_cite_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_code_type
	case Xhtml_code_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_col_type
	case Xhtml_col_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_colgroup_type
	case Xhtml_colgroup_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_dd_type
	case Xhtml_dd_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_dfn_type
	case Xhtml_dfn_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_div_type
	case Xhtml_div_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_dl_type
	case Xhtml_dl_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_dt_type
	case Xhtml_dt_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_edit_type
	case Xhtml_edit_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_em_type
	case Xhtml_em_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h1_type
	case Xhtml_h1_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h2_type
	case Xhtml_h2_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h3_type
	case Xhtml_h3_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h4_type
	case Xhtml_h4_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h5_type
	case Xhtml_h5_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_h6_type
	case Xhtml_h6_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_heading_type
	case Xhtml_heading_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_hr_type
	case Xhtml_hr_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_kbd_type
	case Xhtml_kbd_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_li_type
	case Xhtml_li_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_object_type
	case Xhtml_object_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_ol_type
	case Xhtml_ol_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_p_type
	case Xhtml_p_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_param_type
	case Xhtml_param_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_pre_type
	case Xhtml_pre_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_q_type
	case Xhtml_q_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_samp_type
	case Xhtml_samp_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_span_type
	case Xhtml_span_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_strong_type
	case Xhtml_strong_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_table_type
	case Xhtml_table_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_tbody_type
	case Xhtml_tbody_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_td_type
	case Xhtml_td_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_tfoot_type
	case Xhtml_tfoot_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_th_type
	case Xhtml_th_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_thead_type
	case Xhtml_thead_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_tr_type
	case Xhtml_tr_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_ul_type
	case Xhtml_ul_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Xhtml_var_type
	case Xhtml_var_type:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case ALTERNATIVE_ID:
		res = "ALTERNATIVE_ID"
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		res = "ATTRIBUTE_DEFINITION_BOOLEAN"
	case ATTRIBUTE_DEFINITION_DATE:
		res = "ATTRIBUTE_DEFINITION_DATE"
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		res = "ATTRIBUTE_DEFINITION_ENUMERATION"
	case ATTRIBUTE_DEFINITION_INTEGER:
		res = "ATTRIBUTE_DEFINITION_INTEGER"
	case ATTRIBUTE_DEFINITION_REAL:
		res = "ATTRIBUTE_DEFINITION_REAL"
	case ATTRIBUTE_DEFINITION_STRING:
		res = "ATTRIBUTE_DEFINITION_STRING"
	case ATTRIBUTE_DEFINITION_XHTML:
		res = "ATTRIBUTE_DEFINITION_XHTML"
	case ATTRIBUTE_VALUE_BOOLEAN:
		res = "ATTRIBUTE_VALUE_BOOLEAN"
	case ATTRIBUTE_VALUE_DATE:
		res = "ATTRIBUTE_VALUE_DATE"
	case ATTRIBUTE_VALUE_ENUMERATION:
		res = "ATTRIBUTE_VALUE_ENUMERATION"
	case ATTRIBUTE_VALUE_INTEGER:
		res = "ATTRIBUTE_VALUE_INTEGER"
	case ATTRIBUTE_VALUE_REAL:
		res = "ATTRIBUTE_VALUE_REAL"
	case ATTRIBUTE_VALUE_STRING:
		res = "ATTRIBUTE_VALUE_STRING"
	case ATTRIBUTE_VALUE_XHTML:
		res = "ATTRIBUTE_VALUE_XHTML"
	case AnyType:
		res = "AnyType"
	case DATATYPE_DEFINITION_BOOLEAN:
		res = "DATATYPE_DEFINITION_BOOLEAN"
	case DATATYPE_DEFINITION_DATE:
		res = "DATATYPE_DEFINITION_DATE"
	case DATATYPE_DEFINITION_ENUMERATION:
		res = "DATATYPE_DEFINITION_ENUMERATION"
	case DATATYPE_DEFINITION_INTEGER:
		res = "DATATYPE_DEFINITION_INTEGER"
	case DATATYPE_DEFINITION_REAL:
		res = "DATATYPE_DEFINITION_REAL"
	case DATATYPE_DEFINITION_STRING:
		res = "DATATYPE_DEFINITION_STRING"
	case DATATYPE_DEFINITION_XHTML:
		res = "DATATYPE_DEFINITION_XHTML"
	case EMBEDDED_VALUE:
		res = "EMBEDDED_VALUE"
	case ENUM_VALUE:
		res = "ENUM_VALUE"
	case RELATION_GROUP:
		res = "RELATION_GROUP"
	case RELATION_GROUP_TYPE:
		res = "RELATION_GROUP_TYPE"
	case REQ_IF:
		res = "REQ_IF"
	case REQ_IF_CONTENT:
		res = "REQ_IF_CONTENT"
	case REQ_IF_HEADER:
		res = "REQ_IF_HEADER"
	case REQ_IF_TOOL_EXTENSION:
		res = "REQ_IF_TOOL_EXTENSION"
	case SPECIFICATION:
		res = "SPECIFICATION"
	case SPECIFICATION_TYPE:
		res = "SPECIFICATION_TYPE"
	case SPEC_HIERARCHY:
		res = "SPEC_HIERARCHY"
	case SPEC_OBJECT:
		res = "SPEC_OBJECT"
	case SPEC_OBJECT_TYPE:
		res = "SPEC_OBJECT_TYPE"
	case SPEC_RELATION:
		res = "SPEC_RELATION"
	case SPEC_RELATION_TYPE:
		res = "SPEC_RELATION_TYPE"
	case XHTML_CONTENT:
		res = "XHTML_CONTENT"
	case Xhtml_InlPres_type:
		res = "Xhtml_InlPres_type"
	case Xhtml_a_type:
		res = "Xhtml_a_type"
	case Xhtml_abbr_type:
		res = "Xhtml_abbr_type"
	case Xhtml_acronym_type:
		res = "Xhtml_acronym_type"
	case Xhtml_address_type:
		res = "Xhtml_address_type"
	case Xhtml_blockquote_type:
		res = "Xhtml_blockquote_type"
	case Xhtml_br_type:
		res = "Xhtml_br_type"
	case Xhtml_caption_type:
		res = "Xhtml_caption_type"
	case Xhtml_cite_type:
		res = "Xhtml_cite_type"
	case Xhtml_code_type:
		res = "Xhtml_code_type"
	case Xhtml_col_type:
		res = "Xhtml_col_type"
	case Xhtml_colgroup_type:
		res = "Xhtml_colgroup_type"
	case Xhtml_dd_type:
		res = "Xhtml_dd_type"
	case Xhtml_dfn_type:
		res = "Xhtml_dfn_type"
	case Xhtml_div_type:
		res = "Xhtml_div_type"
	case Xhtml_dl_type:
		res = "Xhtml_dl_type"
	case Xhtml_dt_type:
		res = "Xhtml_dt_type"
	case Xhtml_edit_type:
		res = "Xhtml_edit_type"
	case Xhtml_em_type:
		res = "Xhtml_em_type"
	case Xhtml_h1_type:
		res = "Xhtml_h1_type"
	case Xhtml_h2_type:
		res = "Xhtml_h2_type"
	case Xhtml_h3_type:
		res = "Xhtml_h3_type"
	case Xhtml_h4_type:
		res = "Xhtml_h4_type"
	case Xhtml_h5_type:
		res = "Xhtml_h5_type"
	case Xhtml_h6_type:
		res = "Xhtml_h6_type"
	case Xhtml_heading_type:
		res = "Xhtml_heading_type"
	case Xhtml_hr_type:
		res = "Xhtml_hr_type"
	case Xhtml_kbd_type:
		res = "Xhtml_kbd_type"
	case Xhtml_li_type:
		res = "Xhtml_li_type"
	case Xhtml_object_type:
		res = "Xhtml_object_type"
	case Xhtml_ol_type:
		res = "Xhtml_ol_type"
	case Xhtml_p_type:
		res = "Xhtml_p_type"
	case Xhtml_param_type:
		res = "Xhtml_param_type"
	case Xhtml_pre_type:
		res = "Xhtml_pre_type"
	case Xhtml_q_type:
		res = "Xhtml_q_type"
	case Xhtml_samp_type:
		res = "Xhtml_samp_type"
	case Xhtml_span_type:
		res = "Xhtml_span_type"
	case Xhtml_strong_type:
		res = "Xhtml_strong_type"
	case Xhtml_table_type:
		res = "Xhtml_table_type"
	case Xhtml_tbody_type:
		res = "Xhtml_tbody_type"
	case Xhtml_td_type:
		res = "Xhtml_td_type"
	case Xhtml_tfoot_type:
		res = "Xhtml_tfoot_type"
	case Xhtml_th_type:
		res = "Xhtml_th_type"
	case Xhtml_thead_type:
		res = "Xhtml_thead_type"
	case Xhtml_tr_type:
		res = "Xhtml_tr_type"
	case Xhtml_ul_type:
		res = "Xhtml_ul_type"
	case Xhtml_var_type:
		res = "Xhtml_var_type"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *ALTERNATIVE_ID:
		res = "ALTERNATIVE_ID"
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		res = "ATTRIBUTE_DEFINITION_BOOLEAN"
	case *ATTRIBUTE_DEFINITION_DATE:
		res = "ATTRIBUTE_DEFINITION_DATE"
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		res = "ATTRIBUTE_DEFINITION_ENUMERATION"
	case *ATTRIBUTE_DEFINITION_INTEGER:
		res = "ATTRIBUTE_DEFINITION_INTEGER"
	case *ATTRIBUTE_DEFINITION_REAL:
		res = "ATTRIBUTE_DEFINITION_REAL"
	case *ATTRIBUTE_DEFINITION_STRING:
		res = "ATTRIBUTE_DEFINITION_STRING"
	case *ATTRIBUTE_DEFINITION_XHTML:
		res = "ATTRIBUTE_DEFINITION_XHTML"
	case *ATTRIBUTE_VALUE_BOOLEAN:
		res = "ATTRIBUTE_VALUE_BOOLEAN"
	case *ATTRIBUTE_VALUE_DATE:
		res = "ATTRIBUTE_VALUE_DATE"
	case *ATTRIBUTE_VALUE_ENUMERATION:
		res = "ATTRIBUTE_VALUE_ENUMERATION"
	case *ATTRIBUTE_VALUE_INTEGER:
		res = "ATTRIBUTE_VALUE_INTEGER"
	case *ATTRIBUTE_VALUE_REAL:
		res = "ATTRIBUTE_VALUE_REAL"
	case *ATTRIBUTE_VALUE_STRING:
		res = "ATTRIBUTE_VALUE_STRING"
	case *ATTRIBUTE_VALUE_XHTML:
		res = "ATTRIBUTE_VALUE_XHTML"
	case *AnyType:
		res = "AnyType"
	case *DATATYPE_DEFINITION_BOOLEAN:
		res = "DATATYPE_DEFINITION_BOOLEAN"
	case *DATATYPE_DEFINITION_DATE:
		res = "DATATYPE_DEFINITION_DATE"
	case *DATATYPE_DEFINITION_ENUMERATION:
		res = "DATATYPE_DEFINITION_ENUMERATION"
	case *DATATYPE_DEFINITION_INTEGER:
		res = "DATATYPE_DEFINITION_INTEGER"
	case *DATATYPE_DEFINITION_REAL:
		res = "DATATYPE_DEFINITION_REAL"
	case *DATATYPE_DEFINITION_STRING:
		res = "DATATYPE_DEFINITION_STRING"
	case *DATATYPE_DEFINITION_XHTML:
		res = "DATATYPE_DEFINITION_XHTML"
	case *EMBEDDED_VALUE:
		res = "EMBEDDED_VALUE"
	case *ENUM_VALUE:
		res = "ENUM_VALUE"
	case *RELATION_GROUP:
		res = "RELATION_GROUP"
	case *RELATION_GROUP_TYPE:
		res = "RELATION_GROUP_TYPE"
	case *REQ_IF:
		res = "REQ_IF"
	case *REQ_IF_CONTENT:
		res = "REQ_IF_CONTENT"
	case *REQ_IF_HEADER:
		res = "REQ_IF_HEADER"
	case *REQ_IF_TOOL_EXTENSION:
		res = "REQ_IF_TOOL_EXTENSION"
	case *SPECIFICATION:
		res = "SPECIFICATION"
	case *SPECIFICATION_TYPE:
		res = "SPECIFICATION_TYPE"
	case *SPEC_HIERARCHY:
		res = "SPEC_HIERARCHY"
	case *SPEC_OBJECT:
		res = "SPEC_OBJECT"
	case *SPEC_OBJECT_TYPE:
		res = "SPEC_OBJECT_TYPE"
	case *SPEC_RELATION:
		res = "SPEC_RELATION"
	case *SPEC_RELATION_TYPE:
		res = "SPEC_RELATION_TYPE"
	case *XHTML_CONTENT:
		res = "XHTML_CONTENT"
	case *Xhtml_InlPres_type:
		res = "Xhtml_InlPres_type"
	case *Xhtml_a_type:
		res = "Xhtml_a_type"
	case *Xhtml_abbr_type:
		res = "Xhtml_abbr_type"
	case *Xhtml_acronym_type:
		res = "Xhtml_acronym_type"
	case *Xhtml_address_type:
		res = "Xhtml_address_type"
	case *Xhtml_blockquote_type:
		res = "Xhtml_blockquote_type"
	case *Xhtml_br_type:
		res = "Xhtml_br_type"
	case *Xhtml_caption_type:
		res = "Xhtml_caption_type"
	case *Xhtml_cite_type:
		res = "Xhtml_cite_type"
	case *Xhtml_code_type:
		res = "Xhtml_code_type"
	case *Xhtml_col_type:
		res = "Xhtml_col_type"
	case *Xhtml_colgroup_type:
		res = "Xhtml_colgroup_type"
	case *Xhtml_dd_type:
		res = "Xhtml_dd_type"
	case *Xhtml_dfn_type:
		res = "Xhtml_dfn_type"
	case *Xhtml_div_type:
		res = "Xhtml_div_type"
	case *Xhtml_dl_type:
		res = "Xhtml_dl_type"
	case *Xhtml_dt_type:
		res = "Xhtml_dt_type"
	case *Xhtml_edit_type:
		res = "Xhtml_edit_type"
	case *Xhtml_em_type:
		res = "Xhtml_em_type"
	case *Xhtml_h1_type:
		res = "Xhtml_h1_type"
	case *Xhtml_h2_type:
		res = "Xhtml_h2_type"
	case *Xhtml_h3_type:
		res = "Xhtml_h3_type"
	case *Xhtml_h4_type:
		res = "Xhtml_h4_type"
	case *Xhtml_h5_type:
		res = "Xhtml_h5_type"
	case *Xhtml_h6_type:
		res = "Xhtml_h6_type"
	case *Xhtml_heading_type:
		res = "Xhtml_heading_type"
	case *Xhtml_hr_type:
		res = "Xhtml_hr_type"
	case *Xhtml_kbd_type:
		res = "Xhtml_kbd_type"
	case *Xhtml_li_type:
		res = "Xhtml_li_type"
	case *Xhtml_object_type:
		res = "Xhtml_object_type"
	case *Xhtml_ol_type:
		res = "Xhtml_ol_type"
	case *Xhtml_p_type:
		res = "Xhtml_p_type"
	case *Xhtml_param_type:
		res = "Xhtml_param_type"
	case *Xhtml_pre_type:
		res = "Xhtml_pre_type"
	case *Xhtml_q_type:
		res = "Xhtml_q_type"
	case *Xhtml_samp_type:
		res = "Xhtml_samp_type"
	case *Xhtml_span_type:
		res = "Xhtml_span_type"
	case *Xhtml_strong_type:
		res = "Xhtml_strong_type"
	case *Xhtml_table_type:
		res = "Xhtml_table_type"
	case *Xhtml_tbody_type:
		res = "Xhtml_tbody_type"
	case *Xhtml_td_type:
		res = "Xhtml_td_type"
	case *Xhtml_tfoot_type:
		res = "Xhtml_tfoot_type"
	case *Xhtml_th_type:
		res = "Xhtml_th_type"
	case *Xhtml_thead_type:
		res = "Xhtml_thead_type"
	case *Xhtml_tr_type:
		res = "Xhtml_tr_type"
	case *Xhtml_ul_type:
		res = "Xhtml_ul_type"
	case *Xhtml_var_type:
		res = "Xhtml_var_type"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case ALTERNATIVE_ID:
		res = []string{"Name", "IDENTIFIER"}
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_DEFINITION_DATE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		res = []string{"Name", "DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "MULTI_VALUED"}
	case ATTRIBUTE_DEFINITION_INTEGER:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_DEFINITION_REAL:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_DEFINITION_STRING:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_DEFINITION_XHTML:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_VALUE_BOOLEAN:
		res = []string{"Name", "THE_VALUE"}
	case ATTRIBUTE_VALUE_DATE:
		res = []string{"Name", "THE_VALUE"}
	case ATTRIBUTE_VALUE_ENUMERATION:
		res = []string{"Name"}
	case ATTRIBUTE_VALUE_INTEGER:
		res = []string{"Name"}
	case ATTRIBUTE_VALUE_REAL:
		res = []string{"Name", "THE_VALUE"}
	case ATTRIBUTE_VALUE_STRING:
		res = []string{"Name", "THE_VALUE"}
	case ATTRIBUTE_VALUE_XHTML:
		res = []string{"Name", "THE_VALUE", "THE_ORIGINAL_VALUE", "IS_SIMPLIFIED"}
	case AnyType:
		res = []string{"Name", "InnerXML"}
	case DATATYPE_DEFINITION_BOOLEAN:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case DATATYPE_DEFINITION_DATE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case DATATYPE_DEFINITION_ENUMERATION:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "SPECIFIED_VALUES.ENUM_VALUE", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case DATATYPE_DEFINITION_INTEGER:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case DATATYPE_DEFINITION_REAL:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME", "MAX", "MIN"}
	case DATATYPE_DEFINITION_STRING:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case DATATYPE_DEFINITION_XHTML:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case EMBEDDED_VALUE:
		res = []string{"Name", "OTHER_CONTENT"}
	case ENUM_VALUE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "PROPERTIES.EMBEDDED_VALUE", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case RELATION_GROUP:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case RELATION_GROUP_TYPE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case REQ_IF:
		res = []string{"Name", "THE_HEADER.REQ_IF_HEADER", "CORE_CONTENT.REQ_IF_CONTENT", "TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION", "Lang"}
	case REQ_IF_CONTENT:
		res = []string{"Name", "DATATYPES.DATATYPE_DEFINITION_BOOLEAN", "DATATYPES.DATATYPE_DEFINITION_DATE", "DATATYPES.DATATYPE_DEFINITION_ENUMERATION", "DATATYPES.DATATYPE_DEFINITION_INTEGER", "DATATYPES.DATATYPE_DEFINITION_REAL", "DATATYPES.DATATYPE_DEFINITION_STRING", "DATATYPES.DATATYPE_DEFINITION_XHTML", "SPEC_TYPES.RELATION_GROUP_TYPE", "SPEC_TYPES.SPEC_OBJECT_TYPE", "SPEC_TYPES.SPEC_RELATION_TYPE", "SPEC_TYPES.SPECIFICATION_TYPE", "SPEC_OBJECTS.SPEC_OBJECT", "SPEC_RELATIONS.SPEC_RELATION", "SPECIFICATIONS.SPECIFICATION", "SPEC_RELATION_GROUPS.RELATION_GROUP"}
	case REQ_IF_HEADER:
		res = []string{"Name", "COMMENT", "CREATION_TIME", "REPOSITORY_ID", "REQ_IF_TOOL_ID", "REQ_IF_VERSION", "SOURCE_TOOL_ID", "TITLE"}
	case REQ_IF_TOOL_EXTENSION:
		res = []string{"Name"}
	case SPECIFICATION:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "VALUES.ATTRIBUTE_VALUE_BOOLEAN", "VALUES.ATTRIBUTE_VALUE_DATE", "VALUES.ATTRIBUTE_VALUE_ENUMERATION", "VALUES.ATTRIBUTE_VALUE_INTEGER", "VALUES.ATTRIBUTE_VALUE_REAL", "VALUES.ATTRIBUTE_VALUE_STRING", "VALUES.ATTRIBUTE_VALUE_XHTML", "CHILDREN.SPEC_HIERARCHY", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case SPECIFICATION_TYPE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_HIERARCHY:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "CHILDREN.SPEC_HIERARCHY", "DESC", "IS_EDITABLE", "IS_TABLE_INTERNAL", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_OBJECT:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "VALUES.ATTRIBUTE_VALUE_BOOLEAN", "VALUES.ATTRIBUTE_VALUE_DATE", "VALUES.ATTRIBUTE_VALUE_ENUMERATION", "VALUES.ATTRIBUTE_VALUE_INTEGER", "VALUES.ATTRIBUTE_VALUE_REAL", "VALUES.ATTRIBUTE_VALUE_STRING", "VALUES.ATTRIBUTE_VALUE_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_OBJECT_TYPE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_RELATION:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "VALUES.ATTRIBUTE_VALUE_BOOLEAN", "VALUES.ATTRIBUTE_VALUE_DATE", "VALUES.ATTRIBUTE_VALUE_ENUMERATION", "VALUES.ATTRIBUTE_VALUE_INTEGER", "VALUES.ATTRIBUTE_VALUE_REAL", "VALUES.ATTRIBUTE_VALUE_STRING", "VALUES.ATTRIBUTE_VALUE_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_RELATION_TYPE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case XHTML_CONTENT:
		res = []string{"Name"}
	case Xhtml_InlPres_type:
		res = []string{"Name"}
	case Xhtml_a_type:
		res = []string{"Name"}
	case Xhtml_abbr_type:
		res = []string{"Name"}
	case Xhtml_acronym_type:
		res = []string{"Name"}
	case Xhtml_address_type:
		res = []string{"Name"}
	case Xhtml_blockquote_type:
		res = []string{"Name"}
	case Xhtml_br_type:
		res = []string{"Name"}
	case Xhtml_caption_type:
		res = []string{"Name"}
	case Xhtml_cite_type:
		res = []string{"Name"}
	case Xhtml_code_type:
		res = []string{"Name"}
	case Xhtml_col_type:
		res = []string{"Name"}
	case Xhtml_colgroup_type:
		res = []string{"Name"}
	case Xhtml_dd_type:
		res = []string{"Name"}
	case Xhtml_dfn_type:
		res = []string{"Name"}
	case Xhtml_div_type:
		res = []string{"Name"}
	case Xhtml_dl_type:
		res = []string{"Name"}
	case Xhtml_dt_type:
		res = []string{"Name"}
	case Xhtml_edit_type:
		res = []string{"Name"}
	case Xhtml_em_type:
		res = []string{"Name"}
	case Xhtml_h1_type:
		res = []string{"Name"}
	case Xhtml_h2_type:
		res = []string{"Name"}
	case Xhtml_h3_type:
		res = []string{"Name"}
	case Xhtml_h4_type:
		res = []string{"Name"}
	case Xhtml_h5_type:
		res = []string{"Name"}
	case Xhtml_h6_type:
		res = []string{"Name"}
	case Xhtml_heading_type:
		res = []string{"Name"}
	case Xhtml_hr_type:
		res = []string{"Name"}
	case Xhtml_kbd_type:
		res = []string{"Name"}
	case Xhtml_li_type:
		res = []string{"Name"}
	case Xhtml_object_type:
		res = []string{"Name"}
	case Xhtml_ol_type:
		res = []string{"Name"}
	case Xhtml_p_type:
		res = []string{"Name"}
	case Xhtml_param_type:
		res = []string{"Name"}
	case Xhtml_pre_type:
		res = []string{"Name"}
	case Xhtml_q_type:
		res = []string{"Name"}
	case Xhtml_samp_type:
		res = []string{"Name"}
	case Xhtml_span_type:
		res = []string{"Name"}
	case Xhtml_strong_type:
		res = []string{"Name"}
	case Xhtml_table_type:
		res = []string{"Name"}
	case Xhtml_tbody_type:
		res = []string{"Name"}
	case Xhtml_td_type:
		res = []string{"Name"}
	case Xhtml_tfoot_type:
		res = []string{"Name"}
	case Xhtml_th_type:
		res = []string{"Name"}
	case Xhtml_thead_type:
		res = []string{"Name"}
	case Xhtml_tr_type:
		res = []string{"Name"}
	case Xhtml_ul_type:
		res = []string{"Name"}
	case Xhtml_var_type:
		res = []string{"Name"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case ALTERNATIVE_ID:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "DATATYPE_DEFINITION_BOOLEAN"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "DATATYPE_DEFINITION_DATE"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "DATATYPE_DEFINITION_ENUMERATION"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "DATATYPE_DEFINITION_INTEGER"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "DATATYPE_DEFINITION_REAL"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "DATATYPE_DEFINITION_STRING"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "DATATYPE_DEFINITION_XHTML"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "ENUM_VALUE"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "RELATION_GROUP"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "RELATION_GROUP_TYPE"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION_TYPE"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "SPEC_HIERARCHY"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT_TYPE"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION_TYPE"
		rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
		res = append(res, rf)
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RELATION_GROUP_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
		res = append(res, rf)
	case ATTRIBUTE_DEFINITION_DATE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RELATION_GROUP_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
		res = append(res, rf)
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RELATION_GROUP_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
		res = append(res, rf)
	case ATTRIBUTE_DEFINITION_INTEGER:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RELATION_GROUP_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
		res = append(res, rf)
	case ATTRIBUTE_DEFINITION_REAL:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RELATION_GROUP_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
		res = append(res, rf)
	case ATTRIBUTE_DEFINITION_STRING:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RELATION_GROUP_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
		res = append(res, rf)
	case ATTRIBUTE_DEFINITION_XHTML:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RELATION_GROUP_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION_TYPE"
		rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
		res = append(res, rf)
	case ATTRIBUTE_VALUE_BOOLEAN:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
		rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_BOOLEAN"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_BOOLEAN"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_BOOLEAN"
		res = append(res, rf)
	case ATTRIBUTE_VALUE_DATE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
		rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_DATE"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_DATE"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_DATE"
		res = append(res, rf)
	case ATTRIBUTE_VALUE_ENUMERATION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
		rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_ENUMERATION"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_ENUMERATION"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_ENUMERATION"
		res = append(res, rf)
	case ATTRIBUTE_VALUE_INTEGER:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
		rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_INTEGER"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_INTEGER"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_INTEGER"
		res = append(res, rf)
	case ATTRIBUTE_VALUE_REAL:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
		rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_REAL"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_REAL"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_REAL"
		res = append(res, rf)
	case ATTRIBUTE_VALUE_STRING:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
		rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_STRING"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_STRING"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_STRING"
		res = append(res, rf)
	case ATTRIBUTE_VALUE_XHTML:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
		rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML"
		res = append(res, rf)
		rf.GongstructName = "SPECIFICATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_XHTML"
		res = append(res, rf)
		rf.GongstructName = "SPEC_OBJECT"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_XHTML"
		res = append(res, rf)
		rf.GongstructName = "SPEC_RELATION"
		rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_XHTML"
		res = append(res, rf)
	case AnyType:
		var rf ReverseField
		_ = rf
	case DATATYPE_DEFINITION_BOOLEAN:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_BOOLEAN"
		res = append(res, rf)
	case DATATYPE_DEFINITION_DATE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_DATE"
		res = append(res, rf)
	case DATATYPE_DEFINITION_ENUMERATION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_ENUMERATION"
		res = append(res, rf)
	case DATATYPE_DEFINITION_INTEGER:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_INTEGER"
		res = append(res, rf)
	case DATATYPE_DEFINITION_REAL:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_REAL"
		res = append(res, rf)
	case DATATYPE_DEFINITION_STRING:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_STRING"
		res = append(res, rf)
	case DATATYPE_DEFINITION_XHTML:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_XHTML"
		res = append(res, rf)
	case EMBEDDED_VALUE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ENUM_VALUE"
		rf.Fieldname = "PROPERTIES.EMBEDDED_VALUE"
		res = append(res, rf)
	case ENUM_VALUE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DATATYPE_DEFINITION_ENUMERATION"
		rf.Fieldname = "SPECIFIED_VALUES.ENUM_VALUE"
		res = append(res, rf)
	case RELATION_GROUP:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPEC_RELATION_GROUPS.RELATION_GROUP"
		res = append(res, rf)
	case RELATION_GROUP_TYPE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPEC_TYPES.RELATION_GROUP_TYPE"
		res = append(res, rf)
	case REQ_IF:
		var rf ReverseField
		_ = rf
	case REQ_IF_CONTENT:
		var rf ReverseField
		_ = rf
	case REQ_IF_HEADER:
		var rf ReverseField
		_ = rf
	case REQ_IF_TOOL_EXTENSION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF"
		rf.Fieldname = "TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION"
		res = append(res, rf)
	case SPECIFICATION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPECIFICATIONS.SPECIFICATION"
		res = append(res, rf)
	case SPECIFICATION_TYPE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPEC_TYPES.SPECIFICATION_TYPE"
		res = append(res, rf)
	case SPEC_HIERARCHY:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECIFICATION"
		rf.Fieldname = "CHILDREN.SPEC_HIERARCHY"
		res = append(res, rf)
		rf.GongstructName = "SPEC_HIERARCHY"
		rf.Fieldname = "CHILDREN.SPEC_HIERARCHY"
		res = append(res, rf)
	case SPEC_OBJECT:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPEC_OBJECTS.SPEC_OBJECT"
		res = append(res, rf)
	case SPEC_OBJECT_TYPE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPEC_TYPES.SPEC_OBJECT_TYPE"
		res = append(res, rf)
	case SPEC_RELATION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPEC_RELATIONS.SPEC_RELATION"
		res = append(res, rf)
	case SPEC_RELATION_TYPE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPEC_TYPES.SPEC_RELATION_TYPE"
		res = append(res, rf)
	case XHTML_CONTENT:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
		rf.Fieldname = "THE_VALUE"
		res = append(res, rf)
		rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
		rf.Fieldname = "THE_ORIGINAL_VALUE"
		res = append(res, rf)
	case Xhtml_InlPres_type:
		var rf ReverseField
		_ = rf
	case Xhtml_a_type:
		var rf ReverseField
		_ = rf
	case Xhtml_abbr_type:
		var rf ReverseField
		_ = rf
	case Xhtml_acronym_type:
		var rf ReverseField
		_ = rf
	case Xhtml_address_type:
		var rf ReverseField
		_ = rf
	case Xhtml_blockquote_type:
		var rf ReverseField
		_ = rf
	case Xhtml_br_type:
		var rf ReverseField
		_ = rf
	case Xhtml_caption_type:
		var rf ReverseField
		_ = rf
	case Xhtml_cite_type:
		var rf ReverseField
		_ = rf
	case Xhtml_code_type:
		var rf ReverseField
		_ = rf
	case Xhtml_col_type:
		var rf ReverseField
		_ = rf
	case Xhtml_colgroup_type:
		var rf ReverseField
		_ = rf
	case Xhtml_dd_type:
		var rf ReverseField
		_ = rf
	case Xhtml_dfn_type:
		var rf ReverseField
		_ = rf
	case Xhtml_div_type:
		var rf ReverseField
		_ = rf
	case Xhtml_dl_type:
		var rf ReverseField
		_ = rf
	case Xhtml_dt_type:
		var rf ReverseField
		_ = rf
	case Xhtml_edit_type:
		var rf ReverseField
		_ = rf
	case Xhtml_em_type:
		var rf ReverseField
		_ = rf
	case Xhtml_h1_type:
		var rf ReverseField
		_ = rf
	case Xhtml_h2_type:
		var rf ReverseField
		_ = rf
	case Xhtml_h3_type:
		var rf ReverseField
		_ = rf
	case Xhtml_h4_type:
		var rf ReverseField
		_ = rf
	case Xhtml_h5_type:
		var rf ReverseField
		_ = rf
	case Xhtml_h6_type:
		var rf ReverseField
		_ = rf
	case Xhtml_heading_type:
		var rf ReverseField
		_ = rf
	case Xhtml_hr_type:
		var rf ReverseField
		_ = rf
	case Xhtml_kbd_type:
		var rf ReverseField
		_ = rf
	case Xhtml_li_type:
		var rf ReverseField
		_ = rf
	case Xhtml_object_type:
		var rf ReverseField
		_ = rf
	case Xhtml_ol_type:
		var rf ReverseField
		_ = rf
	case Xhtml_p_type:
		var rf ReverseField
		_ = rf
	case Xhtml_param_type:
		var rf ReverseField
		_ = rf
	case Xhtml_pre_type:
		var rf ReverseField
		_ = rf
	case Xhtml_q_type:
		var rf ReverseField
		_ = rf
	case Xhtml_samp_type:
		var rf ReverseField
		_ = rf
	case Xhtml_span_type:
		var rf ReverseField
		_ = rf
	case Xhtml_strong_type:
		var rf ReverseField
		_ = rf
	case Xhtml_table_type:
		var rf ReverseField
		_ = rf
	case Xhtml_tbody_type:
		var rf ReverseField
		_ = rf
	case Xhtml_td_type:
		var rf ReverseField
		_ = rf
	case Xhtml_tfoot_type:
		var rf ReverseField
		_ = rf
	case Xhtml_th_type:
		var rf ReverseField
		_ = rf
	case Xhtml_thead_type:
		var rf ReverseField
		_ = rf
	case Xhtml_tr_type:
		var rf ReverseField
		_ = rf
	case Xhtml_ul_type:
		var rf ReverseField
		_ = rf
	case Xhtml_var_type:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *ALTERNATIVE_ID:
		res = []string{"Name", "IDENTIFIER"}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_DEFINITION_DATE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		res = []string{"Name", "DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "MULTI_VALUED"}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_DEFINITION_REAL:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_DEFINITION_STRING:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_DEFINITION_XHTML:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML", "DESC", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		res = []string{"Name", "THE_VALUE"}
	case *ATTRIBUTE_VALUE_DATE:
		res = []string{"Name", "THE_VALUE"}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		res = []string{"Name"}
	case *ATTRIBUTE_VALUE_INTEGER:
		res = []string{"Name"}
	case *ATTRIBUTE_VALUE_REAL:
		res = []string{"Name", "THE_VALUE"}
	case *ATTRIBUTE_VALUE_STRING:
		res = []string{"Name", "THE_VALUE"}
	case *ATTRIBUTE_VALUE_XHTML:
		res = []string{"Name", "THE_VALUE", "THE_ORIGINAL_VALUE", "IS_SIMPLIFIED"}
	case *AnyType:
		res = []string{"Name", "InnerXML"}
	case *DATATYPE_DEFINITION_BOOLEAN:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *DATATYPE_DEFINITION_DATE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *DATATYPE_DEFINITION_ENUMERATION:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "SPECIFIED_VALUES.ENUM_VALUE", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *DATATYPE_DEFINITION_INTEGER:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *DATATYPE_DEFINITION_REAL:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME", "MAX", "MIN"}
	case *DATATYPE_DEFINITION_STRING:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *DATATYPE_DEFINITION_XHTML:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *EMBEDDED_VALUE:
		res = []string{"Name", "OTHER_CONTENT"}
	case *ENUM_VALUE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "PROPERTIES.EMBEDDED_VALUE", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *RELATION_GROUP:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *RELATION_GROUP_TYPE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *REQ_IF:
		res = []string{"Name", "THE_HEADER.REQ_IF_HEADER", "CORE_CONTENT.REQ_IF_CONTENT", "TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION", "Lang"}
	case *REQ_IF_CONTENT:
		res = []string{"Name", "DATATYPES.DATATYPE_DEFINITION_BOOLEAN", "DATATYPES.DATATYPE_DEFINITION_DATE", "DATATYPES.DATATYPE_DEFINITION_ENUMERATION", "DATATYPES.DATATYPE_DEFINITION_INTEGER", "DATATYPES.DATATYPE_DEFINITION_REAL", "DATATYPES.DATATYPE_DEFINITION_STRING", "DATATYPES.DATATYPE_DEFINITION_XHTML", "SPEC_TYPES.RELATION_GROUP_TYPE", "SPEC_TYPES.SPEC_OBJECT_TYPE", "SPEC_TYPES.SPEC_RELATION_TYPE", "SPEC_TYPES.SPECIFICATION_TYPE", "SPEC_OBJECTS.SPEC_OBJECT", "SPEC_RELATIONS.SPEC_RELATION", "SPECIFICATIONS.SPECIFICATION", "SPEC_RELATION_GROUPS.RELATION_GROUP"}
	case *REQ_IF_HEADER:
		res = []string{"Name", "COMMENT", "CREATION_TIME", "REPOSITORY_ID", "REQ_IF_TOOL_ID", "REQ_IF_VERSION", "SOURCE_TOOL_ID", "TITLE"}
	case *REQ_IF_TOOL_EXTENSION:
		res = []string{"Name"}
	case *SPECIFICATION:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "VALUES.ATTRIBUTE_VALUE_BOOLEAN", "VALUES.ATTRIBUTE_VALUE_DATE", "VALUES.ATTRIBUTE_VALUE_ENUMERATION", "VALUES.ATTRIBUTE_VALUE_INTEGER", "VALUES.ATTRIBUTE_VALUE_REAL", "VALUES.ATTRIBUTE_VALUE_STRING", "VALUES.ATTRIBUTE_VALUE_XHTML", "CHILDREN.SPEC_HIERARCHY", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *SPECIFICATION_TYPE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_HIERARCHY:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "CHILDREN.SPEC_HIERARCHY", "DESC", "IS_EDITABLE", "IS_TABLE_INTERNAL", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_OBJECT:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "VALUES.ATTRIBUTE_VALUE_BOOLEAN", "VALUES.ATTRIBUTE_VALUE_DATE", "VALUES.ATTRIBUTE_VALUE_ENUMERATION", "VALUES.ATTRIBUTE_VALUE_INTEGER", "VALUES.ATTRIBUTE_VALUE_REAL", "VALUES.ATTRIBUTE_VALUE_STRING", "VALUES.ATTRIBUTE_VALUE_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_OBJECT_TYPE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_RELATION:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "VALUES.ATTRIBUTE_VALUE_BOOLEAN", "VALUES.ATTRIBUTE_VALUE_DATE", "VALUES.ATTRIBUTE_VALUE_ENUMERATION", "VALUES.ATTRIBUTE_VALUE_INTEGER", "VALUES.ATTRIBUTE_VALUE_REAL", "VALUES.ATTRIBUTE_VALUE_STRING", "VALUES.ATTRIBUTE_VALUE_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_RELATION_TYPE:
		res = []string{"Name", "ALTERNATIVE_ID.ALTERNATIVE_ID", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML", "DESC", "LAST_CHANGE", "LONG_NAME"}
	case *XHTML_CONTENT:
		res = []string{"Name"}
	case *Xhtml_InlPres_type:
		res = []string{"Name"}
	case *Xhtml_a_type:
		res = []string{"Name"}
	case *Xhtml_abbr_type:
		res = []string{"Name"}
	case *Xhtml_acronym_type:
		res = []string{"Name"}
	case *Xhtml_address_type:
		res = []string{"Name"}
	case *Xhtml_blockquote_type:
		res = []string{"Name"}
	case *Xhtml_br_type:
		res = []string{"Name"}
	case *Xhtml_caption_type:
		res = []string{"Name"}
	case *Xhtml_cite_type:
		res = []string{"Name"}
	case *Xhtml_code_type:
		res = []string{"Name"}
	case *Xhtml_col_type:
		res = []string{"Name"}
	case *Xhtml_colgroup_type:
		res = []string{"Name"}
	case *Xhtml_dd_type:
		res = []string{"Name"}
	case *Xhtml_dfn_type:
		res = []string{"Name"}
	case *Xhtml_div_type:
		res = []string{"Name"}
	case *Xhtml_dl_type:
		res = []string{"Name"}
	case *Xhtml_dt_type:
		res = []string{"Name"}
	case *Xhtml_edit_type:
		res = []string{"Name"}
	case *Xhtml_em_type:
		res = []string{"Name"}
	case *Xhtml_h1_type:
		res = []string{"Name"}
	case *Xhtml_h2_type:
		res = []string{"Name"}
	case *Xhtml_h3_type:
		res = []string{"Name"}
	case *Xhtml_h4_type:
		res = []string{"Name"}
	case *Xhtml_h5_type:
		res = []string{"Name"}
	case *Xhtml_h6_type:
		res = []string{"Name"}
	case *Xhtml_heading_type:
		res = []string{"Name"}
	case *Xhtml_hr_type:
		res = []string{"Name"}
	case *Xhtml_kbd_type:
		res = []string{"Name"}
	case *Xhtml_li_type:
		res = []string{"Name"}
	case *Xhtml_object_type:
		res = []string{"Name"}
	case *Xhtml_ol_type:
		res = []string{"Name"}
	case *Xhtml_p_type:
		res = []string{"Name"}
	case *Xhtml_param_type:
		res = []string{"Name"}
	case *Xhtml_pre_type:
		res = []string{"Name"}
	case *Xhtml_q_type:
		res = []string{"Name"}
	case *Xhtml_samp_type:
		res = []string{"Name"}
	case *Xhtml_span_type:
		res = []string{"Name"}
	case *Xhtml_strong_type:
		res = []string{"Name"}
	case *Xhtml_table_type:
		res = []string{"Name"}
	case *Xhtml_tbody_type:
		res = []string{"Name"}
	case *Xhtml_td_type:
		res = []string{"Name"}
	case *Xhtml_tfoot_type:
		res = []string{"Name"}
	case *Xhtml_th_type:
		res = []string{"Name"}
	case *Xhtml_thead_type:
		res = []string{"Name"}
	case *Xhtml_tr_type:
		res = []string{"Name"}
	case *Xhtml_ul_type:
		res = []string{"Name"}
	case *Xhtml_var_type:
		res = []string{"Name"}
	}
	return
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt    GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat  GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool   GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeOthers GongFieldValueType = "GongFieldValueTypeOthers"
)

type GongFieldValue struct {
	valueString string
	GongFieldValueType
	valueInt   int
	valueFloat float64
	valueBool  bool
}

func (gongValueField *GongFieldValue) GetValueString() string {
	return gongValueField.valueString
}

func (gongValueField *GongFieldValue) GetValueInt() int {
	return gongValueField.valueInt
}

func (gongValueField *GongFieldValue) GetValueFloat() float64 {
	return gongValueField.valueFloat
}

func (gongValueField *GongFieldValue) GetValueBool() bool {
	return gongValueField.valueBool
}

func GetFieldStringValueFromPointer(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *ALTERNATIVE_ID:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IDENTIFIER":
			res.valueString = inferedInstance.IDENTIFIER
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		case "MULTI_VALUED":
			res.valueString = fmt.Sprintf("%t", inferedInstance.MULTI_VALUED)
			res.valueBool = inferedInstance.MULTI_VALUED
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_VALUE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.THE_VALUE)
			res.valueBool = inferedInstance.THE_VALUE
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case *ATTRIBUTE_VALUE_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_VALUE":
			res.valueString = inferedInstance.THE_VALUE.String()
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *ATTRIBUTE_VALUE_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_VALUE":
			res.valueString = fmt.Sprintf("%f", inferedInstance.THE_VALUE)
			res.valueFloat = inferedInstance.THE_VALUE
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case *ATTRIBUTE_VALUE_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_VALUE":
			res.valueString = inferedInstance.THE_VALUE
		}
	case *ATTRIBUTE_VALUE_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_VALUE":
			for idx, __instance__ := range inferedInstance.THE_VALUE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "THE_ORIGINAL_VALUE":
			for idx, __instance__ := range inferedInstance.THE_ORIGINAL_VALUE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "IS_SIMPLIFIED":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_SIMPLIFIED)
			res.valueBool = inferedInstance.IS_SIMPLIFIED
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case *AnyType:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "InnerXML":
			res.valueString = inferedInstance.InnerXML
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *DATATYPE_DEFINITION_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPECIFIED_VALUES.ENUM_VALUE":
			for idx, __instance__ := range inferedInstance.SPECIFIED_VALUES.ENUM_VALUE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *DATATYPE_DEFINITION_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *DATATYPE_DEFINITION_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		case "MAX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MAX)
			res.valueFloat = inferedInstance.MAX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "MIN":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MIN)
			res.valueFloat = inferedInstance.MIN
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case *DATATYPE_DEFINITION_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *DATATYPE_DEFINITION_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *EMBEDDED_VALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "OTHER_CONTENT":
			res.valueString = inferedInstance.OTHER_CONTENT
		}
	case *ENUM_VALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "PROPERTIES.EMBEDDED_VALUE":
			for idx, __instance__ := range inferedInstance.PROPERTIES.EMBEDDED_VALUE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *RELATION_GROUP:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *RELATION_GROUP_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *REQ_IF:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_HEADER.REQ_IF_HEADER":
			if inferedInstance.THE_HEADER.REQ_IF_HEADER != nil {
				res.valueString = inferedInstance.THE_HEADER.REQ_IF_HEADER.Name
			}
		case "CORE_CONTENT.REQ_IF_CONTENT":
			if inferedInstance.CORE_CONTENT.REQ_IF_CONTENT != nil {
				res.valueString = inferedInstance.CORE_CONTENT.REQ_IF_CONTENT.Name
			}
		case "TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION":
			for idx, __instance__ := range inferedInstance.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Lang":
			res.valueString = inferedInstance.Lang
		}
	case *REQ_IF_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "DATATYPES.DATATYPE_DEFINITION_BOOLEAN":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_DATE":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_ENUMERATION":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_INTEGER":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_REAL":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_STRING":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_XHTML":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_TYPES.RELATION_GROUP_TYPE":
			for idx, __instance__ := range inferedInstance.SPEC_TYPES.RELATION_GROUP_TYPE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_TYPES.SPEC_OBJECT_TYPE":
			for idx, __instance__ := range inferedInstance.SPEC_TYPES.SPEC_OBJECT_TYPE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_TYPES.SPEC_RELATION_TYPE":
			for idx, __instance__ := range inferedInstance.SPEC_TYPES.SPEC_RELATION_TYPE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_TYPES.SPECIFICATION_TYPE":
			for idx, __instance__ := range inferedInstance.SPEC_TYPES.SPECIFICATION_TYPE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_OBJECTS.SPEC_OBJECT":
			for idx, __instance__ := range inferedInstance.SPEC_OBJECTS.SPEC_OBJECT {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_RELATIONS.SPEC_RELATION":
			for idx, __instance__ := range inferedInstance.SPEC_RELATIONS.SPEC_RELATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPECIFICATIONS.SPECIFICATION":
			for idx, __instance__ := range inferedInstance.SPECIFICATIONS.SPECIFICATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_RELATION_GROUPS.RELATION_GROUP":
			for idx, __instance__ := range inferedInstance.SPEC_RELATION_GROUPS.RELATION_GROUP {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *REQ_IF_HEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "COMMENT":
			res.valueString = inferedInstance.COMMENT
		case "CREATION_TIME":
			res.valueString = inferedInstance.CREATION_TIME.String()
		case "REPOSITORY_ID":
			res.valueString = inferedInstance.REPOSITORY_ID
		case "REQ_IF_TOOL_ID":
			res.valueString = inferedInstance.REQ_IF_TOOL_ID
		case "REQ_IF_VERSION":
			res.valueString = inferedInstance.REQ_IF_VERSION
		case "SOURCE_TOOL_ID":
			res.valueString = inferedInstance.SOURCE_TOOL_ID
		case "TITLE":
			res.valueString = inferedInstance.TITLE
		}
	case *REQ_IF_TOOL_EXTENSION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *SPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_DATE":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_INTEGER":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_REAL":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_STRING":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_XHTML":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "CHILDREN.SPEC_HIERARCHY":
			for idx, __instance__ := range inferedInstance.CHILDREN.SPEC_HIERARCHY {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *SPECIFICATION_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *SPEC_HIERARCHY:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "CHILDREN.SPEC_HIERARCHY":
			for idx, __instance__ := range inferedInstance.CHILDREN.SPEC_HIERARCHY {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IS_TABLE_INTERNAL":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_TABLE_INTERNAL)
			res.valueBool = inferedInstance.IS_TABLE_INTERNAL
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *SPEC_OBJECT:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_DATE":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_INTEGER":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_REAL":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_STRING":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_XHTML":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *SPEC_OBJECT_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *SPEC_RELATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_DATE":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_INTEGER":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_REAL":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_STRING":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_XHTML":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *SPEC_RELATION_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case *XHTML_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_InlPres_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_a_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_abbr_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_acronym_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_address_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_blockquote_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_br_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_caption_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_cite_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_code_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_col_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_colgroup_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_dd_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_dfn_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_div_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_dl_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_dt_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_edit_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_em_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_h1_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_h2_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_h3_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_h4_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_h5_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_h6_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_heading_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_hr_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_kbd_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_li_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_object_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_ol_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_p_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_param_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_pre_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_q_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_samp_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_span_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_strong_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_table_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_tbody_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_td_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_tfoot_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_th_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_thead_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_tr_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_ul_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Xhtml_var_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case ALTERNATIVE_ID:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IDENTIFIER":
			res.valueString = inferedInstance.IDENTIFIER
		}
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_DEFINITION_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		case "MULTI_VALUED":
			res.valueString = fmt.Sprintf("%t", inferedInstance.MULTI_VALUED)
			res.valueBool = inferedInstance.MULTI_VALUED
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case ATTRIBUTE_DEFINITION_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_DEFINITION_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_DEFINITION_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_DEFINITION_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML":
			for idx, __instance__ := range inferedInstance.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_VALUE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.THE_VALUE)
			res.valueBool = inferedInstance.THE_VALUE
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case ATTRIBUTE_VALUE_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_VALUE":
			res.valueString = inferedInstance.THE_VALUE.String()
		}
	case ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case ATTRIBUTE_VALUE_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case ATTRIBUTE_VALUE_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_VALUE":
			res.valueString = fmt.Sprintf("%f", inferedInstance.THE_VALUE)
			res.valueFloat = inferedInstance.THE_VALUE
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case ATTRIBUTE_VALUE_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_VALUE":
			res.valueString = inferedInstance.THE_VALUE
		}
	case ATTRIBUTE_VALUE_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_VALUE":
			for idx, __instance__ := range inferedInstance.THE_VALUE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "THE_ORIGINAL_VALUE":
			for idx, __instance__ := range inferedInstance.THE_ORIGINAL_VALUE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "IS_SIMPLIFIED":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_SIMPLIFIED)
			res.valueBool = inferedInstance.IS_SIMPLIFIED
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case AnyType:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "InnerXML":
			res.valueString = inferedInstance.InnerXML
		}
	case DATATYPE_DEFINITION_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case DATATYPE_DEFINITION_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case DATATYPE_DEFINITION_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPECIFIED_VALUES.ENUM_VALUE":
			for idx, __instance__ := range inferedInstance.SPECIFIED_VALUES.ENUM_VALUE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case DATATYPE_DEFINITION_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case DATATYPE_DEFINITION_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		case "MAX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MAX)
			res.valueFloat = inferedInstance.MAX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "MIN":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MIN)
			res.valueFloat = inferedInstance.MIN
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case DATATYPE_DEFINITION_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case DATATYPE_DEFINITION_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case EMBEDDED_VALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "OTHER_CONTENT":
			res.valueString = inferedInstance.OTHER_CONTENT
		}
	case ENUM_VALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "PROPERTIES.EMBEDDED_VALUE":
			for idx, __instance__ := range inferedInstance.PROPERTIES.EMBEDDED_VALUE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case RELATION_GROUP:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case RELATION_GROUP_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case REQ_IF:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "THE_HEADER.REQ_IF_HEADER":
			if inferedInstance.THE_HEADER.REQ_IF_HEADER != nil {
				res.valueString = inferedInstance.THE_HEADER.REQ_IF_HEADER.Name
			}
		case "CORE_CONTENT.REQ_IF_CONTENT":
			if inferedInstance.CORE_CONTENT.REQ_IF_CONTENT != nil {
				res.valueString = inferedInstance.CORE_CONTENT.REQ_IF_CONTENT.Name
			}
		case "TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION":
			for idx, __instance__ := range inferedInstance.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Lang":
			res.valueString = inferedInstance.Lang
		}
	case REQ_IF_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "DATATYPES.DATATYPE_DEFINITION_BOOLEAN":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_DATE":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_ENUMERATION":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_INTEGER":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_REAL":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_STRING":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DATATYPES.DATATYPE_DEFINITION_XHTML":
			for idx, __instance__ := range inferedInstance.DATATYPES.DATATYPE_DEFINITION_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_TYPES.RELATION_GROUP_TYPE":
			for idx, __instance__ := range inferedInstance.SPEC_TYPES.RELATION_GROUP_TYPE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_TYPES.SPEC_OBJECT_TYPE":
			for idx, __instance__ := range inferedInstance.SPEC_TYPES.SPEC_OBJECT_TYPE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_TYPES.SPEC_RELATION_TYPE":
			for idx, __instance__ := range inferedInstance.SPEC_TYPES.SPEC_RELATION_TYPE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_TYPES.SPECIFICATION_TYPE":
			for idx, __instance__ := range inferedInstance.SPEC_TYPES.SPECIFICATION_TYPE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_OBJECTS.SPEC_OBJECT":
			for idx, __instance__ := range inferedInstance.SPEC_OBJECTS.SPEC_OBJECT {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_RELATIONS.SPEC_RELATION":
			for idx, __instance__ := range inferedInstance.SPEC_RELATIONS.SPEC_RELATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPECIFICATIONS.SPECIFICATION":
			for idx, __instance__ := range inferedInstance.SPECIFICATIONS.SPECIFICATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_RELATION_GROUPS.RELATION_GROUP":
			for idx, __instance__ := range inferedInstance.SPEC_RELATION_GROUPS.RELATION_GROUP {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case REQ_IF_HEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "COMMENT":
			res.valueString = inferedInstance.COMMENT
		case "CREATION_TIME":
			res.valueString = inferedInstance.CREATION_TIME.String()
		case "REPOSITORY_ID":
			res.valueString = inferedInstance.REPOSITORY_ID
		case "REQ_IF_TOOL_ID":
			res.valueString = inferedInstance.REQ_IF_TOOL_ID
		case "REQ_IF_VERSION":
			res.valueString = inferedInstance.REQ_IF_VERSION
		case "SOURCE_TOOL_ID":
			res.valueString = inferedInstance.SOURCE_TOOL_ID
		case "TITLE":
			res.valueString = inferedInstance.TITLE
		}
	case REQ_IF_TOOL_EXTENSION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case SPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_DATE":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_INTEGER":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_REAL":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_STRING":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_XHTML":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "CHILDREN.SPEC_HIERARCHY":
			for idx, __instance__ := range inferedInstance.CHILDREN.SPEC_HIERARCHY {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case SPECIFICATION_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case SPEC_HIERARCHY:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "CHILDREN.SPEC_HIERARCHY":
			for idx, __instance__ := range inferedInstance.CHILDREN.SPEC_HIERARCHY {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "IS_EDITABLE":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
			res.valueBool = inferedInstance.IS_EDITABLE
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IS_TABLE_INTERNAL":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IS_TABLE_INTERNAL)
			res.valueBool = inferedInstance.IS_TABLE_INTERNAL
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case SPEC_OBJECT:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_DATE":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_INTEGER":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_REAL":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_STRING":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_XHTML":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case SPEC_OBJECT_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case SPEC_RELATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_DATE":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_INTEGER":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_REAL":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_STRING":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "VALUES.ATTRIBUTE_VALUE_XHTML":
			for idx, __instance__ := range inferedInstance.VALUES.ATTRIBUTE_VALUE_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case SPEC_RELATION_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ALTERNATIVE_ID.ALTERNATIVE_ID":
			for idx, __instance__ := range inferedInstance.ALTERNATIVE_ID.ALTERNATIVE_ID {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			for idx, __instance__ := range inferedInstance.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "DESC":
			res.valueString = inferedInstance.DESC
		case "LAST_CHANGE":
			res.valueString = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res.valueString = inferedInstance.LONG_NAME
		}
	case XHTML_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_InlPres_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_a_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_abbr_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_acronym_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_address_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_blockquote_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_br_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_caption_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_cite_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_code_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_col_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_colgroup_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_dd_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_dfn_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_div_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_dl_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_dt_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_edit_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_em_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_h1_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_h2_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_h3_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_h4_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_h5_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_h6_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_heading_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_hr_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_kbd_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_li_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_object_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_ol_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_p_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_param_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_pre_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_q_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_samp_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_span_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_strong_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_table_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_tbody_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_td_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_tfoot_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_th_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_thead_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_tr_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_ul_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Xhtml_var_type:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
