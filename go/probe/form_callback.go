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
func __gong__New__ALTERNATIVE_IDFormCallback(
	alternative_id *models.ALTERNATIVE_ID,
	probe *Probe,
	formGroup *table.FormGroup,
) (alternative_idFormCallback *ALTERNATIVE_IDFormCallback) {
	alternative_idFormCallback = new(ALTERNATIVE_IDFormCallback)
	alternative_idFormCallback.probe = probe
	alternative_idFormCallback.alternative_id = alternative_id
	alternative_idFormCallback.formGroup = formGroup

	alternative_idFormCallback.CreationMode = (alternative_id == nil)

	return
}

type ALTERNATIVE_IDFormCallback struct {
	alternative_id *models.ALTERNATIVE_ID

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (alternative_idFormCallback *ALTERNATIVE_IDFormCallback) OnSave() {

	log.Println("ALTERNATIVE_IDFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	alternative_idFormCallback.probe.formStage.Checkout()

	if alternative_idFormCallback.alternative_id == nil {
		alternative_idFormCallback.alternative_id = new(models.ALTERNATIVE_ID).Stage(alternative_idFormCallback.probe.stageOfInterest)
	}
	alternative_id_ := alternative_idFormCallback.alternative_id
	_ = alternative_id_

	for _, formDiv := range alternative_idFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(alternative_id_.Name), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(alternative_id_.IDENTIFIER), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_BOOLEANOwner *models.ATTRIBUTE_DEFINITION_BOOLEAN
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_BOOLEANOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_BOOLEAN)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_BOOLEANOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_boolean := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_BOOLEAN](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_boolean.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_BOOLEANOwner := _attribute_definition_boolean // we have a match
						if pastATTRIBUTE_DEFINITION_BOOLEANOwner != nil {
							if newATTRIBUTE_DEFINITION_BOOLEANOwner != pastATTRIBUTE_DEFINITION_BOOLEANOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_DATE:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_DATEOwner *models.ATTRIBUTE_DEFINITION_DATE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_DATEOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_DATE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_DATEOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_date := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_DATE](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_date.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_DATEOwner := _attribute_definition_date // we have a match
						if pastATTRIBUTE_DEFINITION_DATEOwner != nil {
							if newATTRIBUTE_DEFINITION_DATEOwner != pastATTRIBUTE_DEFINITION_DATEOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_ENUMERATIONOwner *models.ATTRIBUTE_DEFINITION_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_ENUMERATIONOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_ENUMERATIONOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_enumeration := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_ENUMERATION](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_ENUMERATIONOwner := _attribute_definition_enumeration // we have a match
						if pastATTRIBUTE_DEFINITION_ENUMERATIONOwner != nil {
							if newATTRIBUTE_DEFINITION_ENUMERATIONOwner != pastATTRIBUTE_DEFINITION_ENUMERATIONOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_INTEGER:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_INTEGEROwner *models.ATTRIBUTE_DEFINITION_INTEGER
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_INTEGEROwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_INTEGER)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_INTEGEROwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_integer := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_INTEGER](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_integer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_INTEGEROwner := _attribute_definition_integer // we have a match
						if pastATTRIBUTE_DEFINITION_INTEGEROwner != nil {
							if newATTRIBUTE_DEFINITION_INTEGEROwner != pastATTRIBUTE_DEFINITION_INTEGEROwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_REAL:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_REALOwner *models.ATTRIBUTE_DEFINITION_REAL
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_REALOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_REAL)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_REALOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_real := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_REAL](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_real.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_REALOwner := _attribute_definition_real // we have a match
						if pastATTRIBUTE_DEFINITION_REALOwner != nil {
							if newATTRIBUTE_DEFINITION_REALOwner != pastATTRIBUTE_DEFINITION_REALOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_STRING:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_STRINGOwner *models.ATTRIBUTE_DEFINITION_STRING
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_STRINGOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_STRING)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_STRINGOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_string := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_STRING](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_string.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_STRINGOwner := _attribute_definition_string // we have a match
						if pastATTRIBUTE_DEFINITION_STRINGOwner != nil {
							if newATTRIBUTE_DEFINITION_STRINGOwner != pastATTRIBUTE_DEFINITION_STRINGOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_XHTML:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_XHTMLOwner *models.ATTRIBUTE_DEFINITION_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_XHTMLOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_XHTMLOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_xhtml := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_XHTML](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_XHTMLOwner := _attribute_definition_xhtml // we have a match
						if pastATTRIBUTE_DEFINITION_XHTMLOwner != nil {
							if newATTRIBUTE_DEFINITION_XHTMLOwner != pastATTRIBUTE_DEFINITION_XHTMLOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_BOOLEAN:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_BOOLEANOwner *models.DATATYPE_DEFINITION_BOOLEAN
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_BOOLEAN"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_BOOLEANOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_BOOLEAN)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_BOOLEANOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_boolean := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_BOOLEAN](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_boolean.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_BOOLEANOwner := _datatype_definition_boolean // we have a match
						if pastDATATYPE_DEFINITION_BOOLEANOwner != nil {
							if newDATATYPE_DEFINITION_BOOLEANOwner != pastDATATYPE_DEFINITION_BOOLEANOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_DATE:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_DATEOwner *models.DATATYPE_DEFINITION_DATE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_DATE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_DATEOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_DATE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_DATEOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_date := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_DATE](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_date.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_DATEOwner := _datatype_definition_date // we have a match
						if pastDATATYPE_DEFINITION_DATEOwner != nil {
							if newDATATYPE_DEFINITION_DATEOwner != pastDATATYPE_DEFINITION_DATEOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_ENUMERATION:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_ENUMERATIONOwner *models.DATATYPE_DEFINITION_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_ENUMERATION"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_ENUMERATIONOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_ENUMERATIONOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_enumeration := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_ENUMERATION](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_ENUMERATIONOwner := _datatype_definition_enumeration // we have a match
						if pastDATATYPE_DEFINITION_ENUMERATIONOwner != nil {
							if newDATATYPE_DEFINITION_ENUMERATIONOwner != pastDATATYPE_DEFINITION_ENUMERATIONOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_INTEGER:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_INTEGEROwner *models.DATATYPE_DEFINITION_INTEGER
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_INTEGER"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_INTEGEROwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_INTEGER)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_INTEGEROwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_integer := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_INTEGER](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_integer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_INTEGEROwner := _datatype_definition_integer // we have a match
						if pastDATATYPE_DEFINITION_INTEGEROwner != nil {
							if newDATATYPE_DEFINITION_INTEGEROwner != pastDATATYPE_DEFINITION_INTEGEROwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_REAL:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_REALOwner *models.DATATYPE_DEFINITION_REAL
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_REAL"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_REALOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_REAL)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_REALOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_real := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_REAL](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_real.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_REALOwner := _datatype_definition_real // we have a match
						if pastDATATYPE_DEFINITION_REALOwner != nil {
							if newDATATYPE_DEFINITION_REALOwner != pastDATATYPE_DEFINITION_REALOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_STRING:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_STRINGOwner *models.DATATYPE_DEFINITION_STRING
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_STRING"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_STRINGOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_STRING)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_STRINGOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_string := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_STRING](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_string.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_STRINGOwner := _datatype_definition_string // we have a match
						if pastDATATYPE_DEFINITION_STRINGOwner != nil {
							if newDATATYPE_DEFINITION_STRINGOwner != pastDATATYPE_DEFINITION_STRINGOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_XHTML:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_XHTMLOwner *models.DATATYPE_DEFINITION_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_XHTML"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_XHTMLOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_XHTMLOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_xhtml := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_XHTML](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_XHTMLOwner := _datatype_definition_xhtml // we have a match
						if pastDATATYPE_DEFINITION_XHTMLOwner != nil {
							if newDATATYPE_DEFINITION_XHTMLOwner != pastDATATYPE_DEFINITION_XHTMLOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "ENUM_VALUE:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastENUM_VALUEOwner *models.ENUM_VALUE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ENUM_VALUE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastENUM_VALUEOwner = reverseFieldOwner.(*models.ENUM_VALUE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastENUM_VALUEOwner != nil {
					idx := slices.Index(pastENUM_VALUEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastENUM_VALUEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastENUM_VALUEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _enum_value := range *models.GetGongstructInstancesSet[models.ENUM_VALUE](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _enum_value.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newENUM_VALUEOwner := _enum_value // we have a match
						if pastENUM_VALUEOwner != nil {
							if newENUM_VALUEOwner != pastENUM_VALUEOwner {
								idx := slices.Index(pastENUM_VALUEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastENUM_VALUEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastENUM_VALUEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newENUM_VALUEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newENUM_VALUEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newENUM_VALUEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newENUM_VALUEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "RELATION_GROUP:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUPOwner *models.RELATION_GROUP
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUPOwner = reverseFieldOwner.(*models.RELATION_GROUP)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUPOwner != nil {
					idx := slices.Index(pastRELATION_GROUPOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastRELATION_GROUPOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastRELATION_GROUPOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group := range *models.GetGongstructInstancesSet[models.RELATION_GROUP](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUPOwner := _relation_group // we have a match
						if pastRELATION_GROUPOwner != nil {
							if newRELATION_GROUPOwner != pastRELATION_GROUPOwner {
								idx := slices.Index(pastRELATION_GROUPOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastRELATION_GROUPOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastRELATION_GROUPOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newRELATION_GROUPOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newRELATION_GROUPOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newRELATION_GROUPOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newRELATION_GROUPOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "RELATION_GROUP_TYPE:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUP_TYPEOwner *models.RELATION_GROUP_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUP_TYPEOwner = reverseFieldOwner.(*models.RELATION_GROUP_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUP_TYPEOwner != nil {
					idx := slices.Index(pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group_type := range *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUP_TYPEOwner := _relation_group_type // we have a match
						if pastRELATION_GROUP_TYPEOwner != nil {
							if newRELATION_GROUP_TYPEOwner != pastRELATION_GROUP_TYPEOwner {
								idx := slices.Index(pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "SPECIFICATION:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastSPECIFICATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPECIFICATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastSPECIFICATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPECIFICATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newSPECIFICATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPECIFICATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newSPECIFICATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPECIFICATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "SPECIFICATION_TYPE:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATION_TYPEOwner *models.SPECIFICATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATION_TYPEOwner = reverseFieldOwner.(*models.SPECIFICATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATION_TYPEOwner != nil {
					idx := slices.Index(pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification_type := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATION_TYPEOwner := _specification_type // we have a match
						if pastSPECIFICATION_TYPEOwner != nil {
							if newSPECIFICATION_TYPEOwner != pastSPECIFICATION_TYPEOwner {
								idx := slices.Index(pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newSPECIFICATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPECIFICATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newSPECIFICATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPECIFICATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "SPEC_HIERARCHY:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPEC_HIERARCHYOwner *models.SPEC_HIERARCHY
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_HIERARCHYOwner = reverseFieldOwner.(*models.SPEC_HIERARCHY)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_HIERARCHYOwner != nil {
					idx := slices.Index(pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_hierarchy := range *models.GetGongstructInstancesSet[models.SPEC_HIERARCHY](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_hierarchy.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_HIERARCHYOwner := _spec_hierarchy // we have a match
						if pastSPEC_HIERARCHYOwner != nil {
							if newSPEC_HIERARCHYOwner != pastSPEC_HIERARCHYOwner {
								idx := slices.Index(pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newSPEC_HIERARCHYOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPEC_HIERARCHYOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newSPEC_HIERARCHYOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPEC_HIERARCHYOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "SPEC_OBJECT:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastSPEC_OBJECTOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPEC_OBJECTOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastSPEC_OBJECTOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPEC_OBJECTOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newSPEC_OBJECTOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPEC_OBJECTOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newSPEC_OBJECTOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPEC_OBJECTOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "SPEC_OBJECT_TYPE:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECT_TYPEOwner *models.SPEC_OBJECT_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECT_TYPEOwner = reverseFieldOwner.(*models.SPEC_OBJECT_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECT_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object_type := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECT_TYPEOwner := _spec_object_type // we have a match
						if pastSPEC_OBJECT_TYPEOwner != nil {
							if newSPEC_OBJECT_TYPEOwner != pastSPEC_OBJECT_TYPEOwner {
								idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "SPEC_RELATION:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastSPEC_RELATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPEC_RELATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastSPEC_RELATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPEC_RELATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newSPEC_RELATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPEC_RELATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newSPEC_RELATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPEC_RELATIONOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		case "SPEC_RELATION_TYPE:ALTERNATIVE_ID.ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATION_TYPEOwner *models.SPEC_RELATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID.ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATION_TYPEOwner = reverseFieldOwner.(*models.SPEC_RELATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATION_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
					pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation_type := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATION_TYPEOwner := _spec_relation_type // we have a match
						if pastSPEC_RELATION_TYPEOwner != nil {
							if newSPEC_RELATION_TYPEOwner != pastSPEC_RELATION_TYPEOwner {
								idx := slices.Index(pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
								pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = slices.Delete(pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, idx, idx+1)
								newSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID = append(newSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if alternative_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		alternative_id_.Unstage(alternative_idFormCallback.probe.stageOfInterest)
	}

	alternative_idFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ALTERNATIVE_ID](
		alternative_idFormCallback.probe,
	)
	alternative_idFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if alternative_idFormCallback.CreationMode || alternative_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		alternative_idFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(alternative_idFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ALTERNATIVE_IDFormCallback(
			nil,
			alternative_idFormCallback.probe,
			newFormGroup,
		)
		alternative_id := new(models.ALTERNATIVE_ID)
		FillUpForm(alternative_id, newFormGroup, alternative_idFormCallback.probe)
		alternative_idFormCallback.probe.formStage.Commit()
	}

	fillUpTree(alternative_idFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
	attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_booleanFormCallback *ATTRIBUTE_DEFINITION_BOOLEANFormCallback) {
	attribute_definition_booleanFormCallback = new(ATTRIBUTE_DEFINITION_BOOLEANFormCallback)
	attribute_definition_booleanFormCallback.probe = probe
	attribute_definition_booleanFormCallback.attribute_definition_boolean = attribute_definition_boolean
	attribute_definition_booleanFormCallback.formGroup = formGroup

	attribute_definition_booleanFormCallback.CreationMode = (attribute_definition_boolean == nil)

	return
}

type ATTRIBUTE_DEFINITION_BOOLEANFormCallback struct {
	attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_booleanFormCallback *ATTRIBUTE_DEFINITION_BOOLEANFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_BOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_booleanFormCallback.probe.formStage.Checkout()

	if attribute_definition_booleanFormCallback.attribute_definition_boolean == nil {
		attribute_definition_booleanFormCallback.attribute_definition_boolean = new(models.ATTRIBUTE_DEFINITION_BOOLEAN).Stage(attribute_definition_booleanFormCallback.probe.stageOfInterest)
	}
	attribute_definition_boolean_ := attribute_definition_booleanFormCallback.attribute_definition_boolean
	_ = attribute_definition_boolean_

	for _, formDiv := range attribute_definition_booleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.DESC), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.LONG_NAME), formDiv)
		case "RELATION_GROUP_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUP_TYPEOwner *models.RELATION_GROUP_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_booleanFormCallback.probe.stageOfInterest,
				attribute_definition_booleanFormCallback.probe.backRepoOfInterest,
				attribute_definition_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUP_TYPEOwner = reverseFieldOwner.(*models.RELATION_GROUP_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUP_TYPEOwner != nil {
					idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
					pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group_type := range *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](attribute_definition_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUP_TYPEOwner := _relation_group_type // we have a match
						if pastRELATION_GROUP_TYPEOwner != nil {
							if newRELATION_GROUP_TYPEOwner != pastRELATION_GROUP_TYPEOwner {
								idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
								pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
								newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
							}
						} else {
							newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
						}
					}
				}
			}
		case "SPECIFICATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATION_TYPEOwner *models.SPECIFICATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_booleanFormCallback.probe.stageOfInterest,
				attribute_definition_booleanFormCallback.probe.backRepoOfInterest,
				attribute_definition_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATION_TYPEOwner = reverseFieldOwner.(*models.SPECIFICATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATION_TYPEOwner != nil {
					idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
					pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification_type := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](attribute_definition_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATION_TYPEOwner := _specification_type // we have a match
						if pastSPECIFICATION_TYPEOwner != nil {
							if newSPECIFICATION_TYPEOwner != pastSPECIFICATION_TYPEOwner {
								idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
								pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
								newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
							}
						} else {
							newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
						}
					}
				}
			}
		case "SPEC_OBJECT_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECT_TYPEOwner *models.SPEC_OBJECT_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_booleanFormCallback.probe.stageOfInterest,
				attribute_definition_booleanFormCallback.probe.backRepoOfInterest,
				attribute_definition_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECT_TYPEOwner = reverseFieldOwner.(*models.SPEC_OBJECT_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECT_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
					pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object_type := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](attribute_definition_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECT_TYPEOwner := _spec_object_type // we have a match
						if pastSPEC_OBJECT_TYPEOwner != nil {
							if newSPEC_OBJECT_TYPEOwner != pastSPEC_OBJECT_TYPEOwner {
								idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
								pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
								newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
							}
						} else {
							newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
						}
					}
				}
			}
		case "SPEC_RELATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATION_TYPEOwner *models.SPEC_RELATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_booleanFormCallback.probe.stageOfInterest,
				attribute_definition_booleanFormCallback.probe.backRepoOfInterest,
				attribute_definition_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATION_TYPEOwner = reverseFieldOwner.(*models.SPEC_RELATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATION_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
					pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation_type := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](attribute_definition_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATION_TYPEOwner := _spec_relation_type // we have a match
						if pastSPEC_RELATION_TYPEOwner != nil {
							if newSPEC_RELATION_TYPEOwner != pastSPEC_RELATION_TYPEOwner {
								idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
								pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
								newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
							}
						} else {
							newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_boolean_.Unstage(attribute_definition_booleanFormCallback.probe.stageOfInterest)
	}

	attribute_definition_booleanFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_BOOLEAN](
		attribute_definition_booleanFormCallback.probe,
	)
	attribute_definition_booleanFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_booleanFormCallback.CreationMode || attribute_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_booleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_booleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
			nil,
			attribute_definition_booleanFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_boolean := new(models.ATTRIBUTE_DEFINITION_BOOLEAN)
		FillUpForm(attribute_definition_boolean, newFormGroup, attribute_definition_booleanFormCallback.probe)
		attribute_definition_booleanFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_booleanFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
	attribute_definition_date *models.ATTRIBUTE_DEFINITION_DATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_dateFormCallback *ATTRIBUTE_DEFINITION_DATEFormCallback) {
	attribute_definition_dateFormCallback = new(ATTRIBUTE_DEFINITION_DATEFormCallback)
	attribute_definition_dateFormCallback.probe = probe
	attribute_definition_dateFormCallback.attribute_definition_date = attribute_definition_date
	attribute_definition_dateFormCallback.formGroup = formGroup

	attribute_definition_dateFormCallback.CreationMode = (attribute_definition_date == nil)

	return
}

type ATTRIBUTE_DEFINITION_DATEFormCallback struct {
	attribute_definition_date *models.ATTRIBUTE_DEFINITION_DATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_dateFormCallback *ATTRIBUTE_DEFINITION_DATEFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_DATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_dateFormCallback.probe.formStage.Checkout()

	if attribute_definition_dateFormCallback.attribute_definition_date == nil {
		attribute_definition_dateFormCallback.attribute_definition_date = new(models.ATTRIBUTE_DEFINITION_DATE).Stage(attribute_definition_dateFormCallback.probe.stageOfInterest)
	}
	attribute_definition_date_ := attribute_definition_dateFormCallback.attribute_definition_date
	_ = attribute_definition_date_

	for _, formDiv := range attribute_definition_dateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_date_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_date_.DESC), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_date_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_date_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_date_.LONG_NAME), formDiv)
		case "RELATION_GROUP_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUP_TYPEOwner *models.RELATION_GROUP_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_dateFormCallback.probe.stageOfInterest,
				attribute_definition_dateFormCallback.probe.backRepoOfInterest,
				attribute_definition_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUP_TYPEOwner = reverseFieldOwner.(*models.RELATION_GROUP_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUP_TYPEOwner != nil {
					idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
					pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group_type := range *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](attribute_definition_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUP_TYPEOwner := _relation_group_type // we have a match
						if pastRELATION_GROUP_TYPEOwner != nil {
							if newRELATION_GROUP_TYPEOwner != pastRELATION_GROUP_TYPEOwner {
								idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
								pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
								newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
							}
						} else {
							newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
						}
					}
				}
			}
		case "SPECIFICATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATION_TYPEOwner *models.SPECIFICATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_dateFormCallback.probe.stageOfInterest,
				attribute_definition_dateFormCallback.probe.backRepoOfInterest,
				attribute_definition_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATION_TYPEOwner = reverseFieldOwner.(*models.SPECIFICATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATION_TYPEOwner != nil {
					idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
					pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification_type := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](attribute_definition_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATION_TYPEOwner := _specification_type // we have a match
						if pastSPECIFICATION_TYPEOwner != nil {
							if newSPECIFICATION_TYPEOwner != pastSPECIFICATION_TYPEOwner {
								idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
								pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
								newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
							}
						} else {
							newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
						}
					}
				}
			}
		case "SPEC_OBJECT_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECT_TYPEOwner *models.SPEC_OBJECT_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_dateFormCallback.probe.stageOfInterest,
				attribute_definition_dateFormCallback.probe.backRepoOfInterest,
				attribute_definition_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECT_TYPEOwner = reverseFieldOwner.(*models.SPEC_OBJECT_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECT_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
					pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object_type := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](attribute_definition_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECT_TYPEOwner := _spec_object_type // we have a match
						if pastSPEC_OBJECT_TYPEOwner != nil {
							if newSPEC_OBJECT_TYPEOwner != pastSPEC_OBJECT_TYPEOwner {
								idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
								pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
								newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
							}
						} else {
							newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
						}
					}
				}
			}
		case "SPEC_RELATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATION_TYPEOwner *models.SPEC_RELATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_dateFormCallback.probe.stageOfInterest,
				attribute_definition_dateFormCallback.probe.backRepoOfInterest,
				attribute_definition_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATION_TYPEOwner = reverseFieldOwner.(*models.SPEC_RELATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATION_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
					pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation_type := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](attribute_definition_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATION_TYPEOwner := _spec_relation_type // we have a match
						if pastSPEC_RELATION_TYPEOwner != nil {
							if newSPEC_RELATION_TYPEOwner != pastSPEC_RELATION_TYPEOwner {
								idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
								pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
								newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
							}
						} else {
							newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_date_.Unstage(attribute_definition_dateFormCallback.probe.stageOfInterest)
	}

	attribute_definition_dateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_DATE](
		attribute_definition_dateFormCallback.probe,
	)
	attribute_definition_dateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_dateFormCallback.CreationMode || attribute_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_dateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_dateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
			nil,
			attribute_definition_dateFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_date := new(models.ATTRIBUTE_DEFINITION_DATE)
		FillUpForm(attribute_definition_date, newFormGroup, attribute_definition_dateFormCallback.probe)
		attribute_definition_dateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_dateFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
	attribute_definition_enumeration *models.ATTRIBUTE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_enumerationFormCallback *ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback) {
	attribute_definition_enumerationFormCallback = new(ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback)
	attribute_definition_enumerationFormCallback.probe = probe
	attribute_definition_enumerationFormCallback.attribute_definition_enumeration = attribute_definition_enumeration
	attribute_definition_enumerationFormCallback.formGroup = formGroup

	attribute_definition_enumerationFormCallback.CreationMode = (attribute_definition_enumeration == nil)

	return
}

type ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback struct {
	attribute_definition_enumeration *models.ATTRIBUTE_DEFINITION_ENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_enumerationFormCallback *ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_enumerationFormCallback.probe.formStage.Checkout()

	if attribute_definition_enumerationFormCallback.attribute_definition_enumeration == nil {
		attribute_definition_enumerationFormCallback.attribute_definition_enumeration = new(models.ATTRIBUTE_DEFINITION_ENUMERATION).Stage(attribute_definition_enumerationFormCallback.probe.stageOfInterest)
	}
	attribute_definition_enumeration_ := attribute_definition_enumerationFormCallback.attribute_definition_enumeration
	_ = attribute_definition_enumeration_

	for _, formDiv := range attribute_definition_enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.DESC), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.LONG_NAME), formDiv)
		case "MULTI_VALUED":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.MULTI_VALUED), formDiv)
		case "RELATION_GROUP_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUP_TYPEOwner *models.RELATION_GROUP_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_enumerationFormCallback.probe.stageOfInterest,
				attribute_definition_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_definition_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUP_TYPEOwner = reverseFieldOwner.(*models.RELATION_GROUP_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUP_TYPEOwner != nil {
					idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
					pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group_type := range *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](attribute_definition_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUP_TYPEOwner := _relation_group_type // we have a match
						if pastRELATION_GROUP_TYPEOwner != nil {
							if newRELATION_GROUP_TYPEOwner != pastRELATION_GROUP_TYPEOwner {
								idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
								pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
								newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
							}
						} else {
							newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
						}
					}
				}
			}
		case "SPECIFICATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATION_TYPEOwner *models.SPECIFICATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_enumerationFormCallback.probe.stageOfInterest,
				attribute_definition_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_definition_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATION_TYPEOwner = reverseFieldOwner.(*models.SPECIFICATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATION_TYPEOwner != nil {
					idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
					pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification_type := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](attribute_definition_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATION_TYPEOwner := _specification_type // we have a match
						if pastSPECIFICATION_TYPEOwner != nil {
							if newSPECIFICATION_TYPEOwner != pastSPECIFICATION_TYPEOwner {
								idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
								pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
								newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
							}
						} else {
							newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
						}
					}
				}
			}
		case "SPEC_OBJECT_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECT_TYPEOwner *models.SPEC_OBJECT_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_enumerationFormCallback.probe.stageOfInterest,
				attribute_definition_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_definition_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECT_TYPEOwner = reverseFieldOwner.(*models.SPEC_OBJECT_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECT_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
					pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object_type := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](attribute_definition_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECT_TYPEOwner := _spec_object_type // we have a match
						if pastSPEC_OBJECT_TYPEOwner != nil {
							if newSPEC_OBJECT_TYPEOwner != pastSPEC_OBJECT_TYPEOwner {
								idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
								pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
								newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
							}
						} else {
							newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
						}
					}
				}
			}
		case "SPEC_RELATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATION_TYPEOwner *models.SPEC_RELATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_enumerationFormCallback.probe.stageOfInterest,
				attribute_definition_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_definition_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATION_TYPEOwner = reverseFieldOwner.(*models.SPEC_RELATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATION_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
					pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation_type := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](attribute_definition_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATION_TYPEOwner := _spec_relation_type // we have a match
						if pastSPEC_RELATION_TYPEOwner != nil {
							if newSPEC_RELATION_TYPEOwner != pastSPEC_RELATION_TYPEOwner {
								idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
								pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
								newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
							}
						} else {
							newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_enumeration_.Unstage(attribute_definition_enumerationFormCallback.probe.stageOfInterest)
	}

	attribute_definition_enumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_ENUMERATION](
		attribute_definition_enumerationFormCallback.probe,
	)
	attribute_definition_enumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_enumerationFormCallback.CreationMode || attribute_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
			nil,
			attribute_definition_enumerationFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_enumeration := new(models.ATTRIBUTE_DEFINITION_ENUMERATION)
		FillUpForm(attribute_definition_enumeration, newFormGroup, attribute_definition_enumerationFormCallback.probe)
		attribute_definition_enumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_enumerationFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
	attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_integerFormCallback *ATTRIBUTE_DEFINITION_INTEGERFormCallback) {
	attribute_definition_integerFormCallback = new(ATTRIBUTE_DEFINITION_INTEGERFormCallback)
	attribute_definition_integerFormCallback.probe = probe
	attribute_definition_integerFormCallback.attribute_definition_integer = attribute_definition_integer
	attribute_definition_integerFormCallback.formGroup = formGroup

	attribute_definition_integerFormCallback.CreationMode = (attribute_definition_integer == nil)

	return
}

type ATTRIBUTE_DEFINITION_INTEGERFormCallback struct {
	attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_integerFormCallback *ATTRIBUTE_DEFINITION_INTEGERFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_INTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_integerFormCallback.probe.formStage.Checkout()

	if attribute_definition_integerFormCallback.attribute_definition_integer == nil {
		attribute_definition_integerFormCallback.attribute_definition_integer = new(models.ATTRIBUTE_DEFINITION_INTEGER).Stage(attribute_definition_integerFormCallback.probe.stageOfInterest)
	}
	attribute_definition_integer_ := attribute_definition_integerFormCallback.attribute_definition_integer
	_ = attribute_definition_integer_

	for _, formDiv := range attribute_definition_integerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_integer_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_integer_.DESC), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_integer_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_integer_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_integer_.LONG_NAME), formDiv)
		case "RELATION_GROUP_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUP_TYPEOwner *models.RELATION_GROUP_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_integerFormCallback.probe.stageOfInterest,
				attribute_definition_integerFormCallback.probe.backRepoOfInterest,
				attribute_definition_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUP_TYPEOwner = reverseFieldOwner.(*models.RELATION_GROUP_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUP_TYPEOwner != nil {
					idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
					pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group_type := range *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](attribute_definition_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUP_TYPEOwner := _relation_group_type // we have a match
						if pastRELATION_GROUP_TYPEOwner != nil {
							if newRELATION_GROUP_TYPEOwner != pastRELATION_GROUP_TYPEOwner {
								idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
								pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
								newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
							}
						} else {
							newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
						}
					}
				}
			}
		case "SPECIFICATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATION_TYPEOwner *models.SPECIFICATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_integerFormCallback.probe.stageOfInterest,
				attribute_definition_integerFormCallback.probe.backRepoOfInterest,
				attribute_definition_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATION_TYPEOwner = reverseFieldOwner.(*models.SPECIFICATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATION_TYPEOwner != nil {
					idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
					pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification_type := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](attribute_definition_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATION_TYPEOwner := _specification_type // we have a match
						if pastSPECIFICATION_TYPEOwner != nil {
							if newSPECIFICATION_TYPEOwner != pastSPECIFICATION_TYPEOwner {
								idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
								pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
								newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
							}
						} else {
							newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
						}
					}
				}
			}
		case "SPEC_OBJECT_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECT_TYPEOwner *models.SPEC_OBJECT_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_integerFormCallback.probe.stageOfInterest,
				attribute_definition_integerFormCallback.probe.backRepoOfInterest,
				attribute_definition_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECT_TYPEOwner = reverseFieldOwner.(*models.SPEC_OBJECT_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECT_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
					pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object_type := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](attribute_definition_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECT_TYPEOwner := _spec_object_type // we have a match
						if pastSPEC_OBJECT_TYPEOwner != nil {
							if newSPEC_OBJECT_TYPEOwner != pastSPEC_OBJECT_TYPEOwner {
								idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
								pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
								newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
							}
						} else {
							newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
						}
					}
				}
			}
		case "SPEC_RELATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATION_TYPEOwner *models.SPEC_RELATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_integerFormCallback.probe.stageOfInterest,
				attribute_definition_integerFormCallback.probe.backRepoOfInterest,
				attribute_definition_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATION_TYPEOwner = reverseFieldOwner.(*models.SPEC_RELATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATION_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
					pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation_type := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](attribute_definition_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATION_TYPEOwner := _spec_relation_type // we have a match
						if pastSPEC_RELATION_TYPEOwner != nil {
							if newSPEC_RELATION_TYPEOwner != pastSPEC_RELATION_TYPEOwner {
								idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
								pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
								newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
							}
						} else {
							newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_integer_.Unstage(attribute_definition_integerFormCallback.probe.stageOfInterest)
	}

	attribute_definition_integerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_INTEGER](
		attribute_definition_integerFormCallback.probe,
	)
	attribute_definition_integerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_integerFormCallback.CreationMode || attribute_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_integerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_integerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
			nil,
			attribute_definition_integerFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_integer := new(models.ATTRIBUTE_DEFINITION_INTEGER)
		FillUpForm(attribute_definition_integer, newFormGroup, attribute_definition_integerFormCallback.probe)
		attribute_definition_integerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_integerFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
	attribute_definition_real *models.ATTRIBUTE_DEFINITION_REAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_realFormCallback *ATTRIBUTE_DEFINITION_REALFormCallback) {
	attribute_definition_realFormCallback = new(ATTRIBUTE_DEFINITION_REALFormCallback)
	attribute_definition_realFormCallback.probe = probe
	attribute_definition_realFormCallback.attribute_definition_real = attribute_definition_real
	attribute_definition_realFormCallback.formGroup = formGroup

	attribute_definition_realFormCallback.CreationMode = (attribute_definition_real == nil)

	return
}

type ATTRIBUTE_DEFINITION_REALFormCallback struct {
	attribute_definition_real *models.ATTRIBUTE_DEFINITION_REAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_realFormCallback *ATTRIBUTE_DEFINITION_REALFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_REALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_realFormCallback.probe.formStage.Checkout()

	if attribute_definition_realFormCallback.attribute_definition_real == nil {
		attribute_definition_realFormCallback.attribute_definition_real = new(models.ATTRIBUTE_DEFINITION_REAL).Stage(attribute_definition_realFormCallback.probe.stageOfInterest)
	}
	attribute_definition_real_ := attribute_definition_realFormCallback.attribute_definition_real
	_ = attribute_definition_real_

	for _, formDiv := range attribute_definition_realFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_real_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_real_.DESC), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_real_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_real_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_real_.LONG_NAME), formDiv)
		case "RELATION_GROUP_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUP_TYPEOwner *models.RELATION_GROUP_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_realFormCallback.probe.stageOfInterest,
				attribute_definition_realFormCallback.probe.backRepoOfInterest,
				attribute_definition_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUP_TYPEOwner = reverseFieldOwner.(*models.RELATION_GROUP_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUP_TYPEOwner != nil {
					idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
					pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group_type := range *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](attribute_definition_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUP_TYPEOwner := _relation_group_type // we have a match
						if pastRELATION_GROUP_TYPEOwner != nil {
							if newRELATION_GROUP_TYPEOwner != pastRELATION_GROUP_TYPEOwner {
								idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
								pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
								newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
							}
						} else {
							newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
						}
					}
				}
			}
		case "SPECIFICATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATION_TYPEOwner *models.SPECIFICATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_realFormCallback.probe.stageOfInterest,
				attribute_definition_realFormCallback.probe.backRepoOfInterest,
				attribute_definition_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATION_TYPEOwner = reverseFieldOwner.(*models.SPECIFICATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATION_TYPEOwner != nil {
					idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
					pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification_type := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](attribute_definition_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATION_TYPEOwner := _specification_type // we have a match
						if pastSPECIFICATION_TYPEOwner != nil {
							if newSPECIFICATION_TYPEOwner != pastSPECIFICATION_TYPEOwner {
								idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
								pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
								newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
							}
						} else {
							newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
						}
					}
				}
			}
		case "SPEC_OBJECT_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECT_TYPEOwner *models.SPEC_OBJECT_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_realFormCallback.probe.stageOfInterest,
				attribute_definition_realFormCallback.probe.backRepoOfInterest,
				attribute_definition_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECT_TYPEOwner = reverseFieldOwner.(*models.SPEC_OBJECT_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECT_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
					pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object_type := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](attribute_definition_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECT_TYPEOwner := _spec_object_type // we have a match
						if pastSPEC_OBJECT_TYPEOwner != nil {
							if newSPEC_OBJECT_TYPEOwner != pastSPEC_OBJECT_TYPEOwner {
								idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
								pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
								newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
							}
						} else {
							newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
						}
					}
				}
			}
		case "SPEC_RELATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATION_TYPEOwner *models.SPEC_RELATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_realFormCallback.probe.stageOfInterest,
				attribute_definition_realFormCallback.probe.backRepoOfInterest,
				attribute_definition_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATION_TYPEOwner = reverseFieldOwner.(*models.SPEC_RELATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATION_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
					pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation_type := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](attribute_definition_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATION_TYPEOwner := _spec_relation_type // we have a match
						if pastSPEC_RELATION_TYPEOwner != nil {
							if newSPEC_RELATION_TYPEOwner != pastSPEC_RELATION_TYPEOwner {
								idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
								pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
								newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
							}
						} else {
							newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_real_.Unstage(attribute_definition_realFormCallback.probe.stageOfInterest)
	}

	attribute_definition_realFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_REAL](
		attribute_definition_realFormCallback.probe,
	)
	attribute_definition_realFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_realFormCallback.CreationMode || attribute_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_realFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_realFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
			nil,
			attribute_definition_realFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_real := new(models.ATTRIBUTE_DEFINITION_REAL)
		FillUpForm(attribute_definition_real, newFormGroup, attribute_definition_realFormCallback.probe)
		attribute_definition_realFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_realFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
	attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_stringFormCallback *ATTRIBUTE_DEFINITION_STRINGFormCallback) {
	attribute_definition_stringFormCallback = new(ATTRIBUTE_DEFINITION_STRINGFormCallback)
	attribute_definition_stringFormCallback.probe = probe
	attribute_definition_stringFormCallback.attribute_definition_string = attribute_definition_string
	attribute_definition_stringFormCallback.formGroup = formGroup

	attribute_definition_stringFormCallback.CreationMode = (attribute_definition_string == nil)

	return
}

type ATTRIBUTE_DEFINITION_STRINGFormCallback struct {
	attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_stringFormCallback *ATTRIBUTE_DEFINITION_STRINGFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_STRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_stringFormCallback.probe.formStage.Checkout()

	if attribute_definition_stringFormCallback.attribute_definition_string == nil {
		attribute_definition_stringFormCallback.attribute_definition_string = new(models.ATTRIBUTE_DEFINITION_STRING).Stage(attribute_definition_stringFormCallback.probe.stageOfInterest)
	}
	attribute_definition_string_ := attribute_definition_stringFormCallback.attribute_definition_string
	_ = attribute_definition_string_

	for _, formDiv := range attribute_definition_stringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_string_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_string_.DESC), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_string_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_string_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_string_.LONG_NAME), formDiv)
		case "RELATION_GROUP_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUP_TYPEOwner *models.RELATION_GROUP_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_stringFormCallback.probe.stageOfInterest,
				attribute_definition_stringFormCallback.probe.backRepoOfInterest,
				attribute_definition_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUP_TYPEOwner = reverseFieldOwner.(*models.RELATION_GROUP_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUP_TYPEOwner != nil {
					idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
					pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group_type := range *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](attribute_definition_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUP_TYPEOwner := _relation_group_type // we have a match
						if pastRELATION_GROUP_TYPEOwner != nil {
							if newRELATION_GROUP_TYPEOwner != pastRELATION_GROUP_TYPEOwner {
								idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
								pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
								newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
							}
						} else {
							newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
						}
					}
				}
			}
		case "SPECIFICATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATION_TYPEOwner *models.SPECIFICATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_stringFormCallback.probe.stageOfInterest,
				attribute_definition_stringFormCallback.probe.backRepoOfInterest,
				attribute_definition_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATION_TYPEOwner = reverseFieldOwner.(*models.SPECIFICATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATION_TYPEOwner != nil {
					idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
					pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification_type := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](attribute_definition_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATION_TYPEOwner := _specification_type // we have a match
						if pastSPECIFICATION_TYPEOwner != nil {
							if newSPECIFICATION_TYPEOwner != pastSPECIFICATION_TYPEOwner {
								idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
								pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
								newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
							}
						} else {
							newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
						}
					}
				}
			}
		case "SPEC_OBJECT_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECT_TYPEOwner *models.SPEC_OBJECT_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_stringFormCallback.probe.stageOfInterest,
				attribute_definition_stringFormCallback.probe.backRepoOfInterest,
				attribute_definition_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECT_TYPEOwner = reverseFieldOwner.(*models.SPEC_OBJECT_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECT_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
					pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object_type := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](attribute_definition_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECT_TYPEOwner := _spec_object_type // we have a match
						if pastSPEC_OBJECT_TYPEOwner != nil {
							if newSPEC_OBJECT_TYPEOwner != pastSPEC_OBJECT_TYPEOwner {
								idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
								pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
								newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
							}
						} else {
							newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
						}
					}
				}
			}
		case "SPEC_RELATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATION_TYPEOwner *models.SPEC_RELATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_stringFormCallback.probe.stageOfInterest,
				attribute_definition_stringFormCallback.probe.backRepoOfInterest,
				attribute_definition_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATION_TYPEOwner = reverseFieldOwner.(*models.SPEC_RELATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATION_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
					pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation_type := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](attribute_definition_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATION_TYPEOwner := _spec_relation_type // we have a match
						if pastSPEC_RELATION_TYPEOwner != nil {
							if newSPEC_RELATION_TYPEOwner != pastSPEC_RELATION_TYPEOwner {
								idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
								pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
								newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
							}
						} else {
							newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_string_.Unstage(attribute_definition_stringFormCallback.probe.stageOfInterest)
	}

	attribute_definition_stringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_STRING](
		attribute_definition_stringFormCallback.probe,
	)
	attribute_definition_stringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_stringFormCallback.CreationMode || attribute_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_stringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_stringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
			nil,
			attribute_definition_stringFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_string := new(models.ATTRIBUTE_DEFINITION_STRING)
		FillUpForm(attribute_definition_string, newFormGroup, attribute_definition_stringFormCallback.probe)
		attribute_definition_stringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_stringFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
	attribute_definition_xhtml *models.ATTRIBUTE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_xhtmlFormCallback *ATTRIBUTE_DEFINITION_XHTMLFormCallback) {
	attribute_definition_xhtmlFormCallback = new(ATTRIBUTE_DEFINITION_XHTMLFormCallback)
	attribute_definition_xhtmlFormCallback.probe = probe
	attribute_definition_xhtmlFormCallback.attribute_definition_xhtml = attribute_definition_xhtml
	attribute_definition_xhtmlFormCallback.formGroup = formGroup

	attribute_definition_xhtmlFormCallback.CreationMode = (attribute_definition_xhtml == nil)

	return
}

type ATTRIBUTE_DEFINITION_XHTMLFormCallback struct {
	attribute_definition_xhtml *models.ATTRIBUTE_DEFINITION_XHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_xhtmlFormCallback *ATTRIBUTE_DEFINITION_XHTMLFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_XHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_xhtmlFormCallback.probe.formStage.Checkout()

	if attribute_definition_xhtmlFormCallback.attribute_definition_xhtml == nil {
		attribute_definition_xhtmlFormCallback.attribute_definition_xhtml = new(models.ATTRIBUTE_DEFINITION_XHTML).Stage(attribute_definition_xhtmlFormCallback.probe.stageOfInterest)
	}
	attribute_definition_xhtml_ := attribute_definition_xhtmlFormCallback.attribute_definition_xhtml
	_ = attribute_definition_xhtml_

	for _, formDiv := range attribute_definition_xhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.DESC), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.LONG_NAME), formDiv)
		case "RELATION_GROUP_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUP_TYPEOwner *models.RELATION_GROUP_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_xhtmlFormCallback.probe.stageOfInterest,
				attribute_definition_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_definition_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUP_TYPEOwner = reverseFieldOwner.(*models.RELATION_GROUP_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUP_TYPEOwner != nil {
					idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
					pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group_type := range *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](attribute_definition_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUP_TYPEOwner := _relation_group_type // we have a match
						if pastRELATION_GROUP_TYPEOwner != nil {
							if newRELATION_GROUP_TYPEOwner != pastRELATION_GROUP_TYPEOwner {
								idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
								pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
								newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
							}
						} else {
							newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
						}
					}
				}
			}
		case "SPECIFICATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATION_TYPEOwner *models.SPECIFICATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_xhtmlFormCallback.probe.stageOfInterest,
				attribute_definition_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_definition_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATION_TYPEOwner = reverseFieldOwner.(*models.SPECIFICATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATION_TYPEOwner != nil {
					idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
					pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification_type := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](attribute_definition_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATION_TYPEOwner := _specification_type // we have a match
						if pastSPECIFICATION_TYPEOwner != nil {
							if newSPECIFICATION_TYPEOwner != pastSPECIFICATION_TYPEOwner {
								idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
								pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
								newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
							}
						} else {
							newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
						}
					}
				}
			}
		case "SPEC_OBJECT_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECT_TYPEOwner *models.SPEC_OBJECT_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_xhtmlFormCallback.probe.stageOfInterest,
				attribute_definition_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_definition_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECT_TYPEOwner = reverseFieldOwner.(*models.SPEC_OBJECT_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECT_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
					pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object_type := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](attribute_definition_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECT_TYPEOwner := _spec_object_type // we have a match
						if pastSPEC_OBJECT_TYPEOwner != nil {
							if newSPEC_OBJECT_TYPEOwner != pastSPEC_OBJECT_TYPEOwner {
								idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
								pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
								newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
							}
						} else {
							newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
						}
					}
				}
			}
		case "SPEC_RELATION_TYPE:SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATION_TYPEOwner *models.SPEC_RELATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_xhtmlFormCallback.probe.stageOfInterest,
				attribute_definition_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_definition_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATION_TYPEOwner = reverseFieldOwner.(*models.SPEC_RELATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATION_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
					pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation_type := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](attribute_definition_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATION_TYPEOwner := _spec_relation_type // we have a match
						if pastSPEC_RELATION_TYPEOwner != nil {
							if newSPEC_RELATION_TYPEOwner != pastSPEC_RELATION_TYPEOwner {
								idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
								pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
								newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
							}
						} else {
							newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_xhtml_.Unstage(attribute_definition_xhtmlFormCallback.probe.stageOfInterest)
	}

	attribute_definition_xhtmlFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_XHTML](
		attribute_definition_xhtmlFormCallback.probe,
	)
	attribute_definition_xhtmlFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_xhtmlFormCallback.CreationMode || attribute_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_xhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_xhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
			nil,
			attribute_definition_xhtmlFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_xhtml := new(models.ATTRIBUTE_DEFINITION_XHTML)
		FillUpForm(attribute_definition_xhtml, newFormGroup, attribute_definition_xhtmlFormCallback.probe)
		attribute_definition_xhtmlFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_xhtmlFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
	attribute_value_boolean *models.ATTRIBUTE_VALUE_BOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_booleanFormCallback *ATTRIBUTE_VALUE_BOOLEANFormCallback) {
	attribute_value_booleanFormCallback = new(ATTRIBUTE_VALUE_BOOLEANFormCallback)
	attribute_value_booleanFormCallback.probe = probe
	attribute_value_booleanFormCallback.attribute_value_boolean = attribute_value_boolean
	attribute_value_booleanFormCallback.formGroup = formGroup

	attribute_value_booleanFormCallback.CreationMode = (attribute_value_boolean == nil)

	return
}

