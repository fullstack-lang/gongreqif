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
var dummy_REQ_IF_TOOL_EXTENSION_sql sql.NullBool
var dummy_REQ_IF_TOOL_EXTENSION_time time.Duration
var dummy_REQ_IF_TOOL_EXTENSION_sort sort.Float64Slice

// REQ_IF_TOOL_EXTENSIONAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model req_if_tool_extensionAPI
type REQ_IF_TOOL_EXTENSIONAPI struct {
	gorm.Model

	models.REQ_IF_TOOL_EXTENSION_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	REQ_IF_TOOL_EXTENSIONPointersEncoding REQ_IF_TOOL_EXTENSIONPointersEncoding
}

// REQ_IF_TOOL_EXTENSIONPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type REQ_IF_TOOL_EXTENSIONPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// REQ_IF_TOOL_EXTENSIONDB describes a req_if_tool_extension in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model req_if_tool_extensionDB
type REQ_IF_TOOL_EXTENSIONDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field req_if_tool_extensionDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	REQ_IF_TOOL_EXTENSIONPointersEncoding
}

// REQ_IF_TOOL_EXTENSIONDBs arrays req_if_tool_extensionDBs
// swagger:response req_if_tool_extensionDBsResponse
type REQ_IF_TOOL_EXTENSIONDBs []REQ_IF_TOOL_EXTENSIONDB

// REQ_IF_TOOL_EXTENSIONDBResponse provides response
// swagger:response req_if_tool_extensionDBResponse
type REQ_IF_TOOL_EXTENSIONDBResponse struct {
	REQ_IF_TOOL_EXTENSIONDB
}

// REQ_IF_TOOL_EXTENSIONWOP is a REQ_IF_TOOL_EXTENSION without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type REQ_IF_TOOL_EXTENSIONWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var REQ_IF_TOOL_EXTENSION_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoREQ_IF_TOOL_EXTENSIONStruct struct {
	// stores REQ_IF_TOOL_EXTENSIONDB according to their gorm ID
	Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB map[uint]*REQ_IF_TOOL_EXTENSIONDB

	// stores REQ_IF_TOOL_EXTENSIONDB ID according to REQ_IF_TOOL_EXTENSION address
	Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID map[*models.REQ_IF_TOOL_EXTENSION]uint

	// stores REQ_IF_TOOL_EXTENSION according to their gorm ID
	Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr map[uint]*models.REQ_IF_TOOL_EXTENSION

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoREQ_IF_TOOL_EXTENSION.stage
	return
}

func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) GetDB() *gorm.DB {
	return backRepoREQ_IF_TOOL_EXTENSION.db
}

// GetREQ_IF_TOOL_EXTENSIONDBFromREQ_IF_TOOL_EXTENSIONPtr is a handy function to access the back repo instance from the stage instance
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) GetREQ_IF_TOOL_EXTENSIONDBFromREQ_IF_TOOL_EXTENSIONPtr(req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION) (req_if_tool_extensionDB *REQ_IF_TOOL_EXTENSIONDB) {
	id := backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID[req_if_tool_extension]
	req_if_tool_extensionDB = backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB[id]
	return
}

// BackRepoREQ_IF_TOOL_EXTENSION.CommitPhaseOne commits all staged instances of REQ_IF_TOOL_EXTENSION to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for req_if_tool_extension := range stage.REQ_IF_TOOL_EXTENSIONs {
		backRepoREQ_IF_TOOL_EXTENSION.CommitPhaseOneInstance(req_if_tool_extension)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, req_if_tool_extension := range backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr {
		if _, ok := stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension]; !ok {
			backRepoREQ_IF_TOOL_EXTENSION.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoREQ_IF_TOOL_EXTENSION.CommitDeleteInstance commits deletion of REQ_IF_TOOL_EXTENSION to the BackRepo
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) CommitDeleteInstance(id uint) (Error error) {

	req_if_tool_extension := backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr[id]

	// req_if_tool_extension is not staged anymore, remove req_if_tool_extensionDB
	req_if_tool_extensionDB := backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB[id]
	query := backRepoREQ_IF_TOOL_EXTENSION.db.Unscoped().Delete(&req_if_tool_extensionDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID, req_if_tool_extension)
	delete(backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr, id)
	delete(backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB, id)

	return
}

