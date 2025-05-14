package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import
	"github.com/fullstack-lang/gongreqif/go/generator"
	gongreqif_models "github.com/fullstack-lang/gongreqif/go/models"
	gongreqif_stack "github.com/fullstack-lang/gongreqif/go/stack"
	gongreqif_static "github.com/fullstack-lang/gongreqif/go/static"

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	pathToReqifFile   = flag.String("pathToReqifFile", "", "Path to the reqif file")
	pathToGoModelFile = flag.String("pathToGoModelFile", "", "Path to the go model file")
)

func main() {

	log.SetPrefix("gongreqif: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongreqif_static.ServeStaticFiles(*logGINFlag)

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	// setup model stack with its probe
	stack := gongreqif_stack.NewStack(r, "gongreqif", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	modelGenerator := &generator.ModelGenerator{
		PathToGoModelFile: *pathToGoModelFile,
		Stage:             stack.Stage,
	}

	// insertion point for call to stager
	stager := gongreqif_models.NewStager(r, splitStage, stack.Stage, *pathToReqifFile, modelGenerator)

	modelGenerator.GenerateModels(stager)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