type ATTRIBUTE_VALUE_BOOLEANFormCallback struct {
	attribute_value_boolean *models.ATTRIBUTE_VALUE_BOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_booleanFormCallback *ATTRIBUTE_VALUE_BOOLEANFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_BOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_booleanFormCallback.probe.formStage.Checkout()

	if attribute_value_booleanFormCallback.attribute_value_boolean == nil {
		attribute_value_booleanFormCallback.attribute_value_boolean = new(models.ATTRIBUTE_VALUE_BOOLEAN).Stage(attribute_value_booleanFormCallback.probe.stageOfInterest)
	}
	attribute_value_boolean_ := attribute_value_booleanFormCallback.attribute_value_boolean
	_ = attribute_value_boolean_

	for _, formDiv := range attribute_value_booleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_boolean_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_boolean_.THE_VALUE), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN:DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_BOOLEANOwner *models.ATTRIBUTE_DEFINITION_BOOLEAN
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_booleanFormCallback.probe.stageOfInterest,
				attribute_value_booleanFormCallback.probe.backRepoOfInterest,
				attribute_value_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_BOOLEANOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_BOOLEAN)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_BOOLEANOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
					pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_boolean := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_BOOLEAN](attribute_value_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_boolean.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_BOOLEANOwner := _attribute_definition_boolean // we have a match
						if pastATTRIBUTE_DEFINITION_BOOLEANOwner != nil {
							if newATTRIBUTE_DEFINITION_BOOLEANOwner != pastATTRIBUTE_DEFINITION_BOOLEANOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
								pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
								newATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN = append(newATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
							}
						} else {
							newATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN = append(newATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
						}
					}
				}
			}
		case "SPECIFICATION:VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_booleanFormCallback.probe.stageOfInterest,
				attribute_value_booleanFormCallback.probe.backRepoOfInterest,
				attribute_value_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
					pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](attribute_value_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
								pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
								newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
							}
						} else {
							newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
						}
					}
				}
			}
		case "SPEC_OBJECT:VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_booleanFormCallback.probe.stageOfInterest,
				attribute_value_booleanFormCallback.probe.backRepoOfInterest,
				attribute_value_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
					pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](attribute_value_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
								pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
								newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
							}
						} else {
							newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
						}
					}
				}
			}
		case "SPEC_RELATION:VALUES.ATTRIBUTE_VALUE_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_booleanFormCallback.probe.stageOfInterest,
				attribute_value_booleanFormCallback.probe.backRepoOfInterest,
				attribute_value_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
					pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](attribute_value_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
								pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
								newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
							}
						} else {
							newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_value_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_boolean_.Unstage(attribute_value_booleanFormCallback.probe.stageOfInterest)
	}

	attribute_value_booleanFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_BOOLEAN](
		attribute_value_booleanFormCallback.probe,
	)
	attribute_value_booleanFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_booleanFormCallback.CreationMode || attribute_value_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_booleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_booleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
			nil,
			attribute_value_booleanFormCallback.probe,
			newFormGroup,
		)
		attribute_value_boolean := new(models.ATTRIBUTE_VALUE_BOOLEAN)
		FillUpForm(attribute_value_boolean, newFormGroup, attribute_value_booleanFormCallback.probe)
		attribute_value_booleanFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_booleanFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
	attribute_value_date *models.ATTRIBUTE_VALUE_DATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_dateFormCallback *ATTRIBUTE_VALUE_DATEFormCallback) {
	attribute_value_dateFormCallback = new(ATTRIBUTE_VALUE_DATEFormCallback)
	attribute_value_dateFormCallback.probe = probe
	attribute_value_dateFormCallback.attribute_value_date = attribute_value_date
	attribute_value_dateFormCallback.formGroup = formGroup

	attribute_value_dateFormCallback.CreationMode = (attribute_value_date == nil)

	return
}

type ATTRIBUTE_VALUE_DATEFormCallback struct {
	attribute_value_date *models.ATTRIBUTE_VALUE_DATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_dateFormCallback *ATTRIBUTE_VALUE_DATEFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_DATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_dateFormCallback.probe.formStage.Checkout()

	if attribute_value_dateFormCallback.attribute_value_date == nil {
		attribute_value_dateFormCallback.attribute_value_date = new(models.ATTRIBUTE_VALUE_DATE).Stage(attribute_value_dateFormCallback.probe.stageOfInterest)
	}
	attribute_value_date_ := attribute_value_dateFormCallback.attribute_value_date
	_ = attribute_value_date_

	for _, formDiv := range attribute_value_dateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_date_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_date_.THE_VALUE), formDiv)
		case "ATTRIBUTE_DEFINITION_DATE:DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_DATEOwner *models.ATTRIBUTE_DEFINITION_DATE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_dateFormCallback.probe.stageOfInterest,
				attribute_value_dateFormCallback.probe.backRepoOfInterest,
				attribute_value_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_DATEOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_DATE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_DATEOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
					pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE = slices.Delete(pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_date := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_DATE](attribute_value_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_date.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_DATEOwner := _attribute_definition_date // we have a match
						if pastATTRIBUTE_DEFINITION_DATEOwner != nil {
							if newATTRIBUTE_DEFINITION_DATEOwner != pastATTRIBUTE_DEFINITION_DATEOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
								pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE = slices.Delete(pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE, idx, idx+1)
								newATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE = append(newATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
							}
						} else {
							newATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE = append(newATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
						}
					}
				}
			}
		case "SPECIFICATION:VALUES.ATTRIBUTE_VALUE_DATE":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_dateFormCallback.probe.stageOfInterest,
				attribute_value_dateFormCallback.probe.backRepoOfInterest,
				attribute_value_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
					pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](attribute_value_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
								pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, idx, idx+1)
								newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
							}
						} else {
							newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
						}
					}
				}
			}
		case "SPEC_OBJECT:VALUES.ATTRIBUTE_VALUE_DATE":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_dateFormCallback.probe.stageOfInterest,
				attribute_value_dateFormCallback.probe.backRepoOfInterest,
				attribute_value_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
					pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_DATE = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](attribute_value_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
								pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_DATE = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_DATE, idx, idx+1)
								newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_DATE = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
							}
						} else {
							newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_DATE = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
						}
					}
				}
			}
		case "SPEC_RELATION:VALUES.ATTRIBUTE_VALUE_DATE":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_dateFormCallback.probe.stageOfInterest,
				attribute_value_dateFormCallback.probe.backRepoOfInterest,
				attribute_value_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
					pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](attribute_value_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
								pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, idx, idx+1)
								newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
							}
						} else {
							newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_value_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_date_.Unstage(attribute_value_dateFormCallback.probe.stageOfInterest)
	}

	attribute_value_dateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_DATE](
		attribute_value_dateFormCallback.probe,
	)
	attribute_value_dateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_dateFormCallback.CreationMode || attribute_value_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_dateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_dateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
			nil,
			attribute_value_dateFormCallback.probe,
			newFormGroup,
		)
		attribute_value_date := new(models.ATTRIBUTE_VALUE_DATE)
		FillUpForm(attribute_value_date, newFormGroup, attribute_value_dateFormCallback.probe)
		attribute_value_dateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_dateFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
	attribute_value_enumeration *models.ATTRIBUTE_VALUE_ENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_enumerationFormCallback *ATTRIBUTE_VALUE_ENUMERATIONFormCallback) {
	attribute_value_enumerationFormCallback = new(ATTRIBUTE_VALUE_ENUMERATIONFormCallback)
	attribute_value_enumerationFormCallback.probe = probe
	attribute_value_enumerationFormCallback.attribute_value_enumeration = attribute_value_enumeration
	attribute_value_enumerationFormCallback.formGroup = formGroup

	attribute_value_enumerationFormCallback.CreationMode = (attribute_value_enumeration == nil)

	return
}

type ATTRIBUTE_VALUE_ENUMERATIONFormCallback struct {
	attribute_value_enumeration *models.ATTRIBUTE_VALUE_ENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_enumerationFormCallback *ATTRIBUTE_VALUE_ENUMERATIONFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_ENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_enumerationFormCallback.probe.formStage.Checkout()

	if attribute_value_enumerationFormCallback.attribute_value_enumeration == nil {
		attribute_value_enumerationFormCallback.attribute_value_enumeration = new(models.ATTRIBUTE_VALUE_ENUMERATION).Stage(attribute_value_enumerationFormCallback.probe.stageOfInterest)
	}
	attribute_value_enumeration_ := attribute_value_enumerationFormCallback.attribute_value_enumeration
	_ = attribute_value_enumeration_

	for _, formDiv := range attribute_value_enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_enumeration_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION:DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_ENUMERATIONOwner *models.ATTRIBUTE_DEFINITION_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_enumerationFormCallback.probe.stageOfInterest,
				attribute_value_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_value_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_ENUMERATIONOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_ENUMERATIONOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
					pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_enumeration := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_ENUMERATION](attribute_value_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_ENUMERATIONOwner := _attribute_definition_enumeration // we have a match
						if pastATTRIBUTE_DEFINITION_ENUMERATIONOwner != nil {
							if newATTRIBUTE_DEFINITION_ENUMERATIONOwner != pastATTRIBUTE_DEFINITION_ENUMERATIONOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
								pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
								newATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION = append(newATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
							}
						} else {
							newATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION = append(newATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
						}
					}
				}
			}
		case "SPECIFICATION:VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_enumerationFormCallback.probe.stageOfInterest,
				attribute_value_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_value_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
					pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](attribute_value_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
								pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
								newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
							}
						} else {
							newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
						}
					}
				}
			}
		case "SPEC_OBJECT:VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_enumerationFormCallback.probe.stageOfInterest,
				attribute_value_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_value_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
					pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](attribute_value_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
								pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
								newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
							}
						} else {
							newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
						}
					}
				}
			}
		case "SPEC_RELATION:VALUES.ATTRIBUTE_VALUE_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_enumerationFormCallback.probe.stageOfInterest,
				attribute_value_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_value_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
					pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](attribute_value_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
								pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
								newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
							}
						} else {
							newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_value_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_enumeration_.Unstage(attribute_value_enumerationFormCallback.probe.stageOfInterest)
	}

	attribute_value_enumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_ENUMERATION](
		attribute_value_enumerationFormCallback.probe,
	)
	attribute_value_enumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_enumerationFormCallback.CreationMode || attribute_value_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
			nil,
			attribute_value_enumerationFormCallback.probe,
			newFormGroup,
		)
		attribute_value_enumeration := new(models.ATTRIBUTE_VALUE_ENUMERATION)
		FillUpForm(attribute_value_enumeration, newFormGroup, attribute_value_enumerationFormCallback.probe)
		attribute_value_enumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_enumerationFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
	attribute_value_integer *models.ATTRIBUTE_VALUE_INTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_integerFormCallback *ATTRIBUTE_VALUE_INTEGERFormCallback) {
	attribute_value_integerFormCallback = new(ATTRIBUTE_VALUE_INTEGERFormCallback)
	attribute_value_integerFormCallback.probe = probe
	attribute_value_integerFormCallback.attribute_value_integer = attribute_value_integer
	attribute_value_integerFormCallback.formGroup = formGroup

	attribute_value_integerFormCallback.CreationMode = (attribute_value_integer == nil)

	return
}