// BackRepoREQ_IF_TOOL_EXTENSION.CommitPhaseOneInstance commits req_if_tool_extension staged instances of REQ_IF_TOOL_EXTENSION to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) CommitPhaseOneInstance(req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION) (Error error) {

	// check if the req_if_tool_extension is not commited yet
	if _, ok := backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID[req_if_tool_extension]; ok {
		return
	}

	// initiate req_if_tool_extension
	var req_if_tool_extensionDB REQ_IF_TOOL_EXTENSIONDB
	req_if_tool_extensionDB.CopyBasicFieldsFromREQ_IF_TOOL_EXTENSION(req_if_tool_extension)

	query := backRepoREQ_IF_TOOL_EXTENSION.db.Create(&req_if_tool_extensionDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID[req_if_tool_extension] = req_if_tool_extensionDB.ID
	backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr[req_if_tool_extensionDB.ID] = req_if_tool_extension
	backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB[req_if_tool_extensionDB.ID] = &req_if_tool_extensionDB

	return
}

// BackRepoREQ_IF_TOOL_EXTENSION.CommitPhaseTwo commits all staged instances of REQ_IF_TOOL_EXTENSION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, req_if_tool_extension := range backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr {
		backRepoREQ_IF_TOOL_EXTENSION.CommitPhaseTwoInstance(backRepo, idx, req_if_tool_extension)
	}

	return
}

// BackRepoREQ_IF_TOOL_EXTENSION.CommitPhaseTwoInstance commits {{structname }} of models.REQ_IF_TOOL_EXTENSION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION) (Error error) {

	// fetch matching req_if_tool_extensionDB
	if req_if_tool_extensionDB, ok := backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB[idx]; ok {

		req_if_tool_extensionDB.CopyBasicFieldsFromREQ_IF_TOOL_EXTENSION(req_if_tool_extension)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoREQ_IF_TOOL_EXTENSION.db.Save(&req_if_tool_extensionDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown REQ_IF_TOOL_EXTENSION intance %s", req_if_tool_extension.Name))
		return err
	}

	return
}

// BackRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) CheckoutPhaseOne() (Error error) {

	req_if_tool_extensionDBArray := make([]REQ_IF_TOOL_EXTENSIONDB, 0)
	query := backRepoREQ_IF_TOOL_EXTENSION.db.Find(&req_if_tool_extensionDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	req_if_tool_extensionInstancesToBeRemovedFromTheStage := make(map[*models.REQ_IF_TOOL_EXTENSION]any)
	for key, value := range backRepoREQ_IF_TOOL_EXTENSION.stage.REQ_IF_TOOL_EXTENSIONs {
		req_if_tool_extensionInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, req_if_tool_extensionDB := range req_if_tool_extensionDBArray {
		backRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseOneInstance(&req_if_tool_extensionDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		req_if_tool_extension, ok := backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr[req_if_tool_extensionDB.ID]
		if ok {
			delete(req_if_tool_extensionInstancesToBeRemovedFromTheStage, req_if_tool_extension)
		}
	}

	// remove from stage and back repo's 3 maps all req_if_tool_extensions that are not in the checkout
	for req_if_tool_extension := range req_if_tool_extensionInstancesToBeRemovedFromTheStage {
		req_if_tool_extension.Unstage(backRepoREQ_IF_TOOL_EXTENSION.GetStage())

		// remove instance from the back repo 3 maps
		req_if_tool_extensionID := backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID[req_if_tool_extension]
		delete(backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID, req_if_tool_extension)
		delete(backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB, req_if_tool_extensionID)
		delete(backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr, req_if_tool_extensionID)
	}

	return
}

// CheckoutPhaseOneInstance takes a req_if_tool_extensionDB that has been found in the DB, updates the backRepo and stages the
// models version of the req_if_tool_extensionDB
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) CheckoutPhaseOneInstance(req_if_tool_extensionDB *REQ_IF_TOOL_EXTENSIONDB) (Error error) {

	req_if_tool_extension, ok := backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr[req_if_tool_extensionDB.ID]
	if !ok {
		req_if_tool_extension = new(models.REQ_IF_TOOL_EXTENSION)

		backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr[req_if_tool_extensionDB.ID] = req_if_tool_extension
		backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID[req_if_tool_extension] = req_if_tool_extensionDB.ID

		// append model store with the new element
		req_if_tool_extension.Name = req_if_tool_extensionDB.Name_Data.String
		req_if_tool_extension.Stage(backRepoREQ_IF_TOOL_EXTENSION.GetStage())
	}
	req_if_tool_extensionDB.CopyBasicFieldsToREQ_IF_TOOL_EXTENSION(req_if_tool_extension)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	req_if_tool_extension.Stage(backRepoREQ_IF_TOOL_EXTENSION.GetStage())

	// preserve pointer to req_if_tool_extensionDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB)[req_if_tool_extensionDB hold variable pointers
	req_if_tool_extensionDB_Data := *req_if_tool_extensionDB
	preservedPtrToREQ_IF_TOOL_EXTENSION := &req_if_tool_extensionDB_Data
	backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB[req_if_tool_extensionDB.ID] = preservedPtrToREQ_IF_TOOL_EXTENSION

	return
}

// BackRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseTwo Checkouts all staged instances of REQ_IF_TOOL_EXTENSION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, req_if_tool_extensionDB := range backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB {
		backRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseTwoInstance(backRepo, req_if_tool_extensionDB)
	}
	return
}

// BackRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseTwoInstance Checkouts staged instances of REQ_IF_TOOL_EXTENSION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, req_if_tool_extensionDB *REQ_IF_TOOL_EXTENSIONDB) (Error error) {

	req_if_tool_extension := backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr[req_if_tool_extensionDB.ID]

	req_if_tool_extensionDB.DecodePointers(backRepo, req_if_tool_extension)

	return
}

