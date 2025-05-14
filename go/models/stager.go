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

	table "github.com/fullstack-lang/gong/lib/table/go/models"
	table_stack "github.com/fullstack-lang/gong/lib/table/go/stack"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage

	buttonStage *button.Stage

	summaryTableStage *table.Stage
	summaryTableName  string

	dataTypeTreeStage *tree.Stage
	dataTypeTreeName  string

	specTypeTreeStage *tree.Stage
	specTypeTreeName  string

	objectTreeStage *tree.Stage
	objectTreeName  string

	rootReqif       *REQ_IF
	pathToReqifFile string

	modelGenerator    ModelGeneratorInterface
	objectTreeUpdater ObjectTreeUpdaterInterface
	objectNamer       ObjectNamerInterface
}

func (stager *Stager) GetObjectTreeStage() (objectTreeStage *tree.Stage) {
	objectTreeStage = stager.objectTreeStage
	return
}

func (stager *Stager) GetObjectTreeName() string {
	return stager.objectTreeName
}

func (stager *Stager) GetRootREQIF() (rootReqif *REQ_IF) {
	rootReqif = stager.rootReqif
	return
}

type ModelGeneratorInterface interface {
	GenerateModels(stager *Stager)
}

type ObjectTreeUpdaterInterface interface {
	UpdateAndCommitObjectTreeStage(stager *Stager)
}

type ObjectNamerInterface interface {
	SetNamesToElements(stage *Stage, reqif *REQ_IF)
}

func NewStager(
	r *gin.Engine,
	splitStage *split.Stage,
	stage *Stage,
	pathToReqifFile string,
	modelGenerator ModelGeneratorInterface,
	objectTreeUpdater ObjectTreeUpdaterInterface,
	objectNamer ObjectNamerInterface) (
	stager *Stager,
) {

	stager = new(Stager)

	stager.stage = stage
	stager.splitStage = splitStage
	stager.pathToReqifFile = pathToReqifFile
	stager.modelGenerator = modelGenerator
	stager.objectTreeUpdater = objectTreeUpdater
	stager.objectNamer = objectNamer

	stager.summaryTableStage = table_stack.NewStack(r, stage.GetName(), "", "", "", false, true).Stage
	stager.summaryTableName = "Summary Table Name"

	stager.dataTypeTreeStage = tree_stack.NewStack(r, stage.GetName()+"datatypes", "", "", "", false, true).Stage
	stager.dataTypeTreeName = "Data Type Tree Name"

	stager.specTypeTreeStage = tree_stack.NewStack(r, stage.GetName()+"spectypes", "", "", "", false, true).Stage
	stager.specTypeTreeName = "Spec Type Tree Name"

	stager.objectTreeStage = tree_stack.NewStack(r, stage.GetName()+"object", "", "", "", false, true).Stage
	stager.specTypeTreeName = "Object Tree Name"

	stager.buttonStage = button_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage

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
							Name: "Summary + button",
							Size: 20,
							AsSplit: (&split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "Summary table",
										ShowNameInHeader: false,
										Size:             75,
										Table: &split.Table{
											StackName: stager.summaryTableStage.GetName(),
											TableName: stager.summaryTableName,
										},
									},
									{
										Name:             "Buttons",
										ShowNameInHeader: false,
										Size:             25,
										Button: &split.Button{
											StackName: stager.buttonStage.GetName(),
										},
									},
								},
							}),
						},

						{
							Size: 15,
							Tree: &split.Tree{
								StackName: stager.dataTypeTreeStage.GetName(),
								TreeName:  stager.dataTypeTreeName,
							},
						},
						{
							Size: 20,
							Tree: &split.Tree{
								StackName: stager.specTypeTreeStage.GetName(),
								TreeName:  stager.specTypeTreeName,
							},
						},
						{
							Size: 25,
							Tree: &split.Tree{
								StackName: stager.objectTreeStage.GetName(),
								TreeName:  stager.objectTreeName,
							},
						},
						{
							Name:             "To be completed",
							ShowNameInHeader: false,

							Size: 15,
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
		Name: "summary table probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.summaryTableStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "data type tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.dataTypeTreeStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "spec type tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.specTypeTreeStage.GetProbeSplitStageName(),
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

	StageBranch(stage, &req_if)

	// fetch the root REQ_IF element and exit otherwise
	for reqif := range *GetGongstructInstancesSet[REQ_IF](stage) {
		stager.rootReqif = reqif
	}

	if stager.rootReqif == nil {
		log.Fatal("No REQ_IF element found in file", pathToReqifFile)
	}

	stager.objectNamer.SetNamesToElements(stage, &req_if)

	stage.Commit()
	stager.updateAndCommitSummaryTableStage()
	stager.updateAndCommit_data_type_tree_stage()
	stager.updateAndCommit_spec_type_tree_stage()
	stager.UpdateAndCommitButtonStage()
	stager.objectTreeUpdater.UpdateAndCommitObjectTreeStage(stager)

	return
}
