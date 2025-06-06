package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongreqif/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Classdiagram__000000_Top_Level := (&models.Classdiagram{}).Stage(stage)
	__Classdiagram__000001_Datatypes := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_REQ_IF := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_REQ_IF_CONTENT := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_REQ_IF_HEADER := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_A_THE_HEADER := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_Default_A_CORE_CONTENT := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000005_Content_REQ_IF_CONTENT := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000006_Content_A_DATATYPES := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_CORE_CONTENT := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_THE_HEADER := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_REQ_IF_HEADER := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000003_REQ_IF_CONTENT := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000004_DATATYPES := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000006_DATATYPE_DEFINITION_DATE := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000009_DATATYPE_DEFINITION_REAL := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000010_DATATYPE_DEFINITION_STRING := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Top_Level.Name = `Top Level`
	__Classdiagram__000000_Top_Level.Description = ``
	__Classdiagram__000000_Top_Level.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Top_Level.IsInRenameMode = false
	__Classdiagram__000000_Top_Level.IsExpanded = true
	__Classdiagram__000000_Top_Level.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Top_Level.NodeGongStructNodeExpansion = `[false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true]`
	__Classdiagram__000000_Top_Level.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Top_Level.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Top_Level.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Top_Level.NodeGongNoteNodeExpansion = ``

	__Classdiagram__000001_Datatypes.Name = `Datatypes`
	__Classdiagram__000001_Datatypes.Description = ``
	__Classdiagram__000001_Datatypes.IsIncludedInStaticWebSite = false
	__Classdiagram__000001_Datatypes.IsInRenameMode = false
	__Classdiagram__000001_Datatypes.IsExpanded = false
	__Classdiagram__000001_Datatypes.NodeGongStructsIsExpanded = true
	__Classdiagram__000001_Datatypes.NodeGongStructNodeExpansion = `[false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false]`
	__Classdiagram__000001_Datatypes.NodeGongEnumsIsExpanded = false
	__Classdiagram__000001_Datatypes.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000001_Datatypes.NodeGongNotesIsExpanded = false
	__Classdiagram__000001_Datatypes.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.Name = `Diagram Package created the 2025-05-07T07:33:18Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_REQ_IF.Name = `Default-REQ_IF`
	__GongStructShape__000000_Default_REQ_IF.X = 9.000000
	__GongStructShape__000000_Default_REQ_IF.Y = 63.000000

	//gong:ident [ref_models.REQ_IF] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_REQ_IF.Identifier = `ref_models.REQ_IF`
	__GongStructShape__000000_Default_REQ_IF.IdentifierMeta = ref_models.REQ_IF{}
	__GongStructShape__000000_Default_REQ_IF.ShowNbInstances = false
	__GongStructShape__000000_Default_REQ_IF.NbInstances = 0
	__GongStructShape__000000_Default_REQ_IF.Width = 240.000000
	__GongStructShape__000000_Default_REQ_IF.Height = 63.000000
	__GongStructShape__000000_Default_REQ_IF.IsSelected = false

	__GongStructShape__000001_Default_REQ_IF_CONTENT.Name = `Default-REQ_IF_CONTENT`
	__GongStructShape__000001_Default_REQ_IF_CONTENT.X = 790.000000
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Y = 333.000000

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Identifier = `ref_models.REQ_IF_CONTENT`
	__GongStructShape__000001_Default_REQ_IF_CONTENT.IdentifierMeta = ref_models.REQ_IF_CONTENT{}
	__GongStructShape__000001_Default_REQ_IF_CONTENT.ShowNbInstances = false
	__GongStructShape__000001_Default_REQ_IF_CONTENT.NbInstances = 0
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Width = 240.000000
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Height = 63.000000
	__GongStructShape__000001_Default_REQ_IF_CONTENT.IsSelected = false

	__GongStructShape__000002_Default_REQ_IF_HEADER.Name = `Default-REQ_IF_HEADER`
	__GongStructShape__000002_Default_REQ_IF_HEADER.X = 787.000000
	__GongStructShape__000002_Default_REQ_IF_HEADER.Y = 69.000000

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_REQ_IF_HEADER.Identifier = `ref_models.REQ_IF_HEADER`
	__GongStructShape__000002_Default_REQ_IF_HEADER.IdentifierMeta = ref_models.REQ_IF_HEADER{}
	__GongStructShape__000002_Default_REQ_IF_HEADER.ShowNbInstances = false
	__GongStructShape__000002_Default_REQ_IF_HEADER.NbInstances = 0
	__GongStructShape__000002_Default_REQ_IF_HEADER.Width = 240.000000
	__GongStructShape__000002_Default_REQ_IF_HEADER.Height = 63.000000
	__GongStructShape__000002_Default_REQ_IF_HEADER.IsSelected = false

	__GongStructShape__000003_Default_A_THE_HEADER.Name = `Default-A_THE_HEADER`
	__GongStructShape__000003_Default_A_THE_HEADER.X = 385.000000
	__GongStructShape__000003_Default_A_THE_HEADER.Y = 64.000000

	//gong:ident [ref_models.A_THE_HEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_A_THE_HEADER.Identifier = `ref_models.A_THE_HEADER`
	__GongStructShape__000003_Default_A_THE_HEADER.IdentifierMeta = ref_models.A_THE_HEADER{}
	__GongStructShape__000003_Default_A_THE_HEADER.ShowNbInstances = false
	__GongStructShape__000003_Default_A_THE_HEADER.NbInstances = 0
	__GongStructShape__000003_Default_A_THE_HEADER.Width = 240.000000
	__GongStructShape__000003_Default_A_THE_HEADER.Height = 63.000000
	__GongStructShape__000003_Default_A_THE_HEADER.IsSelected = false

	__GongStructShape__000004_Default_A_CORE_CONTENT.Name = `Default-A_CORE_CONTENT`
	__GongStructShape__000004_Default_A_CORE_CONTENT.X = 383.000000
	__GongStructShape__000004_Default_A_CORE_CONTENT.Y = 329.000000

	//gong:ident [ref_models.A_CORE_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_A_CORE_CONTENT.Identifier = `ref_models.A_CORE_CONTENT`
	__GongStructShape__000004_Default_A_CORE_CONTENT.IdentifierMeta = ref_models.A_CORE_CONTENT{}
	__GongStructShape__000004_Default_A_CORE_CONTENT.ShowNbInstances = false
	__GongStructShape__000004_Default_A_CORE_CONTENT.NbInstances = 0
	__GongStructShape__000004_Default_A_CORE_CONTENT.Width = 240.000000
	__GongStructShape__000004_Default_A_CORE_CONTENT.Height = 63.000000
	__GongStructShape__000004_Default_A_CORE_CONTENT.IsSelected = false

	__GongStructShape__000005_Content_REQ_IF_CONTENT.Name = `Content-REQ_IF_CONTENT`
	__GongStructShape__000005_Content_REQ_IF_CONTENT.X = 71.000000
	__GongStructShape__000005_Content_REQ_IF_CONTENT.Y = 21.000000

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Content_REQ_IF_CONTENT.Identifier = `ref_models.REQ_IF_CONTENT`
	__GongStructShape__000005_Content_REQ_IF_CONTENT.IdentifierMeta = ref_models.REQ_IF_CONTENT{}
	__GongStructShape__000005_Content_REQ_IF_CONTENT.ShowNbInstances = false
	__GongStructShape__000005_Content_REQ_IF_CONTENT.NbInstances = 0
	__GongStructShape__000005_Content_REQ_IF_CONTENT.Width = 240.000000
	__GongStructShape__000005_Content_REQ_IF_CONTENT.Height = 63.000000
	__GongStructShape__000005_Content_REQ_IF_CONTENT.IsSelected = false

	__GongStructShape__000006_Content_A_DATATYPES.Name = `Content-A_DATATYPES`
	__GongStructShape__000006_Content_A_DATATYPES.X = 59.800000
	__GongStructShape__000006_Content_A_DATATYPES.Y = 305.600000

	//gong:ident [ref_models.A_DATATYPES] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_Content_A_DATATYPES.Identifier = `ref_models.A_DATATYPES`
	__GongStructShape__000006_Content_A_DATATYPES.IdentifierMeta = ref_models.A_DATATYPES{}
	__GongStructShape__000006_Content_A_DATATYPES.ShowNbInstances = false
	__GongStructShape__000006_Content_A_DATATYPES.NbInstances = 0
	__GongStructShape__000006_Content_A_DATATYPES.Width = 240.000000
	__GongStructShape__000006_Content_A_DATATYPES.Height = 63.000000
	__GongStructShape__000006_Content_A_DATATYPES.IsSelected = false

	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN.Name = `Content-DATATYPE_DEFINITION_BOOLEAN`
	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN.X = 907.800000
	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN.Y = 471.800000

	//gong:ident [ref_models.DATATYPE_DEFINITION_BOOLEAN] comment added to overcome the problem with the comment map association
	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN.Identifier = `ref_models.DATATYPE_DEFINITION_BOOLEAN`
	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN.IdentifierMeta = ref_models.DATATYPE_DEFINITION_BOOLEAN{}
	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN.ShowNbInstances = false
	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN.NbInstances = 0
	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN.Width = 358.000000
	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN.Height = 63.000000
	__GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN.IsSelected = false

	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE.Name = `Content-DATATYPE_DEFINITION_DATE`
	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE.X = 906.799939
	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE.Y = 288.800000

	//gong:ident [ref_models.DATATYPE_DEFINITION_DATE] comment added to overcome the problem with the comment map association
	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE.Identifier = `ref_models.DATATYPE_DEFINITION_DATE`
	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE.IdentifierMeta = ref_models.DATATYPE_DEFINITION_DATE{}
	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE.ShowNbInstances = false
	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE.NbInstances = 0
	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE.Width = 367.000000
	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE.Height = 63.000000
	__GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE.IsSelected = false

	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION.Name = `Content-DATATYPE_DEFINITION_ENUMERATION`
	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION.X = 908.600000
	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION.Y = 378.600000

	//gong:ident [ref_models.DATATYPE_DEFINITION_ENUMERATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION.Identifier = `ref_models.DATATYPE_DEFINITION_ENUMERATION`
	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION.IdentifierMeta = ref_models.DATATYPE_DEFINITION_ENUMERATION{}
	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION.ShowNbInstances = false
	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION.NbInstances = 0
	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION.Width = 356.000000
	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION.Height = 63.000000
	__GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION.IsSelected = false

	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER.Name = `Content-DATATYPE_DEFINITION_INTEGER`
	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER.X = 909.599939
	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER.Y = 568.400000

	//gong:ident [ref_models.DATATYPE_DEFINITION_INTEGER] comment added to overcome the problem with the comment map association
	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER.Identifier = `ref_models.DATATYPE_DEFINITION_INTEGER`
	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER.IdentifierMeta = ref_models.DATATYPE_DEFINITION_INTEGER{}
	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER.ShowNbInstances = false
	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER.NbInstances = 0
	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER.Width = 368.000000
	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER.Height = 63.000000
	__GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER.IsSelected = false

	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL.Name = `Content-DATATYPE_DEFINITION_REAL`
	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL.X = 904.400000
	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL.Y = 201.800000

	//gong:ident [ref_models.DATATYPE_DEFINITION_REAL] comment added to overcome the problem with the comment map association
	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL.Identifier = `ref_models.DATATYPE_DEFINITION_REAL`
	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL.IdentifierMeta = ref_models.DATATYPE_DEFINITION_REAL{}
	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL.ShowNbInstances = false
	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL.NbInstances = 0
	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL.Width = 374.000000
	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL.Height = 63.000000
	__GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL.IsSelected = false

	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING.Name = `Content-DATATYPE_DEFINITION_STRING`
	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING.X = 901.400000
	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING.Y = 111.600000

	//gong:ident [ref_models.DATATYPE_DEFINITION_STRING] comment added to overcome the problem with the comment map association
	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING.Identifier = `ref_models.DATATYPE_DEFINITION_STRING`
	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING.IdentifierMeta = ref_models.DATATYPE_DEFINITION_STRING{}
	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING.ShowNbInstances = false
	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING.NbInstances = 0
	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING.Width = 373.000000
	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING.Height = 63.000000
	__GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING.IsSelected = false

	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML.Name = `Content-DATATYPE_DEFINITION_XHTML`
	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML.X = 901.400000
	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML.Y = 24.600000

	//gong:ident [ref_models.DATATYPE_DEFINITION_XHTML] comment added to overcome the problem with the comment map association
	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML.Identifier = `ref_models.DATATYPE_DEFINITION_XHTML`
	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML.IdentifierMeta = ref_models.DATATYPE_DEFINITION_XHTML{}
	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML.ShowNbInstances = false
	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML.NbInstances = 0
	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML.Width = 374.000000
	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML.Height = 63.000000
	__GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML.IsSelected = false

	__LinkShape__000000_CORE_CONTENT.Name = `CORE_CONTENT`

	//gong:ident [ref_models.REQ_IF.CORE_CONTENT] comment added to overcome the problem with the comment map association
	__LinkShape__000000_CORE_CONTENT.Identifier = `ref_models.REQ_IF.CORE_CONTENT`
	__LinkShape__000000_CORE_CONTENT.IdentifierMeta = ref_models.REQ_IF{}.CORE_CONTENT

	//gong:ident [ref_models.A_CORE_CONTENT] comment added to overcome the problem with the comment map association
	__LinkShape__000000_CORE_CONTENT.Fieldtypename = `ref_models.A_CORE_CONTENT`
	__LinkShape__000000_CORE_CONTENT.FieldOffsetX = 0.000000
	__LinkShape__000000_CORE_CONTENT.FieldOffsetY = 0.000000
	__LinkShape__000000_CORE_CONTENT.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000000_CORE_CONTENT.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_CORE_CONTENT.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_CORE_CONTENT.SourceMultiplicity = models.MANY
	__LinkShape__000000_CORE_CONTENT.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_CORE_CONTENT.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_CORE_CONTENT.X = 556.000000
	__LinkShape__000000_CORE_CONTENT.Y = 227.500000
	__LinkShape__000000_CORE_CONTENT.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_CORE_CONTENT.StartRatio = 0.500000
	__LinkShape__000000_CORE_CONTENT.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_CORE_CONTENT.EndRatio = 0.500000
	__LinkShape__000000_CORE_CONTENT.CornerOffsetRatio = 1.380000

	__LinkShape__000001_THE_HEADER.Name = `THE_HEADER`

	//gong:ident [ref_models.REQ_IF.THE_HEADER] comment added to overcome the problem with the comment map association
	__LinkShape__000001_THE_HEADER.Identifier = `ref_models.REQ_IF.THE_HEADER`
	__LinkShape__000001_THE_HEADER.IdentifierMeta = ref_models.REQ_IF{}.THE_HEADER

	//gong:ident [ref_models.A_THE_HEADER] comment added to overcome the problem with the comment map association
	__LinkShape__000001_THE_HEADER.Fieldtypename = `ref_models.A_THE_HEADER`
	__LinkShape__000001_THE_HEADER.FieldOffsetX = 0.000000
	__LinkShape__000001_THE_HEADER.FieldOffsetY = 0.000000
	__LinkShape__000001_THE_HEADER.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000001_THE_HEADER.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_THE_HEADER.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_THE_HEADER.SourceMultiplicity = models.MANY
	__LinkShape__000001_THE_HEADER.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_THE_HEADER.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_THE_HEADER.X = 557.000000
	__LinkShape__000001_THE_HEADER.Y = 95.000000
	__LinkShape__000001_THE_HEADER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_THE_HEADER.StartRatio = 0.500000
	__LinkShape__000001_THE_HEADER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_THE_HEADER.EndRatio = 0.500000
	__LinkShape__000001_THE_HEADER.CornerOffsetRatio = 1.380000

	__LinkShape__000002_REQ_IF_HEADER.Name = `REQ_IF_HEADER`

	//gong:ident [ref_models.A_THE_HEADER.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__LinkShape__000002_REQ_IF_HEADER.Identifier = `ref_models.A_THE_HEADER.REQ_IF_HEADER`
	__LinkShape__000002_REQ_IF_HEADER.IdentifierMeta = ref_models.A_THE_HEADER{}.REQ_IF_HEADER

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__LinkShape__000002_REQ_IF_HEADER.Fieldtypename = `ref_models.REQ_IF_HEADER`
	__LinkShape__000002_REQ_IF_HEADER.FieldOffsetX = 0.000000
	__LinkShape__000002_REQ_IF_HEADER.FieldOffsetY = 0.000000
	__LinkShape__000002_REQ_IF_HEADER.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000002_REQ_IF_HEADER.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_REQ_IF_HEADER.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_REQ_IF_HEADER.SourceMultiplicity = models.MANY
	__LinkShape__000002_REQ_IF_HEADER.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_REQ_IF_HEADER.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_REQ_IF_HEADER.X = 946.000000
	__LinkShape__000002_REQ_IF_HEADER.Y = 98.000000
	__LinkShape__000002_REQ_IF_HEADER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_REQ_IF_HEADER.StartRatio = 0.500000
	__LinkShape__000002_REQ_IF_HEADER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_REQ_IF_HEADER.EndRatio = 0.500000
	__LinkShape__000002_REQ_IF_HEADER.CornerOffsetRatio = 1.380000

	__LinkShape__000003_REQ_IF_CONTENT.Name = `REQ_IF_CONTENT`

	//gong:ident [ref_models.A_CORE_CONTENT.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__LinkShape__000003_REQ_IF_CONTENT.Identifier = `ref_models.A_CORE_CONTENT.REQ_IF_CONTENT`
	__LinkShape__000003_REQ_IF_CONTENT.IdentifierMeta = ref_models.A_CORE_CONTENT{}.REQ_IF_CONTENT

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__LinkShape__000003_REQ_IF_CONTENT.Fieldtypename = `ref_models.REQ_IF_CONTENT`
	__LinkShape__000003_REQ_IF_CONTENT.FieldOffsetX = 0.000000
	__LinkShape__000003_REQ_IF_CONTENT.FieldOffsetY = 0.000000
	__LinkShape__000003_REQ_IF_CONTENT.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000003_REQ_IF_CONTENT.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000003_REQ_IF_CONTENT.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000003_REQ_IF_CONTENT.SourceMultiplicity = models.MANY
	__LinkShape__000003_REQ_IF_CONTENT.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000003_REQ_IF_CONTENT.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000003_REQ_IF_CONTENT.X = 946.500000
	__LinkShape__000003_REQ_IF_CONTENT.Y = 362.500000
	__LinkShape__000003_REQ_IF_CONTENT.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_REQ_IF_CONTENT.StartRatio = 0.500000
	__LinkShape__000003_REQ_IF_CONTENT.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_REQ_IF_CONTENT.EndRatio = 0.500000
	__LinkShape__000003_REQ_IF_CONTENT.CornerOffsetRatio = 1.380000

	__LinkShape__000004_DATATYPES.Name = `DATATYPES`

	//gong:ident [ref_models.REQ_IF_CONTENT.DATATYPES] comment added to overcome the problem with the comment map association
	__LinkShape__000004_DATATYPES.Identifier = `ref_models.REQ_IF_CONTENT.DATATYPES`
	__LinkShape__000004_DATATYPES.IdentifierMeta = ref_models.REQ_IF_CONTENT{}.DATATYPES

	//gong:ident [ref_models.A_DATATYPES] comment added to overcome the problem with the comment map association
	__LinkShape__000004_DATATYPES.Fieldtypename = `ref_models.A_DATATYPES`
	__LinkShape__000004_DATATYPES.FieldTypeIdentifierMeta = ref_models.A_DATATYPES{}
	__LinkShape__000004_DATATYPES.FieldOffsetX = 0.000000
	__LinkShape__000004_DATATYPES.FieldOffsetY = 0.000000
	__LinkShape__000004_DATATYPES.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000004_DATATYPES.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000004_DATATYPES.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000004_DATATYPES.SourceMultiplicity = models.MANY
	__LinkShape__000004_DATATYPES.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000004_DATATYPES.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000004_DATATYPES.X = 628.400000
	__LinkShape__000004_DATATYPES.Y = 52.800000
	__LinkShape__000004_DATATYPES.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000004_DATATYPES.StartRatio = 0.528212
	__LinkShape__000004_DATATYPES.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000004_DATATYPES.EndRatio = 0.566545
	__LinkShape__000004_DATATYPES.CornerOffsetRatio = 2.162698

	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.Name = `DATATYPE_DEFINITION_BOOLEAN`

	//gong:ident [ref_models.A_DATATYPES.DATATYPE_DEFINITION_BOOLEAN] comment added to overcome the problem with the comment map association
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.Identifier = `ref_models.A_DATATYPES.DATATYPE_DEFINITION_BOOLEAN`
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.IdentifierMeta = ref_models.A_DATATYPES{}.DATATYPE_DEFINITION_BOOLEAN

	//gong:ident [ref_models.DATATYPE_DEFINITION_BOOLEAN] comment added to overcome the problem with the comment map association
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.Fieldtypename = `ref_models.DATATYPE_DEFINITION_BOOLEAN`
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.FieldTypeIdentifierMeta = ref_models.DATATYPE_DEFINITION_BOOLEAN{}
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.FieldOffsetX = 0.000000
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.FieldOffsetY = 0.000000
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.TargetMultiplicity = models.MANY
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.SourceMultiplicity = models.MANY
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.X = 1046.800000
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.Y = 278.200000
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.StartRatio = 0.500000
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.EndRatio = 0.500000
	__LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN.CornerOffsetRatio = 1.380000

	__LinkShape__000006_DATATYPE_DEFINITION_DATE.Name = `DATATYPE_DEFINITION_DATE`

	//gong:ident [ref_models.A_DATATYPES.DATATYPE_DEFINITION_DATE] comment added to overcome the problem with the comment map association
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.Identifier = `ref_models.A_DATATYPES.DATATYPE_DEFINITION_DATE`
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.IdentifierMeta = ref_models.A_DATATYPES{}.DATATYPE_DEFINITION_DATE

	//gong:ident [ref_models.DATATYPE_DEFINITION_DATE] comment added to overcome the problem with the comment map association
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.Fieldtypename = `ref_models.DATATYPE_DEFINITION_DATE`
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.FieldTypeIdentifierMeta = ref_models.DATATYPE_DEFINITION_DATE{}
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.FieldOffsetX = 0.000000
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.FieldOffsetY = 0.000000
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.TargetMultiplicity = models.MANY
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.SourceMultiplicity = models.MANY
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.X = 1046.299969
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.Y = 186.700000
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.StartRatio = 0.500000
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.EndRatio = 0.500000
	__LinkShape__000006_DATATYPE_DEFINITION_DATE.CornerOffsetRatio = 1.380000

	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.Name = `DATATYPE_DEFINITION_ENUMERATION`

	//gong:ident [ref_models.A_DATATYPES.DATATYPE_DEFINITION_ENUMERATION] comment added to overcome the problem with the comment map association
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.Identifier = `ref_models.A_DATATYPES.DATATYPE_DEFINITION_ENUMERATION`
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.IdentifierMeta = ref_models.A_DATATYPES{}.DATATYPE_DEFINITION_ENUMERATION

	//gong:ident [ref_models.DATATYPE_DEFINITION_ENUMERATION] comment added to overcome the problem with the comment map association
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.Fieldtypename = `ref_models.DATATYPE_DEFINITION_ENUMERATION`
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.FieldTypeIdentifierMeta = ref_models.DATATYPE_DEFINITION_ENUMERATION{}
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.FieldOffsetX = 0.000000
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.FieldOffsetY = 0.000000
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.TargetMultiplicity = models.MANY
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.SourceMultiplicity = models.MANY
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.X = 1047.200000
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.Y = 231.600000
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.StartRatio = 0.500000
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.EndRatio = 0.500000
	__LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION.CornerOffsetRatio = 1.380000

	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.Name = `DATATYPE_DEFINITION_INTEGER`

	//gong:ident [ref_models.A_DATATYPES.DATATYPE_DEFINITION_INTEGER] comment added to overcome the problem with the comment map association
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.Identifier = `ref_models.A_DATATYPES.DATATYPE_DEFINITION_INTEGER`
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.IdentifierMeta = ref_models.A_DATATYPES{}.DATATYPE_DEFINITION_INTEGER

	//gong:ident [ref_models.DATATYPE_DEFINITION_INTEGER] comment added to overcome the problem with the comment map association
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.Fieldtypename = `ref_models.DATATYPE_DEFINITION_INTEGER`
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.FieldTypeIdentifierMeta = ref_models.DATATYPE_DEFINITION_INTEGER{}
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.FieldOffsetX = 0.000000
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.FieldOffsetY = 0.000000
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.TargetMultiplicity = models.MANY
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.SourceMultiplicity = models.MANY
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.X = 1047.699969
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.Y = 326.500000
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.StartRatio = 0.500000
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.EndRatio = 0.500000
	__LinkShape__000008_DATATYPE_DEFINITION_INTEGER.CornerOffsetRatio = 1.380000

	__LinkShape__000009_DATATYPE_DEFINITION_REAL.Name = `DATATYPE_DEFINITION_REAL`

	//gong:ident [ref_models.A_DATATYPES.DATATYPE_DEFINITION_REAL] comment added to overcome the problem with the comment map association
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.Identifier = `ref_models.A_DATATYPES.DATATYPE_DEFINITION_REAL`
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.IdentifierMeta = ref_models.A_DATATYPES{}.DATATYPE_DEFINITION_REAL

	//gong:ident [ref_models.DATATYPE_DEFINITION_REAL] comment added to overcome the problem with the comment map association
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.Fieldtypename = `ref_models.DATATYPE_DEFINITION_REAL`
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.FieldTypeIdentifierMeta = ref_models.DATATYPE_DEFINITION_REAL{}
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.FieldOffsetX = 0.000000
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.FieldOffsetY = 0.000000
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.TargetMultiplicity = models.MANY
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.SourceMultiplicity = models.MANY
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.X = 1045.100000
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.Y = 143.200000
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.StartRatio = 0.500000
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.EndRatio = 0.500000
	__LinkShape__000009_DATATYPE_DEFINITION_REAL.CornerOffsetRatio = 1.380000

	__LinkShape__000010_DATATYPE_DEFINITION_STRING.Name = `DATATYPE_DEFINITION_STRING`

	//gong:ident [ref_models.A_DATATYPES.DATATYPE_DEFINITION_STRING] comment added to overcome the problem with the comment map association
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.Identifier = `ref_models.A_DATATYPES.DATATYPE_DEFINITION_STRING`
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.IdentifierMeta = ref_models.A_DATATYPES{}.DATATYPE_DEFINITION_STRING

	//gong:ident [ref_models.DATATYPE_DEFINITION_STRING] comment added to overcome the problem with the comment map association
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.Fieldtypename = `ref_models.DATATYPE_DEFINITION_STRING`
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.FieldTypeIdentifierMeta = ref_models.DATATYPE_DEFINITION_STRING{}
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.FieldOffsetX = 0.000000
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.FieldOffsetY = 0.000000
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.TargetMultiplicity = models.MANY
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.SourceMultiplicity = models.MANY
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.X = 1043.600000
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.Y = 98.100000
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.StartRatio = 0.500000
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.EndRatio = 0.500000
	__LinkShape__000010_DATATYPE_DEFINITION_STRING.CornerOffsetRatio = 1.380000

	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.Name = `DATATYPE_DEFINITION_XHTML`

	//gong:ident [ref_models.A_DATATYPES.DATATYPE_DEFINITION_XHTML] comment added to overcome the problem with the comment map association
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.Identifier = `ref_models.A_DATATYPES.DATATYPE_DEFINITION_XHTML`
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.IdentifierMeta = ref_models.A_DATATYPES{}.DATATYPE_DEFINITION_XHTML

	//gong:ident [ref_models.DATATYPE_DEFINITION_XHTML] comment added to overcome the problem with the comment map association
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.Fieldtypename = `ref_models.DATATYPE_DEFINITION_XHTML`
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.FieldTypeIdentifierMeta = ref_models.DATATYPE_DEFINITION_XHTML{}
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.FieldOffsetX = 0.000000
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.FieldOffsetY = 0.000000
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.TargetMultiplicity = models.MANY
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.SourceMultiplicity = models.MANY
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.X = 1043.600000
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.Y = 54.600000
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.StartRatio = 0.500000
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.EndRatio = 0.500000
	__LinkShape__000011_DATATYPE_DEFINITION_XHTML.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Top_Level.GongStructShapes = append(__Classdiagram__000000_Top_Level.GongStructShapes, __GongStructShape__000000_Default_REQ_IF)
	__Classdiagram__000000_Top_Level.GongStructShapes = append(__Classdiagram__000000_Top_Level.GongStructShapes, __GongStructShape__000001_Default_REQ_IF_CONTENT)
	__Classdiagram__000000_Top_Level.GongStructShapes = append(__Classdiagram__000000_Top_Level.GongStructShapes, __GongStructShape__000002_Default_REQ_IF_HEADER)
	__Classdiagram__000000_Top_Level.GongStructShapes = append(__Classdiagram__000000_Top_Level.GongStructShapes, __GongStructShape__000003_Default_A_THE_HEADER)
	__Classdiagram__000000_Top_Level.GongStructShapes = append(__Classdiagram__000000_Top_Level.GongStructShapes, __GongStructShape__000004_Default_A_CORE_CONTENT)
	__Classdiagram__000001_Datatypes.GongStructShapes = append(__Classdiagram__000001_Datatypes.GongStructShapes, __GongStructShape__000005_Content_REQ_IF_CONTENT)
	__Classdiagram__000001_Datatypes.GongStructShapes = append(__Classdiagram__000001_Datatypes.GongStructShapes, __GongStructShape__000006_Content_A_DATATYPES)
	__Classdiagram__000001_Datatypes.GongStructShapes = append(__Classdiagram__000001_Datatypes.GongStructShapes, __GongStructShape__000007_Content_DATATYPE_DEFINITION_BOOLEAN)
	__Classdiagram__000001_Datatypes.GongStructShapes = append(__Classdiagram__000001_Datatypes.GongStructShapes, __GongStructShape__000008_Content_DATATYPE_DEFINITION_DATE)
	__Classdiagram__000001_Datatypes.GongStructShapes = append(__Classdiagram__000001_Datatypes.GongStructShapes, __GongStructShape__000009_Content_DATATYPE_DEFINITION_ENUMERATION)
	__Classdiagram__000001_Datatypes.GongStructShapes = append(__Classdiagram__000001_Datatypes.GongStructShapes, __GongStructShape__000010_Content_DATATYPE_DEFINITION_INTEGER)
	__Classdiagram__000001_Datatypes.GongStructShapes = append(__Classdiagram__000001_Datatypes.GongStructShapes, __GongStructShape__000011_Content_DATATYPE_DEFINITION_REAL)
	__Classdiagram__000001_Datatypes.GongStructShapes = append(__Classdiagram__000001_Datatypes.GongStructShapes, __GongStructShape__000012_Content_DATATYPE_DEFINITION_STRING)
	__Classdiagram__000001_Datatypes.GongStructShapes = append(__Classdiagram__000001_Datatypes.GongStructShapes, __GongStructShape__000013_Content_DATATYPE_DEFINITION_XHTML)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.Classdiagrams, __Classdiagram__000000_Top_Level)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.Classdiagrams, __Classdiagram__000001_Datatypes)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.SelectedClassdiagram = __Classdiagram__000000_Top_Level
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_REQ_IF.LinkShapes = append(__GongStructShape__000000_Default_REQ_IF.LinkShapes, __LinkShape__000000_CORE_CONTENT)
	__GongStructShape__000000_Default_REQ_IF.LinkShapes = append(__GongStructShape__000000_Default_REQ_IF.LinkShapes, __LinkShape__000001_THE_HEADER)
	__GongStructShape__000003_Default_A_THE_HEADER.LinkShapes = append(__GongStructShape__000003_Default_A_THE_HEADER.LinkShapes, __LinkShape__000002_REQ_IF_HEADER)
	__GongStructShape__000004_Default_A_CORE_CONTENT.LinkShapes = append(__GongStructShape__000004_Default_A_CORE_CONTENT.LinkShapes, __LinkShape__000003_REQ_IF_CONTENT)
	__GongStructShape__000005_Content_REQ_IF_CONTENT.LinkShapes = append(__GongStructShape__000005_Content_REQ_IF_CONTENT.LinkShapes, __LinkShape__000004_DATATYPES)
	__GongStructShape__000006_Content_A_DATATYPES.LinkShapes = append(__GongStructShape__000006_Content_A_DATATYPES.LinkShapes, __LinkShape__000005_DATATYPE_DEFINITION_BOOLEAN)
	__GongStructShape__000006_Content_A_DATATYPES.LinkShapes = append(__GongStructShape__000006_Content_A_DATATYPES.LinkShapes, __LinkShape__000006_DATATYPE_DEFINITION_DATE)
	__GongStructShape__000006_Content_A_DATATYPES.LinkShapes = append(__GongStructShape__000006_Content_A_DATATYPES.LinkShapes, __LinkShape__000007_DATATYPE_DEFINITION_ENUMERATION)
	__GongStructShape__000006_Content_A_DATATYPES.LinkShapes = append(__GongStructShape__000006_Content_A_DATATYPES.LinkShapes, __LinkShape__000008_DATATYPE_DEFINITION_INTEGER)
	__GongStructShape__000006_Content_A_DATATYPES.LinkShapes = append(__GongStructShape__000006_Content_A_DATATYPES.LinkShapes, __LinkShape__000009_DATATYPE_DEFINITION_REAL)
	__GongStructShape__000006_Content_A_DATATYPES.LinkShapes = append(__GongStructShape__000006_Content_A_DATATYPES.LinkShapes, __LinkShape__000010_DATATYPE_DEFINITION_STRING)
	__GongStructShape__000006_Content_A_DATATYPES.LinkShapes = append(__GongStructShape__000006_Content_A_DATATYPES.LinkShapes, __LinkShape__000011_DATATYPE_DEFINITION_XHTML)
	// setup of LinkShape instances pointers
}
