package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) UpdateAndCommitButtonStage() {
	stager.buttonStage.Reset()

	layout := new(button.Layout).Stage(stager.buttonStage)

	group1 := new(button.Group).Stage(stager.buttonStage)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	buttonGeneratesModel := button.NewButton(
		// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
		&GenerateModelButtonProxy{
			stager: stager,
		},
		"Generates go model",
		string(buttons.BUTTON_add_comment),
		"Generates go model",
	)

	group1.Buttons = append(group1.Buttons, buttonGeneratesModel)

	buttonExportStaticSite := button.NewButton(
		// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
		&ExportStaticSiteButtonProxy{
			stager: stager,
		},
		"Export Static Web Site",
		string(buttons.BUTTON_web),
		"Export Static Web Site",
	)

	group1.Buttons = append(group1.Buttons, buttonExportStaticSite)

	if stager.pathToOutputReqifFile != "" {
		buttonExportModifiedReqif := button.NewButton(
			// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
			&ExportModifiedReqifButtonProxy{
				stager: stager,
			},
			"Export Modified ReqIF",
			string(buttons.BUTTON_fact_check),
			"Export Modified ReqIF",
		)

		group1.Buttons = append(group1.Buttons, buttonExportModifiedReqif)
	}

	buttonExportRenderingCong := button.NewButton(
		&ExportRenderingConfButtonProxy{
			stager: stager,
		},
		"Export Rendering Configuration",
		string(buttons.BUTTON_fact_check),
		"Export Rendering Configuration",
	)

	group1.Buttons = append(group1.Buttons, buttonExportRenderingCong)

	stager.buttonStage.Commit()
}

type GenerateModelButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *GenerateModelButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.buttonStage
}

// OnAfterUpdateButton implements models.Target.
func (e *GenerateModelButtonProxy) OnAfterUpdateButton() {
	e.stager.modelGenerator.GenerateModels(e.stager)
}

type ExportStaticSiteButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *ExportStaticSiteButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.buttonStage
}

// OnAfterUpdateButton implements models.Target.
func (e *ExportStaticSiteButtonProxy) OnAfterUpdateButton() {

	e.stager.UpdateAndCommitSsgStage()
	e.stager.generatesSiteFromSSGStage()
}

type ExportModifiedReqifButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *ExportModifiedReqifButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.buttonStage
}

// OnAfterUpdateButton implements models.Target.
func (e *ExportModifiedReqifButtonProxy) OnAfterUpdateButton() {

	e.stager.reqifExporter.ExportReqif(e.stager)

	return
}

type ExportRenderingConfButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *ExportRenderingConfButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.buttonStage
}

// OnAfterUpdateButton implements models.Target.
func (e *ExportRenderingConfButtonProxy) OnAfterUpdateButton() {

	conf := e.stager.ToRenderingConfiguration(e.stager.pathToReqifFile)
	_ = conf
	return
}
