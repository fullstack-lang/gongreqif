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

	"github.com/fullstack-lang/gong/lib/doc2/go/db"
	"github.com/fullstack-lang/gong/lib/doc2/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_GongEnumValueShape_sql sql.NullBool
var dummy_GongEnumValueShape_time time.Duration
var dummy_GongEnumValueShape_sort sort.Float64Slice

// GongEnumValueShapeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongenumvalueshapeAPI
type GongEnumValueShapeAPI struct {
	gorm.Model

	models.GongEnumValueShape_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	GongEnumValueShapePointersEncoding GongEnumValueShapePointersEncoding
}

// GongEnumValueShapePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongEnumValueShapePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// GongEnumValueShapeDB describes a gongenumvalueshape in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongenumvalueshapeDB
type GongEnumValueShapeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongenumvalueshapeDB.Name
	Name_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	GongEnumValueShapePointersEncoding
}

// GongEnumValueShapeDBs arrays gongenumvalueshapeDBs
// swagger:response gongenumvalueshapeDBsResponse
type GongEnumValueShapeDBs []GongEnumValueShapeDB

// GongEnumValueShapeDBResponse provides response
// swagger:response gongenumvalueshapeDBResponse
type GongEnumValueShapeDBResponse struct {
	GongEnumValueShapeDB
}

// GongEnumValueShapeWOP is a GongEnumValueShape without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongEnumValueShapeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IdentifierMeta any `xlsx:"2"`
	// insertion for WOP pointer fields
}

var GongEnumValueShape_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IdentifierMeta",
}

type BackRepoGongEnumValueShapeStruct struct {
	// stores GongEnumValueShapeDB according to their gorm ID
	Map_GongEnumValueShapeDBID_GongEnumValueShapeDB map[uint]*GongEnumValueShapeDB

	// stores GongEnumValueShapeDB ID according to GongEnumValueShape address
	Map_GongEnumValueShapePtr_GongEnumValueShapeDBID map[*models.GongEnumValueShape]uint

	// stores GongEnumValueShape according to their gorm ID
	Map_GongEnumValueShapeDBID_GongEnumValueShapePtr map[uint]*models.GongEnumValueShape

	db db.DBInterface

	stage *models.Stage
}

func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) GetStage() (stage *models.Stage) {
	stage = backRepoGongEnumValueShape.stage
	return
}

func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) GetDB() db.DBInterface {
	return backRepoGongEnumValueShape.db
}

// GetGongEnumValueShapeDBFromGongEnumValueShapePtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) GetGongEnumValueShapeDBFromGongEnumValueShapePtr(gongenumvalueshape *models.GongEnumValueShape) (gongenumvalueshapeDB *GongEnumValueShapeDB) {
	id := backRepoGongEnumValueShape.Map_GongEnumValueShapePtr_GongEnumValueShapeDBID[gongenumvalueshape]
	gongenumvalueshapeDB = backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB[id]
	return
}

