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

// data types

type DataTypesTreeUpdaterInterface interface {
	UpdateAndCommitDataTypeTreeStage(stager *Stager)
}

// spec types

type SpecTypesTreeUpdaterInterface interface {
	UpdateAndCommitSpecTypesTreeStage(stager *Stager)
}

// instances

type SpecObjectsTreeUpdaterInterface interface {
	UpdateAndCommitSpecObjectsTreeStage(stager *Stager)
}

type SpecRelationsTreeUpdaterInterface interface {
	UpdateAndCommitSpecRelationsTreeStage(stager *Stager)
}

type SpecificationsTreeUpdaterInterface interface {
	UpdateAndCommitSpecificationsTreeStage(stager *Stager)
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

	// types

	dataTypeTreeStage *tree.Stage
	dataTypeTreeName  string

	specTypesTreeStage *tree.Stage
	specTypesTreeName  string

	// instances

	specObjectsTreeStage *tree.Stage
	specObjectsTreeName  string

	specRelationsTreeStage *tree.Stage
	specRelationsTreeName  string

	specificationsTreeStage *tree.Stage
	specificationsTreeName  string

	rootReqif       *REQ_IF
	pathToReqifFile string

	// those interface allows for processing outside the models
	// package. This is for compilation time sake
	// indeed, since the models is quite complex, the compilation
	// of the models package can last an unusual long time
	// for a go compilation. Therefore, processing is delegated to other
	// package where modification will be compiled in a much faster time
	modelGenerator ModelGeneratorInterface

	dataTypesTreeUpdater      DataTypesTreeUpdaterInterface
	specTypesTreeUpdater      SpecTypesTreeUpdaterInterface
	specObjectsTreeUpdater    SpecObjectsTreeUpdaterInterface
	specRelationsTreeUpdater  SpecRelationsTreeUpdaterInterface
	specificationsTreeUpdater SpecificationsTreeUpdaterInterface
	objectNamer               ObjectNamerInterface

	// maps for navigating the ReqIF data
	Map_id_datatypes_xhtml       map[string]*DATATYPE_DEFINITION_XHTML
	Map_id_datatypes_string      map[string]*DATATYPE_DEFINITION_STRING
	Map_id_datatypes_boolean     map[string]*DATATYPE_DEFINITION_BOOLEAN
	Map_id_datatypes_integer     map[string]*DATATYPE_DEFINITION_INTEGER
	Map_id_datatypes_real        map[string]*DATATYPE_DEFINITION_REAL
	Map_id_datatypes_date        map[string]*DATATYPE_DEFINITION_DATE
	Map_id_datatypes_enumeration map[string]*DATATYPE_DEFINITION_ENUMERATION

	Map_id_specobjectTypes   map[string]*SPEC_OBJECT_TYPE
	Map_id_specificationType map[string]*SPECIFICATION_TYPE
	Map_id_specObject        map[string]*SPEC_OBJECT

	Map_id_attributeDefinitionXHTML   map[string]*ATTRIBUTE_DEFINITION_XHTML
	Map_id_attributeDefinitionString  map[string]*ATTRIBUTE_DEFINITION_STRING
	Map_id_attributeDefinitionBoolean map[string]*ATTRIBUTE_DEFINITION_BOOLEAN
	Map_id_attributeDefinitionInteger map[string]*ATTRIBUTE_DEFINITION_INTEGER
	Map_id_attributeDefinitionDate    map[string]*ATTRIBUTE_DEFINITION_DATE
	Map_id_attributeDefinitionReal    map[string]*ATTRIBUTE_DEFINITION_REAL
	Map_id_attributeDefinitionEnum    map[string]*ATTRIBUTE_DEFINITION_ENUMERATION
	Map_id_enumValues                 map[string]*ENUM_VALUE

	Map_id_specRelationType map[string]*SPEC_RELATION_TYPE
}

func (stager *Stager) GetStage() (stage *Stage) {
	stage = stager.stage
	return
}

// Tree for Data Types

func (stager *Stager) GetDataTypeTreeStage() (dataTypeTreeStage *tree.Stage) {
	dataTypeTreeStage = stager.dataTypeTreeStage
	return
}

func (stager *Stager) GetDataTypeTreeName() string {
	return stager.dataTypeTreeName
}

// Tree for spec types

func (stager *Stager) GetSpecTypesTreeStage() (s *tree.Stage) {
	s = stager.specTypesTreeStage
	return
}

func (stager *Stager) GetSpecTypesTreeName() (s string) {
	s = stager.specTypesTreeName
	return
}

// Tree for spec objects

func (stager *Stager) GetSpecObjectsTreeStage() (s *tree.Stage) {
	s = stager.specObjectsTreeStage
	return
}

func (stager *Stager) GetSpecObjectsTreeName() string {
	return stager.specObjectsTreeName
}

// Tree for spec relations

func (stager *Stager) GetSpecRelationsTreeStage() (s *tree.Stage) {
	s = stager.specRelationsTreeStage
	return
}

func (stager *Stager) GetSpecRelationsTreeName() (s string) {
	s = stager.specRelationsTreeName
	return
}

