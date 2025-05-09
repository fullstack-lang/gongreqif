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
var dummy_RELATION_GROUP_sql sql.NullBool
var dummy_RELATION_GROUP_time time.Duration
var dummy_RELATION_GROUP_sort sort.Float64Slice

// RELATION_GROUPAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model relation_groupAPI
type RELATION_GROUPAPI struct {
	gorm.Model

	models.RELATION_GROUP_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	RELATION_GROUPPointersEncoding RELATION_GROUPPointersEncoding
}

// RELATION_GROUPPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type RELATION_GROUPPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	ALTERNATIVE_ID struct {

		// field ALTERNATIVE_ID is a slice of pointers to another Struct (optional or 0..1)
		ALTERNATIVE_ID IntSlice `gorm:"type:TEXT"`

	} `gorm:"embedded"`
}

// RELATION_GROUPDB describes a relation_group in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model relation_groupDB
type RELATION_GROUPDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field relation_groupDB.Name
	Name_Data sql.NullString

	// Declation for basic field relation_groupDB.DESC
	DESC_Data sql.NullString

	// Declation for basic field relation_groupDB.LAST_CHANGE
	LAST_CHANGE_Data sql.NullTime

	// Declation for basic field relation_groupDB.LONG_NAME
	LONG_NAME_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	RELATION_GROUPPointersEncoding
}

// RELATION_GROUPDBs arrays relation_groupDBs
// swagger:response relation_groupDBsResponse
type RELATION_GROUPDBs []RELATION_GROUPDB

// RELATION_GROUPDBResponse provides response
// swagger:response relation_groupDBResponse
type RELATION_GROUPDBResponse struct {
	RELATION_GROUPDB
}

// RELATION_GROUPWOP is a RELATION_GROUP without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type RELATION_GROUPWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESC string `xlsx:"2"`

	LAST_CHANGE time.Time `xlsx:"3"`

	LONG_NAME string `xlsx:"4"`
	// insertion for WOP pointer fields
}

var RELATION_GROUP_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESC",
	"LAST_CHANGE",
	"LONG_NAME",
}

type BackRepoRELATION_GROUPStruct struct {
	// stores RELATION_GROUPDB according to their gorm ID
	Map_RELATION_GROUPDBID_RELATION_GROUPDB map[uint]*RELATION_GROUPDB

	// stores RELATION_GROUPDB ID according to RELATION_GROUP address
	Map_RELATION_GROUPPtr_RELATION_GROUPDBID map[*models.RELATION_GROUP]uint

	// stores RELATION_GROUP according to their gorm ID
	Map_RELATION_GROUPDBID_RELATION_GROUPPtr map[uint]*models.RELATION_GROUP

	db db.DBInterface

	stage *models.Stage
}

func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) GetStage() (stage *models.Stage) {
	stage = backRepoRELATION_GROUP.stage
	return
}

func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) GetDB() db.DBInterface {
	return backRepoRELATION_GROUP.db
}

// GetRELATION_GROUPDBFromRELATION_GROUPPtr is a handy function to access the back repo instance from the stage instance
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) GetRELATION_GROUPDBFromRELATION_GROUPPtr(relation_group *models.RELATION_GROUP) (relation_groupDB *RELATION_GROUPDB) {
	id := backRepoRELATION_GROUP.Map_RELATION_GROUPPtr_RELATION_GROUPDBID[relation_group]
	relation_groupDB = backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB[id]
	return
}

