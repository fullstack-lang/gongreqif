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
var dummy_ATTRIBUTE_DEFINITION_STRING_sql sql.NullBool
var dummy_ATTRIBUTE_DEFINITION_STRING_time time.Duration
var dummy_ATTRIBUTE_DEFINITION_STRING_sort sort.Float64Slice

// ATTRIBUTE_DEFINITION_STRINGAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model attribute_definition_stringAPI
type ATTRIBUTE_DEFINITION_STRINGAPI struct {
	gorm.Model

	models.ATTRIBUTE_DEFINITION_STRING_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ATTRIBUTE_DEFINITION_STRINGPointersEncoding ATTRIBUTE_DEFINITION_STRINGPointersEncoding
}

// ATTRIBUTE_DEFINITION_STRINGPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ATTRIBUTE_DEFINITION_STRINGPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	ALTERNATIVE_ID struct {

		// field ALTERNATIVE_ID is a slice of pointers to another Struct (optional or 0..1)
		ALTERNATIVE_ID IntSlice `gorm:"type:TEXT"`

	} `gorm:"embedded"`

	DEFAULT_VALUE struct {

		// field ATTRIBUTE_VALUE_STRING is a slice of pointers to another Struct (optional or 0..1)
		ATTRIBUTE_VALUE_STRING IntSlice `gorm:"type:TEXT"`

	} `gorm:"embedded"`
}

// ATTRIBUTE_DEFINITION_STRINGDB describes a attribute_definition_string in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model attribute_definition_stringDB
type ATTRIBUTE_DEFINITION_STRINGDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field attribute_definition_stringDB.Name
	Name_Data sql.NullString

	// Declation for basic field attribute_definition_stringDB.DESC
	DESC_Data sql.NullString

	// Declation for basic field attribute_definition_stringDB.IS_EDITABLE
	// provide the sql storage for the boolan
	IS_EDITABLE_Data sql.NullBool

	// Declation for basic field attribute_definition_stringDB.LAST_CHANGE
	LAST_CHANGE_Data sql.NullTime

	// Declation for basic field attribute_definition_stringDB.LONG_NAME
	LONG_NAME_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ATTRIBUTE_DEFINITION_STRINGPointersEncoding
}

// ATTRIBUTE_DEFINITION_STRINGDBs arrays attribute_definition_stringDBs
// swagger:response attribute_definition_stringDBsResponse
type ATTRIBUTE_DEFINITION_STRINGDBs []ATTRIBUTE_DEFINITION_STRINGDB

// ATTRIBUTE_DEFINITION_STRINGDBResponse provides response
// swagger:response attribute_definition_stringDBResponse
type ATTRIBUTE_DEFINITION_STRINGDBResponse struct {
	ATTRIBUTE_DEFINITION_STRINGDB
}

// ATTRIBUTE_DEFINITION_STRINGWOP is a ATTRIBUTE_DEFINITION_STRING without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ATTRIBUTE_DEFINITION_STRINGWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESC string `xlsx:"2"`

	IS_EDITABLE bool `xlsx:"3"`

	LAST_CHANGE time.Time `xlsx:"4"`

	LONG_NAME string `xlsx:"5"`
	// insertion for WOP pointer fields
}

var ATTRIBUTE_DEFINITION_STRING_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESC",
	"IS_EDITABLE",
	"LAST_CHANGE",
	"LONG_NAME",
}

type BackRepoATTRIBUTE_DEFINITION_STRINGStruct struct {
	// stores ATTRIBUTE_DEFINITION_STRINGDB according to their gorm ID
	Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB map[uint]*ATTRIBUTE_DEFINITION_STRINGDB

	// stores ATTRIBUTE_DEFINITION_STRINGDB ID according to ATTRIBUTE_DEFINITION_STRING address
	Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID map[*models.ATTRIBUTE_DEFINITION_STRING]uint

	// stores ATTRIBUTE_DEFINITION_STRING according to their gorm ID
	Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr map[uint]*models.ATTRIBUTE_DEFINITION_STRING

	db db.DBInterface

	stage *models.Stage
}

func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) GetStage() (stage *models.Stage) {
	stage = backRepoATTRIBUTE_DEFINITION_STRING.stage
	return
}

