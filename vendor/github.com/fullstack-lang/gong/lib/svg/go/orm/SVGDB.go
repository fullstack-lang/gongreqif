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

	"github.com/fullstack-lang/gong/lib/svg/go/db"
	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_SVG_sql sql.NullBool
var dummy_SVG_time time.Duration
var dummy_SVG_sort sort.Float64Slice

// SVGAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model svgAPI
type SVGAPI struct {
	gorm.Model

	models.SVG_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SVGPointersEncoding SVGPointersEncoding
}

// SVGPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SVGPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Layers is a slice of pointers to another Struct (optional or 0..1)
	Layers IntSlice `gorm:"type:TEXT"`

	// field StartRect is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	StartRectID sql.NullInt64

	// field EndRect is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	EndRectID sql.NullInt64
}

// SVGDB describes a svg in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model svgDB
type SVGDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field svgDB.Name
	Name_Data sql.NullString

	// Declation for basic field svgDB.DrawingState
	DrawingState_Data sql.NullString

	// Declation for basic field svgDB.IsEditable
	// provide the sql storage for the boolan
	IsEditable_Data sql.NullBool

	// Declation for basic field svgDB.IsSVGFrontEndFileGenerated
	// provide the sql storage for the boolan
	IsSVGFrontEndFileGenerated_Data sql.NullBool

	// Declation for basic field svgDB.IsSVGBackEndFileGenerated
	// provide the sql storage for the boolan
	IsSVGBackEndFileGenerated_Data sql.NullBool

	// Declation for basic field svgDB.DefaultDirectoryForGeneratedImages
	DefaultDirectoryForGeneratedImages_Data sql.NullString

	// Declation for basic field svgDB.IsControlBannerHidden
	// provide the sql storage for the boolan
	IsControlBannerHidden_Data sql.NullBool

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SVGPointersEncoding
}

// SVGDBs arrays svgDBs
// swagger:response svgDBsResponse
type SVGDBs []SVGDB

// SVGDBResponse provides response
// swagger:response svgDBResponse
type SVGDBResponse struct {
	SVGDB
}

// SVGWOP is a SVG without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SVGWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DrawingState models.DrawingState `xlsx:"2"`

	IsEditable bool `xlsx:"3"`

	IsSVGFrontEndFileGenerated bool `xlsx:"4"`

	IsSVGBackEndFileGenerated bool `xlsx:"5"`

	DefaultDirectoryForGeneratedImages string `xlsx:"6"`

	IsControlBannerHidden bool `xlsx:"7"`
	// insertion for WOP pointer fields
}

var SVG_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DrawingState",
	"IsEditable",
	"IsSVGFrontEndFileGenerated",
	"IsSVGBackEndFileGenerated",
	"DefaultDirectoryForGeneratedImages",
	"IsControlBannerHidden",
}

type BackRepoSVGStruct struct {
	// stores SVGDB according to their gorm ID
	Map_SVGDBID_SVGDB map[uint]*SVGDB

	// stores SVGDB ID according to SVG address
	Map_SVGPtr_SVGDBID map[*models.SVG]uint

	// stores SVG according to their gorm ID
	Map_SVGDBID_SVGPtr map[uint]*models.SVG

	db db.DBInterface

	stage *models.Stage
}

func (backRepoSVG *BackRepoSVGStruct) GetStage() (stage *models.Stage) {
	stage = backRepoSVG.stage
	return
}

func (backRepoSVG *BackRepoSVGStruct) GetDB() db.DBInterface {
	return backRepoSVG.db
}

// GetSVGDBFromSVGPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSVG *BackRepoSVGStruct) GetSVGDBFromSVGPtr(svg *models.SVG) (svgDB *SVGDB) {
	id := backRepoSVG.Map_SVGPtr_SVGDBID[svg]
	svgDB = backRepoSVG.Map_SVGDBID_SVGDB[id]
	return
}

