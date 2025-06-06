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
var dummy_A_RELATION_GROUP_TYPE_REF_sql sql.NullBool
var dummy_A_RELATION_GROUP_TYPE_REF_time time.Duration
var dummy_A_RELATION_GROUP_TYPE_REF_sort sort.Float64Slice

// A_RELATION_GROUP_TYPE_REFAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model a_relation_group_type_refAPI
type A_RELATION_GROUP_TYPE_REFAPI struct {
	gorm.Model

	models.A_RELATION_GROUP_TYPE_REF_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	A_RELATION_GROUP_TYPE_REFPointersEncoding A_RELATION_GROUP_TYPE_REFPointersEncoding
}

// A_RELATION_GROUP_TYPE_REFPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type A_RELATION_GROUP_TYPE_REFPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// A_RELATION_GROUP_TYPE_REFDB describes a a_relation_group_type_ref in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model a_relation_group_type_refDB
type A_RELATION_GROUP_TYPE_REFDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field a_relation_group_type_refDB.Name
	Name_Data sql.NullString

	// Declation for basic field a_relation_group_type_refDB.RELATION_GROUP_TYPE_REF
	RELATION_GROUP_TYPE_REF_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	A_RELATION_GROUP_TYPE_REFPointersEncoding
}

// A_RELATION_GROUP_TYPE_REFDBs arrays a_relation_group_type_refDBs
// swagger:response a_relation_group_type_refDBsResponse
type A_RELATION_GROUP_TYPE_REFDBs []A_RELATION_GROUP_TYPE_REFDB

// A_RELATION_GROUP_TYPE_REFDBResponse provides response
// swagger:response a_relation_group_type_refDBResponse
type A_RELATION_GROUP_TYPE_REFDBResponse struct {
	A_RELATION_GROUP_TYPE_REFDB
}

// A_RELATION_GROUP_TYPE_REFWOP is a A_RELATION_GROUP_TYPE_REF without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type A_RELATION_GROUP_TYPE_REFWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	RELATION_GROUP_TYPE_REF string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var A_RELATION_GROUP_TYPE_REF_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"RELATION_GROUP_TYPE_REF",
}

type BackRepoA_RELATION_GROUP_TYPE_REFStruct struct {
	// stores A_RELATION_GROUP_TYPE_REFDB according to their gorm ID
	Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB map[uint]*A_RELATION_GROUP_TYPE_REFDB

	// stores A_RELATION_GROUP_TYPE_REFDB ID according to A_RELATION_GROUP_TYPE_REF address
	Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID map[*models.A_RELATION_GROUP_TYPE_REF]uint

	// stores A_RELATION_GROUP_TYPE_REF according to their gorm ID
	Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr map[uint]*models.A_RELATION_GROUP_TYPE_REF

	db db.DBInterface

	stage *models.Stage
}

func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) GetStage() (stage *models.Stage) {
	stage = backRepoA_RELATION_GROUP_TYPE_REF.stage
	return
}

func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) GetDB() db.DBInterface {
	return backRepoA_RELATION_GROUP_TYPE_REF.db
}

// GetA_RELATION_GROUP_TYPE_REFDBFromA_RELATION_GROUP_TYPE_REFPtr is a handy function to access the back repo instance from the stage instance
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) GetA_RELATION_GROUP_TYPE_REFDBFromA_RELATION_GROUP_TYPE_REFPtr(a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF) (a_relation_group_type_refDB *A_RELATION_GROUP_TYPE_REFDB) {
	id := backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID[a_relation_group_type_ref]
	a_relation_group_type_refDB = backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB[id]
	return
}

