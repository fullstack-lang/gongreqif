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
var dummy_Xhtml_col_type_sql sql.NullBool
var dummy_Xhtml_col_type_time time.Duration
var dummy_Xhtml_col_type_sort sort.Float64Slice

// Xhtml_col_typeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model xhtml_col_typeAPI
type Xhtml_col_typeAPI struct {
	gorm.Model

	models.Xhtml_col_type_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Xhtml_col_typePointersEncoding Xhtml_col_typePointersEncoding
}

// Xhtml_col_typePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Xhtml_col_typePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Xhtml_col_typeDB describes a xhtml_col_type in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model xhtml_col_typeDB
type Xhtml_col_typeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field xhtml_col_typeDB.Name
	Name_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Xhtml_col_typePointersEncoding
}

// Xhtml_col_typeDBs arrays xhtml_col_typeDBs
// swagger:response xhtml_col_typeDBsResponse
type Xhtml_col_typeDBs []Xhtml_col_typeDB

// Xhtml_col_typeDBResponse provides response
// swagger:response xhtml_col_typeDBResponse
type Xhtml_col_typeDBResponse struct {
	Xhtml_col_typeDB
}

// Xhtml_col_typeWOP is a Xhtml_col_type without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Xhtml_col_typeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Xhtml_col_type_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoXhtml_col_typeStruct struct {
	// stores Xhtml_col_typeDB according to their gorm ID
	Map_Xhtml_col_typeDBID_Xhtml_col_typeDB map[uint]*Xhtml_col_typeDB

	// stores Xhtml_col_typeDB ID according to Xhtml_col_type address
	Map_Xhtml_col_typePtr_Xhtml_col_typeDBID map[*models.Xhtml_col_type]uint

	// stores Xhtml_col_type according to their gorm ID
	Map_Xhtml_col_typeDBID_Xhtml_col_typePtr map[uint]*models.Xhtml_col_type

	db db.DBInterface

	stage *models.Stage
}

func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) GetStage() (stage *models.Stage) {
	stage = backRepoXhtml_col_type.stage
	return
}

func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) GetDB() db.DBInterface {
	return backRepoXhtml_col_type.db
}

// GetXhtml_col_typeDBFromXhtml_col_typePtr is a handy function to access the back repo instance from the stage instance
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) GetXhtml_col_typeDBFromXhtml_col_typePtr(xhtml_col_type *models.Xhtml_col_type) (xhtml_col_typeDB *Xhtml_col_typeDB) {
	id := backRepoXhtml_col_type.Map_Xhtml_col_typePtr_Xhtml_col_typeDBID[xhtml_col_type]
	xhtml_col_typeDB = backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB[id]
	return
}

