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
var dummy_Xhtml_dl_type_sql sql.NullBool
var dummy_Xhtml_dl_type_time time.Duration
var dummy_Xhtml_dl_type_sort sort.Float64Slice

// Xhtml_dl_typeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model xhtml_dl_typeAPI
type Xhtml_dl_typeAPI struct {
	gorm.Model

	models.Xhtml_dl_type_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Xhtml_dl_typePointersEncoding Xhtml_dl_typePointersEncoding
}

// Xhtml_dl_typePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Xhtml_dl_typePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Xhtml_dl_typeDB describes a xhtml_dl_type in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model xhtml_dl_typeDB
type Xhtml_dl_typeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field xhtml_dl_typeDB.Name
	Name_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Xhtml_dl_typePointersEncoding
}

// Xhtml_dl_typeDBs arrays xhtml_dl_typeDBs
// swagger:response xhtml_dl_typeDBsResponse
type Xhtml_dl_typeDBs []Xhtml_dl_typeDB

// Xhtml_dl_typeDBResponse provides response
// swagger:response xhtml_dl_typeDBResponse
type Xhtml_dl_typeDBResponse struct {
	Xhtml_dl_typeDB
}

// Xhtml_dl_typeWOP is a Xhtml_dl_type without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Xhtml_dl_typeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Xhtml_dl_type_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoXhtml_dl_typeStruct struct {
	// stores Xhtml_dl_typeDB according to their gorm ID
	Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB map[uint]*Xhtml_dl_typeDB

	// stores Xhtml_dl_typeDB ID according to Xhtml_dl_type address
	Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID map[*models.Xhtml_dl_type]uint

	// stores Xhtml_dl_type according to their gorm ID
	Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr map[uint]*models.Xhtml_dl_type

	db db.DBInterface

	stage *models.Stage
}

func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) GetStage() (stage *models.Stage) {
	stage = backRepoXhtml_dl_type.stage
	return
}

func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) GetDB() db.DBInterface {
	return backRepoXhtml_dl_type.db
}

// GetXhtml_dl_typeDBFromXhtml_dl_typePtr is a handy function to access the back repo instance from the stage instance
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) GetXhtml_dl_typeDBFromXhtml_dl_typePtr(xhtml_dl_type *models.Xhtml_dl_type) (xhtml_dl_typeDB *Xhtml_dl_typeDB) {
	id := backRepoXhtml_dl_type.Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID[xhtml_dl_type]
	xhtml_dl_typeDB = backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB[id]
	return
}

