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
var dummy_Xhtml_br_type_sql sql.NullBool
var dummy_Xhtml_br_type_time time.Duration
var dummy_Xhtml_br_type_sort sort.Float64Slice

// Xhtml_br_typeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model xhtml_br_typeAPI
type Xhtml_br_typeAPI struct {
	gorm.Model

	models.Xhtml_br_type_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Xhtml_br_typePointersEncoding Xhtml_br_typePointersEncoding
}

// Xhtml_br_typePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Xhtml_br_typePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Xhtml_br_typeDB describes a xhtml_br_type in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model xhtml_br_typeDB
type Xhtml_br_typeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field xhtml_br_typeDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Xhtml_br_typePointersEncoding
}

// Xhtml_br_typeDBs arrays xhtml_br_typeDBs
// swagger:response xhtml_br_typeDBsResponse
type Xhtml_br_typeDBs []Xhtml_br_typeDB

// Xhtml_br_typeDBResponse provides response
// swagger:response xhtml_br_typeDBResponse
type Xhtml_br_typeDBResponse struct {
	Xhtml_br_typeDB
}

// Xhtml_br_typeWOP is a Xhtml_br_type without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Xhtml_br_typeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Xhtml_br_type_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoXhtml_br_typeStruct struct {
	// stores Xhtml_br_typeDB according to their gorm ID
	Map_Xhtml_br_typeDBID_Xhtml_br_typeDB map[uint]*Xhtml_br_typeDB

	// stores Xhtml_br_typeDB ID according to Xhtml_br_type address
	Map_Xhtml_br_typePtr_Xhtml_br_typeDBID map[*models.Xhtml_br_type]uint

	// stores Xhtml_br_type according to their gorm ID
	Map_Xhtml_br_typeDBID_Xhtml_br_typePtr map[uint]*models.Xhtml_br_type

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoXhtml_br_type.stage
	return
}

func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) GetDB() *gorm.DB {
	return backRepoXhtml_br_type.db
}

// GetXhtml_br_typeDBFromXhtml_br_typePtr is a handy function to access the back repo instance from the stage instance
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) GetXhtml_br_typeDBFromXhtml_br_typePtr(xhtml_br_type *models.Xhtml_br_type) (xhtml_br_typeDB *Xhtml_br_typeDB) {
	id := backRepoXhtml_br_type.Map_Xhtml_br_typePtr_Xhtml_br_typeDBID[xhtml_br_type]
	xhtml_br_typeDB = backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB[id]
	return
}

// BackRepoXhtml_br_type.CommitPhaseOne commits all staged instances of Xhtml_br_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for xhtml_br_type := range stage.Xhtml_br_types {
		backRepoXhtml_br_type.CommitPhaseOneInstance(xhtml_br_type)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, xhtml_br_type := range backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr {
		if _, ok := stage.Xhtml_br_types[xhtml_br_type]; !ok {
			backRepoXhtml_br_type.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoXhtml_br_type.CommitDeleteInstance commits deletion of Xhtml_br_type to the BackRepo
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) CommitDeleteInstance(id uint) (Error error) {

	xhtml_br_type := backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr[id]

	// xhtml_br_type is not staged anymore, remove xhtml_br_typeDB
	xhtml_br_typeDB := backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB[id]
	query := backRepoXhtml_br_type.db.Unscoped().Delete(&xhtml_br_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoXhtml_br_type.Map_Xhtml_br_typePtr_Xhtml_br_typeDBID, xhtml_br_type)
	delete(backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr, id)
	delete(backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB, id)

	return
}

// BackRepoXhtml_br_type.CommitPhaseOneInstance commits xhtml_br_type staged instances of Xhtml_br_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) CommitPhaseOneInstance(xhtml_br_type *models.Xhtml_br_type) (Error error) {

	// check if the xhtml_br_type is not commited yet
	if _, ok := backRepoXhtml_br_type.Map_Xhtml_br_typePtr_Xhtml_br_typeDBID[xhtml_br_type]; ok {
		return
	}

	// initiate xhtml_br_type
	var xhtml_br_typeDB Xhtml_br_typeDB
	xhtml_br_typeDB.CopyBasicFieldsFromXhtml_br_type(xhtml_br_type)

	query := backRepoXhtml_br_type.db.Create(&xhtml_br_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoXhtml_br_type.Map_Xhtml_br_typePtr_Xhtml_br_typeDBID[xhtml_br_type] = xhtml_br_typeDB.ID
	backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr[xhtml_br_typeDB.ID] = xhtml_br_type
	backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB[xhtml_br_typeDB.ID] = &xhtml_br_typeDB

	return
}

// BackRepoXhtml_br_type.CommitPhaseTwo commits all staged instances of Xhtml_br_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_br_type := range backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr {
		backRepoXhtml_br_type.CommitPhaseTwoInstance(backRepo, idx, xhtml_br_type)
	}

	return
}

