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
var dummy_Paragraph_sql sql.NullBool
var dummy_Paragraph_time time.Duration
var dummy_Paragraph_sort sort.Float64Slice

// ParagraphAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model paragraphAPI
type ParagraphAPI struct {
	gorm.Model

	models.Paragraph_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ParagraphPointersEncoding ParagraphPointersEncoding
}

// ParagraphPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ParagraphPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Image is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ImageID sql.NullInt64
}

// ParagraphDB describes a paragraph in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model paragraphDB
type ParagraphDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field paragraphDB.Name
	Name_Data sql.NullString

	// Declation for basic field paragraphDB.LegendMarkdownContent
	LegendMarkdownContent_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ParagraphPointersEncoding
}

// ParagraphDBs arrays paragraphDBs
// swagger:response paragraphDBsResponse
type ParagraphDBs []ParagraphDB

// ParagraphDBResponse provides response
// swagger:response paragraphDBResponse
type ParagraphDBResponse struct {
	ParagraphDB
}

// ParagraphWOP is a Paragraph without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ParagraphWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	LegendMarkdownContent string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var Paragraph_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"LegendMarkdownContent",
}

type BackRepoParagraphStruct struct {
	// stores ParagraphDB according to their gorm ID
	Map_ParagraphDBID_ParagraphDB map[uint]*ParagraphDB

	// stores ParagraphDB ID according to Paragraph address
	Map_ParagraphPtr_ParagraphDBID map[*models.Paragraph]uint

	// stores Paragraph according to their gorm ID
	Map_ParagraphDBID_ParagraphPtr map[uint]*models.Paragraph

	db db.DBInterface

	stage *models.Stage
}

func (backRepoParagraph *BackRepoParagraphStruct) GetStage() (stage *models.Stage) {
	stage = backRepoParagraph.stage
	return
}

func (backRepoParagraph *BackRepoParagraphStruct) GetDB() db.DBInterface {
	return backRepoParagraph.db
}

// GetParagraphDBFromParagraphPtr is a handy function to access the back repo instance from the stage instance
func (backRepoParagraph *BackRepoParagraphStruct) GetParagraphDBFromParagraphPtr(paragraph *models.Paragraph) (paragraphDB *ParagraphDB) {
	id := backRepoParagraph.Map_ParagraphPtr_ParagraphDBID[paragraph]
	paragraphDB = backRepoParagraph.Map_ParagraphDBID_ParagraphDB[id]
	return
}