// BackRepoSVG.CommitPhaseOne commits all staged instances of SVG to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSVG *BackRepoSVGStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var svgs []*models.SVG
	for svg := range stage.SVGs {
		svgs = append(svgs, svg)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(svgs, func(i, j int) bool {
		return stage.SVGMap_Staged_Order[svgs[i]] < stage.SVGMap_Staged_Order[svgs[j]]
	})

	for _, svg := range svgs {
		backRepoSVG.CommitPhaseOneInstance(svg)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, svg := range backRepoSVG.Map_SVGDBID_SVGPtr {
		if _, ok := stage.SVGs[svg]; !ok {
			backRepoSVG.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSVG.CommitDeleteInstance commits deletion of SVG to the BackRepo
func (backRepoSVG *BackRepoSVGStruct) CommitDeleteInstance(id uint) (Error error) {

	svg := backRepoSVG.Map_SVGDBID_SVGPtr[id]

	// svg is not staged anymore, remove svgDB
	svgDB := backRepoSVG.Map_SVGDBID_SVGDB[id]
	db, _ := backRepoSVG.db.Unscoped()
	_, err := db.Delete(svgDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoSVG.Map_SVGPtr_SVGDBID, svg)
	delete(backRepoSVG.Map_SVGDBID_SVGPtr, id)
	delete(backRepoSVG.Map_SVGDBID_SVGDB, id)

	return
}

// BackRepoSVG.CommitPhaseOneInstance commits svg staged instances of SVG to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSVG *BackRepoSVGStruct) CommitPhaseOneInstance(svg *models.SVG) (Error error) {

	// check if the svg is not commited yet
	if _, ok := backRepoSVG.Map_SVGPtr_SVGDBID[svg]; ok {
		return
	}

	// initiate svg
	var svgDB SVGDB
	svgDB.CopyBasicFieldsFromSVG(svg)

	_, err := backRepoSVG.db.Create(&svgDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoSVG.Map_SVGPtr_SVGDBID[svg] = svgDB.ID
	backRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID] = svg
	backRepoSVG.Map_SVGDBID_SVGDB[svgDB.ID] = &svgDB

	return
}

// BackRepoSVG.CommitPhaseTwo commits all staged instances of SVG to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVG *BackRepoSVGStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, svg := range backRepoSVG.Map_SVGDBID_SVGPtr {
		backRepoSVG.CommitPhaseTwoInstance(backRepo, idx, svg)
	}

	return
}

// BackRepoSVG.CommitPhaseTwoInstance commits {{structname }} of models.SVG to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVG *BackRepoSVGStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, svg *models.SVG) (Error error) {

	// fetch matching svgDB
	if svgDB, ok := backRepoSVG.Map_SVGDBID_SVGDB[idx]; ok {

		svgDB.CopyBasicFieldsFromSVG(svg)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		svgDB.SVGPointersEncoding.Layers = make([]int, 0)
		// 2. encode
		for _, layerAssocEnd := range svg.Layers {
			layerAssocEnd_DB :=
				backRepo.BackRepoLayer.GetLayerDBFromLayerPtr(layerAssocEnd)
			
			// the stage might be inconsistant, meaning that the layerAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if layerAssocEnd_DB == nil {
				continue
			}
			
			svgDB.SVGPointersEncoding.Layers =
				append(svgDB.SVGPointersEncoding.Layers, int(layerAssocEnd_DB.ID))
		}

		// commit pointer value svg.StartRect translates to updating the svg.StartRectID
		svgDB.StartRectID.Valid = true // allow for a 0 value (nil association)
		if svg.StartRect != nil {
			if StartRectId, ok := backRepo.BackRepoRect.Map_RectPtr_RectDBID[svg.StartRect]; ok {
				svgDB.StartRectID.Int64 = int64(StartRectId)
				svgDB.StartRectID.Valid = true
			}
		} else {
			svgDB.StartRectID.Int64 = 0
			svgDB.StartRectID.Valid = true
		}

		// commit pointer value svg.EndRect translates to updating the svg.EndRectID
		svgDB.EndRectID.Valid = true // allow for a 0 value (nil association)
		if svg.EndRect != nil {
			if EndRectId, ok := backRepo.BackRepoRect.Map_RectPtr_RectDBID[svg.EndRect]; ok {
				svgDB.EndRectID.Int64 = int64(EndRectId)
				svgDB.EndRectID.Valid = true
			}
		} else {
			svgDB.EndRectID.Int64 = 0
			svgDB.EndRectID.Valid = true
		}

		_, err := backRepoSVG.db.Save(svgDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SVG intance %s", svg.Name))
		return err
	}

	return
}

