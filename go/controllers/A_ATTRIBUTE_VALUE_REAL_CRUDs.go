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
var __A_ATTRIBUTE_VALUE_REAL__dummysDeclaration__ models.A_ATTRIBUTE_VALUE_REAL
var __A_ATTRIBUTE_VALUE_REAL_time__dummyDeclaration time.Duration

var mutexA_ATTRIBUTE_VALUE_REAL sync.Mutex

// An A_ATTRIBUTE_VALUE_REALID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ATTRIBUTE_VALUE_REAL updateA_ATTRIBUTE_VALUE_REAL deleteA_ATTRIBUTE_VALUE_REAL
type A_ATTRIBUTE_VALUE_REALID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ATTRIBUTE_VALUE_REALInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ATTRIBUTE_VALUE_REAL updateA_ATTRIBUTE_VALUE_REAL
type A_ATTRIBUTE_VALUE_REALInput struct {
	// The A_ATTRIBUTE_VALUE_REAL to submit or modify
	// in: body
	A_ATTRIBUTE_VALUE_REAL *orm.A_ATTRIBUTE_VALUE_REALAPI
}

// GetA_ATTRIBUTE_VALUE_REALs
//
// swagger:route GET /a_attribute_value_reals a_attribute_value_reals getA_ATTRIBUTE_VALUE_REALs
//
// # Get all a_attribute_value_reals
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_realDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_REALs(c *gin.Context) {

	// source slice
	var a_attribute_value_realDBs []orm.A_ATTRIBUTE_VALUE_REALDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_REALs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_REAL.GetDB()

	_, err := db.Find(&a_attribute_value_realDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_attribute_value_realAPIs := make([]orm.A_ATTRIBUTE_VALUE_REALAPI, 0)

	// for each a_attribute_value_real, update fields from the database nullable fields
	for idx := range a_attribute_value_realDBs {
		a_attribute_value_realDB := &a_attribute_value_realDBs[idx]
		_ = a_attribute_value_realDB
		var a_attribute_value_realAPI orm.A_ATTRIBUTE_VALUE_REALAPI

		// insertion point for updating fields
		a_attribute_value_realAPI.ID = a_attribute_value_realDB.ID
		a_attribute_value_realDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_REAL_WOP(&a_attribute_value_realAPI.A_ATTRIBUTE_VALUE_REAL_WOP)
		a_attribute_value_realAPI.A_ATTRIBUTE_VALUE_REALPointersEncoding = a_attribute_value_realDB.A_ATTRIBUTE_VALUE_REALPointersEncoding
		a_attribute_value_realAPIs = append(a_attribute_value_realAPIs, a_attribute_value_realAPI)
	}

	c.JSON(http.StatusOK, a_attribute_value_realAPIs)
}

// PostA_ATTRIBUTE_VALUE_REAL
//
// swagger:route POST /a_attribute_value_reals a_attribute_value_reals postA_ATTRIBUTE_VALUE_REAL
//
// Creates a a_attribute_value_real
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ATTRIBUTE_VALUE_REAL(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_REAL.Lock()
	defer mutexA_ATTRIBUTE_VALUE_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ATTRIBUTE_VALUE_REALs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_REAL.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_REALAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_attribute_value_real
	a_attribute_value_realDB := orm.A_ATTRIBUTE_VALUE_REALDB{}
	a_attribute_value_realDB.A_ATTRIBUTE_VALUE_REALPointersEncoding = input.A_ATTRIBUTE_VALUE_REALPointersEncoding
	a_attribute_value_realDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_REAL_WOP(&input.A_ATTRIBUTE_VALUE_REAL_WOP)

	_, err = db.Create(&a_attribute_value_realDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ATTRIBUTE_VALUE_REAL.CheckoutPhaseOneInstance(&a_attribute_value_realDB)
	a_attribute_value_real := backRepo.BackRepoA_ATTRIBUTE_VALUE_REAL.Map_A_ATTRIBUTE_VALUE_REALDBID_A_ATTRIBUTE_VALUE_REALPtr[a_attribute_value_realDB.ID]

	if a_attribute_value_real != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_attribute_value_real)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_attribute_value_realDB)
}

