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

	"github.com/fullstack-lang/gongreqif/go/db"
	"github.com/fullstack-lang/gongreqif/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Xhtml_dfn_type_sql sql.NullBool
var dummy_Xhtml_dfn_type_time time.Duration
var dummy_Xhtml_dfn_type_sort sort.Float64Slice

// Xhtml_dfn_typeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model xhtml_dfn_typeAPI
type Xhtml_dfn_typeAPI struct {
	gorm.Model

	models.Xhtml_dfn_type_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Xhtml_dfn_typePointersEncoding Xhtml_dfn_typePointersEncoding
}

// Xhtml_dfn_typePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Xhtml_dfn_typePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Xhtml_dfn_typeDB describes a xhtml_dfn_type in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model xhtml_dfn_typeDB
type Xhtml_dfn_typeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field xhtml_dfn_typeDB.Name
	Name_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Xhtml_dfn_typePointersEncoding
}

// Xhtml_dfn_typeDBs arrays xhtml_dfn_typeDBs
// swagger:response xhtml_dfn_typeDBsResponse
type Xhtml_dfn_typeDBs []Xhtml_dfn_typeDB

// Xhtml_dfn_typeDBResponse provides response
// swagger:response xhtml_dfn_typeDBResponse
type Xhtml_dfn_typeDBResponse struct {
	Xhtml_dfn_typeDB
}

// Xhtml_dfn_typeWOP is a Xhtml_dfn_type without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Xhtml_dfn_typeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Xhtml_dfn_type_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoXhtml_dfn_typeStruct struct {
	// stores Xhtml_dfn_typeDB according to their gorm ID
	Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB map[uint]*Xhtml_dfn_typeDB

	// stores Xhtml_dfn_typeDB ID according to Xhtml_dfn_type address
	Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID map[*models.Xhtml_dfn_type]uint

	// stores Xhtml_dfn_type according to their gorm ID
	Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr map[uint]*models.Xhtml_dfn_type

	db db.DBInterface

	stage *models.Stage
}

func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) GetStage() (stage *models.Stage) {
	stage = backRepoXhtml_dfn_type.stage
	return
}

func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) GetDB() db.DBInterface {
	return backRepoXhtml_dfn_type.db
}

// GetXhtml_dfn_typeDBFromXhtml_dfn_typePtr is a handy function to access the back repo instance from the stage instance
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) GetXhtml_dfn_typeDBFromXhtml_dfn_typePtr(xhtml_dfn_type *models.Xhtml_dfn_type) (xhtml_dfn_typeDB *Xhtml_dfn_typeDB) {
	id := backRepoXhtml_dfn_type.Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID[xhtml_dfn_type]
	xhtml_dfn_typeDB = backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB[id]
	return
}

