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
var __RELATION_GROUP_TYPE__dummysDeclaration__ models.RELATION_GROUP_TYPE
var __RELATION_GROUP_TYPE_time__dummyDeclaration time.Duration

var mutexRELATION_GROUP_TYPE sync.Mutex

// An RELATION_GROUP_TYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRELATION_GROUP_TYPE updateRELATION_GROUP_TYPE deleteRELATION_GROUP_TYPE
type RELATION_GROUP_TYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RELATION_GROUP_TYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRELATION_GROUP_TYPE updateRELATION_GROUP_TYPE
type RELATION_GROUP_TYPEInput struct {
	// The RELATION_GROUP_TYPE to submit or modify
	// in: body
	RELATION_GROUP_TYPE *orm.RELATION_GROUP_TYPEAPI
}

// GetRELATION_GROUP_TYPEs
//
// swagger:route GET /relation_group_types relation_group_types getRELATION_GROUP_TYPEs
//
// # Get all relation_group_types
//
// Responses:
// default: genericError
//
//	200: relation_group_typeDBResponse
func (controller *Controller) GetRELATION_GROUP_TYPEs(c *gin.Context) {

	// source slice
	var relation_group_typeDBs []orm.RELATION_GROUP_TYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRELATION_GROUP_TYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATION_GROUP_TYPE.GetDB()

	_, err := db.Find(&relation_group_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	relation_group_typeAPIs := make([]orm.RELATION_GROUP_TYPEAPI, 0)

	// for each relation_group_type, update fields from the database nullable fields
	for idx := range relation_group_typeDBs {
		relation_group_typeDB := &relation_group_typeDBs[idx]
		_ = relation_group_typeDB
		var relation_group_typeAPI orm.RELATION_GROUP_TYPEAPI

		// insertion point for updating fields
		relation_group_typeAPI.ID = relation_group_typeDB.ID
		relation_group_typeDB.CopyBasicFieldsToRELATION_GROUP_TYPE_WOP(&relation_group_typeAPI.RELATION_GROUP_TYPE_WOP)
		relation_group_typeAPI.RELATION_GROUP_TYPEPointersEncoding = relation_group_typeDB.RELATION_GROUP_TYPEPointersEncoding
		relation_group_typeAPIs = append(relation_group_typeAPIs, relation_group_typeAPI)
	}

	c.JSON(http.StatusOK, relation_group_typeAPIs)
}

// PostRELATION_GROUP_TYPE
//
// swagger:route POST /relation_group_types relation_group_types postRELATION_GROUP_TYPE
//
// Creates a relation_group_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRELATION_GROUP_TYPE(c *gin.Context) {

	mutexRELATION_GROUP_TYPE.Lock()
	defer mutexRELATION_GROUP_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRELATION_GROUP_TYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATION_GROUP_TYPE.GetDB()

	// Validate input
	var input orm.RELATION_GROUP_TYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create relation_group_type
	relation_group_typeDB := orm.RELATION_GROUP_TYPEDB{}
	relation_group_typeDB.RELATION_GROUP_TYPEPointersEncoding = input.RELATION_GROUP_TYPEPointersEncoding
	relation_group_typeDB.CopyBasicFieldsFromRELATION_GROUP_TYPE_WOP(&input.RELATION_GROUP_TYPE_WOP)

	_, err = db.Create(&relation_group_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRELATION_GROUP_TYPE.CheckoutPhaseOneInstance(&relation_group_typeDB)
	relation_group_type := backRepo.BackRepoRELATION_GROUP_TYPE.Map_RELATION_GROUP_TYPEDBID_RELATION_GROUP_TYPEPtr[relation_group_typeDB.ID]

	if relation_group_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), relation_group_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, relation_group_typeDB)
}