// BackRepoGongEnumValueShape.CommitPhaseOne commits all staged instances of GongEnumValueShape to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var gongenumvalueshapes []*models.GongEnumValueShape
	for gongenumvalueshape := range stage.GongEnumValueShapes {
		gongenumvalueshapes = append(gongenumvalueshapes, gongenumvalueshape)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(gongenumvalueshapes, func(i, j int) bool {
		return stage.GongEnumValueShapeMap_Staged_Order[gongenumvalueshapes[i]] < stage.GongEnumValueShapeMap_Staged_Order[gongenumvalueshapes[j]]
	})

	for _, gongenumvalueshape := range gongenumvalueshapes {
		backRepoGongEnumValueShape.CommitPhaseOneInstance(gongenumvalueshape)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongenumvalueshape := range backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr {
		if _, ok := stage.GongEnumValueShapes[gongenumvalueshape]; !ok {
			backRepoGongEnumValueShape.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongEnumValueShape.CommitDeleteInstance commits deletion of GongEnumValueShape to the BackRepo
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) CommitDeleteInstance(id uint) (Error error) {

	gongenumvalueshape := backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr[id]

	// gongenumvalueshape is not staged anymore, remove gongenumvalueshapeDB
	gongenumvalueshapeDB := backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB[id]
	db, _ := backRepoGongEnumValueShape.db.Unscoped()
	_, err := db.Delete(gongenumvalueshapeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoGongEnumValueShape.Map_GongEnumValueShapePtr_GongEnumValueShapeDBID, gongenumvalueshape)
	delete(backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr, id)
	delete(backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB, id)

	return
}

// BackRepoGongEnumValueShape.CommitPhaseOneInstance commits gongenumvalueshape staged instances of GongEnumValueShape to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) CommitPhaseOneInstance(gongenumvalueshape *models.GongEnumValueShape) (Error error) {

	// check if the gongenumvalueshape is not commited yet
	if _, ok := backRepoGongEnumValueShape.Map_GongEnumValueShapePtr_GongEnumValueShapeDBID[gongenumvalueshape]; ok {
		return
	}

	// initiate gongenumvalueshape
	var gongenumvalueshapeDB GongEnumValueShapeDB
	gongenumvalueshapeDB.CopyBasicFieldsFromGongEnumValueShape(gongenumvalueshape)

	_, err := backRepoGongEnumValueShape.db.Create(&gongenumvalueshapeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoGongEnumValueShape.Map_GongEnumValueShapePtr_GongEnumValueShapeDBID[gongenumvalueshape] = gongenumvalueshapeDB.ID
	backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr[gongenumvalueshapeDB.ID] = gongenumvalueshape
	backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB[gongenumvalueshapeDB.ID] = &gongenumvalueshapeDB

	return
}

// BackRepoGongEnumValueShape.CommitPhaseTwo commits all staged instances of GongEnumValueShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongenumvalueshape := range backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr {
		backRepoGongEnumValueShape.CommitPhaseTwoInstance(backRepo, idx, gongenumvalueshape)
	}

	return
}

// BackRepoGongEnumValueShape.CommitPhaseTwoInstance commits {{structname }} of models.GongEnumValueShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongenumvalueshape *models.GongEnumValueShape) (Error error) {

	// fetch matching gongenumvalueshapeDB
	if gongenumvalueshapeDB, ok := backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB[idx]; ok {

		gongenumvalueshapeDB.CopyBasicFieldsFromGongEnumValueShape(gongenumvalueshape)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoGongEnumValueShape.db.Save(gongenumvalueshapeDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongEnumValueShape intance %s", gongenumvalueshape.Name))
		return err
	}

	return
}

// BackRepoGongEnumValueShape.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) CheckoutPhaseOne() (Error error) {

	gongenumvalueshapeDBArray := make([]GongEnumValueShapeDB, 0)
	_, err := backRepoGongEnumValueShape.db.Find(&gongenumvalueshapeDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongenumvalueshapeInstancesToBeRemovedFromTheStage := make(map[*models.GongEnumValueShape]any)
	for key, value := range backRepoGongEnumValueShape.stage.GongEnumValueShapes {
		gongenumvalueshapeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongenumvalueshapeDB := range gongenumvalueshapeDBArray {
		backRepoGongEnumValueShape.CheckoutPhaseOneInstance(&gongenumvalueshapeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongenumvalueshape, ok := backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr[gongenumvalueshapeDB.ID]
		if ok {
			delete(gongenumvalueshapeInstancesToBeRemovedFromTheStage, gongenumvalueshape)
		}
	}

	// remove from stage and back repo's 3 maps all gongenumvalueshapes that are not in the checkout
	for gongenumvalueshape := range gongenumvalueshapeInstancesToBeRemovedFromTheStage {
		gongenumvalueshape.Unstage(backRepoGongEnumValueShape.GetStage())

		// remove instance from the back repo 3 maps
		gongenumvalueshapeID := backRepoGongEnumValueShape.Map_GongEnumValueShapePtr_GongEnumValueShapeDBID[gongenumvalueshape]
		delete(backRepoGongEnumValueShape.Map_GongEnumValueShapePtr_GongEnumValueShapeDBID, gongenumvalueshape)
		delete(backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB, gongenumvalueshapeID)
		delete(backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr, gongenumvalueshapeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongenumvalueshapeDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongenumvalueshapeDB
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) CheckoutPhaseOneInstance(gongenumvalueshapeDB *GongEnumValueShapeDB) (Error error) {

	gongenumvalueshape, ok := backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr[gongenumvalueshapeDB.ID]
	if !ok {
		gongenumvalueshape = new(models.GongEnumValueShape)

		backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr[gongenumvalueshapeDB.ID] = gongenumvalueshape
		backRepoGongEnumValueShape.Map_GongEnumValueShapePtr_GongEnumValueShapeDBID[gongenumvalueshape] = gongenumvalueshapeDB.ID

		// append model store with the new element
		gongenumvalueshape.Name = gongenumvalueshapeDB.Name_Data.String
		gongenumvalueshape.Stage(backRepoGongEnumValueShape.GetStage())
	}
	gongenumvalueshapeDB.CopyBasicFieldsToGongEnumValueShape(gongenumvalueshape)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	gongenumvalueshape.Stage(backRepoGongEnumValueShape.GetStage())

	// preserve pointer to gongenumvalueshapeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongEnumValueShapeDBID_GongEnumValueShapeDB)[gongenumvalueshapeDB hold variable pointers
	gongenumvalueshapeDB_Data := *gongenumvalueshapeDB
	preservedPtrToGongEnumValueShape := &gongenumvalueshapeDB_Data
	backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB[gongenumvalueshapeDB.ID] = preservedPtrToGongEnumValueShape

	return
}

// BackRepoGongEnumValueShape.CheckoutPhaseTwo Checkouts all staged instances of GongEnumValueShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongenumvalueshapeDB := range backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB {
		backRepoGongEnumValueShape.CheckoutPhaseTwoInstance(backRepo, gongenumvalueshapeDB)
	}
	return
}

// BackRepoGongEnumValueShape.CheckoutPhaseTwoInstance Checkouts staged instances of GongEnumValueShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongenumvalueshapeDB *GongEnumValueShapeDB) (Error error) {

	gongenumvalueshape := backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr[gongenumvalueshapeDB.ID]

	gongenumvalueshapeDB.DecodePointers(backRepo, gongenumvalueshape)

	return
}

func (gongenumvalueshapeDB *GongEnumValueShapeDB) DecodePointers(backRepo *BackRepoStruct, gongenumvalueshape *models.GongEnumValueShape) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitGongEnumValueShape allows commit of a single gongenumvalueshape (if already staged)
func (backRepo *BackRepoStruct) CommitGongEnumValueShape(gongenumvalueshape *models.GongEnumValueShape) {
	backRepo.BackRepoGongEnumValueShape.CommitPhaseOneInstance(gongenumvalueshape)
	if id, ok := backRepo.BackRepoGongEnumValueShape.Map_GongEnumValueShapePtr_GongEnumValueShapeDBID[gongenumvalueshape]; ok {
		backRepo.BackRepoGongEnumValueShape.CommitPhaseTwoInstance(backRepo, id, gongenumvalueshape)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitGongEnumValueShape allows checkout of a single gongenumvalueshape (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongEnumValueShape(gongenumvalueshape *models.GongEnumValueShape) {
	// check if the gongenumvalueshape is staged
	if _, ok := backRepo.BackRepoGongEnumValueShape.Map_GongEnumValueShapePtr_GongEnumValueShapeDBID[gongenumvalueshape]; ok {

		if id, ok := backRepo.BackRepoGongEnumValueShape.Map_GongEnumValueShapePtr_GongEnumValueShapeDBID[gongenumvalueshape]; ok {
			var gongenumvalueshapeDB GongEnumValueShapeDB
			gongenumvalueshapeDB.ID = id

			if _, err := backRepo.BackRepoGongEnumValueShape.db.First(&gongenumvalueshapeDB, id); err != nil {
				log.Fatalln("CheckoutGongEnumValueShape : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongEnumValueShape.CheckoutPhaseOneInstance(&gongenumvalueshapeDB)
			backRepo.BackRepoGongEnumValueShape.CheckoutPhaseTwoInstance(backRepo, &gongenumvalueshapeDB)
		}
	}
}

// CopyBasicFieldsFromGongEnumValueShape
func (gongenumvalueshapeDB *GongEnumValueShapeDB) CopyBasicFieldsFromGongEnumValueShape(gongenumvalueshape *models.GongEnumValueShape) {
	// insertion point for fields commit

	gongenumvalueshapeDB.Name_Data.String = gongenumvalueshape.Name
	gongenumvalueshapeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromGongEnumValueShape_WOP
func (gongenumvalueshapeDB *GongEnumValueShapeDB) CopyBasicFieldsFromGongEnumValueShape_WOP(gongenumvalueshape *models.GongEnumValueShape_WOP) {
	// insertion point for fields commit

	gongenumvalueshapeDB.Name_Data.String = gongenumvalueshape.Name
	gongenumvalueshapeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromGongEnumValueShapeWOP
func (gongenumvalueshapeDB *GongEnumValueShapeDB) CopyBasicFieldsFromGongEnumValueShapeWOP(gongenumvalueshape *GongEnumValueShapeWOP) {
	// insertion point for fields commit

	gongenumvalueshapeDB.Name_Data.String = gongenumvalueshape.Name
	gongenumvalueshapeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToGongEnumValueShape
func (gongenumvalueshapeDB *GongEnumValueShapeDB) CopyBasicFieldsToGongEnumValueShape(gongenumvalueshape *models.GongEnumValueShape) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongenumvalueshape.Name = gongenumvalueshapeDB.Name_Data.String
}

// CopyBasicFieldsToGongEnumValueShape_WOP
func (gongenumvalueshapeDB *GongEnumValueShapeDB) CopyBasicFieldsToGongEnumValueShape_WOP(gongenumvalueshape *models.GongEnumValueShape_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongenumvalueshape.Name = gongenumvalueshapeDB.Name_Data.String
}

// CopyBasicFieldsToGongEnumValueShapeWOP
func (gongenumvalueshapeDB *GongEnumValueShapeDB) CopyBasicFieldsToGongEnumValueShapeWOP(gongenumvalueshape *GongEnumValueShapeWOP) {
	gongenumvalueshape.ID = int(gongenumvalueshapeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongenumvalueshape.Name = gongenumvalueshapeDB.Name_Data.String
}

// Backup generates a json file from a slice of all GongEnumValueShapeDB instances in the backrepo
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongEnumValueShapeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongEnumValueShapeDB, 0)
	for _, gongenumvalueshapeDB := range backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB {
		forBackup = append(forBackup, gongenumvalueshapeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json GongEnumValueShape ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json GongEnumValueShape file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongEnumValueShapeDB instances in the backrepo
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongEnumValueShapeDB, 0)
	for _, gongenumvalueshapeDB := range backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB {
		forBackup = append(forBackup, gongenumvalueshapeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongEnumValueShape")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongEnumValueShape_Fields, -1)
	for _, gongenumvalueshapeDB := range forBackup {

		var gongenumvalueshapeWOP GongEnumValueShapeWOP
		gongenumvalueshapeDB.CopyBasicFieldsToGongEnumValueShapeWOP(&gongenumvalueshapeWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongenumvalueshapeWOP, -1)
	}
}

// RestoreXL from the "GongEnumValueShape" sheet all GongEnumValueShapeDB instances
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongEnumValueShapeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongEnumValueShape"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongEnumValueShape.rowVisitorGongEnumValueShape)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) rowVisitorGongEnumValueShape(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongenumvalueshapeWOP GongEnumValueShapeWOP
		row.ReadStruct(&gongenumvalueshapeWOP)

		// add the unmarshalled struct to the stage
		gongenumvalueshapeDB := new(GongEnumValueShapeDB)
		gongenumvalueshapeDB.CopyBasicFieldsFromGongEnumValueShapeWOP(&gongenumvalueshapeWOP)

		gongenumvalueshapeDB_ID_atBackupTime := gongenumvalueshapeDB.ID
		gongenumvalueshapeDB.ID = 0
		_, err := backRepoGongEnumValueShape.db.Create(gongenumvalueshapeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB[gongenumvalueshapeDB.ID] = gongenumvalueshapeDB
		BackRepoGongEnumValueShapeid_atBckpTime_newID[gongenumvalueshapeDB_ID_atBackupTime] = gongenumvalueshapeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongEnumValueShapeDB.json" in dirPath that stores an array
// of GongEnumValueShapeDB and stores it in the database
// the map BackRepoGongEnumValueShapeid_atBckpTime_newID is updated accordingly
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongEnumValueShapeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongEnumValueShapeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json GongEnumValueShape file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongEnumValueShapeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongEnumValueShapeDBID_GongEnumValueShapeDB
	for _, gongenumvalueshapeDB := range forRestore {

		gongenumvalueshapeDB_ID_atBackupTime := gongenumvalueshapeDB.ID
		gongenumvalueshapeDB.ID = 0
		_, err := backRepoGongEnumValueShape.db.Create(gongenumvalueshapeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB[gongenumvalueshapeDB.ID] = gongenumvalueshapeDB
		BackRepoGongEnumValueShapeid_atBckpTime_newID[gongenumvalueshapeDB_ID_atBackupTime] = gongenumvalueshapeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json GongEnumValueShape file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongEnumValueShape>id_atBckpTime_newID
// to compute new index
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) RestorePhaseTwo() {

	for _, gongenumvalueshapeDB := range backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB {

		// next line of code is to avert unused variable compilation error
		_ = gongenumvalueshapeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoGongEnumValueShape.db.Model(gongenumvalueshapeDB)
		_, err := db.Updates(*gongenumvalueshapeDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoGongEnumValueShape.ResetReversePointers commits all staged instances of GongEnumValueShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, gongenumvalueshape := range backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr {
		backRepoGongEnumValueShape.ResetReversePointersInstance(backRepo, idx, gongenumvalueshape)
	}

	return
}

func (backRepoGongEnumValueShape *BackRepoGongEnumValueShapeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, gongenumvalueshape *models.GongEnumValueShape) (Error error) {

	// fetch matching gongenumvalueshapeDB
	if gongenumvalueshapeDB, ok := backRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB[idx]; ok {
		_ = gongenumvalueshapeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongEnumValueShapeid_atBckpTime_newID map[uint]uint