// BackRepoXhtml_dl_type.CommitPhaseOne commits all staged instances of Xhtml_dl_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var xhtml_dl_types []*models.Xhtml_dl_type
	for xhtml_dl_type := range stage.Xhtml_dl_types {
		xhtml_dl_types = append(xhtml_dl_types, xhtml_dl_type)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(xhtml_dl_types, func(i, j int) bool {
		return stage.Xhtml_dl_typeMap_Staged_Order[xhtml_dl_types[i]] < stage.Xhtml_dl_typeMap_Staged_Order[xhtml_dl_types[j]]
	})

	for _, xhtml_dl_type := range xhtml_dl_types {
		backRepoXhtml_dl_type.CommitPhaseOneInstance(xhtml_dl_type)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, xhtml_dl_type := range backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr {
		if _, ok := stage.Xhtml_dl_types[xhtml_dl_type]; !ok {
			backRepoXhtml_dl_type.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoXhtml_dl_type.CommitDeleteInstance commits deletion of Xhtml_dl_type to the BackRepo
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) CommitDeleteInstance(id uint) (Error error) {

	xhtml_dl_type := backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr[id]

	// xhtml_dl_type is not staged anymore, remove xhtml_dl_typeDB
	xhtml_dl_typeDB := backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB[id]
	db, _ := backRepoXhtml_dl_type.db.Unscoped()
	_, err := db.Delete(xhtml_dl_typeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoXhtml_dl_type.Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID, xhtml_dl_type)
	delete(backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr, id)
	delete(backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB, id)

	return
}

// BackRepoXhtml_dl_type.CommitPhaseOneInstance commits xhtml_dl_type staged instances of Xhtml_dl_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) CommitPhaseOneInstance(xhtml_dl_type *models.Xhtml_dl_type) (Error error) {

	// check if the xhtml_dl_type is not commited yet
	if _, ok := backRepoXhtml_dl_type.Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID[xhtml_dl_type]; ok {
		return
	}

	// initiate xhtml_dl_type
	var xhtml_dl_typeDB Xhtml_dl_typeDB
	xhtml_dl_typeDB.CopyBasicFieldsFromXhtml_dl_type(xhtml_dl_type)

	_, err := backRepoXhtml_dl_type.db.Create(&xhtml_dl_typeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoXhtml_dl_type.Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID[xhtml_dl_type] = xhtml_dl_typeDB.ID
	backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr[xhtml_dl_typeDB.ID] = xhtml_dl_type
	backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB[xhtml_dl_typeDB.ID] = &xhtml_dl_typeDB

	return
}

// BackRepoXhtml_dl_type.CommitPhaseTwo commits all staged instances of Xhtml_dl_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_dl_type := range backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr {
		backRepoXhtml_dl_type.CommitPhaseTwoInstance(backRepo, idx, xhtml_dl_type)
	}

	return
}

// BackRepoXhtml_dl_type.CommitPhaseTwoInstance commits {{structname }} of models.Xhtml_dl_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, xhtml_dl_type *models.Xhtml_dl_type) (Error error) {

	// fetch matching xhtml_dl_typeDB
	if xhtml_dl_typeDB, ok := backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB[idx]; ok {

		xhtml_dl_typeDB.CopyBasicFieldsFromXhtml_dl_type(xhtml_dl_type)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoXhtml_dl_type.db.Save(xhtml_dl_typeDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Xhtml_dl_type intance %s", xhtml_dl_type.Name))
		return err
	}

	return
}

// BackRepoXhtml_dl_type.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) CheckoutPhaseOne() (Error error) {

	xhtml_dl_typeDBArray := make([]Xhtml_dl_typeDB, 0)
	_, err := backRepoXhtml_dl_type.db.Find(&xhtml_dl_typeDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	xhtml_dl_typeInstancesToBeRemovedFromTheStage := make(map[*models.Xhtml_dl_type]any)
	for key, value := range backRepoXhtml_dl_type.stage.Xhtml_dl_types {
		xhtml_dl_typeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, xhtml_dl_typeDB := range xhtml_dl_typeDBArray {
		backRepoXhtml_dl_type.CheckoutPhaseOneInstance(&xhtml_dl_typeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		xhtml_dl_type, ok := backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr[xhtml_dl_typeDB.ID]
		if ok {
			delete(xhtml_dl_typeInstancesToBeRemovedFromTheStage, xhtml_dl_type)
		}
	}

	// remove from stage and back repo's 3 maps all xhtml_dl_types that are not in the checkout
	for xhtml_dl_type := range xhtml_dl_typeInstancesToBeRemovedFromTheStage {
		xhtml_dl_type.Unstage(backRepoXhtml_dl_type.GetStage())

		// remove instance from the back repo 3 maps
		xhtml_dl_typeID := backRepoXhtml_dl_type.Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID[xhtml_dl_type]
		delete(backRepoXhtml_dl_type.Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID, xhtml_dl_type)
		delete(backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB, xhtml_dl_typeID)
		delete(backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr, xhtml_dl_typeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a xhtml_dl_typeDB that has been found in the DB, updates the backRepo and stages the
// models version of the xhtml_dl_typeDB
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) CheckoutPhaseOneInstance(xhtml_dl_typeDB *Xhtml_dl_typeDB) (Error error) {

	xhtml_dl_type, ok := backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr[xhtml_dl_typeDB.ID]
	if !ok {
		xhtml_dl_type = new(models.Xhtml_dl_type)

		backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr[xhtml_dl_typeDB.ID] = xhtml_dl_type
		backRepoXhtml_dl_type.Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID[xhtml_dl_type] = xhtml_dl_typeDB.ID

		// append model store with the new element
		xhtml_dl_type.Name = xhtml_dl_typeDB.Name_Data.String
		xhtml_dl_type.Stage(backRepoXhtml_dl_type.GetStage())
	}
	xhtml_dl_typeDB.CopyBasicFieldsToXhtml_dl_type(xhtml_dl_type)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	xhtml_dl_type.Stage(backRepoXhtml_dl_type.GetStage())

	// preserve pointer to xhtml_dl_typeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB)[xhtml_dl_typeDB hold variable pointers
	xhtml_dl_typeDB_Data := *xhtml_dl_typeDB
	preservedPtrToXhtml_dl_type := &xhtml_dl_typeDB_Data
	backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB[xhtml_dl_typeDB.ID] = preservedPtrToXhtml_dl_type

	return
}

// BackRepoXhtml_dl_type.CheckoutPhaseTwo Checkouts all staged instances of Xhtml_dl_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, xhtml_dl_typeDB := range backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB {
		backRepoXhtml_dl_type.CheckoutPhaseTwoInstance(backRepo, xhtml_dl_typeDB)
	}
	return
}

// BackRepoXhtml_dl_type.CheckoutPhaseTwoInstance Checkouts staged instances of Xhtml_dl_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, xhtml_dl_typeDB *Xhtml_dl_typeDB) (Error error) {

	xhtml_dl_type := backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr[xhtml_dl_typeDB.ID]

	xhtml_dl_typeDB.DecodePointers(backRepo, xhtml_dl_type)

	return
}

func (xhtml_dl_typeDB *Xhtml_dl_typeDB) DecodePointers(backRepo *BackRepoStruct, xhtml_dl_type *models.Xhtml_dl_type) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitXhtml_dl_type allows commit of a single xhtml_dl_type (if already staged)
func (backRepo *BackRepoStruct) CommitXhtml_dl_type(xhtml_dl_type *models.Xhtml_dl_type) {
	backRepo.BackRepoXhtml_dl_type.CommitPhaseOneInstance(xhtml_dl_type)
	if id, ok := backRepo.BackRepoXhtml_dl_type.Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID[xhtml_dl_type]; ok {
		backRepo.BackRepoXhtml_dl_type.CommitPhaseTwoInstance(backRepo, id, xhtml_dl_type)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitXhtml_dl_type allows checkout of a single xhtml_dl_type (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutXhtml_dl_type(xhtml_dl_type *models.Xhtml_dl_type) {
	// check if the xhtml_dl_type is staged
	if _, ok := backRepo.BackRepoXhtml_dl_type.Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID[xhtml_dl_type]; ok {

		if id, ok := backRepo.BackRepoXhtml_dl_type.Map_Xhtml_dl_typePtr_Xhtml_dl_typeDBID[xhtml_dl_type]; ok {
			var xhtml_dl_typeDB Xhtml_dl_typeDB
			xhtml_dl_typeDB.ID = id

			if _, err := backRepo.BackRepoXhtml_dl_type.db.First(&xhtml_dl_typeDB, id); err != nil {
				log.Fatalln("CheckoutXhtml_dl_type : Problem with getting object with id:", id)
			}
			backRepo.BackRepoXhtml_dl_type.CheckoutPhaseOneInstance(&xhtml_dl_typeDB)
			backRepo.BackRepoXhtml_dl_type.CheckoutPhaseTwoInstance(backRepo, &xhtml_dl_typeDB)
		}
	}
}

// CopyBasicFieldsFromXhtml_dl_type
func (xhtml_dl_typeDB *Xhtml_dl_typeDB) CopyBasicFieldsFromXhtml_dl_type(xhtml_dl_type *models.Xhtml_dl_type) {
	// insertion point for fields commit

	xhtml_dl_typeDB.Name_Data.String = xhtml_dl_type.Name
	xhtml_dl_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_dl_type_WOP
func (xhtml_dl_typeDB *Xhtml_dl_typeDB) CopyBasicFieldsFromXhtml_dl_type_WOP(xhtml_dl_type *models.Xhtml_dl_type_WOP) {
	// insertion point for fields commit

	xhtml_dl_typeDB.Name_Data.String = xhtml_dl_type.Name
	xhtml_dl_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_dl_typeWOP
func (xhtml_dl_typeDB *Xhtml_dl_typeDB) CopyBasicFieldsFromXhtml_dl_typeWOP(xhtml_dl_type *Xhtml_dl_typeWOP) {
	// insertion point for fields commit

	xhtml_dl_typeDB.Name_Data.String = xhtml_dl_type.Name
	xhtml_dl_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToXhtml_dl_type
func (xhtml_dl_typeDB *Xhtml_dl_typeDB) CopyBasicFieldsToXhtml_dl_type(xhtml_dl_type *models.Xhtml_dl_type) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_dl_type.Name = xhtml_dl_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_dl_type_WOP
func (xhtml_dl_typeDB *Xhtml_dl_typeDB) CopyBasicFieldsToXhtml_dl_type_WOP(xhtml_dl_type *models.Xhtml_dl_type_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_dl_type.Name = xhtml_dl_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_dl_typeWOP
func (xhtml_dl_typeDB *Xhtml_dl_typeDB) CopyBasicFieldsToXhtml_dl_typeWOP(xhtml_dl_type *Xhtml_dl_typeWOP) {
	xhtml_dl_type.ID = int(xhtml_dl_typeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_dl_type.Name = xhtml_dl_typeDB.Name_Data.String
}

// Backup generates a json file from a slice of all Xhtml_dl_typeDB instances in the backrepo
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Xhtml_dl_typeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_dl_typeDB, 0)
	for _, xhtml_dl_typeDB := range backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB {
		forBackup = append(forBackup, xhtml_dl_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Xhtml_dl_type ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Xhtml_dl_type file", err.Error())
	}
}

// Backup generates a json file from a slice of all Xhtml_dl_typeDB instances in the backrepo
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_dl_typeDB, 0)
	for _, xhtml_dl_typeDB := range backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB {
		forBackup = append(forBackup, xhtml_dl_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Xhtml_dl_type")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Xhtml_dl_type_Fields, -1)
	for _, xhtml_dl_typeDB := range forBackup {

		var xhtml_dl_typeWOP Xhtml_dl_typeWOP
		xhtml_dl_typeDB.CopyBasicFieldsToXhtml_dl_typeWOP(&xhtml_dl_typeWOP)

		row := sh.AddRow()
		row.WriteStruct(&xhtml_dl_typeWOP, -1)
	}
}

// RestoreXL from the "Xhtml_dl_type" sheet all Xhtml_dl_typeDB instances
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoXhtml_dl_typeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Xhtml_dl_type"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoXhtml_dl_type.rowVisitorXhtml_dl_type)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) rowVisitorXhtml_dl_type(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var xhtml_dl_typeWOP Xhtml_dl_typeWOP
		row.ReadStruct(&xhtml_dl_typeWOP)

		// add the unmarshalled struct to the stage
		xhtml_dl_typeDB := new(Xhtml_dl_typeDB)
		xhtml_dl_typeDB.CopyBasicFieldsFromXhtml_dl_typeWOP(&xhtml_dl_typeWOP)

		xhtml_dl_typeDB_ID_atBackupTime := xhtml_dl_typeDB.ID
		xhtml_dl_typeDB.ID = 0
		_, err := backRepoXhtml_dl_type.db.Create(xhtml_dl_typeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB[xhtml_dl_typeDB.ID] = xhtml_dl_typeDB
		BackRepoXhtml_dl_typeid_atBckpTime_newID[xhtml_dl_typeDB_ID_atBackupTime] = xhtml_dl_typeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Xhtml_dl_typeDB.json" in dirPath that stores an array
// of Xhtml_dl_typeDB and stores it in the database
// the map BackRepoXhtml_dl_typeid_atBckpTime_newID is updated accordingly
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoXhtml_dl_typeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Xhtml_dl_typeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Xhtml_dl_type file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Xhtml_dl_typeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB
	for _, xhtml_dl_typeDB := range forRestore {

		xhtml_dl_typeDB_ID_atBackupTime := xhtml_dl_typeDB.ID
		xhtml_dl_typeDB.ID = 0
		_, err := backRepoXhtml_dl_type.db.Create(xhtml_dl_typeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB[xhtml_dl_typeDB.ID] = xhtml_dl_typeDB
		BackRepoXhtml_dl_typeid_atBckpTime_newID[xhtml_dl_typeDB_ID_atBackupTime] = xhtml_dl_typeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Xhtml_dl_type file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Xhtml_dl_type>id_atBckpTime_newID
// to compute new index
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) RestorePhaseTwo() {

	for _, xhtml_dl_typeDB := range backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB {

		// next line of code is to avert unused variable compilation error
		_ = xhtml_dl_typeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoXhtml_dl_type.db.Model(xhtml_dl_typeDB)
		_, err := db.Updates(*xhtml_dl_typeDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoXhtml_dl_type.ResetReversePointers commits all staged instances of Xhtml_dl_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_dl_type := range backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr {
		backRepoXhtml_dl_type.ResetReversePointersInstance(backRepo, idx, xhtml_dl_type)
	}

	return
}

func (backRepoXhtml_dl_type *BackRepoXhtml_dl_typeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, xhtml_dl_type *models.Xhtml_dl_type) (Error error) {

	// fetch matching xhtml_dl_typeDB
	if xhtml_dl_typeDB, ok := backRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typeDB[idx]; ok {
		_ = xhtml_dl_typeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoXhtml_dl_typeid_atBckpTime_newID map[uint]uint