// BackRepoA_RELATION_GROUP_TYPE_REF.CommitPhaseOne commits all staged instances of A_RELATION_GROUP_TYPE_REF to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var a_relation_group_type_refs []*models.A_RELATION_GROUP_TYPE_REF
	for a_relation_group_type_ref := range stage.A_RELATION_GROUP_TYPE_REFs {
		a_relation_group_type_refs = append(a_relation_group_type_refs, a_relation_group_type_ref)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(a_relation_group_type_refs, func(i, j int) bool {
		return stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_refs[i]] < stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_refs[j]]
	})

	for _, a_relation_group_type_ref := range a_relation_group_type_refs {
		backRepoA_RELATION_GROUP_TYPE_REF.CommitPhaseOneInstance(a_relation_group_type_ref)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, a_relation_group_type_ref := range backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr {
		if _, ok := stage.A_RELATION_GROUP_TYPE_REFs[a_relation_group_type_ref]; !ok {
			backRepoA_RELATION_GROUP_TYPE_REF.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoA_RELATION_GROUP_TYPE_REF.CommitDeleteInstance commits deletion of A_RELATION_GROUP_TYPE_REF to the BackRepo
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) CommitDeleteInstance(id uint) (Error error) {

	a_relation_group_type_ref := backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr[id]

	// a_relation_group_type_ref is not staged anymore, remove a_relation_group_type_refDB
	a_relation_group_type_refDB := backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB[id]
	db, _ := backRepoA_RELATION_GROUP_TYPE_REF.db.Unscoped()
	_, err := db.Delete(a_relation_group_type_refDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID, a_relation_group_type_ref)
	delete(backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr, id)
	delete(backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB, id)

	return
}

// BackRepoA_RELATION_GROUP_TYPE_REF.CommitPhaseOneInstance commits a_relation_group_type_ref staged instances of A_RELATION_GROUP_TYPE_REF to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) CommitPhaseOneInstance(a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF) (Error error) {

	// check if the a_relation_group_type_ref is not commited yet
	if _, ok := backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID[a_relation_group_type_ref]; ok {
		return
	}

	// initiate a_relation_group_type_ref
	var a_relation_group_type_refDB A_RELATION_GROUP_TYPE_REFDB
	a_relation_group_type_refDB.CopyBasicFieldsFromA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref)

	_, err := backRepoA_RELATION_GROUP_TYPE_REF.db.Create(&a_relation_group_type_refDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID[a_relation_group_type_ref] = a_relation_group_type_refDB.ID
	backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr[a_relation_group_type_refDB.ID] = a_relation_group_type_ref
	backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB[a_relation_group_type_refDB.ID] = &a_relation_group_type_refDB

	return
}

// BackRepoA_RELATION_GROUP_TYPE_REF.CommitPhaseTwo commits all staged instances of A_RELATION_GROUP_TYPE_REF to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, a_relation_group_type_ref := range backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr {
		backRepoA_RELATION_GROUP_TYPE_REF.CommitPhaseTwoInstance(backRepo, idx, a_relation_group_type_ref)
	}

	return
}

// BackRepoA_RELATION_GROUP_TYPE_REF.CommitPhaseTwoInstance commits {{structname }} of models.A_RELATION_GROUP_TYPE_REF to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF) (Error error) {

	// fetch matching a_relation_group_type_refDB
	if a_relation_group_type_refDB, ok := backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB[idx]; ok {

		a_relation_group_type_refDB.CopyBasicFieldsFromA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoA_RELATION_GROUP_TYPE_REF.db.Save(a_relation_group_type_refDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown A_RELATION_GROUP_TYPE_REF intance %s", a_relation_group_type_ref.Name))
		return err
	}

	return
}