func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) GetDB() db.DBInterface {
	return backRepoATTRIBUTE_DEFINITION_STRING.db
}

// GetATTRIBUTE_DEFINITION_STRINGDBFromATTRIBUTE_DEFINITION_STRINGPtr is a handy function to access the back repo instance from the stage instance
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) GetATTRIBUTE_DEFINITION_STRINGDBFromATTRIBUTE_DEFINITION_STRINGPtr(attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING) (attribute_definition_stringDB *ATTRIBUTE_DEFINITION_STRINGDB) {
	id := backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID[attribute_definition_string]
	attribute_definition_stringDB = backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB[id]
	return
}

// BackRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseOne commits all staged instances of ATTRIBUTE_DEFINITION_STRING to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var attribute_definition_strings []*models.ATTRIBUTE_DEFINITION_STRING
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		attribute_definition_strings = append(attribute_definition_strings, attribute_definition_string)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(attribute_definition_strings, func(i, j int) bool {
		return stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_strings[i]] < stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_strings[j]]
	})

	for _, attribute_definition_string := range attribute_definition_strings {
		backRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseOneInstance(attribute_definition_string)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, attribute_definition_string := range backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr {
		if _, ok := stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string]; !ok {
			backRepoATTRIBUTE_DEFINITION_STRING.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoATTRIBUTE_DEFINITION_STRING.CommitDeleteInstance commits deletion of ATTRIBUTE_DEFINITION_STRING to the BackRepo
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) CommitDeleteInstance(id uint) (Error error) {

	attribute_definition_string := backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr[id]

	// attribute_definition_string is not staged anymore, remove attribute_definition_stringDB
	attribute_definition_stringDB := backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB[id]
	db, _ := backRepoATTRIBUTE_DEFINITION_STRING.db.Unscoped()
	_, err := db.Delete(attribute_definition_stringDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID, attribute_definition_string)
	delete(backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr, id)
	delete(backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB, id)

	return
}

// BackRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseOneInstance commits attribute_definition_string staged instances of ATTRIBUTE_DEFINITION_STRING to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) CommitPhaseOneInstance(attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING) (Error error) {

	// check if the attribute_definition_string is not commited yet
	if _, ok := backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID[attribute_definition_string]; ok {
		return
	}

	// initiate attribute_definition_string
	var attribute_definition_stringDB ATTRIBUTE_DEFINITION_STRINGDB
	attribute_definition_stringDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRING(attribute_definition_string)

	_, err := backRepoATTRIBUTE_DEFINITION_STRING.db.Create(&attribute_definition_stringDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID[attribute_definition_string] = attribute_definition_stringDB.ID
	backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr[attribute_definition_stringDB.ID] = attribute_definition_string
	backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB[attribute_definition_stringDB.ID] = &attribute_definition_stringDB

	return
}

// BackRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseTwo commits all staged instances of ATTRIBUTE_DEFINITION_STRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, attribute_definition_string := range backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr {
		backRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseTwoInstance(backRepo, idx, attribute_definition_string)
	}

	return
}

// BackRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseTwoInstance commits {{structname }} of models.ATTRIBUTE_DEFINITION_STRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING) (Error error) {

	// fetch matching attribute_definition_stringDB
	if attribute_definition_stringDB, ok := backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB[idx]; ok {

		attribute_definition_stringDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRING(attribute_definition_string)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID = make([]int, 0)
		// 2. encode
		for _, alternative_idAssocEnd := range attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
			alternative_idAssocEnd_DB :=
				backRepo.BackRepoALTERNATIVE_ID.GetALTERNATIVE_IDDBFromALTERNATIVE_IDPtr(alternative_idAssocEnd)
			
			// the stage might be inconsistant, meaning that the alternative_idAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if alternative_idAssocEnd_DB == nil {
				continue
			}
			
			attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID =
				append(attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID, int(alternative_idAssocEnd_DB.ID))
		}

		// 1. reset
		attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING = make([]int, 0)
		// 2. encode
		for _, attribute_value_stringAssocEnd := range attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
			attribute_value_stringAssocEnd_DB :=
				backRepo.BackRepoATTRIBUTE_VALUE_STRING.GetATTRIBUTE_VALUE_STRINGDBFromATTRIBUTE_VALUE_STRINGPtr(attribute_value_stringAssocEnd)
			
			// the stage might be inconsistant, meaning that the attribute_value_stringAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if attribute_value_stringAssocEnd_DB == nil {
				continue
			}
			
			attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING =
				append(attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, int(attribute_value_stringAssocEnd_DB.ID))
		}

		_, err := backRepoATTRIBUTE_DEFINITION_STRING.db.Save(attribute_definition_stringDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown ATTRIBUTE_DEFINITION_STRING intance %s", attribute_definition_string.Name))
		return err
	}

	return
}

// BackRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) CheckoutPhaseOne() (Error error) {

	attribute_definition_stringDBArray := make([]ATTRIBUTE_DEFINITION_STRINGDB, 0)
	_, err := backRepoATTRIBUTE_DEFINITION_STRING.db.Find(&attribute_definition_stringDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	attribute_definition_stringInstancesToBeRemovedFromTheStage := make(map[*models.ATTRIBUTE_DEFINITION_STRING]any)
	for key, value := range backRepoATTRIBUTE_DEFINITION_STRING.stage.ATTRIBUTE_DEFINITION_STRINGs {
		attribute_definition_stringInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, attribute_definition_stringDB := range attribute_definition_stringDBArray {
		backRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseOneInstance(&attribute_definition_stringDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		attribute_definition_string, ok := backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr[attribute_definition_stringDB.ID]
		if ok {
			delete(attribute_definition_stringInstancesToBeRemovedFromTheStage, attribute_definition_string)
		}
	}

	// remove from stage and back repo's 3 maps all attribute_definition_strings that are not in the checkout
	for attribute_definition_string := range attribute_definition_stringInstancesToBeRemovedFromTheStage {
		attribute_definition_string.Unstage(backRepoATTRIBUTE_DEFINITION_STRING.GetStage())

		// remove instance from the back repo 3 maps
		attribute_definition_stringID := backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID[attribute_definition_string]
		delete(backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID, attribute_definition_string)
		delete(backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB, attribute_definition_stringID)
		delete(backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr, attribute_definition_stringID)
	}

	return
}

// CheckoutPhaseOneInstance takes a attribute_definition_stringDB that has been found in the DB, updates the backRepo and stages the
// models version of the attribute_definition_stringDB
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) CheckoutPhaseOneInstance(attribute_definition_stringDB *ATTRIBUTE_DEFINITION_STRINGDB) (Error error) {

	attribute_definition_string, ok := backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr[attribute_definition_stringDB.ID]
	if !ok {
		attribute_definition_string = new(models.ATTRIBUTE_DEFINITION_STRING)

		backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr[attribute_definition_stringDB.ID] = attribute_definition_string
		backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID[attribute_definition_string] = attribute_definition_stringDB.ID

		// append model store with the new element
		attribute_definition_string.Name = attribute_definition_stringDB.Name_Data.String
		attribute_definition_string.Stage(backRepoATTRIBUTE_DEFINITION_STRING.GetStage())
	}
	attribute_definition_stringDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING(attribute_definition_string)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	attribute_definition_string.Stage(backRepoATTRIBUTE_DEFINITION_STRING.GetStage())

	// preserve pointer to attribute_definition_stringDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB)[attribute_definition_stringDB hold variable pointers
	attribute_definition_stringDB_Data := *attribute_definition_stringDB
	preservedPtrToATTRIBUTE_DEFINITION_STRING := &attribute_definition_stringDB_Data
	backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB[attribute_definition_stringDB.ID] = preservedPtrToATTRIBUTE_DEFINITION_STRING

	return
}

// BackRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseTwo Checkouts all staged instances of ATTRIBUTE_DEFINITION_STRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, attribute_definition_stringDB := range backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB {
		backRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseTwoInstance(backRepo, attribute_definition_stringDB)
	}
	return
}

// BackRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseTwoInstance Checkouts staged instances of ATTRIBUTE_DEFINITION_STRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, attribute_definition_stringDB *ATTRIBUTE_DEFINITION_STRINGDB) (Error error) {

	attribute_definition_string := backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr[attribute_definition_stringDB.ID]

	attribute_definition_stringDB.DecodePointers(backRepo, attribute_definition_string)

	return
}

