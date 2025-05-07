// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gongreqif/go/models"
)

func updateAndCommitTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		updateAndCommitTable[models.ALTERNATIVE_ID](probe)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_BOOLEAN](probe)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_DATE](probe)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_ENUMERATION](probe)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_INTEGER](probe)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_REAL](probe)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_STRING](probe)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		updateAndCommitTable[models.ATTRIBUTE_DEFINITION_XHTML](probe)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		updateAndCommitTable[models.ATTRIBUTE_VALUE_BOOLEAN](probe)
	case *models.ATTRIBUTE_VALUE_DATE:
		updateAndCommitTable[models.ATTRIBUTE_VALUE_DATE](probe)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		updateAndCommitTable[models.ATTRIBUTE_VALUE_ENUMERATION](probe)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		updateAndCommitTable[models.ATTRIBUTE_VALUE_INTEGER](probe)
	case *models.ATTRIBUTE_VALUE_REAL:
		updateAndCommitTable[models.ATTRIBUTE_VALUE_REAL](probe)
	case *models.ATTRIBUTE_VALUE_STRING:
		updateAndCommitTable[models.ATTRIBUTE_VALUE_STRING](probe)
	case *models.ATTRIBUTE_VALUE_XHTML:
		updateAndCommitTable[models.ATTRIBUTE_VALUE_XHTML](probe)
	case *models.AnyType:
		updateAndCommitTable[models.AnyType](probe)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		updateAndCommitTable[models.DATATYPE_DEFINITION_BOOLEAN](probe)
	case *models.DATATYPE_DEFINITION_DATE:
		updateAndCommitTable[models.DATATYPE_DEFINITION_DATE](probe)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		updateAndCommitTable[models.DATATYPE_DEFINITION_ENUMERATION](probe)
	case *models.DATATYPE_DEFINITION_INTEGER:
		updateAndCommitTable[models.DATATYPE_DEFINITION_INTEGER](probe)
	case *models.DATATYPE_DEFINITION_REAL:
		updateAndCommitTable[models.DATATYPE_DEFINITION_REAL](probe)
	case *models.DATATYPE_DEFINITION_STRING:
		updateAndCommitTable[models.DATATYPE_DEFINITION_STRING](probe)
	case *models.DATATYPE_DEFINITION_XHTML:
		updateAndCommitTable[models.DATATYPE_DEFINITION_XHTML](probe)
	case *models.EMBEDDED_VALUE:
		updateAndCommitTable[models.EMBEDDED_VALUE](probe)
	case *models.ENUM_VALUE:
		updateAndCommitTable[models.ENUM_VALUE](probe)
	case *models.RELATION_GROUP:
		updateAndCommitTable[models.RELATION_GROUP](probe)
	case *models.RELATION_GROUP_TYPE:
		updateAndCommitTable[models.RELATION_GROUP_TYPE](probe)
	case *models.REQ_IF:
		updateAndCommitTable[models.REQ_IF](probe)
	case *models.REQ_IF_CONTENT:
		updateAndCommitTable[models.REQ_IF_CONTENT](probe)
	case *models.REQ_IF_HEADER:
		updateAndCommitTable[models.REQ_IF_HEADER](probe)
	case *models.REQ_IF_TOOL_EXTENSION:
		updateAndCommitTable[models.REQ_IF_TOOL_EXTENSION](probe)
	case *models.SPECIFICATION:
		updateAndCommitTable[models.SPECIFICATION](probe)
	case *models.SPECIFICATION_TYPE:
		updateAndCommitTable[models.SPECIFICATION_TYPE](probe)
	case *models.SPEC_HIERARCHY:
		updateAndCommitTable[models.SPEC_HIERARCHY](probe)
	case *models.SPEC_OBJECT:
		updateAndCommitTable[models.SPEC_OBJECT](probe)
	case *models.SPEC_OBJECT_TYPE:
		updateAndCommitTable[models.SPEC_OBJECT_TYPE](probe)
	case *models.SPEC_RELATION:
		updateAndCommitTable[models.SPEC_RELATION](probe)
	case *models.SPEC_RELATION_TYPE:
		updateAndCommitTable[models.SPEC_RELATION_TYPE](probe)
	case *models.XHTML_CONTENT:
		updateAndCommitTable[models.XHTML_CONTENT](probe)
	case *models.Xhtml_InlPres_type:
		updateAndCommitTable[models.Xhtml_InlPres_type](probe)
	case *models.Xhtml_a_type:
		updateAndCommitTable[models.Xhtml_a_type](probe)
	case *models.Xhtml_abbr_type:
		updateAndCommitTable[models.Xhtml_abbr_type](probe)
	case *models.Xhtml_acronym_type:
		updateAndCommitTable[models.Xhtml_acronym_type](probe)
	case *models.Xhtml_address_type:
		updateAndCommitTable[models.Xhtml_address_type](probe)
	case *models.Xhtml_blockquote_type:
		updateAndCommitTable[models.Xhtml_blockquote_type](probe)
	case *models.Xhtml_br_type:
		updateAndCommitTable[models.Xhtml_br_type](probe)
	case *models.Xhtml_caption_type:
		updateAndCommitTable[models.Xhtml_caption_type](probe)
	case *models.Xhtml_cite_type:
		updateAndCommitTable[models.Xhtml_cite_type](probe)
	case *models.Xhtml_code_type:
		updateAndCommitTable[models.Xhtml_code_type](probe)
	case *models.Xhtml_col_type:
		updateAndCommitTable[models.Xhtml_col_type](probe)
	case *models.Xhtml_colgroup_type:
		updateAndCommitTable[models.Xhtml_colgroup_type](probe)
	case *models.Xhtml_dd_type:
		updateAndCommitTable[models.Xhtml_dd_type](probe)
	case *models.Xhtml_dfn_type:
		updateAndCommitTable[models.Xhtml_dfn_type](probe)
	case *models.Xhtml_div_type:
		updateAndCommitTable[models.Xhtml_div_type](probe)
	case *models.Xhtml_dl_type:
		updateAndCommitTable[models.Xhtml_dl_type](probe)
	case *models.Xhtml_dt_type:
		updateAndCommitTable[models.Xhtml_dt_type](probe)
	case *models.Xhtml_edit_type:
		updateAndCommitTable[models.Xhtml_edit_type](probe)
	case *models.Xhtml_em_type:
		updateAndCommitTable[models.Xhtml_em_type](probe)
	case *models.Xhtml_h1_type:
		updateAndCommitTable[models.Xhtml_h1_type](probe)
	case *models.Xhtml_h2_type:
		updateAndCommitTable[models.Xhtml_h2_type](probe)
	case *models.Xhtml_h3_type:
		updateAndCommitTable[models.Xhtml_h3_type](probe)
	case *models.Xhtml_h4_type:
		updateAndCommitTable[models.Xhtml_h4_type](probe)
	case *models.Xhtml_h5_type:
		updateAndCommitTable[models.Xhtml_h5_type](probe)
	case *models.Xhtml_h6_type:
		updateAndCommitTable[models.Xhtml_h6_type](probe)
	case *models.Xhtml_heading_type:
		updateAndCommitTable[models.Xhtml_heading_type](probe)
	case *models.Xhtml_hr_type:
		updateAndCommitTable[models.Xhtml_hr_type](probe)
	case *models.Xhtml_kbd_type:
		updateAndCommitTable[models.Xhtml_kbd_type](probe)
	case *models.Xhtml_li_type:
		updateAndCommitTable[models.Xhtml_li_type](probe)
	case *models.Xhtml_object_type:
		updateAndCommitTable[models.Xhtml_object_type](probe)
	case *models.Xhtml_ol_type:
		updateAndCommitTable[models.Xhtml_ol_type](probe)
	case *models.Xhtml_p_type:
		updateAndCommitTable[models.Xhtml_p_type](probe)
	case *models.Xhtml_param_type:
		updateAndCommitTable[models.Xhtml_param_type](probe)
	case *models.Xhtml_pre_type:
		updateAndCommitTable[models.Xhtml_pre_type](probe)
	case *models.Xhtml_q_type:
		updateAndCommitTable[models.Xhtml_q_type](probe)
	case *models.Xhtml_samp_type:
		updateAndCommitTable[models.Xhtml_samp_type](probe)
	case *models.Xhtml_span_type:
		updateAndCommitTable[models.Xhtml_span_type](probe)
	case *models.Xhtml_strong_type:
		updateAndCommitTable[models.Xhtml_strong_type](probe)
	case *models.Xhtml_table_type:
		updateAndCommitTable[models.Xhtml_table_type](probe)
	case *models.Xhtml_tbody_type:
		updateAndCommitTable[models.Xhtml_tbody_type](probe)
	case *models.Xhtml_td_type:
		updateAndCommitTable[models.Xhtml_td_type](probe)
	case *models.Xhtml_tfoot_type:
		updateAndCommitTable[models.Xhtml_tfoot_type](probe)
	case *models.Xhtml_th_type:
		updateAndCommitTable[models.Xhtml_th_type](probe)
	case *models.Xhtml_thead_type:
		updateAndCommitTable[models.Xhtml_thead_type](probe)
	case *models.Xhtml_tr_type:
		updateAndCommitTable[models.Xhtml_tr_type](probe)
	case *models.Xhtml_ul_type:
		updateAndCommitTable[models.Xhtml_ul_type](probe)
	case *models.Xhtml_var_type:
		updateAndCommitTable[models.Xhtml_var_type](probe)
	default:
		log.Println("unknow type")
	}
}

