// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Test":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Test Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TestFormCallback(
			nil,
			probe,
			formGroup,
		)
		test := new(models.Test)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(test, formGroup, probe)
	}
	formStage.Commit()
}
