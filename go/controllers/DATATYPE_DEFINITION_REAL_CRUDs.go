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
var __DATATYPE_DEFINITION_REAL__dummysDeclaration__ models.DATATYPE_DEFINITION_REAL
var __DATATYPE_DEFINITION_REAL_time__dummyDeclaration time.Duration

var mutexDATATYPE_DEFINITION_REAL sync.Mutex

// An DATATYPE_DEFINITION_REALID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPE_DEFINITION_REAL updateDATATYPE_DEFINITION_REAL deleteDATATYPE_DEFINITION_REAL
type DATATYPE_DEFINITION_REALID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPE_DEFINITION_REALInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPE_DEFINITION_REAL updateDATATYPE_DEFINITION_REAL
type DATATYPE_DEFINITION_REALInput struct {
	// The DATATYPE_DEFINITION_REAL to submit or modify
	// in: body
	DATATYPE_DEFINITION_REAL *orm.DATATYPE_DEFINITION_REALAPI
}

// GetDATATYPE_DEFINITION_REALs
//
// swagger:route GET /datatype_definition_reals datatype_definition_reals getDATATYPE_DEFINITION_REALs
//
// # Get all datatype_definition_reals
//
// Responses:
// default: genericError
//
//	200: datatype_definition_realDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_REALs(c *gin.Context) {

	// source slice
	var datatype_definition_realDBs []orm.DATATYPE_DEFINITION_REALDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_REALs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_REAL.GetDB()

	_, err := db.Find(&datatype_definition_realDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatype_definition_realAPIs := make([]orm.DATATYPE_DEFINITION_REALAPI, 0)

	// for each datatype_definition_real, update fields from the database nullable fields
	for idx := range datatype_definition_realDBs {
		datatype_definition_realDB := &datatype_definition_realDBs[idx]
		_ = datatype_definition_realDB
		var datatype_definition_realAPI orm.DATATYPE_DEFINITION_REALAPI

		// insertion point for updating fields
		datatype_definition_realAPI.ID = datatype_definition_realDB.ID
		datatype_definition_realDB.CopyBasicFieldsToDATATYPE_DEFINITION_REAL_WOP(&datatype_definition_realAPI.DATATYPE_DEFINITION_REAL_WOP)
		datatype_definition_realAPI.DATATYPE_DEFINITION_REALPointersEncoding = datatype_definition_realDB.DATATYPE_DEFINITION_REALPointersEncoding
		datatype_definition_realAPIs = append(datatype_definition_realAPIs, datatype_definition_realAPI)
	}

	c.JSON(http.StatusOK, datatype_definition_realAPIs)
}

// PostDATATYPE_DEFINITION_REAL
//
// swagger:route POST /datatype_definition_reals datatype_definition_reals postDATATYPE_DEFINITION_REAL
//
// Creates a datatype_definition_real
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPE_DEFINITION_REAL(c *gin.Context) {

	mutexDATATYPE_DEFINITION_REAL.Lock()
	defer mutexDATATYPE_DEFINITION_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPE_DEFINITION_REALs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_REAL.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_REALAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatype_definition_real
	datatype_definition_realDB := orm.DATATYPE_DEFINITION_REALDB{}
	datatype_definition_realDB.DATATYPE_DEFINITION_REALPointersEncoding = input.DATATYPE_DEFINITION_REALPointersEncoding
	datatype_definition_realDB.CopyBasicFieldsFromDATATYPE_DEFINITION_REAL_WOP(&input.DATATYPE_DEFINITION_REAL_WOP)

	_, err = db.Create(&datatype_definition_realDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.CheckoutPhaseOneInstance(&datatype_definition_realDB)
	datatype_definition_real := backRepo.BackRepoDATATYPE_DEFINITION_REAL.Map_DATATYPE_DEFINITION_REALDBID_DATATYPE_DEFINITION_REALPtr[datatype_definition_realDB.ID]

	if datatype_definition_real != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatype_definition_real)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatype_definition_realDB)
}

