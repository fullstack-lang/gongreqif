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
var __REQ_IF_HEADER__dummysDeclaration__ models.REQ_IF_HEADER
var __REQ_IF_HEADER_time__dummyDeclaration time.Duration

var mutexREQ_IF_HEADER sync.Mutex

// An REQ_IF_HEADERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getREQ_IF_HEADER updateREQ_IF_HEADER deleteREQ_IF_HEADER
type REQ_IF_HEADERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// REQ_IF_HEADERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postREQ_IF_HEADER updateREQ_IF_HEADER
type REQ_IF_HEADERInput struct {
	// The REQ_IF_HEADER to submit or modify
	// in: body
	REQ_IF_HEADER *orm.REQ_IF_HEADERAPI
}

// GetREQ_IF_HEADERs
//
// swagger:route GET /req_if_headers req_if_headers getREQ_IF_HEADERs
//
// # Get all req_if_headers
//
// Responses:
// default: genericError
//
//	200: req_if_headerDBResponse
func (controller *Controller) GetREQ_IF_HEADERs(c *gin.Context) {

	// source slice
	var req_if_headerDBs []orm.REQ_IF_HEADERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQ_IF_HEADERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQ_IF_HEADER.GetDB()

	query := db.Find(&req_if_headerDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	req_if_headerAPIs := make([]orm.REQ_IF_HEADERAPI, 0)

	// for each req_if_header, update fields from the database nullable fields
	for idx := range req_if_headerDBs {
		req_if_headerDB := &req_if_headerDBs[idx]
		_ = req_if_headerDB
		var req_if_headerAPI orm.REQ_IF_HEADERAPI

		// insertion point for updating fields
		req_if_headerAPI.ID = req_if_headerDB.ID
		req_if_headerDB.CopyBasicFieldsToREQ_IF_HEADER_WOP(&req_if_headerAPI.REQ_IF_HEADER_WOP)
		req_if_headerAPI.REQ_IF_HEADERPointersEncoding = req_if_headerDB.REQ_IF_HEADERPointersEncoding
		req_if_headerAPIs = append(req_if_headerAPIs, req_if_headerAPI)
	}

	c.JSON(http.StatusOK, req_if_headerAPIs)
}

// PostREQ_IF_HEADER
//
// swagger:route POST /req_if_headers req_if_headers postREQ_IF_HEADER
//
// Creates a req_if_header
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostREQ_IF_HEADER(c *gin.Context) {

	mutexREQ_IF_HEADER.Lock()
	defer mutexREQ_IF_HEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostREQ_IF_HEADERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQ_IF_HEADER.GetDB()

	// Validate input
	var input orm.REQ_IF_HEADERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create req_if_header
	req_if_headerDB := orm.REQ_IF_HEADERDB{}
	req_if_headerDB.REQ_IF_HEADERPointersEncoding = input.REQ_IF_HEADERPointersEncoding
	req_if_headerDB.CopyBasicFieldsFromREQ_IF_HEADER_WOP(&input.REQ_IF_HEADER_WOP)

	query := db.Create(&req_if_headerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoREQ_IF_HEADER.CheckoutPhaseOneInstance(&req_if_headerDB)
	req_if_header := backRepo.BackRepoREQ_IF_HEADER.Map_REQ_IF_HEADERDBID_REQ_IF_HEADERPtr[req_if_headerDB.ID]

	if req_if_header != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), req_if_header)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, req_if_headerDB)
}

// GetREQ_IF_HEADER
//
// swagger:route GET /req_if_headers/{ID} req_if_headers getREQ_IF_HEADER
//
// Gets the details for a req_if_header.
//
// Responses:
// default: genericError
//
//	200: req_if_headerDBResponse
func (controller *Controller) GetREQ_IF_HEADER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQ_IF_HEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQ_IF_HEADER.GetDB()

	// Get req_if_headerDB in DB
	var req_if_headerDB orm.REQ_IF_HEADERDB
	if err := db.First(&req_if_headerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var req_if_headerAPI orm.REQ_IF_HEADERAPI
	req_if_headerAPI.ID = req_if_headerDB.ID
	req_if_headerAPI.REQ_IF_HEADERPointersEncoding = req_if_headerDB.REQ_IF_HEADERPointersEncoding
	req_if_headerDB.CopyBasicFieldsToREQ_IF_HEADER_WOP(&req_if_headerAPI.REQ_IF_HEADER_WOP)

	c.JSON(http.StatusOK, req_if_headerAPI)
}

// UpdateREQ_IF_HEADER
//
// swagger:route PATCH /req_if_headers/{ID} req_if_headers updateREQ_IF_HEADER
//
// # Update a req_if_header
//
// Responses:
// default: genericError
//
//	200: req_if_headerDBResponse
func (controller *Controller) UpdateREQ_IF_HEADER(c *gin.Context) {

	mutexREQ_IF_HEADER.Lock()
	defer mutexREQ_IF_HEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateREQ_IF_HEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQ_IF_HEADER.GetDB()

	// Validate input
	var input orm.REQ_IF_HEADERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var req_if_headerDB orm.REQ_IF_HEADERDB

	// fetch the req_if_header
	query := db.First(&req_if_headerDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	req_if_headerDB.CopyBasicFieldsFromREQ_IF_HEADER_WOP(&input.REQ_IF_HEADER_WOP)
	req_if_headerDB.REQ_IF_HEADERPointersEncoding = input.REQ_IF_HEADERPointersEncoding

	query = db.Model(&req_if_headerDB).Updates(req_if_headerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	req_if_headerNew := new(models.REQ_IF_HEADER)
	req_if_headerDB.CopyBasicFieldsToREQ_IF_HEADER(req_if_headerNew)

	// redeem pointers
	req_if_headerDB.DecodePointers(backRepo, req_if_headerNew)

	// get stage instance from DB instance, and call callback function
	req_if_headerOld := backRepo.BackRepoREQ_IF_HEADER.Map_REQ_IF_HEADERDBID_REQ_IF_HEADERPtr[req_if_headerDB.ID]
	if req_if_headerOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), req_if_headerOld, req_if_headerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the req_if_headerDB
	c.JSON(http.StatusOK, req_if_headerDB)
}

// DeleteREQ_IF_HEADER
//
// swagger:route DELETE /req_if_headers/{ID} req_if_headers deleteREQ_IF_HEADER
//
// # Delete a req_if_header
//
// default: genericError
//
//	200: req_if_headerDBResponse
func (controller *Controller) DeleteREQ_IF_HEADER(c *gin.Context) {

	mutexREQ_IF_HEADER.Lock()
	defer mutexREQ_IF_HEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteREQ_IF_HEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQ_IF_HEADER.GetDB()

	// Get model if exist
	var req_if_headerDB orm.REQ_IF_HEADERDB
	if err := db.First(&req_if_headerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&req_if_headerDB)

	// get an instance (not staged) from DB instance, and call callback function
	req_if_headerDeleted := new(models.REQ_IF_HEADER)
	req_if_headerDB.CopyBasicFieldsToREQ_IF_HEADER(req_if_headerDeleted)

	// get stage instance from DB instance, and call callback function
	req_if_headerStaged := backRepo.BackRepoREQ_IF_HEADER.Map_REQ_IF_HEADERDBID_REQ_IF_HEADERPtr[req_if_headerDB.ID]
	if req_if_headerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), req_if_headerStaged, req_if_headerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
