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

type ModelGeneratorInterface interface {
	GenerateModels(stager *Stager)
}

type DataTypeTreeUpdaterInterface interface {
	UpdateAndCommitDataTypeTreeStage(stager *Stager)
}

type SpecTreeUpdaterInterface interface {
	UpdateAndCommitSpecTreeStage(stager *Stager)
}

type ObjectTreeUpdaterInterface interface {
	UpdateAndCommitObjectTreeStage(stager *Stager)
}

type ObjectNamerInterface interface {
	SetNamesToElements(stage *Stage, reqif *REQ_IF)
}

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

	// those interface allows for processing outside the models
	// package. This is for compilation time sake
	// indeed, since the models is quite complex, the compilation
	// of the models package can last an unusual long time
	// for a go compilation. Therefore, processing is delegated to other
	// package where modification will be compiled in a much faster time
	modelGenerator      ModelGeneratorInterface
	dataTypeTreeUpdater DataTypeTreeUpdaterInterface
	specsTreeUpdater    SpecTreeUpdaterInterface
	objectTreeUpdater   ObjectTreeUpdaterInterface
	objectNamer         ObjectNamerInterface
}

func (stager *Stager) GetStage() (stage *Stage) {
	stage = stager.stage
	return
}

func (stager *Stager) GetDataTypeTreeStage() (dataTypeTreeStage *tree.Stage) {
	dataTypeTreeStage = stager.dataTypeTreeStage
	return
}

func (stager *Stager) GetDataTypeTreeName() string {
	return stager.dataTypeTreeName
}

func (stager *Stager) GetObjectTreeStage() (objectTreeStage *tree.Stage) {
	objectTreeStage = stager.objectTreeStage
	return
}

func (stager *Stager) GetObjectTreeName() string {
	return stager.objectTreeName
}

func (stager *Stager) GetSpecTreeStage() (specTypeTreeStage *tree.Stage) {
	specTypeTreeStage = stager.specTypeTreeStage
	return
}

func (stager *Stager) GetSpecTreeName() (specTypeTreeName string) {
	specTypeTreeName = stager.specTypeTreeName
	return
}

func (stager *Stager) GetRootREQIF() (rootReqif *REQ_IF) {
	rootReqif = stager.rootReqif
	return
}

func (stager *Stager) SetModelGenerator(modelGenerator ModelGeneratorInterface) {
	stager.modelGenerator = modelGenerator
}

func NewStager(
	r *gin.Engine,
	splitStage *split.Stage,
	stage *Stage,
	pathToReqifFile string,

	dataTypeTreeUpdater DataTypeTreeUpdaterInterface,
	specsTreeUpdater SpecTreeUpdaterInterface,
	objectTreeUpdater ObjectTreeUpdaterInterface,

	objectNamer ObjectNamerInterface) (
	stager *Stager,
) {

	stager = new(Stager)

	stager.stage = stage
	stager.splitStage = splitStage
	stager.pathToReqifFile = pathToReqifFile
	stager.dataTypeTreeUpdater = dataTypeTreeUpdater
	stager.specsTreeUpdater = specsTreeUpdater
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
		Name: "REQIF Data",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "room for tree",
				ShowNameInHeader: false,
				Size:             100,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name: "Summary + button",
							Size: 15,
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
							Size: 15,
							Tree: &split.Tree{
								StackName: stager.specTypeTreeStage.GetName(),
								TreeName:  stager.specTypeTreeName,
							},
						},
						{
							Size: 55,
							Tree: &split.Tree{
								StackName: stager.objectTreeStage.GetName(),
								TreeName:  stager.objectTreeName,
							},
						},
						// {
						// 	Name:             "To be completed",
						// 	ShowNameInHeader: false,

						// 	Size: 15,
						// },
					},
				}),
			},
			{
				Size: 0,
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}),
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "REQIF Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}),
			}),
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
	stager.UpdateAndCommitButtonStage()

	stager.dataTypeTreeUpdater.UpdateAndCommitDataTypeTreeStage(stager)
	stager.specsTreeUpdater.UpdateAndCommitSpecTreeStage(stager)
	stager.objectTreeUpdater.UpdateAndCommitObjectTreeStage(stager)

	return
}
