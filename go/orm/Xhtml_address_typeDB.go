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
var dummy_Xhtml_address_type_sql sql.NullBool
var dummy_Xhtml_address_type_time time.Duration
var dummy_Xhtml_address_type_sort sort.Float64Slice

// Xhtml_address_typeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model xhtml_address_typeAPI
type Xhtml_address_typeAPI struct {
	gorm.Model

	models.Xhtml_address_type_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Xhtml_address_typePointersEncoding Xhtml_address_typePointersEncoding
}

// Xhtml_address_typePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Xhtml_address_typePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Xhtml_address_typeDB describes a xhtml_address_type in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model xhtml_address_typeDB
type Xhtml_address_typeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field xhtml_address_typeDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Xhtml_address_typePointersEncoding
}

// Xhtml_address_typeDBs arrays xhtml_address_typeDBs
// swagger:response xhtml_address_typeDBsResponse
type Xhtml_address_typeDBs []Xhtml_address_typeDB

// Xhtml_address_typeDBResponse provides response
// swagger:response xhtml_address_typeDBResponse
type Xhtml_address_typeDBResponse struct {
	Xhtml_address_typeDB
}

// Xhtml_address_typeWOP is a Xhtml_address_type without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Xhtml_address_typeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Xhtml_address_type_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoXhtml_address_typeStruct struct {
	// stores Xhtml_address_typeDB according to their gorm ID
	Map_Xhtml_address_typeDBID_Xhtml_address_typeDB map[uint]*Xhtml_address_typeDB

	// stores Xhtml_address_typeDB ID according to Xhtml_address_type address
	Map_Xhtml_address_typePtr_Xhtml_address_typeDBID map[*models.Xhtml_address_type]uint

	// stores Xhtml_address_type according to their gorm ID
	Map_Xhtml_address_typeDBID_Xhtml_address_typePtr map[uint]*models.Xhtml_address_type

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoXhtml_address_type.stage
	return
}

func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) GetDB() *gorm.DB {
	return backRepoXhtml_address_type.db
}

// GetXhtml_address_typeDBFromXhtml_address_typePtr is a handy function to access the back repo instance from the stage instance
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) GetXhtml_address_typeDBFromXhtml_address_typePtr(xhtml_address_type *models.Xhtml_address_type) (xhtml_address_typeDB *Xhtml_address_typeDB) {
	id := backRepoXhtml_address_type.Map_Xhtml_address_typePtr_Xhtml_address_typeDBID[xhtml_address_type]
	xhtml_address_typeDB = backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB[id]
	return
}

// BackRepoXhtml_address_type.CommitPhaseOne commits all staged instances of Xhtml_address_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for xhtml_address_type := range stage.Xhtml_address_types {
		backRepoXhtml_address_type.CommitPhaseOneInstance(xhtml_address_type)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, xhtml_address_type := range backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr {
		if _, ok := stage.Xhtml_address_types[xhtml_address_type]; !ok {
			backRepoXhtml_address_type.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoXhtml_address_type.CommitDeleteInstance commits deletion of Xhtml_address_type to the BackRepo
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) CommitDeleteInstance(id uint) (Error error) {

	xhtml_address_type := backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr[id]

	// xhtml_address_type is not staged anymore, remove xhtml_address_typeDB
	xhtml_address_typeDB := backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB[id]
	query := backRepoXhtml_address_type.db.Unscoped().Delete(&xhtml_address_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoXhtml_address_type.Map_Xhtml_address_typePtr_Xhtml_address_typeDBID, xhtml_address_type)
	delete(backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr, id)
	delete(backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB, id)

	return
}

// BackRepoXhtml_address_type.CommitPhaseOneInstance commits xhtml_address_type staged instances of Xhtml_address_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) CommitPhaseOneInstance(xhtml_address_type *models.Xhtml_address_type) (Error error) {

	// check if the xhtml_address_type is not commited yet
	if _, ok := backRepoXhtml_address_type.Map_Xhtml_address_typePtr_Xhtml_address_typeDBID[xhtml_address_type]; ok {
		return
	}

	// initiate xhtml_address_type
	var xhtml_address_typeDB Xhtml_address_typeDB
	xhtml_address_typeDB.CopyBasicFieldsFromXhtml_address_type(xhtml_address_type)

	query := backRepoXhtml_address_type.db.Create(&xhtml_address_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoXhtml_address_type.Map_Xhtml_address_typePtr_Xhtml_address_typeDBID[xhtml_address_type] = xhtml_address_typeDB.ID
	backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr[xhtml_address_typeDB.ID] = xhtml_address_type
	backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB[xhtml_address_typeDB.ID] = &xhtml_address_typeDB

	return
}

// BackRepoXhtml_address_type.CommitPhaseTwo commits all staged instances of Xhtml_address_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_address_type := range backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr {
		backRepoXhtml_address_type.CommitPhaseTwoInstance(backRepo, idx, xhtml_address_type)
	}

	return
}

