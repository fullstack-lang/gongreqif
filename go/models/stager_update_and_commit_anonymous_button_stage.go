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

	group1.Buttons = append(group1.Buttons,
		button.NewButton(
			&AnonymousButtonProxy{
				stager: stager,
			},
			"Export blanked version",
			string(buttons.BUTTON_shuffle),
			"Export blanked version",
		))

	group1.Buttons = append(group1.Buttons,
		button.NewButton(
			&LoadSampleButtonProxy{
				stager: stager,
			},
			"Load sample (collecting drone)",
			string(buttons.BUTTON_shuffle),
			"Load sample (collecting drone)",
		))

	stage.Commit()
}

type AnonymousButtonProxy struct {
	stager *Stager
}

func (e *AnonymousButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.anonymousButtonStage
}

func (e *AnonymousButtonProxy) OnAfterUpdateButton() {
	e.stager.reqifExporter.ExportAnonymousReqif(e.stager)
}

type LoadSampleButtonProxy struct {
	stager *Stager
}

func (e *LoadSampleButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.anonymousButtonStage
}

func (e *LoadSampleButtonProxy) OnAfterUpdateButton() {
	// e.stager.reqifExporter.ExportLoadSampleReqif(e.stager)

}
