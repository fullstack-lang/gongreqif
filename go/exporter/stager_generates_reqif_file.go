package exporter

import (
	"encoding/xml"
	"fmt"
	"log"

	load "github.com/fullstack-lang/gong/lib/load/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
)

type ReqifExporter struct {
}

func (reqifExport *ReqifExporter) ExportReqif(stager *models.Stager) {

	log.Println("Exporting the ReqIF file")

	if stager.GetRootREQIF() == nil {
		log.Fatal("No REQ_IF element to marshall")
	}

	// // Get the path to the system's temporary directory.
	// // os.TempDir() provides a cross-platform way to get this path.
	// // For example:
	// // - Linux/macOS: /tmp
	// // - Windows: C:\Users\YourUser\AppData\Local\Temp
	// tempDir := os.TempDir()
	// filePath := filepath.Join(tempDir, stager.pathToOutputReqifFile)

	// // Write the final byte slice to the specified file.
	// // os.WriteFile handles creating/truncating the file and writing the data.
	// // 0644 is a standard file permission.
	// err = os.WriteFile(filePath, outputData, 0644)
	// if err != nil {
	// 	fmt.Printf("Error writing to file %s: %v\n", filePath, err)
	// 	return
	// }

	// fmt.Printf("Successfully wrote REQ_IF data to %s\n", filePath)

	// get the root of the SPEC_TYPE slice
	var specTypes *models.A_SPEC_TYPES

	specTypesInstances := models.GetGongstrucsSorted[*models.A_SPEC_TYPES](stager.GetStage())

	if len(specTypesInstances) != 1 {
		log.Println("Problem, there should be one models.A_SPEC_TYPES instance")
		return
	}

	specTypes = specTypesInstances[0]
	_ = specTypes

	if len(specTypes.SPEC_OBJECT_TYPE) != 1 {
		log.Println("Problem, there should be one SPEC_OBJECT_TYPE in DOORS legacy")
		return
	}
	legacyDoorsObjectType := specTypes.SPEC_OBJECT_TYPE[0]

	// add a spec object type
	newSpecObjectType := new(models.SPEC_OBJECT_TYPE)
	specTypes.SPEC_OBJECT_TYPE = append(specTypes.SPEC_OBJECT_TYPE, newSpecObjectType)
	newSpecObjectType.Name = legacyDoorsObjectType.Name + "-Text"
	newSpecObjectType.IDENTIFIER = legacyDoorsObjectType.IDENTIFIER + "-Text"
	newSpecObjectType.SPEC_ATTRIBUTES = legacyDoorsObjectType.SPEC_ATTRIBUTES

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

	fileToDownload.Name = stager.GetPathToOutputReqifFile()
	fileToDownload.Content = string(outputData)

	stager.GetLoadStage().Commit()

	log.Println("Finished exporting the ReqIF file", stager.GetPathToOutputReqifFile())

}
