package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	// insertion point for models import

	// the location of the following go package is important
	// They are NOT within the "github.com/fullstack-lang/gongreqif/go/models"
	// because this package is within the "//go:embed models" directive in
	// the "github.com/fullstack-lang/gongreqif/go" package
	// Therefore any change to those packge would pro
	"github.com/fullstack-lang/gongreqif/go/datatypes"
	"github.com/fullstack-lang/gongreqif/go/generator"
	"github.com/fullstack-lang/gongreqif/go/namer"
	"github.com/fullstack-lang/gongreqif/go/reqifz"
	"github.com/fullstack-lang/gongreqif/go/specifications"
	"github.com/fullstack-lang/gongreqif/go/specobjects"
	"github.com/fullstack-lang/gongreqif/go/specrelations"
	"github.com/fullstack-lang/gongreqif/go/spectypes"

	gongreqif_models "github.com/fullstack-lang/gongreqif/go/models"
	gongreqif_stack "github.com/fullstack-lang/gongreqif/go/stack"
	gongreqif_static "github.com/fullstack-lang/gongreqif/go/static"

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", true, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	pathToReqifFile   = flag.String("pathToReqifFile", "", "Path to the reqif file")
	pathToGoModelFile = flag.String("pathToGoModelFile", "", "Path to the go model file")
	pathToXLFile      = flag.String("pathToXLFile", "", "Path to the go model file")
)

func main() {

	log.SetPrefix("gongreqif: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	if flag.NArg() > 0 {
		pathToReqifOrReqifz := flag.Arg(0)
		isReqifz, err := reqifz.ConvertReQIFzToReqIF(pathToReqifOrReqifz)
		if err != nil {
			log.Fatalln("Not able to convert reqif file", err.Error())
		}

		if isReqifz {
			*pathToReqifFile = strings.TrimSuffix(pathToReqifOrReqifz, ".reqifz") + ".reqif"
		}
	}

	// setup the static file server and get the controller
	r := gongreqif_static.ServeStaticFiles(*logGINFlag)

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	// setup model stack with its probe
	stack := gongreqif_stack.NewStack(r, "gongreqif", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// insertion point for call to stager
	stager := gongreqif_models.NewStager(r,
		splitStage,
		stack.Stage,
		*pathToReqifFile,
		&datatypes.DataTypeTreeStageUpdater{},
		&spectypes.SpecTypesTreeStageUpdater{},
		&specobjects.SpecObjectsTreeStageUpdater{},
		&specrelations.SpecRelationsTreeStageUpdater{},
		&specifications.SpecificationsTreeStageUpdater{},
		&namer.ObjectNamer{})

	if *pathToGoModelFile != "" {
		modelGenerator := &generator.ModelGenerator{
			PathToGoModelFile: *pathToGoModelFile,
		}
		stager.SetModelGenerator(modelGenerator)
		modelGenerator.GenerateModels(stager)
	}

	if *pathToXLFile != "" {
		gongreqif_models.SerializeStage(stack.Stage, *pathToXLFile)
	}

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
