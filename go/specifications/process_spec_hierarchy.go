package specifications

import (
	"fmt"
	"log"
	"sort"
	"strings"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	m "github.com/fullstack-lang/gongreqif/go/models"
	"github.com/fullstack-lang/gongreqif/go/specobjects"
	"github.com/fullstack-lang/gongreqif/go/spectypes"
)

func processSpecHierarchy(
	stager *m.Stager,
	specHierarchy *m.SPEC_HIERARCHY,
	hierarchyParentNode *tree.Node,
	outerDepth int,
	markDownContent *string,
	digitPrefix string,
) {
	stage := stager.GetStage()

	specObject, ok := stager.Map_id_SPEC_OBJECT[specHierarchy.OBJECT.SPEC_OBJECT_REF]
	if !ok {
		log.Panic("specHierarchy.OBJECT.SPEC_OBJECT_REF", specHierarchy.OBJECT.SPEC_OBJECT_REF,
			"unknown ref")
	}

	specObjectType, ok := stager.Map_id_SPEC_OBJECT_TYPE[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
	if !ok {
		log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
			"unknown ref")
	}

	markdownBoldStartingMark := `
**`
	markdownBoldEndingMark := `**
`

	specObjectTypeRendering := spectypes.GetSpecObjectTypeRendering(stage, specObjectType)

	if specObjectTypeRendering.IsHeading {

		*markDownContent += `
`

		if specHierarchy.CHILDREN == nil || len(specHierarchy.CHILDREN.SPEC_HIERARCHY) == 0 {
			*markDownContent += "#"
		}

		if true {
			for range outerDepth {
				*markDownContent += "#"
			}
			*markDownContent += " "

			if digitPrefix != "" {
				*markDownContent += digitPrefix + " "
			}
		}
	} else {
		*markDownContent += markdownBoldStartingMark
	}

	specobjects.AddIdentifierAndNameToTitle(stager, specObjectType, markDownContent, specObject)

	// Fetch sorted attributes based on Rank
	sortedAttributes := getSortedAttributes(stager, specObject)

	// Generate Title Complement
	var titleBuilder strings.Builder
	firstTitle := true
	for _, attr := range sortedAttributes {
		if attr.ShowInTitle {
			if !firstTitle {
				titleBuilder.WriteString(" - ")
			}
			titleBuilder.WriteString(attr.Value)
			firstTitle = false
		}
	}
	titleComplement := titleBuilder.String()

	// Generate Subject Complement
	var subjectBuilder strings.Builder
	for _, attr := range sortedAttributes {
		if attr.ShowInSubject {
			if subjectBuilder.Len() > 0 {
				subjectBuilder.WriteString("\n\n")
			}
			subjectBuilder.WriteString(attr.Value)
		}
	}
	subjectComplement := subjectBuilder.String()

	if !specObjectTypeRendering.IsHeading && !strings.HasSuffix(*markDownContent, "**") && titleComplement != "" {
		*markDownContent += " - "
	}

	*markDownContent += titleComplement

	if specObjectTypeRendering.IsHeading {
		*markDownContent += `
`
	} else {
		// ending mark for bold
		*markDownContent += markdownBoldEndingMark
	}

	// remove "****" if no title is present
	if strings.HasSuffix(*markDownContent, markdownBoldStartingMark+markdownBoldEndingMark) {
		*markDownContent = strings.TrimSuffix(*markDownContent, markdownBoldStartingMark+markdownBoldEndingMark)
	}

	// add the subject after the title
	if subjectComplement != "" {
		*markDownContent += `
` + subjectComplement + `
`
	}

	specObjectNode := &tree.Node{
		Name: specObject.Name + " : " + specObjectType.Name,
	}
	hierarchyParentNode.Children = append(hierarchyParentNode.Children, specObjectNode)

	specobjects.AddAttributeNodes(stager, specObjectNode, specObject)

	// Append Attributes to Markdown (Table)
	{
		var tableBuilder strings.Builder
		hasTableRows := false

		// Header
		tableHeader := "\n| Attribute | Value |\n| --- | --- |\n"

		for _, attr := range sortedAttributes {
			if attr.ShowInTable {
				if !hasTableRows {
					tableBuilder.WriteString(tableHeader)
					hasTableRows = true
				}
				// Escape pipes in value to avoid breaking table
				safeValue := strings.ReplaceAll(attr.Value, "|", "\\|")
				safeValue = strings.ReplaceAll(safeValue, "\n", "<br>")
				tableBuilder.WriteString(fmt.Sprintf("| %s | %s |\n", attr.Name, safeValue))
			}
		}

		if hasTableRows {
			*markDownContent += tableBuilder.String() + "\n"
		}
	}

	m.AddIconForEditabilityOfAttribute(specHierarchy.IS_EDITABLE, specObject.Name, specObjectNode)

	if specHierarchy.CHILDREN != nil {
		for i, specHierarchyChildren := range specHierarchy.CHILDREN.SPEC_HIERARCHY {
			childPrefix := ""
			if digitPrefix != "" {
				childPrefix = digitPrefix + "." + fmt.Sprintf("%d", i+1)
			}
			processSpecHierarchy(
				stager,
				specHierarchyChildren,
				specObjectNode,
				outerDepth+1,
				markDownContent,
				childPrefix)
		}
	}
}

