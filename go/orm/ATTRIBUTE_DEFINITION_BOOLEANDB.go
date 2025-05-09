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
var dummy_ATTRIBUTE_DEFINITION_BOOLEAN_sql sql.NullBool
var dummy_ATTRIBUTE_DEFINITION_BOOLEAN_time time.Duration
var dummy_ATTRIBUTE_DEFINITION_BOOLEAN_sort sort.Float64Slice

// ATTRIBUTE_DEFINITION_BOOLEANAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model attribute_definition_booleanAPI
type ATTRIBUTE_DEFINITION_BOOLEANAPI struct {
	gorm.Model

	models.ATTRIBUTE_DEFINITION_BOOLEAN_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding
}

// ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	ALTERNATIVE_ID struct {

		// field ALTERNATIVE_ID is a slice of pointers to another Struct (optional or 0..1)
		ALTERNATIVE_ID IntSlice `gorm:"type:TEXT"`

	} `gorm:"embedded"`

	DEFAULT_VALUE struct {

		// field ATTRIBUTE_VALUE_BOOLEAN is a slice of pointers to another Struct (optional or 0..1)
		ATTRIBUTE_VALUE_BOOLEAN IntSlice `gorm:"type:TEXT"`

	} `gorm:"embedded"`
}

// ATTRIBUTE_DEFINITION_BOOLEANDB describes a attribute_definition_boolean in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model attribute_definition_booleanDB
type ATTRIBUTE_DEFINITION_BOOLEANDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field attribute_definition_booleanDB.Name
	Name_Data sql.NullString

	// Declation for basic field attribute_definition_booleanDB.DESC
	DESC_Data sql.NullString

	// Declation for basic field attribute_definition_booleanDB.IS_EDITABLE
	// provide the sql storage for the boolan
	IS_EDITABLE_Data sql.NullBool

	// Declation for basic field attribute_definition_booleanDB.LAST_CHANGE
	LAST_CHANGE_Data sql.NullTime

	// Declation for basic field attribute_definition_booleanDB.LONG_NAME
	LONG_NAME_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding
}

// ATTRIBUTE_DEFINITION_BOOLEANDBs arrays attribute_definition_booleanDBs
// swagger:response attribute_definition_booleanDBsResponse
type ATTRIBUTE_DEFINITION_BOOLEANDBs []ATTRIBUTE_DEFINITION_BOOLEANDB

// ATTRIBUTE_DEFINITION_BOOLEANDBResponse provides response
// swagger:response attribute_definition_booleanDBResponse
type ATTRIBUTE_DEFINITION_BOOLEANDBResponse struct {
	ATTRIBUTE_DEFINITION_BOOLEANDB
}

// ATTRIBUTE_DEFINITION_BOOLEANWOP is a ATTRIBUTE_DEFINITION_BOOLEAN without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ATTRIBUTE_DEFINITION_BOOLEANWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESC string `xlsx:"2"`

	IS_EDITABLE bool `xlsx:"3"`

	LAST_CHANGE time.Time `xlsx:"4"`

	LONG_NAME string `xlsx:"5"`
	// insertion for WOP pointer fields
}

var ATTRIBUTE_DEFINITION_BOOLEAN_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESC",
	"IS_EDITABLE",
	"LAST_CHANGE",
	"LONG_NAME",
}

type BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct struct {
	// stores ATTRIBUTE_DEFINITION_BOOLEANDB according to their gorm ID
	Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB map[uint]*ATTRIBUTE_DEFINITION_BOOLEANDB

	// stores ATTRIBUTE_DEFINITION_BOOLEANDB ID according to ATTRIBUTE_DEFINITION_BOOLEAN address
	Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID map[*models.ATTRIBUTE_DEFINITION_BOOLEAN]uint

	// stores ATTRIBUTE_DEFINITION_BOOLEAN according to their gorm ID
	Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr map[uint]*models.ATTRIBUTE_DEFINITION_BOOLEAN

	db db.DBInterface

	stage *models.Stage
}

func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) GetStage() (stage *models.Stage) {
	stage = backRepoATTRIBUTE_DEFINITION_BOOLEAN.stage
	return
}

func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) GetDB() db.DBInterface {
	return backRepoATTRIBUTE_DEFINITION_BOOLEAN.db
}

