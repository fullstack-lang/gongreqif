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
var dummy_A_SPEC_TYPES_sql sql.NullBool
var dummy_A_SPEC_TYPES_time time.Duration
var dummy_A_SPEC_TYPES_sort sort.Float64Slice

// A_SPEC_TYPESAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model a_spec_typesAPI
type A_SPEC_TYPESAPI struct {
	gorm.Model

	models.A_SPEC_TYPES_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	A_SPEC_TYPESPointersEncoding A_SPEC_TYPESPointersEncoding
}

// A_SPEC_TYPESPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type A_SPEC_TYPESPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field RELATION_GROUP_TYPE is a slice of pointers to another Struct (optional or 0..1)
	RELATION_GROUP_TYPE IntSlice `gorm:"type:TEXT"`

	// field SPEC_OBJECT_TYPE is a slice of pointers to another Struct (optional or 0..1)
	SPEC_OBJECT_TYPE IntSlice `gorm:"type:TEXT"`

	// field SPEC_RELATION_TYPE is a slice of pointers to another Struct (optional or 0..1)
	SPEC_RELATION_TYPE IntSlice `gorm:"type:TEXT"`

	// field SPECIFICATION_TYPE is a slice of pointers to another Struct (optional or 0..1)
	SPECIFICATION_TYPE IntSlice `gorm:"type:TEXT"`
}

// A_SPEC_TYPESDB describes a a_spec_types in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model a_spec_typesDB
type A_SPEC_TYPESDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field a_spec_typesDB.Name
	Name_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	A_SPEC_TYPESPointersEncoding
}

// A_SPEC_TYPESDBs arrays a_spec_typesDBs
// swagger:response a_spec_typesDBsResponse
type A_SPEC_TYPESDBs []A_SPEC_TYPESDB

// A_SPEC_TYPESDBResponse provides response
// swagger:response a_spec_typesDBResponse
type A_SPEC_TYPESDBResponse struct {
	A_SPEC_TYPESDB
}

// A_SPEC_TYPESWOP is a A_SPEC_TYPES without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type A_SPEC_TYPESWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var A_SPEC_TYPES_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoA_SPEC_TYPESStruct struct {
	// stores A_SPEC_TYPESDB according to their gorm ID
	Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB map[uint]*A_SPEC_TYPESDB

	// stores A_SPEC_TYPESDB ID according to A_SPEC_TYPES address
	Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID map[*models.A_SPEC_TYPES]uint

	// stores A_SPEC_TYPES according to their gorm ID
	Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr map[uint]*models.A_SPEC_TYPES

	db db.DBInterface

	stage *models.Stage
}

func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) GetStage() (stage *models.Stage) {
	stage = backRepoA_SPEC_TYPES.stage
	return
}

func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) GetDB() db.DBInterface {
	return backRepoA_SPEC_TYPES.db
}

// GetA_SPEC_TYPESDBFromA_SPEC_TYPESPtr is a handy function to access the back repo instance from the stage instance
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) GetA_SPEC_TYPESDBFromA_SPEC_TYPESPtr(a_spec_types *models.A_SPEC_TYPES) (a_spec_typesDB *A_SPEC_TYPESDB) {
	id := backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID[a_spec_types]
	a_spec_typesDB = backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB[id]
	return
}