// GetDATATYPE_DEFINITION_REAL
//
// swagger:route GET /datatype_definition_reals/{ID} datatype_definition_reals getDATATYPE_DEFINITION_REAL
//
// Gets the details for a datatype_definition_real.
//
// Responses:
// default: genericError
//
//	200: datatype_definition_realDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_REAL(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_REAL", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_REAL.GetDB()

	// Get datatype_definition_realDB in DB
	var datatype_definition_realDB orm.DATATYPE_DEFINITION_REALDB
	if _, err := db.First(&datatype_definition_realDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatype_definition_realAPI orm.DATATYPE_DEFINITION_REALAPI
	datatype_definition_realAPI.ID = datatype_definition_realDB.ID
	datatype_definition_realAPI.DATATYPE_DEFINITION_REALPointersEncoding = datatype_definition_realDB.DATATYPE_DEFINITION_REALPointersEncoding
	datatype_definition_realDB.CopyBasicFieldsToDATATYPE_DEFINITION_REAL_WOP(&datatype_definition_realAPI.DATATYPE_DEFINITION_REAL_WOP)

	c.JSON(http.StatusOK, datatype_definition_realAPI)
}

// UpdateDATATYPE_DEFINITION_REAL
//
// swagger:route PATCH /datatype_definition_reals/{ID} datatype_definition_reals updateDATATYPE_DEFINITION_REAL
//
// # Update a datatype_definition_real
//
// Responses:
// default: genericError
//
//	200: datatype_definition_realDBResponse
func (controller *Controller) UpdateDATATYPE_DEFINITION_REAL(c *gin.Context) {

	mutexDATATYPE_DEFINITION_REAL.Lock()
	defer mutexDATATYPE_DEFINITION_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDATATYPE_DEFINITION_REAL", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_REAL.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_REALAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatype_definition_realDB orm.DATATYPE_DEFINITION_REALDB

	// fetch the datatype_definition_real
	_, err := db.First(&datatype_definition_realDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatype_definition_realDB.CopyBasicFieldsFromDATATYPE_DEFINITION_REAL_WOP(&input.DATATYPE_DEFINITION_REAL_WOP)
	datatype_definition_realDB.DATATYPE_DEFINITION_REALPointersEncoding = input.DATATYPE_DEFINITION_REALPointersEncoding

	db, _ = db.Model(&datatype_definition_realDB)
	_, err = db.Updates(&datatype_definition_realDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_realNew := new(models.DATATYPE_DEFINITION_REAL)
	datatype_definition_realDB.CopyBasicFieldsToDATATYPE_DEFINITION_REAL(datatype_definition_realNew)

	// redeem pointers
	datatype_definition_realDB.DecodePointers(backRepo, datatype_definition_realNew)

	// get stage instance from DB instance, and call callback function
	datatype_definition_realOld := backRepo.BackRepoDATATYPE_DEFINITION_REAL.Map_DATATYPE_DEFINITION_REALDBID_DATATYPE_DEFINITION_REALPtr[datatype_definition_realDB.ID]
	if datatype_definition_realOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), datatype_definition_realOld, datatype_definition_realNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatype_definition_realDB
	c.JSON(http.StatusOK, datatype_definition_realDB)
}

// DeleteDATATYPE_DEFINITION_REAL
//
// swagger:route DELETE /datatype_definition_reals/{ID} datatype_definition_reals deleteDATATYPE_DEFINITION_REAL
//
// # Delete a datatype_definition_real
//
// default: genericError
//
//	200: datatype_definition_realDBResponse
func (controller *Controller) DeleteDATATYPE_DEFINITION_REAL(c *gin.Context) {

	mutexDATATYPE_DEFINITION_REAL.Lock()
	defer mutexDATATYPE_DEFINITION_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPE_DEFINITION_REAL", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_REAL.GetDB()

	// Get model if exist
	var datatype_definition_realDB orm.DATATYPE_DEFINITION_REALDB
	if _, err := db.First(&datatype_definition_realDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&datatype_definition_realDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_realDeleted := new(models.DATATYPE_DEFINITION_REAL)
	datatype_definition_realDB.CopyBasicFieldsToDATATYPE_DEFINITION_REAL(datatype_definition_realDeleted)

	// get stage instance from DB instance, and call callback function
	datatype_definition_realStaged := backRepo.BackRepoDATATYPE_DEFINITION_REAL.Map_DATATYPE_DEFINITION_REALDBID_DATATYPE_DEFINITION_REALPtr[datatype_definition_realDB.ID]
	if datatype_definition_realStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatype_definition_realStaged, datatype_definition_realDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
