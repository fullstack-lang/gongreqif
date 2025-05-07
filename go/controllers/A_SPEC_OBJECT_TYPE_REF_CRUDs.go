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
var __A_SPEC_OBJECT_TYPE_REF__dummysDeclaration__ models.A_SPEC_OBJECT_TYPE_REF
var __A_SPEC_OBJECT_TYPE_REF_time__dummyDeclaration time.Duration

var mutexA_SPEC_OBJECT_TYPE_REF sync.Mutex

// An A_SPEC_OBJECT_TYPE_REFID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SPEC_OBJECT_TYPE_REF updateA_SPEC_OBJECT_TYPE_REF deleteA_SPEC_OBJECT_TYPE_REF
type A_SPEC_OBJECT_TYPE_REFID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SPEC_OBJECT_TYPE_REFInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SPEC_OBJECT_TYPE_REF updateA_SPEC_OBJECT_TYPE_REF
type A_SPEC_OBJECT_TYPE_REFInput struct {
	// The A_SPEC_OBJECT_TYPE_REF to submit or modify
	// in: body
	A_SPEC_OBJECT_TYPE_REF *orm.A_SPEC_OBJECT_TYPE_REFAPI
}

// GetA_SPEC_OBJECT_TYPE_REFs
//
// swagger:route GET /a_spec_object_type_refs a_spec_object_type_refs getA_SPEC_OBJECT_TYPE_REFs
//
// # Get all a_spec_object_type_refs
//
// Responses:
// default: genericError
//
//	200: a_spec_object_type_refDBResponse
func (controller *Controller) GetA_SPEC_OBJECT_TYPE_REFs(c *gin.Context) {

	// source slice
	var a_spec_object_type_refDBs []orm.A_SPEC_OBJECT_TYPE_REFDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_OBJECT_TYPE_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_OBJECT_TYPE_REF.GetDB()

	_, err := db.Find(&a_spec_object_type_refDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_spec_object_type_refAPIs := make([]orm.A_SPEC_OBJECT_TYPE_REFAPI, 0)

	// for each a_spec_object_type_ref, update fields from the database nullable fields
	for idx := range a_spec_object_type_refDBs {
		a_spec_object_type_refDB := &a_spec_object_type_refDBs[idx]
		_ = a_spec_object_type_refDB
		var a_spec_object_type_refAPI orm.A_SPEC_OBJECT_TYPE_REFAPI

		// insertion point for updating fields
		a_spec_object_type_refAPI.ID = a_spec_object_type_refDB.ID
		a_spec_object_type_refDB.CopyBasicFieldsToA_SPEC_OBJECT_TYPE_REF_WOP(&a_spec_object_type_refAPI.A_SPEC_OBJECT_TYPE_REF_WOP)
		a_spec_object_type_refAPI.A_SPEC_OBJECT_TYPE_REFPointersEncoding = a_spec_object_type_refDB.A_SPEC_OBJECT_TYPE_REFPointersEncoding
		a_spec_object_type_refAPIs = append(a_spec_object_type_refAPIs, a_spec_object_type_refAPI)
	}

	c.JSON(http.StatusOK, a_spec_object_type_refAPIs)
}

// PostA_SPEC_OBJECT_TYPE_REF
//
// swagger:route POST /a_spec_object_type_refs a_spec_object_type_refs postA_SPEC_OBJECT_TYPE_REF
//
// Creates a a_spec_object_type_ref
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SPEC_OBJECT_TYPE_REF(c *gin.Context) {

	mutexA_SPEC_OBJECT_TYPE_REF.Lock()
	defer mutexA_SPEC_OBJECT_TYPE_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SPEC_OBJECT_TYPE_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_OBJECT_TYPE_REF.GetDB()

	// Validate input
	var input orm.A_SPEC_OBJECT_TYPE_REFAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_spec_object_type_ref
	a_spec_object_type_refDB := orm.A_SPEC_OBJECT_TYPE_REFDB{}
	a_spec_object_type_refDB.A_SPEC_OBJECT_TYPE_REFPointersEncoding = input.A_SPEC_OBJECT_TYPE_REFPointersEncoding
	a_spec_object_type_refDB.CopyBasicFieldsFromA_SPEC_OBJECT_TYPE_REF_WOP(&input.A_SPEC_OBJECT_TYPE_REF_WOP)

	_, err = db.Create(&a_spec_object_type_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SPEC_OBJECT_TYPE_REF.CheckoutPhaseOneInstance(&a_spec_object_type_refDB)
	a_spec_object_type_ref := backRepo.BackRepoA_SPEC_OBJECT_TYPE_REF.Map_A_SPEC_OBJECT_TYPE_REFDBID_A_SPEC_OBJECT_TYPE_REFPtr[a_spec_object_type_refDB.ID]

	if a_spec_object_type_ref != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_spec_object_type_ref)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_spec_object_type_refDB)
}