// BackRepoA_SPEC_TYPES.CommitPhaseOne commits all staged instances of A_SPEC_TYPES to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var a_spec_typess []*models.A_SPEC_TYPES
	for a_spec_types := range stage.A_SPEC_TYPESs {
		a_spec_typess = append(a_spec_typess, a_spec_types)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(a_spec_typess, func(i, j int) bool {
		return stage.A_SPEC_TYPESMap_Staged_Order[a_spec_typess[i]] < stage.A_SPEC_TYPESMap_Staged_Order[a_spec_typess[j]]
	})

	for _, a_spec_types := range a_spec_typess {
		backRepoA_SPEC_TYPES.CommitPhaseOneInstance(a_spec_types)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, a_spec_types := range backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr {
		if _, ok := stage.A_SPEC_TYPESs[a_spec_types]; !ok {
			backRepoA_SPEC_TYPES.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoA_SPEC_TYPES.CommitDeleteInstance commits deletion of A_SPEC_TYPES to the BackRepo
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) CommitDeleteInstance(id uint) (Error error) {

	a_spec_types := backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr[id]

	// a_spec_types is not staged anymore, remove a_spec_typesDB
	a_spec_typesDB := backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB[id]
	db, _ := backRepoA_SPEC_TYPES.db.Unscoped()
	_, err := db.Delete(a_spec_typesDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID, a_spec_types)
	delete(backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr, id)
	delete(backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB, id)

	return
}

// BackRepoA_SPEC_TYPES.CommitPhaseOneInstance commits a_spec_types staged instances of A_SPEC_TYPES to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) CommitPhaseOneInstance(a_spec_types *models.A_SPEC_TYPES) (Error error) {

	// check if the a_spec_types is not commited yet
	if _, ok := backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID[a_spec_types]; ok {
		return
	}

	// initiate a_spec_types
	var a_spec_typesDB A_SPEC_TYPESDB
	a_spec_typesDB.CopyBasicFieldsFromA_SPEC_TYPES(a_spec_types)

	_, err := backRepoA_SPEC_TYPES.db.Create(&a_spec_typesDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID[a_spec_types] = a_spec_typesDB.ID
	backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr[a_spec_typesDB.ID] = a_spec_types
	backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB[a_spec_typesDB.ID] = &a_spec_typesDB

	return
}

// BackRepoA_SPEC_TYPES.CommitPhaseTwo commits all staged instances of A_SPEC_TYPES to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, a_spec_types := range backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr {
		backRepoA_SPEC_TYPES.CommitPhaseTwoInstance(backRepo, idx, a_spec_types)
	}

	return
}

// BackRepoA_SPEC_TYPES.CommitPhaseTwoInstance commits {{structname }} of models.A_SPEC_TYPES to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, a_spec_types *models.A_SPEC_TYPES) (Error error) {

	// fetch matching a_spec_typesDB
	if a_spec_typesDB, ok := backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB[idx]; ok {

		a_spec_typesDB.CopyBasicFieldsFromA_SPEC_TYPES(a_spec_types)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		a_spec_typesDB.A_SPEC_TYPESPointersEncoding.RELATION_GROUP_TYPE = make([]int, 0)
		// 2. encode
		for _, relation_group_typeAssocEnd := range a_spec_types.RELATION_GROUP_TYPE {
			relation_group_typeAssocEnd_DB :=
				backRepo.BackRepoRELATION_GROUP_TYPE.GetRELATION_GROUP_TYPEDBFromRELATION_GROUP_TYPEPtr(relation_group_typeAssocEnd)
			
			// the stage might be inconsistant, meaning that the relation_group_typeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if relation_group_typeAssocEnd_DB == nil {
				continue
			}
			
			a_spec_typesDB.A_SPEC_TYPESPointersEncoding.RELATION_GROUP_TYPE =
				append(a_spec_typesDB.A_SPEC_TYPESPointersEncoding.RELATION_GROUP_TYPE, int(relation_group_typeAssocEnd_DB.ID))
		}

		// 1. reset
		a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPEC_OBJECT_TYPE = make([]int, 0)
		// 2. encode
		for _, spec_object_typeAssocEnd := range a_spec_types.SPEC_OBJECT_TYPE {
			spec_object_typeAssocEnd_DB :=
				backRepo.BackRepoSPEC_OBJECT_TYPE.GetSPEC_OBJECT_TYPEDBFromSPEC_OBJECT_TYPEPtr(spec_object_typeAssocEnd)
			
			// the stage might be inconsistant, meaning that the spec_object_typeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if spec_object_typeAssocEnd_DB == nil {
				continue
			}
			
			a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPEC_OBJECT_TYPE =
				append(a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPEC_OBJECT_TYPE, int(spec_object_typeAssocEnd_DB.ID))
		}

		// 1. reset
		a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPEC_RELATION_TYPE = make([]int, 0)
		// 2. encode
		for _, spec_relation_typeAssocEnd := range a_spec_types.SPEC_RELATION_TYPE {
			spec_relation_typeAssocEnd_DB :=
				backRepo.BackRepoSPEC_RELATION_TYPE.GetSPEC_RELATION_TYPEDBFromSPEC_RELATION_TYPEPtr(spec_relation_typeAssocEnd)
			
			// the stage might be inconsistant, meaning that the spec_relation_typeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if spec_relation_typeAssocEnd_DB == nil {
				continue
			}
			
			a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPEC_RELATION_TYPE =
				append(a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPEC_RELATION_TYPE, int(spec_relation_typeAssocEnd_DB.ID))
		}

		// 1. reset
		a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPECIFICATION_TYPE = make([]int, 0)
		// 2. encode
		for _, specification_typeAssocEnd := range a_spec_types.SPECIFICATION_TYPE {
			specification_typeAssocEnd_DB :=
				backRepo.BackRepoSPECIFICATION_TYPE.GetSPECIFICATION_TYPEDBFromSPECIFICATION_TYPEPtr(specification_typeAssocEnd)
			
			// the stage might be inconsistant, meaning that the specification_typeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if specification_typeAssocEnd_DB == nil {
				continue
			}
			
			a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPECIFICATION_TYPE =
				append(a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPECIFICATION_TYPE, int(specification_typeAssocEnd_DB.ID))
		}

		_, err := backRepoA_SPEC_TYPES.db.Save(a_spec_typesDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown A_SPEC_TYPES intance %s", a_spec_types.Name))
		return err
	}

	return
}

