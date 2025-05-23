// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gong/lib/table/go/db"
	"github.com/fullstack-lang/gong/lib/table/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_FormFieldInt_sql sql.NullBool
var dummy_FormFieldInt_time time.Duration
var dummy_FormFieldInt_sort sort.Float64Slice

// FormFieldIntAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model formfieldintAPI
type FormFieldIntAPI struct {
	gorm.Model

	models.FormFieldInt_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	FormFieldIntPointersEncoding FormFieldIntPointersEncoding
}

// FormFieldIntPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type FormFieldIntPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// FormFieldIntDB describes a formfieldint in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model formfieldintDB
type FormFieldIntDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field formfieldintDB.Name
	Name_Data sql.NullString

	// Declation for basic field formfieldintDB.Value
	Value_Data sql.NullInt64

	// Declation for basic field formfieldintDB.HasMinValidator
	// provide the sql storage for the boolan
	HasMinValidator_Data sql.NullBool

	// Declation for basic field formfieldintDB.MinValue
	MinValue_Data sql.NullInt64

	// Declation for basic field formfieldintDB.HasMaxValidator
	// provide the sql storage for the boolan
	HasMaxValidator_Data sql.NullBool

	// Declation for basic field formfieldintDB.MaxValue
	MaxValue_Data sql.NullInt64

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	FormFieldIntPointersEncoding
}

// FormFieldIntDBs arrays formfieldintDBs
// swagger:response formfieldintDBsResponse
type FormFieldIntDBs []FormFieldIntDB

// FormFieldIntDBResponse provides response
// swagger:response formfieldintDBResponse
type FormFieldIntDBResponse struct {
	FormFieldIntDB
}

// FormFieldIntWOP is a FormFieldInt without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type FormFieldIntWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Value int `xlsx:"2"`

	HasMinValidator bool `xlsx:"3"`

	MinValue int `xlsx:"4"`

	HasMaxValidator bool `xlsx:"5"`

	MaxValue int `xlsx:"6"`
	// insertion for WOP pointer fields
}

var FormFieldInt_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Value",
	"HasMinValidator",
	"MinValue",
	"HasMaxValidator",
	"MaxValue",
}

type BackRepoFormFieldIntStruct struct {
	// stores FormFieldIntDB according to their gorm ID
	Map_FormFieldIntDBID_FormFieldIntDB map[uint]*FormFieldIntDB

	// stores FormFieldIntDB ID according to FormFieldInt address
	Map_FormFieldIntPtr_FormFieldIntDBID map[*models.FormFieldInt]uint

	// stores FormFieldInt according to their gorm ID
	Map_FormFieldIntDBID_FormFieldIntPtr map[uint]*models.FormFieldInt

	db db.DBInterface

	stage *models.Stage
}

func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) GetStage() (stage *models.Stage) {
	stage = backRepoFormFieldInt.stage
	return
}

func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) GetDB() db.DBInterface {
	return backRepoFormFieldInt.db
}

// GetFormFieldIntDBFromFormFieldIntPtr is a handy function to access the back repo instance from the stage instance
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) GetFormFieldIntDBFromFormFieldIntPtr(formfieldint *models.FormFieldInt) (formfieldintDB *FormFieldIntDB) {
	id := backRepoFormFieldInt.Map_FormFieldIntPtr_FormFieldIntDBID[formfieldint]
	formfieldintDB = backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB[id]
	return
}