// BackRepoA_RELATION_GROUP_TYPE_REF.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) CheckoutPhaseOne() (Error error) {

	a_relation_group_type_refDBArray := make([]A_RELATION_GROUP_TYPE_REFDB, 0)
	_, err := backRepoA_RELATION_GROUP_TYPE_REF.db.Find(&a_relation_group_type_refDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	a_relation_group_type_refInstancesToBeRemovedFromTheStage := make(map[*models.A_RELATION_GROUP_TYPE_REF]any)
	for key, value := range backRepoA_RELATION_GROUP_TYPE_REF.stage.A_RELATION_GROUP_TYPE_REFs {
		a_relation_group_type_refInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, a_relation_group_type_refDB := range a_relation_group_type_refDBArray {
		backRepoA_RELATION_GROUP_TYPE_REF.CheckoutPhaseOneInstance(&a_relation_group_type_refDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		a_relation_group_type_ref, ok := backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr[a_relation_group_type_refDB.ID]
		if ok {
			delete(a_relation_group_type_refInstancesToBeRemovedFromTheStage, a_relation_group_type_ref)
		}
	}

	// remove from stage and back repo's 3 maps all a_relation_group_type_refs that are not in the checkout
	for a_relation_group_type_ref := range a_relation_group_type_refInstancesToBeRemovedFromTheStage {
		a_relation_group_type_ref.Unstage(backRepoA_RELATION_GROUP_TYPE_REF.GetStage())

		// remove instance from the back repo 3 maps
		a_relation_group_type_refID := backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID[a_relation_group_type_ref]
		delete(backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID, a_relation_group_type_ref)
		delete(backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB, a_relation_group_type_refID)
		delete(backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr, a_relation_group_type_refID)
	}

	return
}

// CheckoutPhaseOneInstance takes a a_relation_group_type_refDB that has been found in the DB, updates the backRepo and stages the
// models version of the a_relation_group_type_refDB
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) CheckoutPhaseOneInstance(a_relation_group_type_refDB *A_RELATION_GROUP_TYPE_REFDB) (Error error) {

	a_relation_group_type_ref, ok := backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr[a_relation_group_type_refDB.ID]
	if !ok {
		a_relation_group_type_ref = new(models.A_RELATION_GROUP_TYPE_REF)

		backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr[a_relation_group_type_refDB.ID] = a_relation_group_type_ref
		backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID[a_relation_group_type_ref] = a_relation_group_type_refDB.ID

		// append model store with the new element
		a_relation_group_type_ref.Name = a_relation_group_type_refDB.Name_Data.String
		a_relation_group_type_ref.Stage(backRepoA_RELATION_GROUP_TYPE_REF.GetStage())
	}
	a_relation_group_type_refDB.CopyBasicFieldsToA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	a_relation_group_type_ref.Stage(backRepoA_RELATION_GROUP_TYPE_REF.GetStage())

	// preserve pointer to a_relation_group_type_refDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB)[a_relation_group_type_refDB hold variable pointers
	a_relation_group_type_refDB_Data := *a_relation_group_type_refDB
	preservedPtrToA_RELATION_GROUP_TYPE_REF := &a_relation_group_type_refDB_Data
	backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB[a_relation_group_type_refDB.ID] = preservedPtrToA_RELATION_GROUP_TYPE_REF

	return
}

// BackRepoA_RELATION_GROUP_TYPE_REF.CheckoutPhaseTwo Checkouts all staged instances of A_RELATION_GROUP_TYPE_REF to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, a_relation_group_type_refDB := range backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB {
		backRepoA_RELATION_GROUP_TYPE_REF.CheckoutPhaseTwoInstance(backRepo, a_relation_group_type_refDB)
	}
	return
}

// BackRepoA_RELATION_GROUP_TYPE_REF.CheckoutPhaseTwoInstance Checkouts staged instances of A_RELATION_GROUP_TYPE_REF to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, a_relation_group_type_refDB *A_RELATION_GROUP_TYPE_REFDB) (Error error) {

	a_relation_group_type_ref := backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr[a_relation_group_type_refDB.ID]

	a_relation_group_type_refDB.DecodePointers(backRepo, a_relation_group_type_ref)

	return
}

