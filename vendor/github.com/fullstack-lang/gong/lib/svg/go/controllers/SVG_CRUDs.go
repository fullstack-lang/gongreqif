// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __SVG__dummysDeclaration__ models.SVG
var __SVG_time__dummyDeclaration time.Duration

var mutexSVG sync.Mutex

// An SVGID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSVG updateSVG deleteSVG
type SVGID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SVGInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSVG updateSVG
type SVGInput struct {
	// The SVG to submit or modify
	// in: body
	SVG *orm.SVGAPI
}

// GetSVGs
//
// swagger:route GET /svgs svgs getSVGs
//
// # Get all svgs
//
// Responses:
// default: genericError
//
//	200: svgDBResponse
func (controller *Controller) GetSVGs(c *gin.Context) {

	// source slice
	var svgDBs []orm.SVGDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSVGs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSVG.GetDB()

	_, err := db.Find(&svgDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	svgAPIs := make([]orm.SVGAPI, 0)

	// for each svg, update fields from the database nullable fields
	for idx := range svgDBs {
		svgDB := &svgDBs[idx]
		_ = svgDB
		var svgAPI orm.SVGAPI

		// insertion point for updating fields
		svgAPI.ID = svgDB.ID
		svgDB.CopyBasicFieldsToSVG_WOP(&svgAPI.SVG_WOP)
		svgAPI.SVGPointersEncoding = svgDB.SVGPointersEncoding
		svgAPIs = append(svgAPIs, svgAPI)
	}

	c.JSON(http.StatusOK, svgAPIs)
}

// PostSVG
//
// swagger:route POST /svgs svgs postSVG
//
// Creates a svg
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSVG(c *gin.Context) {

	mutexSVG.Lock()
	defer mutexSVG.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSVGs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSVG.GetDB()

	// Validate input
	var input orm.SVGAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create svg
	svgDB := orm.SVGDB{}
	svgDB.SVGPointersEncoding = input.SVGPointersEncoding
	svgDB.CopyBasicFieldsFromSVG_WOP(&input.SVG_WOP)

	_, err = db.Create(&svgDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSVG.CheckoutPhaseOneInstance(&svgDB)
	svg := backRepo.BackRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID]

	if svg != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), svg)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, svgDB)
}

// GetSVG
//
// swagger:route GET /svgs/{ID} svgs getSVG
//
// Gets the details for a svg.
//
// Responses:
// default: genericError
//
//	200: svgDBResponse
func (controller *Controller) GetSVG(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSVG", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSVG.GetDB()

	// Get svgDB in DB
	var svgDB orm.SVGDB
	if _, err := db.First(&svgDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var svgAPI orm.SVGAPI
	svgAPI.ID = svgDB.ID
	svgAPI.SVGPointersEncoding = svgDB.SVGPointersEncoding
	svgDB.CopyBasicFieldsToSVG_WOP(&svgAPI.SVG_WOP)

	c.JSON(http.StatusOK, svgAPI)
}

// UpdateSVG
//
// swagger:route PATCH /svgs/{ID} svgs updateSVG
//
// # Update a svg
//
// Responses:
// default: genericError
//
//	200: svgDBResponse
func (controller *Controller) UpdateSVG(c *gin.Context) {

	mutexSVG.Lock()
	defer mutexSVG.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSVG", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSVG.GetDB()

	// Validate input
	var input orm.SVGAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var svgDB orm.SVGDB

	// fetch the svg
	_, err := db.First(&svgDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	svgDB.CopyBasicFieldsFromSVG_WOP(&input.SVG_WOP)
	svgDB.SVGPointersEncoding = input.SVGPointersEncoding

	db, _ = db.Model(&svgDB)
	_, err = db.Updates(&svgDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	svgNew := new(models.SVG)
	svgDB.CopyBasicFieldsToSVG(svgNew)

	// redeem pointers
	svgDB.DecodePointers(backRepo, svgNew)

	// get stage instance from DB instance, and call callback function
	svgOld := backRepo.BackRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID]
	if svgOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), svgOld, svgNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the svgDB
	c.JSON(http.StatusOK, svgDB)
}

// DeleteSVG
//
// swagger:route DELETE /svgs/{ID} svgs deleteSVG
//
// # Delete a svg
//
// default: genericError
//
//	200: svgDBResponse
func (controller *Controller) DeleteSVG(c *gin.Context) {

	mutexSVG.Lock()
	defer mutexSVG.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSVG", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoSVG.GetDB()

	// Get model if exist
	var svgDB orm.SVGDB
	if _, err := db.First(&svgDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&svgDB)

	// get an instance (not staged) from DB instance, and call callback function
	svgDeleted := new(models.SVG)
	svgDB.CopyBasicFieldsToSVG(svgDeleted)

	// get stage instance from DB instance, and call callback function
	svgStaged := backRepo.BackRepoSVG.Map_SVGDBID_SVGPtr[svgDB.ID]
	if svgStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), svgStaged, svgDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
