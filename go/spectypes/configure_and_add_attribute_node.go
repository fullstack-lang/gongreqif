package spectypes

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/gongreqif/go/icons"
	m "github.com/fullstack-lang/gongreqif/go/models"
)

// configureAndAddAttributeNode configures a new attribute node, sets its icon,
// and returns the rank for sorting.
// It DOES NOT add the node to the parent 'nodeSpecType' anymore.
func configureAndAddAttributeNode[AttrDef m.AttributeDefinition](
	stager *m.Stager,
	nodeAttribute *tree.Node,
	nbInstances int,
	isEditable bool,
	longName string,
	attributeDefinition AttrDef,
) (rank int) {
	if nbInstances > 0 {
		nodeAttribute.IsWithPreceedingIcon = true
		nodeAttribute.PreceedingIcon = string(buttons.BUTTON_check_circle)
	}
	m.AddIconForEditabilityOfAttribute(isEditable, longName, nodeAttribute)

	attrDefRendering := GetAttrDefRendering(stager, attributeDefinition)
	rank = *attrDefRendering.GetRankPtr()

	{
		button := &tree.Button{
			Name: longName + ": show attribute on/off in title",
			Impl: &toggleButtonProxy{
				stager:      stager,
				toggleValue: attrDefRendering.GetShowInTitlePtr(),
			},
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
		}

		if *attrDefRendering.GetShowInTitlePtr() {
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
			Impl: &toggleButtonProxy{
				stager:      stager,
				toggleValue: attrDefRendering.GetShowInTablePtr(),
			},
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
		}

		if *attrDefRendering.GetShowInTablePtr() {
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
			Impl: &toggleButtonProxy{
				stager:      stager,
				toggleValue: attrDefRendering.GetShowInSubjectPtr(),
			},
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
		}

		if *attrDefRendering.GetShowInSubjectPtr() {
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

	return
}
