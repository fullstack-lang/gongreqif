package models

import (
	table "github.com/fullstack-lang/gong/lib/table/go/models"
)

func (stager *Stager) updateAndCommitSummaryTableStage() {

	stager.summaryTableStage.Reset()

	table.StageBranch(stager.summaryTableStage,
		&table.Table{
			Name: stager.summaryTableName,
			DisplayedColumns: []*table.DisplayedColumn{
				{
					Name: "Property",
				},
				{
					Name: "Value",
				},
			},
			Rows: []*table.Row{
				{
					Cells: []*table.Cell{
						{
							CellString: &table.CellString{
								Value: "Filename",
							},
						},
						{
							CellString: &table.CellString{
								Value: stager.pathToReqifFile,
							},
						},
					},
				},
				{
					Cells: []*table.Cell{
						{
							CellString: &table.CellString{
								Value: "Creation time",
							},
						},
						{
							CellString: &table.CellString{
								Value: stager.rootReqif.THE_HEADER.REQ_IF_HEADER.CREATION_TIME,
							},
						},
					},
				},
				{
					Cells: []*table.Cell{
						{
							CellString: &table.CellString{
								Value: "Tool",
							},
						},
						{
							CellString: &table.CellString{
								Value: stager.rootReqif.THE_HEADER.REQ_IF_HEADER.TITLE,
							},
						},
					},
				},
			},
		})

	stager.summaryTableStage.Commit()
}
