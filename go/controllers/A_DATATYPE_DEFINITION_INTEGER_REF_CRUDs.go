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
var __A_DATATYPE_DEFINITION_INTEGER_REF__dummysDeclaration__ models.A_DATATYPE_DEFINITION_INTEGER_REF
var __A_DATATYPE_DEFINITION_INTEGER_REF_time__dummyDeclaration time.Duration

var mutexA_DATATYPE_DEFINITION_INTEGER_REF sync.Mutex

// An A_DATATYPE_DEFINITION_INTEGER_REFID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DATATYPE_DEFINITION_INTEGER_REF updateA_DATATYPE_DEFINITION_INTEGER_REF deleteA_DATATYPE_DEFINITION_INTEGER_REF
type A_DATATYPE_DEFINITION_INTEGER_REFID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DATATYPE_DEFINITION_INTEGER_REFInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DATATYPE_DEFINITION_INTEGER_REF updateA_DATATYPE_DEFINITION_INTEGER_REF
type A_DATATYPE_DEFINITION_INTEGER_REFInput struct {
	// The A_DATATYPE_DEFINITION_INTEGER_REF to submit or modify
	// in: body
	A_DATATYPE_DEFINITION_INTEGER_REF *orm.A_DATATYPE_DEFINITION_INTEGER_REFAPI
}

// GetA_DATATYPE_DEFINITION_INTEGER_REFs
//
// swagger:route GET /a_datatype_definition_integer_refs a_datatype_definition_integer_refs getA_DATATYPE_DEFINITION_INTEGER_REFs
//
// # Get all a_datatype_definition_integer_refs
//
// Responses:
// default: genericError
//
//	200: a_datatype_definition_integer_refDBResponse
func (controller *Controller) GetA_DATATYPE_DEFINITION_INTEGER_REFs(c *gin.Context) {

	// source slice
	var a_datatype_definition_integer_refDBs []orm.A_DATATYPE_DEFINITION_INTEGER_REFDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DATATYPE_DEFINITION_INTEGER_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPE_DEFINITION_INTEGER_REF.GetDB()

	_, err := db.Find(&a_datatype_definition_integer_refDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_datatype_definition_integer_refAPIs := make([]orm.A_DATATYPE_DEFINITION_INTEGER_REFAPI, 0)

	// for each a_datatype_definition_integer_ref, update fields from the database nullable fields
	for idx := range a_datatype_definition_integer_refDBs {
		a_datatype_definition_integer_refDB := &a_datatype_definition_integer_refDBs[idx]
		_ = a_datatype_definition_integer_refDB
		var a_datatype_definition_integer_refAPI orm.A_DATATYPE_DEFINITION_INTEGER_REFAPI

		// insertion point for updating fields
		a_datatype_definition_integer_refAPI.ID = a_datatype_definition_integer_refDB.ID
		a_datatype_definition_integer_refDB.CopyBasicFieldsToA_DATATYPE_DEFINITION_INTEGER_REF_WOP(&a_datatype_definition_integer_refAPI.A_DATATYPE_DEFINITION_INTEGER_REF_WOP)
		a_datatype_definition_integer_refAPI.A_DATATYPE_DEFINITION_INTEGER_REFPointersEncoding = a_datatype_definition_integer_refDB.A_DATATYPE_DEFINITION_INTEGER_REFPointersEncoding
		a_datatype_definition_integer_refAPIs = append(a_datatype_definition_integer_refAPIs, a_datatype_definition_integer_refAPI)
	}

	c.JSON(http.StatusOK, a_datatype_definition_integer_refAPIs)
}

// PostA_DATATYPE_DEFINITION_INTEGER_REF
//
// swagger:route POST /a_datatype_definition_integer_refs a_datatype_definition_integer_refs postA_DATATYPE_DEFINITION_INTEGER_REF
//
// Creates a a_datatype_definition_integer_ref
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DATATYPE_DEFINITION_INTEGER_REF(c *gin.Context) {

	mutexA_DATATYPE_DEFINITION_INTEGER_REF.Lock()
	defer mutexA_DATATYPE_DEFINITION_INTEGER_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DATATYPE_DEFINITION_INTEGER_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPE_DEFINITION_INTEGER_REF.GetDB()

	// Validate input
	var input orm.A_DATATYPE_DEFINITION_INTEGER_REFAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_datatype_definition_integer_ref
	a_datatype_definition_integer_refDB := orm.A_DATATYPE_DEFINITION_INTEGER_REFDB{}
	a_datatype_definition_integer_refDB.A_DATATYPE_DEFINITION_INTEGER_REFPointersEncoding = input.A_DATATYPE_DEFINITION_INTEGER_REFPointersEncoding
	a_datatype_definition_integer_refDB.CopyBasicFieldsFromA_DATATYPE_DEFINITION_INTEGER_REF_WOP(&input.A_DATATYPE_DEFINITION_INTEGER_REF_WOP)

	_, err = db.Create(&a_datatype_definition_integer_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DATATYPE_DEFINITION_INTEGER_REF.CheckoutPhaseOneInstance(&a_datatype_definition_integer_refDB)
	a_datatype_definition_integer_ref := backRepo.BackRepoA_DATATYPE_DEFINITION_INTEGER_REF.Map_A_DATATYPE_DEFINITION_INTEGER_REFDBID_A_DATATYPE_DEFINITION_INTEGER_REFPtr[a_datatype_definition_integer_refDB.ID]

	if a_datatype_definition_integer_ref != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_datatype_definition_integer_ref)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_datatype_definition_integer_refDB)
}

