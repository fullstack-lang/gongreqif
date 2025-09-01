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
var __StaticWebSiteImage__dummysDeclaration__ models.StaticWebSiteImage
var __StaticWebSiteImage_time__dummyDeclaration time.Duration

var mutexStaticWebSiteImage sync.Mutex

// An StaticWebSiteImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStaticWebSiteImage updateStaticWebSiteImage deleteStaticWebSiteImage
type StaticWebSiteImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StaticWebSiteImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStaticWebSiteImage updateStaticWebSiteImage
type StaticWebSiteImageInput struct {
	// The StaticWebSiteImage to submit or modify
	// in: body
	StaticWebSiteImage *orm.StaticWebSiteImageAPI
}

// GetStaticWebSiteImages
//
// swagger:route GET /staticwebsiteimages staticwebsiteimages getStaticWebSiteImages
//
// # Get all staticwebsiteimages
//
// Responses:
// default: genericError
//
//	200: staticwebsiteimageDBResponse
func (controller *Controller) GetStaticWebSiteImages(c *gin.Context) {

	// source slice
	var staticwebsiteimageDBs []orm.StaticWebSiteImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaticWebSiteImages", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteImage.GetDB()

	_, err := db.Find(&staticwebsiteimageDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	staticwebsiteimageAPIs := make([]orm.StaticWebSiteImageAPI, 0)

	// for each staticwebsiteimage, update fields from the database nullable fields
	for idx := range staticwebsiteimageDBs {
		staticwebsiteimageDB := &staticwebsiteimageDBs[idx]
		_ = staticwebsiteimageDB
		var staticwebsiteimageAPI orm.StaticWebSiteImageAPI

		// insertion point for updating fields
		staticwebsiteimageAPI.ID = staticwebsiteimageDB.ID
		staticwebsiteimageDB.CopyBasicFieldsToStaticWebSiteImage_WOP(&staticwebsiteimageAPI.StaticWebSiteImage_WOP)
		staticwebsiteimageAPI.StaticWebSiteImagePointersEncoding = staticwebsiteimageDB.StaticWebSiteImagePointersEncoding
		staticwebsiteimageAPIs = append(staticwebsiteimageAPIs, staticwebsiteimageAPI)
	}

	c.JSON(http.StatusOK, staticwebsiteimageAPIs)
}

// PostStaticWebSiteImage
//
// swagger:route POST /staticwebsiteimages staticwebsiteimages postStaticWebSiteImage
//
// Creates a staticwebsiteimage
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStaticWebSiteImage(c *gin.Context) {

	mutexStaticWebSiteImage.Lock()
	defer mutexStaticWebSiteImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStaticWebSiteImages", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteImage.GetDB()

	// Validate input
	var input orm.StaticWebSiteImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create staticwebsiteimage
	staticwebsiteimageDB := orm.StaticWebSiteImageDB{}
	staticwebsiteimageDB.StaticWebSiteImagePointersEncoding = input.StaticWebSiteImagePointersEncoding
	staticwebsiteimageDB.CopyBasicFieldsFromStaticWebSiteImage_WOP(&input.StaticWebSiteImage_WOP)

	_, err = db.Create(&staticwebsiteimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStaticWebSiteImage.CheckoutPhaseOneInstance(&staticwebsiteimageDB)
	staticwebsiteimage := backRepo.BackRepoStaticWebSiteImage.Map_StaticWebSiteImageDBID_StaticWebSiteImagePtr[staticwebsiteimageDB.ID]

	if staticwebsiteimage != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), staticwebsiteimage)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, staticwebsiteimageDB)
}

