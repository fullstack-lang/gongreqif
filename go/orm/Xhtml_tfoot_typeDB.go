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
var dummy_Xhtml_tfoot_type_sql sql.NullBool
var dummy_Xhtml_tfoot_type_time time.Duration
var dummy_Xhtml_tfoot_type_sort sort.Float64Slice

// Xhtml_tfoot_typeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model xhtml_tfoot_typeAPI
type Xhtml_tfoot_typeAPI struct {
	gorm.Model

	models.Xhtml_tfoot_type_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Xhtml_tfoot_typePointersEncoding Xhtml_tfoot_typePointersEncoding
}

// Xhtml_tfoot_typePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Xhtml_tfoot_typePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Xhtml_tfoot_typeDB describes a xhtml_tfoot_type in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model xhtml_tfoot_typeDB
type Xhtml_tfoot_typeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field xhtml_tfoot_typeDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Xhtml_tfoot_typePointersEncoding
}

// Xhtml_tfoot_typeDBs arrays xhtml_tfoot_typeDBs
// swagger:response xhtml_tfoot_typeDBsResponse
type Xhtml_tfoot_typeDBs []Xhtml_tfoot_typeDB

// Xhtml_tfoot_typeDBResponse provides response
// swagger:response xhtml_tfoot_typeDBResponse
type Xhtml_tfoot_typeDBResponse struct {
	Xhtml_tfoot_typeDB
}

// Xhtml_tfoot_typeWOP is a Xhtml_tfoot_type without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Xhtml_tfoot_typeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Xhtml_tfoot_type_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoXhtml_tfoot_typeStruct struct {
	// stores Xhtml_tfoot_typeDB according to their gorm ID
	Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB map[uint]*Xhtml_tfoot_typeDB

	// stores Xhtml_tfoot_typeDB ID according to Xhtml_tfoot_type address
	Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID map[*models.Xhtml_tfoot_type]uint

	// stores Xhtml_tfoot_type according to their gorm ID
	Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr map[uint]*models.Xhtml_tfoot_type

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoXhtml_tfoot_type.stage
	return
}

func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) GetDB() *gorm.DB {
	return backRepoXhtml_tfoot_type.db
}

// GetXhtml_tfoot_typeDBFromXhtml_tfoot_typePtr is a handy function to access the back repo instance from the stage instance
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) GetXhtml_tfoot_typeDBFromXhtml_tfoot_typePtr(xhtml_tfoot_type *models.Xhtml_tfoot_type) (xhtml_tfoot_typeDB *Xhtml_tfoot_typeDB) {
	id := backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID[xhtml_tfoot_type]
	xhtml_tfoot_typeDB = backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB[id]
	return
}

// BackRepoXhtml_tfoot_type.CommitPhaseOne commits all staged instances of Xhtml_tfoot_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for xhtml_tfoot_type := range stage.Xhtml_tfoot_types {
		backRepoXhtml_tfoot_type.CommitPhaseOneInstance(xhtml_tfoot_type)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, xhtml_tfoot_type := range backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr {
		if _, ok := stage.Xhtml_tfoot_types[xhtml_tfoot_type]; !ok {
			backRepoXhtml_tfoot_type.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoXhtml_tfoot_type.CommitDeleteInstance commits deletion of Xhtml_tfoot_type to the BackRepo
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) CommitDeleteInstance(id uint) (Error error) {

	xhtml_tfoot_type := backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr[id]

	// xhtml_tfoot_type is not staged anymore, remove xhtml_tfoot_typeDB
	xhtml_tfoot_typeDB := backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB[id]
	query := backRepoXhtml_tfoot_type.db.Unscoped().Delete(&xhtml_tfoot_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID, xhtml_tfoot_type)
	delete(backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr, id)
	delete(backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB, id)

	return
}

// BackRepoXhtml_tfoot_type.CommitPhaseOneInstance commits xhtml_tfoot_type staged instances of Xhtml_tfoot_type to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) CommitPhaseOneInstance(xhtml_tfoot_type *models.Xhtml_tfoot_type) (Error error) {

	// check if the xhtml_tfoot_type is not commited yet
	if _, ok := backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID[xhtml_tfoot_type]; ok {
		return
	}

	// initiate xhtml_tfoot_type
	var xhtml_tfoot_typeDB Xhtml_tfoot_typeDB
	xhtml_tfoot_typeDB.CopyBasicFieldsFromXhtml_tfoot_type(xhtml_tfoot_type)

	query := backRepoXhtml_tfoot_type.db.Create(&xhtml_tfoot_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID[xhtml_tfoot_type] = xhtml_tfoot_typeDB.ID
	backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr[xhtml_tfoot_typeDB.ID] = xhtml_tfoot_type
	backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB[xhtml_tfoot_typeDB.ID] = &xhtml_tfoot_typeDB

	return
}

// BackRepoXhtml_tfoot_type.CommitPhaseTwo commits all staged instances of Xhtml_tfoot_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_tfoot_type := range backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr {
		backRepoXhtml_tfoot_type.CommitPhaseTwoInstance(backRepo, idx, xhtml_tfoot_type)
	}

	return
}

