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
var dummy_ATTRIBUTE_DEFINITION_INTEGER_sql sql.NullBool
var dummy_ATTRIBUTE_DEFINITION_INTEGER_time time.Duration
var dummy_ATTRIBUTE_DEFINITION_INTEGER_sort sort.Float64Slice

// ATTRIBUTE_DEFINITION_INTEGERAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model attribute_definition_integerAPI
type ATTRIBUTE_DEFINITION_INTEGERAPI struct {
	gorm.Model

	models.ATTRIBUTE_DEFINITION_INTEGER_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ATTRIBUTE_DEFINITION_INTEGERPointersEncoding ATTRIBUTE_DEFINITION_INTEGERPointersEncoding
}

// ATTRIBUTE_DEFINITION_INTEGERPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ATTRIBUTE_DEFINITION_INTEGERPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	ALTERNATIVE_ID struct {

		// field ALTERNATIVE_ID is a slice of pointers to another Struct (optional or 0..1)
		ALTERNATIVE_ID IntSlice `gorm:"type:TEXT"`

	} `gorm:"embedded"`

	DEFAULT_VALUE struct {

		// field ATTRIBUTE_VALUE_INTEGER is a slice of pointers to another Struct (optional or 0..1)
		ATTRIBUTE_VALUE_INTEGER IntSlice `gorm:"type:TEXT"`

	} `gorm:"embedded"`
}

// ATTRIBUTE_DEFINITION_INTEGERDB describes a attribute_definition_integer in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model attribute_definition_integerDB
type ATTRIBUTE_DEFINITION_INTEGERDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field attribute_definition_integerDB.Name
	Name_Data sql.NullString

	// Declation for basic field attribute_definition_integerDB.DESC
	DESC_Data sql.NullString

	// Declation for basic field attribute_definition_integerDB.IS_EDITABLE
	// provide the sql storage for the boolan
	IS_EDITABLE_Data sql.NullBool

	// Declation for basic field attribute_definition_integerDB.LAST_CHANGE
	LAST_CHANGE_Data sql.NullTime

	// Declation for basic field attribute_definition_integerDB.LONG_NAME
	LONG_NAME_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ATTRIBUTE_DEFINITION_INTEGERPointersEncoding
}

// ATTRIBUTE_DEFINITION_INTEGERDBs arrays attribute_definition_integerDBs
// swagger:response attribute_definition_integerDBsResponse
type ATTRIBUTE_DEFINITION_INTEGERDBs []ATTRIBUTE_DEFINITION_INTEGERDB

// ATTRIBUTE_DEFINITION_INTEGERDBResponse provides response
// swagger:response attribute_definition_integerDBResponse
type ATTRIBUTE_DEFINITION_INTEGERDBResponse struct {
	ATTRIBUTE_DEFINITION_INTEGERDB
}

// ATTRIBUTE_DEFINITION_INTEGERWOP is a ATTRIBUTE_DEFINITION_INTEGER without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ATTRIBUTE_DEFINITION_INTEGERWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESC string `xlsx:"2"`

	IS_EDITABLE bool `xlsx:"3"`

	LAST_CHANGE time.Time `xlsx:"4"`

	LONG_NAME string `xlsx:"5"`
	// insertion for WOP pointer fields
}

var ATTRIBUTE_DEFINITION_INTEGER_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESC",
	"IS_EDITABLE",
	"LAST_CHANGE",
	"LONG_NAME",
}

type BackRepoATTRIBUTE_DEFINITION_INTEGERStruct struct {
	// stores ATTRIBUTE_DEFINITION_INTEGERDB according to their gorm ID
	Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB map[uint]*ATTRIBUTE_DEFINITION_INTEGERDB

	// stores ATTRIBUTE_DEFINITION_INTEGERDB ID according to ATTRIBUTE_DEFINITION_INTEGER address
	Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID map[*models.ATTRIBUTE_DEFINITION_INTEGER]uint

	// stores ATTRIBUTE_DEFINITION_INTEGER according to their gorm ID
	Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr map[uint]*models.ATTRIBUTE_DEFINITION_INTEGER

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoATTRIBUTE_DEFINITION_INTEGER.stage
	return
}

