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
var dummy_Image_sql sql.NullBool
var dummy_Image_time time.Duration
var dummy_Image_sort sort.Float64Slice

// ImageAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model imageAPI
type ImageAPI struct {
	gorm.Model

	models.Image_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ImagePointersEncoding ImagePointersEncoding
}

// ImagePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ImagePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// ImageDB describes a image in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model imageDB
type ImageDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field imageDB.Name
	Name_Data sql.NullString

	// Declation for basic field imageDB.SourceDirectoryPath
	SourceDirectoryPath_Data sql.NullString

	// Declation for basic field imageDB.Height
	Height_Data sql.NullInt64

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ImagePointersEncoding
}

// ImageDBs arrays imageDBs
// swagger:response imageDBsResponse
type ImageDBs []ImageDB

// ImageDBResponse provides response
// swagger:response imageDBResponse
type ImageDBResponse struct {
	ImageDB
}

// ImageWOP is a Image without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ImageWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	SourceDirectoryPath string `xlsx:"2"`

	Height int `xlsx:"3"`
	// insertion for WOP pointer fields
}

var Image_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"SourceDirectoryPath",
	"Height",
}

type BackRepoImageStruct struct {
	// stores ImageDB according to their gorm ID
	Map_ImageDBID_ImageDB map[uint]*ImageDB

	// stores ImageDB ID according to Image address
	Map_ImagePtr_ImageDBID map[*models.Image]uint

	// stores Image according to their gorm ID
	Map_ImageDBID_ImagePtr map[uint]*models.Image

	db db.DBInterface

	stage *models.Stage
}

func (backRepoImage *BackRepoImageStruct) GetStage() (stage *models.Stage) {
	stage = backRepoImage.stage
	return
}

func (backRepoImage *BackRepoImageStruct) GetDB() db.DBInterface {
	return backRepoImage.db
}

// GetImageDBFromImagePtr is a handy function to access the back repo instance from the stage instance
func (backRepoImage *BackRepoImageStruct) GetImageDBFromImagePtr(image *models.Image) (imageDB *ImageDB) {
	id := backRepoImage.Map_ImagePtr_ImageDBID[image]
	imageDB = backRepoImage.Map_ImageDBID_ImageDB[id]
	return
}

