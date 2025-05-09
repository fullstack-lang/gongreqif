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
var dummy_CheckBox_sql sql.NullBool
var dummy_CheckBox_time time.Duration
var dummy_CheckBox_sort sort.Float64Slice

// CheckBoxAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model checkboxAPI
type CheckBoxAPI struct {
	gorm.Model

	models.CheckBox_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	CheckBoxPointersEncoding CheckBoxPointersEncoding
}

// CheckBoxPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type CheckBoxPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// CheckBoxDB describes a checkbox in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model checkboxDB
type CheckBoxDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field checkboxDB.Name
	Name_Data sql.NullString

	// Declation for basic field checkboxDB.Value
	// provide the sql storage for the boolan
	Value_Data sql.NullBool

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	CheckBoxPointersEncoding
}

// CheckBoxDBs arrays checkboxDBs
// swagger:response checkboxDBsResponse
type CheckBoxDBs []CheckBoxDB

// CheckBoxDBResponse provides response
// swagger:response checkboxDBResponse
type CheckBoxDBResponse struct {
	CheckBoxDB
}

// CheckBoxWOP is a CheckBox without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type CheckBoxWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Value bool `xlsx:"2"`
	// insertion for WOP pointer fields
}

var CheckBox_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Value",
}

type BackRepoCheckBoxStruct struct {
	// stores CheckBoxDB according to their gorm ID
	Map_CheckBoxDBID_CheckBoxDB map[uint]*CheckBoxDB

	// stores CheckBoxDB ID according to CheckBox address
	Map_CheckBoxPtr_CheckBoxDBID map[*models.CheckBox]uint

	// stores CheckBox according to their gorm ID
	Map_CheckBoxDBID_CheckBoxPtr map[uint]*models.CheckBox

	db db.DBInterface

	stage *models.Stage
}

func (backRepoCheckBox *BackRepoCheckBoxStruct) GetStage() (stage *models.Stage) {
	stage = backRepoCheckBox.stage
	return
}

func (backRepoCheckBox *BackRepoCheckBoxStruct) GetDB() db.DBInterface {
	return backRepoCheckBox.db
}

// GetCheckBoxDBFromCheckBoxPtr is a handy function to access the back repo instance from the stage instance
func (backRepoCheckBox *BackRepoCheckBoxStruct) GetCheckBoxDBFromCheckBoxPtr(checkbox *models.CheckBox) (checkboxDB *CheckBoxDB) {
	id := backRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID[checkbox]
	checkboxDB = backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB[id]
	return
}