// BackRepoFormFieldInt.CommitPhaseOne commits all staged instances of FormFieldInt to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var formfieldints []*models.FormFieldInt
	for formfieldint := range stage.FormFieldInts {
		formfieldints = append(formfieldints, formfieldint)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(formfieldints, func(i, j int) bool {
		return stage.FormFieldIntMap_Staged_Order[formfieldints[i]] < stage.FormFieldIntMap_Staged_Order[formfieldints[j]]
	})

	for _, formfieldint := range formfieldints {
		backRepoFormFieldInt.CommitPhaseOneInstance(formfieldint)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, formfieldint := range backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr {
		if _, ok := stage.FormFieldInts[formfieldint]; !ok {
			backRepoFormFieldInt.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFormFieldInt.CommitDeleteInstance commits deletion of FormFieldInt to the BackRepo
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) CommitDeleteInstance(id uint) (Error error) {

	formfieldint := backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr[id]

	// formfieldint is not staged anymore, remove formfieldintDB
	formfieldintDB := backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB[id]
	db, _ := backRepoFormFieldInt.db.Unscoped()
	_, err := db.Delete(formfieldintDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoFormFieldInt.Map_FormFieldIntPtr_FormFieldIntDBID, formfieldint)
	delete(backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr, id)
	delete(backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB, id)

	return
}

// BackRepoFormFieldInt.CommitPhaseOneInstance commits formfieldint staged instances of FormFieldInt to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) CommitPhaseOneInstance(formfieldint *models.FormFieldInt) (Error error) {

	// check if the formfieldint is not commited yet
	if _, ok := backRepoFormFieldInt.Map_FormFieldIntPtr_FormFieldIntDBID[formfieldint]; ok {
		return
	}

	// initiate formfieldint
	var formfieldintDB FormFieldIntDB
	formfieldintDB.CopyBasicFieldsFromFormFieldInt(formfieldint)

	_, err := backRepoFormFieldInt.db.Create(&formfieldintDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoFormFieldInt.Map_FormFieldIntPtr_FormFieldIntDBID[formfieldint] = formfieldintDB.ID
	backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr[formfieldintDB.ID] = formfieldint
	backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB[formfieldintDB.ID] = &formfieldintDB

	return
}

// BackRepoFormFieldInt.CommitPhaseTwo commits all staged instances of FormFieldInt to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, formfieldint := range backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr {
		backRepoFormFieldInt.CommitPhaseTwoInstance(backRepo, idx, formfieldint)
	}

	return
}

// BackRepoFormFieldInt.CommitPhaseTwoInstance commits {{structname }} of models.FormFieldInt to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, formfieldint *models.FormFieldInt) (Error error) {

	// fetch matching formfieldintDB
	if formfieldintDB, ok := backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB[idx]; ok {

		formfieldintDB.CopyBasicFieldsFromFormFieldInt(formfieldint)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoFormFieldInt.db.Save(formfieldintDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown FormFieldInt intance %s", formfieldint.Name))
		return err
	}

	return
}

// BackRepoFormFieldInt.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) CheckoutPhaseOne() (Error error) {

	formfieldintDBArray := make([]FormFieldIntDB, 0)
	_, err := backRepoFormFieldInt.db.Find(&formfieldintDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	formfieldintInstancesToBeRemovedFromTheStage := make(map[*models.FormFieldInt]any)
	for key, value := range backRepoFormFieldInt.stage.FormFieldInts {
		formfieldintInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, formfieldintDB := range formfieldintDBArray {
		backRepoFormFieldInt.CheckoutPhaseOneInstance(&formfieldintDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		formfieldint, ok := backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr[formfieldintDB.ID]
		if ok {
			delete(formfieldintInstancesToBeRemovedFromTheStage, formfieldint)
		}
	}

	// remove from stage and back repo's 3 maps all formfieldints that are not in the checkout
	for formfieldint := range formfieldintInstancesToBeRemovedFromTheStage {
		formfieldint.Unstage(backRepoFormFieldInt.GetStage())

		// remove instance from the back repo 3 maps
		formfieldintID := backRepoFormFieldInt.Map_FormFieldIntPtr_FormFieldIntDBID[formfieldint]
		delete(backRepoFormFieldInt.Map_FormFieldIntPtr_FormFieldIntDBID, formfieldint)
		delete(backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB, formfieldintID)
		delete(backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr, formfieldintID)
	}

	return
}

// CheckoutPhaseOneInstance takes a formfieldintDB that has been found in the DB, updates the backRepo and stages the
// models version of the formfieldintDB
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) CheckoutPhaseOneInstance(formfieldintDB *FormFieldIntDB) (Error error) {

	formfieldint, ok := backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr[formfieldintDB.ID]
	if !ok {
		formfieldint = new(models.FormFieldInt)

		backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr[formfieldintDB.ID] = formfieldint
		backRepoFormFieldInt.Map_FormFieldIntPtr_FormFieldIntDBID[formfieldint] = formfieldintDB.ID

		// append model store with the new element
		formfieldint.Name = formfieldintDB.Name_Data.String
		formfieldint.Stage(backRepoFormFieldInt.GetStage())
	}
	formfieldintDB.CopyBasicFieldsToFormFieldInt(formfieldint)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	formfieldint.Stage(backRepoFormFieldInt.GetStage())

	// preserve pointer to formfieldintDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_FormFieldIntDBID_FormFieldIntDB)[formfieldintDB hold variable pointers
	formfieldintDB_Data := *formfieldintDB
	preservedPtrToFormFieldInt := &formfieldintDB_Data
	backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB[formfieldintDB.ID] = preservedPtrToFormFieldInt

	return
}

// BackRepoFormFieldInt.CheckoutPhaseTwo Checkouts all staged instances of FormFieldInt to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, formfieldintDB := range backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB {
		backRepoFormFieldInt.CheckoutPhaseTwoInstance(backRepo, formfieldintDB)
	}
	return
}

