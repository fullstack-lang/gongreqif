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
var __EmbeddedSvgImage__dummysDeclaration__ models.EmbeddedSvgImage
var __EmbeddedSvgImage_time__dummyDeclaration time.Duration

var mutexEmbeddedSvgImage sync.Mutex

// An EmbeddedSvgImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmbeddedSvgImage updateEmbeddedSvgImage deleteEmbeddedSvgImage
type EmbeddedSvgImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EmbeddedSvgImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmbeddedSvgImage updateEmbeddedSvgImage
type EmbeddedSvgImageInput struct {
	// The EmbeddedSvgImage to submit or modify
	// in: body
	EmbeddedSvgImage *orm.EmbeddedSvgImageAPI
}

// GetEmbeddedSvgImages
//
// swagger:route GET /embeddedsvgimages embeddedsvgimages getEmbeddedSvgImages
//
// # Get all embeddedsvgimages
//
// Responses:
// default: genericError
//
//	200: embeddedsvgimageDBResponse
func (controller *Controller) GetEmbeddedSvgImages(c *gin.Context) {

	// source slice
	var embeddedsvgimageDBs []orm.EmbeddedSvgImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmbeddedSvgImages", "Name", stackPath)
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
	db := backRepo.BackRepoEmbeddedSvgImage.GetDB()

	_, err := db.Find(&embeddedsvgimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	embeddedsvgimageAPIs := make([]orm.EmbeddedSvgImageAPI, 0)

	// for each embeddedsvgimage, update fields from the database nullable fields
	for idx := range embeddedsvgimageDBs {
		embeddedsvgimageDB := &embeddedsvgimageDBs[idx]
		_ = embeddedsvgimageDB
		var embeddedsvgimageAPI orm.EmbeddedSvgImageAPI

		// insertion point for updating fields
		embeddedsvgimageAPI.ID = embeddedsvgimageDB.ID
		embeddedsvgimageDB.CopyBasicFieldsToEmbeddedSvgImage_WOP(&embeddedsvgimageAPI.EmbeddedSvgImage_WOP)
		embeddedsvgimageAPI.EmbeddedSvgImagePointersEncoding = embeddedsvgimageDB.EmbeddedSvgImagePointersEncoding
		embeddedsvgimageAPIs = append(embeddedsvgimageAPIs, embeddedsvgimageAPI)
	}

	c.JSON(http.StatusOK, embeddedsvgimageAPIs)
}

// PostEmbeddedSvgImage
//
// swagger:route POST /embeddedsvgimages embeddedsvgimages postEmbeddedSvgImage
//
// Creates a embeddedsvgimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmbeddedSvgImage(c *gin.Context) {

	mutexEmbeddedSvgImage.Lock()
	defer mutexEmbeddedSvgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmbeddedSvgImages", "Name", stackPath)
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
	db := backRepo.BackRepoEmbeddedSvgImage.GetDB()

	// Validate input
	var input orm.EmbeddedSvgImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create embeddedsvgimage
	embeddedsvgimageDB := orm.EmbeddedSvgImageDB{}
	embeddedsvgimageDB.EmbeddedSvgImagePointersEncoding = input.EmbeddedSvgImagePointersEncoding
	embeddedsvgimageDB.CopyBasicFieldsFromEmbeddedSvgImage_WOP(&input.EmbeddedSvgImage_WOP)

	_, err = db.Create(&embeddedsvgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmbeddedSvgImage.CheckoutPhaseOneInstance(&embeddedsvgimageDB)
	embeddedsvgimage := backRepo.BackRepoEmbeddedSvgImage.Map_EmbeddedSvgImageDBID_EmbeddedSvgImagePtr[embeddedsvgimageDB.ID]

	if embeddedsvgimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), embeddedsvgimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, embeddedsvgimageDB)
}