// Tree for specification

func (stager *Stager) GetSpecificationsTreeStage() (s *tree.Stage) {
	s = stager.specificationsTreeStage
	return
}

func (stager *Stager) GetSpecificationsTreeName() (s string) {
	s = stager.specificationsTreeName
	return
}

// OTHERS

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

	dataTypesTreeUpdater DataTypesTreeUpdaterInterface,
	specTypesTreeUpdater SpecTypesTreeUpdaterInterface,

	specObjectsTreeUpdater SpecObjectsTreeUpdaterInterface,
	specRelationsTreeUpdater SpecRelationsTreeUpdaterInterface,
	specificationsTreeUpdater SpecificationsTreeUpdaterInterface,

	objectNamer ObjectNamerInterface) (
	stager *Stager,
) {

	stager = new(Stager)

	stager.stage = stage
	stager.splitStage = splitStage
	stager.pathToReqifFile = pathToReqifFile

	stager.dataTypesTreeUpdater = dataTypesTreeUpdater
	stager.specTypesTreeUpdater = specTypesTreeUpdater
	stager.specObjectsTreeUpdater = specObjectsTreeUpdater
	stager.specRelationsTreeUpdater = specRelationsTreeUpdater
	stager.specificationsTreeUpdater = specificationsTreeUpdater

	stager.objectNamer = objectNamer

	stager.summaryTableStage = table_stack.NewStack(r, stage.GetName(), "", "", "", false, true).Stage
	stager.summaryTableName = "Summary Table Name"

	stager.dataTypeTreeStage = tree_stack.NewStack(r, stage.GetName()+"data types", "", "", "", false, true).Stage
	stager.dataTypeTreeName = "Data Type Tree Name"

	stager.specTypesTreeStage = tree_stack.NewStack(r, stage.GetName()+"spec types", "", "", "", false, true).Stage
	stager.specTypesTreeName = "Spec Object Type Tree Name"

	// instancess
	//

	stager.specObjectsTreeStage = tree_stack.NewStack(r, stage.GetName()+"spec objects", "", "", "", false, true).Stage
	stager.specTypesTreeName = "Object Tree Name"

	stager.specRelationsTreeStage = tree_stack.NewStack(r, stage.GetName()+"spec relations", "", "", "", false, true).Stage
	stager.specRelationsTreeName = "Spec Relations Tree Name"

	stager.specificationsTreeStage = tree_stack.NewStack(r, stage.GetName()+"specifications", "", "", "", false, true).Stage
	stager.specificationsTreeName = "Specifications Tree Name"

	stager.buttonStage = button_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage

	// StageBranch will stage on the the first argument
	// all instances related to the second argument
	split.StageBranch(stager.splitStage, &split.View{
		Name: "REQIF Specification",
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
							Size: 25,
							Tree: &split.Tree{
								StackName: stager.dataTypeTreeStage.GetName(),
								TreeName:  stager.dataTypeTreeName,
							},
						},
						{
							Size: 25,
							Tree: &split.Tree{
								StackName: stager.specTypesTreeStage.GetName(),
								TreeName:  stager.specTypesTreeName,
							},
						},
						{
							Size: 10,
							Tree: &split.Tree{
								StackName: stager.specObjectsTreeStage.GetName(),
								TreeName:  stager.specObjectsTreeName,
							},
						},
						{
							Size: 10,
							Tree: &split.Tree{
								StackName: stager.specRelationsTreeStage.GetName(),
								TreeName:  stager.specRelationsTreeName,
							},
						},
						{
							Size: 10,
							Tree: &split.Tree{
								StackName: stager.specificationsTreeStage.GetName(),
								TreeName:  stager.specificationsTreeName,
							},
						},
					},
				}),
			},
		},
	})

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
							Size: 33,
							Tree: &split.Tree{
								StackName: stager.specObjectsTreeStage.GetName(),
								TreeName:  stager.specObjectsTreeName,
							},
						},
						{
							Size: 33,
							Tree: &split.Tree{
								StackName: stager.specRelationsTreeStage.GetName(),
								TreeName:  stager.specRelationsTreeName,
							},
						},
						{
							Size: 34,
							Tree: &split.Tree{
								StackName: stager.specificationsTreeStage.GetName(),
								TreeName:  stager.specificationsTreeName,
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
					StackName: stager.specTypesTreeStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "specifications tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.specificationsTreeStage.GetProbeSplitStageName(),
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

	stager.initMaps()

	stage.Commit()
	stager.updateAndCommitSummaryTableStage()
	stager.UpdateAndCommitButtonStage()

	stager.dataTypesTreeUpdater.UpdateAndCommitDataTypeTreeStage(stager)
	stager.specTypesTreeUpdater.UpdateAndCommitSpecTypesTreeStage(stager)

	stager.specObjectsTreeUpdater.UpdateAndCommitSpecObjectsTreeStage(stager)
	stager.specRelationsTreeUpdater.UpdateAndCommitSpecRelationsTreeStage(stager)
	stager.specificationsTreeUpdater.UpdateAndCommitSpecificationsTreeStage(stager)

	return
}
