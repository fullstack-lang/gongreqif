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

	"github.com/fullstack-lang/gongreqif/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Xhtml_samp_type_sql sql.NullBool
var dummy_Xhtml_samp_type_time time.Duration
var dummy_Xhtml_samp_type_sort sort.Float64Slice

// Xhtml_samp_typeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model xhtml_samp_typeAPI
type Xhtml_samp_typeAPI struct {
	gorm.Model

	models.Xhtml_samp_type_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Xhtml_samp_typePointersEncoding Xhtml_samp_typePointersEncoding
}

// Xhtml_samp_typePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Xhtml_samp_typePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Xhtml_samp_typeDB describes a xhtml_samp_type in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model xhtml_samp_typeDB
type Xhtml_samp_typeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field xhtml_samp_typeDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Xhtml_samp_typePointersEncoding
}

// Xhtml_samp_typeDBs arrays xhtml_samp_typeDBs
// swagger:response xhtml_samp_typeDBsResponse
type Xhtml_samp_typeDBs []Xhtml_samp_typeDB

// Xhtml_samp_typeDBResponse provides response
// swagger:response xhtml_samp_typeDBResponse
type Xhtml_samp_typeDBResponse struct {
	Xhtml_samp_typeDB
}

// Xhtml_samp_typeWOP is a Xhtml_samp_type without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Xhtml_samp_typeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Xhtml_samp_type_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoXhtml_samp_typeStruct struct {
	// stores Xhtml_samp_typeDB according to their gorm ID
	Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB map[uint]*Xhtml_samp_typeDB

	// stores Xhtml_samp_typeDB ID according to Xhtml_samp_type address
	Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID map[*models.Xhtml_samp_type]uint

	// stores Xhtml_samp_type according to their gorm ID
	Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr map[uint]*models.Xhtml_samp_type

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoXhtml_samp_type.stage
	return
}

func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) GetDB() *gorm.DB {
	return backRepoXhtml_samp_type.db
}

// GetXhtml_samp_typeDBFromXhtml_samp_typePtr is a handy function to access the back repo instance from the stage instance
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) GetXhtml_samp_typeDBFromXhtml_samp_typePtr(xhtml_samp_type *models.Xhtml_samp_type) (xhtml_samp_typeDB *Xhtml_samp_typeDB) {
	id := backRepoXhtml_samp_type.Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID[xhtml_samp_type]
	xhtml_samp_typeDB = backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB[id]
	return
}

// BackRepoXhtml_samp_type.CommitPhaseOne commits all staged instances of Xhtml_samp_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for xhtml_samp_type := range stage.Xhtml_samp_types {
		backRepoXhtml_samp_type.CommitPhaseOneInstance(xhtml_samp_type)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, xhtml_samp_type := range backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr {
		if _, ok := stage.Xhtml_samp_types[xhtml_samp_type]; !ok {
			backRepoXhtml_samp_type.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoXhtml_samp_type.CommitDeleteInstance commits deletion of Xhtml_samp_type to the BackRepo
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) CommitDeleteInstance(id uint) (Error error) {

	xhtml_samp_type := backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr[id]

	// xhtml_samp_type is not staged anymore, remove xhtml_samp_typeDB
	xhtml_samp_typeDB := backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB[id]
	query := backRepoXhtml_samp_type.db.Unscoped().Delete(&xhtml_samp_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoXhtml_samp_type.Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID, xhtml_samp_type)
	delete(backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr, id)
	delete(backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB, id)

	return
}

// BackRepoXhtml_samp_type.CommitPhaseOneInstance commits xhtml_samp_type staged instances of Xhtml_samp_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) CommitPhaseOneInstance(xhtml_samp_type *models.Xhtml_samp_type) (Error error) {

	// check if the xhtml_samp_type is not commited yet
	if _, ok := backRepoXhtml_samp_type.Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID[xhtml_samp_type]; ok {
		return
	}

	// initiate xhtml_samp_type
	var xhtml_samp_typeDB Xhtml_samp_typeDB
	xhtml_samp_typeDB.CopyBasicFieldsFromXhtml_samp_type(xhtml_samp_type)

	query := backRepoXhtml_samp_type.db.Create(&xhtml_samp_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoXhtml_samp_type.Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID[xhtml_samp_type] = xhtml_samp_typeDB.ID
	backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr[xhtml_samp_typeDB.ID] = xhtml_samp_type
	backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB[xhtml_samp_typeDB.ID] = &xhtml_samp_typeDB

	return
}

// BackRepoXhtml_samp_type.CommitPhaseTwo commits all staged instances of Xhtml_samp_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_samp_type := range backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr {
		backRepoXhtml_samp_type.CommitPhaseTwoInstance(backRepo, idx, xhtml_samp_type)
	}

	return
}

