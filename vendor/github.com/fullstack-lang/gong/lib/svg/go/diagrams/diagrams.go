package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/svg/go/models"
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

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_Content := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_LinkAnchoredText := (&models.GongStructShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`

	//gong:ident [ref_models.LinkAnchoredText.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000000_Name.Identifier = `ref_models.LinkAnchoredText.Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.LinkAnchoredText{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `LinkAnchoredText`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_Content.Name = `Content`

	//gong:ident [ref_models.LinkAnchoredText.Content] comment added to overcome the problem with the comment map association
	__AttributeShape__000001_Content.Identifier = `ref_models.LinkAnchoredText.Content`
	__AttributeShape__000001_Content.IdentifierMeta = ref_models.LinkAnchoredText{}.Content
	__AttributeShape__000001_Content.FieldTypeAsString = ``
	__AttributeShape__000001_Content.Structname = `LinkAnchoredText`
	__AttributeShape__000001_Content.Fieldtypename = `string`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[false,false,false,false,false,false,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Name = `Diagram Package created the 2025-05-05T06:12:28Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_LinkAnchoredText.Name = `Default-LinkAnchoredText`
	__GongStructShape__000000_Default_LinkAnchoredText.X = 25.000000
	__GongStructShape__000000_Default_LinkAnchoredText.Y = 53.000000

	//gong:ident [.............] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_LinkAnchoredText.Identifier = `.............`
	__GongStructShape__000000_Default_LinkAnchoredText.IdentifierMeta = ref_models.LinkAnchoredText{}
	__GongStructShape__000000_Default_LinkAnchoredText.ShowNbInstances = false
	__GongStructShape__000000_Default_LinkAnchoredText.NbInstances = 0
	__GongStructShape__000000_Default_LinkAnchoredText.Width = 240.000000
	__GongStructShape__000000_Default_LinkAnchoredText.Height = 103.000000
	__GongStructShape__000000_Default_LinkAnchoredText.IsSelected = false

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_LinkAnchoredText)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_LinkAnchoredText.AttributeShapes = append(__GongStructShape__000000_Default_LinkAnchoredText.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_LinkAnchoredText.AttributeShapes = append(__GongStructShape__000000_Default_LinkAnchoredText.AttributeShapes, __AttributeShape__000001_Content)
}