func (a_relation_group_type_refDB *A_RELATION_GROUP_TYPE_REFDB) DecodePointers(backRepo *BackRepoStruct, a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitA_RELATION_GROUP_TYPE_REF allows commit of a single a_relation_group_type_ref (if already staged)
func (backRepo *BackRepoStruct) CommitA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF) {
	backRepo.BackRepoA_RELATION_GROUP_TYPE_REF.CommitPhaseOneInstance(a_relation_group_type_ref)
	if id, ok := backRepo.BackRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID[a_relation_group_type_ref]; ok {
		backRepo.BackRepoA_RELATION_GROUP_TYPE_REF.CommitPhaseTwoInstance(backRepo, id, a_relation_group_type_ref)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitA_RELATION_GROUP_TYPE_REF allows checkout of a single a_relation_group_type_ref (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF) {
	// check if the a_relation_group_type_ref is staged
	if _, ok := backRepo.BackRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID[a_relation_group_type_ref]; ok {

		if id, ok := backRepo.BackRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFPtr_A_RELATION_GROUP_TYPE_REFDBID[a_relation_group_type_ref]; ok {
			var a_relation_group_type_refDB A_RELATION_GROUP_TYPE_REFDB
			a_relation_group_type_refDB.ID = id

			if _, err := backRepo.BackRepoA_RELATION_GROUP_TYPE_REF.db.First(&a_relation_group_type_refDB, id); err != nil {
				log.Fatalln("CheckoutA_RELATION_GROUP_TYPE_REF : Problem with getting object with id:", id)
			}
			backRepo.BackRepoA_RELATION_GROUP_TYPE_REF.CheckoutPhaseOneInstance(&a_relation_group_type_refDB)
			backRepo.BackRepoA_RELATION_GROUP_TYPE_REF.CheckoutPhaseTwoInstance(backRepo, &a_relation_group_type_refDB)
		}
	}
}

// CopyBasicFieldsFromA_RELATION_GROUP_TYPE_REF
func (a_relation_group_type_refDB *A_RELATION_GROUP_TYPE_REFDB) CopyBasicFieldsFromA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF) {
	// insertion point for fields commit

	a_relation_group_type_refDB.Name_Data.String = a_relation_group_type_ref.Name
	a_relation_group_type_refDB.Name_Data.Valid = true

	a_relation_group_type_refDB.RELATION_GROUP_TYPE_REF_Data.String = a_relation_group_type_ref.RELATION_GROUP_TYPE_REF
	a_relation_group_type_refDB.RELATION_GROUP_TYPE_REF_Data.Valid = true
}

// CopyBasicFieldsFromA_RELATION_GROUP_TYPE_REF_WOP
func (a_relation_group_type_refDB *A_RELATION_GROUP_TYPE_REFDB) CopyBasicFieldsFromA_RELATION_GROUP_TYPE_REF_WOP(a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF_WOP) {
	// insertion point for fields commit

	a_relation_group_type_refDB.Name_Data.String = a_relation_group_type_ref.Name
	a_relation_group_type_refDB.Name_Data.Valid = true

	a_relation_group_type_refDB.RELATION_GROUP_TYPE_REF_Data.String = a_relation_group_type_ref.RELATION_GROUP_TYPE_REF
	a_relation_group_type_refDB.RELATION_GROUP_TYPE_REF_Data.Valid = true
}

// CopyBasicFieldsFromA_RELATION_GROUP_TYPE_REFWOP
func (a_relation_group_type_refDB *A_RELATION_GROUP_TYPE_REFDB) CopyBasicFieldsFromA_RELATION_GROUP_TYPE_REFWOP(a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REFWOP) {
	// insertion point for fields commit

	a_relation_group_type_refDB.Name_Data.String = a_relation_group_type_ref.Name
	a_relation_group_type_refDB.Name_Data.Valid = true

	a_relation_group_type_refDB.RELATION_GROUP_TYPE_REF_Data.String = a_relation_group_type_ref.RELATION_GROUP_TYPE_REF
	a_relation_group_type_refDB.RELATION_GROUP_TYPE_REF_Data.Valid = true
}

// CopyBasicFieldsToA_RELATION_GROUP_TYPE_REF
func (a_relation_group_type_refDB *A_RELATION_GROUP_TYPE_REFDB) CopyBasicFieldsToA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF) {
	// insertion point for checkout of basic fields (back repo to stage)
	a_relation_group_type_ref.Name = a_relation_group_type_refDB.Name_Data.String
	a_relation_group_type_ref.RELATION_GROUP_TYPE_REF = a_relation_group_type_refDB.RELATION_GROUP_TYPE_REF_Data.String
}

// CopyBasicFieldsToA_RELATION_GROUP_TYPE_REF_WOP
func (a_relation_group_type_refDB *A_RELATION_GROUP_TYPE_REFDB) CopyBasicFieldsToA_RELATION_GROUP_TYPE_REF_WOP(a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	a_relation_group_type_ref.Name = a_relation_group_type_refDB.Name_Data.String
	a_relation_group_type_ref.RELATION_GROUP_TYPE_REF = a_relation_group_type_refDB.RELATION_GROUP_TYPE_REF_Data.String
}

// CopyBasicFieldsToA_RELATION_GROUP_TYPE_REFWOP
func (a_relation_group_type_refDB *A_RELATION_GROUP_TYPE_REFDB) CopyBasicFieldsToA_RELATION_GROUP_TYPE_REFWOP(a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REFWOP) {
	a_relation_group_type_ref.ID = int(a_relation_group_type_refDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	a_relation_group_type_ref.Name = a_relation_group_type_refDB.Name_Data.String
	a_relation_group_type_ref.RELATION_GROUP_TYPE_REF = a_relation_group_type_refDB.RELATION_GROUP_TYPE_REF_Data.String
}

// Backup generates a json file from a slice of all A_RELATION_GROUP_TYPE_REFDB instances in the backrepo
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "A_RELATION_GROUP_TYPE_REFDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*A_RELATION_GROUP_TYPE_REFDB, 0)
	for _, a_relation_group_type_refDB := range backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB {
		forBackup = append(forBackup, a_relation_group_type_refDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json A_RELATION_GROUP_TYPE_REF ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json A_RELATION_GROUP_TYPE_REF file", err.Error())
	}
}

// Backup generates a json file from a slice of all A_RELATION_GROUP_TYPE_REFDB instances in the backrepo
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*A_RELATION_GROUP_TYPE_REFDB, 0)
	for _, a_relation_group_type_refDB := range backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB {
		forBackup = append(forBackup, a_relation_group_type_refDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("A_RELATION_GROUP_TYPE_REF")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&A_RELATION_GROUP_TYPE_REF_Fields, -1)
	for _, a_relation_group_type_refDB := range forBackup {

		var a_relation_group_type_refWOP A_RELATION_GROUP_TYPE_REFWOP
		a_relation_group_type_refDB.CopyBasicFieldsToA_RELATION_GROUP_TYPE_REFWOP(&a_relation_group_type_refWOP)

		row := sh.AddRow()
		row.WriteStruct(&a_relation_group_type_refWOP, -1)
	}
}