// BackRepoRELATION_GROUP.CommitPhaseOne commits all staged instances of RELATION_GROUP to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var relation_groups []*models.RELATION_GROUP
	for relation_group := range stage.RELATION_GROUPs {
		relation_groups = append(relation_groups, relation_group)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(relation_groups, func(i, j int) bool {
		return stage.RELATION_GROUPMap_Staged_Order[relation_groups[i]] < stage.RELATION_GROUPMap_Staged_Order[relation_groups[j]]
	})

	for _, relation_group := range relation_groups {
		backRepoRELATION_GROUP.CommitPhaseOneInstance(relation_group)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, relation_group := range backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr {
		if _, ok := stage.RELATION_GROUPs[relation_group]; !ok {
			backRepoRELATION_GROUP.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoRELATION_GROUP.CommitDeleteInstance commits deletion of RELATION_GROUP to the BackRepo
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) CommitDeleteInstance(id uint) (Error error) {

	relation_group := backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr[id]

	// relation_group is not staged anymore, remove relation_groupDB
	relation_groupDB := backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB[id]
	db, _ := backRepoRELATION_GROUP.db.Unscoped()
	_, err := db.Delete(relation_groupDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoRELATION_GROUP.Map_RELATION_GROUPPtr_RELATION_GROUPDBID, relation_group)
	delete(backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr, id)
	delete(backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB, id)

	return
}

// BackRepoRELATION_GROUP.CommitPhaseOneInstance commits relation_group staged instances of RELATION_GROUP to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) CommitPhaseOneInstance(relation_group *models.RELATION_GROUP) (Error error) {

	// check if the relation_group is not commited yet
	if _, ok := backRepoRELATION_GROUP.Map_RELATION_GROUPPtr_RELATION_GROUPDBID[relation_group]; ok {
		return
	}

	// initiate relation_group
	var relation_groupDB RELATION_GROUPDB
	relation_groupDB.CopyBasicFieldsFromRELATION_GROUP(relation_group)

	_, err := backRepoRELATION_GROUP.db.Create(&relation_groupDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoRELATION_GROUP.Map_RELATION_GROUPPtr_RELATION_GROUPDBID[relation_group] = relation_groupDB.ID
	backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr[relation_groupDB.ID] = relation_group
	backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB[relation_groupDB.ID] = &relation_groupDB

	return
}

// BackRepoRELATION_GROUP.CommitPhaseTwo commits all staged instances of RELATION_GROUP to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, relation_group := range backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr {
		backRepoRELATION_GROUP.CommitPhaseTwoInstance(backRepo, idx, relation_group)
	}

	return
}

// BackRepoRELATION_GROUP.CommitPhaseTwoInstance commits {{structname }} of models.RELATION_GROUP to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, relation_group *models.RELATION_GROUP) (Error error) {

	// fetch matching relation_groupDB
	if relation_groupDB, ok := backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB[idx]; ok {

		relation_groupDB.CopyBasicFieldsFromRELATION_GROUP(relation_group)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		relation_groupDB.RELATION_GROUPPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID = make([]int, 0)
		// 2. encode
		for _, alternative_idAssocEnd := range relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID {
			alternative_idAssocEnd_DB :=
				backRepo.BackRepoALTERNATIVE_ID.GetALTERNATIVE_IDDBFromALTERNATIVE_IDPtr(alternative_idAssocEnd)
			
			// the stage might be inconsistant, meaning that the alternative_idAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if alternative_idAssocEnd_DB == nil {
				continue
			}
			
			relation_groupDB.RELATION_GROUPPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID =
				append(relation_groupDB.RELATION_GROUPPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID, int(alternative_idAssocEnd_DB.ID))
		}

		_, err := backRepoRELATION_GROUP.db.Save(relation_groupDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown RELATION_GROUP intance %s", relation_group.Name))
		return err
	}

	return
}

