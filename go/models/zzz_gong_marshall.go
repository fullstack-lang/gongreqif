// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const MetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(modelsPackageName, packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(modelsPackageName, packageName string) (res string, err error) {

	res = marshallRes
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_ALTERNATIVE_ID_Identifiers := make(map[*ALTERNATIVE_ID]string)
	_ = map_ALTERNATIVE_ID_Identifiers

	alternative_idOrdered := []*ALTERNATIVE_ID{}
	for alternative_id := range stage.ALTERNATIVE_IDs {
		alternative_idOrdered = append(alternative_idOrdered, alternative_id)
	}
	sort.Slice(alternative_idOrdered[:], func(i, j int) bool {
		alternative_idi := alternative_idOrdered[i]
		alternative_idj := alternative_idOrdered[j]
		alternative_idi_order, oki := stage.ALTERNATIVE_IDMap_Staged_Order[alternative_idi]
		alternative_idj_order, okj := stage.ALTERNATIVE_IDMap_Staged_Order[alternative_idj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return alternative_idi_order < alternative_idj_order
	})
	if len(alternative_idOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, alternative_id := range alternative_idOrdered {

		id = generatesIdentifier("ALTERNATIVE_ID", int(stage.ALTERNATIVE_IDMap_Staged_Order[alternative_id]), alternative_id.Name)
		map_ALTERNATIVE_ID_Identifiers[alternative_id] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ALTERNATIVE_ID")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", alternative_id.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(alternative_id.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(alternative_id.IDENTIFIER))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers := make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]string)
	_ = map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers

	attribute_definition_booleanOrdered := []*ATTRIBUTE_DEFINITION_BOOLEAN{}
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		attribute_definition_booleanOrdered = append(attribute_definition_booleanOrdered, attribute_definition_boolean)
	}
	sort.Slice(attribute_definition_booleanOrdered[:], func(i, j int) bool {
		attribute_definition_booleani := attribute_definition_booleanOrdered[i]
		attribute_definition_booleanj := attribute_definition_booleanOrdered[j]
		attribute_definition_booleani_order, oki := stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_booleani]
		attribute_definition_booleanj_order, okj := stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_booleanj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_booleani_order < attribute_definition_booleanj_order
	})
	if len(attribute_definition_booleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_boolean := range attribute_definition_booleanOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_BOOLEAN", int(stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_boolean]), attribute_definition_boolean.Name)
		map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[attribute_definition_boolean] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_BOOLEAN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_boolean.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_boolean.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_Identifiers := make(map[*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering]string)
	_ = map_ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_Identifiers

	attribute_definition_boolean_renderingOrdered := []*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering{}
	for attribute_definition_boolean_rendering := range stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings {
		attribute_definition_boolean_renderingOrdered = append(attribute_definition_boolean_renderingOrdered, attribute_definition_boolean_rendering)
	}
	sort.Slice(attribute_definition_boolean_renderingOrdered[:], func(i, j int) bool {
		attribute_definition_boolean_renderingi := attribute_definition_boolean_renderingOrdered[i]
		attribute_definition_boolean_renderingj := attribute_definition_boolean_renderingOrdered[j]
		attribute_definition_boolean_renderingi_order, oki := stage.ATTRIBUTE_DEFINITION_BOOLEAN_RenderingMap_Staged_Order[attribute_definition_boolean_renderingi]
		attribute_definition_boolean_renderingj_order, okj := stage.ATTRIBUTE_DEFINITION_BOOLEAN_RenderingMap_Staged_Order[attribute_definition_boolean_renderingj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_boolean_renderingi_order < attribute_definition_boolean_renderingj_order
	})
	if len(attribute_definition_boolean_renderingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_boolean_rendering := range attribute_definition_boolean_renderingOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_BOOLEAN_Rendering", int(stage.ATTRIBUTE_DEFINITION_BOOLEAN_RenderingMap_Staged_Order[attribute_definition_boolean_rendering]), attribute_definition_boolean_rendering.Name)
		map_ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_Identifiers[attribute_definition_boolean_rendering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_BOOLEAN_Rendering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_boolean_rendering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean_rendering.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_boolean_rendering.ShowInTable))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTitle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_boolean_rendering.ShowInTitle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInSubject")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_boolean_rendering.ShowInSubject))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_DATE_Identifiers := make(map[*ATTRIBUTE_DEFINITION_DATE]string)
	_ = map_ATTRIBUTE_DEFINITION_DATE_Identifiers

	attribute_definition_dateOrdered := []*ATTRIBUTE_DEFINITION_DATE{}
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		attribute_definition_dateOrdered = append(attribute_definition_dateOrdered, attribute_definition_date)
	}
	sort.Slice(attribute_definition_dateOrdered[:], func(i, j int) bool {
		attribute_definition_datei := attribute_definition_dateOrdered[i]
		attribute_definition_datej := attribute_definition_dateOrdered[j]
		attribute_definition_datei_order, oki := stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[attribute_definition_datei]
		attribute_definition_datej_order, okj := stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[attribute_definition_datej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_datei_order < attribute_definition_datej_order
	})
	if len(attribute_definition_dateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_date := range attribute_definition_dateOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_DATE", int(stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[attribute_definition_date]), attribute_definition_date.Name)
		map_ATTRIBUTE_DEFINITION_DATE_Identifiers[attribute_definition_date] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_DATE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_date.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_date.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_DATE_Rendering_Identifiers := make(map[*ATTRIBUTE_DEFINITION_DATE_Rendering]string)
	_ = map_ATTRIBUTE_DEFINITION_DATE_Rendering_Identifiers

	attribute_definition_date_renderingOrdered := []*ATTRIBUTE_DEFINITION_DATE_Rendering{}
	for attribute_definition_date_rendering := range stage.ATTRIBUTE_DEFINITION_DATE_Renderings {
		attribute_definition_date_renderingOrdered = append(attribute_definition_date_renderingOrdered, attribute_definition_date_rendering)
	}
	sort.Slice(attribute_definition_date_renderingOrdered[:], func(i, j int) bool {
		attribute_definition_date_renderingi := attribute_definition_date_renderingOrdered[i]
		attribute_definition_date_renderingj := attribute_definition_date_renderingOrdered[j]
		attribute_definition_date_renderingi_order, oki := stage.ATTRIBUTE_DEFINITION_DATE_RenderingMap_Staged_Order[attribute_definition_date_renderingi]
		attribute_definition_date_renderingj_order, okj := stage.ATTRIBUTE_DEFINITION_DATE_RenderingMap_Staged_Order[attribute_definition_date_renderingj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_date_renderingi_order < attribute_definition_date_renderingj_order
	})
	if len(attribute_definition_date_renderingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_date_rendering := range attribute_definition_date_renderingOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_DATE_Rendering", int(stage.ATTRIBUTE_DEFINITION_DATE_RenderingMap_Staged_Order[attribute_definition_date_rendering]), attribute_definition_date_rendering.Name)
		map_ATTRIBUTE_DEFINITION_DATE_Rendering_Identifiers[attribute_definition_date_rendering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_DATE_Rendering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_date_rendering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date_rendering.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_date_rendering.ShowInTable))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTitle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_date_rendering.ShowInTitle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInSubject")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_date_rendering.ShowInSubject))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers := make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]string)
	_ = map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers

	attribute_definition_enumerationOrdered := []*ATTRIBUTE_DEFINITION_ENUMERATION{}
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		attribute_definition_enumerationOrdered = append(attribute_definition_enumerationOrdered, attribute_definition_enumeration)
	}
	sort.Slice(attribute_definition_enumerationOrdered[:], func(i, j int) bool {
		attribute_definition_enumerationi := attribute_definition_enumerationOrdered[i]
		attribute_definition_enumerationj := attribute_definition_enumerationOrdered[j]
		attribute_definition_enumerationi_order, oki := stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[attribute_definition_enumerationi]
		attribute_definition_enumerationj_order, okj := stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[attribute_definition_enumerationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_enumerationi_order < attribute_definition_enumerationj_order
	})
	if len(attribute_definition_enumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_enumeration := range attribute_definition_enumerationOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_ENUMERATION", int(stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[attribute_definition_enumeration]), attribute_definition_enumeration.Name)
		map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[attribute_definition_enumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_ENUMERATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_enumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_enumeration.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.LONG_NAME))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MULTI_VALUED")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_enumeration.MULTI_VALUED))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_Identifiers := make(map[*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering]string)
	_ = map_ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_Identifiers

	attribute_definition_enumeration_renderingOrdered := []*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering{}
	for attribute_definition_enumeration_rendering := range stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings {
		attribute_definition_enumeration_renderingOrdered = append(attribute_definition_enumeration_renderingOrdered, attribute_definition_enumeration_rendering)
	}
	sort.Slice(attribute_definition_enumeration_renderingOrdered[:], func(i, j int) bool {
		attribute_definition_enumeration_renderingi := attribute_definition_enumeration_renderingOrdered[i]
		attribute_definition_enumeration_renderingj := attribute_definition_enumeration_renderingOrdered[j]
		attribute_definition_enumeration_renderingi_order, oki := stage.ATTRIBUTE_DEFINITION_ENUMERATION_RenderingMap_Staged_Order[attribute_definition_enumeration_renderingi]
		attribute_definition_enumeration_renderingj_order, okj := stage.ATTRIBUTE_DEFINITION_ENUMERATION_RenderingMap_Staged_Order[attribute_definition_enumeration_renderingj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_enumeration_renderingi_order < attribute_definition_enumeration_renderingj_order
	})
	if len(attribute_definition_enumeration_renderingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_enumeration_rendering := range attribute_definition_enumeration_renderingOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_ENUMERATION_Rendering", int(stage.ATTRIBUTE_DEFINITION_ENUMERATION_RenderingMap_Staged_Order[attribute_definition_enumeration_rendering]), attribute_definition_enumeration_rendering.Name)
		map_ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_Identifiers[attribute_definition_enumeration_rendering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_ENUMERATION_Rendering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_enumeration_rendering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration_rendering.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_enumeration_rendering.ShowInTable))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTitle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_enumeration_rendering.ShowInTitle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInSubject")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_enumeration_rendering.ShowInSubject))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers := make(map[*ATTRIBUTE_DEFINITION_INTEGER]string)
	_ = map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers

	attribute_definition_integerOrdered := []*ATTRIBUTE_DEFINITION_INTEGER{}
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		attribute_definition_integerOrdered = append(attribute_definition_integerOrdered, attribute_definition_integer)
	}
	sort.Slice(attribute_definition_integerOrdered[:], func(i, j int) bool {
		attribute_definition_integeri := attribute_definition_integerOrdered[i]
		attribute_definition_integerj := attribute_definition_integerOrdered[j]
		attribute_definition_integeri_order, oki := stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[attribute_definition_integeri]
		attribute_definition_integerj_order, okj := stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[attribute_definition_integerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_integeri_order < attribute_definition_integerj_order
	})
	if len(attribute_definition_integerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_integer := range attribute_definition_integerOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_INTEGER", int(stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[attribute_definition_integer]), attribute_definition_integer.Name)
		map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[attribute_definition_integer] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_INTEGER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_integer.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_integer.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_INTEGER_Rendering_Identifiers := make(map[*ATTRIBUTE_DEFINITION_INTEGER_Rendering]string)
	_ = map_ATTRIBUTE_DEFINITION_INTEGER_Rendering_Identifiers

	attribute_definition_integer_renderingOrdered := []*ATTRIBUTE_DEFINITION_INTEGER_Rendering{}
	for attribute_definition_integer_rendering := range stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings {
		attribute_definition_integer_renderingOrdered = append(attribute_definition_integer_renderingOrdered, attribute_definition_integer_rendering)
	}
	sort.Slice(attribute_definition_integer_renderingOrdered[:], func(i, j int) bool {
		attribute_definition_integer_renderingi := attribute_definition_integer_renderingOrdered[i]
		attribute_definition_integer_renderingj := attribute_definition_integer_renderingOrdered[j]
		attribute_definition_integer_renderingi_order, oki := stage.ATTRIBUTE_DEFINITION_INTEGER_RenderingMap_Staged_Order[attribute_definition_integer_renderingi]
		attribute_definition_integer_renderingj_order, okj := stage.ATTRIBUTE_DEFINITION_INTEGER_RenderingMap_Staged_Order[attribute_definition_integer_renderingj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_integer_renderingi_order < attribute_definition_integer_renderingj_order
	})
	if len(attribute_definition_integer_renderingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_integer_rendering := range attribute_definition_integer_renderingOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_INTEGER_Rendering", int(stage.ATTRIBUTE_DEFINITION_INTEGER_RenderingMap_Staged_Order[attribute_definition_integer_rendering]), attribute_definition_integer_rendering.Name)
		map_ATTRIBUTE_DEFINITION_INTEGER_Rendering_Identifiers[attribute_definition_integer_rendering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_INTEGER_Rendering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_integer_rendering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer_rendering.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_integer_rendering.ShowInTable))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTitle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_integer_rendering.ShowInTitle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInSubject")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_integer_rendering.ShowInSubject))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_REAL_Identifiers := make(map[*ATTRIBUTE_DEFINITION_REAL]string)
	_ = map_ATTRIBUTE_DEFINITION_REAL_Identifiers

	attribute_definition_realOrdered := []*ATTRIBUTE_DEFINITION_REAL{}
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		attribute_definition_realOrdered = append(attribute_definition_realOrdered, attribute_definition_real)
	}
	sort.Slice(attribute_definition_realOrdered[:], func(i, j int) bool {
		attribute_definition_reali := attribute_definition_realOrdered[i]
		attribute_definition_realj := attribute_definition_realOrdered[j]
		attribute_definition_reali_order, oki := stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[attribute_definition_reali]
		attribute_definition_realj_order, okj := stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[attribute_definition_realj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_reali_order < attribute_definition_realj_order
	})
	if len(attribute_definition_realOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_real := range attribute_definition_realOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_REAL", int(stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[attribute_definition_real]), attribute_definition_real.Name)
		map_ATTRIBUTE_DEFINITION_REAL_Identifiers[attribute_definition_real] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_REAL")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_real.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_real.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_REAL_Rendering_Identifiers := make(map[*ATTRIBUTE_DEFINITION_REAL_Rendering]string)
	_ = map_ATTRIBUTE_DEFINITION_REAL_Rendering_Identifiers

	attribute_definition_real_renderingOrdered := []*ATTRIBUTE_DEFINITION_REAL_Rendering{}
	for attribute_definition_real_rendering := range stage.ATTRIBUTE_DEFINITION_REAL_Renderings {
		attribute_definition_real_renderingOrdered = append(attribute_definition_real_renderingOrdered, attribute_definition_real_rendering)
	}
	sort.Slice(attribute_definition_real_renderingOrdered[:], func(i, j int) bool {
		attribute_definition_real_renderingi := attribute_definition_real_renderingOrdered[i]
		attribute_definition_real_renderingj := attribute_definition_real_renderingOrdered[j]
		attribute_definition_real_renderingi_order, oki := stage.ATTRIBUTE_DEFINITION_REAL_RenderingMap_Staged_Order[attribute_definition_real_renderingi]
		attribute_definition_real_renderingj_order, okj := stage.ATTRIBUTE_DEFINITION_REAL_RenderingMap_Staged_Order[attribute_definition_real_renderingj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_real_renderingi_order < attribute_definition_real_renderingj_order
	})
	if len(attribute_definition_real_renderingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_real_rendering := range attribute_definition_real_renderingOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_REAL_Rendering", int(stage.ATTRIBUTE_DEFINITION_REAL_RenderingMap_Staged_Order[attribute_definition_real_rendering]), attribute_definition_real_rendering.Name)
		map_ATTRIBUTE_DEFINITION_REAL_Rendering_Identifiers[attribute_definition_real_rendering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_REAL_Rendering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_real_rendering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real_rendering.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_real_rendering.ShowInTable))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTitle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_real_rendering.ShowInTitle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInSubject")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_real_rendering.ShowInSubject))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_Rendering_Identifiers := make(map[*ATTRIBUTE_DEFINITION_Rendering]string)
	_ = map_ATTRIBUTE_DEFINITION_Rendering_Identifiers

	attribute_definition_renderingOrdered := []*ATTRIBUTE_DEFINITION_Rendering{}
	for attribute_definition_rendering := range stage.ATTRIBUTE_DEFINITION_Renderings {
		attribute_definition_renderingOrdered = append(attribute_definition_renderingOrdered, attribute_definition_rendering)
	}
	sort.Slice(attribute_definition_renderingOrdered[:], func(i, j int) bool {
		attribute_definition_renderingi := attribute_definition_renderingOrdered[i]
		attribute_definition_renderingj := attribute_definition_renderingOrdered[j]
		attribute_definition_renderingi_order, oki := stage.ATTRIBUTE_DEFINITION_RenderingMap_Staged_Order[attribute_definition_renderingi]
		attribute_definition_renderingj_order, okj := stage.ATTRIBUTE_DEFINITION_RenderingMap_Staged_Order[attribute_definition_renderingj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_renderingi_order < attribute_definition_renderingj_order
	})
	if len(attribute_definition_renderingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_rendering := range attribute_definition_renderingOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_Rendering", int(stage.ATTRIBUTE_DEFINITION_RenderingMap_Staged_Order[attribute_definition_rendering]), attribute_definition_rendering.Name)
		map_ATTRIBUTE_DEFINITION_Rendering_Identifiers[attribute_definition_rendering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_Rendering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_rendering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_rendering.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_rendering.ShowInTable))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTitle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_rendering.ShowInTitle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInSubject")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_rendering.ShowInSubject))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_STRING_Identifiers := make(map[*ATTRIBUTE_DEFINITION_STRING]string)
	_ = map_ATTRIBUTE_DEFINITION_STRING_Identifiers

	attribute_definition_stringOrdered := []*ATTRIBUTE_DEFINITION_STRING{}
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		attribute_definition_stringOrdered = append(attribute_definition_stringOrdered, attribute_definition_string)
	}
	sort.Slice(attribute_definition_stringOrdered[:], func(i, j int) bool {
		attribute_definition_stringi := attribute_definition_stringOrdered[i]
		attribute_definition_stringj := attribute_definition_stringOrdered[j]
		attribute_definition_stringi_order, oki := stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_stringi]
		attribute_definition_stringj_order, okj := stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_stringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_stringi_order < attribute_definition_stringj_order
	})
	if len(attribute_definition_stringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_string := range attribute_definition_stringOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_STRING", int(stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_string]), attribute_definition_string.Name)
		map_ATTRIBUTE_DEFINITION_STRING_Identifiers[attribute_definition_string] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_STRING")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_string.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_string.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_STRING_Rendering_Identifiers := make(map[*ATTRIBUTE_DEFINITION_STRING_Rendering]string)
	_ = map_ATTRIBUTE_DEFINITION_STRING_Rendering_Identifiers

	attribute_definition_string_renderingOrdered := []*ATTRIBUTE_DEFINITION_STRING_Rendering{}
	for attribute_definition_string_rendering := range stage.ATTRIBUTE_DEFINITION_STRING_Renderings {
		attribute_definition_string_renderingOrdered = append(attribute_definition_string_renderingOrdered, attribute_definition_string_rendering)
	}
	sort.Slice(attribute_definition_string_renderingOrdered[:], func(i, j int) bool {
		attribute_definition_string_renderingi := attribute_definition_string_renderingOrdered[i]
		attribute_definition_string_renderingj := attribute_definition_string_renderingOrdered[j]
		attribute_definition_string_renderingi_order, oki := stage.ATTRIBUTE_DEFINITION_STRING_RenderingMap_Staged_Order[attribute_definition_string_renderingi]
		attribute_definition_string_renderingj_order, okj := stage.ATTRIBUTE_DEFINITION_STRING_RenderingMap_Staged_Order[attribute_definition_string_renderingj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_string_renderingi_order < attribute_definition_string_renderingj_order
	})
	if len(attribute_definition_string_renderingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_string_rendering := range attribute_definition_string_renderingOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_STRING_Rendering", int(stage.ATTRIBUTE_DEFINITION_STRING_RenderingMap_Staged_Order[attribute_definition_string_rendering]), attribute_definition_string_rendering.Name)
		map_ATTRIBUTE_DEFINITION_STRING_Rendering_Identifiers[attribute_definition_string_rendering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_STRING_Rendering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_string_rendering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string_rendering.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_string_rendering.ShowInTable))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTitle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_string_rendering.ShowInTitle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInSubject")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_string_rendering.ShowInSubject))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_XHTML_Identifiers := make(map[*ATTRIBUTE_DEFINITION_XHTML]string)
	_ = map_ATTRIBUTE_DEFINITION_XHTML_Identifiers

	attribute_definition_xhtmlOrdered := []*ATTRIBUTE_DEFINITION_XHTML{}
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		attribute_definition_xhtmlOrdered = append(attribute_definition_xhtmlOrdered, attribute_definition_xhtml)
	}
	sort.Slice(attribute_definition_xhtmlOrdered[:], func(i, j int) bool {
		attribute_definition_xhtmli := attribute_definition_xhtmlOrdered[i]
		attribute_definition_xhtmlj := attribute_definition_xhtmlOrdered[j]
		attribute_definition_xhtmli_order, oki := stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[attribute_definition_xhtmli]
		attribute_definition_xhtmlj_order, okj := stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[attribute_definition_xhtmlj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_xhtmli_order < attribute_definition_xhtmlj_order
	})
	if len(attribute_definition_xhtmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_xhtml := range attribute_definition_xhtmlOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_XHTML", int(stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[attribute_definition_xhtml]), attribute_definition_xhtml.Name)
		map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[attribute_definition_xhtml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_XHTML")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_xhtml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_xhtml.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_XHTML_Rendering_Identifiers := make(map[*ATTRIBUTE_DEFINITION_XHTML_Rendering]string)
	_ = map_ATTRIBUTE_DEFINITION_XHTML_Rendering_Identifiers

	attribute_definition_xhtml_renderingOrdered := []*ATTRIBUTE_DEFINITION_XHTML_Rendering{}
	for attribute_definition_xhtml_rendering := range stage.ATTRIBUTE_DEFINITION_XHTML_Renderings {
		attribute_definition_xhtml_renderingOrdered = append(attribute_definition_xhtml_renderingOrdered, attribute_definition_xhtml_rendering)
	}
	sort.Slice(attribute_definition_xhtml_renderingOrdered[:], func(i, j int) bool {
		attribute_definition_xhtml_renderingi := attribute_definition_xhtml_renderingOrdered[i]
		attribute_definition_xhtml_renderingj := attribute_definition_xhtml_renderingOrdered[j]
		attribute_definition_xhtml_renderingi_order, oki := stage.ATTRIBUTE_DEFINITION_XHTML_RenderingMap_Staged_Order[attribute_definition_xhtml_renderingi]
		attribute_definition_xhtml_renderingj_order, okj := stage.ATTRIBUTE_DEFINITION_XHTML_RenderingMap_Staged_Order[attribute_definition_xhtml_renderingj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_xhtml_renderingi_order < attribute_definition_xhtml_renderingj_order
	})
	if len(attribute_definition_xhtml_renderingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_definition_xhtml_rendering := range attribute_definition_xhtml_renderingOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_XHTML_Rendering", int(stage.ATTRIBUTE_DEFINITION_XHTML_RenderingMap_Staged_Order[attribute_definition_xhtml_rendering]), attribute_definition_xhtml_rendering.Name)
		map_ATTRIBUTE_DEFINITION_XHTML_Rendering_Identifiers[attribute_definition_xhtml_rendering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_XHTML_Rendering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_xhtml_rendering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml_rendering.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_xhtml_rendering.ShowInTable))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInTitle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_xhtml_rendering.ShowInTitle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInSubject")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_xhtml_rendering.ShowInSubject))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers := make(map[*ATTRIBUTE_VALUE_BOOLEAN]string)
	_ = map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers

	attribute_value_booleanOrdered := []*ATTRIBUTE_VALUE_BOOLEAN{}
	for attribute_value_boolean := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		attribute_value_booleanOrdered = append(attribute_value_booleanOrdered, attribute_value_boolean)
	}
	sort.Slice(attribute_value_booleanOrdered[:], func(i, j int) bool {
		attribute_value_booleani := attribute_value_booleanOrdered[i]
		attribute_value_booleanj := attribute_value_booleanOrdered[j]
		attribute_value_booleani_order, oki := stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[attribute_value_booleani]
		attribute_value_booleanj_order, okj := stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[attribute_value_booleanj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_booleani_order < attribute_value_booleanj_order
	})
	if len(attribute_value_booleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_value_boolean := range attribute_value_booleanOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_BOOLEAN", int(stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[attribute_value_boolean]), attribute_value_boolean.Name)
		map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[attribute_value_boolean] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_BOOLEAN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_boolean.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_boolean.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_value_boolean.THE_VALUE))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_DATE_Identifiers := make(map[*ATTRIBUTE_VALUE_DATE]string)
	_ = map_ATTRIBUTE_VALUE_DATE_Identifiers

	attribute_value_dateOrdered := []*ATTRIBUTE_VALUE_DATE{}
	for attribute_value_date := range stage.ATTRIBUTE_VALUE_DATEs {
		attribute_value_dateOrdered = append(attribute_value_dateOrdered, attribute_value_date)
	}
	sort.Slice(attribute_value_dateOrdered[:], func(i, j int) bool {
		attribute_value_datei := attribute_value_dateOrdered[i]
		attribute_value_datej := attribute_value_dateOrdered[j]
		attribute_value_datei_order, oki := stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[attribute_value_datei]
		attribute_value_datej_order, okj := stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[attribute_value_datej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_datei_order < attribute_value_datej_order
	})
	if len(attribute_value_dateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_value_date := range attribute_value_dateOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_DATE", int(stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[attribute_value_date]), attribute_value_date.Name)
		map_ATTRIBUTE_VALUE_DATE_Identifiers[attribute_value_date] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_DATE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_date.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_date.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_date.THE_VALUE))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers := make(map[*ATTRIBUTE_VALUE_ENUMERATION]string)
	_ = map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers

	attribute_value_enumerationOrdered := []*ATTRIBUTE_VALUE_ENUMERATION{}
	for attribute_value_enumeration := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		attribute_value_enumerationOrdered = append(attribute_value_enumerationOrdered, attribute_value_enumeration)
	}
	sort.Slice(attribute_value_enumerationOrdered[:], func(i, j int) bool {
		attribute_value_enumerationi := attribute_value_enumerationOrdered[i]
		attribute_value_enumerationj := attribute_value_enumerationOrdered[j]
		attribute_value_enumerationi_order, oki := stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[attribute_value_enumerationi]
		attribute_value_enumerationj_order, okj := stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[attribute_value_enumerationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_enumerationi_order < attribute_value_enumerationj_order
	})
	if len(attribute_value_enumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_value_enumeration := range attribute_value_enumerationOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_ENUMERATION", int(stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[attribute_value_enumeration]), attribute_value_enumeration.Name)
		map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[attribute_value_enumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_ENUMERATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_enumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_enumeration.Name))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_INTEGER_Identifiers := make(map[*ATTRIBUTE_VALUE_INTEGER]string)
	_ = map_ATTRIBUTE_VALUE_INTEGER_Identifiers

	attribute_value_integerOrdered := []*ATTRIBUTE_VALUE_INTEGER{}
	for attribute_value_integer := range stage.ATTRIBUTE_VALUE_INTEGERs {
		attribute_value_integerOrdered = append(attribute_value_integerOrdered, attribute_value_integer)
	}
	sort.Slice(attribute_value_integerOrdered[:], func(i, j int) bool {
		attribute_value_integeri := attribute_value_integerOrdered[i]
		attribute_value_integerj := attribute_value_integerOrdered[j]
		attribute_value_integeri_order, oki := stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[attribute_value_integeri]
		attribute_value_integerj_order, okj := stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[attribute_value_integerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_integeri_order < attribute_value_integerj_order
	})
	if len(attribute_value_integerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_value_integer := range attribute_value_integerOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_INTEGER", int(stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[attribute_value_integer]), attribute_value_integer.Name)
		map_ATTRIBUTE_VALUE_INTEGER_Identifiers[attribute_value_integer] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_INTEGER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_integer.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_integer.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", attribute_value_integer.THE_VALUE))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_REAL_Identifiers := make(map[*ATTRIBUTE_VALUE_REAL]string)
	_ = map_ATTRIBUTE_VALUE_REAL_Identifiers

	attribute_value_realOrdered := []*ATTRIBUTE_VALUE_REAL{}
	for attribute_value_real := range stage.ATTRIBUTE_VALUE_REALs {
		attribute_value_realOrdered = append(attribute_value_realOrdered, attribute_value_real)
	}
	sort.Slice(attribute_value_realOrdered[:], func(i, j int) bool {
		attribute_value_reali := attribute_value_realOrdered[i]
		attribute_value_realj := attribute_value_realOrdered[j]
		attribute_value_reali_order, oki := stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[attribute_value_reali]
		attribute_value_realj_order, okj := stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[attribute_value_realj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_reali_order < attribute_value_realj_order
	})
	if len(attribute_value_realOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_value_real := range attribute_value_realOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_REAL", int(stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[attribute_value_real]), attribute_value_real.Name)
		map_ATTRIBUTE_VALUE_REAL_Identifiers[attribute_value_real] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_REAL")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_real.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_real.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", attribute_value_real.THE_VALUE))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_STRING_Identifiers := make(map[*ATTRIBUTE_VALUE_STRING]string)
	_ = map_ATTRIBUTE_VALUE_STRING_Identifiers

	attribute_value_stringOrdered := []*ATTRIBUTE_VALUE_STRING{}
	for attribute_value_string := range stage.ATTRIBUTE_VALUE_STRINGs {
		attribute_value_stringOrdered = append(attribute_value_stringOrdered, attribute_value_string)
	}
	sort.Slice(attribute_value_stringOrdered[:], func(i, j int) bool {
		attribute_value_stringi := attribute_value_stringOrdered[i]
		attribute_value_stringj := attribute_value_stringOrdered[j]
		attribute_value_stringi_order, oki := stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[attribute_value_stringi]
		attribute_value_stringj_order, okj := stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[attribute_value_stringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_stringi_order < attribute_value_stringj_order
	})
	if len(attribute_value_stringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_value_string := range attribute_value_stringOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_STRING", int(stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[attribute_value_string]), attribute_value_string.Name)
		map_ATTRIBUTE_VALUE_STRING_Identifiers[attribute_value_string] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_STRING")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_string.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_string.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_string.THE_VALUE))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_XHTML_Identifiers := make(map[*ATTRIBUTE_VALUE_XHTML]string)
	_ = map_ATTRIBUTE_VALUE_XHTML_Identifiers

	attribute_value_xhtmlOrdered := []*ATTRIBUTE_VALUE_XHTML{}
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		attribute_value_xhtmlOrdered = append(attribute_value_xhtmlOrdered, attribute_value_xhtml)
	}
	sort.Slice(attribute_value_xhtmlOrdered[:], func(i, j int) bool {
		attribute_value_xhtmli := attribute_value_xhtmlOrdered[i]
		attribute_value_xhtmlj := attribute_value_xhtmlOrdered[j]
		attribute_value_xhtmli_order, oki := stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtmli]
		attribute_value_xhtmlj_order, okj := stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtmlj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_xhtmli_order < attribute_value_xhtmlj_order
	})
	if len(attribute_value_xhtmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, attribute_value_xhtml := range attribute_value_xhtmlOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_XHTML", int(stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtml]), attribute_value_xhtml.Name)
		map_ATTRIBUTE_VALUE_XHTML_Identifiers[attribute_value_xhtml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_XHTML")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_xhtml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_xhtml.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_SIMPLIFIED")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_value_xhtml.IS_SIMPLIFIED))
		initializerStatements += setValueField

	}

	map_A_ALTERNATIVE_ID_Identifiers := make(map[*A_ALTERNATIVE_ID]string)
	_ = map_A_ALTERNATIVE_ID_Identifiers

	a_alternative_idOrdered := []*A_ALTERNATIVE_ID{}
	for a_alternative_id := range stage.A_ALTERNATIVE_IDs {
		a_alternative_idOrdered = append(a_alternative_idOrdered, a_alternative_id)
	}
	sort.Slice(a_alternative_idOrdered[:], func(i, j int) bool {
		a_alternative_idi := a_alternative_idOrdered[i]
		a_alternative_idj := a_alternative_idOrdered[j]
		a_alternative_idi_order, oki := stage.A_ALTERNATIVE_IDMap_Staged_Order[a_alternative_idi]
		a_alternative_idj_order, okj := stage.A_ALTERNATIVE_IDMap_Staged_Order[a_alternative_idj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_alternative_idi_order < a_alternative_idj_order
	})
	if len(a_alternative_idOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_alternative_id := range a_alternative_idOrdered {

		id = generatesIdentifier("A_ALTERNATIVE_ID", int(stage.A_ALTERNATIVE_IDMap_Staged_Order[a_alternative_id]), a_alternative_id.Name)
		map_A_ALTERNATIVE_ID_Identifiers[a_alternative_id] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ALTERNATIVE_ID")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_alternative_id.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_alternative_id.Name))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Identifiers := make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]string)
	_ = map_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Identifiers

	a_attribute_definition_boolean_refOrdered := []*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF{}
	for a_attribute_definition_boolean_ref := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
		a_attribute_definition_boolean_refOrdered = append(a_attribute_definition_boolean_refOrdered, a_attribute_definition_boolean_ref)
	}
	sort.Slice(a_attribute_definition_boolean_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_boolean_refi := a_attribute_definition_boolean_refOrdered[i]
		a_attribute_definition_boolean_refj := a_attribute_definition_boolean_refOrdered[j]
		a_attribute_definition_boolean_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_attribute_definition_boolean_refi]
		a_attribute_definition_boolean_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_attribute_definition_boolean_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_boolean_refi_order < a_attribute_definition_boolean_refj_order
	})
	if len(a_attribute_definition_boolean_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_definition_boolean_ref := range a_attribute_definition_boolean_refOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_BOOLEAN_REF", int(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_attribute_definition_boolean_ref]), a_attribute_definition_boolean_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Identifiers[a_attribute_definition_boolean_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_boolean_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_boolean_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_BOOLEAN_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_boolean_ref.ATTRIBUTE_DEFINITION_BOOLEAN_REF))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_DEFINITION_DATE_REF_Identifiers := make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF]string)
	_ = map_A_ATTRIBUTE_DEFINITION_DATE_REF_Identifiers

	a_attribute_definition_date_refOrdered := []*A_ATTRIBUTE_DEFINITION_DATE_REF{}
	for a_attribute_definition_date_ref := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
		a_attribute_definition_date_refOrdered = append(a_attribute_definition_date_refOrdered, a_attribute_definition_date_ref)
	}
	sort.Slice(a_attribute_definition_date_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_date_refi := a_attribute_definition_date_refOrdered[i]
		a_attribute_definition_date_refj := a_attribute_definition_date_refOrdered[j]
		a_attribute_definition_date_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order[a_attribute_definition_date_refi]
		a_attribute_definition_date_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order[a_attribute_definition_date_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_date_refi_order < a_attribute_definition_date_refj_order
	})
	if len(a_attribute_definition_date_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_definition_date_ref := range a_attribute_definition_date_refOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_DATE_REF", int(stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order[a_attribute_definition_date_ref]), a_attribute_definition_date_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_DATE_REF_Identifiers[a_attribute_definition_date_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_DATE_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_date_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_date_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_DATE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_date_ref.ATTRIBUTE_DEFINITION_DATE_REF))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Identifiers := make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]string)
	_ = map_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Identifiers

	a_attribute_definition_enumeration_refOrdered := []*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF{}
	for a_attribute_definition_enumeration_ref := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
		a_attribute_definition_enumeration_refOrdered = append(a_attribute_definition_enumeration_refOrdered, a_attribute_definition_enumeration_ref)
	}
	sort.Slice(a_attribute_definition_enumeration_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_enumeration_refi := a_attribute_definition_enumeration_refOrdered[i]
		a_attribute_definition_enumeration_refj := a_attribute_definition_enumeration_refOrdered[j]
		a_attribute_definition_enumeration_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_attribute_definition_enumeration_refi]
		a_attribute_definition_enumeration_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_attribute_definition_enumeration_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_enumeration_refi_order < a_attribute_definition_enumeration_refj_order
	})
	if len(a_attribute_definition_enumeration_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_definition_enumeration_ref := range a_attribute_definition_enumeration_refOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_ENUMERATION_REF", int(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_attribute_definition_enumeration_ref]), a_attribute_definition_enumeration_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Identifiers[a_attribute_definition_enumeration_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_enumeration_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_enumeration_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_ENUMERATION_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_enumeration_ref.ATTRIBUTE_DEFINITION_ENUMERATION_REF))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Identifiers := make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]string)
	_ = map_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Identifiers

	a_attribute_definition_integer_refOrdered := []*A_ATTRIBUTE_DEFINITION_INTEGER_REF{}
	for a_attribute_definition_integer_ref := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
		a_attribute_definition_integer_refOrdered = append(a_attribute_definition_integer_refOrdered, a_attribute_definition_integer_ref)
	}
	sort.Slice(a_attribute_definition_integer_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_integer_refi := a_attribute_definition_integer_refOrdered[i]
		a_attribute_definition_integer_refj := a_attribute_definition_integer_refOrdered[j]
		a_attribute_definition_integer_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order[a_attribute_definition_integer_refi]
		a_attribute_definition_integer_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order[a_attribute_definition_integer_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_integer_refi_order < a_attribute_definition_integer_refj_order
	})
	if len(a_attribute_definition_integer_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_definition_integer_ref := range a_attribute_definition_integer_refOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_INTEGER_REF", int(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order[a_attribute_definition_integer_ref]), a_attribute_definition_integer_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Identifiers[a_attribute_definition_integer_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_INTEGER_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_integer_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_integer_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_INTEGER_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_integer_ref.ATTRIBUTE_DEFINITION_INTEGER_REF))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_DEFINITION_REAL_REF_Identifiers := make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF]string)
	_ = map_A_ATTRIBUTE_DEFINITION_REAL_REF_Identifiers

	a_attribute_definition_real_refOrdered := []*A_ATTRIBUTE_DEFINITION_REAL_REF{}
	for a_attribute_definition_real_ref := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
		a_attribute_definition_real_refOrdered = append(a_attribute_definition_real_refOrdered, a_attribute_definition_real_ref)
	}
	sort.Slice(a_attribute_definition_real_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_real_refi := a_attribute_definition_real_refOrdered[i]
		a_attribute_definition_real_refj := a_attribute_definition_real_refOrdered[j]
		a_attribute_definition_real_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order[a_attribute_definition_real_refi]
		a_attribute_definition_real_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order[a_attribute_definition_real_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_real_refi_order < a_attribute_definition_real_refj_order
	})
	if len(a_attribute_definition_real_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_definition_real_ref := range a_attribute_definition_real_refOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_REAL_REF", int(stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order[a_attribute_definition_real_ref]), a_attribute_definition_real_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_REAL_REF_Identifiers[a_attribute_definition_real_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_REAL_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_real_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_real_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_REAL_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_real_ref.ATTRIBUTE_DEFINITION_REAL_REF))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_DEFINITION_STRING_REF_Identifiers := make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF]string)
	_ = map_A_ATTRIBUTE_DEFINITION_STRING_REF_Identifiers

	a_attribute_definition_string_refOrdered := []*A_ATTRIBUTE_DEFINITION_STRING_REF{}
	for a_attribute_definition_string_ref := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
		a_attribute_definition_string_refOrdered = append(a_attribute_definition_string_refOrdered, a_attribute_definition_string_ref)
	}
	sort.Slice(a_attribute_definition_string_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_string_refi := a_attribute_definition_string_refOrdered[i]
		a_attribute_definition_string_refj := a_attribute_definition_string_refOrdered[j]
		a_attribute_definition_string_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order[a_attribute_definition_string_refi]
		a_attribute_definition_string_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order[a_attribute_definition_string_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_string_refi_order < a_attribute_definition_string_refj_order
	})
	if len(a_attribute_definition_string_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_definition_string_ref := range a_attribute_definition_string_refOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_STRING_REF", int(stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order[a_attribute_definition_string_ref]), a_attribute_definition_string_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_STRING_REF_Identifiers[a_attribute_definition_string_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_STRING_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_string_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_string_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_STRING_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_string_ref.ATTRIBUTE_DEFINITION_STRING_REF))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_DEFINITION_XHTML_REF_Identifiers := make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]string)
	_ = map_A_ATTRIBUTE_DEFINITION_XHTML_REF_Identifiers

	a_attribute_definition_xhtml_refOrdered := []*A_ATTRIBUTE_DEFINITION_XHTML_REF{}
	for a_attribute_definition_xhtml_ref := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
		a_attribute_definition_xhtml_refOrdered = append(a_attribute_definition_xhtml_refOrdered, a_attribute_definition_xhtml_ref)
	}
	sort.Slice(a_attribute_definition_xhtml_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_xhtml_refi := a_attribute_definition_xhtml_refOrdered[i]
		a_attribute_definition_xhtml_refj := a_attribute_definition_xhtml_refOrdered[j]
		a_attribute_definition_xhtml_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order[a_attribute_definition_xhtml_refi]
		a_attribute_definition_xhtml_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order[a_attribute_definition_xhtml_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_xhtml_refi_order < a_attribute_definition_xhtml_refj_order
	})
	if len(a_attribute_definition_xhtml_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_definition_xhtml_ref := range a_attribute_definition_xhtml_refOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_XHTML_REF", int(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order[a_attribute_definition_xhtml_ref]), a_attribute_definition_xhtml_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_XHTML_REF_Identifiers[a_attribute_definition_xhtml_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_XHTML_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_xhtml_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_xhtml_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_XHTML_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_xhtml_ref.ATTRIBUTE_DEFINITION_XHTML_REF))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_VALUE_BOOLEAN_Identifiers := make(map[*A_ATTRIBUTE_VALUE_BOOLEAN]string)
	_ = map_A_ATTRIBUTE_VALUE_BOOLEAN_Identifiers

	a_attribute_value_booleanOrdered := []*A_ATTRIBUTE_VALUE_BOOLEAN{}
	for a_attribute_value_boolean := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		a_attribute_value_booleanOrdered = append(a_attribute_value_booleanOrdered, a_attribute_value_boolean)
	}
	sort.Slice(a_attribute_value_booleanOrdered[:], func(i, j int) bool {
		a_attribute_value_booleani := a_attribute_value_booleanOrdered[i]
		a_attribute_value_booleanj := a_attribute_value_booleanOrdered[j]
		a_attribute_value_booleani_order, oki := stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[a_attribute_value_booleani]
		a_attribute_value_booleanj_order, okj := stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[a_attribute_value_booleanj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_booleani_order < a_attribute_value_booleanj_order
	})
	if len(a_attribute_value_booleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_value_boolean := range a_attribute_value_booleanOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_BOOLEAN", int(stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[a_attribute_value_boolean]), a_attribute_value_boolean.Name)
		map_A_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[a_attribute_value_boolean] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_BOOLEAN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_boolean.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_value_boolean.Name))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_VALUE_DATE_Identifiers := make(map[*A_ATTRIBUTE_VALUE_DATE]string)
	_ = map_A_ATTRIBUTE_VALUE_DATE_Identifiers

	a_attribute_value_dateOrdered := []*A_ATTRIBUTE_VALUE_DATE{}
	for a_attribute_value_date := range stage.A_ATTRIBUTE_VALUE_DATEs {
		a_attribute_value_dateOrdered = append(a_attribute_value_dateOrdered, a_attribute_value_date)
	}
	sort.Slice(a_attribute_value_dateOrdered[:], func(i, j int) bool {
		a_attribute_value_datei := a_attribute_value_dateOrdered[i]
		a_attribute_value_datej := a_attribute_value_dateOrdered[j]
		a_attribute_value_datei_order, oki := stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order[a_attribute_value_datei]
		a_attribute_value_datej_order, okj := stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order[a_attribute_value_datej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_datei_order < a_attribute_value_datej_order
	})
	if len(a_attribute_value_dateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_value_date := range a_attribute_value_dateOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_DATE", int(stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order[a_attribute_value_date]), a_attribute_value_date.Name)
		map_A_ATTRIBUTE_VALUE_DATE_Identifiers[a_attribute_value_date] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_DATE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_date.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_value_date.Name))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_VALUE_ENUMERATION_Identifiers := make(map[*A_ATTRIBUTE_VALUE_ENUMERATION]string)
	_ = map_A_ATTRIBUTE_VALUE_ENUMERATION_Identifiers

	a_attribute_value_enumerationOrdered := []*A_ATTRIBUTE_VALUE_ENUMERATION{}
	for a_attribute_value_enumeration := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		a_attribute_value_enumerationOrdered = append(a_attribute_value_enumerationOrdered, a_attribute_value_enumeration)
	}
	sort.Slice(a_attribute_value_enumerationOrdered[:], func(i, j int) bool {
		a_attribute_value_enumerationi := a_attribute_value_enumerationOrdered[i]
		a_attribute_value_enumerationj := a_attribute_value_enumerationOrdered[j]
		a_attribute_value_enumerationi_order, oki := stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[a_attribute_value_enumerationi]
		a_attribute_value_enumerationj_order, okj := stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[a_attribute_value_enumerationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_enumerationi_order < a_attribute_value_enumerationj_order
	})
	if len(a_attribute_value_enumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_value_enumeration := range a_attribute_value_enumerationOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_ENUMERATION", int(stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[a_attribute_value_enumeration]), a_attribute_value_enumeration.Name)
		map_A_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[a_attribute_value_enumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_ENUMERATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_enumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_value_enumeration.Name))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_VALUE_INTEGER_Identifiers := make(map[*A_ATTRIBUTE_VALUE_INTEGER]string)
	_ = map_A_ATTRIBUTE_VALUE_INTEGER_Identifiers

	a_attribute_value_integerOrdered := []*A_ATTRIBUTE_VALUE_INTEGER{}
	for a_attribute_value_integer := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		a_attribute_value_integerOrdered = append(a_attribute_value_integerOrdered, a_attribute_value_integer)
	}
	sort.Slice(a_attribute_value_integerOrdered[:], func(i, j int) bool {
		a_attribute_value_integeri := a_attribute_value_integerOrdered[i]
		a_attribute_value_integerj := a_attribute_value_integerOrdered[j]
		a_attribute_value_integeri_order, oki := stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[a_attribute_value_integeri]
		a_attribute_value_integerj_order, okj := stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[a_attribute_value_integerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_integeri_order < a_attribute_value_integerj_order
	})
	if len(a_attribute_value_integerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_value_integer := range a_attribute_value_integerOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_INTEGER", int(stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[a_attribute_value_integer]), a_attribute_value_integer.Name)
		map_A_ATTRIBUTE_VALUE_INTEGER_Identifiers[a_attribute_value_integer] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_INTEGER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_integer.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_value_integer.Name))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_VALUE_REAL_Identifiers := make(map[*A_ATTRIBUTE_VALUE_REAL]string)
	_ = map_A_ATTRIBUTE_VALUE_REAL_Identifiers

	a_attribute_value_realOrdered := []*A_ATTRIBUTE_VALUE_REAL{}
	for a_attribute_value_real := range stage.A_ATTRIBUTE_VALUE_REALs {
		a_attribute_value_realOrdered = append(a_attribute_value_realOrdered, a_attribute_value_real)
	}
	sort.Slice(a_attribute_value_realOrdered[:], func(i, j int) bool {
		a_attribute_value_reali := a_attribute_value_realOrdered[i]
		a_attribute_value_realj := a_attribute_value_realOrdered[j]
		a_attribute_value_reali_order, oki := stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order[a_attribute_value_reali]
		a_attribute_value_realj_order, okj := stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order[a_attribute_value_realj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_reali_order < a_attribute_value_realj_order
	})
	if len(a_attribute_value_realOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_value_real := range a_attribute_value_realOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_REAL", int(stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order[a_attribute_value_real]), a_attribute_value_real.Name)
		map_A_ATTRIBUTE_VALUE_REAL_Identifiers[a_attribute_value_real] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_REAL")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_real.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_value_real.Name))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_VALUE_STRING_Identifiers := make(map[*A_ATTRIBUTE_VALUE_STRING]string)
	_ = map_A_ATTRIBUTE_VALUE_STRING_Identifiers

	a_attribute_value_stringOrdered := []*A_ATTRIBUTE_VALUE_STRING{}
	for a_attribute_value_string := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		a_attribute_value_stringOrdered = append(a_attribute_value_stringOrdered, a_attribute_value_string)
	}
	sort.Slice(a_attribute_value_stringOrdered[:], func(i, j int) bool {
		a_attribute_value_stringi := a_attribute_value_stringOrdered[i]
		a_attribute_value_stringj := a_attribute_value_stringOrdered[j]
		a_attribute_value_stringi_order, oki := stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order[a_attribute_value_stringi]
		a_attribute_value_stringj_order, okj := stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order[a_attribute_value_stringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_stringi_order < a_attribute_value_stringj_order
	})
	if len(a_attribute_value_stringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_value_string := range a_attribute_value_stringOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_STRING", int(stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order[a_attribute_value_string]), a_attribute_value_string.Name)
		map_A_ATTRIBUTE_VALUE_STRING_Identifiers[a_attribute_value_string] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_STRING")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_string.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_value_string.Name))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_VALUE_XHTML_Identifiers := make(map[*A_ATTRIBUTE_VALUE_XHTML]string)
	_ = map_A_ATTRIBUTE_VALUE_XHTML_Identifiers

	a_attribute_value_xhtmlOrdered := []*A_ATTRIBUTE_VALUE_XHTML{}
	for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		a_attribute_value_xhtmlOrdered = append(a_attribute_value_xhtmlOrdered, a_attribute_value_xhtml)
	}
	sort.Slice(a_attribute_value_xhtmlOrdered[:], func(i, j int) bool {
		a_attribute_value_xhtmli := a_attribute_value_xhtmlOrdered[i]
		a_attribute_value_xhtmlj := a_attribute_value_xhtmlOrdered[j]
		a_attribute_value_xhtmli_order, oki := stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[a_attribute_value_xhtmli]
		a_attribute_value_xhtmlj_order, okj := stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[a_attribute_value_xhtmlj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_xhtmli_order < a_attribute_value_xhtmlj_order
	})
	if len(a_attribute_value_xhtmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_value_xhtml := range a_attribute_value_xhtmlOrdered {

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_XHTML", int(stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[a_attribute_value_xhtml]), a_attribute_value_xhtml.Name)
		map_A_ATTRIBUTE_VALUE_XHTML_Identifiers[a_attribute_value_xhtml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_XHTML")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_xhtml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_value_xhtml.Name))
		initializerStatements += setValueField

	}

	map_A_ATTRIBUTE_VALUE_XHTML_1_Identifiers := make(map[*A_ATTRIBUTE_VALUE_XHTML_1]string)
	_ = map_A_ATTRIBUTE_VALUE_XHTML_1_Identifiers

	a_attribute_value_xhtml_1Ordered := []*A_ATTRIBUTE_VALUE_XHTML_1{}
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		a_attribute_value_xhtml_1Ordered = append(a_attribute_value_xhtml_1Ordered, a_attribute_value_xhtml_1)
	}
	sort.Slice(a_attribute_value_xhtml_1Ordered[:], func(i, j int) bool {
		a_attribute_value_xhtml_1i := a_attribute_value_xhtml_1Ordered[i]
		a_attribute_value_xhtml_1j := a_attribute_value_xhtml_1Ordered[j]
		a_attribute_value_xhtml_1i_order, oki := stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order[a_attribute_value_xhtml_1i]
		a_attribute_value_xhtml_1j_order, okj := stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order[a_attribute_value_xhtml_1j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_xhtml_1i_order < a_attribute_value_xhtml_1j_order
	})
	if len(a_attribute_value_xhtml_1Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_attribute_value_xhtml_1 := range a_attribute_value_xhtml_1Ordered {

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_XHTML_1", int(stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order[a_attribute_value_xhtml_1]), a_attribute_value_xhtml_1.Name)
		map_A_ATTRIBUTE_VALUE_XHTML_1_Identifiers[a_attribute_value_xhtml_1] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_XHTML_1")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_xhtml_1.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_attribute_value_xhtml_1.Name))
		initializerStatements += setValueField

	}

	map_A_CHILDREN_Identifiers := make(map[*A_CHILDREN]string)
	_ = map_A_CHILDREN_Identifiers

	a_childrenOrdered := []*A_CHILDREN{}
	for a_children := range stage.A_CHILDRENs {
		a_childrenOrdered = append(a_childrenOrdered, a_children)
	}
	sort.Slice(a_childrenOrdered[:], func(i, j int) bool {
		a_childreni := a_childrenOrdered[i]
		a_childrenj := a_childrenOrdered[j]
		a_childreni_order, oki := stage.A_CHILDRENMap_Staged_Order[a_childreni]
		a_childrenj_order, okj := stage.A_CHILDRENMap_Staged_Order[a_childrenj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_childreni_order < a_childrenj_order
	})
	if len(a_childrenOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_children := range a_childrenOrdered {

		id = generatesIdentifier("A_CHILDREN", int(stage.A_CHILDRENMap_Staged_Order[a_children]), a_children.Name)
		map_A_CHILDREN_Identifiers[a_children] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_CHILDREN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_children.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_children.Name))
		initializerStatements += setValueField

	}

	map_A_CORE_CONTENT_Identifiers := make(map[*A_CORE_CONTENT]string)
	_ = map_A_CORE_CONTENT_Identifiers

	a_core_contentOrdered := []*A_CORE_CONTENT{}
	for a_core_content := range stage.A_CORE_CONTENTs {
		a_core_contentOrdered = append(a_core_contentOrdered, a_core_content)
	}
	sort.Slice(a_core_contentOrdered[:], func(i, j int) bool {
		a_core_contenti := a_core_contentOrdered[i]
		a_core_contentj := a_core_contentOrdered[j]
		a_core_contenti_order, oki := stage.A_CORE_CONTENTMap_Staged_Order[a_core_contenti]
		a_core_contentj_order, okj := stage.A_CORE_CONTENTMap_Staged_Order[a_core_contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_core_contenti_order < a_core_contentj_order
	})
	if len(a_core_contentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_core_content := range a_core_contentOrdered {

		id = generatesIdentifier("A_CORE_CONTENT", int(stage.A_CORE_CONTENTMap_Staged_Order[a_core_content]), a_core_content.Name)
		map_A_CORE_CONTENT_Identifiers[a_core_content] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_CORE_CONTENT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_core_content.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_core_content.Name))
		initializerStatements += setValueField

	}

	map_A_DATATYPES_Identifiers := make(map[*A_DATATYPES]string)
	_ = map_A_DATATYPES_Identifiers

	a_datatypesOrdered := []*A_DATATYPES{}
	for a_datatypes := range stage.A_DATATYPESs {
		a_datatypesOrdered = append(a_datatypesOrdered, a_datatypes)
	}
	sort.Slice(a_datatypesOrdered[:], func(i, j int) bool {
		a_datatypesi := a_datatypesOrdered[i]
		a_datatypesj := a_datatypesOrdered[j]
		a_datatypesi_order, oki := stage.A_DATATYPESMap_Staged_Order[a_datatypesi]
		a_datatypesj_order, okj := stage.A_DATATYPESMap_Staged_Order[a_datatypesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatypesi_order < a_datatypesj_order
	})
	if len(a_datatypesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_datatypes := range a_datatypesOrdered {

		id = generatesIdentifier("A_DATATYPES", int(stage.A_DATATYPESMap_Staged_Order[a_datatypes]), a_datatypes.Name)
		map_A_DATATYPES_Identifiers[a_datatypes] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatypes.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatypes.Name))
		initializerStatements += setValueField

	}

	map_A_DATATYPE_DEFINITION_BOOLEAN_REF_Identifiers := make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]string)
	_ = map_A_DATATYPE_DEFINITION_BOOLEAN_REF_Identifiers

	a_datatype_definition_boolean_refOrdered := []*A_DATATYPE_DEFINITION_BOOLEAN_REF{}
	for a_datatype_definition_boolean_ref := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
		a_datatype_definition_boolean_refOrdered = append(a_datatype_definition_boolean_refOrdered, a_datatype_definition_boolean_ref)
	}
	sort.Slice(a_datatype_definition_boolean_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_boolean_refi := a_datatype_definition_boolean_refOrdered[i]
		a_datatype_definition_boolean_refj := a_datatype_definition_boolean_refOrdered[j]
		a_datatype_definition_boolean_refi_order, oki := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_datatype_definition_boolean_refi]
		a_datatype_definition_boolean_refj_order, okj := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_datatype_definition_boolean_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_boolean_refi_order < a_datatype_definition_boolean_refj_order
	})
	if len(a_datatype_definition_boolean_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_datatype_definition_boolean_ref := range a_datatype_definition_boolean_refOrdered {

		id = generatesIdentifier("A_DATATYPE_DEFINITION_BOOLEAN_REF", int(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_datatype_definition_boolean_ref]), a_datatype_definition_boolean_ref.Name)
		map_A_DATATYPE_DEFINITION_BOOLEAN_REF_Identifiers[a_datatype_definition_boolean_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_BOOLEAN_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_boolean_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_boolean_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_BOOLEAN_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_boolean_ref.DATATYPE_DEFINITION_BOOLEAN_REF))
		initializerStatements += setValueField

	}

	map_A_DATATYPE_DEFINITION_DATE_REF_Identifiers := make(map[*A_DATATYPE_DEFINITION_DATE_REF]string)
	_ = map_A_DATATYPE_DEFINITION_DATE_REF_Identifiers

	a_datatype_definition_date_refOrdered := []*A_DATATYPE_DEFINITION_DATE_REF{}
	for a_datatype_definition_date_ref := range stage.A_DATATYPE_DEFINITION_DATE_REFs {
		a_datatype_definition_date_refOrdered = append(a_datatype_definition_date_refOrdered, a_datatype_definition_date_ref)
	}
	sort.Slice(a_datatype_definition_date_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_date_refi := a_datatype_definition_date_refOrdered[i]
		a_datatype_definition_date_refj := a_datatype_definition_date_refOrdered[j]
		a_datatype_definition_date_refi_order, oki := stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order[a_datatype_definition_date_refi]
		a_datatype_definition_date_refj_order, okj := stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order[a_datatype_definition_date_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_date_refi_order < a_datatype_definition_date_refj_order
	})
	if len(a_datatype_definition_date_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_datatype_definition_date_ref := range a_datatype_definition_date_refOrdered {

		id = generatesIdentifier("A_DATATYPE_DEFINITION_DATE_REF", int(stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order[a_datatype_definition_date_ref]), a_datatype_definition_date_ref.Name)
		map_A_DATATYPE_DEFINITION_DATE_REF_Identifiers[a_datatype_definition_date_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_DATE_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_date_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_date_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_DATE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_date_ref.DATATYPE_DEFINITION_DATE_REF))
		initializerStatements += setValueField

	}

	map_A_DATATYPE_DEFINITION_ENUMERATION_REF_Identifiers := make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]string)
	_ = map_A_DATATYPE_DEFINITION_ENUMERATION_REF_Identifiers

	a_datatype_definition_enumeration_refOrdered := []*A_DATATYPE_DEFINITION_ENUMERATION_REF{}
	for a_datatype_definition_enumeration_ref := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
		a_datatype_definition_enumeration_refOrdered = append(a_datatype_definition_enumeration_refOrdered, a_datatype_definition_enumeration_ref)
	}
	sort.Slice(a_datatype_definition_enumeration_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_enumeration_refi := a_datatype_definition_enumeration_refOrdered[i]
		a_datatype_definition_enumeration_refj := a_datatype_definition_enumeration_refOrdered[j]
		a_datatype_definition_enumeration_refi_order, oki := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_datatype_definition_enumeration_refi]
		a_datatype_definition_enumeration_refj_order, okj := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_datatype_definition_enumeration_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_enumeration_refi_order < a_datatype_definition_enumeration_refj_order
	})
	if len(a_datatype_definition_enumeration_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_datatype_definition_enumeration_ref := range a_datatype_definition_enumeration_refOrdered {

		id = generatesIdentifier("A_DATATYPE_DEFINITION_ENUMERATION_REF", int(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_datatype_definition_enumeration_ref]), a_datatype_definition_enumeration_ref.Name)
		map_A_DATATYPE_DEFINITION_ENUMERATION_REF_Identifiers[a_datatype_definition_enumeration_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_ENUMERATION_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_enumeration_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_enumeration_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_ENUMERATION_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_enumeration_ref.DATATYPE_DEFINITION_ENUMERATION_REF))
		initializerStatements += setValueField

	}

	map_A_DATATYPE_DEFINITION_INTEGER_REF_Identifiers := make(map[*A_DATATYPE_DEFINITION_INTEGER_REF]string)
	_ = map_A_DATATYPE_DEFINITION_INTEGER_REF_Identifiers

	a_datatype_definition_integer_refOrdered := []*A_DATATYPE_DEFINITION_INTEGER_REF{}
	for a_datatype_definition_integer_ref := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
		a_datatype_definition_integer_refOrdered = append(a_datatype_definition_integer_refOrdered, a_datatype_definition_integer_ref)
	}
	sort.Slice(a_datatype_definition_integer_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_integer_refi := a_datatype_definition_integer_refOrdered[i]
		a_datatype_definition_integer_refj := a_datatype_definition_integer_refOrdered[j]
		a_datatype_definition_integer_refi_order, oki := stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order[a_datatype_definition_integer_refi]
		a_datatype_definition_integer_refj_order, okj := stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order[a_datatype_definition_integer_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_integer_refi_order < a_datatype_definition_integer_refj_order
	})
	if len(a_datatype_definition_integer_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_datatype_definition_integer_ref := range a_datatype_definition_integer_refOrdered {

		id = generatesIdentifier("A_DATATYPE_DEFINITION_INTEGER_REF", int(stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order[a_datatype_definition_integer_ref]), a_datatype_definition_integer_ref.Name)
		map_A_DATATYPE_DEFINITION_INTEGER_REF_Identifiers[a_datatype_definition_integer_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_INTEGER_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_integer_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_integer_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_INTEGER_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_integer_ref.DATATYPE_DEFINITION_INTEGER_REF))
		initializerStatements += setValueField

	}

	map_A_DATATYPE_DEFINITION_REAL_REF_Identifiers := make(map[*A_DATATYPE_DEFINITION_REAL_REF]string)
	_ = map_A_DATATYPE_DEFINITION_REAL_REF_Identifiers

	a_datatype_definition_real_refOrdered := []*A_DATATYPE_DEFINITION_REAL_REF{}
	for a_datatype_definition_real_ref := range stage.A_DATATYPE_DEFINITION_REAL_REFs {
		a_datatype_definition_real_refOrdered = append(a_datatype_definition_real_refOrdered, a_datatype_definition_real_ref)
	}
	sort.Slice(a_datatype_definition_real_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_real_refi := a_datatype_definition_real_refOrdered[i]
		a_datatype_definition_real_refj := a_datatype_definition_real_refOrdered[j]
		a_datatype_definition_real_refi_order, oki := stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order[a_datatype_definition_real_refi]
		a_datatype_definition_real_refj_order, okj := stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order[a_datatype_definition_real_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_real_refi_order < a_datatype_definition_real_refj_order
	})
	if len(a_datatype_definition_real_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_datatype_definition_real_ref := range a_datatype_definition_real_refOrdered {

		id = generatesIdentifier("A_DATATYPE_DEFINITION_REAL_REF", int(stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order[a_datatype_definition_real_ref]), a_datatype_definition_real_ref.Name)
		map_A_DATATYPE_DEFINITION_REAL_REF_Identifiers[a_datatype_definition_real_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_REAL_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_real_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_real_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_REAL_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_real_ref.DATATYPE_DEFINITION_REAL_REF))
		initializerStatements += setValueField

	}

	map_A_DATATYPE_DEFINITION_STRING_REF_Identifiers := make(map[*A_DATATYPE_DEFINITION_STRING_REF]string)
	_ = map_A_DATATYPE_DEFINITION_STRING_REF_Identifiers

	a_datatype_definition_string_refOrdered := []*A_DATATYPE_DEFINITION_STRING_REF{}
	for a_datatype_definition_string_ref := range stage.A_DATATYPE_DEFINITION_STRING_REFs {
		a_datatype_definition_string_refOrdered = append(a_datatype_definition_string_refOrdered, a_datatype_definition_string_ref)
	}
	sort.Slice(a_datatype_definition_string_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_string_refi := a_datatype_definition_string_refOrdered[i]
		a_datatype_definition_string_refj := a_datatype_definition_string_refOrdered[j]
		a_datatype_definition_string_refi_order, oki := stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order[a_datatype_definition_string_refi]
		a_datatype_definition_string_refj_order, okj := stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order[a_datatype_definition_string_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_string_refi_order < a_datatype_definition_string_refj_order
	})
	if len(a_datatype_definition_string_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_datatype_definition_string_ref := range a_datatype_definition_string_refOrdered {

		id = generatesIdentifier("A_DATATYPE_DEFINITION_STRING_REF", int(stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order[a_datatype_definition_string_ref]), a_datatype_definition_string_ref.Name)
		map_A_DATATYPE_DEFINITION_STRING_REF_Identifiers[a_datatype_definition_string_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_STRING_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_string_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_string_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_STRING_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_string_ref.DATATYPE_DEFINITION_STRING_REF))
		initializerStatements += setValueField

	}

	map_A_DATATYPE_DEFINITION_XHTML_REF_Identifiers := make(map[*A_DATATYPE_DEFINITION_XHTML_REF]string)
	_ = map_A_DATATYPE_DEFINITION_XHTML_REF_Identifiers

	a_datatype_definition_xhtml_refOrdered := []*A_DATATYPE_DEFINITION_XHTML_REF{}
	for a_datatype_definition_xhtml_ref := range stage.A_DATATYPE_DEFINITION_XHTML_REFs {
		a_datatype_definition_xhtml_refOrdered = append(a_datatype_definition_xhtml_refOrdered, a_datatype_definition_xhtml_ref)
	}
	sort.Slice(a_datatype_definition_xhtml_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_xhtml_refi := a_datatype_definition_xhtml_refOrdered[i]
		a_datatype_definition_xhtml_refj := a_datatype_definition_xhtml_refOrdered[j]
		a_datatype_definition_xhtml_refi_order, oki := stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order[a_datatype_definition_xhtml_refi]
		a_datatype_definition_xhtml_refj_order, okj := stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order[a_datatype_definition_xhtml_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_xhtml_refi_order < a_datatype_definition_xhtml_refj_order
	})
	if len(a_datatype_definition_xhtml_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_datatype_definition_xhtml_ref := range a_datatype_definition_xhtml_refOrdered {

		id = generatesIdentifier("A_DATATYPE_DEFINITION_XHTML_REF", int(stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order[a_datatype_definition_xhtml_ref]), a_datatype_definition_xhtml_ref.Name)
		map_A_DATATYPE_DEFINITION_XHTML_REF_Identifiers[a_datatype_definition_xhtml_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_XHTML_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_xhtml_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_xhtml_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_XHTML_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_xhtml_ref.DATATYPE_DEFINITION_XHTML_REF))
		initializerStatements += setValueField

	}

	map_A_EDITABLE_ATTS_Identifiers := make(map[*A_EDITABLE_ATTS]string)
	_ = map_A_EDITABLE_ATTS_Identifiers

	a_editable_attsOrdered := []*A_EDITABLE_ATTS{}
	for a_editable_atts := range stage.A_EDITABLE_ATTSs {
		a_editable_attsOrdered = append(a_editable_attsOrdered, a_editable_atts)
	}
	sort.Slice(a_editable_attsOrdered[:], func(i, j int) bool {
		a_editable_attsi := a_editable_attsOrdered[i]
		a_editable_attsj := a_editable_attsOrdered[j]
		a_editable_attsi_order, oki := stage.A_EDITABLE_ATTSMap_Staged_Order[a_editable_attsi]
		a_editable_attsj_order, okj := stage.A_EDITABLE_ATTSMap_Staged_Order[a_editable_attsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_editable_attsi_order < a_editable_attsj_order
	})
	if len(a_editable_attsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_editable_atts := range a_editable_attsOrdered {

		id = generatesIdentifier("A_EDITABLE_ATTS", int(stage.A_EDITABLE_ATTSMap_Staged_Order[a_editable_atts]), a_editable_atts.Name)
		map_A_EDITABLE_ATTS_Identifiers[a_editable_atts] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_EDITABLE_ATTS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_editable_atts.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_BOOLEAN_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_BOOLEAN_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_DATE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_DATE_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_ENUMERATION_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_ENUMERATION_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_INTEGER_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_INTEGER_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_REAL_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_REAL_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_STRING_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_STRING_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_XHTML_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_XHTML_REF))
		initializerStatements += setValueField

	}

	map_A_ENUM_VALUE_REF_Identifiers := make(map[*A_ENUM_VALUE_REF]string)
	_ = map_A_ENUM_VALUE_REF_Identifiers

	a_enum_value_refOrdered := []*A_ENUM_VALUE_REF{}
	for a_enum_value_ref := range stage.A_ENUM_VALUE_REFs {
		a_enum_value_refOrdered = append(a_enum_value_refOrdered, a_enum_value_ref)
	}
	sort.Slice(a_enum_value_refOrdered[:], func(i, j int) bool {
		a_enum_value_refi := a_enum_value_refOrdered[i]
		a_enum_value_refj := a_enum_value_refOrdered[j]
		a_enum_value_refi_order, oki := stage.A_ENUM_VALUE_REFMap_Staged_Order[a_enum_value_refi]
		a_enum_value_refj_order, okj := stage.A_ENUM_VALUE_REFMap_Staged_Order[a_enum_value_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_enum_value_refi_order < a_enum_value_refj_order
	})
	if len(a_enum_value_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_enum_value_ref := range a_enum_value_refOrdered {

		id = generatesIdentifier("A_ENUM_VALUE_REF", int(stage.A_ENUM_VALUE_REFMap_Staged_Order[a_enum_value_ref]), a_enum_value_ref.Name)
		map_A_ENUM_VALUE_REF_Identifiers[a_enum_value_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ENUM_VALUE_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_enum_value_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_enum_value_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ENUM_VALUE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_enum_value_ref.ENUM_VALUE_REF))
		initializerStatements += setValueField

	}

	map_A_OBJECT_Identifiers := make(map[*A_OBJECT]string)
	_ = map_A_OBJECT_Identifiers

	a_objectOrdered := []*A_OBJECT{}
	for a_object := range stage.A_OBJECTs {
		a_objectOrdered = append(a_objectOrdered, a_object)
	}
	sort.Slice(a_objectOrdered[:], func(i, j int) bool {
		a_objecti := a_objectOrdered[i]
		a_objectj := a_objectOrdered[j]
		a_objecti_order, oki := stage.A_OBJECTMap_Staged_Order[a_objecti]
		a_objectj_order, okj := stage.A_OBJECTMap_Staged_Order[a_objectj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_objecti_order < a_objectj_order
	})
	if len(a_objectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_object := range a_objectOrdered {

		id = generatesIdentifier("A_OBJECT", int(stage.A_OBJECTMap_Staged_Order[a_object]), a_object.Name)
		map_A_OBJECT_Identifiers[a_object] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_OBJECT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_object.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_object.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPEC_OBJECT_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_object.SPEC_OBJECT_REF))
		initializerStatements += setValueField

	}

	map_A_PROPERTIES_Identifiers := make(map[*A_PROPERTIES]string)
	_ = map_A_PROPERTIES_Identifiers

	a_propertiesOrdered := []*A_PROPERTIES{}
	for a_properties := range stage.A_PROPERTIESs {
		a_propertiesOrdered = append(a_propertiesOrdered, a_properties)
	}
	sort.Slice(a_propertiesOrdered[:], func(i, j int) bool {
		a_propertiesi := a_propertiesOrdered[i]
		a_propertiesj := a_propertiesOrdered[j]
		a_propertiesi_order, oki := stage.A_PROPERTIESMap_Staged_Order[a_propertiesi]
		a_propertiesj_order, okj := stage.A_PROPERTIESMap_Staged_Order[a_propertiesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_propertiesi_order < a_propertiesj_order
	})
	if len(a_propertiesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_properties := range a_propertiesOrdered {

		id = generatesIdentifier("A_PROPERTIES", int(stage.A_PROPERTIESMap_Staged_Order[a_properties]), a_properties.Name)
		map_A_PROPERTIES_Identifiers[a_properties] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_PROPERTIES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_properties.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_properties.Name))
		initializerStatements += setValueField

	}

	map_A_RELATION_GROUP_TYPE_REF_Identifiers := make(map[*A_RELATION_GROUP_TYPE_REF]string)
	_ = map_A_RELATION_GROUP_TYPE_REF_Identifiers

	a_relation_group_type_refOrdered := []*A_RELATION_GROUP_TYPE_REF{}
	for a_relation_group_type_ref := range stage.A_RELATION_GROUP_TYPE_REFs {
		a_relation_group_type_refOrdered = append(a_relation_group_type_refOrdered, a_relation_group_type_ref)
	}
	sort.Slice(a_relation_group_type_refOrdered[:], func(i, j int) bool {
		a_relation_group_type_refi := a_relation_group_type_refOrdered[i]
		a_relation_group_type_refj := a_relation_group_type_refOrdered[j]
		a_relation_group_type_refi_order, oki := stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_refi]
		a_relation_group_type_refj_order, okj := stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_relation_group_type_refi_order < a_relation_group_type_refj_order
	})
	if len(a_relation_group_type_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_relation_group_type_ref := range a_relation_group_type_refOrdered {

		id = generatesIdentifier("A_RELATION_GROUP_TYPE_REF", int(stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_ref]), a_relation_group_type_ref.Name)
		map_A_RELATION_GROUP_TYPE_REF_Identifiers[a_relation_group_type_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_RELATION_GROUP_TYPE_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_relation_group_type_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_relation_group_type_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RELATION_GROUP_TYPE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_relation_group_type_ref.RELATION_GROUP_TYPE_REF))
		initializerStatements += setValueField

	}

	map_A_SOURCE_1_Identifiers := make(map[*A_SOURCE_1]string)
	_ = map_A_SOURCE_1_Identifiers

	a_source_1Ordered := []*A_SOURCE_1{}
	for a_source_1 := range stage.A_SOURCE_1s {
		a_source_1Ordered = append(a_source_1Ordered, a_source_1)
	}
	sort.Slice(a_source_1Ordered[:], func(i, j int) bool {
		a_source_1i := a_source_1Ordered[i]
		a_source_1j := a_source_1Ordered[j]
		a_source_1i_order, oki := stage.A_SOURCE_1Map_Staged_Order[a_source_1i]
		a_source_1j_order, okj := stage.A_SOURCE_1Map_Staged_Order[a_source_1j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_source_1i_order < a_source_1j_order
	})
	if len(a_source_1Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_source_1 := range a_source_1Ordered {

		id = generatesIdentifier("A_SOURCE_1", int(stage.A_SOURCE_1Map_Staged_Order[a_source_1]), a_source_1.Name)
		map_A_SOURCE_1_Identifiers[a_source_1] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SOURCE_1")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_source_1.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_source_1.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPEC_OBJECT_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_source_1.SPEC_OBJECT_REF))
		initializerStatements += setValueField

	}

	map_A_SOURCE_SPECIFICATION_1_Identifiers := make(map[*A_SOURCE_SPECIFICATION_1]string)
	_ = map_A_SOURCE_SPECIFICATION_1_Identifiers

	a_source_specification_1Ordered := []*A_SOURCE_SPECIFICATION_1{}
	for a_source_specification_1 := range stage.A_SOURCE_SPECIFICATION_1s {
		a_source_specification_1Ordered = append(a_source_specification_1Ordered, a_source_specification_1)
	}
	sort.Slice(a_source_specification_1Ordered[:], func(i, j int) bool {
		a_source_specification_1i := a_source_specification_1Ordered[i]
		a_source_specification_1j := a_source_specification_1Ordered[j]
		a_source_specification_1i_order, oki := stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order[a_source_specification_1i]
		a_source_specification_1j_order, okj := stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order[a_source_specification_1j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_source_specification_1i_order < a_source_specification_1j_order
	})
	if len(a_source_specification_1Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_source_specification_1 := range a_source_specification_1Ordered {

		id = generatesIdentifier("A_SOURCE_SPECIFICATION_1", int(stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order[a_source_specification_1]), a_source_specification_1.Name)
		map_A_SOURCE_SPECIFICATION_1_Identifiers[a_source_specification_1] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SOURCE_SPECIFICATION_1")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_source_specification_1.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_source_specification_1.Name))
		initializerStatements += setValueField

		if a_source_specification_1.SPECIFICATION_REF != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPECIFICATION_REF")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+a_source_specification_1.SPECIFICATION_REF.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_A_SPECIFICATIONS_Identifiers := make(map[*A_SPECIFICATIONS]string)
	_ = map_A_SPECIFICATIONS_Identifiers

	a_specificationsOrdered := []*A_SPECIFICATIONS{}
	for a_specifications := range stage.A_SPECIFICATIONSs {
		a_specificationsOrdered = append(a_specificationsOrdered, a_specifications)
	}
	sort.Slice(a_specificationsOrdered[:], func(i, j int) bool {
		a_specificationsi := a_specificationsOrdered[i]
		a_specificationsj := a_specificationsOrdered[j]
		a_specificationsi_order, oki := stage.A_SPECIFICATIONSMap_Staged_Order[a_specificationsi]
		a_specificationsj_order, okj := stage.A_SPECIFICATIONSMap_Staged_Order[a_specificationsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_specificationsi_order < a_specificationsj_order
	})
	if len(a_specificationsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_specifications := range a_specificationsOrdered {

		id = generatesIdentifier("A_SPECIFICATIONS", int(stage.A_SPECIFICATIONSMap_Staged_Order[a_specifications]), a_specifications.Name)
		map_A_SPECIFICATIONS_Identifiers[a_specifications] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFICATIONS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_specifications.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_specifications.Name))
		initializerStatements += setValueField

	}

	map_A_SPECIFICATION_TYPE_REF_Identifiers := make(map[*A_SPECIFICATION_TYPE_REF]string)
	_ = map_A_SPECIFICATION_TYPE_REF_Identifiers

	a_specification_type_refOrdered := []*A_SPECIFICATION_TYPE_REF{}
	for a_specification_type_ref := range stage.A_SPECIFICATION_TYPE_REFs {
		a_specification_type_refOrdered = append(a_specification_type_refOrdered, a_specification_type_ref)
	}
	sort.Slice(a_specification_type_refOrdered[:], func(i, j int) bool {
		a_specification_type_refi := a_specification_type_refOrdered[i]
		a_specification_type_refj := a_specification_type_refOrdered[j]
		a_specification_type_refi_order, oki := stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order[a_specification_type_refi]
		a_specification_type_refj_order, okj := stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order[a_specification_type_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_specification_type_refi_order < a_specification_type_refj_order
	})
	if len(a_specification_type_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_specification_type_ref := range a_specification_type_refOrdered {

		id = generatesIdentifier("A_SPECIFICATION_TYPE_REF", int(stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order[a_specification_type_ref]), a_specification_type_ref.Name)
		map_A_SPECIFICATION_TYPE_REF_Identifiers[a_specification_type_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFICATION_TYPE_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_specification_type_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_specification_type_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPECIFICATION_TYPE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_specification_type_ref.SPECIFICATION_TYPE_REF))
		initializerStatements += setValueField

	}

	map_A_SPECIFIED_VALUES_Identifiers := make(map[*A_SPECIFIED_VALUES]string)
	_ = map_A_SPECIFIED_VALUES_Identifiers

	a_specified_valuesOrdered := []*A_SPECIFIED_VALUES{}
	for a_specified_values := range stage.A_SPECIFIED_VALUESs {
		a_specified_valuesOrdered = append(a_specified_valuesOrdered, a_specified_values)
	}
	sort.Slice(a_specified_valuesOrdered[:], func(i, j int) bool {
		a_specified_valuesi := a_specified_valuesOrdered[i]
		a_specified_valuesj := a_specified_valuesOrdered[j]
		a_specified_valuesi_order, oki := stage.A_SPECIFIED_VALUESMap_Staged_Order[a_specified_valuesi]
		a_specified_valuesj_order, okj := stage.A_SPECIFIED_VALUESMap_Staged_Order[a_specified_valuesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_specified_valuesi_order < a_specified_valuesj_order
	})
	if len(a_specified_valuesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_specified_values := range a_specified_valuesOrdered {

		id = generatesIdentifier("A_SPECIFIED_VALUES", int(stage.A_SPECIFIED_VALUESMap_Staged_Order[a_specified_values]), a_specified_values.Name)
		map_A_SPECIFIED_VALUES_Identifiers[a_specified_values] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFIED_VALUES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_specified_values.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_specified_values.Name))
		initializerStatements += setValueField

	}

	map_A_SPEC_ATTRIBUTES_Identifiers := make(map[*A_SPEC_ATTRIBUTES]string)
	_ = map_A_SPEC_ATTRIBUTES_Identifiers

	a_spec_attributesOrdered := []*A_SPEC_ATTRIBUTES{}
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		a_spec_attributesOrdered = append(a_spec_attributesOrdered, a_spec_attributes)
	}
	sort.Slice(a_spec_attributesOrdered[:], func(i, j int) bool {
		a_spec_attributesi := a_spec_attributesOrdered[i]
		a_spec_attributesj := a_spec_attributesOrdered[j]
		a_spec_attributesi_order, oki := stage.A_SPEC_ATTRIBUTESMap_Staged_Order[a_spec_attributesi]
		a_spec_attributesj_order, okj := stage.A_SPEC_ATTRIBUTESMap_Staged_Order[a_spec_attributesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_attributesi_order < a_spec_attributesj_order
	})
	if len(a_spec_attributesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_spec_attributes := range a_spec_attributesOrdered {

		id = generatesIdentifier("A_SPEC_ATTRIBUTES", int(stage.A_SPEC_ATTRIBUTESMap_Staged_Order[a_spec_attributes]), a_spec_attributes.Name)
		map_A_SPEC_ATTRIBUTES_Identifiers[a_spec_attributes] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_ATTRIBUTES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_attributes.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_attributes.Name))
		initializerStatements += setValueField

	}

	map_A_SPEC_OBJECTS_Identifiers := make(map[*A_SPEC_OBJECTS]string)
	_ = map_A_SPEC_OBJECTS_Identifiers

	a_spec_objectsOrdered := []*A_SPEC_OBJECTS{}
	for a_spec_objects := range stage.A_SPEC_OBJECTSs {
		a_spec_objectsOrdered = append(a_spec_objectsOrdered, a_spec_objects)
	}
	sort.Slice(a_spec_objectsOrdered[:], func(i, j int) bool {
		a_spec_objectsi := a_spec_objectsOrdered[i]
		a_spec_objectsj := a_spec_objectsOrdered[j]
		a_spec_objectsi_order, oki := stage.A_SPEC_OBJECTSMap_Staged_Order[a_spec_objectsi]
		a_spec_objectsj_order, okj := stage.A_SPEC_OBJECTSMap_Staged_Order[a_spec_objectsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_objectsi_order < a_spec_objectsj_order
	})
	if len(a_spec_objectsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_spec_objects := range a_spec_objectsOrdered {

		id = generatesIdentifier("A_SPEC_OBJECTS", int(stage.A_SPEC_OBJECTSMap_Staged_Order[a_spec_objects]), a_spec_objects.Name)
		map_A_SPEC_OBJECTS_Identifiers[a_spec_objects] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_OBJECTS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_objects.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_objects.Name))
		initializerStatements += setValueField

	}

	map_A_SPEC_OBJECT_TYPE_REF_Identifiers := make(map[*A_SPEC_OBJECT_TYPE_REF]string)
	_ = map_A_SPEC_OBJECT_TYPE_REF_Identifiers

	a_spec_object_type_refOrdered := []*A_SPEC_OBJECT_TYPE_REF{}
	for a_spec_object_type_ref := range stage.A_SPEC_OBJECT_TYPE_REFs {
		a_spec_object_type_refOrdered = append(a_spec_object_type_refOrdered, a_spec_object_type_ref)
	}
	sort.Slice(a_spec_object_type_refOrdered[:], func(i, j int) bool {
		a_spec_object_type_refi := a_spec_object_type_refOrdered[i]
		a_spec_object_type_refj := a_spec_object_type_refOrdered[j]
		a_spec_object_type_refi_order, oki := stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order[a_spec_object_type_refi]
		a_spec_object_type_refj_order, okj := stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order[a_spec_object_type_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_object_type_refi_order < a_spec_object_type_refj_order
	})
	if len(a_spec_object_type_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_spec_object_type_ref := range a_spec_object_type_refOrdered {

		id = generatesIdentifier("A_SPEC_OBJECT_TYPE_REF", int(stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order[a_spec_object_type_ref]), a_spec_object_type_ref.Name)
		map_A_SPEC_OBJECT_TYPE_REF_Identifiers[a_spec_object_type_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_OBJECT_TYPE_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_object_type_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_object_type_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPEC_OBJECT_TYPE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_object_type_ref.SPEC_OBJECT_TYPE_REF))
		initializerStatements += setValueField

	}

	map_A_SPEC_RELATIONS_Identifiers := make(map[*A_SPEC_RELATIONS]string)
	_ = map_A_SPEC_RELATIONS_Identifiers

	a_spec_relationsOrdered := []*A_SPEC_RELATIONS{}
	for a_spec_relations := range stage.A_SPEC_RELATIONSs {
		a_spec_relationsOrdered = append(a_spec_relationsOrdered, a_spec_relations)
	}
	sort.Slice(a_spec_relationsOrdered[:], func(i, j int) bool {
		a_spec_relationsi := a_spec_relationsOrdered[i]
		a_spec_relationsj := a_spec_relationsOrdered[j]
		a_spec_relationsi_order, oki := stage.A_SPEC_RELATIONSMap_Staged_Order[a_spec_relationsi]
		a_spec_relationsj_order, okj := stage.A_SPEC_RELATIONSMap_Staged_Order[a_spec_relationsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_relationsi_order < a_spec_relationsj_order
	})
	if len(a_spec_relationsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_spec_relations := range a_spec_relationsOrdered {

		id = generatesIdentifier("A_SPEC_RELATIONS", int(stage.A_SPEC_RELATIONSMap_Staged_Order[a_spec_relations]), a_spec_relations.Name)
		map_A_SPEC_RELATIONS_Identifiers[a_spec_relations] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATIONS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relations.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_relations.Name))
		initializerStatements += setValueField

	}

	map_A_SPEC_RELATION_GROUPS_Identifiers := make(map[*A_SPEC_RELATION_GROUPS]string)
	_ = map_A_SPEC_RELATION_GROUPS_Identifiers

	a_spec_relation_groupsOrdered := []*A_SPEC_RELATION_GROUPS{}
	for a_spec_relation_groups := range stage.A_SPEC_RELATION_GROUPSs {
		a_spec_relation_groupsOrdered = append(a_spec_relation_groupsOrdered, a_spec_relation_groups)
	}
	sort.Slice(a_spec_relation_groupsOrdered[:], func(i, j int) bool {
		a_spec_relation_groupsi := a_spec_relation_groupsOrdered[i]
		a_spec_relation_groupsj := a_spec_relation_groupsOrdered[j]
		a_spec_relation_groupsi_order, oki := stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groupsi]
		a_spec_relation_groupsj_order, okj := stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groupsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_relation_groupsi_order < a_spec_relation_groupsj_order
	})
	if len(a_spec_relation_groupsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_spec_relation_groups := range a_spec_relation_groupsOrdered {

		id = generatesIdentifier("A_SPEC_RELATION_GROUPS", int(stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groups]), a_spec_relation_groups.Name)
		map_A_SPEC_RELATION_GROUPS_Identifiers[a_spec_relation_groups] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATION_GROUPS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relation_groups.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_relation_groups.Name))
		initializerStatements += setValueField

	}

	map_A_SPEC_RELATION_REF_Identifiers := make(map[*A_SPEC_RELATION_REF]string)
	_ = map_A_SPEC_RELATION_REF_Identifiers

	a_spec_relation_refOrdered := []*A_SPEC_RELATION_REF{}
	for a_spec_relation_ref := range stage.A_SPEC_RELATION_REFs {
		a_spec_relation_refOrdered = append(a_spec_relation_refOrdered, a_spec_relation_ref)
	}
	sort.Slice(a_spec_relation_refOrdered[:], func(i, j int) bool {
		a_spec_relation_refi := a_spec_relation_refOrdered[i]
		a_spec_relation_refj := a_spec_relation_refOrdered[j]
		a_spec_relation_refi_order, oki := stage.A_SPEC_RELATION_REFMap_Staged_Order[a_spec_relation_refi]
		a_spec_relation_refj_order, okj := stage.A_SPEC_RELATION_REFMap_Staged_Order[a_spec_relation_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_relation_refi_order < a_spec_relation_refj_order
	})
	if len(a_spec_relation_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_spec_relation_ref := range a_spec_relation_refOrdered {

		id = generatesIdentifier("A_SPEC_RELATION_REF", int(stage.A_SPEC_RELATION_REFMap_Staged_Order[a_spec_relation_ref]), a_spec_relation_ref.Name)
		map_A_SPEC_RELATION_REF_Identifiers[a_spec_relation_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATION_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relation_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_relation_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPEC_RELATION_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_relation_ref.SPEC_RELATION_REF))
		initializerStatements += setValueField

	}

	map_A_SPEC_RELATION_TYPE_REF_Identifiers := make(map[*A_SPEC_RELATION_TYPE_REF]string)
	_ = map_A_SPEC_RELATION_TYPE_REF_Identifiers

	a_spec_relation_type_refOrdered := []*A_SPEC_RELATION_TYPE_REF{}
	for a_spec_relation_type_ref := range stage.A_SPEC_RELATION_TYPE_REFs {
		a_spec_relation_type_refOrdered = append(a_spec_relation_type_refOrdered, a_spec_relation_type_ref)
	}
	sort.Slice(a_spec_relation_type_refOrdered[:], func(i, j int) bool {
		a_spec_relation_type_refi := a_spec_relation_type_refOrdered[i]
		a_spec_relation_type_refj := a_spec_relation_type_refOrdered[j]
		a_spec_relation_type_refi_order, oki := stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order[a_spec_relation_type_refi]
		a_spec_relation_type_refj_order, okj := stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order[a_spec_relation_type_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_relation_type_refi_order < a_spec_relation_type_refj_order
	})
	if len(a_spec_relation_type_refOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_spec_relation_type_ref := range a_spec_relation_type_refOrdered {

		id = generatesIdentifier("A_SPEC_RELATION_TYPE_REF", int(stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order[a_spec_relation_type_ref]), a_spec_relation_type_ref.Name)
		map_A_SPEC_RELATION_TYPE_REF_Identifiers[a_spec_relation_type_ref] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATION_TYPE_REF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relation_type_ref.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_relation_type_ref.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPEC_RELATION_TYPE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_relation_type_ref.SPEC_RELATION_TYPE_REF))
		initializerStatements += setValueField

	}

	map_A_SPEC_TYPES_Identifiers := make(map[*A_SPEC_TYPES]string)
	_ = map_A_SPEC_TYPES_Identifiers

	a_spec_typesOrdered := []*A_SPEC_TYPES{}
	for a_spec_types := range stage.A_SPEC_TYPESs {
		a_spec_typesOrdered = append(a_spec_typesOrdered, a_spec_types)
	}
	sort.Slice(a_spec_typesOrdered[:], func(i, j int) bool {
		a_spec_typesi := a_spec_typesOrdered[i]
		a_spec_typesj := a_spec_typesOrdered[j]
		a_spec_typesi_order, oki := stage.A_SPEC_TYPESMap_Staged_Order[a_spec_typesi]
		a_spec_typesj_order, okj := stage.A_SPEC_TYPESMap_Staged_Order[a_spec_typesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_typesi_order < a_spec_typesj_order
	})
	if len(a_spec_typesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_spec_types := range a_spec_typesOrdered {

		id = generatesIdentifier("A_SPEC_TYPES", int(stage.A_SPEC_TYPESMap_Staged_Order[a_spec_types]), a_spec_types.Name)
		map_A_SPEC_TYPES_Identifiers[a_spec_types] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_TYPES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_types.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_types.Name))
		initializerStatements += setValueField

	}

	map_A_THE_HEADER_Identifiers := make(map[*A_THE_HEADER]string)
	_ = map_A_THE_HEADER_Identifiers

	a_the_headerOrdered := []*A_THE_HEADER{}
	for a_the_header := range stage.A_THE_HEADERs {
		a_the_headerOrdered = append(a_the_headerOrdered, a_the_header)
	}
	sort.Slice(a_the_headerOrdered[:], func(i, j int) bool {
		a_the_headeri := a_the_headerOrdered[i]
		a_the_headerj := a_the_headerOrdered[j]
		a_the_headeri_order, oki := stage.A_THE_HEADERMap_Staged_Order[a_the_headeri]
		a_the_headerj_order, okj := stage.A_THE_HEADERMap_Staged_Order[a_the_headerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_the_headeri_order < a_the_headerj_order
	})
	if len(a_the_headerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_the_header := range a_the_headerOrdered {

		id = generatesIdentifier("A_THE_HEADER", int(stage.A_THE_HEADERMap_Staged_Order[a_the_header]), a_the_header.Name)
		map_A_THE_HEADER_Identifiers[a_the_header] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_THE_HEADER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_the_header.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_the_header.Name))
		initializerStatements += setValueField

	}

	map_A_TOOL_EXTENSIONS_Identifiers := make(map[*A_TOOL_EXTENSIONS]string)
	_ = map_A_TOOL_EXTENSIONS_Identifiers

	a_tool_extensionsOrdered := []*A_TOOL_EXTENSIONS{}
	for a_tool_extensions := range stage.A_TOOL_EXTENSIONSs {
		a_tool_extensionsOrdered = append(a_tool_extensionsOrdered, a_tool_extensions)
	}
	sort.Slice(a_tool_extensionsOrdered[:], func(i, j int) bool {
		a_tool_extensionsi := a_tool_extensionsOrdered[i]
		a_tool_extensionsj := a_tool_extensionsOrdered[j]
		a_tool_extensionsi_order, oki := stage.A_TOOL_EXTENSIONSMap_Staged_Order[a_tool_extensionsi]
		a_tool_extensionsj_order, okj := stage.A_TOOL_EXTENSIONSMap_Staged_Order[a_tool_extensionsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_tool_extensionsi_order < a_tool_extensionsj_order
	})
	if len(a_tool_extensionsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, a_tool_extensions := range a_tool_extensionsOrdered {

		id = generatesIdentifier("A_TOOL_EXTENSIONS", int(stage.A_TOOL_EXTENSIONSMap_Staged_Order[a_tool_extensions]), a_tool_extensions.Name)
		map_A_TOOL_EXTENSIONS_Identifiers[a_tool_extensions] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TOOL_EXTENSIONS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_tool_extensions.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_tool_extensions.Name))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_BOOLEAN_Identifiers := make(map[*DATATYPE_DEFINITION_BOOLEAN]string)
	_ = map_DATATYPE_DEFINITION_BOOLEAN_Identifiers

	datatype_definition_booleanOrdered := []*DATATYPE_DEFINITION_BOOLEAN{}
	for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
		datatype_definition_booleanOrdered = append(datatype_definition_booleanOrdered, datatype_definition_boolean)
	}
	sort.Slice(datatype_definition_booleanOrdered[:], func(i, j int) bool {
		datatype_definition_booleani := datatype_definition_booleanOrdered[i]
		datatype_definition_booleanj := datatype_definition_booleanOrdered[j]
		datatype_definition_booleani_order, oki := stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[datatype_definition_booleani]
		datatype_definition_booleanj_order, okj := stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[datatype_definition_booleanj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_booleani_order < datatype_definition_booleanj_order
	})
	if len(datatype_definition_booleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, datatype_definition_boolean := range datatype_definition_booleanOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_BOOLEAN", int(stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[datatype_definition_boolean]), datatype_definition_boolean.Name)
		map_DATATYPE_DEFINITION_BOOLEAN_Identifiers[datatype_definition_boolean] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_BOOLEAN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_boolean.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.LONG_NAME))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_DATE_Identifiers := make(map[*DATATYPE_DEFINITION_DATE]string)
	_ = map_DATATYPE_DEFINITION_DATE_Identifiers

	datatype_definition_dateOrdered := []*DATATYPE_DEFINITION_DATE{}
	for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
		datatype_definition_dateOrdered = append(datatype_definition_dateOrdered, datatype_definition_date)
	}
	sort.Slice(datatype_definition_dateOrdered[:], func(i, j int) bool {
		datatype_definition_datei := datatype_definition_dateOrdered[i]
		datatype_definition_datej := datatype_definition_dateOrdered[j]
		datatype_definition_datei_order, oki := stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[datatype_definition_datei]
		datatype_definition_datej_order, okj := stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[datatype_definition_datej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_datei_order < datatype_definition_datej_order
	})
	if len(datatype_definition_dateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, datatype_definition_date := range datatype_definition_dateOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_DATE", int(stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[datatype_definition_date]), datatype_definition_date.Name)
		map_DATATYPE_DEFINITION_DATE_Identifiers[datatype_definition_date] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_DATE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_date.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.LONG_NAME))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_ENUMERATION_Identifiers := make(map[*DATATYPE_DEFINITION_ENUMERATION]string)
	_ = map_DATATYPE_DEFINITION_ENUMERATION_Identifiers

	datatype_definition_enumerationOrdered := []*DATATYPE_DEFINITION_ENUMERATION{}
	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		datatype_definition_enumerationOrdered = append(datatype_definition_enumerationOrdered, datatype_definition_enumeration)
	}
	sort.Slice(datatype_definition_enumerationOrdered[:], func(i, j int) bool {
		datatype_definition_enumerationi := datatype_definition_enumerationOrdered[i]
		datatype_definition_enumerationj := datatype_definition_enumerationOrdered[j]
		datatype_definition_enumerationi_order, oki := stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[datatype_definition_enumerationi]
		datatype_definition_enumerationj_order, okj := stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[datatype_definition_enumerationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_enumerationi_order < datatype_definition_enumerationj_order
	})
	if len(datatype_definition_enumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, datatype_definition_enumeration := range datatype_definition_enumerationOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_ENUMERATION", int(stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[datatype_definition_enumeration]), datatype_definition_enumeration.Name)
		map_DATATYPE_DEFINITION_ENUMERATION_Identifiers[datatype_definition_enumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_ENUMERATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_enumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.LONG_NAME))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_INTEGER_Identifiers := make(map[*DATATYPE_DEFINITION_INTEGER]string)
	_ = map_DATATYPE_DEFINITION_INTEGER_Identifiers

	datatype_definition_integerOrdered := []*DATATYPE_DEFINITION_INTEGER{}
	for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
		datatype_definition_integerOrdered = append(datatype_definition_integerOrdered, datatype_definition_integer)
	}
	sort.Slice(datatype_definition_integerOrdered[:], func(i, j int) bool {
		datatype_definition_integeri := datatype_definition_integerOrdered[i]
		datatype_definition_integerj := datatype_definition_integerOrdered[j]
		datatype_definition_integeri_order, oki := stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[datatype_definition_integeri]
		datatype_definition_integerj_order, okj := stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[datatype_definition_integerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_integeri_order < datatype_definition_integerj_order
	})
	if len(datatype_definition_integerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, datatype_definition_integer := range datatype_definition_integerOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_INTEGER", int(stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[datatype_definition_integer]), datatype_definition_integer.Name)
		map_DATATYPE_DEFINITION_INTEGER_Identifiers[datatype_definition_integer] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_INTEGER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_integer.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.LONG_NAME))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MAX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_integer.MAX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MIN")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_integer.MIN))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_REAL_Identifiers := make(map[*DATATYPE_DEFINITION_REAL]string)
	_ = map_DATATYPE_DEFINITION_REAL_Identifiers

	datatype_definition_realOrdered := []*DATATYPE_DEFINITION_REAL{}
	for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
		datatype_definition_realOrdered = append(datatype_definition_realOrdered, datatype_definition_real)
	}
	sort.Slice(datatype_definition_realOrdered[:], func(i, j int) bool {
		datatype_definition_reali := datatype_definition_realOrdered[i]
		datatype_definition_realj := datatype_definition_realOrdered[j]
		datatype_definition_reali_order, oki := stage.DATATYPE_DEFINITION_REALMap_Staged_Order[datatype_definition_reali]
		datatype_definition_realj_order, okj := stage.DATATYPE_DEFINITION_REALMap_Staged_Order[datatype_definition_realj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_reali_order < datatype_definition_realj_order
	})
	if len(datatype_definition_realOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, datatype_definition_real := range datatype_definition_realOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_REAL", int(stage.DATATYPE_DEFINITION_REALMap_Staged_Order[datatype_definition_real]), datatype_definition_real.Name)
		map_DATATYPE_DEFINITION_REAL_Identifiers[datatype_definition_real] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_REAL")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_real.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ACCURACY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_real.ACCURACY))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.LONG_NAME))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MAX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", datatype_definition_real.MAX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MIN")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", datatype_definition_real.MIN))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_STRING_Identifiers := make(map[*DATATYPE_DEFINITION_STRING]string)
	_ = map_DATATYPE_DEFINITION_STRING_Identifiers

	datatype_definition_stringOrdered := []*DATATYPE_DEFINITION_STRING{}
	for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
		datatype_definition_stringOrdered = append(datatype_definition_stringOrdered, datatype_definition_string)
	}
	sort.Slice(datatype_definition_stringOrdered[:], func(i, j int) bool {
		datatype_definition_stringi := datatype_definition_stringOrdered[i]
		datatype_definition_stringj := datatype_definition_stringOrdered[j]
		datatype_definition_stringi_order, oki := stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[datatype_definition_stringi]
		datatype_definition_stringj_order, okj := stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[datatype_definition_stringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_stringi_order < datatype_definition_stringj_order
	})
	if len(datatype_definition_stringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, datatype_definition_string := range datatype_definition_stringOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_STRING", int(stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[datatype_definition_string]), datatype_definition_string.Name)
		map_DATATYPE_DEFINITION_STRING_Identifiers[datatype_definition_string] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_STRING")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_string.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.LONG_NAME))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MAX_LENGTH")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_string.MAX_LENGTH))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_XHTML_Identifiers := make(map[*DATATYPE_DEFINITION_XHTML]string)
	_ = map_DATATYPE_DEFINITION_XHTML_Identifiers

	datatype_definition_xhtmlOrdered := []*DATATYPE_DEFINITION_XHTML{}
	for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
		datatype_definition_xhtmlOrdered = append(datatype_definition_xhtmlOrdered, datatype_definition_xhtml)
	}
	sort.Slice(datatype_definition_xhtmlOrdered[:], func(i, j int) bool {
		datatype_definition_xhtmli := datatype_definition_xhtmlOrdered[i]
		datatype_definition_xhtmlj := datatype_definition_xhtmlOrdered[j]
		datatype_definition_xhtmli_order, oki := stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[datatype_definition_xhtmli]
		datatype_definition_xhtmlj_order, okj := stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[datatype_definition_xhtmlj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_xhtmli_order < datatype_definition_xhtmlj_order
	})
	if len(datatype_definition_xhtmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, datatype_definition_xhtml := range datatype_definition_xhtmlOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_XHTML", int(stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[datatype_definition_xhtml]), datatype_definition_xhtml.Name)
		map_DATATYPE_DEFINITION_XHTML_Identifiers[datatype_definition_xhtml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_XHTML")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_xhtml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.LONG_NAME))
		initializerStatements += setValueField

	}

	map_EMBEDDED_VALUE_Identifiers := make(map[*EMBEDDED_VALUE]string)
	_ = map_EMBEDDED_VALUE_Identifiers

	embedded_valueOrdered := []*EMBEDDED_VALUE{}
	for embedded_value := range stage.EMBEDDED_VALUEs {
		embedded_valueOrdered = append(embedded_valueOrdered, embedded_value)
	}
	sort.Slice(embedded_valueOrdered[:], func(i, j int) bool {
		embedded_valuei := embedded_valueOrdered[i]
		embedded_valuej := embedded_valueOrdered[j]
		embedded_valuei_order, oki := stage.EMBEDDED_VALUEMap_Staged_Order[embedded_valuei]
		embedded_valuej_order, okj := stage.EMBEDDED_VALUEMap_Staged_Order[embedded_valuej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return embedded_valuei_order < embedded_valuej_order
	})
	if len(embedded_valueOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, embedded_value := range embedded_valueOrdered {

		id = generatesIdentifier("EMBEDDED_VALUE", int(stage.EMBEDDED_VALUEMap_Staged_Order[embedded_value]), embedded_value.Name)
		map_EMBEDDED_VALUE_Identifiers[embedded_value] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EMBEDDED_VALUE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", embedded_value.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embedded_value.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "KEY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", embedded_value.KEY))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OTHER_CONTENT")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embedded_value.OTHER_CONTENT))
		initializerStatements += setValueField

	}

	map_ENUM_VALUE_Identifiers := make(map[*ENUM_VALUE]string)
	_ = map_ENUM_VALUE_Identifiers

	enum_valueOrdered := []*ENUM_VALUE{}
	for enum_value := range stage.ENUM_VALUEs {
		enum_valueOrdered = append(enum_valueOrdered, enum_value)
	}
	sort.Slice(enum_valueOrdered[:], func(i, j int) bool {
		enum_valuei := enum_valueOrdered[i]
		enum_valuej := enum_valueOrdered[j]
		enum_valuei_order, oki := stage.ENUM_VALUEMap_Staged_Order[enum_valuei]
		enum_valuej_order, okj := stage.ENUM_VALUEMap_Staged_Order[enum_valuej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return enum_valuei_order < enum_valuej_order
	})
	if len(enum_valueOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, enum_value := range enum_valueOrdered {

		id = generatesIdentifier("ENUM_VALUE", int(stage.ENUM_VALUEMap_Staged_Order[enum_value]), enum_value.Name)
		map_ENUM_VALUE_Identifiers[enum_value] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ENUM_VALUE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", enum_value.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.LONG_NAME))
		initializerStatements += setValueField

	}

	map_EmbeddedJpgImage_Identifiers := make(map[*EmbeddedJpgImage]string)
	_ = map_EmbeddedJpgImage_Identifiers

	embeddedjpgimageOrdered := []*EmbeddedJpgImage{}
	for embeddedjpgimage := range stage.EmbeddedJpgImages {
		embeddedjpgimageOrdered = append(embeddedjpgimageOrdered, embeddedjpgimage)
	}
	sort.Slice(embeddedjpgimageOrdered[:], func(i, j int) bool {
		embeddedjpgimagei := embeddedjpgimageOrdered[i]
		embeddedjpgimagej := embeddedjpgimageOrdered[j]
		embeddedjpgimagei_order, oki := stage.EmbeddedJpgImageMap_Staged_Order[embeddedjpgimagei]
		embeddedjpgimagej_order, okj := stage.EmbeddedJpgImageMap_Staged_Order[embeddedjpgimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return embeddedjpgimagei_order < embeddedjpgimagej_order
	})
	if len(embeddedjpgimageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, embeddedjpgimage := range embeddedjpgimageOrdered {

		id = generatesIdentifier("EmbeddedJpgImage", int(stage.EmbeddedJpgImageMap_Staged_Order[embeddedjpgimage]), embeddedjpgimage.Name)
		map_EmbeddedJpgImage_Identifiers[embeddedjpgimage] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EmbeddedJpgImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", embeddedjpgimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embeddedjpgimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Base64Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embeddedjpgimage.Base64Content))
		initializerStatements += setValueField

	}

	map_EmbeddedPngImage_Identifiers := make(map[*EmbeddedPngImage]string)
	_ = map_EmbeddedPngImage_Identifiers

	embeddedpngimageOrdered := []*EmbeddedPngImage{}
	for embeddedpngimage := range stage.EmbeddedPngImages {
		embeddedpngimageOrdered = append(embeddedpngimageOrdered, embeddedpngimage)
	}
	sort.Slice(embeddedpngimageOrdered[:], func(i, j int) bool {
		embeddedpngimagei := embeddedpngimageOrdered[i]
		embeddedpngimagej := embeddedpngimageOrdered[j]
		embeddedpngimagei_order, oki := stage.EmbeddedPngImageMap_Staged_Order[embeddedpngimagei]
		embeddedpngimagej_order, okj := stage.EmbeddedPngImageMap_Staged_Order[embeddedpngimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return embeddedpngimagei_order < embeddedpngimagej_order
	})
	if len(embeddedpngimageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, embeddedpngimage := range embeddedpngimageOrdered {

		id = generatesIdentifier("EmbeddedPngImage", int(stage.EmbeddedPngImageMap_Staged_Order[embeddedpngimage]), embeddedpngimage.Name)
		map_EmbeddedPngImage_Identifiers[embeddedpngimage] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EmbeddedPngImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", embeddedpngimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embeddedpngimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Base64Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embeddedpngimage.Base64Content))
		initializerStatements += setValueField

	}

	map_EmbeddedSvgImage_Identifiers := make(map[*EmbeddedSvgImage]string)
	_ = map_EmbeddedSvgImage_Identifiers

	embeddedsvgimageOrdered := []*EmbeddedSvgImage{}
	for embeddedsvgimage := range stage.EmbeddedSvgImages {
		embeddedsvgimageOrdered = append(embeddedsvgimageOrdered, embeddedsvgimage)
	}
	sort.Slice(embeddedsvgimageOrdered[:], func(i, j int) bool {
		embeddedsvgimagei := embeddedsvgimageOrdered[i]
		embeddedsvgimagej := embeddedsvgimageOrdered[j]
		embeddedsvgimagei_order, oki := stage.EmbeddedSvgImageMap_Staged_Order[embeddedsvgimagei]
		embeddedsvgimagej_order, okj := stage.EmbeddedSvgImageMap_Staged_Order[embeddedsvgimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return embeddedsvgimagei_order < embeddedsvgimagej_order
	})
	if len(embeddedsvgimageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, embeddedsvgimage := range embeddedsvgimageOrdered {

		id = generatesIdentifier("EmbeddedSvgImage", int(stage.EmbeddedSvgImageMap_Staged_Order[embeddedsvgimage]), embeddedsvgimage.Name)
		map_EmbeddedSvgImage_Identifiers[embeddedsvgimage] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EmbeddedSvgImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", embeddedsvgimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embeddedsvgimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embeddedsvgimage.Content))
		initializerStatements += setValueField

	}

	map_Kill_Identifiers := make(map[*Kill]string)
	_ = map_Kill_Identifiers

	killOrdered := []*Kill{}
	for kill := range stage.Kills {
		killOrdered = append(killOrdered, kill)
	}
	sort.Slice(killOrdered[:], func(i, j int) bool {
		killi := killOrdered[i]
		killj := killOrdered[j]
		killi_order, oki := stage.KillMap_Staged_Order[killi]
		killj_order, okj := stage.KillMap_Staged_Order[killj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return killi_order < killj_order
	})
	if len(killOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, kill := range killOrdered {

		id = generatesIdentifier("Kill", int(stage.KillMap_Staged_Order[kill]), kill.Name)
		map_Kill_Identifiers[kill] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Kill")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", kill.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(kill.Name))
		initializerStatements += setValueField

	}

	map_Map_identifier_bool_Identifiers := make(map[*Map_identifier_bool]string)
	_ = map_Map_identifier_bool_Identifiers

	map_identifier_boolOrdered := []*Map_identifier_bool{}
	for map_identifier_bool := range stage.Map_identifier_bools {
		map_identifier_boolOrdered = append(map_identifier_boolOrdered, map_identifier_bool)
	}
	sort.Slice(map_identifier_boolOrdered[:], func(i, j int) bool {
		map_identifier_booli := map_identifier_boolOrdered[i]
		map_identifier_boolj := map_identifier_boolOrdered[j]
		map_identifier_booli_order, oki := stage.Map_identifier_boolMap_Staged_Order[map_identifier_booli]
		map_identifier_boolj_order, okj := stage.Map_identifier_boolMap_Staged_Order[map_identifier_boolj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return map_identifier_booli_order < map_identifier_boolj_order
	})
	if len(map_identifier_boolOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, map_identifier_bool := range map_identifier_boolOrdered {

		id = generatesIdentifier("Map_identifier_bool", int(stage.Map_identifier_boolMap_Staged_Order[map_identifier_bool]), map_identifier_bool.Name)
		map_Map_identifier_bool_Identifiers[map_identifier_bool] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Map_identifier_bool")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", map_identifier_bool.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(map_identifier_bool.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", map_identifier_bool.Value))
		initializerStatements += setValueField

	}

	map_RELATION_GROUP_Identifiers := make(map[*RELATION_GROUP]string)
	_ = map_RELATION_GROUP_Identifiers

	relation_groupOrdered := []*RELATION_GROUP{}
	for relation_group := range stage.RELATION_GROUPs {
		relation_groupOrdered = append(relation_groupOrdered, relation_group)
	}
	sort.Slice(relation_groupOrdered[:], func(i, j int) bool {
		relation_groupi := relation_groupOrdered[i]
		relation_groupj := relation_groupOrdered[j]
		relation_groupi_order, oki := stage.RELATION_GROUPMap_Staged_Order[relation_groupi]
		relation_groupj_order, okj := stage.RELATION_GROUPMap_Staged_Order[relation_groupj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return relation_groupi_order < relation_groupj_order
	})
	if len(relation_groupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, relation_group := range relation_groupOrdered {

		id = generatesIdentifier("RELATION_GROUP", int(stage.RELATION_GROUPMap_Staged_Order[relation_group]), relation_group.Name)
		map_RELATION_GROUP_Identifiers[relation_group] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RELATION_GROUP")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", relation_group.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group.LONG_NAME))
		initializerStatements += setValueField

	}

	map_RELATION_GROUP_TYPE_Identifiers := make(map[*RELATION_GROUP_TYPE]string)
	_ = map_RELATION_GROUP_TYPE_Identifiers

	relation_group_typeOrdered := []*RELATION_GROUP_TYPE{}
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		relation_group_typeOrdered = append(relation_group_typeOrdered, relation_group_type)
	}
	sort.Slice(relation_group_typeOrdered[:], func(i, j int) bool {
		relation_group_typei := relation_group_typeOrdered[i]
		relation_group_typej := relation_group_typeOrdered[j]
		relation_group_typei_order, oki := stage.RELATION_GROUP_TYPEMap_Staged_Order[relation_group_typei]
		relation_group_typej_order, okj := stage.RELATION_GROUP_TYPEMap_Staged_Order[relation_group_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return relation_group_typei_order < relation_group_typej_order
	})
	if len(relation_group_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, relation_group_type := range relation_group_typeOrdered {

		id = generatesIdentifier("RELATION_GROUP_TYPE", int(stage.RELATION_GROUP_TYPEMap_Staged_Order[relation_group_type]), relation_group_type.Name)
		map_RELATION_GROUP_TYPE_Identifiers[relation_group_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RELATION_GROUP_TYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", relation_group_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group_type.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group_type.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group_type.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group_type.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group_type.LONG_NAME))
		initializerStatements += setValueField

	}

	map_REQ_IF_Identifiers := make(map[*REQ_IF]string)
	_ = map_REQ_IF_Identifiers

	req_ifOrdered := []*REQ_IF{}
	for req_if := range stage.REQ_IFs {
		req_ifOrdered = append(req_ifOrdered, req_if)
	}
	sort.Slice(req_ifOrdered[:], func(i, j int) bool {
		req_ifi := req_ifOrdered[i]
		req_ifj := req_ifOrdered[j]
		req_ifi_order, oki := stage.REQ_IFMap_Staged_Order[req_ifi]
		req_ifj_order, okj := stage.REQ_IFMap_Staged_Order[req_ifj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return req_ifi_order < req_ifj_order
	})
	if len(req_ifOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, req_if := range req_ifOrdered {

		id = generatesIdentifier("REQ_IF", int(stage.REQ_IFMap_Staged_Order[req_if]), req_if.Name)
		map_REQ_IF_Identifiers[req_if] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lang")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if.Lang))
		initializerStatements += setValueField

	}

	map_REQ_IF_CONTENT_Identifiers := make(map[*REQ_IF_CONTENT]string)
	_ = map_REQ_IF_CONTENT_Identifiers

	req_if_contentOrdered := []*REQ_IF_CONTENT{}
	for req_if_content := range stage.REQ_IF_CONTENTs {
		req_if_contentOrdered = append(req_if_contentOrdered, req_if_content)
	}
	sort.Slice(req_if_contentOrdered[:], func(i, j int) bool {
		req_if_contenti := req_if_contentOrdered[i]
		req_if_contentj := req_if_contentOrdered[j]
		req_if_contenti_order, oki := stage.REQ_IF_CONTENTMap_Staged_Order[req_if_contenti]
		req_if_contentj_order, okj := stage.REQ_IF_CONTENTMap_Staged_Order[req_if_contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return req_if_contenti_order < req_if_contentj_order
	})
	if len(req_if_contentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, req_if_content := range req_if_contentOrdered {

		id = generatesIdentifier("REQ_IF_CONTENT", int(stage.REQ_IF_CONTENTMap_Staged_Order[req_if_content]), req_if_content.Name)
		map_REQ_IF_CONTENT_Identifiers[req_if_content] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_CONTENT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if_content.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_content.Name))
		initializerStatements += setValueField

	}

	map_REQ_IF_HEADER_Identifiers := make(map[*REQ_IF_HEADER]string)
	_ = map_REQ_IF_HEADER_Identifiers

	req_if_headerOrdered := []*REQ_IF_HEADER{}
	for req_if_header := range stage.REQ_IF_HEADERs {
		req_if_headerOrdered = append(req_if_headerOrdered, req_if_header)
	}
	sort.Slice(req_if_headerOrdered[:], func(i, j int) bool {
		req_if_headeri := req_if_headerOrdered[i]
		req_if_headerj := req_if_headerOrdered[j]
		req_if_headeri_order, oki := stage.REQ_IF_HEADERMap_Staged_Order[req_if_headeri]
		req_if_headerj_order, okj := stage.REQ_IF_HEADERMap_Staged_Order[req_if_headerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return req_if_headeri_order < req_if_headerj_order
	})
	if len(req_if_headerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, req_if_header := range req_if_headerOrdered {

		id = generatesIdentifier("REQ_IF_HEADER", int(stage.REQ_IF_HEADERMap_Staged_Order[req_if_header]), req_if_header.Name)
		map_REQ_IF_HEADER_Identifiers[req_if_header] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_HEADER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if_header.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "COMMENT")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.COMMENT))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CREATION_TIME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.CREATION_TIME))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "REPOSITORY_ID")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.REPOSITORY_ID))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "REQ_IF_TOOL_ID")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.REQ_IF_TOOL_ID))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "REQ_IF_VERSION")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.REQ_IF_VERSION))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SOURCE_TOOL_ID")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.SOURCE_TOOL_ID))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TITLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.TITLE))
		initializerStatements += setValueField

	}

	map_REQ_IF_TOOL_EXTENSION_Identifiers := make(map[*REQ_IF_TOOL_EXTENSION]string)
	_ = map_REQ_IF_TOOL_EXTENSION_Identifiers

	req_if_tool_extensionOrdered := []*REQ_IF_TOOL_EXTENSION{}
	for req_if_tool_extension := range stage.REQ_IF_TOOL_EXTENSIONs {
		req_if_tool_extensionOrdered = append(req_if_tool_extensionOrdered, req_if_tool_extension)
	}
	sort.Slice(req_if_tool_extensionOrdered[:], func(i, j int) bool {
		req_if_tool_extensioni := req_if_tool_extensionOrdered[i]
		req_if_tool_extensionj := req_if_tool_extensionOrdered[j]
		req_if_tool_extensioni_order, oki := stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[req_if_tool_extensioni]
		req_if_tool_extensionj_order, okj := stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[req_if_tool_extensionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return req_if_tool_extensioni_order < req_if_tool_extensionj_order
	})
	if len(req_if_tool_extensionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, req_if_tool_extension := range req_if_tool_extensionOrdered {

		id = generatesIdentifier("REQ_IF_TOOL_EXTENSION", int(stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[req_if_tool_extension]), req_if_tool_extension.Name)
		map_REQ_IF_TOOL_EXTENSION_Identifiers[req_if_tool_extension] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_TOOL_EXTENSION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if_tool_extension.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_tool_extension.Name))
		initializerStatements += setValueField

	}

	map_SPECIFICATION_Identifiers := make(map[*SPECIFICATION]string)
	_ = map_SPECIFICATION_Identifiers

	specificationOrdered := []*SPECIFICATION{}
	for specification := range stage.SPECIFICATIONs {
		specificationOrdered = append(specificationOrdered, specification)
	}
	sort.Slice(specificationOrdered[:], func(i, j int) bool {
		specificationi := specificationOrdered[i]
		specificationj := specificationOrdered[j]
		specificationi_order, oki := stage.SPECIFICATIONMap_Staged_Order[specificationi]
		specificationj_order, okj := stage.SPECIFICATIONMap_Staged_Order[specificationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return specificationi_order < specificationj_order
	})
	if len(specificationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, specification := range specificationOrdered {

		id = generatesIdentifier("SPECIFICATION", int(stage.SPECIFICATIONMap_Staged_Order[specification]), specification.Name)
		map_SPECIFICATION_Identifiers[specification] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specification.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPECIFICATION_Rendering_Identifiers := make(map[*SPECIFICATION_Rendering]string)
	_ = map_SPECIFICATION_Rendering_Identifiers

	specification_renderingOrdered := []*SPECIFICATION_Rendering{}
	for specification_rendering := range stage.SPECIFICATION_Renderings {
		specification_renderingOrdered = append(specification_renderingOrdered, specification_rendering)
	}
	sort.Slice(specification_renderingOrdered[:], func(i, j int) bool {
		specification_renderingi := specification_renderingOrdered[i]
		specification_renderingj := specification_renderingOrdered[j]
		specification_renderingi_order, oki := stage.SPECIFICATION_RenderingMap_Staged_Order[specification_renderingi]
		specification_renderingj_order, okj := stage.SPECIFICATION_RenderingMap_Staged_Order[specification_renderingj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return specification_renderingi_order < specification_renderingj_order
	})
	if len(specification_renderingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, specification_rendering := range specification_renderingOrdered {

		id = generatesIdentifier("SPECIFICATION_Rendering", int(stage.SPECIFICATION_RenderingMap_Staged_Order[specification_rendering]), specification_rendering.Name)
		map_SPECIFICATION_Rendering_Identifiers[specification_rendering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION_Rendering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specification_rendering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_rendering.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", specification_rendering.IsNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", specification_rendering.IsSelected))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsWithHeadingNumbering")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", specification_rendering.IsWithHeadingNumbering))
		initializerStatements += setValueField

	}

	map_SPECIFICATION_TYPE_Identifiers := make(map[*SPECIFICATION_TYPE]string)
	_ = map_SPECIFICATION_TYPE_Identifiers

	specification_typeOrdered := []*SPECIFICATION_TYPE{}
	for specification_type := range stage.SPECIFICATION_TYPEs {
		specification_typeOrdered = append(specification_typeOrdered, specification_type)
	}
	sort.Slice(specification_typeOrdered[:], func(i, j int) bool {
		specification_typei := specification_typeOrdered[i]
		specification_typej := specification_typeOrdered[j]
		specification_typei_order, oki := stage.SPECIFICATION_TYPEMap_Staged_Order[specification_typei]
		specification_typej_order, okj := stage.SPECIFICATION_TYPEMap_Staged_Order[specification_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return specification_typei_order < specification_typej_order
	})
	if len(specification_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, specification_type := range specification_typeOrdered {

		id = generatesIdentifier("SPECIFICATION_TYPE", int(stage.SPECIFICATION_TYPEMap_Staged_Order[specification_type]), specification_type.Name)
		map_SPECIFICATION_TYPE_Identifiers[specification_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION_TYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specification_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_type.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_type.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_type.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_type.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_type.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPEC_HIERARCHY_Identifiers := make(map[*SPEC_HIERARCHY]string)
	_ = map_SPEC_HIERARCHY_Identifiers

	spec_hierarchyOrdered := []*SPEC_HIERARCHY{}
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		spec_hierarchyOrdered = append(spec_hierarchyOrdered, spec_hierarchy)
	}
	sort.Slice(spec_hierarchyOrdered[:], func(i, j int) bool {
		spec_hierarchyi := spec_hierarchyOrdered[i]
		spec_hierarchyj := spec_hierarchyOrdered[j]
		spec_hierarchyi_order, oki := stage.SPEC_HIERARCHYMap_Staged_Order[spec_hierarchyi]
		spec_hierarchyj_order, okj := stage.SPEC_HIERARCHYMap_Staged_Order[spec_hierarchyj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_hierarchyi_order < spec_hierarchyj_order
	})
	if len(spec_hierarchyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spec_hierarchy := range spec_hierarchyOrdered {

		id = generatesIdentifier("SPEC_HIERARCHY", int(stage.SPEC_HIERARCHYMap_Staged_Order[spec_hierarchy]), spec_hierarchy.Name)
		map_SPEC_HIERARCHY_Identifiers[spec_hierarchy] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_HIERARCHY")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_hierarchy.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_hierarchy.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_TABLE_INTERNAL")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_hierarchy.IS_TABLE_INTERNAL))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPEC_OBJECT_Identifiers := make(map[*SPEC_OBJECT]string)
	_ = map_SPEC_OBJECT_Identifiers

	spec_objectOrdered := []*SPEC_OBJECT{}
	for spec_object := range stage.SPEC_OBJECTs {
		spec_objectOrdered = append(spec_objectOrdered, spec_object)
	}
	sort.Slice(spec_objectOrdered[:], func(i, j int) bool {
		spec_objecti := spec_objectOrdered[i]
		spec_objectj := spec_objectOrdered[j]
		spec_objecti_order, oki := stage.SPEC_OBJECTMap_Staged_Order[spec_objecti]
		spec_objectj_order, okj := stage.SPEC_OBJECTMap_Staged_Order[spec_objectj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_objecti_order < spec_objectj_order
	})
	if len(spec_objectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spec_object := range spec_objectOrdered {

		id = generatesIdentifier("SPEC_OBJECT", int(stage.SPEC_OBJECTMap_Staged_Order[spec_object]), spec_object.Name)
		map_SPEC_OBJECT_Identifiers[spec_object] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_OBJECT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_object.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPEC_OBJECT_TYPE_Identifiers := make(map[*SPEC_OBJECT_TYPE]string)
	_ = map_SPEC_OBJECT_TYPE_Identifiers

	spec_object_typeOrdered := []*SPEC_OBJECT_TYPE{}
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		spec_object_typeOrdered = append(spec_object_typeOrdered, spec_object_type)
	}
	sort.Slice(spec_object_typeOrdered[:], func(i, j int) bool {
		spec_object_typei := spec_object_typeOrdered[i]
		spec_object_typej := spec_object_typeOrdered[j]
		spec_object_typei_order, oki := stage.SPEC_OBJECT_TYPEMap_Staged_Order[spec_object_typei]
		spec_object_typej_order, okj := stage.SPEC_OBJECT_TYPEMap_Staged_Order[spec_object_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_object_typei_order < spec_object_typej_order
	})
	if len(spec_object_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spec_object_type := range spec_object_typeOrdered {

		id = generatesIdentifier("SPEC_OBJECT_TYPE", int(stage.SPEC_OBJECT_TYPEMap_Staged_Order[spec_object_type]), spec_object_type.Name)
		map_SPEC_OBJECT_TYPE_Identifiers[spec_object_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_OBJECT_TYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_object_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPEC_OBJECT_TYPE_Rendering_Identifiers := make(map[*SPEC_OBJECT_TYPE_Rendering]string)
	_ = map_SPEC_OBJECT_TYPE_Rendering_Identifiers

	spec_object_type_renderingOrdered := []*SPEC_OBJECT_TYPE_Rendering{}
	for spec_object_type_rendering := range stage.SPEC_OBJECT_TYPE_Renderings {
		spec_object_type_renderingOrdered = append(spec_object_type_renderingOrdered, spec_object_type_rendering)
	}
	sort.Slice(spec_object_type_renderingOrdered[:], func(i, j int) bool {
		spec_object_type_renderingi := spec_object_type_renderingOrdered[i]
		spec_object_type_renderingj := spec_object_type_renderingOrdered[j]
		spec_object_type_renderingi_order, oki := stage.SPEC_OBJECT_TYPE_RenderingMap_Staged_Order[spec_object_type_renderingi]
		spec_object_type_renderingj_order, okj := stage.SPEC_OBJECT_TYPE_RenderingMap_Staged_Order[spec_object_type_renderingj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_object_type_renderingi_order < spec_object_type_renderingj_order
	})
	if len(spec_object_type_renderingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spec_object_type_rendering := range spec_object_type_renderingOrdered {

		id = generatesIdentifier("SPEC_OBJECT_TYPE_Rendering", int(stage.SPEC_OBJECT_TYPE_RenderingMap_Staged_Order[spec_object_type_rendering]), spec_object_type_rendering.Name)
		map_SPEC_OBJECT_TYPE_Rendering_Identifiers[spec_object_type_rendering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_OBJECT_TYPE_Rendering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_object_type_rendering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type_rendering.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_object_type_rendering.IsNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowIdentifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_object_type_rendering.ShowIdentifier))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_object_type_rendering.ShowName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowRelations")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_object_type_rendering.ShowRelations))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsHeading")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_object_type_rendering.IsHeading))
		initializerStatements += setValueField

	}

	map_SPEC_RELATION_Identifiers := make(map[*SPEC_RELATION]string)
	_ = map_SPEC_RELATION_Identifiers

	spec_relationOrdered := []*SPEC_RELATION{}
	for spec_relation := range stage.SPEC_RELATIONs {
		spec_relationOrdered = append(spec_relationOrdered, spec_relation)
	}
	sort.Slice(spec_relationOrdered[:], func(i, j int) bool {
		spec_relationi := spec_relationOrdered[i]
		spec_relationj := spec_relationOrdered[j]
		spec_relationi_order, oki := stage.SPEC_RELATIONMap_Staged_Order[spec_relationi]
		spec_relationj_order, okj := stage.SPEC_RELATIONMap_Staged_Order[spec_relationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_relationi_order < spec_relationj_order
	})
	if len(spec_relationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spec_relation := range spec_relationOrdered {

		id = generatesIdentifier("SPEC_RELATION", int(stage.SPEC_RELATIONMap_Staged_Order[spec_relation]), spec_relation.Name)
		map_SPEC_RELATION_Identifiers[spec_relation] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_RELATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_relation.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPEC_RELATION_TYPE_Identifiers := make(map[*SPEC_RELATION_TYPE]string)
	_ = map_SPEC_RELATION_TYPE_Identifiers

	spec_relation_typeOrdered := []*SPEC_RELATION_TYPE{}
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		spec_relation_typeOrdered = append(spec_relation_typeOrdered, spec_relation_type)
	}
	sort.Slice(spec_relation_typeOrdered[:], func(i, j int) bool {
		spec_relation_typei := spec_relation_typeOrdered[i]
		spec_relation_typej := spec_relation_typeOrdered[j]
		spec_relation_typei_order, oki := stage.SPEC_RELATION_TYPEMap_Staged_Order[spec_relation_typei]
		spec_relation_typej_order, okj := stage.SPEC_RELATION_TYPEMap_Staged_Order[spec_relation_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_relation_typei_order < spec_relation_typej_order
	})
	if len(spec_relation_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spec_relation_type := range spec_relation_typeOrdered {

		id = generatesIdentifier("SPEC_RELATION_TYPE", int(stage.SPEC_RELATION_TYPEMap_Staged_Order[spec_relation_type]), spec_relation_type.Name)
		map_SPEC_RELATION_TYPE_Identifiers[spec_relation_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_RELATION_TYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_relation_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.LONG_NAME))
		initializerStatements += setValueField

	}

	map_StaticWebSite_Identifiers := make(map[*StaticWebSite]string)
	_ = map_StaticWebSite_Identifiers

	staticwebsiteOrdered := []*StaticWebSite{}
	for staticwebsite := range stage.StaticWebSites {
		staticwebsiteOrdered = append(staticwebsiteOrdered, staticwebsite)
	}
	sort.Slice(staticwebsiteOrdered[:], func(i, j int) bool {
		staticwebsitei := staticwebsiteOrdered[i]
		staticwebsitej := staticwebsiteOrdered[j]
		staticwebsitei_order, oki := stage.StaticWebSiteMap_Staged_Order[staticwebsitei]
		staticwebsitej_order, okj := stage.StaticWebSiteMap_Staged_Order[staticwebsitej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return staticwebsitei_order < staticwebsitej_order
	})
	if len(staticwebsiteOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, staticwebsite := range staticwebsiteOrdered {

		id = generatesIdentifier("StaticWebSite", int(stage.StaticWebSiteMap_Staged_Order[staticwebsite]), staticwebsite.Name)
		map_StaticWebSite_Identifiers[staticwebsite] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StaticWebSite")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", staticwebsite.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsite.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MarkdownContent")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsite.MarkdownContent))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InputImagesDir")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsite.InputImagesDir))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OutputStaticWebDir")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsite.OutputStaticWebDir))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "VersionInfo")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsite.VersionInfo))
		initializerStatements += setValueField

	}

	map_StaticWebSiteChapter_Identifiers := make(map[*StaticWebSiteChapter]string)
	_ = map_StaticWebSiteChapter_Identifiers

	staticwebsitechapterOrdered := []*StaticWebSiteChapter{}
	for staticwebsitechapter := range stage.StaticWebSiteChapters {
		staticwebsitechapterOrdered = append(staticwebsitechapterOrdered, staticwebsitechapter)
	}
	sort.Slice(staticwebsitechapterOrdered[:], func(i, j int) bool {
		staticwebsitechapteri := staticwebsitechapterOrdered[i]
		staticwebsitechapterj := staticwebsitechapterOrdered[j]
		staticwebsitechapteri_order, oki := stage.StaticWebSiteChapterMap_Staged_Order[staticwebsitechapteri]
		staticwebsitechapterj_order, okj := stage.StaticWebSiteChapterMap_Staged_Order[staticwebsitechapterj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return staticwebsitechapteri_order < staticwebsitechapterj_order
	})
	if len(staticwebsitechapterOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, staticwebsitechapter := range staticwebsitechapterOrdered {

		id = generatesIdentifier("StaticWebSiteChapter", int(stage.StaticWebSiteChapterMap_Staged_Order[staticwebsitechapter]), staticwebsitechapter.Name)
		map_StaticWebSiteChapter_Identifiers[staticwebsitechapter] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StaticWebSiteChapter")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", staticwebsitechapter.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsitechapter.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MarkdownContent")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsitechapter.MarkdownContent))
		initializerStatements += setValueField

	}

	map_StaticWebSiteGeneratedImage_Identifiers := make(map[*StaticWebSiteGeneratedImage]string)
	_ = map_StaticWebSiteGeneratedImage_Identifiers

	staticwebsitegeneratedimageOrdered := []*StaticWebSiteGeneratedImage{}
	for staticwebsitegeneratedimage := range stage.StaticWebSiteGeneratedImages {
		staticwebsitegeneratedimageOrdered = append(staticwebsitegeneratedimageOrdered, staticwebsitegeneratedimage)
	}
	sort.Slice(staticwebsitegeneratedimageOrdered[:], func(i, j int) bool {
		staticwebsitegeneratedimagei := staticwebsitegeneratedimageOrdered[i]
		staticwebsitegeneratedimagej := staticwebsitegeneratedimageOrdered[j]
		staticwebsitegeneratedimagei_order, oki := stage.StaticWebSiteGeneratedImageMap_Staged_Order[staticwebsitegeneratedimagei]
		staticwebsitegeneratedimagej_order, okj := stage.StaticWebSiteGeneratedImageMap_Staged_Order[staticwebsitegeneratedimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return staticwebsitegeneratedimagei_order < staticwebsitegeneratedimagej_order
	})
	if len(staticwebsitegeneratedimageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, staticwebsitegeneratedimage := range staticwebsitegeneratedimageOrdered {

		id = generatesIdentifier("StaticWebSiteGeneratedImage", int(stage.StaticWebSiteGeneratedImageMap_Staged_Order[staticwebsitegeneratedimage]), staticwebsitegeneratedimage.Name)
		map_StaticWebSiteGeneratedImage_Identifiers[staticwebsitegeneratedimage] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StaticWebSiteGeneratedImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", staticwebsitegeneratedimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsitegeneratedimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceDirectoryPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsitegeneratedimage.SourceDirectoryPath))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", staticwebsitegeneratedimage.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", staticwebsitegeneratedimage.Height))
		initializerStatements += setValueField

	}

	map_StaticWebSiteImage_Identifiers := make(map[*StaticWebSiteImage]string)
	_ = map_StaticWebSiteImage_Identifiers

	staticwebsiteimageOrdered := []*StaticWebSiteImage{}
	for staticwebsiteimage := range stage.StaticWebSiteImages {
		staticwebsiteimageOrdered = append(staticwebsiteimageOrdered, staticwebsiteimage)
	}
	sort.Slice(staticwebsiteimageOrdered[:], func(i, j int) bool {
		staticwebsiteimagei := staticwebsiteimageOrdered[i]
		staticwebsiteimagej := staticwebsiteimageOrdered[j]
		staticwebsiteimagei_order, oki := stage.StaticWebSiteImageMap_Staged_Order[staticwebsiteimagei]
		staticwebsiteimagej_order, okj := stage.StaticWebSiteImageMap_Staged_Order[staticwebsiteimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return staticwebsiteimagei_order < staticwebsiteimagej_order
	})
	if len(staticwebsiteimageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, staticwebsiteimage := range staticwebsiteimageOrdered {

		id = generatesIdentifier("StaticWebSiteImage", int(stage.StaticWebSiteImageMap_Staged_Order[staticwebsiteimage]), staticwebsiteimage.Name)
		map_StaticWebSiteImage_Identifiers[staticwebsiteimage] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StaticWebSiteImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", staticwebsiteimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsiteimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceDirectoryPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsiteimage.SourceDirectoryPath))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", staticwebsiteimage.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", staticwebsiteimage.Height))
		initializerStatements += setValueField

	}

	map_StaticWebSiteParagraph_Identifiers := make(map[*StaticWebSiteParagraph]string)
	_ = map_StaticWebSiteParagraph_Identifiers

	staticwebsiteparagraphOrdered := []*StaticWebSiteParagraph{}
	for staticwebsiteparagraph := range stage.StaticWebSiteParagraphs {
		staticwebsiteparagraphOrdered = append(staticwebsiteparagraphOrdered, staticwebsiteparagraph)
	}
	sort.Slice(staticwebsiteparagraphOrdered[:], func(i, j int) bool {
		staticwebsiteparagraphi := staticwebsiteparagraphOrdered[i]
		staticwebsiteparagraphj := staticwebsiteparagraphOrdered[j]
		staticwebsiteparagraphi_order, oki := stage.StaticWebSiteParagraphMap_Staged_Order[staticwebsiteparagraphi]
		staticwebsiteparagraphj_order, okj := stage.StaticWebSiteParagraphMap_Staged_Order[staticwebsiteparagraphj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return staticwebsiteparagraphi_order < staticwebsiteparagraphj_order
	})
	if len(staticwebsiteparagraphOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, staticwebsiteparagraph := range staticwebsiteparagraphOrdered {

		id = generatesIdentifier("StaticWebSiteParagraph", int(stage.StaticWebSiteParagraphMap_Staged_Order[staticwebsiteparagraph]), staticwebsiteparagraph.Name)
		map_StaticWebSiteParagraph_Identifiers[staticwebsiteparagraph] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StaticWebSiteParagraph")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", staticwebsiteparagraph.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsiteparagraph.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LegendMarkdownContent")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staticwebsiteparagraph.LegendMarkdownContent))
		initializerStatements += setValueField

	}

	map_XHTML_CONTENT_Identifiers := make(map[*XHTML_CONTENT]string)
	_ = map_XHTML_CONTENT_Identifiers

	xhtml_contentOrdered := []*XHTML_CONTENT{}
	for xhtml_content := range stage.XHTML_CONTENTs {
		xhtml_contentOrdered = append(xhtml_contentOrdered, xhtml_content)
	}
	sort.Slice(xhtml_contentOrdered[:], func(i, j int) bool {
		xhtml_contenti := xhtml_contentOrdered[i]
		xhtml_contentj := xhtml_contentOrdered[j]
		xhtml_contenti_order, oki := stage.XHTML_CONTENTMap_Staged_Order[xhtml_contenti]
		xhtml_contentj_order, okj := stage.XHTML_CONTENTMap_Staged_Order[xhtml_contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_contenti_order < xhtml_contentj_order
	})
	if len(xhtml_contentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, xhtml_content := range xhtml_contentOrdered {

		id = generatesIdentifier("XHTML_CONTENT", int(stage.XHTML_CONTENTMap_Staged_Order[xhtml_content]), xhtml_content.Name)
		map_XHTML_CONTENT_Identifiers[xhtml_content] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XHTML_CONTENT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_content.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_content.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EnclosedText")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_content.EnclosedText))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PureText")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_content.PureText))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(alternative_idOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ALTERNATIVE_ID instances pointers"
	}
	for _, alternative_id := range alternative_idOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ALTERNATIVE_ID", int(stage.ALTERNATIVE_IDMap_Staged_Order[alternative_id]), alternative_id.Name)
		map_ALTERNATIVE_ID_Identifiers[alternative_id] = id

		// Initialisation of values
	}

	if len(attribute_definition_booleanOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_BOOLEAN instances pointers"
	}
	for _, attribute_definition_boolean := range attribute_definition_booleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_BOOLEAN", int(stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_boolean]), attribute_definition_boolean.Name)
		map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[attribute_definition_boolean] = id

		// Initialisation of values
		if attribute_definition_boolean.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[attribute_definition_boolean.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_boolean.DEFAULT_VALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[attribute_definition_boolean.DEFAULT_VALUE])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_boolean.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DATATYPE_DEFINITION_BOOLEAN_REF_Identifiers[attribute_definition_boolean.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_boolean_renderingOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_BOOLEAN_Rendering instances pointers"
	}
	for _, attribute_definition_boolean_rendering := range attribute_definition_boolean_renderingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_BOOLEAN_Rendering", int(stage.ATTRIBUTE_DEFINITION_BOOLEAN_RenderingMap_Staged_Order[attribute_definition_boolean_rendering]), attribute_definition_boolean_rendering.Name)
		map_ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_Identifiers[attribute_definition_boolean_rendering] = id

		// Initialisation of values
	}

	if len(attribute_definition_dateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_DATE instances pointers"
	}
	for _, attribute_definition_date := range attribute_definition_dateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_DATE", int(stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[attribute_definition_date]), attribute_definition_date.Name)
		map_ATTRIBUTE_DEFINITION_DATE_Identifiers[attribute_definition_date] = id

		// Initialisation of values
		if attribute_definition_date.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[attribute_definition_date.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_date.DEFAULT_VALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_VALUE_DATE_Identifiers[attribute_definition_date.DEFAULT_VALUE])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_date.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DATATYPE_DEFINITION_DATE_REF_Identifiers[attribute_definition_date.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_date_renderingOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_DATE_Rendering instances pointers"
	}
	for _, attribute_definition_date_rendering := range attribute_definition_date_renderingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_DATE_Rendering", int(stage.ATTRIBUTE_DEFINITION_DATE_RenderingMap_Staged_Order[attribute_definition_date_rendering]), attribute_definition_date_rendering.Name)
		map_ATTRIBUTE_DEFINITION_DATE_Rendering_Identifiers[attribute_definition_date_rendering] = id

		// Initialisation of values
	}

	if len(attribute_definition_enumerationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_ENUMERATION instances pointers"
	}
	for _, attribute_definition_enumeration := range attribute_definition_enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_ENUMERATION", int(stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[attribute_definition_enumeration]), attribute_definition_enumeration.Name)
		map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[attribute_definition_enumeration] = id

		// Initialisation of values
		if attribute_definition_enumeration.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[attribute_definition_enumeration.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_enumeration.DEFAULT_VALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[attribute_definition_enumeration.DEFAULT_VALUE])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_enumeration.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DATATYPE_DEFINITION_ENUMERATION_REF_Identifiers[attribute_definition_enumeration.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_enumeration_renderingOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_ENUMERATION_Rendering instances pointers"
	}
	for _, attribute_definition_enumeration_rendering := range attribute_definition_enumeration_renderingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_ENUMERATION_Rendering", int(stage.ATTRIBUTE_DEFINITION_ENUMERATION_RenderingMap_Staged_Order[attribute_definition_enumeration_rendering]), attribute_definition_enumeration_rendering.Name)
		map_ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_Identifiers[attribute_definition_enumeration_rendering] = id

		// Initialisation of values
	}

	if len(attribute_definition_integerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_INTEGER instances pointers"
	}
	for _, attribute_definition_integer := range attribute_definition_integerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_INTEGER", int(stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[attribute_definition_integer]), attribute_definition_integer.Name)
		map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[attribute_definition_integer] = id

		// Initialisation of values
		if attribute_definition_integer.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[attribute_definition_integer.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_integer.DEFAULT_VALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_VALUE_INTEGER_Identifiers[attribute_definition_integer.DEFAULT_VALUE])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_integer.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DATATYPE_DEFINITION_INTEGER_REF_Identifiers[attribute_definition_integer.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_integer_renderingOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_INTEGER_Rendering instances pointers"
	}
	for _, attribute_definition_integer_rendering := range attribute_definition_integer_renderingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_INTEGER_Rendering", int(stage.ATTRIBUTE_DEFINITION_INTEGER_RenderingMap_Staged_Order[attribute_definition_integer_rendering]), attribute_definition_integer_rendering.Name)
		map_ATTRIBUTE_DEFINITION_INTEGER_Rendering_Identifiers[attribute_definition_integer_rendering] = id

		// Initialisation of values
	}

	if len(attribute_definition_realOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_REAL instances pointers"
	}
	for _, attribute_definition_real := range attribute_definition_realOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_REAL", int(stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[attribute_definition_real]), attribute_definition_real.Name)
		map_ATTRIBUTE_DEFINITION_REAL_Identifiers[attribute_definition_real] = id

		// Initialisation of values
		if attribute_definition_real.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[attribute_definition_real.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_real.DEFAULT_VALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_VALUE_REAL_Identifiers[attribute_definition_real.DEFAULT_VALUE])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_real.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DATATYPE_DEFINITION_REAL_REF_Identifiers[attribute_definition_real.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_real_renderingOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_REAL_Rendering instances pointers"
	}
	for _, attribute_definition_real_rendering := range attribute_definition_real_renderingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_REAL_Rendering", int(stage.ATTRIBUTE_DEFINITION_REAL_RenderingMap_Staged_Order[attribute_definition_real_rendering]), attribute_definition_real_rendering.Name)
		map_ATTRIBUTE_DEFINITION_REAL_Rendering_Identifiers[attribute_definition_real_rendering] = id

		// Initialisation of values
	}

	if len(attribute_definition_renderingOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_Rendering instances pointers"
	}
	for _, attribute_definition_rendering := range attribute_definition_renderingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_Rendering", int(stage.ATTRIBUTE_DEFINITION_RenderingMap_Staged_Order[attribute_definition_rendering]), attribute_definition_rendering.Name)
		map_ATTRIBUTE_DEFINITION_Rendering_Identifiers[attribute_definition_rendering] = id

		// Initialisation of values
	}

	if len(attribute_definition_stringOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_STRING instances pointers"
	}
	for _, attribute_definition_string := range attribute_definition_stringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_STRING", int(stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_string]), attribute_definition_string.Name)
		map_ATTRIBUTE_DEFINITION_STRING_Identifiers[attribute_definition_string] = id

		// Initialisation of values
		if attribute_definition_string.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[attribute_definition_string.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_string.DEFAULT_VALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_VALUE_STRING_Identifiers[attribute_definition_string.DEFAULT_VALUE])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_string.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DATATYPE_DEFINITION_STRING_REF_Identifiers[attribute_definition_string.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_string_renderingOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_STRING_Rendering instances pointers"
	}
	for _, attribute_definition_string_rendering := range attribute_definition_string_renderingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_STRING_Rendering", int(stage.ATTRIBUTE_DEFINITION_STRING_RenderingMap_Staged_Order[attribute_definition_string_rendering]), attribute_definition_string_rendering.Name)
		map_ATTRIBUTE_DEFINITION_STRING_Rendering_Identifiers[attribute_definition_string_rendering] = id

		// Initialisation of values
	}

	if len(attribute_definition_xhtmlOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_XHTML instances pointers"
	}
	for _, attribute_definition_xhtml := range attribute_definition_xhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_XHTML", int(stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[attribute_definition_xhtml]), attribute_definition_xhtml.Name)
		map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[attribute_definition_xhtml] = id

		// Initialisation of values
		if attribute_definition_xhtml.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[attribute_definition_xhtml.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_xhtml.DEFAULT_VALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_VALUE_XHTML_Identifiers[attribute_definition_xhtml.DEFAULT_VALUE])
			pointersInitializesStatements += setPointerField
		}

		if attribute_definition_xhtml.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DATATYPE_DEFINITION_XHTML_REF_Identifiers[attribute_definition_xhtml.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_xhtml_renderingOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_XHTML_Rendering instances pointers"
	}
	for _, attribute_definition_xhtml_rendering := range attribute_definition_xhtml_renderingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_XHTML_Rendering", int(stage.ATTRIBUTE_DEFINITION_XHTML_RenderingMap_Staged_Order[attribute_definition_xhtml_rendering]), attribute_definition_xhtml_rendering.Name)
		map_ATTRIBUTE_DEFINITION_XHTML_Rendering_Identifiers[attribute_definition_xhtml_rendering] = id

		// Initialisation of values
	}

	if len(attribute_value_booleanOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_BOOLEAN instances pointers"
	}
	for _, attribute_value_boolean := range attribute_value_booleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_BOOLEAN", int(stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[attribute_value_boolean]), attribute_value_boolean.Name)
		map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[attribute_value_boolean] = id

		// Initialisation of values
		if attribute_value_boolean.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Identifiers[attribute_value_boolean.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_value_dateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_DATE instances pointers"
	}
	for _, attribute_value_date := range attribute_value_dateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_DATE", int(stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[attribute_value_date]), attribute_value_date.Name)
		map_ATTRIBUTE_VALUE_DATE_Identifiers[attribute_value_date] = id

		// Initialisation of values
		if attribute_value_date.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_DEFINITION_DATE_REF_Identifiers[attribute_value_date.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_value_enumerationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_ENUMERATION instances pointers"
	}
	for _, attribute_value_enumeration := range attribute_value_enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_ENUMERATION", int(stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[attribute_value_enumeration]), attribute_value_enumeration.Name)
		map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[attribute_value_enumeration] = id

		// Initialisation of values
		if attribute_value_enumeration.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Identifiers[attribute_value_enumeration.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

		if attribute_value_enumeration.VALUES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ENUM_VALUE_REF_Identifiers[attribute_value_enumeration.VALUES])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_value_integerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_INTEGER instances pointers"
	}
	for _, attribute_value_integer := range attribute_value_integerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_INTEGER", int(stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[attribute_value_integer]), attribute_value_integer.Name)
		map_ATTRIBUTE_VALUE_INTEGER_Identifiers[attribute_value_integer] = id

		// Initialisation of values
		if attribute_value_integer.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Identifiers[attribute_value_integer.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_value_realOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_REAL instances pointers"
	}
	for _, attribute_value_real := range attribute_value_realOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_REAL", int(stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[attribute_value_real]), attribute_value_real.Name)
		map_ATTRIBUTE_VALUE_REAL_Identifiers[attribute_value_real] = id

		// Initialisation of values
		if attribute_value_real.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_DEFINITION_REAL_REF_Identifiers[attribute_value_real.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_value_stringOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_STRING instances pointers"
	}
	for _, attribute_value_string := range attribute_value_stringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_STRING", int(stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[attribute_value_string]), attribute_value_string.Name)
		map_ATTRIBUTE_VALUE_STRING_Identifiers[attribute_value_string] = id

		// Initialisation of values
		if attribute_value_string.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_DEFINITION_STRING_REF_Identifiers[attribute_value_string.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_value_xhtmlOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_XHTML instances pointers"
	}
	for _, attribute_value_xhtml := range attribute_value_xhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_XHTML", int(stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtml]), attribute_value_xhtml.Name)
		map_ATTRIBUTE_VALUE_XHTML_Identifiers[attribute_value_xhtml] = id

		// Initialisation of values
		if attribute_value_xhtml.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_DEFINITION_XHTML_REF_Identifiers[attribute_value_xhtml.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

		if attribute_value_xhtml.THE_VALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THE_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XHTML_CONTENT_Identifiers[attribute_value_xhtml.THE_VALUE])
			pointersInitializesStatements += setPointerField
		}

		if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THE_ORIGINAL_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XHTML_CONTENT_Identifiers[attribute_value_xhtml.THE_ORIGINAL_VALUE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_alternative_idOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ALTERNATIVE_ID instances pointers"
	}
	for _, a_alternative_id := range a_alternative_idOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ALTERNATIVE_ID", int(stage.A_ALTERNATIVE_IDMap_Staged_Order[a_alternative_id]), a_alternative_id.Name)
		map_A_ALTERNATIVE_ID_Identifiers[a_alternative_id] = id

		// Initialisation of values
		if a_alternative_id.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[a_alternative_id.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_attribute_definition_boolean_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_DEFINITION_BOOLEAN_REF instances pointers"
	}
	for _, a_attribute_definition_boolean_ref := range a_attribute_definition_boolean_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_BOOLEAN_REF", int(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_attribute_definition_boolean_ref]), a_attribute_definition_boolean_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Identifiers[a_attribute_definition_boolean_ref] = id

		// Initialisation of values
	}

	if len(a_attribute_definition_date_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_DEFINITION_DATE_REF instances pointers"
	}
	for _, a_attribute_definition_date_ref := range a_attribute_definition_date_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_DATE_REF", int(stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order[a_attribute_definition_date_ref]), a_attribute_definition_date_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_DATE_REF_Identifiers[a_attribute_definition_date_ref] = id

		// Initialisation of values
	}

	if len(a_attribute_definition_enumeration_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_DEFINITION_ENUMERATION_REF instances pointers"
	}
	for _, a_attribute_definition_enumeration_ref := range a_attribute_definition_enumeration_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_ENUMERATION_REF", int(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_attribute_definition_enumeration_ref]), a_attribute_definition_enumeration_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Identifiers[a_attribute_definition_enumeration_ref] = id

		// Initialisation of values
	}

	if len(a_attribute_definition_integer_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_DEFINITION_INTEGER_REF instances pointers"
	}
	for _, a_attribute_definition_integer_ref := range a_attribute_definition_integer_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_INTEGER_REF", int(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order[a_attribute_definition_integer_ref]), a_attribute_definition_integer_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Identifiers[a_attribute_definition_integer_ref] = id

		// Initialisation of values
	}

	if len(a_attribute_definition_real_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_DEFINITION_REAL_REF instances pointers"
	}
	for _, a_attribute_definition_real_ref := range a_attribute_definition_real_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_REAL_REF", int(stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order[a_attribute_definition_real_ref]), a_attribute_definition_real_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_REAL_REF_Identifiers[a_attribute_definition_real_ref] = id

		// Initialisation of values
	}

	if len(a_attribute_definition_string_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_DEFINITION_STRING_REF instances pointers"
	}
	for _, a_attribute_definition_string_ref := range a_attribute_definition_string_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_STRING_REF", int(stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order[a_attribute_definition_string_ref]), a_attribute_definition_string_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_STRING_REF_Identifiers[a_attribute_definition_string_ref] = id

		// Initialisation of values
	}

	if len(a_attribute_definition_xhtml_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_DEFINITION_XHTML_REF instances pointers"
	}
	for _, a_attribute_definition_xhtml_ref := range a_attribute_definition_xhtml_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_DEFINITION_XHTML_REF", int(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order[a_attribute_definition_xhtml_ref]), a_attribute_definition_xhtml_ref.Name)
		map_A_ATTRIBUTE_DEFINITION_XHTML_REF_Identifiers[a_attribute_definition_xhtml_ref] = id

		// Initialisation of values
	}

	if len(a_attribute_value_booleanOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_VALUE_BOOLEAN instances pointers"
	}
	for _, a_attribute_value_boolean := range a_attribute_value_booleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_BOOLEAN", int(stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[a_attribute_value_boolean]), a_attribute_value_boolean.Name)
		map_A_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[a_attribute_value_boolean] = id

		// Initialisation of values
		for _, _attribute_value_boolean := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[_attribute_value_boolean])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_attribute_value_dateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_VALUE_DATE instances pointers"
	}
	for _, a_attribute_value_date := range a_attribute_value_dateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_DATE", int(stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order[a_attribute_value_date]), a_attribute_value_date.Name)
		map_A_ATTRIBUTE_VALUE_DATE_Identifiers[a_attribute_value_date] = id

		// Initialisation of values
		for _, _attribute_value_date := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_DATE_Identifiers[_attribute_value_date])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_attribute_value_enumerationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_VALUE_ENUMERATION instances pointers"
	}
	for _, a_attribute_value_enumeration := range a_attribute_value_enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_ENUMERATION", int(stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[a_attribute_value_enumeration]), a_attribute_value_enumeration.Name)
		map_A_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[a_attribute_value_enumeration] = id

		// Initialisation of values
		for _, _attribute_value_enumeration := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[_attribute_value_enumeration])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_attribute_value_integerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_VALUE_INTEGER instances pointers"
	}
	for _, a_attribute_value_integer := range a_attribute_value_integerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_INTEGER", int(stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[a_attribute_value_integer]), a_attribute_value_integer.Name)
		map_A_ATTRIBUTE_VALUE_INTEGER_Identifiers[a_attribute_value_integer] = id

		// Initialisation of values
		for _, _attribute_value_integer := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_INTEGER_Identifiers[_attribute_value_integer])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_attribute_value_realOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_VALUE_REAL instances pointers"
	}
	for _, a_attribute_value_real := range a_attribute_value_realOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_REAL", int(stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order[a_attribute_value_real]), a_attribute_value_real.Name)
		map_A_ATTRIBUTE_VALUE_REAL_Identifiers[a_attribute_value_real] = id

		// Initialisation of values
		for _, _attribute_value_real := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_REAL_Identifiers[_attribute_value_real])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_attribute_value_stringOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_VALUE_STRING instances pointers"
	}
	for _, a_attribute_value_string := range a_attribute_value_stringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_STRING", int(stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order[a_attribute_value_string]), a_attribute_value_string.Name)
		map_A_ATTRIBUTE_VALUE_STRING_Identifiers[a_attribute_value_string] = id

		// Initialisation of values
		for _, _attribute_value_string := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_STRING_Identifiers[_attribute_value_string])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_attribute_value_xhtmlOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_VALUE_XHTML instances pointers"
	}
	for _, a_attribute_value_xhtml := range a_attribute_value_xhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_XHTML", int(stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[a_attribute_value_xhtml]), a_attribute_value_xhtml.Name)
		map_A_ATTRIBUTE_VALUE_XHTML_Identifiers[a_attribute_value_xhtml] = id

		// Initialisation of values
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_XHTML_Identifiers[_attribute_value_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_attribute_value_xhtml_1Ordered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ATTRIBUTE_VALUE_XHTML_1 instances pointers"
	}
	for _, a_attribute_value_xhtml_1 := range a_attribute_value_xhtml_1Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ATTRIBUTE_VALUE_XHTML_1", int(stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order[a_attribute_value_xhtml_1]), a_attribute_value_xhtml_1.Name)
		map_A_ATTRIBUTE_VALUE_XHTML_1_Identifiers[a_attribute_value_xhtml_1] = id

		// Initialisation of values
		for _, _attribute_value_boolean := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[_attribute_value_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_date := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_DATE_Identifiers[_attribute_value_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_enumeration := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[_attribute_value_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_integer := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_INTEGER_Identifiers[_attribute_value_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_real := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_REAL_Identifiers[_attribute_value_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_string := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_STRING_Identifiers[_attribute_value_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_xhtml := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_XHTML_Identifiers[_attribute_value_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_childrenOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_CHILDREN instances pointers"
	}
	for _, a_children := range a_childrenOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_CHILDREN", int(stage.A_CHILDRENMap_Staged_Order[a_children]), a_children.Name)
		map_A_CHILDREN_Identifiers[a_children] = id

		// Initialisation of values
		for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_HIERARCHY")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_HIERARCHY_Identifiers[_spec_hierarchy])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_core_contentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_CORE_CONTENT instances pointers"
	}
	for _, a_core_content := range a_core_contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_CORE_CONTENT", int(stage.A_CORE_CONTENTMap_Staged_Order[a_core_content]), a_core_content.Name)
		map_A_CORE_CONTENT_Identifiers[a_core_content] = id

		// Initialisation of values
		if a_core_content.REQ_IF_CONTENT != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "REQ_IF_CONTENT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQ_IF_CONTENT_Identifiers[a_core_content.REQ_IF_CONTENT])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_datatypesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_DATATYPES instances pointers"
	}
	for _, a_datatypes := range a_datatypesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DATATYPES", int(stage.A_DATATYPESMap_Staged_Order[a_datatypes]), a_datatypes.Name)
		map_A_DATATYPES_Identifiers[a_datatypes] = id

		// Initialisation of values
		for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_BOOLEAN_Identifiers[_datatype_definition_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_DATE_Identifiers[_datatype_definition_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_ENUMERATION_Identifiers[_datatype_definition_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_INTEGER_Identifiers[_datatype_definition_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_REAL_Identifiers[_datatype_definition_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_STRING_Identifiers[_datatype_definition_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_XHTML_Identifiers[_datatype_definition_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_datatype_definition_boolean_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_DATATYPE_DEFINITION_BOOLEAN_REF instances pointers"
	}
	for _, a_datatype_definition_boolean_ref := range a_datatype_definition_boolean_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DATATYPE_DEFINITION_BOOLEAN_REF", int(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_datatype_definition_boolean_ref]), a_datatype_definition_boolean_ref.Name)
		map_A_DATATYPE_DEFINITION_BOOLEAN_REF_Identifiers[a_datatype_definition_boolean_ref] = id

		// Initialisation of values
	}

	if len(a_datatype_definition_date_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_DATATYPE_DEFINITION_DATE_REF instances pointers"
	}
	for _, a_datatype_definition_date_ref := range a_datatype_definition_date_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DATATYPE_DEFINITION_DATE_REF", int(stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order[a_datatype_definition_date_ref]), a_datatype_definition_date_ref.Name)
		map_A_DATATYPE_DEFINITION_DATE_REF_Identifiers[a_datatype_definition_date_ref] = id

		// Initialisation of values
	}

	if len(a_datatype_definition_enumeration_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_DATATYPE_DEFINITION_ENUMERATION_REF instances pointers"
	}
	for _, a_datatype_definition_enumeration_ref := range a_datatype_definition_enumeration_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DATATYPE_DEFINITION_ENUMERATION_REF", int(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_datatype_definition_enumeration_ref]), a_datatype_definition_enumeration_ref.Name)
		map_A_DATATYPE_DEFINITION_ENUMERATION_REF_Identifiers[a_datatype_definition_enumeration_ref] = id

		// Initialisation of values
	}

	if len(a_datatype_definition_integer_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_DATATYPE_DEFINITION_INTEGER_REF instances pointers"
	}
	for _, a_datatype_definition_integer_ref := range a_datatype_definition_integer_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DATATYPE_DEFINITION_INTEGER_REF", int(stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order[a_datatype_definition_integer_ref]), a_datatype_definition_integer_ref.Name)
		map_A_DATATYPE_DEFINITION_INTEGER_REF_Identifiers[a_datatype_definition_integer_ref] = id

		// Initialisation of values
	}

	if len(a_datatype_definition_real_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_DATATYPE_DEFINITION_REAL_REF instances pointers"
	}
	for _, a_datatype_definition_real_ref := range a_datatype_definition_real_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DATATYPE_DEFINITION_REAL_REF", int(stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order[a_datatype_definition_real_ref]), a_datatype_definition_real_ref.Name)
		map_A_DATATYPE_DEFINITION_REAL_REF_Identifiers[a_datatype_definition_real_ref] = id

		// Initialisation of values
	}

	if len(a_datatype_definition_string_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_DATATYPE_DEFINITION_STRING_REF instances pointers"
	}
	for _, a_datatype_definition_string_ref := range a_datatype_definition_string_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DATATYPE_DEFINITION_STRING_REF", int(stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order[a_datatype_definition_string_ref]), a_datatype_definition_string_ref.Name)
		map_A_DATATYPE_DEFINITION_STRING_REF_Identifiers[a_datatype_definition_string_ref] = id

		// Initialisation of values
	}

	if len(a_datatype_definition_xhtml_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_DATATYPE_DEFINITION_XHTML_REF instances pointers"
	}
	for _, a_datatype_definition_xhtml_ref := range a_datatype_definition_xhtml_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DATATYPE_DEFINITION_XHTML_REF", int(stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order[a_datatype_definition_xhtml_ref]), a_datatype_definition_xhtml_ref.Name)
		map_A_DATATYPE_DEFINITION_XHTML_REF_Identifiers[a_datatype_definition_xhtml_ref] = id

		// Initialisation of values
	}

	if len(a_editable_attsOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_EDITABLE_ATTS instances pointers"
	}
	for _, a_editable_atts := range a_editable_attsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_EDITABLE_ATTS", int(stage.A_EDITABLE_ATTSMap_Staged_Order[a_editable_atts]), a_editable_atts.Name)
		map_A_EDITABLE_ATTS_Identifiers[a_editable_atts] = id

		// Initialisation of values
	}

	if len(a_enum_value_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_ENUM_VALUE_REF instances pointers"
	}
	for _, a_enum_value_ref := range a_enum_value_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ENUM_VALUE_REF", int(stage.A_ENUM_VALUE_REFMap_Staged_Order[a_enum_value_ref]), a_enum_value_ref.Name)
		map_A_ENUM_VALUE_REF_Identifiers[a_enum_value_ref] = id

		// Initialisation of values
	}

	if len(a_objectOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_OBJECT instances pointers"
	}
	for _, a_object := range a_objectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_OBJECT", int(stage.A_OBJECTMap_Staged_Order[a_object]), a_object.Name)
		map_A_OBJECT_Identifiers[a_object] = id

		// Initialisation of values
	}

	if len(a_propertiesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_PROPERTIES instances pointers"
	}
	for _, a_properties := range a_propertiesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_PROPERTIES", int(stage.A_PROPERTIESMap_Staged_Order[a_properties]), a_properties.Name)
		map_A_PROPERTIES_Identifiers[a_properties] = id

		// Initialisation of values
		if a_properties.EMBEDDED_VALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EMBEDDED_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_EMBEDDED_VALUE_Identifiers[a_properties.EMBEDDED_VALUE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_relation_group_type_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_RELATION_GROUP_TYPE_REF instances pointers"
	}
	for _, a_relation_group_type_ref := range a_relation_group_type_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_RELATION_GROUP_TYPE_REF", int(stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_ref]), a_relation_group_type_ref.Name)
		map_A_RELATION_GROUP_TYPE_REF_Identifiers[a_relation_group_type_ref] = id

		// Initialisation of values
	}

	if len(a_source_1Ordered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SOURCE_1 instances pointers"
	}
	for _, a_source_1 := range a_source_1Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SOURCE_1", int(stage.A_SOURCE_1Map_Staged_Order[a_source_1]), a_source_1.Name)
		map_A_SOURCE_1_Identifiers[a_source_1] = id

		// Initialisation of values
	}

	if len(a_source_specification_1Ordered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SOURCE_SPECIFICATION_1 instances pointers"
	}
	for _, a_source_specification_1 := range a_source_specification_1Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SOURCE_SPECIFICATION_1", int(stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order[a_source_specification_1]), a_source_specification_1.Name)
		map_A_SOURCE_SPECIFICATION_1_Identifiers[a_source_specification_1] = id

		// Initialisation of values
	}

	if len(a_specificationsOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPECIFICATIONS instances pointers"
	}
	for _, a_specifications := range a_specificationsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPECIFICATIONS", int(stage.A_SPECIFICATIONSMap_Staged_Order[a_specifications]), a_specifications.Name)
		map_A_SPECIFICATIONS_Identifiers[a_specifications] = id

		// Initialisation of values
		for _, _specification := range a_specifications.SPECIFICATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFICATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECIFICATION_Identifiers[_specification])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_specification_type_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPECIFICATION_TYPE_REF instances pointers"
	}
	for _, a_specification_type_ref := range a_specification_type_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPECIFICATION_TYPE_REF", int(stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order[a_specification_type_ref]), a_specification_type_ref.Name)
		map_A_SPECIFICATION_TYPE_REF_Identifiers[a_specification_type_ref] = id

		// Initialisation of values
	}

	if len(a_specified_valuesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPECIFIED_VALUES instances pointers"
	}
	for _, a_specified_values := range a_specified_valuesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPECIFIED_VALUES", int(stage.A_SPECIFIED_VALUESMap_Staged_Order[a_specified_values]), a_specified_values.Name)
		map_A_SPECIFIED_VALUES_Identifiers[a_specified_values] = id

		// Initialisation of values
		for _, _enum_value := range a_specified_values.ENUM_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ENUM_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ENUM_VALUE_Identifiers[_enum_value])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_spec_attributesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPEC_ATTRIBUTES instances pointers"
	}
	for _, a_spec_attributes := range a_spec_attributesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_ATTRIBUTES", int(stage.A_SPEC_ATTRIBUTESMap_Staged_Order[a_spec_attributes]), a_spec_attributes.Name)
		map_A_SPEC_ATTRIBUTES_Identifiers[a_spec_attributes] = id

		// Initialisation of values
		for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[_attribute_definition_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_DATE_Identifiers[_attribute_definition_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[_attribute_definition_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[_attribute_definition_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_REAL_Identifiers[_attribute_definition_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_STRING_Identifiers[_attribute_definition_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[_attribute_definition_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_spec_objectsOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPEC_OBJECTS instances pointers"
	}
	for _, a_spec_objects := range a_spec_objectsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_OBJECTS", int(stage.A_SPEC_OBJECTSMap_Staged_Order[a_spec_objects]), a_spec_objects.Name)
		map_A_SPEC_OBJECTS_Identifiers[a_spec_objects] = id

		// Initialisation of values
		for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_OBJECT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_OBJECT_Identifiers[_spec_object])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_spec_object_type_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPEC_OBJECT_TYPE_REF instances pointers"
	}
	for _, a_spec_object_type_ref := range a_spec_object_type_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_OBJECT_TYPE_REF", int(stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order[a_spec_object_type_ref]), a_spec_object_type_ref.Name)
		map_A_SPEC_OBJECT_TYPE_REF_Identifiers[a_spec_object_type_ref] = id

		// Initialisation of values
	}

	if len(a_spec_relationsOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPEC_RELATIONS instances pointers"
	}
	for _, a_spec_relations := range a_spec_relationsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_RELATIONS", int(stage.A_SPEC_RELATIONSMap_Staged_Order[a_spec_relations]), a_spec_relations.Name)
		map_A_SPEC_RELATIONS_Identifiers[a_spec_relations] = id

		// Initialisation of values
		for _, _spec_relation := range a_spec_relations.SPEC_RELATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_RELATION_Identifiers[_spec_relation])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_spec_relation_groupsOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPEC_RELATION_GROUPS instances pointers"
	}
	for _, a_spec_relation_groups := range a_spec_relation_groupsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_RELATION_GROUPS", int(stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groups]), a_spec_relation_groups.Name)
		map_A_SPEC_RELATION_GROUPS_Identifiers[a_spec_relation_groups] = id

		// Initialisation of values
		for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RELATION_GROUP")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RELATION_GROUP_Identifiers[_relation_group])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_spec_relation_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPEC_RELATION_REF instances pointers"
	}
	for _, a_spec_relation_ref := range a_spec_relation_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_RELATION_REF", int(stage.A_SPEC_RELATION_REFMap_Staged_Order[a_spec_relation_ref]), a_spec_relation_ref.Name)
		map_A_SPEC_RELATION_REF_Identifiers[a_spec_relation_ref] = id

		// Initialisation of values
	}

	if len(a_spec_relation_type_refOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPEC_RELATION_TYPE_REF instances pointers"
	}
	for _, a_spec_relation_type_ref := range a_spec_relation_type_refOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_RELATION_TYPE_REF", int(stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order[a_spec_relation_type_ref]), a_spec_relation_type_ref.Name)
		map_A_SPEC_RELATION_TYPE_REF_Identifiers[a_spec_relation_type_ref] = id

		// Initialisation of values
	}

	if len(a_spec_typesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_SPEC_TYPES instances pointers"
	}
	for _, a_spec_types := range a_spec_typesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_TYPES", int(stage.A_SPEC_TYPESMap_Staged_Order[a_spec_types]), a_spec_types.Name)
		map_A_SPEC_TYPES_Identifiers[a_spec_types] = id

		// Initialisation of values
		for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RELATION_GROUP_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RELATION_GROUP_TYPE_Identifiers[_relation_group_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_OBJECT_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_OBJECT_TYPE_Identifiers[_spec_object_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATION_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_RELATION_TYPE_Identifiers[_spec_relation_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFICATION_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECIFICATION_TYPE_Identifiers[_specification_type])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_the_headerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_THE_HEADER instances pointers"
	}
	for _, a_the_header := range a_the_headerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_THE_HEADER", int(stage.A_THE_HEADERMap_Staged_Order[a_the_header]), a_the_header.Name)
		map_A_THE_HEADER_Identifiers[a_the_header] = id

		// Initialisation of values
		if a_the_header.REQ_IF_HEADER != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "REQ_IF_HEADER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQ_IF_HEADER_Identifiers[a_the_header.REQ_IF_HEADER])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(a_tool_extensionsOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of A_TOOL_EXTENSIONS instances pointers"
	}
	for _, a_tool_extensions := range a_tool_extensionsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TOOL_EXTENSIONS", int(stage.A_TOOL_EXTENSIONSMap_Staged_Order[a_tool_extensions]), a_tool_extensions.Name)
		map_A_TOOL_EXTENSIONS_Identifiers[a_tool_extensions] = id

		// Initialisation of values
		for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "REQ_IF_TOOL_EXTENSION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQ_IF_TOOL_EXTENSION_Identifiers[_req_if_tool_extension])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_booleanOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_BOOLEAN instances pointers"
	}
	for _, datatype_definition_boolean := range datatype_definition_booleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_BOOLEAN", int(stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[datatype_definition_boolean]), datatype_definition_boolean.Name)
		map_DATATYPE_DEFINITION_BOOLEAN_Identifiers[datatype_definition_boolean] = id

		// Initialisation of values
		if datatype_definition_boolean.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[datatype_definition_boolean.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_dateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_DATE instances pointers"
	}
	for _, datatype_definition_date := range datatype_definition_dateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_DATE", int(stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[datatype_definition_date]), datatype_definition_date.Name)
		map_DATATYPE_DEFINITION_DATE_Identifiers[datatype_definition_date] = id

		// Initialisation of values
		if datatype_definition_date.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[datatype_definition_date.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_enumerationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_ENUMERATION instances pointers"
	}
	for _, datatype_definition_enumeration := range datatype_definition_enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_ENUMERATION", int(stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[datatype_definition_enumeration]), datatype_definition_enumeration.Name)
		map_DATATYPE_DEFINITION_ENUMERATION_Identifiers[datatype_definition_enumeration] = id

		// Initialisation of values
		if datatype_definition_enumeration.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[datatype_definition_enumeration.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if datatype_definition_enumeration.SPECIFIED_VALUES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFIED_VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPECIFIED_VALUES_Identifiers[datatype_definition_enumeration.SPECIFIED_VALUES])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_integerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_INTEGER instances pointers"
	}
	for _, datatype_definition_integer := range datatype_definition_integerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_INTEGER", int(stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[datatype_definition_integer]), datatype_definition_integer.Name)
		map_DATATYPE_DEFINITION_INTEGER_Identifiers[datatype_definition_integer] = id

		// Initialisation of values
		if datatype_definition_integer.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[datatype_definition_integer.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_realOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_REAL instances pointers"
	}
	for _, datatype_definition_real := range datatype_definition_realOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_REAL", int(stage.DATATYPE_DEFINITION_REALMap_Staged_Order[datatype_definition_real]), datatype_definition_real.Name)
		map_DATATYPE_DEFINITION_REAL_Identifiers[datatype_definition_real] = id

		// Initialisation of values
		if datatype_definition_real.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[datatype_definition_real.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_stringOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_STRING instances pointers"
	}
	for _, datatype_definition_string := range datatype_definition_stringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_STRING", int(stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[datatype_definition_string]), datatype_definition_string.Name)
		map_DATATYPE_DEFINITION_STRING_Identifiers[datatype_definition_string] = id

		// Initialisation of values
		if datatype_definition_string.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[datatype_definition_string.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_xhtmlOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_XHTML instances pointers"
	}
	for _, datatype_definition_xhtml := range datatype_definition_xhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_XHTML", int(stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[datatype_definition_xhtml]), datatype_definition_xhtml.Name)
		map_DATATYPE_DEFINITION_XHTML_Identifiers[datatype_definition_xhtml] = id

		// Initialisation of values
		if datatype_definition_xhtml.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[datatype_definition_xhtml.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(embedded_valueOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of EMBEDDED_VALUE instances pointers"
	}
	for _, embedded_value := range embedded_valueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("EMBEDDED_VALUE", int(stage.EMBEDDED_VALUEMap_Staged_Order[embedded_value]), embedded_value.Name)
		map_EMBEDDED_VALUE_Identifiers[embedded_value] = id

		// Initialisation of values
	}

	if len(enum_valueOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ENUM_VALUE instances pointers"
	}
	for _, enum_value := range enum_valueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ENUM_VALUE", int(stage.ENUM_VALUEMap_Staged_Order[enum_value]), enum_value.Name)
		map_ENUM_VALUE_Identifiers[enum_value] = id

		// Initialisation of values
		if enum_value.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[enum_value.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if enum_value.PROPERTIES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "PROPERTIES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_PROPERTIES_Identifiers[enum_value.PROPERTIES])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(embeddedjpgimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of EmbeddedJpgImage instances pointers"
	}
	for _, embeddedjpgimage := range embeddedjpgimageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("EmbeddedJpgImage", int(stage.EmbeddedJpgImageMap_Staged_Order[embeddedjpgimage]), embeddedjpgimage.Name)
		map_EmbeddedJpgImage_Identifiers[embeddedjpgimage] = id

		// Initialisation of values
	}

	if len(embeddedpngimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of EmbeddedPngImage instances pointers"
	}
	for _, embeddedpngimage := range embeddedpngimageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("EmbeddedPngImage", int(stage.EmbeddedPngImageMap_Staged_Order[embeddedpngimage]), embeddedpngimage.Name)
		map_EmbeddedPngImage_Identifiers[embeddedpngimage] = id

		// Initialisation of values
	}

	if len(embeddedsvgimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of EmbeddedSvgImage instances pointers"
	}
	for _, embeddedsvgimage := range embeddedsvgimageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("EmbeddedSvgImage", int(stage.EmbeddedSvgImageMap_Staged_Order[embeddedsvgimage]), embeddedsvgimage.Name)
		map_EmbeddedSvgImage_Identifiers[embeddedsvgimage] = id

		// Initialisation of values
	}

	if len(killOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Kill instances pointers"
	}
	for _, kill := range killOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Kill", int(stage.KillMap_Staged_Order[kill]), kill.Name)
		map_Kill_Identifiers[kill] = id

		// Initialisation of values
	}

	if len(map_identifier_boolOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Map_identifier_bool instances pointers"
	}
	for _, map_identifier_bool := range map_identifier_boolOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Map_identifier_bool", int(stage.Map_identifier_boolMap_Staged_Order[map_identifier_bool]), map_identifier_bool.Name)
		map_Map_identifier_bool_Identifiers[map_identifier_bool] = id

		// Initialisation of values
	}

	if len(relation_groupOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of RELATION_GROUP instances pointers"
	}
	for _, relation_group := range relation_groupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RELATION_GROUP", int(stage.RELATION_GROUPMap_Staged_Order[relation_group]), relation_group.Name)
		map_RELATION_GROUP_Identifiers[relation_group] = id

		// Initialisation of values
		if relation_group.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[relation_group.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if relation_group.SOURCE_SPECIFICATION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SOURCE_SPECIFICATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SOURCE_SPECIFICATION_1_Identifiers[relation_group.SOURCE_SPECIFICATION])
			pointersInitializesStatements += setPointerField
		}

		if relation_group.SPEC_RELATIONS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_RELATION_REF_Identifiers[relation_group.SPEC_RELATIONS])
			pointersInitializesStatements += setPointerField
		}

		if relation_group.TARGET_SPECIFICATION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TARGET_SPECIFICATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SOURCE_SPECIFICATION_1_Identifiers[relation_group.TARGET_SPECIFICATION])
			pointersInitializesStatements += setPointerField
		}

		if relation_group.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_RELATION_GROUP_TYPE_REF_Identifiers[relation_group.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(relation_group_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of RELATION_GROUP_TYPE instances pointers"
	}
	for _, relation_group_type := range relation_group_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RELATION_GROUP_TYPE", int(stage.RELATION_GROUP_TYPEMap_Staged_Order[relation_group_type]), relation_group_type.Name)
		map_RELATION_GROUP_TYPE_Identifiers[relation_group_type] = id

		// Initialisation of values
		if relation_group_type.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[relation_group_type.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if relation_group_type.SPEC_ATTRIBUTES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_ATTRIBUTES_Identifiers[relation_group_type.SPEC_ATTRIBUTES])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(req_ifOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of REQ_IF instances pointers"
	}
	for _, req_if := range req_ifOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF", int(stage.REQ_IFMap_Staged_Order[req_if]), req_if.Name)
		map_REQ_IF_Identifiers[req_if] = id

		// Initialisation of values
		if req_if.THE_HEADER != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THE_HEADER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_THE_HEADER_Identifiers[req_if.THE_HEADER])
			pointersInitializesStatements += setPointerField
		}

		if req_if.CORE_CONTENT != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CORE_CONTENT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_CORE_CONTENT_Identifiers[req_if.CORE_CONTENT])
			pointersInitializesStatements += setPointerField
		}

		if req_if.TOOL_EXTENSIONS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TOOL_EXTENSIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TOOL_EXTENSIONS_Identifiers[req_if.TOOL_EXTENSIONS])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(req_if_contentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of REQ_IF_CONTENT instances pointers"
	}
	for _, req_if_content := range req_if_contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF_CONTENT", int(stage.REQ_IF_CONTENTMap_Staged_Order[req_if_content]), req_if_content.Name)
		map_REQ_IF_CONTENT_Identifiers[req_if_content] = id

		// Initialisation of values
		if req_if_content.DATATYPES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DATATYPES_Identifiers[req_if_content.DATATYPES])
			pointersInitializesStatements += setPointerField
		}

		if req_if_content.SPEC_TYPES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_TYPES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_TYPES_Identifiers[req_if_content.SPEC_TYPES])
			pointersInitializesStatements += setPointerField
		}

		if req_if_content.SPEC_OBJECTS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_OBJECTS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_OBJECTS_Identifiers[req_if_content.SPEC_OBJECTS])
			pointersInitializesStatements += setPointerField
		}

		if req_if_content.SPEC_RELATIONS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_RELATIONS_Identifiers[req_if_content.SPEC_RELATIONS])
			pointersInitializesStatements += setPointerField
		}

		if req_if_content.SPECIFICATIONS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFICATIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPECIFICATIONS_Identifiers[req_if_content.SPECIFICATIONS])
			pointersInitializesStatements += setPointerField
		}

		if req_if_content.SPEC_RELATION_GROUPS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATION_GROUPS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_RELATION_GROUPS_Identifiers[req_if_content.SPEC_RELATION_GROUPS])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(req_if_headerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of REQ_IF_HEADER instances pointers"
	}
	for _, req_if_header := range req_if_headerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF_HEADER", int(stage.REQ_IF_HEADERMap_Staged_Order[req_if_header]), req_if_header.Name)
		map_REQ_IF_HEADER_Identifiers[req_if_header] = id

		// Initialisation of values
	}

	if len(req_if_tool_extensionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of REQ_IF_TOOL_EXTENSION instances pointers"
	}
	for _, req_if_tool_extension := range req_if_tool_extensionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF_TOOL_EXTENSION", int(stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[req_if_tool_extension]), req_if_tool_extension.Name)
		map_REQ_IF_TOOL_EXTENSION_Identifiers[req_if_tool_extension] = id

		// Initialisation of values
	}

	if len(specificationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPECIFICATION instances pointers"
	}
	for _, specification := range specificationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFICATION", int(stage.SPECIFICATIONMap_Staged_Order[specification]), specification.Name)
		map_SPECIFICATION_Identifiers[specification] = id

		// Initialisation of values
		if specification.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[specification.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if specification.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPECIFICATION_TYPE_REF_Identifiers[specification.TYPE])
			pointersInitializesStatements += setPointerField
		}

		if specification.CHILDREN != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CHILDREN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_CHILDREN_Identifiers[specification.CHILDREN])
			pointersInitializesStatements += setPointerField
		}

		if specification.VALUES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_VALUE_XHTML_1_Identifiers[specification.VALUES])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(specification_renderingOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPECIFICATION_Rendering instances pointers"
	}
	for _, specification_rendering := range specification_renderingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFICATION_Rendering", int(stage.SPECIFICATION_RenderingMap_Staged_Order[specification_rendering]), specification_rendering.Name)
		map_SPECIFICATION_Rendering_Identifiers[specification_rendering] = id

		// Initialisation of values
	}

	if len(specification_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPECIFICATION_TYPE instances pointers"
	}
	for _, specification_type := range specification_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFICATION_TYPE", int(stage.SPECIFICATION_TYPEMap_Staged_Order[specification_type]), specification_type.Name)
		map_SPECIFICATION_TYPE_Identifiers[specification_type] = id

		// Initialisation of values
		if specification_type.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[specification_type.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if specification_type.SPEC_ATTRIBUTES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_ATTRIBUTES_Identifiers[specification_type.SPEC_ATTRIBUTES])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spec_hierarchyOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_HIERARCHY instances pointers"
	}
	for _, spec_hierarchy := range spec_hierarchyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_HIERARCHY", int(stage.SPEC_HIERARCHYMap_Staged_Order[spec_hierarchy]), spec_hierarchy.Name)
		map_SPEC_HIERARCHY_Identifiers[spec_hierarchy] = id

		// Initialisation of values
		if spec_hierarchy.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[spec_hierarchy.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if spec_hierarchy.OBJECT != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OBJECT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_OBJECT_Identifiers[spec_hierarchy.OBJECT])
			pointersInitializesStatements += setPointerField
		}

		if spec_hierarchy.CHILDREN != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CHILDREN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_CHILDREN_Identifiers[spec_hierarchy.CHILDREN])
			pointersInitializesStatements += setPointerField
		}

		if spec_hierarchy.EDITABLE_ATTS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EDITABLE_ATTS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_EDITABLE_ATTS_Identifiers[spec_hierarchy.EDITABLE_ATTS])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spec_objectOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_OBJECT instances pointers"
	}
	for _, spec_object := range spec_objectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_OBJECT", int(stage.SPEC_OBJECTMap_Staged_Order[spec_object]), spec_object.Name)
		map_SPEC_OBJECT_Identifiers[spec_object] = id

		// Initialisation of values
		if spec_object.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[spec_object.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if spec_object.VALUES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_VALUE_XHTML_1_Identifiers[spec_object.VALUES])
			pointersInitializesStatements += setPointerField
		}

		if spec_object.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_OBJECT_TYPE_REF_Identifiers[spec_object.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spec_object_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_OBJECT_TYPE instances pointers"
	}
	for _, spec_object_type := range spec_object_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_OBJECT_TYPE", int(stage.SPEC_OBJECT_TYPEMap_Staged_Order[spec_object_type]), spec_object_type.Name)
		map_SPEC_OBJECT_TYPE_Identifiers[spec_object_type] = id

		// Initialisation of values
		if spec_object_type.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[spec_object_type.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if spec_object_type.SPEC_ATTRIBUTES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_ATTRIBUTES_Identifiers[spec_object_type.SPEC_ATTRIBUTES])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spec_object_type_renderingOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_OBJECT_TYPE_Rendering instances pointers"
	}
	for _, spec_object_type_rendering := range spec_object_type_renderingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_OBJECT_TYPE_Rendering", int(stage.SPEC_OBJECT_TYPE_RenderingMap_Staged_Order[spec_object_type_rendering]), spec_object_type_rendering.Name)
		map_SPEC_OBJECT_TYPE_Rendering_Identifiers[spec_object_type_rendering] = id

		// Initialisation of values
	}

	if len(spec_relationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_RELATION instances pointers"
	}
	for _, spec_relation := range spec_relationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_RELATION", int(stage.SPEC_RELATIONMap_Staged_Order[spec_relation]), spec_relation.Name)
		map_SPEC_RELATION_Identifiers[spec_relation] = id

		// Initialisation of values
		if spec_relation.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[spec_relation.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if spec_relation.VALUES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ATTRIBUTE_VALUE_XHTML_1_Identifiers[spec_relation.VALUES])
			pointersInitializesStatements += setPointerField
		}

		if spec_relation.SOURCE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SOURCE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SOURCE_1_Identifiers[spec_relation.SOURCE])
			pointersInitializesStatements += setPointerField
		}

		if spec_relation.TARGET != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TARGET")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SOURCE_1_Identifiers[spec_relation.TARGET])
			pointersInitializesStatements += setPointerField
		}

		if spec_relation.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_RELATION_TYPE_REF_Identifiers[spec_relation.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spec_relation_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_RELATION_TYPE instances pointers"
	}
	for _, spec_relation_type := range spec_relation_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_RELATION_TYPE", int(stage.SPEC_RELATION_TYPEMap_Staged_Order[spec_relation_type]), spec_relation_type.Name)
		map_SPEC_RELATION_TYPE_Identifiers[spec_relation_type] = id

		// Initialisation of values
		if spec_relation_type.ALTERNATIVE_ID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[spec_relation_type.ALTERNATIVE_ID])
			pointersInitializesStatements += setPointerField
		}

		if spec_relation_type.SPEC_ATTRIBUTES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_ATTRIBUTES_Identifiers[spec_relation_type.SPEC_ATTRIBUTES])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(staticwebsiteOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of StaticWebSite instances pointers"
	}
	for _, staticwebsite := range staticwebsiteOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("StaticWebSite", int(stage.StaticWebSiteMap_Staged_Order[staticwebsite]), staticwebsite.Name)
		map_StaticWebSite_Identifiers[staticwebsite] = id

		// Initialisation of values
		for _, _staticwebsitechapter := range staticwebsite.Chapters {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Chapters")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_StaticWebSiteChapter_Identifiers[_staticwebsitechapter])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(staticwebsitechapterOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of StaticWebSiteChapter instances pointers"
	}
	for _, staticwebsitechapter := range staticwebsitechapterOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("StaticWebSiteChapter", int(stage.StaticWebSiteChapterMap_Staged_Order[staticwebsitechapter]), staticwebsitechapter.Name)
		map_StaticWebSiteChapter_Identifiers[staticwebsitechapter] = id

		// Initialisation of values
		for _, _staticwebsiteparagraph := range staticwebsitechapter.Paragraphs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Paragraphs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_StaticWebSiteParagraph_Identifiers[_staticwebsiteparagraph])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(staticwebsitegeneratedimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of StaticWebSiteGeneratedImage instances pointers"
	}
	for _, staticwebsitegeneratedimage := range staticwebsitegeneratedimageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("StaticWebSiteGeneratedImage", int(stage.StaticWebSiteGeneratedImageMap_Staged_Order[staticwebsitegeneratedimage]), staticwebsitegeneratedimage.Name)
		map_StaticWebSiteGeneratedImage_Identifiers[staticwebsitegeneratedimage] = id

		// Initialisation of values
	}

	if len(staticwebsiteimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of StaticWebSiteImage instances pointers"
	}
	for _, staticwebsiteimage := range staticwebsiteimageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("StaticWebSiteImage", int(stage.StaticWebSiteImageMap_Staged_Order[staticwebsiteimage]), staticwebsiteimage.Name)
		map_StaticWebSiteImage_Identifiers[staticwebsiteimage] = id

		// Initialisation of values
	}

	if len(staticwebsiteparagraphOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of StaticWebSiteParagraph instances pointers"
	}
	for _, staticwebsiteparagraph := range staticwebsiteparagraphOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("StaticWebSiteParagraph", int(stage.StaticWebSiteParagraphMap_Staged_Order[staticwebsiteparagraph]), staticwebsiteparagraph.Name)
		map_StaticWebSiteParagraph_Identifiers[staticwebsiteparagraph] = id

		// Initialisation of values
		if staticwebsiteparagraph.Image != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Image")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_StaticWebSiteImage_Identifiers[staticwebsiteparagraph.Image])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(xhtml_contentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of XHTML_CONTENT instances pointers"
	}
	for _, xhtml_content := range xhtml_contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("XHTML_CONTENT", int(stage.XHTML_CONTENTMap_Staged_Order[xhtml_content]), xhtml_content.Name)
		map_XHTML_CONTENT_Identifiers[xhtml_content] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}
	return
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%08d_%s", gongStructName, idx, processedString)

	return
}
