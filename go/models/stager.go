// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	table "github.com/fullstack-lang/gong/lib/table/go/models"
	table_stack "github.com/fullstack-lang/gong/lib/table/go/stack"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"
)

type Stager struct {
	stage             *Stage
	splitStage        *split.Stage
	summaryTableStage *table.Stage
	summaryTableName  string

	metaTypeTreeStage *tree.Stage
	metaTypeTreeName  string

	rootReqif       *REQ_IF
	pathToReqifFile string
}

func NewStager(r *gin.Engine, stage *Stage, pathToReqifFile string) (
	stager *Stager,
) {

	stager = new(Stager)

	stager.stage = stage
	stager.pathToReqifFile = pathToReqifFile

	stager.summaryTableStage = table_stack.NewStack(r, stage.GetName(), "", "", "", false, true).Stage
	stager.summaryTableName = "Summary Table Name"

	stager.metaTypeTreeStage = tree_stack.NewStack(r, stage.GetName(), "", "", "", false, true).Stage
	stager.metaTypeTreeName = "Summary Meta Tree Name"

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

	// StageBranch will stage on the the first argument
	// all instances related to the second argument
	split.StageBranch(stager.splitStage, &split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "room for tree",
				ShowNameInHeader: false,
				Size:             70,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name:             "Summary table",
							ShowNameInHeader: false,
							Size:             25,
							Table: &split.Table{
								StackName: stager.summaryTableStage.GetName(),
								TableName: stager.summaryTableName,
							},
						},
						{
							Name:             "Summary table",
							ShowNameInHeader: false,
							Size:             25,
							Tree: &split.Tree{
								StackName: stager.metaTypeTreeStage.GetName(),
								TreeName:  stager.metaTypeTreeName,
							},
						},
						{
							Name:             "To be completed",
							ShowNameInHeader: false,
							Size:             50,
						},
					},
				}),
			},
			{
				Size: 30,
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}),
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.summaryTableStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	stager.splitStage.Commit()

	// Open the XML file
	xmlFile, err := os.Open(pathToReqifFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	// Read the XML file
	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the XML into the Reqif struct
	var req_if REQ_IF
	err = xml.Unmarshal(byteValue, &req_if)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	SetNamesToREQIF_Elements(stage)

	StageBranch(stage, &req_if)

	// fetch the root REQ_IF element and exit otherwise
	for reqif := range *GetGongstructInstancesSet[REQ_IF](stage) {
		stager.rootReqif = reqif
	}

	if stager.rootReqif == nil {
		log.Fatal("No REQ_IF element found in file", pathToReqifFile)
	}

	stage.Commit()
	stager.updateAndCommitSummaryTableStage()

	return
}
