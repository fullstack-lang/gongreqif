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
var dummy_StaticWebSiteChapter_sql sql.NullBool
var dummy_StaticWebSiteChapter_time time.Duration
var dummy_StaticWebSiteChapter_sort sort.Float64Slice

// StaticWebSiteChapterAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model staticwebsitechapterAPI
type StaticWebSiteChapterAPI struct {
	gorm.Model

	models.StaticWebSiteChapter_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	StaticWebSiteChapterPointersEncoding StaticWebSiteChapterPointersEncoding
}

// StaticWebSiteChapterPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type StaticWebSiteChapterPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Paragraphs is a slice of pointers to another Struct (optional or 0..1)
	Paragraphs IntSlice `gorm:"type:TEXT"`
}

// StaticWebSiteChapterDB describes a staticwebsitechapter in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model staticwebsitechapterDB
type StaticWebSiteChapterDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field staticwebsitechapterDB.Name
	Name_Data sql.NullString

	// Declation for basic field staticwebsitechapterDB.MarkdownContent
	MarkdownContent_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	StaticWebSiteChapterPointersEncoding
}

// StaticWebSiteChapterDBs arrays staticwebsitechapterDBs
// swagger:response staticwebsitechapterDBsResponse
type StaticWebSiteChapterDBs []StaticWebSiteChapterDB

// StaticWebSiteChapterDBResponse provides response
// swagger:response staticwebsitechapterDBResponse
type StaticWebSiteChapterDBResponse struct {
	StaticWebSiteChapterDB
}

// StaticWebSiteChapterWOP is a StaticWebSiteChapter without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type StaticWebSiteChapterWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	MarkdownContent string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var StaticWebSiteChapter_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"MarkdownContent",
}

type BackRepoStaticWebSiteChapterStruct struct {
	// stores StaticWebSiteChapterDB according to their gorm ID
	Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB map[uint]*StaticWebSiteChapterDB

	// stores StaticWebSiteChapterDB ID according to StaticWebSiteChapter address
	Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID map[*models.StaticWebSiteChapter]uint

	// stores StaticWebSiteChapter according to their gorm ID
	Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr map[uint]*models.StaticWebSiteChapter

	db db.DBInterface

	stage *models.Stage
}

func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) GetStage() (stage *models.Stage) {
	stage = backRepoStaticWebSiteChapter.stage
	return
}

func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) GetDB() db.DBInterface {
	return backRepoStaticWebSiteChapter.db
}

// GetStaticWebSiteChapterDBFromStaticWebSiteChapterPtr is a handy function to access the back repo instance from the stage instance
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) GetStaticWebSiteChapterDBFromStaticWebSiteChapterPtr(staticwebsitechapter *models.StaticWebSiteChapter) (staticwebsitechapterDB *StaticWebSiteChapterDB) {
	id := backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID[staticwebsitechapter]
	staticwebsitechapterDB = backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB[id]
	return
}

