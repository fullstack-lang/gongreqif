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
var __A_SOURCE_SPECIFICATION_1__dummysDeclaration__ models.A_SOURCE_SPECIFICATION_1
var __A_SOURCE_SPECIFICATION_1_time__dummyDeclaration time.Duration

var mutexA_SOURCE_SPECIFICATION_1 sync.Mutex

// An A_SOURCE_SPECIFICATION_1ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SOURCE_SPECIFICATION_1 updateA_SOURCE_SPECIFICATION_1 deleteA_SOURCE_SPECIFICATION_1
type A_SOURCE_SPECIFICATION_1ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SOURCE_SPECIFICATION_1Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SOURCE_SPECIFICATION_1 updateA_SOURCE_SPECIFICATION_1
type A_SOURCE_SPECIFICATION_1Input struct {
	// The A_SOURCE_SPECIFICATION_1 to submit or modify
	// in: body
	A_SOURCE_SPECIFICATION_1 *orm.A_SOURCE_SPECIFICATION_1API
}

// GetA_SOURCE_SPECIFICATION_1s
//
// swagger:route GET /a_source_specification_1s a_source_specification_1s getA_SOURCE_SPECIFICATION_1s
//
// # Get all a_source_specification_1s
//
// Responses:
// default: genericError
//
//	200: a_source_specification_1DBResponse
func (controller *Controller) GetA_SOURCE_SPECIFICATION_1s(c *gin.Context) {

	// source slice
	var a_source_specification_1DBs []orm.A_SOURCE_SPECIFICATION_1DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SOURCE_SPECIFICATION_1s", "Name", stackPath)
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
	db := backRepo.BackRepoA_SOURCE_SPECIFICATION_1.GetDB()

	_, err := db.Find(&a_source_specification_1DBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_source_specification_1APIs := make([]orm.A_SOURCE_SPECIFICATION_1API, 0)

	// for each a_source_specification_1, update fields from the database nullable fields
	for idx := range a_source_specification_1DBs {
		a_source_specification_1DB := &a_source_specification_1DBs[idx]
		_ = a_source_specification_1DB
		var a_source_specification_1API orm.A_SOURCE_SPECIFICATION_1API

		// insertion point for updating fields
		a_source_specification_1API.ID = a_source_specification_1DB.ID
		a_source_specification_1DB.CopyBasicFieldsToA_SOURCE_SPECIFICATION_1_WOP(&a_source_specification_1API.A_SOURCE_SPECIFICATION_1_WOP)
		a_source_specification_1API.A_SOURCE_SPECIFICATION_1PointersEncoding = a_source_specification_1DB.A_SOURCE_SPECIFICATION_1PointersEncoding
		a_source_specification_1APIs = append(a_source_specification_1APIs, a_source_specification_1API)
	}

	c.JSON(http.StatusOK, a_source_specification_1APIs)
}

// PostA_SOURCE_SPECIFICATION_1
//
// swagger:route POST /a_source_specification_1s a_source_specification_1s postA_SOURCE_SPECIFICATION_1
//
// Creates a a_source_specification_1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SOURCE_SPECIFICATION_1(c *gin.Context) {

	mutexA_SOURCE_SPECIFICATION_1.Lock()
	defer mutexA_SOURCE_SPECIFICATION_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SOURCE_SPECIFICATION_1s", "Name", stackPath)
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
	db := backRepo.BackRepoA_SOURCE_SPECIFICATION_1.GetDB()

	// Validate input
	var input orm.A_SOURCE_SPECIFICATION_1API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_source_specification_1
	a_source_specification_1DB := orm.A_SOURCE_SPECIFICATION_1DB{}
	a_source_specification_1DB.A_SOURCE_SPECIFICATION_1PointersEncoding = input.A_SOURCE_SPECIFICATION_1PointersEncoding
	a_source_specification_1DB.CopyBasicFieldsFromA_SOURCE_SPECIFICATION_1_WOP(&input.A_SOURCE_SPECIFICATION_1_WOP)

	_, err = db.Create(&a_source_specification_1DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SOURCE_SPECIFICATION_1.CheckoutPhaseOneInstance(&a_source_specification_1DB)
	a_source_specification_1 := backRepo.BackRepoA_SOURCE_SPECIFICATION_1.Map_A_SOURCE_SPECIFICATION_1DBID_A_SOURCE_SPECIFICATION_1Ptr[a_source_specification_1DB.ID]

	if a_source_specification_1 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_source_specification_1)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_source_specification_1DB)
}

