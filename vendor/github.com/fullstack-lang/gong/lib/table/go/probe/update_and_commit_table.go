// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/lib/table/go/models"
)

func updateAndCommitTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.Cell:
		updateAndCommitTable[models.Cell](probe)
	case *models.CellBoolean:
		updateAndCommitTable[models.CellBoolean](probe)
	case *models.CellFloat64:
		updateAndCommitTable[models.CellFloat64](probe)
	case *models.CellIcon:
		updateAndCommitTable[models.CellIcon](probe)
	case *models.CellInt:
		updateAndCommitTable[models.CellInt](probe)
	case *models.CellString:
		updateAndCommitTable[models.CellString](probe)
	case *models.CheckBox:
		updateAndCommitTable[models.CheckBox](probe)
	case *models.DisplayedColumn:
		updateAndCommitTable[models.DisplayedColumn](probe)
	case *models.FormDiv:
		updateAndCommitTable[models.FormDiv](probe)
	case *models.FormEditAssocButton:
		updateAndCommitTable[models.FormEditAssocButton](probe)
	case *models.FormField:
		updateAndCommitTable[models.FormField](probe)
	case *models.FormFieldDate:
		updateAndCommitTable[models.FormFieldDate](probe)
	case *models.FormFieldDateTime:
		updateAndCommitTable[models.FormFieldDateTime](probe)
	case *models.FormFieldFloat64:
		updateAndCommitTable[models.FormFieldFloat64](probe)
	case *models.FormFieldInt:
		updateAndCommitTable[models.FormFieldInt](probe)
	case *models.FormFieldSelect:
		updateAndCommitTable[models.FormFieldSelect](probe)
	case *models.FormFieldString:
		updateAndCommitTable[models.FormFieldString](probe)
	case *models.FormFieldTime:
		updateAndCommitTable[models.FormFieldTime](probe)
	case *models.FormGroup:
		updateAndCommitTable[models.FormGroup](probe)
	case *models.FormSortAssocButton:
		updateAndCommitTable[models.FormSortAssocButton](probe)
	case *models.Option:
		updateAndCommitTable[models.Option](probe)
	case *models.Row:
		updateAndCommitTable[models.Row](probe)
	case *models.Table:
		updateAndCommitTable[models.Table](probe)
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
