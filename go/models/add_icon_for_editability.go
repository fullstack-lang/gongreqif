package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func AddIconForEditabilityOfAttribute(attributeIS_EDITABLE bool, attributeLONG_NAME string, nodeAttribute *tree.Node) {
	if attributeIS_EDITABLE {
		nodeAttribute.Buttons = append(nodeAttribute.Buttons,
			&tree.Button{
				Name:            attributeLONG_NAME + " " + string(buttons.BUTTON_edit),
				Icon:            string(buttons.BUTTON_edit),
				HasToolTip:      true,
				ToolTipText:     "REQIF specifies that this element can be modified in a receiving requirements management tool after a data exchange",
				ToolTipPosition: tree.Above,
			})
	} else {
		nodeAttribute.Buttons = append(nodeAttribute.Buttons,
			&tree.Button{
				Name:            attributeLONG_NAME + " " + string(buttons.BUTTON_edit_off),
				Icon:            string(buttons.BUTTON_edit_off),
				HasToolTip:      true,
				ToolTipText:     "REQIF specifies that this element cannot be modified in a receiving requirements management tool after a data exchange",
				ToolTipPosition: tree.Above,
			})
	}
}
