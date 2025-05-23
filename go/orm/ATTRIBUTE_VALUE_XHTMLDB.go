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
var dummy_ATTRIBUTE_VALUE_XHTML_sql sql.NullBool
var dummy_ATTRIBUTE_VALUE_XHTML_time time.Duration
var dummy_ATTRIBUTE_VALUE_XHTML_sort sort.Float64Slice

// ATTRIBUTE_VALUE_XHTMLAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model attribute_value_xhtmlAPI
type ATTRIBUTE_VALUE_XHTMLAPI struct {
	gorm.Model

	models.ATTRIBUTE_VALUE_XHTML_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ATTRIBUTE_VALUE_XHTMLPointersEncoding ATTRIBUTE_VALUE_XHTMLPointersEncoding
}

// ATTRIBUTE_VALUE_XHTMLPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ATTRIBUTE_VALUE_XHTMLPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field THE_VALUE is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	THE_VALUEID sql.NullInt64

	// field THE_ORIGINAL_VALUE is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	THE_ORIGINAL_VALUEID sql.NullInt64

	// field DEFINITION is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	DEFINITIONID sql.NullInt64
}

// ATTRIBUTE_VALUE_XHTMLDB describes a attribute_value_xhtml in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model attribute_value_xhtmlDB
type ATTRIBUTE_VALUE_XHTMLDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field attribute_value_xhtmlDB.Name
	Name_Data sql.NullString

	// Declation for basic field attribute_value_xhtmlDB.IS_SIMPLIFIED
	// provide the sql storage for the boolan
	IS_SIMPLIFIED_Data sql.NullBool

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ATTRIBUTE_VALUE_XHTMLPointersEncoding
}

// ATTRIBUTE_VALUE_XHTMLDBs arrays attribute_value_xhtmlDBs
// swagger:response attribute_value_xhtmlDBsResponse
type ATTRIBUTE_VALUE_XHTMLDBs []ATTRIBUTE_VALUE_XHTMLDB

// ATTRIBUTE_VALUE_XHTMLDBResponse provides response
// swagger:response attribute_value_xhtmlDBResponse
type ATTRIBUTE_VALUE_XHTMLDBResponse struct {
	ATTRIBUTE_VALUE_XHTMLDB
}

// ATTRIBUTE_VALUE_XHTMLWOP is a ATTRIBUTE_VALUE_XHTML without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ATTRIBUTE_VALUE_XHTMLWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IS_SIMPLIFIED bool `xlsx:"2"`
	// insertion for WOP pointer fields
}

var ATTRIBUTE_VALUE_XHTML_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IS_SIMPLIFIED",
}

type BackRepoATTRIBUTE_VALUE_XHTMLStruct struct {
	// stores ATTRIBUTE_VALUE_XHTMLDB according to their gorm ID
	Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB map[uint]*ATTRIBUTE_VALUE_XHTMLDB

	// stores ATTRIBUTE_VALUE_XHTMLDB ID according to ATTRIBUTE_VALUE_XHTML address
	Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID map[*models.ATTRIBUTE_VALUE_XHTML]uint

	// stores ATTRIBUTE_VALUE_XHTML according to their gorm ID
	Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr map[uint]*models.ATTRIBUTE_VALUE_XHTML

	db db.DBInterface

	stage *models.Stage
}

func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) GetStage() (stage *models.Stage) {
	stage = backRepoATTRIBUTE_VALUE_XHTML.stage
	return
}

func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) GetDB() db.DBInterface {
	return backRepoATTRIBUTE_VALUE_XHTML.db
}

// GetATTRIBUTE_VALUE_XHTMLDBFromATTRIBUTE_VALUE_XHTMLPtr is a handy function to access the back repo instance from the stage instance
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) GetATTRIBUTE_VALUE_XHTMLDBFromATTRIBUTE_VALUE_XHTMLPtr(attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML) (attribute_value_xhtmlDB *ATTRIBUTE_VALUE_XHTMLDB) {
	id := backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID[attribute_value_xhtml]
	attribute_value_xhtmlDB = backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB[id]
	return
}