// GetATTRIBUTE_DEFINITION_BOOLEANDBFromATTRIBUTE_DEFINITION_BOOLEANPtr is a handy function to access the back repo instance from the stage instance
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) GetATTRIBUTE_DEFINITION_BOOLEANDBFromATTRIBUTE_DEFINITION_BOOLEANPtr(attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN) (attribute_definition_booleanDB *ATTRIBUTE_DEFINITION_BOOLEANDB) {
	id := backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID[attribute_definition_boolean]
	attribute_definition_booleanDB = backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB[id]
	return
}

// BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseOne commits all staged instances of ATTRIBUTE_DEFINITION_BOOLEAN to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var attribute_definition_booleans []*models.ATTRIBUTE_DEFINITION_BOOLEAN
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		attribute_definition_booleans = append(attribute_definition_booleans, attribute_definition_boolean)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(attribute_definition_booleans, func(i, j int) bool {
		return stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_booleans[i]] < stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_booleans[j]]
	})

	for _, attribute_definition_boolean := range attribute_definition_booleans {
		backRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseOneInstance(attribute_definition_boolean)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, attribute_definition_boolean := range backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr {
		if _, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean]; !ok {
			backRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitDeleteInstance commits deletion of ATTRIBUTE_DEFINITION_BOOLEAN to the BackRepo
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) CommitDeleteInstance(id uint) (Error error) {

	attribute_definition_boolean := backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr[id]

	// attribute_definition_boolean is not staged anymore, remove attribute_definition_booleanDB
	attribute_definition_booleanDB := backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB[id]
	db, _ := backRepoATTRIBUTE_DEFINITION_BOOLEAN.db.Unscoped()
	_, err := db.Delete(attribute_definition_booleanDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID, attribute_definition_boolean)
	delete(backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr, id)
	delete(backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB, id)

	return
}

// BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseOneInstance commits attribute_definition_boolean staged instances of ATTRIBUTE_DEFINITION_BOOLEAN to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) CommitPhaseOneInstance(attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN) (Error error) {

	// check if the attribute_definition_boolean is not commited yet
	if _, ok := backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID[attribute_definition_boolean]; ok {
		return
	}

	// initiate attribute_definition_boolean
	var attribute_definition_booleanDB ATTRIBUTE_DEFINITION_BOOLEANDB
	attribute_definition_booleanDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean)

	_, err := backRepoATTRIBUTE_DEFINITION_BOOLEAN.db.Create(&attribute_definition_booleanDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID[attribute_definition_boolean] = attribute_definition_booleanDB.ID
	backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr[attribute_definition_booleanDB.ID] = attribute_definition_boolean
	backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB[attribute_definition_booleanDB.ID] = &attribute_definition_booleanDB

	return
}

// BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseTwo commits all staged instances of ATTRIBUTE_DEFINITION_BOOLEAN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, attribute_definition_boolean := range backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr {
		backRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseTwoInstance(backRepo, idx, attribute_definition_boolean)
	}

	return
}

// BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseTwoInstance commits {{structname }} of models.ATTRIBUTE_DEFINITION_BOOLEAN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN) (Error error) {

	// fetch matching attribute_definition_booleanDB
	if attribute_definition_booleanDB, ok := backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB[idx]; ok {

		attribute_definition_booleanDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID = make([]int, 0)
		// 2. encode
		for _, alternative_idAssocEnd := range attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
			alternative_idAssocEnd_DB :=
				backRepo.BackRepoALTERNATIVE_ID.GetALTERNATIVE_IDDBFromALTERNATIVE_IDPtr(alternative_idAssocEnd)
			
			// the stage might be inconsistant, meaning that the alternative_idAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if alternative_idAssocEnd_DB == nil {
				continue
			}
			
			attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID =
				append(attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID, int(alternative_idAssocEnd_DB.ID))
		}

		// 1. reset
		attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN = make([]int, 0)
		// 2. encode
		for _, attribute_value_booleanAssocEnd := range attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
			attribute_value_booleanAssocEnd_DB :=
				backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.GetATTRIBUTE_VALUE_BOOLEANDBFromATTRIBUTE_VALUE_BOOLEANPtr(attribute_value_booleanAssocEnd)
			
			// the stage might be inconsistant, meaning that the attribute_value_booleanAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if attribute_value_booleanAssocEnd_DB == nil {
				continue
			}
			
			attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN =
				append(attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, int(attribute_value_booleanAssocEnd_DB.ID))
		}

		_, err := backRepoATTRIBUTE_DEFINITION_BOOLEAN.db.Save(attribute_definition_booleanDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown ATTRIBUTE_DEFINITION_BOOLEAN intance %s", attribute_definition_boolean.Name))
		return err
	}

	return
}

// BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) CheckoutPhaseOne() (Error error) {

	attribute_definition_booleanDBArray := make([]ATTRIBUTE_DEFINITION_BOOLEANDB, 0)
	_, err := backRepoATTRIBUTE_DEFINITION_BOOLEAN.db.Find(&attribute_definition_booleanDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	attribute_definition_booleanInstancesToBeRemovedFromTheStage := make(map[*models.ATTRIBUTE_DEFINITION_BOOLEAN]any)
	for key, value := range backRepoATTRIBUTE_DEFINITION_BOOLEAN.stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		attribute_definition_booleanInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, attribute_definition_booleanDB := range attribute_definition_booleanDBArray {
		backRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseOneInstance(&attribute_definition_booleanDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		attribute_definition_boolean, ok := backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr[attribute_definition_booleanDB.ID]
		if ok {
			delete(attribute_definition_booleanInstancesToBeRemovedFromTheStage, attribute_definition_boolean)
		}
	}

	// remove from stage and back repo's 3 maps all attribute_definition_booleans that are not in the checkout
	for attribute_definition_boolean := range attribute_definition_booleanInstancesToBeRemovedFromTheStage {
		attribute_definition_boolean.Unstage(backRepoATTRIBUTE_DEFINITION_BOOLEAN.GetStage())

		// remove instance from the back repo 3 maps
		attribute_definition_booleanID := backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID[attribute_definition_boolean]
		delete(backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID, attribute_definition_boolean)
		delete(backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB, attribute_definition_booleanID)
		delete(backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr, attribute_definition_booleanID)
	}

	return
}

// CheckoutPhaseOneInstance takes a attribute_definition_booleanDB that has been found in the DB, updates the backRepo and stages the
// models version of the attribute_definition_booleanDB
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) CheckoutPhaseOneInstance(attribute_definition_booleanDB *ATTRIBUTE_DEFINITION_BOOLEANDB) (Error error) {

	attribute_definition_boolean, ok := backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr[attribute_definition_booleanDB.ID]
	if !ok {
		attribute_definition_boolean = new(models.ATTRIBUTE_DEFINITION_BOOLEAN)

		backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr[attribute_definition_booleanDB.ID] = attribute_definition_boolean
		backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID[attribute_definition_boolean] = attribute_definition_booleanDB.ID

		// append model store with the new element
		attribute_definition_boolean.Name = attribute_definition_booleanDB.Name_Data.String
		attribute_definition_boolean.Stage(backRepoATTRIBUTE_DEFINITION_BOOLEAN.GetStage())
	}
	attribute_definition_booleanDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	attribute_definition_boolean.Stage(backRepoATTRIBUTE_DEFINITION_BOOLEAN.GetStage())

	// preserve pointer to attribute_definition_booleanDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB)[attribute_definition_booleanDB hold variable pointers
	attribute_definition_booleanDB_Data := *attribute_definition_booleanDB
	preservedPtrToATTRIBUTE_DEFINITION_BOOLEAN := &attribute_definition_booleanDB_Data
	backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB[attribute_definition_booleanDB.ID] = preservedPtrToATTRIBUTE_DEFINITION_BOOLEAN

	return
}

// BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseTwo Checkouts all staged instances of ATTRIBUTE_DEFINITION_BOOLEAN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, attribute_definition_booleanDB := range backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB {
		backRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseTwoInstance(backRepo, attribute_definition_booleanDB)
	}
	return
}

// BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseTwoInstance Checkouts staged instances of ATTRIBUTE_DEFINITION_BOOLEAN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, attribute_definition_booleanDB *ATTRIBUTE_DEFINITION_BOOLEANDB) (Error error) {

	attribute_definition_boolean := backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr[attribute_definition_booleanDB.ID]

	attribute_definition_booleanDB.DecodePointers(backRepo, attribute_definition_boolean)

	return
}

