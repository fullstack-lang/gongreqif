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
var __A_ENUM_VALUE_REF__dummysDeclaration__ models.A_ENUM_VALUE_REF
var __A_ENUM_VALUE_REF_time__dummyDeclaration time.Duration

var mutexA_ENUM_VALUE_REF sync.Mutex

// An A_ENUM_VALUE_REFID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ENUM_VALUE_REF updateA_ENUM_VALUE_REF deleteA_ENUM_VALUE_REF
type A_ENUM_VALUE_REFID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ENUM_VALUE_REFInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ENUM_VALUE_REF updateA_ENUM_VALUE_REF
type A_ENUM_VALUE_REFInput struct {
	// The A_ENUM_VALUE_REF to submit or modify
	// in: body
	A_ENUM_VALUE_REF *orm.A_ENUM_VALUE_REFAPI
}

// GetA_ENUM_VALUE_REFs
//
// swagger:route GET /a_enum_value_refs a_enum_value_refs getA_ENUM_VALUE_REFs
//
// # Get all a_enum_value_refs
//
// Responses:
// default: genericError
//
//	200: a_enum_value_refDBResponse
func (controller *Controller) GetA_ENUM_VALUE_REFs(c *gin.Context) {

	// source slice
	var a_enum_value_refDBs []orm.A_ENUM_VALUE_REFDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ENUM_VALUE_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ENUM_VALUE_REF.GetDB()

	_, err := db.Find(&a_enum_value_refDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_enum_value_refAPIs := make([]orm.A_ENUM_VALUE_REFAPI, 0)

	// for each a_enum_value_ref, update fields from the database nullable fields
	for idx := range a_enum_value_refDBs {
		a_enum_value_refDB := &a_enum_value_refDBs[idx]
		_ = a_enum_value_refDB
		var a_enum_value_refAPI orm.A_ENUM_VALUE_REFAPI

		// insertion point for updating fields
		a_enum_value_refAPI.ID = a_enum_value_refDB.ID
		a_enum_value_refDB.CopyBasicFieldsToA_ENUM_VALUE_REF_WOP(&a_enum_value_refAPI.A_ENUM_VALUE_REF_WOP)
		a_enum_value_refAPI.A_ENUM_VALUE_REFPointersEncoding = a_enum_value_refDB.A_ENUM_VALUE_REFPointersEncoding
		a_enum_value_refAPIs = append(a_enum_value_refAPIs, a_enum_value_refAPI)
	}

	c.JSON(http.StatusOK, a_enum_value_refAPIs)
}

// PostA_ENUM_VALUE_REF
//
// swagger:route POST /a_enum_value_refs a_enum_value_refs postA_ENUM_VALUE_REF
//
// Creates a a_enum_value_ref
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ENUM_VALUE_REF(c *gin.Context) {

	mutexA_ENUM_VALUE_REF.Lock()
	defer mutexA_ENUM_VALUE_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ENUM_VALUE_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ENUM_VALUE_REF.GetDB()

	// Validate input
	var input orm.A_ENUM_VALUE_REFAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_enum_value_ref
	a_enum_value_refDB := orm.A_ENUM_VALUE_REFDB{}
	a_enum_value_refDB.A_ENUM_VALUE_REFPointersEncoding = input.A_ENUM_VALUE_REFPointersEncoding
	a_enum_value_refDB.CopyBasicFieldsFromA_ENUM_VALUE_REF_WOP(&input.A_ENUM_VALUE_REF_WOP)

	_, err = db.Create(&a_enum_value_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ENUM_VALUE_REF.CheckoutPhaseOneInstance(&a_enum_value_refDB)
	a_enum_value_ref := backRepo.BackRepoA_ENUM_VALUE_REF.Map_A_ENUM_VALUE_REFDBID_A_ENUM_VALUE_REFPtr[a_enum_value_refDB.ID]

	if a_enum_value_ref != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_enum_value_ref)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_enum_value_refDB)
}

