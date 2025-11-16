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
var __EmbeddedJpgImage__dummysDeclaration__ models.EmbeddedJpgImage
var __EmbeddedJpgImage_time__dummyDeclaration time.Duration

var mutexEmbeddedJpgImage sync.Mutex

// An EmbeddedJpgImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmbeddedJpgImage updateEmbeddedJpgImage deleteEmbeddedJpgImage
type EmbeddedJpgImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EmbeddedJpgImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmbeddedJpgImage updateEmbeddedJpgImage
type EmbeddedJpgImageInput struct {
	// The EmbeddedJpgImage to submit or modify
	// in: body
	EmbeddedJpgImage *orm.EmbeddedJpgImageAPI
}

// GetEmbeddedJpgImages
//
// swagger:route GET /embeddedjpgimages embeddedjpgimages getEmbeddedJpgImages
//
// # Get all embeddedjpgimages
//
// Responses:
// default: genericError
//
//	200: embeddedjpgimageDBResponse
func (controller *Controller) GetEmbeddedJpgImages(c *gin.Context) {

	// source slice
	var embeddedjpgimageDBs []orm.EmbeddedJpgImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmbeddedJpgImages", "Name", stackPath)
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
	db := backRepo.BackRepoEmbeddedJpgImage.GetDB()

	_, err := db.Find(&embeddedjpgimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	embeddedjpgimageAPIs := make([]orm.EmbeddedJpgImageAPI, 0)

	// for each embeddedjpgimage, update fields from the database nullable fields
	for idx := range embeddedjpgimageDBs {
		embeddedjpgimageDB := &embeddedjpgimageDBs[idx]
		_ = embeddedjpgimageDB
		var embeddedjpgimageAPI orm.EmbeddedJpgImageAPI

		// insertion point for updating fields
		embeddedjpgimageAPI.ID = embeddedjpgimageDB.ID
		embeddedjpgimageDB.CopyBasicFieldsToEmbeddedJpgImage_WOP(&embeddedjpgimageAPI.EmbeddedJpgImage_WOP)
		embeddedjpgimageAPI.EmbeddedJpgImagePointersEncoding = embeddedjpgimageDB.EmbeddedJpgImagePointersEncoding
		embeddedjpgimageAPIs = append(embeddedjpgimageAPIs, embeddedjpgimageAPI)
	}

	c.JSON(http.StatusOK, embeddedjpgimageAPIs)
}

// PostEmbeddedJpgImage
//
// swagger:route POST /embeddedjpgimages embeddedjpgimages postEmbeddedJpgImage
//
// Creates a embeddedjpgimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmbeddedJpgImage(c *gin.Context) {

	mutexEmbeddedJpgImage.Lock()
	defer mutexEmbeddedJpgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmbeddedJpgImages", "Name", stackPath)
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
	db := backRepo.BackRepoEmbeddedJpgImage.GetDB()

	// Validate input
	var input orm.EmbeddedJpgImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create embeddedjpgimage
	embeddedjpgimageDB := orm.EmbeddedJpgImageDB{}
	embeddedjpgimageDB.EmbeddedJpgImagePointersEncoding = input.EmbeddedJpgImagePointersEncoding
	embeddedjpgimageDB.CopyBasicFieldsFromEmbeddedJpgImage_WOP(&input.EmbeddedJpgImage_WOP)

	_, err = db.Create(&embeddedjpgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmbeddedJpgImage.CheckoutPhaseOneInstance(&embeddedjpgimageDB)
	embeddedjpgimage := backRepo.BackRepoEmbeddedJpgImage.Map_EmbeddedJpgImageDBID_EmbeddedJpgImagePtr[embeddedjpgimageDB.ID]

	if embeddedjpgimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), embeddedjpgimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, embeddedjpgimageDB)
}