// BackRepoCheckBox.CommitPhaseOne commits all staged instances of CheckBox to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoCheckBox *BackRepoCheckBoxStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var checkboxs []*models.CheckBox
	for checkbox := range stage.CheckBoxs {
		checkboxs = append(checkboxs, checkbox)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(checkboxs, func(i, j int) bool {
		return stage.CheckBoxMap_Staged_Order[checkboxs[i]] < stage.CheckBoxMap_Staged_Order[checkboxs[j]]
	})

	for _, checkbox := range checkboxs {
		backRepoCheckBox.CommitPhaseOneInstance(checkbox)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, checkbox := range backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr {
		if _, ok := stage.CheckBoxs[checkbox]; !ok {
			backRepoCheckBox.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoCheckBox.CommitDeleteInstance commits deletion of CheckBox to the BackRepo
func (backRepoCheckBox *BackRepoCheckBoxStruct) CommitDeleteInstance(id uint) (Error error) {

	checkbox := backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[id]

	// checkbox is not staged anymore, remove checkboxDB
	checkboxDB := backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB[id]
	db, _ := backRepoCheckBox.db.Unscoped()
	_, err := db.Delete(checkboxDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID, checkbox)
	delete(backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr, id)
	delete(backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB, id)

	return
}

// BackRepoCheckBox.CommitPhaseOneInstance commits checkbox staged instances of CheckBox to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoCheckBox *BackRepoCheckBoxStruct) CommitPhaseOneInstance(checkbox *models.CheckBox) (Error error) {

	// check if the checkbox is not commited yet
	if _, ok := backRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID[checkbox]; ok {
		return
	}

	// initiate checkbox
	var checkboxDB CheckBoxDB
	checkboxDB.CopyBasicFieldsFromCheckBox(checkbox)

	_, err := backRepoCheckBox.db.Create(&checkboxDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID[checkbox] = checkboxDB.ID
	backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[checkboxDB.ID] = checkbox
	backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB[checkboxDB.ID] = &checkboxDB

	return
}

// BackRepoCheckBox.CommitPhaseTwo commits all staged instances of CheckBox to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCheckBox *BackRepoCheckBoxStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, checkbox := range backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr {
		backRepoCheckBox.CommitPhaseTwoInstance(backRepo, idx, checkbox)
	}

	return
}

// BackRepoCheckBox.CommitPhaseTwoInstance commits {{structname }} of models.CheckBox to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCheckBox *BackRepoCheckBoxStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, checkbox *models.CheckBox) (Error error) {

	// fetch matching checkboxDB
	if checkboxDB, ok := backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB[idx]; ok {

		checkboxDB.CopyBasicFieldsFromCheckBox(checkbox)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoCheckBox.db.Save(checkboxDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown CheckBox intance %s", checkbox.Name))
		return err
	}

	return
}

// BackRepoCheckBox.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoCheckBox *BackRepoCheckBoxStruct) CheckoutPhaseOne() (Error error) {

	checkboxDBArray := make([]CheckBoxDB, 0)
	_, err := backRepoCheckBox.db.Find(&checkboxDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	checkboxInstancesToBeRemovedFromTheStage := make(map[*models.CheckBox]any)
	for key, value := range backRepoCheckBox.stage.CheckBoxs {
		checkboxInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, checkboxDB := range checkboxDBArray {
		backRepoCheckBox.CheckoutPhaseOneInstance(&checkboxDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		checkbox, ok := backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[checkboxDB.ID]
		if ok {
			delete(checkboxInstancesToBeRemovedFromTheStage, checkbox)
		}
	}

	// remove from stage and back repo's 3 maps all checkboxs that are not in the checkout
	for checkbox := range checkboxInstancesToBeRemovedFromTheStage {
		checkbox.Unstage(backRepoCheckBox.GetStage())

		// remove instance from the back repo 3 maps
		checkboxID := backRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID[checkbox]
		delete(backRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID, checkbox)
		delete(backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB, checkboxID)
		delete(backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr, checkboxID)
	}

	return
}

// CheckoutPhaseOneInstance takes a checkboxDB that has been found in the DB, updates the backRepo and stages the
// models version of the checkboxDB
func (backRepoCheckBox *BackRepoCheckBoxStruct) CheckoutPhaseOneInstance(checkboxDB *CheckBoxDB) (Error error) {

	checkbox, ok := backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[checkboxDB.ID]
	if !ok {
		checkbox = new(models.CheckBox)

		backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[checkboxDB.ID] = checkbox
		backRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID[checkbox] = checkboxDB.ID

		// append model store with the new element
		checkbox.Name = checkboxDB.Name_Data.String
		checkbox.Stage(backRepoCheckBox.GetStage())
	}
	checkboxDB.CopyBasicFieldsToCheckBox(checkbox)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	checkbox.Stage(backRepoCheckBox.GetStage())

	// preserve pointer to checkboxDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_CheckBoxDBID_CheckBoxDB)[checkboxDB hold variable pointers
	checkboxDB_Data := *checkboxDB
	preservedPtrToCheckBox := &checkboxDB_Data
	backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB[checkboxDB.ID] = preservedPtrToCheckBox

	return
}

// BackRepoCheckBox.CheckoutPhaseTwo Checkouts all staged instances of CheckBox to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCheckBox *BackRepoCheckBoxStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, checkboxDB := range backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB {
		backRepoCheckBox.CheckoutPhaseTwoInstance(backRepo, checkboxDB)
	}
	return
}

// BackRepoCheckBox.CheckoutPhaseTwoInstance Checkouts staged instances of CheckBox to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCheckBox *BackRepoCheckBoxStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, checkboxDB *CheckBoxDB) (Error error) {

	checkbox := backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[checkboxDB.ID]

	checkboxDB.DecodePointers(backRepo, checkbox)

	return
}

func (checkboxDB *CheckBoxDB) DecodePointers(backRepo *BackRepoStruct, checkbox *models.CheckBox) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitCheckBox allows commit of a single checkbox (if already staged)
func (backRepo *BackRepoStruct) CommitCheckBox(checkbox *models.CheckBox) {
	backRepo.BackRepoCheckBox.CommitPhaseOneInstance(checkbox)
	if id, ok := backRepo.BackRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID[checkbox]; ok {
		backRepo.BackRepoCheckBox.CommitPhaseTwoInstance(backRepo, id, checkbox)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitCheckBox allows checkout of a single checkbox (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutCheckBox(checkbox *models.CheckBox) {
	// check if the checkbox is staged
	if _, ok := backRepo.BackRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID[checkbox]; ok {

		if id, ok := backRepo.BackRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID[checkbox]; ok {
			var checkboxDB CheckBoxDB
			checkboxDB.ID = id

			if _, err := backRepo.BackRepoCheckBox.db.First(&checkboxDB, id); err != nil {
				log.Fatalln("CheckoutCheckBox : Problem with getting object with id:", id)
			}
			backRepo.BackRepoCheckBox.CheckoutPhaseOneInstance(&checkboxDB)
			backRepo.BackRepoCheckBox.CheckoutPhaseTwoInstance(backRepo, &checkboxDB)
		}
	}
}

// CopyBasicFieldsFromCheckBox
func (checkboxDB *CheckBoxDB) CopyBasicFieldsFromCheckBox(checkbox *models.CheckBox) {
	// insertion point for fields commit

	checkboxDB.Name_Data.String = checkbox.Name
	checkboxDB.Name_Data.Valid = true

	checkboxDB.Value_Data.Bool = checkbox.Value
	checkboxDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromCheckBox_WOP
func (checkboxDB *CheckBoxDB) CopyBasicFieldsFromCheckBox_WOP(checkbox *models.CheckBox_WOP) {
	// insertion point for fields commit

	checkboxDB.Name_Data.String = checkbox.Name
	checkboxDB.Name_Data.Valid = true

	checkboxDB.Value_Data.Bool = checkbox.Value
	checkboxDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromCheckBoxWOP
func (checkboxDB *CheckBoxDB) CopyBasicFieldsFromCheckBoxWOP(checkbox *CheckBoxWOP) {
	// insertion point for fields commit

	checkboxDB.Name_Data.String = checkbox.Name
	checkboxDB.Name_Data.Valid = true

	checkboxDB.Value_Data.Bool = checkbox.Value
	checkboxDB.Value_Data.Valid = true
}

// CopyBasicFieldsToCheckBox
func (checkboxDB *CheckBoxDB) CopyBasicFieldsToCheckBox(checkbox *models.CheckBox) {
	// insertion point for checkout of basic fields (back repo to stage)
	checkbox.Name = checkboxDB.Name_Data.String
	checkbox.Value = checkboxDB.Value_Data.Bool
}

// CopyBasicFieldsToCheckBox_WOP
func (checkboxDB *CheckBoxDB) CopyBasicFieldsToCheckBox_WOP(checkbox *models.CheckBox_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	checkbox.Name = checkboxDB.Name_Data.String
	checkbox.Value = checkboxDB.Value_Data.Bool
}

// CopyBasicFieldsToCheckBoxWOP
func (checkboxDB *CheckBoxDB) CopyBasicFieldsToCheckBoxWOP(checkbox *CheckBoxWOP) {
	checkbox.ID = int(checkboxDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	checkbox.Name = checkboxDB.Name_Data.String
	checkbox.Value = checkboxDB.Value_Data.Bool
}

// Backup generates a json file from a slice of all CheckBoxDB instances in the backrepo
func (backRepoCheckBox *BackRepoCheckBoxStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "CheckBoxDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*CheckBoxDB, 0)
	for _, checkboxDB := range backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB {
		forBackup = append(forBackup, checkboxDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json CheckBox ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json CheckBox file", err.Error())
	}
}

// Backup generates a json file from a slice of all CheckBoxDB instances in the backrepo
func (backRepoCheckBox *BackRepoCheckBoxStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*CheckBoxDB, 0)
	for _, checkboxDB := range backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB {
		forBackup = append(forBackup, checkboxDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("CheckBox")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&CheckBox_Fields, -1)
	for _, checkboxDB := range forBackup {

		var checkboxWOP CheckBoxWOP
		checkboxDB.CopyBasicFieldsToCheckBoxWOP(&checkboxWOP)

		row := sh.AddRow()
		row.WriteStruct(&checkboxWOP, -1)
	}
}

// RestoreXL from the "CheckBox" sheet all CheckBoxDB instances
func (backRepoCheckBox *BackRepoCheckBoxStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoCheckBoxid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["CheckBox"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoCheckBox.rowVisitorCheckBox)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoCheckBox *BackRepoCheckBoxStruct) rowVisitorCheckBox(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var checkboxWOP CheckBoxWOP
		row.ReadStruct(&checkboxWOP)

		// add the unmarshalled struct to the stage
		checkboxDB := new(CheckBoxDB)
		checkboxDB.CopyBasicFieldsFromCheckBoxWOP(&checkboxWOP)

		checkboxDB_ID_atBackupTime := checkboxDB.ID
		checkboxDB.ID = 0
		_, err := backRepoCheckBox.db.Create(checkboxDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB[checkboxDB.ID] = checkboxDB
		BackRepoCheckBoxid_atBckpTime_newID[checkboxDB_ID_atBackupTime] = checkboxDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "CheckBoxDB.json" in dirPath that stores an array
// of CheckBoxDB and stores it in the database
// the map BackRepoCheckBoxid_atBckpTime_newID is updated accordingly
func (backRepoCheckBox *BackRepoCheckBoxStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoCheckBoxid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "CheckBoxDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json CheckBox file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*CheckBoxDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_CheckBoxDBID_CheckBoxDB
	for _, checkboxDB := range forRestore {

		checkboxDB_ID_atBackupTime := checkboxDB.ID
		checkboxDB.ID = 0
		_, err := backRepoCheckBox.db.Create(checkboxDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB[checkboxDB.ID] = checkboxDB
		BackRepoCheckBoxid_atBckpTime_newID[checkboxDB_ID_atBackupTime] = checkboxDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json CheckBox file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<CheckBox>id_atBckpTime_newID
// to compute new index
func (backRepoCheckBox *BackRepoCheckBoxStruct) RestorePhaseTwo() {

	for _, checkboxDB := range backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB {

		// next line of code is to avert unused variable compilation error
		_ = checkboxDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoCheckBox.db.Model(checkboxDB)
		_, err := db.Updates(*checkboxDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoCheckBox.ResetReversePointers commits all staged instances of CheckBox to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCheckBox *BackRepoCheckBoxStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, checkbox := range backRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr {
		backRepoCheckBox.ResetReversePointersInstance(backRepo, idx, checkbox)
	}

	return
}

func (backRepoCheckBox *BackRepoCheckBoxStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, checkbox *models.CheckBox) (Error error) {

	// fetch matching checkboxDB
	if checkboxDB, ok := backRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB[idx]; ok {
		_ = checkboxDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoCheckBoxid_atBckpTime_newID map[uint]uint
