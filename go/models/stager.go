// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"encoding/xml"
	"fmt"
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

	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
	ssg_stack "github.com/fullstack-lang/gong/lib/ssg/go/stack"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	load_stack "github.com/fullstack-lang/gong/lib/load/go/stack"

	markdown "github.com/fullstack-lang/gong/lib/markdown/go/models"
	markdown_stack "github.com/fullstack-lang/gong/lib/markdown/go/stack"
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
	UpdateAndCommitSpecificationsMarkdownStage(stager *Stager)
	UpdateAttributeDefinitionNb(
		stager *Stager,
	)
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

	markdownStage *markdown.Stage

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
	Map_id_DATATYPE_DEFINITION_XHTML       map[string]*DATATYPE_DEFINITION_XHTML
	Map_id_DATATYPE_DEFINITION_STRING      map[string]*DATATYPE_DEFINITION_STRING
	Map_id_DATATYPE_DEFINITION_BOOLEAN     map[string]*DATATYPE_DEFINITION_BOOLEAN
	Map_id_DATATYPE_DEFINITION_INTEGER     map[string]*DATATYPE_DEFINITION_INTEGER
	Map_id_DATATYPE_DEFINITION_REAL        map[string]*DATATYPE_DEFINITION_REAL
	Map_id_DATATYPE_DEFINITION_DATE        map[string]*DATATYPE_DEFINITION_DATE
	Map_id_DATATYPE_DEFINITION_ENUMERATION map[string]*DATATYPE_DEFINITION_ENUMERATION

	Map_id_SPEC_OBJECT_TYPE   map[string]*SPEC_OBJECT_TYPE
	Map_id_SPECIFICATION_TYPE map[string]*SPECIFICATION_TYPE
	Map_id_SPEC_OBJECT        map[string]*SPEC_OBJECT

	Map_id_ATTRIBUTE_DEFINITION_XHTML       map[string]*ATTRIBUTE_DEFINITION_XHTML
	Map_id_ATTRIBUTE_DEFINITION_STRING      map[string]*ATTRIBUTE_DEFINITION_STRING
	Map_id_ATTRIBUTE_DEFINITION_BOOLEAN     map[string]*ATTRIBUTE_DEFINITION_BOOLEAN
	Map_id_ATTRIBUTE_DEFINITION_INTEGER     map[string]*ATTRIBUTE_DEFINITION_INTEGER
	Map_id_ATTRIBUTE_DEFINITION_DATE        map[string]*ATTRIBUTE_DEFINITION_DATE
	Map_id_ATTRIBUTE_DEFINITION_REAL        map[string]*ATTRIBUTE_DEFINITION_REAL
	Map_id_ATTRIBUTE_DEFINITION_ENUMERATION map[string]*ATTRIBUTE_DEFINITION_ENUMERATION

	Map_SPEC_OBJECT_TYPE_Spec_nbInstance   map[*SPEC_OBJECT_TYPE]int
	Map_SPECIFICATION_TYPE_Spec_nbInstance map[*SPECIFICATION_TYPE]int
	Map_SPEC_RELATION_TYPE_Spec_nbInstance map[*SPEC_RELATION_TYPE]int

	Map_SPEC_OBJECT_TYPE_isHeading      map[*SPEC_OBJECT_TYPE]bool
	Map_SPEC_OBJECT_TYPE_isNodeExpanded map[*SPEC_OBJECT_TYPE]bool // is the tree node expanded in the UX

	Map_ATTRIBUTE_DEFINITION_XHTML_Spec_nbInstance       map[*ATTRIBUTE_DEFINITION_XHTML]int
	Map_ATTRIBUTE_DEFINITION_STRING_Spec_nbInstance      map[*ATTRIBUTE_DEFINITION_STRING]int
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_Spec_nbInstance     map[*ATTRIBUTE_DEFINITION_BOOLEAN]int
	Map_ATTRIBUTE_DEFINITION_INTEGER_Spec_nbInstance     map[*ATTRIBUTE_DEFINITION_INTEGER]int
	Map_ATTRIBUTE_DEFINITION_DATE_Spec_nbInstance        map[*ATTRIBUTE_DEFINITION_DATE]int
	Map_ATTRIBUTE_DEFINITION_REAL_Spec_nbInstance        map[*ATTRIBUTE_DEFINITION_REAL]int
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_Spec_nbInstance map[*ATTRIBUTE_DEFINITION_ENUMERATION]int

	Map_ENUM_VALUE_Spec_nbInstance map[*ENUM_VALUE]int

	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle       map[*ATTRIBUTE_DEFINITION_XHTML]bool
	Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitle      map[*ATTRIBUTE_DEFINITION_STRING]bool
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle     map[*ATTRIBUTE_DEFINITION_BOOLEAN]bool
	Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle     map[*ATTRIBUTE_DEFINITION_INTEGER]bool
	Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitle        map[*ATTRIBUTE_DEFINITION_DATE]bool
	Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitle        map[*ATTRIBUTE_DEFINITION_REAL]bool
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle map[*ATTRIBUTE_DEFINITION_ENUMERATION]bool

	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTable       map[*ATTRIBUTE_DEFINITION_XHTML]bool
	Map_ATTRIBUTE_DEFINITION_STRING_ShowInTable      map[*ATTRIBUTE_DEFINITION_STRING]bool
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable     map[*ATTRIBUTE_DEFINITION_BOOLEAN]bool
	Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable     map[*ATTRIBUTE_DEFINITION_INTEGER]bool
	Map_ATTRIBUTE_DEFINITION_DATE_ShowInTable        map[*ATTRIBUTE_DEFINITION_DATE]bool
	Map_ATTRIBUTE_DEFINITION_REAL_ShowInTable        map[*ATTRIBUTE_DEFINITION_REAL]bool
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable map[*ATTRIBUTE_DEFINITION_ENUMERATION]bool

	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject       map[*ATTRIBUTE_DEFINITION_XHTML]bool
	Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubject      map[*ATTRIBUTE_DEFINITION_STRING]bool
	Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject     map[*ATTRIBUTE_DEFINITION_BOOLEAN]bool
	Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject     map[*ATTRIBUTE_DEFINITION_INTEGER]bool
	Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubject        map[*ATTRIBUTE_DEFINITION_DATE]bool
	Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubject        map[*ATTRIBUTE_DEFINITION_REAL]bool
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject map[*ATTRIBUTE_DEFINITION_ENUMERATION]bool

	Map_id_ENUM_VALUE map[string]*ENUM_VALUE

	Map_id_SPEC_RELATION_TYPE map[string]*SPEC_RELATION_TYPE

	// attributes for generating static web site
	ssgStage *ssg.Stage

	// for generating the static site, one needs to
	// (1) extract the default files of SSG into a directory
	pathToExtractedGongSsgDefaultStaticFiles string
	pathToExtractedGongSsgDefaultLayoutFiles string

	// (2) copy those files to a directory
	// where images generated by one will be added
	pathToFilesUsedForSsgGeneration string

	// (3) the will serve for the generation of markdown files
	rootPathToImageInputs                   string
	pathToFilesForSsgGeneratedMarkdownFiles string

	// (4) that will be used for the path to the ssg site
	rootPathToStaticOutputs string

	loadStage *load.Stage

	selectedSpecification               *SPECIFICATION
	Map_SPECIFICATION_Nodes_expanded    map[*SPECIFICATION]bool
	Map_SPEC_OBJECT_TYPE_showIdentifier map[*SPEC_OBJECT_TYPE]bool
	Map_SPEC_OBJECT_TYPE_showName       map[*SPEC_OBJECT_TYPE]bool
}