// BackRepoXhtml_samp_type.CommitPhaseTwoInstance commits {{structname }} of models.Xhtml_samp_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, xhtml_samp_type *models.Xhtml_samp_type) (Error error) {

	// fetch matching xhtml_samp_typeDB
	if xhtml_samp_typeDB, ok := backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB[idx]; ok {

		xhtml_samp_typeDB.CopyBasicFieldsFromXhtml_samp_type(xhtml_samp_type)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoXhtml_samp_type.db.Save(&xhtml_samp_typeDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Xhtml_samp_type intance %s", xhtml_samp_type.Name))
		return err
	}

	return
}

// BackRepoXhtml_samp_type.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) CheckoutPhaseOne() (Error error) {

	xhtml_samp_typeDBArray := make([]Xhtml_samp_typeDB, 0)
	query := backRepoXhtml_samp_type.db.Find(&xhtml_samp_typeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	xhtml_samp_typeInstancesToBeRemovedFromTheStage := make(map[*models.Xhtml_samp_type]any)
	for key, value := range backRepoXhtml_samp_type.stage.Xhtml_samp_types {
		xhtml_samp_typeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, xhtml_samp_typeDB := range xhtml_samp_typeDBArray {
		backRepoXhtml_samp_type.CheckoutPhaseOneInstance(&xhtml_samp_typeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		xhtml_samp_type, ok := backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr[xhtml_samp_typeDB.ID]
		if ok {
			delete(xhtml_samp_typeInstancesToBeRemovedFromTheStage, xhtml_samp_type)
		}
	}

	// remove from stage and back repo's 3 maps all xhtml_samp_types that are not in the checkout
	for xhtml_samp_type := range xhtml_samp_typeInstancesToBeRemovedFromTheStage {
		xhtml_samp_type.Unstage(backRepoXhtml_samp_type.GetStage())

		// remove instance from the back repo 3 maps
		xhtml_samp_typeID := backRepoXhtml_samp_type.Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID[xhtml_samp_type]
		delete(backRepoXhtml_samp_type.Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID, xhtml_samp_type)
		delete(backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB, xhtml_samp_typeID)
		delete(backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr, xhtml_samp_typeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a xhtml_samp_typeDB that has been found in the DB, updates the backRepo and stages the
// models version of the xhtml_samp_typeDB
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) CheckoutPhaseOneInstance(xhtml_samp_typeDB *Xhtml_samp_typeDB) (Error error) {

	xhtml_samp_type, ok := backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr[xhtml_samp_typeDB.ID]
	if !ok {
		xhtml_samp_type = new(models.Xhtml_samp_type)

		backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr[xhtml_samp_typeDB.ID] = xhtml_samp_type
		backRepoXhtml_samp_type.Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID[xhtml_samp_type] = xhtml_samp_typeDB.ID

		// append model store with the new element
		xhtml_samp_type.Name = xhtml_samp_typeDB.Name_Data.String
		xhtml_samp_type.Stage(backRepoXhtml_samp_type.GetStage())
	}
	xhtml_samp_typeDB.CopyBasicFieldsToXhtml_samp_type(xhtml_samp_type)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	xhtml_samp_type.Stage(backRepoXhtml_samp_type.GetStage())

	// preserve pointer to xhtml_samp_typeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB)[xhtml_samp_typeDB hold variable pointers
	xhtml_samp_typeDB_Data := *xhtml_samp_typeDB
	preservedPtrToXhtml_samp_type := &xhtml_samp_typeDB_Data
	backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB[xhtml_samp_typeDB.ID] = preservedPtrToXhtml_samp_type

	return
}

// BackRepoXhtml_samp_type.CheckoutPhaseTwo Checkouts all staged instances of Xhtml_samp_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, xhtml_samp_typeDB := range backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB {
		backRepoXhtml_samp_type.CheckoutPhaseTwoInstance(backRepo, xhtml_samp_typeDB)
	}
	return
}

// BackRepoXhtml_samp_type.CheckoutPhaseTwoInstance Checkouts staged instances of Xhtml_samp_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, xhtml_samp_typeDB *Xhtml_samp_typeDB) (Error error) {

	xhtml_samp_type := backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr[xhtml_samp_typeDB.ID]

	xhtml_samp_typeDB.DecodePointers(backRepo, xhtml_samp_type)

	return
}

func (xhtml_samp_typeDB *Xhtml_samp_typeDB) DecodePointers(backRepo *BackRepoStruct, xhtml_samp_type *models.Xhtml_samp_type) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitXhtml_samp_type allows commit of a single xhtml_samp_type (if already staged)
func (backRepo *BackRepoStruct) CommitXhtml_samp_type(xhtml_samp_type *models.Xhtml_samp_type) {
	backRepo.BackRepoXhtml_samp_type.CommitPhaseOneInstance(xhtml_samp_type)
	if id, ok := backRepo.BackRepoXhtml_samp_type.Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID[xhtml_samp_type]; ok {
		backRepo.BackRepoXhtml_samp_type.CommitPhaseTwoInstance(backRepo, id, xhtml_samp_type)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitXhtml_samp_type allows checkout of a single xhtml_samp_type (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutXhtml_samp_type(xhtml_samp_type *models.Xhtml_samp_type) {
	// check if the xhtml_samp_type is staged
	if _, ok := backRepo.BackRepoXhtml_samp_type.Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID[xhtml_samp_type]; ok {

		if id, ok := backRepo.BackRepoXhtml_samp_type.Map_Xhtml_samp_typePtr_Xhtml_samp_typeDBID[xhtml_samp_type]; ok {
			var xhtml_samp_typeDB Xhtml_samp_typeDB
			xhtml_samp_typeDB.ID = id

			if err := backRepo.BackRepoXhtml_samp_type.db.First(&xhtml_samp_typeDB, id).Error; err != nil {
				log.Fatalln("CheckoutXhtml_samp_type : Problem with getting object with id:", id)
			}
			backRepo.BackRepoXhtml_samp_type.CheckoutPhaseOneInstance(&xhtml_samp_typeDB)
			backRepo.BackRepoXhtml_samp_type.CheckoutPhaseTwoInstance(backRepo, &xhtml_samp_typeDB)
		}
	}
}

// CopyBasicFieldsFromXhtml_samp_type
func (xhtml_samp_typeDB *Xhtml_samp_typeDB) CopyBasicFieldsFromXhtml_samp_type(xhtml_samp_type *models.Xhtml_samp_type) {
	// insertion point for fields commit

	xhtml_samp_typeDB.Name_Data.String = xhtml_samp_type.Name
	xhtml_samp_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_samp_type_WOP
func (xhtml_samp_typeDB *Xhtml_samp_typeDB) CopyBasicFieldsFromXhtml_samp_type_WOP(xhtml_samp_type *models.Xhtml_samp_type_WOP) {
	// insertion point for fields commit

	xhtml_samp_typeDB.Name_Data.String = xhtml_samp_type.Name
	xhtml_samp_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_samp_typeWOP
func (xhtml_samp_typeDB *Xhtml_samp_typeDB) CopyBasicFieldsFromXhtml_samp_typeWOP(xhtml_samp_type *Xhtml_samp_typeWOP) {
	// insertion point for fields commit

	xhtml_samp_typeDB.Name_Data.String = xhtml_samp_type.Name
	xhtml_samp_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToXhtml_samp_type
func (xhtml_samp_typeDB *Xhtml_samp_typeDB) CopyBasicFieldsToXhtml_samp_type(xhtml_samp_type *models.Xhtml_samp_type) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_samp_type.Name = xhtml_samp_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_samp_type_WOP
func (xhtml_samp_typeDB *Xhtml_samp_typeDB) CopyBasicFieldsToXhtml_samp_type_WOP(xhtml_samp_type *models.Xhtml_samp_type_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_samp_type.Name = xhtml_samp_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_samp_typeWOP
func (xhtml_samp_typeDB *Xhtml_samp_typeDB) CopyBasicFieldsToXhtml_samp_typeWOP(xhtml_samp_type *Xhtml_samp_typeWOP) {
	xhtml_samp_type.ID = int(xhtml_samp_typeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_samp_type.Name = xhtml_samp_typeDB.Name_Data.String
}

// Backup generates a json file from a slice of all Xhtml_samp_typeDB instances in the backrepo
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Xhtml_samp_typeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_samp_typeDB, 0)
	for _, xhtml_samp_typeDB := range backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB {
		forBackup = append(forBackup, xhtml_samp_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Xhtml_samp_type ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Xhtml_samp_type file", err.Error())
	}
}

// Backup generates a json file from a slice of all Xhtml_samp_typeDB instances in the backrepo
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_samp_typeDB, 0)
	for _, xhtml_samp_typeDB := range backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB {
		forBackup = append(forBackup, xhtml_samp_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Xhtml_samp_type")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Xhtml_samp_type_Fields, -1)
	for _, xhtml_samp_typeDB := range forBackup {

		var xhtml_samp_typeWOP Xhtml_samp_typeWOP
		xhtml_samp_typeDB.CopyBasicFieldsToXhtml_samp_typeWOP(&xhtml_samp_typeWOP)

		row := sh.AddRow()
		row.WriteStruct(&xhtml_samp_typeWOP, -1)
	}
}

// RestoreXL from the "Xhtml_samp_type" sheet all Xhtml_samp_typeDB instances
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoXhtml_samp_typeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Xhtml_samp_type"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoXhtml_samp_type.rowVisitorXhtml_samp_type)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) rowVisitorXhtml_samp_type(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var xhtml_samp_typeWOP Xhtml_samp_typeWOP
		row.ReadStruct(&xhtml_samp_typeWOP)

		// add the unmarshalled struct to the stage
		xhtml_samp_typeDB := new(Xhtml_samp_typeDB)
		xhtml_samp_typeDB.CopyBasicFieldsFromXhtml_samp_typeWOP(&xhtml_samp_typeWOP)

		xhtml_samp_typeDB_ID_atBackupTime := xhtml_samp_typeDB.ID
		xhtml_samp_typeDB.ID = 0
		query := backRepoXhtml_samp_type.db.Create(xhtml_samp_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB[xhtml_samp_typeDB.ID] = xhtml_samp_typeDB
		BackRepoXhtml_samp_typeid_atBckpTime_newID[xhtml_samp_typeDB_ID_atBackupTime] = xhtml_samp_typeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Xhtml_samp_typeDB.json" in dirPath that stores an array
// of Xhtml_samp_typeDB and stores it in the database
// the map BackRepoXhtml_samp_typeid_atBckpTime_newID is updated accordingly
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoXhtml_samp_typeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Xhtml_samp_typeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Xhtml_samp_type file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Xhtml_samp_typeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB
	for _, xhtml_samp_typeDB := range forRestore {

		xhtml_samp_typeDB_ID_atBackupTime := xhtml_samp_typeDB.ID
		xhtml_samp_typeDB.ID = 0
		query := backRepoXhtml_samp_type.db.Create(xhtml_samp_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB[xhtml_samp_typeDB.ID] = xhtml_samp_typeDB
		BackRepoXhtml_samp_typeid_atBckpTime_newID[xhtml_samp_typeDB_ID_atBackupTime] = xhtml_samp_typeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Xhtml_samp_type file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Xhtml_samp_type>id_atBckpTime_newID
// to compute new index
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) RestorePhaseTwo() {

	for _, xhtml_samp_typeDB := range backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB {

		// next line of code is to avert unused variable compilation error
		_ = xhtml_samp_typeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoXhtml_samp_type.db.Model(xhtml_samp_typeDB).Updates(*xhtml_samp_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoXhtml_samp_type.ResetReversePointers commits all staged instances of Xhtml_samp_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_samp_type := range backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr {
		backRepoXhtml_samp_type.ResetReversePointersInstance(backRepo, idx, xhtml_samp_type)
	}

	return
}

func (backRepoXhtml_samp_type *BackRepoXhtml_samp_typeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, xhtml_samp_type *models.Xhtml_samp_type) (Error error) {

	// fetch matching xhtml_samp_typeDB
	if xhtml_samp_typeDB, ok := backRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typeDB[idx]; ok {
		_ = xhtml_samp_typeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoXhtml_samp_typeid_atBckpTime_newID map[uint]uint