// BackRepoRELATION_GROUP.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) CheckoutPhaseOne() (Error error) {

	relation_groupDBArray := make([]RELATION_GROUPDB, 0)
	_, err := backRepoRELATION_GROUP.db.Find(&relation_groupDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	relation_groupInstancesToBeRemovedFromTheStage := make(map[*models.RELATION_GROUP]any)
	for key, value := range backRepoRELATION_GROUP.stage.RELATION_GROUPs {
		relation_groupInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, relation_groupDB := range relation_groupDBArray {
		backRepoRELATION_GROUP.CheckoutPhaseOneInstance(&relation_groupDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		relation_group, ok := backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr[relation_groupDB.ID]
		if ok {
			delete(relation_groupInstancesToBeRemovedFromTheStage, relation_group)
		}
	}

	// remove from stage and back repo's 3 maps all relation_groups that are not in the checkout
	for relation_group := range relation_groupInstancesToBeRemovedFromTheStage {
		relation_group.Unstage(backRepoRELATION_GROUP.GetStage())

		// remove instance from the back repo 3 maps
		relation_groupID := backRepoRELATION_GROUP.Map_RELATION_GROUPPtr_RELATION_GROUPDBID[relation_group]
		delete(backRepoRELATION_GROUP.Map_RELATION_GROUPPtr_RELATION_GROUPDBID, relation_group)
		delete(backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB, relation_groupID)
		delete(backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr, relation_groupID)
	}

	return
}

// CheckoutPhaseOneInstance takes a relation_groupDB that has been found in the DB, updates the backRepo and stages the
// models version of the relation_groupDB
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) CheckoutPhaseOneInstance(relation_groupDB *RELATION_GROUPDB) (Error error) {

	relation_group, ok := backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr[relation_groupDB.ID]
	if !ok {
		relation_group = new(models.RELATION_GROUP)

		backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr[relation_groupDB.ID] = relation_group
		backRepoRELATION_GROUP.Map_RELATION_GROUPPtr_RELATION_GROUPDBID[relation_group] = relation_groupDB.ID

		// append model store with the new element
		relation_group.Name = relation_groupDB.Name_Data.String
		relation_group.Stage(backRepoRELATION_GROUP.GetStage())
	}
	relation_groupDB.CopyBasicFieldsToRELATION_GROUP(relation_group)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	relation_group.Stage(backRepoRELATION_GROUP.GetStage())

	// preserve pointer to relation_groupDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_RELATION_GROUPDBID_RELATION_GROUPDB)[relation_groupDB hold variable pointers
	relation_groupDB_Data := *relation_groupDB
	preservedPtrToRELATION_GROUP := &relation_groupDB_Data
	backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB[relation_groupDB.ID] = preservedPtrToRELATION_GROUP

	return
}

// BackRepoRELATION_GROUP.CheckoutPhaseTwo Checkouts all staged instances of RELATION_GROUP to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, relation_groupDB := range backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB {
		backRepoRELATION_GROUP.CheckoutPhaseTwoInstance(backRepo, relation_groupDB)
	}
	return
}

// BackRepoRELATION_GROUP.CheckoutPhaseTwoInstance Checkouts staged instances of RELATION_GROUP to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, relation_groupDB *RELATION_GROUPDB) (Error error) {

	relation_group := backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr[relation_groupDB.ID]

	relation_groupDB.DecodePointers(backRepo, relation_group)

	return
}

func (relation_groupDB *RELATION_GROUPDB) DecodePointers(backRepo *BackRepoStruct, relation_group *models.RELATION_GROUP) {

	// insertion point for checkout of pointer encoding
	// This loop redeem relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID in the stage from the encode in the back repo
	// It parses all ALTERNATIVE_IDDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID = relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID[:0]
	for _, _ALTERNATIVE_IDid := range relation_groupDB.RELATION_GROUPPointersEncoding.ALTERNATIVE_ID.ALTERNATIVE_ID {
		relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID = append(relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID, backRepo.BackRepoALTERNATIVE_ID.Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDPtr[uint(_ALTERNATIVE_IDid)])
	}

	return
}

