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
var dummy_A_SPEC_RELATION_GROUPS_sql sql.NullBool
var dummy_A_SPEC_RELATION_GROUPS_time time.Duration
var dummy_A_SPEC_RELATION_GROUPS_sort sort.Float64Slice

// A_SPEC_RELATION_GROUPSAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model a_spec_relation_groupsAPI
type A_SPEC_RELATION_GROUPSAPI struct {
	gorm.Model

	models.A_SPEC_RELATION_GROUPS_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	A_SPEC_RELATION_GROUPSPointersEncoding A_SPEC_RELATION_GROUPSPointersEncoding
}

// A_SPEC_RELATION_GROUPSPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type A_SPEC_RELATION_GROUPSPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field RELATION_GROUP is a slice of pointers to another Struct (optional or 0..1)
	RELATION_GROUP IntSlice `gorm:"type:TEXT"`
}

// A_SPEC_RELATION_GROUPSDB describes a a_spec_relation_groups in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model a_spec_relation_groupsDB
type A_SPEC_RELATION_GROUPSDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field a_spec_relation_groupsDB.Name
	Name_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	A_SPEC_RELATION_GROUPSPointersEncoding
}

// A_SPEC_RELATION_GROUPSDBs arrays a_spec_relation_groupsDBs
// swagger:response a_spec_relation_groupsDBsResponse
type A_SPEC_RELATION_GROUPSDBs []A_SPEC_RELATION_GROUPSDB

// A_SPEC_RELATION_GROUPSDBResponse provides response
// swagger:response a_spec_relation_groupsDBResponse
type A_SPEC_RELATION_GROUPSDBResponse struct {
	A_SPEC_RELATION_GROUPSDB
}

// A_SPEC_RELATION_GROUPSWOP is a A_SPEC_RELATION_GROUPS without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type A_SPEC_RELATION_GROUPSWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var A_SPEC_RELATION_GROUPS_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoA_SPEC_RELATION_GROUPSStruct struct {
	// stores A_SPEC_RELATION_GROUPSDB according to their gorm ID
	Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB map[uint]*A_SPEC_RELATION_GROUPSDB

	// stores A_SPEC_RELATION_GROUPSDB ID according to A_SPEC_RELATION_GROUPS address
	Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID map[*models.A_SPEC_RELATION_GROUPS]uint

	// stores A_SPEC_RELATION_GROUPS according to their gorm ID
	Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr map[uint]*models.A_SPEC_RELATION_GROUPS

	db db.DBInterface

	stage *models.Stage
}

func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) GetStage() (stage *models.Stage) {
	stage = backRepoA_SPEC_RELATION_GROUPS.stage
	return
}

func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) GetDB() db.DBInterface {
	return backRepoA_SPEC_RELATION_GROUPS.db
}

// GetA_SPEC_RELATION_GROUPSDBFromA_SPEC_RELATION_GROUPSPtr is a handy function to access the back repo instance from the stage instance
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) GetA_SPEC_RELATION_GROUPSDBFromA_SPEC_RELATION_GROUPSPtr(a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS) (a_spec_relation_groupsDB *A_SPEC_RELATION_GROUPSDB) {
	id := backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID[a_spec_relation_groups]
	a_spec_relation_groupsDB = backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB[id]
	return
}

