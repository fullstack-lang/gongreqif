package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_REQ_IF := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_REQ_IF_CONTENT := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_REQ_IF_HEADER := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`

	//gong:ident [ref_models.REQ_IF.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000000_Name.Identifier = `ref_models.REQ_IF.Name`
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `REQ_IF`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.Name = `Diagram Package created the 2025-05-07T07:33:18Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_REQ_IF.Name = `Default-REQ_IF`
	__GongStructShape__000000_Default_REQ_IF.X = 42.000000
	__GongStructShape__000000_Default_REQ_IF.Y = 25.000000

	//gong:ident [ref_models.REQ_IF] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_REQ_IF.Identifier = `ref_models.REQ_IF`
	__GongStructShape__000000_Default_REQ_IF.ShowNbInstances = false
	__GongStructShape__000000_Default_REQ_IF.NbInstances = 0
	__GongStructShape__000000_Default_REQ_IF.Width = 240.000000
	__GongStructShape__000000_Default_REQ_IF.Height = 83.000000
	__GongStructShape__000000_Default_REQ_IF.IsSelected = false

	__GongStructShape__000001_Default_REQ_IF_CONTENT.Name = `Default-REQ_IF_CONTENT`
	__GongStructShape__000001_Default_REQ_IF_CONTENT.X = 747.000000
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Y = 190.000000

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Identifier = `ref_models.REQ_IF_CONTENT`
	__GongStructShape__000001_Default_REQ_IF_CONTENT.ShowNbInstances = false
	__GongStructShape__000001_Default_REQ_IF_CONTENT.NbInstances = 0
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Width = 240.000000
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Height = 63.000000
	__GongStructShape__000001_Default_REQ_IF_CONTENT.IsSelected = false

	__GongStructShape__000002_Default_REQ_IF_HEADER.Name = `Default-REQ_IF_HEADER`
	__GongStructShape__000002_Default_REQ_IF_HEADER.X = 741.000000
	__GongStructShape__000002_Default_REQ_IF_HEADER.Y = 48.000000

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_REQ_IF_HEADER.Identifier = `ref_models.REQ_IF_HEADER`
	__GongStructShape__000002_Default_REQ_IF_HEADER.ShowNbInstances = false
	__GongStructShape__000002_Default_REQ_IF_HEADER.NbInstances = 0
	__GongStructShape__000002_Default_REQ_IF_HEADER.Width = 240.000000
	__GongStructShape__000002_Default_REQ_IF_HEADER.Height = 63.000000
	__GongStructShape__000002_Default_REQ_IF_HEADER.IsSelected = false

	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.Name = `CORE_CONTENT.REQ_IF_CONTENT`

	//gong:ident [ref_models.REQ_IF.CORE_CONTENT.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.Identifier = `ref_models.REQ_IF.CORE_CONTENT.REQ_IF_CONTENT`

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.Fieldtypename = `ref_models.REQ_IF_CONTENT`
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.FieldOffsetX = 0.000000
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.FieldOffsetY = 0.000000
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.SourceMultiplicity = models.MANY
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.X = 412.500000
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.Y = 62.500000
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.StartRatio = 0.500000
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.EndRatio = 0.500000
	__LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT.CornerOffsetRatio = 1.380000

	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.Name = `THE_HEADER.REQ_IF_HEADER`

	//gong:ident [ref_models.REQ_IF.THE_HEADER.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.Identifier = `ref_models.REQ_IF.THE_HEADER.REQ_IF_HEADER`

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.Fieldtypename = `ref_models.REQ_IF_HEADER`
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.FieldOffsetX = 0.000000
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.FieldOffsetY = 0.000000
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.SourceMultiplicity = models.MANY
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.X = 751.500000
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.Y = 78.000000
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.StartRatio = 0.500000
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.EndRatio = 0.500000
	__LinkShape__000001_THE_HEADER_REQ_IF_HEADER.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_REQ_IF)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_REQ_IF_CONTENT)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_REQ_IF_HEADER)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_33_18Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_REQ_IF.AttributeShapes = append(__GongStructShape__000000_Default_REQ_IF.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_REQ_IF.LinkShapes = append(__GongStructShape__000000_Default_REQ_IF.LinkShapes, __LinkShape__000000_CORE_CONTENT_REQ_IF_CONTENT)
	__GongStructShape__000000_Default_REQ_IF.LinkShapes = append(__GongStructShape__000000_Default_REQ_IF.LinkShapes, __LinkShape__000001_THE_HEADER_REQ_IF_HEADER)
	// setup of LinkShape instances pointers
}
