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
var __ALTERNATIVE_ID__dummysDeclaration__ models.ALTERNATIVE_ID
var __ALTERNATIVE_ID_time__dummyDeclaration time.Duration

var mutexALTERNATIVE_ID sync.Mutex

// An ALTERNATIVE_IDID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getALTERNATIVE_ID updateALTERNATIVE_ID deleteALTERNATIVE_ID
type ALTERNATIVE_IDID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ALTERNATIVE_IDInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postALTERNATIVE_ID updateALTERNATIVE_ID
type ALTERNATIVE_IDInput struct {
	// The ALTERNATIVE_ID to submit or modify
	// in: body
	ALTERNATIVE_ID *orm.ALTERNATIVE_IDAPI
}

// GetALTERNATIVE_IDs
//
// swagger:route GET /alternative_ids alternative_ids getALTERNATIVE_IDs
//
// # Get all alternative_ids
//
// Responses:
// default: genericError
//
//	200: alternative_idDBResponse
func (controller *Controller) GetALTERNATIVE_IDs(c *gin.Context) {

	// source slice
	var alternative_idDBs []orm.ALTERNATIVE_IDDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetALTERNATIVE_IDs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoALTERNATIVE_ID.GetDB()

	_, err := db.Find(&alternative_idDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	alternative_idAPIs := make([]orm.ALTERNATIVE_IDAPI, 0)

	// for each alternative_id, update fields from the database nullable fields
	for idx := range alternative_idDBs {
		alternative_idDB := &alternative_idDBs[idx]
		_ = alternative_idDB
		var alternative_idAPI orm.ALTERNATIVE_IDAPI

		// insertion point for updating fields
		alternative_idAPI.ID = alternative_idDB.ID
		alternative_idDB.CopyBasicFieldsToALTERNATIVE_ID_WOP(&alternative_idAPI.ALTERNATIVE_ID_WOP)
		alternative_idAPI.ALTERNATIVE_IDPointersEncoding = alternative_idDB.ALTERNATIVE_IDPointersEncoding
		alternative_idAPIs = append(alternative_idAPIs, alternative_idAPI)
	}

	c.JSON(http.StatusOK, alternative_idAPIs)
}

// PostALTERNATIVE_ID
//
// swagger:route POST /alternative_ids alternative_ids postALTERNATIVE_ID
//
// Creates a alternative_id
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostALTERNATIVE_ID(c *gin.Context) {

	mutexALTERNATIVE_ID.Lock()
	defer mutexALTERNATIVE_ID.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostALTERNATIVE_IDs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoALTERNATIVE_ID.GetDB()

	// Validate input
	var input orm.ALTERNATIVE_IDAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create alternative_id
	alternative_idDB := orm.ALTERNATIVE_IDDB{}
	alternative_idDB.ALTERNATIVE_IDPointersEncoding = input.ALTERNATIVE_IDPointersEncoding
	alternative_idDB.CopyBasicFieldsFromALTERNATIVE_ID_WOP(&input.ALTERNATIVE_ID_WOP)

	_, err = db.Create(&alternative_idDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoALTERNATIVE_ID.CheckoutPhaseOneInstance(&alternative_idDB)
	alternative_id := backRepo.BackRepoALTERNATIVE_ID.Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDPtr[alternative_idDB.ID]

	if alternative_id != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), alternative_id)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, alternative_idDB)
}

// GetALTERNATIVE_ID
//
// swagger:route GET /alternative_ids/{ID} alternative_ids getALTERNATIVE_ID
//
// Gets the details for a alternative_id.
//
// Responses:
// default: genericError
//
//	200: alternative_idDBResponse
func (controller *Controller) GetALTERNATIVE_ID(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetALTERNATIVE_ID", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoALTERNATIVE_ID.GetDB()

	// Get alternative_idDB in DB
	var alternative_idDB orm.ALTERNATIVE_IDDB
	if _, err := db.First(&alternative_idDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var alternative_idAPI orm.ALTERNATIVE_IDAPI
	alternative_idAPI.ID = alternative_idDB.ID
	alternative_idAPI.ALTERNATIVE_IDPointersEncoding = alternative_idDB.ALTERNATIVE_IDPointersEncoding
	alternative_idDB.CopyBasicFieldsToALTERNATIVE_ID_WOP(&alternative_idAPI.ALTERNATIVE_ID_WOP)

	c.JSON(http.StatusOK, alternative_idAPI)
}

// UpdateALTERNATIVE_ID
//
// swagger:route PATCH /alternative_ids/{ID} alternative_ids updateALTERNATIVE_ID
//
// # Update a alternative_id
//
// Responses:
// default: genericError
//
//	200: alternative_idDBResponse
func (controller *Controller) UpdateALTERNATIVE_ID(c *gin.Context) {

	mutexALTERNATIVE_ID.Lock()
	defer mutexALTERNATIVE_ID.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateALTERNATIVE_ID", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoALTERNATIVE_ID.GetDB()

	// Validate input
	var input orm.ALTERNATIVE_IDAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var alternative_idDB orm.ALTERNATIVE_IDDB

	// fetch the alternative_id
	_, err := db.First(&alternative_idDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	alternative_idDB.CopyBasicFieldsFromALTERNATIVE_ID_WOP(&input.ALTERNATIVE_ID_WOP)
	alternative_idDB.ALTERNATIVE_IDPointersEncoding = input.ALTERNATIVE_IDPointersEncoding

	db, _ = db.Model(&alternative_idDB)
	_, err = db.Updates(&alternative_idDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	alternative_idNew := new(models.ALTERNATIVE_ID)
	alternative_idDB.CopyBasicFieldsToALTERNATIVE_ID(alternative_idNew)

	// redeem pointers
	alternative_idDB.DecodePointers(backRepo, alternative_idNew)

	// get stage instance from DB instance, and call callback function
	alternative_idOld := backRepo.BackRepoALTERNATIVE_ID.Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDPtr[alternative_idDB.ID]
	if alternative_idOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), alternative_idOld, alternative_idNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the alternative_idDB
	c.JSON(http.StatusOK, alternative_idDB)
}

// DeleteALTERNATIVE_ID
//
// swagger:route DELETE /alternative_ids/{ID} alternative_ids deleteALTERNATIVE_ID
//
// # Delete a alternative_id
//
// default: genericError
//
//	200: alternative_idDBResponse
func (controller *Controller) DeleteALTERNATIVE_ID(c *gin.Context) {

	mutexALTERNATIVE_ID.Lock()
	defer mutexALTERNATIVE_ID.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteALTERNATIVE_ID", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoALTERNATIVE_ID.GetDB()

	// Get model if exist
	var alternative_idDB orm.ALTERNATIVE_IDDB
	if _, err := db.First(&alternative_idDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&alternative_idDB)

	// get an instance (not staged) from DB instance, and call callback function
	alternative_idDeleted := new(models.ALTERNATIVE_ID)
	alternative_idDB.CopyBasicFieldsToALTERNATIVE_ID(alternative_idDeleted)

	// get stage instance from DB instance, and call callback function
	alternative_idStaged := backRepo.BackRepoALTERNATIVE_ID.Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDPtr[alternative_idDB.ID]
	if alternative_idStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), alternative_idStaged, alternative_idDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
