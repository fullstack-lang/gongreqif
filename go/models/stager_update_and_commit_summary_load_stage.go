package models

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

type FileToUploadProxy struct {
	stager *Stager
}

func (fileToUploadProxy *FileToUploadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	fileName := uploadedFile.GetName()
	fileExt := strings.ToLower(filepath.Ext(fileName))

	var contentToProcess []byte
	var err error

	// Direct processing for .reqif files
	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}

	switch fileExt {
	case ".reqif":

		contentToProcess = decodedBytes

	case ".reqifz":
		// Unzip and extract .reqif content for .reqifz files
		contentToProcess, err = extractReqifFromZip(decodedBytes)
		if err != nil {
			return fmt.Errorf("failed to extract REQIF from zip: %w", err)
		}

	default:
		return fmt.Errorf("unsupported file extension: %s", fileExt)
	}

	// Process the content (either direct .reqif or extracted from .reqifz)
	fileToUploadProxy.stager.processReqifData(contentToProcess, fileName)
	return nil
}

// extractReqifFromZip extracts the first .reqif file found in the zip archive
func extractReqifFromZip(zipData []byte) ([]byte, error) {
	// Create a reader from the zip data
	zipReader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil, fmt.Errorf("failed to create zip reader: %w", err)
	}

	// Look for .reqif files in the archive
	for _, file := range zipReader.File {
		if strings.ToLower(filepath.Ext(file.Name)) == ".reqif" {
			// Open the file
			rc, err := file.Open()
			if err != nil {
				return nil, fmt.Errorf("failed to open file %s in zip: %w", file.Name, err)
			}
			defer rc.Close()

			// Read the content
			content, err := io.ReadAll(rc)
			if err != nil {
				return nil, fmt.Errorf("failed to read file %s from zip: %w", file.Name, err)
			}

			return content, nil
		}
	}

	return nil, fmt.Errorf("no .reqif file found in the zip archive")
}
func (stager *Stager) updateAndCommitSummaryLoadStage() {

	stager.loadStage.Reset()

	var fileToUpload = &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &FileToUploadProxy{
			stager: stager,
		},
	}

	load.StageBranch(stager.loadStage,
		fileToUpload,
	)

	message := &load.Message{
		Name: "Drop your .reqif or .reqifz file here or ",
	}

	message.Stage(stager.loadStage)

	stager.loadStage.Commit()
}
