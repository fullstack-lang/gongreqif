// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongreqif/go/models"
	"github.com/fullstack-lang/gongreqif/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __EmbeddedPngImage__dummysDeclaration__ models.EmbeddedPngImage
var __EmbeddedPngImage_time__dummyDeclaration time.Duration

var mutexEmbeddedPngImage sync.Mutex

// An EmbeddedPngImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmbeddedPngImage updateEmbeddedPngImage deleteEmbeddedPngImage
type EmbeddedPngImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EmbeddedPngImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmbeddedPngImage updateEmbeddedPngImage
type EmbeddedPngImageInput struct {
	// The EmbeddedPngImage to submit or modify
	// in: body
	EmbeddedPngImage *orm.EmbeddedPngImageAPI
}

// GetEmbeddedPngImages
//
// swagger:route GET /embeddedpngimages embeddedpngimages getEmbeddedPngImages
//
// # Get all embeddedpngimages
//
// Responses:
// default: genericError
//
//	200: embeddedpngimageDBResponse
func (controller *Controller) GetEmbeddedPngImages(c *gin.Context) {

	// source slice
	var embeddedpngimageDBs []orm.EmbeddedPngImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmbeddedPngImages", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoEmbeddedPngImage.GetDB()

	_, err := db.Find(&embeddedpngimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	embeddedpngimageAPIs := make([]orm.EmbeddedPngImageAPI, 0)

	// for each embeddedpngimage, update fields from the database nullable fields
	for idx := range embeddedpngimageDBs {
		embeddedpngimageDB := &embeddedpngimageDBs[idx]
		_ = embeddedpngimageDB
		var embeddedpngimageAPI orm.EmbeddedPngImageAPI

		// insertion point for updating fields
		embeddedpngimageAPI.ID = embeddedpngimageDB.ID
		embeddedpngimageDB.CopyBasicFieldsToEmbeddedPngImage_WOP(&embeddedpngimageAPI.EmbeddedPngImage_WOP)
		embeddedpngimageAPI.EmbeddedPngImagePointersEncoding = embeddedpngimageDB.EmbeddedPngImagePointersEncoding
		embeddedpngimageAPIs = append(embeddedpngimageAPIs, embeddedpngimageAPI)
	}

	c.JSON(http.StatusOK, embeddedpngimageAPIs)
}

// PostEmbeddedPngImage
//
// swagger:route POST /embeddedpngimages embeddedpngimages postEmbeddedPngImage
//
// Creates a embeddedpngimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmbeddedPngImage(c *gin.Context) {

	mutexEmbeddedPngImage.Lock()
	defer mutexEmbeddedPngImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmbeddedPngImages", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoEmbeddedPngImage.GetDB()

	// Validate input
	var input orm.EmbeddedPngImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create embeddedpngimage
	embeddedpngimageDB := orm.EmbeddedPngImageDB{}
	embeddedpngimageDB.EmbeddedPngImagePointersEncoding = input.EmbeddedPngImagePointersEncoding
	embeddedpngimageDB.CopyBasicFieldsFromEmbeddedPngImage_WOP(&input.EmbeddedPngImage_WOP)

	_, err = db.Create(&embeddedpngimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmbeddedPngImage.CheckoutPhaseOneInstance(&embeddedpngimageDB)
	embeddedpngimage := backRepo.BackRepoEmbeddedPngImage.Map_EmbeddedPngImageDBID_EmbeddedPngImagePtr[embeddedpngimageDB.ID]

	if embeddedpngimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), embeddedpngimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, embeddedpngimageDB)
}