func (attribute_definition_booleanDB *ATTRIBUTE_DEFINITION_BOOLEANDB) DecodePointers(backRepo *BackRepoStruct, attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN) {

	// insertion point for checkout of pointer encoding
	// This loop redeem attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID in the stage from the encode in the back repo
	// It parses all ALTERNATIVE_IDDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID = attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID[:0]
	for _, _ALTERNATIVE_IDid := range attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID {
		attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID = append(attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID, backRepo.BackRepoALTERNATIVE_ID.Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDPtr[uint(_ALTERNATIVE_IDid)])
	}

	// This loop redeem attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN in the stage from the encode in the back repo
	// It parses all ATTRIBUTE_VALUE_BOOLEANDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN = attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN[:0]
	for _, _ATTRIBUTE_VALUE_BOOLEANid := range attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
		attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN = append(attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.Map_ATTRIBUTE_VALUE_BOOLEANDBID_ATTRIBUTE_VALUE_BOOLEANPtr[uint(_ATTRIBUTE_VALUE_BOOLEANid)])
	}

	return
}

// CommitATTRIBUTE_DEFINITION_BOOLEAN allows commit of a single attribute_definition_boolean (if already staged)
func (backRepo *BackRepoStruct) CommitATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN) {
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseOneInstance(attribute_definition_boolean)
	if id, ok := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID[attribute_definition_boolean]; ok {
		backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseTwoInstance(backRepo, id, attribute_definition_boolean)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitATTRIBUTE_DEFINITION_BOOLEAN allows checkout of a single attribute_definition_boolean (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN) {
	// check if the attribute_definition_boolean is staged
	if _, ok := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID[attribute_definition_boolean]; ok {

		if id, ok := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID[attribute_definition_boolean]; ok {
			var attribute_definition_booleanDB ATTRIBUTE_DEFINITION_BOOLEANDB
			attribute_definition_booleanDB.ID = id

			if _, err := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.db.First(&attribute_definition_booleanDB, id); err != nil {
				log.Fatalln("CheckoutATTRIBUTE_DEFINITION_BOOLEAN : Problem with getting object with id:", id)
			}
			backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseOneInstance(&attribute_definition_booleanDB)
			backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseTwoInstance(backRepo, &attribute_definition_booleanDB)
		}
	}
}

// CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEAN
func (attribute_definition_booleanDB *ATTRIBUTE_DEFINITION_BOOLEANDB) CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN) {
	// insertion point for fields commit

	attribute_definition_booleanDB.Name_Data.String = attribute_definition_boolean.Name
	attribute_definition_booleanDB.Name_Data.Valid = true

	attribute_definition_booleanDB.DESC_Data.String = attribute_definition_boolean.DESC
	attribute_definition_booleanDB.DESC_Data.Valid = true

	attribute_definition_booleanDB.IS_EDITABLE_Data.Bool = attribute_definition_boolean.IS_EDITABLE
	attribute_definition_booleanDB.IS_EDITABLE_Data.Valid = true

	attribute_definition_booleanDB.LAST_CHANGE_Data.Time = attribute_definition_boolean.LAST_CHANGE
	attribute_definition_booleanDB.LAST_CHANGE_Data.Valid = true

	attribute_definition_booleanDB.LONG_NAME_Data.String = attribute_definition_boolean.LONG_NAME
	attribute_definition_booleanDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEAN_WOP
func (attribute_definition_booleanDB *ATTRIBUTE_DEFINITION_BOOLEANDB) CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEAN_WOP(attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN_WOP) {
	// insertion point for fields commit

	attribute_definition_booleanDB.Name_Data.String = attribute_definition_boolean.Name
	attribute_definition_booleanDB.Name_Data.Valid = true

	attribute_definition_booleanDB.DESC_Data.String = attribute_definition_boolean.DESC
	attribute_definition_booleanDB.DESC_Data.Valid = true

	attribute_definition_booleanDB.IS_EDITABLE_Data.Bool = attribute_definition_boolean.IS_EDITABLE
	attribute_definition_booleanDB.IS_EDITABLE_Data.Valid = true

	attribute_definition_booleanDB.LAST_CHANGE_Data.Time = attribute_definition_boolean.LAST_CHANGE
	attribute_definition_booleanDB.LAST_CHANGE_Data.Valid = true

	attribute_definition_booleanDB.LONG_NAME_Data.String = attribute_definition_boolean.LONG_NAME
	attribute_definition_booleanDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEANWOP
func (attribute_definition_booleanDB *ATTRIBUTE_DEFINITION_BOOLEANDB) CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEANWOP(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEANWOP) {
	// insertion point for fields commit

	attribute_definition_booleanDB.Name_Data.String = attribute_definition_boolean.Name
	attribute_definition_booleanDB.Name_Data.Valid = true

	attribute_definition_booleanDB.DESC_Data.String = attribute_definition_boolean.DESC
	attribute_definition_booleanDB.DESC_Data.Valid = true

	attribute_definition_booleanDB.IS_EDITABLE_Data.Bool = attribute_definition_boolean.IS_EDITABLE
	attribute_definition_booleanDB.IS_EDITABLE_Data.Valid = true

	attribute_definition_booleanDB.LAST_CHANGE_Data.Time = attribute_definition_boolean.LAST_CHANGE
	attribute_definition_booleanDB.LAST_CHANGE_Data.Valid = true

	attribute_definition_booleanDB.LONG_NAME_Data.String = attribute_definition_boolean.LONG_NAME
	attribute_definition_booleanDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN
func (attribute_definition_booleanDB *ATTRIBUTE_DEFINITION_BOOLEANDB) CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN) {
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_definition_boolean.Name = attribute_definition_booleanDB.Name_Data.String
	attribute_definition_boolean.DESC = attribute_definition_booleanDB.DESC_Data.String
	attribute_definition_boolean.IS_EDITABLE = attribute_definition_booleanDB.IS_EDITABLE_Data.Bool
	attribute_definition_boolean.LAST_CHANGE = attribute_definition_booleanDB.LAST_CHANGE_Data.Time
	attribute_definition_boolean.LONG_NAME = attribute_definition_booleanDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN_WOP
func (attribute_definition_booleanDB *ATTRIBUTE_DEFINITION_BOOLEANDB) CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN_WOP(attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_definition_boolean.Name = attribute_definition_booleanDB.Name_Data.String
	attribute_definition_boolean.DESC = attribute_definition_booleanDB.DESC_Data.String
	attribute_definition_boolean.IS_EDITABLE = attribute_definition_booleanDB.IS_EDITABLE_Data.Bool
	attribute_definition_boolean.LAST_CHANGE = attribute_definition_booleanDB.LAST_CHANGE_Data.Time
	attribute_definition_boolean.LONG_NAME = attribute_definition_booleanDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEANWOP
func (attribute_definition_booleanDB *ATTRIBUTE_DEFINITION_BOOLEANDB) CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEANWOP(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEANWOP) {
	attribute_definition_boolean.ID = int(attribute_definition_booleanDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_definition_boolean.Name = attribute_definition_booleanDB.Name_Data.String
	attribute_definition_boolean.DESC = attribute_definition_booleanDB.DESC_Data.String
	attribute_definition_boolean.IS_EDITABLE = attribute_definition_booleanDB.IS_EDITABLE_Data.Bool
	attribute_definition_boolean.LAST_CHANGE = attribute_definition_booleanDB.LAST_CHANGE_Data.Time
	attribute_definition_boolean.LONG_NAME = attribute_definition_booleanDB.LONG_NAME_Data.String
}

// Backup generates a json file from a slice of all ATTRIBUTE_DEFINITION_BOOLEANDB instances in the backrepo
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ATTRIBUTE_DEFINITION_BOOLEANDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTE_DEFINITION_BOOLEANDB, 0)
	for _, attribute_definition_booleanDB := range backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB {
		forBackup = append(forBackup, attribute_definition_booleanDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json ATTRIBUTE_DEFINITION_BOOLEAN ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json ATTRIBUTE_DEFINITION_BOOLEAN file", err.Error())
	}
}

// Backup generates a json file from a slice of all ATTRIBUTE_DEFINITION_BOOLEANDB instances in the backrepo
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTE_DEFINITION_BOOLEANDB, 0)
	for _, attribute_definition_booleanDB := range backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB {
		forBackup = append(forBackup, attribute_definition_booleanDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("ATTRIBUTE_DEFINITION_BOOLEAN")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&ATTRIBUTE_DEFINITION_BOOLEAN_Fields, -1)
	for _, attribute_definition_booleanDB := range forBackup {

		var attribute_definition_booleanWOP ATTRIBUTE_DEFINITION_BOOLEANWOP
		attribute_definition_booleanDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEANWOP(&attribute_definition_booleanWOP)

		row := sh.AddRow()
		row.WriteStruct(&attribute_definition_booleanWOP, -1)
	}
}

// RestoreXL from the "ATTRIBUTE_DEFINITION_BOOLEAN" sheet all ATTRIBUTE_DEFINITION_BOOLEANDB instances
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoATTRIBUTE_DEFINITION_BOOLEANid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["ATTRIBUTE_DEFINITION_BOOLEAN"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoATTRIBUTE_DEFINITION_BOOLEAN.rowVisitorATTRIBUTE_DEFINITION_BOOLEAN)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) rowVisitorATTRIBUTE_DEFINITION_BOOLEAN(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var attribute_definition_booleanWOP ATTRIBUTE_DEFINITION_BOOLEANWOP
		row.ReadStruct(&attribute_definition_booleanWOP)

		// add the unmarshalled struct to the stage
		attribute_definition_booleanDB := new(ATTRIBUTE_DEFINITION_BOOLEANDB)
		attribute_definition_booleanDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEANWOP(&attribute_definition_booleanWOP)

		attribute_definition_booleanDB_ID_atBackupTime := attribute_definition_booleanDB.ID
		attribute_definition_booleanDB.ID = 0
		_, err := backRepoATTRIBUTE_DEFINITION_BOOLEAN.db.Create(attribute_definition_booleanDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB[attribute_definition_booleanDB.ID] = attribute_definition_booleanDB
		BackRepoATTRIBUTE_DEFINITION_BOOLEANid_atBckpTime_newID[attribute_definition_booleanDB_ID_atBackupTime] = attribute_definition_booleanDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ATTRIBUTE_DEFINITION_BOOLEANDB.json" in dirPath that stores an array
// of ATTRIBUTE_DEFINITION_BOOLEANDB and stores it in the database
// the map BackRepoATTRIBUTE_DEFINITION_BOOLEANid_atBckpTime_newID is updated accordingly
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoATTRIBUTE_DEFINITION_BOOLEANid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ATTRIBUTE_DEFINITION_BOOLEANDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json ATTRIBUTE_DEFINITION_BOOLEAN file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ATTRIBUTE_DEFINITION_BOOLEANDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB
	for _, attribute_definition_booleanDB := range forRestore {

		attribute_definition_booleanDB_ID_atBackupTime := attribute_definition_booleanDB.ID
		attribute_definition_booleanDB.ID = 0
		_, err := backRepoATTRIBUTE_DEFINITION_BOOLEAN.db.Create(attribute_definition_booleanDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB[attribute_definition_booleanDB.ID] = attribute_definition_booleanDB
		BackRepoATTRIBUTE_DEFINITION_BOOLEANid_atBckpTime_newID[attribute_definition_booleanDB_ID_atBackupTime] = attribute_definition_booleanDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json ATTRIBUTE_DEFINITION_BOOLEAN file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<ATTRIBUTE_DEFINITION_BOOLEAN>id_atBckpTime_newID
// to compute new index
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) RestorePhaseTwo() {

	for _, attribute_definition_booleanDB := range backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB {

		// next line of code is to avert unused variable compilation error
		_ = attribute_definition_booleanDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoATTRIBUTE_DEFINITION_BOOLEAN.db.Model(attribute_definition_booleanDB)
		_, err := db.Updates(*attribute_definition_booleanDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoATTRIBUTE_DEFINITION_BOOLEAN.ResetReversePointers commits all staged instances of ATTRIBUTE_DEFINITION_BOOLEAN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, attribute_definition_boolean := range backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr {
		backRepoATTRIBUTE_DEFINITION_BOOLEAN.ResetReversePointersInstance(backRepo, idx, attribute_definition_boolean)
	}

	return
}

func (backRepoATTRIBUTE_DEFINITION_BOOLEAN *BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN) (Error error) {

	// fetch matching attribute_definition_booleanDB
	if attribute_definition_booleanDB, ok := backRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB[idx]; ok {
		_ = attribute_definition_booleanDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoATTRIBUTE_DEFINITION_BOOLEANid_atBckpTime_newID map[uint]uint