// BackRepoA_SPEC_TYPES.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) CheckoutPhaseOne() (Error error) {

	a_spec_typesDBArray := make([]A_SPEC_TYPESDB, 0)
	_, err := backRepoA_SPEC_TYPES.db.Find(&a_spec_typesDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	a_spec_typesInstancesToBeRemovedFromTheStage := make(map[*models.A_SPEC_TYPES]any)
	for key, value := range backRepoA_SPEC_TYPES.stage.A_SPEC_TYPESs {
		a_spec_typesInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, a_spec_typesDB := range a_spec_typesDBArray {
		backRepoA_SPEC_TYPES.CheckoutPhaseOneInstance(&a_spec_typesDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		a_spec_types, ok := backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr[a_spec_typesDB.ID]
		if ok {
			delete(a_spec_typesInstancesToBeRemovedFromTheStage, a_spec_types)
		}
	}

	// remove from stage and back repo's 3 maps all a_spec_typess that are not in the checkout
	for a_spec_types := range a_spec_typesInstancesToBeRemovedFromTheStage {
		a_spec_types.Unstage(backRepoA_SPEC_TYPES.GetStage())

		// remove instance from the back repo 3 maps
		a_spec_typesID := backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID[a_spec_types]
		delete(backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID, a_spec_types)
		delete(backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB, a_spec_typesID)
		delete(backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr, a_spec_typesID)
	}

	return
}

// CheckoutPhaseOneInstance takes a a_spec_typesDB that has been found in the DB, updates the backRepo and stages the
// models version of the a_spec_typesDB
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) CheckoutPhaseOneInstance(a_spec_typesDB *A_SPEC_TYPESDB) (Error error) {

	a_spec_types, ok := backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr[a_spec_typesDB.ID]
	if !ok {
		a_spec_types = new(models.A_SPEC_TYPES)

		backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr[a_spec_typesDB.ID] = a_spec_types
		backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID[a_spec_types] = a_spec_typesDB.ID

		// append model store with the new element
		a_spec_types.Name = a_spec_typesDB.Name_Data.String
		a_spec_types.Stage(backRepoA_SPEC_TYPES.GetStage())
	}
	a_spec_typesDB.CopyBasicFieldsToA_SPEC_TYPES(a_spec_types)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	a_spec_types.Stage(backRepoA_SPEC_TYPES.GetStage())

	// preserve pointer to a_spec_typesDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB)[a_spec_typesDB hold variable pointers
	a_spec_typesDB_Data := *a_spec_typesDB
	preservedPtrToA_SPEC_TYPES := &a_spec_typesDB_Data
	backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB[a_spec_typesDB.ID] = preservedPtrToA_SPEC_TYPES

	return
}

// BackRepoA_SPEC_TYPES.CheckoutPhaseTwo Checkouts all staged instances of A_SPEC_TYPES to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, a_spec_typesDB := range backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB {
		backRepoA_SPEC_TYPES.CheckoutPhaseTwoInstance(backRepo, a_spec_typesDB)
	}
	return
}