// BackRepoStaticWebSiteChapter.CommitPhaseOne commits all staged instances of StaticWebSiteChapter to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var staticwebsitechapters []*models.StaticWebSiteChapter
	for staticwebsitechapter := range stage.StaticWebSiteChapters {
		staticwebsitechapters = append(staticwebsitechapters, staticwebsitechapter)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(staticwebsitechapters, func(i, j int) bool {
		return stage.StaticWebSiteChapterMap_Staged_Order[staticwebsitechapters[i]] < stage.StaticWebSiteChapterMap_Staged_Order[staticwebsitechapters[j]]
	})

	for _, staticwebsitechapter := range staticwebsitechapters {
		backRepoStaticWebSiteChapter.CommitPhaseOneInstance(staticwebsitechapter)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, staticwebsitechapter := range backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr {
		if _, ok := stage.StaticWebSiteChapters[staticwebsitechapter]; !ok {
			backRepoStaticWebSiteChapter.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoStaticWebSiteChapter.CommitDeleteInstance commits deletion of StaticWebSiteChapter to the BackRepo
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) CommitDeleteInstance(id uint) (Error error) {

	staticwebsitechapter := backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr[id]

	// staticwebsitechapter is not staged anymore, remove staticwebsitechapterDB
	staticwebsitechapterDB := backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB[id]
	db, _ := backRepoStaticWebSiteChapter.db.Unscoped()
	_, err := db.Delete(staticwebsitechapterDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID, staticwebsitechapter)
	delete(backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr, id)
	delete(backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB, id)

	return
}

// BackRepoStaticWebSiteChapter.CommitPhaseOneInstance commits staticwebsitechapter staged instances of StaticWebSiteChapter to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) CommitPhaseOneInstance(staticwebsitechapter *models.StaticWebSiteChapter) (Error error) {

	// check if the staticwebsitechapter is not commited yet
	if _, ok := backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID[staticwebsitechapter]; ok {
		return
	}

	// initiate staticwebsitechapter
	var staticwebsitechapterDB StaticWebSiteChapterDB
	staticwebsitechapterDB.CopyBasicFieldsFromStaticWebSiteChapter(staticwebsitechapter)

	_, err := backRepoStaticWebSiteChapter.db.Create(&staticwebsitechapterDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID[staticwebsitechapter] = staticwebsitechapterDB.ID
	backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr[staticwebsitechapterDB.ID] = staticwebsitechapter
	backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB[staticwebsitechapterDB.ID] = &staticwebsitechapterDB

	return
}

// BackRepoStaticWebSiteChapter.CommitPhaseTwo commits all staged instances of StaticWebSiteChapter to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, staticwebsitechapter := range backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr {
		backRepoStaticWebSiteChapter.CommitPhaseTwoInstance(backRepo, idx, staticwebsitechapter)
	}

	return
}

// BackRepoStaticWebSiteChapter.CommitPhaseTwoInstance commits {{structname }} of models.StaticWebSiteChapter to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, staticwebsitechapter *models.StaticWebSiteChapter) (Error error) {

	// fetch matching staticwebsitechapterDB
	if staticwebsitechapterDB, ok := backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB[idx]; ok {

		staticwebsitechapterDB.CopyBasicFieldsFromStaticWebSiteChapter(staticwebsitechapter)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		staticwebsitechapterDB.StaticWebSiteChapterPointersEncoding.Paragraphs = make([]int, 0)
		// 2. encode
		for _, staticwebsiteparagraphAssocEnd := range staticwebsitechapter.Paragraphs {
			staticwebsiteparagraphAssocEnd_DB :=
				backRepo.BackRepoStaticWebSiteParagraph.GetStaticWebSiteParagraphDBFromStaticWebSiteParagraphPtr(staticwebsiteparagraphAssocEnd)
			
			// the stage might be inconsistant, meaning that the staticwebsiteparagraphAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if staticwebsiteparagraphAssocEnd_DB == nil {
				continue
			}
			
			staticwebsitechapterDB.StaticWebSiteChapterPointersEncoding.Paragraphs =
				append(staticwebsitechapterDB.StaticWebSiteChapterPointersEncoding.Paragraphs, int(staticwebsiteparagraphAssocEnd_DB.ID))
		}

		_, err := backRepoStaticWebSiteChapter.db.Save(staticwebsitechapterDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown StaticWebSiteChapter intance %s", staticwebsitechapter.Name))
		return err
	}

	return
}