func (stager *Stager) SetSelectedSpecification(selectedSpecification *SPECIFICATION) {
	stager.selectedSpecification = selectedSpecification
}

func (stager *Stager) GetSelectedSpecification() (selectedSpecification *SPECIFICATION) {
	return stager.selectedSpecification
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

func (stager *Stager) GetMarkdownStage() (s *markdown.Stage) {
	s = stager.markdownStage
	return
}

func (stager *Stager) GetSsgStage() (s *ssg.Stage) {
	s = stager.ssgStage
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

	stager.ssgStage = ssg_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage
	stager.loadStage = load_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage
	stager.markdownStage = markdown_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage

	stager.summaryTableStage = table_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage
	stager.summaryTableName = "Summary Table Name"

	stager.dataTypeTreeStage = tree_stack.NewStack(r, stage.GetName()+"-data types", "", "", "", true, true).Stage
	stager.dataTypeTreeName = "Data Type Tree Name"

	stager.specTypesTreeStage = tree_stack.NewStack(r, stage.GetName()+"-spec types", "", "", "", true, true).Stage
	stager.specTypesTreeName = "Spec Object Type Tree Name"

	// instancess
	//

	stager.specObjectsTreeStage = tree_stack.NewStack(r, stage.GetName()+"-spec objects", "", "", "", true, true).Stage
	stager.specTypesTreeName = "Object Tree Name"

	stager.specRelationsTreeStage = tree_stack.NewStack(r, stage.GetName()+"-spec relations", "", "", "", true, true).Stage
	stager.specRelationsTreeName = "Spec Relations Tree Name"

	stager.specificationsTreeStage = tree_stack.NewStack(r, stage.GetName()+"-specifications", "", "", "", true, true).Stage
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
							Name: "Summary + upload + button",
							Size: 20,
							AsSplit: (&split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "Summary table",
										ShowNameInHeader: false,
										Size:             50,
										Table: &split.Table{
											StackName: stager.summaryTableStage.GetName(),
											TableName: stager.summaryTableName,
										},
									},
									(&split.AsSplitArea{
										Name: "Upload",
										Size: 25,
										AsSplit: (&split.AsSplit{
											Direction: split.Horizontal,
											AsSplitAreas: []*split.AsSplitArea{
												load.NewStager(r, stager.loadStage, stager.splitStage).GetAsSplitArea()},
										}),
									}),
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
							Size: 40,
							Tree: &split.Tree{
								StackName: stager.dataTypeTreeStage.GetName(),
								TreeName:  stager.dataTypeTreeName,
							},
						},
						{
							Size: 40,
							Tree: &split.Tree{
								StackName: stager.specTypesTreeStage.GetName(),
								TreeName:  stager.specTypesTreeName,
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
		Name: "REQIF Render",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "room for tree",
				ShowNameInHeader: false,
				Size:             100,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{

						{
							Size: 19.1,
							Tree: &split.Tree{
								StackName: stager.specificationsTreeStage.GetName(),
								TreeName:  stager.specificationsTreeName,
							},
						},
						{
							Size: 19.1,
							Tree: &split.Tree{
								StackName: stager.specTypesTreeStage.GetName(),
								TreeName:  stager.specTypesTreeName,
							},
						},
						{
							Size: 61.8,
							Markdown: &split.Markdown{
								StackName: stager.markdownStage.GetName(),
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
		Name: "(Dev) REQIF Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) summary table probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.summaryTableStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) data type tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.dataTypeTreeStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) spec type tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.specTypesTreeStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "(Dev) markdownStage",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.markdownStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	stager.splitStage.Commit()

	if pathToReqifFile != "" {
		// Open the XML file
		reqifData, err := os.ReadFile(pathToReqifFile) // os.ReadFile is simpler
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", pathToReqifFile, err)
		}
		stager.processReqifData(reqifData, pathToReqifFile)
	}

	stager.updateAndCommitSummaryLoadStage()

	return
}

func (stager *Stager) processReqifData(reqifData []byte, pathToReqifFile string) {

	stager.stage.Reset()
	stager.pathToReqifFile = pathToReqifFile

	// Unmarshal the XML into the Reqif struct
	var req_if REQ_IF
	err := xml.Unmarshal(reqifData, &req_if)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	StageBranch(stager.stage, &req_if)

	// fetch the root REQ_IF element and exit otherwise
	for reqif := range *GetGongstructInstancesSet[REQ_IF](stager.stage) {
		stager.rootReqif = reqif
	}

	if stager.rootReqif == nil {
		log.Fatal("No REQ_IF element found in file", pathToReqifFile)
	}

	stager.objectNamer.SetNamesToElements(stager.stage, &req_if)

	stager.initMaps()

	stager.stage.Commit()

	stager.updateAndCommitSummaryTableStage()

	stager.dataTypesTreeUpdater.UpdateAndCommitDataTypeTreeStage(stager)
	stager.specTypesTreeUpdater.UpdateAndCommitSpecTypesTreeStage(stager)

	stager.specObjectsTreeUpdater.UpdateAndCommitSpecObjectsTreeStage(stager)
	stager.specRelationsTreeUpdater.UpdateAndCommitSpecRelationsTreeStage(stager)
	stager.specificationsTreeUpdater.UpdateAndCommitSpecificationsTreeStage(stager)
	stager.specificationsTreeUpdater.UpdateAndCommitSpecificationsMarkdownStage(stager)

	// stager.UpdateAndCommitButtonStage()

}

func (stager *Stager) GetSpecificationsTreeUpdater() (specificationsTreeUpdater SpecificationsTreeUpdaterInterface) {
	return stager.specificationsTreeUpdater
}

func (stager *Stager) GetSpecTypesTreeUpdater() (specTypesTreeUpdater SpecTypesTreeUpdaterInterface) {
	return stager.specTypesTreeUpdater
}
