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
var __A_ALTERNATIVE_ID__dummysDeclaration__ models.A_ALTERNATIVE_ID
var __A_ALTERNATIVE_ID_time__dummyDeclaration time.Duration

var mutexA_ALTERNATIVE_ID sync.Mutex

// An A_ALTERNATIVE_IDID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ALTERNATIVE_ID updateA_ALTERNATIVE_ID deleteA_ALTERNATIVE_ID
type A_ALTERNATIVE_IDID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ALTERNATIVE_IDInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ALTERNATIVE_ID updateA_ALTERNATIVE_ID
type A_ALTERNATIVE_IDInput struct {
	// The A_ALTERNATIVE_ID to submit or modify
	// in: body
	A_ALTERNATIVE_ID *orm.A_ALTERNATIVE_IDAPI
}

// GetA_ALTERNATIVE_IDs
//
// swagger:route GET /a_alternative_ids a_alternative_ids getA_ALTERNATIVE_IDs
//
// # Get all a_alternative_ids
//
// Responses:
// default: genericError
//
//	200: a_alternative_idDBResponse
func (controller *Controller) GetA_ALTERNATIVE_IDs(c *gin.Context) {

	// source slice
	var a_alternative_idDBs []orm.A_ALTERNATIVE_IDDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ALTERNATIVE_IDs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ALTERNATIVE_ID.GetDB()

	_, err := db.Find(&a_alternative_idDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_alternative_idAPIs := make([]orm.A_ALTERNATIVE_IDAPI, 0)

	// for each a_alternative_id, update fields from the database nullable fields
	for idx := range a_alternative_idDBs {
		a_alternative_idDB := &a_alternative_idDBs[idx]
		_ = a_alternative_idDB
		var a_alternative_idAPI orm.A_ALTERNATIVE_IDAPI

		// insertion point for updating fields
		a_alternative_idAPI.ID = a_alternative_idDB.ID
		a_alternative_idDB.CopyBasicFieldsToA_ALTERNATIVE_ID_WOP(&a_alternative_idAPI.A_ALTERNATIVE_ID_WOP)
		a_alternative_idAPI.A_ALTERNATIVE_IDPointersEncoding = a_alternative_idDB.A_ALTERNATIVE_IDPointersEncoding
		a_alternative_idAPIs = append(a_alternative_idAPIs, a_alternative_idAPI)
	}

	c.JSON(http.StatusOK, a_alternative_idAPIs)
}

// PostA_ALTERNATIVE_ID
//
// swagger:route POST /a_alternative_ids a_alternative_ids postA_ALTERNATIVE_ID
//
// Creates a a_alternative_id
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ALTERNATIVE_ID(c *gin.Context) {

	mutexA_ALTERNATIVE_ID.Lock()
	defer mutexA_ALTERNATIVE_ID.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ALTERNATIVE_IDs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ALTERNATIVE_ID.GetDB()

	// Validate input
	var input orm.A_ALTERNATIVE_IDAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_alternative_id
	a_alternative_idDB := orm.A_ALTERNATIVE_IDDB{}
	a_alternative_idDB.A_ALTERNATIVE_IDPointersEncoding = input.A_ALTERNATIVE_IDPointersEncoding
	a_alternative_idDB.CopyBasicFieldsFromA_ALTERNATIVE_ID_WOP(&input.A_ALTERNATIVE_ID_WOP)

	_, err = db.Create(&a_alternative_idDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ALTERNATIVE_ID.CheckoutPhaseOneInstance(&a_alternative_idDB)
	a_alternative_id := backRepo.BackRepoA_ALTERNATIVE_ID.Map_A_ALTERNATIVE_IDDBID_A_ALTERNATIVE_IDPtr[a_alternative_idDB.ID]

	if a_alternative_id != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_alternative_id)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_alternative_idDB)
}