// BackRepoFormFieldInt.CheckoutPhaseTwoInstance Checkouts staged instances of FormFieldInt to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, formfieldintDB *FormFieldIntDB) (Error error) {

	formfieldint := backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr[formfieldintDB.ID]

	formfieldintDB.DecodePointers(backRepo, formfieldint)

	return
}

func (formfieldintDB *FormFieldIntDB) DecodePointers(backRepo *BackRepoStruct, formfieldint *models.FormFieldInt) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitFormFieldInt allows commit of a single formfieldint (if already staged)
func (backRepo *BackRepoStruct) CommitFormFieldInt(formfieldint *models.FormFieldInt) {
	backRepo.BackRepoFormFieldInt.CommitPhaseOneInstance(formfieldint)
	if id, ok := backRepo.BackRepoFormFieldInt.Map_FormFieldIntPtr_FormFieldIntDBID[formfieldint]; ok {
		backRepo.BackRepoFormFieldInt.CommitPhaseTwoInstance(backRepo, id, formfieldint)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFormFieldInt allows checkout of a single formfieldint (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFormFieldInt(formfieldint *models.FormFieldInt) {
	// check if the formfieldint is staged
	if _, ok := backRepo.BackRepoFormFieldInt.Map_FormFieldIntPtr_FormFieldIntDBID[formfieldint]; ok {

		if id, ok := backRepo.BackRepoFormFieldInt.Map_FormFieldIntPtr_FormFieldIntDBID[formfieldint]; ok {
			var formfieldintDB FormFieldIntDB
			formfieldintDB.ID = id

			if _, err := backRepo.BackRepoFormFieldInt.db.First(&formfieldintDB, id); err != nil {
				log.Fatalln("CheckoutFormFieldInt : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFormFieldInt.CheckoutPhaseOneInstance(&formfieldintDB)
			backRepo.BackRepoFormFieldInt.CheckoutPhaseTwoInstance(backRepo, &formfieldintDB)
		}
	}
}

// CopyBasicFieldsFromFormFieldInt
func (formfieldintDB *FormFieldIntDB) CopyBasicFieldsFromFormFieldInt(formfieldint *models.FormFieldInt) {
	// insertion point for fields commit

	formfieldintDB.Name_Data.String = formfieldint.Name
	formfieldintDB.Name_Data.Valid = true

	formfieldintDB.Value_Data.Int64 = int64(formfieldint.Value)
	formfieldintDB.Value_Data.Valid = true

	formfieldintDB.HasMinValidator_Data.Bool = formfieldint.HasMinValidator
	formfieldintDB.HasMinValidator_Data.Valid = true

	formfieldintDB.MinValue_Data.Int64 = int64(formfieldint.MinValue)
	formfieldintDB.MinValue_Data.Valid = true

	formfieldintDB.HasMaxValidator_Data.Bool = formfieldint.HasMaxValidator
	formfieldintDB.HasMaxValidator_Data.Valid = true

	formfieldintDB.MaxValue_Data.Int64 = int64(formfieldint.MaxValue)
	formfieldintDB.MaxValue_Data.Valid = true
}

// CopyBasicFieldsFromFormFieldInt_WOP
func (formfieldintDB *FormFieldIntDB) CopyBasicFieldsFromFormFieldInt_WOP(formfieldint *models.FormFieldInt_WOP) {
	// insertion point for fields commit

	formfieldintDB.Name_Data.String = formfieldint.Name
	formfieldintDB.Name_Data.Valid = true

	formfieldintDB.Value_Data.Int64 = int64(formfieldint.Value)
	formfieldintDB.Value_Data.Valid = true

	formfieldintDB.HasMinValidator_Data.Bool = formfieldint.HasMinValidator
	formfieldintDB.HasMinValidator_Data.Valid = true

	formfieldintDB.MinValue_Data.Int64 = int64(formfieldint.MinValue)
	formfieldintDB.MinValue_Data.Valid = true

	formfieldintDB.HasMaxValidator_Data.Bool = formfieldint.HasMaxValidator
	formfieldintDB.HasMaxValidator_Data.Valid = true

	formfieldintDB.MaxValue_Data.Int64 = int64(formfieldint.MaxValue)
	formfieldintDB.MaxValue_Data.Valid = true
}

// CopyBasicFieldsFromFormFieldIntWOP
func (formfieldintDB *FormFieldIntDB) CopyBasicFieldsFromFormFieldIntWOP(formfieldint *FormFieldIntWOP) {
	// insertion point for fields commit

	formfieldintDB.Name_Data.String = formfieldint.Name
	formfieldintDB.Name_Data.Valid = true

	formfieldintDB.Value_Data.Int64 = int64(formfieldint.Value)
	formfieldintDB.Value_Data.Valid = true

	formfieldintDB.HasMinValidator_Data.Bool = formfieldint.HasMinValidator
	formfieldintDB.HasMinValidator_Data.Valid = true

	formfieldintDB.MinValue_Data.Int64 = int64(formfieldint.MinValue)
	formfieldintDB.MinValue_Data.Valid = true

	formfieldintDB.HasMaxValidator_Data.Bool = formfieldint.HasMaxValidator
	formfieldintDB.HasMaxValidator_Data.Valid = true

	formfieldintDB.MaxValue_Data.Int64 = int64(formfieldint.MaxValue)
	formfieldintDB.MaxValue_Data.Valid = true
}

// CopyBasicFieldsToFormFieldInt
func (formfieldintDB *FormFieldIntDB) CopyBasicFieldsToFormFieldInt(formfieldint *models.FormFieldInt) {
	// insertion point for checkout of basic fields (back repo to stage)
	formfieldint.Name = formfieldintDB.Name_Data.String
	formfieldint.Value = int(formfieldintDB.Value_Data.Int64)
	formfieldint.HasMinValidator = formfieldintDB.HasMinValidator_Data.Bool
	formfieldint.MinValue = int(formfieldintDB.MinValue_Data.Int64)
	formfieldint.HasMaxValidator = formfieldintDB.HasMaxValidator_Data.Bool
	formfieldint.MaxValue = int(formfieldintDB.MaxValue_Data.Int64)
}

// CopyBasicFieldsToFormFieldInt_WOP
func (formfieldintDB *FormFieldIntDB) CopyBasicFieldsToFormFieldInt_WOP(formfieldint *models.FormFieldInt_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	formfieldint.Name = formfieldintDB.Name_Data.String
	formfieldint.Value = int(formfieldintDB.Value_Data.Int64)
	formfieldint.HasMinValidator = formfieldintDB.HasMinValidator_Data.Bool
	formfieldint.MinValue = int(formfieldintDB.MinValue_Data.Int64)
	formfieldint.HasMaxValidator = formfieldintDB.HasMaxValidator_Data.Bool
	formfieldint.MaxValue = int(formfieldintDB.MaxValue_Data.Int64)
}

// CopyBasicFieldsToFormFieldIntWOP
func (formfieldintDB *FormFieldIntDB) CopyBasicFieldsToFormFieldIntWOP(formfieldint *FormFieldIntWOP) {
	formfieldint.ID = int(formfieldintDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	formfieldint.Name = formfieldintDB.Name_Data.String
	formfieldint.Value = int(formfieldintDB.Value_Data.Int64)
	formfieldint.HasMinValidator = formfieldintDB.HasMinValidator_Data.Bool
	formfieldint.MinValue = int(formfieldintDB.MinValue_Data.Int64)
	formfieldint.HasMaxValidator = formfieldintDB.HasMaxValidator_Data.Bool
	formfieldint.MaxValue = int(formfieldintDB.MaxValue_Data.Int64)
}

// Backup generates a json file from a slice of all FormFieldIntDB instances in the backrepo
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "FormFieldIntDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormFieldIntDB, 0)
	for _, formfieldintDB := range backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB {
		forBackup = append(forBackup, formfieldintDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json FormFieldInt ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json FormFieldInt file", err.Error())
	}
}

// Backup generates a json file from a slice of all FormFieldIntDB instances in the backrepo
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormFieldIntDB, 0)
	for _, formfieldintDB := range backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB {
		forBackup = append(forBackup, formfieldintDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("FormFieldInt")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&FormFieldInt_Fields, -1)
	for _, formfieldintDB := range forBackup {

		var formfieldintWOP FormFieldIntWOP
		formfieldintDB.CopyBasicFieldsToFormFieldIntWOP(&formfieldintWOP)

		row := sh.AddRow()
		row.WriteStruct(&formfieldintWOP, -1)
	}
}

// RestoreXL from the "FormFieldInt" sheet all FormFieldIntDB instances
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFormFieldIntid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["FormFieldInt"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFormFieldInt.rowVisitorFormFieldInt)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) rowVisitorFormFieldInt(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var formfieldintWOP FormFieldIntWOP
		row.ReadStruct(&formfieldintWOP)

		// add the unmarshalled struct to the stage
		formfieldintDB := new(FormFieldIntDB)
		formfieldintDB.CopyBasicFieldsFromFormFieldIntWOP(&formfieldintWOP)

		formfieldintDB_ID_atBackupTime := formfieldintDB.ID
		formfieldintDB.ID = 0
		_, err := backRepoFormFieldInt.db.Create(formfieldintDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB[formfieldintDB.ID] = formfieldintDB
		BackRepoFormFieldIntid_atBckpTime_newID[formfieldintDB_ID_atBackupTime] = formfieldintDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "FormFieldIntDB.json" in dirPath that stores an array
// of FormFieldIntDB and stores it in the database
// the map BackRepoFormFieldIntid_atBckpTime_newID is updated accordingly
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFormFieldIntid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "FormFieldIntDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json FormFieldInt file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*FormFieldIntDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_FormFieldIntDBID_FormFieldIntDB
	for _, formfieldintDB := range forRestore {

		formfieldintDB_ID_atBackupTime := formfieldintDB.ID
		formfieldintDB.ID = 0
		_, err := backRepoFormFieldInt.db.Create(formfieldintDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB[formfieldintDB.ID] = formfieldintDB
		BackRepoFormFieldIntid_atBckpTime_newID[formfieldintDB_ID_atBackupTime] = formfieldintDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json FormFieldInt file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<FormFieldInt>id_atBckpTime_newID
// to compute new index
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) RestorePhaseTwo() {

	for _, formfieldintDB := range backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB {

		// next line of code is to avert unused variable compilation error
		_ = formfieldintDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoFormFieldInt.db.Model(formfieldintDB)
		_, err := db.Updates(*formfieldintDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoFormFieldInt.ResetReversePointers commits all staged instances of FormFieldInt to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, formfieldint := range backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr {
		backRepoFormFieldInt.ResetReversePointersInstance(backRepo, idx, formfieldint)
	}

	return
}

func (backRepoFormFieldInt *BackRepoFormFieldIntStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, formfieldint *models.FormFieldInt) (Error error) {

	// fetch matching formfieldintDB
	if formfieldintDB, ok := backRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB[idx]; ok {
		_ = formfieldintDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFormFieldIntid_atBckpTime_newID map[uint]uint
