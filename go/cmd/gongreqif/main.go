package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import
	gongreqif_models "github.com/fullstack-lang/gongreqif/go/models"
	gongreqif_stack "github.com/fullstack-lang/gongreqif/go/stack"
	gongreqif_static "github.com/fullstack-lang/gongreqif/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("gongreqif: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongreqif_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := gongreqif_stack.NewStack(r, "gongreqif", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// insertion point for call to stager
	gongreqif_models.NewStager(r, stack.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