// GetA_ENUM_VALUE_REF
//
// swagger:route GET /a_enum_value_refs/{ID} a_enum_value_refs getA_ENUM_VALUE_REF
//
// Gets the details for a a_enum_value_ref.
//
// Responses:
// default: genericError
//
//	200: a_enum_value_refDBResponse
func (controller *Controller) GetA_ENUM_VALUE_REF(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ENUM_VALUE_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_ENUM_VALUE_REF.GetDB()

	// Get a_enum_value_refDB in DB
	var a_enum_value_refDB orm.A_ENUM_VALUE_REFDB
	if _, err := db.First(&a_enum_value_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_enum_value_refAPI orm.A_ENUM_VALUE_REFAPI
	a_enum_value_refAPI.ID = a_enum_value_refDB.ID
	a_enum_value_refAPI.A_ENUM_VALUE_REFPointersEncoding = a_enum_value_refDB.A_ENUM_VALUE_REFPointersEncoding
	a_enum_value_refDB.CopyBasicFieldsToA_ENUM_VALUE_REF_WOP(&a_enum_value_refAPI.A_ENUM_VALUE_REF_WOP)

	c.JSON(http.StatusOK, a_enum_value_refAPI)
}

// UpdateA_ENUM_VALUE_REF
//
// swagger:route PATCH /a_enum_value_refs/{ID} a_enum_value_refs updateA_ENUM_VALUE_REF
//
// # Update a a_enum_value_ref
//
// Responses:
// default: genericError
//
//	200: a_enum_value_refDBResponse
func (controller *Controller) UpdateA_ENUM_VALUE_REF(c *gin.Context) {

	mutexA_ENUM_VALUE_REF.Lock()
	defer mutexA_ENUM_VALUE_REF.Unlock()

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
	db := backRepo.BackRepoA_ENUM_VALUE_REF.GetDB()

	// Validate input
	var input orm.A_ENUM_VALUE_REFAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_enum_value_refDB orm.A_ENUM_VALUE_REFDB

	// fetch the a_enum_value_ref
	_, err := db.First(&a_enum_value_refDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_enum_value_refDB.CopyBasicFieldsFromA_ENUM_VALUE_REF_WOP(&input.A_ENUM_VALUE_REF_WOP)
	a_enum_value_refDB.A_ENUM_VALUE_REFPointersEncoding = input.A_ENUM_VALUE_REFPointersEncoding

	db, _ = db.Model(&a_enum_value_refDB)
	_, err = db.Updates(&a_enum_value_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_enum_value_refNew := new(models.A_ENUM_VALUE_REF)
	a_enum_value_refDB.CopyBasicFieldsToA_ENUM_VALUE_REF(a_enum_value_refNew)

	// redeem pointers
	a_enum_value_refDB.DecodePointers(backRepo, a_enum_value_refNew)

	// get stage instance from DB instance, and call callback function
	a_enum_value_refOld := backRepo.BackRepoA_ENUM_VALUE_REF.Map_A_ENUM_VALUE_REFDBID_A_ENUM_VALUE_REFPtr[a_enum_value_refDB.ID]
	if a_enum_value_refOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), a_enum_value_refOld, a_enum_value_refNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_enum_value_refDB
	c.JSON(http.StatusOK, a_enum_value_refDB)
}

// DeleteA_ENUM_VALUE_REF
//
// swagger:route DELETE /a_enum_value_refs/{ID} a_enum_value_refs deleteA_ENUM_VALUE_REF
//
// # Delete a a_enum_value_ref
//
// default: genericError
//
//	200: a_enum_value_refDBResponse
func (controller *Controller) DeleteA_ENUM_VALUE_REF(c *gin.Context) {

	mutexA_ENUM_VALUE_REF.Lock()
	defer mutexA_ENUM_VALUE_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ENUM_VALUE_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_ENUM_VALUE_REF.GetDB()

	// Get model if exist
	var a_enum_value_refDB orm.A_ENUM_VALUE_REFDB
	if _, err := db.First(&a_enum_value_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_enum_value_refDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_enum_value_refDeleted := new(models.A_ENUM_VALUE_REF)
	a_enum_value_refDB.CopyBasicFieldsToA_ENUM_VALUE_REF(a_enum_value_refDeleted)

	// get stage instance from DB instance, and call callback function
	a_enum_value_refStaged := backRepo.BackRepoA_ENUM_VALUE_REF.Map_A_ENUM_VALUE_REFDBID_A_ENUM_VALUE_REFPtr[a_enum_value_refDB.ID]
	if a_enum_value_refStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_enum_value_refStaged, a_enum_value_refDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