// GetStaticWebSiteImage
//
// swagger:route GET /staticwebsiteimages/{ID} staticwebsiteimages getStaticWebSiteImage
//
// Gets the details for a staticwebsiteimage.
//
// Responses:
// default: genericError
//
//	200: staticwebsiteimageDBResponse
func (controller *Controller) GetStaticWebSiteImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaticWebSiteImage", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteImage.GetDB()

	// Get staticwebsiteimageDB in DB
	var staticwebsiteimageDB orm.StaticWebSiteImageDB
	if _, err := db.First(&staticwebsiteimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var staticwebsiteimageAPI orm.StaticWebSiteImageAPI
	staticwebsiteimageAPI.ID = staticwebsiteimageDB.ID
	staticwebsiteimageAPI.StaticWebSiteImagePointersEncoding = staticwebsiteimageDB.StaticWebSiteImagePointersEncoding
	staticwebsiteimageDB.CopyBasicFieldsToStaticWebSiteImage_WOP(&staticwebsiteimageAPI.StaticWebSiteImage_WOP)

	c.JSON(http.StatusOK, staticwebsiteimageAPI)
}

// UpdateStaticWebSiteImage
//
// swagger:route PATCH /staticwebsiteimages/{ID} staticwebsiteimages updateStaticWebSiteImage
//
// # Update a staticwebsiteimage
//
// Responses:
// default: genericError
//
//	200: staticwebsiteimageDBResponse
func (controller *Controller) UpdateStaticWebSiteImage(c *gin.Context) {

	mutexStaticWebSiteImage.Lock()
	defer mutexStaticWebSiteImage.Unlock()

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
	db := backRepo.BackRepoStaticWebSiteImage.GetDB()

	// Validate input
	var input orm.StaticWebSiteImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var staticwebsiteimageDB orm.StaticWebSiteImageDB

	// fetch the staticwebsiteimage
	_, err := db.First(&staticwebsiteimageDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	staticwebsiteimageDB.CopyBasicFieldsFromStaticWebSiteImage_WOP(&input.StaticWebSiteImage_WOP)
	staticwebsiteimageDB.StaticWebSiteImagePointersEncoding = input.StaticWebSiteImagePointersEncoding

	db, _ = db.Model(&staticwebsiteimageDB)
	_, err = db.Updates(&staticwebsiteimageDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	staticwebsiteimageNew := new(models.StaticWebSiteImage)
	staticwebsiteimageDB.CopyBasicFieldsToStaticWebSiteImage(staticwebsiteimageNew)

	// redeem pointers
	staticwebsiteimageDB.DecodePointers(backRepo, staticwebsiteimageNew)

	// get stage instance from DB instance, and call callback function
	staticwebsiteimageOld := backRepo.BackRepoStaticWebSiteImage.Map_StaticWebSiteImageDBID_StaticWebSiteImagePtr[staticwebsiteimageDB.ID]
	if staticwebsiteimageOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), staticwebsiteimageOld, staticwebsiteimageNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), staticwebsiteimageOld, staticwebsiteimageNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the staticwebsiteimageDB
	c.JSON(http.StatusOK, staticwebsiteimageDB)
}

// DeleteStaticWebSiteImage
//
// swagger:route DELETE /staticwebsiteimages/{ID} staticwebsiteimages deleteStaticWebSiteImage
//
// # Delete a staticwebsiteimage
//
// default: genericError
//
//	200: staticwebsiteimageDBResponse
func (controller *Controller) DeleteStaticWebSiteImage(c *gin.Context) {

	mutexStaticWebSiteImage.Lock()
	defer mutexStaticWebSiteImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStaticWebSiteImage", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteImage.GetDB()

	// Get model if exist
	var staticwebsiteimageDB orm.StaticWebSiteImageDB
	if _, err := db.First(&staticwebsiteimageDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&staticwebsiteimageDB)

	// get an instance (not staged) from DB instance, and call callback function
	staticwebsiteimageDeleted := new(models.StaticWebSiteImage)
	staticwebsiteimageDB.CopyBasicFieldsToStaticWebSiteImage(staticwebsiteimageDeleted)

	// get stage instance from DB instance, and call callback function
	staticwebsiteimageStaged := backRepo.BackRepoStaticWebSiteImage.Map_StaticWebSiteImageDBID_StaticWebSiteImagePtr[staticwebsiteimageDB.ID]
	if staticwebsiteimageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), staticwebsiteimageStaged, staticwebsiteimageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
