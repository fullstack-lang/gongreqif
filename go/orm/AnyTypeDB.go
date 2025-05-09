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
var dummy_AnyType_sql sql.NullBool
var dummy_AnyType_time time.Duration
var dummy_AnyType_sort sort.Float64Slice

// AnyTypeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model anytypeAPI
type AnyTypeAPI struct {
	gorm.Model

	models.AnyType_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	AnyTypePointersEncoding AnyTypePointersEncoding
}

// AnyTypePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type AnyTypePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// AnyTypeDB describes a anytype in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model anytypeDB
type AnyTypeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field anytypeDB.Name
	Name_Data sql.NullString

	// Declation for basic field anytypeDB.InnerXML
	InnerXML_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	AnyTypePointersEncoding
}

// AnyTypeDBs arrays anytypeDBs
// swagger:response anytypeDBsResponse
type AnyTypeDBs []AnyTypeDB

// AnyTypeDBResponse provides response
// swagger:response anytypeDBResponse
type AnyTypeDBResponse struct {
	AnyTypeDB
}

// AnyTypeWOP is a AnyType without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type AnyTypeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	InnerXML string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var AnyType_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"InnerXML",
}

type BackRepoAnyTypeStruct struct {
	// stores AnyTypeDB according to their gorm ID
	Map_AnyTypeDBID_AnyTypeDB map[uint]*AnyTypeDB

	// stores AnyTypeDB ID according to AnyType address
	Map_AnyTypePtr_AnyTypeDBID map[*models.AnyType]uint

	// stores AnyType according to their gorm ID
	Map_AnyTypeDBID_AnyTypePtr map[uint]*models.AnyType

	db db.DBInterface

	stage *models.Stage
}

func (backRepoAnyType *BackRepoAnyTypeStruct) GetStage() (stage *models.Stage) {
	stage = backRepoAnyType.stage
	return
}

func (backRepoAnyType *BackRepoAnyTypeStruct) GetDB() db.DBInterface {
	return backRepoAnyType.db
}

// GetAnyTypeDBFromAnyTypePtr is a handy function to access the back repo instance from the stage instance
func (backRepoAnyType *BackRepoAnyTypeStruct) GetAnyTypeDBFromAnyTypePtr(anytype *models.AnyType) (anytypeDB *AnyTypeDB) {
	id := backRepoAnyType.Map_AnyTypePtr_AnyTypeDBID[anytype]
	anytypeDB = backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB[id]
	return
}