// GetEmbeddedSvgImage
//
// swagger:route GET /embeddedsvgimages/{ID} embeddedsvgimages getEmbeddedSvgImage
//
// Gets the details for a embeddedsvgimage.
//
// Responses:
// default: genericError
//
//	200: embeddedsvgimageDBResponse
func (controller *Controller) GetEmbeddedSvgImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmbeddedSvgImage", "Name", stackPath)
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
	db := backRepo.BackRepoEmbeddedSvgImage.GetDB()

	// Get embeddedsvgimageDB in DB
	var embeddedsvgimageDB orm.EmbeddedSvgImageDB
	if _, err := db.First(&embeddedsvgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var embeddedsvgimageAPI orm.EmbeddedSvgImageAPI
	embeddedsvgimageAPI.ID = embeddedsvgimageDB.ID
	embeddedsvgimageAPI.EmbeddedSvgImagePointersEncoding = embeddedsvgimageDB.EmbeddedSvgImagePointersEncoding
	embeddedsvgimageDB.CopyBasicFieldsToEmbeddedSvgImage_WOP(&embeddedsvgimageAPI.EmbeddedSvgImage_WOP)

	c.JSON(http.StatusOK, embeddedsvgimageAPI)
}

// UpdateEmbeddedSvgImage
//
// swagger:route PATCH /embeddedsvgimages/{ID} embeddedsvgimages updateEmbeddedSvgImage
//
// # Update a embeddedsvgimage
//
// Responses:
// default: genericError
//
//	200: embeddedsvgimageDBResponse
func (controller *Controller) UpdateEmbeddedSvgImage(c *gin.Context) {

	mutexEmbeddedSvgImage.Lock()
	defer mutexEmbeddedSvgImage.Unlock()

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
	db := backRepo.BackRepoEmbeddedSvgImage.GetDB()

	// Validate input
	var input orm.EmbeddedSvgImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var embeddedsvgimageDB orm.EmbeddedSvgImageDB

	// fetch the embeddedsvgimage
	_, err := db.First(&embeddedsvgimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	embeddedsvgimageDB.CopyBasicFieldsFromEmbeddedSvgImage_WOP(&input.EmbeddedSvgImage_WOP)
	embeddedsvgimageDB.EmbeddedSvgImagePointersEncoding = input.EmbeddedSvgImagePointersEncoding

	db, _ = db.Model(&embeddedsvgimageDB)
	_, err = db.Updates(&embeddedsvgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	embeddedsvgimageNew := new(models.EmbeddedSvgImage)
	embeddedsvgimageDB.CopyBasicFieldsToEmbeddedSvgImage(embeddedsvgimageNew)

	// redeem pointers
	embeddedsvgimageDB.DecodePointers(backRepo, embeddedsvgimageNew)

	// get stage instance from DB instance, and call callback function
	embeddedsvgimageOld := backRepo.BackRepoEmbeddedSvgImage.Map_EmbeddedSvgImageDBID_EmbeddedSvgImagePtr[embeddedsvgimageDB.ID]
	if embeddedsvgimageOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), embeddedsvgimageOld, embeddedsvgimageNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), embeddedsvgimageOld, embeddedsvgimageNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the embeddedsvgimageDB
	c.JSON(http.StatusOK, embeddedsvgimageDB)
}

// DeleteEmbeddedSvgImage
//
// swagger:route DELETE /embeddedsvgimages/{ID} embeddedsvgimages deleteEmbeddedSvgImage
//
// # Delete a embeddedsvgimage
//
// default: genericError
//
//	200: embeddedsvgimageDBResponse
func (controller *Controller) DeleteEmbeddedSvgImage(c *gin.Context) {

	mutexEmbeddedSvgImage.Lock()
	defer mutexEmbeddedSvgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmbeddedSvgImage", "Name", stackPath)
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
	db := backRepo.BackRepoEmbeddedSvgImage.GetDB()

	// Get model if exist
	var embeddedsvgimageDB orm.EmbeddedSvgImageDB
	if _, err := db.First(&embeddedsvgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&embeddedsvgimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	embeddedsvgimageDeleted := new(models.EmbeddedSvgImage)
	embeddedsvgimageDB.CopyBasicFieldsToEmbeddedSvgImage(embeddedsvgimageDeleted)

	// get stage instance from DB instance, and call callback function
	embeddedsvgimageStaged := backRepo.BackRepoEmbeddedSvgImage.Map_EmbeddedSvgImageDBID_EmbeddedSvgImagePtr[embeddedsvgimageDB.ID]
	if embeddedsvgimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), embeddedsvgimageStaged, embeddedsvgimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