// GetEmbeddedPngImage
//
// swagger:route GET /embeddedpngimages/{ID} embeddedpngimages getEmbeddedPngImage
//
// Gets the details for a embeddedpngimage.
//
// Responses:
// default: genericError
//
//	200: embeddedpngimageDBResponse
func (controller *Controller) GetEmbeddedPngImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmbeddedPngImage", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoEmbeddedPngImage.GetDB()

	// Get embeddedpngimageDB in DB
	var embeddedpngimageDB orm.EmbeddedPngImageDB
	if _, err := db.First(&embeddedpngimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var embeddedpngimageAPI orm.EmbeddedPngImageAPI
	embeddedpngimageAPI.ID = embeddedpngimageDB.ID
	embeddedpngimageAPI.EmbeddedPngImagePointersEncoding = embeddedpngimageDB.EmbeddedPngImagePointersEncoding
	embeddedpngimageDB.CopyBasicFieldsToEmbeddedPngImage_WOP(&embeddedpngimageAPI.EmbeddedPngImage_WOP)

	c.JSON(http.StatusOK, embeddedpngimageAPI)
}

// UpdateEmbeddedPngImage
//
// swagger:route PATCH /embeddedpngimages/{ID} embeddedpngimages updateEmbeddedPngImage
//
// # Update a embeddedpngimage
//
// Responses:
// default: genericError
//
//	200: embeddedpngimageDBResponse
func (controller *Controller) UpdateEmbeddedPngImage(c *gin.Context) {

	mutexEmbeddedPngImage.Lock()
	defer mutexEmbeddedPngImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	hasMouseEvent := false
	shiftKey := false
	_ = shiftKey
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	if len(_values) >= 2 {
		hasMouseEvent = true
		_shiftKeyValues := _values["shiftKey"]
		if len(_shiftKeyValues) == 1 {
			shiftKey = _shiftKeyValues[0] == "true"
		}
	}

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoEmbeddedPngImage.GetDB()

	// Validate input
	var input orm.EmbeddedPngImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var embeddedpngimageDB orm.EmbeddedPngImageDB

	// fetch the embeddedpngimage
	_, err := db.First(&embeddedpngimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	embeddedpngimageDB.CopyBasicFieldsFromEmbeddedPngImage_WOP(&input.EmbeddedPngImage_WOP)
	embeddedpngimageDB.EmbeddedPngImagePointersEncoding = input.EmbeddedPngImagePointersEncoding

	db, _ = db.Model(&embeddedpngimageDB)
	_, err = db.Updates(&embeddedpngimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	embeddedpngimageNew := new(models.EmbeddedPngImage)
	embeddedpngimageDB.CopyBasicFieldsToEmbeddedPngImage(embeddedpngimageNew)

	// redeem pointers
	embeddedpngimageDB.DecodePointers(backRepo, embeddedpngimageNew)

	// get stage instance from DB instance, and call callback function
	embeddedpngimageOld := backRepo.BackRepoEmbeddedPngImage.Map_EmbeddedPngImageDBID_EmbeddedPngImagePtr[embeddedpngimageDB.ID]
	if embeddedpngimageOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), embeddedpngimageOld, embeddedpngimageNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), embeddedpngimageOld, embeddedpngimageNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the embeddedpngimageDB
	c.JSON(http.StatusOK, embeddedpngimageDB)
}

// DeleteEmbeddedPngImage
//
// swagger:route DELETE /embeddedpngimages/{ID} embeddedpngimages deleteEmbeddedPngImage
//
// # Delete a embeddedpngimage
//
// default: genericError
//
//	200: embeddedpngimageDBResponse
func (controller *Controller) DeleteEmbeddedPngImage(c *gin.Context) {

	mutexEmbeddedPngImage.Lock()
	defer mutexEmbeddedPngImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmbeddedPngImage", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoEmbeddedPngImage.GetDB()

	// Get model if exist
	var embeddedpngimageDB orm.EmbeddedPngImageDB
	if _, err := db.First(&embeddedpngimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&embeddedpngimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	embeddedpngimageDeleted := new(models.EmbeddedPngImage)
	embeddedpngimageDB.CopyBasicFieldsToEmbeddedPngImage(embeddedpngimageDeleted)

	// get stage instance from DB instance, and call callback function
	embeddedpngimageStaged := backRepo.BackRepoEmbeddedPngImage.Map_EmbeddedPngImageDBID_EmbeddedPngImagePtr[embeddedpngimageDB.ID]
	if embeddedpngimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), embeddedpngimageStaged, embeddedpngimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