// GetA_SPEC_OBJECT_TYPE_REF
//
// swagger:route GET /a_spec_object_type_refs/{ID} a_spec_object_type_refs getA_SPEC_OBJECT_TYPE_REF
//
// Gets the details for a a_spec_object_type_ref.
//
// Responses:
// default: genericError
//
//	200: a_spec_object_type_refDBResponse
func (controller *Controller) GetA_SPEC_OBJECT_TYPE_REF(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_OBJECT_TYPE_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_OBJECT_TYPE_REF.GetDB()

	// Get a_spec_object_type_refDB in DB
	var a_spec_object_type_refDB orm.A_SPEC_OBJECT_TYPE_REFDB
	if _, err := db.First(&a_spec_object_type_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_spec_object_type_refAPI orm.A_SPEC_OBJECT_TYPE_REFAPI
	a_spec_object_type_refAPI.ID = a_spec_object_type_refDB.ID
	a_spec_object_type_refAPI.A_SPEC_OBJECT_TYPE_REFPointersEncoding = a_spec_object_type_refDB.A_SPEC_OBJECT_TYPE_REFPointersEncoding
	a_spec_object_type_refDB.CopyBasicFieldsToA_SPEC_OBJECT_TYPE_REF_WOP(&a_spec_object_type_refAPI.A_SPEC_OBJECT_TYPE_REF_WOP)

	c.JSON(http.StatusOK, a_spec_object_type_refAPI)
}

// UpdateA_SPEC_OBJECT_TYPE_REF
//
// swagger:route PATCH /a_spec_object_type_refs/{ID} a_spec_object_type_refs updateA_SPEC_OBJECT_TYPE_REF
//
// # Update a a_spec_object_type_ref
//
// Responses:
// default: genericError
//
//	200: a_spec_object_type_refDBResponse
func (controller *Controller) UpdateA_SPEC_OBJECT_TYPE_REF(c *gin.Context) {

	mutexA_SPEC_OBJECT_TYPE_REF.Lock()
	defer mutexA_SPEC_OBJECT_TYPE_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_SPEC_OBJECT_TYPE_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_OBJECT_TYPE_REF.GetDB()

	// Validate input
	var input orm.A_SPEC_OBJECT_TYPE_REFAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_spec_object_type_refDB orm.A_SPEC_OBJECT_TYPE_REFDB

	// fetch the a_spec_object_type_ref
	_, err := db.First(&a_spec_object_type_refDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_spec_object_type_refDB.CopyBasicFieldsFromA_SPEC_OBJECT_TYPE_REF_WOP(&input.A_SPEC_OBJECT_TYPE_REF_WOP)
	a_spec_object_type_refDB.A_SPEC_OBJECT_TYPE_REFPointersEncoding = input.A_SPEC_OBJECT_TYPE_REFPointersEncoding

	db, _ = db.Model(&a_spec_object_type_refDB)
	_, err = db.Updates(&a_spec_object_type_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_object_type_refNew := new(models.A_SPEC_OBJECT_TYPE_REF)
	a_spec_object_type_refDB.CopyBasicFieldsToA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_refNew)

	// redeem pointers
	a_spec_object_type_refDB.DecodePointers(backRepo, a_spec_object_type_refNew)

	// get stage instance from DB instance, and call callback function
	a_spec_object_type_refOld := backRepo.BackRepoA_SPEC_OBJECT_TYPE_REF.Map_A_SPEC_OBJECT_TYPE_REFDBID_A_SPEC_OBJECT_TYPE_REFPtr[a_spec_object_type_refDB.ID]
	if a_spec_object_type_refOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_spec_object_type_refOld, a_spec_object_type_refNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_spec_object_type_refDB
	c.JSON(http.StatusOK, a_spec_object_type_refDB)
}

// DeleteA_SPEC_OBJECT_TYPE_REF
//
// swagger:route DELETE /a_spec_object_type_refs/{ID} a_spec_object_type_refs deleteA_SPEC_OBJECT_TYPE_REF
//
// # Delete a a_spec_object_type_ref
//
// default: genericError
//
//	200: a_spec_object_type_refDBResponse
func (controller *Controller) DeleteA_SPEC_OBJECT_TYPE_REF(c *gin.Context) {

	mutexA_SPEC_OBJECT_TYPE_REF.Lock()
	defer mutexA_SPEC_OBJECT_TYPE_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SPEC_OBJECT_TYPE_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_OBJECT_TYPE_REF.GetDB()

	// Get model if exist
	var a_spec_object_type_refDB orm.A_SPEC_OBJECT_TYPE_REFDB
	if _, err := db.First(&a_spec_object_type_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_spec_object_type_refDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_object_type_refDeleted := new(models.A_SPEC_OBJECT_TYPE_REF)
	a_spec_object_type_refDB.CopyBasicFieldsToA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_refDeleted)

	// get stage instance from DB instance, and call callback function
	a_spec_object_type_refStaged := backRepo.BackRepoA_SPEC_OBJECT_TYPE_REF.Map_A_SPEC_OBJECT_TYPE_REFDBID_A_SPEC_OBJECT_TYPE_REFPtr[a_spec_object_type_refDB.ID]
	if a_spec_object_type_refStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_spec_object_type_refStaged, a_spec_object_type_refDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