// BackRepoXhtml_br_type.CommitPhaseTwoInstance commits {{structname }} of models.Xhtml_br_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, xhtml_br_type *models.Xhtml_br_type) (Error error) {

	// fetch matching xhtml_br_typeDB
	if xhtml_br_typeDB, ok := backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB[idx]; ok {

		xhtml_br_typeDB.CopyBasicFieldsFromXhtml_br_type(xhtml_br_type)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoXhtml_br_type.db.Save(&xhtml_br_typeDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Xhtml_br_type intance %s", xhtml_br_type.Name))
		return err
	}

	return
}

// BackRepoXhtml_br_type.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) CheckoutPhaseOne() (Error error) {

	xhtml_br_typeDBArray := make([]Xhtml_br_typeDB, 0)
	query := backRepoXhtml_br_type.db.Find(&xhtml_br_typeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	xhtml_br_typeInstancesToBeRemovedFromTheStage := make(map[*models.Xhtml_br_type]any)
	for key, value := range backRepoXhtml_br_type.stage.Xhtml_br_types {
		xhtml_br_typeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, xhtml_br_typeDB := range xhtml_br_typeDBArray {
		backRepoXhtml_br_type.CheckoutPhaseOneInstance(&xhtml_br_typeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		xhtml_br_type, ok := backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr[xhtml_br_typeDB.ID]
		if ok {
			delete(xhtml_br_typeInstancesToBeRemovedFromTheStage, xhtml_br_type)
		}
	}

	// remove from stage and back repo's 3 maps all xhtml_br_types that are not in the checkout
	for xhtml_br_type := range xhtml_br_typeInstancesToBeRemovedFromTheStage {
		xhtml_br_type.Unstage(backRepoXhtml_br_type.GetStage())

		// remove instance from the back repo 3 maps
		xhtml_br_typeID := backRepoXhtml_br_type.Map_Xhtml_br_typePtr_Xhtml_br_typeDBID[xhtml_br_type]
		delete(backRepoXhtml_br_type.Map_Xhtml_br_typePtr_Xhtml_br_typeDBID, xhtml_br_type)
		delete(backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB, xhtml_br_typeID)
		delete(backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr, xhtml_br_typeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a xhtml_br_typeDB that has been found in the DB, updates the backRepo and stages the
// models version of the xhtml_br_typeDB
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) CheckoutPhaseOneInstance(xhtml_br_typeDB *Xhtml_br_typeDB) (Error error) {

	xhtml_br_type, ok := backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr[xhtml_br_typeDB.ID]
	if !ok {
		xhtml_br_type = new(models.Xhtml_br_type)

		backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr[xhtml_br_typeDB.ID] = xhtml_br_type
		backRepoXhtml_br_type.Map_Xhtml_br_typePtr_Xhtml_br_typeDBID[xhtml_br_type] = xhtml_br_typeDB.ID

		// append model store with the new element
		xhtml_br_type.Name = xhtml_br_typeDB.Name_Data.String
		xhtml_br_type.Stage(backRepoXhtml_br_type.GetStage())
	}
	xhtml_br_typeDB.CopyBasicFieldsToXhtml_br_type(xhtml_br_type)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	xhtml_br_type.Stage(backRepoXhtml_br_type.GetStage())

	// preserve pointer to xhtml_br_typeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Xhtml_br_typeDBID_Xhtml_br_typeDB)[xhtml_br_typeDB hold variable pointers
	xhtml_br_typeDB_Data := *xhtml_br_typeDB
	preservedPtrToXhtml_br_type := &xhtml_br_typeDB_Data
	backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB[xhtml_br_typeDB.ID] = preservedPtrToXhtml_br_type

	return
}

// BackRepoXhtml_br_type.CheckoutPhaseTwo Checkouts all staged instances of Xhtml_br_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, xhtml_br_typeDB := range backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB {
		backRepoXhtml_br_type.CheckoutPhaseTwoInstance(backRepo, xhtml_br_typeDB)
	}
	return
}

// BackRepoXhtml_br_type.CheckoutPhaseTwoInstance Checkouts staged instances of Xhtml_br_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, xhtml_br_typeDB *Xhtml_br_typeDB) (Error error) {

	xhtml_br_type := backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr[xhtml_br_typeDB.ID]

	xhtml_br_typeDB.DecodePointers(backRepo, xhtml_br_type)

	return
}

func (xhtml_br_typeDB *Xhtml_br_typeDB) DecodePointers(backRepo *BackRepoStruct, xhtml_br_type *models.Xhtml_br_type) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitXhtml_br_type allows commit of a single xhtml_br_type (if already staged)
func (backRepo *BackRepoStruct) CommitXhtml_br_type(xhtml_br_type *models.Xhtml_br_type) {
	backRepo.BackRepoXhtml_br_type.CommitPhaseOneInstance(xhtml_br_type)
	if id, ok := backRepo.BackRepoXhtml_br_type.Map_Xhtml_br_typePtr_Xhtml_br_typeDBID[xhtml_br_type]; ok {
		backRepo.BackRepoXhtml_br_type.CommitPhaseTwoInstance(backRepo, id, xhtml_br_type)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitXhtml_br_type allows checkout of a single xhtml_br_type (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutXhtml_br_type(xhtml_br_type *models.Xhtml_br_type) {
	// check if the xhtml_br_type is staged
	if _, ok := backRepo.BackRepoXhtml_br_type.Map_Xhtml_br_typePtr_Xhtml_br_typeDBID[xhtml_br_type]; ok {

		if id, ok := backRepo.BackRepoXhtml_br_type.Map_Xhtml_br_typePtr_Xhtml_br_typeDBID[xhtml_br_type]; ok {
			var xhtml_br_typeDB Xhtml_br_typeDB
			xhtml_br_typeDB.ID = id

			if err := backRepo.BackRepoXhtml_br_type.db.First(&xhtml_br_typeDB, id).Error; err != nil {
				log.Fatalln("CheckoutXhtml_br_type : Problem with getting object with id:", id)
			}
			backRepo.BackRepoXhtml_br_type.CheckoutPhaseOneInstance(&xhtml_br_typeDB)
			backRepo.BackRepoXhtml_br_type.CheckoutPhaseTwoInstance(backRepo, &xhtml_br_typeDB)
		}
	}
}

// CopyBasicFieldsFromXhtml_br_type
func (xhtml_br_typeDB *Xhtml_br_typeDB) CopyBasicFieldsFromXhtml_br_type(xhtml_br_type *models.Xhtml_br_type) {
	// insertion point for fields commit

	xhtml_br_typeDB.Name_Data.String = xhtml_br_type.Name
	xhtml_br_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_br_type_WOP
func (xhtml_br_typeDB *Xhtml_br_typeDB) CopyBasicFieldsFromXhtml_br_type_WOP(xhtml_br_type *models.Xhtml_br_type_WOP) {
	// insertion point for fields commit

	xhtml_br_typeDB.Name_Data.String = xhtml_br_type.Name
	xhtml_br_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_br_typeWOP
func (xhtml_br_typeDB *Xhtml_br_typeDB) CopyBasicFieldsFromXhtml_br_typeWOP(xhtml_br_type *Xhtml_br_typeWOP) {
	// insertion point for fields commit

	xhtml_br_typeDB.Name_Data.String = xhtml_br_type.Name
	xhtml_br_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToXhtml_br_type
func (xhtml_br_typeDB *Xhtml_br_typeDB) CopyBasicFieldsToXhtml_br_type(xhtml_br_type *models.Xhtml_br_type) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_br_type.Name = xhtml_br_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_br_type_WOP
func (xhtml_br_typeDB *Xhtml_br_typeDB) CopyBasicFieldsToXhtml_br_type_WOP(xhtml_br_type *models.Xhtml_br_type_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_br_type.Name = xhtml_br_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_br_typeWOP
func (xhtml_br_typeDB *Xhtml_br_typeDB) CopyBasicFieldsToXhtml_br_typeWOP(xhtml_br_type *Xhtml_br_typeWOP) {
	xhtml_br_type.ID = int(xhtml_br_typeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_br_type.Name = xhtml_br_typeDB.Name_Data.String
}

// Backup generates a json file from a slice of all Xhtml_br_typeDB instances in the backrepo
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Xhtml_br_typeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_br_typeDB, 0)
	for _, xhtml_br_typeDB := range backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB {
		forBackup = append(forBackup, xhtml_br_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Xhtml_br_type ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Xhtml_br_type file", err.Error())
	}
}

// Backup generates a json file from a slice of all Xhtml_br_typeDB instances in the backrepo
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_br_typeDB, 0)
	for _, xhtml_br_typeDB := range backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB {
		forBackup = append(forBackup, xhtml_br_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Xhtml_br_type")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Xhtml_br_type_Fields, -1)
	for _, xhtml_br_typeDB := range forBackup {

		var xhtml_br_typeWOP Xhtml_br_typeWOP
		xhtml_br_typeDB.CopyBasicFieldsToXhtml_br_typeWOP(&xhtml_br_typeWOP)

		row := sh.AddRow()
		row.WriteStruct(&xhtml_br_typeWOP, -1)
	}
}

// RestoreXL from the "Xhtml_br_type" sheet all Xhtml_br_typeDB instances
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoXhtml_br_typeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Xhtml_br_type"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoXhtml_br_type.rowVisitorXhtml_br_type)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) rowVisitorXhtml_br_type(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var xhtml_br_typeWOP Xhtml_br_typeWOP
		row.ReadStruct(&xhtml_br_typeWOP)

		// add the unmarshalled struct to the stage
		xhtml_br_typeDB := new(Xhtml_br_typeDB)
		xhtml_br_typeDB.CopyBasicFieldsFromXhtml_br_typeWOP(&xhtml_br_typeWOP)

		xhtml_br_typeDB_ID_atBackupTime := xhtml_br_typeDB.ID
		xhtml_br_typeDB.ID = 0
		query := backRepoXhtml_br_type.db.Create(xhtml_br_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB[xhtml_br_typeDB.ID] = xhtml_br_typeDB
		BackRepoXhtml_br_typeid_atBckpTime_newID[xhtml_br_typeDB_ID_atBackupTime] = xhtml_br_typeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Xhtml_br_typeDB.json" in dirPath that stores an array
// of Xhtml_br_typeDB and stores it in the database
// the map BackRepoXhtml_br_typeid_atBckpTime_newID is updated accordingly
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoXhtml_br_typeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Xhtml_br_typeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Xhtml_br_type file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Xhtml_br_typeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Xhtml_br_typeDBID_Xhtml_br_typeDB
	for _, xhtml_br_typeDB := range forRestore {

		xhtml_br_typeDB_ID_atBackupTime := xhtml_br_typeDB.ID
		xhtml_br_typeDB.ID = 0
		query := backRepoXhtml_br_type.db.Create(xhtml_br_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB[xhtml_br_typeDB.ID] = xhtml_br_typeDB
		BackRepoXhtml_br_typeid_atBckpTime_newID[xhtml_br_typeDB_ID_atBackupTime] = xhtml_br_typeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Xhtml_br_type file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Xhtml_br_type>id_atBckpTime_newID
// to compute new index
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) RestorePhaseTwo() {

	for _, xhtml_br_typeDB := range backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB {

		// next line of code is to avert unused variable compilation error
		_ = xhtml_br_typeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoXhtml_br_type.db.Model(xhtml_br_typeDB).Updates(*xhtml_br_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoXhtml_br_type.ResetReversePointers commits all staged instances of Xhtml_br_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_br_type := range backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr {
		backRepoXhtml_br_type.ResetReversePointersInstance(backRepo, idx, xhtml_br_type)
	}

	return
}

func (backRepoXhtml_br_type *BackRepoXhtml_br_typeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, xhtml_br_type *models.Xhtml_br_type) (Error error) {

	// fetch matching xhtml_br_typeDB
	if xhtml_br_typeDB, ok := backRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typeDB[idx]; ok {
		_ = xhtml_br_typeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoXhtml_br_typeid_atBckpTime_newID map[uint]uint