// GetEmbeddedJpgImage
//
// swagger:route GET /embeddedjpgimages/{ID} embeddedjpgimages getEmbeddedJpgImage
//
// Gets the details for a embeddedjpgimage.
//
// Responses:
// default: genericError
//
//	200: embeddedjpgimageDBResponse
func (controller *Controller) GetEmbeddedJpgImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmbeddedJpgImage", "Name", stackPath)
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
	db := backRepo.BackRepoEmbeddedJpgImage.GetDB()

	// Get embeddedjpgimageDB in DB
	var embeddedjpgimageDB orm.EmbeddedJpgImageDB
	if _, err := db.First(&embeddedjpgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var embeddedjpgimageAPI orm.EmbeddedJpgImageAPI
	embeddedjpgimageAPI.ID = embeddedjpgimageDB.ID
	embeddedjpgimageAPI.EmbeddedJpgImagePointersEncoding = embeddedjpgimageDB.EmbeddedJpgImagePointersEncoding
	embeddedjpgimageDB.CopyBasicFieldsToEmbeddedJpgImage_WOP(&embeddedjpgimageAPI.EmbeddedJpgImage_WOP)

	c.JSON(http.StatusOK, embeddedjpgimageAPI)
}

// UpdateEmbeddedJpgImage
//
// swagger:route PATCH /embeddedjpgimages/{ID} embeddedjpgimages updateEmbeddedJpgImage
//
// # Update a embeddedjpgimage
//
// Responses:
// default: genericError
//
//	200: embeddedjpgimageDBResponse
func (controller *Controller) UpdateEmbeddedJpgImage(c *gin.Context) {

	mutexEmbeddedJpgImage.Lock()
	defer mutexEmbeddedJpgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
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
	db := backRepo.BackRepoEmbeddedJpgImage.GetDB()

	// Validate input
	var input orm.EmbeddedJpgImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var embeddedjpgimageDB orm.EmbeddedJpgImageDB

	// fetch the embeddedjpgimage
	_, err := db.First(&embeddedjpgimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	embeddedjpgimageDB.CopyBasicFieldsFromEmbeddedJpgImage_WOP(&input.EmbeddedJpgImage_WOP)
	embeddedjpgimageDB.EmbeddedJpgImagePointersEncoding = input.EmbeddedJpgImagePointersEncoding

	db, _ = db.Model(&embeddedjpgimageDB)
	_, err = db.Updates(&embeddedjpgimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	embeddedjpgimageNew := new(models.EmbeddedJpgImage)
	embeddedjpgimageDB.CopyBasicFieldsToEmbeddedJpgImage(embeddedjpgimageNew)

	// redeem pointers
	embeddedjpgimageDB.DecodePointers(backRepo, embeddedjpgimageNew)

	// get stage instance from DB instance, and call callback function
	embeddedjpgimageOld := backRepo.BackRepoEmbeddedJpgImage.Map_EmbeddedJpgImageDBID_EmbeddedJpgImagePtr[embeddedjpgimageDB.ID]
	if embeddedjpgimageOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), embeddedjpgimageOld, embeddedjpgimageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the embeddedjpgimageDB
	c.JSON(http.StatusOK, embeddedjpgimageDB)
}

// DeleteEmbeddedJpgImage
//
// swagger:route DELETE /embeddedjpgimages/{ID} embeddedjpgimages deleteEmbeddedJpgImage
//
// # Delete a embeddedjpgimage
//
// default: genericError
//
//	200: embeddedjpgimageDBResponse
func (controller *Controller) DeleteEmbeddedJpgImage(c *gin.Context) {

	mutexEmbeddedJpgImage.Lock()
	defer mutexEmbeddedJpgImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmbeddedJpgImage", "Name", stackPath)
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
	db := backRepo.BackRepoEmbeddedJpgImage.GetDB()

	// Get model if exist
	var embeddedjpgimageDB orm.EmbeddedJpgImageDB
	if _, err := db.First(&embeddedjpgimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&embeddedjpgimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	embeddedjpgimageDeleted := new(models.EmbeddedJpgImage)
	embeddedjpgimageDB.CopyBasicFieldsToEmbeddedJpgImage(embeddedjpgimageDeleted)

	// get stage instance from DB instance, and call callback function
	embeddedjpgimageStaged := backRepo.BackRepoEmbeddedJpgImage.Map_EmbeddedJpgImageDBID_EmbeddedJpgImagePtr[embeddedjpgimageDB.ID]
	if embeddedjpgimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), embeddedjpgimageStaged, embeddedjpgimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
