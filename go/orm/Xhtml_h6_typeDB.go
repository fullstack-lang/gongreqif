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
var dummy_Xhtml_h6_type_sql sql.NullBool
var dummy_Xhtml_h6_type_time time.Duration
var dummy_Xhtml_h6_type_sort sort.Float64Slice

// Xhtml_h6_typeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model xhtml_h6_typeAPI
type Xhtml_h6_typeAPI struct {
	gorm.Model

	models.Xhtml_h6_type_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Xhtml_h6_typePointersEncoding Xhtml_h6_typePointersEncoding
}

// Xhtml_h6_typePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Xhtml_h6_typePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Xhtml_h6_typeDB describes a xhtml_h6_type in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model xhtml_h6_typeDB
type Xhtml_h6_typeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field xhtml_h6_typeDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Xhtml_h6_typePointersEncoding
}

// Xhtml_h6_typeDBs arrays xhtml_h6_typeDBs
// swagger:response xhtml_h6_typeDBsResponse
type Xhtml_h6_typeDBs []Xhtml_h6_typeDB

// Xhtml_h6_typeDBResponse provides response
// swagger:response xhtml_h6_typeDBResponse
type Xhtml_h6_typeDBResponse struct {
	Xhtml_h6_typeDB
}

// Xhtml_h6_typeWOP is a Xhtml_h6_type without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Xhtml_h6_typeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Xhtml_h6_type_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoXhtml_h6_typeStruct struct {
	// stores Xhtml_h6_typeDB according to their gorm ID
	Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB map[uint]*Xhtml_h6_typeDB

	// stores Xhtml_h6_typeDB ID according to Xhtml_h6_type address
	Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID map[*models.Xhtml_h6_type]uint

	// stores Xhtml_h6_type according to their gorm ID
	Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr map[uint]*models.Xhtml_h6_type

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoXhtml_h6_type.stage
	return
}

func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) GetDB() *gorm.DB {
	return backRepoXhtml_h6_type.db
}

// GetXhtml_h6_typeDBFromXhtml_h6_typePtr is a handy function to access the back repo instance from the stage instance
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) GetXhtml_h6_typeDBFromXhtml_h6_typePtr(xhtml_h6_type *models.Xhtml_h6_type) (xhtml_h6_typeDB *Xhtml_h6_typeDB) {
	id := backRepoXhtml_h6_type.Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID[xhtml_h6_type]
	xhtml_h6_typeDB = backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB[id]
	return
}

// BackRepoXhtml_h6_type.CommitPhaseOne commits all staged instances of Xhtml_h6_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for xhtml_h6_type := range stage.Xhtml_h6_types {
		backRepoXhtml_h6_type.CommitPhaseOneInstance(xhtml_h6_type)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, xhtml_h6_type := range backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr {
		if _, ok := stage.Xhtml_h6_types[xhtml_h6_type]; !ok {
			backRepoXhtml_h6_type.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoXhtml_h6_type.CommitDeleteInstance commits deletion of Xhtml_h6_type to the BackRepo
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) CommitDeleteInstance(id uint) (Error error) {

	xhtml_h6_type := backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr[id]

	// xhtml_h6_type is not staged anymore, remove xhtml_h6_typeDB
	xhtml_h6_typeDB := backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB[id]
	query := backRepoXhtml_h6_type.db.Unscoped().Delete(&xhtml_h6_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoXhtml_h6_type.Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID, xhtml_h6_type)
	delete(backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr, id)
	delete(backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB, id)

	return
}

// BackRepoXhtml_h6_type.CommitPhaseOneInstance commits xhtml_h6_type staged instances of Xhtml_h6_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) CommitPhaseOneInstance(xhtml_h6_type *models.Xhtml_h6_type) (Error error) {

	// check if the xhtml_h6_type is not commited yet
	if _, ok := backRepoXhtml_h6_type.Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID[xhtml_h6_type]; ok {
		return
	}

	// initiate xhtml_h6_type
	var xhtml_h6_typeDB Xhtml_h6_typeDB
	xhtml_h6_typeDB.CopyBasicFieldsFromXhtml_h6_type(xhtml_h6_type)

	query := backRepoXhtml_h6_type.db.Create(&xhtml_h6_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoXhtml_h6_type.Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID[xhtml_h6_type] = xhtml_h6_typeDB.ID
	backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr[xhtml_h6_typeDB.ID] = xhtml_h6_type
	backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB[xhtml_h6_typeDB.ID] = &xhtml_h6_typeDB

	return
}

// BackRepoXhtml_h6_type.CommitPhaseTwo commits all staged instances of Xhtml_h6_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_h6_type := range backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr {
		backRepoXhtml_h6_type.CommitPhaseTwoInstance(backRepo, idx, xhtml_h6_type)
	}

	return
}

