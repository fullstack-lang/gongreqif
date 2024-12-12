// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongreqif/go/db"
	"github.com/fullstack-lang/gongreqif/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gongreqif/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoALTERNATIVE_ID BackRepoALTERNATIVE_IDStruct

	BackRepoATTRIBUTE_DEFINITION_BOOLEAN BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct

	BackRepoATTRIBUTE_DEFINITION_DATE BackRepoATTRIBUTE_DEFINITION_DATEStruct

	BackRepoATTRIBUTE_DEFINITION_ENUMERATION BackRepoATTRIBUTE_DEFINITION_ENUMERATIONStruct

	BackRepoATTRIBUTE_DEFINITION_INTEGER BackRepoATTRIBUTE_DEFINITION_INTEGERStruct

	BackRepoATTRIBUTE_DEFINITION_REAL BackRepoATTRIBUTE_DEFINITION_REALStruct

	BackRepoATTRIBUTE_DEFINITION_STRING BackRepoATTRIBUTE_DEFINITION_STRINGStruct

	BackRepoATTRIBUTE_DEFINITION_XHTML BackRepoATTRIBUTE_DEFINITION_XHTMLStruct

	BackRepoATTRIBUTE_VALUE_BOOLEAN BackRepoATTRIBUTE_VALUE_BOOLEANStruct

	BackRepoATTRIBUTE_VALUE_DATE BackRepoATTRIBUTE_VALUE_DATEStruct

	BackRepoATTRIBUTE_VALUE_ENUMERATION BackRepoATTRIBUTE_VALUE_ENUMERATIONStruct

	BackRepoATTRIBUTE_VALUE_INTEGER BackRepoATTRIBUTE_VALUE_INTEGERStruct

	BackRepoATTRIBUTE_VALUE_REAL BackRepoATTRIBUTE_VALUE_REALStruct

	BackRepoATTRIBUTE_VALUE_STRING BackRepoATTRIBUTE_VALUE_STRINGStruct

	BackRepoATTRIBUTE_VALUE_XHTML BackRepoATTRIBUTE_VALUE_XHTMLStruct

	BackRepoAnyType BackRepoAnyTypeStruct

	BackRepoDATATYPE_DEFINITION_BOOLEAN BackRepoDATATYPE_DEFINITION_BOOLEANStruct

	BackRepoDATATYPE_DEFINITION_DATE BackRepoDATATYPE_DEFINITION_DATEStruct

	BackRepoDATATYPE_DEFINITION_ENUMERATION BackRepoDATATYPE_DEFINITION_ENUMERATIONStruct

	BackRepoDATATYPE_DEFINITION_INTEGER BackRepoDATATYPE_DEFINITION_INTEGERStruct

	BackRepoDATATYPE_DEFINITION_REAL BackRepoDATATYPE_DEFINITION_REALStruct

	BackRepoDATATYPE_DEFINITION_STRING BackRepoDATATYPE_DEFINITION_STRINGStruct

	BackRepoDATATYPE_DEFINITION_XHTML BackRepoDATATYPE_DEFINITION_XHTMLStruct

	BackRepoEMBEDDED_VALUE BackRepoEMBEDDED_VALUEStruct

	BackRepoENUM_VALUE BackRepoENUM_VALUEStruct

	BackRepoRELATION_GROUP BackRepoRELATION_GROUPStruct

	BackRepoRELATION_GROUP_TYPE BackRepoRELATION_GROUP_TYPEStruct

	BackRepoREQ_IF BackRepoREQ_IFStruct

	BackRepoREQ_IF_CONTENT BackRepoREQ_IF_CONTENTStruct

	BackRepoREQ_IF_HEADER BackRepoREQ_IF_HEADERStruct

	BackRepoREQ_IF_TOOL_EXTENSION BackRepoREQ_IF_TOOL_EXTENSIONStruct

	BackRepoSPECIFICATION BackRepoSPECIFICATIONStruct

	BackRepoSPECIFICATION_TYPE BackRepoSPECIFICATION_TYPEStruct

	BackRepoSPEC_HIERARCHY BackRepoSPEC_HIERARCHYStruct

	BackRepoSPEC_OBJECT BackRepoSPEC_OBJECTStruct

	BackRepoSPEC_OBJECT_TYPE BackRepoSPEC_OBJECT_TYPEStruct

	BackRepoSPEC_RELATION BackRepoSPEC_RELATIONStruct

	BackRepoSPEC_RELATION_TYPE BackRepoSPEC_RELATION_TYPEStruct

	BackRepoXHTML_CONTENT BackRepoXHTML_CONTENTStruct

	BackRepoXhtml_InlPres_type BackRepoXhtml_InlPres_typeStruct

	BackRepoXhtml_a_type BackRepoXhtml_a_typeStruct

	BackRepoXhtml_abbr_type BackRepoXhtml_abbr_typeStruct

	BackRepoXhtml_acronym_type BackRepoXhtml_acronym_typeStruct

	BackRepoXhtml_address_type BackRepoXhtml_address_typeStruct

	BackRepoXhtml_blockquote_type BackRepoXhtml_blockquote_typeStruct

	BackRepoXhtml_br_type BackRepoXhtml_br_typeStruct

	BackRepoXhtml_caption_type BackRepoXhtml_caption_typeStruct

	BackRepoXhtml_cite_type BackRepoXhtml_cite_typeStruct

	BackRepoXhtml_code_type BackRepoXhtml_code_typeStruct

	BackRepoXhtml_col_type BackRepoXhtml_col_typeStruct

	BackRepoXhtml_colgroup_type BackRepoXhtml_colgroup_typeStruct

	BackRepoXhtml_dd_type BackRepoXhtml_dd_typeStruct

	BackRepoXhtml_dfn_type BackRepoXhtml_dfn_typeStruct

	BackRepoXhtml_div_type BackRepoXhtml_div_typeStruct

	BackRepoXhtml_dl_type BackRepoXhtml_dl_typeStruct

	BackRepoXhtml_dt_type BackRepoXhtml_dt_typeStruct

	BackRepoXhtml_edit_type BackRepoXhtml_edit_typeStruct

	BackRepoXhtml_em_type BackRepoXhtml_em_typeStruct

	BackRepoXhtml_h1_type BackRepoXhtml_h1_typeStruct

	BackRepoXhtml_h2_type BackRepoXhtml_h2_typeStruct

	BackRepoXhtml_h3_type BackRepoXhtml_h3_typeStruct

	BackRepoXhtml_h4_type BackRepoXhtml_h4_typeStruct

	BackRepoXhtml_h5_type BackRepoXhtml_h5_typeStruct

	BackRepoXhtml_h6_type BackRepoXhtml_h6_typeStruct

	BackRepoXhtml_heading_type BackRepoXhtml_heading_typeStruct

	BackRepoXhtml_hr_type BackRepoXhtml_hr_typeStruct

	BackRepoXhtml_kbd_type BackRepoXhtml_kbd_typeStruct

	BackRepoXhtml_li_type BackRepoXhtml_li_typeStruct

	BackRepoXhtml_object_type BackRepoXhtml_object_typeStruct

	BackRepoXhtml_ol_type BackRepoXhtml_ol_typeStruct

	BackRepoXhtml_p_type BackRepoXhtml_p_typeStruct

	BackRepoXhtml_param_type BackRepoXhtml_param_typeStruct

	BackRepoXhtml_pre_type BackRepoXhtml_pre_typeStruct

	BackRepoXhtml_q_type BackRepoXhtml_q_typeStruct

	BackRepoXhtml_samp_type BackRepoXhtml_samp_typeStruct

	BackRepoXhtml_span_type BackRepoXhtml_span_typeStruct

	BackRepoXhtml_strong_type BackRepoXhtml_strong_typeStruct

	BackRepoXhtml_table_type BackRepoXhtml_table_typeStruct

	BackRepoXhtml_tbody_type BackRepoXhtml_tbody_typeStruct

	BackRepoXhtml_td_type BackRepoXhtml_td_typeStruct

	BackRepoXhtml_tfoot_type BackRepoXhtml_tfoot_typeStruct

	BackRepoXhtml_th_type BackRepoXhtml_th_typeStruct

	BackRepoXhtml_thead_type BackRepoXhtml_thead_typeStruct

	BackRepoXhtml_tr_type BackRepoXhtml_tr_typeStruct

	BackRepoXhtml_ul_type BackRepoXhtml_ul_typeStruct

	BackRepoXhtml_var_type BackRepoXhtml_var_typeStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex

	subscribersRwMutex sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	var db db.DBInterface

	db = NewDBLite()

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gongreqif_go",
		&ALTERNATIVE_IDDB{},
		&ATTRIBUTE_DEFINITION_BOOLEANDB{},
		&ATTRIBUTE_DEFINITION_DATEDB{},
		&ATTRIBUTE_DEFINITION_ENUMERATIONDB{},
		&ATTRIBUTE_DEFINITION_INTEGERDB{},
		&ATTRIBUTE_DEFINITION_REALDB{},
		&ATTRIBUTE_DEFINITION_STRINGDB{},
		&ATTRIBUTE_DEFINITION_XHTMLDB{},
		&ATTRIBUTE_VALUE_BOOLEANDB{},
		&ATTRIBUTE_VALUE_DATEDB{},
		&ATTRIBUTE_VALUE_ENUMERATIONDB{},
		&ATTRIBUTE_VALUE_INTEGERDB{},
		&ATTRIBUTE_VALUE_REALDB{},
		&ATTRIBUTE_VALUE_STRINGDB{},
		&ATTRIBUTE_VALUE_XHTMLDB{},
		&AnyTypeDB{},
		&DATATYPE_DEFINITION_BOOLEANDB{},
		&DATATYPE_DEFINITION_DATEDB{},
		&DATATYPE_DEFINITION_ENUMERATIONDB{},
		&DATATYPE_DEFINITION_INTEGERDB{},
		&DATATYPE_DEFINITION_REALDB{},
		&DATATYPE_DEFINITION_STRINGDB{},
		&DATATYPE_DEFINITION_XHTMLDB{},
		&EMBEDDED_VALUEDB{},
		&ENUM_VALUEDB{},
		&RELATION_GROUPDB{},
		&RELATION_GROUP_TYPEDB{},
		&REQ_IFDB{},
		&REQ_IF_CONTENTDB{},
		&REQ_IF_HEADERDB{},
		&REQ_IF_TOOL_EXTENSIONDB{},
		&SPECIFICATIONDB{},
		&SPECIFICATION_TYPEDB{},
		&SPEC_HIERARCHYDB{},
		&SPEC_OBJECTDB{},
		&SPEC_OBJECT_TYPEDB{},
		&SPEC_RELATIONDB{},
		&SPEC_RELATION_TYPEDB{},
		&XHTML_CONTENTDB{},
		&Xhtml_InlPres_typeDB{},
		&Xhtml_a_typeDB{},
		&Xhtml_abbr_typeDB{},
		&Xhtml_acronym_typeDB{},
		&Xhtml_address_typeDB{},
		&Xhtml_blockquote_typeDB{},
		&Xhtml_br_typeDB{},
		&Xhtml_caption_typeDB{},
		&Xhtml_cite_typeDB{},
		&Xhtml_code_typeDB{},
		&Xhtml_col_typeDB{},
		&Xhtml_colgroup_typeDB{},
		&Xhtml_dd_typeDB{},
		&Xhtml_dfn_typeDB{},
		&Xhtml_div_typeDB{},
		&Xhtml_dl_typeDB{},
		&Xhtml_dt_typeDB{},
		&Xhtml_edit_typeDB{},
		&Xhtml_em_typeDB{},
		&Xhtml_h1_typeDB{},
		&Xhtml_h2_typeDB{},
		&Xhtml_h3_typeDB{},
		&Xhtml_h4_typeDB{},
		&Xhtml_h5_typeDB{},
		&Xhtml_h6_typeDB{},
		&Xhtml_heading_typeDB{},
		&Xhtml_hr_typeDB{},
		&Xhtml_kbd_typeDB{},
		&Xhtml_li_typeDB{},
		&Xhtml_object_typeDB{},
		&Xhtml_ol_typeDB{},
		&Xhtml_p_typeDB{},
		&Xhtml_param_typeDB{},
		&Xhtml_pre_typeDB{},
		&Xhtml_q_typeDB{},
		&Xhtml_samp_typeDB{},
		&Xhtml_span_typeDB{},
		&Xhtml_strong_typeDB{},
		&Xhtml_table_typeDB{},
		&Xhtml_tbody_typeDB{},
		&Xhtml_td_typeDB{},
		&Xhtml_tfoot_typeDB{},
		&Xhtml_th_typeDB{},
		&Xhtml_thead_typeDB{},
		&Xhtml_tr_typeDB{},
		&Xhtml_ul_typeDB{},
		&Xhtml_var_typeDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoALTERNATIVE_ID = BackRepoALTERNATIVE_IDStruct{
		Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDPtr: make(map[uint]*models.ALTERNATIVE_ID, 0),
		Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDDB:  make(map[uint]*ALTERNATIVE_IDDB, 0),
		Map_ALTERNATIVE_IDPtr_ALTERNATIVE_IDDBID: make(map[*models.ALTERNATIVE_ID]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN = BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct{
		Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_BOOLEAN, 0),
		Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB:  make(map[uint]*ATTRIBUTE_DEFINITION_BOOLEANDB, 0),
		Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID: make(map[*models.ATTRIBUTE_DEFINITION_BOOLEAN]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE = BackRepoATTRIBUTE_DEFINITION_DATEStruct{
		Map_ATTRIBUTE_DEFINITION_DATEDBID_ATTRIBUTE_DEFINITION_DATEPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_DATE, 0),
		Map_ATTRIBUTE_DEFINITION_DATEDBID_ATTRIBUTE_DEFINITION_DATEDB:  make(map[uint]*ATTRIBUTE_DEFINITION_DATEDB, 0),
		Map_ATTRIBUTE_DEFINITION_DATEPtr_ATTRIBUTE_DEFINITION_DATEDBID: make(map[*models.ATTRIBUTE_DEFINITION_DATE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION = BackRepoATTRIBUTE_DEFINITION_ENUMERATIONStruct{
		Map_ATTRIBUTE_DEFINITION_ENUMERATIONDBID_ATTRIBUTE_DEFINITION_ENUMERATIONPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_ENUMERATION, 0),
		Map_ATTRIBUTE_DEFINITION_ENUMERATIONDBID_ATTRIBUTE_DEFINITION_ENUMERATIONDB:  make(map[uint]*ATTRIBUTE_DEFINITION_ENUMERATIONDB, 0),
		Map_ATTRIBUTE_DEFINITION_ENUMERATIONPtr_ATTRIBUTE_DEFINITION_ENUMERATIONDBID: make(map[*models.ATTRIBUTE_DEFINITION_ENUMERATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER = BackRepoATTRIBUTE_DEFINITION_INTEGERStruct{
		Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_INTEGER, 0),
		Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB:  make(map[uint]*ATTRIBUTE_DEFINITION_INTEGERDB, 0),
		Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID: make(map[*models.ATTRIBUTE_DEFINITION_INTEGER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL = BackRepoATTRIBUTE_DEFINITION_REALStruct{
		Map_ATTRIBUTE_DEFINITION_REALDBID_ATTRIBUTE_DEFINITION_REALPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_REAL, 0),
		Map_ATTRIBUTE_DEFINITION_REALDBID_ATTRIBUTE_DEFINITION_REALDB:  make(map[uint]*ATTRIBUTE_DEFINITION_REALDB, 0),
		Map_ATTRIBUTE_DEFINITION_REALPtr_ATTRIBUTE_DEFINITION_REALDBID: make(map[*models.ATTRIBUTE_DEFINITION_REAL]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING = BackRepoATTRIBUTE_DEFINITION_STRINGStruct{
		Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_STRING, 0),
		Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB:  make(map[uint]*ATTRIBUTE_DEFINITION_STRINGDB, 0),
		Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID: make(map[*models.ATTRIBUTE_DEFINITION_STRING]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML = BackRepoATTRIBUTE_DEFINITION_XHTMLStruct{
		Map_ATTRIBUTE_DEFINITION_XHTMLDBID_ATTRIBUTE_DEFINITION_XHTMLPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_XHTML, 0),
		Map_ATTRIBUTE_DEFINITION_XHTMLDBID_ATTRIBUTE_DEFINITION_XHTMLDB:  make(map[uint]*ATTRIBUTE_DEFINITION_XHTMLDB, 0),
		Map_ATTRIBUTE_DEFINITION_XHTMLPtr_ATTRIBUTE_DEFINITION_XHTMLDBID: make(map[*models.ATTRIBUTE_DEFINITION_XHTML]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN = BackRepoATTRIBUTE_VALUE_BOOLEANStruct{
		Map_ATTRIBUTE_VALUE_BOOLEANDBID_ATTRIBUTE_VALUE_BOOLEANPtr: make(map[uint]*models.ATTRIBUTE_VALUE_BOOLEAN, 0),
		Map_ATTRIBUTE_VALUE_BOOLEANDBID_ATTRIBUTE_VALUE_BOOLEANDB:  make(map[uint]*ATTRIBUTE_VALUE_BOOLEANDB, 0),
		Map_ATTRIBUTE_VALUE_BOOLEANPtr_ATTRIBUTE_VALUE_BOOLEANDBID: make(map[*models.ATTRIBUTE_VALUE_BOOLEAN]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_DATE = BackRepoATTRIBUTE_VALUE_DATEStruct{
		Map_ATTRIBUTE_VALUE_DATEDBID_ATTRIBUTE_VALUE_DATEPtr: make(map[uint]*models.ATTRIBUTE_VALUE_DATE, 0),
		Map_ATTRIBUTE_VALUE_DATEDBID_ATTRIBUTE_VALUE_DATEDB:  make(map[uint]*ATTRIBUTE_VALUE_DATEDB, 0),
		Map_ATTRIBUTE_VALUE_DATEPtr_ATTRIBUTE_VALUE_DATEDBID: make(map[*models.ATTRIBUTE_VALUE_DATE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION = BackRepoATTRIBUTE_VALUE_ENUMERATIONStruct{
		Map_ATTRIBUTE_VALUE_ENUMERATIONDBID_ATTRIBUTE_VALUE_ENUMERATIONPtr: make(map[uint]*models.ATTRIBUTE_VALUE_ENUMERATION, 0),
		Map_ATTRIBUTE_VALUE_ENUMERATIONDBID_ATTRIBUTE_VALUE_ENUMERATIONDB:  make(map[uint]*ATTRIBUTE_VALUE_ENUMERATIONDB, 0),
		Map_ATTRIBUTE_VALUE_ENUMERATIONPtr_ATTRIBUTE_VALUE_ENUMERATIONDBID: make(map[*models.ATTRIBUTE_VALUE_ENUMERATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER = BackRepoATTRIBUTE_VALUE_INTEGERStruct{
		Map_ATTRIBUTE_VALUE_INTEGERDBID_ATTRIBUTE_VALUE_INTEGERPtr: make(map[uint]*models.ATTRIBUTE_VALUE_INTEGER, 0),
		Map_ATTRIBUTE_VALUE_INTEGERDBID_ATTRIBUTE_VALUE_INTEGERDB:  make(map[uint]*ATTRIBUTE_VALUE_INTEGERDB, 0),
		Map_ATTRIBUTE_VALUE_INTEGERPtr_ATTRIBUTE_VALUE_INTEGERDBID: make(map[*models.ATTRIBUTE_VALUE_INTEGER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_REAL = BackRepoATTRIBUTE_VALUE_REALStruct{
		Map_ATTRIBUTE_VALUE_REALDBID_ATTRIBUTE_VALUE_REALPtr: make(map[uint]*models.ATTRIBUTE_VALUE_REAL, 0),
		Map_ATTRIBUTE_VALUE_REALDBID_ATTRIBUTE_VALUE_REALDB:  make(map[uint]*ATTRIBUTE_VALUE_REALDB, 0),
		Map_ATTRIBUTE_VALUE_REALPtr_ATTRIBUTE_VALUE_REALDBID: make(map[*models.ATTRIBUTE_VALUE_REAL]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_STRING = BackRepoATTRIBUTE_VALUE_STRINGStruct{
		Map_ATTRIBUTE_VALUE_STRINGDBID_ATTRIBUTE_VALUE_STRINGPtr: make(map[uint]*models.ATTRIBUTE_VALUE_STRING, 0),
		Map_ATTRIBUTE_VALUE_STRINGDBID_ATTRIBUTE_VALUE_STRINGDB:  make(map[uint]*ATTRIBUTE_VALUE_STRINGDB, 0),
		Map_ATTRIBUTE_VALUE_STRINGPtr_ATTRIBUTE_VALUE_STRINGDBID: make(map[*models.ATTRIBUTE_VALUE_STRING]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML = BackRepoATTRIBUTE_VALUE_XHTMLStruct{
		Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr: make(map[uint]*models.ATTRIBUTE_VALUE_XHTML, 0),
		Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB:  make(map[uint]*ATTRIBUTE_VALUE_XHTMLDB, 0),
		Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID: make(map[*models.ATTRIBUTE_VALUE_XHTML]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAnyType = BackRepoAnyTypeStruct{
		Map_AnyTypeDBID_AnyTypePtr: make(map[uint]*models.AnyType, 0),
		Map_AnyTypeDBID_AnyTypeDB:  make(map[uint]*AnyTypeDB, 0),
		Map_AnyTypePtr_AnyTypeDBID: make(map[*models.AnyType]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN = BackRepoDATATYPE_DEFINITION_BOOLEANStruct{
		Map_DATATYPE_DEFINITION_BOOLEANDBID_DATATYPE_DEFINITION_BOOLEANPtr: make(map[uint]*models.DATATYPE_DEFINITION_BOOLEAN, 0),
		Map_DATATYPE_DEFINITION_BOOLEANDBID_DATATYPE_DEFINITION_BOOLEANDB:  make(map[uint]*DATATYPE_DEFINITION_BOOLEANDB, 0),
		Map_DATATYPE_DEFINITION_BOOLEANPtr_DATATYPE_DEFINITION_BOOLEANDBID: make(map[*models.DATATYPE_DEFINITION_BOOLEAN]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_DATE = BackRepoDATATYPE_DEFINITION_DATEStruct{
		Map_DATATYPE_DEFINITION_DATEDBID_DATATYPE_DEFINITION_DATEPtr: make(map[uint]*models.DATATYPE_DEFINITION_DATE, 0),
		Map_DATATYPE_DEFINITION_DATEDBID_DATATYPE_DEFINITION_DATEDB:  make(map[uint]*DATATYPE_DEFINITION_DATEDB, 0),
		Map_DATATYPE_DEFINITION_DATEPtr_DATATYPE_DEFINITION_DATEDBID: make(map[*models.DATATYPE_DEFINITION_DATE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION = BackRepoDATATYPE_DEFINITION_ENUMERATIONStruct{
		Map_DATATYPE_DEFINITION_ENUMERATIONDBID_DATATYPE_DEFINITION_ENUMERATIONPtr: make(map[uint]*models.DATATYPE_DEFINITION_ENUMERATION, 0),
		Map_DATATYPE_DEFINITION_ENUMERATIONDBID_DATATYPE_DEFINITION_ENUMERATIONDB:  make(map[uint]*DATATYPE_DEFINITION_ENUMERATIONDB, 0),
		Map_DATATYPE_DEFINITION_ENUMERATIONPtr_DATATYPE_DEFINITION_ENUMERATIONDBID: make(map[*models.DATATYPE_DEFINITION_ENUMERATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER = BackRepoDATATYPE_DEFINITION_INTEGERStruct{
		Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr: make(map[uint]*models.DATATYPE_DEFINITION_INTEGER, 0),
		Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB:  make(map[uint]*DATATYPE_DEFINITION_INTEGERDB, 0),
		Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID: make(map[*models.DATATYPE_DEFINITION_INTEGER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_REAL = BackRepoDATATYPE_DEFINITION_REALStruct{
		Map_DATATYPE_DEFINITION_REALDBID_DATATYPE_DEFINITION_REALPtr: make(map[uint]*models.DATATYPE_DEFINITION_REAL, 0),
		Map_DATATYPE_DEFINITION_REALDBID_DATATYPE_DEFINITION_REALDB:  make(map[uint]*DATATYPE_DEFINITION_REALDB, 0),
		Map_DATATYPE_DEFINITION_REALPtr_DATATYPE_DEFINITION_REALDBID: make(map[*models.DATATYPE_DEFINITION_REAL]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_STRING = BackRepoDATATYPE_DEFINITION_STRINGStruct{
		Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr: make(map[uint]*models.DATATYPE_DEFINITION_STRING, 0),
		Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB:  make(map[uint]*DATATYPE_DEFINITION_STRINGDB, 0),
		Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID: make(map[*models.DATATYPE_DEFINITION_STRING]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML = BackRepoDATATYPE_DEFINITION_XHTMLStruct{
		Map_DATATYPE_DEFINITION_XHTMLDBID_DATATYPE_DEFINITION_XHTMLPtr: make(map[uint]*models.DATATYPE_DEFINITION_XHTML, 0),
		Map_DATATYPE_DEFINITION_XHTMLDBID_DATATYPE_DEFINITION_XHTMLDB:  make(map[uint]*DATATYPE_DEFINITION_XHTMLDB, 0),
		Map_DATATYPE_DEFINITION_XHTMLPtr_DATATYPE_DEFINITION_XHTMLDBID: make(map[*models.DATATYPE_DEFINITION_XHTML]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEMBEDDED_VALUE = BackRepoEMBEDDED_VALUEStruct{
		Map_EMBEDDED_VALUEDBID_EMBEDDED_VALUEPtr: make(map[uint]*models.EMBEDDED_VALUE, 0),
		Map_EMBEDDED_VALUEDBID_EMBEDDED_VALUEDB:  make(map[uint]*EMBEDDED_VALUEDB, 0),
		Map_EMBEDDED_VALUEPtr_EMBEDDED_VALUEDBID: make(map[*models.EMBEDDED_VALUE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoENUM_VALUE = BackRepoENUM_VALUEStruct{
		Map_ENUM_VALUEDBID_ENUM_VALUEPtr: make(map[uint]*models.ENUM_VALUE, 0),
		Map_ENUM_VALUEDBID_ENUM_VALUEDB:  make(map[uint]*ENUM_VALUEDB, 0),
		Map_ENUM_VALUEPtr_ENUM_VALUEDBID: make(map[*models.ENUM_VALUE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRELATION_GROUP = BackRepoRELATION_GROUPStruct{
		Map_RELATION_GROUPDBID_RELATION_GROUPPtr: make(map[uint]*models.RELATION_GROUP, 0),
		Map_RELATION_GROUPDBID_RELATION_GROUPDB:  make(map[uint]*RELATION_GROUPDB, 0),
		Map_RELATION_GROUPPtr_RELATION_GROUPDBID: make(map[*models.RELATION_GROUP]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRELATION_GROUP_TYPE = BackRepoRELATION_GROUP_TYPEStruct{
		Map_RELATION_GROUP_TYPEDBID_RELATION_GROUP_TYPEPtr: make(map[uint]*models.RELATION_GROUP_TYPE, 0),
		Map_RELATION_GROUP_TYPEDBID_RELATION_GROUP_TYPEDB:  make(map[uint]*RELATION_GROUP_TYPEDB, 0),
		Map_RELATION_GROUP_TYPEPtr_RELATION_GROUP_TYPEDBID: make(map[*models.RELATION_GROUP_TYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQ_IF = BackRepoREQ_IFStruct{
		Map_REQ_IFDBID_REQ_IFPtr: make(map[uint]*models.REQ_IF, 0),
		Map_REQ_IFDBID_REQ_IFDB:  make(map[uint]*REQ_IFDB, 0),
		Map_REQ_IFPtr_REQ_IFDBID: make(map[*models.REQ_IF]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQ_IF_CONTENT = BackRepoREQ_IF_CONTENTStruct{
		Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr: make(map[uint]*models.REQ_IF_CONTENT, 0),
		Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB:  make(map[uint]*REQ_IF_CONTENTDB, 0),
		Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID: make(map[*models.REQ_IF_CONTENT]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQ_IF_HEADER = BackRepoREQ_IF_HEADERStruct{
		Map_REQ_IF_HEADERDBID_REQ_IF_HEADERPtr: make(map[uint]*models.REQ_IF_HEADER, 0),
		Map_REQ_IF_HEADERDBID_REQ_IF_HEADERDB:  make(map[uint]*REQ_IF_HEADERDB, 0),
		Map_REQ_IF_HEADERPtr_REQ_IF_HEADERDBID: make(map[*models.REQ_IF_HEADER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION = BackRepoREQ_IF_TOOL_EXTENSIONStruct{
		Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr: make(map[uint]*models.REQ_IF_TOOL_EXTENSION, 0),
		Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB:  make(map[uint]*REQ_IF_TOOL_EXTENSIONDB, 0),
		Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID: make(map[*models.REQ_IF_TOOL_EXTENSION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECIFICATION = BackRepoSPECIFICATIONStruct{
		Map_SPECIFICATIONDBID_SPECIFICATIONPtr: make(map[uint]*models.SPECIFICATION, 0),
		Map_SPECIFICATIONDBID_SPECIFICATIONDB:  make(map[uint]*SPECIFICATIONDB, 0),
		Map_SPECIFICATIONPtr_SPECIFICATIONDBID: make(map[*models.SPECIFICATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECIFICATION_TYPE = BackRepoSPECIFICATION_TYPEStruct{
		Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEPtr: make(map[uint]*models.SPECIFICATION_TYPE, 0),
		Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEDB:  make(map[uint]*SPECIFICATION_TYPEDB, 0),
		Map_SPECIFICATION_TYPEPtr_SPECIFICATION_TYPEDBID: make(map[*models.SPECIFICATION_TYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPEC_HIERARCHY = BackRepoSPEC_HIERARCHYStruct{
		Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr: make(map[uint]*models.SPEC_HIERARCHY, 0),
		Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB:  make(map[uint]*SPEC_HIERARCHYDB, 0),
		Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID: make(map[*models.SPEC_HIERARCHY]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPEC_OBJECT = BackRepoSPEC_OBJECTStruct{
		Map_SPEC_OBJECTDBID_SPEC_OBJECTPtr: make(map[uint]*models.SPEC_OBJECT, 0),
		Map_SPEC_OBJECTDBID_SPEC_OBJECTDB:  make(map[uint]*SPEC_OBJECTDB, 0),
		Map_SPEC_OBJECTPtr_SPEC_OBJECTDBID: make(map[*models.SPEC_OBJECT]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPEC_OBJECT_TYPE = BackRepoSPEC_OBJECT_TYPEStruct{
		Map_SPEC_OBJECT_TYPEDBID_SPEC_OBJECT_TYPEPtr: make(map[uint]*models.SPEC_OBJECT_TYPE, 0),
		Map_SPEC_OBJECT_TYPEDBID_SPEC_OBJECT_TYPEDB:  make(map[uint]*SPEC_OBJECT_TYPEDB, 0),
		Map_SPEC_OBJECT_TYPEPtr_SPEC_OBJECT_TYPEDBID: make(map[*models.SPEC_OBJECT_TYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPEC_RELATION = BackRepoSPEC_RELATIONStruct{
		Map_SPEC_RELATIONDBID_SPEC_RELATIONPtr: make(map[uint]*models.SPEC_RELATION, 0),
		Map_SPEC_RELATIONDBID_SPEC_RELATIONDB:  make(map[uint]*SPEC_RELATIONDB, 0),
		Map_SPEC_RELATIONPtr_SPEC_RELATIONDBID: make(map[*models.SPEC_RELATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPEC_RELATION_TYPE = BackRepoSPEC_RELATION_TYPEStruct{
		Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr: make(map[uint]*models.SPEC_RELATION_TYPE, 0),
		Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB:  make(map[uint]*SPEC_RELATION_TYPEDB, 0),
		Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID: make(map[*models.SPEC_RELATION_TYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXHTML_CONTENT = BackRepoXHTML_CONTENTStruct{
		Map_XHTML_CONTENTDBID_XHTML_CONTENTPtr: make(map[uint]*models.XHTML_CONTENT, 0),
		Map_XHTML_CONTENTDBID_XHTML_CONTENTDB:  make(map[uint]*XHTML_CONTENTDB, 0),
		Map_XHTML_CONTENTPtr_XHTML_CONTENTDBID: make(map[*models.XHTML_CONTENT]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_InlPres_type = BackRepoXhtml_InlPres_typeStruct{
		Map_Xhtml_InlPres_typeDBID_Xhtml_InlPres_typePtr: make(map[uint]*models.Xhtml_InlPres_type, 0),
		Map_Xhtml_InlPres_typeDBID_Xhtml_InlPres_typeDB:  make(map[uint]*Xhtml_InlPres_typeDB, 0),
		Map_Xhtml_InlPres_typePtr_Xhtml_InlPres_typeDBID: make(map[*models.Xhtml_InlPres_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_a_type = BackRepoXhtml_a_typeStruct{
		Map_Xhtml_a_typeDBID_Xhtml_a_typePtr: make(map[uint]*models.Xhtml_a_type, 0),
		Map_Xhtml_a_typeDBID_Xhtml_a_typeDB:  make(map[uint]*Xhtml_a_typeDB, 0),
		Map_Xhtml_a_typePtr_Xhtml_a_typeDBID: make(map[*models.Xhtml_a_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_abbr_type = BackRepoXhtml_abbr_typeStruct{
		Map_Xhtml_abbr_typeDBID_Xhtml_abbr_typePtr: make(map[uint]*models.Xhtml_abbr_type, 0),
		Map_Xhtml_abbr_typeDBID_Xhtml_abbr_typeDB:  make(map[uint]*Xhtml_abbr_typeDB, 0),
		Map_Xhtml_abbr_typePtr_Xhtml_abbr_typeDBID: make(map[*models.Xhtml_abbr_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_acronym_type = BackRepoXhtml_acronym_typeStruct{
		Map_Xhtml_acronym_typeDBID_Xhtml_acronym_typePtr: make(map[uint]*models.Xhtml_acronym_type, 0),
		Map_Xhtml_acronym_typeDBID_Xhtml_acronym_typeDB:  make(map[uint]*Xhtml_acronym_typeDB, 0),
		Map_Xhtml_acronym_typePtr_Xhtml_acronym_typeDBID: make(map[*models.Xhtml_acronym_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_address_type = BackRepoXhtml_address_typeStruct{
		Map_Xhtml_address_typeDBID_Xhtml_address_typePtr: make(map[uint]*models.Xhtml_address_type, 0),
		Map_Xhtml_address_typeDBID_Xhtml_address_typeDB:  make(map[uint]*Xhtml_address_typeDB, 0),
		Map_Xhtml_address_typePtr_Xhtml_address_typeDBID: make(map[*models.Xhtml_address_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_blockquote_type = BackRepoXhtml_blockquote_typeStruct{
		Map_Xhtml_blockquote_typeDBID_Xhtml_blockquote_typePtr: make(map[uint]*models.Xhtml_blockquote_type, 0),
		Map_Xhtml_blockquote_typeDBID_Xhtml_blockquote_typeDB:  make(map[uint]*Xhtml_blockquote_typeDB, 0),
		Map_Xhtml_blockquote_typePtr_Xhtml_blockquote_typeDBID: make(map[*models.Xhtml_blockquote_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_br_type = BackRepoXhtml_br_typeStruct{
		Map_Xhtml_br_typeDBID_Xhtml_br_typePtr: make(map[uint]*models.Xhtml_br_type, 0),
		Map_Xhtml_br_typeDBID_Xhtml_br_typeDB:  make(map[uint]*Xhtml_br_typeDB, 0),
		Map_Xhtml_br_typePtr_Xhtml_br_typeDBID: make(map[*models.Xhtml_br_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_caption_type = BackRepoXhtml_caption_typeStruct{
		Map_Xhtml_caption_typeDBID_Xhtml_caption_typePtr: make(map[uint]*models.Xhtml_caption_type, 0),
		Map_Xhtml_caption_typeDBID_Xhtml_caption_typeDB:  make(map[uint]*Xhtml_caption_typeDB, 0),
		Map_Xhtml_caption_typePtr_Xhtml_caption_typeDBID: make(map[*models.Xhtml_caption_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_cite_type = BackRepoXhtml_cite_typeStruct{
		Map_Xhtml_cite_typeDBID_Xhtml_cite_typePtr: make(map[uint]*models.Xhtml_cite_type, 0),
		Map_Xhtml_cite_typeDBID_Xhtml_cite_typeDB:  make(map[uint]*Xhtml_cite_typeDB, 0),
		Map_Xhtml_cite_typePtr_Xhtml_cite_typeDBID: make(map[*models.Xhtml_cite_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_code_type = BackRepoXhtml_code_typeStruct{
		Map_Xhtml_code_typeDBID_Xhtml_code_typePtr: make(map[uint]*models.Xhtml_code_type, 0),
		Map_Xhtml_code_typeDBID_Xhtml_code_typeDB:  make(map[uint]*Xhtml_code_typeDB, 0),
		Map_Xhtml_code_typePtr_Xhtml_code_typeDBID: make(map[*models.Xhtml_code_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_col_type = BackRepoXhtml_col_typeStruct{
		Map_Xhtml_col_typeDBID_Xhtml_col_typePtr: make(map[uint]*models.Xhtml_col_type, 0),
		Map_Xhtml_col_typeDBID_Xhtml_col_typeDB:  make(map[uint]*Xhtml_col_typeDB, 0),
		Map_Xhtml_col_typePtr_Xhtml_col_typeDBID: make(map[*models.Xhtml_col_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_colgroup_type = BackRepoXhtml_colgroup_typeStruct{
		Map_Xhtml_colgroup_typeDBID_Xhtml_colgroup_typePtr: make(map[uint]*models.Xhtml_colgroup_type, 0),
		Map_Xhtml_colgroup_typeDBID_Xhtml_colgroup_typeDB:  make(map[uint]*Xhtml_colgroup_typeDB, 0),
		Map_Xhtml_colgroup_typePtr_Xhtml_colgroup_typeDBID: make(map[*models.Xhtml_colgroup_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_dd_type = BackRepoXhtml_dd_typeStruct{
		Map_Xhtml_dd_typeDBID_Xhtml_dd_typePtr: make(map[uint]*models.Xhtml_dd_type, 0),
		Map_Xhtml_dd_typeDBID_Xhtml_dd_typeDB:  make(map[uint]*Xhtml_dd_typeDB, 0),
		Map_Xhtml_dd_typePtr_Xhtml_dd_typeDBID: make(map[*models.Xhtml_dd_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_dfn_type = BackRepoXhtml_dfn_typeStruct{
		Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr: make(map[uint]*models.Xhtml_dfn_type, 0),
		Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB:  make(map[uint]*Xhtml_dfn_typeDB, 0),
		Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID: make(map[*models.Xhtml_dfn_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_div_type = BackRepoXhtml_div_typeStruct{
		Map_Xhtml_div_typeDBID_Xhtml_div_typePtr: make(map[uint]*models.Xhtml_div_type, 0),
		Map_Xhtml_div_typeDBID_Xhtml_div_typeDB:  make(map[uint]*Xhtml_div_typeDB, 0),
		Map_Xhtml_div_typePtr_Xhtml_div_typeDBID: make(map[*models.Xhtml_div_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_dl_type = BackRepoXhtml_dl_typeStruct{
		Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr: make(map[uint]*models.Xhtml_dl_type, 0),
		Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB:  make(map[uint]*Xhtml_dl_typeDB, 0),
		Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID: make(map[*models.Xhtml_dl_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_dt_type = BackRepoXhtml_dt_typeStruct{
		Map_Xhtml_dt_typeDBID_Xhtml_dt_typePtr: make(map[uint]*models.Xhtml_dt_type, 0),
		Map_Xhtml_dt_typeDBID_Xhtml_dt_typeDB:  make(map[uint]*Xhtml_dt_typeDB, 0),
		Map_Xhtml_dt_typePtr_Xhtml_dt_typeDBID: make(map[*models.Xhtml_dt_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_edit_type = BackRepoXhtml_edit_typeStruct{
		Map_Xhtml_edit_typeDBID_Xhtml_edit_typePtr: make(map[uint]*models.Xhtml_edit_type, 0),
		Map_Xhtml_edit_typeDBID_Xhtml_edit_typeDB:  make(map[uint]*Xhtml_edit_typeDB, 0),
		Map_Xhtml_edit_typePtr_Xhtml_edit_typeDBID: make(map[*models.Xhtml_edit_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_em_type = BackRepoXhtml_em_typeStruct{
		Map_Xhtml_em_typeDBID_Xhtml_em_typePtr: make(map[uint]*models.Xhtml_em_type, 0),
		Map_Xhtml_em_typeDBID_Xhtml_em_typeDB:  make(map[uint]*Xhtml_em_typeDB, 0),
		Map_Xhtml_em_typePtr_Xhtml_em_typeDBID: make(map[*models.Xhtml_em_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_h1_type = BackRepoXhtml_h1_typeStruct{
		Map_Xhtml_h1_typeDBID_Xhtml_h1_typePtr: make(map[uint]*models.Xhtml_h1_type, 0),
		Map_Xhtml_h1_typeDBID_Xhtml_h1_typeDB:  make(map[uint]*Xhtml_h1_typeDB, 0),
		Map_Xhtml_h1_typePtr_Xhtml_h1_typeDBID: make(map[*models.Xhtml_h1_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_h2_type = BackRepoXhtml_h2_typeStruct{
		Map_Xhtml_h2_typeDBID_Xhtml_h2_typePtr: make(map[uint]*models.Xhtml_h2_type, 0),
		Map_Xhtml_h2_typeDBID_Xhtml_h2_typeDB:  make(map[uint]*Xhtml_h2_typeDB, 0),
		Map_Xhtml_h2_typePtr_Xhtml_h2_typeDBID: make(map[*models.Xhtml_h2_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_h3_type = BackRepoXhtml_h3_typeStruct{
		Map_Xhtml_h3_typeDBID_Xhtml_h3_typePtr: make(map[uint]*models.Xhtml_h3_type, 0),
		Map_Xhtml_h3_typeDBID_Xhtml_h3_typeDB:  make(map[uint]*Xhtml_h3_typeDB, 0),
		Map_Xhtml_h3_typePtr_Xhtml_h3_typeDBID: make(map[*models.Xhtml_h3_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_h4_type = BackRepoXhtml_h4_typeStruct{
		Map_Xhtml_h4_typeDBID_Xhtml_h4_typePtr: make(map[uint]*models.Xhtml_h4_type, 0),
		Map_Xhtml_h4_typeDBID_Xhtml_h4_typeDB:  make(map[uint]*Xhtml_h4_typeDB, 0),
		Map_Xhtml_h4_typePtr_Xhtml_h4_typeDBID: make(map[*models.Xhtml_h4_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_h5_type = BackRepoXhtml_h5_typeStruct{
		Map_Xhtml_h5_typeDBID_Xhtml_h5_typePtr: make(map[uint]*models.Xhtml_h5_type, 0),
		Map_Xhtml_h5_typeDBID_Xhtml_h5_typeDB:  make(map[uint]*Xhtml_h5_typeDB, 0),
		Map_Xhtml_h5_typePtr_Xhtml_h5_typeDBID: make(map[*models.Xhtml_h5_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_h6_type = BackRepoXhtml_h6_typeStruct{
		Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr: make(map[uint]*models.Xhtml_h6_type, 0),
		Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB:  make(map[uint]*Xhtml_h6_typeDB, 0),
		Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID: make(map[*models.Xhtml_h6_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_heading_type = BackRepoXhtml_heading_typeStruct{
		Map_Xhtml_heading_typeDBID_Xhtml_heading_typePtr: make(map[uint]*models.Xhtml_heading_type, 0),
		Map_Xhtml_heading_typeDBID_Xhtml_heading_typeDB:  make(map[uint]*Xhtml_heading_typeDB, 0),
		Map_Xhtml_heading_typePtr_Xhtml_heading_typeDBID: make(map[*models.Xhtml_heading_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_hr_type = BackRepoXhtml_hr_typeStruct{
		Map_Xhtml_hr_typeDBID_Xhtml_hr_typePtr: make(map[uint]*models.Xhtml_hr_type, 0),
		Map_Xhtml_hr_typeDBID_Xhtml_hr_typeDB:  make(map[uint]*Xhtml_hr_typeDB, 0),
		Map_Xhtml_hr_typePtr_Xhtml_hr_typeDBID: make(map[*models.Xhtml_hr_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_kbd_type = BackRepoXhtml_kbd_typeStruct{
		Map_Xhtml_kbd_typeDBID_Xhtml_kbd_typePtr: make(map[uint]*models.Xhtml_kbd_type, 0),
		Map_Xhtml_kbd_typeDBID_Xhtml_kbd_typeDB:  make(map[uint]*Xhtml_kbd_typeDB, 0),
		Map_Xhtml_kbd_typePtr_Xhtml_kbd_typeDBID: make(map[*models.Xhtml_kbd_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_li_type = BackRepoXhtml_li_typeStruct{
		Map_Xhtml_li_typeDBID_Xhtml_li_typePtr: make(map[uint]*models.Xhtml_li_type, 0),
		Map_Xhtml_li_typeDBID_Xhtml_li_typeDB:  make(map[uint]*Xhtml_li_typeDB, 0),
		Map_Xhtml_li_typePtr_Xhtml_li_typeDBID: make(map[*models.Xhtml_li_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_object_type = BackRepoXhtml_object_typeStruct{
		Map_Xhtml_object_typeDBID_Xhtml_object_typePtr: make(map[uint]*models.Xhtml_object_type, 0),
		Map_Xhtml_object_typeDBID_Xhtml_object_typeDB:  make(map[uint]*Xhtml_object_typeDB, 0),
		Map_Xhtml_object_typePtr_Xhtml_object_typeDBID: make(map[*models.Xhtml_object_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_ol_type = BackRepoXhtml_ol_typeStruct{
		Map_Xhtml_ol_typeDBID_Xhtml_ol_typePtr: make(map[uint]*models.Xhtml_ol_type, 0),
		Map_Xhtml_ol_typeDBID_Xhtml_ol_typeDB:  make(map[uint]*Xhtml_ol_typeDB, 0),
		Map_Xhtml_ol_typePtr_Xhtml_ol_typeDBID: make(map[*models.Xhtml_ol_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_p_type = BackRepoXhtml_p_typeStruct{
		Map_Xhtml_p_typeDBID_Xhtml_p_typePtr: make(map[uint]*models.Xhtml_p_type, 0),
		Map_Xhtml_p_typeDBID_Xhtml_p_typeDB:  make(map[uint]*Xhtml_p_typeDB, 0),
		Map_Xhtml_p_typePtr_Xhtml_p_typeDBID: make(map[*models.Xhtml_p_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_param_type = BackRepoXhtml_param_typeStruct{
		Map_Xhtml_param_typeDBID_Xhtml_param_typePtr: make(map[uint]*models.Xhtml_param_type, 0),
		Map_Xhtml_param_typeDBID_Xhtml_param_typeDB:  make(map[uint]*Xhtml_param_typeDB, 0),
		Map_Xhtml_param_typePtr_Xhtml_param_typeDBID: make(map[*models.Xhtml_param_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_pre_type = BackRepoXhtml_pre_typeStruct{
		Map_Xhtml_pre_typeDBID_Xhtml_pre_typePtr: make(map[uint]*models.Xhtml_pre_type, 0),
		Map_Xhtml_pre_typeDBID_Xhtml_pre_typeDB:  make(map[uint]*Xhtml_pre_typeDB, 0),
		Map_Xhtml_pre_typePtr_Xhtml_pre_typeDBID: make(map[*models.Xhtml_pre_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_q_type = BackRepoXhtml_q_typeStruct{
		Map_Xhtml_q_typeDBID_Xhtml_q_typePtr: make(map[uint]*models.Xhtml_q_type, 0),
		Map_Xhtml_q_typeDBID_Xhtml_q_typeDB:  make(map[uint]*Xhtml_q_typeDB, 0),
		Map_Xhtml_q_typePtr_Xhtml_q_typeDBID: make(map[*models.Xhtml_q_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_samp_type = BackRepoXhtml_samp_typeStruct{
		Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr: make(map[uint]*models.Xhtml_samp_type, 0),
		Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB:  make(map[uint]*Xhtml_samp_typeDB, 0),
		Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID: make(map[*models.Xhtml_samp_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_span_type = BackRepoXhtml_span_typeStruct{
		Map_Xhtml_span_typeDBID_Xhtml_span_typePtr: make(map[uint]*models.Xhtml_span_type, 0),
		Map_Xhtml_span_typeDBID_Xhtml_span_typeDB:  make(map[uint]*Xhtml_span_typeDB, 0),
		Map_Xhtml_span_typePtr_Xhtml_span_typeDBID: make(map[*models.Xhtml_span_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_strong_type = BackRepoXhtml_strong_typeStruct{
		Map_Xhtml_strong_typeDBID_Xhtml_strong_typePtr: make(map[uint]*models.Xhtml_strong_type, 0),
		Map_Xhtml_strong_typeDBID_Xhtml_strong_typeDB:  make(map[uint]*Xhtml_strong_typeDB, 0),
		Map_Xhtml_strong_typePtr_Xhtml_strong_typeDBID: make(map[*models.Xhtml_strong_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_table_type = BackRepoXhtml_table_typeStruct{
		Map_Xhtml_table_typeDBID_Xhtml_table_typePtr: make(map[uint]*models.Xhtml_table_type, 0),
		Map_Xhtml_table_typeDBID_Xhtml_table_typeDB:  make(map[uint]*Xhtml_table_typeDB, 0),
		Map_Xhtml_table_typePtr_Xhtml_table_typeDBID: make(map[*models.Xhtml_table_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_tbody_type = BackRepoXhtml_tbody_typeStruct{
		Map_Xhtml_tbody_typeDBID_Xhtml_tbody_typePtr: make(map[uint]*models.Xhtml_tbody_type, 0),
		Map_Xhtml_tbody_typeDBID_Xhtml_tbody_typeDB:  make(map[uint]*Xhtml_tbody_typeDB, 0),
		Map_Xhtml_tbody_typePtr_Xhtml_tbody_typeDBID: make(map[*models.Xhtml_tbody_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_td_type = BackRepoXhtml_td_typeStruct{
		Map_Xhtml_td_typeDBID_Xhtml_td_typePtr: make(map[uint]*models.Xhtml_td_type, 0),
		Map_Xhtml_td_typeDBID_Xhtml_td_typeDB:  make(map[uint]*Xhtml_td_typeDB, 0),
		Map_Xhtml_td_typePtr_Xhtml_td_typeDBID: make(map[*models.Xhtml_td_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_tfoot_type = BackRepoXhtml_tfoot_typeStruct{
		Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr: make(map[uint]*models.Xhtml_tfoot_type, 0),
		Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB:  make(map[uint]*Xhtml_tfoot_typeDB, 0),
		Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID: make(map[*models.Xhtml_tfoot_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_th_type = BackRepoXhtml_th_typeStruct{
		Map_Xhtml_th_typeDBID_Xhtml_th_typePtr: make(map[uint]*models.Xhtml_th_type, 0),
		Map_Xhtml_th_typeDBID_Xhtml_th_typeDB:  make(map[uint]*Xhtml_th_typeDB, 0),
		Map_Xhtml_th_typePtr_Xhtml_th_typeDBID: make(map[*models.Xhtml_th_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_thead_type = BackRepoXhtml_thead_typeStruct{
		Map_Xhtml_thead_typeDBID_Xhtml_thead_typePtr: make(map[uint]*models.Xhtml_thead_type, 0),
		Map_Xhtml_thead_typeDBID_Xhtml_thead_typeDB:  make(map[uint]*Xhtml_thead_typeDB, 0),
		Map_Xhtml_thead_typePtr_Xhtml_thead_typeDBID: make(map[*models.Xhtml_thead_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_tr_type = BackRepoXhtml_tr_typeStruct{
		Map_Xhtml_tr_typeDBID_Xhtml_tr_typePtr: make(map[uint]*models.Xhtml_tr_type, 0),
		Map_Xhtml_tr_typeDBID_Xhtml_tr_typeDB:  make(map[uint]*Xhtml_tr_typeDB, 0),
		Map_Xhtml_tr_typePtr_Xhtml_tr_typeDBID: make(map[*models.Xhtml_tr_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_ul_type = BackRepoXhtml_ul_typeStruct{
		Map_Xhtml_ul_typeDBID_Xhtml_ul_typePtr: make(map[uint]*models.Xhtml_ul_type, 0),
		Map_Xhtml_ul_typeDBID_Xhtml_ul_typeDB:  make(map[uint]*Xhtml_ul_typeDB, 0),
		Map_Xhtml_ul_typePtr_Xhtml_ul_typeDBID: make(map[*models.Xhtml_ul_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXhtml_var_type = BackRepoXhtml_var_typeStruct{
		Map_Xhtml_var_typeDBID_Xhtml_var_typePtr: make(map[uint]*models.Xhtml_var_type, 0),
		Map_Xhtml_var_typeDBID_Xhtml_var_typeDB:  make(map[uint]*Xhtml_var_typeDB, 0),
		Map_Xhtml_var_typePtr_Xhtml_var_typeDBID: make(map[*models.Xhtml_var_type]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1

	backRepo.broadcastNbCommitToBack()

	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {

	// forbid read of back repo during commit
	backRepo.rwMutex.Lock()
	defer backRepo.rwMutex.Unlock()

	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoALTERNATIVE_ID.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CommitPhaseOne(stage)
	backRepo.BackRepoAnyType.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.CommitPhaseOne(stage)
	backRepo.BackRepoEMBEDDED_VALUE.CommitPhaseOne(stage)
	backRepo.BackRepoENUM_VALUE.CommitPhaseOne(stage)
	backRepo.BackRepoRELATION_GROUP.CommitPhaseOne(stage)
	backRepo.BackRepoRELATION_GROUP_TYPE.CommitPhaseOne(stage)
	backRepo.BackRepoREQ_IF.CommitPhaseOne(stage)
	backRepo.BackRepoREQ_IF_CONTENT.CommitPhaseOne(stage)
	backRepo.BackRepoREQ_IF_HEADER.CommitPhaseOne(stage)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CommitPhaseOne(stage)
	backRepo.BackRepoSPECIFICATION.CommitPhaseOne(stage)
	backRepo.BackRepoSPECIFICATION_TYPE.CommitPhaseOne(stage)
	backRepo.BackRepoSPEC_HIERARCHY.CommitPhaseOne(stage)
	backRepo.BackRepoSPEC_OBJECT.CommitPhaseOne(stage)
	backRepo.BackRepoSPEC_OBJECT_TYPE.CommitPhaseOne(stage)
	backRepo.BackRepoSPEC_RELATION.CommitPhaseOne(stage)
	backRepo.BackRepoSPEC_RELATION_TYPE.CommitPhaseOne(stage)
	backRepo.BackRepoXHTML_CONTENT.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_InlPres_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_a_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_abbr_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_acronym_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_address_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_blockquote_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_br_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_caption_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_cite_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_code_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_col_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_colgroup_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_dd_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_dfn_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_div_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_dl_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_dt_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_edit_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_em_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_h1_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_h2_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_h3_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_h4_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_h5_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_h6_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_heading_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_hr_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_kbd_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_li_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_object_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_ol_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_p_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_param_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_pre_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_q_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_samp_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_span_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_strong_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_table_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_tbody_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_td_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_tfoot_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_th_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_thead_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_tr_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_ul_type.CommitPhaseOne(stage)
	backRepo.BackRepoXhtml_var_type.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoALTERNATIVE_ID.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAnyType.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEMBEDDED_VALUE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoENUM_VALUE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRELATION_GROUP.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRELATION_GROUP_TYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_CONTENT.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_HEADER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION_TYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_HIERARCHY.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_OBJECT.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_OBJECT_TYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_RELATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_RELATION_TYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXHTML_CONTENT.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_InlPres_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_a_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_abbr_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_acronym_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_address_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_blockquote_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_br_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_caption_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_cite_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_code_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_col_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_colgroup_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_dd_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_dfn_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_div_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_dl_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_dt_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_edit_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_em_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h1_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h2_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h3_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h4_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h5_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h6_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_heading_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_hr_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_kbd_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_li_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_object_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_ol_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_p_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_param_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_pre_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_q_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_samp_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_span_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_strong_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_table_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_tbody_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_td_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_tfoot_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_th_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_thead_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_tr_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_ul_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_var_type.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoALTERNATIVE_ID.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseOne()
	backRepo.BackRepoAnyType.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.CheckoutPhaseOne()
	backRepo.BackRepoEMBEDDED_VALUE.CheckoutPhaseOne()
	backRepo.BackRepoENUM_VALUE.CheckoutPhaseOne()
	backRepo.BackRepoRELATION_GROUP.CheckoutPhaseOne()
	backRepo.BackRepoRELATION_GROUP_TYPE.CheckoutPhaseOne()
	backRepo.BackRepoREQ_IF.CheckoutPhaseOne()
	backRepo.BackRepoREQ_IF_CONTENT.CheckoutPhaseOne()
	backRepo.BackRepoREQ_IF_HEADER.CheckoutPhaseOne()
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseOne()
	backRepo.BackRepoSPECIFICATION.CheckoutPhaseOne()
	backRepo.BackRepoSPECIFICATION_TYPE.CheckoutPhaseOne()
	backRepo.BackRepoSPEC_HIERARCHY.CheckoutPhaseOne()
	backRepo.BackRepoSPEC_OBJECT.CheckoutPhaseOne()
	backRepo.BackRepoSPEC_OBJECT_TYPE.CheckoutPhaseOne()
	backRepo.BackRepoSPEC_RELATION.CheckoutPhaseOne()
	backRepo.BackRepoSPEC_RELATION_TYPE.CheckoutPhaseOne()
	backRepo.BackRepoXHTML_CONTENT.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_InlPres_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_a_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_abbr_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_acronym_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_address_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_blockquote_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_br_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_caption_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_cite_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_code_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_col_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_colgroup_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_dd_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_dfn_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_div_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_dl_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_dt_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_edit_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_em_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_h1_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_h2_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_h3_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_h4_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_h5_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_h6_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_heading_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_hr_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_kbd_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_li_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_object_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_ol_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_p_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_param_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_pre_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_q_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_samp_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_span_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_strong_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_table_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_tbody_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_td_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_tfoot_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_th_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_thead_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_tr_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_ul_type.CheckoutPhaseOne()
	backRepo.BackRepoXhtml_var_type.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoALTERNATIVE_ID.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAnyType.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEMBEDDED_VALUE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoENUM_VALUE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRELATION_GROUP.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRELATION_GROUP_TYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_CONTENT.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_HEADER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION_TYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_HIERARCHY.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_OBJECT.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_OBJECT_TYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_RELATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_RELATION_TYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXHTML_CONTENT.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_InlPres_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_a_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_abbr_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_acronym_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_address_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_blockquote_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_br_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_caption_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_cite_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_code_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_col_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_colgroup_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_dd_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_dfn_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_div_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_dl_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_dt_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_edit_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_em_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h1_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h2_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h3_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h4_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h5_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_h6_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_heading_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_hr_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_kbd_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_li_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_object_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_ol_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_p_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_param_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_pre_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_q_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_samp_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_span_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_strong_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_table_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_tbody_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_td_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_tfoot_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_th_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_thead_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_tr_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_ul_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXhtml_var_type.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVE_ID.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.Backup(dirPath)
	backRepo.BackRepoAnyType.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.Backup(dirPath)
	backRepo.BackRepoEMBEDDED_VALUE.Backup(dirPath)
	backRepo.BackRepoENUM_VALUE.Backup(dirPath)
	backRepo.BackRepoRELATION_GROUP.Backup(dirPath)
	backRepo.BackRepoRELATION_GROUP_TYPE.Backup(dirPath)
	backRepo.BackRepoREQ_IF.Backup(dirPath)
	backRepo.BackRepoREQ_IF_CONTENT.Backup(dirPath)
	backRepo.BackRepoREQ_IF_HEADER.Backup(dirPath)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.Backup(dirPath)
	backRepo.BackRepoSPECIFICATION.Backup(dirPath)
	backRepo.BackRepoSPECIFICATION_TYPE.Backup(dirPath)
	backRepo.BackRepoSPEC_HIERARCHY.Backup(dirPath)
	backRepo.BackRepoSPEC_OBJECT.Backup(dirPath)
	backRepo.BackRepoSPEC_OBJECT_TYPE.Backup(dirPath)
	backRepo.BackRepoSPEC_RELATION.Backup(dirPath)
	backRepo.BackRepoSPEC_RELATION_TYPE.Backup(dirPath)
	backRepo.BackRepoXHTML_CONTENT.Backup(dirPath)
	backRepo.BackRepoXhtml_InlPres_type.Backup(dirPath)
	backRepo.BackRepoXhtml_a_type.Backup(dirPath)
	backRepo.BackRepoXhtml_abbr_type.Backup(dirPath)
	backRepo.BackRepoXhtml_acronym_type.Backup(dirPath)
	backRepo.BackRepoXhtml_address_type.Backup(dirPath)
	backRepo.BackRepoXhtml_blockquote_type.Backup(dirPath)
	backRepo.BackRepoXhtml_br_type.Backup(dirPath)
	backRepo.BackRepoXhtml_caption_type.Backup(dirPath)
	backRepo.BackRepoXhtml_cite_type.Backup(dirPath)
	backRepo.BackRepoXhtml_code_type.Backup(dirPath)
	backRepo.BackRepoXhtml_col_type.Backup(dirPath)
	backRepo.BackRepoXhtml_colgroup_type.Backup(dirPath)
	backRepo.BackRepoXhtml_dd_type.Backup(dirPath)
	backRepo.BackRepoXhtml_dfn_type.Backup(dirPath)
	backRepo.BackRepoXhtml_div_type.Backup(dirPath)
	backRepo.BackRepoXhtml_dl_type.Backup(dirPath)
	backRepo.BackRepoXhtml_dt_type.Backup(dirPath)
	backRepo.BackRepoXhtml_edit_type.Backup(dirPath)
	backRepo.BackRepoXhtml_em_type.Backup(dirPath)
	backRepo.BackRepoXhtml_h1_type.Backup(dirPath)
	backRepo.BackRepoXhtml_h2_type.Backup(dirPath)
	backRepo.BackRepoXhtml_h3_type.Backup(dirPath)
	backRepo.BackRepoXhtml_h4_type.Backup(dirPath)
	backRepo.BackRepoXhtml_h5_type.Backup(dirPath)
	backRepo.BackRepoXhtml_h6_type.Backup(dirPath)
	backRepo.BackRepoXhtml_heading_type.Backup(dirPath)
	backRepo.BackRepoXhtml_hr_type.Backup(dirPath)
	backRepo.BackRepoXhtml_kbd_type.Backup(dirPath)
	backRepo.BackRepoXhtml_li_type.Backup(dirPath)
	backRepo.BackRepoXhtml_object_type.Backup(dirPath)
	backRepo.BackRepoXhtml_ol_type.Backup(dirPath)
	backRepo.BackRepoXhtml_p_type.Backup(dirPath)
	backRepo.BackRepoXhtml_param_type.Backup(dirPath)
	backRepo.BackRepoXhtml_pre_type.Backup(dirPath)
	backRepo.BackRepoXhtml_q_type.Backup(dirPath)
	backRepo.BackRepoXhtml_samp_type.Backup(dirPath)
	backRepo.BackRepoXhtml_span_type.Backup(dirPath)
	backRepo.BackRepoXhtml_strong_type.Backup(dirPath)
	backRepo.BackRepoXhtml_table_type.Backup(dirPath)
	backRepo.BackRepoXhtml_tbody_type.Backup(dirPath)
	backRepo.BackRepoXhtml_td_type.Backup(dirPath)
	backRepo.BackRepoXhtml_tfoot_type.Backup(dirPath)
	backRepo.BackRepoXhtml_th_type.Backup(dirPath)
	backRepo.BackRepoXhtml_thead_type.Backup(dirPath)
	backRepo.BackRepoXhtml_tr_type.Backup(dirPath)
	backRepo.BackRepoXhtml_ul_type.Backup(dirPath)
	backRepo.BackRepoXhtml_var_type.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVE_ID.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.BackupXL(file)
	backRepo.BackRepoAnyType.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.BackupXL(file)
	backRepo.BackRepoEMBEDDED_VALUE.BackupXL(file)
	backRepo.BackRepoENUM_VALUE.BackupXL(file)
	backRepo.BackRepoRELATION_GROUP.BackupXL(file)
	backRepo.BackRepoRELATION_GROUP_TYPE.BackupXL(file)
	backRepo.BackRepoREQ_IF.BackupXL(file)
	backRepo.BackRepoREQ_IF_CONTENT.BackupXL(file)
	backRepo.BackRepoREQ_IF_HEADER.BackupXL(file)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.BackupXL(file)
	backRepo.BackRepoSPECIFICATION.BackupXL(file)
	backRepo.BackRepoSPECIFICATION_TYPE.BackupXL(file)
	backRepo.BackRepoSPEC_HIERARCHY.BackupXL(file)
	backRepo.BackRepoSPEC_OBJECT.BackupXL(file)
	backRepo.BackRepoSPEC_OBJECT_TYPE.BackupXL(file)
	backRepo.BackRepoSPEC_RELATION.BackupXL(file)
	backRepo.BackRepoSPEC_RELATION_TYPE.BackupXL(file)
	backRepo.BackRepoXHTML_CONTENT.BackupXL(file)
	backRepo.BackRepoXhtml_InlPres_type.BackupXL(file)
	backRepo.BackRepoXhtml_a_type.BackupXL(file)
	backRepo.BackRepoXhtml_abbr_type.BackupXL(file)
	backRepo.BackRepoXhtml_acronym_type.BackupXL(file)
	backRepo.BackRepoXhtml_address_type.BackupXL(file)
	backRepo.BackRepoXhtml_blockquote_type.BackupXL(file)
	backRepo.BackRepoXhtml_br_type.BackupXL(file)
	backRepo.BackRepoXhtml_caption_type.BackupXL(file)
	backRepo.BackRepoXhtml_cite_type.BackupXL(file)
	backRepo.BackRepoXhtml_code_type.BackupXL(file)
	backRepo.BackRepoXhtml_col_type.BackupXL(file)
	backRepo.BackRepoXhtml_colgroup_type.BackupXL(file)
	backRepo.BackRepoXhtml_dd_type.BackupXL(file)
	backRepo.BackRepoXhtml_dfn_type.BackupXL(file)
	backRepo.BackRepoXhtml_div_type.BackupXL(file)
	backRepo.BackRepoXhtml_dl_type.BackupXL(file)
	backRepo.BackRepoXhtml_dt_type.BackupXL(file)
	backRepo.BackRepoXhtml_edit_type.BackupXL(file)
	backRepo.BackRepoXhtml_em_type.BackupXL(file)
	backRepo.BackRepoXhtml_h1_type.BackupXL(file)
	backRepo.BackRepoXhtml_h2_type.BackupXL(file)
	backRepo.BackRepoXhtml_h3_type.BackupXL(file)
	backRepo.BackRepoXhtml_h4_type.BackupXL(file)
	backRepo.BackRepoXhtml_h5_type.BackupXL(file)
	backRepo.BackRepoXhtml_h6_type.BackupXL(file)
	backRepo.BackRepoXhtml_heading_type.BackupXL(file)
	backRepo.BackRepoXhtml_hr_type.BackupXL(file)
	backRepo.BackRepoXhtml_kbd_type.BackupXL(file)
	backRepo.BackRepoXhtml_li_type.BackupXL(file)
	backRepo.BackRepoXhtml_object_type.BackupXL(file)
	backRepo.BackRepoXhtml_ol_type.BackupXL(file)
	backRepo.BackRepoXhtml_p_type.BackupXL(file)
	backRepo.BackRepoXhtml_param_type.BackupXL(file)
	backRepo.BackRepoXhtml_pre_type.BackupXL(file)
	backRepo.BackRepoXhtml_q_type.BackupXL(file)
	backRepo.BackRepoXhtml_samp_type.BackupXL(file)
	backRepo.BackRepoXhtml_span_type.BackupXL(file)
	backRepo.BackRepoXhtml_strong_type.BackupXL(file)
	backRepo.BackRepoXhtml_table_type.BackupXL(file)
	backRepo.BackRepoXhtml_tbody_type.BackupXL(file)
	backRepo.BackRepoXhtml_td_type.BackupXL(file)
	backRepo.BackRepoXhtml_tfoot_type.BackupXL(file)
	backRepo.BackRepoXhtml_th_type.BackupXL(file)
	backRepo.BackRepoXhtml_thead_type.BackupXL(file)
	backRepo.BackRepoXhtml_tr_type.BackupXL(file)
	backRepo.BackRepoXhtml_ul_type.BackupXL(file)
	backRepo.BackRepoXhtml_var_type.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVE_ID.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.RestorePhaseOne(dirPath)
	backRepo.BackRepoAnyType.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.RestorePhaseOne(dirPath)
	backRepo.BackRepoEMBEDDED_VALUE.RestorePhaseOne(dirPath)
	backRepo.BackRepoENUM_VALUE.RestorePhaseOne(dirPath)
	backRepo.BackRepoRELATION_GROUP.RestorePhaseOne(dirPath)
	backRepo.BackRepoRELATION_GROUP_TYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQ_IF.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQ_IF_CONTENT.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQ_IF_HEADER.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECIFICATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECIFICATION_TYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPEC_HIERARCHY.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPEC_OBJECT.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPEC_OBJECT_TYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPEC_RELATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPEC_RELATION_TYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoXHTML_CONTENT.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_InlPres_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_a_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_abbr_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_acronym_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_address_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_blockquote_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_br_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_caption_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_cite_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_code_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_col_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_colgroup_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_dd_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_dfn_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_div_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_dl_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_dt_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_edit_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_em_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_h1_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_h2_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_h3_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_h4_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_h5_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_h6_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_heading_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_hr_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_kbd_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_li_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_object_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_ol_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_p_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_param_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_pre_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_q_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_samp_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_span_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_strong_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_table_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_tbody_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_td_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_tfoot_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_th_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_thead_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_tr_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_ul_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoXhtml_var_type.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVE_ID.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.RestorePhaseTwo()
	backRepo.BackRepoAnyType.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.RestorePhaseTwo()
	backRepo.BackRepoEMBEDDED_VALUE.RestorePhaseTwo()
	backRepo.BackRepoENUM_VALUE.RestorePhaseTwo()
	backRepo.BackRepoRELATION_GROUP.RestorePhaseTwo()
	backRepo.BackRepoRELATION_GROUP_TYPE.RestorePhaseTwo()
	backRepo.BackRepoREQ_IF.RestorePhaseTwo()
	backRepo.BackRepoREQ_IF_CONTENT.RestorePhaseTwo()
	backRepo.BackRepoREQ_IF_HEADER.RestorePhaseTwo()
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.RestorePhaseTwo()
	backRepo.BackRepoSPECIFICATION.RestorePhaseTwo()
	backRepo.BackRepoSPECIFICATION_TYPE.RestorePhaseTwo()
	backRepo.BackRepoSPEC_HIERARCHY.RestorePhaseTwo()
	backRepo.BackRepoSPEC_OBJECT.RestorePhaseTwo()
	backRepo.BackRepoSPEC_OBJECT_TYPE.RestorePhaseTwo()
	backRepo.BackRepoSPEC_RELATION.RestorePhaseTwo()
	backRepo.BackRepoSPEC_RELATION_TYPE.RestorePhaseTwo()
	backRepo.BackRepoXHTML_CONTENT.RestorePhaseTwo()
	backRepo.BackRepoXhtml_InlPres_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_a_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_abbr_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_acronym_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_address_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_blockquote_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_br_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_caption_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_cite_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_code_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_col_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_colgroup_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_dd_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_dfn_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_div_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_dl_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_dt_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_edit_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_em_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_h1_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_h2_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_h3_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_h4_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_h5_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_h6_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_heading_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_hr_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_kbd_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_li_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_object_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_ol_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_p_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_param_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_pre_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_q_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_samp_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_span_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_strong_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_table_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_tbody_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_td_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_tfoot_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_th_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_thead_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_tr_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_ul_type.RestorePhaseTwo()
	backRepo.BackRepoXhtml_var_type.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVE_ID.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.RestoreXLPhaseOne(file)
	backRepo.BackRepoAnyType.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.RestoreXLPhaseOne(file)
	backRepo.BackRepoEMBEDDED_VALUE.RestoreXLPhaseOne(file)
	backRepo.BackRepoENUM_VALUE.RestoreXLPhaseOne(file)
	backRepo.BackRepoRELATION_GROUP.RestoreXLPhaseOne(file)
	backRepo.BackRepoRELATION_GROUP_TYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQ_IF.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQ_IF_CONTENT.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQ_IF_HEADER.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECIFICATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECIFICATION_TYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPEC_HIERARCHY.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPEC_OBJECT.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPEC_OBJECT_TYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPEC_RELATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPEC_RELATION_TYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoXHTML_CONTENT.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_InlPres_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_a_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_abbr_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_acronym_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_address_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_blockquote_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_br_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_caption_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_cite_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_code_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_col_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_colgroup_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_dd_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_dfn_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_div_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_dl_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_dt_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_edit_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_em_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_h1_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_h2_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_h3_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_h4_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_h5_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_h6_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_heading_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_hr_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_kbd_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_li_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_object_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_ol_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_p_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_param_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_pre_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_q_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_samp_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_span_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_strong_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_table_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_tbody_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_td_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_tfoot_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_th_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_thead_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_tr_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_ul_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoXhtml_var_type.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb(ctx context.Context) <-chan int {
	ch := make(chan int)

	backRepoStruct.subscribersRwMutex.Lock()
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	backRepoStruct.subscribersRwMutex.Unlock()

	// Goroutine to remove subscriber when context is done
	go func() {
		<-ctx.Done()
		backRepoStruct.unsubscribe(ch)
	}()
	return ch
}

// unsubscribe removes a subscriber's channel from the subscribers slice.
func (backRepoStruct *BackRepoStruct) unsubscribe(ch chan int) {
	backRepoStruct.subscribersRwMutex.Lock()
	defer backRepoStruct.subscribersRwMutex.Unlock()
	for i, subscriber := range backRepoStruct.subscribers {
		if subscriber == ch {
			backRepoStruct.subscribers =
				append(backRepoStruct.subscribers[:i],
					backRepoStruct.subscribers[i+1:]...)
			close(ch) // Close the channel to signal completion
			break
		}
	}
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.subscribersRwMutex.RLock()
	subscribers := make([]chan int, len(backRepoStruct.subscribers))
	copy(subscribers, backRepoStruct.subscribers)
	backRepoStruct.subscribersRwMutex.RUnlock()

	for _, ch := range subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			// Successfully sent commit from back
		default:
			// Subscriber is not ready to receive; skip to avoid blocking
		}
	}
}