// BackRepoSVG.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSVG *BackRepoSVGStruct) CheckoutPhaseOne() (Error error) {

	svgDBArray := make([]SVGDB, 0)
	_, err := backRepoSVG.db.Find(&svgDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	svgInstancesToBeRemovedFromTheStage := make(map[*models.SVG]any)
	for key, value := range backRepoSVG.stage.SVGs {
		svgInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, svgDB := range svgDBArray {
		backRepoSVG.CheckoutPhaseOneInstance(&svgDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		svg, ok := backRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID]
		if ok {
			delete(svgInstancesToBeRemovedFromTheStage, svg)
		}
	}

	// remove from stage and back repo's 3 maps all svgs that are not in the checkout
	for svg := range svgInstancesToBeRemovedFromTheStage {
		svg.Unstage(backRepoSVG.GetStage())

		// remove instance from the back repo 3 maps
		svgID := backRepoSVG.Map_SVGPtr_SVGDBID[svg]
		delete(backRepoSVG.Map_SVGPtr_SVGDBID, svg)
		delete(backRepoSVG.Map_SVGDBID_SVGDB, svgID)
		delete(backRepoSVG.Map_SVGDBID_SVGPtr, svgID)
	}

	return
}

// CheckoutPhaseOneInstance takes a svgDB that has been found in the DB, updates the backRepo and stages the
// models version of the svgDB
func (backRepoSVG *BackRepoSVGStruct) CheckoutPhaseOneInstance(svgDB *SVGDB) (Error error) {

	svg, ok := backRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID]
	if !ok {
		svg = new(models.SVG)

		backRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID] = svg
		backRepoSVG.Map_SVGPtr_SVGDBID[svg] = svgDB.ID

		// append model store with the new element
		svg.Name = svgDB.Name_Data.String
		svg.Stage(backRepoSVG.GetStage())
	}
	svgDB.CopyBasicFieldsToSVG(svg)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	svg.Stage(backRepoSVG.GetStage())

	// preserve pointer to svgDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SVGDBID_SVGDB)[svgDB hold variable pointers
	svgDB_Data := *svgDB
	preservedPtrToSVG := &svgDB_Data
	backRepoSVG.Map_SVGDBID_SVGDB[svgDB.ID] = preservedPtrToSVG

	return
}

// BackRepoSVG.CheckoutPhaseTwo Checkouts all staged instances of SVG to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVG *BackRepoSVGStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, svgDB := range backRepoSVG.Map_SVGDBID_SVGDB {
		backRepoSVG.CheckoutPhaseTwoInstance(backRepo, svgDB)
	}
	return
}

// BackRepoSVG.CheckoutPhaseTwoInstance Checkouts staged instances of SVG to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVG *BackRepoSVGStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, svgDB *SVGDB) (Error error) {

	svg := backRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID]

	svgDB.DecodePointers(backRepo, svg)

	return
}

func (svgDB *SVGDB) DecodePointers(backRepo *BackRepoStruct, svg *models.SVG) {

	// insertion point for checkout of pointer encoding
	// This loop redeem svg.Layers in the stage from the encode in the back repo
	// It parses all LayerDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	svg.Layers = svg.Layers[:0]
	for _, _Layerid := range svgDB.SVGPointersEncoding.Layers {
		svg.Layers = append(svg.Layers, backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[uint(_Layerid)])
	}

	// StartRect field	
	{
		id := svgDB.StartRectID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoRect.Map_RectDBID_RectPtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: svg.StartRect, unknown pointer id", id)
				svg.StartRect = nil
			} else {
				// updates only if field has changed
				if svg.StartRect == nil || svg.StartRect != tmp {
					svg.StartRect = tmp
				}
			}
		} else {
			svg.StartRect = nil
		}
	}
	
	// EndRect field	
	{
		id := svgDB.EndRectID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoRect.Map_RectDBID_RectPtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: svg.EndRect, unknown pointer id", id)
				svg.EndRect = nil
			} else {
				// updates only if field has changed
				if svg.EndRect == nil || svg.EndRect != tmp {
					svg.EndRect = tmp
				}
			}
		} else {
			svg.EndRect = nil
		}
	}
	
	return
}

