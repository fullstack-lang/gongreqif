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

	if stager.GetRootREQIF() == nil {
		log.Fatal("No REQ_IF element to marshall")
	}

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

	// add a spec object type

	stager.GetLoadStage().Reset()

	fileToDownload := new(load.FileToDownload).Stage(stager.GetLoadStage())

	fileToDownload.Name = stager.GetPathToOutputReqifFile()
	fileToDownload.Content = string(outputData)

	stager.GetLoadStage().Commit()
}