// BackRepoXhtml_tfoot_type.CommitPhaseTwoInstance commits {{structname }} of models.Xhtml_tfoot_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, xhtml_tfoot_type *models.Xhtml_tfoot_type) (Error error) {

	// fetch matching xhtml_tfoot_typeDB
	if xhtml_tfoot_typeDB, ok := backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB[idx]; ok {

		xhtml_tfoot_typeDB.CopyBasicFieldsFromXhtml_tfoot_type(xhtml_tfoot_type)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoXhtml_tfoot_type.db.Save(&xhtml_tfoot_typeDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Xhtml_tfoot_type intance %s", xhtml_tfoot_type.Name))
		return err
	}

	return
}

// BackRepoXhtml_tfoot_type.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) CheckoutPhaseOne() (Error error) {

	xhtml_tfoot_typeDBArray := make([]Xhtml_tfoot_typeDB, 0)
	query := backRepoXhtml_tfoot_type.db.Find(&xhtml_tfoot_typeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	xhtml_tfoot_typeInstancesToBeRemovedFromTheStage := make(map[*models.Xhtml_tfoot_type]any)
	for key, value := range backRepoXhtml_tfoot_type.stage.Xhtml_tfoot_types {
		xhtml_tfoot_typeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, xhtml_tfoot_typeDB := range xhtml_tfoot_typeDBArray {
		backRepoXhtml_tfoot_type.CheckoutPhaseOneInstance(&xhtml_tfoot_typeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		xhtml_tfoot_type, ok := backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr[xhtml_tfoot_typeDB.ID]
		if ok {
			delete(xhtml_tfoot_typeInstancesToBeRemovedFromTheStage, xhtml_tfoot_type)
		}
	}

	// remove from stage and back repo's 3 maps all xhtml_tfoot_types that are not in the checkout
	for xhtml_tfoot_type := range xhtml_tfoot_typeInstancesToBeRemovedFromTheStage {
		xhtml_tfoot_type.Unstage(backRepoXhtml_tfoot_type.GetStage())

		// remove instance from the back repo 3 maps
		xhtml_tfoot_typeID := backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID[xhtml_tfoot_type]
		delete(backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID, xhtml_tfoot_type)
		delete(backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB, xhtml_tfoot_typeID)
		delete(backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr, xhtml_tfoot_typeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a xhtml_tfoot_typeDB that has been found in the DB, updates the backRepo and stages the
// models version of the xhtml_tfoot_typeDB
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) CheckoutPhaseOneInstance(xhtml_tfoot_typeDB *Xhtml_tfoot_typeDB) (Error error) {

	xhtml_tfoot_type, ok := backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr[xhtml_tfoot_typeDB.ID]
	if !ok {
		xhtml_tfoot_type = new(models.Xhtml_tfoot_type)

		backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr[xhtml_tfoot_typeDB.ID] = xhtml_tfoot_type
		backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID[xhtml_tfoot_type] = xhtml_tfoot_typeDB.ID

		// append model store with the new element
		xhtml_tfoot_type.Name = xhtml_tfoot_typeDB.Name_Data.String
		xhtml_tfoot_type.Stage(backRepoXhtml_tfoot_type.GetStage())
	}
	xhtml_tfoot_typeDB.CopyBasicFieldsToXhtml_tfoot_type(xhtml_tfoot_type)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	xhtml_tfoot_type.Stage(backRepoXhtml_tfoot_type.GetStage())

	// preserve pointer to xhtml_tfoot_typeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB)[xhtml_tfoot_typeDB hold variable pointers
	xhtml_tfoot_typeDB_Data := *xhtml_tfoot_typeDB
	preservedPtrToXhtml_tfoot_type := &xhtml_tfoot_typeDB_Data
	backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB[xhtml_tfoot_typeDB.ID] = preservedPtrToXhtml_tfoot_type

	return
}

// BackRepoXhtml_tfoot_type.CheckoutPhaseTwo Checkouts all staged instances of Xhtml_tfoot_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, xhtml_tfoot_typeDB := range backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB {
		backRepoXhtml_tfoot_type.CheckoutPhaseTwoInstance(backRepo, xhtml_tfoot_typeDB)
	}
	return
}

// BackRepoXhtml_tfoot_type.CheckoutPhaseTwoInstance Checkouts staged instances of Xhtml_tfoot_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, xhtml_tfoot_typeDB *Xhtml_tfoot_typeDB) (Error error) {

	xhtml_tfoot_type := backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr[xhtml_tfoot_typeDB.ID]

	xhtml_tfoot_typeDB.DecodePointers(backRepo, xhtml_tfoot_type)

	return
}

func (xhtml_tfoot_typeDB *Xhtml_tfoot_typeDB) DecodePointers(backRepo *BackRepoStruct, xhtml_tfoot_type *models.Xhtml_tfoot_type) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitXhtml_tfoot_type allows commit of a single xhtml_tfoot_type (if already staged)
func (backRepo *BackRepoStruct) CommitXhtml_tfoot_type(xhtml_tfoot_type *models.Xhtml_tfoot_type) {
	backRepo.BackRepoXhtml_tfoot_type.CommitPhaseOneInstance(xhtml_tfoot_type)
	if id, ok := backRepo.BackRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID[xhtml_tfoot_type]; ok {
		backRepo.BackRepoXhtml_tfoot_type.CommitPhaseTwoInstance(backRepo, id, xhtml_tfoot_type)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitXhtml_tfoot_type allows checkout of a single xhtml_tfoot_type (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutXhtml_tfoot_type(xhtml_tfoot_type *models.Xhtml_tfoot_type) {
	// check if the xhtml_tfoot_type is staged
	if _, ok := backRepo.BackRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID[xhtml_tfoot_type]; ok {

		if id, ok := backRepo.BackRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typePtr_Xhtml_tfoot_typeDBID[xhtml_tfoot_type]; ok {
			var xhtml_tfoot_typeDB Xhtml_tfoot_typeDB
			xhtml_tfoot_typeDB.ID = id

			if err := backRepo.BackRepoXhtml_tfoot_type.db.First(&xhtml_tfoot_typeDB, id).Error; err != nil {
				log.Fatalln("CheckoutXhtml_tfoot_type : Problem with getting object with id:", id)
			}
			backRepo.BackRepoXhtml_tfoot_type.CheckoutPhaseOneInstance(&xhtml_tfoot_typeDB)
			backRepo.BackRepoXhtml_tfoot_type.CheckoutPhaseTwoInstance(backRepo, &xhtml_tfoot_typeDB)
		}
	}
}

// CopyBasicFieldsFromXhtml_tfoot_type
func (xhtml_tfoot_typeDB *Xhtml_tfoot_typeDB) CopyBasicFieldsFromXhtml_tfoot_type(xhtml_tfoot_type *models.Xhtml_tfoot_type) {
	// insertion point for fields commit

	xhtml_tfoot_typeDB.Name_Data.String = xhtml_tfoot_type.Name
	xhtml_tfoot_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_tfoot_type_WOP
func (xhtml_tfoot_typeDB *Xhtml_tfoot_typeDB) CopyBasicFieldsFromXhtml_tfoot_type_WOP(xhtml_tfoot_type *models.Xhtml_tfoot_type_WOP) {
	// insertion point for fields commit

	xhtml_tfoot_typeDB.Name_Data.String = xhtml_tfoot_type.Name
	xhtml_tfoot_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromXhtml_tfoot_typeWOP
func (xhtml_tfoot_typeDB *Xhtml_tfoot_typeDB) CopyBasicFieldsFromXhtml_tfoot_typeWOP(xhtml_tfoot_type *Xhtml_tfoot_typeWOP) {
	// insertion point for fields commit

	xhtml_tfoot_typeDB.Name_Data.String = xhtml_tfoot_type.Name
	xhtml_tfoot_typeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToXhtml_tfoot_type
func (xhtml_tfoot_typeDB *Xhtml_tfoot_typeDB) CopyBasicFieldsToXhtml_tfoot_type(xhtml_tfoot_type *models.Xhtml_tfoot_type) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_tfoot_type.Name = xhtml_tfoot_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_tfoot_type_WOP
func (xhtml_tfoot_typeDB *Xhtml_tfoot_typeDB) CopyBasicFieldsToXhtml_tfoot_type_WOP(xhtml_tfoot_type *models.Xhtml_tfoot_type_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_tfoot_type.Name = xhtml_tfoot_typeDB.Name_Data.String
}

// CopyBasicFieldsToXhtml_tfoot_typeWOP
func (xhtml_tfoot_typeDB *Xhtml_tfoot_typeDB) CopyBasicFieldsToXhtml_tfoot_typeWOP(xhtml_tfoot_type *Xhtml_tfoot_typeWOP) {
	xhtml_tfoot_type.ID = int(xhtml_tfoot_typeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	xhtml_tfoot_type.Name = xhtml_tfoot_typeDB.Name_Data.String
}

// Backup generates a json file from a slice of all Xhtml_tfoot_typeDB instances in the backrepo
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Xhtml_tfoot_typeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_tfoot_typeDB, 0)
	for _, xhtml_tfoot_typeDB := range backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB {
		forBackup = append(forBackup, xhtml_tfoot_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Xhtml_tfoot_type ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Xhtml_tfoot_type file", err.Error())
	}
}

// Backup generates a json file from a slice of all Xhtml_tfoot_typeDB instances in the backrepo
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Xhtml_tfoot_typeDB, 0)
	for _, xhtml_tfoot_typeDB := range backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB {
		forBackup = append(forBackup, xhtml_tfoot_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Xhtml_tfoot_type")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Xhtml_tfoot_type_Fields, -1)
	for _, xhtml_tfoot_typeDB := range forBackup {

		var xhtml_tfoot_typeWOP Xhtml_tfoot_typeWOP
		xhtml_tfoot_typeDB.CopyBasicFieldsToXhtml_tfoot_typeWOP(&xhtml_tfoot_typeWOP)

		row := sh.AddRow()
		row.WriteStruct(&xhtml_tfoot_typeWOP, -1)
	}
}

// RestoreXL from the "Xhtml_tfoot_type" sheet all Xhtml_tfoot_typeDB instances
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoXhtml_tfoot_typeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Xhtml_tfoot_type"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoXhtml_tfoot_type.rowVisitorXhtml_tfoot_type)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) rowVisitorXhtml_tfoot_type(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var xhtml_tfoot_typeWOP Xhtml_tfoot_typeWOP
		row.ReadStruct(&xhtml_tfoot_typeWOP)

		// add the unmarshalled struct to the stage
		xhtml_tfoot_typeDB := new(Xhtml_tfoot_typeDB)
		xhtml_tfoot_typeDB.CopyBasicFieldsFromXhtml_tfoot_typeWOP(&xhtml_tfoot_typeWOP)

		xhtml_tfoot_typeDB_ID_atBackupTime := xhtml_tfoot_typeDB.ID
		xhtml_tfoot_typeDB.ID = 0
		query := backRepoXhtml_tfoot_type.db.Create(xhtml_tfoot_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB[xhtml_tfoot_typeDB.ID] = xhtml_tfoot_typeDB
		BackRepoXhtml_tfoot_typeid_atBckpTime_newID[xhtml_tfoot_typeDB_ID_atBackupTime] = xhtml_tfoot_typeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Xhtml_tfoot_typeDB.json" in dirPath that stores an array
// of Xhtml_tfoot_typeDB and stores it in the database
// the map BackRepoXhtml_tfoot_typeid_atBckpTime_newID is updated accordingly
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoXhtml_tfoot_typeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Xhtml_tfoot_typeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Xhtml_tfoot_type file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Xhtml_tfoot_typeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB
	for _, xhtml_tfoot_typeDB := range forRestore {

		xhtml_tfoot_typeDB_ID_atBackupTime := xhtml_tfoot_typeDB.ID
		xhtml_tfoot_typeDB.ID = 0
		query := backRepoXhtml_tfoot_type.db.Create(xhtml_tfoot_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB[xhtml_tfoot_typeDB.ID] = xhtml_tfoot_typeDB
		BackRepoXhtml_tfoot_typeid_atBckpTime_newID[xhtml_tfoot_typeDB_ID_atBackupTime] = xhtml_tfoot_typeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Xhtml_tfoot_type file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Xhtml_tfoot_type>id_atBckpTime_newID
// to compute new index
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) RestorePhaseTwo() {

	for _, xhtml_tfoot_typeDB := range backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB {

		// next line of code is to avert unused variable compilation error
		_ = xhtml_tfoot_typeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoXhtml_tfoot_type.db.Model(xhtml_tfoot_typeDB).Updates(*xhtml_tfoot_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoXhtml_tfoot_type.ResetReversePointers commits all staged instances of Xhtml_tfoot_type to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, xhtml_tfoot_type := range backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr {
		backRepoXhtml_tfoot_type.ResetReversePointersInstance(backRepo, idx, xhtml_tfoot_type)
	}

	return
}

func (backRepoXhtml_tfoot_type *BackRepoXhtml_tfoot_typeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, xhtml_tfoot_type *models.Xhtml_tfoot_type) (Error error) {

	// fetch matching xhtml_tfoot_typeDB
	if xhtml_tfoot_typeDB, ok := backRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typeDB[idx]; ok {
		_ = xhtml_tfoot_typeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoXhtml_tfoot_typeid_atBckpTime_newID map[uint]uint