// BackRepoImage.CommitPhaseOne commits all staged instances of Image to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoImage *BackRepoImageStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var images []*models.Image
	for image := range stage.Images {
		images = append(images, image)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(images, func(i, j int) bool {
		return stage.ImageMap_Staged_Order[images[i]] < stage.ImageMap_Staged_Order[images[j]]
	})

	for _, image := range images {
		backRepoImage.CommitPhaseOneInstance(image)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, image := range backRepoImage.Map_ImageDBID_ImagePtr {
		if _, ok := stage.Images[image]; !ok {
			backRepoImage.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoImage.CommitDeleteInstance commits deletion of Image to the BackRepo
func (backRepoImage *BackRepoImageStruct) CommitDeleteInstance(id uint) (Error error) {

	image := backRepoImage.Map_ImageDBID_ImagePtr[id]

	// image is not staged anymore, remove imageDB
	imageDB := backRepoImage.Map_ImageDBID_ImageDB[id]
	db, _ := backRepoImage.db.Unscoped()
	_, err := db.Delete(imageDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoImage.Map_ImagePtr_ImageDBID, image)
	delete(backRepoImage.Map_ImageDBID_ImagePtr, id)
	delete(backRepoImage.Map_ImageDBID_ImageDB, id)

	return
}

// BackRepoImage.CommitPhaseOneInstance commits image staged instances of Image to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoImage *BackRepoImageStruct) CommitPhaseOneInstance(image *models.Image) (Error error) {

	// check if the image is not commited yet
	if _, ok := backRepoImage.Map_ImagePtr_ImageDBID[image]; ok {
		return
	}

	// initiate image
	var imageDB ImageDB
	imageDB.CopyBasicFieldsFromImage(image)

	_, err := backRepoImage.db.Create(&imageDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoImage.Map_ImagePtr_ImageDBID[image] = imageDB.ID
	backRepoImage.Map_ImageDBID_ImagePtr[imageDB.ID] = image
	backRepoImage.Map_ImageDBID_ImageDB[imageDB.ID] = &imageDB

	return
}

// BackRepoImage.CommitPhaseTwo commits all staged instances of Image to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoImage *BackRepoImageStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, image := range backRepoImage.Map_ImageDBID_ImagePtr {
		backRepoImage.CommitPhaseTwoInstance(backRepo, idx, image)
	}

	return
}

// BackRepoImage.CommitPhaseTwoInstance commits {{structname }} of models.Image to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoImage *BackRepoImageStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, image *models.Image) (Error error) {

	// fetch matching imageDB
	if imageDB, ok := backRepoImage.Map_ImageDBID_ImageDB[idx]; ok {

		imageDB.CopyBasicFieldsFromImage(image)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoImage.db.Save(imageDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Image intance %s", image.Name))
		return err
	}

	return
}

// BackRepoImage.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoImage *BackRepoImageStruct) CheckoutPhaseOne() (Error error) {

	imageDBArray := make([]ImageDB, 0)
	_, err := backRepoImage.db.Find(&imageDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	imageInstancesToBeRemovedFromTheStage := make(map[*models.Image]any)
	for key, value := range backRepoImage.stage.Images {
		imageInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, imageDB := range imageDBArray {
		backRepoImage.CheckoutPhaseOneInstance(&imageDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		image, ok := backRepoImage.Map_ImageDBID_ImagePtr[imageDB.ID]
		if ok {
			delete(imageInstancesToBeRemovedFromTheStage, image)
		}
	}

	// remove from stage and back repo's 3 maps all images that are not in the checkout
	for image := range imageInstancesToBeRemovedFromTheStage {
		image.Unstage(backRepoImage.GetStage())

		// remove instance from the back repo 3 maps
		imageID := backRepoImage.Map_ImagePtr_ImageDBID[image]
		delete(backRepoImage.Map_ImagePtr_ImageDBID, image)
		delete(backRepoImage.Map_ImageDBID_ImageDB, imageID)
		delete(backRepoImage.Map_ImageDBID_ImagePtr, imageID)
	}

	return
}

// CheckoutPhaseOneInstance takes a imageDB that has been found in the DB, updates the backRepo and stages the
// models version of the imageDB
func (backRepoImage *BackRepoImageStruct) CheckoutPhaseOneInstance(imageDB *ImageDB) (Error error) {

	image, ok := backRepoImage.Map_ImageDBID_ImagePtr[imageDB.ID]
	if !ok {
		image = new(models.Image)

		backRepoImage.Map_ImageDBID_ImagePtr[imageDB.ID] = image
		backRepoImage.Map_ImagePtr_ImageDBID[image] = imageDB.ID

		// append model store with the new element
		image.Name = imageDB.Name_Data.String
		image.Stage(backRepoImage.GetStage())
	}
	imageDB.CopyBasicFieldsToImage(image)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	image.Stage(backRepoImage.GetStage())

	// preserve pointer to imageDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ImageDBID_ImageDB)[imageDB hold variable pointers
	imageDB_Data := *imageDB
	preservedPtrToImage := &imageDB_Data
	backRepoImage.Map_ImageDBID_ImageDB[imageDB.ID] = preservedPtrToImage

	return
}

// BackRepoImage.CheckoutPhaseTwo Checkouts all staged instances of Image to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoImage *BackRepoImageStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, imageDB := range backRepoImage.Map_ImageDBID_ImageDB {
		backRepoImage.CheckoutPhaseTwoInstance(backRepo, imageDB)
	}
	return
}

// BackRepoImage.CheckoutPhaseTwoInstance Checkouts staged instances of Image to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoImage *BackRepoImageStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, imageDB *ImageDB) (Error error) {

	image := backRepoImage.Map_ImageDBID_ImagePtr[imageDB.ID]

	imageDB.DecodePointers(backRepo, image)

	return
}

func (imageDB *ImageDB) DecodePointers(backRepo *BackRepoStruct, image *models.Image) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitImage allows commit of a single image (if already staged)
func (backRepo *BackRepoStruct) CommitImage(image *models.Image) {
	backRepo.BackRepoImage.CommitPhaseOneInstance(image)
	if id, ok := backRepo.BackRepoImage.Map_ImagePtr_ImageDBID[image]; ok {
		backRepo.BackRepoImage.CommitPhaseTwoInstance(backRepo, id, image)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitImage allows checkout of a single image (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutImage(image *models.Image) {
	// check if the image is staged
	if _, ok := backRepo.BackRepoImage.Map_ImagePtr_ImageDBID[image]; ok {

		if id, ok := backRepo.BackRepoImage.Map_ImagePtr_ImageDBID[image]; ok {
			var imageDB ImageDB
			imageDB.ID = id

			if _, err := backRepo.BackRepoImage.db.First(&imageDB, id); err != nil {
				log.Fatalln("CheckoutImage : Problem with getting object with id:", id)
			}
			backRepo.BackRepoImage.CheckoutPhaseOneInstance(&imageDB)
			backRepo.BackRepoImage.CheckoutPhaseTwoInstance(backRepo, &imageDB)
		}
	}
}

// CopyBasicFieldsFromImage
func (imageDB *ImageDB) CopyBasicFieldsFromImage(image *models.Image) {
	// insertion point for fields commit

	imageDB.Name_Data.String = image.Name
	imageDB.Name_Data.Valid = true

	imageDB.SourceDirectoryPath_Data.String = image.SourceDirectoryPath
	imageDB.SourceDirectoryPath_Data.Valid = true

	imageDB.Height_Data.Int64 = int64(image.Height)
	imageDB.Height_Data.Valid = true
}

// CopyBasicFieldsFromImage_WOP
func (imageDB *ImageDB) CopyBasicFieldsFromImage_WOP(image *models.Image_WOP) {
	// insertion point for fields commit

	imageDB.Name_Data.String = image.Name
	imageDB.Name_Data.Valid = true

	imageDB.SourceDirectoryPath_Data.String = image.SourceDirectoryPath
	imageDB.SourceDirectoryPath_Data.Valid = true

	imageDB.Height_Data.Int64 = int64(image.Height)
	imageDB.Height_Data.Valid = true
}

// CopyBasicFieldsFromImageWOP
func (imageDB *ImageDB) CopyBasicFieldsFromImageWOP(image *ImageWOP) {
	// insertion point for fields commit

	imageDB.Name_Data.String = image.Name
	imageDB.Name_Data.Valid = true

	imageDB.SourceDirectoryPath_Data.String = image.SourceDirectoryPath
	imageDB.SourceDirectoryPath_Data.Valid = true

	imageDB.Height_Data.Int64 = int64(image.Height)
	imageDB.Height_Data.Valid = true
}

// CopyBasicFieldsToImage
func (imageDB *ImageDB) CopyBasicFieldsToImage(image *models.Image) {
	// insertion point for checkout of basic fields (back repo to stage)
	image.Name = imageDB.Name_Data.String
	image.SourceDirectoryPath = imageDB.SourceDirectoryPath_Data.String
	image.Height = int(imageDB.Height_Data.Int64)
}

// CopyBasicFieldsToImage_WOP
func (imageDB *ImageDB) CopyBasicFieldsToImage_WOP(image *models.Image_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	image.Name = imageDB.Name_Data.String
	image.SourceDirectoryPath = imageDB.SourceDirectoryPath_Data.String
	image.Height = int(imageDB.Height_Data.Int64)
}

// CopyBasicFieldsToImageWOP
func (imageDB *ImageDB) CopyBasicFieldsToImageWOP(image *ImageWOP) {
	image.ID = int(imageDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	image.Name = imageDB.Name_Data.String
	image.SourceDirectoryPath = imageDB.SourceDirectoryPath_Data.String
	image.Height = int(imageDB.Height_Data.Int64)
}

// Backup generates a json file from a slice of all ImageDB instances in the backrepo
func (backRepoImage *BackRepoImageStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ImageDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ImageDB, 0)
	for _, imageDB := range backRepoImage.Map_ImageDBID_ImageDB {
		forBackup = append(forBackup, imageDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Image ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Image file", err.Error())
	}
}

// Backup generates a json file from a slice of all ImageDB instances in the backrepo
func (backRepoImage *BackRepoImageStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ImageDB, 0)
	for _, imageDB := range backRepoImage.Map_ImageDBID_ImageDB {
		forBackup = append(forBackup, imageDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Image")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Image_Fields, -1)
	for _, imageDB := range forBackup {

		var imageWOP ImageWOP
		imageDB.CopyBasicFieldsToImageWOP(&imageWOP)

		row := sh.AddRow()
		row.WriteStruct(&imageWOP, -1)
	}
}

// RestoreXL from the "Image" sheet all ImageDB instances
func (backRepoImage *BackRepoImageStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoImageid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Image"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoImage.rowVisitorImage)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoImage *BackRepoImageStruct) rowVisitorImage(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var imageWOP ImageWOP
		row.ReadStruct(&imageWOP)

		// add the unmarshalled struct to the stage
		imageDB := new(ImageDB)
		imageDB.CopyBasicFieldsFromImageWOP(&imageWOP)

		imageDB_ID_atBackupTime := imageDB.ID
		imageDB.ID = 0
		_, err := backRepoImage.db.Create(imageDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoImage.Map_ImageDBID_ImageDB[imageDB.ID] = imageDB
		BackRepoImageid_atBckpTime_newID[imageDB_ID_atBackupTime] = imageDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ImageDB.json" in dirPath that stores an array
// of ImageDB and stores it in the database
// the map BackRepoImageid_atBckpTime_newID is updated accordingly
func (backRepoImage *BackRepoImageStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoImageid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ImageDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Image file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ImageDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ImageDBID_ImageDB
	for _, imageDB := range forRestore {

		imageDB_ID_atBackupTime := imageDB.ID
		imageDB.ID = 0
		_, err := backRepoImage.db.Create(imageDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoImage.Map_ImageDBID_ImageDB[imageDB.ID] = imageDB
		BackRepoImageid_atBckpTime_newID[imageDB_ID_atBackupTime] = imageDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Image file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Image>id_atBckpTime_newID
// to compute new index
func (backRepoImage *BackRepoImageStruct) RestorePhaseTwo() {

	for _, imageDB := range backRepoImage.Map_ImageDBID_ImageDB {

		// next line of code is to avert unused variable compilation error
		_ = imageDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoImage.db.Model(imageDB)
		_, err := db.Updates(*imageDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoImage.ResetReversePointers commits all staged instances of Image to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoImage *BackRepoImageStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, image := range backRepoImage.Map_ImageDBID_ImagePtr {
		backRepoImage.ResetReversePointersInstance(backRepo, idx, image)
	}

	return
}

func (backRepoImage *BackRepoImageStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, image *models.Image) (Error error) {

	// fetch matching imageDB
	if imageDB, ok := backRepoImage.Map_ImageDBID_ImageDB[idx]; ok {
		_ = imageDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoImageid_atBckpTime_newID map[uint]uint