// BackRepoATTRIBUTE_VALUE_XHTML.CommitPhaseOne commits all staged instances of ATTRIBUTE_VALUE_XHTML to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var attribute_value_xhtmls []*models.ATTRIBUTE_VALUE_XHTML
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		attribute_value_xhtmls = append(attribute_value_xhtmls, attribute_value_xhtml)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(attribute_value_xhtmls, func(i, j int) bool {
		return stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtmls[i]] < stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtmls[j]]
	})

	for _, attribute_value_xhtml := range attribute_value_xhtmls {
		backRepoATTRIBUTE_VALUE_XHTML.CommitPhaseOneInstance(attribute_value_xhtml)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, attribute_value_xhtml := range backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr {
		if _, ok := stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml]; !ok {
			backRepoATTRIBUTE_VALUE_XHTML.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoATTRIBUTE_VALUE_XHTML.CommitDeleteInstance commits deletion of ATTRIBUTE_VALUE_XHTML to the BackRepo
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) CommitDeleteInstance(id uint) (Error error) {

	attribute_value_xhtml := backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr[id]

	// attribute_value_xhtml is not staged anymore, remove attribute_value_xhtmlDB
	attribute_value_xhtmlDB := backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB[id]
	db, _ := backRepoATTRIBUTE_VALUE_XHTML.db.Unscoped()
	_, err := db.Delete(attribute_value_xhtmlDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID, attribute_value_xhtml)
	delete(backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr, id)
	delete(backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB, id)

	return
}

// BackRepoATTRIBUTE_VALUE_XHTML.CommitPhaseOneInstance commits attribute_value_xhtml staged instances of ATTRIBUTE_VALUE_XHTML to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) CommitPhaseOneInstance(attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML) (Error error) {

	// check if the attribute_value_xhtml is not commited yet
	if _, ok := backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID[attribute_value_xhtml]; ok {
		return
	}

	// initiate attribute_value_xhtml
	var attribute_value_xhtmlDB ATTRIBUTE_VALUE_XHTMLDB
	attribute_value_xhtmlDB.CopyBasicFieldsFromATTRIBUTE_VALUE_XHTML(attribute_value_xhtml)

	_, err := backRepoATTRIBUTE_VALUE_XHTML.db.Create(&attribute_value_xhtmlDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID[attribute_value_xhtml] = attribute_value_xhtmlDB.ID
	backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr[attribute_value_xhtmlDB.ID] = attribute_value_xhtml
	backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB[attribute_value_xhtmlDB.ID] = &attribute_value_xhtmlDB

	return
}

// BackRepoATTRIBUTE_VALUE_XHTML.CommitPhaseTwo commits all staged instances of ATTRIBUTE_VALUE_XHTML to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, attribute_value_xhtml := range backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr {
		backRepoATTRIBUTE_VALUE_XHTML.CommitPhaseTwoInstance(backRepo, idx, attribute_value_xhtml)
	}

	return
}