// BackRepoA_SPEC_TYPES.CheckoutPhaseTwoInstance Checkouts staged instances of A_SPEC_TYPES to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, a_spec_typesDB *A_SPEC_TYPESDB) (Error error) {

	a_spec_types := backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr[a_spec_typesDB.ID]

	a_spec_typesDB.DecodePointers(backRepo, a_spec_types)

	return
}

func (a_spec_typesDB *A_SPEC_TYPESDB) DecodePointers(backRepo *BackRepoStruct, a_spec_types *models.A_SPEC_TYPES) {

	// insertion point for checkout of pointer encoding
	// This loop redeem a_spec_types.RELATION_GROUP_TYPE in the stage from the encode in the back repo
	// It parses all RELATION_GROUP_TYPEDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	a_spec_types.RELATION_GROUP_TYPE = a_spec_types.RELATION_GROUP_TYPE[:0]
	for _, _RELATION_GROUP_TYPEid := range a_spec_typesDB.A_SPEC_TYPESPointersEncoding.RELATION_GROUP_TYPE {
		a_spec_types.RELATION_GROUP_TYPE = append(a_spec_types.RELATION_GROUP_TYPE, backRepo.BackRepoRELATION_GROUP_TYPE.Map_RELATION_GROUP_TYPEDBID_RELATION_GROUP_TYPEPtr[uint(_RELATION_GROUP_TYPEid)])
	}

	// This loop redeem a_spec_types.SPEC_OBJECT_TYPE in the stage from the encode in the back repo
	// It parses all SPEC_OBJECT_TYPEDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	a_spec_types.SPEC_OBJECT_TYPE = a_spec_types.SPEC_OBJECT_TYPE[:0]
	for _, _SPEC_OBJECT_TYPEid := range a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPEC_OBJECT_TYPE {
		a_spec_types.SPEC_OBJECT_TYPE = append(a_spec_types.SPEC_OBJECT_TYPE, backRepo.BackRepoSPEC_OBJECT_TYPE.Map_SPEC_OBJECT_TYPEDBID_SPEC_OBJECT_TYPEPtr[uint(_SPEC_OBJECT_TYPEid)])
	}

	// This loop redeem a_spec_types.SPEC_RELATION_TYPE in the stage from the encode in the back repo
	// It parses all SPEC_RELATION_TYPEDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	a_spec_types.SPEC_RELATION_TYPE = a_spec_types.SPEC_RELATION_TYPE[:0]
	for _, _SPEC_RELATION_TYPEid := range a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPEC_RELATION_TYPE {
		a_spec_types.SPEC_RELATION_TYPE = append(a_spec_types.SPEC_RELATION_TYPE, backRepo.BackRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr[uint(_SPEC_RELATION_TYPEid)])
	}

	// This loop redeem a_spec_types.SPECIFICATION_TYPE in the stage from the encode in the back repo
	// It parses all SPECIFICATION_TYPEDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	a_spec_types.SPECIFICATION_TYPE = a_spec_types.SPECIFICATION_TYPE[:0]
	for _, _SPECIFICATION_TYPEid := range a_spec_typesDB.A_SPEC_TYPESPointersEncoding.SPECIFICATION_TYPE {
		a_spec_types.SPECIFICATION_TYPE = append(a_spec_types.SPECIFICATION_TYPE, backRepo.BackRepoSPECIFICATION_TYPE.Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEPtr[uint(_SPECIFICATION_TYPEid)])
	}

	return
}