// BackRepoStaticWebSiteChapter.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) CheckoutPhaseOne() (Error error) {

	staticwebsitechapterDBArray := make([]StaticWebSiteChapterDB, 0)
	_, err := backRepoStaticWebSiteChapter.db.Find(&staticwebsitechapterDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	staticwebsitechapterInstancesToBeRemovedFromTheStage := make(map[*models.StaticWebSiteChapter]any)
	for key, value := range backRepoStaticWebSiteChapter.stage.StaticWebSiteChapters {
		staticwebsitechapterInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, staticwebsitechapterDB := range staticwebsitechapterDBArray {
		backRepoStaticWebSiteChapter.CheckoutPhaseOneInstance(&staticwebsitechapterDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		staticwebsitechapter, ok := backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr[staticwebsitechapterDB.ID]
		if ok {
			delete(staticwebsitechapterInstancesToBeRemovedFromTheStage, staticwebsitechapter)
		}
	}

	// remove from stage and back repo's 3 maps all staticwebsitechapters that are not in the checkout
	for staticwebsitechapter := range staticwebsitechapterInstancesToBeRemovedFromTheStage {
		staticwebsitechapter.Unstage(backRepoStaticWebSiteChapter.GetStage())

		// remove instance from the back repo 3 maps
		staticwebsitechapterID := backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID[staticwebsitechapter]
		delete(backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID, staticwebsitechapter)
		delete(backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB, staticwebsitechapterID)
		delete(backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr, staticwebsitechapterID)
	}

	return
}

// CheckoutPhaseOneInstance takes a staticwebsitechapterDB that has been found in the DB, updates the backRepo and stages the
// models version of the staticwebsitechapterDB
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) CheckoutPhaseOneInstance(staticwebsitechapterDB *StaticWebSiteChapterDB) (Error error) {

	staticwebsitechapter, ok := backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr[staticwebsitechapterDB.ID]
	if !ok {
		staticwebsitechapter = new(models.StaticWebSiteChapter)

		backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr[staticwebsitechapterDB.ID] = staticwebsitechapter
		backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID[staticwebsitechapter] = staticwebsitechapterDB.ID

		// append model store with the new element
		staticwebsitechapter.Name = staticwebsitechapterDB.Name_Data.String
		staticwebsitechapter.Stage(backRepoStaticWebSiteChapter.GetStage())
	}
	staticwebsitechapterDB.CopyBasicFieldsToStaticWebSiteChapter(staticwebsitechapter)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	staticwebsitechapter.Stage(backRepoStaticWebSiteChapter.GetStage())

	// preserve pointer to staticwebsitechapterDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB)[staticwebsitechapterDB hold variable pointers
	staticwebsitechapterDB_Data := *staticwebsitechapterDB
	preservedPtrToStaticWebSiteChapter := &staticwebsitechapterDB_Data
	backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB[staticwebsitechapterDB.ID] = preservedPtrToStaticWebSiteChapter

	return
}

// BackRepoStaticWebSiteChapter.CheckoutPhaseTwo Checkouts all staged instances of StaticWebSiteChapter to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, staticwebsitechapterDB := range backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB {
		backRepoStaticWebSiteChapter.CheckoutPhaseTwoInstance(backRepo, staticwebsitechapterDB)
	}
	return
}

// BackRepoStaticWebSiteChapter.CheckoutPhaseTwoInstance Checkouts staged instances of StaticWebSiteChapter to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, staticwebsitechapterDB *StaticWebSiteChapterDB) (Error error) {

	staticwebsitechapter := backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr[staticwebsitechapterDB.ID]

	staticwebsitechapterDB.DecodePointers(backRepo, staticwebsitechapter)

	return
}

func (staticwebsitechapterDB *StaticWebSiteChapterDB) DecodePointers(backRepo *BackRepoStruct, staticwebsitechapter *models.StaticWebSiteChapter) {

	// insertion point for checkout of pointer encoding
	// This loop redeem staticwebsitechapter.Paragraphs in the stage from the encode in the back repo
	// It parses all StaticWebSiteParagraphDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	staticwebsitechapter.Paragraphs = staticwebsitechapter.Paragraphs[:0]
	for _, _StaticWebSiteParagraphid := range staticwebsitechapterDB.StaticWebSiteChapterPointersEncoding.Paragraphs {
		staticwebsitechapter.Paragraphs = append(staticwebsitechapter.Paragraphs, backRepo.BackRepoStaticWebSiteParagraph.Map_StaticWebSiteParagraphDBID_StaticWebSiteParagraphPtr[uint(_StaticWebSiteParagraphid)])
	}

	return
}

