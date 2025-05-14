package generator

import (
	"log"
	"os"
	"strings"

	"github.com/fullstack-lang/gongreqif/go/models"
)

type ModelGenerator struct {
	PathToGoModelFile string
}

// GenerateModels implements models.ModelGeneratorInterface.
func (m *ModelGenerator) GenerateModels(stager *models.Stager) {
	var sb strings.Builder

	sb.WriteString(`// Generated code, do not edit
package models
`)

	result := sb.String()
	err := os.WriteFile(m.PathToGoModelFile, []byte(result), 0644)
	if err != nil {
		log.Panicln("Generate go models", err.Error())
	}
}