const TableName = "Table"

func updateAndCommitTable[T models.Gongstruct](
	probe *Probe,
) {

	probe.tableStage.Reset()

	table := new(gongtable.Table).Stage(probe.tableStage)
	table.Name = TableName
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	reverseFields := models.GetReverseFields[T]()

	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	probe.stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](probe.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return models.GetOrder(probe.stageOfInterest, sliceOfGongStructsSorted[i]) <
			models.GetOrder(probe.stageOfInterest, sliceOfGongStructsSorted[j])
	})

	column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, reverseField := range reverseFields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(probe.tableStage)
		value := models.GetFieldStringValue(*structInstance, "Name")
		row.Name = value.GetValueString()

		updater := NewRowUpdate[T](structInstance, probe)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: int(models.GetOrder(
				probe.stageOfInterest,
				structInstance,
			)),
		}).Stage(probe.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: fmt.Sprintf("Delete Icon %d", models.GetOrder(
				probe.stageOfInterest,
				structInstance,
			)),
			Icon: string(maticons.BUTTON_delete),
		}).Stage(probe.tableStage)
		cellIcon.Impl = NewCellDeleteIconImpl[T](structInstance, probe)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue(*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value.GetValueString()
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			switch value.GongFieldValueType {
			case models.GongFieldValueTypeInt:
				cellInt := (&gongtable.CellInt{
					Name:  name,
					Value: value.GetValueInt(),
				}).Stage(probe.tableStage)
				cell.CellInt = cellInt
			case models.GongFieldValueTypeFloat:
				cellFloat := (&gongtable.CellFloat64{
					Name:  name,
					Value: value.GetValueFloat(),
				}).Stage(probe.tableStage)
				cell.CellFloat64 = cellFloat
			case models.GongFieldValueTypeBool:
				cellBool := (&gongtable.CellBoolean{
					Name:  name,
					Value: value.GetValueBool(),
				}).Stage(probe.tableStage)
				cell.CellBool = cellBool
			default:
				cellString := (&gongtable.CellString{
					Name:  name,
					Value: value.GetValueString(),
				}).Stage(probe.tableStage)
				cell.CellString = cellString

			}
		}
		for _, reverseField := range reverseFields {

			value := models.GetReverseFieldOwnerName(
				probe.stageOfInterest,
				structInstance,
				&reverseField)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(probe.tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.probe = probe
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance *T
	probe    *Probe
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.Stage, row, updatedRow *gongtable.Row) {
	// log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.probe)
}
