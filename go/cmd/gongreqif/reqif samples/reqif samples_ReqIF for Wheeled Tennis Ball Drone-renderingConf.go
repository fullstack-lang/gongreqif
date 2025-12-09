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

	// Declaration of instances to stage

	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000000_ats_system_req_verified := (&models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000001_ats_spec_approved := (&models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000002_ats_trace_approved := (&models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering{}).Stage(stage)

	__ATTRIBUTE_DEFINITION_DATE_Rendering__000000_ats_system_req_duedate := (&models.ATTRIBUTE_DEFINITION_DATE_Rendering{}).Stage(stage)

	__ATTRIBUTE_DEFINITION_ENUMERATION_Rendering__000000_ats_relation_status := (&models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_ENUMERATION_Rendering__000001_ats_system_req_status := (&models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering{}).Stage(stage)

	__ATTRIBUTE_DEFINITION_INTEGER_Rendering__000000_ats_system_req_priority := (&models.ATTRIBUTE_DEFINITION_INTEGER_Rendering{}).Stage(stage)

	__ATTRIBUTE_DEFINITION_REAL_Rendering__000000_ats_system_req_cost := (&models.ATTRIBUTE_DEFINITION_REAL_Rendering{}).Stage(stage)

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000000_ats_relation_validation := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000001_ats_trace_version := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000002_ats_heading_title := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000003_ats_stakeholder_req_text := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000004_ats_system_req_text := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000005_ats_spec_version := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000006_ats_trace_owner := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000007_ats_stakeholder_req_title := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000008_ats_spec_owner := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{}).Stage(stage)
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000009_ats_system_req_title := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{}).Stage(stage)

	__ATTRIBUTE_DEFINITION_XHTML_Rendering__000000_ats_text_content := (&models.ATTRIBUTE_DEFINITION_XHTML_Rendering{}).Stage(stage)

	__SPECIFICATION_Rendering__000000_spec_stakeholder := (&models.SPECIFICATION_Rendering{}).Stage(stage)
	__SPECIFICATION_Rendering__000001_spec_system := (&models.SPECIFICATION_Rendering{}).Stage(stage)
	__SPECIFICATION_Rendering__000002_spec_traceability_matrix := (&models.SPECIFICATION_Rendering{}).Stage(stage)

	__SPEC_OBJECT_TYPE_Rendering__000000_st_descriptive_text := (&models.SPEC_OBJECT_TYPE_Rendering{}).Stage(stage)
	__SPEC_OBJECT_TYPE_Rendering__000001_st_heading := (&models.SPEC_OBJECT_TYPE_Rendering{}).Stage(stage)
	__SPEC_OBJECT_TYPE_Rendering__000002_st_stakeholder_req := (&models.SPEC_OBJECT_TYPE_Rendering{}).Stage(stage)
	__SPEC_OBJECT_TYPE_Rendering__000003_st_system_req := (&models.SPEC_OBJECT_TYPE_Rendering{}).Stage(stage)

	// Setup of values

	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000000_ats_system_req_verified.Name = `ats-system-req-verified`
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000000_ats_system_req_verified.ShowInTable = false
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000000_ats_system_req_verified.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000000_ats_system_req_verified.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000001_ats_spec_approved.Name = `ats-spec-approved`
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000001_ats_spec_approved.ShowInTable = false
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000001_ats_spec_approved.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000001_ats_spec_approved.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000002_ats_trace_approved.Name = `ats-trace-approved`
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000002_ats_trace_approved.ShowInTable = false
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000002_ats_trace_approved.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_BOOLEAN_Rendering__000002_ats_trace_approved.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_DATE_Rendering__000000_ats_system_req_duedate.Name = `ats-system-req-duedate`
	__ATTRIBUTE_DEFINITION_DATE_Rendering__000000_ats_system_req_duedate.ShowInTable = false
	__ATTRIBUTE_DEFINITION_DATE_Rendering__000000_ats_system_req_duedate.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_DATE_Rendering__000000_ats_system_req_duedate.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_ENUMERATION_Rendering__000000_ats_relation_status.Name = `ats-relation-status`
	__ATTRIBUTE_DEFINITION_ENUMERATION_Rendering__000000_ats_relation_status.ShowInTable = false
	__ATTRIBUTE_DEFINITION_ENUMERATION_Rendering__000000_ats_relation_status.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_ENUMERATION_Rendering__000000_ats_relation_status.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_ENUMERATION_Rendering__000001_ats_system_req_status.Name = `ats-system-req-status`
	__ATTRIBUTE_DEFINITION_ENUMERATION_Rendering__000001_ats_system_req_status.ShowInTable = false
	__ATTRIBUTE_DEFINITION_ENUMERATION_Rendering__000001_ats_system_req_status.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_ENUMERATION_Rendering__000001_ats_system_req_status.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_INTEGER_Rendering__000000_ats_system_req_priority.Name = `ats-system-req-priority`
	__ATTRIBUTE_DEFINITION_INTEGER_Rendering__000000_ats_system_req_priority.ShowInTable = false
	__ATTRIBUTE_DEFINITION_INTEGER_Rendering__000000_ats_system_req_priority.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_INTEGER_Rendering__000000_ats_system_req_priority.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_REAL_Rendering__000000_ats_system_req_cost.Name = `ats-system-req-cost`
	__ATTRIBUTE_DEFINITION_REAL_Rendering__000000_ats_system_req_cost.ShowInTable = false
	__ATTRIBUTE_DEFINITION_REAL_Rendering__000000_ats_system_req_cost.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_REAL_Rendering__000000_ats_system_req_cost.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000000_ats_relation_validation.Name = `ats-relation-validation`
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000000_ats_relation_validation.ShowInTable = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000000_ats_relation_validation.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000000_ats_relation_validation.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000001_ats_trace_version.Name = `ats-trace-version`
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000001_ats_trace_version.ShowInTable = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000001_ats_trace_version.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000001_ats_trace_version.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000002_ats_heading_title.Name = `ats-heading-title`
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000002_ats_heading_title.ShowInTable = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000002_ats_heading_title.ShowInTitle = true
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000002_ats_heading_title.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000003_ats_stakeholder_req_text.Name = `ats-stakeholder-req-text`
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000003_ats_stakeholder_req_text.ShowInTable = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000003_ats_stakeholder_req_text.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000003_ats_stakeholder_req_text.ShowInSubject = true

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000004_ats_system_req_text.Name = `ats-system-req-text`
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000004_ats_system_req_text.ShowInTable = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000004_ats_system_req_text.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000004_ats_system_req_text.ShowInSubject = true

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000005_ats_spec_version.Name = `ats-spec-version`
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000005_ats_spec_version.ShowInTable = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000005_ats_spec_version.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000005_ats_spec_version.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000006_ats_trace_owner.Name = `ats-trace-owner`
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000006_ats_trace_owner.ShowInTable = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000006_ats_trace_owner.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000006_ats_trace_owner.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000007_ats_stakeholder_req_title.Name = `ats-stakeholder-req-title`
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000007_ats_stakeholder_req_title.ShowInTable = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000007_ats_stakeholder_req_title.ShowInTitle = true
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000007_ats_stakeholder_req_title.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000008_ats_spec_owner.Name = `ats-spec-owner`
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000008_ats_spec_owner.ShowInTable = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000008_ats_spec_owner.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000008_ats_spec_owner.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_STRING_Rendering__000009_ats_system_req_title.Name = `ats-system-req-title`
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000009_ats_system_req_title.ShowInTable = false
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000009_ats_system_req_title.ShowInTitle = true
	__ATTRIBUTE_DEFINITION_STRING_Rendering__000009_ats_system_req_title.ShowInSubject = false

	__ATTRIBUTE_DEFINITION_XHTML_Rendering__000000_ats_text_content.Name = `ats-text-content`
	__ATTRIBUTE_DEFINITION_XHTML_Rendering__000000_ats_text_content.ShowInTable = false
	__ATTRIBUTE_DEFINITION_XHTML_Rendering__000000_ats_text_content.ShowInTitle = false
	__ATTRIBUTE_DEFINITION_XHTML_Rendering__000000_ats_text_content.ShowInSubject = true

	__SPECIFICATION_Rendering__000000_spec_stakeholder.Name = `spec-stakeholder`
	__SPECIFICATION_Rendering__000000_spec_stakeholder.IsNodeExpanded = false
	__SPECIFICATION_Rendering__000000_spec_stakeholder.IsSelected = true

	__SPECIFICATION_Rendering__000001_spec_system.Name = `spec-system`
	__SPECIFICATION_Rendering__000001_spec_system.IsNodeExpanded = false
	__SPECIFICATION_Rendering__000001_spec_system.IsSelected = false

	__SPECIFICATION_Rendering__000002_spec_traceability_matrix.Name = `spec-traceability-matrix`
	__SPECIFICATION_Rendering__000002_spec_traceability_matrix.IsNodeExpanded = false
	__SPECIFICATION_Rendering__000002_spec_traceability_matrix.IsSelected = false

	__SPEC_OBJECT_TYPE_Rendering__000000_st_descriptive_text.Name = `st-descriptive-text`
	__SPEC_OBJECT_TYPE_Rendering__000000_st_descriptive_text.IsNodeExpanded = true
	__SPEC_OBJECT_TYPE_Rendering__000000_st_descriptive_text.ShowIdentifier = false
	__SPEC_OBJECT_TYPE_Rendering__000000_st_descriptive_text.ShowName = false
	__SPEC_OBJECT_TYPE_Rendering__000000_st_descriptive_text.ShowRelations = false
	__SPEC_OBJECT_TYPE_Rendering__000000_st_descriptive_text.IsHeading = false

	__SPEC_OBJECT_TYPE_Rendering__000001_st_heading.Name = `st-heading`
	__SPEC_OBJECT_TYPE_Rendering__000001_st_heading.IsNodeExpanded = true
	__SPEC_OBJECT_TYPE_Rendering__000001_st_heading.ShowIdentifier = false
	__SPEC_OBJECT_TYPE_Rendering__000001_st_heading.ShowName = false
	__SPEC_OBJECT_TYPE_Rendering__000001_st_heading.ShowRelations = false
	__SPEC_OBJECT_TYPE_Rendering__000001_st_heading.IsHeading = true

	__SPEC_OBJECT_TYPE_Rendering__000002_st_stakeholder_req.Name = `st-stakeholder-req`
	__SPEC_OBJECT_TYPE_Rendering__000002_st_stakeholder_req.IsNodeExpanded = true
	__SPEC_OBJECT_TYPE_Rendering__000002_st_stakeholder_req.ShowIdentifier = true
	__SPEC_OBJECT_TYPE_Rendering__000002_st_stakeholder_req.ShowName = false
	__SPEC_OBJECT_TYPE_Rendering__000002_st_stakeholder_req.ShowRelations = true
	__SPEC_OBJECT_TYPE_Rendering__000002_st_stakeholder_req.IsHeading = false

	__SPEC_OBJECT_TYPE_Rendering__000003_st_system_req.Name = `st-system-req`
	__SPEC_OBJECT_TYPE_Rendering__000003_st_system_req.IsNodeExpanded = true
	__SPEC_OBJECT_TYPE_Rendering__000003_st_system_req.ShowIdentifier = true
	__SPEC_OBJECT_TYPE_Rendering__000003_st_system_req.ShowName = false
	__SPEC_OBJECT_TYPE_Rendering__000003_st_system_req.ShowRelations = true
	__SPEC_OBJECT_TYPE_Rendering__000003_st_system_req.IsHeading = false

	// Setup of pointers
	// setup of ATTRIBUTE_DEFINITION_BOOLEAN_Rendering instances pointers
	// setup of ATTRIBUTE_DEFINITION_DATE_Rendering instances pointers
	// setup of ATTRIBUTE_DEFINITION_ENUMERATION_Rendering instances pointers
	// setup of ATTRIBUTE_DEFINITION_INTEGER_Rendering instances pointers
	// setup of ATTRIBUTE_DEFINITION_REAL_Rendering instances pointers
	// setup of ATTRIBUTE_DEFINITION_STRING_Rendering instances pointers
	// setup of ATTRIBUTE_DEFINITION_XHTML_Rendering instances pointers
	// setup of SPECIFICATION_Rendering instances pointers
	// setup of SPEC_OBJECT_TYPE_Rendering instances pointers
}