// BackRepoA_SPEC_RELATION_GROUPS.CommitPhaseOne commits all staged instances of A_SPEC_RELATION_GROUPS to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var a_spec_relation_groupss []*models.A_SPEC_RELATION_GROUPS
	for a_spec_relation_groups := range stage.A_SPEC_RELATION_GROUPSs {
		a_spec_relation_groupss = append(a_spec_relation_groupss, a_spec_relation_groups)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(a_spec_relation_groupss, func(i, j int) bool {
		return stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groupss[i]] < stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groupss[j]]
	})

	for _, a_spec_relation_groups := range a_spec_relation_groupss {
		backRepoA_SPEC_RELATION_GROUPS.CommitPhaseOneInstance(a_spec_relation_groups)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, a_spec_relation_groups := range backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr {
		if _, ok := stage.A_SPEC_RELATION_GROUPSs[a_spec_relation_groups]; !ok {
			backRepoA_SPEC_RELATION_GROUPS.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoA_SPEC_RELATION_GROUPS.CommitDeleteInstance commits deletion of A_SPEC_RELATION_GROUPS to the BackRepo
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) CommitDeleteInstance(id uint) (Error error) {

	a_spec_relation_groups := backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr[id]

	// a_spec_relation_groups is not staged anymore, remove a_spec_relation_groupsDB
	a_spec_relation_groupsDB := backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB[id]
	db, _ := backRepoA_SPEC_RELATION_GROUPS.db.Unscoped()
	_, err := db.Delete(a_spec_relation_groupsDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID, a_spec_relation_groups)
	delete(backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr, id)
	delete(backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB, id)

	return
}

// BackRepoA_SPEC_RELATION_GROUPS.CommitPhaseOneInstance commits a_spec_relation_groups staged instances of A_SPEC_RELATION_GROUPS to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) CommitPhaseOneInstance(a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS) (Error error) {

	// check if the a_spec_relation_groups is not commited yet
	if _, ok := backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID[a_spec_relation_groups]; ok {
		return
	}

	// initiate a_spec_relation_groups
	var a_spec_relation_groupsDB A_SPEC_RELATION_GROUPSDB
	a_spec_relation_groupsDB.CopyBasicFieldsFromA_SPEC_RELATION_GROUPS(a_spec_relation_groups)

	_, err := backRepoA_SPEC_RELATION_GROUPS.db.Create(&a_spec_relation_groupsDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID[a_spec_relation_groups] = a_spec_relation_groupsDB.ID
	backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr[a_spec_relation_groupsDB.ID] = a_spec_relation_groups
	backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB[a_spec_relation_groupsDB.ID] = &a_spec_relation_groupsDB

	return
}

// BackRepoA_SPEC_RELATION_GROUPS.CommitPhaseTwo commits all staged instances of A_SPEC_RELATION_GROUPS to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, a_spec_relation_groups := range backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr {
		backRepoA_SPEC_RELATION_GROUPS.CommitPhaseTwoInstance(backRepo, idx, a_spec_relation_groups)
	}

	return
}

// BackRepoA_SPEC_RELATION_GROUPS.CommitPhaseTwoInstance commits {{structname }} of models.A_SPEC_RELATION_GROUPS to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS) (Error error) {

	// fetch matching a_spec_relation_groupsDB
	if a_spec_relation_groupsDB, ok := backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB[idx]; ok {

		a_spec_relation_groupsDB.CopyBasicFieldsFromA_SPEC_RELATION_GROUPS(a_spec_relation_groups)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		a_spec_relation_groupsDB.A_SPEC_RELATION_GROUPSPointersEncoding.RELATION_GROUP = make([]int, 0)
		// 2. encode
		for _, relation_groupAssocEnd := range a_spec_relation_groups.RELATION_GROUP {
			relation_groupAssocEnd_DB :=
				backRepo.BackRepoRELATION_GROUP.GetRELATION_GROUPDBFromRELATION_GROUPPtr(relation_groupAssocEnd)
			
			// the stage might be inconsistant, meaning that the relation_groupAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if relation_groupAssocEnd_DB == nil {
				continue
			}
			
			a_spec_relation_groupsDB.A_SPEC_RELATION_GROUPSPointersEncoding.RELATION_GROUP =
				append(a_spec_relation_groupsDB.A_SPEC_RELATION_GROUPSPointersEncoding.RELATION_GROUP, int(relation_groupAssocEnd_DB.ID))
		}

		_, err := backRepoA_SPEC_RELATION_GROUPS.db.Save(a_spec_relation_groupsDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown A_SPEC_RELATION_GROUPS intance %s", a_spec_relation_groups.Name))
		return err
	}

	return
}