// BackRepoAnyType.CommitPhaseOne commits all staged instances of AnyType to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoAnyType *BackRepoAnyTypeStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var anytypes []*models.AnyType
	for anytype := range stage.AnyTypes {
		anytypes = append(anytypes, anytype)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(anytypes, func(i, j int) bool {
		return stage.AnyTypeMap_Staged_Order[anytypes[i]] < stage.AnyTypeMap_Staged_Order[anytypes[j]]
	})

	for _, anytype := range anytypes {
		backRepoAnyType.CommitPhaseOneInstance(anytype)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, anytype := range backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr {
		if _, ok := stage.AnyTypes[anytype]; !ok {
			backRepoAnyType.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoAnyType.CommitDeleteInstance commits deletion of AnyType to the BackRepo
func (backRepoAnyType *BackRepoAnyTypeStruct) CommitDeleteInstance(id uint) (Error error) {

	anytype := backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr[id]

	// anytype is not staged anymore, remove anytypeDB
	anytypeDB := backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB[id]
	db, _ := backRepoAnyType.db.Unscoped()
	_, err := db.Delete(anytypeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoAnyType.Map_AnyTypePtr_AnyTypeDBID, anytype)
	delete(backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr, id)
	delete(backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB, id)

	return
}

// BackRepoAnyType.CommitPhaseOneInstance commits anytype staged instances of AnyType to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoAnyType *BackRepoAnyTypeStruct) CommitPhaseOneInstance(anytype *models.AnyType) (Error error) {

	// check if the anytype is not commited yet
	if _, ok := backRepoAnyType.Map_AnyTypePtr_AnyTypeDBID[anytype]; ok {
		return
	}

	// initiate anytype
	var anytypeDB AnyTypeDB
	anytypeDB.CopyBasicFieldsFromAnyType(anytype)

	_, err := backRepoAnyType.db.Create(&anytypeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoAnyType.Map_AnyTypePtr_AnyTypeDBID[anytype] = anytypeDB.ID
	backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr[anytypeDB.ID] = anytype
	backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB[anytypeDB.ID] = &anytypeDB

	return
}

// BackRepoAnyType.CommitPhaseTwo commits all staged instances of AnyType to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAnyType *BackRepoAnyTypeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, anytype := range backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr {
		backRepoAnyType.CommitPhaseTwoInstance(backRepo, idx, anytype)
	}

	return
}

// BackRepoAnyType.CommitPhaseTwoInstance commits {{structname }} of models.AnyType to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAnyType *BackRepoAnyTypeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, anytype *models.AnyType) (Error error) {

	// fetch matching anytypeDB
	if anytypeDB, ok := backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB[idx]; ok {

		anytypeDB.CopyBasicFieldsFromAnyType(anytype)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoAnyType.db.Save(anytypeDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown AnyType intance %s", anytype.Name))
		return err
	}

	return
}

// BackRepoAnyType.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoAnyType *BackRepoAnyTypeStruct) CheckoutPhaseOne() (Error error) {

	anytypeDBArray := make([]AnyTypeDB, 0)
	_, err := backRepoAnyType.db.Find(&anytypeDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	anytypeInstancesToBeRemovedFromTheStage := make(map[*models.AnyType]any)
	for key, value := range backRepoAnyType.stage.AnyTypes {
		anytypeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, anytypeDB := range anytypeDBArray {
		backRepoAnyType.CheckoutPhaseOneInstance(&anytypeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		anytype, ok := backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr[anytypeDB.ID]
		if ok {
			delete(anytypeInstancesToBeRemovedFromTheStage, anytype)
		}
	}

	// remove from stage and back repo's 3 maps all anytypes that are not in the checkout
	for anytype := range anytypeInstancesToBeRemovedFromTheStage {
		anytype.Unstage(backRepoAnyType.GetStage())

		// remove instance from the back repo 3 maps
		anytypeID := backRepoAnyType.Map_AnyTypePtr_AnyTypeDBID[anytype]
		delete(backRepoAnyType.Map_AnyTypePtr_AnyTypeDBID, anytype)
		delete(backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB, anytypeID)
		delete(backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr, anytypeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a anytypeDB that has been found in the DB, updates the backRepo and stages the
// models version of the anytypeDB
func (backRepoAnyType *BackRepoAnyTypeStruct) CheckoutPhaseOneInstance(anytypeDB *AnyTypeDB) (Error error) {

	anytype, ok := backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr[anytypeDB.ID]
	if !ok {
		anytype = new(models.AnyType)

		backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr[anytypeDB.ID] = anytype
		backRepoAnyType.Map_AnyTypePtr_AnyTypeDBID[anytype] = anytypeDB.ID

		// append model store with the new element
		anytype.Name = anytypeDB.Name_Data.String
		anytype.Stage(backRepoAnyType.GetStage())
	}
	anytypeDB.CopyBasicFieldsToAnyType(anytype)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	anytype.Stage(backRepoAnyType.GetStage())

	// preserve pointer to anytypeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_AnyTypeDBID_AnyTypeDB)[anytypeDB hold variable pointers
	anytypeDB_Data := *anytypeDB
	preservedPtrToAnyType := &anytypeDB_Data
	backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB[anytypeDB.ID] = preservedPtrToAnyType

	return
}

// BackRepoAnyType.CheckoutPhaseTwo Checkouts all staged instances of AnyType to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAnyType *BackRepoAnyTypeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, anytypeDB := range backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB {
		backRepoAnyType.CheckoutPhaseTwoInstance(backRepo, anytypeDB)
	}
	return
}

// BackRepoAnyType.CheckoutPhaseTwoInstance Checkouts staged instances of AnyType to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAnyType *BackRepoAnyTypeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, anytypeDB *AnyTypeDB) (Error error) {

	anytype := backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr[anytypeDB.ID]

	anytypeDB.DecodePointers(backRepo, anytype)

	return
}

func (anytypeDB *AnyTypeDB) DecodePointers(backRepo *BackRepoStruct, anytype *models.AnyType) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitAnyType allows commit of a single anytype (if already staged)
func (backRepo *BackRepoStruct) CommitAnyType(anytype *models.AnyType) {
	backRepo.BackRepoAnyType.CommitPhaseOneInstance(anytype)
	if id, ok := backRepo.BackRepoAnyType.Map_AnyTypePtr_AnyTypeDBID[anytype]; ok {
		backRepo.BackRepoAnyType.CommitPhaseTwoInstance(backRepo, id, anytype)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitAnyType allows checkout of a single anytype (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutAnyType(anytype *models.AnyType) {
	// check if the anytype is staged
	if _, ok := backRepo.BackRepoAnyType.Map_AnyTypePtr_AnyTypeDBID[anytype]; ok {

		if id, ok := backRepo.BackRepoAnyType.Map_AnyTypePtr_AnyTypeDBID[anytype]; ok {
			var anytypeDB AnyTypeDB
			anytypeDB.ID = id

			if _, err := backRepo.BackRepoAnyType.db.First(&anytypeDB, id); err != nil {
				log.Fatalln("CheckoutAnyType : Problem with getting object with id:", id)
			}
			backRepo.BackRepoAnyType.CheckoutPhaseOneInstance(&anytypeDB)
			backRepo.BackRepoAnyType.CheckoutPhaseTwoInstance(backRepo, &anytypeDB)
		}
	}
}

// CopyBasicFieldsFromAnyType
func (anytypeDB *AnyTypeDB) CopyBasicFieldsFromAnyType(anytype *models.AnyType) {
	// insertion point for fields commit

	anytypeDB.Name_Data.String = anytype.Name
	anytypeDB.Name_Data.Valid = true

	anytypeDB.InnerXML_Data.String = anytype.InnerXML
	anytypeDB.InnerXML_Data.Valid = true
}

// CopyBasicFieldsFromAnyType_WOP
func (anytypeDB *AnyTypeDB) CopyBasicFieldsFromAnyType_WOP(anytype *models.AnyType_WOP) {
	// insertion point for fields commit

	anytypeDB.Name_Data.String = anytype.Name
	anytypeDB.Name_Data.Valid = true

	anytypeDB.InnerXML_Data.String = anytype.InnerXML
	anytypeDB.InnerXML_Data.Valid = true
}

// CopyBasicFieldsFromAnyTypeWOP
func (anytypeDB *AnyTypeDB) CopyBasicFieldsFromAnyTypeWOP(anytype *AnyTypeWOP) {
	// insertion point for fields commit

	anytypeDB.Name_Data.String = anytype.Name
	anytypeDB.Name_Data.Valid = true

	anytypeDB.InnerXML_Data.String = anytype.InnerXML
	anytypeDB.InnerXML_Data.Valid = true
}

// CopyBasicFieldsToAnyType
func (anytypeDB *AnyTypeDB) CopyBasicFieldsToAnyType(anytype *models.AnyType) {
	// insertion point for checkout of basic fields (back repo to stage)
	anytype.Name = anytypeDB.Name_Data.String
	anytype.InnerXML = anytypeDB.InnerXML_Data.String
}

// CopyBasicFieldsToAnyType_WOP
func (anytypeDB *AnyTypeDB) CopyBasicFieldsToAnyType_WOP(anytype *models.AnyType_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	anytype.Name = anytypeDB.Name_Data.String
	anytype.InnerXML = anytypeDB.InnerXML_Data.String
}

// CopyBasicFieldsToAnyTypeWOP
func (anytypeDB *AnyTypeDB) CopyBasicFieldsToAnyTypeWOP(anytype *AnyTypeWOP) {
	anytype.ID = int(anytypeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	anytype.Name = anytypeDB.Name_Data.String
	anytype.InnerXML = anytypeDB.InnerXML_Data.String
}

// Backup generates a json file from a slice of all AnyTypeDB instances in the backrepo
func (backRepoAnyType *BackRepoAnyTypeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "AnyTypeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AnyTypeDB, 0)
	for _, anytypeDB := range backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB {
		forBackup = append(forBackup, anytypeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json AnyType ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json AnyType file", err.Error())
	}
}

// Backup generates a json file from a slice of all AnyTypeDB instances in the backrepo
func (backRepoAnyType *BackRepoAnyTypeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AnyTypeDB, 0)
	for _, anytypeDB := range backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB {
		forBackup = append(forBackup, anytypeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("AnyType")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&AnyType_Fields, -1)
	for _, anytypeDB := range forBackup {

		var anytypeWOP AnyTypeWOP
		anytypeDB.CopyBasicFieldsToAnyTypeWOP(&anytypeWOP)

		row := sh.AddRow()
		row.WriteStruct(&anytypeWOP, -1)
	}
}

// RestoreXL from the "AnyType" sheet all AnyTypeDB instances
func (backRepoAnyType *BackRepoAnyTypeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoAnyTypeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["AnyType"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoAnyType.rowVisitorAnyType)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoAnyType *BackRepoAnyTypeStruct) rowVisitorAnyType(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var anytypeWOP AnyTypeWOP
		row.ReadStruct(&anytypeWOP)

		// add the unmarshalled struct to the stage
		anytypeDB := new(AnyTypeDB)
		anytypeDB.CopyBasicFieldsFromAnyTypeWOP(&anytypeWOP)

		anytypeDB_ID_atBackupTime := anytypeDB.ID
		anytypeDB.ID = 0
		_, err := backRepoAnyType.db.Create(anytypeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB[anytypeDB.ID] = anytypeDB
		BackRepoAnyTypeid_atBckpTime_newID[anytypeDB_ID_atBackupTime] = anytypeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "AnyTypeDB.json" in dirPath that stores an array
// of AnyTypeDB and stores it in the database
// the map BackRepoAnyTypeid_atBckpTime_newID is updated accordingly
func (backRepoAnyType *BackRepoAnyTypeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoAnyTypeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "AnyTypeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json AnyType file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*AnyTypeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_AnyTypeDBID_AnyTypeDB
	for _, anytypeDB := range forRestore {

		anytypeDB_ID_atBackupTime := anytypeDB.ID
		anytypeDB.ID = 0
		_, err := backRepoAnyType.db.Create(anytypeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB[anytypeDB.ID] = anytypeDB
		BackRepoAnyTypeid_atBckpTime_newID[anytypeDB_ID_atBackupTime] = anytypeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json AnyType file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<AnyType>id_atBckpTime_newID
// to compute new index
func (backRepoAnyType *BackRepoAnyTypeStruct) RestorePhaseTwo() {

	for _, anytypeDB := range backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB {

		// next line of code is to avert unused variable compilation error
		_ = anytypeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoAnyType.db.Model(anytypeDB)
		_, err := db.Updates(*anytypeDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoAnyType.ResetReversePointers commits all staged instances of AnyType to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAnyType *BackRepoAnyTypeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, anytype := range backRepoAnyType.Map_AnyTypeDBID_AnyTypePtr {
		backRepoAnyType.ResetReversePointersInstance(backRepo, idx, anytype)
	}

	return
}

func (backRepoAnyType *BackRepoAnyTypeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, anytype *models.AnyType) (Error error) {

	// fetch matching anytypeDB
	if anytypeDB, ok := backRepoAnyType.Map_AnyTypeDBID_AnyTypeDB[idx]; ok {
		_ = anytypeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoAnyTypeid_atBckpTime_newID map[uint]uint
