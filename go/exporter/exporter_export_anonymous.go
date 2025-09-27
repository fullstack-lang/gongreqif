package exporter

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
)

func (exporter *Exporter) ExportAnonymousReqif(stager *models.Stager) {

	log.Println("Exporting the ReqIF file")

	var rootReqif *models.REQ_IF
	if rootReqif = stager.GetRootREQIF(); rootReqif == nil {
		log.Fatal("No REQ_IF element to marshall")
	}
	_ = rootReqif

	for _, specObject := range rootReqif.CORE_CONTENT.REQ_IF_CONTENT.SPEC_OBJECTS.SPEC_OBJECT {
		_ = specObject

	}

	// parse all spec objects and if the spec object has a chapter name field, then
	// change the type of the object

	// Marshal the struct into an indented XML byte slice.
	// Using MarshalIndent makes the output file human-readable.
	// The "" means no line prefix, and "  " means use two spaces for indentation.
	xmlData, err := xml.MarshalIndent(stager.GetRootREQIF(), "", "  ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}

	// Prepend the standard XML header to the marshalled data.
	// This makes it a valid XML file.
	outputData := []byte(xml.Header + string(xmlData))

	stager.GetLoadStage().Reset()

	fileToDownload := new(load.FileToDownload).Stage(stager.GetLoadStage())

	filename := stager.GetPathToReqifFile()

	if strings.HasSuffix(filename, ".reqifz") {
		filename = strings.TrimSuffix(filename, "z")
	}

	fileToDownload.Name = filename
	fileToDownload.Content = string(outputData)

	stager.GetLoadStage().Commit()

	time.Sleep(1 * time.Second)
	stager.UpdateAndCommitLoadReqifStage()

	log.Println("Finished exporting the anonymous ReqIF file", filename)

}
