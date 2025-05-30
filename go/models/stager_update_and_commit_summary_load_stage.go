package models

import (
	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

type FileToUploadProxy struct {
	stager *Stager
}

func (fileToUploadProxy *FileToUploadProxy) OnFileUpload(front *load.FileToUpload) {

	fileToUploadProxy.stager.processReqifData([]byte(front.Content), front.GetName())
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

	stager.loadStage.Commit()
}