// RestoreXL from the "A_RELATION_GROUP_TYPE_REF" sheet all A_RELATION_GROUP_TYPE_REFDB instances
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoA_RELATION_GROUP_TYPE_REFid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["A_RELATION_GROUP_TYPE_REF"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoA_RELATION_GROUP_TYPE_REF.rowVisitorA_RELATION_GROUP_TYPE_REF)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) rowVisitorA_RELATION_GROUP_TYPE_REF(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var a_relation_group_type_refWOP A_RELATION_GROUP_TYPE_REFWOP
		row.ReadStruct(&a_relation_group_type_refWOP)

		// add the unmarshalled struct to the stage
		a_relation_group_type_refDB := new(A_RELATION_GROUP_TYPE_REFDB)
		a_relation_group_type_refDB.CopyBasicFieldsFromA_RELATION_GROUP_TYPE_REFWOP(&a_relation_group_type_refWOP)

		a_relation_group_type_refDB_ID_atBackupTime := a_relation_group_type_refDB.ID
		a_relation_group_type_refDB.ID = 0
		_, err := backRepoA_RELATION_GROUP_TYPE_REF.db.Create(a_relation_group_type_refDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB[a_relation_group_type_refDB.ID] = a_relation_group_type_refDB
		BackRepoA_RELATION_GROUP_TYPE_REFid_atBckpTime_newID[a_relation_group_type_refDB_ID_atBackupTime] = a_relation_group_type_refDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "A_RELATION_GROUP_TYPE_REFDB.json" in dirPath that stores an array
// of A_RELATION_GROUP_TYPE_REFDB and stores it in the database
// the map BackRepoA_RELATION_GROUP_TYPE_REFid_atBckpTime_newID is updated accordingly
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoA_RELATION_GROUP_TYPE_REFid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "A_RELATION_GROUP_TYPE_REFDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json A_RELATION_GROUP_TYPE_REF file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*A_RELATION_GROUP_TYPE_REFDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB
	for _, a_relation_group_type_refDB := range forRestore {

		a_relation_group_type_refDB_ID_atBackupTime := a_relation_group_type_refDB.ID
		a_relation_group_type_refDB.ID = 0
		_, err := backRepoA_RELATION_GROUP_TYPE_REF.db.Create(a_relation_group_type_refDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB[a_relation_group_type_refDB.ID] = a_relation_group_type_refDB
		BackRepoA_RELATION_GROUP_TYPE_REFid_atBckpTime_newID[a_relation_group_type_refDB_ID_atBackupTime] = a_relation_group_type_refDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json A_RELATION_GROUP_TYPE_REF file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<A_RELATION_GROUP_TYPE_REF>id_atBckpTime_newID
// to compute new index
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) RestorePhaseTwo() {

	for _, a_relation_group_type_refDB := range backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB {

		// next line of code is to avert unused variable compilation error
		_ = a_relation_group_type_refDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoA_RELATION_GROUP_TYPE_REF.db.Model(a_relation_group_type_refDB)
		_, err := db.Updates(*a_relation_group_type_refDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoA_RELATION_GROUP_TYPE_REF.ResetReversePointers commits all staged instances of A_RELATION_GROUP_TYPE_REF to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, a_relation_group_type_ref := range backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFPtr {
		backRepoA_RELATION_GROUP_TYPE_REF.ResetReversePointersInstance(backRepo, idx, a_relation_group_type_ref)
	}

	return
}

func (backRepoA_RELATION_GROUP_TYPE_REF *BackRepoA_RELATION_GROUP_TYPE_REFStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF) (Error error) {

	// fetch matching a_relation_group_type_refDB
	if a_relation_group_type_refDB, ok := backRepoA_RELATION_GROUP_TYPE_REF.Map_A_RELATION_GROUP_TYPE_REFDBID_A_RELATION_GROUP_TYPE_REFDB[idx]; ok {
		_ = a_relation_group_type_refDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoA_RELATION_GROUP_TYPE_REFid_atBckpTime_newID map[uint]uint