// BackRepoXhtml_h6_type.CommitPhaseTwoInstance commits {{structname }} of models.Xhtml_h6_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, xhtml_h6_type *models.Xhtml_h6_type) (Error error) {

	// fetch matching xhtml_h6_typeDB
	if xhtml_h6_typeDB, ok := backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB[idx]; ok {

		xhtml_h6_typeDB.CopyBasicFieldsFromXhtml_h6_type(xhtml_h6_type)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoXhtml_h6_type.db.Save(&xhtml_h6_typeDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Xhtml_h6_type intance %s", xhtml_h6_type.Name))
		return err
	}

	return
}

// BackRepoXhtml_h6_type.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) CheckoutPhaseOne() (Error error) {

	xhtml_h6_typeDBArray := make([]Xhtml_h6_typeDB, 0)
	query := backRepoXhtml_h6_type.db.Find(&xhtml_h6_typeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	xhtml_h6_typeInstancesToBeRemovedFromTheStage := make(map[*models.Xhtml_h6_type]any)
	for key, value := range backRepoXhtml_h6_type.stage.Xhtml_h6_types {
		xhtml_h6_typeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, xhtml_h6_typeDB := range xhtml_h6_typeDBArray {
		backRepoXhtml_h6_type.CheckoutPhaseOneInstance(&xhtml_h6_typeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		xhtml_h6_type, ok := backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr[xhtml_h6_typeDB.ID]
		if ok {
			delete(xhtml_h6_typeInstancesToBeRemovedFromTheStage, xhtml_h6_type)
		}
	}

	// remove from stage and back repo's 3 maps all xhtml_h6_types that are not in the checkout
	for xhtml_h6_type := range xhtml_h6_typeInstancesToBeRemovedFromTheStage {
		xhtml_h6_type.Unstage(backRepoXhtml_h6_type.GetStage())

		// remove instance from the back repo 3 maps
		xhtml_h6_typeID := backRepoXhtml_h6_type.Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID[xhtml_h6_type]
		delete(backRepoXhtml_h6_type.Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID, xhtml_h6_type)
		delete(backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB, xhtml_h6_typeID)
		delete(backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr, xhtml_h6_typeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a xhtml_h6_typeDB that has been found in the DB, updates the backRepo and stages the
// models version of the xhtml_h6_typeDB
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) CheckoutPhaseOneInstance(xhtml_h6_typeDB *Xhtml_h6_typeDB) (Error error) {

	xhtml_h6_type, ok := backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr[xhtml_h6_typeDB.ID]
	if !ok {
		xhtml_h6_type = new(models.Xhtml_h6_type)

		backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr[xhtml_h6_typeDB.ID] = xhtml_h6_type
		backRepoXhtml_h6_type.Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID[xhtml_h6_type] = xhtml_h6_typeDB.ID

		// append model store with the new element
		xhtml_h6_type.Name = xhtml_h6_typeDB.Name_Data.String
		xhtml_h6_type.Stage(backRepoXhtml_h6_type.GetStage())
	}
	xhtml_h6_typeDB.CopyBasicFieldsToXhtml_h6_type(xhtml_h6_type)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	xhtml_h6_type.Stage(backRepoXhtml_h6_type.GetStage())

	// preserve pointer to xhtml_h6_typeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB)[xhtml_h6_typeDB hold variable pointers
	xhtml_h6_typeDB_Data := *xhtml_h6_typeDB
	preservedPtrToXhtml_h6_type := &xhtml_h6_typeDB_Data
	backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB[xhtml_h6_typeDB.ID] = preservedPtrToXhtml_h6_type

	return
}

// BackRepoXhtml_h6_type.CheckoutPhaseTwo Checkouts all staged instances of Xhtml_h6_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, xhtml_h6_typeDB := range backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB {
		backRepoXhtml_h6_type.CheckoutPhaseTwoInstance(backRepo, xhtml_h6_typeDB)
	}
	return
}

// BackRepoXhtml_h6_type.CheckoutPhaseTwoInstance Checkouts staged instances of Xhtml_h6_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, xhtml_h6_typeDB *Xhtml_h6_typeDB) (Error error) {

	xhtml_h6_type := backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr[xhtml_h6_typeDB.ID]

	xhtml_h6_typeDB.DecodePointers(backRepo, xhtml_h6_type)

	return
}

func (xhtml_h6_typeDB *Xhtml_h6_typeDB) DecodePointers(backRepo *BackRepoStruct, xhtml_h6_type *models.Xhtml_h6_type) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitXhtml_h6_type allows commit of a single xhtml_h6_type (if already staged)
func (backRepo *BackRepoStruct) CommitXhtml_h6_type(xhtml_h6_type *models.Xhtml_h6_type) {
	backRepo.BackRepoXhtml_h6_type.CommitPhaseOneInstance(xhtml_h6_type)
	if id, ok := backRepo.BackRepoXhtml_h6_type.Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID[xhtml_h6_type]; ok {
		backRepo.BackRepoXhtml_h6_type.CommitPhaseTwoInstance(backRepo, id, xhtml_h6_type)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitXhtml_h6_type allows checkout of a single xhtml_h6_type (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutXhtml_h6_type(xhtml_h6_type *models.Xhtml_h6_type) {
	// check if the xhtml_h6_type is staged
	if _, ok := backRepo.BackRepoXhtml_h6_type.Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID[xhtml_h6_type]; ok {

		if id, ok := backRepo.BackRepoXhtml_h6_type.Map_Xhtml_h6_typePtr_Xhtml_h6_typeDBID[xhtml_h6_type]; ok {
			var xhtml_h6_typeDB Xhtml_h6_typeDB
			xhtml_h6_typeDB.ID = id

			if err := backRepo.BackRepoXhtml_h6_type.db.First(&xhtml_h6_typeDB, id).Error; err != nil {
				log.Fatalln("CheckoutXhtml_h6_type : Problem with getting object with id:", id)
			}
			backRepo.BackRepoXhtml_h6_type.CheckoutPhaseOneInstance(&xhtml_h6_typeDB)
			backRepo.BackRepoXhtml_h6_type.CheckoutPhaseTwoInstance(backRepo, &xhtml_h6_typeDB)
		}
	}
}

// CopyBasicFieldsFromXhtml_h6_type
func (xhtml_h6_typeDB *Xhtml_h6_typeDB) CopyBasicFieldsFromXhtml_h6_type(xhtml_h6_type *models.Xhtml_h6_type) {
	// insertion point for fields commit

	xhtml_h6_typeDB.Name_Data.String = xhtml_h6_type.Name
	xhtml_h6_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_h6_type_WOP
func (xhtml_h6_typeDB *Xhtml_h6_typeDB) CopyBasicFieldsFromXhtml_h6_type_WOP(xhtml_h6_type *models.Xhtml_h6_type_WOP) {
	// insertion point for fields commit

	xhtml_h6_typeDB.Name_Data.String = xhtml_h6_type.Name
	xhtml_h6_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_h6_typeWOP
func (xhtml_h6_typeDB *Xhtml_h6_typeDB) CopyBasicFieldsFromXhtml_h6_typeWOP(xhtml_h6_type *Xhtml_h6_typeWOP) {
	// insertion point for fields commit

	xhtml_h6_typeDB.Name_Data.String = xhtml_h6_type.Name
	xhtml_h6_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToXhtml_h6_type
func (xhtml_h6_typeDB *Xhtml_h6_typeDB) CopyBasicFieldsToXhtml_h6_type(xhtml_h6_type *models.Xhtml_h6_type) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_h6_type.Name = xhtml_h6_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_h6_type_WOP
func (xhtml_h6_typeDB *Xhtml_h6_typeDB) CopyBasicFieldsToXhtml_h6_type_WOP(xhtml_h6_type *models.Xhtml_h6_type_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_h6_type.Name = xhtml_h6_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_h6_typeWOP
func (xhtml_h6_typeDB *Xhtml_h6_typeDB) CopyBasicFieldsToXhtml_h6_typeWOP(xhtml_h6_type *Xhtml_h6_typeWOP) {
	xhtml_h6_type.ID = int(xhtml_h6_typeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_h6_type.Name = xhtml_h6_typeDB.Name_Data.String
}

// Backup generates a json file from a slice of all Xhtml_h6_typeDB instances in the backrepo
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Xhtml_h6_typeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_h6_typeDB, 0)
	for _, xhtml_h6_typeDB := range backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB {
		forBackup = append(forBackup, xhtml_h6_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Xhtml_h6_type ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Xhtml_h6_type file", err.Error())
	}
}

// Backup generates a json file from a slice of all Xhtml_h6_typeDB instances in the backrepo
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_h6_typeDB, 0)
	for _, xhtml_h6_typeDB := range backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB {
		forBackup = append(forBackup, xhtml_h6_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Xhtml_h6_type")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Xhtml_h6_type_Fields, -1)
	for _, xhtml_h6_typeDB := range forBackup {

		var xhtml_h6_typeWOP Xhtml_h6_typeWOP
		xhtml_h6_typeDB.CopyBasicFieldsToXhtml_h6_typeWOP(&xhtml_h6_typeWOP)

		row := sh.AddRow()
		row.WriteStruct(&xhtml_h6_typeWOP, -1)
	}
}

// RestoreXL from the "Xhtml_h6_type" sheet all Xhtml_h6_typeDB instances
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoXhtml_h6_typeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Xhtml_h6_type"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoXhtml_h6_type.rowVisitorXhtml_h6_type)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) rowVisitorXhtml_h6_type(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var xhtml_h6_typeWOP Xhtml_h6_typeWOP
		row.ReadStruct(&xhtml_h6_typeWOP)

		// add the unmarshalled struct to the stage
		xhtml_h6_typeDB := new(Xhtml_h6_typeDB)
		xhtml_h6_typeDB.CopyBasicFieldsFromXhtml_h6_typeWOP(&xhtml_h6_typeWOP)

		xhtml_h6_typeDB_ID_atBackupTime := xhtml_h6_typeDB.ID
		xhtml_h6_typeDB.ID = 0
		query := backRepoXhtml_h6_type.db.Create(xhtml_h6_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB[xhtml_h6_typeDB.ID] = xhtml_h6_typeDB
		BackRepoXhtml_h6_typeid_atBckpTime_newID[xhtml_h6_typeDB_ID_atBackupTime] = xhtml_h6_typeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Xhtml_h6_typeDB.json" in dirPath that stores an array
// of Xhtml_h6_typeDB and stores it in the database
// the map BackRepoXhtml_h6_typeid_atBckpTime_newID is updated accordingly
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoXhtml_h6_typeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Xhtml_h6_typeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Xhtml_h6_type file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Xhtml_h6_typeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB
	for _, xhtml_h6_typeDB := range forRestore {

		xhtml_h6_typeDB_ID_atBackupTime := xhtml_h6_typeDB.ID
		xhtml_h6_typeDB.ID = 0
		query := backRepoXhtml_h6_type.db.Create(xhtml_h6_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB[xhtml_h6_typeDB.ID] = xhtml_h6_typeDB
		BackRepoXhtml_h6_typeid_atBckpTime_newID[xhtml_h6_typeDB_ID_atBackupTime] = xhtml_h6_typeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Xhtml_h6_type file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Xhtml_h6_type>id_atBckpTime_newID
// to compute new index
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) RestorePhaseTwo() {

	for _, xhtml_h6_typeDB := range backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB {

		// next line of code is to avert unused variable compilation error
		_ = xhtml_h6_typeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoXhtml_h6_type.db.Model(xhtml_h6_typeDB).Updates(*xhtml_h6_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoXhtml_h6_type.ResetReversePointers commits all staged instances of Xhtml_h6_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_h6_type := range backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr {
		backRepoXhtml_h6_type.ResetReversePointersInstance(backRepo, idx, xhtml_h6_type)
	}

	return
}

func (backRepoXhtml_h6_type *BackRepoXhtml_h6_typeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, xhtml_h6_type *models.Xhtml_h6_type) (Error error) {

	// fetch matching xhtml_h6_typeDB
	if xhtml_h6_typeDB, ok := backRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typeDB[idx]; ok {
		_ = xhtml_h6_typeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoXhtml_h6_typeid_atBckpTime_newID map[uint]uint