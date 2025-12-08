package spectypes

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/gongreqif/go/icons"
	m "github.com/fullstack-lang/gongreqif/go/models"
)

// configureAndAddAttributeNode configures a new attribute node, sets its icon,
// and adds it to the parent node in the tree.
func configureAndAddAttributeNode[AttrDef m.AttributeDefinition](
	stager *m.Stager,
	nodeSpecType *tree.Node,
	nodeAttribute *tree.Node,
	nbInstances int,
	isEditable bool,
	longName string,
	attributeDefinition AttrDef,
) {
	if nbInstances > 0 {
		nodeAttribute.IsWithPreceedingIcon = true
		nodeAttribute.PreceedingIcon = string(buttons.BUTTON_check_circle)
	}
	nodeSpecType.Children = append(nodeSpecType.Children, nodeAttribute)
	m.AddIconForEditabilityOfAttribute(isEditable, longName, nodeAttribute)

	{
		button := &tree.Button{
			Name: longName + ": show attribute on/off in title",
			Impl: &ButtonToggleShowAttributeFieldInTitleProxy[AttrDef]{
				stager:              stager,
				attributeDefinition: attributeDefinition,
			},
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
		}

		var showInTitle bool
		switch any(attributeDefinition).(type) {
		case *m.ATTRIBUTE_DEFINITION_XHTML:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_XHTML_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTitle = attrDefRendering.ShowInTitle
		case *m.ATTRIBUTE_DEFINITION_STRING:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_STRING_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTitle = attrDefRendering.ShowInTitle
		case *m.ATTRIBUTE_DEFINITION_BOOLEAN:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTitle = attrDefRendering.ShowInTitle
		case *m.ATTRIBUTE_DEFINITION_INTEGER:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_INTEGER_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTitle = attrDefRendering.ShowInTitle
		case *m.ATTRIBUTE_DEFINITION_DATE:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_DATE_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTitle = attrDefRendering.ShowInTitle
		case *m.ATTRIBUTE_DEFINITION_REAL:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_REAL_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTitle = attrDefRendering.ShowInTitle
		case *m.ATTRIBUTE_DEFINITION_ENUMERATION:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTitle = attrDefRendering.ShowInTitle
		}

		if showInTitle {
			button.SVGIcon = icons.SvgIconTitleOff
			button.ToolTipText = "Click to remove attributes from the title"
		} else {
			button.SVGIcon = icons.SvgIconTitle
			button.ToolTipText = "Click to add attribute to the title"
		}
		nodeAttribute.Buttons = append(nodeAttribute.Buttons,
			button,
		)
	}

	{
		button := &tree.Button{
			Name: longName + ": show attribute on/off in table",
			Impl: &ButtonToggleShowAttributeFieldInTableProxy[AttrDef]{
				stager:              stager,
				attributeDefinition: attributeDefinition,
			},
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
		}

		var showInTable bool
		switch any(attributeDefinition).(type) {
		case *m.ATTRIBUTE_DEFINITION_XHTML:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_XHTML_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTable = attrDefRendering.ShowInTable
		case *m.ATTRIBUTE_DEFINITION_STRING:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_STRING_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTable = attrDefRendering.ShowInTable
		case *m.ATTRIBUTE_DEFINITION_BOOLEAN:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTable = attrDefRendering.ShowInTable
		case *m.ATTRIBUTE_DEFINITION_INTEGER:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_INTEGER_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTable = attrDefRendering.ShowInTable
		case *m.ATTRIBUTE_DEFINITION_DATE:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_DATE_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTable = attrDefRendering.ShowInTable
		case *m.ATTRIBUTE_DEFINITION_REAL:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_REAL_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTable = attrDefRendering.ShowInTable
		case *m.ATTRIBUTE_DEFINITION_ENUMERATION:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stager.GetStage(),
				attributeDefinition)
			showInTable = attrDefRendering.ShowInTable
		}

		if showInTable {
			button.SVGIcon = icons.SvgIconTableOff
			button.ToolTipText = "Click to remove attributes from the table"
		} else {
			button.SVGIcon = icons.SvgIconTable
			button.ToolTipText = "Click to have attributes in the table"
		}
		nodeAttribute.Buttons = append(nodeAttribute.Buttons,
			button,
		)
	}

	{
		button := &tree.Button{
			Name: longName + ": show attribute on/off in subject",
			Impl: &ButtonToggleShowAttributeFieldInSubjectProxy[AttrDef]{
				stager:              stager,
				attributeDefinition: attributeDefinition,
			},
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
		}

		var showInSubject bool
		switch any(attributeDefinition).(type) {
		case *m.ATTRIBUTE_DEFINITION_XHTML:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_XHTML_Rendering](stager.GetStage(),
				attributeDefinition)
			showInSubject = attrDefRendering.ShowInSubject
		case *m.ATTRIBUTE_DEFINITION_STRING:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_STRING_Rendering](stager.GetStage(),
				attributeDefinition)
			showInSubject = attrDefRendering.ShowInSubject
		case *m.ATTRIBUTE_DEFINITION_BOOLEAN:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stager.GetStage(),
				attributeDefinition)
			showInSubject = attrDefRendering.ShowInSubject
		case *m.ATTRIBUTE_DEFINITION_INTEGER:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_INTEGER_Rendering](stager.GetStage(),
				attributeDefinition)
			showInSubject = attrDefRendering.ShowInSubject
		case *m.ATTRIBUTE_DEFINITION_DATE:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_DATE_Rendering](stager.GetStage(),
				attributeDefinition)
			showInSubject = attrDefRendering.ShowInSubject
		case *m.ATTRIBUTE_DEFINITION_REAL:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_REAL_Rendering](stager.GetStage(),
				attributeDefinition)
			showInSubject = attrDefRendering.ShowInSubject
		case *m.ATTRIBUTE_DEFINITION_ENUMERATION:
			attrDefRendering := GetSpecAttributeDefinitionRendering[
				AttrDef, *m.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stager.GetStage(),
				attributeDefinition)
			showInSubject = attrDefRendering.ShowInSubject
		}

		if showInSubject {
			button.SVGIcon = icons.SvgIconSubjectOff
			button.ToolTipText = "Click to remove attributes from the subject"
		} else {
			button.SVGIcon = icons.SvgIconSubject
			button.ToolTipText = "Click to have attributes in the subject"
		}
		nodeAttribute.Buttons = append(nodeAttribute.Buttons,
			button,
		)
	}
}
