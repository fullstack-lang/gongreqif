package exporter

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
)

func (exporter *Exporter) ExportRenderingConf(renderingConf *models.RenderingConfiguration, stager *models.Stager) {

	log.Println("Exporting the rendering configuration")

	stager.GetLoadStage().Reset()

	fileToDownload := new(load.FileToDownload).Stage(stager.GetLoadStage())

	fileToDownload.Name = strings.TrimSuffix(renderingConf.Name, ".reqif") + "-renderingConf.go"

	// 0. we create a new stage for just the marshall of the rendering configuration
	stageForRenderinfConf := models.NewStage("renderingConf")

	models.StageBranch(stageForRenderinfConf, renderingConf)

	fileName := filepath.Base(fileToDownload.Name)

	tempFile, err := os.CreateTemp("", fileName)
	if err != nil {
		log.Fatalf("Failed to create temporary file: %v", err)
	}

	// 2. Defer the removal of the file to ensure it's cleaned up.
	defer os.Remove(tempFile.Name())

	// 3. Marshall the data into the temporary file.
	stageForRenderinfConf.Marshall(tempFile, "github.com/fullstack-lang/gongreqif/go/models", "main")

	// 4. Read the content back from the file.
	// os.ReadFile needs a file path, so we use the name.
	// First, ensure the file is closed so all data is flushed to disk.
	if err := tempFile.Close(); err != nil {
		log.Fatalf("Failed to close temp file: %v", err)
	}

	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		log.Fatalf("Failed to read back temp file: %v", err)
	}

	// 5. The content is now a string in memory, and the file will be deleted.
	resultString := string(content)
	fileToDownload.Content = resultString

	stager.GetLoadStage().Commit()

	time.Sleep(1 * time.Second)
	stager.UpdateAndCommitLoadReqifStage()

	log.Println("Finished exporting the rendering configuration file", tempFile.Name())

}