type ATTRIBUTE_VALUE_INTEGERFormCallback struct {
	attribute_value_integer *models.ATTRIBUTE_VALUE_INTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_integerFormCallback *ATTRIBUTE_VALUE_INTEGERFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_INTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_integerFormCallback.probe.formStage.Checkout()

	if attribute_value_integerFormCallback.attribute_value_integer == nil {
		attribute_value_integerFormCallback.attribute_value_integer = new(models.ATTRIBUTE_VALUE_INTEGER).Stage(attribute_value_integerFormCallback.probe.stageOfInterest)
	}
	attribute_value_integer_ := attribute_value_integerFormCallback.attribute_value_integer
	_ = attribute_value_integer_

	for _, formDiv := range attribute_value_integerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_integer_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER:DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_INTEGEROwner *models.ATTRIBUTE_DEFINITION_INTEGER
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_integerFormCallback.probe.stageOfInterest,
				attribute_value_integerFormCallback.probe.backRepoOfInterest,
				attribute_value_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_INTEGEROwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_INTEGER)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_INTEGEROwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
					pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_integer := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_INTEGER](attribute_value_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_integer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_INTEGEROwner := _attribute_definition_integer // we have a match
						if pastATTRIBUTE_DEFINITION_INTEGEROwner != nil {
							if newATTRIBUTE_DEFINITION_INTEGEROwner != pastATTRIBUTE_DEFINITION_INTEGEROwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
								pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
								newATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER = append(newATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
							}
						} else {
							newATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER = append(newATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
						}
					}
				}
			}
		case "SPECIFICATION:VALUES.ATTRIBUTE_VALUE_INTEGER":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_integerFormCallback.probe.stageOfInterest,
				attribute_value_integerFormCallback.probe.backRepoOfInterest,
				attribute_value_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
					pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](attribute_value_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
								pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
								newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
							}
						} else {
							newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
						}
					}
				}
			}
		case "SPEC_OBJECT:VALUES.ATTRIBUTE_VALUE_INTEGER":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_integerFormCallback.probe.stageOfInterest,
				attribute_value_integerFormCallback.probe.backRepoOfInterest,
				attribute_value_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
					pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](attribute_value_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
								pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
								newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
							}
						} else {
							newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
						}
					}
				}
			}
		case "SPEC_RELATION:VALUES.ATTRIBUTE_VALUE_INTEGER":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_integerFormCallback.probe.stageOfInterest,
				attribute_value_integerFormCallback.probe.backRepoOfInterest,
				attribute_value_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
					pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](attribute_value_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
								pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
								newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
							}
						} else {
							newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_value_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_integer_.Unstage(attribute_value_integerFormCallback.probe.stageOfInterest)
	}

	attribute_value_integerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_INTEGER](
		attribute_value_integerFormCallback.probe,
	)
	attribute_value_integerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_integerFormCallback.CreationMode || attribute_value_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_integerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_integerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
			nil,
			attribute_value_integerFormCallback.probe,
			newFormGroup,
		)
		attribute_value_integer := new(models.ATTRIBUTE_VALUE_INTEGER)
		FillUpForm(attribute_value_integer, newFormGroup, attribute_value_integerFormCallback.probe)
		attribute_value_integerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_integerFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
	attribute_value_real *models.ATTRIBUTE_VALUE_REAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_realFormCallback *ATTRIBUTE_VALUE_REALFormCallback) {
	attribute_value_realFormCallback = new(ATTRIBUTE_VALUE_REALFormCallback)
	attribute_value_realFormCallback.probe = probe
	attribute_value_realFormCallback.attribute_value_real = attribute_value_real
	attribute_value_realFormCallback.formGroup = formGroup

	attribute_value_realFormCallback.CreationMode = (attribute_value_real == nil)

	return
}

type ATTRIBUTE_VALUE_REALFormCallback struct {
	attribute_value_real *models.ATTRIBUTE_VALUE_REAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_realFormCallback *ATTRIBUTE_VALUE_REALFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_REALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_realFormCallback.probe.formStage.Checkout()

	if attribute_value_realFormCallback.attribute_value_real == nil {
		attribute_value_realFormCallback.attribute_value_real = new(models.ATTRIBUTE_VALUE_REAL).Stage(attribute_value_realFormCallback.probe.stageOfInterest)
	}
	attribute_value_real_ := attribute_value_realFormCallback.attribute_value_real
	_ = attribute_value_real_

	for _, formDiv := range attribute_value_realFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_real_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_real_.THE_VALUE), formDiv)
		case "ATTRIBUTE_DEFINITION_REAL:DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_REALOwner *models.ATTRIBUTE_DEFINITION_REAL
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_realFormCallback.probe.stageOfInterest,
				attribute_value_realFormCallback.probe.backRepoOfInterest,
				attribute_value_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_REALOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_REAL)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_REALOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
					pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL = slices.Delete(pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_real := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_REAL](attribute_value_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_real.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_REALOwner := _attribute_definition_real // we have a match
						if pastATTRIBUTE_DEFINITION_REALOwner != nil {
							if newATTRIBUTE_DEFINITION_REALOwner != pastATTRIBUTE_DEFINITION_REALOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
								pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL = slices.Delete(pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL, idx, idx+1)
								newATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL = append(newATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
							}
						} else {
							newATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL = append(newATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
						}
					}
				}
			}
		case "SPECIFICATION:VALUES.ATTRIBUTE_VALUE_REAL":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_realFormCallback.probe.stageOfInterest,
				attribute_value_realFormCallback.probe.backRepoOfInterest,
				attribute_value_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
					pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](attribute_value_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
								pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, idx, idx+1)
								newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
							}
						} else {
							newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
						}
					}
				}
			}
		case "SPEC_OBJECT:VALUES.ATTRIBUTE_VALUE_REAL":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_realFormCallback.probe.stageOfInterest,
				attribute_value_realFormCallback.probe.backRepoOfInterest,
				attribute_value_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
					pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_REAL = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](attribute_value_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
								pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_REAL = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_REAL, idx, idx+1)
								newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_REAL = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
							}
						} else {
							newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_REAL = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
						}
					}
				}
			}
		case "SPEC_RELATION:VALUES.ATTRIBUTE_VALUE_REAL":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_realFormCallback.probe.stageOfInterest,
				attribute_value_realFormCallback.probe.backRepoOfInterest,
				attribute_value_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
					pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](attribute_value_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
								pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, idx, idx+1)
								newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
							}
						} else {
							newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_value_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_real_.Unstage(attribute_value_realFormCallback.probe.stageOfInterest)
	}

	attribute_value_realFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_REAL](
		attribute_value_realFormCallback.probe,
	)
	attribute_value_realFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_realFormCallback.CreationMode || attribute_value_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_realFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_realFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
			nil,
			attribute_value_realFormCallback.probe,
			newFormGroup,
		)
		attribute_value_real := new(models.ATTRIBUTE_VALUE_REAL)
		FillUpForm(attribute_value_real, newFormGroup, attribute_value_realFormCallback.probe)
		attribute_value_realFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_realFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
	attribute_value_string *models.ATTRIBUTE_VALUE_STRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_stringFormCallback *ATTRIBUTE_VALUE_STRINGFormCallback) {
	attribute_value_stringFormCallback = new(ATTRIBUTE_VALUE_STRINGFormCallback)
	attribute_value_stringFormCallback.probe = probe
	attribute_value_stringFormCallback.attribute_value_string = attribute_value_string
	attribute_value_stringFormCallback.formGroup = formGroup

	attribute_value_stringFormCallback.CreationMode = (attribute_value_string == nil)

	return
}

type ATTRIBUTE_VALUE_STRINGFormCallback struct {
	attribute_value_string *models.ATTRIBUTE_VALUE_STRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_stringFormCallback *ATTRIBUTE_VALUE_STRINGFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_STRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_stringFormCallback.probe.formStage.Checkout()

	if attribute_value_stringFormCallback.attribute_value_string == nil {
		attribute_value_stringFormCallback.attribute_value_string = new(models.ATTRIBUTE_VALUE_STRING).Stage(attribute_value_stringFormCallback.probe.stageOfInterest)
	}
	attribute_value_string_ := attribute_value_stringFormCallback.attribute_value_string
	_ = attribute_value_string_

	for _, formDiv := range attribute_value_stringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_string_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_string_.THE_VALUE), formDiv)
		case "ATTRIBUTE_DEFINITION_STRING:DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_STRINGOwner *models.ATTRIBUTE_DEFINITION_STRING
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_stringFormCallback.probe.stageOfInterest,
				attribute_value_stringFormCallback.probe.backRepoOfInterest,
				attribute_value_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_STRINGOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_STRING)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_STRINGOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
					pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING = slices.Delete(pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_string := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_STRING](attribute_value_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_string.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_STRINGOwner := _attribute_definition_string // we have a match
						if pastATTRIBUTE_DEFINITION_STRINGOwner != nil {
							if newATTRIBUTE_DEFINITION_STRINGOwner != pastATTRIBUTE_DEFINITION_STRINGOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
								pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING = slices.Delete(pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, idx, idx+1)
								newATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING = append(newATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
							}
						} else {
							newATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING = append(newATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
						}
					}
				}
			}
		case "SPECIFICATION:VALUES.ATTRIBUTE_VALUE_STRING":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_stringFormCallback.probe.stageOfInterest,
				attribute_value_stringFormCallback.probe.backRepoOfInterest,
				attribute_value_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
					pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](attribute_value_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
								pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, idx, idx+1)
								newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
							}
						} else {
							newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
						}
					}
				}
			}
		case "SPEC_OBJECT:VALUES.ATTRIBUTE_VALUE_STRING":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_stringFormCallback.probe.stageOfInterest,
				attribute_value_stringFormCallback.probe.backRepoOfInterest,
				attribute_value_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
					pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_STRING = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](attribute_value_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
								pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_STRING = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_STRING, idx, idx+1)
								newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_STRING = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
							}
						} else {
							newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_STRING = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
						}
					}
				}
			}
		case "SPEC_RELATION:VALUES.ATTRIBUTE_VALUE_STRING":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_stringFormCallback.probe.stageOfInterest,
				attribute_value_stringFormCallback.probe.backRepoOfInterest,
				attribute_value_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
					pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](attribute_value_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
								pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, idx, idx+1)
								newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
							}
						} else {
							newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_value_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_string_.Unstage(attribute_value_stringFormCallback.probe.stageOfInterest)
	}

	attribute_value_stringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_STRING](
		attribute_value_stringFormCallback.probe,
	)
	attribute_value_stringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_stringFormCallback.CreationMode || attribute_value_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_stringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_stringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
			nil,
			attribute_value_stringFormCallback.probe,
			newFormGroup,
		)
		attribute_value_string := new(models.ATTRIBUTE_VALUE_STRING)
		FillUpForm(attribute_value_string, newFormGroup, attribute_value_stringFormCallback.probe)
		attribute_value_stringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_stringFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
	attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_xhtmlFormCallback *ATTRIBUTE_VALUE_XHTMLFormCallback) {
	attribute_value_xhtmlFormCallback = new(ATTRIBUTE_VALUE_XHTMLFormCallback)
	attribute_value_xhtmlFormCallback.probe = probe
	attribute_value_xhtmlFormCallback.attribute_value_xhtml = attribute_value_xhtml
	attribute_value_xhtmlFormCallback.formGroup = formGroup

	attribute_value_xhtmlFormCallback.CreationMode = (attribute_value_xhtml == nil)

	return
}

type ATTRIBUTE_VALUE_XHTMLFormCallback struct {
	attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_xhtmlFormCallback *ATTRIBUTE_VALUE_XHTMLFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_XHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_xhtmlFormCallback.probe.formStage.Checkout()

	if attribute_value_xhtmlFormCallback.attribute_value_xhtml == nil {
		attribute_value_xhtmlFormCallback.attribute_value_xhtml = new(models.ATTRIBUTE_VALUE_XHTML).Stage(attribute_value_xhtmlFormCallback.probe.stageOfInterest)
	}
	attribute_value_xhtml_ := attribute_value_xhtmlFormCallback.attribute_value_xhtml
	_ = attribute_value_xhtml_

	for _, formDiv := range attribute_value_xhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_xhtml_.Name), formDiv)
		case "IS_SIMPLIFIED":
			FormDivBasicFieldToField(&(attribute_value_xhtml_.IS_SIMPLIFIED), formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML:DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_XHTMLOwner *models.ATTRIBUTE_DEFINITION_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
			rf.Fieldname = "DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_xhtmlFormCallback.probe.stageOfInterest,
				attribute_value_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_value_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_XHTMLOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_XHTMLOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
					pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_xhtml := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_XHTML](attribute_value_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_XHTMLOwner := _attribute_definition_xhtml // we have a match
						if pastATTRIBUTE_DEFINITION_XHTMLOwner != nil {
							if newATTRIBUTE_DEFINITION_XHTMLOwner != pastATTRIBUTE_DEFINITION_XHTMLOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
								pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
								newATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML = append(newATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
							}
						} else {
							newATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML = append(newATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
						}
					}
				}
			}
		case "SPECIFICATION:VALUES.ATTRIBUTE_VALUE_XHTML":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_xhtmlFormCallback.probe.stageOfInterest,
				attribute_value_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_value_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
					pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](attribute_value_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
								pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
								newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
							}
						} else {
							newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML = append(newSPECIFICATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
						}
					}
				}
			}
		case "SPEC_OBJECT:VALUES.ATTRIBUTE_VALUE_XHTML":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_xhtmlFormCallback.probe.stageOfInterest,
				attribute_value_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_value_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
					pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](attribute_value_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
								pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
								newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_XHTML = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
							}
						} else {
							newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_XHTML = append(newSPEC_OBJECTOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
						}
					}
				}
			}
		case "SPEC_RELATION:VALUES.ATTRIBUTE_VALUE_XHTML":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES.ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_xhtmlFormCallback.probe.stageOfInterest,
				attribute_value_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_value_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
					pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](attribute_value_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
								pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
								newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
							}
						} else {
							newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML = append(newSPEC_RELATIONOwner.VALUES.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attribute_value_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_xhtml_.Unstage(attribute_value_xhtmlFormCallback.probe.stageOfInterest)
	}

	attribute_value_xhtmlFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_XHTML](
		attribute_value_xhtmlFormCallback.probe,
	)
	attribute_value_xhtmlFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_xhtmlFormCallback.CreationMode || attribute_value_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_xhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_xhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
			nil,
			attribute_value_xhtmlFormCallback.probe,
			newFormGroup,
		)
		attribute_value_xhtml := new(models.ATTRIBUTE_VALUE_XHTML)
		FillUpForm(attribute_value_xhtml, newFormGroup, attribute_value_xhtmlFormCallback.probe)
		attribute_value_xhtmlFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_xhtmlFormCallback.probe)
}
func __gong__New__AnyTypeFormCallback(
	anytype *models.AnyType,
	probe *Probe,
	formGroup *table.FormGroup,
) (anytypeFormCallback *AnyTypeFormCallback) {
	anytypeFormCallback = new(AnyTypeFormCallback)
	anytypeFormCallback.probe = probe
	anytypeFormCallback.anytype = anytype
	anytypeFormCallback.formGroup = formGroup

	anytypeFormCallback.CreationMode = (anytype == nil)

	return
}

type AnyTypeFormCallback struct {
	anytype *models.AnyType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (anytypeFormCallback *AnyTypeFormCallback) OnSave() {

	log.Println("AnyTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	anytypeFormCallback.probe.formStage.Checkout()

	if anytypeFormCallback.anytype == nil {
		anytypeFormCallback.anytype = new(models.AnyType).Stage(anytypeFormCallback.probe.stageOfInterest)
	}
	anytype_ := anytypeFormCallback.anytype
	_ = anytype_

	for _, formDiv := range anytypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(anytype_.Name), formDiv)
		case "InnerXML":
			FormDivBasicFieldToField(&(anytype_.InnerXML), formDiv)
		}
	}

	// manage the suppress operation
	if anytypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		anytype_.Unstage(anytypeFormCallback.probe.stageOfInterest)
	}

	anytypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AnyType](
		anytypeFormCallback.probe,
	)
	anytypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if anytypeFormCallback.CreationMode || anytypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		anytypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(anytypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AnyTypeFormCallback(
			nil,
			anytypeFormCallback.probe,
			newFormGroup,
		)
		anytype := new(models.AnyType)
		FillUpForm(anytype, newFormGroup, anytypeFormCallback.probe)
		anytypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(anytypeFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
	datatype_definition_boolean *models.DATATYPE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_booleanFormCallback *DATATYPE_DEFINITION_BOOLEANFormCallback) {
	datatype_definition_booleanFormCallback = new(DATATYPE_DEFINITION_BOOLEANFormCallback)
	datatype_definition_booleanFormCallback.probe = probe
	datatype_definition_booleanFormCallback.datatype_definition_boolean = datatype_definition_boolean
	datatype_definition_booleanFormCallback.formGroup = formGroup

	datatype_definition_booleanFormCallback.CreationMode = (datatype_definition_boolean == nil)

	return
}

type DATATYPE_DEFINITION_BOOLEANFormCallback struct {
	datatype_definition_boolean *models.DATATYPE_DEFINITION_BOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_booleanFormCallback *DATATYPE_DEFINITION_BOOLEANFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_BOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_booleanFormCallback.probe.formStage.Checkout()

	if datatype_definition_booleanFormCallback.datatype_definition_boolean == nil {
		datatype_definition_booleanFormCallback.datatype_definition_boolean = new(models.DATATYPE_DEFINITION_BOOLEAN).Stage(datatype_definition_booleanFormCallback.probe.stageOfInterest)
	}
	datatype_definition_boolean_ := datatype_definition_booleanFormCallback.datatype_definition_boolean
	_ = datatype_definition_boolean_

	for _, formDiv := range datatype_definition_booleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:DATATYPES.DATATYPE_DEFINITION_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_booleanFormCallback.probe.stageOfInterest,
				datatype_definition_booleanFormCallback.probe.backRepoOfInterest,
				datatype_definition_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_BOOLEAN, datatype_definition_boolean_)
					pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_BOOLEAN = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](datatype_definition_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_BOOLEAN, datatype_definition_boolean_)
								pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_BOOLEAN = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_BOOLEAN, idx, idx+1)
								newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_BOOLEAN = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_BOOLEAN, datatype_definition_boolean_)
							}
						} else {
							newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_BOOLEAN = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_BOOLEAN, datatype_definition_boolean_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatype_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_boolean_.Unstage(datatype_definition_booleanFormCallback.probe.stageOfInterest)
	}

	datatype_definition_booleanFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_BOOLEAN](
		datatype_definition_booleanFormCallback.probe,
	)
	datatype_definition_booleanFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_booleanFormCallback.CreationMode || datatype_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_booleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_booleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
			nil,
			datatype_definition_booleanFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_boolean := new(models.DATATYPE_DEFINITION_BOOLEAN)
		FillUpForm(datatype_definition_boolean, newFormGroup, datatype_definition_booleanFormCallback.probe)
		datatype_definition_booleanFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_booleanFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
	datatype_definition_date *models.DATATYPE_DEFINITION_DATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_dateFormCallback *DATATYPE_DEFINITION_DATEFormCallback) {
	datatype_definition_dateFormCallback = new(DATATYPE_DEFINITION_DATEFormCallback)
	datatype_definition_dateFormCallback.probe = probe
	datatype_definition_dateFormCallback.datatype_definition_date = datatype_definition_date
	datatype_definition_dateFormCallback.formGroup = formGroup

	datatype_definition_dateFormCallback.CreationMode = (datatype_definition_date == nil)

	return
}

type DATATYPE_DEFINITION_DATEFormCallback struct {
	datatype_definition_date *models.DATATYPE_DEFINITION_DATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_dateFormCallback *DATATYPE_DEFINITION_DATEFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_DATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_dateFormCallback.probe.formStage.Checkout()

	if datatype_definition_dateFormCallback.datatype_definition_date == nil {
		datatype_definition_dateFormCallback.datatype_definition_date = new(models.DATATYPE_DEFINITION_DATE).Stage(datatype_definition_dateFormCallback.probe.stageOfInterest)
	}
	datatype_definition_date_ := datatype_definition_dateFormCallback.datatype_definition_date
	_ = datatype_definition_date_

	for _, formDiv := range datatype_definition_dateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_date_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_date_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_date_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_date_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:DATATYPES.DATATYPE_DEFINITION_DATE":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_dateFormCallback.probe.stageOfInterest,
				datatype_definition_dateFormCallback.probe.backRepoOfInterest,
				datatype_definition_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_DATE, datatype_definition_date_)
					pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_DATE = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](datatype_definition_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_DATE, datatype_definition_date_)
								pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_DATE = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_DATE, idx, idx+1)
								newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_DATE = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_DATE, datatype_definition_date_)
							}
						} else {
							newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_DATE = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_DATE, datatype_definition_date_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatype_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_date_.Unstage(datatype_definition_dateFormCallback.probe.stageOfInterest)
	}

	datatype_definition_dateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_DATE](
		datatype_definition_dateFormCallback.probe,
	)
	datatype_definition_dateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_dateFormCallback.CreationMode || datatype_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_dateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_dateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
			nil,
			datatype_definition_dateFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_date := new(models.DATATYPE_DEFINITION_DATE)
		FillUpForm(datatype_definition_date, newFormGroup, datatype_definition_dateFormCallback.probe)
		datatype_definition_dateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_dateFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
	datatype_definition_enumeration *models.DATATYPE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_enumerationFormCallback *DATATYPE_DEFINITION_ENUMERATIONFormCallback) {
	datatype_definition_enumerationFormCallback = new(DATATYPE_DEFINITION_ENUMERATIONFormCallback)
	datatype_definition_enumerationFormCallback.probe = probe
	datatype_definition_enumerationFormCallback.datatype_definition_enumeration = datatype_definition_enumeration
	datatype_definition_enumerationFormCallback.formGroup = formGroup

	datatype_definition_enumerationFormCallback.CreationMode = (datatype_definition_enumeration == nil)

	return
}

type DATATYPE_DEFINITION_ENUMERATIONFormCallback struct {
	datatype_definition_enumeration *models.DATATYPE_DEFINITION_ENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_enumerationFormCallback *DATATYPE_DEFINITION_ENUMERATIONFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_ENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_enumerationFormCallback.probe.formStage.Checkout()

	if datatype_definition_enumerationFormCallback.datatype_definition_enumeration == nil {
		datatype_definition_enumerationFormCallback.datatype_definition_enumeration = new(models.DATATYPE_DEFINITION_ENUMERATION).Stage(datatype_definition_enumerationFormCallback.probe.stageOfInterest)
	}
	datatype_definition_enumeration_ := datatype_definition_enumerationFormCallback.datatype_definition_enumeration
	_ = datatype_definition_enumeration_

	for _, formDiv := range datatype_definition_enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:DATATYPES.DATATYPE_DEFINITION_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_enumerationFormCallback.probe.stageOfInterest,
				datatype_definition_enumerationFormCallback.probe.backRepoOfInterest,
				datatype_definition_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_ENUMERATION, datatype_definition_enumeration_)
					pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_ENUMERATION = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](datatype_definition_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_ENUMERATION, datatype_definition_enumeration_)
								pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_ENUMERATION = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_ENUMERATION, idx, idx+1)
								newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_ENUMERATION = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_ENUMERATION, datatype_definition_enumeration_)
							}
						} else {
							newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_ENUMERATION = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_ENUMERATION, datatype_definition_enumeration_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatype_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_enumeration_.Unstage(datatype_definition_enumerationFormCallback.probe.stageOfInterest)
	}

	datatype_definition_enumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_ENUMERATION](
		datatype_definition_enumerationFormCallback.probe,
	)
	datatype_definition_enumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_enumerationFormCallback.CreationMode || datatype_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
			nil,
			datatype_definition_enumerationFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_enumeration := new(models.DATATYPE_DEFINITION_ENUMERATION)
		FillUpForm(datatype_definition_enumeration, newFormGroup, datatype_definition_enumerationFormCallback.probe)
		datatype_definition_enumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_enumerationFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
	datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_integerFormCallback *DATATYPE_DEFINITION_INTEGERFormCallback) {
	datatype_definition_integerFormCallback = new(DATATYPE_DEFINITION_INTEGERFormCallback)
	datatype_definition_integerFormCallback.probe = probe
	datatype_definition_integerFormCallback.datatype_definition_integer = datatype_definition_integer
	datatype_definition_integerFormCallback.formGroup = formGroup

	datatype_definition_integerFormCallback.CreationMode = (datatype_definition_integer == nil)

	return
}

type DATATYPE_DEFINITION_INTEGERFormCallback struct {
	datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_integerFormCallback *DATATYPE_DEFINITION_INTEGERFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_INTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_integerFormCallback.probe.formStage.Checkout()

	if datatype_definition_integerFormCallback.datatype_definition_integer == nil {
		datatype_definition_integerFormCallback.datatype_definition_integer = new(models.DATATYPE_DEFINITION_INTEGER).Stage(datatype_definition_integerFormCallback.probe.stageOfInterest)
	}
	datatype_definition_integer_ := datatype_definition_integerFormCallback.datatype_definition_integer
	_ = datatype_definition_integer_

	for _, formDiv := range datatype_definition_integerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_integer_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_integer_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_integer_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_integer_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:DATATYPES.DATATYPE_DEFINITION_INTEGER":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_integerFormCallback.probe.stageOfInterest,
				datatype_definition_integerFormCallback.probe.backRepoOfInterest,
				datatype_definition_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_INTEGER, datatype_definition_integer_)
					pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_INTEGER = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](datatype_definition_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_INTEGER, datatype_definition_integer_)
								pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_INTEGER = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_INTEGER, idx, idx+1)
								newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_INTEGER = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_INTEGER, datatype_definition_integer_)
							}
						} else {
							newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_INTEGER = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_INTEGER, datatype_definition_integer_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatype_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_integer_.Unstage(datatype_definition_integerFormCallback.probe.stageOfInterest)
	}

	datatype_definition_integerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_INTEGER](
		datatype_definition_integerFormCallback.probe,
	)
	datatype_definition_integerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_integerFormCallback.CreationMode || datatype_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_integerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_integerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
			nil,
			datatype_definition_integerFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_integer := new(models.DATATYPE_DEFINITION_INTEGER)
		FillUpForm(datatype_definition_integer, newFormGroup, datatype_definition_integerFormCallback.probe)
		datatype_definition_integerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_integerFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_REALFormCallback(
	datatype_definition_real *models.DATATYPE_DEFINITION_REAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_realFormCallback *DATATYPE_DEFINITION_REALFormCallback) {
	datatype_definition_realFormCallback = new(DATATYPE_DEFINITION_REALFormCallback)
	datatype_definition_realFormCallback.probe = probe
	datatype_definition_realFormCallback.datatype_definition_real = datatype_definition_real
	datatype_definition_realFormCallback.formGroup = formGroup

	datatype_definition_realFormCallback.CreationMode = (datatype_definition_real == nil)

	return
}

