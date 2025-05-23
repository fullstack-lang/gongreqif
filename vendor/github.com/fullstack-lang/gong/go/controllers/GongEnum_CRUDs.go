// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongEnum__dummysDeclaration__ models.GongEnum
var __GongEnum_time__dummyDeclaration time.Duration

var mutexGongEnum sync.Mutex

// An GongEnumID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongEnum updateGongEnum deleteGongEnum
type GongEnumID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongEnumInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongEnum updateGongEnum
type GongEnumInput struct {
	// The GongEnum to submit or modify
	// in: body
	GongEnum *orm.GongEnumAPI
}

// GetGongEnums
//
// swagger:route GET /gongenums gongenums getGongEnums
//
// # Get all gongenums
//
// Responses:
// default: genericError
//
//	200: gongenumDBResponse
func (controller *Controller) GetGongEnums(c *gin.Context) {

	// source slice
	var gongenumDBs []orm.GongEnumDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongEnums", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGongEnum.GetDB()

	_, err := db.Find(&gongenumDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongenumAPIs := make([]orm.GongEnumAPI, 0)

	// for each gongenum, update fields from the database nullable fields
	for idx := range gongenumDBs {
		gongenumDB := &gongenumDBs[idx]
		_ = gongenumDB
		var gongenumAPI orm.GongEnumAPI

		// insertion point for updating fields
		gongenumAPI.ID = gongenumDB.ID
		gongenumDB.CopyBasicFieldsToGongEnum_WOP(&gongenumAPI.GongEnum_WOP)
		gongenumAPI.GongEnumPointersEncoding = gongenumDB.GongEnumPointersEncoding
		gongenumAPIs = append(gongenumAPIs, gongenumAPI)
	}

	c.JSON(http.StatusOK, gongenumAPIs)
}

// PostGongEnum
//
// swagger:route POST /gongenums gongenums postGongEnum
//
// Creates a gongenum
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongEnum(c *gin.Context) {

	mutexGongEnum.Lock()
	defer mutexGongEnum.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongEnums", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGongEnum.GetDB()

	// Validate input
	var input orm.GongEnumAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongenum
	gongenumDB := orm.GongEnumDB{}
	gongenumDB.GongEnumPointersEncoding = input.GongEnumPointersEncoding
	gongenumDB.CopyBasicFieldsFromGongEnum_WOP(&input.GongEnum_WOP)

	_, err = db.Create(&gongenumDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongEnum.CheckoutPhaseOneInstance(&gongenumDB)
	gongenum := backRepo.BackRepoGongEnum.Map_GongEnumDBID_GongEnumPtr[gongenumDB.ID]

	if gongenum != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongenum)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongenumDB)
}

// GetGongEnum
//
// swagger:route GET /gongenums/{ID} gongenums getGongEnum
//
// Gets the details for a gongenum.
//
// Responses:
// default: genericError
//
//	200: gongenumDBResponse
func (controller *Controller) GetGongEnum(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongEnum", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGongEnum.GetDB()

	// Get gongenumDB in DB
	var gongenumDB orm.GongEnumDB
	if _, err := db.First(&gongenumDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongenumAPI orm.GongEnumAPI
	gongenumAPI.ID = gongenumDB.ID
	gongenumAPI.GongEnumPointersEncoding = gongenumDB.GongEnumPointersEncoding
	gongenumDB.CopyBasicFieldsToGongEnum_WOP(&gongenumAPI.GongEnum_WOP)

	c.JSON(http.StatusOK, gongenumAPI)
}

// UpdateGongEnum
//
// swagger:route PATCH /gongenums/{ID} gongenums updateGongEnum
//
// # Update a gongenum
//
// Responses:
// default: genericError
//
//	200: gongenumDBResponse
func (controller *Controller) UpdateGongEnum(c *gin.Context) {

	mutexGongEnum.Lock()
	defer mutexGongEnum.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongEnum", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGongEnum.GetDB()

	// Validate input
	var input orm.GongEnumAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongenumDB orm.GongEnumDB

	// fetch the gongenum
	_, err := db.First(&gongenumDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongenumDB.CopyBasicFieldsFromGongEnum_WOP(&input.GongEnum_WOP)
	gongenumDB.GongEnumPointersEncoding = input.GongEnumPointersEncoding

	db, _ = db.Model(&gongenumDB)
	_, err = db.Updates(&gongenumDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongenumNew := new(models.GongEnum)
	gongenumDB.CopyBasicFieldsToGongEnum(gongenumNew)

	// redeem pointers
	gongenumDB.DecodePointers(backRepo, gongenumNew)

	// get stage instance from DB instance, and call callback function
	gongenumOld := backRepo.BackRepoGongEnum.Map_GongEnumDBID_GongEnumPtr[gongenumDB.ID]
	if gongenumOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongenumOld, gongenumNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongenumDB
	c.JSON(http.StatusOK, gongenumDB)
}

// DeleteGongEnum
//
// swagger:route DELETE /gongenums/{ID} gongenums deleteGongEnum
//
// # Delete a gongenum
//
// default: genericError
//
//	200: gongenumDBResponse
func (controller *Controller) DeleteGongEnum(c *gin.Context) {

	mutexGongEnum.Lock()
	defer mutexGongEnum.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongEnum", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGongEnum.GetDB()

	// Get model if exist
	var gongenumDB orm.GongEnumDB
	if _, err := db.First(&gongenumDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongenumDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongenumDeleted := new(models.GongEnum)
	gongenumDB.CopyBasicFieldsToGongEnum(gongenumDeleted)

	// get stage instance from DB instance, and call callback function
	gongenumStaged := backRepo.BackRepoGongEnum.Map_GongEnumDBID_GongEnumPtr[gongenumDB.ID]
	if gongenumStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongenumStaged, gongenumDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
