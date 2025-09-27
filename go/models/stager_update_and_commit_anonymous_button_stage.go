package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) UpdateAndCommitAnonymousButtonStage() {
	stager.anonymousButtonStage.Reset()
	stage := stager.anonymousButtonStage

	layout := new(button.Layout).Stage(stage)

	group1 := new(button.Group).Stage(stage)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	buttonExportRenderingCong := button.NewButton(
		&AnonymousButtonProxy{
			stager: stager,
		},
		"Export anonymoused version",
		string(buttons.BUTTON_shuffle),
		"Export anonymoused version",
	)

	group1.Buttons = append(group1.Buttons, buttonExportRenderingCong)

	stage.Commit()
}

type AnonymousButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *AnonymousButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.anonymousButtonStage
}

// OnAfterUpdateButton implements models.Target.
func (e *AnonymousButtonProxy) OnAfterUpdateButton() {

	e.stager.reqifExporter.ExportAnonymousReqif(e.stager)

}