func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) GetDB() *gorm.DB {
	return backRepoATTRIBUTE_DEFINITION_INTEGER.db
}

// GetATTRIBUTE_DEFINITION_INTEGERDBFromATTRIBUTE_DEFINITION_INTEGERPtr is a handy function to access the back repo instance from the stage instance
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) GetATTRIBUTE_DEFINITION_INTEGERDBFromATTRIBUTE_DEFINITION_INTEGERPtr(attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER) (attribute_definition_integerDB *ATTRIBUTE_DEFINITION_INTEGERDB) {
	id := backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID[attribute_definition_integer]
	attribute_definition_integerDB = backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB[id]
	return
}

// BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseOne commits all staged instances of ATTRIBUTE_DEFINITION_INTEGER to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		backRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseOneInstance(attribute_definition_integer)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, attribute_definition_integer := range backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr {
		if _, ok := stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer]; !ok {
			backRepoATTRIBUTE_DEFINITION_INTEGER.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitDeleteInstance commits deletion of ATTRIBUTE_DEFINITION_INTEGER to the BackRepo
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) CommitDeleteInstance(id uint) (Error error) {

	attribute_definition_integer := backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr[id]

	// attribute_definition_integer is not staged anymore, remove attribute_definition_integerDB
	attribute_definition_integerDB := backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB[id]
	query := backRepoATTRIBUTE_DEFINITION_INTEGER.db.Unscoped().Delete(&attribute_definition_integerDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID, attribute_definition_integer)
	delete(backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr, id)
	delete(backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB, id)

	return
}

// BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseOneInstance commits attribute_definition_integer staged instances of ATTRIBUTE_DEFINITION_INTEGER to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) CommitPhaseOneInstance(attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER) (Error error) {

	// check if the attribute_definition_integer is not commited yet
	if _, ok := backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID[attribute_definition_integer]; ok {
		return
	}

	// initiate attribute_definition_integer
	var attribute_definition_integerDB ATTRIBUTE_DEFINITION_INTEGERDB
	attribute_definition_integerDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer)

	query := backRepoATTRIBUTE_DEFINITION_INTEGER.db.Create(&attribute_definition_integerDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID[attribute_definition_integer] = attribute_definition_integerDB.ID
	backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr[attribute_definition_integerDB.ID] = attribute_definition_integer
	backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB[attribute_definition_integerDB.ID] = &attribute_definition_integerDB

	return
}

// BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseTwo commits all staged instances of ATTRIBUTE_DEFINITION_INTEGER to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, attribute_definition_integer := range backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr {
		backRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseTwoInstance(backRepo, idx, attribute_definition_integer)
	}

	return
}

// BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseTwoInstance commits {{structname }} of models.ATTRIBUTE_DEFINITION_INTEGER to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER) (Error error) {

	// fetch matching attribute_definition_integerDB
	if attribute_definition_integerDB, ok := backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB[idx]; ok {

		attribute_definition_integerDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID = make([]int, 0)
		// 2. encode
		for _, alternative_idAssocEnd := range attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
			alternative_idAssocEnd_DB :=
				backRepo.BackRepoALTERNATIVE_ID.GetALTERNATIVE_IDDBFromALTERNATIVE_IDPtr(alternative_idAssocEnd)
			
			// the stage might be inconsistant, meaning that the alternative_idAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if alternative_idAssocEnd_DB == nil {
				continue
			}
			
			attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID =
				append(attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID, int(alternative_idAssocEnd_DB.ID))
		}

		// 1. reset
		attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER = make([]int, 0)
		// 2. encode
		for _, attribute_value_integerAssocEnd := range attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
			attribute_value_integerAssocEnd_DB :=
				backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.GetATTRIBUTE_VALUE_INTEGERDBFromATTRIBUTE_VALUE_INTEGERPtr(attribute_value_integerAssocEnd)
			
			// the stage might be inconsistant, meaning that the attribute_value_integerAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if attribute_value_integerAssocEnd_DB == nil {
				continue
			}
			
			attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER =
				append(attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, int(attribute_value_integerAssocEnd_DB.ID))
		}

		query := backRepoATTRIBUTE_DEFINITION_INTEGER.db.Save(&attribute_definition_integerDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown ATTRIBUTE_DEFINITION_INTEGER intance %s", attribute_definition_integer.Name))
		return err
	}

	return
}

// BackRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) CheckoutPhaseOne() (Error error) {

	attribute_definition_integerDBArray := make([]ATTRIBUTE_DEFINITION_INTEGERDB, 0)
	query := backRepoATTRIBUTE_DEFINITION_INTEGER.db.Find(&attribute_definition_integerDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	attribute_definition_integerInstancesToBeRemovedFromTheStage := make(map[*models.ATTRIBUTE_DEFINITION_INTEGER]any)
	for key, value := range backRepoATTRIBUTE_DEFINITION_INTEGER.stage.ATTRIBUTE_DEFINITION_INTEGERs {
		attribute_definition_integerInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, attribute_definition_integerDB := range attribute_definition_integerDBArray {
		backRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseOneInstance(&attribute_definition_integerDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		attribute_definition_integer, ok := backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr[attribute_definition_integerDB.ID]
		if ok {
			delete(attribute_definition_integerInstancesToBeRemovedFromTheStage, attribute_definition_integer)
		}
	}

	// remove from stage and back repo's 3 maps all attribute_definition_integers that are not in the checkout
	for attribute_definition_integer := range attribute_definition_integerInstancesToBeRemovedFromTheStage {
		attribute_definition_integer.Unstage(backRepoATTRIBUTE_DEFINITION_INTEGER.GetStage())

		// remove instance from the back repo 3 maps
		attribute_definition_integerID := backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID[attribute_definition_integer]
		delete(backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID, attribute_definition_integer)
		delete(backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB, attribute_definition_integerID)
		delete(backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr, attribute_definition_integerID)
	}

	return
}

// CheckoutPhaseOneInstance takes a attribute_definition_integerDB that has been found in the DB, updates the backRepo and stages the
// models version of the attribute_definition_integerDB
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) CheckoutPhaseOneInstance(attribute_definition_integerDB *ATTRIBUTE_DEFINITION_INTEGERDB) (Error error) {

	attribute_definition_integer, ok := backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr[attribute_definition_integerDB.ID]
	if !ok {
		attribute_definition_integer = new(models.ATTRIBUTE_DEFINITION_INTEGER)

		backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr[attribute_definition_integerDB.ID] = attribute_definition_integer
		backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID[attribute_definition_integer] = attribute_definition_integerDB.ID

		// append model store with the new element
		attribute_definition_integer.Name = attribute_definition_integerDB.Name_Data.String
		attribute_definition_integer.Stage(backRepoATTRIBUTE_DEFINITION_INTEGER.GetStage())
	}
	attribute_definition_integerDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	attribute_definition_integer.Stage(backRepoATTRIBUTE_DEFINITION_INTEGER.GetStage())

	// preserve pointer to attribute_definition_integerDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB)[attribute_definition_integerDB hold variable pointers
	attribute_definition_integerDB_Data := *attribute_definition_integerDB
	preservedPtrToATTRIBUTE_DEFINITION_INTEGER := &attribute_definition_integerDB_Data
	backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB[attribute_definition_integerDB.ID] = preservedPtrToATTRIBUTE_DEFINITION_INTEGER

	return
}

// BackRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseTwo Checkouts all staged instances of ATTRIBUTE_DEFINITION_INTEGER to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, attribute_definition_integerDB := range backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB {
		backRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseTwoInstance(backRepo, attribute_definition_integerDB)
	}
	return
}

// BackRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseTwoInstance Checkouts staged instances of ATTRIBUTE_DEFINITION_INTEGER to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, attribute_definition_integerDB *ATTRIBUTE_DEFINITION_INTEGERDB) (Error error) {

	attribute_definition_integer := backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr[attribute_definition_integerDB.ID]

	attribute_definition_integerDB.DecodePointers(backRepo, attribute_definition_integer)

	return
}