// GetA_ATTRIBUTE_VALUE_REAL
//
// swagger:route GET /a_attribute_value_reals/{ID} a_attribute_value_reals getA_ATTRIBUTE_VALUE_REAL
//
// Gets the details for a a_attribute_value_real.
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_realDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_REAL(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_REAL", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_REAL.GetDB()

	// Get a_attribute_value_realDB in DB
	var a_attribute_value_realDB orm.A_ATTRIBUTE_VALUE_REALDB
	if _, err := db.First(&a_attribute_value_realDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_attribute_value_realAPI orm.A_ATTRIBUTE_VALUE_REALAPI
	a_attribute_value_realAPI.ID = a_attribute_value_realDB.ID
	a_attribute_value_realAPI.A_ATTRIBUTE_VALUE_REALPointersEncoding = a_attribute_value_realDB.A_ATTRIBUTE_VALUE_REALPointersEncoding
	a_attribute_value_realDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_REAL_WOP(&a_attribute_value_realAPI.A_ATTRIBUTE_VALUE_REAL_WOP)

	c.JSON(http.StatusOK, a_attribute_value_realAPI)
}

// UpdateA_ATTRIBUTE_VALUE_REAL
//
// swagger:route PATCH /a_attribute_value_reals/{ID} a_attribute_value_reals updateA_ATTRIBUTE_VALUE_REAL
//
// # Update a a_attribute_value_real
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_realDBResponse
func (controller *Controller) UpdateA_ATTRIBUTE_VALUE_REAL(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_REAL.Lock()
	defer mutexA_ATTRIBUTE_VALUE_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_ATTRIBUTE_VALUE_REAL", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_REAL.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_REALAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_attribute_value_realDB orm.A_ATTRIBUTE_VALUE_REALDB

	// fetch the a_attribute_value_real
	_, err := db.First(&a_attribute_value_realDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_attribute_value_realDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_REAL_WOP(&input.A_ATTRIBUTE_VALUE_REAL_WOP)
	a_attribute_value_realDB.A_ATTRIBUTE_VALUE_REALPointersEncoding = input.A_ATTRIBUTE_VALUE_REALPointersEncoding

	db, _ = db.Model(&a_attribute_value_realDB)
	_, err = db.Updates(&a_attribute_value_realDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_realNew := new(models.A_ATTRIBUTE_VALUE_REAL)
	a_attribute_value_realDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_REAL(a_attribute_value_realNew)

	// redeem pointers
	a_attribute_value_realDB.DecodePointers(backRepo, a_attribute_value_realNew)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_realOld := backRepo.BackRepoA_ATTRIBUTE_VALUE_REAL.Map_A_ATTRIBUTE_VALUE_REALDBID_A_ATTRIBUTE_VALUE_REALPtr[a_attribute_value_realDB.ID]
	if a_attribute_value_realOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_realOld, a_attribute_value_realNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_attribute_value_realDB
	c.JSON(http.StatusOK, a_attribute_value_realDB)
}

// DeleteA_ATTRIBUTE_VALUE_REAL
//
// swagger:route DELETE /a_attribute_value_reals/{ID} a_attribute_value_reals deleteA_ATTRIBUTE_VALUE_REAL
//
// # Delete a a_attribute_value_real
//
// default: genericError
//
//	200: a_attribute_value_realDBResponse
func (controller *Controller) DeleteA_ATTRIBUTE_VALUE_REAL(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_REAL.Lock()
	defer mutexA_ATTRIBUTE_VALUE_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ATTRIBUTE_VALUE_REAL", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_REAL.GetDB()

	// Get model if exist
	var a_attribute_value_realDB orm.A_ATTRIBUTE_VALUE_REALDB
	if _, err := db.First(&a_attribute_value_realDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_attribute_value_realDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_realDeleted := new(models.A_ATTRIBUTE_VALUE_REAL)
	a_attribute_value_realDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_REAL(a_attribute_value_realDeleted)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_realStaged := backRepo.BackRepoA_ATTRIBUTE_VALUE_REAL.Map_A_ATTRIBUTE_VALUE_REALDBID_A_ATTRIBUTE_VALUE_REALPtr[a_attribute_value_realDB.ID]
	if a_attribute_value_realStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_attribute_value_realStaged, a_attribute_value_realDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