// BackRepoATTRIBUTE_VALUE_XHTML.CommitPhaseTwoInstance commits {{structname }} of models.ATTRIBUTE_VALUE_XHTML to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML) (Error error) {

	// fetch matching attribute_value_xhtmlDB
	if attribute_value_xhtmlDB, ok := backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB[idx]; ok {

		attribute_value_xhtmlDB.CopyBasicFieldsFromATTRIBUTE_VALUE_XHTML(attribute_value_xhtml)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value attribute_value_xhtml.THE_VALUE translates to updating the attribute_value_xhtml.THE_VALUEID
		attribute_value_xhtmlDB.THE_VALUEID.Valid = true // allow for a 0 value (nil association)
		if attribute_value_xhtml.THE_VALUE != nil {
			if THE_VALUEId, ok := backRepo.BackRepoXHTML_CONTENT.Map_XHTML_CONTENTPtr_XHTML_CONTENTDBID[attribute_value_xhtml.THE_VALUE]; ok {
				attribute_value_xhtmlDB.THE_VALUEID.Int64 = int64(THE_VALUEId)
				attribute_value_xhtmlDB.THE_VALUEID.Valid = true
			}
		} else {
			attribute_value_xhtmlDB.THE_VALUEID.Int64 = 0
			attribute_value_xhtmlDB.THE_VALUEID.Valid = true
		}

		// commit pointer value attribute_value_xhtml.THE_ORIGINAL_VALUE translates to updating the attribute_value_xhtml.THE_ORIGINAL_VALUEID
		attribute_value_xhtmlDB.THE_ORIGINAL_VALUEID.Valid = true // allow for a 0 value (nil association)
		if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
			if THE_ORIGINAL_VALUEId, ok := backRepo.BackRepoXHTML_CONTENT.Map_XHTML_CONTENTPtr_XHTML_CONTENTDBID[attribute_value_xhtml.THE_ORIGINAL_VALUE]; ok {
				attribute_value_xhtmlDB.THE_ORIGINAL_VALUEID.Int64 = int64(THE_ORIGINAL_VALUEId)
				attribute_value_xhtmlDB.THE_ORIGINAL_VALUEID.Valid = true
			}
		} else {
			attribute_value_xhtmlDB.THE_ORIGINAL_VALUEID.Int64 = 0
			attribute_value_xhtmlDB.THE_ORIGINAL_VALUEID.Valid = true
		}

		// commit pointer value attribute_value_xhtml.DEFINITION translates to updating the attribute_value_xhtml.DEFINITIONID
		attribute_value_xhtmlDB.DEFINITIONID.Valid = true // allow for a 0 value (nil association)
		if attribute_value_xhtml.DEFINITION != nil {
			if DEFINITIONId, ok := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_XHTML_REF.Map_A_ATTRIBUTE_DEFINITION_XHTML_REFPtr_A_ATTRIBUTE_DEFINITION_XHTML_REFDBID[attribute_value_xhtml.DEFINITION]; ok {
				attribute_value_xhtmlDB.DEFINITIONID.Int64 = int64(DEFINITIONId)
				attribute_value_xhtmlDB.DEFINITIONID.Valid = true
			}
		} else {
			attribute_value_xhtmlDB.DEFINITIONID.Int64 = 0
			attribute_value_xhtmlDB.DEFINITIONID.Valid = true
		}

		_, err := backRepoATTRIBUTE_VALUE_XHTML.db.Save(attribute_value_xhtmlDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown ATTRIBUTE_VALUE_XHTML intance %s", attribute_value_xhtml.Name))
		return err
	}

	return
}

// BackRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) CheckoutPhaseOne() (Error error) {

	attribute_value_xhtmlDBArray := make([]ATTRIBUTE_VALUE_XHTMLDB, 0)
	_, err := backRepoATTRIBUTE_VALUE_XHTML.db.Find(&attribute_value_xhtmlDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	attribute_value_xhtmlInstancesToBeRemovedFromTheStage := make(map[*models.ATTRIBUTE_VALUE_XHTML]any)
	for key, value := range backRepoATTRIBUTE_VALUE_XHTML.stage.ATTRIBUTE_VALUE_XHTMLs {
		attribute_value_xhtmlInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, attribute_value_xhtmlDB := range attribute_value_xhtmlDBArray {
		backRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseOneInstance(&attribute_value_xhtmlDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		attribute_value_xhtml, ok := backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr[attribute_value_xhtmlDB.ID]
		if ok {
			delete(attribute_value_xhtmlInstancesToBeRemovedFromTheStage, attribute_value_xhtml)
		}
	}

	// remove from stage and back repo's 3 maps all attribute_value_xhtmls that are not in the checkout
	for attribute_value_xhtml := range attribute_value_xhtmlInstancesToBeRemovedFromTheStage {
		attribute_value_xhtml.Unstage(backRepoATTRIBUTE_VALUE_XHTML.GetStage())

		// remove instance from the back repo 3 maps
		attribute_value_xhtmlID := backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID[attribute_value_xhtml]
		delete(backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID, attribute_value_xhtml)
		delete(backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB, attribute_value_xhtmlID)
		delete(backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr, attribute_value_xhtmlID)
	}

	return
}

// CheckoutPhaseOneInstance takes a attribute_value_xhtmlDB that has been found in the DB, updates the backRepo and stages the
// models version of the attribute_value_xhtmlDB
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) CheckoutPhaseOneInstance(attribute_value_xhtmlDB *ATTRIBUTE_VALUE_XHTMLDB) (Error error) {

	attribute_value_xhtml, ok := backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr[attribute_value_xhtmlDB.ID]
	if !ok {
		attribute_value_xhtml = new(models.ATTRIBUTE_VALUE_XHTML)

		backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr[attribute_value_xhtmlDB.ID] = attribute_value_xhtml
		backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID[attribute_value_xhtml] = attribute_value_xhtmlDB.ID

		// append model store with the new element
		attribute_value_xhtml.Name = attribute_value_xhtmlDB.Name_Data.String
		attribute_value_xhtml.Stage(backRepoATTRIBUTE_VALUE_XHTML.GetStage())
	}
	attribute_value_xhtmlDB.CopyBasicFieldsToATTRIBUTE_VALUE_XHTML(attribute_value_xhtml)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	attribute_value_xhtml.Stage(backRepoATTRIBUTE_VALUE_XHTML.GetStage())

	// preserve pointer to attribute_value_xhtmlDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB)[attribute_value_xhtmlDB hold variable pointers
	attribute_value_xhtmlDB_Data := *attribute_value_xhtmlDB
	preservedPtrToATTRIBUTE_VALUE_XHTML := &attribute_value_xhtmlDB_Data
	backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB[attribute_value_xhtmlDB.ID] = preservedPtrToATTRIBUTE_VALUE_XHTML

	return
}

// BackRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseTwo Checkouts all staged instances of ATTRIBUTE_VALUE_XHTML to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, attribute_value_xhtmlDB := range backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB {
		backRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseTwoInstance(backRepo, attribute_value_xhtmlDB)
	}
	return
}

// BackRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseTwoInstance Checkouts staged instances of ATTRIBUTE_VALUE_XHTML to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, attribute_value_xhtmlDB *ATTRIBUTE_VALUE_XHTMLDB) (Error error) {

	attribute_value_xhtml := backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr[attribute_value_xhtmlDB.ID]

	attribute_value_xhtmlDB.DecodePointers(backRepo, attribute_value_xhtml)

	return
}

func (attribute_value_xhtmlDB *ATTRIBUTE_VALUE_XHTMLDB) DecodePointers(backRepo *BackRepoStruct, attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML) {

	// insertion point for checkout of pointer encoding
	// THE_VALUE field	
	{
		id := attribute_value_xhtmlDB.THE_VALUEID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoXHTML_CONTENT.Map_XHTML_CONTENTDBID_XHTML_CONTENTPtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: attribute_value_xhtml.THE_VALUE, unknown pointer id", id)
				attribute_value_xhtml.THE_VALUE = nil
			} else {
				// updates only if field has changed
				if attribute_value_xhtml.THE_VALUE == nil || attribute_value_xhtml.THE_VALUE != tmp {
					attribute_value_xhtml.THE_VALUE = tmp
				}
			}
		} else {
			attribute_value_xhtml.THE_VALUE = nil
		}
	}
	
	// THE_ORIGINAL_VALUE field	
	{
		id := attribute_value_xhtmlDB.THE_ORIGINAL_VALUEID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoXHTML_CONTENT.Map_XHTML_CONTENTDBID_XHTML_CONTENTPtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: attribute_value_xhtml.THE_ORIGINAL_VALUE, unknown pointer id", id)
				attribute_value_xhtml.THE_ORIGINAL_VALUE = nil
			} else {
				// updates only if field has changed
				if attribute_value_xhtml.THE_ORIGINAL_VALUE == nil || attribute_value_xhtml.THE_ORIGINAL_VALUE != tmp {
					attribute_value_xhtml.THE_ORIGINAL_VALUE = tmp
				}
			}
		} else {
			attribute_value_xhtml.THE_ORIGINAL_VALUE = nil
		}
	}
	
	// DEFINITION field	
	{
		id := attribute_value_xhtmlDB.DEFINITIONID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_XHTML_REF.Map_A_ATTRIBUTE_DEFINITION_XHTML_REFDBID_A_ATTRIBUTE_DEFINITION_XHTML_REFPtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: attribute_value_xhtml.DEFINITION, unknown pointer id", id)
				attribute_value_xhtml.DEFINITION = nil
			} else {
				// updates only if field has changed
				if attribute_value_xhtml.DEFINITION == nil || attribute_value_xhtml.DEFINITION != tmp {
					attribute_value_xhtml.DEFINITION = tmp
				}
			}
		} else {
			attribute_value_xhtml.DEFINITION = nil
		}
	}
	
	return
}