// GetA_DATATYPE_DEFINITION_INTEGER_REF
//
// swagger:route GET /a_datatype_definition_integer_refs/{ID} a_datatype_definition_integer_refs getA_DATATYPE_DEFINITION_INTEGER_REF
//
// Gets the details for a a_datatype_definition_integer_ref.
//
// Responses:
// default: genericError
//
//	200: a_datatype_definition_integer_refDBResponse
func (controller *Controller) GetA_DATATYPE_DEFINITION_INTEGER_REF(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DATATYPE_DEFINITION_INTEGER_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPE_DEFINITION_INTEGER_REF.GetDB()

	// Get a_datatype_definition_integer_refDB in DB
	var a_datatype_definition_integer_refDB orm.A_DATATYPE_DEFINITION_INTEGER_REFDB
	if _, err := db.First(&a_datatype_definition_integer_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_datatype_definition_integer_refAPI orm.A_DATATYPE_DEFINITION_INTEGER_REFAPI
	a_datatype_definition_integer_refAPI.ID = a_datatype_definition_integer_refDB.ID
	a_datatype_definition_integer_refAPI.A_DATATYPE_DEFINITION_INTEGER_REFPointersEncoding = a_datatype_definition_integer_refDB.A_DATATYPE_DEFINITION_INTEGER_REFPointersEncoding
	a_datatype_definition_integer_refDB.CopyBasicFieldsToA_DATATYPE_DEFINITION_INTEGER_REF_WOP(&a_datatype_definition_integer_refAPI.A_DATATYPE_DEFINITION_INTEGER_REF_WOP)

	c.JSON(http.StatusOK, a_datatype_definition_integer_refAPI)
}

// UpdateA_DATATYPE_DEFINITION_INTEGER_REF
//
// swagger:route PATCH /a_datatype_definition_integer_refs/{ID} a_datatype_definition_integer_refs updateA_DATATYPE_DEFINITION_INTEGER_REF
//
// # Update a a_datatype_definition_integer_ref
//
// Responses:
// default: genericError
//
//	200: a_datatype_definition_integer_refDBResponse
func (controller *Controller) UpdateA_DATATYPE_DEFINITION_INTEGER_REF(c *gin.Context) {

	mutexA_DATATYPE_DEFINITION_INTEGER_REF.Lock()
	defer mutexA_DATATYPE_DEFINITION_INTEGER_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DATATYPE_DEFINITION_INTEGER_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPE_DEFINITION_INTEGER_REF.GetDB()

	// Validate input
	var input orm.A_DATATYPE_DEFINITION_INTEGER_REFAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_datatype_definition_integer_refDB orm.A_DATATYPE_DEFINITION_INTEGER_REFDB

	// fetch the a_datatype_definition_integer_ref
	_, err := db.First(&a_datatype_definition_integer_refDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_datatype_definition_integer_refDB.CopyBasicFieldsFromA_DATATYPE_DEFINITION_INTEGER_REF_WOP(&input.A_DATATYPE_DEFINITION_INTEGER_REF_WOP)
	a_datatype_definition_integer_refDB.A_DATATYPE_DEFINITION_INTEGER_REFPointersEncoding = input.A_DATATYPE_DEFINITION_INTEGER_REFPointersEncoding

	db, _ = db.Model(&a_datatype_definition_integer_refDB)
	_, err = db.Updates(&a_datatype_definition_integer_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_datatype_definition_integer_refNew := new(models.A_DATATYPE_DEFINITION_INTEGER_REF)
	a_datatype_definition_integer_refDB.CopyBasicFieldsToA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_refNew)

	// redeem pointers
	a_datatype_definition_integer_refDB.DecodePointers(backRepo, a_datatype_definition_integer_refNew)

	// get stage instance from DB instance, and call callback function
	a_datatype_definition_integer_refOld := backRepo.BackRepoA_DATATYPE_DEFINITION_INTEGER_REF.Map_A_DATATYPE_DEFINITION_INTEGER_REFDBID_A_DATATYPE_DEFINITION_INTEGER_REFPtr[a_datatype_definition_integer_refDB.ID]
	if a_datatype_definition_integer_refOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_datatype_definition_integer_refOld, a_datatype_definition_integer_refNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_datatype_definition_integer_refDB
	c.JSON(http.StatusOK, a_datatype_definition_integer_refDB)
}

// DeleteA_DATATYPE_DEFINITION_INTEGER_REF
//
// swagger:route DELETE /a_datatype_definition_integer_refs/{ID} a_datatype_definition_integer_refs deleteA_DATATYPE_DEFINITION_INTEGER_REF
//
// # Delete a a_datatype_definition_integer_ref
//
// default: genericError
//
//	200: a_datatype_definition_integer_refDBResponse
func (controller *Controller) DeleteA_DATATYPE_DEFINITION_INTEGER_REF(c *gin.Context) {

	mutexA_DATATYPE_DEFINITION_INTEGER_REF.Lock()
	defer mutexA_DATATYPE_DEFINITION_INTEGER_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DATATYPE_DEFINITION_INTEGER_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPE_DEFINITION_INTEGER_REF.GetDB()

	// Get model if exist
	var a_datatype_definition_integer_refDB orm.A_DATATYPE_DEFINITION_INTEGER_REFDB
	if _, err := db.First(&a_datatype_definition_integer_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_datatype_definition_integer_refDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_datatype_definition_integer_refDeleted := new(models.A_DATATYPE_DEFINITION_INTEGER_REF)
	a_datatype_definition_integer_refDB.CopyBasicFieldsToA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_refDeleted)

	// get stage instance from DB instance, and call callback function
	a_datatype_definition_integer_refStaged := backRepo.BackRepoA_DATATYPE_DEFINITION_INTEGER_REF.Map_A_DATATYPE_DEFINITION_INTEGER_REFDBID_A_DATATYPE_DEFINITION_INTEGER_REFPtr[a_datatype_definition_integer_refDB.ID]
	if a_datatype_definition_integer_refStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_datatype_definition_integer_refStaged, a_datatype_definition_integer_refDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