// CommitStaticWebSiteChapter allows commit of a single staticwebsitechapter (if already staged)
func (backRepo *BackRepoStruct) CommitStaticWebSiteChapter(staticwebsitechapter *models.StaticWebSiteChapter) {
	backRepo.BackRepoStaticWebSiteChapter.CommitPhaseOneInstance(staticwebsitechapter)
	if id, ok := backRepo.BackRepoStaticWebSiteChapter.Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID[staticwebsitechapter]; ok {
		backRepo.BackRepoStaticWebSiteChapter.CommitPhaseTwoInstance(backRepo, id, staticwebsitechapter)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitStaticWebSiteChapter allows checkout of a single staticwebsitechapter (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutStaticWebSiteChapter(staticwebsitechapter *models.StaticWebSiteChapter) {
	// check if the staticwebsitechapter is staged
	if _, ok := backRepo.BackRepoStaticWebSiteChapter.Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID[staticwebsitechapter]; ok {

		if id, ok := backRepo.BackRepoStaticWebSiteChapter.Map_StaticWebSiteChapterPtr_StaticWebSiteChapterDBID[staticwebsitechapter]; ok {
			var staticwebsitechapterDB StaticWebSiteChapterDB
			staticwebsitechapterDB.ID = id

			if _, err := backRepo.BackRepoStaticWebSiteChapter.db.First(&staticwebsitechapterDB, id); err != nil {
				log.Fatalln("CheckoutStaticWebSiteChapter : Problem with getting object with id:", id)
			}
			backRepo.BackRepoStaticWebSiteChapter.CheckoutPhaseOneInstance(&staticwebsitechapterDB)
			backRepo.BackRepoStaticWebSiteChapter.CheckoutPhaseTwoInstance(backRepo, &staticwebsitechapterDB)
		}
	}
}

// CopyBasicFieldsFromStaticWebSiteChapter
func (staticwebsitechapterDB *StaticWebSiteChapterDB) CopyBasicFieldsFromStaticWebSiteChapter(staticwebsitechapter *models.StaticWebSiteChapter) {
	// insertion point for fields commit

	staticwebsitechapterDB.Name_Data.String = staticwebsitechapter.Name
	staticwebsitechapterDB.Name_Data.Valid = true

	staticwebsitechapterDB.MarkdownContent_Data.String = staticwebsitechapter.MarkdownContent
	staticwebsitechapterDB.MarkdownContent_Data.Valid = true
}

// CopyBasicFieldsFromStaticWebSiteChapter_WOP
func (staticwebsitechapterDB *StaticWebSiteChapterDB) CopyBasicFieldsFromStaticWebSiteChapter_WOP(staticwebsitechapter *models.StaticWebSiteChapter_WOP) {
	// insertion point for fields commit

	staticwebsitechapterDB.Name_Data.String = staticwebsitechapter.Name
	staticwebsitechapterDB.Name_Data.Valid = true

	staticwebsitechapterDB.MarkdownContent_Data.String = staticwebsitechapter.MarkdownContent
	staticwebsitechapterDB.MarkdownContent_Data.Valid = true
}

// CopyBasicFieldsFromStaticWebSiteChapterWOP
func (staticwebsitechapterDB *StaticWebSiteChapterDB) CopyBasicFieldsFromStaticWebSiteChapterWOP(staticwebsitechapter *StaticWebSiteChapterWOP) {
	// insertion point for fields commit

	staticwebsitechapterDB.Name_Data.String = staticwebsitechapter.Name
	staticwebsitechapterDB.Name_Data.Valid = true

	staticwebsitechapterDB.MarkdownContent_Data.String = staticwebsitechapter.MarkdownContent
	staticwebsitechapterDB.MarkdownContent_Data.Valid = true
}

// CopyBasicFieldsToStaticWebSiteChapter
func (staticwebsitechapterDB *StaticWebSiteChapterDB) CopyBasicFieldsToStaticWebSiteChapter(staticwebsitechapter *models.StaticWebSiteChapter) {
	// insertion point for checkout of basic fields (back repo to stage)
	staticwebsitechapter.Name = staticwebsitechapterDB.Name_Data.String
	staticwebsitechapter.MarkdownContent = staticwebsitechapterDB.MarkdownContent_Data.String
}

// CopyBasicFieldsToStaticWebSiteChapter_WOP
func (staticwebsitechapterDB *StaticWebSiteChapterDB) CopyBasicFieldsToStaticWebSiteChapter_WOP(staticwebsitechapter *models.StaticWebSiteChapter_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	staticwebsitechapter.Name = staticwebsitechapterDB.Name_Data.String
	staticwebsitechapter.MarkdownContent = staticwebsitechapterDB.MarkdownContent_Data.String
}

// CopyBasicFieldsToStaticWebSiteChapterWOP
func (staticwebsitechapterDB *StaticWebSiteChapterDB) CopyBasicFieldsToStaticWebSiteChapterWOP(staticwebsitechapter *StaticWebSiteChapterWOP) {
	staticwebsitechapter.ID = int(staticwebsitechapterDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	staticwebsitechapter.Name = staticwebsitechapterDB.Name_Data.String
	staticwebsitechapter.MarkdownContent = staticwebsitechapterDB.MarkdownContent_Data.String
}

// Backup generates a json file from a slice of all StaticWebSiteChapterDB instances in the backrepo
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "StaticWebSiteChapterDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*StaticWebSiteChapterDB, 0)
	for _, staticwebsitechapterDB := range backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB {
		forBackup = append(forBackup, staticwebsitechapterDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json StaticWebSiteChapter ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json StaticWebSiteChapter file", err.Error())
	}
}

// Backup generates a json file from a slice of all StaticWebSiteChapterDB instances in the backrepo
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*StaticWebSiteChapterDB, 0)
	for _, staticwebsitechapterDB := range backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB {
		forBackup = append(forBackup, staticwebsitechapterDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("StaticWebSiteChapter")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&StaticWebSiteChapter_Fields, -1)
	for _, staticwebsitechapterDB := range forBackup {

		var staticwebsitechapterWOP StaticWebSiteChapterWOP
		staticwebsitechapterDB.CopyBasicFieldsToStaticWebSiteChapterWOP(&staticwebsitechapterWOP)

		row := sh.AddRow()
		row.WriteStruct(&staticwebsitechapterWOP, -1)
	}
}

// RestoreXL from the "StaticWebSiteChapter" sheet all StaticWebSiteChapterDB instances
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoStaticWebSiteChapterid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["StaticWebSiteChapter"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoStaticWebSiteChapter.rowVisitorStaticWebSiteChapter)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) rowVisitorStaticWebSiteChapter(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var staticwebsitechapterWOP StaticWebSiteChapterWOP
		row.ReadStruct(&staticwebsitechapterWOP)

		// add the unmarshalled struct to the stage
		staticwebsitechapterDB := new(StaticWebSiteChapterDB)
		staticwebsitechapterDB.CopyBasicFieldsFromStaticWebSiteChapterWOP(&staticwebsitechapterWOP)

		staticwebsitechapterDB_ID_atBackupTime := staticwebsitechapterDB.ID
		staticwebsitechapterDB.ID = 0
		_, err := backRepoStaticWebSiteChapter.db.Create(staticwebsitechapterDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB[staticwebsitechapterDB.ID] = staticwebsitechapterDB
		BackRepoStaticWebSiteChapterid_atBckpTime_newID[staticwebsitechapterDB_ID_atBackupTime] = staticwebsitechapterDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "StaticWebSiteChapterDB.json" in dirPath that stores an array
// of StaticWebSiteChapterDB and stores it in the database
// the map BackRepoStaticWebSiteChapterid_atBckpTime_newID is updated accordingly
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoStaticWebSiteChapterid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "StaticWebSiteChapterDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json StaticWebSiteChapter file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*StaticWebSiteChapterDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB
	for _, staticwebsitechapterDB := range forRestore {

		staticwebsitechapterDB_ID_atBackupTime := staticwebsitechapterDB.ID
		staticwebsitechapterDB.ID = 0
		_, err := backRepoStaticWebSiteChapter.db.Create(staticwebsitechapterDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB[staticwebsitechapterDB.ID] = staticwebsitechapterDB
		BackRepoStaticWebSiteChapterid_atBckpTime_newID[staticwebsitechapterDB_ID_atBackupTime] = staticwebsitechapterDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json StaticWebSiteChapter file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<StaticWebSiteChapter>id_atBckpTime_newID
// to compute new index
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) RestorePhaseTwo() {

	for _, staticwebsitechapterDB := range backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB {

		// next line of code is to avert unused variable compilation error
		_ = staticwebsitechapterDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoStaticWebSiteChapter.db.Model(staticwebsitechapterDB)
		_, err := db.Updates(*staticwebsitechapterDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoStaticWebSiteChapter.ResetReversePointers commits all staged instances of StaticWebSiteChapter to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, staticwebsitechapter := range backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr {
		backRepoStaticWebSiteChapter.ResetReversePointersInstance(backRepo, idx, staticwebsitechapter)
	}

	return
}

func (backRepoStaticWebSiteChapter *BackRepoStaticWebSiteChapterStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, staticwebsitechapter *models.StaticWebSiteChapter) (Error error) {

	// fetch matching staticwebsitechapterDB
	if staticwebsitechapterDB, ok := backRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterDB[idx]; ok {
		_ = staticwebsitechapterDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoStaticWebSiteChapterid_atBckpTime_newID map[uint]uint