type DATATYPE_DEFINITION_REALFormCallback struct {
	datatype_definition_real *models.DATATYPE_DEFINITION_REAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_realFormCallback *DATATYPE_DEFINITION_REALFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_REALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_realFormCallback.probe.formStage.Checkout()

	if datatype_definition_realFormCallback.datatype_definition_real == nil {
		datatype_definition_realFormCallback.datatype_definition_real = new(models.DATATYPE_DEFINITION_REAL).Stage(datatype_definition_realFormCallback.probe.stageOfInterest)
	}
	datatype_definition_real_ := datatype_definition_realFormCallback.datatype_definition_real
	_ = datatype_definition_real_

	for _, formDiv := range datatype_definition_realFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_real_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_real_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_real_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_real_.LONG_NAME), formDiv)
		case "MAX":
			FormDivBasicFieldToField(&(datatype_definition_real_.MAX), formDiv)
		case "MIN":
			FormDivBasicFieldToField(&(datatype_definition_real_.MIN), formDiv)
		case "REQ_IF_CONTENT:DATATYPES.DATATYPE_DEFINITION_REAL":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_realFormCallback.probe.stageOfInterest,
				datatype_definition_realFormCallback.probe.backRepoOfInterest,
				datatype_definition_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_REAL, datatype_definition_real_)
					pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_REAL = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](datatype_definition_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_REAL, datatype_definition_real_)
								pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_REAL = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_REAL, idx, idx+1)
								newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_REAL = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_REAL, datatype_definition_real_)
							}
						} else {
							newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_REAL = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_REAL, datatype_definition_real_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatype_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_real_.Unstage(datatype_definition_realFormCallback.probe.stageOfInterest)
	}

	datatype_definition_realFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_REAL](
		datatype_definition_realFormCallback.probe,
	)
	datatype_definition_realFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_realFormCallback.CreationMode || datatype_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_realFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_realFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_REALFormCallback(
			nil,
			datatype_definition_realFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_real := new(models.DATATYPE_DEFINITION_REAL)
		FillUpForm(datatype_definition_real, newFormGroup, datatype_definition_realFormCallback.probe)
		datatype_definition_realFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_realFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
	datatype_definition_string *models.DATATYPE_DEFINITION_STRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_stringFormCallback *DATATYPE_DEFINITION_STRINGFormCallback) {
	datatype_definition_stringFormCallback = new(DATATYPE_DEFINITION_STRINGFormCallback)
	datatype_definition_stringFormCallback.probe = probe
	datatype_definition_stringFormCallback.datatype_definition_string = datatype_definition_string
	datatype_definition_stringFormCallback.formGroup = formGroup

	datatype_definition_stringFormCallback.CreationMode = (datatype_definition_string == nil)

	return
}

type DATATYPE_DEFINITION_STRINGFormCallback struct {
	datatype_definition_string *models.DATATYPE_DEFINITION_STRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_stringFormCallback *DATATYPE_DEFINITION_STRINGFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_STRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_stringFormCallback.probe.formStage.Checkout()

	if datatype_definition_stringFormCallback.datatype_definition_string == nil {
		datatype_definition_stringFormCallback.datatype_definition_string = new(models.DATATYPE_DEFINITION_STRING).Stage(datatype_definition_stringFormCallback.probe.stageOfInterest)
	}
	datatype_definition_string_ := datatype_definition_stringFormCallback.datatype_definition_string
	_ = datatype_definition_string_

	for _, formDiv := range datatype_definition_stringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_string_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_string_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_string_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_string_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:DATATYPES.DATATYPE_DEFINITION_STRING":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_stringFormCallback.probe.stageOfInterest,
				datatype_definition_stringFormCallback.probe.backRepoOfInterest,
				datatype_definition_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_STRING, datatype_definition_string_)
					pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_STRING = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](datatype_definition_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_STRING, datatype_definition_string_)
								pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_STRING = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_STRING, idx, idx+1)
								newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_STRING = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_STRING, datatype_definition_string_)
							}
						} else {
							newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_STRING = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_STRING, datatype_definition_string_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatype_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_string_.Unstage(datatype_definition_stringFormCallback.probe.stageOfInterest)
	}

	datatype_definition_stringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_STRING](
		datatype_definition_stringFormCallback.probe,
	)
	datatype_definition_stringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_stringFormCallback.CreationMode || datatype_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_stringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_stringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
			nil,
			datatype_definition_stringFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_string := new(models.DATATYPE_DEFINITION_STRING)
		FillUpForm(datatype_definition_string, newFormGroup, datatype_definition_stringFormCallback.probe)
		datatype_definition_stringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_stringFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
	datatype_definition_xhtml *models.DATATYPE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_xhtmlFormCallback *DATATYPE_DEFINITION_XHTMLFormCallback) {
	datatype_definition_xhtmlFormCallback = new(DATATYPE_DEFINITION_XHTMLFormCallback)
	datatype_definition_xhtmlFormCallback.probe = probe
	datatype_definition_xhtmlFormCallback.datatype_definition_xhtml = datatype_definition_xhtml
	datatype_definition_xhtmlFormCallback.formGroup = formGroup

	datatype_definition_xhtmlFormCallback.CreationMode = (datatype_definition_xhtml == nil)

	return
}

type DATATYPE_DEFINITION_XHTMLFormCallback struct {
	datatype_definition_xhtml *models.DATATYPE_DEFINITION_XHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_xhtmlFormCallback *DATATYPE_DEFINITION_XHTMLFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_XHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_xhtmlFormCallback.probe.formStage.Checkout()

	if datatype_definition_xhtmlFormCallback.datatype_definition_xhtml == nil {
		datatype_definition_xhtmlFormCallback.datatype_definition_xhtml = new(models.DATATYPE_DEFINITION_XHTML).Stage(datatype_definition_xhtmlFormCallback.probe.stageOfInterest)
	}
	datatype_definition_xhtml_ := datatype_definition_xhtmlFormCallback.datatype_definition_xhtml
	_ = datatype_definition_xhtml_

	for _, formDiv := range datatype_definition_xhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:DATATYPES.DATATYPE_DEFINITION_XHTML":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES.DATATYPE_DEFINITION_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_xhtmlFormCallback.probe.stageOfInterest,
				datatype_definition_xhtmlFormCallback.probe.backRepoOfInterest,
				datatype_definition_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_XHTML, datatype_definition_xhtml_)
					pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_XHTML = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](datatype_definition_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_XHTML, datatype_definition_xhtml_)
								pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_XHTML = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_XHTML, idx, idx+1)
								newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_XHTML = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_XHTML, datatype_definition_xhtml_)
							}
						} else {
							newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_XHTML = append(newREQ_IF_CONTENTOwner.DATATYPES.DATATYPE_DEFINITION_XHTML, datatype_definition_xhtml_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatype_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_xhtml_.Unstage(datatype_definition_xhtmlFormCallback.probe.stageOfInterest)
	}

	datatype_definition_xhtmlFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_XHTML](
		datatype_definition_xhtmlFormCallback.probe,
	)
	datatype_definition_xhtmlFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_xhtmlFormCallback.CreationMode || datatype_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_xhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_xhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
			nil,
			datatype_definition_xhtmlFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_xhtml := new(models.DATATYPE_DEFINITION_XHTML)
		FillUpForm(datatype_definition_xhtml, newFormGroup, datatype_definition_xhtmlFormCallback.probe)
		datatype_definition_xhtmlFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_xhtmlFormCallback.probe)
}
func __gong__New__EMBEDDED_VALUEFormCallback(
	embedded_value *models.EMBEDDED_VALUE,
	probe *Probe,
	formGroup *table.FormGroup,
) (embedded_valueFormCallback *EMBEDDED_VALUEFormCallback) {
	embedded_valueFormCallback = new(EMBEDDED_VALUEFormCallback)
	embedded_valueFormCallback.probe = probe
	embedded_valueFormCallback.embedded_value = embedded_value
	embedded_valueFormCallback.formGroup = formGroup

	embedded_valueFormCallback.CreationMode = (embedded_value == nil)

	return
}

