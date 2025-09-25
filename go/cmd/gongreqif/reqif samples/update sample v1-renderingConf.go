package main

import (
	"time"

	"github.com/fullstack-lang/gongreqif/go/models"
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

	const __write__local_time = "0001-01-01 00:00:00.000000 UTC"
	const __write__utc_time__ = "0001-01-01 00:00:00.000000 UTC"

	const __commitId__ = "0000000000"

	// Declaration of instances to stage

	__Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry__000000__attr_reqif_text := (&models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry{}).Stage(stage)

	__Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry__000000__attr_reqif_text := (&models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry{}).Stage(stage)

	__Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry__000000__attr_reqif_text := (&models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry{}).Stage(stage)

	__Map_SPECIFICATION_Nodes_expandedEntry__000000__spec_001 := (&models.Map_SPECIFICATION_Nodes_expandedEntry{}).Stage(stage)

	__Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry__000000__spectype_requirement := (&models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry{}).Stage(stage)

	__Map_SPEC_OBJECT_TYPE_showIdentifierEntry__000000__spectype_requirement := (&models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry{}).Stage(stage)

	__Map_SPEC_OBJECT_TYPE_showNameEntry__000000__spectype_requirement := (&models.Map_SPEC_OBJECT_TYPE_showNameEntry{}).Stage(stage)

	__RenderingConfiguration__000000_update_sample_v1_reqif := (&models.RenderingConfiguration{}).Stage(stage)

	// Setup of values

	__Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry__000000__attr_reqif_text.Name = `_attr_reqif_text`
	__Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry__000000__attr_reqif_text.Value = true

	__Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry__000000__attr_reqif_text.Name = `_attr_reqif_text`
	__Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry__000000__attr_reqif_text.Value = false

	__Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry__000000__attr_reqif_text.Name = `_attr_reqif_text`
	__Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry__000000__attr_reqif_text.Value = false

	__Map_SPECIFICATION_Nodes_expandedEntry__000000__spec_001.Name = `_spec_001`
	__Map_SPECIFICATION_Nodes_expandedEntry__000000__spec_001.Value = true

	__Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry__000000__spectype_requirement.Name = `_spectype_requirement`
	__Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry__000000__spectype_requirement.Value = true

	__Map_SPEC_OBJECT_TYPE_showIdentifierEntry__000000__spectype_requirement.Name = `_spectype_requirement`
	__Map_SPEC_OBJECT_TYPE_showIdentifierEntry__000000__spectype_requirement.Value = true

	__Map_SPEC_OBJECT_TYPE_showNameEntry__000000__spectype_requirement.Name = `_spectype_requirement`
	__Map_SPEC_OBJECT_TYPE_showNameEntry__000000__spectype_requirement.Value = false

	__RenderingConfiguration__000000_update_sample_v1_reqif.Name = `update sample v1.reqif`

	// Setup of pointers
	// setup of Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry instances pointers
	// setup of Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry instances pointers
	// setup of Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry instances pointers
	// setup of Map_SPECIFICATION_Nodes_expandedEntry instances pointers
	// setup of Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry instances pointers
	// setup of Map_SPEC_OBJECT_TYPE_showIdentifierEntry instances pointers
	// setup of Map_SPEC_OBJECT_TYPE_showNameEntry instances pointers
	// setup of RenderingConfiguration instances pointers
	__RenderingConfiguration__000000_update_sample_v1_reqif.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries = append(__RenderingConfiguration__000000_update_sample_v1_reqif.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries, __Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry__000000__spectype_requirement)
	__RenderingConfiguration__000000_update_sample_v1_reqif.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries = append(__RenderingConfiguration__000000_update_sample_v1_reqif.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries, __Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry__000000__attr_reqif_text)
	__RenderingConfiguration__000000_update_sample_v1_reqif.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries = append(__RenderingConfiguration__000000_update_sample_v1_reqif.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries, __Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry__000000__attr_reqif_text)
	__RenderingConfiguration__000000_update_sample_v1_reqif.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries = append(__RenderingConfiguration__000000_update_sample_v1_reqif.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries, __Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry__000000__attr_reqif_text)
	__RenderingConfiguration__000000_update_sample_v1_reqif.Map_SPECIFICATION_Nodes_expandedEntries = append(__RenderingConfiguration__000000_update_sample_v1_reqif.Map_SPECIFICATION_Nodes_expandedEntries, __Map_SPECIFICATION_Nodes_expandedEntry__000000__spec_001)
	__RenderingConfiguration__000000_update_sample_v1_reqif.Map_SPEC_OBJECT_TYPE_showIdentifierEntries = append(__RenderingConfiguration__000000_update_sample_v1_reqif.Map_SPEC_OBJECT_TYPE_showIdentifierEntries, __Map_SPEC_OBJECT_TYPE_showIdentifierEntry__000000__spectype_requirement)
	__RenderingConfiguration__000000_update_sample_v1_reqif.Map_SPEC_OBJECT_TYPE_showNameEntries = append(__RenderingConfiguration__000000_update_sample_v1_reqif.Map_SPEC_OBJECT_TYPE_showNameEntries, __Map_SPEC_OBJECT_TYPE_showNameEntry__000000__spectype_requirement)
}