// BackRepoA_SPEC_RELATION_GROUPS.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) CheckoutPhaseOne() (Error error) {

	a_spec_relation_groupsDBArray := make([]A_SPEC_RELATION_GROUPSDB, 0)
	_, err := backRepoA_SPEC_RELATION_GROUPS.db.Find(&a_spec_relation_groupsDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	a_spec_relation_groupsInstancesToBeRemovedFromTheStage := make(map[*models.A_SPEC_RELATION_GROUPS]any)
	for key, value := range backRepoA_SPEC_RELATION_GROUPS.stage.A_SPEC_RELATION_GROUPSs {
		a_spec_relation_groupsInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, a_spec_relation_groupsDB := range a_spec_relation_groupsDBArray {
		backRepoA_SPEC_RELATION_GROUPS.CheckoutPhaseOneInstance(&a_spec_relation_groupsDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		a_spec_relation_groups, ok := backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr[a_spec_relation_groupsDB.ID]
		if ok {
			delete(a_spec_relation_groupsInstancesToBeRemovedFromTheStage, a_spec_relation_groups)
		}
	}

	// remove from stage and back repo's 3 maps all a_spec_relation_groupss that are not in the checkout
	for a_spec_relation_groups := range a_spec_relation_groupsInstancesToBeRemovedFromTheStage {
		a_spec_relation_groups.Unstage(backRepoA_SPEC_RELATION_GROUPS.GetStage())

		// remove instance from the back repo 3 maps
		a_spec_relation_groupsID := backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID[a_spec_relation_groups]
		delete(backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID, a_spec_relation_groups)
		delete(backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB, a_spec_relation_groupsID)
		delete(backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr, a_spec_relation_groupsID)
	}

	return
}

// CheckoutPhaseOneInstance takes a a_spec_relation_groupsDB that has been found in the DB, updates the backRepo and stages the
// models version of the a_spec_relation_groupsDB
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) CheckoutPhaseOneInstance(a_spec_relation_groupsDB *A_SPEC_RELATION_GROUPSDB) (Error error) {

	a_spec_relation_groups, ok := backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr[a_spec_relation_groupsDB.ID]
	if !ok {
		a_spec_relation_groups = new(models.A_SPEC_RELATION_GROUPS)

		backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr[a_spec_relation_groupsDB.ID] = a_spec_relation_groups
		backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID[a_spec_relation_groups] = a_spec_relation_groupsDB.ID

		// append model store with the new element
		a_spec_relation_groups.Name = a_spec_relation_groupsDB.Name_Data.String
		a_spec_relation_groups.Stage(backRepoA_SPEC_RELATION_GROUPS.GetStage())
	}
	a_spec_relation_groupsDB.CopyBasicFieldsToA_SPEC_RELATION_GROUPS(a_spec_relation_groups)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	a_spec_relation_groups.Stage(backRepoA_SPEC_RELATION_GROUPS.GetStage())

	// preserve pointer to a_spec_relation_groupsDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB)[a_spec_relation_groupsDB hold variable pointers
	a_spec_relation_groupsDB_Data := *a_spec_relation_groupsDB
	preservedPtrToA_SPEC_RELATION_GROUPS := &a_spec_relation_groupsDB_Data
	backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB[a_spec_relation_groupsDB.ID] = preservedPtrToA_SPEC_RELATION_GROUPS

	return
}

// BackRepoA_SPEC_RELATION_GROUPS.CheckoutPhaseTwo Checkouts all staged instances of A_SPEC_RELATION_GROUPS to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, a_spec_relation_groupsDB := range backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB {
		backRepoA_SPEC_RELATION_GROUPS.CheckoutPhaseTwoInstance(backRepo, a_spec_relation_groupsDB)
	}
	return
}

// BackRepoA_SPEC_RELATION_GROUPS.CheckoutPhaseTwoInstance Checkouts staged instances of A_SPEC_RELATION_GROUPS to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, a_spec_relation_groupsDB *A_SPEC_RELATION_GROUPSDB) (Error error) {

	a_spec_relation_groups := backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr[a_spec_relation_groupsDB.ID]

	a_spec_relation_groupsDB.DecodePointers(backRepo, a_spec_relation_groups)

	return
}

func (a_spec_relation_groupsDB *A_SPEC_RELATION_GROUPSDB) DecodePointers(backRepo *BackRepoStruct, a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS) {

	// insertion point for checkout of pointer encoding
	// This loop redeem a_spec_relation_groups.RELATION_GROUP in the stage from the encode in the back repo
	// It parses all RELATION_GROUPDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	a_spec_relation_groups.RELATION_GROUP = a_spec_relation_groups.RELATION_GROUP[:0]
	for _, _RELATION_GROUPid := range a_spec_relation_groupsDB.A_SPEC_RELATION_GROUPSPointersEncoding.RELATION_GROUP {
		a_spec_relation_groups.RELATION_GROUP = append(a_spec_relation_groups.RELATION_GROUP, backRepo.BackRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr[uint(_RELATION_GROUPid)])
	}

	return
}