// AttributeDefinitionRendering interface definition
type AttributeDefinitionRendering interface {
	m.Gongstruct
	GetName() string
	GetShowInTablePtr() *bool
	GetShowInTitlePtr() *bool
	GetShowInSubjectPtr() *bool
	GetRankPtr() *int
}

type SortedAttribute struct {
	Rank          int
	Name          string
	Value         string
	ShowInTable   bool
	ShowInTitle   bool
	ShowInSubject bool
}

func getSortedAttributes(stager *m.Stager, specObject *m.SPEC_OBJECT) []SortedAttribute {
	var attributes []SortedAttribute
	stage := stager.GetStage()

	// Helper to extract rank and flags safely
	extractRenderingInfo := func(rendering interface{}) (rank int, table, title, subject bool) {
		table = true // Default for table
		if rendering == nil {
			return
		}
		if r, ok := rendering.(AttributeDefinitionRendering); ok {
			if r.GetRankPtr() != nil {
				rank = *r.GetRankPtr()
			}
			if r.GetShowInTablePtr() != nil {
				table = *r.GetShowInTablePtr()
			}
			if r.GetShowInTitlePtr() != nil {
				title = *r.GetShowInTitlePtr()
			}
			if r.GetShowInSubjectPtr() != nil {
				subject = *r.GetShowInSubjectPtr()
			}
		}
		return
	}

	// 1. STRING
	for _, att := range specObject.VALUES.ATTRIBUTE_VALUE_STRING {
		defID := att.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF
		if def, ok := stager.Map_id_ATTRIBUTE_DEFINITION_STRING[defID]; ok {
			renderMap := m.GongGetMap[*m.ATTRIBUTE_DEFINITION_STRING_Rendering](stage)
			var rendering interface{}
			if r, ok := renderMap[defID]; ok {
				rendering = r
			}
			rank, table, title, subject := extractRenderingInfo(rendering)
			attributes = append(attributes, SortedAttribute{
				Rank: rank, Name: def.Name, Value: att.THE_VALUE,
				ShowInTable: table, ShowInTitle: title, ShowInSubject: subject,
			})
		}
	}

	// 2. INTEGER
	for _, att := range specObject.VALUES.ATTRIBUTE_VALUE_INTEGER {
		defID := att.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF
		if def, ok := stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER[defID]; ok {
			renderMap := m.GongGetMap[*m.ATTRIBUTE_DEFINITION_INTEGER_Rendering](stage)
			var rendering interface{}
			if r, ok := renderMap[defID]; ok {
				rendering = r
			}
			rank, table, title, subject := extractRenderingInfo(rendering)
			attributes = append(attributes, SortedAttribute{
				Rank: rank, Name: def.Name, Value: fmt.Sprintf("%d", att.THE_VALUE),
				ShowInTable: table, ShowInTitle: title, ShowInSubject: subject,
			})
		}
	}

	// 3. REAL
	for _, att := range specObject.VALUES.ATTRIBUTE_VALUE_REAL {
		defID := att.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF
		if def, ok := stager.Map_id_ATTRIBUTE_DEFINITION_REAL[defID]; ok {
			renderMap := m.GongGetMap[*m.ATTRIBUTE_DEFINITION_REAL_Rendering](stage)
			var rendering interface{}
			if r, ok := renderMap[defID]; ok {
				rendering = r
			}
			rank, table, title, subject := extractRenderingInfo(rendering)
			attributes = append(attributes, SortedAttribute{
				Rank: rank, Name: def.Name, Value: fmt.Sprintf("%f", att.THE_VALUE),
				ShowInTable: table, ShowInTitle: title, ShowInSubject: subject,
			})
		}
	}

	// 4. BOOLEAN
	for _, att := range specObject.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		defID := att.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF
		if def, ok := stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN[defID]; ok {
			renderMap := m.GongGetMap[*m.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stage)
			var rendering interface{}
			if r, ok := renderMap[defID]; ok {
				rendering = r
			}
			rank, table, title, subject := extractRenderingInfo(rendering)
			attributes = append(attributes, SortedAttribute{
				Rank: rank, Name: def.Name, Value: fmt.Sprintf("%t", att.THE_VALUE),
				ShowInTable: table, ShowInTitle: title, ShowInSubject: subject,
			})
		}
	}

	// 5. DATE
	for _, att := range specObject.VALUES.ATTRIBUTE_VALUE_DATE {
		defID := att.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF
		if def, ok := stager.Map_id_ATTRIBUTE_DEFINITION_DATE[defID]; ok {
			renderMap := m.GongGetMap[*m.ATTRIBUTE_DEFINITION_DATE_Rendering](stage)
			var rendering interface{}
			if r, ok := renderMap[defID]; ok {
				rendering = r
			}
			rank, table, title, subject := extractRenderingInfo(rendering)
			attributes = append(attributes, SortedAttribute{
				Rank: rank, Name: def.Name, Value: att.THE_VALUE,
				ShowInTable: table, ShowInTitle: title, ShowInSubject: subject,
			})
		}
	}

	// 6. XHTML
	for _, att := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
		defID := att.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF
		if def, ok := stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[defID]; ok {
			renderMap := m.GongGetMap[*m.ATTRIBUTE_DEFINITION_XHTML_Rendering](stage)
			var rendering interface{}
			if r, ok := renderMap[defID]; ok {
				rendering = r
			}
			rank, table, title, subject := extractRenderingInfo(rendering)

			// Clean XHTML
			enclosedText := att.THE_VALUE.EnclosedText

			// 1. Strip namespaces so proper HTML tags are used (e.g. <xhtml:div> -> <div>)
			enclosedText = strings.ReplaceAll(enclosedText, "<xhtml:", "<")
			enclosedText = strings.ReplaceAll(enclosedText, "</xhtml:", "</")
			enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:", "<")
			enclosedText = strings.ReplaceAll(enclosedText, "</reqif-xhtml:", "</")

			// 2. Remove indentation.
			// ReqIF XML is often pretty-printed. If this indentation (spaces) is left,
			// Markdown interprets it as a Code Block, preventing the HTML from rendering.
			lines := strings.Split(enclosedText, "\n")
			for i, line := range lines {
				lines[i] = strings.TrimSpace(line)
			}
			enclosedText = strings.Join(lines, "\n")

			attributes = append(attributes, SortedAttribute{
				Rank: rank, Name: def.Name, Value: enclosedText,
				ShowInTable: table, ShowInTitle: title, ShowInSubject: subject,
			})
		}
	}

	// 7. ENUMERATION
	for _, att := range specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		defID := att.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF
		if def, ok := stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION[defID]; ok {
			renderMap := m.GongGetMap[*m.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stage)
			var rendering interface{}
			if r, ok := renderMap[defID]; ok {
				rendering = r
			}
			rank, table, title, subject := extractRenderingInfo(rendering)

			var values []string
			valRef := att.VALUES.ENUM_VALUE_REF
			if valRef != "" {
				if enumVal, ok := stager.Map_id_ENUM_VALUE[valRef]; ok {
					values = append(values, enumVal.LONG_NAME)
				}
			}
			valStr := strings.Join(values, ", ")

			attributes = append(attributes, SortedAttribute{
				Rank: rank, Name: def.Name, Value: valStr,
				ShowInTable: table, ShowInTitle: title, ShowInSubject: subject,
			})
		}
	}

	// Sort
	sort.Slice(attributes, func(i, j int) bool {
		if attributes[i].Rank != attributes[j].Rank {
			return attributes[i].Rank < attributes[j].Rank
		}
		return attributes[i].Name < attributes[j].Name
	})

	return attributes
}