// GetA_SOURCE_SPECIFICATION_1
//
// swagger:route GET /a_source_specification_1s/{ID} a_source_specification_1s getA_SOURCE_SPECIFICATION_1
//
// Gets the details for a a_source_specification_1.
//
// Responses:
// default: genericError
//
//	200: a_source_specification_1DBResponse
func (controller *Controller) GetA_SOURCE_SPECIFICATION_1(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SOURCE_SPECIFICATION_1", "Name", stackPath)
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
	db := backRepo.BackRepoA_SOURCE_SPECIFICATION_1.GetDB()

	// Get a_source_specification_1DB in DB
	var a_source_specification_1DB orm.A_SOURCE_SPECIFICATION_1DB
	if _, err := db.First(&a_source_specification_1DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_source_specification_1API orm.A_SOURCE_SPECIFICATION_1API
	a_source_specification_1API.ID = a_source_specification_1DB.ID
	a_source_specification_1API.A_SOURCE_SPECIFICATION_1PointersEncoding = a_source_specification_1DB.A_SOURCE_SPECIFICATION_1PointersEncoding
	a_source_specification_1DB.CopyBasicFieldsToA_SOURCE_SPECIFICATION_1_WOP(&a_source_specification_1API.A_SOURCE_SPECIFICATION_1_WOP)

	c.JSON(http.StatusOK, a_source_specification_1API)
}

// UpdateA_SOURCE_SPECIFICATION_1
//
// swagger:route PATCH /a_source_specification_1s/{ID} a_source_specification_1s updateA_SOURCE_SPECIFICATION_1
//
// # Update a a_source_specification_1
//
// Responses:
// default: genericError
//
//	200: a_source_specification_1DBResponse
func (controller *Controller) UpdateA_SOURCE_SPECIFICATION_1(c *gin.Context) {

	mutexA_SOURCE_SPECIFICATION_1.Lock()
	defer mutexA_SOURCE_SPECIFICATION_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_SOURCE_SPECIFICATION_1", "Name", stackPath)
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
	db := backRepo.BackRepoA_SOURCE_SPECIFICATION_1.GetDB()

	// Validate input
	var input orm.A_SOURCE_SPECIFICATION_1API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_source_specification_1DB orm.A_SOURCE_SPECIFICATION_1DB

	// fetch the a_source_specification_1
	_, err := db.First(&a_source_specification_1DB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_source_specification_1DB.CopyBasicFieldsFromA_SOURCE_SPECIFICATION_1_WOP(&input.A_SOURCE_SPECIFICATION_1_WOP)
	a_source_specification_1DB.A_SOURCE_SPECIFICATION_1PointersEncoding = input.A_SOURCE_SPECIFICATION_1PointersEncoding

	db, _ = db.Model(&a_source_specification_1DB)
	_, err = db.Updates(&a_source_specification_1DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_source_specification_1New := new(models.A_SOURCE_SPECIFICATION_1)
	a_source_specification_1DB.CopyBasicFieldsToA_SOURCE_SPECIFICATION_1(a_source_specification_1New)

	// redeem pointers
	a_source_specification_1DB.DecodePointers(backRepo, a_source_specification_1New)

	// get stage instance from DB instance, and call callback function
	a_source_specification_1Old := backRepo.BackRepoA_SOURCE_SPECIFICATION_1.Map_A_SOURCE_SPECIFICATION_1DBID_A_SOURCE_SPECIFICATION_1Ptr[a_source_specification_1DB.ID]
	if a_source_specification_1Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_source_specification_1Old, a_source_specification_1New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_source_specification_1DB
	c.JSON(http.StatusOK, a_source_specification_1DB)
}

// DeleteA_SOURCE_SPECIFICATION_1
//
// swagger:route DELETE /a_source_specification_1s/{ID} a_source_specification_1s deleteA_SOURCE_SPECIFICATION_1
//
// # Delete a a_source_specification_1
//
// default: genericError
//
//	200: a_source_specification_1DBResponse
func (controller *Controller) DeleteA_SOURCE_SPECIFICATION_1(c *gin.Context) {

	mutexA_SOURCE_SPECIFICATION_1.Lock()
	defer mutexA_SOURCE_SPECIFICATION_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SOURCE_SPECIFICATION_1", "Name", stackPath)
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
	db := backRepo.BackRepoA_SOURCE_SPECIFICATION_1.GetDB()

	// Get model if exist
	var a_source_specification_1DB orm.A_SOURCE_SPECIFICATION_1DB
	if _, err := db.First(&a_source_specification_1DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_source_specification_1DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_source_specification_1Deleted := new(models.A_SOURCE_SPECIFICATION_1)
	a_source_specification_1DB.CopyBasicFieldsToA_SOURCE_SPECIFICATION_1(a_source_specification_1Deleted)

	// get stage instance from DB instance, and call callback function
	a_source_specification_1Staged := backRepo.BackRepoA_SOURCE_SPECIFICATION_1.Map_A_SOURCE_SPECIFICATION_1DBID_A_SOURCE_SPECIFICATION_1Ptr[a_source_specification_1DB.ID]
	if a_source_specification_1Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_source_specification_1Staged, a_source_specification_1Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