// CommitATTRIBUTE_VALUE_XHTML allows commit of a single attribute_value_xhtml (if already staged)
func (backRepo *BackRepoStruct) CommitATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML) {
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CommitPhaseOneInstance(attribute_value_xhtml)
	if id, ok := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID[attribute_value_xhtml]; ok {
		backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CommitPhaseTwoInstance(backRepo, id, attribute_value_xhtml)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitATTRIBUTE_VALUE_XHTML allows checkout of a single attribute_value_xhtml (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML) {
	// check if the attribute_value_xhtml is staged
	if _, ok := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID[attribute_value_xhtml]; ok {

		if id, ok := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID[attribute_value_xhtml]; ok {
			var attribute_value_xhtmlDB ATTRIBUTE_VALUE_XHTMLDB
			attribute_value_xhtmlDB.ID = id

			if _, err := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.db.First(&attribute_value_xhtmlDB, id); err != nil {
				log.Fatalln("CheckoutATTRIBUTE_VALUE_XHTML : Problem with getting object with id:", id)
			}
			backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseOneInstance(&attribute_value_xhtmlDB)
			backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseTwoInstance(backRepo, &attribute_value_xhtmlDB)
		}
	}
}

// CopyBasicFieldsFromATTRIBUTE_VALUE_XHTML
func (attribute_value_xhtmlDB *ATTRIBUTE_VALUE_XHTMLDB) CopyBasicFieldsFromATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML) {
	// insertion point for fields commit

	attribute_value_xhtmlDB.Name_Data.String = attribute_value_xhtml.Name
	attribute_value_xhtmlDB.Name_Data.Valid = true

	attribute_value_xhtmlDB.IS_SIMPLIFIED_Data.Bool = attribute_value_xhtml.IS_SIMPLIFIED
	attribute_value_xhtmlDB.IS_SIMPLIFIED_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTE_VALUE_XHTML_WOP
func (attribute_value_xhtmlDB *ATTRIBUTE_VALUE_XHTMLDB) CopyBasicFieldsFromATTRIBUTE_VALUE_XHTML_WOP(attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML_WOP) {
	// insertion point for fields commit

	attribute_value_xhtmlDB.Name_Data.String = attribute_value_xhtml.Name
	attribute_value_xhtmlDB.Name_Data.Valid = true

	attribute_value_xhtmlDB.IS_SIMPLIFIED_Data.Bool = attribute_value_xhtml.IS_SIMPLIFIED
	attribute_value_xhtmlDB.IS_SIMPLIFIED_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTE_VALUE_XHTMLWOP
func (attribute_value_xhtmlDB *ATTRIBUTE_VALUE_XHTMLDB) CopyBasicFieldsFromATTRIBUTE_VALUE_XHTMLWOP(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTMLWOP) {
	// insertion point for fields commit

	attribute_value_xhtmlDB.Name_Data.String = attribute_value_xhtml.Name
	attribute_value_xhtmlDB.Name_Data.Valid = true

	attribute_value_xhtmlDB.IS_SIMPLIFIED_Data.Bool = attribute_value_xhtml.IS_SIMPLIFIED
	attribute_value_xhtmlDB.IS_SIMPLIFIED_Data.Valid = true
}

// CopyBasicFieldsToATTRIBUTE_VALUE_XHTML
func (attribute_value_xhtmlDB *ATTRIBUTE_VALUE_XHTMLDB) CopyBasicFieldsToATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML) {
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_value_xhtml.Name = attribute_value_xhtmlDB.Name_Data.String
	attribute_value_xhtml.IS_SIMPLIFIED = attribute_value_xhtmlDB.IS_SIMPLIFIED_Data.Bool
}

// CopyBasicFieldsToATTRIBUTE_VALUE_XHTML_WOP
func (attribute_value_xhtmlDB *ATTRIBUTE_VALUE_XHTMLDB) CopyBasicFieldsToATTRIBUTE_VALUE_XHTML_WOP(attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_value_xhtml.Name = attribute_value_xhtmlDB.Name_Data.String
	attribute_value_xhtml.IS_SIMPLIFIED = attribute_value_xhtmlDB.IS_SIMPLIFIED_Data.Bool
}

// CopyBasicFieldsToATTRIBUTE_VALUE_XHTMLWOP
func (attribute_value_xhtmlDB *ATTRIBUTE_VALUE_XHTMLDB) CopyBasicFieldsToATTRIBUTE_VALUE_XHTMLWOP(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTMLWOP) {
	attribute_value_xhtml.ID = int(attribute_value_xhtmlDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	attribute_value_xhtml.Name = attribute_value_xhtmlDB.Name_Data.String
	attribute_value_xhtml.IS_SIMPLIFIED = attribute_value_xhtmlDB.IS_SIMPLIFIED_Data.Bool
}

// Backup generates a json file from a slice of all ATTRIBUTE_VALUE_XHTMLDB instances in the backrepo
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ATTRIBUTE_VALUE_XHTMLDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTE_VALUE_XHTMLDB, 0)
	for _, attribute_value_xhtmlDB := range backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB {
		forBackup = append(forBackup, attribute_value_xhtmlDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json ATTRIBUTE_VALUE_XHTML ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json ATTRIBUTE_VALUE_XHTML file", err.Error())
	}
}

// Backup generates a json file from a slice of all ATTRIBUTE_VALUE_XHTMLDB instances in the backrepo
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTE_VALUE_XHTMLDB, 0)
	for _, attribute_value_xhtmlDB := range backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB {
		forBackup = append(forBackup, attribute_value_xhtmlDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("ATTRIBUTE_VALUE_XHTML")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&ATTRIBUTE_VALUE_XHTML_Fields, -1)
	for _, attribute_value_xhtmlDB := range forBackup {

		var attribute_value_xhtmlWOP ATTRIBUTE_VALUE_XHTMLWOP
		attribute_value_xhtmlDB.CopyBasicFieldsToATTRIBUTE_VALUE_XHTMLWOP(&attribute_value_xhtmlWOP)

		row := sh.AddRow()
		row.WriteStruct(&attribute_value_xhtmlWOP, -1)
	}
}

// RestoreXL from the "ATTRIBUTE_VALUE_XHTML" sheet all ATTRIBUTE_VALUE_XHTMLDB instances
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoATTRIBUTE_VALUE_XHTMLid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["ATTRIBUTE_VALUE_XHTML"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoATTRIBUTE_VALUE_XHTML.rowVisitorATTRIBUTE_VALUE_XHTML)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) rowVisitorATTRIBUTE_VALUE_XHTML(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var attribute_value_xhtmlWOP ATTRIBUTE_VALUE_XHTMLWOP
		row.ReadStruct(&attribute_value_xhtmlWOP)

		// add the unmarshalled struct to the stage
		attribute_value_xhtmlDB := new(ATTRIBUTE_VALUE_XHTMLDB)
		attribute_value_xhtmlDB.CopyBasicFieldsFromATTRIBUTE_VALUE_XHTMLWOP(&attribute_value_xhtmlWOP)

		attribute_value_xhtmlDB_ID_atBackupTime := attribute_value_xhtmlDB.ID
		attribute_value_xhtmlDB.ID = 0
		_, err := backRepoATTRIBUTE_VALUE_XHTML.db.Create(attribute_value_xhtmlDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB[attribute_value_xhtmlDB.ID] = attribute_value_xhtmlDB
		BackRepoATTRIBUTE_VALUE_XHTMLid_atBckpTime_newID[attribute_value_xhtmlDB_ID_atBackupTime] = attribute_value_xhtmlDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ATTRIBUTE_VALUE_XHTMLDB.json" in dirPath that stores an array
// of ATTRIBUTE_VALUE_XHTMLDB and stores it in the database
// the map BackRepoATTRIBUTE_VALUE_XHTMLid_atBckpTime_newID is updated accordingly
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoATTRIBUTE_VALUE_XHTMLid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ATTRIBUTE_VALUE_XHTMLDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json ATTRIBUTE_VALUE_XHTML file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ATTRIBUTE_VALUE_XHTMLDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB
	for _, attribute_value_xhtmlDB := range forRestore {

		attribute_value_xhtmlDB_ID_atBackupTime := attribute_value_xhtmlDB.ID
		attribute_value_xhtmlDB.ID = 0
		_, err := backRepoATTRIBUTE_VALUE_XHTML.db.Create(attribute_value_xhtmlDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB[attribute_value_xhtmlDB.ID] = attribute_value_xhtmlDB
		BackRepoATTRIBUTE_VALUE_XHTMLid_atBckpTime_newID[attribute_value_xhtmlDB_ID_atBackupTime] = attribute_value_xhtmlDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json ATTRIBUTE_VALUE_XHTML file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<ATTRIBUTE_VALUE_XHTML>id_atBckpTime_newID
// to compute new index
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) RestorePhaseTwo() {

	for _, attribute_value_xhtmlDB := range backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB {

		// next line of code is to avert unused variable compilation error
		_ = attribute_value_xhtmlDB

		// insertion point for reindexing pointers encoding
		// reindexing THE_VALUE field
		if attribute_value_xhtmlDB.THE_VALUEID.Int64 != 0 {
			attribute_value_xhtmlDB.THE_VALUEID.Int64 = int64(BackRepoXHTML_CONTENTid_atBckpTime_newID[uint(attribute_value_xhtmlDB.THE_VALUEID.Int64)])
			attribute_value_xhtmlDB.THE_VALUEID.Valid = true
		}

		// reindexing THE_ORIGINAL_VALUE field
		if attribute_value_xhtmlDB.THE_ORIGINAL_VALUEID.Int64 != 0 {
			attribute_value_xhtmlDB.THE_ORIGINAL_VALUEID.Int64 = int64(BackRepoXHTML_CONTENTid_atBckpTime_newID[uint(attribute_value_xhtmlDB.THE_ORIGINAL_VALUEID.Int64)])
			attribute_value_xhtmlDB.THE_ORIGINAL_VALUEID.Valid = true
		}

		// reindexing DEFINITION field
		if attribute_value_xhtmlDB.DEFINITIONID.Int64 != 0 {
			attribute_value_xhtmlDB.DEFINITIONID.Int64 = int64(BackRepoA_ATTRIBUTE_DEFINITION_XHTML_REFid_atBckpTime_newID[uint(attribute_value_xhtmlDB.DEFINITIONID.Int64)])
			attribute_value_xhtmlDB.DEFINITIONID.Valid = true
		}

		// update databse with new index encoding
		db, _ := backRepoATTRIBUTE_VALUE_XHTML.db.Model(attribute_value_xhtmlDB)
		_, err := db.Updates(*attribute_value_xhtmlDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoATTRIBUTE_VALUE_XHTML.ResetReversePointers commits all staged instances of ATTRIBUTE_VALUE_XHTML to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, attribute_value_xhtml := range backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr {
		backRepoATTRIBUTE_VALUE_XHTML.ResetReversePointersInstance(backRepo, idx, attribute_value_xhtml)
	}

	return
}

func (backRepoATTRIBUTE_VALUE_XHTML *BackRepoATTRIBUTE_VALUE_XHTMLStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML) (Error error) {

	// fetch matching attribute_value_xhtmlDB
	if attribute_value_xhtmlDB, ok := backRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB[idx]; ok {
		_ = attribute_value_xhtmlDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoATTRIBUTE_VALUE_XHTMLid_atBckpTime_newID map[uint]uint