// BackRepoXhtml_dfn_type.CommitPhaseOne commits all staged instances of Xhtml_dfn_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var xhtml_dfn_types []*models.Xhtml_dfn_type
	for xhtml_dfn_type := range stage.Xhtml_dfn_types {
		xhtml_dfn_types = append(xhtml_dfn_types, xhtml_dfn_type)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(xhtml_dfn_types, func(i, j int) bool {
		return stage.Xhtml_dfn_typeMap_Staged_Order[xhtml_dfn_types[i]] < stage.Xhtml_dfn_typeMap_Staged_Order[xhtml_dfn_types[j]]
	})

	for _, xhtml_dfn_type := range xhtml_dfn_types {
		backRepoXhtml_dfn_type.CommitPhaseOneInstance(xhtml_dfn_type)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, xhtml_dfn_type := range backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr {
		if _, ok := stage.Xhtml_dfn_types[xhtml_dfn_type]; !ok {
			backRepoXhtml_dfn_type.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoXhtml_dfn_type.CommitDeleteInstance commits deletion of Xhtml_dfn_type to the BackRepo
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) CommitDeleteInstance(id uint) (Error error) {

	xhtml_dfn_type := backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr[id]

	// xhtml_dfn_type is not staged anymore, remove xhtml_dfn_typeDB
	xhtml_dfn_typeDB := backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB[id]
	db, _ := backRepoXhtml_dfn_type.db.Unscoped()
	_, err := db.Delete(xhtml_dfn_typeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoXhtml_dfn_type.Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID, xhtml_dfn_type)
	delete(backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr, id)
	delete(backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB, id)

	return
}

// BackRepoXhtml_dfn_type.CommitPhaseOneInstance commits xhtml_dfn_type staged instances of Xhtml_dfn_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) CommitPhaseOneInstance(xhtml_dfn_type *models.Xhtml_dfn_type) (Error error) {

	// check if the xhtml_dfn_type is not commited yet
	if _, ok := backRepoXhtml_dfn_type.Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID[xhtml_dfn_type]; ok {
		return
	}

	// initiate xhtml_dfn_type
	var xhtml_dfn_typeDB Xhtml_dfn_typeDB
	xhtml_dfn_typeDB.CopyBasicFieldsFromXhtml_dfn_type(xhtml_dfn_type)

	_, err := backRepoXhtml_dfn_type.db.Create(&xhtml_dfn_typeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoXhtml_dfn_type.Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID[xhtml_dfn_type] = xhtml_dfn_typeDB.ID
	backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr[xhtml_dfn_typeDB.ID] = xhtml_dfn_type
	backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB[xhtml_dfn_typeDB.ID] = &xhtml_dfn_typeDB

	return
}

// BackRepoXhtml_dfn_type.CommitPhaseTwo commits all staged instances of Xhtml_dfn_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_dfn_type := range backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr {
		backRepoXhtml_dfn_type.CommitPhaseTwoInstance(backRepo, idx, xhtml_dfn_type)
	}

	return
}

// BackRepoXhtml_dfn_type.CommitPhaseTwoInstance commits {{structname }} of models.Xhtml_dfn_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, xhtml_dfn_type *models.Xhtml_dfn_type) (Error error) {

	// fetch matching xhtml_dfn_typeDB
	if xhtml_dfn_typeDB, ok := backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB[idx]; ok {

		xhtml_dfn_typeDB.CopyBasicFieldsFromXhtml_dfn_type(xhtml_dfn_type)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoXhtml_dfn_type.db.Save(xhtml_dfn_typeDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Xhtml_dfn_type intance %s", xhtml_dfn_type.Name))
		return err
	}

	return
}

// BackRepoXhtml_dfn_type.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) CheckoutPhaseOne() (Error error) {

	xhtml_dfn_typeDBArray := make([]Xhtml_dfn_typeDB, 0)
	_, err := backRepoXhtml_dfn_type.db.Find(&xhtml_dfn_typeDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	xhtml_dfn_typeInstancesToBeRemovedFromTheStage := make(map[*models.Xhtml_dfn_type]any)
	for key, value := range backRepoXhtml_dfn_type.stage.Xhtml_dfn_types {
		xhtml_dfn_typeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, xhtml_dfn_typeDB := range xhtml_dfn_typeDBArray {
		backRepoXhtml_dfn_type.CheckoutPhaseOneInstance(&xhtml_dfn_typeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		xhtml_dfn_type, ok := backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr[xhtml_dfn_typeDB.ID]
		if ok {
			delete(xhtml_dfn_typeInstancesToBeRemovedFromTheStage, xhtml_dfn_type)
		}
	}

	// remove from stage and back repo's 3 maps all xhtml_dfn_types that are not in the checkout
	for xhtml_dfn_type := range xhtml_dfn_typeInstancesToBeRemovedFromTheStage {
		xhtml_dfn_type.Unstage(backRepoXhtml_dfn_type.GetStage())

		// remove instance from the back repo 3 maps
		xhtml_dfn_typeID := backRepoXhtml_dfn_type.Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID[xhtml_dfn_type]
		delete(backRepoXhtml_dfn_type.Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID, xhtml_dfn_type)
		delete(backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB, xhtml_dfn_typeID)
		delete(backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr, xhtml_dfn_typeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a xhtml_dfn_typeDB that has been found in the DB, updates the backRepo and stages the
// models version of the xhtml_dfn_typeDB
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) CheckoutPhaseOneInstance(xhtml_dfn_typeDB *Xhtml_dfn_typeDB) (Error error) {

	xhtml_dfn_type, ok := backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr[xhtml_dfn_typeDB.ID]
	if !ok {
		xhtml_dfn_type = new(models.Xhtml_dfn_type)

		backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr[xhtml_dfn_typeDB.ID] = xhtml_dfn_type
		backRepoXhtml_dfn_type.Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID[xhtml_dfn_type] = xhtml_dfn_typeDB.ID

		// append model store with the new element
		xhtml_dfn_type.Name = xhtml_dfn_typeDB.Name_Data.String
		xhtml_dfn_type.Stage(backRepoXhtml_dfn_type.GetStage())
	}
	xhtml_dfn_typeDB.CopyBasicFieldsToXhtml_dfn_type(xhtml_dfn_type)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	xhtml_dfn_type.Stage(backRepoXhtml_dfn_type.GetStage())

	// preserve pointer to xhtml_dfn_typeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB)[xhtml_dfn_typeDB hold variable pointers
	xhtml_dfn_typeDB_Data := *xhtml_dfn_typeDB
	preservedPtrToXhtml_dfn_type := &xhtml_dfn_typeDB_Data
	backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB[xhtml_dfn_typeDB.ID] = preservedPtrToXhtml_dfn_type

	return
}

// BackRepoXhtml_dfn_type.CheckoutPhaseTwo Checkouts all staged instances of Xhtml_dfn_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, xhtml_dfn_typeDB := range backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB {
		backRepoXhtml_dfn_type.CheckoutPhaseTwoInstance(backRepo, xhtml_dfn_typeDB)
	}
	return
}

// BackRepoXhtml_dfn_type.CheckoutPhaseTwoInstance Checkouts staged instances of Xhtml_dfn_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, xhtml_dfn_typeDB *Xhtml_dfn_typeDB) (Error error) {

	xhtml_dfn_type := backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr[xhtml_dfn_typeDB.ID]

	xhtml_dfn_typeDB.DecodePointers(backRepo, xhtml_dfn_type)

	return
}

func (xhtml_dfn_typeDB *Xhtml_dfn_typeDB) DecodePointers(backRepo *BackRepoStruct, xhtml_dfn_type *models.Xhtml_dfn_type) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitXhtml_dfn_type allows commit of a single xhtml_dfn_type (if already staged)
func (backRepo *BackRepoStruct) CommitXhtml_dfn_type(xhtml_dfn_type *models.Xhtml_dfn_type) {
	backRepo.BackRepoXhtml_dfn_type.CommitPhaseOneInstance(xhtml_dfn_type)
	if id, ok := backRepo.BackRepoXhtml_dfn_type.Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID[xhtml_dfn_type]; ok {
		backRepo.BackRepoXhtml_dfn_type.CommitPhaseTwoInstance(backRepo, id, xhtml_dfn_type)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitXhtml_dfn_type allows checkout of a single xhtml_dfn_type (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutXhtml_dfn_type(xhtml_dfn_type *models.Xhtml_dfn_type) {
	// check if the xhtml_dfn_type is staged
	if _, ok := backRepo.BackRepoXhtml_dfn_type.Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID[xhtml_dfn_type]; ok {

		if id, ok := backRepo.BackRepoXhtml_dfn_type.Map_Xhtml_dfn_typePtr_Xhtml_dfn_typeDBID[xhtml_dfn_type]; ok {
			var xhtml_dfn_typeDB Xhtml_dfn_typeDB
			xhtml_dfn_typeDB.ID = id

			if _, err := backRepo.BackRepoXhtml_dfn_type.db.First(&xhtml_dfn_typeDB, id); err != nil {
				log.Fatalln("CheckoutXhtml_dfn_type : Problem with getting object with id:", id)
			}
			backRepo.BackRepoXhtml_dfn_type.CheckoutPhaseOneInstance(&xhtml_dfn_typeDB)
			backRepo.BackRepoXhtml_dfn_type.CheckoutPhaseTwoInstance(backRepo, &xhtml_dfn_typeDB)
		}
	}
}

// CopyBasicFieldsFromXhtml_dfn_type
func (xhtml_dfn_typeDB *Xhtml_dfn_typeDB) CopyBasicFieldsFromXhtml_dfn_type(xhtml_dfn_type *models.Xhtml_dfn_type) {
	// insertion point for fields commit

	xhtml_dfn_typeDB.Name_Data.String = xhtml_dfn_type.Name
	xhtml_dfn_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_dfn_type_WOP
func (xhtml_dfn_typeDB *Xhtml_dfn_typeDB) CopyBasicFieldsFromXhtml_dfn_type_WOP(xhtml_dfn_type *models.Xhtml_dfn_type_WOP) {
	// insertion point for fields commit

	xhtml_dfn_typeDB.Name_Data.String = xhtml_dfn_type.Name
	xhtml_dfn_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_dfn_typeWOP
func (xhtml_dfn_typeDB *Xhtml_dfn_typeDB) CopyBasicFieldsFromXhtml_dfn_typeWOP(xhtml_dfn_type *Xhtml_dfn_typeWOP) {
	// insertion point for fields commit

	xhtml_dfn_typeDB.Name_Data.String = xhtml_dfn_type.Name
	xhtml_dfn_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToXhtml_dfn_type
func (xhtml_dfn_typeDB *Xhtml_dfn_typeDB) CopyBasicFieldsToXhtml_dfn_type(xhtml_dfn_type *models.Xhtml_dfn_type) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_dfn_type.Name = xhtml_dfn_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_dfn_type_WOP
func (xhtml_dfn_typeDB *Xhtml_dfn_typeDB) CopyBasicFieldsToXhtml_dfn_type_WOP(xhtml_dfn_type *models.Xhtml_dfn_type_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_dfn_type.Name = xhtml_dfn_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_dfn_typeWOP
func (xhtml_dfn_typeDB *Xhtml_dfn_typeDB) CopyBasicFieldsToXhtml_dfn_typeWOP(xhtml_dfn_type *Xhtml_dfn_typeWOP) {
	xhtml_dfn_type.ID = int(xhtml_dfn_typeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_dfn_type.Name = xhtml_dfn_typeDB.Name_Data.String
}

// Backup generates a json file from a slice of all Xhtml_dfn_typeDB instances in the backrepo
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Xhtml_dfn_typeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_dfn_typeDB, 0)
	for _, xhtml_dfn_typeDB := range backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB {
		forBackup = append(forBackup, xhtml_dfn_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Xhtml_dfn_type ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Xhtml_dfn_type file", err.Error())
	}
}

// Backup generates a json file from a slice of all Xhtml_dfn_typeDB instances in the backrepo
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_dfn_typeDB, 0)
	for _, xhtml_dfn_typeDB := range backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB {
		forBackup = append(forBackup, xhtml_dfn_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Xhtml_dfn_type")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Xhtml_dfn_type_Fields, -1)
	for _, xhtml_dfn_typeDB := range forBackup {

		var xhtml_dfn_typeWOP Xhtml_dfn_typeWOP
		xhtml_dfn_typeDB.CopyBasicFieldsToXhtml_dfn_typeWOP(&xhtml_dfn_typeWOP)

		row := sh.AddRow()
		row.WriteStruct(&xhtml_dfn_typeWOP, -1)
	}
}

// RestoreXL from the "Xhtml_dfn_type" sheet all Xhtml_dfn_typeDB instances
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoXhtml_dfn_typeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Xhtml_dfn_type"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoXhtml_dfn_type.rowVisitorXhtml_dfn_type)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) rowVisitorXhtml_dfn_type(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var xhtml_dfn_typeWOP Xhtml_dfn_typeWOP
		row.ReadStruct(&xhtml_dfn_typeWOP)

		// add the unmarshalled struct to the stage
		xhtml_dfn_typeDB := new(Xhtml_dfn_typeDB)
		xhtml_dfn_typeDB.CopyBasicFieldsFromXhtml_dfn_typeWOP(&xhtml_dfn_typeWOP)

		xhtml_dfn_typeDB_ID_atBackupTime := xhtml_dfn_typeDB.ID
		xhtml_dfn_typeDB.ID = 0
		_, err := backRepoXhtml_dfn_type.db.Create(xhtml_dfn_typeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB[xhtml_dfn_typeDB.ID] = xhtml_dfn_typeDB
		BackRepoXhtml_dfn_typeid_atBckpTime_newID[xhtml_dfn_typeDB_ID_atBackupTime] = xhtml_dfn_typeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Xhtml_dfn_typeDB.json" in dirPath that stores an array
// of Xhtml_dfn_typeDB and stores it in the database
// the map BackRepoXhtml_dfn_typeid_atBckpTime_newID is updated accordingly
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoXhtml_dfn_typeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Xhtml_dfn_typeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Xhtml_dfn_type file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Xhtml_dfn_typeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB
	for _, xhtml_dfn_typeDB := range forRestore {

		xhtml_dfn_typeDB_ID_atBackupTime := xhtml_dfn_typeDB.ID
		xhtml_dfn_typeDB.ID = 0
		_, err := backRepoXhtml_dfn_type.db.Create(xhtml_dfn_typeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB[xhtml_dfn_typeDB.ID] = xhtml_dfn_typeDB
		BackRepoXhtml_dfn_typeid_atBckpTime_newID[xhtml_dfn_typeDB_ID_atBackupTime] = xhtml_dfn_typeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Xhtml_dfn_type file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Xhtml_dfn_type>id_atBckpTime_newID
// to compute new index
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) RestorePhaseTwo() {

	for _, xhtml_dfn_typeDB := range backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB {

		// next line of code is to avert unused variable compilation error
		_ = xhtml_dfn_typeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoXhtml_dfn_type.db.Model(xhtml_dfn_typeDB)
		_, err := db.Updates(*xhtml_dfn_typeDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoXhtml_dfn_type.ResetReversePointers commits all staged instances of Xhtml_dfn_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_dfn_type := range backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr {
		backRepoXhtml_dfn_type.ResetReversePointersInstance(backRepo, idx, xhtml_dfn_type)
	}

	return
}

func (backRepoXhtml_dfn_type *BackRepoXhtml_dfn_typeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, xhtml_dfn_type *models.Xhtml_dfn_type) (Error error) {

	// fetch matching xhtml_dfn_typeDB
	if xhtml_dfn_typeDB, ok := backRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typeDB[idx]; ok {
		_ = xhtml_dfn_typeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoXhtml_dfn_typeid_atBckpTime_newID map[uint]uint
