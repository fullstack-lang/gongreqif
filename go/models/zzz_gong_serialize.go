// generated code - do not edit
package models

import (
	"cmp"
	"fmt"
	"log"
	"slices"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

func SerializeStage(stage *Stage, filename string) {

	f := excelize.NewFile()
	{
		// insertion point
		SerializeExcelizePointerToGongstruct[*ALTERNATIVE_ID](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_DEFINITION_BOOLEAN](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_DEFINITION_DATE](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_DEFINITION_ENUMERATION](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_DEFINITION_INTEGER](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_DEFINITION_REAL](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_DEFINITION_STRING](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_DEFINITION_XHTML](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_VALUE_BOOLEAN](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_VALUE_DATE](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_VALUE_ENUMERATION](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_VALUE_INTEGER](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_VALUE_REAL](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_VALUE_STRING](stage, f)
		SerializeExcelizePointerToGongstruct[*ATTRIBUTE_VALUE_XHTML](stage, f)
		SerializeExcelizePointerToGongstruct[*AnyType](stage, f)
		SerializeExcelizePointerToGongstruct[*DATATYPE_DEFINITION_BOOLEAN](stage, f)
		SerializeExcelizePointerToGongstruct[*DATATYPE_DEFINITION_DATE](stage, f)
		SerializeExcelizePointerToGongstruct[*DATATYPE_DEFINITION_ENUMERATION](stage, f)
		SerializeExcelizePointerToGongstruct[*DATATYPE_DEFINITION_INTEGER](stage, f)
		SerializeExcelizePointerToGongstruct[*DATATYPE_DEFINITION_REAL](stage, f)
		SerializeExcelizePointerToGongstruct[*DATATYPE_DEFINITION_STRING](stage, f)
		SerializeExcelizePointerToGongstruct[*DATATYPE_DEFINITION_XHTML](stage, f)
		SerializeExcelizePointerToGongstruct[*EMBEDDED_VALUE](stage, f)
		SerializeExcelizePointerToGongstruct[*ENUM_VALUE](stage, f)
		SerializeExcelizePointerToGongstruct[*RELATION_GROUP](stage, f)
		SerializeExcelizePointerToGongstruct[*RELATION_GROUP_TYPE](stage, f)
		SerializeExcelizePointerToGongstruct[*REQ_IF](stage, f)
		SerializeExcelizePointerToGongstruct[*REQ_IF_CONTENT](stage, f)
		SerializeExcelizePointerToGongstruct[*REQ_IF_HEADER](stage, f)
		SerializeExcelizePointerToGongstruct[*REQ_IF_TOOL_EXTENSION](stage, f)
		SerializeExcelizePointerToGongstruct[*SPECIFICATION](stage, f)
		SerializeExcelizePointerToGongstruct[*SPECIFICATION_TYPE](stage, f)
		SerializeExcelizePointerToGongstruct[*SPEC_HIERARCHY](stage, f)
		SerializeExcelizePointerToGongstruct[*SPEC_OBJECT](stage, f)
		SerializeExcelizePointerToGongstruct[*SPEC_OBJECT_TYPE](stage, f)
		SerializeExcelizePointerToGongstruct[*SPEC_RELATION](stage, f)
		SerializeExcelizePointerToGongstruct[*SPEC_RELATION_TYPE](stage, f)
		SerializeExcelizePointerToGongstruct[*XHTML_CONTENT](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_InlPres_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_a_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_abbr_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_acronym_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_address_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_blockquote_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_br_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_caption_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_cite_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_code_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_col_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_colgroup_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_dd_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_dfn_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_div_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_dl_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_dt_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_edit_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_em_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_h1_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_h2_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_h3_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_h4_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_h5_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_h6_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_heading_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_hr_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_kbd_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_li_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_object_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_ol_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_p_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_param_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_pre_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_q_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_samp_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_span_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_strong_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_table_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_tbody_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_td_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_tfoot_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_th_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_thead_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_tr_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_ul_type](stage, f)
		SerializeExcelizePointerToGongstruct[*Xhtml_var_type](stage, f)
	}

	// Create a style with wrap text enabled
	wrapStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	_ = wrapStyle
	if err != nil {
		fmt.Println("failed to create style:", err)
		return
	}

	// Create a style with bold text
	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	_ = boldStyle
	if err != nil {
		fmt.Println("failed to create bold style:", err)
		return
	}

	// Get all sheet names
	sheetList := f.GetSheetList()

	for _, sheet := range sheetList {
		// Read all rows of the current sheet
		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Printf("failed to get rows for sheet %q: %v\n", sheet, err)
			continue
		}

		// If there's no data at all, skip this sheet
		if len(rows) == 0 {
			continue
		}

		// The first row of the sheet
		firstRow := rows[0]

		// Track the first and last “used” column in the first row,
		// so we can later apply an AutoFilter from the first to last used col
		var firstUsedColIdx, lastUsedColIdx int

		for colIdx, cellValue := range firstRow {
			if cellValue == "" {
				// Skip columns with empty first-row cells
				continue
			}

			// Convert zero-based colIdx to 1-based for Excelize,
			// then get the column name (A, B, C, etc.)
			colName, err := excelize.ColumnNumberToName(colIdx + 1)
			if err != nil {
				fmt.Printf("failed to convert column number: %v\n", err)
				continue
			}

			// Apply wrap-text style to this entire column
			colRange := colName + ":" + colName
			if err := f.SetColStyle(sheet, colRange, wrapStyle); err != nil {
				fmt.Printf("failed to set col style on %s: %v\n", colRange, err)
				continue
			}

			// Make the first row (cell in row 1) bold in this column
			cellRef := fmt.Sprintf("%s1", colName)
			if err := f.SetCellStyle(sheet, cellRef, cellRef, boldStyle); err != nil {
				fmt.Printf("failed to set cell style on %s: %v\n", cellRef, err)
				continue
			}

			// Update our “first used” and “last used” column indices
			if firstUsedColIdx == 0 {
				firstUsedColIdx = colIdx + 1
			}
			if colIdx+1 > lastUsedColIdx {
				lastUsedColIdx = colIdx + 1
			}
		}

		// If we found at least one non-empty column in row 1, enable AutoFilter
		if firstUsedColIdx != 0 && lastUsedColIdx >= firstUsedColIdx {
			startCol, _ := excelize.ColumnNumberToName(firstUsedColIdx)
			endCol, _ := excelize.ColumnNumberToName(lastUsedColIdx)
			styleRange := fmt.Sprintf("%s:%s", startCol, endCol)
			autoFilterRange := fmt.Sprintf("%s1:%s1", startCol, endCol)
			startCellString := fmt.Sprintf("%s1", startCol)
			endCellString := fmt.Sprintf("%s1", endCol)

			if err := f.SetColStyle(sheet, styleRange, wrapStyle); err != nil {
				fmt.Println("failed to set column style:", err)
				return
			}

			// Apply the bold style to the first row (A1:XFD1)
			if err := f.SetCellStyle(sheet, startCellString, endCellString, boldStyle); err != nil {
				fmt.Println("failed to set bold style:", err)
				return
			}

			var opts []excelize.AutoFilterOptions
			if err := f.AutoFilter(sheet, autoFilterRange, opts); err != nil {
				fmt.Printf("failed to enable auto filter on range %s: %v\n", autoFilterRange, err)
			}
		}
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

func Serialize[Type Gongstruct](stage *Stage, tab Tabulator) {
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
				any(*instance).(Type), fieldName).valueString)
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

func SerializeExcelizePointerToGongstruct[Type PointerToGongstruct](stage *Stage, f *excelize.File) {
	sheetName := GetPointerToGongstructName[Type]()

	// Create a new sheet.
	f.NewSheet(sheetName)

	set := *GetGongstructInstancesSetFromPointerType[Type](stage)

	var sortedSlice []Type
	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, func(a, b Type) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	line := 1

	for index, fieldName := range GetFieldsFromPointer[Type]() {
		f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", IntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for _, instance := range sortedSlice {
		line = line + 1
		for index, fieldName := range GetFieldsFromPointer[Type]() {
			fieldStringValue := GetFieldStringValueFromPointer(instance, fieldName)
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
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

func SerializeExcelize[Type Gongstruct](stage *Stage, f *excelize.File) {
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
			fieldStringValue := GetFieldStringValue(any(*instance).(Type), fieldName)
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
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