// BackRepoXhtml_address_type.CommitPhaseTwoInstance commits {{structname }} of models.Xhtml_address_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, xhtml_address_type *models.Xhtml_address_type) (Error error) {

	// fetch matching xhtml_address_typeDB
	if xhtml_address_typeDB, ok := backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB[idx]; ok {

		xhtml_address_typeDB.CopyBasicFieldsFromXhtml_address_type(xhtml_address_type)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoXhtml_address_type.db.Save(&xhtml_address_typeDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Xhtml_address_type intance %s", xhtml_address_type.Name))
		return err
	}

	return
}

// BackRepoXhtml_address_type.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) CheckoutPhaseOne() (Error error) {

	xhtml_address_typeDBArray := make([]Xhtml_address_typeDB, 0)
	query := backRepoXhtml_address_type.db.Find(&xhtml_address_typeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	xhtml_address_typeInstancesToBeRemovedFromTheStage := make(map[*models.Xhtml_address_type]any)
	for key, value := range backRepoXhtml_address_type.stage.Xhtml_address_types {
		xhtml_address_typeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, xhtml_address_typeDB := range xhtml_address_typeDBArray {
		backRepoXhtml_address_type.CheckoutPhaseOneInstance(&xhtml_address_typeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		xhtml_address_type, ok := backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr[xhtml_address_typeDB.ID]
		if ok {
			delete(xhtml_address_typeInstancesToBeRemovedFromTheStage, xhtml_address_type)
		}
	}

	// remove from stage and back repo's 3 maps all xhtml_address_types that are not in the checkout
	for xhtml_address_type := range xhtml_address_typeInstancesToBeRemovedFromTheStage {
		xhtml_address_type.Unstage(backRepoXhtml_address_type.GetStage())

		// remove instance from the back repo 3 maps
		xhtml_address_typeID := backRepoXhtml_address_type.Map_Xhtml_address_typePtr_Xhtml_address_typeDBID[xhtml_address_type]
		delete(backRepoXhtml_address_type.Map_Xhtml_address_typePtr_Xhtml_address_typeDBID, xhtml_address_type)
		delete(backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB, xhtml_address_typeID)
		delete(backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr, xhtml_address_typeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a xhtml_address_typeDB that has been found in the DB, updates the backRepo and stages the
// models version of the xhtml_address_typeDB
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) CheckoutPhaseOneInstance(xhtml_address_typeDB *Xhtml_address_typeDB) (Error error) {

	xhtml_address_type, ok := backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr[xhtml_address_typeDB.ID]
	if !ok {
		xhtml_address_type = new(models.Xhtml_address_type)

		backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr[xhtml_address_typeDB.ID] = xhtml_address_type
		backRepoXhtml_address_type.Map_Xhtml_address_typePtr_Xhtml_address_typeDBID[xhtml_address_type] = xhtml_address_typeDB.ID

		// append model store with the new element
		xhtml_address_type.Name = xhtml_address_typeDB.Name_Data.String
		xhtml_address_type.Stage(backRepoXhtml_address_type.GetStage())
	}
	xhtml_address_typeDB.CopyBasicFieldsToXhtml_address_type(xhtml_address_type)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	xhtml_address_type.Stage(backRepoXhtml_address_type.GetStage())

	// preserve pointer to xhtml_address_typeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Xhtml_address_typeDBID_Xhtml_address_typeDB)[xhtml_address_typeDB hold variable pointers
	xhtml_address_typeDB_Data := *xhtml_address_typeDB
	preservedPtrToXhtml_address_type := &xhtml_address_typeDB_Data
	backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB[xhtml_address_typeDB.ID] = preservedPtrToXhtml_address_type

	return
}

// BackRepoXhtml_address_type.CheckoutPhaseTwo Checkouts all staged instances of Xhtml_address_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, xhtml_address_typeDB := range backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB {
		backRepoXhtml_address_type.CheckoutPhaseTwoInstance(backRepo, xhtml_address_typeDB)
	}
	return
}

// BackRepoXhtml_address_type.CheckoutPhaseTwoInstance Checkouts staged instances of Xhtml_address_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, xhtml_address_typeDB *Xhtml_address_typeDB) (Error error) {

	xhtml_address_type := backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr[xhtml_address_typeDB.ID]

	xhtml_address_typeDB.DecodePointers(backRepo, xhtml_address_type)

	return
}

func (xhtml_address_typeDB *Xhtml_address_typeDB) DecodePointers(backRepo *BackRepoStruct, xhtml_address_type *models.Xhtml_address_type) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitXhtml_address_type allows commit of a single xhtml_address_type (if already staged)
func (backRepo *BackRepoStruct) CommitXhtml_address_type(xhtml_address_type *models.Xhtml_address_type) {
	backRepo.BackRepoXhtml_address_type.CommitPhaseOneInstance(xhtml_address_type)
	if id, ok := backRepo.BackRepoXhtml_address_type.Map_Xhtml_address_typePtr_Xhtml_address_typeDBID[xhtml_address_type]; ok {
		backRepo.BackRepoXhtml_address_type.CommitPhaseTwoInstance(backRepo, id, xhtml_address_type)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitXhtml_address_type allows checkout of a single xhtml_address_type (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutXhtml_address_type(xhtml_address_type *models.Xhtml_address_type) {
	// check if the xhtml_address_type is staged
	if _, ok := backRepo.BackRepoXhtml_address_type.Map_Xhtml_address_typePtr_Xhtml_address_typeDBID[xhtml_address_type]; ok {

		if id, ok := backRepo.BackRepoXhtml_address_type.Map_Xhtml_address_typePtr_Xhtml_address_typeDBID[xhtml_address_type]; ok {
			var xhtml_address_typeDB Xhtml_address_typeDB
			xhtml_address_typeDB.ID = id

			if err := backRepo.BackRepoXhtml_address_type.db.First(&xhtml_address_typeDB, id).Error; err != nil {
				log.Fatalln("CheckoutXhtml_address_type : Problem with getting object with id:", id)
			}
			backRepo.BackRepoXhtml_address_type.CheckoutPhaseOneInstance(&xhtml_address_typeDB)
			backRepo.BackRepoXhtml_address_type.CheckoutPhaseTwoInstance(backRepo, &xhtml_address_typeDB)
		}
	}
}

// CopyBasicFieldsFromXhtml_address_type
func (xhtml_address_typeDB *Xhtml_address_typeDB) CopyBasicFieldsFromXhtml_address_type(xhtml_address_type *models.Xhtml_address_type) {
	// insertion point for fields commit

	xhtml_address_typeDB.Name_Data.String = xhtml_address_type.Name
	xhtml_address_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_address_type_WOP
func (xhtml_address_typeDB *Xhtml_address_typeDB) CopyBasicFieldsFromXhtml_address_type_WOP(xhtml_address_type *models.Xhtml_address_type_WOP) {
	// insertion point for fields commit

	xhtml_address_typeDB.Name_Data.String = xhtml_address_type.Name
	xhtml_address_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_address_typeWOP
func (xhtml_address_typeDB *Xhtml_address_typeDB) CopyBasicFieldsFromXhtml_address_typeWOP(xhtml_address_type *Xhtml_address_typeWOP) {
	// insertion point for fields commit

	xhtml_address_typeDB.Name_Data.String = xhtml_address_type.Name
	xhtml_address_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToXhtml_address_type
func (xhtml_address_typeDB *Xhtml_address_typeDB) CopyBasicFieldsToXhtml_address_type(xhtml_address_type *models.Xhtml_address_type) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_address_type.Name = xhtml_address_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_address_type_WOP
func (xhtml_address_typeDB *Xhtml_address_typeDB) CopyBasicFieldsToXhtml_address_type_WOP(xhtml_address_type *models.Xhtml_address_type_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_address_type.Name = xhtml_address_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_address_typeWOP
func (xhtml_address_typeDB *Xhtml_address_typeDB) CopyBasicFieldsToXhtml_address_typeWOP(xhtml_address_type *Xhtml_address_typeWOP) {
	xhtml_address_type.ID = int(xhtml_address_typeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_address_type.Name = xhtml_address_typeDB.Name_Data.String
}

// Backup generates a json file from a slice of all Xhtml_address_typeDB instances in the backrepo
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Xhtml_address_typeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_address_typeDB, 0)
	for _, xhtml_address_typeDB := range backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB {
		forBackup = append(forBackup, xhtml_address_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Xhtml_address_type ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Xhtml_address_type file", err.Error())
	}
}

// Backup generates a json file from a slice of all Xhtml_address_typeDB instances in the backrepo
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_address_typeDB, 0)
	for _, xhtml_address_typeDB := range backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB {
		forBackup = append(forBackup, xhtml_address_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Xhtml_address_type")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Xhtml_address_type_Fields, -1)
	for _, xhtml_address_typeDB := range forBackup {

		var xhtml_address_typeWOP Xhtml_address_typeWOP
		xhtml_address_typeDB.CopyBasicFieldsToXhtml_address_typeWOP(&xhtml_address_typeWOP)

		row := sh.AddRow()
		row.WriteStruct(&xhtml_address_typeWOP, -1)
	}
}

// RestoreXL from the "Xhtml_address_type" sheet all Xhtml_address_typeDB instances
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoXhtml_address_typeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Xhtml_address_type"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoXhtml_address_type.rowVisitorXhtml_address_type)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) rowVisitorXhtml_address_type(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var xhtml_address_typeWOP Xhtml_address_typeWOP
		row.ReadStruct(&xhtml_address_typeWOP)

		// add the unmarshalled struct to the stage
		xhtml_address_typeDB := new(Xhtml_address_typeDB)
		xhtml_address_typeDB.CopyBasicFieldsFromXhtml_address_typeWOP(&xhtml_address_typeWOP)

		xhtml_address_typeDB_ID_atBackupTime := xhtml_address_typeDB.ID
		xhtml_address_typeDB.ID = 0
		query := backRepoXhtml_address_type.db.Create(xhtml_address_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB[xhtml_address_typeDB.ID] = xhtml_address_typeDB
		BackRepoXhtml_address_typeid_atBckpTime_newID[xhtml_address_typeDB_ID_atBackupTime] = xhtml_address_typeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Xhtml_address_typeDB.json" in dirPath that stores an array
// of Xhtml_address_typeDB and stores it in the database
// the map BackRepoXhtml_address_typeid_atBckpTime_newID is updated accordingly
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoXhtml_address_typeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Xhtml_address_typeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Xhtml_address_type file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Xhtml_address_typeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Xhtml_address_typeDBID_Xhtml_address_typeDB
	for _, xhtml_address_typeDB := range forRestore {

		xhtml_address_typeDB_ID_atBackupTime := xhtml_address_typeDB.ID
		xhtml_address_typeDB.ID = 0
		query := backRepoXhtml_address_type.db.Create(xhtml_address_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB[xhtml_address_typeDB.ID] = xhtml_address_typeDB
		BackRepoXhtml_address_typeid_atBckpTime_newID[xhtml_address_typeDB_ID_atBackupTime] = xhtml_address_typeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Xhtml_address_type file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Xhtml_address_type>id_atBckpTime_newID
// to compute new index
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) RestorePhaseTwo() {

	for _, xhtml_address_typeDB := range backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB {

		// next line of code is to avert unused variable compilation error
		_ = xhtml_address_typeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoXhtml_address_type.db.Model(xhtml_address_typeDB).Updates(*xhtml_address_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoXhtml_address_type.ResetReversePointers commits all staged instances of Xhtml_address_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_address_type := range backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr {
		backRepoXhtml_address_type.ResetReversePointersInstance(backRepo, idx, xhtml_address_type)
	}

	return
}

func (backRepoXhtml_address_type *BackRepoXhtml_address_typeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, xhtml_address_type *models.Xhtml_address_type) (Error error) {

	// fetch matching xhtml_address_typeDB
	if xhtml_address_typeDB, ok := backRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typeDB[idx]; ok {
		_ = xhtml_address_typeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoXhtml_address_typeid_atBckpTime_newID map[uint]uint