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
	case *models.A_ALTERNATIVE_ID:
		updateAndCommitTable[models.A_ALTERNATIVE_ID](probe)
	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_DATE_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_REAL_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_STRING_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		updateAndCommitTable[models.A_ATTRIBUTE_DEFINITION_XHTML_REF](probe)
	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_BOOLEAN](probe)
	case *models.A_ATTRIBUTE_VALUE_DATE:
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_DATE](probe)
	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_ENUMERATION](probe)
	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_INTEGER](probe)
	case *models.A_ATTRIBUTE_VALUE_REAL:
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_REAL](probe)
	case *models.A_ATTRIBUTE_VALUE_STRING:
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_STRING](probe)
	case *models.A_ATTRIBUTE_VALUE_XHTML:
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_XHTML](probe)
	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		updateAndCommitTable[models.A_ATTRIBUTE_VALUE_XHTML_1](probe)
	case *models.A_CHILDREN:
		updateAndCommitTable[models.A_CHILDREN](probe)
	case *models.A_CORE_CONTENT:
		updateAndCommitTable[models.A_CORE_CONTENT](probe)
	case *models.A_DATATYPES:
		updateAndCommitTable[models.A_DATATYPES](probe)
	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_BOOLEAN_REF](probe)
	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_DATE_REF](probe)
	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_ENUMERATION_REF](probe)
	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_INTEGER_REF](probe)
	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_REAL_REF](probe)
	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_STRING_REF](probe)
	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		updateAndCommitTable[models.A_DATATYPE_DEFINITION_XHTML_REF](probe)
	case *models.A_EDITABLE_ATTS:
		updateAndCommitTable[models.A_EDITABLE_ATTS](probe)
	case *models.A_ENUM_VALUE_REF:
		updateAndCommitTable[models.A_ENUM_VALUE_REF](probe)
	case *models.A_OBJECT:
		updateAndCommitTable[models.A_OBJECT](probe)
	case *models.A_PROPERTIES:
		updateAndCommitTable[models.A_PROPERTIES](probe)
	case *models.A_RELATION_GROUP_TYPE_REF:
		updateAndCommitTable[models.A_RELATION_GROUP_TYPE_REF](probe)
	case *models.A_SOURCE_1:
		updateAndCommitTable[models.A_SOURCE_1](probe)
	case *models.A_SOURCE_SPECIFICATION_1:
		updateAndCommitTable[models.A_SOURCE_SPECIFICATION_1](probe)
	case *models.A_SPECIFICATIONS:
		updateAndCommitTable[models.A_SPECIFICATIONS](probe)
	case *models.A_SPECIFICATION_TYPE_REF:
		updateAndCommitTable[models.A_SPECIFICATION_TYPE_REF](probe)
	case *models.A_SPECIFIED_VALUES:
		updateAndCommitTable[models.A_SPECIFIED_VALUES](probe)
	case *models.A_SPEC_ATTRIBUTES:
		updateAndCommitTable[models.A_SPEC_ATTRIBUTES](probe)
	case *models.A_SPEC_OBJECTS:
		updateAndCommitTable[models.A_SPEC_OBJECTS](probe)
	case *models.A_SPEC_OBJECT_TYPE_REF:
		updateAndCommitTable[models.A_SPEC_OBJECT_TYPE_REF](probe)
	case *models.A_SPEC_RELATIONS:
		updateAndCommitTable[models.A_SPEC_RELATIONS](probe)
	case *models.A_SPEC_RELATION_GROUPS:
		updateAndCommitTable[models.A_SPEC_RELATION_GROUPS](probe)
	case *models.A_SPEC_RELATION_REF:
		updateAndCommitTable[models.A_SPEC_RELATION_REF](probe)
	case *models.A_SPEC_RELATION_TYPE_REF:
		updateAndCommitTable[models.A_SPEC_RELATION_TYPE_REF](probe)
	case *models.A_SPEC_TYPES:
		updateAndCommitTable[models.A_SPEC_TYPES](probe)
	case *models.A_THE_HEADER:
		updateAndCommitTable[models.A_THE_HEADER](probe)
	case *models.A_TOOL_EXTENSIONS:
		updateAndCommitTable[models.A_TOOL_EXTENSIONS](probe)
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
	case *models.GeneratedImageMetamodel:
		updateAndCommitTable[models.GeneratedImageMetamodel](probe)
	case *models.Image:
		updateAndCommitTable[models.Image](probe)
	case *models.Paragraph:
		updateAndCommitTable[models.Paragraph](probe)
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
	case *models.StaticWebSite:
		updateAndCommitTable[models.StaticWebSite](probe)
	case *models.StaticWebSiteChapter:
		updateAndCommitTable[models.StaticWebSiteChapter](probe)
	case *models.XHTML_CONTENT:
		updateAndCommitTable[models.XHTML_CONTENT](probe)
	default:
		log.Println("unknow type")
	}
}

const TableName = "Table"

func updateAndCommitTable[T models.Gongstruct](
	probe *Probe,
) {

	probe.tableStage.Reset()

	table := new(gongtable.Table)
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

	column := new(gongtable.DisplayedColumn)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, reverseField := range reverseFields {
		column := new(gongtable.DisplayedColumn)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row)
		value := models.GetFieldStringValue(*structInstance, "Name")
		row.Name = value.GetValueString()

		updater := NewRowUpdate(structInstance, probe)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := &gongtable.Cell{
			Name: "ID",
		}
		row.Cells = append(row.Cells, cell)
		cellInt := &gongtable.CellInt{
			Name: "ID",
			Value: int(models.GetOrder(
				probe.stageOfInterest,
				structInstance,
			)),
		}
		cell.CellInt = cellInt

		cell = &gongtable.Cell{
			Name: "Delete Icon",
		}
		row.Cells = append(row.Cells, cell)
		cellIcon := &gongtable.CellIcon{
			Name: fmt.Sprintf("Delete Icon %d", models.GetOrder(
				probe.stageOfInterest,
				structInstance,
			)),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm tou want to delete this instance ?",
		}
		cellIcon.Impl = NewCellDeleteIconImpl(structInstance, probe)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue(*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value.GetValueString()
			fieldIndex++
			// log.Println(fieldName, value)
			cell := &gongtable.Cell{
				Name: name,
			}
			row.Cells = append(row.Cells, cell)

			switch value.GongFieldValueType {
			case models.GongFieldValueTypeInt:
				cellInt := &gongtable.CellInt{
					Name:  name,
					Value: value.GetValueInt(),
				}
				cell.CellInt = cellInt
			case models.GongFieldValueTypeFloat:
				cellFloat := &gongtable.CellFloat64{
					Name:  name,
					Value: value.GetValueFloat(),
				}
				cell.CellFloat64 = cellFloat
			case models.GongFieldValueTypeBool:
				cellBool := &gongtable.CellBoolean{
					Name:  name,
					Value: value.GetValueBool(),
				}
				cell.CellBool = cellBool
			default:
				cellString := &gongtable.CellString{
					Name:  name,
					Value: value.GetValueString(),
				}
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
			cell := &gongtable.Cell{
				Name: name,
			}
			row.Cells = append(row.Cells, cell)

			cellString := &gongtable.CellString{
				Name:  name,
				Value: value,
			}
			cell.CellString = cellString
		}
	}

	gongtable.StageBranch(probe.tableStage, table)

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