// BackRepoParagraph.CommitPhaseOne commits all staged instances of Paragraph to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoParagraph *BackRepoParagraphStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var paragraphs []*models.Paragraph
	for paragraph := range stage.Paragraphs {
		paragraphs = append(paragraphs, paragraph)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(paragraphs, func(i, j int) bool {
		return stage.ParagraphMap_Staged_Order[paragraphs[i]] < stage.ParagraphMap_Staged_Order[paragraphs[j]]
	})

	for _, paragraph := range paragraphs {
		backRepoParagraph.CommitPhaseOneInstance(paragraph)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, paragraph := range backRepoParagraph.Map_ParagraphDBID_ParagraphPtr {
		if _, ok := stage.Paragraphs[paragraph]; !ok {
			backRepoParagraph.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoParagraph.CommitDeleteInstance commits deletion of Paragraph to the BackRepo
func (backRepoParagraph *BackRepoParagraphStruct) CommitDeleteInstance(id uint) (Error error) {

	paragraph := backRepoParagraph.Map_ParagraphDBID_ParagraphPtr[id]

	// paragraph is not staged anymore, remove paragraphDB
	paragraphDB := backRepoParagraph.Map_ParagraphDBID_ParagraphDB[id]
	db, _ := backRepoParagraph.db.Unscoped()
	_, err := db.Delete(paragraphDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoParagraph.Map_ParagraphPtr_ParagraphDBID, paragraph)
	delete(backRepoParagraph.Map_ParagraphDBID_ParagraphPtr, id)
	delete(backRepoParagraph.Map_ParagraphDBID_ParagraphDB, id)

	return
}

// BackRepoParagraph.CommitPhaseOneInstance commits paragraph staged instances of Paragraph to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoParagraph *BackRepoParagraphStruct) CommitPhaseOneInstance(paragraph *models.Paragraph) (Error error) {

	// check if the paragraph is not commited yet
	if _, ok := backRepoParagraph.Map_ParagraphPtr_ParagraphDBID[paragraph]; ok {
		return
	}

	// initiate paragraph
	var paragraphDB ParagraphDB
	paragraphDB.CopyBasicFieldsFromParagraph(paragraph)

	_, err := backRepoParagraph.db.Create(&paragraphDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoParagraph.Map_ParagraphPtr_ParagraphDBID[paragraph] = paragraphDB.ID
	backRepoParagraph.Map_ParagraphDBID_ParagraphPtr[paragraphDB.ID] = paragraph
	backRepoParagraph.Map_ParagraphDBID_ParagraphDB[paragraphDB.ID] = &paragraphDB

	return
}

// BackRepoParagraph.CommitPhaseTwo commits all staged instances of Paragraph to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoParagraph *BackRepoParagraphStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, paragraph := range backRepoParagraph.Map_ParagraphDBID_ParagraphPtr {
		backRepoParagraph.CommitPhaseTwoInstance(backRepo, idx, paragraph)
	}

	return
}

// BackRepoParagraph.CommitPhaseTwoInstance commits {{structname }} of models.Paragraph to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoParagraph *BackRepoParagraphStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, paragraph *models.Paragraph) (Error error) {

	// fetch matching paragraphDB
	if paragraphDB, ok := backRepoParagraph.Map_ParagraphDBID_ParagraphDB[idx]; ok {

		paragraphDB.CopyBasicFieldsFromParagraph(paragraph)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value paragraph.Image translates to updating the paragraph.ImageID
		paragraphDB.ImageID.Valid = true // allow for a 0 value (nil association)
		if paragraph.Image != nil {
			if ImageId, ok := backRepo.BackRepoImage.Map_ImagePtr_ImageDBID[paragraph.Image]; ok {
				paragraphDB.ImageID.Int64 = int64(ImageId)
				paragraphDB.ImageID.Valid = true
			}
		} else {
			paragraphDB.ImageID.Int64 = 0
			paragraphDB.ImageID.Valid = true
		}

		_, err := backRepoParagraph.db.Save(paragraphDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Paragraph intance %s", paragraph.Name))
		return err
	}

	return
}

// BackRepoParagraph.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoParagraph *BackRepoParagraphStruct) CheckoutPhaseOne() (Error error) {

	paragraphDBArray := make([]ParagraphDB, 0)
	_, err := backRepoParagraph.db.Find(&paragraphDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	paragraphInstancesToBeRemovedFromTheStage := make(map[*models.Paragraph]any)
	for key, value := range backRepoParagraph.stage.Paragraphs {
		paragraphInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, paragraphDB := range paragraphDBArray {
		backRepoParagraph.CheckoutPhaseOneInstance(&paragraphDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		paragraph, ok := backRepoParagraph.Map_ParagraphDBID_ParagraphPtr[paragraphDB.ID]
		if ok {
			delete(paragraphInstancesToBeRemovedFromTheStage, paragraph)
		}
	}

	// remove from stage and back repo's 3 maps all paragraphs that are not in the checkout
	for paragraph := range paragraphInstancesToBeRemovedFromTheStage {
		paragraph.Unstage(backRepoParagraph.GetStage())

		// remove instance from the back repo 3 maps
		paragraphID := backRepoParagraph.Map_ParagraphPtr_ParagraphDBID[paragraph]
		delete(backRepoParagraph.Map_ParagraphPtr_ParagraphDBID, paragraph)
		delete(backRepoParagraph.Map_ParagraphDBID_ParagraphDB, paragraphID)
		delete(backRepoParagraph.Map_ParagraphDBID_ParagraphPtr, paragraphID)
	}

	return
}

// CheckoutPhaseOneInstance takes a paragraphDB that has been found in the DB, updates the backRepo and stages the
// models version of the paragraphDB
func (backRepoParagraph *BackRepoParagraphStruct) CheckoutPhaseOneInstance(paragraphDB *ParagraphDB) (Error error) {

	paragraph, ok := backRepoParagraph.Map_ParagraphDBID_ParagraphPtr[paragraphDB.ID]
	if !ok {
		paragraph = new(models.Paragraph)

		backRepoParagraph.Map_ParagraphDBID_ParagraphPtr[paragraphDB.ID] = paragraph
		backRepoParagraph.Map_ParagraphPtr_ParagraphDBID[paragraph] = paragraphDB.ID

		// append model store with the new element
		paragraph.Name = paragraphDB.Name_Data.String
		paragraph.Stage(backRepoParagraph.GetStage())
	}
	paragraphDB.CopyBasicFieldsToParagraph(paragraph)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	paragraph.Stage(backRepoParagraph.GetStage())

	// preserve pointer to paragraphDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ParagraphDBID_ParagraphDB)[paragraphDB hold variable pointers
	paragraphDB_Data := *paragraphDB
	preservedPtrToParagraph := &paragraphDB_Data
	backRepoParagraph.Map_ParagraphDBID_ParagraphDB[paragraphDB.ID] = preservedPtrToParagraph

	return
}

// BackRepoParagraph.CheckoutPhaseTwo Checkouts all staged instances of Paragraph to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoParagraph *BackRepoParagraphStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, paragraphDB := range backRepoParagraph.Map_ParagraphDBID_ParagraphDB {
		backRepoParagraph.CheckoutPhaseTwoInstance(backRepo, paragraphDB)
	}
	return
}

// BackRepoParagraph.CheckoutPhaseTwoInstance Checkouts staged instances of Paragraph to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoParagraph *BackRepoParagraphStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, paragraphDB *ParagraphDB) (Error error) {

	paragraph := backRepoParagraph.Map_ParagraphDBID_ParagraphPtr[paragraphDB.ID]

	paragraphDB.DecodePointers(backRepo, paragraph)

	return
}

func (paragraphDB *ParagraphDB) DecodePointers(backRepo *BackRepoStruct, paragraph *models.Paragraph) {

	// insertion point for checkout of pointer encoding
	// Image field	
	{
		id := paragraphDB.ImageID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoImage.Map_ImageDBID_ImagePtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: paragraph.Image, unknown pointer id", id)
				paragraph.Image = nil
			} else {
				// updates only if field has changed
				if paragraph.Image == nil || paragraph.Image != tmp {
					paragraph.Image = tmp
				}
			}
		} else {
			paragraph.Image = nil
		}
	}
	
	return
}