func (req_if_tool_extensionDB *REQ_IF_TOOL_EXTENSIONDB) DecodePointers(backRepo *BackRepoStruct, req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitREQ_IF_TOOL_EXTENSION allows commit of a single req_if_tool_extension (if already staged)
func (backRepo *BackRepoStruct) CommitREQ_IF_TOOL_EXTENSION(req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION) {
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CommitPhaseOneInstance(req_if_tool_extension)
	if id, ok := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID[req_if_tool_extension]; ok {
		backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CommitPhaseTwoInstance(backRepo, id, req_if_tool_extension)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitREQ_IF_TOOL_EXTENSION allows checkout of a single req_if_tool_extension (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutREQ_IF_TOOL_EXTENSION(req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION) {
	// check if the req_if_tool_extension is staged
	if _, ok := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID[req_if_tool_extension]; ok {

		if id, ok := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID[req_if_tool_extension]; ok {
			var req_if_tool_extensionDB REQ_IF_TOOL_EXTENSIONDB
			req_if_tool_extensionDB.ID = id

			if err := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.db.First(&req_if_tool_extensionDB, id).Error; err != nil {
				log.Fatalln("CheckoutREQ_IF_TOOL_EXTENSION : Problem with getting object with id:", id)
			}
			backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseOneInstance(&req_if_tool_extensionDB)
			backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseTwoInstance(backRepo, &req_if_tool_extensionDB)
		}
	}
}

// CopyBasicFieldsFromREQ_IF_TOOL_EXTENSION
func (req_if_tool_extensionDB *REQ_IF_TOOL_EXTENSIONDB) CopyBasicFieldsFromREQ_IF_TOOL_EXTENSION(req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION) {
	// insertion point for fields commit

	req_if_tool_extensionDB.Name_Data.String = req_if_tool_extension.Name
	req_if_tool_extensionDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromREQ_IF_TOOL_EXTENSION_WOP
func (req_if_tool_extensionDB *REQ_IF_TOOL_EXTENSIONDB) CopyBasicFieldsFromREQ_IF_TOOL_EXTENSION_WOP(req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION_WOP) {
	// insertion point for fields commit

	req_if_tool_extensionDB.Name_Data.String = req_if_tool_extension.Name
	req_if_tool_extensionDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromREQ_IF_TOOL_EXTENSIONWOP
func (req_if_tool_extensionDB *REQ_IF_TOOL_EXTENSIONDB) CopyBasicFieldsFromREQ_IF_TOOL_EXTENSIONWOP(req_if_tool_extension *REQ_IF_TOOL_EXTENSIONWOP) {
	// insertion point for fields commit

	req_if_tool_extensionDB.Name_Data.String = req_if_tool_extension.Name
	req_if_tool_extensionDB.Name_Data.Valid = true
}

// CopyBasicFieldsToREQ_IF_TOOL_EXTENSION
func (req_if_tool_extensionDB *REQ_IF_TOOL_EXTENSIONDB) CopyBasicFieldsToREQ_IF_TOOL_EXTENSION(req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION) {
	// insertion point for checkout of basic fields (back repo to stage)
	req_if_tool_extension.Name = req_if_tool_extensionDB.Name_Data.String
}

// CopyBasicFieldsToREQ_IF_TOOL_EXTENSION_WOP
func (req_if_tool_extensionDB *REQ_IF_TOOL_EXTENSIONDB) CopyBasicFieldsToREQ_IF_TOOL_EXTENSION_WOP(req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	req_if_tool_extension.Name = req_if_tool_extensionDB.Name_Data.String
}

// CopyBasicFieldsToREQ_IF_TOOL_EXTENSIONWOP
func (req_if_tool_extensionDB *REQ_IF_TOOL_EXTENSIONDB) CopyBasicFieldsToREQ_IF_TOOL_EXTENSIONWOP(req_if_tool_extension *REQ_IF_TOOL_EXTENSIONWOP) {
	req_if_tool_extension.ID = int(req_if_tool_extensionDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	req_if_tool_extension.Name = req_if_tool_extensionDB.Name_Data.String
}

// Backup generates a json file from a slice of all REQ_IF_TOOL_EXTENSIONDB instances in the backrepo
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "REQ_IF_TOOL_EXTENSIONDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*REQ_IF_TOOL_EXTENSIONDB, 0)
	for _, req_if_tool_extensionDB := range backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB {
		forBackup = append(forBackup, req_if_tool_extensionDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json REQ_IF_TOOL_EXTENSION ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json REQ_IF_TOOL_EXTENSION file", err.Error())
	}
}

// Backup generates a json file from a slice of all REQ_IF_TOOL_EXTENSIONDB instances in the backrepo
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*REQ_IF_TOOL_EXTENSIONDB, 0)
	for _, req_if_tool_extensionDB := range backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB {
		forBackup = append(forBackup, req_if_tool_extensionDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("REQ_IF_TOOL_EXTENSION")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&REQ_IF_TOOL_EXTENSION_Fields, -1)
	for _, req_if_tool_extensionDB := range forBackup {

		var req_if_tool_extensionWOP REQ_IF_TOOL_EXTENSIONWOP
		req_if_tool_extensionDB.CopyBasicFieldsToREQ_IF_TOOL_EXTENSIONWOP(&req_if_tool_extensionWOP)

		row := sh.AddRow()
		row.WriteStruct(&req_if_tool_extensionWOP, -1)
	}
}

// RestoreXL from the "REQ_IF_TOOL_EXTENSION" sheet all REQ_IF_TOOL_EXTENSIONDB instances
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoREQ_IF_TOOL_EXTENSIONid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["REQ_IF_TOOL_EXTENSION"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoREQ_IF_TOOL_EXTENSION.rowVisitorREQ_IF_TOOL_EXTENSION)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) rowVisitorREQ_IF_TOOL_EXTENSION(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var req_if_tool_extensionWOP REQ_IF_TOOL_EXTENSIONWOP
		row.ReadStruct(&req_if_tool_extensionWOP)

		// add the unmarshalled struct to the stage
		req_if_tool_extensionDB := new(REQ_IF_TOOL_EXTENSIONDB)
		req_if_tool_extensionDB.CopyBasicFieldsFromREQ_IF_TOOL_EXTENSIONWOP(&req_if_tool_extensionWOP)

		req_if_tool_extensionDB_ID_atBackupTime := req_if_tool_extensionDB.ID
		req_if_tool_extensionDB.ID = 0
		query := backRepoREQ_IF_TOOL_EXTENSION.db.Create(req_if_tool_extensionDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB[req_if_tool_extensionDB.ID] = req_if_tool_extensionDB
		BackRepoREQ_IF_TOOL_EXTENSIONid_atBckpTime_newID[req_if_tool_extensionDB_ID_atBackupTime] = req_if_tool_extensionDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "REQ_IF_TOOL_EXTENSIONDB.json" in dirPath that stores an array
// of REQ_IF_TOOL_EXTENSIONDB and stores it in the database
// the map BackRepoREQ_IF_TOOL_EXTENSIONid_atBckpTime_newID is updated accordingly
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoREQ_IF_TOOL_EXTENSIONid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "REQ_IF_TOOL_EXTENSIONDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json REQ_IF_TOOL_EXTENSION file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*REQ_IF_TOOL_EXTENSIONDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB
	for _, req_if_tool_extensionDB := range forRestore {

		req_if_tool_extensionDB_ID_atBackupTime := req_if_tool_extensionDB.ID
		req_if_tool_extensionDB.ID = 0
		query := backRepoREQ_IF_TOOL_EXTENSION.db.Create(req_if_tool_extensionDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB[req_if_tool_extensionDB.ID] = req_if_tool_extensionDB
		BackRepoREQ_IF_TOOL_EXTENSIONid_atBckpTime_newID[req_if_tool_extensionDB_ID_atBackupTime] = req_if_tool_extensionDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json REQ_IF_TOOL_EXTENSION file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<REQ_IF_TOOL_EXTENSION>id_atBckpTime_newID
// to compute new index
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) RestorePhaseTwo() {

	for _, req_if_tool_extensionDB := range backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB {

		// next line of code is to avert unused variable compilation error
		_ = req_if_tool_extensionDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoREQ_IF_TOOL_EXTENSION.db.Model(req_if_tool_extensionDB).Updates(*req_if_tool_extensionDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoREQ_IF_TOOL_EXTENSION.ResetReversePointers commits all staged instances of REQ_IF_TOOL_EXTENSION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, req_if_tool_extension := range backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr {
		backRepoREQ_IF_TOOL_EXTENSION.ResetReversePointersInstance(backRepo, idx, req_if_tool_extension)
	}

	return
}

func (backRepoREQ_IF_TOOL_EXTENSION *BackRepoREQ_IF_TOOL_EXTENSIONStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION) (Error error) {

	// fetch matching req_if_tool_extensionDB
	if req_if_tool_extensionDB, ok := backRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB[idx]; ok {
		_ = req_if_tool_extensionDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoREQ_IF_TOOL_EXTENSIONid_atBckpTime_newID map[uint]uint