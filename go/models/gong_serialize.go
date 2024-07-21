// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

func SerializeStage(stage *StageStruct, filename string) {

	f := excelize.NewFile()
	{
		// insertion point
		SerializeExcelize[ALTERNATIVE_ID](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_BOOLEAN](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_DATE](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_ENUMERATION](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_INTEGER](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_REAL](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_STRING](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_XHTML](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_BOOLEAN](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_DATE](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_ENUMERATION](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_INTEGER](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_REAL](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_STRING](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_XHTML](stage, f)
		SerializeExcelize[AnyType](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_BOOLEAN](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_DATE](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_ENUMERATION](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_INTEGER](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_REAL](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_STRING](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_XHTML](stage, f)
		SerializeExcelize[EMBEDDED_VALUE](stage, f)
		SerializeExcelize[ENUM_VALUE](stage, f)
		SerializeExcelize[RELATION_GROUP](stage, f)
		SerializeExcelize[RELATION_GROUP_TYPE](stage, f)
		SerializeExcelize[REQ_IF](stage, f)
		SerializeExcelize[REQ_IF_CONTENT](stage, f)
		SerializeExcelize[REQ_IF_HEADER](stage, f)
		SerializeExcelize[REQ_IF_TOOL_EXTENSION](stage, f)
		SerializeExcelize[SPECIFICATION](stage, f)
		SerializeExcelize[SPECIFICATION_TYPE](stage, f)
		SerializeExcelize[SPEC_HIERARCHY](stage, f)
		SerializeExcelize[SPEC_OBJECT](stage, f)
		SerializeExcelize[SPEC_OBJECT_TYPE](stage, f)
		SerializeExcelize[SPEC_RELATION](stage, f)
		SerializeExcelize[SPEC_RELATION_TYPE](stage, f)
		SerializeExcelize[XHTML_CONTENT](stage, f)
		SerializeExcelize[Xhtml_InlPres_type](stage, f)
		SerializeExcelize[Xhtml_a_type](stage, f)
		SerializeExcelize[Xhtml_abbr_type](stage, f)
		SerializeExcelize[Xhtml_acronym_type](stage, f)
		SerializeExcelize[Xhtml_address_type](stage, f)
		SerializeExcelize[Xhtml_blockquote_type](stage, f)
		SerializeExcelize[Xhtml_br_type](stage, f)
		SerializeExcelize[Xhtml_caption_type](stage, f)
		SerializeExcelize[Xhtml_cite_type](stage, f)
		SerializeExcelize[Xhtml_code_type](stage, f)
		SerializeExcelize[Xhtml_col_type](stage, f)
		SerializeExcelize[Xhtml_colgroup_type](stage, f)
		SerializeExcelize[Xhtml_dd_type](stage, f)
		SerializeExcelize[Xhtml_dfn_type](stage, f)
		SerializeExcelize[Xhtml_div_type](stage, f)
		SerializeExcelize[Xhtml_dl_type](stage, f)
		SerializeExcelize[Xhtml_dt_type](stage, f)
		SerializeExcelize[Xhtml_edit_type](stage, f)
		SerializeExcelize[Xhtml_em_type](stage, f)
		SerializeExcelize[Xhtml_h1_type](stage, f)
		SerializeExcelize[Xhtml_h2_type](stage, f)
		SerializeExcelize[Xhtml_h3_type](stage, f)
		SerializeExcelize[Xhtml_h4_type](stage, f)
		SerializeExcelize[Xhtml_h5_type](stage, f)
		SerializeExcelize[Xhtml_h6_type](stage, f)
		SerializeExcelize[Xhtml_heading_type](stage, f)
		SerializeExcelize[Xhtml_hr_type](stage, f)
		SerializeExcelize[Xhtml_kbd_type](stage, f)
		SerializeExcelize[Xhtml_li_type](stage, f)
		SerializeExcelize[Xhtml_object_type](stage, f)
		SerializeExcelize[Xhtml_ol_type](stage, f)
		SerializeExcelize[Xhtml_p_type](stage, f)
		SerializeExcelize[Xhtml_param_type](stage, f)
		SerializeExcelize[Xhtml_pre_type](stage, f)
		SerializeExcelize[Xhtml_q_type](stage, f)
		SerializeExcelize[Xhtml_samp_type](stage, f)
		SerializeExcelize[Xhtml_span_type](stage, f)
		SerializeExcelize[Xhtml_strong_type](stage, f)
		SerializeExcelize[Xhtml_table_type](stage, f)
		SerializeExcelize[Xhtml_tbody_type](stage, f)
		SerializeExcelize[Xhtml_td_type](stage, f)
		SerializeExcelize[Xhtml_tfoot_type](stage, f)
		SerializeExcelize[Xhtml_th_type](stage, f)
		SerializeExcelize[Xhtml_thead_type](stage, f)
		SerializeExcelize[Xhtml_tr_type](stage, f)
		SerializeExcelize[Xhtml_ul_type](stage, f)
		SerializeExcelize[Xhtml_var_type](stage, f)
	}

	var tab ExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
		if err := f.SaveAs(filename); err != nil {
			fmt.Println("cannot write xl file : ", err)
		}
	}

}

// Tabulator is an interface for writing to a table strings
type Tabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

func Serialize[Type Gongstruct](stage *StageStruct, tab Tabulator) {
	sheetName := GetGongstructName[Type]()

	// Create a new sheet.
	tab.AddSheet(sheetName)

	headerRowIndex := tab.AddRow(sheetName)
	for colIndex, fieldName := range GetFields[Type]() {
		tab.AddCell(sheetName, headerRowIndex, colIndex, fieldName)
		// f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}

	set := *GetGongstructInstancesSet[Type](stage)
	for instance := range set {
		line := tab.AddRow(sheetName)
		for index, fieldName := range GetFields[Type]() {
			tab.AddCell(sheetName, line, index, GetFieldStringValue(
				any(*instance).(Type), fieldName))
			// f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), GetFieldStringValue(
			// 	any(*instance).(Type), fieldName))
		}
	}
}

type ExcelizeTabulator struct {
	f *excelize.File
}

func (tab *ExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *ExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *ExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *ExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

func SerializeExcelize[Type Gongstruct](stage *StageStruct, f *excelize.File) {
	sheetName := GetGongstructName[Type]()

	// Create a new sheet.
	f.NewSheet(sheetName)

	set := *GetGongstructInstancesSet[Type](stage)
	line := 1

	for index, fieldName := range GetFields[Type]() {
		f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", IntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for instance := range set {
		line = line + 1
		for index, fieldName := range GetFields[Type]() {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), GetFieldStringValue(
				any(*instance).(Type), fieldName))
		}
	}

	// Autofit all columns according to their text content
	cols, err := f.GetCols(sheetName)
	if err != nil {
		log.Panicln("SerializeExcelize")
	}
	for idx, col := range cols {
		largestWidth := 0
		for _, rowCell := range col {
			cellWidth := utf8.RuneCountInString(rowCell) + 2 // + 2 for margin
			if cellWidth > largestWidth {
				largestWidth = cellWidth
			}
		}
		name, err := excelize.ColumnNumberToName(idx + 1)
		if err != nil {
			log.Panicln("SerializeExcelize")
		}
		f.SetColWidth(sheetName, name, name, float64(largestWidth))
	}
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}