func (attribute_definition_integerDB *ATTRIBUTE_DEFINITION_INTEGERDB) DecodePointers(backRepo *BackRepoStruct, attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER) {

	// insertion point for checkout of pointer encoding
	// This loop redeem attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID in the stage from the encode in the back repo
	// It parses all ALTERNATIVE_IDDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID = attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID[:0]
	for _, _ALTERNATIVE_IDid := range attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID {
		attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID = append(attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID, backRepo.BackRepoALTERNATIVE_ID.Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDPtr[uint(_ALTERNATIVE_IDid)])
	}

	// This loop redeem attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER in the stage from the encode in the back repo
	// It parses all ATTRIBUTE_VALUE_INTEGERDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER = attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER[:0]
	for _, _ATTRIBUTE_VALUE_INTEGERid := range attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
		attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER = append(attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.Map_ATTRIBUTE_VALUE_INTEGERDBID_ATTRIBUTE_VALUE_INTEGERPtr[uint(_ATTRIBUTE_VALUE_INTEGERid)])
	}

	return
}

// CommitATTRIBUTE_DEFINITION_INTEGER allows commit of a single attribute_definition_integer (if already staged)
func (backRepo *BackRepoStruct) CommitATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER) {
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseOneInstance(attribute_definition_integer)
	if id, ok := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID[attribute_definition_integer]; ok {
		backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseTwoInstance(backRepo, id, attribute_definition_integer)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitATTRIBUTE_DEFINITION_INTEGER allows checkout of a single attribute_definition_integer (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER) {
	// check if the attribute_definition_integer is staged
	if _, ok := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID[attribute_definition_integer]; ok {

		if id, ok := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID[attribute_definition_integer]; ok {
			var attribute_definition_integerDB ATTRIBUTE_DEFINITION_INTEGERDB
			attribute_definition_integerDB.ID = id

			if err := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.db.First(&attribute_definition_integerDB, id).Error; err != nil {
				log.Fatalln("CheckoutATTRIBUTE_DEFINITION_INTEGER : Problem with getting object with id:", id)
			}
			backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseOneInstance(&attribute_definition_integerDB)
			backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseTwoInstance(backRepo, &attribute_definition_integerDB)
		}
	}
}

// CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGER
func (attribute_definition_integerDB *ATTRIBUTE_DEFINITION_INTEGERDB) CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER) {
	// insertion point for fields commit

	attribute_definition_integerDB.Name_Data.String = attribute_definition_integer.Name
	attribute_definition_integerDB.Name_Data.Valid = true

	attribute_definition_integerDB.DESC_Data.String = attribute_definition_integer.DESC
	attribute_definition_integerDB.DESC_Data.Valid = true

	attribute_definition_integerDB.IS_EDITABLE_Data.Bool = attribute_definition_integer.IS_EDITABLE
	attribute_definition_integerDB.IS_EDITABLE_Data.Valid = true

	attribute_definition_integerDB.LAST_CHANGE_Data.Time = attribute_definition_integer.LAST_CHANGE
	attribute_definition_integerDB.LAST_CHANGE_Data.Valid = true

	attribute_definition_integerDB.LONG_NAME_Data.String = attribute_definition_integer.LONG_NAME
	attribute_definition_integerDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGER_WOP
func (attribute_definition_integerDB *ATTRIBUTE_DEFINITION_INTEGERDB) CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGER_WOP(attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER_WOP) {
	// insertion point for fields commit

	attribute_definition_integerDB.Name_Data.String = attribute_definition_integer.Name
	attribute_definition_integerDB.Name_Data.Valid = true

	attribute_definition_integerDB.DESC_Data.String = attribute_definition_integer.DESC
	attribute_definition_integerDB.DESC_Data.Valid = true

	attribute_definition_integerDB.IS_EDITABLE_Data.Bool = attribute_definition_integer.IS_EDITABLE
	attribute_definition_integerDB.IS_EDITABLE_Data.Valid = true

	attribute_definition_integerDB.LAST_CHANGE_Data.Time = attribute_definition_integer.LAST_CHANGE
	attribute_definition_integerDB.LAST_CHANGE_Data.Valid = true

	attribute_definition_integerDB.LONG_NAME_Data.String = attribute_definition_integer.LONG_NAME
	attribute_definition_integerDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGERWOP
func (attribute_definition_integerDB *ATTRIBUTE_DEFINITION_INTEGERDB) CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGERWOP(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGERWOP) {
	// insertion point for fields commit

	attribute_definition_integerDB.Name_Data.String = attribute_definition_integer.Name
	attribute_definition_integerDB.Name_Data.Valid = true

	attribute_definition_integerDB.DESC_Data.String = attribute_definition_integer.DESC
	attribute_definition_integerDB.DESC_Data.Valid = true

	attribute_definition_integerDB.IS_EDITABLE_Data.Bool = attribute_definition_integer.IS_EDITABLE
	attribute_definition_integerDB.IS_EDITABLE_Data.Valid = true

	attribute_definition_integerDB.LAST_CHANGE_Data.Time = attribute_definition_integer.LAST_CHANGE
	attribute_definition_integerDB.LAST_CHANGE_Data.Valid = true

	attribute_definition_integerDB.LONG_NAME_Data.String = attribute_definition_integer.LONG_NAME
	attribute_definition_integerDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER
func (attribute_definition_integerDB *ATTRIBUTE_DEFINITION_INTEGERDB) CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER) {
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_definition_integer.Name = attribute_definition_integerDB.Name_Data.String
	attribute_definition_integer.DESC = attribute_definition_integerDB.DESC_Data.String
	attribute_definition_integer.IS_EDITABLE = attribute_definition_integerDB.IS_EDITABLE_Data.Bool
	attribute_definition_integer.LAST_CHANGE = attribute_definition_integerDB.LAST_CHANGE_Data.Time
	attribute_definition_integer.LONG_NAME = attribute_definition_integerDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER_WOP
func (attribute_definition_integerDB *ATTRIBUTE_DEFINITION_INTEGERDB) CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER_WOP(attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_definition_integer.Name = attribute_definition_integerDB.Name_Data.String
	attribute_definition_integer.DESC = attribute_definition_integerDB.DESC_Data.String
	attribute_definition_integer.IS_EDITABLE = attribute_definition_integerDB.IS_EDITABLE_Data.Bool
	attribute_definition_integer.LAST_CHANGE = attribute_definition_integerDB.LAST_CHANGE_Data.Time
	attribute_definition_integer.LONG_NAME = attribute_definition_integerDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGERWOP
func (attribute_definition_integerDB *ATTRIBUTE_DEFINITION_INTEGERDB) CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGERWOP(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGERWOP) {
	attribute_definition_integer.ID = int(attribute_definition_integerDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_definition_integer.Name = attribute_definition_integerDB.Name_Data.String
	attribute_definition_integer.DESC = attribute_definition_integerDB.DESC_Data.String
	attribute_definition_integer.IS_EDITABLE = attribute_definition_integerDB.IS_EDITABLE_Data.Bool
	attribute_definition_integer.LAST_CHANGE = attribute_definition_integerDB.LAST_CHANGE_Data.Time
	attribute_definition_integer.LONG_NAME = attribute_definition_integerDB.LONG_NAME_Data.String
}

// Backup generates a json file from a slice of all ATTRIBUTE_DEFINITION_INTEGERDB instances in the backrepo
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ATTRIBUTE_DEFINITION_INTEGERDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTE_DEFINITION_INTEGERDB, 0)
	for _, attribute_definition_integerDB := range backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB {
		forBackup = append(forBackup, attribute_definition_integerDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json ATTRIBUTE_DEFINITION_INTEGER ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json ATTRIBUTE_DEFINITION_INTEGER file", err.Error())
	}
}

// Backup generates a json file from a slice of all ATTRIBUTE_DEFINITION_INTEGERDB instances in the backrepo
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTE_DEFINITION_INTEGERDB, 0)
	for _, attribute_definition_integerDB := range backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB {
		forBackup = append(forBackup, attribute_definition_integerDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("ATTRIBUTE_DEFINITION_INTEGER")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&ATTRIBUTE_DEFINITION_INTEGER_Fields, -1)
	for _, attribute_definition_integerDB := range forBackup {

		var attribute_definition_integerWOP ATTRIBUTE_DEFINITION_INTEGERWOP
		attribute_definition_integerDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGERWOP(&attribute_definition_integerWOP)

		row := sh.AddRow()
		row.WriteStruct(&attribute_definition_integerWOP, -1)
	}
}

// RestoreXL from the "ATTRIBUTE_DEFINITION_INTEGER" sheet all ATTRIBUTE_DEFINITION_INTEGERDB instances
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoATTRIBUTE_DEFINITION_INTEGERid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["ATTRIBUTE_DEFINITION_INTEGER"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoATTRIBUTE_DEFINITION_INTEGER.rowVisitorATTRIBUTE_DEFINITION_INTEGER)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) rowVisitorATTRIBUTE_DEFINITION_INTEGER(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var attribute_definition_integerWOP ATTRIBUTE_DEFINITION_INTEGERWOP
		row.ReadStruct(&attribute_definition_integerWOP)

		// add the unmarshalled struct to the stage
		attribute_definition_integerDB := new(ATTRIBUTE_DEFINITION_INTEGERDB)
		attribute_definition_integerDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGERWOP(&attribute_definition_integerWOP)

		attribute_definition_integerDB_ID_atBackupTime := attribute_definition_integerDB.ID
		attribute_definition_integerDB.ID = 0
		query := backRepoATTRIBUTE_DEFINITION_INTEGER.db.Create(attribute_definition_integerDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB[attribute_definition_integerDB.ID] = attribute_definition_integerDB
		BackRepoATTRIBUTE_DEFINITION_INTEGERid_atBckpTime_newID[attribute_definition_integerDB_ID_atBackupTime] = attribute_definition_integerDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ATTRIBUTE_DEFINITION_INTEGERDB.json" in dirPath that stores an array
// of ATTRIBUTE_DEFINITION_INTEGERDB and stores it in the database
// the map BackRepoATTRIBUTE_DEFINITION_INTEGERid_atBckpTime_newID is updated accordingly
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoATTRIBUTE_DEFINITION_INTEGERid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ATTRIBUTE_DEFINITION_INTEGERDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json ATTRIBUTE_DEFINITION_INTEGER file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ATTRIBUTE_DEFINITION_INTEGERDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB
	for _, attribute_definition_integerDB := range forRestore {

		attribute_definition_integerDB_ID_atBackupTime := attribute_definition_integerDB.ID
		attribute_definition_integerDB.ID = 0
		query := backRepoATTRIBUTE_DEFINITION_INTEGER.db.Create(attribute_definition_integerDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB[attribute_definition_integerDB.ID] = attribute_definition_integerDB
		BackRepoATTRIBUTE_DEFINITION_INTEGERid_atBckpTime_newID[attribute_definition_integerDB_ID_atBackupTime] = attribute_definition_integerDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json ATTRIBUTE_DEFINITION_INTEGER file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<ATTRIBUTE_DEFINITION_INTEGER>id_atBckpTime_newID
// to compute new index
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) RestorePhaseTwo() {

	for _, attribute_definition_integerDB := range backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB {

		// next line of code is to avert unused variable compilation error
		_ = attribute_definition_integerDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoATTRIBUTE_DEFINITION_INTEGER.db.Model(attribute_definition_integerDB).Updates(*attribute_definition_integerDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoATTRIBUTE_DEFINITION_INTEGER.ResetReversePointers commits all staged instances of ATTRIBUTE_DEFINITION_INTEGER to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, attribute_definition_integer := range backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr {
		backRepoATTRIBUTE_DEFINITION_INTEGER.ResetReversePointersInstance(backRepo, idx, attribute_definition_integer)
	}

	return
}

func (backRepoATTRIBUTE_DEFINITION_INTEGER *BackRepoATTRIBUTE_DEFINITION_INTEGERStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER) (Error error) {

	// fetch matching attribute_definition_integerDB
	if attribute_definition_integerDB, ok := backRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB[idx]; ok {
		_ = attribute_definition_integerDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoATTRIBUTE_DEFINITION_INTEGERid_atBckpTime_newID map[uint]uint