type EMBEDDED_VALUEFormCallback struct {
	embedded_value *models.EMBEDDED_VALUE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (embedded_valueFormCallback *EMBEDDED_VALUEFormCallback) OnSave() {

	log.Println("EMBEDDED_VALUEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	embedded_valueFormCallback.probe.formStage.Checkout()

	if embedded_valueFormCallback.embedded_value == nil {
		embedded_valueFormCallback.embedded_value = new(models.EMBEDDED_VALUE).Stage(embedded_valueFormCallback.probe.stageOfInterest)
	}
	embedded_value_ := embedded_valueFormCallback.embedded_value
	_ = embedded_value_

	for _, formDiv := range embedded_valueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(embedded_value_.Name), formDiv)
		case "OTHER_CONTENT":
			FormDivBasicFieldToField(&(embedded_value_.OTHER_CONTENT), formDiv)
		case "ENUM_VALUE:PROPERTIES.EMBEDDED_VALUE":
			// we need to retrieve the field owner before the change
			var pastENUM_VALUEOwner *models.ENUM_VALUE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ENUM_VALUE"
			rf.Fieldname = "PROPERTIES.EMBEDDED_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				embedded_valueFormCallback.probe.stageOfInterest,
				embedded_valueFormCallback.probe.backRepoOfInterest,
				embedded_value_,
				&rf)

			if reverseFieldOwner != nil {
				pastENUM_VALUEOwner = reverseFieldOwner.(*models.ENUM_VALUE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastENUM_VALUEOwner != nil {
					idx := slices.Index(pastENUM_VALUEOwner.PROPERTIES.EMBEDDED_VALUE, embedded_value_)
					pastENUM_VALUEOwner.PROPERTIES.EMBEDDED_VALUE = slices.Delete(pastENUM_VALUEOwner.PROPERTIES.EMBEDDED_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _enum_value := range *models.GetGongstructInstancesSet[models.ENUM_VALUE](embedded_valueFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _enum_value.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newENUM_VALUEOwner := _enum_value // we have a match
						if pastENUM_VALUEOwner != nil {
							if newENUM_VALUEOwner != pastENUM_VALUEOwner {
								idx := slices.Index(pastENUM_VALUEOwner.PROPERTIES.EMBEDDED_VALUE, embedded_value_)
								pastENUM_VALUEOwner.PROPERTIES.EMBEDDED_VALUE = slices.Delete(pastENUM_VALUEOwner.PROPERTIES.EMBEDDED_VALUE, idx, idx+1)
								newENUM_VALUEOwner.PROPERTIES.EMBEDDED_VALUE = append(newENUM_VALUEOwner.PROPERTIES.EMBEDDED_VALUE, embedded_value_)
							}
						} else {
							newENUM_VALUEOwner.PROPERTIES.EMBEDDED_VALUE = append(newENUM_VALUEOwner.PROPERTIES.EMBEDDED_VALUE, embedded_value_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if embedded_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		embedded_value_.Unstage(embedded_valueFormCallback.probe.stageOfInterest)
	}

	embedded_valueFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.EMBEDDED_VALUE](
		embedded_valueFormCallback.probe,
	)
	embedded_valueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if embedded_valueFormCallback.CreationMode || embedded_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		embedded_valueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(embedded_valueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EMBEDDED_VALUEFormCallback(
			nil,
			embedded_valueFormCallback.probe,
			newFormGroup,
		)
		embedded_value := new(models.EMBEDDED_VALUE)
		FillUpForm(embedded_value, newFormGroup, embedded_valueFormCallback.probe)
		embedded_valueFormCallback.probe.formStage.Commit()
	}

	fillUpTree(embedded_valueFormCallback.probe)
}
func __gong__New__ENUM_VALUEFormCallback(
	enum_value *models.ENUM_VALUE,
	probe *Probe,
	formGroup *table.FormGroup,
) (enum_valueFormCallback *ENUM_VALUEFormCallback) {
	enum_valueFormCallback = new(ENUM_VALUEFormCallback)
	enum_valueFormCallback.probe = probe
	enum_valueFormCallback.enum_value = enum_value
	enum_valueFormCallback.formGroup = formGroup

	enum_valueFormCallback.CreationMode = (enum_value == nil)

	return
}

type ENUM_VALUEFormCallback struct {
	enum_value *models.ENUM_VALUE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (enum_valueFormCallback *ENUM_VALUEFormCallback) OnSave() {

	log.Println("ENUM_VALUEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	enum_valueFormCallback.probe.formStage.Checkout()

	if enum_valueFormCallback.enum_value == nil {
		enum_valueFormCallback.enum_value = new(models.ENUM_VALUE).Stage(enum_valueFormCallback.probe.stageOfInterest)
	}
	enum_value_ := enum_valueFormCallback.enum_value
	_ = enum_value_

	for _, formDiv := range enum_valueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(enum_value_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(enum_value_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(enum_value_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(enum_value_.LONG_NAME), formDiv)
		case "DATATYPE_DEFINITION_ENUMERATION:SPECIFIED_VALUES.ENUM_VALUE":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_ENUMERATIONOwner *models.DATATYPE_DEFINITION_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_ENUMERATION"
			rf.Fieldname = "SPECIFIED_VALUES.ENUM_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				enum_valueFormCallback.probe.stageOfInterest,
				enum_valueFormCallback.probe.backRepoOfInterest,
				enum_value_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_ENUMERATIONOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_ENUMERATIONOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES.ENUM_VALUE, enum_value_)
					pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES.ENUM_VALUE = slices.Delete(pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES.ENUM_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_enumeration := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_ENUMERATION](enum_valueFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_ENUMERATIONOwner := _datatype_definition_enumeration // we have a match
						if pastDATATYPE_DEFINITION_ENUMERATIONOwner != nil {
							if newDATATYPE_DEFINITION_ENUMERATIONOwner != pastDATATYPE_DEFINITION_ENUMERATIONOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES.ENUM_VALUE, enum_value_)
								pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES.ENUM_VALUE = slices.Delete(pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES.ENUM_VALUE, idx, idx+1)
								newDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES.ENUM_VALUE = append(newDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES.ENUM_VALUE, enum_value_)
							}
						} else {
							newDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES.ENUM_VALUE = append(newDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES.ENUM_VALUE, enum_value_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if enum_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enum_value_.Unstage(enum_valueFormCallback.probe.stageOfInterest)
	}

	enum_valueFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ENUM_VALUE](
		enum_valueFormCallback.probe,
	)
	enum_valueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if enum_valueFormCallback.CreationMode || enum_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enum_valueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(enum_valueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ENUM_VALUEFormCallback(
			nil,
			enum_valueFormCallback.probe,
			newFormGroup,
		)
		enum_value := new(models.ENUM_VALUE)
		FillUpForm(enum_value, newFormGroup, enum_valueFormCallback.probe)
		enum_valueFormCallback.probe.formStage.Commit()
	}

	fillUpTree(enum_valueFormCallback.probe)
}
func __gong__New__RELATION_GROUPFormCallback(
	relation_group *models.RELATION_GROUP,
	probe *Probe,
	formGroup *table.FormGroup,
) (relation_groupFormCallback *RELATION_GROUPFormCallback) {
	relation_groupFormCallback = new(RELATION_GROUPFormCallback)
	relation_groupFormCallback.probe = probe
	relation_groupFormCallback.relation_group = relation_group
	relation_groupFormCallback.formGroup = formGroup

	relation_groupFormCallback.CreationMode = (relation_group == nil)

	return
}

type RELATION_GROUPFormCallback struct {
	relation_group *models.RELATION_GROUP

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (relation_groupFormCallback *RELATION_GROUPFormCallback) OnSave() {

	log.Println("RELATION_GROUPFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	relation_groupFormCallback.probe.formStage.Checkout()

	if relation_groupFormCallback.relation_group == nil {
		relation_groupFormCallback.relation_group = new(models.RELATION_GROUP).Stage(relation_groupFormCallback.probe.stageOfInterest)
	}
	relation_group_ := relation_groupFormCallback.relation_group
	_ = relation_group_

	for _, formDiv := range relation_groupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(relation_group_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(relation_group_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(relation_group_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(relation_group_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:SPEC_RELATION_GROUPS.RELATION_GROUP":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_RELATION_GROUPS.RELATION_GROUP"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				relation_groupFormCallback.probe.stageOfInterest,
				relation_groupFormCallback.probe.backRepoOfInterest,
				relation_group_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS.RELATION_GROUP, relation_group_)
					pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS.RELATION_GROUP = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS.RELATION_GROUP, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](relation_groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS.RELATION_GROUP, relation_group_)
								pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS.RELATION_GROUP = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS.RELATION_GROUP, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS.RELATION_GROUP = append(newREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS.RELATION_GROUP, relation_group_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS.RELATION_GROUP = append(newREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS.RELATION_GROUP, relation_group_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if relation_groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_group_.Unstage(relation_groupFormCallback.probe.stageOfInterest)
	}

	relation_groupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.RELATION_GROUP](
		relation_groupFormCallback.probe,
	)
	relation_groupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if relation_groupFormCallback.CreationMode || relation_groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_groupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(relation_groupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RELATION_GROUPFormCallback(
			nil,
			relation_groupFormCallback.probe,
			newFormGroup,
		)
		relation_group := new(models.RELATION_GROUP)
		FillUpForm(relation_group, newFormGroup, relation_groupFormCallback.probe)
		relation_groupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(relation_groupFormCallback.probe)
}
func __gong__New__RELATION_GROUP_TYPEFormCallback(
	relation_group_type *models.RELATION_GROUP_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (relation_group_typeFormCallback *RELATION_GROUP_TYPEFormCallback) {
	relation_group_typeFormCallback = new(RELATION_GROUP_TYPEFormCallback)
	relation_group_typeFormCallback.probe = probe
	relation_group_typeFormCallback.relation_group_type = relation_group_type
	relation_group_typeFormCallback.formGroup = formGroup

	relation_group_typeFormCallback.CreationMode = (relation_group_type == nil)

	return
}

type RELATION_GROUP_TYPEFormCallback struct {
	relation_group_type *models.RELATION_GROUP_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (relation_group_typeFormCallback *RELATION_GROUP_TYPEFormCallback) OnSave() {

	log.Println("RELATION_GROUP_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	relation_group_typeFormCallback.probe.formStage.Checkout()

	if relation_group_typeFormCallback.relation_group_type == nil {
		relation_group_typeFormCallback.relation_group_type = new(models.RELATION_GROUP_TYPE).Stage(relation_group_typeFormCallback.probe.stageOfInterest)
	}
	relation_group_type_ := relation_group_typeFormCallback.relation_group_type
	_ = relation_group_type_

	for _, formDiv := range relation_group_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(relation_group_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(relation_group_type_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(relation_group_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(relation_group_type_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:SPEC_TYPES.RELATION_GROUP_TYPE":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_TYPES.RELATION_GROUP_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				relation_group_typeFormCallback.probe.stageOfInterest,
				relation_group_typeFormCallback.probe.backRepoOfInterest,
				relation_group_type_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_TYPES.RELATION_GROUP_TYPE, relation_group_type_)
					pastREQ_IF_CONTENTOwner.SPEC_TYPES.RELATION_GROUP_TYPE = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_TYPES.RELATION_GROUP_TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](relation_group_typeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_TYPES.RELATION_GROUP_TYPE, relation_group_type_)
								pastREQ_IF_CONTENTOwner.SPEC_TYPES.RELATION_GROUP_TYPE = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_TYPES.RELATION_GROUP_TYPE, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_TYPES.RELATION_GROUP_TYPE = append(newREQ_IF_CONTENTOwner.SPEC_TYPES.RELATION_GROUP_TYPE, relation_group_type_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_TYPES.RELATION_GROUP_TYPE = append(newREQ_IF_CONTENTOwner.SPEC_TYPES.RELATION_GROUP_TYPE, relation_group_type_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if relation_group_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_group_type_.Unstage(relation_group_typeFormCallback.probe.stageOfInterest)
	}

	relation_group_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.RELATION_GROUP_TYPE](
		relation_group_typeFormCallback.probe,
	)
	relation_group_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if relation_group_typeFormCallback.CreationMode || relation_group_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_group_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(relation_group_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RELATION_GROUP_TYPEFormCallback(
			nil,
			relation_group_typeFormCallback.probe,
			newFormGroup,
		)
		relation_group_type := new(models.RELATION_GROUP_TYPE)
		FillUpForm(relation_group_type, newFormGroup, relation_group_typeFormCallback.probe)
		relation_group_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(relation_group_typeFormCallback.probe)
}
func __gong__New__REQ_IFFormCallback(
	req_if *models.REQ_IF,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_ifFormCallback *REQ_IFFormCallback) {
	req_ifFormCallback = new(REQ_IFFormCallback)
	req_ifFormCallback.probe = probe
	req_ifFormCallback.req_if = req_if
	req_ifFormCallback.formGroup = formGroup

	req_ifFormCallback.CreationMode = (req_if == nil)

	return
}

type REQ_IFFormCallback struct {
	req_if *models.REQ_IF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_ifFormCallback *REQ_IFFormCallback) OnSave() {

	log.Println("REQ_IFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_ifFormCallback.probe.formStage.Checkout()

	if req_ifFormCallback.req_if == nil {
		req_ifFormCallback.req_if = new(models.REQ_IF).Stage(req_ifFormCallback.probe.stageOfInterest)
	}
	req_if_ := req_ifFormCallback.req_if
	_ = req_if_

	for _, formDiv := range req_ifFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_.Name), formDiv)
		case "THE_HEADER.REQ_IF_HEADER":
			FormDivSelectFieldToField(&(req_if_.THE_HEADER.REQ_IF_HEADER), req_ifFormCallback.probe.stageOfInterest, formDiv)
		case "CORE_CONTENT.REQ_IF_CONTENT":
			FormDivSelectFieldToField(&(req_if_.CORE_CONTENT.REQ_IF_CONTENT), req_ifFormCallback.probe.stageOfInterest, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(req_if_.Lang), formDiv)
		}
	}

	// manage the suppress operation
	if req_ifFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_.Unstage(req_ifFormCallback.probe.stageOfInterest)
	}

	req_ifFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQ_IF](
		req_ifFormCallback.probe,
	)
	req_ifFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if req_ifFormCallback.CreationMode || req_ifFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_ifFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(req_ifFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IFFormCallback(
			nil,
			req_ifFormCallback.probe,
			newFormGroup,
		)
		req_if := new(models.REQ_IF)
		FillUpForm(req_if, newFormGroup, req_ifFormCallback.probe)
		req_ifFormCallback.probe.formStage.Commit()
	}

	fillUpTree(req_ifFormCallback.probe)
}
func __gong__New__REQ_IF_CONTENTFormCallback(
	req_if_content *models.REQ_IF_CONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_contentFormCallback *REQ_IF_CONTENTFormCallback) {
	req_if_contentFormCallback = new(REQ_IF_CONTENTFormCallback)
	req_if_contentFormCallback.probe = probe
	req_if_contentFormCallback.req_if_content = req_if_content
	req_if_contentFormCallback.formGroup = formGroup

	req_if_contentFormCallback.CreationMode = (req_if_content == nil)

	return
}

type REQ_IF_CONTENTFormCallback struct {
	req_if_content *models.REQ_IF_CONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_contentFormCallback *REQ_IF_CONTENTFormCallback) OnSave() {

	log.Println("REQ_IF_CONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_contentFormCallback.probe.formStage.Checkout()

	if req_if_contentFormCallback.req_if_content == nil {
		req_if_contentFormCallback.req_if_content = new(models.REQ_IF_CONTENT).Stage(req_if_contentFormCallback.probe.stageOfInterest)
	}
	req_if_content_ := req_if_contentFormCallback.req_if_content
	_ = req_if_content_

	for _, formDiv := range req_if_contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_content_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if req_if_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_content_.Unstage(req_if_contentFormCallback.probe.stageOfInterest)
	}

	req_if_contentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQ_IF_CONTENT](
		req_if_contentFormCallback.probe,
	)
	req_if_contentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if req_if_contentFormCallback.CreationMode || req_if_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(req_if_contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_CONTENTFormCallback(
			nil,
			req_if_contentFormCallback.probe,
			newFormGroup,
		)
		req_if_content := new(models.REQ_IF_CONTENT)
		FillUpForm(req_if_content, newFormGroup, req_if_contentFormCallback.probe)
		req_if_contentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(req_if_contentFormCallback.probe)
}
func __gong__New__REQ_IF_HEADERFormCallback(
	req_if_header *models.REQ_IF_HEADER,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_headerFormCallback *REQ_IF_HEADERFormCallback) {
	req_if_headerFormCallback = new(REQ_IF_HEADERFormCallback)
	req_if_headerFormCallback.probe = probe
	req_if_headerFormCallback.req_if_header = req_if_header
	req_if_headerFormCallback.formGroup = formGroup

	req_if_headerFormCallback.CreationMode = (req_if_header == nil)

	return
}

type REQ_IF_HEADERFormCallback struct {
	req_if_header *models.REQ_IF_HEADER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_headerFormCallback *REQ_IF_HEADERFormCallback) OnSave() {

	log.Println("REQ_IF_HEADERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_headerFormCallback.probe.formStage.Checkout()

	if req_if_headerFormCallback.req_if_header == nil {
		req_if_headerFormCallback.req_if_header = new(models.REQ_IF_HEADER).Stage(req_if_headerFormCallback.probe.stageOfInterest)
	}
	req_if_header_ := req_if_headerFormCallback.req_if_header
	_ = req_if_header_

	for _, formDiv := range req_if_headerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_header_.Name), formDiv)
		case "COMMENT":
			FormDivBasicFieldToField(&(req_if_header_.COMMENT), formDiv)
		case "CREATION_TIME":
			FormDivBasicFieldToField(&(req_if_header_.CREATION_TIME), formDiv)
		case "REPOSITORY_ID":
			FormDivBasicFieldToField(&(req_if_header_.REPOSITORY_ID), formDiv)
		case "REQ_IF_TOOL_ID":
			FormDivBasicFieldToField(&(req_if_header_.REQ_IF_TOOL_ID), formDiv)
		case "REQ_IF_VERSION":
			FormDivBasicFieldToField(&(req_if_header_.REQ_IF_VERSION), formDiv)
		case "SOURCE_TOOL_ID":
			FormDivBasicFieldToField(&(req_if_header_.SOURCE_TOOL_ID), formDiv)
		case "TITLE":
			FormDivBasicFieldToField(&(req_if_header_.TITLE), formDiv)
		}
	}

	// manage the suppress operation
	if req_if_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_header_.Unstage(req_if_headerFormCallback.probe.stageOfInterest)
	}

	req_if_headerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQ_IF_HEADER](
		req_if_headerFormCallback.probe,
	)
	req_if_headerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if req_if_headerFormCallback.CreationMode || req_if_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_headerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(req_if_headerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_HEADERFormCallback(
			nil,
			req_if_headerFormCallback.probe,
			newFormGroup,
		)
		req_if_header := new(models.REQ_IF_HEADER)
		FillUpForm(req_if_header, newFormGroup, req_if_headerFormCallback.probe)
		req_if_headerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(req_if_headerFormCallback.probe)
}
func __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
	req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_tool_extensionFormCallback *REQ_IF_TOOL_EXTENSIONFormCallback) {
	req_if_tool_extensionFormCallback = new(REQ_IF_TOOL_EXTENSIONFormCallback)
	req_if_tool_extensionFormCallback.probe = probe
	req_if_tool_extensionFormCallback.req_if_tool_extension = req_if_tool_extension
	req_if_tool_extensionFormCallback.formGroup = formGroup

	req_if_tool_extensionFormCallback.CreationMode = (req_if_tool_extension == nil)

	return
}

type REQ_IF_TOOL_EXTENSIONFormCallback struct {
	req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_tool_extensionFormCallback *REQ_IF_TOOL_EXTENSIONFormCallback) OnSave() {

	log.Println("REQ_IF_TOOL_EXTENSIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_tool_extensionFormCallback.probe.formStage.Checkout()

	if req_if_tool_extensionFormCallback.req_if_tool_extension == nil {
		req_if_tool_extensionFormCallback.req_if_tool_extension = new(models.REQ_IF_TOOL_EXTENSION).Stage(req_if_tool_extensionFormCallback.probe.stageOfInterest)
	}
	req_if_tool_extension_ := req_if_tool_extensionFormCallback.req_if_tool_extension
	_ = req_if_tool_extension_

	for _, formDiv := range req_if_tool_extensionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_tool_extension_.Name), formDiv)
		case "REQ_IF:TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION":
			// we need to retrieve the field owner before the change
			var pastREQ_IFOwner *models.REQ_IF
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF"
			rf.Fieldname = "TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				req_if_tool_extensionFormCallback.probe.stageOfInterest,
				req_if_tool_extensionFormCallback.probe.backRepoOfInterest,
				req_if_tool_extension_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IFOwner = reverseFieldOwner.(*models.REQ_IF)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IFOwner != nil {
					idx := slices.Index(pastREQ_IFOwner.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION, req_if_tool_extension_)
					pastREQ_IFOwner.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION = slices.Delete(pastREQ_IFOwner.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if := range *models.GetGongstructInstancesSet[models.REQ_IF](req_if_tool_extensionFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IFOwner := _req_if // we have a match
						if pastREQ_IFOwner != nil {
							if newREQ_IFOwner != pastREQ_IFOwner {
								idx := slices.Index(pastREQ_IFOwner.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION, req_if_tool_extension_)
								pastREQ_IFOwner.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION = slices.Delete(pastREQ_IFOwner.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION, idx, idx+1)
								newREQ_IFOwner.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION = append(newREQ_IFOwner.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION, req_if_tool_extension_)
							}
						} else {
							newREQ_IFOwner.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION = append(newREQ_IFOwner.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION, req_if_tool_extension_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if req_if_tool_extensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_tool_extension_.Unstage(req_if_tool_extensionFormCallback.probe.stageOfInterest)
	}

	req_if_tool_extensionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQ_IF_TOOL_EXTENSION](
		req_if_tool_extensionFormCallback.probe,
	)
	req_if_tool_extensionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if req_if_tool_extensionFormCallback.CreationMode || req_if_tool_extensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_tool_extensionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(req_if_tool_extensionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
			nil,
			req_if_tool_extensionFormCallback.probe,
			newFormGroup,
		)
		req_if_tool_extension := new(models.REQ_IF_TOOL_EXTENSION)
		FillUpForm(req_if_tool_extension, newFormGroup, req_if_tool_extensionFormCallback.probe)
		req_if_tool_extensionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(req_if_tool_extensionFormCallback.probe)
}
func __gong__New__SPECIFICATIONFormCallback(
	specification *models.SPECIFICATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (specificationFormCallback *SPECIFICATIONFormCallback) {
	specificationFormCallback = new(SPECIFICATIONFormCallback)
	specificationFormCallback.probe = probe
	specificationFormCallback.specification = specification
	specificationFormCallback.formGroup = formGroup

	specificationFormCallback.CreationMode = (specification == nil)

	return
}

type SPECIFICATIONFormCallback struct {
	specification *models.SPECIFICATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specificationFormCallback *SPECIFICATIONFormCallback) OnSave() {

	log.Println("SPECIFICATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specificationFormCallback.probe.formStage.Checkout()

	if specificationFormCallback.specification == nil {
		specificationFormCallback.specification = new(models.SPECIFICATION).Stage(specificationFormCallback.probe.stageOfInterest)
	}
	specification_ := specificationFormCallback.specification
	_ = specification_

	for _, formDiv := range specificationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specification_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(specification_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(specification_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(specification_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:SPECIFICATIONS.SPECIFICATION":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPECIFICATIONS.SPECIFICATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				specificationFormCallback.probe.stageOfInterest,
				specificationFormCallback.probe.backRepoOfInterest,
				specification_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPECIFICATIONS.SPECIFICATION, specification_)
					pastREQ_IF_CONTENTOwner.SPECIFICATIONS.SPECIFICATION = slices.Delete(pastREQ_IF_CONTENTOwner.SPECIFICATIONS.SPECIFICATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](specificationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPECIFICATIONS.SPECIFICATION, specification_)
								pastREQ_IF_CONTENTOwner.SPECIFICATIONS.SPECIFICATION = slices.Delete(pastREQ_IF_CONTENTOwner.SPECIFICATIONS.SPECIFICATION, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPECIFICATIONS.SPECIFICATION = append(newREQ_IF_CONTENTOwner.SPECIFICATIONS.SPECIFICATION, specification_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPECIFICATIONS.SPECIFICATION = append(newREQ_IF_CONTENTOwner.SPECIFICATIONS.SPECIFICATION, specification_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if specificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specification_.Unstage(specificationFormCallback.probe.stageOfInterest)
	}

	specificationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECIFICATION](
		specificationFormCallback.probe,
	)
	specificationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specificationFormCallback.CreationMode || specificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specificationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specificationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECIFICATIONFormCallback(
			nil,
			specificationFormCallback.probe,
			newFormGroup,
		)
		specification := new(models.SPECIFICATION)
		FillUpForm(specification, newFormGroup, specificationFormCallback.probe)
		specificationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specificationFormCallback.probe)
}
func __gong__New__SPECIFICATION_TYPEFormCallback(
	specification_type *models.SPECIFICATION_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (specification_typeFormCallback *SPECIFICATION_TYPEFormCallback) {
	specification_typeFormCallback = new(SPECIFICATION_TYPEFormCallback)
	specification_typeFormCallback.probe = probe
	specification_typeFormCallback.specification_type = specification_type
	specification_typeFormCallback.formGroup = formGroup

	specification_typeFormCallback.CreationMode = (specification_type == nil)

	return
}

type SPECIFICATION_TYPEFormCallback struct {
	specification_type *models.SPECIFICATION_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specification_typeFormCallback *SPECIFICATION_TYPEFormCallback) OnSave() {

	log.Println("SPECIFICATION_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specification_typeFormCallback.probe.formStage.Checkout()

	if specification_typeFormCallback.specification_type == nil {
		specification_typeFormCallback.specification_type = new(models.SPECIFICATION_TYPE).Stage(specification_typeFormCallback.probe.stageOfInterest)
	}
	specification_type_ := specification_typeFormCallback.specification_type
	_ = specification_type_

	for _, formDiv := range specification_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specification_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(specification_type_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(specification_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(specification_type_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:SPEC_TYPES.SPECIFICATION_TYPE":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_TYPES.SPECIFICATION_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				specification_typeFormCallback.probe.stageOfInterest,
				specification_typeFormCallback.probe.backRepoOfInterest,
				specification_type_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPECIFICATION_TYPE, specification_type_)
					pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPECIFICATION_TYPE = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPECIFICATION_TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](specification_typeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPECIFICATION_TYPE, specification_type_)
								pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPECIFICATION_TYPE = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPECIFICATION_TYPE, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_TYPES.SPECIFICATION_TYPE = append(newREQ_IF_CONTENTOwner.SPEC_TYPES.SPECIFICATION_TYPE, specification_type_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_TYPES.SPECIFICATION_TYPE = append(newREQ_IF_CONTENTOwner.SPEC_TYPES.SPECIFICATION_TYPE, specification_type_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if specification_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specification_type_.Unstage(specification_typeFormCallback.probe.stageOfInterest)
	}

	specification_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECIFICATION_TYPE](
		specification_typeFormCallback.probe,
	)
	specification_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specification_typeFormCallback.CreationMode || specification_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specification_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specification_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECIFICATION_TYPEFormCallback(
			nil,
			specification_typeFormCallback.probe,
			newFormGroup,
		)
		specification_type := new(models.SPECIFICATION_TYPE)
		FillUpForm(specification_type, newFormGroup, specification_typeFormCallback.probe)
		specification_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specification_typeFormCallback.probe)
}
func __gong__New__SPEC_HIERARCHYFormCallback(
	spec_hierarchy *models.SPEC_HIERARCHY,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_hierarchyFormCallback *SPEC_HIERARCHYFormCallback) {
	spec_hierarchyFormCallback = new(SPEC_HIERARCHYFormCallback)
	spec_hierarchyFormCallback.probe = probe
	spec_hierarchyFormCallback.spec_hierarchy = spec_hierarchy
	spec_hierarchyFormCallback.formGroup = formGroup

	spec_hierarchyFormCallback.CreationMode = (spec_hierarchy == nil)

	return
}

type SPEC_HIERARCHYFormCallback struct {
	spec_hierarchy *models.SPEC_HIERARCHY

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_hierarchyFormCallback *SPEC_HIERARCHYFormCallback) OnSave() {

	log.Println("SPEC_HIERARCHYFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_hierarchyFormCallback.probe.formStage.Checkout()

	if spec_hierarchyFormCallback.spec_hierarchy == nil {
		spec_hierarchyFormCallback.spec_hierarchy = new(models.SPEC_HIERARCHY).Stage(spec_hierarchyFormCallback.probe.stageOfInterest)
	}
	spec_hierarchy_ := spec_hierarchyFormCallback.spec_hierarchy
	_ = spec_hierarchy_

	for _, formDiv := range spec_hierarchyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_hierarchy_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_hierarchy_.DESC), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(spec_hierarchy_.IS_EDITABLE), formDiv)
		case "IS_TABLE_INTERNAL":
			FormDivBasicFieldToField(&(spec_hierarchy_.IS_TABLE_INTERNAL), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_hierarchy_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_hierarchy_.LONG_NAME), formDiv)
		case "SPECIFICATION:CHILDREN.SPEC_HIERARCHY":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "CHILDREN.SPEC_HIERARCHY"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_hierarchyFormCallback.probe.stageOfInterest,
				spec_hierarchyFormCallback.probe.backRepoOfInterest,
				spec_hierarchy_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.CHILDREN.SPEC_HIERARCHY, spec_hierarchy_)
					pastSPECIFICATIONOwner.CHILDREN.SPEC_HIERARCHY = slices.Delete(pastSPECIFICATIONOwner.CHILDREN.SPEC_HIERARCHY, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](spec_hierarchyFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.CHILDREN.SPEC_HIERARCHY, spec_hierarchy_)
								pastSPECIFICATIONOwner.CHILDREN.SPEC_HIERARCHY = slices.Delete(pastSPECIFICATIONOwner.CHILDREN.SPEC_HIERARCHY, idx, idx+1)
								newSPECIFICATIONOwner.CHILDREN.SPEC_HIERARCHY = append(newSPECIFICATIONOwner.CHILDREN.SPEC_HIERARCHY, spec_hierarchy_)
							}
						} else {
							newSPECIFICATIONOwner.CHILDREN.SPEC_HIERARCHY = append(newSPECIFICATIONOwner.CHILDREN.SPEC_HIERARCHY, spec_hierarchy_)
						}
					}
				}
			}
		case "SPEC_HIERARCHY:CHILDREN.SPEC_HIERARCHY":
			// we need to retrieve the field owner before the change
			var pastSPEC_HIERARCHYOwner *models.SPEC_HIERARCHY
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "CHILDREN.SPEC_HIERARCHY"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_hierarchyFormCallback.probe.stageOfInterest,
				spec_hierarchyFormCallback.probe.backRepoOfInterest,
				spec_hierarchy_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_HIERARCHYOwner = reverseFieldOwner.(*models.SPEC_HIERARCHY)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_HIERARCHYOwner != nil {
					idx := slices.Index(pastSPEC_HIERARCHYOwner.CHILDREN.SPEC_HIERARCHY, spec_hierarchy_)
					pastSPEC_HIERARCHYOwner.CHILDREN.SPEC_HIERARCHY = slices.Delete(pastSPEC_HIERARCHYOwner.CHILDREN.SPEC_HIERARCHY, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_hierarchy := range *models.GetGongstructInstancesSet[models.SPEC_HIERARCHY](spec_hierarchyFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_hierarchy.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_HIERARCHYOwner := _spec_hierarchy // we have a match
						if pastSPEC_HIERARCHYOwner != nil {
							if newSPEC_HIERARCHYOwner != pastSPEC_HIERARCHYOwner {
								idx := slices.Index(pastSPEC_HIERARCHYOwner.CHILDREN.SPEC_HIERARCHY, spec_hierarchy_)
								pastSPEC_HIERARCHYOwner.CHILDREN.SPEC_HIERARCHY = slices.Delete(pastSPEC_HIERARCHYOwner.CHILDREN.SPEC_HIERARCHY, idx, idx+1)
								newSPEC_HIERARCHYOwner.CHILDREN.SPEC_HIERARCHY = append(newSPEC_HIERARCHYOwner.CHILDREN.SPEC_HIERARCHY, spec_hierarchy_)
							}
						} else {
							newSPEC_HIERARCHYOwner.CHILDREN.SPEC_HIERARCHY = append(newSPEC_HIERARCHYOwner.CHILDREN.SPEC_HIERARCHY, spec_hierarchy_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if spec_hierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_hierarchy_.Unstage(spec_hierarchyFormCallback.probe.stageOfInterest)
	}

	spec_hierarchyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_HIERARCHY](
		spec_hierarchyFormCallback.probe,
	)
	spec_hierarchyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_hierarchyFormCallback.CreationMode || spec_hierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_hierarchyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_hierarchyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_HIERARCHYFormCallback(
			nil,
			spec_hierarchyFormCallback.probe,
			newFormGroup,
		)
		spec_hierarchy := new(models.SPEC_HIERARCHY)
		FillUpForm(spec_hierarchy, newFormGroup, spec_hierarchyFormCallback.probe)
		spec_hierarchyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_hierarchyFormCallback.probe)
}
func __gong__New__SPEC_OBJECTFormCallback(
	spec_object *models.SPEC_OBJECT,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_objectFormCallback *SPEC_OBJECTFormCallback) {
	spec_objectFormCallback = new(SPEC_OBJECTFormCallback)
	spec_objectFormCallback.probe = probe
	spec_objectFormCallback.spec_object = spec_object
	spec_objectFormCallback.formGroup = formGroup

	spec_objectFormCallback.CreationMode = (spec_object == nil)

	return
}

type SPEC_OBJECTFormCallback struct {
	spec_object *models.SPEC_OBJECT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_objectFormCallback *SPEC_OBJECTFormCallback) OnSave() {

	log.Println("SPEC_OBJECTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_objectFormCallback.probe.formStage.Checkout()

	if spec_objectFormCallback.spec_object == nil {
		spec_objectFormCallback.spec_object = new(models.SPEC_OBJECT).Stage(spec_objectFormCallback.probe.stageOfInterest)
	}
	spec_object_ := spec_objectFormCallback.spec_object
	_ = spec_object_

	for _, formDiv := range spec_objectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_object_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_object_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_object_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_object_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:SPEC_OBJECTS.SPEC_OBJECT":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_OBJECTS.SPEC_OBJECT"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_objectFormCallback.probe.stageOfInterest,
				spec_objectFormCallback.probe.backRepoOfInterest,
				spec_object_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_OBJECTS.SPEC_OBJECT, spec_object_)
					pastREQ_IF_CONTENTOwner.SPEC_OBJECTS.SPEC_OBJECT = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_OBJECTS.SPEC_OBJECT, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](spec_objectFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_OBJECTS.SPEC_OBJECT, spec_object_)
								pastREQ_IF_CONTENTOwner.SPEC_OBJECTS.SPEC_OBJECT = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_OBJECTS.SPEC_OBJECT, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_OBJECTS.SPEC_OBJECT = append(newREQ_IF_CONTENTOwner.SPEC_OBJECTS.SPEC_OBJECT, spec_object_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_OBJECTS.SPEC_OBJECT = append(newREQ_IF_CONTENTOwner.SPEC_OBJECTS.SPEC_OBJECT, spec_object_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if spec_objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_object_.Unstage(spec_objectFormCallback.probe.stageOfInterest)
	}

	spec_objectFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_OBJECT](
		spec_objectFormCallback.probe,
	)
	spec_objectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_objectFormCallback.CreationMode || spec_objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_objectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_objectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_OBJECTFormCallback(
			nil,
			spec_objectFormCallback.probe,
			newFormGroup,
		)
		spec_object := new(models.SPEC_OBJECT)
		FillUpForm(spec_object, newFormGroup, spec_objectFormCallback.probe)
		spec_objectFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_objectFormCallback.probe)
}
func __gong__New__SPEC_OBJECT_TYPEFormCallback(
	spec_object_type *models.SPEC_OBJECT_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_object_typeFormCallback *SPEC_OBJECT_TYPEFormCallback) {
	spec_object_typeFormCallback = new(SPEC_OBJECT_TYPEFormCallback)
	spec_object_typeFormCallback.probe = probe
	spec_object_typeFormCallback.spec_object_type = spec_object_type
	spec_object_typeFormCallback.formGroup = formGroup

	spec_object_typeFormCallback.CreationMode = (spec_object_type == nil)

	return
}

type SPEC_OBJECT_TYPEFormCallback struct {
	spec_object_type *models.SPEC_OBJECT_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_object_typeFormCallback *SPEC_OBJECT_TYPEFormCallback) OnSave() {

	log.Println("SPEC_OBJECT_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_object_typeFormCallback.probe.formStage.Checkout()

	if spec_object_typeFormCallback.spec_object_type == nil {
		spec_object_typeFormCallback.spec_object_type = new(models.SPEC_OBJECT_TYPE).Stage(spec_object_typeFormCallback.probe.stageOfInterest)
	}
	spec_object_type_ := spec_object_typeFormCallback.spec_object_type
	_ = spec_object_type_

	for _, formDiv := range spec_object_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_object_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_object_type_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_object_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_object_type_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:SPEC_TYPES.SPEC_OBJECT_TYPE":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_TYPES.SPEC_OBJECT_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_object_typeFormCallback.probe.stageOfInterest,
				spec_object_typeFormCallback.probe.backRepoOfInterest,
				spec_object_type_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_OBJECT_TYPE, spec_object_type_)
					pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_OBJECT_TYPE = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_OBJECT_TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](spec_object_typeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_OBJECT_TYPE, spec_object_type_)
								pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_OBJECT_TYPE = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_OBJECT_TYPE, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_OBJECT_TYPE = append(newREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_OBJECT_TYPE, spec_object_type_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_OBJECT_TYPE = append(newREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_OBJECT_TYPE, spec_object_type_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if spec_object_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_object_type_.Unstage(spec_object_typeFormCallback.probe.stageOfInterest)
	}

	spec_object_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_OBJECT_TYPE](
		spec_object_typeFormCallback.probe,
	)
	spec_object_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_object_typeFormCallback.CreationMode || spec_object_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_object_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_object_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_OBJECT_TYPEFormCallback(
			nil,
			spec_object_typeFormCallback.probe,
			newFormGroup,
		)
		spec_object_type := new(models.SPEC_OBJECT_TYPE)
		FillUpForm(spec_object_type, newFormGroup, spec_object_typeFormCallback.probe)
		spec_object_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_object_typeFormCallback.probe)
}
func __gong__New__SPEC_RELATIONFormCallback(
	spec_relation *models.SPEC_RELATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_relationFormCallback *SPEC_RELATIONFormCallback) {
	spec_relationFormCallback = new(SPEC_RELATIONFormCallback)
	spec_relationFormCallback.probe = probe
	spec_relationFormCallback.spec_relation = spec_relation
	spec_relationFormCallback.formGroup = formGroup

	spec_relationFormCallback.CreationMode = (spec_relation == nil)

	return
}

type SPEC_RELATIONFormCallback struct {
	spec_relation *models.SPEC_RELATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_relationFormCallback *SPEC_RELATIONFormCallback) OnSave() {

	log.Println("SPEC_RELATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_relationFormCallback.probe.formStage.Checkout()

	if spec_relationFormCallback.spec_relation == nil {
		spec_relationFormCallback.spec_relation = new(models.SPEC_RELATION).Stage(spec_relationFormCallback.probe.stageOfInterest)
	}
	spec_relation_ := spec_relationFormCallback.spec_relation
	_ = spec_relation_

	for _, formDiv := range spec_relationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_relation_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_relation_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_relation_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_relation_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:SPEC_RELATIONS.SPEC_RELATION":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_RELATIONS.SPEC_RELATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_relationFormCallback.probe.stageOfInterest,
				spec_relationFormCallback.probe.backRepoOfInterest,
				spec_relation_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_RELATIONS.SPEC_RELATION, spec_relation_)
					pastREQ_IF_CONTENTOwner.SPEC_RELATIONS.SPEC_RELATION = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_RELATIONS.SPEC_RELATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](spec_relationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_RELATIONS.SPEC_RELATION, spec_relation_)
								pastREQ_IF_CONTENTOwner.SPEC_RELATIONS.SPEC_RELATION = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_RELATIONS.SPEC_RELATION, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_RELATIONS.SPEC_RELATION = append(newREQ_IF_CONTENTOwner.SPEC_RELATIONS.SPEC_RELATION, spec_relation_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_RELATIONS.SPEC_RELATION = append(newREQ_IF_CONTENTOwner.SPEC_RELATIONS.SPEC_RELATION, spec_relation_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if spec_relationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relation_.Unstage(spec_relationFormCallback.probe.stageOfInterest)
	}

	spec_relationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_RELATION](
		spec_relationFormCallback.probe,
	)
	spec_relationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_relationFormCallback.CreationMode || spec_relationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_relationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_RELATIONFormCallback(
			nil,
			spec_relationFormCallback.probe,
			newFormGroup,
		)
		spec_relation := new(models.SPEC_RELATION)
		FillUpForm(spec_relation, newFormGroup, spec_relationFormCallback.probe)
		spec_relationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_relationFormCallback.probe)
}
func __gong__New__SPEC_RELATION_TYPEFormCallback(
	spec_relation_type *models.SPEC_RELATION_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_relation_typeFormCallback *SPEC_RELATION_TYPEFormCallback) {
	spec_relation_typeFormCallback = new(SPEC_RELATION_TYPEFormCallback)
	spec_relation_typeFormCallback.probe = probe
	spec_relation_typeFormCallback.spec_relation_type = spec_relation_type
	spec_relation_typeFormCallback.formGroup = formGroup

	spec_relation_typeFormCallback.CreationMode = (spec_relation_type == nil)

	return
}

type SPEC_RELATION_TYPEFormCallback struct {
	spec_relation_type *models.SPEC_RELATION_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_relation_typeFormCallback *SPEC_RELATION_TYPEFormCallback) OnSave() {

	log.Println("SPEC_RELATION_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_relation_typeFormCallback.probe.formStage.Checkout()

	if spec_relation_typeFormCallback.spec_relation_type == nil {
		spec_relation_typeFormCallback.spec_relation_type = new(models.SPEC_RELATION_TYPE).Stage(spec_relation_typeFormCallback.probe.stageOfInterest)
	}
	spec_relation_type_ := spec_relation_typeFormCallback.spec_relation_type
	_ = spec_relation_type_

	for _, formDiv := range spec_relation_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_relation_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_relation_type_.DESC), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_relation_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_relation_type_.LONG_NAME), formDiv)
		case "REQ_IF_CONTENT:SPEC_TYPES.SPEC_RELATION_TYPE":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_TYPES.SPEC_RELATION_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_relation_typeFormCallback.probe.stageOfInterest,
				spec_relation_typeFormCallback.probe.backRepoOfInterest,
				spec_relation_type_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_RELATION_TYPE, spec_relation_type_)
					pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_RELATION_TYPE = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_RELATION_TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](spec_relation_typeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_RELATION_TYPE, spec_relation_type_)
								pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_RELATION_TYPE = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_RELATION_TYPE, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_RELATION_TYPE = append(newREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_RELATION_TYPE, spec_relation_type_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_RELATION_TYPE = append(newREQ_IF_CONTENTOwner.SPEC_TYPES.SPEC_RELATION_TYPE, spec_relation_type_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if spec_relation_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relation_type_.Unstage(spec_relation_typeFormCallback.probe.stageOfInterest)
	}

	spec_relation_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_RELATION_TYPE](
		spec_relation_typeFormCallback.probe,
	)
	spec_relation_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_relation_typeFormCallback.CreationMode || spec_relation_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relation_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_relation_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_RELATION_TYPEFormCallback(
			nil,
			spec_relation_typeFormCallback.probe,
			newFormGroup,
		)
		spec_relation_type := new(models.SPEC_RELATION_TYPE)
		FillUpForm(spec_relation_type, newFormGroup, spec_relation_typeFormCallback.probe)
		spec_relation_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_relation_typeFormCallback.probe)
}
func __gong__New__XHTML_CONTENTFormCallback(
	xhtml_content *models.XHTML_CONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_contentFormCallback *XHTML_CONTENTFormCallback) {
	xhtml_contentFormCallback = new(XHTML_CONTENTFormCallback)
	xhtml_contentFormCallback.probe = probe
	xhtml_contentFormCallback.xhtml_content = xhtml_content
	xhtml_contentFormCallback.formGroup = formGroup

	xhtml_contentFormCallback.CreationMode = (xhtml_content == nil)

	return
}

type XHTML_CONTENTFormCallback struct {
	xhtml_content *models.XHTML_CONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_contentFormCallback *XHTML_CONTENTFormCallback) OnSave() {

	log.Println("XHTML_CONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_contentFormCallback.probe.formStage.Checkout()

	if xhtml_contentFormCallback.xhtml_content == nil {
		xhtml_contentFormCallback.xhtml_content = new(models.XHTML_CONTENT).Stage(xhtml_contentFormCallback.probe.stageOfInterest)
	}
	xhtml_content_ := xhtml_contentFormCallback.xhtml_content
	_ = xhtml_content_

	for _, formDiv := range xhtml_contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_content_.Name), formDiv)
		case "ATTRIBUTE_VALUE_XHTML:THE_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_XHTMLOwner *models.ATTRIBUTE_VALUE_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "THE_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				xhtml_contentFormCallback.probe.stageOfInterest,
				xhtml_contentFormCallback.probe.backRepoOfInterest,
				xhtml_content_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_XHTMLOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_XHTMLOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, xhtml_content_)
					pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE = slices.Delete(pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_xhtml := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_XHTML](xhtml_contentFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_XHTMLOwner := _attribute_value_xhtml // we have a match
						if pastATTRIBUTE_VALUE_XHTMLOwner != nil {
							if newATTRIBUTE_VALUE_XHTMLOwner != pastATTRIBUTE_VALUE_XHTMLOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, xhtml_content_)
								pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE = slices.Delete(pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, idx, idx+1)
								newATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE = append(newATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, xhtml_content_)
							}
						} else {
							newATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE = append(newATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, xhtml_content_)
						}
					}
				}
			}
		case "ATTRIBUTE_VALUE_XHTML:THE_ORIGINAL_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_XHTMLOwner *models.ATTRIBUTE_VALUE_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "THE_ORIGINAL_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				xhtml_contentFormCallback.probe.stageOfInterest,
				xhtml_contentFormCallback.probe.backRepoOfInterest,
				xhtml_content_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_XHTMLOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_XHTMLOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, xhtml_content_)
					pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE = slices.Delete(pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_xhtml := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_XHTML](xhtml_contentFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_XHTMLOwner := _attribute_value_xhtml // we have a match
						if pastATTRIBUTE_VALUE_XHTMLOwner != nil {
							if newATTRIBUTE_VALUE_XHTMLOwner != pastATTRIBUTE_VALUE_XHTMLOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, xhtml_content_)
								pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE = slices.Delete(pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, idx, idx+1)
								newATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE = append(newATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, xhtml_content_)
							}
						} else {
							newATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE = append(newATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, xhtml_content_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if xhtml_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_content_.Unstage(xhtml_contentFormCallback.probe.stageOfInterest)
	}

	xhtml_contentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.XHTML_CONTENT](
		xhtml_contentFormCallback.probe,
	)
	xhtml_contentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_contentFormCallback.CreationMode || xhtml_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__XHTML_CONTENTFormCallback(
			nil,
			xhtml_contentFormCallback.probe,
			newFormGroup,
		)
		xhtml_content := new(models.XHTML_CONTENT)
		FillUpForm(xhtml_content, newFormGroup, xhtml_contentFormCallback.probe)
		xhtml_contentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_contentFormCallback.probe)
}
func __gong__New__Xhtml_InlPres_typeFormCallback(
	xhtml_inlpres_type *models.Xhtml_InlPres_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_inlpres_typeFormCallback *Xhtml_InlPres_typeFormCallback) {
	xhtml_inlpres_typeFormCallback = new(Xhtml_InlPres_typeFormCallback)
	xhtml_inlpres_typeFormCallback.probe = probe
	xhtml_inlpres_typeFormCallback.xhtml_inlpres_type = xhtml_inlpres_type
	xhtml_inlpres_typeFormCallback.formGroup = formGroup

	xhtml_inlpres_typeFormCallback.CreationMode = (xhtml_inlpres_type == nil)

	return
}