// BackRepoXhtml_col_type.CommitPhaseOne commits all staged instances of Xhtml_col_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var xhtml_col_types []*models.Xhtml_col_type
	for xhtml_col_type := range stage.Xhtml_col_types {
		xhtml_col_types = append(xhtml_col_types, xhtml_col_type)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(xhtml_col_types, func(i, j int) bool {
		return stage.Xhtml_col_typeMap_Staged_Order[xhtml_col_types[i]] < stage.Xhtml_col_typeMap_Staged_Order[xhtml_col_types[j]]
	})

	for _, xhtml_col_type := range xhtml_col_types {
		backRepoXhtml_col_type.CommitPhaseOneInstance(xhtml_col_type)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, xhtml_col_type := range backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr {
		if _, ok := stage.Xhtml_col_types[xhtml_col_type]; !ok {
			backRepoXhtml_col_type.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoXhtml_col_type.CommitDeleteInstance commits deletion of Xhtml_col_type to the BackRepo
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) CommitDeleteInstance(id uint) (Error error) {

	xhtml_col_type := backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr[id]

	// xhtml_col_type is not staged anymore, remove xhtml_col_typeDB
	xhtml_col_typeDB := backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB[id]
	db, _ := backRepoXhtml_col_type.db.Unscoped()
	_, err := db.Delete(xhtml_col_typeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoXhtml_col_type.Map_Xhtml_col_typePtr_Xhtml_col_typeDBID, xhtml_col_type)
	delete(backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr, id)
	delete(backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB, id)

	return
}

// BackRepoXhtml_col_type.CommitPhaseOneInstance commits xhtml_col_type staged instances of Xhtml_col_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) CommitPhaseOneInstance(xhtml_col_type *models.Xhtml_col_type) (Error error) {

	// check if the xhtml_col_type is not commited yet
	if _, ok := backRepoXhtml_col_type.Map_Xhtml_col_typePtr_Xhtml_col_typeDBID[xhtml_col_type]; ok {
		return
	}

	// initiate xhtml_col_type
	var xhtml_col_typeDB Xhtml_col_typeDB
	xhtml_col_typeDB.CopyBasicFieldsFromXhtml_col_type(xhtml_col_type)

	_, err := backRepoXhtml_col_type.db.Create(&xhtml_col_typeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoXhtml_col_type.Map_Xhtml_col_typePtr_Xhtml_col_typeDBID[xhtml_col_type] = xhtml_col_typeDB.ID
	backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr[xhtml_col_typeDB.ID] = xhtml_col_type
	backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB[xhtml_col_typeDB.ID] = &xhtml_col_typeDB

	return
}

// BackRepoXhtml_col_type.CommitPhaseTwo commits all staged instances of Xhtml_col_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_col_type := range backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr {
		backRepoXhtml_col_type.CommitPhaseTwoInstance(backRepo, idx, xhtml_col_type)
	}

	return
}

// BackRepoXhtml_col_type.CommitPhaseTwoInstance commits {{structname }} of models.Xhtml_col_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, xhtml_col_type *models.Xhtml_col_type) (Error error) {

	// fetch matching xhtml_col_typeDB
	if xhtml_col_typeDB, ok := backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB[idx]; ok {

		xhtml_col_typeDB.CopyBasicFieldsFromXhtml_col_type(xhtml_col_type)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoXhtml_col_type.db.Save(xhtml_col_typeDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Xhtml_col_type intance %s", xhtml_col_type.Name))
		return err
	}

	return
}

// BackRepoXhtml_col_type.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) CheckoutPhaseOne() (Error error) {

	xhtml_col_typeDBArray := make([]Xhtml_col_typeDB, 0)
	_, err := backRepoXhtml_col_type.db.Find(&xhtml_col_typeDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	xhtml_col_typeInstancesToBeRemovedFromTheStage := make(map[*models.Xhtml_col_type]any)
	for key, value := range backRepoXhtml_col_type.stage.Xhtml_col_types {
		xhtml_col_typeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, xhtml_col_typeDB := range xhtml_col_typeDBArray {
		backRepoXhtml_col_type.CheckoutPhaseOneInstance(&xhtml_col_typeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		xhtml_col_type, ok := backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr[xhtml_col_typeDB.ID]
		if ok {
			delete(xhtml_col_typeInstancesToBeRemovedFromTheStage, xhtml_col_type)
		}
	}

	// remove from stage and back repo's 3 maps all xhtml_col_types that are not in the checkout
	for xhtml_col_type := range xhtml_col_typeInstancesToBeRemovedFromTheStage {
		xhtml_col_type.Unstage(backRepoXhtml_col_type.GetStage())

		// remove instance from the back repo 3 maps
		xhtml_col_typeID := backRepoXhtml_col_type.Map_Xhtml_col_typePtr_Xhtml_col_typeDBID[xhtml_col_type]
		delete(backRepoXhtml_col_type.Map_Xhtml_col_typePtr_Xhtml_col_typeDBID, xhtml_col_type)
		delete(backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB, xhtml_col_typeID)
		delete(backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr, xhtml_col_typeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a xhtml_col_typeDB that has been found in the DB, updates the backRepo and stages the
// models version of the xhtml_col_typeDB
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) CheckoutPhaseOneInstance(xhtml_col_typeDB *Xhtml_col_typeDB) (Error error) {

	xhtml_col_type, ok := backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr[xhtml_col_typeDB.ID]
	if !ok {
		xhtml_col_type = new(models.Xhtml_col_type)

		backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr[xhtml_col_typeDB.ID] = xhtml_col_type
		backRepoXhtml_col_type.Map_Xhtml_col_typePtr_Xhtml_col_typeDBID[xhtml_col_type] = xhtml_col_typeDB.ID

		// append model store with the new element
		xhtml_col_type.Name = xhtml_col_typeDB.Name_Data.String
		xhtml_col_type.Stage(backRepoXhtml_col_type.GetStage())
	}
	xhtml_col_typeDB.CopyBasicFieldsToXhtml_col_type(xhtml_col_type)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	xhtml_col_type.Stage(backRepoXhtml_col_type.GetStage())

	// preserve pointer to xhtml_col_typeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Xhtml_col_typeDBID_Xhtml_col_typeDB)[xhtml_col_typeDB hold variable pointers
	xhtml_col_typeDB_Data := *xhtml_col_typeDB
	preservedPtrToXhtml_col_type := &xhtml_col_typeDB_Data
	backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB[xhtml_col_typeDB.ID] = preservedPtrToXhtml_col_type

	return
}

// BackRepoXhtml_col_type.CheckoutPhaseTwo Checkouts all staged instances of Xhtml_col_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, xhtml_col_typeDB := range backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB {
		backRepoXhtml_col_type.CheckoutPhaseTwoInstance(backRepo, xhtml_col_typeDB)
	}
	return
}

// BackRepoXhtml_col_type.CheckoutPhaseTwoInstance Checkouts staged instances of Xhtml_col_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, xhtml_col_typeDB *Xhtml_col_typeDB) (Error error) {

	xhtml_col_type := backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr[xhtml_col_typeDB.ID]

	xhtml_col_typeDB.DecodePointers(backRepo, xhtml_col_type)

	return
}

func (xhtml_col_typeDB *Xhtml_col_typeDB) DecodePointers(backRepo *BackRepoStruct, xhtml_col_type *models.Xhtml_col_type) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitXhtml_col_type allows commit of a single xhtml_col_type (if already staged)
func (backRepo *BackRepoStruct) CommitXhtml_col_type(xhtml_col_type *models.Xhtml_col_type) {
	backRepo.BackRepoXhtml_col_type.CommitPhaseOneInstance(xhtml_col_type)
	if id, ok := backRepo.BackRepoXhtml_col_type.Map_Xhtml_col_typePtr_Xhtml_col_typeDBID[xhtml_col_type]; ok {
		backRepo.BackRepoXhtml_col_type.CommitPhaseTwoInstance(backRepo, id, xhtml_col_type)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitXhtml_col_type allows checkout of a single xhtml_col_type (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutXhtml_col_type(xhtml_col_type *models.Xhtml_col_type) {
	// check if the xhtml_col_type is staged
	if _, ok := backRepo.BackRepoXhtml_col_type.Map_Xhtml_col_typePtr_Xhtml_col_typeDBID[xhtml_col_type]; ok {

		if id, ok := backRepo.BackRepoXhtml_col_type.Map_Xhtml_col_typePtr_Xhtml_col_typeDBID[xhtml_col_type]; ok {
			var xhtml_col_typeDB Xhtml_col_typeDB
			xhtml_col_typeDB.ID = id

			if _, err := backRepo.BackRepoXhtml_col_type.db.First(&xhtml_col_typeDB, id); err != nil {
				log.Fatalln("CheckoutXhtml_col_type : Problem with getting object with id:", id)
			}
			backRepo.BackRepoXhtml_col_type.CheckoutPhaseOneInstance(&xhtml_col_typeDB)
			backRepo.BackRepoXhtml_col_type.CheckoutPhaseTwoInstance(backRepo, &xhtml_col_typeDB)
		}
	}
}

// CopyBasicFieldsFromXhtml_col_type
func (xhtml_col_typeDB *Xhtml_col_typeDB) CopyBasicFieldsFromXhtml_col_type(xhtml_col_type *models.Xhtml_col_type) {
	// insertion point for fields commit

	xhtml_col_typeDB.Name_Data.String = xhtml_col_type.Name
	xhtml_col_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_col_type_WOP
func (xhtml_col_typeDB *Xhtml_col_typeDB) CopyBasicFieldsFromXhtml_col_type_WOP(xhtml_col_type *models.Xhtml_col_type_WOP) {
	// insertion point for fields commit

	xhtml_col_typeDB.Name_Data.String = xhtml_col_type.Name
	xhtml_col_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_col_typeWOP
func (xhtml_col_typeDB *Xhtml_col_typeDB) CopyBasicFieldsFromXhtml_col_typeWOP(xhtml_col_type *Xhtml_col_typeWOP) {
	// insertion point for fields commit

	xhtml_col_typeDB.Name_Data.String = xhtml_col_type.Name
	xhtml_col_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToXhtml_col_type
func (xhtml_col_typeDB *Xhtml_col_typeDB) CopyBasicFieldsToXhtml_col_type(xhtml_col_type *models.Xhtml_col_type) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_col_type.Name = xhtml_col_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_col_type_WOP
func (xhtml_col_typeDB *Xhtml_col_typeDB) CopyBasicFieldsToXhtml_col_type_WOP(xhtml_col_type *models.Xhtml_col_type_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_col_type.Name = xhtml_col_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_col_typeWOP
func (xhtml_col_typeDB *Xhtml_col_typeDB) CopyBasicFieldsToXhtml_col_typeWOP(xhtml_col_type *Xhtml_col_typeWOP) {
	xhtml_col_type.ID = int(xhtml_col_typeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_col_type.Name = xhtml_col_typeDB.Name_Data.String
}

// Backup generates a json file from a slice of all Xhtml_col_typeDB instances in the backrepo
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Xhtml_col_typeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_col_typeDB, 0)
	for _, xhtml_col_typeDB := range backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB {
		forBackup = append(forBackup, xhtml_col_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Xhtml_col_type ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Xhtml_col_type file", err.Error())
	}
}

// Backup generates a json file from a slice of all Xhtml_col_typeDB instances in the backrepo
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_col_typeDB, 0)
	for _, xhtml_col_typeDB := range backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB {
		forBackup = append(forBackup, xhtml_col_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Xhtml_col_type")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Xhtml_col_type_Fields, -1)
	for _, xhtml_col_typeDB := range forBackup {

		var xhtml_col_typeWOP Xhtml_col_typeWOP
		xhtml_col_typeDB.CopyBasicFieldsToXhtml_col_typeWOP(&xhtml_col_typeWOP)

		row := sh.AddRow()
		row.WriteStruct(&xhtml_col_typeWOP, -1)
	}
}

// RestoreXL from the "Xhtml_col_type" sheet all Xhtml_col_typeDB instances
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoXhtml_col_typeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Xhtml_col_type"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoXhtml_col_type.rowVisitorXhtml_col_type)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) rowVisitorXhtml_col_type(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var xhtml_col_typeWOP Xhtml_col_typeWOP
		row.ReadStruct(&xhtml_col_typeWOP)

		// add the unmarshalled struct to the stage
		xhtml_col_typeDB := new(Xhtml_col_typeDB)
		xhtml_col_typeDB.CopyBasicFieldsFromXhtml_col_typeWOP(&xhtml_col_typeWOP)

		xhtml_col_typeDB_ID_atBackupTime := xhtml_col_typeDB.ID
		xhtml_col_typeDB.ID = 0
		_, err := backRepoXhtml_col_type.db.Create(xhtml_col_typeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB[xhtml_col_typeDB.ID] = xhtml_col_typeDB
		BackRepoXhtml_col_typeid_atBckpTime_newID[xhtml_col_typeDB_ID_atBackupTime] = xhtml_col_typeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Xhtml_col_typeDB.json" in dirPath that stores an array
// of Xhtml_col_typeDB and stores it in the database
// the map BackRepoXhtml_col_typeid_atBckpTime_newID is updated accordingly
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoXhtml_col_typeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Xhtml_col_typeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Xhtml_col_type file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Xhtml_col_typeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Xhtml_col_typeDBID_Xhtml_col_typeDB
	for _, xhtml_col_typeDB := range forRestore {

		xhtml_col_typeDB_ID_atBackupTime := xhtml_col_typeDB.ID
		xhtml_col_typeDB.ID = 0
		_, err := backRepoXhtml_col_type.db.Create(xhtml_col_typeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB[xhtml_col_typeDB.ID] = xhtml_col_typeDB
		BackRepoXhtml_col_typeid_atBckpTime_newID[xhtml_col_typeDB_ID_atBackupTime] = xhtml_col_typeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Xhtml_col_type file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Xhtml_col_type>id_atBckpTime_newID
// to compute new index
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) RestorePhaseTwo() {

	for _, xhtml_col_typeDB := range backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB {

		// next line of code is to avert unused variable compilation error
		_ = xhtml_col_typeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoXhtml_col_type.db.Model(xhtml_col_typeDB)
		_, err := db.Updates(*xhtml_col_typeDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoXhtml_col_type.ResetReversePointers commits all staged instances of Xhtml_col_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_col_type := range backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr {
		backRepoXhtml_col_type.ResetReversePointersInstance(backRepo, idx, xhtml_col_type)
	}

	return
}

func (backRepoXhtml_col_type *BackRepoXhtml_col_typeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, xhtml_col_type *models.Xhtml_col_type) (Error error) {

	// fetch matching xhtml_col_typeDB
	if xhtml_col_typeDB, ok := backRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typeDB[idx]; ok {
		_ = xhtml_col_typeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoXhtml_col_typeid_atBckpTime_newID map[uint]uint
