// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gongreqif/go/models"
	"github.com/fullstack-lang/gongreqif/go/orm"
)

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		fillUpTable[models.ALTERNATIVE_ID](probe)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		fillUpTable[models.ATTRIBUTE_DEFINITION_BOOLEAN](probe)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		fillUpTable[models.ATTRIBUTE_DEFINITION_DATE](probe)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		fillUpTable[models.ATTRIBUTE_DEFINITION_ENUMERATION](probe)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		fillUpTable[models.ATTRIBUTE_DEFINITION_INTEGER](probe)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		fillUpTable[models.ATTRIBUTE_DEFINITION_REAL](probe)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		fillUpTable[models.ATTRIBUTE_DEFINITION_STRING](probe)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		fillUpTable[models.ATTRIBUTE_DEFINITION_XHTML](probe)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		fillUpTable[models.ATTRIBUTE_VALUE_BOOLEAN](probe)
	case *models.ATTRIBUTE_VALUE_DATE:
		fillUpTable[models.ATTRIBUTE_VALUE_DATE](probe)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		fillUpTable[models.ATTRIBUTE_VALUE_ENUMERATION](probe)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		fillUpTable[models.ATTRIBUTE_VALUE_INTEGER](probe)
	case *models.ATTRIBUTE_VALUE_REAL:
		fillUpTable[models.ATTRIBUTE_VALUE_REAL](probe)
	case *models.ATTRIBUTE_VALUE_STRING:
		fillUpTable[models.ATTRIBUTE_VALUE_STRING](probe)
	case *models.ATTRIBUTE_VALUE_XHTML:
		fillUpTable[models.ATTRIBUTE_VALUE_XHTML](probe)
	case *models.AnyType:
		fillUpTable[models.AnyType](probe)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		fillUpTable[models.DATATYPE_DEFINITION_BOOLEAN](probe)
	case *models.DATATYPE_DEFINITION_DATE:
		fillUpTable[models.DATATYPE_DEFINITION_DATE](probe)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		fillUpTable[models.DATATYPE_DEFINITION_ENUMERATION](probe)
	case *models.DATATYPE_DEFINITION_INTEGER:
		fillUpTable[models.DATATYPE_DEFINITION_INTEGER](probe)
	case *models.DATATYPE_DEFINITION_REAL:
		fillUpTable[models.DATATYPE_DEFINITION_REAL](probe)
	case *models.DATATYPE_DEFINITION_STRING:
		fillUpTable[models.DATATYPE_DEFINITION_STRING](probe)
	case *models.DATATYPE_DEFINITION_XHTML:
		fillUpTable[models.DATATYPE_DEFINITION_XHTML](probe)
	case *models.EMBEDDED_VALUE:
		fillUpTable[models.EMBEDDED_VALUE](probe)
	case *models.ENUM_VALUE:
		fillUpTable[models.ENUM_VALUE](probe)
	case *models.RELATION_GROUP:
		fillUpTable[models.RELATION_GROUP](probe)
	case *models.RELATION_GROUP_TYPE:
		fillUpTable[models.RELATION_GROUP_TYPE](probe)
	case *models.REQ_IF:
		fillUpTable[models.REQ_IF](probe)
	case *models.REQ_IF_CONTENT:
		fillUpTable[models.REQ_IF_CONTENT](probe)
	case *models.REQ_IF_HEADER:
		fillUpTable[models.REQ_IF_HEADER](probe)
	case *models.REQ_IF_TOOL_EXTENSION:
		fillUpTable[models.REQ_IF_TOOL_EXTENSION](probe)
	case *models.SPECIFICATION:
		fillUpTable[models.SPECIFICATION](probe)
	case *models.SPECIFICATION_TYPE:
		fillUpTable[models.SPECIFICATION_TYPE](probe)
	case *models.SPEC_HIERARCHY:
		fillUpTable[models.SPEC_HIERARCHY](probe)
	case *models.SPEC_OBJECT:
		fillUpTable[models.SPEC_OBJECT](probe)
	case *models.SPEC_OBJECT_TYPE:
		fillUpTable[models.SPEC_OBJECT_TYPE](probe)
	case *models.SPEC_RELATION:
		fillUpTable[models.SPEC_RELATION](probe)
	case *models.SPEC_RELATION_TYPE:
		fillUpTable[models.SPEC_RELATION_TYPE](probe)
	case *models.XHTML_CONTENT:
		fillUpTable[models.XHTML_CONTENT](probe)
	case *models.Xhtml_InlPres_type:
		fillUpTable[models.Xhtml_InlPres_type](probe)
	case *models.Xhtml_a_type:
		fillUpTable[models.Xhtml_a_type](probe)
	case *models.Xhtml_abbr_type:
		fillUpTable[models.Xhtml_abbr_type](probe)
	case *models.Xhtml_acronym_type:
		fillUpTable[models.Xhtml_acronym_type](probe)
	case *models.Xhtml_address_type:
		fillUpTable[models.Xhtml_address_type](probe)
	case *models.Xhtml_blockquote_type:
		fillUpTable[models.Xhtml_blockquote_type](probe)
	case *models.Xhtml_br_type:
		fillUpTable[models.Xhtml_br_type](probe)
	case *models.Xhtml_caption_type:
		fillUpTable[models.Xhtml_caption_type](probe)
	case *models.Xhtml_cite_type:
		fillUpTable[models.Xhtml_cite_type](probe)
	case *models.Xhtml_code_type:
		fillUpTable[models.Xhtml_code_type](probe)
	case *models.Xhtml_col_type:
		fillUpTable[models.Xhtml_col_type](probe)
	case *models.Xhtml_colgroup_type:
		fillUpTable[models.Xhtml_colgroup_type](probe)
	case *models.Xhtml_dd_type:
		fillUpTable[models.Xhtml_dd_type](probe)
	case *models.Xhtml_dfn_type:
		fillUpTable[models.Xhtml_dfn_type](probe)
	case *models.Xhtml_div_type:
		fillUpTable[models.Xhtml_div_type](probe)
	case *models.Xhtml_dl_type:
		fillUpTable[models.Xhtml_dl_type](probe)
	case *models.Xhtml_dt_type:
		fillUpTable[models.Xhtml_dt_type](probe)
	case *models.Xhtml_edit_type:
		fillUpTable[models.Xhtml_edit_type](probe)
	case *models.Xhtml_em_type:
		fillUpTable[models.Xhtml_em_type](probe)
	case *models.Xhtml_h1_type:
		fillUpTable[models.Xhtml_h1_type](probe)
	case *models.Xhtml_h2_type:
		fillUpTable[models.Xhtml_h2_type](probe)
	case *models.Xhtml_h3_type:
		fillUpTable[models.Xhtml_h3_type](probe)
	case *models.Xhtml_h4_type:
		fillUpTable[models.Xhtml_h4_type](probe)
	case *models.Xhtml_h5_type:
		fillUpTable[models.Xhtml_h5_type](probe)
	case *models.Xhtml_h6_type:
		fillUpTable[models.Xhtml_h6_type](probe)
	case *models.Xhtml_heading_type:
		fillUpTable[models.Xhtml_heading_type](probe)
	case *models.Xhtml_hr_type:
		fillUpTable[models.Xhtml_hr_type](probe)
	case *models.Xhtml_kbd_type:
		fillUpTable[models.Xhtml_kbd_type](probe)
	case *models.Xhtml_li_type:
		fillUpTable[models.Xhtml_li_type](probe)
	case *models.Xhtml_object_type:
		fillUpTable[models.Xhtml_object_type](probe)
	case *models.Xhtml_ol_type:
		fillUpTable[models.Xhtml_ol_type](probe)
	case *models.Xhtml_p_type:
		fillUpTable[models.Xhtml_p_type](probe)
	case *models.Xhtml_param_type:
		fillUpTable[models.Xhtml_param_type](probe)
	case *models.Xhtml_pre_type:
		fillUpTable[models.Xhtml_pre_type](probe)
	case *models.Xhtml_q_type:
		fillUpTable[models.Xhtml_q_type](probe)
	case *models.Xhtml_samp_type:
		fillUpTable[models.Xhtml_samp_type](probe)
	case *models.Xhtml_span_type:
		fillUpTable[models.Xhtml_span_type](probe)
	case *models.Xhtml_strong_type:
		fillUpTable[models.Xhtml_strong_type](probe)
	case *models.Xhtml_table_type:
		fillUpTable[models.Xhtml_table_type](probe)
	case *models.Xhtml_tbody_type:
		fillUpTable[models.Xhtml_tbody_type](probe)
	case *models.Xhtml_td_type:
		fillUpTable[models.Xhtml_td_type](probe)
	case *models.Xhtml_tfoot_type:
		fillUpTable[models.Xhtml_tfoot_type](probe)
	case *models.Xhtml_th_type:
		fillUpTable[models.Xhtml_th_type](probe)
	case *models.Xhtml_thead_type:
		fillUpTable[models.Xhtml_thead_type](probe)
	case *models.Xhtml_tr_type:
		fillUpTable[models.Xhtml_tr_type](probe)
	case *models.Xhtml_ul_type:
		fillUpTable[models.Xhtml_ul_type](probe)
	case *models.Xhtml_var_type:
		fillUpTable[models.Xhtml_var_type](probe)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	probe *Probe,
) {

	probe.tableStage.Reset()
	probe.tableStage.Commit()

	table := new(gongtable.Table).Stage(probe.tableStage)
	table.Name = "Table"
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
		return orm.GetID(
			probe.stageOfInterest,
			probe.backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
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
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

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
			Value: orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
			),
		}).Stage(probe.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: fmt.Sprintf("Delete Icon %d", orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
			)),
			Icon: string(maticons.BUTTON_delete),
		}).Stage(probe.tableStage)
		cellIcon.Impl = NewCellDeleteIconImpl[T](structInstance, probe)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
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
		for _, reverseField := range reverseFields {

			value := orm.GetReverseFieldOwnerName[T](
				probe.stageOfInterest,
				probe.backRepoOfInterest,
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

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.probe)
}
