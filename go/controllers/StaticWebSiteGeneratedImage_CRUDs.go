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
var __StaticWebSiteGeneratedImage__dummysDeclaration__ models.StaticWebSiteGeneratedImage
var __StaticWebSiteGeneratedImage_time__dummyDeclaration time.Duration

var mutexStaticWebSiteGeneratedImage sync.Mutex

// An StaticWebSiteGeneratedImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStaticWebSiteGeneratedImage updateStaticWebSiteGeneratedImage deleteStaticWebSiteGeneratedImage
type StaticWebSiteGeneratedImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StaticWebSiteGeneratedImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStaticWebSiteGeneratedImage updateStaticWebSiteGeneratedImage
type StaticWebSiteGeneratedImageInput struct {
	// The StaticWebSiteGeneratedImage to submit or modify
	// in: body
	StaticWebSiteGeneratedImage *orm.StaticWebSiteGeneratedImageAPI
}

// GetStaticWebSiteGeneratedImages
//
// swagger:route GET /staticwebsitegeneratedimages staticwebsitegeneratedimages getStaticWebSiteGeneratedImages
//
// # Get all staticwebsitegeneratedimages
//
// Responses:
// default: genericError
//
//	200: staticwebsitegeneratedimageDBResponse
func (controller *Controller) GetStaticWebSiteGeneratedImages(c *gin.Context) {

	// source slice
	var staticwebsitegeneratedimageDBs []orm.StaticWebSiteGeneratedImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaticWebSiteGeneratedImages", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteGeneratedImage.GetDB()

	_, err := db.Find(&staticwebsitegeneratedimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	staticwebsitegeneratedimageAPIs := make([]orm.StaticWebSiteGeneratedImageAPI, 0)

	// for each staticwebsitegeneratedimage, update fields from the database nullable fields
	for idx := range staticwebsitegeneratedimageDBs {
		staticwebsitegeneratedimageDB := &staticwebsitegeneratedimageDBs[idx]
		_ = staticwebsitegeneratedimageDB
		var staticwebsitegeneratedimageAPI orm.StaticWebSiteGeneratedImageAPI

		// insertion point for updating fields
		staticwebsitegeneratedimageAPI.ID = staticwebsitegeneratedimageDB.ID
		staticwebsitegeneratedimageDB.CopyBasicFieldsToStaticWebSiteGeneratedImage_WOP(&staticwebsitegeneratedimageAPI.StaticWebSiteGeneratedImage_WOP)
		staticwebsitegeneratedimageAPI.StaticWebSiteGeneratedImagePointersEncoding = staticwebsitegeneratedimageDB.StaticWebSiteGeneratedImagePointersEncoding
		staticwebsitegeneratedimageAPIs = append(staticwebsitegeneratedimageAPIs, staticwebsitegeneratedimageAPI)
	}

	c.JSON(http.StatusOK, staticwebsitegeneratedimageAPIs)
}

// PostStaticWebSiteGeneratedImage
//
// swagger:route POST /staticwebsitegeneratedimages staticwebsitegeneratedimages postStaticWebSiteGeneratedImage
//
// Creates a staticwebsitegeneratedimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStaticWebSiteGeneratedImage(c *gin.Context) {

	mutexStaticWebSiteGeneratedImage.Lock()
	defer mutexStaticWebSiteGeneratedImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStaticWebSiteGeneratedImages", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteGeneratedImage.GetDB()

	// Validate input
	var input orm.StaticWebSiteGeneratedImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create staticwebsitegeneratedimage
	staticwebsitegeneratedimageDB := orm.StaticWebSiteGeneratedImageDB{}
	staticwebsitegeneratedimageDB.StaticWebSiteGeneratedImagePointersEncoding = input.StaticWebSiteGeneratedImagePointersEncoding
	staticwebsitegeneratedimageDB.CopyBasicFieldsFromStaticWebSiteGeneratedImage_WOP(&input.StaticWebSiteGeneratedImage_WOP)

	_, err = db.Create(&staticwebsitegeneratedimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStaticWebSiteGeneratedImage.CheckoutPhaseOneInstance(&staticwebsitegeneratedimageDB)
	staticwebsitegeneratedimage := backRepo.BackRepoStaticWebSiteGeneratedImage.Map_StaticWebSiteGeneratedImageDBID_StaticWebSiteGeneratedImagePtr[staticwebsitegeneratedimageDB.ID]

	if staticwebsitegeneratedimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), staticwebsitegeneratedimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, staticwebsitegeneratedimageDB)
}