// CommitA_SPEC_TYPES allows commit of a single a_spec_types (if already staged)
func (backRepo *BackRepoStruct) CommitA_SPEC_TYPES(a_spec_types *models.A_SPEC_TYPES) {
	backRepo.BackRepoA_SPEC_TYPES.CommitPhaseOneInstance(a_spec_types)
	if id, ok := backRepo.BackRepoA_SPEC_TYPES.Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID[a_spec_types]; ok {
		backRepo.BackRepoA_SPEC_TYPES.CommitPhaseTwoInstance(backRepo, id, a_spec_types)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitA_SPEC_TYPES allows checkout of a single a_spec_types (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutA_SPEC_TYPES(a_spec_types *models.A_SPEC_TYPES) {
	// check if the a_spec_types is staged
	if _, ok := backRepo.BackRepoA_SPEC_TYPES.Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID[a_spec_types]; ok {

		if id, ok := backRepo.BackRepoA_SPEC_TYPES.Map_A_SPEC_TYPESPtr_A_SPEC_TYPESDBID[a_spec_types]; ok {
			var a_spec_typesDB A_SPEC_TYPESDB
			a_spec_typesDB.ID = id

			if _, err := backRepo.BackRepoA_SPEC_TYPES.db.First(&a_spec_typesDB, id); err != nil {
				log.Fatalln("CheckoutA_SPEC_TYPES : Problem with getting object with id:", id)
			}
			backRepo.BackRepoA_SPEC_TYPES.CheckoutPhaseOneInstance(&a_spec_typesDB)
			backRepo.BackRepoA_SPEC_TYPES.CheckoutPhaseTwoInstance(backRepo, &a_spec_typesDB)
		}
	}
}

// CopyBasicFieldsFromA_SPEC_TYPES
func (a_spec_typesDB *A_SPEC_TYPESDB) CopyBasicFieldsFromA_SPEC_TYPES(a_spec_types *models.A_SPEC_TYPES) {
	// insertion point for fields commit

	a_spec_typesDB.Name_Data.String = a_spec_types.Name
	a_spec_typesDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromA_SPEC_TYPES_WOP
func (a_spec_typesDB *A_SPEC_TYPESDB) CopyBasicFieldsFromA_SPEC_TYPES_WOP(a_spec_types *models.A_SPEC_TYPES_WOP) {
	// insertion point for fields commit

	a_spec_typesDB.Name_Data.String = a_spec_types.Name
	a_spec_typesDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromA_SPEC_TYPESWOP
func (a_spec_typesDB *A_SPEC_TYPESDB) CopyBasicFieldsFromA_SPEC_TYPESWOP(a_spec_types *A_SPEC_TYPESWOP) {
	// insertion point for fields commit

	a_spec_typesDB.Name_Data.String = a_spec_types.Name
	a_spec_typesDB.Name_Data.Valid = true
}

// CopyBasicFieldsToA_SPEC_TYPES
func (a_spec_typesDB *A_SPEC_TYPESDB) CopyBasicFieldsToA_SPEC_TYPES(a_spec_types *models.A_SPEC_TYPES) {
	// insertion point for checkout of basic fields (back repo to stage)
	a_spec_types.Name = a_spec_typesDB.Name_Data.String
}

// CopyBasicFieldsToA_SPEC_TYPES_WOP
func (a_spec_typesDB *A_SPEC_TYPESDB) CopyBasicFieldsToA_SPEC_TYPES_WOP(a_spec_types *models.A_SPEC_TYPES_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	a_spec_types.Name = a_spec_typesDB.Name_Data.String
}

// CopyBasicFieldsToA_SPEC_TYPESWOP
func (a_spec_typesDB *A_SPEC_TYPESDB) CopyBasicFieldsToA_SPEC_TYPESWOP(a_spec_types *A_SPEC_TYPESWOP) {
	a_spec_types.ID = int(a_spec_typesDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	a_spec_types.Name = a_spec_typesDB.Name_Data.String
}

// Backup generates a json file from a slice of all A_SPEC_TYPESDB instances in the backrepo
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "A_SPEC_TYPESDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*A_SPEC_TYPESDB, 0)
	for _, a_spec_typesDB := range backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB {
		forBackup = append(forBackup, a_spec_typesDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json A_SPEC_TYPES ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json A_SPEC_TYPES file", err.Error())
	}
}

// Backup generates a json file from a slice of all A_SPEC_TYPESDB instances in the backrepo
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*A_SPEC_TYPESDB, 0)
	for _, a_spec_typesDB := range backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB {
		forBackup = append(forBackup, a_spec_typesDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("A_SPEC_TYPES")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&A_SPEC_TYPES_Fields, -1)
	for _, a_spec_typesDB := range forBackup {

		var a_spec_typesWOP A_SPEC_TYPESWOP
		a_spec_typesDB.CopyBasicFieldsToA_SPEC_TYPESWOP(&a_spec_typesWOP)

		row := sh.AddRow()
		row.WriteStruct(&a_spec_typesWOP, -1)
	}
}

// RestoreXL from the "A_SPEC_TYPES" sheet all A_SPEC_TYPESDB instances
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoA_SPEC_TYPESid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["A_SPEC_TYPES"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoA_SPEC_TYPES.rowVisitorA_SPEC_TYPES)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) rowVisitorA_SPEC_TYPES(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var a_spec_typesWOP A_SPEC_TYPESWOP
		row.ReadStruct(&a_spec_typesWOP)

		// add the unmarshalled struct to the stage
		a_spec_typesDB := new(A_SPEC_TYPESDB)
		a_spec_typesDB.CopyBasicFieldsFromA_SPEC_TYPESWOP(&a_spec_typesWOP)

		a_spec_typesDB_ID_atBackupTime := a_spec_typesDB.ID
		a_spec_typesDB.ID = 0
		_, err := backRepoA_SPEC_TYPES.db.Create(a_spec_typesDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB[a_spec_typesDB.ID] = a_spec_typesDB
		BackRepoA_SPEC_TYPESid_atBckpTime_newID[a_spec_typesDB_ID_atBackupTime] = a_spec_typesDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "A_SPEC_TYPESDB.json" in dirPath that stores an array
// of A_SPEC_TYPESDB and stores it in the database
// the map BackRepoA_SPEC_TYPESid_atBckpTime_newID is updated accordingly
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoA_SPEC_TYPESid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "A_SPEC_TYPESDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json A_SPEC_TYPES file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*A_SPEC_TYPESDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB
	for _, a_spec_typesDB := range forRestore {

		a_spec_typesDB_ID_atBackupTime := a_spec_typesDB.ID
		a_spec_typesDB.ID = 0
		_, err := backRepoA_SPEC_TYPES.db.Create(a_spec_typesDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB[a_spec_typesDB.ID] = a_spec_typesDB
		BackRepoA_SPEC_TYPESid_atBckpTime_newID[a_spec_typesDB_ID_atBackupTime] = a_spec_typesDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json A_SPEC_TYPES file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<A_SPEC_TYPES>id_atBckpTime_newID
// to compute new index
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) RestorePhaseTwo() {

	for _, a_spec_typesDB := range backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB {

		// next line of code is to avert unused variable compilation error
		_ = a_spec_typesDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoA_SPEC_TYPES.db.Model(a_spec_typesDB)
		_, err := db.Updates(*a_spec_typesDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoA_SPEC_TYPES.ResetReversePointers commits all staged instances of A_SPEC_TYPES to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, a_spec_types := range backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr {
		backRepoA_SPEC_TYPES.ResetReversePointersInstance(backRepo, idx, a_spec_types)
	}

	return
}

func (backRepoA_SPEC_TYPES *BackRepoA_SPEC_TYPESStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, a_spec_types *models.A_SPEC_TYPES) (Error error) {

	// fetch matching a_spec_typesDB
	if a_spec_typesDB, ok := backRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB[idx]; ok {
		_ = a_spec_typesDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoA_SPEC_TYPESid_atBckpTime_newID map[uint]uint