// GetA_ALTERNATIVE_ID
//
// swagger:route GET /a_alternative_ids/{ID} a_alternative_ids getA_ALTERNATIVE_ID
//
// Gets the details for a a_alternative_id.
//
// Responses:
// default: genericError
//
//	200: a_alternative_idDBResponse
func (controller *Controller) GetA_ALTERNATIVE_ID(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ALTERNATIVE_ID", "Name", stackPath)
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
	db := backRepo.BackRepoA_ALTERNATIVE_ID.GetDB()

	// Get a_alternative_idDB in DB
	var a_alternative_idDB orm.A_ALTERNATIVE_IDDB
	if _, err := db.First(&a_alternative_idDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_alternative_idAPI orm.A_ALTERNATIVE_IDAPI
	a_alternative_idAPI.ID = a_alternative_idDB.ID
	a_alternative_idAPI.A_ALTERNATIVE_IDPointersEncoding = a_alternative_idDB.A_ALTERNATIVE_IDPointersEncoding
	a_alternative_idDB.CopyBasicFieldsToA_ALTERNATIVE_ID_WOP(&a_alternative_idAPI.A_ALTERNATIVE_ID_WOP)

	c.JSON(http.StatusOK, a_alternative_idAPI)
}

// UpdateA_ALTERNATIVE_ID
//
// swagger:route PATCH /a_alternative_ids/{ID} a_alternative_ids updateA_ALTERNATIVE_ID
//
// # Update a a_alternative_id
//
// Responses:
// default: genericError
//
//	200: a_alternative_idDBResponse
func (controller *Controller) UpdateA_ALTERNATIVE_ID(c *gin.Context) {

	mutexA_ALTERNATIVE_ID.Lock()
	defer mutexA_ALTERNATIVE_ID.Unlock()

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
	db := backRepo.BackRepoA_ALTERNATIVE_ID.GetDB()

	// Validate input
	var input orm.A_ALTERNATIVE_IDAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_alternative_idDB orm.A_ALTERNATIVE_IDDB

	// fetch the a_alternative_id
	_, err := db.First(&a_alternative_idDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_alternative_idDB.CopyBasicFieldsFromA_ALTERNATIVE_ID_WOP(&input.A_ALTERNATIVE_ID_WOP)
	a_alternative_idDB.A_ALTERNATIVE_IDPointersEncoding = input.A_ALTERNATIVE_IDPointersEncoding

	db, _ = db.Model(&a_alternative_idDB)
	_, err = db.Updates(&a_alternative_idDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_alternative_idNew := new(models.A_ALTERNATIVE_ID)
	a_alternative_idDB.CopyBasicFieldsToA_ALTERNATIVE_ID(a_alternative_idNew)

	// redeem pointers
	a_alternative_idDB.DecodePointers(backRepo, a_alternative_idNew)

	// get stage instance from DB instance, and call callback function
	a_alternative_idOld := backRepo.BackRepoA_ALTERNATIVE_ID.Map_A_ALTERNATIVE_IDDBID_A_ALTERNATIVE_IDPtr[a_alternative_idDB.ID]
	if a_alternative_idOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), a_alternative_idOld, a_alternative_idNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_alternative_idDB
	c.JSON(http.StatusOK, a_alternative_idDB)
}

// DeleteA_ALTERNATIVE_ID
//
// swagger:route DELETE /a_alternative_ids/{ID} a_alternative_ids deleteA_ALTERNATIVE_ID
//
// # Delete a a_alternative_id
//
// default: genericError
//
//	200: a_alternative_idDBResponse
func (controller *Controller) DeleteA_ALTERNATIVE_ID(c *gin.Context) {

	mutexA_ALTERNATIVE_ID.Lock()
	defer mutexA_ALTERNATIVE_ID.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ALTERNATIVE_ID", "Name", stackPath)
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
	db := backRepo.BackRepoA_ALTERNATIVE_ID.GetDB()

	// Get model if exist
	var a_alternative_idDB orm.A_ALTERNATIVE_IDDB
	if _, err := db.First(&a_alternative_idDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_alternative_idDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_alternative_idDeleted := new(models.A_ALTERNATIVE_ID)
	a_alternative_idDB.CopyBasicFieldsToA_ALTERNATIVE_ID(a_alternative_idDeleted)

	// get stage instance from DB instance, and call callback function
	a_alternative_idStaged := backRepo.BackRepoA_ALTERNATIVE_ID.Map_A_ALTERNATIVE_IDDBID_A_ALTERNATIVE_IDPtr[a_alternative_idDB.ID]
	if a_alternative_idStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_alternative_idStaged, a_alternative_idDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
