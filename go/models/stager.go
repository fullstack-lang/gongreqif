// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
}

func NewStager(r *gin.Engine, stage *Stage, pathToReqifFile string) (
	stager *Stager,
) {

	stager = new(Stager)

	stager.stage = stage

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

	stage.Commit()

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
				ShowNameInHeader: true,
				Size:             50,
			},
			{
				Size: 50,
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}),
			},
		},
	})

	stager.splitStage.Commit()

	return
}