// CommitSVG allows commit of a single svg (if already staged)
func (backRepo *BackRepoStruct) CommitSVG(svg *models.SVG) {
	backRepo.BackRepoSVG.CommitPhaseOneInstance(svg)
	if id, ok := backRepo.BackRepoSVG.Map_SVGPtr_SVGDBID[svg]; ok {
		backRepo.BackRepoSVG.CommitPhaseTwoInstance(backRepo, id, svg)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSVG allows checkout of a single svg (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSVG(svg *models.SVG) {
	// check if the svg is staged
	if _, ok := backRepo.BackRepoSVG.Map_SVGPtr_SVGDBID[svg]; ok {

		if id, ok := backRepo.BackRepoSVG.Map_SVGPtr_SVGDBID[svg]; ok {
			var svgDB SVGDB
			svgDB.ID = id

			if _, err := backRepo.BackRepoSVG.db.First(&svgDB, id); err != nil {
				log.Fatalln("CheckoutSVG : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSVG.CheckoutPhaseOneInstance(&svgDB)
			backRepo.BackRepoSVG.CheckoutPhaseTwoInstance(backRepo, &svgDB)
		}
	}
}

// CopyBasicFieldsFromSVG
func (svgDB *SVGDB) CopyBasicFieldsFromSVG(svg *models.SVG) {
	// insertion point for fields commit

	svgDB.Name_Data.String = svg.Name
	svgDB.Name_Data.Valid = true

	svgDB.DrawingState_Data.String = svg.DrawingState.ToString()
	svgDB.DrawingState_Data.Valid = true

	svgDB.IsEditable_Data.Bool = svg.IsEditable
	svgDB.IsEditable_Data.Valid = true

	svgDB.IsSVGFrontEndFileGenerated_Data.Bool = svg.IsSVGFrontEndFileGenerated
	svgDB.IsSVGFrontEndFileGenerated_Data.Valid = true

	svgDB.IsSVGBackEndFileGenerated_Data.Bool = svg.IsSVGBackEndFileGenerated
	svgDB.IsSVGBackEndFileGenerated_Data.Valid = true

	svgDB.DefaultDirectoryForGeneratedImages_Data.String = svg.DefaultDirectoryForGeneratedImages
	svgDB.DefaultDirectoryForGeneratedImages_Data.Valid = true

	svgDB.IsControlBannerHidden_Data.Bool = svg.IsControlBannerHidden
	svgDB.IsControlBannerHidden_Data.Valid = true
}

// CopyBasicFieldsFromSVG_WOP
func (svgDB *SVGDB) CopyBasicFieldsFromSVG_WOP(svg *models.SVG_WOP) {
	// insertion point for fields commit

	svgDB.Name_Data.String = svg.Name
	svgDB.Name_Data.Valid = true

	svgDB.DrawingState_Data.String = svg.DrawingState.ToString()
	svgDB.DrawingState_Data.Valid = true

	svgDB.IsEditable_Data.Bool = svg.IsEditable
	svgDB.IsEditable_Data.Valid = true

	svgDB.IsSVGFrontEndFileGenerated_Data.Bool = svg.IsSVGFrontEndFileGenerated
	svgDB.IsSVGFrontEndFileGenerated_Data.Valid = true

	svgDB.IsSVGBackEndFileGenerated_Data.Bool = svg.IsSVGBackEndFileGenerated
	svgDB.IsSVGBackEndFileGenerated_Data.Valid = true

	svgDB.DefaultDirectoryForGeneratedImages_Data.String = svg.DefaultDirectoryForGeneratedImages
	svgDB.DefaultDirectoryForGeneratedImages_Data.Valid = true

	svgDB.IsControlBannerHidden_Data.Bool = svg.IsControlBannerHidden
	svgDB.IsControlBannerHidden_Data.Valid = true
}

// CopyBasicFieldsFromSVGWOP
func (svgDB *SVGDB) CopyBasicFieldsFromSVGWOP(svg *SVGWOP) {
	// insertion point for fields commit

	svgDB.Name_Data.String = svg.Name
	svgDB.Name_Data.Valid = true

	svgDB.DrawingState_Data.String = svg.DrawingState.ToString()
	svgDB.DrawingState_Data.Valid = true

	svgDB.IsEditable_Data.Bool = svg.IsEditable
	svgDB.IsEditable_Data.Valid = true

	svgDB.IsSVGFrontEndFileGenerated_Data.Bool = svg.IsSVGFrontEndFileGenerated
	svgDB.IsSVGFrontEndFileGenerated_Data.Valid = true

	svgDB.IsSVGBackEndFileGenerated_Data.Bool = svg.IsSVGBackEndFileGenerated
	svgDB.IsSVGBackEndFileGenerated_Data.Valid = true

	svgDB.DefaultDirectoryForGeneratedImages_Data.String = svg.DefaultDirectoryForGeneratedImages
	svgDB.DefaultDirectoryForGeneratedImages_Data.Valid = true

	svgDB.IsControlBannerHidden_Data.Bool = svg.IsControlBannerHidden
	svgDB.IsControlBannerHidden_Data.Valid = true
}

// CopyBasicFieldsToSVG
func (svgDB *SVGDB) CopyBasicFieldsToSVG(svg *models.SVG) {
	// insertion point for checkout of basic fields (back repo to stage)
	svg.Name = svgDB.Name_Data.String
	svg.DrawingState.FromString(svgDB.DrawingState_Data.String)
	svg.IsEditable = svgDB.IsEditable_Data.Bool
	svg.IsSVGFrontEndFileGenerated = svgDB.IsSVGFrontEndFileGenerated_Data.Bool
	svg.IsSVGBackEndFileGenerated = svgDB.IsSVGBackEndFileGenerated_Data.Bool
	svg.DefaultDirectoryForGeneratedImages = svgDB.DefaultDirectoryForGeneratedImages_Data.String
	svg.IsControlBannerHidden = svgDB.IsControlBannerHidden_Data.Bool
}

// CopyBasicFieldsToSVG_WOP
func (svgDB *SVGDB) CopyBasicFieldsToSVG_WOP(svg *models.SVG_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	svg.Name = svgDB.Name_Data.String
	svg.DrawingState.FromString(svgDB.DrawingState_Data.String)
	svg.IsEditable = svgDB.IsEditable_Data.Bool
	svg.IsSVGFrontEndFileGenerated = svgDB.IsSVGFrontEndFileGenerated_Data.Bool
	svg.IsSVGBackEndFileGenerated = svgDB.IsSVGBackEndFileGenerated_Data.Bool
	svg.DefaultDirectoryForGeneratedImages = svgDB.DefaultDirectoryForGeneratedImages_Data.String
	svg.IsControlBannerHidden = svgDB.IsControlBannerHidden_Data.Bool
}

// CopyBasicFieldsToSVGWOP
func (svgDB *SVGDB) CopyBasicFieldsToSVGWOP(svg *SVGWOP) {
	svg.ID = int(svgDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	svg.Name = svgDB.Name_Data.String
	svg.DrawingState.FromString(svgDB.DrawingState_Data.String)
	svg.IsEditable = svgDB.IsEditable_Data.Bool
	svg.IsSVGFrontEndFileGenerated = svgDB.IsSVGFrontEndFileGenerated_Data.Bool
	svg.IsSVGBackEndFileGenerated = svgDB.IsSVGBackEndFileGenerated_Data.Bool
	svg.DefaultDirectoryForGeneratedImages = svgDB.DefaultDirectoryForGeneratedImages_Data.String
	svg.IsControlBannerHidden = svgDB.IsControlBannerHidden_Data.Bool
}

// Backup generates a json file from a slice of all SVGDB instances in the backrepo
func (backRepoSVG *BackRepoSVGStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SVGDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SVGDB, 0)
	for _, svgDB := range backRepoSVG.Map_SVGDBID_SVGDB {
		forBackup = append(forBackup, svgDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SVG ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SVG file", err.Error())
	}
}

// Backup generates a json file from a slice of all SVGDB instances in the backrepo
func (backRepoSVG *BackRepoSVGStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SVGDB, 0)
	for _, svgDB := range backRepoSVG.Map_SVGDBID_SVGDB {
		forBackup = append(forBackup, svgDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SVG")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SVG_Fields, -1)
	for _, svgDB := range forBackup {

		var svgWOP SVGWOP
		svgDB.CopyBasicFieldsToSVGWOP(&svgWOP)

		row := sh.AddRow()
		row.WriteStruct(&svgWOP, -1)
	}
}

// RestoreXL from the "SVG" sheet all SVGDB instances
func (backRepoSVG *BackRepoSVGStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSVGid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SVG"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSVG.rowVisitorSVG)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSVG *BackRepoSVGStruct) rowVisitorSVG(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var svgWOP SVGWOP
		row.ReadStruct(&svgWOP)

		// add the unmarshalled struct to the stage
		svgDB := new(SVGDB)
		svgDB.CopyBasicFieldsFromSVGWOP(&svgWOP)

		svgDB_ID_atBackupTime := svgDB.ID
		svgDB.ID = 0
		_, err := backRepoSVG.db.Create(svgDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoSVG.Map_SVGDBID_SVGDB[svgDB.ID] = svgDB
		BackRepoSVGid_atBckpTime_newID[svgDB_ID_atBackupTime] = svgDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SVGDB.json" in dirPath that stores an array
// of SVGDB and stores it in the database
// the map BackRepoSVGid_atBckpTime_newID is updated accordingly
func (backRepoSVG *BackRepoSVGStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSVGid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SVGDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SVG file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SVGDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SVGDBID_SVGDB
	for _, svgDB := range forRestore {

		svgDB_ID_atBackupTime := svgDB.ID
		svgDB.ID = 0
		_, err := backRepoSVG.db.Create(svgDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoSVG.Map_SVGDBID_SVGDB[svgDB.ID] = svgDB
		BackRepoSVGid_atBckpTime_newID[svgDB_ID_atBackupTime] = svgDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SVG file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SVG>id_atBckpTime_newID
// to compute new index
func (backRepoSVG *BackRepoSVGStruct) RestorePhaseTwo() {

	for _, svgDB := range backRepoSVG.Map_SVGDBID_SVGDB {

		// next line of code is to avert unused variable compilation error
		_ = svgDB

		// insertion point for reindexing pointers encoding
		// reindexing StartRect field
		if svgDB.StartRectID.Int64 != 0 {
			svgDB.StartRectID.Int64 = int64(BackRepoRectid_atBckpTime_newID[uint(svgDB.StartRectID.Int64)])
			svgDB.StartRectID.Valid = true
		}

		// reindexing EndRect field
		if svgDB.EndRectID.Int64 != 0 {
			svgDB.EndRectID.Int64 = int64(BackRepoRectid_atBckpTime_newID[uint(svgDB.EndRectID.Int64)])
			svgDB.EndRectID.Valid = true
		}

		// update databse with new index encoding
		db, _ := backRepoSVG.db.Model(svgDB)
		_, err := db.Updates(*svgDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoSVG.ResetReversePointers commits all staged instances of SVG to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVG *BackRepoSVGStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, svg := range backRepoSVG.Map_SVGDBID_SVGPtr {
		backRepoSVG.ResetReversePointersInstance(backRepo, idx, svg)
	}

	return
}

func (backRepoSVG *BackRepoSVGStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, svg *models.SVG) (Error error) {

	// fetch matching svgDB
	if svgDB, ok := backRepoSVG.Map_SVGDBID_SVGDB[idx]; ok {
		_ = svgDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSVGid_atBckpTime_newID map[uint]uint
