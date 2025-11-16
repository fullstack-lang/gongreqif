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
var __A_ATTRIBUTE_VALUE_XHTML_1__dummysDeclaration__ models.A_ATTRIBUTE_VALUE_XHTML_1
var __A_ATTRIBUTE_VALUE_XHTML_1_time__dummyDeclaration time.Duration

var mutexA_ATTRIBUTE_VALUE_XHTML_1 sync.Mutex

// An A_ATTRIBUTE_VALUE_XHTML_1ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ATTRIBUTE_VALUE_XHTML_1 updateA_ATTRIBUTE_VALUE_XHTML_1 deleteA_ATTRIBUTE_VALUE_XHTML_1
type A_ATTRIBUTE_VALUE_XHTML_1ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ATTRIBUTE_VALUE_XHTML_1Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ATTRIBUTE_VALUE_XHTML_1 updateA_ATTRIBUTE_VALUE_XHTML_1
type A_ATTRIBUTE_VALUE_XHTML_1Input struct {
	// The A_ATTRIBUTE_VALUE_XHTML_1 to submit or modify
	// in: body
	A_ATTRIBUTE_VALUE_XHTML_1 *orm.A_ATTRIBUTE_VALUE_XHTML_1API
}

// GetA_ATTRIBUTE_VALUE_XHTML_1s
//
// swagger:route GET /a_attribute_value_xhtml_1s a_attribute_value_xhtml_1s getA_ATTRIBUTE_VALUE_XHTML_1s
//
// # Get all a_attribute_value_xhtml_1s
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_xhtml_1DBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_XHTML_1s(c *gin.Context) {

	// source slice
	var a_attribute_value_xhtml_1DBs []orm.A_ATTRIBUTE_VALUE_XHTML_1DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_XHTML_1s", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML_1.GetDB()

	_, err := db.Find(&a_attribute_value_xhtml_1DBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_attribute_value_xhtml_1APIs := make([]orm.A_ATTRIBUTE_VALUE_XHTML_1API, 0)

	// for each a_attribute_value_xhtml_1, update fields from the database nullable fields
	for idx := range a_attribute_value_xhtml_1DBs {
		a_attribute_value_xhtml_1DB := &a_attribute_value_xhtml_1DBs[idx]
		_ = a_attribute_value_xhtml_1DB
		var a_attribute_value_xhtml_1API orm.A_ATTRIBUTE_VALUE_XHTML_1API

		// insertion point for updating fields
		a_attribute_value_xhtml_1API.ID = a_attribute_value_xhtml_1DB.ID
		a_attribute_value_xhtml_1DB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_XHTML_1_WOP(&a_attribute_value_xhtml_1API.A_ATTRIBUTE_VALUE_XHTML_1_WOP)
		a_attribute_value_xhtml_1API.A_ATTRIBUTE_VALUE_XHTML_1PointersEncoding = a_attribute_value_xhtml_1DB.A_ATTRIBUTE_VALUE_XHTML_1PointersEncoding
		a_attribute_value_xhtml_1APIs = append(a_attribute_value_xhtml_1APIs, a_attribute_value_xhtml_1API)
	}

	c.JSON(http.StatusOK, a_attribute_value_xhtml_1APIs)
}

// PostA_ATTRIBUTE_VALUE_XHTML_1
//
// swagger:route POST /a_attribute_value_xhtml_1s a_attribute_value_xhtml_1s postA_ATTRIBUTE_VALUE_XHTML_1
//
// Creates a a_attribute_value_xhtml_1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ATTRIBUTE_VALUE_XHTML_1(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_XHTML_1.Lock()
	defer mutexA_ATTRIBUTE_VALUE_XHTML_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ATTRIBUTE_VALUE_XHTML_1s", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML_1.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_XHTML_1API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_attribute_value_xhtml_1
	a_attribute_value_xhtml_1DB := orm.A_ATTRIBUTE_VALUE_XHTML_1DB{}
	a_attribute_value_xhtml_1DB.A_ATTRIBUTE_VALUE_XHTML_1PointersEncoding = input.A_ATTRIBUTE_VALUE_XHTML_1PointersEncoding
	a_attribute_value_xhtml_1DB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_XHTML_1_WOP(&input.A_ATTRIBUTE_VALUE_XHTML_1_WOP)

	_, err = db.Create(&a_attribute_value_xhtml_1DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML_1.CheckoutPhaseOneInstance(&a_attribute_value_xhtml_1DB)
	a_attribute_value_xhtml_1 := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML_1.Map_A_ATTRIBUTE_VALUE_XHTML_1DBID_A_ATTRIBUTE_VALUE_XHTML_1Ptr[a_attribute_value_xhtml_1DB.ID]

	if a_attribute_value_xhtml_1 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_attribute_value_xhtml_1)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_attribute_value_xhtml_1DB)
}