type Xhtml_InlPres_typeFormCallback struct {
	xhtml_inlpres_type *models.Xhtml_InlPres_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_inlpres_typeFormCallback *Xhtml_InlPres_typeFormCallback) OnSave() {

	log.Println("Xhtml_InlPres_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_inlpres_typeFormCallback.probe.formStage.Checkout()

	if xhtml_inlpres_typeFormCallback.xhtml_inlpres_type == nil {
		xhtml_inlpres_typeFormCallback.xhtml_inlpres_type = new(models.Xhtml_InlPres_type).Stage(xhtml_inlpres_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_inlpres_type_ := xhtml_inlpres_typeFormCallback.xhtml_inlpres_type
	_ = xhtml_inlpres_type_

	for _, formDiv := range xhtml_inlpres_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_inlpres_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_inlpres_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_inlpres_type_.Unstage(xhtml_inlpres_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_inlpres_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_InlPres_type](
		xhtml_inlpres_typeFormCallback.probe,
	)
	xhtml_inlpres_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_inlpres_typeFormCallback.CreationMode || xhtml_inlpres_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_inlpres_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_inlpres_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_InlPres_typeFormCallback(
			nil,
			xhtml_inlpres_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_inlpres_type := new(models.Xhtml_InlPres_type)
		FillUpForm(xhtml_inlpres_type, newFormGroup, xhtml_inlpres_typeFormCallback.probe)
		xhtml_inlpres_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_inlpres_typeFormCallback.probe)
}
func __gong__New__Xhtml_a_typeFormCallback(
	xhtml_a_type *models.Xhtml_a_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_a_typeFormCallback *Xhtml_a_typeFormCallback) {
	xhtml_a_typeFormCallback = new(Xhtml_a_typeFormCallback)
	xhtml_a_typeFormCallback.probe = probe
	xhtml_a_typeFormCallback.xhtml_a_type = xhtml_a_type
	xhtml_a_typeFormCallback.formGroup = formGroup

	xhtml_a_typeFormCallback.CreationMode = (xhtml_a_type == nil)

	return
}

type Xhtml_a_typeFormCallback struct {
	xhtml_a_type *models.Xhtml_a_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_a_typeFormCallback *Xhtml_a_typeFormCallback) OnSave() {

	log.Println("Xhtml_a_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_a_typeFormCallback.probe.formStage.Checkout()

	if xhtml_a_typeFormCallback.xhtml_a_type == nil {
		xhtml_a_typeFormCallback.xhtml_a_type = new(models.Xhtml_a_type).Stage(xhtml_a_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_a_type_ := xhtml_a_typeFormCallback.xhtml_a_type
	_ = xhtml_a_type_

	for _, formDiv := range xhtml_a_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_a_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_a_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_a_type_.Unstage(xhtml_a_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_a_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_a_type](
		xhtml_a_typeFormCallback.probe,
	)
	xhtml_a_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_a_typeFormCallback.CreationMode || xhtml_a_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_a_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_a_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_a_typeFormCallback(
			nil,
			xhtml_a_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_a_type := new(models.Xhtml_a_type)
		FillUpForm(xhtml_a_type, newFormGroup, xhtml_a_typeFormCallback.probe)
		xhtml_a_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_a_typeFormCallback.probe)
}
func __gong__New__Xhtml_abbr_typeFormCallback(
	xhtml_abbr_type *models.Xhtml_abbr_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_abbr_typeFormCallback *Xhtml_abbr_typeFormCallback) {
	xhtml_abbr_typeFormCallback = new(Xhtml_abbr_typeFormCallback)
	xhtml_abbr_typeFormCallback.probe = probe
	xhtml_abbr_typeFormCallback.xhtml_abbr_type = xhtml_abbr_type
	xhtml_abbr_typeFormCallback.formGroup = formGroup

	xhtml_abbr_typeFormCallback.CreationMode = (xhtml_abbr_type == nil)

	return
}

type Xhtml_abbr_typeFormCallback struct {
	xhtml_abbr_type *models.Xhtml_abbr_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_abbr_typeFormCallback *Xhtml_abbr_typeFormCallback) OnSave() {

	log.Println("Xhtml_abbr_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_abbr_typeFormCallback.probe.formStage.Checkout()

	if xhtml_abbr_typeFormCallback.xhtml_abbr_type == nil {
		xhtml_abbr_typeFormCallback.xhtml_abbr_type = new(models.Xhtml_abbr_type).Stage(xhtml_abbr_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_abbr_type_ := xhtml_abbr_typeFormCallback.xhtml_abbr_type
	_ = xhtml_abbr_type_

	for _, formDiv := range xhtml_abbr_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_abbr_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_abbr_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_abbr_type_.Unstage(xhtml_abbr_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_abbr_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_abbr_type](
		xhtml_abbr_typeFormCallback.probe,
	)
	xhtml_abbr_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_abbr_typeFormCallback.CreationMode || xhtml_abbr_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_abbr_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_abbr_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_abbr_typeFormCallback(
			nil,
			xhtml_abbr_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_abbr_type := new(models.Xhtml_abbr_type)
		FillUpForm(xhtml_abbr_type, newFormGroup, xhtml_abbr_typeFormCallback.probe)
		xhtml_abbr_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_abbr_typeFormCallback.probe)
}
func __gong__New__Xhtml_acronym_typeFormCallback(
	xhtml_acronym_type *models.Xhtml_acronym_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_acronym_typeFormCallback *Xhtml_acronym_typeFormCallback) {
	xhtml_acronym_typeFormCallback = new(Xhtml_acronym_typeFormCallback)
	xhtml_acronym_typeFormCallback.probe = probe
	xhtml_acronym_typeFormCallback.xhtml_acronym_type = xhtml_acronym_type
	xhtml_acronym_typeFormCallback.formGroup = formGroup

	xhtml_acronym_typeFormCallback.CreationMode = (xhtml_acronym_type == nil)

	return
}

type Xhtml_acronym_typeFormCallback struct {
	xhtml_acronym_type *models.Xhtml_acronym_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_acronym_typeFormCallback *Xhtml_acronym_typeFormCallback) OnSave() {

	log.Println("Xhtml_acronym_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_acronym_typeFormCallback.probe.formStage.Checkout()

	if xhtml_acronym_typeFormCallback.xhtml_acronym_type == nil {
		xhtml_acronym_typeFormCallback.xhtml_acronym_type = new(models.Xhtml_acronym_type).Stage(xhtml_acronym_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_acronym_type_ := xhtml_acronym_typeFormCallback.xhtml_acronym_type
	_ = xhtml_acronym_type_

	for _, formDiv := range xhtml_acronym_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_acronym_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_acronym_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_acronym_type_.Unstage(xhtml_acronym_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_acronym_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_acronym_type](
		xhtml_acronym_typeFormCallback.probe,
	)
	xhtml_acronym_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_acronym_typeFormCallback.CreationMode || xhtml_acronym_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_acronym_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_acronym_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_acronym_typeFormCallback(
			nil,
			xhtml_acronym_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_acronym_type := new(models.Xhtml_acronym_type)
		FillUpForm(xhtml_acronym_type, newFormGroup, xhtml_acronym_typeFormCallback.probe)
		xhtml_acronym_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_acronym_typeFormCallback.probe)
}
func __gong__New__Xhtml_address_typeFormCallback(
	xhtml_address_type *models.Xhtml_address_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_address_typeFormCallback *Xhtml_address_typeFormCallback) {
	xhtml_address_typeFormCallback = new(Xhtml_address_typeFormCallback)
	xhtml_address_typeFormCallback.probe = probe
	xhtml_address_typeFormCallback.xhtml_address_type = xhtml_address_type
	xhtml_address_typeFormCallback.formGroup = formGroup

	xhtml_address_typeFormCallback.CreationMode = (xhtml_address_type == nil)

	return
}

type Xhtml_address_typeFormCallback struct {
	xhtml_address_type *models.Xhtml_address_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_address_typeFormCallback *Xhtml_address_typeFormCallback) OnSave() {

	log.Println("Xhtml_address_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_address_typeFormCallback.probe.formStage.Checkout()

	if xhtml_address_typeFormCallback.xhtml_address_type == nil {
		xhtml_address_typeFormCallback.xhtml_address_type = new(models.Xhtml_address_type).Stage(xhtml_address_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_address_type_ := xhtml_address_typeFormCallback.xhtml_address_type
	_ = xhtml_address_type_

	for _, formDiv := range xhtml_address_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_address_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_address_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_address_type_.Unstage(xhtml_address_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_address_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_address_type](
		xhtml_address_typeFormCallback.probe,
	)
	xhtml_address_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_address_typeFormCallback.CreationMode || xhtml_address_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_address_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_address_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_address_typeFormCallback(
			nil,
			xhtml_address_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_address_type := new(models.Xhtml_address_type)
		FillUpForm(xhtml_address_type, newFormGroup, xhtml_address_typeFormCallback.probe)
		xhtml_address_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_address_typeFormCallback.probe)
}
func __gong__New__Xhtml_blockquote_typeFormCallback(
	xhtml_blockquote_type *models.Xhtml_blockquote_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_blockquote_typeFormCallback *Xhtml_blockquote_typeFormCallback) {
	xhtml_blockquote_typeFormCallback = new(Xhtml_blockquote_typeFormCallback)
	xhtml_blockquote_typeFormCallback.probe = probe
	xhtml_blockquote_typeFormCallback.xhtml_blockquote_type = xhtml_blockquote_type
	xhtml_blockquote_typeFormCallback.formGroup = formGroup

	xhtml_blockquote_typeFormCallback.CreationMode = (xhtml_blockquote_type == nil)

	return
}

type Xhtml_blockquote_typeFormCallback struct {
	xhtml_blockquote_type *models.Xhtml_blockquote_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_blockquote_typeFormCallback *Xhtml_blockquote_typeFormCallback) OnSave() {

	log.Println("Xhtml_blockquote_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_blockquote_typeFormCallback.probe.formStage.Checkout()

	if xhtml_blockquote_typeFormCallback.xhtml_blockquote_type == nil {
		xhtml_blockquote_typeFormCallback.xhtml_blockquote_type = new(models.Xhtml_blockquote_type).Stage(xhtml_blockquote_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_blockquote_type_ := xhtml_blockquote_typeFormCallback.xhtml_blockquote_type
	_ = xhtml_blockquote_type_

	for _, formDiv := range xhtml_blockquote_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_blockquote_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_blockquote_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_blockquote_type_.Unstage(xhtml_blockquote_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_blockquote_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_blockquote_type](
		xhtml_blockquote_typeFormCallback.probe,
	)
	xhtml_blockquote_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_blockquote_typeFormCallback.CreationMode || xhtml_blockquote_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_blockquote_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_blockquote_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_blockquote_typeFormCallback(
			nil,
			xhtml_blockquote_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_blockquote_type := new(models.Xhtml_blockquote_type)
		FillUpForm(xhtml_blockquote_type, newFormGroup, xhtml_blockquote_typeFormCallback.probe)
		xhtml_blockquote_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_blockquote_typeFormCallback.probe)
}
func __gong__New__Xhtml_br_typeFormCallback(
	xhtml_br_type *models.Xhtml_br_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_br_typeFormCallback *Xhtml_br_typeFormCallback) {
	xhtml_br_typeFormCallback = new(Xhtml_br_typeFormCallback)
	xhtml_br_typeFormCallback.probe = probe
	xhtml_br_typeFormCallback.xhtml_br_type = xhtml_br_type
	xhtml_br_typeFormCallback.formGroup = formGroup

	xhtml_br_typeFormCallback.CreationMode = (xhtml_br_type == nil)

	return
}

type Xhtml_br_typeFormCallback struct {
	xhtml_br_type *models.Xhtml_br_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_br_typeFormCallback *Xhtml_br_typeFormCallback) OnSave() {

	log.Println("Xhtml_br_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_br_typeFormCallback.probe.formStage.Checkout()

	if xhtml_br_typeFormCallback.xhtml_br_type == nil {
		xhtml_br_typeFormCallback.xhtml_br_type = new(models.Xhtml_br_type).Stage(xhtml_br_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_br_type_ := xhtml_br_typeFormCallback.xhtml_br_type
	_ = xhtml_br_type_

	for _, formDiv := range xhtml_br_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_br_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_br_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_br_type_.Unstage(xhtml_br_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_br_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_br_type](
		xhtml_br_typeFormCallback.probe,
	)
	xhtml_br_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_br_typeFormCallback.CreationMode || xhtml_br_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_br_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_br_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_br_typeFormCallback(
			nil,
			xhtml_br_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_br_type := new(models.Xhtml_br_type)
		FillUpForm(xhtml_br_type, newFormGroup, xhtml_br_typeFormCallback.probe)
		xhtml_br_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_br_typeFormCallback.probe)
}
func __gong__New__Xhtml_caption_typeFormCallback(
	xhtml_caption_type *models.Xhtml_caption_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_caption_typeFormCallback *Xhtml_caption_typeFormCallback) {
	xhtml_caption_typeFormCallback = new(Xhtml_caption_typeFormCallback)
	xhtml_caption_typeFormCallback.probe = probe
	xhtml_caption_typeFormCallback.xhtml_caption_type = xhtml_caption_type
	xhtml_caption_typeFormCallback.formGroup = formGroup

	xhtml_caption_typeFormCallback.CreationMode = (xhtml_caption_type == nil)

	return
}

type Xhtml_caption_typeFormCallback struct {
	xhtml_caption_type *models.Xhtml_caption_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_caption_typeFormCallback *Xhtml_caption_typeFormCallback) OnSave() {

	log.Println("Xhtml_caption_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_caption_typeFormCallback.probe.formStage.Checkout()

	if xhtml_caption_typeFormCallback.xhtml_caption_type == nil {
		xhtml_caption_typeFormCallback.xhtml_caption_type = new(models.Xhtml_caption_type).Stage(xhtml_caption_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_caption_type_ := xhtml_caption_typeFormCallback.xhtml_caption_type
	_ = xhtml_caption_type_

	for _, formDiv := range xhtml_caption_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_caption_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_caption_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_caption_type_.Unstage(xhtml_caption_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_caption_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_caption_type](
		xhtml_caption_typeFormCallback.probe,
	)
	xhtml_caption_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_caption_typeFormCallback.CreationMode || xhtml_caption_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_caption_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_caption_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_caption_typeFormCallback(
			nil,
			xhtml_caption_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_caption_type := new(models.Xhtml_caption_type)
		FillUpForm(xhtml_caption_type, newFormGroup, xhtml_caption_typeFormCallback.probe)
		xhtml_caption_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_caption_typeFormCallback.probe)
}
func __gong__New__Xhtml_cite_typeFormCallback(
	xhtml_cite_type *models.Xhtml_cite_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_cite_typeFormCallback *Xhtml_cite_typeFormCallback) {
	xhtml_cite_typeFormCallback = new(Xhtml_cite_typeFormCallback)
	xhtml_cite_typeFormCallback.probe = probe
	xhtml_cite_typeFormCallback.xhtml_cite_type = xhtml_cite_type
	xhtml_cite_typeFormCallback.formGroup = formGroup

	xhtml_cite_typeFormCallback.CreationMode = (xhtml_cite_type == nil)

	return
}

type Xhtml_cite_typeFormCallback struct {
	xhtml_cite_type *models.Xhtml_cite_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_cite_typeFormCallback *Xhtml_cite_typeFormCallback) OnSave() {

	log.Println("Xhtml_cite_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_cite_typeFormCallback.probe.formStage.Checkout()

	if xhtml_cite_typeFormCallback.xhtml_cite_type == nil {
		xhtml_cite_typeFormCallback.xhtml_cite_type = new(models.Xhtml_cite_type).Stage(xhtml_cite_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_cite_type_ := xhtml_cite_typeFormCallback.xhtml_cite_type
	_ = xhtml_cite_type_

	for _, formDiv := range xhtml_cite_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_cite_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_cite_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_cite_type_.Unstage(xhtml_cite_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_cite_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_cite_type](
		xhtml_cite_typeFormCallback.probe,
	)
	xhtml_cite_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_cite_typeFormCallback.CreationMode || xhtml_cite_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_cite_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_cite_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_cite_typeFormCallback(
			nil,
			xhtml_cite_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_cite_type := new(models.Xhtml_cite_type)
		FillUpForm(xhtml_cite_type, newFormGroup, xhtml_cite_typeFormCallback.probe)
		xhtml_cite_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_cite_typeFormCallback.probe)
}
func __gong__New__Xhtml_code_typeFormCallback(
	xhtml_code_type *models.Xhtml_code_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_code_typeFormCallback *Xhtml_code_typeFormCallback) {
	xhtml_code_typeFormCallback = new(Xhtml_code_typeFormCallback)
	xhtml_code_typeFormCallback.probe = probe
	xhtml_code_typeFormCallback.xhtml_code_type = xhtml_code_type
	xhtml_code_typeFormCallback.formGroup = formGroup

	xhtml_code_typeFormCallback.CreationMode = (xhtml_code_type == nil)

	return
}

type Xhtml_code_typeFormCallback struct {
	xhtml_code_type *models.Xhtml_code_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_code_typeFormCallback *Xhtml_code_typeFormCallback) OnSave() {

	log.Println("Xhtml_code_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_code_typeFormCallback.probe.formStage.Checkout()

	if xhtml_code_typeFormCallback.xhtml_code_type == nil {
		xhtml_code_typeFormCallback.xhtml_code_type = new(models.Xhtml_code_type).Stage(xhtml_code_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_code_type_ := xhtml_code_typeFormCallback.xhtml_code_type
	_ = xhtml_code_type_

	for _, formDiv := range xhtml_code_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_code_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_code_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_code_type_.Unstage(xhtml_code_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_code_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_code_type](
		xhtml_code_typeFormCallback.probe,
	)
	xhtml_code_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_code_typeFormCallback.CreationMode || xhtml_code_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_code_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_code_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_code_typeFormCallback(
			nil,
			xhtml_code_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_code_type := new(models.Xhtml_code_type)
		FillUpForm(xhtml_code_type, newFormGroup, xhtml_code_typeFormCallback.probe)
		xhtml_code_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_code_typeFormCallback.probe)
}
func __gong__New__Xhtml_col_typeFormCallback(
	xhtml_col_type *models.Xhtml_col_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_col_typeFormCallback *Xhtml_col_typeFormCallback) {
	xhtml_col_typeFormCallback = new(Xhtml_col_typeFormCallback)
	xhtml_col_typeFormCallback.probe = probe
	xhtml_col_typeFormCallback.xhtml_col_type = xhtml_col_type
	xhtml_col_typeFormCallback.formGroup = formGroup

	xhtml_col_typeFormCallback.CreationMode = (xhtml_col_type == nil)

	return
}

type Xhtml_col_typeFormCallback struct {
	xhtml_col_type *models.Xhtml_col_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_col_typeFormCallback *Xhtml_col_typeFormCallback) OnSave() {

	log.Println("Xhtml_col_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_col_typeFormCallback.probe.formStage.Checkout()

	if xhtml_col_typeFormCallback.xhtml_col_type == nil {
		xhtml_col_typeFormCallback.xhtml_col_type = new(models.Xhtml_col_type).Stage(xhtml_col_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_col_type_ := xhtml_col_typeFormCallback.xhtml_col_type
	_ = xhtml_col_type_

	for _, formDiv := range xhtml_col_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_col_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_col_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_col_type_.Unstage(xhtml_col_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_col_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_col_type](
		xhtml_col_typeFormCallback.probe,
	)
	xhtml_col_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_col_typeFormCallback.CreationMode || xhtml_col_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_col_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_col_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_col_typeFormCallback(
			nil,
			xhtml_col_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_col_type := new(models.Xhtml_col_type)
		FillUpForm(xhtml_col_type, newFormGroup, xhtml_col_typeFormCallback.probe)
		xhtml_col_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_col_typeFormCallback.probe)
}
func __gong__New__Xhtml_colgroup_typeFormCallback(
	xhtml_colgroup_type *models.Xhtml_colgroup_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_colgroup_typeFormCallback *Xhtml_colgroup_typeFormCallback) {
	xhtml_colgroup_typeFormCallback = new(Xhtml_colgroup_typeFormCallback)
	xhtml_colgroup_typeFormCallback.probe = probe
	xhtml_colgroup_typeFormCallback.xhtml_colgroup_type = xhtml_colgroup_type
	xhtml_colgroup_typeFormCallback.formGroup = formGroup

	xhtml_colgroup_typeFormCallback.CreationMode = (xhtml_colgroup_type == nil)

	return
}

type Xhtml_colgroup_typeFormCallback struct {
	xhtml_colgroup_type *models.Xhtml_colgroup_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_colgroup_typeFormCallback *Xhtml_colgroup_typeFormCallback) OnSave() {

	log.Println("Xhtml_colgroup_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_colgroup_typeFormCallback.probe.formStage.Checkout()

	if xhtml_colgroup_typeFormCallback.xhtml_colgroup_type == nil {
		xhtml_colgroup_typeFormCallback.xhtml_colgroup_type = new(models.Xhtml_colgroup_type).Stage(xhtml_colgroup_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_colgroup_type_ := xhtml_colgroup_typeFormCallback.xhtml_colgroup_type
	_ = xhtml_colgroup_type_

	for _, formDiv := range xhtml_colgroup_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_colgroup_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_colgroup_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_colgroup_type_.Unstage(xhtml_colgroup_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_colgroup_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_colgroup_type](
		xhtml_colgroup_typeFormCallback.probe,
	)
	xhtml_colgroup_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_colgroup_typeFormCallback.CreationMode || xhtml_colgroup_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_colgroup_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_colgroup_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_colgroup_typeFormCallback(
			nil,
			xhtml_colgroup_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_colgroup_type := new(models.Xhtml_colgroup_type)
		FillUpForm(xhtml_colgroup_type, newFormGroup, xhtml_colgroup_typeFormCallback.probe)
		xhtml_colgroup_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_colgroup_typeFormCallback.probe)
}
func __gong__New__Xhtml_dd_typeFormCallback(
	xhtml_dd_type *models.Xhtml_dd_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_dd_typeFormCallback *Xhtml_dd_typeFormCallback) {
	xhtml_dd_typeFormCallback = new(Xhtml_dd_typeFormCallback)
	xhtml_dd_typeFormCallback.probe = probe
	xhtml_dd_typeFormCallback.xhtml_dd_type = xhtml_dd_type
	xhtml_dd_typeFormCallback.formGroup = formGroup

	xhtml_dd_typeFormCallback.CreationMode = (xhtml_dd_type == nil)

	return
}

type Xhtml_dd_typeFormCallback struct {
	xhtml_dd_type *models.Xhtml_dd_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_dd_typeFormCallback *Xhtml_dd_typeFormCallback) OnSave() {

	log.Println("Xhtml_dd_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_dd_typeFormCallback.probe.formStage.Checkout()

	if xhtml_dd_typeFormCallback.xhtml_dd_type == nil {
		xhtml_dd_typeFormCallback.xhtml_dd_type = new(models.Xhtml_dd_type).Stage(xhtml_dd_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_dd_type_ := xhtml_dd_typeFormCallback.xhtml_dd_type
	_ = xhtml_dd_type_

	for _, formDiv := range xhtml_dd_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_dd_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_dd_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_dd_type_.Unstage(xhtml_dd_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_dd_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_dd_type](
		xhtml_dd_typeFormCallback.probe,
	)
	xhtml_dd_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_dd_typeFormCallback.CreationMode || xhtml_dd_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_dd_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_dd_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_dd_typeFormCallback(
			nil,
			xhtml_dd_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_dd_type := new(models.Xhtml_dd_type)
		FillUpForm(xhtml_dd_type, newFormGroup, xhtml_dd_typeFormCallback.probe)
		xhtml_dd_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_dd_typeFormCallback.probe)
}
func __gong__New__Xhtml_dfn_typeFormCallback(
	xhtml_dfn_type *models.Xhtml_dfn_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_dfn_typeFormCallback *Xhtml_dfn_typeFormCallback) {
	xhtml_dfn_typeFormCallback = new(Xhtml_dfn_typeFormCallback)
	xhtml_dfn_typeFormCallback.probe = probe
	xhtml_dfn_typeFormCallback.xhtml_dfn_type = xhtml_dfn_type
	xhtml_dfn_typeFormCallback.formGroup = formGroup

	xhtml_dfn_typeFormCallback.CreationMode = (xhtml_dfn_type == nil)

	return
}

type Xhtml_dfn_typeFormCallback struct {
	xhtml_dfn_type *models.Xhtml_dfn_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_dfn_typeFormCallback *Xhtml_dfn_typeFormCallback) OnSave() {

	log.Println("Xhtml_dfn_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_dfn_typeFormCallback.probe.formStage.Checkout()

	if xhtml_dfn_typeFormCallback.xhtml_dfn_type == nil {
		xhtml_dfn_typeFormCallback.xhtml_dfn_type = new(models.Xhtml_dfn_type).Stage(xhtml_dfn_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_dfn_type_ := xhtml_dfn_typeFormCallback.xhtml_dfn_type
	_ = xhtml_dfn_type_

	for _, formDiv := range xhtml_dfn_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_dfn_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_dfn_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_dfn_type_.Unstage(xhtml_dfn_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_dfn_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_dfn_type](
		xhtml_dfn_typeFormCallback.probe,
	)
	xhtml_dfn_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_dfn_typeFormCallback.CreationMode || xhtml_dfn_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_dfn_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_dfn_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_dfn_typeFormCallback(
			nil,
			xhtml_dfn_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_dfn_type := new(models.Xhtml_dfn_type)
		FillUpForm(xhtml_dfn_type, newFormGroup, xhtml_dfn_typeFormCallback.probe)
		xhtml_dfn_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_dfn_typeFormCallback.probe)
}
func __gong__New__Xhtml_div_typeFormCallback(
	xhtml_div_type *models.Xhtml_div_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_div_typeFormCallback *Xhtml_div_typeFormCallback) {
	xhtml_div_typeFormCallback = new(Xhtml_div_typeFormCallback)
	xhtml_div_typeFormCallback.probe = probe
	xhtml_div_typeFormCallback.xhtml_div_type = xhtml_div_type
	xhtml_div_typeFormCallback.formGroup = formGroup

	xhtml_div_typeFormCallback.CreationMode = (xhtml_div_type == nil)

	return
}

type Xhtml_div_typeFormCallback struct {
	xhtml_div_type *models.Xhtml_div_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_div_typeFormCallback *Xhtml_div_typeFormCallback) OnSave() {

	log.Println("Xhtml_div_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_div_typeFormCallback.probe.formStage.Checkout()

	if xhtml_div_typeFormCallback.xhtml_div_type == nil {
		xhtml_div_typeFormCallback.xhtml_div_type = new(models.Xhtml_div_type).Stage(xhtml_div_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_div_type_ := xhtml_div_typeFormCallback.xhtml_div_type
	_ = xhtml_div_type_

	for _, formDiv := range xhtml_div_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_div_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_div_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_div_type_.Unstage(xhtml_div_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_div_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_div_type](
		xhtml_div_typeFormCallback.probe,
	)
	xhtml_div_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_div_typeFormCallback.CreationMode || xhtml_div_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_div_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_div_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_div_typeFormCallback(
			nil,
			xhtml_div_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_div_type := new(models.Xhtml_div_type)
		FillUpForm(xhtml_div_type, newFormGroup, xhtml_div_typeFormCallback.probe)
		xhtml_div_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_div_typeFormCallback.probe)
}
func __gong__New__Xhtml_dl_typeFormCallback(
	xhtml_dl_type *models.Xhtml_dl_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_dl_typeFormCallback *Xhtml_dl_typeFormCallback) {
	xhtml_dl_typeFormCallback = new(Xhtml_dl_typeFormCallback)
	xhtml_dl_typeFormCallback.probe = probe
	xhtml_dl_typeFormCallback.xhtml_dl_type = xhtml_dl_type
	xhtml_dl_typeFormCallback.formGroup = formGroup

	xhtml_dl_typeFormCallback.CreationMode = (xhtml_dl_type == nil)

	return
}

type Xhtml_dl_typeFormCallback struct {
	xhtml_dl_type *models.Xhtml_dl_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_dl_typeFormCallback *Xhtml_dl_typeFormCallback) OnSave() {

	log.Println("Xhtml_dl_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_dl_typeFormCallback.probe.formStage.Checkout()

	if xhtml_dl_typeFormCallback.xhtml_dl_type == nil {
		xhtml_dl_typeFormCallback.xhtml_dl_type = new(models.Xhtml_dl_type).Stage(xhtml_dl_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_dl_type_ := xhtml_dl_typeFormCallback.xhtml_dl_type
	_ = xhtml_dl_type_

	for _, formDiv := range xhtml_dl_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_dl_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_dl_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_dl_type_.Unstage(xhtml_dl_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_dl_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_dl_type](
		xhtml_dl_typeFormCallback.probe,
	)
	xhtml_dl_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_dl_typeFormCallback.CreationMode || xhtml_dl_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_dl_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_dl_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_dl_typeFormCallback(
			nil,
			xhtml_dl_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_dl_type := new(models.Xhtml_dl_type)
		FillUpForm(xhtml_dl_type, newFormGroup, xhtml_dl_typeFormCallback.probe)
		xhtml_dl_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_dl_typeFormCallback.probe)
}
func __gong__New__Xhtml_dt_typeFormCallback(
	xhtml_dt_type *models.Xhtml_dt_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_dt_typeFormCallback *Xhtml_dt_typeFormCallback) {
	xhtml_dt_typeFormCallback = new(Xhtml_dt_typeFormCallback)
	xhtml_dt_typeFormCallback.probe = probe
	xhtml_dt_typeFormCallback.xhtml_dt_type = xhtml_dt_type
	xhtml_dt_typeFormCallback.formGroup = formGroup

	xhtml_dt_typeFormCallback.CreationMode = (xhtml_dt_type == nil)

	return
}

type Xhtml_dt_typeFormCallback struct {
	xhtml_dt_type *models.Xhtml_dt_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_dt_typeFormCallback *Xhtml_dt_typeFormCallback) OnSave() {

	log.Println("Xhtml_dt_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_dt_typeFormCallback.probe.formStage.Checkout()

	if xhtml_dt_typeFormCallback.xhtml_dt_type == nil {
		xhtml_dt_typeFormCallback.xhtml_dt_type = new(models.Xhtml_dt_type).Stage(xhtml_dt_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_dt_type_ := xhtml_dt_typeFormCallback.xhtml_dt_type
	_ = xhtml_dt_type_

	for _, formDiv := range xhtml_dt_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_dt_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_dt_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_dt_type_.Unstage(xhtml_dt_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_dt_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_dt_type](
		xhtml_dt_typeFormCallback.probe,
	)
	xhtml_dt_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_dt_typeFormCallback.CreationMode || xhtml_dt_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_dt_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_dt_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_dt_typeFormCallback(
			nil,
			xhtml_dt_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_dt_type := new(models.Xhtml_dt_type)
		FillUpForm(xhtml_dt_type, newFormGroup, xhtml_dt_typeFormCallback.probe)
		xhtml_dt_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_dt_typeFormCallback.probe)
}
func __gong__New__Xhtml_edit_typeFormCallback(
	xhtml_edit_type *models.Xhtml_edit_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_edit_typeFormCallback *Xhtml_edit_typeFormCallback) {
	xhtml_edit_typeFormCallback = new(Xhtml_edit_typeFormCallback)
	xhtml_edit_typeFormCallback.probe = probe
	xhtml_edit_typeFormCallback.xhtml_edit_type = xhtml_edit_type
	xhtml_edit_typeFormCallback.formGroup = formGroup

	xhtml_edit_typeFormCallback.CreationMode = (xhtml_edit_type == nil)

	return
}

type Xhtml_edit_typeFormCallback struct {
	xhtml_edit_type *models.Xhtml_edit_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_edit_typeFormCallback *Xhtml_edit_typeFormCallback) OnSave() {

	log.Println("Xhtml_edit_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_edit_typeFormCallback.probe.formStage.Checkout()

	if xhtml_edit_typeFormCallback.xhtml_edit_type == nil {
		xhtml_edit_typeFormCallback.xhtml_edit_type = new(models.Xhtml_edit_type).Stage(xhtml_edit_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_edit_type_ := xhtml_edit_typeFormCallback.xhtml_edit_type
	_ = xhtml_edit_type_

	for _, formDiv := range xhtml_edit_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_edit_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_edit_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_edit_type_.Unstage(xhtml_edit_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_edit_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_edit_type](
		xhtml_edit_typeFormCallback.probe,
	)
	xhtml_edit_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_edit_typeFormCallback.CreationMode || xhtml_edit_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_edit_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_edit_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_edit_typeFormCallback(
			nil,
			xhtml_edit_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_edit_type := new(models.Xhtml_edit_type)
		FillUpForm(xhtml_edit_type, newFormGroup, xhtml_edit_typeFormCallback.probe)
		xhtml_edit_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_edit_typeFormCallback.probe)
}
func __gong__New__Xhtml_em_typeFormCallback(
	xhtml_em_type *models.Xhtml_em_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_em_typeFormCallback *Xhtml_em_typeFormCallback) {
	xhtml_em_typeFormCallback = new(Xhtml_em_typeFormCallback)
	xhtml_em_typeFormCallback.probe = probe
	xhtml_em_typeFormCallback.xhtml_em_type = xhtml_em_type
	xhtml_em_typeFormCallback.formGroup = formGroup

	xhtml_em_typeFormCallback.CreationMode = (xhtml_em_type == nil)

	return
}

type Xhtml_em_typeFormCallback struct {
	xhtml_em_type *models.Xhtml_em_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_em_typeFormCallback *Xhtml_em_typeFormCallback) OnSave() {

	log.Println("Xhtml_em_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_em_typeFormCallback.probe.formStage.Checkout()

	if xhtml_em_typeFormCallback.xhtml_em_type == nil {
		xhtml_em_typeFormCallback.xhtml_em_type = new(models.Xhtml_em_type).Stage(xhtml_em_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_em_type_ := xhtml_em_typeFormCallback.xhtml_em_type
	_ = xhtml_em_type_

	for _, formDiv := range xhtml_em_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_em_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_em_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_em_type_.Unstage(xhtml_em_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_em_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_em_type](
		xhtml_em_typeFormCallback.probe,
	)
	xhtml_em_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_em_typeFormCallback.CreationMode || xhtml_em_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_em_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_em_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_em_typeFormCallback(
			nil,
			xhtml_em_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_em_type := new(models.Xhtml_em_type)
		FillUpForm(xhtml_em_type, newFormGroup, xhtml_em_typeFormCallback.probe)
		xhtml_em_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_em_typeFormCallback.probe)
}
func __gong__New__Xhtml_h1_typeFormCallback(
	xhtml_h1_type *models.Xhtml_h1_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_h1_typeFormCallback *Xhtml_h1_typeFormCallback) {
	xhtml_h1_typeFormCallback = new(Xhtml_h1_typeFormCallback)
	xhtml_h1_typeFormCallback.probe = probe
	xhtml_h1_typeFormCallback.xhtml_h1_type = xhtml_h1_type
	xhtml_h1_typeFormCallback.formGroup = formGroup

	xhtml_h1_typeFormCallback.CreationMode = (xhtml_h1_type == nil)

	return
}

type Xhtml_h1_typeFormCallback struct {
	xhtml_h1_type *models.Xhtml_h1_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_h1_typeFormCallback *Xhtml_h1_typeFormCallback) OnSave() {

	log.Println("Xhtml_h1_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_h1_typeFormCallback.probe.formStage.Checkout()

	if xhtml_h1_typeFormCallback.xhtml_h1_type == nil {
		xhtml_h1_typeFormCallback.xhtml_h1_type = new(models.Xhtml_h1_type).Stage(xhtml_h1_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_h1_type_ := xhtml_h1_typeFormCallback.xhtml_h1_type
	_ = xhtml_h1_type_

	for _, formDiv := range xhtml_h1_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_h1_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_h1_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h1_type_.Unstage(xhtml_h1_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_h1_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_h1_type](
		xhtml_h1_typeFormCallback.probe,
	)
	xhtml_h1_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_h1_typeFormCallback.CreationMode || xhtml_h1_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h1_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_h1_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_h1_typeFormCallback(
			nil,
			xhtml_h1_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_h1_type := new(models.Xhtml_h1_type)
		FillUpForm(xhtml_h1_type, newFormGroup, xhtml_h1_typeFormCallback.probe)
		xhtml_h1_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_h1_typeFormCallback.probe)
}
func __gong__New__Xhtml_h2_typeFormCallback(
	xhtml_h2_type *models.Xhtml_h2_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_h2_typeFormCallback *Xhtml_h2_typeFormCallback) {
	xhtml_h2_typeFormCallback = new(Xhtml_h2_typeFormCallback)
	xhtml_h2_typeFormCallback.probe = probe
	xhtml_h2_typeFormCallback.xhtml_h2_type = xhtml_h2_type
	xhtml_h2_typeFormCallback.formGroup = formGroup

	xhtml_h2_typeFormCallback.CreationMode = (xhtml_h2_type == nil)

	return
}

type Xhtml_h2_typeFormCallback struct {
	xhtml_h2_type *models.Xhtml_h2_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_h2_typeFormCallback *Xhtml_h2_typeFormCallback) OnSave() {

	log.Println("Xhtml_h2_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_h2_typeFormCallback.probe.formStage.Checkout()

	if xhtml_h2_typeFormCallback.xhtml_h2_type == nil {
		xhtml_h2_typeFormCallback.xhtml_h2_type = new(models.Xhtml_h2_type).Stage(xhtml_h2_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_h2_type_ := xhtml_h2_typeFormCallback.xhtml_h2_type
	_ = xhtml_h2_type_

	for _, formDiv := range xhtml_h2_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_h2_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_h2_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h2_type_.Unstage(xhtml_h2_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_h2_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_h2_type](
		xhtml_h2_typeFormCallback.probe,
	)
	xhtml_h2_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_h2_typeFormCallback.CreationMode || xhtml_h2_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h2_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_h2_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_h2_typeFormCallback(
			nil,
			xhtml_h2_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_h2_type := new(models.Xhtml_h2_type)
		FillUpForm(xhtml_h2_type, newFormGroup, xhtml_h2_typeFormCallback.probe)
		xhtml_h2_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_h2_typeFormCallback.probe)
}
func __gong__New__Xhtml_h3_typeFormCallback(
	xhtml_h3_type *models.Xhtml_h3_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_h3_typeFormCallback *Xhtml_h3_typeFormCallback) {
	xhtml_h3_typeFormCallback = new(Xhtml_h3_typeFormCallback)
	xhtml_h3_typeFormCallback.probe = probe
	xhtml_h3_typeFormCallback.xhtml_h3_type = xhtml_h3_type
	xhtml_h3_typeFormCallback.formGroup = formGroup

	xhtml_h3_typeFormCallback.CreationMode = (xhtml_h3_type == nil)

	return
}

type Xhtml_h3_typeFormCallback struct {
	xhtml_h3_type *models.Xhtml_h3_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_h3_typeFormCallback *Xhtml_h3_typeFormCallback) OnSave() {

	log.Println("Xhtml_h3_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_h3_typeFormCallback.probe.formStage.Checkout()

	if xhtml_h3_typeFormCallback.xhtml_h3_type == nil {
		xhtml_h3_typeFormCallback.xhtml_h3_type = new(models.Xhtml_h3_type).Stage(xhtml_h3_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_h3_type_ := xhtml_h3_typeFormCallback.xhtml_h3_type
	_ = xhtml_h3_type_

	for _, formDiv := range xhtml_h3_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_h3_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_h3_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h3_type_.Unstage(xhtml_h3_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_h3_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_h3_type](
		xhtml_h3_typeFormCallback.probe,
	)
	xhtml_h3_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_h3_typeFormCallback.CreationMode || xhtml_h3_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h3_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_h3_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_h3_typeFormCallback(
			nil,
			xhtml_h3_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_h3_type := new(models.Xhtml_h3_type)
		FillUpForm(xhtml_h3_type, newFormGroup, xhtml_h3_typeFormCallback.probe)
		xhtml_h3_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_h3_typeFormCallback.probe)
}
func __gong__New__Xhtml_h4_typeFormCallback(
	xhtml_h4_type *models.Xhtml_h4_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_h4_typeFormCallback *Xhtml_h4_typeFormCallback) {
	xhtml_h4_typeFormCallback = new(Xhtml_h4_typeFormCallback)
	xhtml_h4_typeFormCallback.probe = probe
	xhtml_h4_typeFormCallback.xhtml_h4_type = xhtml_h4_type
	xhtml_h4_typeFormCallback.formGroup = formGroup

	xhtml_h4_typeFormCallback.CreationMode = (xhtml_h4_type == nil)

	return
}

type Xhtml_h4_typeFormCallback struct {
	xhtml_h4_type *models.Xhtml_h4_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_h4_typeFormCallback *Xhtml_h4_typeFormCallback) OnSave() {

	log.Println("Xhtml_h4_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_h4_typeFormCallback.probe.formStage.Checkout()

	if xhtml_h4_typeFormCallback.xhtml_h4_type == nil {
		xhtml_h4_typeFormCallback.xhtml_h4_type = new(models.Xhtml_h4_type).Stage(xhtml_h4_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_h4_type_ := xhtml_h4_typeFormCallback.xhtml_h4_type
	_ = xhtml_h4_type_

	for _, formDiv := range xhtml_h4_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_h4_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_h4_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h4_type_.Unstage(xhtml_h4_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_h4_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_h4_type](
		xhtml_h4_typeFormCallback.probe,
	)
	xhtml_h4_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_h4_typeFormCallback.CreationMode || xhtml_h4_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h4_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_h4_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_h4_typeFormCallback(
			nil,
			xhtml_h4_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_h4_type := new(models.Xhtml_h4_type)
		FillUpForm(xhtml_h4_type, newFormGroup, xhtml_h4_typeFormCallback.probe)
		xhtml_h4_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_h4_typeFormCallback.probe)
}
func __gong__New__Xhtml_h5_typeFormCallback(
	xhtml_h5_type *models.Xhtml_h5_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_h5_typeFormCallback *Xhtml_h5_typeFormCallback) {
	xhtml_h5_typeFormCallback = new(Xhtml_h5_typeFormCallback)
	xhtml_h5_typeFormCallback.probe = probe
	xhtml_h5_typeFormCallback.xhtml_h5_type = xhtml_h5_type
	xhtml_h5_typeFormCallback.formGroup = formGroup

	xhtml_h5_typeFormCallback.CreationMode = (xhtml_h5_type == nil)

	return
}

type Xhtml_h5_typeFormCallback struct {
	xhtml_h5_type *models.Xhtml_h5_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_h5_typeFormCallback *Xhtml_h5_typeFormCallback) OnSave() {

	log.Println("Xhtml_h5_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_h5_typeFormCallback.probe.formStage.Checkout()

	if xhtml_h5_typeFormCallback.xhtml_h5_type == nil {
		xhtml_h5_typeFormCallback.xhtml_h5_type = new(models.Xhtml_h5_type).Stage(xhtml_h5_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_h5_type_ := xhtml_h5_typeFormCallback.xhtml_h5_type
	_ = xhtml_h5_type_

	for _, formDiv := range xhtml_h5_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_h5_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_h5_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h5_type_.Unstage(xhtml_h5_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_h5_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_h5_type](
		xhtml_h5_typeFormCallback.probe,
	)
	xhtml_h5_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_h5_typeFormCallback.CreationMode || xhtml_h5_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h5_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_h5_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_h5_typeFormCallback(
			nil,
			xhtml_h5_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_h5_type := new(models.Xhtml_h5_type)
		FillUpForm(xhtml_h5_type, newFormGroup, xhtml_h5_typeFormCallback.probe)
		xhtml_h5_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_h5_typeFormCallback.probe)
}
func __gong__New__Xhtml_h6_typeFormCallback(
	xhtml_h6_type *models.Xhtml_h6_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_h6_typeFormCallback *Xhtml_h6_typeFormCallback) {
	xhtml_h6_typeFormCallback = new(Xhtml_h6_typeFormCallback)
	xhtml_h6_typeFormCallback.probe = probe
	xhtml_h6_typeFormCallback.xhtml_h6_type = xhtml_h6_type
	xhtml_h6_typeFormCallback.formGroup = formGroup

	xhtml_h6_typeFormCallback.CreationMode = (xhtml_h6_type == nil)

	return
}

type Xhtml_h6_typeFormCallback struct {
	xhtml_h6_type *models.Xhtml_h6_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_h6_typeFormCallback *Xhtml_h6_typeFormCallback) OnSave() {

	log.Println("Xhtml_h6_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_h6_typeFormCallback.probe.formStage.Checkout()

	if xhtml_h6_typeFormCallback.xhtml_h6_type == nil {
		xhtml_h6_typeFormCallback.xhtml_h6_type = new(models.Xhtml_h6_type).Stage(xhtml_h6_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_h6_type_ := xhtml_h6_typeFormCallback.xhtml_h6_type
	_ = xhtml_h6_type_

	for _, formDiv := range xhtml_h6_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_h6_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_h6_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h6_type_.Unstage(xhtml_h6_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_h6_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_h6_type](
		xhtml_h6_typeFormCallback.probe,
	)
	xhtml_h6_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_h6_typeFormCallback.CreationMode || xhtml_h6_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_h6_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_h6_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_h6_typeFormCallback(
			nil,
			xhtml_h6_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_h6_type := new(models.Xhtml_h6_type)
		FillUpForm(xhtml_h6_type, newFormGroup, xhtml_h6_typeFormCallback.probe)
		xhtml_h6_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_h6_typeFormCallback.probe)
}
func __gong__New__Xhtml_heading_typeFormCallback(
	xhtml_heading_type *models.Xhtml_heading_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_heading_typeFormCallback *Xhtml_heading_typeFormCallback) {
	xhtml_heading_typeFormCallback = new(Xhtml_heading_typeFormCallback)
	xhtml_heading_typeFormCallback.probe = probe
	xhtml_heading_typeFormCallback.xhtml_heading_type = xhtml_heading_type
	xhtml_heading_typeFormCallback.formGroup = formGroup

	xhtml_heading_typeFormCallback.CreationMode = (xhtml_heading_type == nil)

	return
}

type Xhtml_heading_typeFormCallback struct {
	xhtml_heading_type *models.Xhtml_heading_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_heading_typeFormCallback *Xhtml_heading_typeFormCallback) OnSave() {

	log.Println("Xhtml_heading_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_heading_typeFormCallback.probe.formStage.Checkout()

	if xhtml_heading_typeFormCallback.xhtml_heading_type == nil {
		xhtml_heading_typeFormCallback.xhtml_heading_type = new(models.Xhtml_heading_type).Stage(xhtml_heading_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_heading_type_ := xhtml_heading_typeFormCallback.xhtml_heading_type
	_ = xhtml_heading_type_

	for _, formDiv := range xhtml_heading_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_heading_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_heading_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_heading_type_.Unstage(xhtml_heading_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_heading_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_heading_type](
		xhtml_heading_typeFormCallback.probe,
	)
	xhtml_heading_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_heading_typeFormCallback.CreationMode || xhtml_heading_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_heading_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_heading_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_heading_typeFormCallback(
			nil,
			xhtml_heading_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_heading_type := new(models.Xhtml_heading_type)
		FillUpForm(xhtml_heading_type, newFormGroup, xhtml_heading_typeFormCallback.probe)
		xhtml_heading_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_heading_typeFormCallback.probe)
}
func __gong__New__Xhtml_hr_typeFormCallback(
	xhtml_hr_type *models.Xhtml_hr_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_hr_typeFormCallback *Xhtml_hr_typeFormCallback) {
	xhtml_hr_typeFormCallback = new(Xhtml_hr_typeFormCallback)
	xhtml_hr_typeFormCallback.probe = probe
	xhtml_hr_typeFormCallback.xhtml_hr_type = xhtml_hr_type
	xhtml_hr_typeFormCallback.formGroup = formGroup

	xhtml_hr_typeFormCallback.CreationMode = (xhtml_hr_type == nil)

	return
}

type Xhtml_hr_typeFormCallback struct {
	xhtml_hr_type *models.Xhtml_hr_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_hr_typeFormCallback *Xhtml_hr_typeFormCallback) OnSave() {

	log.Println("Xhtml_hr_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_hr_typeFormCallback.probe.formStage.Checkout()

	if xhtml_hr_typeFormCallback.xhtml_hr_type == nil {
		xhtml_hr_typeFormCallback.xhtml_hr_type = new(models.Xhtml_hr_type).Stage(xhtml_hr_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_hr_type_ := xhtml_hr_typeFormCallback.xhtml_hr_type
	_ = xhtml_hr_type_

	for _, formDiv := range xhtml_hr_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_hr_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_hr_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_hr_type_.Unstage(xhtml_hr_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_hr_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_hr_type](
		xhtml_hr_typeFormCallback.probe,
	)
	xhtml_hr_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_hr_typeFormCallback.CreationMode || xhtml_hr_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_hr_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_hr_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_hr_typeFormCallback(
			nil,
			xhtml_hr_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_hr_type := new(models.Xhtml_hr_type)
		FillUpForm(xhtml_hr_type, newFormGroup, xhtml_hr_typeFormCallback.probe)
		xhtml_hr_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_hr_typeFormCallback.probe)
}
func __gong__New__Xhtml_kbd_typeFormCallback(
	xhtml_kbd_type *models.Xhtml_kbd_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_kbd_typeFormCallback *Xhtml_kbd_typeFormCallback) {
	xhtml_kbd_typeFormCallback = new(Xhtml_kbd_typeFormCallback)
	xhtml_kbd_typeFormCallback.probe = probe
	xhtml_kbd_typeFormCallback.xhtml_kbd_type = xhtml_kbd_type
	xhtml_kbd_typeFormCallback.formGroup = formGroup

	xhtml_kbd_typeFormCallback.CreationMode = (xhtml_kbd_type == nil)

	return
}

type Xhtml_kbd_typeFormCallback struct {
	xhtml_kbd_type *models.Xhtml_kbd_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_kbd_typeFormCallback *Xhtml_kbd_typeFormCallback) OnSave() {

	log.Println("Xhtml_kbd_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_kbd_typeFormCallback.probe.formStage.Checkout()

	if xhtml_kbd_typeFormCallback.xhtml_kbd_type == nil {
		xhtml_kbd_typeFormCallback.xhtml_kbd_type = new(models.Xhtml_kbd_type).Stage(xhtml_kbd_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_kbd_type_ := xhtml_kbd_typeFormCallback.xhtml_kbd_type
	_ = xhtml_kbd_type_

	for _, formDiv := range xhtml_kbd_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_kbd_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_kbd_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_kbd_type_.Unstage(xhtml_kbd_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_kbd_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_kbd_type](
		xhtml_kbd_typeFormCallback.probe,
	)
	xhtml_kbd_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_kbd_typeFormCallback.CreationMode || xhtml_kbd_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_kbd_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_kbd_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_kbd_typeFormCallback(
			nil,
			xhtml_kbd_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_kbd_type := new(models.Xhtml_kbd_type)
		FillUpForm(xhtml_kbd_type, newFormGroup, xhtml_kbd_typeFormCallback.probe)
		xhtml_kbd_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_kbd_typeFormCallback.probe)
}
func __gong__New__Xhtml_li_typeFormCallback(
	xhtml_li_type *models.Xhtml_li_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_li_typeFormCallback *Xhtml_li_typeFormCallback) {
	xhtml_li_typeFormCallback = new(Xhtml_li_typeFormCallback)
	xhtml_li_typeFormCallback.probe = probe
	xhtml_li_typeFormCallback.xhtml_li_type = xhtml_li_type
	xhtml_li_typeFormCallback.formGroup = formGroup

	xhtml_li_typeFormCallback.CreationMode = (xhtml_li_type == nil)

	return
}

type Xhtml_li_typeFormCallback struct {
	xhtml_li_type *models.Xhtml_li_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_li_typeFormCallback *Xhtml_li_typeFormCallback) OnSave() {

	log.Println("Xhtml_li_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_li_typeFormCallback.probe.formStage.Checkout()

	if xhtml_li_typeFormCallback.xhtml_li_type == nil {
		xhtml_li_typeFormCallback.xhtml_li_type = new(models.Xhtml_li_type).Stage(xhtml_li_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_li_type_ := xhtml_li_typeFormCallback.xhtml_li_type
	_ = xhtml_li_type_

	for _, formDiv := range xhtml_li_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_li_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_li_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_li_type_.Unstage(xhtml_li_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_li_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_li_type](
		xhtml_li_typeFormCallback.probe,
	)
	xhtml_li_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_li_typeFormCallback.CreationMode || xhtml_li_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_li_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_li_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_li_typeFormCallback(
			nil,
			xhtml_li_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_li_type := new(models.Xhtml_li_type)
		FillUpForm(xhtml_li_type, newFormGroup, xhtml_li_typeFormCallback.probe)
		xhtml_li_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_li_typeFormCallback.probe)
}
func __gong__New__Xhtml_object_typeFormCallback(
	xhtml_object_type *models.Xhtml_object_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_object_typeFormCallback *Xhtml_object_typeFormCallback) {
	xhtml_object_typeFormCallback = new(Xhtml_object_typeFormCallback)
	xhtml_object_typeFormCallback.probe = probe
	xhtml_object_typeFormCallback.xhtml_object_type = xhtml_object_type
	xhtml_object_typeFormCallback.formGroup = formGroup

	xhtml_object_typeFormCallback.CreationMode = (xhtml_object_type == nil)

	return
}

type Xhtml_object_typeFormCallback struct {
	xhtml_object_type *models.Xhtml_object_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_object_typeFormCallback *Xhtml_object_typeFormCallback) OnSave() {

	log.Println("Xhtml_object_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_object_typeFormCallback.probe.formStage.Checkout()

	if xhtml_object_typeFormCallback.xhtml_object_type == nil {
		xhtml_object_typeFormCallback.xhtml_object_type = new(models.Xhtml_object_type).Stage(xhtml_object_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_object_type_ := xhtml_object_typeFormCallback.xhtml_object_type
	_ = xhtml_object_type_

	for _, formDiv := range xhtml_object_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_object_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_object_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_object_type_.Unstage(xhtml_object_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_object_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_object_type](
		xhtml_object_typeFormCallback.probe,
	)
	xhtml_object_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_object_typeFormCallback.CreationMode || xhtml_object_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_object_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_object_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_object_typeFormCallback(
			nil,
			xhtml_object_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_object_type := new(models.Xhtml_object_type)
		FillUpForm(xhtml_object_type, newFormGroup, xhtml_object_typeFormCallback.probe)
		xhtml_object_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_object_typeFormCallback.probe)
}
func __gong__New__Xhtml_ol_typeFormCallback(
	xhtml_ol_type *models.Xhtml_ol_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_ol_typeFormCallback *Xhtml_ol_typeFormCallback) {
	xhtml_ol_typeFormCallback = new(Xhtml_ol_typeFormCallback)
	xhtml_ol_typeFormCallback.probe = probe
	xhtml_ol_typeFormCallback.xhtml_ol_type = xhtml_ol_type
	xhtml_ol_typeFormCallback.formGroup = formGroup

	xhtml_ol_typeFormCallback.CreationMode = (xhtml_ol_type == nil)

	return
}

type Xhtml_ol_typeFormCallback struct {
	xhtml_ol_type *models.Xhtml_ol_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_ol_typeFormCallback *Xhtml_ol_typeFormCallback) OnSave() {

	log.Println("Xhtml_ol_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_ol_typeFormCallback.probe.formStage.Checkout()

	if xhtml_ol_typeFormCallback.xhtml_ol_type == nil {
		xhtml_ol_typeFormCallback.xhtml_ol_type = new(models.Xhtml_ol_type).Stage(xhtml_ol_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_ol_type_ := xhtml_ol_typeFormCallback.xhtml_ol_type
	_ = xhtml_ol_type_

	for _, formDiv := range xhtml_ol_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_ol_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_ol_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_ol_type_.Unstage(xhtml_ol_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_ol_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_ol_type](
		xhtml_ol_typeFormCallback.probe,
	)
	xhtml_ol_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_ol_typeFormCallback.CreationMode || xhtml_ol_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_ol_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_ol_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_ol_typeFormCallback(
			nil,
			xhtml_ol_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_ol_type := new(models.Xhtml_ol_type)
		FillUpForm(xhtml_ol_type, newFormGroup, xhtml_ol_typeFormCallback.probe)
		xhtml_ol_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_ol_typeFormCallback.probe)
}
func __gong__New__Xhtml_p_typeFormCallback(
	xhtml_p_type *models.Xhtml_p_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_p_typeFormCallback *Xhtml_p_typeFormCallback) {
	xhtml_p_typeFormCallback = new(Xhtml_p_typeFormCallback)
	xhtml_p_typeFormCallback.probe = probe
	xhtml_p_typeFormCallback.xhtml_p_type = xhtml_p_type
	xhtml_p_typeFormCallback.formGroup = formGroup

	xhtml_p_typeFormCallback.CreationMode = (xhtml_p_type == nil)

	return
}

type Xhtml_p_typeFormCallback struct {
	xhtml_p_type *models.Xhtml_p_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_p_typeFormCallback *Xhtml_p_typeFormCallback) OnSave() {

	log.Println("Xhtml_p_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_p_typeFormCallback.probe.formStage.Checkout()

	if xhtml_p_typeFormCallback.xhtml_p_type == nil {
		xhtml_p_typeFormCallback.xhtml_p_type = new(models.Xhtml_p_type).Stage(xhtml_p_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_p_type_ := xhtml_p_typeFormCallback.xhtml_p_type
	_ = xhtml_p_type_

	for _, formDiv := range xhtml_p_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_p_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_p_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_p_type_.Unstage(xhtml_p_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_p_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_p_type](
		xhtml_p_typeFormCallback.probe,
	)
	xhtml_p_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_p_typeFormCallback.CreationMode || xhtml_p_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_p_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_p_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_p_typeFormCallback(
			nil,
			xhtml_p_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_p_type := new(models.Xhtml_p_type)
		FillUpForm(xhtml_p_type, newFormGroup, xhtml_p_typeFormCallback.probe)
		xhtml_p_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_p_typeFormCallback.probe)
}
func __gong__New__Xhtml_param_typeFormCallback(
	xhtml_param_type *models.Xhtml_param_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_param_typeFormCallback *Xhtml_param_typeFormCallback) {
	xhtml_param_typeFormCallback = new(Xhtml_param_typeFormCallback)
	xhtml_param_typeFormCallback.probe = probe
	xhtml_param_typeFormCallback.xhtml_param_type = xhtml_param_type
	xhtml_param_typeFormCallback.formGroup = formGroup

	xhtml_param_typeFormCallback.CreationMode = (xhtml_param_type == nil)

	return
}

type Xhtml_param_typeFormCallback struct {
	xhtml_param_type *models.Xhtml_param_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_param_typeFormCallback *Xhtml_param_typeFormCallback) OnSave() {

	log.Println("Xhtml_param_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_param_typeFormCallback.probe.formStage.Checkout()

	if xhtml_param_typeFormCallback.xhtml_param_type == nil {
		xhtml_param_typeFormCallback.xhtml_param_type = new(models.Xhtml_param_type).Stage(xhtml_param_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_param_type_ := xhtml_param_typeFormCallback.xhtml_param_type
	_ = xhtml_param_type_

	for _, formDiv := range xhtml_param_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_param_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_param_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_param_type_.Unstage(xhtml_param_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_param_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_param_type](
		xhtml_param_typeFormCallback.probe,
	)
	xhtml_param_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_param_typeFormCallback.CreationMode || xhtml_param_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_param_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_param_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_param_typeFormCallback(
			nil,
			xhtml_param_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_param_type := new(models.Xhtml_param_type)
		FillUpForm(xhtml_param_type, newFormGroup, xhtml_param_typeFormCallback.probe)
		xhtml_param_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_param_typeFormCallback.probe)
}
func __gong__New__Xhtml_pre_typeFormCallback(
	xhtml_pre_type *models.Xhtml_pre_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_pre_typeFormCallback *Xhtml_pre_typeFormCallback) {
	xhtml_pre_typeFormCallback = new(Xhtml_pre_typeFormCallback)
	xhtml_pre_typeFormCallback.probe = probe
	xhtml_pre_typeFormCallback.xhtml_pre_type = xhtml_pre_type
	xhtml_pre_typeFormCallback.formGroup = formGroup

	xhtml_pre_typeFormCallback.CreationMode = (xhtml_pre_type == nil)

	return
}

type Xhtml_pre_typeFormCallback struct {
	xhtml_pre_type *models.Xhtml_pre_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_pre_typeFormCallback *Xhtml_pre_typeFormCallback) OnSave() {

	log.Println("Xhtml_pre_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_pre_typeFormCallback.probe.formStage.Checkout()

	if xhtml_pre_typeFormCallback.xhtml_pre_type == nil {
		xhtml_pre_typeFormCallback.xhtml_pre_type = new(models.Xhtml_pre_type).Stage(xhtml_pre_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_pre_type_ := xhtml_pre_typeFormCallback.xhtml_pre_type
	_ = xhtml_pre_type_

	for _, formDiv := range xhtml_pre_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_pre_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_pre_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_pre_type_.Unstage(xhtml_pre_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_pre_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_pre_type](
		xhtml_pre_typeFormCallback.probe,
	)
	xhtml_pre_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_pre_typeFormCallback.CreationMode || xhtml_pre_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_pre_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_pre_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_pre_typeFormCallback(
			nil,
			xhtml_pre_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_pre_type := new(models.Xhtml_pre_type)
		FillUpForm(xhtml_pre_type, newFormGroup, xhtml_pre_typeFormCallback.probe)
		xhtml_pre_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_pre_typeFormCallback.probe)
}
func __gong__New__Xhtml_q_typeFormCallback(
	xhtml_q_type *models.Xhtml_q_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_q_typeFormCallback *Xhtml_q_typeFormCallback) {
	xhtml_q_typeFormCallback = new(Xhtml_q_typeFormCallback)
	xhtml_q_typeFormCallback.probe = probe
	xhtml_q_typeFormCallback.xhtml_q_type = xhtml_q_type
	xhtml_q_typeFormCallback.formGroup = formGroup

	xhtml_q_typeFormCallback.CreationMode = (xhtml_q_type == nil)

	return
}

type Xhtml_q_typeFormCallback struct {
	xhtml_q_type *models.Xhtml_q_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_q_typeFormCallback *Xhtml_q_typeFormCallback) OnSave() {

	log.Println("Xhtml_q_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_q_typeFormCallback.probe.formStage.Checkout()

	if xhtml_q_typeFormCallback.xhtml_q_type == nil {
		xhtml_q_typeFormCallback.xhtml_q_type = new(models.Xhtml_q_type).Stage(xhtml_q_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_q_type_ := xhtml_q_typeFormCallback.xhtml_q_type
	_ = xhtml_q_type_

	for _, formDiv := range xhtml_q_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_q_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_q_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_q_type_.Unstage(xhtml_q_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_q_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_q_type](
		xhtml_q_typeFormCallback.probe,
	)
	xhtml_q_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_q_typeFormCallback.CreationMode || xhtml_q_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_q_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_q_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_q_typeFormCallback(
			nil,
			xhtml_q_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_q_type := new(models.Xhtml_q_type)
		FillUpForm(xhtml_q_type, newFormGroup, xhtml_q_typeFormCallback.probe)
		xhtml_q_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_q_typeFormCallback.probe)
}
func __gong__New__Xhtml_samp_typeFormCallback(
	xhtml_samp_type *models.Xhtml_samp_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_samp_typeFormCallback *Xhtml_samp_typeFormCallback) {
	xhtml_samp_typeFormCallback = new(Xhtml_samp_typeFormCallback)
	xhtml_samp_typeFormCallback.probe = probe
	xhtml_samp_typeFormCallback.xhtml_samp_type = xhtml_samp_type
	xhtml_samp_typeFormCallback.formGroup = formGroup

	xhtml_samp_typeFormCallback.CreationMode = (xhtml_samp_type == nil)

	return
}

type Xhtml_samp_typeFormCallback struct {
	xhtml_samp_type *models.Xhtml_samp_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_samp_typeFormCallback *Xhtml_samp_typeFormCallback) OnSave() {

	log.Println("Xhtml_samp_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_samp_typeFormCallback.probe.formStage.Checkout()

	if xhtml_samp_typeFormCallback.xhtml_samp_type == nil {
		xhtml_samp_typeFormCallback.xhtml_samp_type = new(models.Xhtml_samp_type).Stage(xhtml_samp_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_samp_type_ := xhtml_samp_typeFormCallback.xhtml_samp_type
	_ = xhtml_samp_type_

	for _, formDiv := range xhtml_samp_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_samp_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_samp_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_samp_type_.Unstage(xhtml_samp_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_samp_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_samp_type](
		xhtml_samp_typeFormCallback.probe,
	)
	xhtml_samp_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_samp_typeFormCallback.CreationMode || xhtml_samp_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_samp_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_samp_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_samp_typeFormCallback(
			nil,
			xhtml_samp_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_samp_type := new(models.Xhtml_samp_type)
		FillUpForm(xhtml_samp_type, newFormGroup, xhtml_samp_typeFormCallback.probe)
		xhtml_samp_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_samp_typeFormCallback.probe)
}
func __gong__New__Xhtml_span_typeFormCallback(
	xhtml_span_type *models.Xhtml_span_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_span_typeFormCallback *Xhtml_span_typeFormCallback) {
	xhtml_span_typeFormCallback = new(Xhtml_span_typeFormCallback)
	xhtml_span_typeFormCallback.probe = probe
	xhtml_span_typeFormCallback.xhtml_span_type = xhtml_span_type
	xhtml_span_typeFormCallback.formGroup = formGroup

	xhtml_span_typeFormCallback.CreationMode = (xhtml_span_type == nil)

	return
}

type Xhtml_span_typeFormCallback struct {
	xhtml_span_type *models.Xhtml_span_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_span_typeFormCallback *Xhtml_span_typeFormCallback) OnSave() {

	log.Println("Xhtml_span_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_span_typeFormCallback.probe.formStage.Checkout()

	if xhtml_span_typeFormCallback.xhtml_span_type == nil {
		xhtml_span_typeFormCallback.xhtml_span_type = new(models.Xhtml_span_type).Stage(xhtml_span_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_span_type_ := xhtml_span_typeFormCallback.xhtml_span_type
	_ = xhtml_span_type_

	for _, formDiv := range xhtml_span_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_span_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_span_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_span_type_.Unstage(xhtml_span_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_span_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_span_type](
		xhtml_span_typeFormCallback.probe,
	)
	xhtml_span_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_span_typeFormCallback.CreationMode || xhtml_span_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_span_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_span_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_span_typeFormCallback(
			nil,
			xhtml_span_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_span_type := new(models.Xhtml_span_type)
		FillUpForm(xhtml_span_type, newFormGroup, xhtml_span_typeFormCallback.probe)
		xhtml_span_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_span_typeFormCallback.probe)
}
func __gong__New__Xhtml_strong_typeFormCallback(
	xhtml_strong_type *models.Xhtml_strong_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_strong_typeFormCallback *Xhtml_strong_typeFormCallback) {
	xhtml_strong_typeFormCallback = new(Xhtml_strong_typeFormCallback)
	xhtml_strong_typeFormCallback.probe = probe
	xhtml_strong_typeFormCallback.xhtml_strong_type = xhtml_strong_type
	xhtml_strong_typeFormCallback.formGroup = formGroup

	xhtml_strong_typeFormCallback.CreationMode = (xhtml_strong_type == nil)

	return
}

type Xhtml_strong_typeFormCallback struct {
	xhtml_strong_type *models.Xhtml_strong_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_strong_typeFormCallback *Xhtml_strong_typeFormCallback) OnSave() {

	log.Println("Xhtml_strong_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_strong_typeFormCallback.probe.formStage.Checkout()

	if xhtml_strong_typeFormCallback.xhtml_strong_type == nil {
		xhtml_strong_typeFormCallback.xhtml_strong_type = new(models.Xhtml_strong_type).Stage(xhtml_strong_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_strong_type_ := xhtml_strong_typeFormCallback.xhtml_strong_type
	_ = xhtml_strong_type_

	for _, formDiv := range xhtml_strong_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_strong_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_strong_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_strong_type_.Unstage(xhtml_strong_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_strong_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_strong_type](
		xhtml_strong_typeFormCallback.probe,
	)
	xhtml_strong_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_strong_typeFormCallback.CreationMode || xhtml_strong_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_strong_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_strong_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_strong_typeFormCallback(
			nil,
			xhtml_strong_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_strong_type := new(models.Xhtml_strong_type)
		FillUpForm(xhtml_strong_type, newFormGroup, xhtml_strong_typeFormCallback.probe)
		xhtml_strong_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_strong_typeFormCallback.probe)
}
func __gong__New__Xhtml_table_typeFormCallback(
	xhtml_table_type *models.Xhtml_table_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_table_typeFormCallback *Xhtml_table_typeFormCallback) {
	xhtml_table_typeFormCallback = new(Xhtml_table_typeFormCallback)
	xhtml_table_typeFormCallback.probe = probe
	xhtml_table_typeFormCallback.xhtml_table_type = xhtml_table_type
	xhtml_table_typeFormCallback.formGroup = formGroup

	xhtml_table_typeFormCallback.CreationMode = (xhtml_table_type == nil)

	return
}

type Xhtml_table_typeFormCallback struct {
	xhtml_table_type *models.Xhtml_table_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_table_typeFormCallback *Xhtml_table_typeFormCallback) OnSave() {

	log.Println("Xhtml_table_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_table_typeFormCallback.probe.formStage.Checkout()

	if xhtml_table_typeFormCallback.xhtml_table_type == nil {
		xhtml_table_typeFormCallback.xhtml_table_type = new(models.Xhtml_table_type).Stage(xhtml_table_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_table_type_ := xhtml_table_typeFormCallback.xhtml_table_type
	_ = xhtml_table_type_

	for _, formDiv := range xhtml_table_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_table_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_table_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_table_type_.Unstage(xhtml_table_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_table_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_table_type](
		xhtml_table_typeFormCallback.probe,
	)
	xhtml_table_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_table_typeFormCallback.CreationMode || xhtml_table_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_table_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_table_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_table_typeFormCallback(
			nil,
			xhtml_table_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_table_type := new(models.Xhtml_table_type)
		FillUpForm(xhtml_table_type, newFormGroup, xhtml_table_typeFormCallback.probe)
		xhtml_table_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_table_typeFormCallback.probe)
}
func __gong__New__Xhtml_tbody_typeFormCallback(
	xhtml_tbody_type *models.Xhtml_tbody_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_tbody_typeFormCallback *Xhtml_tbody_typeFormCallback) {
	xhtml_tbody_typeFormCallback = new(Xhtml_tbody_typeFormCallback)
	xhtml_tbody_typeFormCallback.probe = probe
	xhtml_tbody_typeFormCallback.xhtml_tbody_type = xhtml_tbody_type
	xhtml_tbody_typeFormCallback.formGroup = formGroup

	xhtml_tbody_typeFormCallback.CreationMode = (xhtml_tbody_type == nil)

	return
}

type Xhtml_tbody_typeFormCallback struct {
	xhtml_tbody_type *models.Xhtml_tbody_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_tbody_typeFormCallback *Xhtml_tbody_typeFormCallback) OnSave() {

	log.Println("Xhtml_tbody_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_tbody_typeFormCallback.probe.formStage.Checkout()

	if xhtml_tbody_typeFormCallback.xhtml_tbody_type == nil {
		xhtml_tbody_typeFormCallback.xhtml_tbody_type = new(models.Xhtml_tbody_type).Stage(xhtml_tbody_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_tbody_type_ := xhtml_tbody_typeFormCallback.xhtml_tbody_type
	_ = xhtml_tbody_type_

	for _, formDiv := range xhtml_tbody_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_tbody_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_tbody_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_tbody_type_.Unstage(xhtml_tbody_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_tbody_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_tbody_type](
		xhtml_tbody_typeFormCallback.probe,
	)
	xhtml_tbody_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_tbody_typeFormCallback.CreationMode || xhtml_tbody_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_tbody_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_tbody_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_tbody_typeFormCallback(
			nil,
			xhtml_tbody_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_tbody_type := new(models.Xhtml_tbody_type)
		FillUpForm(xhtml_tbody_type, newFormGroup, xhtml_tbody_typeFormCallback.probe)
		xhtml_tbody_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_tbody_typeFormCallback.probe)
}
func __gong__New__Xhtml_td_typeFormCallback(
	xhtml_td_type *models.Xhtml_td_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_td_typeFormCallback *Xhtml_td_typeFormCallback) {
	xhtml_td_typeFormCallback = new(Xhtml_td_typeFormCallback)
	xhtml_td_typeFormCallback.probe = probe
	xhtml_td_typeFormCallback.xhtml_td_type = xhtml_td_type
	xhtml_td_typeFormCallback.formGroup = formGroup

	xhtml_td_typeFormCallback.CreationMode = (xhtml_td_type == nil)

	return
}

type Xhtml_td_typeFormCallback struct {
	xhtml_td_type *models.Xhtml_td_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_td_typeFormCallback *Xhtml_td_typeFormCallback) OnSave() {

	log.Println("Xhtml_td_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_td_typeFormCallback.probe.formStage.Checkout()

	if xhtml_td_typeFormCallback.xhtml_td_type == nil {
		xhtml_td_typeFormCallback.xhtml_td_type = new(models.Xhtml_td_type).Stage(xhtml_td_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_td_type_ := xhtml_td_typeFormCallback.xhtml_td_type
	_ = xhtml_td_type_

	for _, formDiv := range xhtml_td_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_td_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_td_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_td_type_.Unstage(xhtml_td_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_td_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_td_type](
		xhtml_td_typeFormCallback.probe,
	)
	xhtml_td_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_td_typeFormCallback.CreationMode || xhtml_td_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_td_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_td_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_td_typeFormCallback(
			nil,
			xhtml_td_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_td_type := new(models.Xhtml_td_type)
		FillUpForm(xhtml_td_type, newFormGroup, xhtml_td_typeFormCallback.probe)
		xhtml_td_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_td_typeFormCallback.probe)
}
func __gong__New__Xhtml_tfoot_typeFormCallback(
	xhtml_tfoot_type *models.Xhtml_tfoot_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_tfoot_typeFormCallback *Xhtml_tfoot_typeFormCallback) {
	xhtml_tfoot_typeFormCallback = new(Xhtml_tfoot_typeFormCallback)
	xhtml_tfoot_typeFormCallback.probe = probe
	xhtml_tfoot_typeFormCallback.xhtml_tfoot_type = xhtml_tfoot_type
	xhtml_tfoot_typeFormCallback.formGroup = formGroup

	xhtml_tfoot_typeFormCallback.CreationMode = (xhtml_tfoot_type == nil)

	return
}

type Xhtml_tfoot_typeFormCallback struct {
	xhtml_tfoot_type *models.Xhtml_tfoot_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_tfoot_typeFormCallback *Xhtml_tfoot_typeFormCallback) OnSave() {

	log.Println("Xhtml_tfoot_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_tfoot_typeFormCallback.probe.formStage.Checkout()

	if xhtml_tfoot_typeFormCallback.xhtml_tfoot_type == nil {
		xhtml_tfoot_typeFormCallback.xhtml_tfoot_type = new(models.Xhtml_tfoot_type).Stage(xhtml_tfoot_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_tfoot_type_ := xhtml_tfoot_typeFormCallback.xhtml_tfoot_type
	_ = xhtml_tfoot_type_

	for _, formDiv := range xhtml_tfoot_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_tfoot_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_tfoot_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_tfoot_type_.Unstage(xhtml_tfoot_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_tfoot_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_tfoot_type](
		xhtml_tfoot_typeFormCallback.probe,
	)
	xhtml_tfoot_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_tfoot_typeFormCallback.CreationMode || xhtml_tfoot_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_tfoot_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_tfoot_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_tfoot_typeFormCallback(
			nil,
			xhtml_tfoot_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_tfoot_type := new(models.Xhtml_tfoot_type)
		FillUpForm(xhtml_tfoot_type, newFormGroup, xhtml_tfoot_typeFormCallback.probe)
		xhtml_tfoot_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_tfoot_typeFormCallback.probe)
}
func __gong__New__Xhtml_th_typeFormCallback(
	xhtml_th_type *models.Xhtml_th_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_th_typeFormCallback *Xhtml_th_typeFormCallback) {
	xhtml_th_typeFormCallback = new(Xhtml_th_typeFormCallback)
	xhtml_th_typeFormCallback.probe = probe
	xhtml_th_typeFormCallback.xhtml_th_type = xhtml_th_type
	xhtml_th_typeFormCallback.formGroup = formGroup

	xhtml_th_typeFormCallback.CreationMode = (xhtml_th_type == nil)

	return
}

type Xhtml_th_typeFormCallback struct {
	xhtml_th_type *models.Xhtml_th_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_th_typeFormCallback *Xhtml_th_typeFormCallback) OnSave() {

	log.Println("Xhtml_th_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_th_typeFormCallback.probe.formStage.Checkout()

	if xhtml_th_typeFormCallback.xhtml_th_type == nil {
		xhtml_th_typeFormCallback.xhtml_th_type = new(models.Xhtml_th_type).Stage(xhtml_th_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_th_type_ := xhtml_th_typeFormCallback.xhtml_th_type
	_ = xhtml_th_type_

	for _, formDiv := range xhtml_th_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_th_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_th_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_th_type_.Unstage(xhtml_th_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_th_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_th_type](
		xhtml_th_typeFormCallback.probe,
	)
	xhtml_th_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_th_typeFormCallback.CreationMode || xhtml_th_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_th_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_th_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_th_typeFormCallback(
			nil,
			xhtml_th_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_th_type := new(models.Xhtml_th_type)
		FillUpForm(xhtml_th_type, newFormGroup, xhtml_th_typeFormCallback.probe)
		xhtml_th_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_th_typeFormCallback.probe)
}
func __gong__New__Xhtml_thead_typeFormCallback(
	xhtml_thead_type *models.Xhtml_thead_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_thead_typeFormCallback *Xhtml_thead_typeFormCallback) {
	xhtml_thead_typeFormCallback = new(Xhtml_thead_typeFormCallback)
	xhtml_thead_typeFormCallback.probe = probe
	xhtml_thead_typeFormCallback.xhtml_thead_type = xhtml_thead_type
	xhtml_thead_typeFormCallback.formGroup = formGroup

	xhtml_thead_typeFormCallback.CreationMode = (xhtml_thead_type == nil)

	return
}

type Xhtml_thead_typeFormCallback struct {
	xhtml_thead_type *models.Xhtml_thead_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_thead_typeFormCallback *Xhtml_thead_typeFormCallback) OnSave() {

	log.Println("Xhtml_thead_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_thead_typeFormCallback.probe.formStage.Checkout()

	if xhtml_thead_typeFormCallback.xhtml_thead_type == nil {
		xhtml_thead_typeFormCallback.xhtml_thead_type = new(models.Xhtml_thead_type).Stage(xhtml_thead_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_thead_type_ := xhtml_thead_typeFormCallback.xhtml_thead_type
	_ = xhtml_thead_type_

	for _, formDiv := range xhtml_thead_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_thead_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_thead_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_thead_type_.Unstage(xhtml_thead_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_thead_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_thead_type](
		xhtml_thead_typeFormCallback.probe,
	)
	xhtml_thead_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_thead_typeFormCallback.CreationMode || xhtml_thead_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_thead_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_thead_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_thead_typeFormCallback(
			nil,
			xhtml_thead_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_thead_type := new(models.Xhtml_thead_type)
		FillUpForm(xhtml_thead_type, newFormGroup, xhtml_thead_typeFormCallback.probe)
		xhtml_thead_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_thead_typeFormCallback.probe)
}
func __gong__New__Xhtml_tr_typeFormCallback(
	xhtml_tr_type *models.Xhtml_tr_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_tr_typeFormCallback *Xhtml_tr_typeFormCallback) {
	xhtml_tr_typeFormCallback = new(Xhtml_tr_typeFormCallback)
	xhtml_tr_typeFormCallback.probe = probe
	xhtml_tr_typeFormCallback.xhtml_tr_type = xhtml_tr_type
	xhtml_tr_typeFormCallback.formGroup = formGroup

	xhtml_tr_typeFormCallback.CreationMode = (xhtml_tr_type == nil)

	return
}

type Xhtml_tr_typeFormCallback struct {
	xhtml_tr_type *models.Xhtml_tr_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_tr_typeFormCallback *Xhtml_tr_typeFormCallback) OnSave() {

	log.Println("Xhtml_tr_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_tr_typeFormCallback.probe.formStage.Checkout()

	if xhtml_tr_typeFormCallback.xhtml_tr_type == nil {
		xhtml_tr_typeFormCallback.xhtml_tr_type = new(models.Xhtml_tr_type).Stage(xhtml_tr_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_tr_type_ := xhtml_tr_typeFormCallback.xhtml_tr_type
	_ = xhtml_tr_type_

	for _, formDiv := range xhtml_tr_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_tr_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_tr_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_tr_type_.Unstage(xhtml_tr_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_tr_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_tr_type](
		xhtml_tr_typeFormCallback.probe,
	)
	xhtml_tr_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_tr_typeFormCallback.CreationMode || xhtml_tr_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_tr_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_tr_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_tr_typeFormCallback(
			nil,
			xhtml_tr_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_tr_type := new(models.Xhtml_tr_type)
		FillUpForm(xhtml_tr_type, newFormGroup, xhtml_tr_typeFormCallback.probe)
		xhtml_tr_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_tr_typeFormCallback.probe)
}
func __gong__New__Xhtml_ul_typeFormCallback(
	xhtml_ul_type *models.Xhtml_ul_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_ul_typeFormCallback *Xhtml_ul_typeFormCallback) {
	xhtml_ul_typeFormCallback = new(Xhtml_ul_typeFormCallback)
	xhtml_ul_typeFormCallback.probe = probe
	xhtml_ul_typeFormCallback.xhtml_ul_type = xhtml_ul_type
	xhtml_ul_typeFormCallback.formGroup = formGroup

	xhtml_ul_typeFormCallback.CreationMode = (xhtml_ul_type == nil)

	return
}

type Xhtml_ul_typeFormCallback struct {
	xhtml_ul_type *models.Xhtml_ul_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_ul_typeFormCallback *Xhtml_ul_typeFormCallback) OnSave() {

	log.Println("Xhtml_ul_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_ul_typeFormCallback.probe.formStage.Checkout()

	if xhtml_ul_typeFormCallback.xhtml_ul_type == nil {
		xhtml_ul_typeFormCallback.xhtml_ul_type = new(models.Xhtml_ul_type).Stage(xhtml_ul_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_ul_type_ := xhtml_ul_typeFormCallback.xhtml_ul_type
	_ = xhtml_ul_type_

	for _, formDiv := range xhtml_ul_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_ul_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_ul_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_ul_type_.Unstage(xhtml_ul_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_ul_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_ul_type](
		xhtml_ul_typeFormCallback.probe,
	)
	xhtml_ul_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_ul_typeFormCallback.CreationMode || xhtml_ul_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_ul_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_ul_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_ul_typeFormCallback(
			nil,
			xhtml_ul_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_ul_type := new(models.Xhtml_ul_type)
		FillUpForm(xhtml_ul_type, newFormGroup, xhtml_ul_typeFormCallback.probe)
		xhtml_ul_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_ul_typeFormCallback.probe)
}
func __gong__New__Xhtml_var_typeFormCallback(
	xhtml_var_type *models.Xhtml_var_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_var_typeFormCallback *Xhtml_var_typeFormCallback) {
	xhtml_var_typeFormCallback = new(Xhtml_var_typeFormCallback)
	xhtml_var_typeFormCallback.probe = probe
	xhtml_var_typeFormCallback.xhtml_var_type = xhtml_var_type
	xhtml_var_typeFormCallback.formGroup = formGroup

	xhtml_var_typeFormCallback.CreationMode = (xhtml_var_type == nil)

	return
}

type Xhtml_var_typeFormCallback struct {
	xhtml_var_type *models.Xhtml_var_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_var_typeFormCallback *Xhtml_var_typeFormCallback) OnSave() {

	log.Println("Xhtml_var_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_var_typeFormCallback.probe.formStage.Checkout()

	if xhtml_var_typeFormCallback.xhtml_var_type == nil {
		xhtml_var_typeFormCallback.xhtml_var_type = new(models.Xhtml_var_type).Stage(xhtml_var_typeFormCallback.probe.stageOfInterest)
	}
	xhtml_var_type_ := xhtml_var_typeFormCallback.xhtml_var_type
	_ = xhtml_var_type_

	for _, formDiv := range xhtml_var_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_var_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_var_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_var_type_.Unstage(xhtml_var_typeFormCallback.probe.stageOfInterest)
	}

	xhtml_var_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Xhtml_var_type](
		xhtml_var_typeFormCallback.probe,
	)
	xhtml_var_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_var_typeFormCallback.CreationMode || xhtml_var_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_var_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_var_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Xhtml_var_typeFormCallback(
			nil,
			xhtml_var_typeFormCallback.probe,
			newFormGroup,
		)
		xhtml_var_type := new(models.Xhtml_var_type)
		FillUpForm(xhtml_var_type, newFormGroup, xhtml_var_typeFormCallback.probe)
		xhtml_var_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_var_typeFormCallback.probe)
}