// GetStaticWebSiteGeneratedImage
//
// swagger:route GET /staticwebsitegeneratedimages/{ID} staticwebsitegeneratedimages getStaticWebSiteGeneratedImage
//
// Gets the details for a staticwebsitegeneratedimage.
//
// Responses:
// default: genericError
//
//	200: staticwebsitegeneratedimageDBResponse
func (controller *Controller) GetStaticWebSiteGeneratedImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaticWebSiteGeneratedImage", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteGeneratedImage.GetDB()

	// Get staticwebsitegeneratedimageDB in DB
	var staticwebsitegeneratedimageDB orm.StaticWebSiteGeneratedImageDB
	if _, err := db.First(&staticwebsitegeneratedimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var staticwebsitegeneratedimageAPI orm.StaticWebSiteGeneratedImageAPI
	staticwebsitegeneratedimageAPI.ID = staticwebsitegeneratedimageDB.ID
	staticwebsitegeneratedimageAPI.StaticWebSiteGeneratedImagePointersEncoding = staticwebsitegeneratedimageDB.StaticWebSiteGeneratedImagePointersEncoding
	staticwebsitegeneratedimageDB.CopyBasicFieldsToStaticWebSiteGeneratedImage_WOP(&staticwebsitegeneratedimageAPI.StaticWebSiteGeneratedImage_WOP)

	c.JSON(http.StatusOK, staticwebsitegeneratedimageAPI)
}

// UpdateStaticWebSiteGeneratedImage
//
// swagger:route PATCH /staticwebsitegeneratedimages/{ID} staticwebsitegeneratedimages updateStaticWebSiteGeneratedImage
//
// # Update a staticwebsitegeneratedimage
//
// Responses:
// default: genericError
//
//	200: staticwebsitegeneratedimageDBResponse
func (controller *Controller) UpdateStaticWebSiteGeneratedImage(c *gin.Context) {

	mutexStaticWebSiteGeneratedImage.Lock()
	defer mutexStaticWebSiteGeneratedImage.Unlock()

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
	db := backRepo.BackRepoStaticWebSiteGeneratedImage.GetDB()

	// Validate input
	var input orm.StaticWebSiteGeneratedImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var staticwebsitegeneratedimageDB orm.StaticWebSiteGeneratedImageDB

	// fetch the staticwebsitegeneratedimage
	_, err := db.First(&staticwebsitegeneratedimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	staticwebsitegeneratedimageDB.CopyBasicFieldsFromStaticWebSiteGeneratedImage_WOP(&input.StaticWebSiteGeneratedImage_WOP)
	staticwebsitegeneratedimageDB.StaticWebSiteGeneratedImagePointersEncoding = input.StaticWebSiteGeneratedImagePointersEncoding

	db, _ = db.Model(&staticwebsitegeneratedimageDB)
	_, err = db.Updates(&staticwebsitegeneratedimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	staticwebsitegeneratedimageNew := new(models.StaticWebSiteGeneratedImage)
	staticwebsitegeneratedimageDB.CopyBasicFieldsToStaticWebSiteGeneratedImage(staticwebsitegeneratedimageNew)

	// redeem pointers
	staticwebsitegeneratedimageDB.DecodePointers(backRepo, staticwebsitegeneratedimageNew)

	// get stage instance from DB instance, and call callback function
	staticwebsitegeneratedimageOld := backRepo.BackRepoStaticWebSiteGeneratedImage.Map_StaticWebSiteGeneratedImageDBID_StaticWebSiteGeneratedImagePtr[staticwebsitegeneratedimageDB.ID]
	if staticwebsitegeneratedimageOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), staticwebsitegeneratedimageOld, staticwebsitegeneratedimageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the staticwebsitegeneratedimageDB
	c.JSON(http.StatusOK, staticwebsitegeneratedimageDB)
}

// DeleteStaticWebSiteGeneratedImage
//
// swagger:route DELETE /staticwebsitegeneratedimages/{ID} staticwebsitegeneratedimages deleteStaticWebSiteGeneratedImage
//
// # Delete a staticwebsitegeneratedimage
//
// default: genericError
//
//	200: staticwebsitegeneratedimageDBResponse
func (controller *Controller) DeleteStaticWebSiteGeneratedImage(c *gin.Context) {

	mutexStaticWebSiteGeneratedImage.Lock()
	defer mutexStaticWebSiteGeneratedImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStaticWebSiteGeneratedImage", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteGeneratedImage.GetDB()

	// Get model if exist
	var staticwebsitegeneratedimageDB orm.StaticWebSiteGeneratedImageDB
	if _, err := db.First(&staticwebsitegeneratedimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&staticwebsitegeneratedimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	staticwebsitegeneratedimageDeleted := new(models.StaticWebSiteGeneratedImage)
	staticwebsitegeneratedimageDB.CopyBasicFieldsToStaticWebSiteGeneratedImage(staticwebsitegeneratedimageDeleted)

	// get stage instance from DB instance, and call callback function
	staticwebsitegeneratedimageStaged := backRepo.BackRepoStaticWebSiteGeneratedImage.Map_StaticWebSiteGeneratedImageDBID_StaticWebSiteGeneratedImagePtr[staticwebsitegeneratedimageDB.ID]
	if staticwebsitegeneratedimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), staticwebsitegeneratedimageStaged, staticwebsitegeneratedimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