// CommitRELATION_GROUP allows commit of a single relation_group (if already staged)
func (backRepo *BackRepoStruct) CommitRELATION_GROUP(relation_group *models.RELATION_GROUP) {
	backRepo.BackRepoRELATION_GROUP.CommitPhaseOneInstance(relation_group)
	if id, ok := backRepo.BackRepoRELATION_GROUP.Map_RELATION_GROUPPtr_RELATION_GROUPDBID[relation_group]; ok {
		backRepo.BackRepoRELATION_GROUP.CommitPhaseTwoInstance(backRepo, id, relation_group)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitRELATION_GROUP allows checkout of a single relation_group (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutRELATION_GROUP(relation_group *models.RELATION_GROUP) {
	// check if the relation_group is staged
	if _, ok := backRepo.BackRepoRELATION_GROUP.Map_RELATION_GROUPPtr_RELATION_GROUPDBID[relation_group]; ok {

		if id, ok := backRepo.BackRepoRELATION_GROUP.Map_RELATION_GROUPPtr_RELATION_GROUPDBID[relation_group]; ok {
			var relation_groupDB RELATION_GROUPDB
			relation_groupDB.ID = id

			if _, err := backRepo.BackRepoRELATION_GROUP.db.First(&relation_groupDB, id); err != nil {
				log.Fatalln("CheckoutRELATION_GROUP : Problem with getting object with id:", id)
			}
			backRepo.BackRepoRELATION_GROUP.CheckoutPhaseOneInstance(&relation_groupDB)
			backRepo.BackRepoRELATION_GROUP.CheckoutPhaseTwoInstance(backRepo, &relation_groupDB)
		}
	}
}

// CopyBasicFieldsFromRELATION_GROUP
func (relation_groupDB *RELATION_GROUPDB) CopyBasicFieldsFromRELATION_GROUP(relation_group *models.RELATION_GROUP) {
	// insertion point for fields commit

	relation_groupDB.Name_Data.String = relation_group.Name
	relation_groupDB.Name_Data.Valid = true

	relation_groupDB.DESC_Data.String = relation_group.DESC
	relation_groupDB.DESC_Data.Valid = true

	relation_groupDB.LAST_CHANGE_Data.Time = relation_group.LAST_CHANGE
	relation_groupDB.LAST_CHANGE_Data.Valid = true

	relation_groupDB.LONG_NAME_Data.String = relation_group.LONG_NAME
	relation_groupDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromRELATION_GROUP_WOP
func (relation_groupDB *RELATION_GROUPDB) CopyBasicFieldsFromRELATION_GROUP_WOP(relation_group *models.RELATION_GROUP_WOP) {
	// insertion point for fields commit

	relation_groupDB.Name_Data.String = relation_group.Name
	relation_groupDB.Name_Data.Valid = true

	relation_groupDB.DESC_Data.String = relation_group.DESC
	relation_groupDB.DESC_Data.Valid = true

	relation_groupDB.LAST_CHANGE_Data.Time = relation_group.LAST_CHANGE
	relation_groupDB.LAST_CHANGE_Data.Valid = true

	relation_groupDB.LONG_NAME_Data.String = relation_group.LONG_NAME
	relation_groupDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromRELATION_GROUPWOP
func (relation_groupDB *RELATION_GROUPDB) CopyBasicFieldsFromRELATION_GROUPWOP(relation_group *RELATION_GROUPWOP) {
	// insertion point for fields commit

	relation_groupDB.Name_Data.String = relation_group.Name
	relation_groupDB.Name_Data.Valid = true

	relation_groupDB.DESC_Data.String = relation_group.DESC
	relation_groupDB.DESC_Data.Valid = true

	relation_groupDB.LAST_CHANGE_Data.Time = relation_group.LAST_CHANGE
	relation_groupDB.LAST_CHANGE_Data.Valid = true

	relation_groupDB.LONG_NAME_Data.String = relation_group.LONG_NAME
	relation_groupDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsToRELATION_GROUP
func (relation_groupDB *RELATION_GROUPDB) CopyBasicFieldsToRELATION_GROUP(relation_group *models.RELATION_GROUP) {
	// insertion point for checkout of basic fields (back repo to stage)
	relation_group.Name = relation_groupDB.Name_Data.String
	relation_group.DESC = relation_groupDB.DESC_Data.String
	relation_group.LAST_CHANGE = relation_groupDB.LAST_CHANGE_Data.Time
	relation_group.LONG_NAME = relation_groupDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToRELATION_GROUP_WOP
func (relation_groupDB *RELATION_GROUPDB) CopyBasicFieldsToRELATION_GROUP_WOP(relation_group *models.RELATION_GROUP_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	relation_group.Name = relation_groupDB.Name_Data.String
	relation_group.DESC = relation_groupDB.DESC_Data.String
	relation_group.LAST_CHANGE = relation_groupDB.LAST_CHANGE_Data.Time
	relation_group.LONG_NAME = relation_groupDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToRELATION_GROUPWOP
func (relation_groupDB *RELATION_GROUPDB) CopyBasicFieldsToRELATION_GROUPWOP(relation_group *RELATION_GROUPWOP) {
	relation_group.ID = int(relation_groupDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	relation_group.Name = relation_groupDB.Name_Data.String
	relation_group.DESC = relation_groupDB.DESC_Data.String
	relation_group.LAST_CHANGE = relation_groupDB.LAST_CHANGE_Data.Time
	relation_group.LONG_NAME = relation_groupDB.LONG_NAME_Data.String
}

// Backup generates a json file from a slice of all RELATION_GROUPDB instances in the backrepo
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "RELATION_GROUPDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*RELATION_GROUPDB, 0)
	for _, relation_groupDB := range backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB {
		forBackup = append(forBackup, relation_groupDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json RELATION_GROUP ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json RELATION_GROUP file", err.Error())
	}
}

// Backup generates a json file from a slice of all RELATION_GROUPDB instances in the backrepo
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*RELATION_GROUPDB, 0)
	for _, relation_groupDB := range backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB {
		forBackup = append(forBackup, relation_groupDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("RELATION_GROUP")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&RELATION_GROUP_Fields, -1)
	for _, relation_groupDB := range forBackup {

		var relation_groupWOP RELATION_GROUPWOP
		relation_groupDB.CopyBasicFieldsToRELATION_GROUPWOP(&relation_groupWOP)

		row := sh.AddRow()
		row.WriteStruct(&relation_groupWOP, -1)
	}
}

// RestoreXL from the "RELATION_GROUP" sheet all RELATION_GROUPDB instances
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoRELATION_GROUPid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["RELATION_GROUP"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoRELATION_GROUP.rowVisitorRELATION_GROUP)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) rowVisitorRELATION_GROUP(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var relation_groupWOP RELATION_GROUPWOP
		row.ReadStruct(&relation_groupWOP)

		// add the unmarshalled struct to the stage
		relation_groupDB := new(RELATION_GROUPDB)
		relation_groupDB.CopyBasicFieldsFromRELATION_GROUPWOP(&relation_groupWOP)

		relation_groupDB_ID_atBackupTime := relation_groupDB.ID
		relation_groupDB.ID = 0
		_, err := backRepoRELATION_GROUP.db.Create(relation_groupDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB[relation_groupDB.ID] = relation_groupDB
		BackRepoRELATION_GROUPid_atBckpTime_newID[relation_groupDB_ID_atBackupTime] = relation_groupDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "RELATION_GROUPDB.json" in dirPath that stores an array
// of RELATION_GROUPDB and stores it in the database
// the map BackRepoRELATION_GROUPid_atBckpTime_newID is updated accordingly
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoRELATION_GROUPid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "RELATION_GROUPDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json RELATION_GROUP file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*RELATION_GROUPDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_RELATION_GROUPDBID_RELATION_GROUPDB
	for _, relation_groupDB := range forRestore {

		relation_groupDB_ID_atBackupTime := relation_groupDB.ID
		relation_groupDB.ID = 0
		_, err := backRepoRELATION_GROUP.db.Create(relation_groupDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB[relation_groupDB.ID] = relation_groupDB
		BackRepoRELATION_GROUPid_atBckpTime_newID[relation_groupDB_ID_atBackupTime] = relation_groupDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json RELATION_GROUP file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<RELATION_GROUP>id_atBckpTime_newID
// to compute new index
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) RestorePhaseTwo() {

	for _, relation_groupDB := range backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB {

		// next line of code is to avert unused variable compilation error
		_ = relation_groupDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoRELATION_GROUP.db.Model(relation_groupDB)
		_, err := db.Updates(*relation_groupDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoRELATION_GROUP.ResetReversePointers commits all staged instances of RELATION_GROUP to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, relation_group := range backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr {
		backRepoRELATION_GROUP.ResetReversePointersInstance(backRepo, idx, relation_group)
	}

	return
}

func (backRepoRELATION_GROUP *BackRepoRELATION_GROUPStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, relation_group *models.RELATION_GROUP) (Error error) {

	// fetch matching relation_groupDB
	if relation_groupDB, ok := backRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB[idx]; ok {
		_ = relation_groupDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoRELATION_GROUPid_atBckpTime_newID map[uint]uint
