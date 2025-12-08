package models

import (
	"encoding/base64"
	"fmt"
	"os"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

type RenderingConfFileToUploadProxy struct {
	stager *Stager
}

func (proxy *RenderingConfFileToUploadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	fileName := uploadedFile.GetName()
	_ = fileName

	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}

	stageForRenderingConf := NewStage("renderingConf")
	ParseAstFromBytes(stageForRenderingConf, decodedBytes)

	stage := proxy.stager.GetStage()

	// remove existing
	for o := range *GetGongstructInstancesSetFromPointerType[*SPECIFICATION_Rendering](stage) {
		o.Unstage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*SPEC_OBJECT_TYPE_Rendering](stage) {
		o.Unstage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stage) {
		o.Unstage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_DATE_Rendering](stage) {
		o.Unstage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stage) {
		o.Unstage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_INTEGER_Rendering](stage) {
		o.Unstage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_REAL_Rendering](stage) {
		o.Unstage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_STRING_Rendering](stage) {
		o.Unstage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_XHTML_Rendering](stage) {
		o.Unstage(stage)
	}

	// stage those in storage
	for o := range *GetGongstructInstancesSetFromPointerType[*SPECIFICATION_Rendering](stageForRenderingConf) {
		o.Stage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*SPEC_OBJECT_TYPE_Rendering](stageForRenderingConf) {
		o.Stage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stageForRenderingConf) {
		o.Stage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_DATE_Rendering](stageForRenderingConf) {
		o.Stage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stageForRenderingConf) {
		o.Stage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_INTEGER_Rendering](stageForRenderingConf) {
		o.Stage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_REAL_Rendering](stageForRenderingConf) {
		o.Stage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_STRING_Rendering](stageForRenderingConf) {
		o.Stage(stage)
	}
	for o := range *GetGongstructInstancesSetFromPointerType[*ATTRIBUTE_DEFINITION_XHTML_Rendering](stageForRenderingConf) {
		o.Stage(stage)
	}

	proxy.stager.enforceRenderingConfigurationSemantic()

	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.updateAndCommitLoadRenderingConfStage()
	proxy.stager.UpdateAndCommitLoadReqifStage()

	return nil
}

func (stager *Stager) updateAndCommitLoadRenderingConfStage() {

	stager.loadRenderingConfStage.Reset()

	var fileToUpload = &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &RenderingConfFileToUploadProxy{
			stager: stager,
		},
	}

	load.StageBranch(stager.loadRenderingConfStage,
		fileToUpload,
	)

	message := &load.Message{
		Name: "Drop your rendering conf .go file here or ",
	}

	message.Stage(stager.loadRenderingConfStage)

	stager.loadRenderingConfStage.Commit()

}

func ParseAstFromBytes(stage *Stage, input []byte) error {
	// 1. Create a temporary file. The "" for dir means use the OS default.
	// The pattern "ast-*.tmp" helps in identifying the file if it's left behind.
	tempFile, err := os.CreateTemp("", "ast-*.tmp")
	if err != nil {
		return fmt.Errorf("could not create temporary file: %w", err)
	}

	// 2. IMPORTANT: Schedule the file to be removed when the function returns.
	// `defer` ensures this runs even if subsequent operations fail.
	defer os.Remove(tempFile.Name())

	fmt.Printf("Created temporary file: %s\n", tempFile.Name())

	// 3. Write the input bytes to the temporary file.
	if _, err := tempFile.Write(input); err != nil {
		return fmt.Errorf("could not write to temporary file: %w", err)
	}

	// 4. Close the file to ensure all data is flushed to disk before parsing.
	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("could not close temporary file: %w", err)
	}

	// 5. Now, call your original function with the path to our new temp file.
	return ParseAstFile(stage, tempFile.Name())
}
