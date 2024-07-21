// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
	"github.com/fullstack-lang/gongreqif/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__TestFormCallback(
	test *models.Test,
	probe *Probe,
	formGroup *table.FormGroup,
) (testFormCallback *TestFormCallback) {
	testFormCallback = new(TestFormCallback)
	testFormCallback.probe = probe
	testFormCallback.test = test
	testFormCallback.formGroup = formGroup

	testFormCallback.CreationMode = (test == nil)

	return
}

type TestFormCallback struct {
	test *models.Test

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (testFormCallback *TestFormCallback) OnSave() {

	log.Println("TestFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	testFormCallback.probe.formStage.Checkout()

	if testFormCallback.test == nil {
		testFormCallback.test = new(models.Test).Stage(testFormCallback.probe.stageOfInterest)
	}
	test_ := testFormCallback.test
	_ = test_

	for _, formDiv := range testFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(test_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if testFormCallback.formGroup.HasSuppressButtonBeenPressed {
		test_.Unstage(testFormCallback.probe.stageOfInterest)
	}

	testFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Test](
		testFormCallback.probe,
	)
	testFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if testFormCallback.CreationMode || testFormCallback.formGroup.HasSuppressButtonBeenPressed {
		testFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(testFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TestFormCallback(
			nil,
			testFormCallback.probe,
			newFormGroup,
		)
		test := new(models.Test)
		FillUpForm(test, newFormGroup, testFormCallback.probe)
		testFormCallback.probe.formStage.Commit()
	}

	fillUpTree(testFormCallback.probe)
}