// CommitA_SPEC_RELATION_GROUPS allows commit of a single a_spec_relation_groups (if already staged)
func (backRepo *BackRepoStruct) CommitA_SPEC_RELATION_GROUPS(a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS) {
	backRepo.BackRepoA_SPEC_RELATION_GROUPS.CommitPhaseOneInstance(a_spec_relation_groups)
	if id, ok := backRepo.BackRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID[a_spec_relation_groups]; ok {
		backRepo.BackRepoA_SPEC_RELATION_GROUPS.CommitPhaseTwoInstance(backRepo, id, a_spec_relation_groups)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitA_SPEC_RELATION_GROUPS allows checkout of a single a_spec_relation_groups (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutA_SPEC_RELATION_GROUPS(a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS) {
	// check if the a_spec_relation_groups is staged
	if _, ok := backRepo.BackRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID[a_spec_relation_groups]; ok {

		if id, ok := backRepo.BackRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSPtr_A_SPEC_RELATION_GROUPSDBID[a_spec_relation_groups]; ok {
			var a_spec_relation_groupsDB A_SPEC_RELATION_GROUPSDB
			a_spec_relation_groupsDB.ID = id

			if _, err := backRepo.BackRepoA_SPEC_RELATION_GROUPS.db.First(&a_spec_relation_groupsDB, id); err != nil {
				log.Fatalln("CheckoutA_SPEC_RELATION_GROUPS : Problem with getting object with id:", id)
			}
			backRepo.BackRepoA_SPEC_RELATION_GROUPS.CheckoutPhaseOneInstance(&a_spec_relation_groupsDB)
			backRepo.BackRepoA_SPEC_RELATION_GROUPS.CheckoutPhaseTwoInstance(backRepo, &a_spec_relation_groupsDB)
		}
	}
}

// CopyBasicFieldsFromA_SPEC_RELATION_GROUPS
func (a_spec_relation_groupsDB *A_SPEC_RELATION_GROUPSDB) CopyBasicFieldsFromA_SPEC_RELATION_GROUPS(a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS) {
	// insertion point for fields commit

	a_spec_relation_groupsDB.Name_Data.String = a_spec_relation_groups.Name
	a_spec_relation_groupsDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromA_SPEC_RELATION_GROUPS_WOP
func (a_spec_relation_groupsDB *A_SPEC_RELATION_GROUPSDB) CopyBasicFieldsFromA_SPEC_RELATION_GROUPS_WOP(a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS_WOP) {
	// insertion point for fields commit

	a_spec_relation_groupsDB.Name_Data.String = a_spec_relation_groups.Name
	a_spec_relation_groupsDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromA_SPEC_RELATION_GROUPSWOP
func (a_spec_relation_groupsDB *A_SPEC_RELATION_GROUPSDB) CopyBasicFieldsFromA_SPEC_RELATION_GROUPSWOP(a_spec_relation_groups *A_SPEC_RELATION_GROUPSWOP) {
	// insertion point for fields commit

	a_spec_relation_groupsDB.Name_Data.String = a_spec_relation_groups.Name
	a_spec_relation_groupsDB.Name_Data.Valid = true
}

// CopyBasicFieldsToA_SPEC_RELATION_GROUPS
func (a_spec_relation_groupsDB *A_SPEC_RELATION_GROUPSDB) CopyBasicFieldsToA_SPEC_RELATION_GROUPS(a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS) {
	// insertion point for checkout of basic fields (back repo to stage)
	a_spec_relation_groups.Name = a_spec_relation_groupsDB.Name_Data.String
}

// CopyBasicFieldsToA_SPEC_RELATION_GROUPS_WOP
func (a_spec_relation_groupsDB *A_SPEC_RELATION_GROUPSDB) CopyBasicFieldsToA_SPEC_RELATION_GROUPS_WOP(a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	a_spec_relation_groups.Name = a_spec_relation_groupsDB.Name_Data.String
}

// CopyBasicFieldsToA_SPEC_RELATION_GROUPSWOP
func (a_spec_relation_groupsDB *A_SPEC_RELATION_GROUPSDB) CopyBasicFieldsToA_SPEC_RELATION_GROUPSWOP(a_spec_relation_groups *A_SPEC_RELATION_GROUPSWOP) {
	a_spec_relation_groups.ID = int(a_spec_relation_groupsDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	a_spec_relation_groups.Name = a_spec_relation_groupsDB.Name_Data.String
}

// Backup generates a json file from a slice of all A_SPEC_RELATION_GROUPSDB instances in the backrepo
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "A_SPEC_RELATION_GROUPSDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*A_SPEC_RELATION_GROUPSDB, 0)
	for _, a_spec_relation_groupsDB := range backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB {
		forBackup = append(forBackup, a_spec_relation_groupsDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json A_SPEC_RELATION_GROUPS ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json A_SPEC_RELATION_GROUPS file", err.Error())
	}
}

// Backup generates a json file from a slice of all A_SPEC_RELATION_GROUPSDB instances in the backrepo
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*A_SPEC_RELATION_GROUPSDB, 0)
	for _, a_spec_relation_groupsDB := range backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB {
		forBackup = append(forBackup, a_spec_relation_groupsDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("A_SPEC_RELATION_GROUPS")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&A_SPEC_RELATION_GROUPS_Fields, -1)
	for _, a_spec_relation_groupsDB := range forBackup {

		var a_spec_relation_groupsWOP A_SPEC_RELATION_GROUPSWOP
		a_spec_relation_groupsDB.CopyBasicFieldsToA_SPEC_RELATION_GROUPSWOP(&a_spec_relation_groupsWOP)

		row := sh.AddRow()
		row.WriteStruct(&a_spec_relation_groupsWOP, -1)
	}
}

// RestoreXL from the "A_SPEC_RELATION_GROUPS" sheet all A_SPEC_RELATION_GROUPSDB instances
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoA_SPEC_RELATION_GROUPSid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["A_SPEC_RELATION_GROUPS"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoA_SPEC_RELATION_GROUPS.rowVisitorA_SPEC_RELATION_GROUPS)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) rowVisitorA_SPEC_RELATION_GROUPS(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var a_spec_relation_groupsWOP A_SPEC_RELATION_GROUPSWOP
		row.ReadStruct(&a_spec_relation_groupsWOP)

		// add the unmarshalled struct to the stage
		a_spec_relation_groupsDB := new(A_SPEC_RELATION_GROUPSDB)
		a_spec_relation_groupsDB.CopyBasicFieldsFromA_SPEC_RELATION_GROUPSWOP(&a_spec_relation_groupsWOP)

		a_spec_relation_groupsDB_ID_atBackupTime := a_spec_relation_groupsDB.ID
		a_spec_relation_groupsDB.ID = 0
		_, err := backRepoA_SPEC_RELATION_GROUPS.db.Create(a_spec_relation_groupsDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB[a_spec_relation_groupsDB.ID] = a_spec_relation_groupsDB
		BackRepoA_SPEC_RELATION_GROUPSid_atBckpTime_newID[a_spec_relation_groupsDB_ID_atBackupTime] = a_spec_relation_groupsDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "A_SPEC_RELATION_GROUPSDB.json" in dirPath that stores an array
// of A_SPEC_RELATION_GROUPSDB and stores it in the database
// the map BackRepoA_SPEC_RELATION_GROUPSid_atBckpTime_newID is updated accordingly
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoA_SPEC_RELATION_GROUPSid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "A_SPEC_RELATION_GROUPSDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json A_SPEC_RELATION_GROUPS file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*A_SPEC_RELATION_GROUPSDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB
	for _, a_spec_relation_groupsDB := range forRestore {

		a_spec_relation_groupsDB_ID_atBackupTime := a_spec_relation_groupsDB.ID
		a_spec_relation_groupsDB.ID = 0
		_, err := backRepoA_SPEC_RELATION_GROUPS.db.Create(a_spec_relation_groupsDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB[a_spec_relation_groupsDB.ID] = a_spec_relation_groupsDB
		BackRepoA_SPEC_RELATION_GROUPSid_atBckpTime_newID[a_spec_relation_groupsDB_ID_atBackupTime] = a_spec_relation_groupsDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json A_SPEC_RELATION_GROUPS file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<A_SPEC_RELATION_GROUPS>id_atBckpTime_newID
// to compute new index
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) RestorePhaseTwo() {

	for _, a_spec_relation_groupsDB := range backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB {

		// next line of code is to avert unused variable compilation error
		_ = a_spec_relation_groupsDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoA_SPEC_RELATION_GROUPS.db.Model(a_spec_relation_groupsDB)
		_, err := db.Updates(*a_spec_relation_groupsDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoA_SPEC_RELATION_GROUPS.ResetReversePointers commits all staged instances of A_SPEC_RELATION_GROUPS to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, a_spec_relation_groups := range backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr {
		backRepoA_SPEC_RELATION_GROUPS.ResetReversePointersInstance(backRepo, idx, a_spec_relation_groups)
	}

	return
}

func (backRepoA_SPEC_RELATION_GROUPS *BackRepoA_SPEC_RELATION_GROUPSStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS) (Error error) {

	// fetch matching a_spec_relation_groupsDB
	if a_spec_relation_groupsDB, ok := backRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB[idx]; ok {
		_ = a_spec_relation_groupsDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoA_SPEC_RELATION_GROUPSid_atBckpTime_newID map[uint]uint