// CommitParagraph allows commit of a single paragraph (if already staged)
func (backRepo *BackRepoStruct) CommitParagraph(paragraph *models.Paragraph) {
	backRepo.BackRepoParagraph.CommitPhaseOneInstance(paragraph)
	if id, ok := backRepo.BackRepoParagraph.Map_ParagraphPtr_ParagraphDBID[paragraph]; ok {
		backRepo.BackRepoParagraph.CommitPhaseTwoInstance(backRepo, id, paragraph)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitParagraph allows checkout of a single paragraph (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutParagraph(paragraph *models.Paragraph) {
	// check if the paragraph is staged
	if _, ok := backRepo.BackRepoParagraph.Map_ParagraphPtr_ParagraphDBID[paragraph]; ok {

		if id, ok := backRepo.BackRepoParagraph.Map_ParagraphPtr_ParagraphDBID[paragraph]; ok {
			var paragraphDB ParagraphDB
			paragraphDB.ID = id

			if _, err := backRepo.BackRepoParagraph.db.First(&paragraphDB, id); err != nil {
				log.Fatalln("CheckoutParagraph : Problem with getting object with id:", id)
			}
			backRepo.BackRepoParagraph.CheckoutPhaseOneInstance(&paragraphDB)
			backRepo.BackRepoParagraph.CheckoutPhaseTwoInstance(backRepo, &paragraphDB)
		}
	}
}

// CopyBasicFieldsFromParagraph
func (paragraphDB *ParagraphDB) CopyBasicFieldsFromParagraph(paragraph *models.Paragraph) {
	// insertion point for fields commit

	paragraphDB.Name_Data.String = paragraph.Name
	paragraphDB.Name_Data.Valid = true

	paragraphDB.LegendMarkdownContent_Data.String = paragraph.LegendMarkdownContent
	paragraphDB.LegendMarkdownContent_Data.Valid = true
}

// CopyBasicFieldsFromParagraph_WOP
func (paragraphDB *ParagraphDB) CopyBasicFieldsFromParagraph_WOP(paragraph *models.Paragraph_WOP) {
	// insertion point for fields commit

	paragraphDB.Name_Data.String = paragraph.Name
	paragraphDB.Name_Data.Valid = true

	paragraphDB.LegendMarkdownContent_Data.String = paragraph.LegendMarkdownContent
	paragraphDB.LegendMarkdownContent_Data.Valid = true
}

// CopyBasicFieldsFromParagraphWOP
func (paragraphDB *ParagraphDB) CopyBasicFieldsFromParagraphWOP(paragraph *ParagraphWOP) {
	// insertion point for fields commit

	paragraphDB.Name_Data.String = paragraph.Name
	paragraphDB.Name_Data.Valid = true

	paragraphDB.LegendMarkdownContent_Data.String = paragraph.LegendMarkdownContent
	paragraphDB.LegendMarkdownContent_Data.Valid = true
}

// CopyBasicFieldsToParagraph
func (paragraphDB *ParagraphDB) CopyBasicFieldsToParagraph(paragraph *models.Paragraph) {
	// insertion point for checkout of basic fields (back repo to stage)
	paragraph.Name = paragraphDB.Name_Data.String
	paragraph.LegendMarkdownContent = paragraphDB.LegendMarkdownContent_Data.String
}

// CopyBasicFieldsToParagraph_WOP
func (paragraphDB *ParagraphDB) CopyBasicFieldsToParagraph_WOP(paragraph *models.Paragraph_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	paragraph.Name = paragraphDB.Name_Data.String
	paragraph.LegendMarkdownContent = paragraphDB.LegendMarkdownContent_Data.String
}

// CopyBasicFieldsToParagraphWOP
func (paragraphDB *ParagraphDB) CopyBasicFieldsToParagraphWOP(paragraph *ParagraphWOP) {
	paragraph.ID = int(paragraphDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	paragraph.Name = paragraphDB.Name_Data.String
	paragraph.LegendMarkdownContent = paragraphDB.LegendMarkdownContent_Data.String
}

// Backup generates a json file from a slice of all ParagraphDB instances in the backrepo
func (backRepoParagraph *BackRepoParagraphStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ParagraphDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ParagraphDB, 0)
	for _, paragraphDB := range backRepoParagraph.Map_ParagraphDBID_ParagraphDB {
		forBackup = append(forBackup, paragraphDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Paragraph ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Paragraph file", err.Error())
	}
}

// Backup generates a json file from a slice of all ParagraphDB instances in the backrepo
func (backRepoParagraph *BackRepoParagraphStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ParagraphDB, 0)
	for _, paragraphDB := range backRepoParagraph.Map_ParagraphDBID_ParagraphDB {
		forBackup = append(forBackup, paragraphDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Paragraph")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Paragraph_Fields, -1)
	for _, paragraphDB := range forBackup {

		var paragraphWOP ParagraphWOP
		paragraphDB.CopyBasicFieldsToParagraphWOP(&paragraphWOP)

		row := sh.AddRow()
		row.WriteStruct(&paragraphWOP, -1)
	}
}

// RestoreXL from the "Paragraph" sheet all ParagraphDB instances
func (backRepoParagraph *BackRepoParagraphStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoParagraphid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Paragraph"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoParagraph.rowVisitorParagraph)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoParagraph *BackRepoParagraphStruct) rowVisitorParagraph(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var paragraphWOP ParagraphWOP
		row.ReadStruct(&paragraphWOP)

		// add the unmarshalled struct to the stage
		paragraphDB := new(ParagraphDB)
		paragraphDB.CopyBasicFieldsFromParagraphWOP(&paragraphWOP)

		paragraphDB_ID_atBackupTime := paragraphDB.ID
		paragraphDB.ID = 0
		_, err := backRepoParagraph.db.Create(paragraphDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoParagraph.Map_ParagraphDBID_ParagraphDB[paragraphDB.ID] = paragraphDB
		BackRepoParagraphid_atBckpTime_newID[paragraphDB_ID_atBackupTime] = paragraphDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ParagraphDB.json" in dirPath that stores an array
// of ParagraphDB and stores it in the database
// the map BackRepoParagraphid_atBckpTime_newID is updated accordingly
func (backRepoParagraph *BackRepoParagraphStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoParagraphid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ParagraphDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Paragraph file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ParagraphDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ParagraphDBID_ParagraphDB
	for _, paragraphDB := range forRestore {

		paragraphDB_ID_atBackupTime := paragraphDB.ID
		paragraphDB.ID = 0
		_, err := backRepoParagraph.db.Create(paragraphDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoParagraph.Map_ParagraphDBID_ParagraphDB[paragraphDB.ID] = paragraphDB
		BackRepoParagraphid_atBckpTime_newID[paragraphDB_ID_atBackupTime] = paragraphDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Paragraph file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Paragraph>id_atBckpTime_newID
// to compute new index
func (backRepoParagraph *BackRepoParagraphStruct) RestorePhaseTwo() {

	for _, paragraphDB := range backRepoParagraph.Map_ParagraphDBID_ParagraphDB {

		// next line of code is to avert unused variable compilation error
		_ = paragraphDB

		// insertion point for reindexing pointers encoding
		// reindexing Image field
		if paragraphDB.ImageID.Int64 != 0 {
			paragraphDB.ImageID.Int64 = int64(BackRepoImageid_atBckpTime_newID[uint(paragraphDB.ImageID.Int64)])
			paragraphDB.ImageID.Valid = true
		}

		// update databse with new index encoding
		db, _ := backRepoParagraph.db.Model(paragraphDB)
		_, err := db.Updates(*paragraphDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoParagraph.ResetReversePointers commits all staged instances of Paragraph to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoParagraph *BackRepoParagraphStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, paragraph := range backRepoParagraph.Map_ParagraphDBID_ParagraphPtr {
		backRepoParagraph.ResetReversePointersInstance(backRepo, idx, paragraph)
	}

	return
}

func (backRepoParagraph *BackRepoParagraphStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, paragraph *models.Paragraph) (Error error) {

	// fetch matching paragraphDB
	if paragraphDB, ok := backRepoParagraph.Map_ParagraphDBID_ParagraphDB[idx]; ok {
		_ = paragraphDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoParagraphid_atBckpTime_newID map[uint]uint