// GetA_ATTRIBUTE_VALUE_XHTML_1
//
// swagger:route GET /a_attribute_value_xhtml_1s/{ID} a_attribute_value_xhtml_1s getA_ATTRIBUTE_VALUE_XHTML_1
//
// Gets the details for a a_attribute_value_xhtml_1.
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_xhtml_1DBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_XHTML_1(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_XHTML_1", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML_1.GetDB()

	// Get a_attribute_value_xhtml_1DB in DB
	var a_attribute_value_xhtml_1DB orm.A_ATTRIBUTE_VALUE_XHTML_1DB
	if _, err := db.First(&a_attribute_value_xhtml_1DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_attribute_value_xhtml_1API orm.A_ATTRIBUTE_VALUE_XHTML_1API
	a_attribute_value_xhtml_1API.ID = a_attribute_value_xhtml_1DB.ID
	a_attribute_value_xhtml_1API.A_ATTRIBUTE_VALUE_XHTML_1PointersEncoding = a_attribute_value_xhtml_1DB.A_ATTRIBUTE_VALUE_XHTML_1PointersEncoding
	a_attribute_value_xhtml_1DB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_XHTML_1_WOP(&a_attribute_value_xhtml_1API.A_ATTRIBUTE_VALUE_XHTML_1_WOP)

	c.JSON(http.StatusOK, a_attribute_value_xhtml_1API)
}

// UpdateA_ATTRIBUTE_VALUE_XHTML_1
//
// swagger:route PATCH /a_attribute_value_xhtml_1s/{ID} a_attribute_value_xhtml_1s updateA_ATTRIBUTE_VALUE_XHTML_1
//
// # Update a a_attribute_value_xhtml_1
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_xhtml_1DBResponse
func (controller *Controller) UpdateA_ATTRIBUTE_VALUE_XHTML_1(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_XHTML_1.Lock()
	defer mutexA_ATTRIBUTE_VALUE_XHTML_1.Unlock()

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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML_1.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_XHTML_1API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_attribute_value_xhtml_1DB orm.A_ATTRIBUTE_VALUE_XHTML_1DB

	// fetch the a_attribute_value_xhtml_1
	_, err := db.First(&a_attribute_value_xhtml_1DB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_attribute_value_xhtml_1DB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_XHTML_1_WOP(&input.A_ATTRIBUTE_VALUE_XHTML_1_WOP)
	a_attribute_value_xhtml_1DB.A_ATTRIBUTE_VALUE_XHTML_1PointersEncoding = input.A_ATTRIBUTE_VALUE_XHTML_1PointersEncoding

	db, _ = db.Model(&a_attribute_value_xhtml_1DB)
	_, err = db.Updates(&a_attribute_value_xhtml_1DB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_xhtml_1New := new(models.A_ATTRIBUTE_VALUE_XHTML_1)
	a_attribute_value_xhtml_1DB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1New)

	// redeem pointers
	a_attribute_value_xhtml_1DB.DecodePointers(backRepo, a_attribute_value_xhtml_1New)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_xhtml_1Old := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML_1.Map_A_ATTRIBUTE_VALUE_XHTML_1DBID_A_ATTRIBUTE_VALUE_XHTML_1Ptr[a_attribute_value_xhtml_1DB.ID]
	if a_attribute_value_xhtml_1Old != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_xhtml_1Old, a_attribute_value_xhtml_1New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_attribute_value_xhtml_1DB
	c.JSON(http.StatusOK, a_attribute_value_xhtml_1DB)
}

// DeleteA_ATTRIBUTE_VALUE_XHTML_1
//
// swagger:route DELETE /a_attribute_value_xhtml_1s/{ID} a_attribute_value_xhtml_1s deleteA_ATTRIBUTE_VALUE_XHTML_1
//
// # Delete a a_attribute_value_xhtml_1
//
// default: genericError
//
//	200: a_attribute_value_xhtml_1DBResponse
func (controller *Controller) DeleteA_ATTRIBUTE_VALUE_XHTML_1(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_XHTML_1.Lock()
	defer mutexA_ATTRIBUTE_VALUE_XHTML_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ATTRIBUTE_VALUE_XHTML_1", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML_1.GetDB()

	// Get model if exist
	var a_attribute_value_xhtml_1DB orm.A_ATTRIBUTE_VALUE_XHTML_1DB
	if _, err := db.First(&a_attribute_value_xhtml_1DB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_attribute_value_xhtml_1DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_xhtml_1Deleted := new(models.A_ATTRIBUTE_VALUE_XHTML_1)
	a_attribute_value_xhtml_1DB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1Deleted)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_xhtml_1Staged := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML_1.Map_A_ATTRIBUTE_VALUE_XHTML_1DBID_A_ATTRIBUTE_VALUE_XHTML_1Ptr[a_attribute_value_xhtml_1DB.ID]
	if a_attribute_value_xhtml_1Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_attribute_value_xhtml_1Staged, a_attribute_value_xhtml_1Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
