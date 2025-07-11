// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gong/lib/svg/go/db"
	"github.com/fullstack-lang/gong/lib/svg/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gong/lib/svg/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAnimate BackRepoAnimateStruct

	BackRepoCircle BackRepoCircleStruct

	BackRepoEllipse BackRepoEllipseStruct

	BackRepoLayer BackRepoLayerStruct

	BackRepoLine BackRepoLineStruct

	BackRepoLink BackRepoLinkStruct

	BackRepoLinkAnchoredText BackRepoLinkAnchoredTextStruct

	BackRepoPath BackRepoPathStruct

	BackRepoPoint BackRepoPointStruct

	BackRepoPolygone BackRepoPolygoneStruct

	BackRepoPolyline BackRepoPolylineStruct

	BackRepoRect BackRepoRectStruct

	BackRepoRectAnchoredPath BackRepoRectAnchoredPathStruct

	BackRepoRectAnchoredRect BackRepoRectAnchoredRectStruct

	BackRepoRectAnchoredText BackRepoRectAnchoredTextStruct

	BackRepoRectLinkLink BackRepoRectLinkLinkStruct

	BackRepoSVG BackRepoSVGStruct

	BackRepoSvgText BackRepoSvgTextStruct

	BackRepoText BackRepoTextStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.Stage

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex

	subscribersRwMutex sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.Stage, filename string) (backRepo *BackRepoStruct) {

	var db db.DBInterface

	db = NewDBLite()

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gong_lib_svg_go",
		&AnimateDB{},
		&CircleDB{},
		&EllipseDB{},
		&LayerDB{},
		&LineDB{},
		&LinkDB{},
		&LinkAnchoredTextDB{},
		&PathDB{},
		&PointDB{},
		&PolygoneDB{},
		&PolylineDB{},
		&RectDB{},
		&RectAnchoredPathDB{},
		&RectAnchoredRectDB{},
		&RectAnchoredTextDB{},
		&RectLinkLinkDB{},
		&SVGDB{},
		&SvgTextDB{},
		&TextDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoAnimate = BackRepoAnimateStruct{
		Map_AnimateDBID_AnimatePtr: make(map[uint]*models.Animate, 0),
		Map_AnimateDBID_AnimateDB:  make(map[uint]*AnimateDB, 0),
		Map_AnimatePtr_AnimateDBID: make(map[*models.Animate]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCircle = BackRepoCircleStruct{
		Map_CircleDBID_CirclePtr: make(map[uint]*models.Circle, 0),
		Map_CircleDBID_CircleDB:  make(map[uint]*CircleDB, 0),
		Map_CirclePtr_CircleDBID: make(map[*models.Circle]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEllipse = BackRepoEllipseStruct{
		Map_EllipseDBID_EllipsePtr: make(map[uint]*models.Ellipse, 0),
		Map_EllipseDBID_EllipseDB:  make(map[uint]*EllipseDB, 0),
		Map_EllipsePtr_EllipseDBID: make(map[*models.Ellipse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLayer = BackRepoLayerStruct{
		Map_LayerDBID_LayerPtr: make(map[uint]*models.Layer, 0),
		Map_LayerDBID_LayerDB:  make(map[uint]*LayerDB, 0),
		Map_LayerPtr_LayerDBID: make(map[*models.Layer]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLine = BackRepoLineStruct{
		Map_LineDBID_LinePtr: make(map[uint]*models.Line, 0),
		Map_LineDBID_LineDB:  make(map[uint]*LineDB, 0),
		Map_LinePtr_LineDBID: make(map[*models.Line]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLink = BackRepoLinkStruct{
		Map_LinkDBID_LinkPtr: make(map[uint]*models.Link, 0),
		Map_LinkDBID_LinkDB:  make(map[uint]*LinkDB, 0),
		Map_LinkPtr_LinkDBID: make(map[*models.Link]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLinkAnchoredText = BackRepoLinkAnchoredTextStruct{
		Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr: make(map[uint]*models.LinkAnchoredText, 0),
		Map_LinkAnchoredTextDBID_LinkAnchoredTextDB:  make(map[uint]*LinkAnchoredTextDB, 0),
		Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID: make(map[*models.LinkAnchoredText]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPath = BackRepoPathStruct{
		Map_PathDBID_PathPtr: make(map[uint]*models.Path, 0),
		Map_PathDBID_PathDB:  make(map[uint]*PathDB, 0),
		Map_PathPtr_PathDBID: make(map[*models.Path]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPoint = BackRepoPointStruct{
		Map_PointDBID_PointPtr: make(map[uint]*models.Point, 0),
		Map_PointDBID_PointDB:  make(map[uint]*PointDB, 0),
		Map_PointPtr_PointDBID: make(map[*models.Point]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPolygone = BackRepoPolygoneStruct{
		Map_PolygoneDBID_PolygonePtr: make(map[uint]*models.Polygone, 0),
		Map_PolygoneDBID_PolygoneDB:  make(map[uint]*PolygoneDB, 0),
		Map_PolygonePtr_PolygoneDBID: make(map[*models.Polygone]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPolyline = BackRepoPolylineStruct{
		Map_PolylineDBID_PolylinePtr: make(map[uint]*models.Polyline, 0),
		Map_PolylineDBID_PolylineDB:  make(map[uint]*PolylineDB, 0),
		Map_PolylinePtr_PolylineDBID: make(map[*models.Polyline]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRect = BackRepoRectStruct{
		Map_RectDBID_RectPtr: make(map[uint]*models.Rect, 0),
		Map_RectDBID_RectDB:  make(map[uint]*RectDB, 0),
		Map_RectPtr_RectDBID: make(map[*models.Rect]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRectAnchoredPath = BackRepoRectAnchoredPathStruct{
		Map_RectAnchoredPathDBID_RectAnchoredPathPtr: make(map[uint]*models.RectAnchoredPath, 0),
		Map_RectAnchoredPathDBID_RectAnchoredPathDB:  make(map[uint]*RectAnchoredPathDB, 0),
		Map_RectAnchoredPathPtr_RectAnchoredPathDBID: make(map[*models.RectAnchoredPath]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRectAnchoredRect = BackRepoRectAnchoredRectStruct{
		Map_RectAnchoredRectDBID_RectAnchoredRectPtr: make(map[uint]*models.RectAnchoredRect, 0),
		Map_RectAnchoredRectDBID_RectAnchoredRectDB:  make(map[uint]*RectAnchoredRectDB, 0),
		Map_RectAnchoredRectPtr_RectAnchoredRectDBID: make(map[*models.RectAnchoredRect]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRectAnchoredText = BackRepoRectAnchoredTextStruct{
		Map_RectAnchoredTextDBID_RectAnchoredTextPtr: make(map[uint]*models.RectAnchoredText, 0),
		Map_RectAnchoredTextDBID_RectAnchoredTextDB:  make(map[uint]*RectAnchoredTextDB, 0),
		Map_RectAnchoredTextPtr_RectAnchoredTextDBID: make(map[*models.RectAnchoredText]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRectLinkLink = BackRepoRectLinkLinkStruct{
		Map_RectLinkLinkDBID_RectLinkLinkPtr: make(map[uint]*models.RectLinkLink, 0),
		Map_RectLinkLinkDBID_RectLinkLinkDB:  make(map[uint]*RectLinkLinkDB, 0),
		Map_RectLinkLinkPtr_RectLinkLinkDBID: make(map[*models.RectLinkLink]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSVG = BackRepoSVGStruct{
		Map_SVGDBID_SVGPtr: make(map[uint]*models.SVG, 0),
		Map_SVGDBID_SVGDB:  make(map[uint]*SVGDB, 0),
		Map_SVGPtr_SVGDBID: make(map[*models.SVG]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSvgText = BackRepoSvgTextStruct{
		Map_SvgTextDBID_SvgTextPtr: make(map[uint]*models.SvgText, 0),
		Map_SvgTextDBID_SvgTextDB:  make(map[uint]*SvgTextDB, 0),
		Map_SvgTextPtr_SvgTextDBID: make(map[*models.SvgText]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoText = BackRepoTextStruct{
		Map_TextDBID_TextPtr: make(map[uint]*models.Text, 0),
		Map_TextDBID_TextDB:  make(map[uint]*TextDB, 0),
		Map_TextPtr_TextDBID: make(map[*models.Text]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.Stage) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1

	backRepo.broadcastNbCommitToBack()

	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.Stage) {

	// forbid read of back repo during commit
	backRepo.rwMutex.Lock()

	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAnimate.CommitPhaseOne(stage)
	backRepo.BackRepoCircle.CommitPhaseOne(stage)
	backRepo.BackRepoEllipse.CommitPhaseOne(stage)
	backRepo.BackRepoLayer.CommitPhaseOne(stage)
	backRepo.BackRepoLine.CommitPhaseOne(stage)
	backRepo.BackRepoLink.CommitPhaseOne(stage)
	backRepo.BackRepoLinkAnchoredText.CommitPhaseOne(stage)
	backRepo.BackRepoPath.CommitPhaseOne(stage)
	backRepo.BackRepoPoint.CommitPhaseOne(stage)
	backRepo.BackRepoPolygone.CommitPhaseOne(stage)
	backRepo.BackRepoPolyline.CommitPhaseOne(stage)
	backRepo.BackRepoRect.CommitPhaseOne(stage)
	backRepo.BackRepoRectAnchoredPath.CommitPhaseOne(stage)
	backRepo.BackRepoRectAnchoredRect.CommitPhaseOne(stage)
	backRepo.BackRepoRectAnchoredText.CommitPhaseOne(stage)
	backRepo.BackRepoRectLinkLink.CommitPhaseOne(stage)
	backRepo.BackRepoSVG.CommitPhaseOne(stage)
	backRepo.BackRepoSvgText.CommitPhaseOne(stage)
	backRepo.BackRepoText.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAnimate.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCircle.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEllipse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLayer.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLine.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLink.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLinkAnchoredText.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPath.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPoint.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPolygone.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPolyline.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRect.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRectAnchoredPath.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRectAnchoredRect.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRectAnchoredText.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRectLinkLink.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSVG.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSvgText.CommitPhaseTwo(backRepo)
	backRepo.BackRepoText.CommitPhaseTwo(backRepo)

	// important to release the mutex before calls to IncrementCommitFromBackNb
	// because it will block otherwise
	backRepo.rwMutex.Unlock()

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.Stage) {

	backRepo.rwMutex.Lock()
	defer backRepo.rwMutex.Unlock()
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAnimate.CheckoutPhaseOne()
	backRepo.BackRepoCircle.CheckoutPhaseOne()
	backRepo.BackRepoEllipse.CheckoutPhaseOne()
	backRepo.BackRepoLayer.CheckoutPhaseOne()
	backRepo.BackRepoLine.CheckoutPhaseOne()
	backRepo.BackRepoLink.CheckoutPhaseOne()
	backRepo.BackRepoLinkAnchoredText.CheckoutPhaseOne()
	backRepo.BackRepoPath.CheckoutPhaseOne()
	backRepo.BackRepoPoint.CheckoutPhaseOne()
	backRepo.BackRepoPolygone.CheckoutPhaseOne()
	backRepo.BackRepoPolyline.CheckoutPhaseOne()
	backRepo.BackRepoRect.CheckoutPhaseOne()
	backRepo.BackRepoRectAnchoredPath.CheckoutPhaseOne()
	backRepo.BackRepoRectAnchoredRect.CheckoutPhaseOne()
	backRepo.BackRepoRectAnchoredText.CheckoutPhaseOne()
	backRepo.BackRepoRectLinkLink.CheckoutPhaseOne()
	backRepo.BackRepoSVG.CheckoutPhaseOne()
	backRepo.BackRepoSvgText.CheckoutPhaseOne()
	backRepo.BackRepoText.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAnimate.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCircle.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEllipse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLayer.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLine.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLink.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLinkAnchoredText.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPath.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPoint.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPolygone.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPolyline.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRect.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRectAnchoredPath.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRectAnchoredRect.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRectAnchoredText.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRectLinkLink.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSVG.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSvgText.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoText.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAnimate.Backup(dirPath)
	backRepo.BackRepoCircle.Backup(dirPath)
	backRepo.BackRepoEllipse.Backup(dirPath)
	backRepo.BackRepoLayer.Backup(dirPath)
	backRepo.BackRepoLine.Backup(dirPath)
	backRepo.BackRepoLink.Backup(dirPath)
	backRepo.BackRepoLinkAnchoredText.Backup(dirPath)
	backRepo.BackRepoPath.Backup(dirPath)
	backRepo.BackRepoPoint.Backup(dirPath)
	backRepo.BackRepoPolygone.Backup(dirPath)
	backRepo.BackRepoPolyline.Backup(dirPath)
	backRepo.BackRepoRect.Backup(dirPath)
	backRepo.BackRepoRectAnchoredPath.Backup(dirPath)
	backRepo.BackRepoRectAnchoredRect.Backup(dirPath)
	backRepo.BackRepoRectAnchoredText.Backup(dirPath)
	backRepo.BackRepoRectLinkLink.Backup(dirPath)
	backRepo.BackRepoSVG.Backup(dirPath)
	backRepo.BackRepoSvgText.Backup(dirPath)
	backRepo.BackRepoText.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAnimate.BackupXL(file)
	backRepo.BackRepoCircle.BackupXL(file)
	backRepo.BackRepoEllipse.BackupXL(file)
	backRepo.BackRepoLayer.BackupXL(file)
	backRepo.BackRepoLine.BackupXL(file)
	backRepo.BackRepoLink.BackupXL(file)
	backRepo.BackRepoLinkAnchoredText.BackupXL(file)
	backRepo.BackRepoPath.BackupXL(file)
	backRepo.BackRepoPoint.BackupXL(file)
	backRepo.BackRepoPolygone.BackupXL(file)
	backRepo.BackRepoPolyline.BackupXL(file)
	backRepo.BackRepoRect.BackupXL(file)
	backRepo.BackRepoRectAnchoredPath.BackupXL(file)
	backRepo.BackRepoRectAnchoredRect.BackupXL(file)
	backRepo.BackRepoRectAnchoredText.BackupXL(file)
	backRepo.BackRepoRectLinkLink.BackupXL(file)
	backRepo.BackRepoSVG.BackupXL(file)
	backRepo.BackRepoSvgText.BackupXL(file)
	backRepo.BackRepoText.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.Stage, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAnimate.RestorePhaseOne(dirPath)
	backRepo.BackRepoCircle.RestorePhaseOne(dirPath)
	backRepo.BackRepoEllipse.RestorePhaseOne(dirPath)
	backRepo.BackRepoLayer.RestorePhaseOne(dirPath)
	backRepo.BackRepoLine.RestorePhaseOne(dirPath)
	backRepo.BackRepoLink.RestorePhaseOne(dirPath)
	backRepo.BackRepoLinkAnchoredText.RestorePhaseOne(dirPath)
	backRepo.BackRepoPath.RestorePhaseOne(dirPath)
	backRepo.BackRepoPoint.RestorePhaseOne(dirPath)
	backRepo.BackRepoPolygone.RestorePhaseOne(dirPath)
	backRepo.BackRepoPolyline.RestorePhaseOne(dirPath)
	backRepo.BackRepoRect.RestorePhaseOne(dirPath)
	backRepo.BackRepoRectAnchoredPath.RestorePhaseOne(dirPath)
	backRepo.BackRepoRectAnchoredRect.RestorePhaseOne(dirPath)
	backRepo.BackRepoRectAnchoredText.RestorePhaseOne(dirPath)
	backRepo.BackRepoRectLinkLink.RestorePhaseOne(dirPath)
	backRepo.BackRepoSVG.RestorePhaseOne(dirPath)
	backRepo.BackRepoSvgText.RestorePhaseOne(dirPath)
	backRepo.BackRepoText.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAnimate.RestorePhaseTwo()
	backRepo.BackRepoCircle.RestorePhaseTwo()
	backRepo.BackRepoEllipse.RestorePhaseTwo()
	backRepo.BackRepoLayer.RestorePhaseTwo()
	backRepo.BackRepoLine.RestorePhaseTwo()
	backRepo.BackRepoLink.RestorePhaseTwo()
	backRepo.BackRepoLinkAnchoredText.RestorePhaseTwo()
	backRepo.BackRepoPath.RestorePhaseTwo()
	backRepo.BackRepoPoint.RestorePhaseTwo()
	backRepo.BackRepoPolygone.RestorePhaseTwo()
	backRepo.BackRepoPolyline.RestorePhaseTwo()
	backRepo.BackRepoRect.RestorePhaseTwo()
	backRepo.BackRepoRectAnchoredPath.RestorePhaseTwo()
	backRepo.BackRepoRectAnchoredRect.RestorePhaseTwo()
	backRepo.BackRepoRectAnchoredText.RestorePhaseTwo()
	backRepo.BackRepoRectLinkLink.RestorePhaseTwo()
	backRepo.BackRepoSVG.RestorePhaseTwo()
	backRepo.BackRepoSvgText.RestorePhaseTwo()
	backRepo.BackRepoText.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.Stage, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAnimate.RestoreXLPhaseOne(file)
	backRepo.BackRepoCircle.RestoreXLPhaseOne(file)
	backRepo.BackRepoEllipse.RestoreXLPhaseOne(file)
	backRepo.BackRepoLayer.RestoreXLPhaseOne(file)
	backRepo.BackRepoLine.RestoreXLPhaseOne(file)
	backRepo.BackRepoLink.RestoreXLPhaseOne(file)
	backRepo.BackRepoLinkAnchoredText.RestoreXLPhaseOne(file)
	backRepo.BackRepoPath.RestoreXLPhaseOne(file)
	backRepo.BackRepoPoint.RestoreXLPhaseOne(file)
	backRepo.BackRepoPolygone.RestoreXLPhaseOne(file)
	backRepo.BackRepoPolyline.RestoreXLPhaseOne(file)
	backRepo.BackRepoRect.RestoreXLPhaseOne(file)
	backRepo.BackRepoRectAnchoredPath.RestoreXLPhaseOne(file)
	backRepo.BackRepoRectAnchoredRect.RestoreXLPhaseOne(file)
	backRepo.BackRepoRectAnchoredText.RestoreXLPhaseOne(file)
	backRepo.BackRepoRectLinkLink.RestoreXLPhaseOne(file)
	backRepo.BackRepoSVG.RestoreXLPhaseOne(file)
	backRepo.BackRepoSvgText.RestoreXLPhaseOne(file)
	backRepo.BackRepoText.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb(ctx context.Context) <-chan int {
	ch := make(chan int)

	backRepoStruct.subscribersRwMutex.Lock()
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	backRepoStruct.subscribersRwMutex.Unlock()

	// Goroutine to remove subscriber when context is done
	go func() {
		<-ctx.Done()
		backRepoStruct.unsubscribe(ch)
	}()
	return ch
}

// unsubscribe removes a subscriber's channel from the subscribers slice.
func (backRepoStruct *BackRepoStruct) unsubscribe(ch chan int) {
	backRepoStruct.subscribersRwMutex.Lock()
	defer backRepoStruct.subscribersRwMutex.Unlock()
	for i, subscriber := range backRepoStruct.subscribers {
		if subscriber == ch {
			backRepoStruct.subscribers =
				append(backRepoStruct.subscribers[:i],
					backRepoStruct.subscribers[i+1:]...)
			close(ch) // Close the channel to signal completion
			break
		}
	}
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.subscribersRwMutex.RLock()
	subscribers := make([]chan int, len(backRepoStruct.subscribers))
	copy(subscribers, backRepoStruct.subscribers)
	backRepoStruct.subscribersRwMutex.RUnlock()

	// if len(subscribers) == 0 {
	// 	log.Println(backRepoStruct.stage.GetType(), backRepoStruct.stage.GetName(), "no subsribers to broadcast to")
	// }


	for _, ch := range subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			// Successfully sent commit from back
		default:
			// Subscriber is not ready to receive; skip to avoid blocking
		}
	}
}
