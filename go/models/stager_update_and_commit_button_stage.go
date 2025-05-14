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

	stager.buttonStage.Commit()

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
	e.stager.ModelGenerator.GenerateModels(e.stager)
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
}