// GetRELATION_GROUP_TYPE
//
// swagger:route GET /relation_group_types/{ID} relation_group_types getRELATION_GROUP_TYPE
//
// Gets the details for a relation_group_type.
//
// Responses:
// default: genericError
//
//	200: relation_group_typeDBResponse
func (controller *Controller) GetRELATION_GROUP_TYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRELATION_GROUP_TYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATION_GROUP_TYPE.GetDB()

	// Get relation_group_typeDB in DB
	var relation_group_typeDB orm.RELATION_GROUP_TYPEDB
	if _, err := db.First(&relation_group_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var relation_group_typeAPI orm.RELATION_GROUP_TYPEAPI
	relation_group_typeAPI.ID = relation_group_typeDB.ID
	relation_group_typeAPI.RELATION_GROUP_TYPEPointersEncoding = relation_group_typeDB.RELATION_GROUP_TYPEPointersEncoding
	relation_group_typeDB.CopyBasicFieldsToRELATION_GROUP_TYPE_WOP(&relation_group_typeAPI.RELATION_GROUP_TYPE_WOP)

	c.JSON(http.StatusOK, relation_group_typeAPI)
}

// UpdateRELATION_GROUP_TYPE
//
// swagger:route PATCH /relation_group_types/{ID} relation_group_types updateRELATION_GROUP_TYPE
//
// # Update a relation_group_type
//
// Responses:
// default: genericError
//
//	200: relation_group_typeDBResponse
func (controller *Controller) UpdateRELATION_GROUP_TYPE(c *gin.Context) {

	mutexRELATION_GROUP_TYPE.Lock()
	defer mutexRELATION_GROUP_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRELATION_GROUP_TYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATION_GROUP_TYPE.GetDB()

	// Validate input
	var input orm.RELATION_GROUP_TYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var relation_group_typeDB orm.RELATION_GROUP_TYPEDB

	// fetch the relation_group_type
	_, err := db.First(&relation_group_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	relation_group_typeDB.CopyBasicFieldsFromRELATION_GROUP_TYPE_WOP(&input.RELATION_GROUP_TYPE_WOP)
	relation_group_typeDB.RELATION_GROUP_TYPEPointersEncoding = input.RELATION_GROUP_TYPEPointersEncoding

	db, _ = db.Model(&relation_group_typeDB)
	_, err = db.Updates(&relation_group_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	relation_group_typeNew := new(models.RELATION_GROUP_TYPE)
	relation_group_typeDB.CopyBasicFieldsToRELATION_GROUP_TYPE(relation_group_typeNew)

	// redeem pointers
	relation_group_typeDB.DecodePointers(backRepo, relation_group_typeNew)

	// get stage instance from DB instance, and call callback function
	relation_group_typeOld := backRepo.BackRepoRELATION_GROUP_TYPE.Map_RELATION_GROUP_TYPEDBID_RELATION_GROUP_TYPEPtr[relation_group_typeDB.ID]
	if relation_group_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), relation_group_typeOld, relation_group_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the relation_group_typeDB
	c.JSON(http.StatusOK, relation_group_typeDB)
}

// DeleteRELATION_GROUP_TYPE
//
// swagger:route DELETE /relation_group_types/{ID} relation_group_types deleteRELATION_GROUP_TYPE
//
// # Delete a relation_group_type
//
// default: genericError
//
//	200: relation_group_typeDBResponse
func (controller *Controller) DeleteRELATION_GROUP_TYPE(c *gin.Context) {

	mutexRELATION_GROUP_TYPE.Lock()
	defer mutexRELATION_GROUP_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRELATION_GROUP_TYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATION_GROUP_TYPE.GetDB()

	// Get model if exist
	var relation_group_typeDB orm.RELATION_GROUP_TYPEDB
	if _, err := db.First(&relation_group_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&relation_group_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	relation_group_typeDeleted := new(models.RELATION_GROUP_TYPE)
	relation_group_typeDB.CopyBasicFieldsToRELATION_GROUP_TYPE(relation_group_typeDeleted)

	// get stage instance from DB instance, and call callback function
	relation_group_typeStaged := backRepo.BackRepoRELATION_GROUP_TYPE.Map_RELATION_GROUP_TYPEDBID_RELATION_GROUP_TYPEPtr[relation_group_typeDB.ID]
	if relation_group_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), relation_group_typeStaged, relation_group_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