func (attribute_definition_stringDB *ATTRIBUTE_DEFINITION_STRINGDB) DecodePointers(backRepo *BackRepoStruct, attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING) {

	// insertion point for checkout of pointer encoding
	// This loop redeem attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID in the stage from the encode in the back repo
	// It parses all ALTERNATIVE_IDDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID = attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID[:0]
	for _, _ALTERNATIVE_IDid := range attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID {
		attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID = append(attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID, backRepo.BackRepoALTERNATIVE_ID.Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDPtr[uint(_ALTERNATIVE_IDid)])
	}

	// This loop redeem attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING in the stage from the encode in the back repo
	// It parses all ATTRIBUTE_VALUE_STRINGDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING = attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING[:0]
	for _, _ATTRIBUTE_VALUE_STRINGid := range attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
		attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING = append(attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, backRepo.BackRepoATTRIBUTE_VALUE_STRING.Map_ATTRIBUTE_VALUE_STRINGDBID_ATTRIBUTE_VALUE_STRINGPtr[uint(_ATTRIBUTE_VALUE_STRINGid)])
	}

	return
}

// CommitATTRIBUTE_DEFINITION_STRING allows commit of a single attribute_definition_string (if already staged)
func (backRepo *BackRepoStruct) CommitATTRIBUTE_DEFINITION_STRING(attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING) {
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseOneInstance(attribute_definition_string)
	if id, ok := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID[attribute_definition_string]; ok {
		backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseTwoInstance(backRepo, id, attribute_definition_string)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitATTRIBUTE_DEFINITION_STRING allows checkout of a single attribute_definition_string (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutATTRIBUTE_DEFINITION_STRING(attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING) {
	// check if the attribute_definition_string is staged
	if _, ok := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID[attribute_definition_string]; ok {

		if id, ok := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID[attribute_definition_string]; ok {
			var attribute_definition_stringDB ATTRIBUTE_DEFINITION_STRINGDB
			attribute_definition_stringDB.ID = id

			if _, err := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.db.First(&attribute_definition_stringDB, id); err != nil {
				log.Fatalln("CheckoutATTRIBUTE_DEFINITION_STRING : Problem with getting object with id:", id)
			}
			backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseOneInstance(&attribute_definition_stringDB)
			backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseTwoInstance(backRepo, &attribute_definition_stringDB)
		}
	}
}

// CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRING
func (attribute_definition_stringDB *ATTRIBUTE_DEFINITION_STRINGDB) CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRING(attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING) {
	// insertion point for fields commit

	attribute_definition_stringDB.Name_Data.String = attribute_definition_string.Name
	attribute_definition_stringDB.Name_Data.Valid = true

	attribute_definition_stringDB.DESC_Data.String = attribute_definition_string.DESC
	attribute_definition_stringDB.DESC_Data.Valid = true

	attribute_definition_stringDB.IS_EDITABLE_Data.Bool = attribute_definition_string.IS_EDITABLE
	attribute_definition_stringDB.IS_EDITABLE_Data.Valid = true

	attribute_definition_stringDB.LAST_CHANGE_Data.Time = attribute_definition_string.LAST_CHANGE
	attribute_definition_stringDB.LAST_CHANGE_Data.Valid = true

	attribute_definition_stringDB.LONG_NAME_Data.String = attribute_definition_string.LONG_NAME
	attribute_definition_stringDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRING_WOP
func (attribute_definition_stringDB *ATTRIBUTE_DEFINITION_STRINGDB) CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRING_WOP(attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING_WOP) {
	// insertion point for fields commit

	attribute_definition_stringDB.Name_Data.String = attribute_definition_string.Name
	attribute_definition_stringDB.Name_Data.Valid = true

	attribute_definition_stringDB.DESC_Data.String = attribute_definition_string.DESC
	attribute_definition_stringDB.DESC_Data.Valid = true

	attribute_definition_stringDB.IS_EDITABLE_Data.Bool = attribute_definition_string.IS_EDITABLE
	attribute_definition_stringDB.IS_EDITABLE_Data.Valid = true

	attribute_definition_stringDB.LAST_CHANGE_Data.Time = attribute_definition_string.LAST_CHANGE
	attribute_definition_stringDB.LAST_CHANGE_Data.Valid = true

	attribute_definition_stringDB.LONG_NAME_Data.String = attribute_definition_string.LONG_NAME
	attribute_definition_stringDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRINGWOP
func (attribute_definition_stringDB *ATTRIBUTE_DEFINITION_STRINGDB) CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRINGWOP(attribute_definition_string *ATTRIBUTE_DEFINITION_STRINGWOP) {
	// insertion point for fields commit

	attribute_definition_stringDB.Name_Data.String = attribute_definition_string.Name
	attribute_definition_stringDB.Name_Data.Valid = true

	attribute_definition_stringDB.DESC_Data.String = attribute_definition_string.DESC
	attribute_definition_stringDB.DESC_Data.Valid = true

	attribute_definition_stringDB.IS_EDITABLE_Data.Bool = attribute_definition_string.IS_EDITABLE
	attribute_definition_stringDB.IS_EDITABLE_Data.Valid = true

	attribute_definition_stringDB.LAST_CHANGE_Data.Time = attribute_definition_string.LAST_CHANGE
	attribute_definition_stringDB.LAST_CHANGE_Data.Valid = true

	attribute_definition_stringDB.LONG_NAME_Data.String = attribute_definition_string.LONG_NAME
	attribute_definition_stringDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING
func (attribute_definition_stringDB *ATTRIBUTE_DEFINITION_STRINGDB) CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING(attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING) {
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_definition_string.Name = attribute_definition_stringDB.Name_Data.String
	attribute_definition_string.DESC = attribute_definition_stringDB.DESC_Data.String
	attribute_definition_string.IS_EDITABLE = attribute_definition_stringDB.IS_EDITABLE_Data.Bool
	attribute_definition_string.LAST_CHANGE = attribute_definition_stringDB.LAST_CHANGE_Data.Time
	attribute_definition_string.LONG_NAME = attribute_definition_stringDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING_WOP
func (attribute_definition_stringDB *ATTRIBUTE_DEFINITION_STRINGDB) CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING_WOP(attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_definition_string.Name = attribute_definition_stringDB.Name_Data.String
	attribute_definition_string.DESC = attribute_definition_stringDB.DESC_Data.String
	attribute_definition_string.IS_EDITABLE = attribute_definition_stringDB.IS_EDITABLE_Data.Bool
	attribute_definition_string.LAST_CHANGE = attribute_definition_stringDB.LAST_CHANGE_Data.Time
	attribute_definition_string.LONG_NAME = attribute_definition_stringDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToATTRIBUTE_DEFINITION_STRINGWOP
func (attribute_definition_stringDB *ATTRIBUTE_DEFINITION_STRINGDB) CopyBasicFieldsToATTRIBUTE_DEFINITION_STRINGWOP(attribute_definition_string *ATTRIBUTE_DEFINITION_STRINGWOP) {
	attribute_definition_string.ID = int(attribute_definition_stringDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_definition_string.Name = attribute_definition_stringDB.Name_Data.String
	attribute_definition_string.DESC = attribute_definition_stringDB.DESC_Data.String
	attribute_definition_string.IS_EDITABLE = attribute_definition_stringDB.IS_EDITABLE_Data.Bool
	attribute_definition_string.LAST_CHANGE = attribute_definition_stringDB.LAST_CHANGE_Data.Time
	attribute_definition_string.LONG_NAME = attribute_definition_stringDB.LONG_NAME_Data.String
}

// Backup generates a json file from a slice of all ATTRIBUTE_DEFINITION_STRINGDB instances in the backrepo
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ATTRIBUTE_DEFINITION_STRINGDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTE_DEFINITION_STRINGDB, 0)
	for _, attribute_definition_stringDB := range backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB {
		forBackup = append(forBackup, attribute_definition_stringDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json ATTRIBUTE_DEFINITION_STRING ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json ATTRIBUTE_DEFINITION_STRING file", err.Error())
	}
}

// Backup generates a json file from a slice of all ATTRIBUTE_DEFINITION_STRINGDB instances in the backrepo
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTE_DEFINITION_STRINGDB, 0)
	for _, attribute_definition_stringDB := range backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB {
		forBackup = append(forBackup, attribute_definition_stringDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("ATTRIBUTE_DEFINITION_STRING")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&ATTRIBUTE_DEFINITION_STRING_Fields, -1)
	for _, attribute_definition_stringDB := range forBackup {

		var attribute_definition_stringWOP ATTRIBUTE_DEFINITION_STRINGWOP
		attribute_definition_stringDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_STRINGWOP(&attribute_definition_stringWOP)

		row := sh.AddRow()
		row.WriteStruct(&attribute_definition_stringWOP, -1)
	}
}

// RestoreXL from the "ATTRIBUTE_DEFINITION_STRING" sheet all ATTRIBUTE_DEFINITION_STRINGDB instances
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoATTRIBUTE_DEFINITION_STRINGid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["ATTRIBUTE_DEFINITION_STRING"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoATTRIBUTE_DEFINITION_STRING.rowVisitorATTRIBUTE_DEFINITION_STRING)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) rowVisitorATTRIBUTE_DEFINITION_STRING(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var attribute_definition_stringWOP ATTRIBUTE_DEFINITION_STRINGWOP
		row.ReadStruct(&attribute_definition_stringWOP)

		// add the unmarshalled struct to the stage
		attribute_definition_stringDB := new(ATTRIBUTE_DEFINITION_STRINGDB)
		attribute_definition_stringDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRINGWOP(&attribute_definition_stringWOP)

		attribute_definition_stringDB_ID_atBackupTime := attribute_definition_stringDB.ID
		attribute_definition_stringDB.ID = 0
		_, err := backRepoATTRIBUTE_DEFINITION_STRING.db.Create(attribute_definition_stringDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB[attribute_definition_stringDB.ID] = attribute_definition_stringDB
		BackRepoATTRIBUTE_DEFINITION_STRINGid_atBckpTime_newID[attribute_definition_stringDB_ID_atBackupTime] = attribute_definition_stringDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ATTRIBUTE_DEFINITION_STRINGDB.json" in dirPath that stores an array
// of ATTRIBUTE_DEFINITION_STRINGDB and stores it in the database
// the map BackRepoATTRIBUTE_DEFINITION_STRINGid_atBckpTime_newID is updated accordingly
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoATTRIBUTE_DEFINITION_STRINGid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ATTRIBUTE_DEFINITION_STRINGDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json ATTRIBUTE_DEFINITION_STRING file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ATTRIBUTE_DEFINITION_STRINGDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB
	for _, attribute_definition_stringDB := range forRestore {

		attribute_definition_stringDB_ID_atBackupTime := attribute_definition_stringDB.ID
		attribute_definition_stringDB.ID = 0
		_, err := backRepoATTRIBUTE_DEFINITION_STRING.db.Create(attribute_definition_stringDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB[attribute_definition_stringDB.ID] = attribute_definition_stringDB
		BackRepoATTRIBUTE_DEFINITION_STRINGid_atBckpTime_newID[attribute_definition_stringDB_ID_atBackupTime] = attribute_definition_stringDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json ATTRIBUTE_DEFINITION_STRING file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<ATTRIBUTE_DEFINITION_STRING>id_atBckpTime_newID
// to compute new index
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) RestorePhaseTwo() {

	for _, attribute_definition_stringDB := range backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB {

		// next line of code is to avert unused variable compilation error
		_ = attribute_definition_stringDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoATTRIBUTE_DEFINITION_STRING.db.Model(attribute_definition_stringDB)
		_, err := db.Updates(*attribute_definition_stringDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoATTRIBUTE_DEFINITION_STRING.ResetReversePointers commits all staged instances of ATTRIBUTE_DEFINITION_STRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, attribute_definition_string := range backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr {
		backRepoATTRIBUTE_DEFINITION_STRING.ResetReversePointersInstance(backRepo, idx, attribute_definition_string)
	}

	return
}

func (backRepoATTRIBUTE_DEFINITION_STRING *BackRepoATTRIBUTE_DEFINITION_STRINGStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING) (Error error) {

	// fetch matching attribute_definition_stringDB
	if attribute_definition_stringDB, ok := backRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB[idx]; ok {
		_ = attribute_definition_stringDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoATTRIBUTE_DEFINITION_STRINGid_atBckpTime_newID map[uint]uint
