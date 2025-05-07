// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
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
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

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
		log.Fatalln(name + " is not a go filename")
	}

	log.Printf("%s Marshalling %s", time.Now().Format("2006-01-02 15:04:05.000000"), name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
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
	for idx, alternative_id := range alternative_idOrdered {

		id = generatesIdentifier("ALTERNATIVE_ID", idx, alternative_id.Name)
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
	for idx, attribute_definition_boolean := range attribute_definition_booleanOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_BOOLEAN", idx, attribute_definition_boolean.Name)
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_boolean.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", attribute_definition_boolean.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.LONG_NAME))
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
	for idx, attribute_definition_date := range attribute_definition_dateOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_DATE", idx, attribute_definition_date.Name)
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_date.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", attribute_definition_date.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.LONG_NAME))
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
	for idx, attribute_definition_enumeration := range attribute_definition_enumerationOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_ENUMERATION", idx, attribute_definition_enumeration.Name)
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_enumeration.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", attribute_definition_enumeration.LAST_CHANGE.String())
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
	for idx, attribute_definition_integer := range attribute_definition_integerOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_INTEGER", idx, attribute_definition_integer.Name)
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_integer.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", attribute_definition_integer.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.LONG_NAME))
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
	for idx, attribute_definition_real := range attribute_definition_realOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_REAL", idx, attribute_definition_real.Name)
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_real.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", attribute_definition_real.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.LONG_NAME))
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
	for idx, attribute_definition_string := range attribute_definition_stringOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_STRING", idx, attribute_definition_string.Name)
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_string.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", attribute_definition_string.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.LONG_NAME))
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
	for idx, attribute_definition_xhtml := range attribute_definition_xhtmlOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_XHTML", idx, attribute_definition_xhtml.Name)
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_xhtml.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", attribute_definition_xhtml.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.LONG_NAME))
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
	for idx, attribute_value_boolean := range attribute_value_booleanOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_BOOLEAN", idx, attribute_value_boolean.Name)
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
	for idx, attribute_value_date := range attribute_value_dateOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_DATE", idx, attribute_value_date.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", attribute_value_date.THE_VALUE.String())
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
	for idx, attribute_value_enumeration := range attribute_value_enumerationOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_ENUMERATION", idx, attribute_value_enumeration.Name)
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
	for idx, attribute_value_integer := range attribute_value_integerOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_INTEGER", idx, attribute_value_integer.Name)
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
	for idx, attribute_value_real := range attribute_value_realOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_REAL", idx, attribute_value_real.Name)
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
	for idx, attribute_value_string := range attribute_value_stringOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_STRING", idx, attribute_value_string.Name)
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
	for idx, attribute_value_xhtml := range attribute_value_xhtmlOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_XHTML", idx, attribute_value_xhtml.Name)
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

	map_AnyType_Identifiers := make(map[*AnyType]string)
	_ = map_AnyType_Identifiers

	anytypeOrdered := []*AnyType{}
	for anytype := range stage.AnyTypes {
		anytypeOrdered = append(anytypeOrdered, anytype)
	}
	sort.Slice(anytypeOrdered[:], func(i, j int) bool {
		anytypei := anytypeOrdered[i]
		anytypej := anytypeOrdered[j]
		anytypei_order, oki := stage.AnyTypeMap_Staged_Order[anytypei]
		anytypej_order, okj := stage.AnyTypeMap_Staged_Order[anytypej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return anytypei_order < anytypej_order
	})
	if len(anytypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, anytype := range anytypeOrdered {

		id = generatesIdentifier("AnyType", idx, anytype.Name)
		map_AnyType_Identifiers[anytype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AnyType")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", anytype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(anytype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InnerXML")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(anytype.InnerXML))
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
	for idx, datatype_definition_boolean := range datatype_definition_booleanOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_BOOLEAN", idx, datatype_definition_boolean.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", datatype_definition_boolean.LAST_CHANGE.String())
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
	for idx, datatype_definition_date := range datatype_definition_dateOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_DATE", idx, datatype_definition_date.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", datatype_definition_date.LAST_CHANGE.String())
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
	for idx, datatype_definition_enumeration := range datatype_definition_enumerationOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_ENUMERATION", idx, datatype_definition_enumeration.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", datatype_definition_enumeration.LAST_CHANGE.String())
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
	for idx, datatype_definition_integer := range datatype_definition_integerOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_INTEGER", idx, datatype_definition_integer.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", datatype_definition_integer.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.LONG_NAME))
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
	for idx, datatype_definition_real := range datatype_definition_realOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_REAL", idx, datatype_definition_real.Name)
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

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.DESC))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", datatype_definition_real.LAST_CHANGE.String())
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
	for idx, datatype_definition_string := range datatype_definition_stringOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_STRING", idx, datatype_definition_string.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", datatype_definition_string.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.LONG_NAME))
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
	for idx, datatype_definition_xhtml := range datatype_definition_xhtmlOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_XHTML", idx, datatype_definition_xhtml.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", datatype_definition_xhtml.LAST_CHANGE.String())
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
	for idx, embedded_value := range embedded_valueOrdered {

		id = generatesIdentifier("EMBEDDED_VALUE", idx, embedded_value.Name)
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
	for idx, enum_value := range enum_valueOrdered {

		id = generatesIdentifier("ENUM_VALUE", idx, enum_value.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", enum_value.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.LONG_NAME))
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
	for idx, relation_group := range relation_groupOrdered {

		id = generatesIdentifier("RELATION_GROUP", idx, relation_group.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", relation_group.LAST_CHANGE.String())
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
	for idx, relation_group_type := range relation_group_typeOrdered {

		id = generatesIdentifier("RELATION_GROUP_TYPE", idx, relation_group_type.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", relation_group_type.LAST_CHANGE.String())
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
	for idx, req_if := range req_ifOrdered {

		id = generatesIdentifier("REQ_IF", idx, req_if.Name)
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
	for idx, req_if_content := range req_if_contentOrdered {

		id = generatesIdentifier("REQ_IF_CONTENT", idx, req_if_content.Name)
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
	for idx, req_if_header := range req_if_headerOrdered {

		id = generatesIdentifier("REQ_IF_HEADER", idx, req_if_header.Name)
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "COMMENT")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.COMMENT))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CREATION_TIME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", req_if_header.CREATION_TIME.String())
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
	for idx, req_if_tool_extension := range req_if_tool_extensionOrdered {

		id = generatesIdentifier("REQ_IF_TOOL_EXTENSION", idx, req_if_tool_extension.Name)
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
	for idx, specification := range specificationOrdered {

		id = generatesIdentifier("SPECIFICATION", idx, specification.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", specification.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.LONG_NAME))
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
	for idx, specification_type := range specification_typeOrdered {

		id = generatesIdentifier("SPECIFICATION_TYPE", idx, specification_type.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", specification_type.LAST_CHANGE.String())
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
	for idx, spec_hierarchy := range spec_hierarchyOrdered {

		id = generatesIdentifier("SPEC_HIERARCHY", idx, spec_hierarchy.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", spec_hierarchy.LAST_CHANGE.String())
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
	for idx, spec_object := range spec_objectOrdered {

		id = generatesIdentifier("SPEC_OBJECT", idx, spec_object.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", spec_object.LAST_CHANGE.String())
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
	for idx, spec_object_type := range spec_object_typeOrdered {

		id = generatesIdentifier("SPEC_OBJECT_TYPE", idx, spec_object_type.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", spec_object_type.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.LONG_NAME))
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
	for idx, spec_relation := range spec_relationOrdered {

		id = generatesIdentifier("SPEC_RELATION", idx, spec_relation.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", spec_relation.LAST_CHANGE.String())
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
	for idx, spec_relation_type := range spec_relation_typeOrdered {

		id = generatesIdentifier("SPEC_RELATION_TYPE", idx, spec_relation_type.Name)
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

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", spec_relation_type.LAST_CHANGE.String())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.LONG_NAME))
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
	for idx, xhtml_content := range xhtml_contentOrdered {

		id = generatesIdentifier("XHTML_CONTENT", idx, xhtml_content.Name)
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

	}

	map_Xhtml_InlPres_type_Identifiers := make(map[*Xhtml_InlPres_type]string)
	_ = map_Xhtml_InlPres_type_Identifiers

	xhtml_inlpres_typeOrdered := []*Xhtml_InlPres_type{}
	for xhtml_inlpres_type := range stage.Xhtml_InlPres_types {
		xhtml_inlpres_typeOrdered = append(xhtml_inlpres_typeOrdered, xhtml_inlpres_type)
	}
	sort.Slice(xhtml_inlpres_typeOrdered[:], func(i, j int) bool {
		xhtml_inlpres_typei := xhtml_inlpres_typeOrdered[i]
		xhtml_inlpres_typej := xhtml_inlpres_typeOrdered[j]
		xhtml_inlpres_typei_order, oki := stage.Xhtml_InlPres_typeMap_Staged_Order[xhtml_inlpres_typei]
		xhtml_inlpres_typej_order, okj := stage.Xhtml_InlPres_typeMap_Staged_Order[xhtml_inlpres_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_inlpres_typei_order < xhtml_inlpres_typej_order
	})
	if len(xhtml_inlpres_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_inlpres_type := range xhtml_inlpres_typeOrdered {

		id = generatesIdentifier("Xhtml_InlPres_type", idx, xhtml_inlpres_type.Name)
		map_Xhtml_InlPres_type_Identifiers[xhtml_inlpres_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_InlPres_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_inlpres_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_inlpres_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_a_type_Identifiers := make(map[*Xhtml_a_type]string)
	_ = map_Xhtml_a_type_Identifiers

	xhtml_a_typeOrdered := []*Xhtml_a_type{}
	for xhtml_a_type := range stage.Xhtml_a_types {
		xhtml_a_typeOrdered = append(xhtml_a_typeOrdered, xhtml_a_type)
	}
	sort.Slice(xhtml_a_typeOrdered[:], func(i, j int) bool {
		xhtml_a_typei := xhtml_a_typeOrdered[i]
		xhtml_a_typej := xhtml_a_typeOrdered[j]
		xhtml_a_typei_order, oki := stage.Xhtml_a_typeMap_Staged_Order[xhtml_a_typei]
		xhtml_a_typej_order, okj := stage.Xhtml_a_typeMap_Staged_Order[xhtml_a_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_a_typei_order < xhtml_a_typej_order
	})
	if len(xhtml_a_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_a_type := range xhtml_a_typeOrdered {

		id = generatesIdentifier("Xhtml_a_type", idx, xhtml_a_type.Name)
		map_Xhtml_a_type_Identifiers[xhtml_a_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_a_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_a_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_a_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_abbr_type_Identifiers := make(map[*Xhtml_abbr_type]string)
	_ = map_Xhtml_abbr_type_Identifiers

	xhtml_abbr_typeOrdered := []*Xhtml_abbr_type{}
	for xhtml_abbr_type := range stage.Xhtml_abbr_types {
		xhtml_abbr_typeOrdered = append(xhtml_abbr_typeOrdered, xhtml_abbr_type)
	}
	sort.Slice(xhtml_abbr_typeOrdered[:], func(i, j int) bool {
		xhtml_abbr_typei := xhtml_abbr_typeOrdered[i]
		xhtml_abbr_typej := xhtml_abbr_typeOrdered[j]
		xhtml_abbr_typei_order, oki := stage.Xhtml_abbr_typeMap_Staged_Order[xhtml_abbr_typei]
		xhtml_abbr_typej_order, okj := stage.Xhtml_abbr_typeMap_Staged_Order[xhtml_abbr_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_abbr_typei_order < xhtml_abbr_typej_order
	})
	if len(xhtml_abbr_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_abbr_type := range xhtml_abbr_typeOrdered {

		id = generatesIdentifier("Xhtml_abbr_type", idx, xhtml_abbr_type.Name)
		map_Xhtml_abbr_type_Identifiers[xhtml_abbr_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_abbr_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_abbr_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_abbr_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_acronym_type_Identifiers := make(map[*Xhtml_acronym_type]string)
	_ = map_Xhtml_acronym_type_Identifiers

	xhtml_acronym_typeOrdered := []*Xhtml_acronym_type{}
	for xhtml_acronym_type := range stage.Xhtml_acronym_types {
		xhtml_acronym_typeOrdered = append(xhtml_acronym_typeOrdered, xhtml_acronym_type)
	}
	sort.Slice(xhtml_acronym_typeOrdered[:], func(i, j int) bool {
		xhtml_acronym_typei := xhtml_acronym_typeOrdered[i]
		xhtml_acronym_typej := xhtml_acronym_typeOrdered[j]
		xhtml_acronym_typei_order, oki := stage.Xhtml_acronym_typeMap_Staged_Order[xhtml_acronym_typei]
		xhtml_acronym_typej_order, okj := stage.Xhtml_acronym_typeMap_Staged_Order[xhtml_acronym_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_acronym_typei_order < xhtml_acronym_typej_order
	})
	if len(xhtml_acronym_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_acronym_type := range xhtml_acronym_typeOrdered {

		id = generatesIdentifier("Xhtml_acronym_type", idx, xhtml_acronym_type.Name)
		map_Xhtml_acronym_type_Identifiers[xhtml_acronym_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_acronym_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_acronym_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_acronym_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_address_type_Identifiers := make(map[*Xhtml_address_type]string)
	_ = map_Xhtml_address_type_Identifiers

	xhtml_address_typeOrdered := []*Xhtml_address_type{}
	for xhtml_address_type := range stage.Xhtml_address_types {
		xhtml_address_typeOrdered = append(xhtml_address_typeOrdered, xhtml_address_type)
	}
	sort.Slice(xhtml_address_typeOrdered[:], func(i, j int) bool {
		xhtml_address_typei := xhtml_address_typeOrdered[i]
		xhtml_address_typej := xhtml_address_typeOrdered[j]
		xhtml_address_typei_order, oki := stage.Xhtml_address_typeMap_Staged_Order[xhtml_address_typei]
		xhtml_address_typej_order, okj := stage.Xhtml_address_typeMap_Staged_Order[xhtml_address_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_address_typei_order < xhtml_address_typej_order
	})
	if len(xhtml_address_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_address_type := range xhtml_address_typeOrdered {

		id = generatesIdentifier("Xhtml_address_type", idx, xhtml_address_type.Name)
		map_Xhtml_address_type_Identifiers[xhtml_address_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_address_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_address_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_address_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_blockquote_type_Identifiers := make(map[*Xhtml_blockquote_type]string)
	_ = map_Xhtml_blockquote_type_Identifiers

	xhtml_blockquote_typeOrdered := []*Xhtml_blockquote_type{}
	for xhtml_blockquote_type := range stage.Xhtml_blockquote_types {
		xhtml_blockquote_typeOrdered = append(xhtml_blockquote_typeOrdered, xhtml_blockquote_type)
	}
	sort.Slice(xhtml_blockquote_typeOrdered[:], func(i, j int) bool {
		xhtml_blockquote_typei := xhtml_blockquote_typeOrdered[i]
		xhtml_blockquote_typej := xhtml_blockquote_typeOrdered[j]
		xhtml_blockquote_typei_order, oki := stage.Xhtml_blockquote_typeMap_Staged_Order[xhtml_blockquote_typei]
		xhtml_blockquote_typej_order, okj := stage.Xhtml_blockquote_typeMap_Staged_Order[xhtml_blockquote_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_blockquote_typei_order < xhtml_blockquote_typej_order
	})
	if len(xhtml_blockquote_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_blockquote_type := range xhtml_blockquote_typeOrdered {

		id = generatesIdentifier("Xhtml_blockquote_type", idx, xhtml_blockquote_type.Name)
		map_Xhtml_blockquote_type_Identifiers[xhtml_blockquote_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_blockquote_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_blockquote_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_blockquote_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_br_type_Identifiers := make(map[*Xhtml_br_type]string)
	_ = map_Xhtml_br_type_Identifiers

	xhtml_br_typeOrdered := []*Xhtml_br_type{}
	for xhtml_br_type := range stage.Xhtml_br_types {
		xhtml_br_typeOrdered = append(xhtml_br_typeOrdered, xhtml_br_type)
	}
	sort.Slice(xhtml_br_typeOrdered[:], func(i, j int) bool {
		xhtml_br_typei := xhtml_br_typeOrdered[i]
		xhtml_br_typej := xhtml_br_typeOrdered[j]
		xhtml_br_typei_order, oki := stage.Xhtml_br_typeMap_Staged_Order[xhtml_br_typei]
		xhtml_br_typej_order, okj := stage.Xhtml_br_typeMap_Staged_Order[xhtml_br_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_br_typei_order < xhtml_br_typej_order
	})
	if len(xhtml_br_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_br_type := range xhtml_br_typeOrdered {

		id = generatesIdentifier("Xhtml_br_type", idx, xhtml_br_type.Name)
		map_Xhtml_br_type_Identifiers[xhtml_br_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_br_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_br_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_br_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_caption_type_Identifiers := make(map[*Xhtml_caption_type]string)
	_ = map_Xhtml_caption_type_Identifiers

	xhtml_caption_typeOrdered := []*Xhtml_caption_type{}
	for xhtml_caption_type := range stage.Xhtml_caption_types {
		xhtml_caption_typeOrdered = append(xhtml_caption_typeOrdered, xhtml_caption_type)
	}
	sort.Slice(xhtml_caption_typeOrdered[:], func(i, j int) bool {
		xhtml_caption_typei := xhtml_caption_typeOrdered[i]
		xhtml_caption_typej := xhtml_caption_typeOrdered[j]
		xhtml_caption_typei_order, oki := stage.Xhtml_caption_typeMap_Staged_Order[xhtml_caption_typei]
		xhtml_caption_typej_order, okj := stage.Xhtml_caption_typeMap_Staged_Order[xhtml_caption_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_caption_typei_order < xhtml_caption_typej_order
	})
	if len(xhtml_caption_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_caption_type := range xhtml_caption_typeOrdered {

		id = generatesIdentifier("Xhtml_caption_type", idx, xhtml_caption_type.Name)
		map_Xhtml_caption_type_Identifiers[xhtml_caption_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_caption_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_caption_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_caption_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_cite_type_Identifiers := make(map[*Xhtml_cite_type]string)
	_ = map_Xhtml_cite_type_Identifiers

	xhtml_cite_typeOrdered := []*Xhtml_cite_type{}
	for xhtml_cite_type := range stage.Xhtml_cite_types {
		xhtml_cite_typeOrdered = append(xhtml_cite_typeOrdered, xhtml_cite_type)
	}
	sort.Slice(xhtml_cite_typeOrdered[:], func(i, j int) bool {
		xhtml_cite_typei := xhtml_cite_typeOrdered[i]
		xhtml_cite_typej := xhtml_cite_typeOrdered[j]
		xhtml_cite_typei_order, oki := stage.Xhtml_cite_typeMap_Staged_Order[xhtml_cite_typei]
		xhtml_cite_typej_order, okj := stage.Xhtml_cite_typeMap_Staged_Order[xhtml_cite_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_cite_typei_order < xhtml_cite_typej_order
	})
	if len(xhtml_cite_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_cite_type := range xhtml_cite_typeOrdered {

		id = generatesIdentifier("Xhtml_cite_type", idx, xhtml_cite_type.Name)
		map_Xhtml_cite_type_Identifiers[xhtml_cite_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_cite_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_cite_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_cite_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_code_type_Identifiers := make(map[*Xhtml_code_type]string)
	_ = map_Xhtml_code_type_Identifiers

	xhtml_code_typeOrdered := []*Xhtml_code_type{}
	for xhtml_code_type := range stage.Xhtml_code_types {
		xhtml_code_typeOrdered = append(xhtml_code_typeOrdered, xhtml_code_type)
	}
	sort.Slice(xhtml_code_typeOrdered[:], func(i, j int) bool {
		xhtml_code_typei := xhtml_code_typeOrdered[i]
		xhtml_code_typej := xhtml_code_typeOrdered[j]
		xhtml_code_typei_order, oki := stage.Xhtml_code_typeMap_Staged_Order[xhtml_code_typei]
		xhtml_code_typej_order, okj := stage.Xhtml_code_typeMap_Staged_Order[xhtml_code_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_code_typei_order < xhtml_code_typej_order
	})
	if len(xhtml_code_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_code_type := range xhtml_code_typeOrdered {

		id = generatesIdentifier("Xhtml_code_type", idx, xhtml_code_type.Name)
		map_Xhtml_code_type_Identifiers[xhtml_code_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_code_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_code_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_code_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_col_type_Identifiers := make(map[*Xhtml_col_type]string)
	_ = map_Xhtml_col_type_Identifiers

	xhtml_col_typeOrdered := []*Xhtml_col_type{}
	for xhtml_col_type := range stage.Xhtml_col_types {
		xhtml_col_typeOrdered = append(xhtml_col_typeOrdered, xhtml_col_type)
	}
	sort.Slice(xhtml_col_typeOrdered[:], func(i, j int) bool {
		xhtml_col_typei := xhtml_col_typeOrdered[i]
		xhtml_col_typej := xhtml_col_typeOrdered[j]
		xhtml_col_typei_order, oki := stage.Xhtml_col_typeMap_Staged_Order[xhtml_col_typei]
		xhtml_col_typej_order, okj := stage.Xhtml_col_typeMap_Staged_Order[xhtml_col_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_col_typei_order < xhtml_col_typej_order
	})
	if len(xhtml_col_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_col_type := range xhtml_col_typeOrdered {

		id = generatesIdentifier("Xhtml_col_type", idx, xhtml_col_type.Name)
		map_Xhtml_col_type_Identifiers[xhtml_col_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_col_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_col_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_col_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_colgroup_type_Identifiers := make(map[*Xhtml_colgroup_type]string)
	_ = map_Xhtml_colgroup_type_Identifiers

	xhtml_colgroup_typeOrdered := []*Xhtml_colgroup_type{}
	for xhtml_colgroup_type := range stage.Xhtml_colgroup_types {
		xhtml_colgroup_typeOrdered = append(xhtml_colgroup_typeOrdered, xhtml_colgroup_type)
	}
	sort.Slice(xhtml_colgroup_typeOrdered[:], func(i, j int) bool {
		xhtml_colgroup_typei := xhtml_colgroup_typeOrdered[i]
		xhtml_colgroup_typej := xhtml_colgroup_typeOrdered[j]
		xhtml_colgroup_typei_order, oki := stage.Xhtml_colgroup_typeMap_Staged_Order[xhtml_colgroup_typei]
		xhtml_colgroup_typej_order, okj := stage.Xhtml_colgroup_typeMap_Staged_Order[xhtml_colgroup_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_colgroup_typei_order < xhtml_colgroup_typej_order
	})
	if len(xhtml_colgroup_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_colgroup_type := range xhtml_colgroup_typeOrdered {

		id = generatesIdentifier("Xhtml_colgroup_type", idx, xhtml_colgroup_type.Name)
		map_Xhtml_colgroup_type_Identifiers[xhtml_colgroup_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_colgroup_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_colgroup_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_colgroup_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_dd_type_Identifiers := make(map[*Xhtml_dd_type]string)
	_ = map_Xhtml_dd_type_Identifiers

	xhtml_dd_typeOrdered := []*Xhtml_dd_type{}
	for xhtml_dd_type := range stage.Xhtml_dd_types {
		xhtml_dd_typeOrdered = append(xhtml_dd_typeOrdered, xhtml_dd_type)
	}
	sort.Slice(xhtml_dd_typeOrdered[:], func(i, j int) bool {
		xhtml_dd_typei := xhtml_dd_typeOrdered[i]
		xhtml_dd_typej := xhtml_dd_typeOrdered[j]
		xhtml_dd_typei_order, oki := stage.Xhtml_dd_typeMap_Staged_Order[xhtml_dd_typei]
		xhtml_dd_typej_order, okj := stage.Xhtml_dd_typeMap_Staged_Order[xhtml_dd_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_dd_typei_order < xhtml_dd_typej_order
	})
	if len(xhtml_dd_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_dd_type := range xhtml_dd_typeOrdered {

		id = generatesIdentifier("Xhtml_dd_type", idx, xhtml_dd_type.Name)
		map_Xhtml_dd_type_Identifiers[xhtml_dd_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_dd_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_dd_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_dd_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_dfn_type_Identifiers := make(map[*Xhtml_dfn_type]string)
	_ = map_Xhtml_dfn_type_Identifiers

	xhtml_dfn_typeOrdered := []*Xhtml_dfn_type{}
	for xhtml_dfn_type := range stage.Xhtml_dfn_types {
		xhtml_dfn_typeOrdered = append(xhtml_dfn_typeOrdered, xhtml_dfn_type)
	}
	sort.Slice(xhtml_dfn_typeOrdered[:], func(i, j int) bool {
		xhtml_dfn_typei := xhtml_dfn_typeOrdered[i]
		xhtml_dfn_typej := xhtml_dfn_typeOrdered[j]
		xhtml_dfn_typei_order, oki := stage.Xhtml_dfn_typeMap_Staged_Order[xhtml_dfn_typei]
		xhtml_dfn_typej_order, okj := stage.Xhtml_dfn_typeMap_Staged_Order[xhtml_dfn_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_dfn_typei_order < xhtml_dfn_typej_order
	})
	if len(xhtml_dfn_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_dfn_type := range xhtml_dfn_typeOrdered {

		id = generatesIdentifier("Xhtml_dfn_type", idx, xhtml_dfn_type.Name)
		map_Xhtml_dfn_type_Identifiers[xhtml_dfn_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_dfn_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_dfn_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_dfn_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_div_type_Identifiers := make(map[*Xhtml_div_type]string)
	_ = map_Xhtml_div_type_Identifiers

	xhtml_div_typeOrdered := []*Xhtml_div_type{}
	for xhtml_div_type := range stage.Xhtml_div_types {
		xhtml_div_typeOrdered = append(xhtml_div_typeOrdered, xhtml_div_type)
	}
	sort.Slice(xhtml_div_typeOrdered[:], func(i, j int) bool {
		xhtml_div_typei := xhtml_div_typeOrdered[i]
		xhtml_div_typej := xhtml_div_typeOrdered[j]
		xhtml_div_typei_order, oki := stage.Xhtml_div_typeMap_Staged_Order[xhtml_div_typei]
		xhtml_div_typej_order, okj := stage.Xhtml_div_typeMap_Staged_Order[xhtml_div_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_div_typei_order < xhtml_div_typej_order
	})
	if len(xhtml_div_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_div_type := range xhtml_div_typeOrdered {

		id = generatesIdentifier("Xhtml_div_type", idx, xhtml_div_type.Name)
		map_Xhtml_div_type_Identifiers[xhtml_div_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_div_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_div_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_div_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_dl_type_Identifiers := make(map[*Xhtml_dl_type]string)
	_ = map_Xhtml_dl_type_Identifiers

	xhtml_dl_typeOrdered := []*Xhtml_dl_type{}
	for xhtml_dl_type := range stage.Xhtml_dl_types {
		xhtml_dl_typeOrdered = append(xhtml_dl_typeOrdered, xhtml_dl_type)
	}
	sort.Slice(xhtml_dl_typeOrdered[:], func(i, j int) bool {
		xhtml_dl_typei := xhtml_dl_typeOrdered[i]
		xhtml_dl_typej := xhtml_dl_typeOrdered[j]
		xhtml_dl_typei_order, oki := stage.Xhtml_dl_typeMap_Staged_Order[xhtml_dl_typei]
		xhtml_dl_typej_order, okj := stage.Xhtml_dl_typeMap_Staged_Order[xhtml_dl_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_dl_typei_order < xhtml_dl_typej_order
	})
	if len(xhtml_dl_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_dl_type := range xhtml_dl_typeOrdered {

		id = generatesIdentifier("Xhtml_dl_type", idx, xhtml_dl_type.Name)
		map_Xhtml_dl_type_Identifiers[xhtml_dl_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_dl_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_dl_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_dl_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_dt_type_Identifiers := make(map[*Xhtml_dt_type]string)
	_ = map_Xhtml_dt_type_Identifiers

	xhtml_dt_typeOrdered := []*Xhtml_dt_type{}
	for xhtml_dt_type := range stage.Xhtml_dt_types {
		xhtml_dt_typeOrdered = append(xhtml_dt_typeOrdered, xhtml_dt_type)
	}
	sort.Slice(xhtml_dt_typeOrdered[:], func(i, j int) bool {
		xhtml_dt_typei := xhtml_dt_typeOrdered[i]
		xhtml_dt_typej := xhtml_dt_typeOrdered[j]
		xhtml_dt_typei_order, oki := stage.Xhtml_dt_typeMap_Staged_Order[xhtml_dt_typei]
		xhtml_dt_typej_order, okj := stage.Xhtml_dt_typeMap_Staged_Order[xhtml_dt_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_dt_typei_order < xhtml_dt_typej_order
	})
	if len(xhtml_dt_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_dt_type := range xhtml_dt_typeOrdered {

		id = generatesIdentifier("Xhtml_dt_type", idx, xhtml_dt_type.Name)
		map_Xhtml_dt_type_Identifiers[xhtml_dt_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_dt_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_dt_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_dt_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_edit_type_Identifiers := make(map[*Xhtml_edit_type]string)
	_ = map_Xhtml_edit_type_Identifiers

	xhtml_edit_typeOrdered := []*Xhtml_edit_type{}
	for xhtml_edit_type := range stage.Xhtml_edit_types {
		xhtml_edit_typeOrdered = append(xhtml_edit_typeOrdered, xhtml_edit_type)
	}
	sort.Slice(xhtml_edit_typeOrdered[:], func(i, j int) bool {
		xhtml_edit_typei := xhtml_edit_typeOrdered[i]
		xhtml_edit_typej := xhtml_edit_typeOrdered[j]
		xhtml_edit_typei_order, oki := stage.Xhtml_edit_typeMap_Staged_Order[xhtml_edit_typei]
		xhtml_edit_typej_order, okj := stage.Xhtml_edit_typeMap_Staged_Order[xhtml_edit_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_edit_typei_order < xhtml_edit_typej_order
	})
	if len(xhtml_edit_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_edit_type := range xhtml_edit_typeOrdered {

		id = generatesIdentifier("Xhtml_edit_type", idx, xhtml_edit_type.Name)
		map_Xhtml_edit_type_Identifiers[xhtml_edit_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_edit_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_edit_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_edit_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_em_type_Identifiers := make(map[*Xhtml_em_type]string)
	_ = map_Xhtml_em_type_Identifiers

	xhtml_em_typeOrdered := []*Xhtml_em_type{}
	for xhtml_em_type := range stage.Xhtml_em_types {
		xhtml_em_typeOrdered = append(xhtml_em_typeOrdered, xhtml_em_type)
	}
	sort.Slice(xhtml_em_typeOrdered[:], func(i, j int) bool {
		xhtml_em_typei := xhtml_em_typeOrdered[i]
		xhtml_em_typej := xhtml_em_typeOrdered[j]
		xhtml_em_typei_order, oki := stage.Xhtml_em_typeMap_Staged_Order[xhtml_em_typei]
		xhtml_em_typej_order, okj := stage.Xhtml_em_typeMap_Staged_Order[xhtml_em_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_em_typei_order < xhtml_em_typej_order
	})
	if len(xhtml_em_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_em_type := range xhtml_em_typeOrdered {

		id = generatesIdentifier("Xhtml_em_type", idx, xhtml_em_type.Name)
		map_Xhtml_em_type_Identifiers[xhtml_em_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_em_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_em_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_em_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_h1_type_Identifiers := make(map[*Xhtml_h1_type]string)
	_ = map_Xhtml_h1_type_Identifiers

	xhtml_h1_typeOrdered := []*Xhtml_h1_type{}
	for xhtml_h1_type := range stage.Xhtml_h1_types {
		xhtml_h1_typeOrdered = append(xhtml_h1_typeOrdered, xhtml_h1_type)
	}
	sort.Slice(xhtml_h1_typeOrdered[:], func(i, j int) bool {
		xhtml_h1_typei := xhtml_h1_typeOrdered[i]
		xhtml_h1_typej := xhtml_h1_typeOrdered[j]
		xhtml_h1_typei_order, oki := stage.Xhtml_h1_typeMap_Staged_Order[xhtml_h1_typei]
		xhtml_h1_typej_order, okj := stage.Xhtml_h1_typeMap_Staged_Order[xhtml_h1_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_h1_typei_order < xhtml_h1_typej_order
	})
	if len(xhtml_h1_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_h1_type := range xhtml_h1_typeOrdered {

		id = generatesIdentifier("Xhtml_h1_type", idx, xhtml_h1_type.Name)
		map_Xhtml_h1_type_Identifiers[xhtml_h1_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_h1_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_h1_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_h1_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_h2_type_Identifiers := make(map[*Xhtml_h2_type]string)
	_ = map_Xhtml_h2_type_Identifiers

	xhtml_h2_typeOrdered := []*Xhtml_h2_type{}
	for xhtml_h2_type := range stage.Xhtml_h2_types {
		xhtml_h2_typeOrdered = append(xhtml_h2_typeOrdered, xhtml_h2_type)
	}
	sort.Slice(xhtml_h2_typeOrdered[:], func(i, j int) bool {
		xhtml_h2_typei := xhtml_h2_typeOrdered[i]
		xhtml_h2_typej := xhtml_h2_typeOrdered[j]
		xhtml_h2_typei_order, oki := stage.Xhtml_h2_typeMap_Staged_Order[xhtml_h2_typei]
		xhtml_h2_typej_order, okj := stage.Xhtml_h2_typeMap_Staged_Order[xhtml_h2_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_h2_typei_order < xhtml_h2_typej_order
	})
	if len(xhtml_h2_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_h2_type := range xhtml_h2_typeOrdered {

		id = generatesIdentifier("Xhtml_h2_type", idx, xhtml_h2_type.Name)
		map_Xhtml_h2_type_Identifiers[xhtml_h2_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_h2_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_h2_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_h2_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_h3_type_Identifiers := make(map[*Xhtml_h3_type]string)
	_ = map_Xhtml_h3_type_Identifiers

	xhtml_h3_typeOrdered := []*Xhtml_h3_type{}
	for xhtml_h3_type := range stage.Xhtml_h3_types {
		xhtml_h3_typeOrdered = append(xhtml_h3_typeOrdered, xhtml_h3_type)
	}
	sort.Slice(xhtml_h3_typeOrdered[:], func(i, j int) bool {
		xhtml_h3_typei := xhtml_h3_typeOrdered[i]
		xhtml_h3_typej := xhtml_h3_typeOrdered[j]
		xhtml_h3_typei_order, oki := stage.Xhtml_h3_typeMap_Staged_Order[xhtml_h3_typei]
		xhtml_h3_typej_order, okj := stage.Xhtml_h3_typeMap_Staged_Order[xhtml_h3_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_h3_typei_order < xhtml_h3_typej_order
	})
	if len(xhtml_h3_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_h3_type := range xhtml_h3_typeOrdered {

		id = generatesIdentifier("Xhtml_h3_type", idx, xhtml_h3_type.Name)
		map_Xhtml_h3_type_Identifiers[xhtml_h3_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_h3_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_h3_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_h3_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_h4_type_Identifiers := make(map[*Xhtml_h4_type]string)
	_ = map_Xhtml_h4_type_Identifiers

	xhtml_h4_typeOrdered := []*Xhtml_h4_type{}
	for xhtml_h4_type := range stage.Xhtml_h4_types {
		xhtml_h4_typeOrdered = append(xhtml_h4_typeOrdered, xhtml_h4_type)
	}
	sort.Slice(xhtml_h4_typeOrdered[:], func(i, j int) bool {
		xhtml_h4_typei := xhtml_h4_typeOrdered[i]
		xhtml_h4_typej := xhtml_h4_typeOrdered[j]
		xhtml_h4_typei_order, oki := stage.Xhtml_h4_typeMap_Staged_Order[xhtml_h4_typei]
		xhtml_h4_typej_order, okj := stage.Xhtml_h4_typeMap_Staged_Order[xhtml_h4_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_h4_typei_order < xhtml_h4_typej_order
	})
	if len(xhtml_h4_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_h4_type := range xhtml_h4_typeOrdered {

		id = generatesIdentifier("Xhtml_h4_type", idx, xhtml_h4_type.Name)
		map_Xhtml_h4_type_Identifiers[xhtml_h4_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_h4_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_h4_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_h4_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_h5_type_Identifiers := make(map[*Xhtml_h5_type]string)
	_ = map_Xhtml_h5_type_Identifiers

	xhtml_h5_typeOrdered := []*Xhtml_h5_type{}
	for xhtml_h5_type := range stage.Xhtml_h5_types {
		xhtml_h5_typeOrdered = append(xhtml_h5_typeOrdered, xhtml_h5_type)
	}
	sort.Slice(xhtml_h5_typeOrdered[:], func(i, j int) bool {
		xhtml_h5_typei := xhtml_h5_typeOrdered[i]
		xhtml_h5_typej := xhtml_h5_typeOrdered[j]
		xhtml_h5_typei_order, oki := stage.Xhtml_h5_typeMap_Staged_Order[xhtml_h5_typei]
		xhtml_h5_typej_order, okj := stage.Xhtml_h5_typeMap_Staged_Order[xhtml_h5_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_h5_typei_order < xhtml_h5_typej_order
	})
	if len(xhtml_h5_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_h5_type := range xhtml_h5_typeOrdered {

		id = generatesIdentifier("Xhtml_h5_type", idx, xhtml_h5_type.Name)
		map_Xhtml_h5_type_Identifiers[xhtml_h5_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_h5_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_h5_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_h5_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_h6_type_Identifiers := make(map[*Xhtml_h6_type]string)
	_ = map_Xhtml_h6_type_Identifiers

	xhtml_h6_typeOrdered := []*Xhtml_h6_type{}
	for xhtml_h6_type := range stage.Xhtml_h6_types {
		xhtml_h6_typeOrdered = append(xhtml_h6_typeOrdered, xhtml_h6_type)
	}
	sort.Slice(xhtml_h6_typeOrdered[:], func(i, j int) bool {
		xhtml_h6_typei := xhtml_h6_typeOrdered[i]
		xhtml_h6_typej := xhtml_h6_typeOrdered[j]
		xhtml_h6_typei_order, oki := stage.Xhtml_h6_typeMap_Staged_Order[xhtml_h6_typei]
		xhtml_h6_typej_order, okj := stage.Xhtml_h6_typeMap_Staged_Order[xhtml_h6_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_h6_typei_order < xhtml_h6_typej_order
	})
	if len(xhtml_h6_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_h6_type := range xhtml_h6_typeOrdered {

		id = generatesIdentifier("Xhtml_h6_type", idx, xhtml_h6_type.Name)
		map_Xhtml_h6_type_Identifiers[xhtml_h6_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_h6_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_h6_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_h6_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_heading_type_Identifiers := make(map[*Xhtml_heading_type]string)
	_ = map_Xhtml_heading_type_Identifiers

	xhtml_heading_typeOrdered := []*Xhtml_heading_type{}
	for xhtml_heading_type := range stage.Xhtml_heading_types {
		xhtml_heading_typeOrdered = append(xhtml_heading_typeOrdered, xhtml_heading_type)
	}
	sort.Slice(xhtml_heading_typeOrdered[:], func(i, j int) bool {
		xhtml_heading_typei := xhtml_heading_typeOrdered[i]
		xhtml_heading_typej := xhtml_heading_typeOrdered[j]
		xhtml_heading_typei_order, oki := stage.Xhtml_heading_typeMap_Staged_Order[xhtml_heading_typei]
		xhtml_heading_typej_order, okj := stage.Xhtml_heading_typeMap_Staged_Order[xhtml_heading_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_heading_typei_order < xhtml_heading_typej_order
	})
	if len(xhtml_heading_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_heading_type := range xhtml_heading_typeOrdered {

		id = generatesIdentifier("Xhtml_heading_type", idx, xhtml_heading_type.Name)
		map_Xhtml_heading_type_Identifiers[xhtml_heading_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_heading_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_heading_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_heading_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_hr_type_Identifiers := make(map[*Xhtml_hr_type]string)
	_ = map_Xhtml_hr_type_Identifiers

	xhtml_hr_typeOrdered := []*Xhtml_hr_type{}
	for xhtml_hr_type := range stage.Xhtml_hr_types {
		xhtml_hr_typeOrdered = append(xhtml_hr_typeOrdered, xhtml_hr_type)
	}
	sort.Slice(xhtml_hr_typeOrdered[:], func(i, j int) bool {
		xhtml_hr_typei := xhtml_hr_typeOrdered[i]
		xhtml_hr_typej := xhtml_hr_typeOrdered[j]
		xhtml_hr_typei_order, oki := stage.Xhtml_hr_typeMap_Staged_Order[xhtml_hr_typei]
		xhtml_hr_typej_order, okj := stage.Xhtml_hr_typeMap_Staged_Order[xhtml_hr_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_hr_typei_order < xhtml_hr_typej_order
	})
	if len(xhtml_hr_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_hr_type := range xhtml_hr_typeOrdered {

		id = generatesIdentifier("Xhtml_hr_type", idx, xhtml_hr_type.Name)
		map_Xhtml_hr_type_Identifiers[xhtml_hr_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_hr_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_hr_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_hr_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_kbd_type_Identifiers := make(map[*Xhtml_kbd_type]string)
	_ = map_Xhtml_kbd_type_Identifiers

	xhtml_kbd_typeOrdered := []*Xhtml_kbd_type{}
	for xhtml_kbd_type := range stage.Xhtml_kbd_types {
		xhtml_kbd_typeOrdered = append(xhtml_kbd_typeOrdered, xhtml_kbd_type)
	}
	sort.Slice(xhtml_kbd_typeOrdered[:], func(i, j int) bool {
		xhtml_kbd_typei := xhtml_kbd_typeOrdered[i]
		xhtml_kbd_typej := xhtml_kbd_typeOrdered[j]
		xhtml_kbd_typei_order, oki := stage.Xhtml_kbd_typeMap_Staged_Order[xhtml_kbd_typei]
		xhtml_kbd_typej_order, okj := stage.Xhtml_kbd_typeMap_Staged_Order[xhtml_kbd_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_kbd_typei_order < xhtml_kbd_typej_order
	})
	if len(xhtml_kbd_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_kbd_type := range xhtml_kbd_typeOrdered {

		id = generatesIdentifier("Xhtml_kbd_type", idx, xhtml_kbd_type.Name)
		map_Xhtml_kbd_type_Identifiers[xhtml_kbd_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_kbd_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_kbd_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_kbd_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_li_type_Identifiers := make(map[*Xhtml_li_type]string)
	_ = map_Xhtml_li_type_Identifiers

	xhtml_li_typeOrdered := []*Xhtml_li_type{}
	for xhtml_li_type := range stage.Xhtml_li_types {
		xhtml_li_typeOrdered = append(xhtml_li_typeOrdered, xhtml_li_type)
	}
	sort.Slice(xhtml_li_typeOrdered[:], func(i, j int) bool {
		xhtml_li_typei := xhtml_li_typeOrdered[i]
		xhtml_li_typej := xhtml_li_typeOrdered[j]
		xhtml_li_typei_order, oki := stage.Xhtml_li_typeMap_Staged_Order[xhtml_li_typei]
		xhtml_li_typej_order, okj := stage.Xhtml_li_typeMap_Staged_Order[xhtml_li_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_li_typei_order < xhtml_li_typej_order
	})
	if len(xhtml_li_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_li_type := range xhtml_li_typeOrdered {

		id = generatesIdentifier("Xhtml_li_type", idx, xhtml_li_type.Name)
		map_Xhtml_li_type_Identifiers[xhtml_li_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_li_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_li_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_li_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_object_type_Identifiers := make(map[*Xhtml_object_type]string)
	_ = map_Xhtml_object_type_Identifiers

	xhtml_object_typeOrdered := []*Xhtml_object_type{}
	for xhtml_object_type := range stage.Xhtml_object_types {
		xhtml_object_typeOrdered = append(xhtml_object_typeOrdered, xhtml_object_type)
	}
	sort.Slice(xhtml_object_typeOrdered[:], func(i, j int) bool {
		xhtml_object_typei := xhtml_object_typeOrdered[i]
		xhtml_object_typej := xhtml_object_typeOrdered[j]
		xhtml_object_typei_order, oki := stage.Xhtml_object_typeMap_Staged_Order[xhtml_object_typei]
		xhtml_object_typej_order, okj := stage.Xhtml_object_typeMap_Staged_Order[xhtml_object_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_object_typei_order < xhtml_object_typej_order
	})
	if len(xhtml_object_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_object_type := range xhtml_object_typeOrdered {

		id = generatesIdentifier("Xhtml_object_type", idx, xhtml_object_type.Name)
		map_Xhtml_object_type_Identifiers[xhtml_object_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_object_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_object_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_object_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_ol_type_Identifiers := make(map[*Xhtml_ol_type]string)
	_ = map_Xhtml_ol_type_Identifiers

	xhtml_ol_typeOrdered := []*Xhtml_ol_type{}
	for xhtml_ol_type := range stage.Xhtml_ol_types {
		xhtml_ol_typeOrdered = append(xhtml_ol_typeOrdered, xhtml_ol_type)
	}
	sort.Slice(xhtml_ol_typeOrdered[:], func(i, j int) bool {
		xhtml_ol_typei := xhtml_ol_typeOrdered[i]
		xhtml_ol_typej := xhtml_ol_typeOrdered[j]
		xhtml_ol_typei_order, oki := stage.Xhtml_ol_typeMap_Staged_Order[xhtml_ol_typei]
		xhtml_ol_typej_order, okj := stage.Xhtml_ol_typeMap_Staged_Order[xhtml_ol_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_ol_typei_order < xhtml_ol_typej_order
	})
	if len(xhtml_ol_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_ol_type := range xhtml_ol_typeOrdered {

		id = generatesIdentifier("Xhtml_ol_type", idx, xhtml_ol_type.Name)
		map_Xhtml_ol_type_Identifiers[xhtml_ol_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_ol_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_ol_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_ol_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_p_type_Identifiers := make(map[*Xhtml_p_type]string)
	_ = map_Xhtml_p_type_Identifiers

	xhtml_p_typeOrdered := []*Xhtml_p_type{}
	for xhtml_p_type := range stage.Xhtml_p_types {
		xhtml_p_typeOrdered = append(xhtml_p_typeOrdered, xhtml_p_type)
	}
	sort.Slice(xhtml_p_typeOrdered[:], func(i, j int) bool {
		xhtml_p_typei := xhtml_p_typeOrdered[i]
		xhtml_p_typej := xhtml_p_typeOrdered[j]
		xhtml_p_typei_order, oki := stage.Xhtml_p_typeMap_Staged_Order[xhtml_p_typei]
		xhtml_p_typej_order, okj := stage.Xhtml_p_typeMap_Staged_Order[xhtml_p_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_p_typei_order < xhtml_p_typej_order
	})
	if len(xhtml_p_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_p_type := range xhtml_p_typeOrdered {

		id = generatesIdentifier("Xhtml_p_type", idx, xhtml_p_type.Name)
		map_Xhtml_p_type_Identifiers[xhtml_p_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_p_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_p_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_p_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_param_type_Identifiers := make(map[*Xhtml_param_type]string)
	_ = map_Xhtml_param_type_Identifiers

	xhtml_param_typeOrdered := []*Xhtml_param_type{}
	for xhtml_param_type := range stage.Xhtml_param_types {
		xhtml_param_typeOrdered = append(xhtml_param_typeOrdered, xhtml_param_type)
	}
	sort.Slice(xhtml_param_typeOrdered[:], func(i, j int) bool {
		xhtml_param_typei := xhtml_param_typeOrdered[i]
		xhtml_param_typej := xhtml_param_typeOrdered[j]
		xhtml_param_typei_order, oki := stage.Xhtml_param_typeMap_Staged_Order[xhtml_param_typei]
		xhtml_param_typej_order, okj := stage.Xhtml_param_typeMap_Staged_Order[xhtml_param_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_param_typei_order < xhtml_param_typej_order
	})
	if len(xhtml_param_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_param_type := range xhtml_param_typeOrdered {

		id = generatesIdentifier("Xhtml_param_type", idx, xhtml_param_type.Name)
		map_Xhtml_param_type_Identifiers[xhtml_param_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_param_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_param_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_param_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_pre_type_Identifiers := make(map[*Xhtml_pre_type]string)
	_ = map_Xhtml_pre_type_Identifiers

	xhtml_pre_typeOrdered := []*Xhtml_pre_type{}
	for xhtml_pre_type := range stage.Xhtml_pre_types {
		xhtml_pre_typeOrdered = append(xhtml_pre_typeOrdered, xhtml_pre_type)
	}
	sort.Slice(xhtml_pre_typeOrdered[:], func(i, j int) bool {
		xhtml_pre_typei := xhtml_pre_typeOrdered[i]
		xhtml_pre_typej := xhtml_pre_typeOrdered[j]
		xhtml_pre_typei_order, oki := stage.Xhtml_pre_typeMap_Staged_Order[xhtml_pre_typei]
		xhtml_pre_typej_order, okj := stage.Xhtml_pre_typeMap_Staged_Order[xhtml_pre_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_pre_typei_order < xhtml_pre_typej_order
	})
	if len(xhtml_pre_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_pre_type := range xhtml_pre_typeOrdered {

		id = generatesIdentifier("Xhtml_pre_type", idx, xhtml_pre_type.Name)
		map_Xhtml_pre_type_Identifiers[xhtml_pre_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_pre_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_pre_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_pre_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_q_type_Identifiers := make(map[*Xhtml_q_type]string)
	_ = map_Xhtml_q_type_Identifiers

	xhtml_q_typeOrdered := []*Xhtml_q_type{}
	for xhtml_q_type := range stage.Xhtml_q_types {
		xhtml_q_typeOrdered = append(xhtml_q_typeOrdered, xhtml_q_type)
	}
	sort.Slice(xhtml_q_typeOrdered[:], func(i, j int) bool {
		xhtml_q_typei := xhtml_q_typeOrdered[i]
		xhtml_q_typej := xhtml_q_typeOrdered[j]
		xhtml_q_typei_order, oki := stage.Xhtml_q_typeMap_Staged_Order[xhtml_q_typei]
		xhtml_q_typej_order, okj := stage.Xhtml_q_typeMap_Staged_Order[xhtml_q_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_q_typei_order < xhtml_q_typej_order
	})
	if len(xhtml_q_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_q_type := range xhtml_q_typeOrdered {

		id = generatesIdentifier("Xhtml_q_type", idx, xhtml_q_type.Name)
		map_Xhtml_q_type_Identifiers[xhtml_q_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_q_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_q_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_q_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_samp_type_Identifiers := make(map[*Xhtml_samp_type]string)
	_ = map_Xhtml_samp_type_Identifiers

	xhtml_samp_typeOrdered := []*Xhtml_samp_type{}
	for xhtml_samp_type := range stage.Xhtml_samp_types {
		xhtml_samp_typeOrdered = append(xhtml_samp_typeOrdered, xhtml_samp_type)
	}
	sort.Slice(xhtml_samp_typeOrdered[:], func(i, j int) bool {
		xhtml_samp_typei := xhtml_samp_typeOrdered[i]
		xhtml_samp_typej := xhtml_samp_typeOrdered[j]
		xhtml_samp_typei_order, oki := stage.Xhtml_samp_typeMap_Staged_Order[xhtml_samp_typei]
		xhtml_samp_typej_order, okj := stage.Xhtml_samp_typeMap_Staged_Order[xhtml_samp_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_samp_typei_order < xhtml_samp_typej_order
	})
	if len(xhtml_samp_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_samp_type := range xhtml_samp_typeOrdered {

		id = generatesIdentifier("Xhtml_samp_type", idx, xhtml_samp_type.Name)
		map_Xhtml_samp_type_Identifiers[xhtml_samp_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_samp_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_samp_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_samp_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_span_type_Identifiers := make(map[*Xhtml_span_type]string)
	_ = map_Xhtml_span_type_Identifiers

	xhtml_span_typeOrdered := []*Xhtml_span_type{}
	for xhtml_span_type := range stage.Xhtml_span_types {
		xhtml_span_typeOrdered = append(xhtml_span_typeOrdered, xhtml_span_type)
	}
	sort.Slice(xhtml_span_typeOrdered[:], func(i, j int) bool {
		xhtml_span_typei := xhtml_span_typeOrdered[i]
		xhtml_span_typej := xhtml_span_typeOrdered[j]
		xhtml_span_typei_order, oki := stage.Xhtml_span_typeMap_Staged_Order[xhtml_span_typei]
		xhtml_span_typej_order, okj := stage.Xhtml_span_typeMap_Staged_Order[xhtml_span_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_span_typei_order < xhtml_span_typej_order
	})
	if len(xhtml_span_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_span_type := range xhtml_span_typeOrdered {

		id = generatesIdentifier("Xhtml_span_type", idx, xhtml_span_type.Name)
		map_Xhtml_span_type_Identifiers[xhtml_span_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_span_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_span_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_span_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_strong_type_Identifiers := make(map[*Xhtml_strong_type]string)
	_ = map_Xhtml_strong_type_Identifiers

	xhtml_strong_typeOrdered := []*Xhtml_strong_type{}
	for xhtml_strong_type := range stage.Xhtml_strong_types {
		xhtml_strong_typeOrdered = append(xhtml_strong_typeOrdered, xhtml_strong_type)
	}
	sort.Slice(xhtml_strong_typeOrdered[:], func(i, j int) bool {
		xhtml_strong_typei := xhtml_strong_typeOrdered[i]
		xhtml_strong_typej := xhtml_strong_typeOrdered[j]
		xhtml_strong_typei_order, oki := stage.Xhtml_strong_typeMap_Staged_Order[xhtml_strong_typei]
		xhtml_strong_typej_order, okj := stage.Xhtml_strong_typeMap_Staged_Order[xhtml_strong_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_strong_typei_order < xhtml_strong_typej_order
	})
	if len(xhtml_strong_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_strong_type := range xhtml_strong_typeOrdered {

		id = generatesIdentifier("Xhtml_strong_type", idx, xhtml_strong_type.Name)
		map_Xhtml_strong_type_Identifiers[xhtml_strong_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_strong_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_strong_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_strong_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_table_type_Identifiers := make(map[*Xhtml_table_type]string)
	_ = map_Xhtml_table_type_Identifiers

	xhtml_table_typeOrdered := []*Xhtml_table_type{}
	for xhtml_table_type := range stage.Xhtml_table_types {
		xhtml_table_typeOrdered = append(xhtml_table_typeOrdered, xhtml_table_type)
	}
	sort.Slice(xhtml_table_typeOrdered[:], func(i, j int) bool {
		xhtml_table_typei := xhtml_table_typeOrdered[i]
		xhtml_table_typej := xhtml_table_typeOrdered[j]
		xhtml_table_typei_order, oki := stage.Xhtml_table_typeMap_Staged_Order[xhtml_table_typei]
		xhtml_table_typej_order, okj := stage.Xhtml_table_typeMap_Staged_Order[xhtml_table_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_table_typei_order < xhtml_table_typej_order
	})
	if len(xhtml_table_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_table_type := range xhtml_table_typeOrdered {

		id = generatesIdentifier("Xhtml_table_type", idx, xhtml_table_type.Name)
		map_Xhtml_table_type_Identifiers[xhtml_table_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_table_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_table_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_table_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_tbody_type_Identifiers := make(map[*Xhtml_tbody_type]string)
	_ = map_Xhtml_tbody_type_Identifiers

	xhtml_tbody_typeOrdered := []*Xhtml_tbody_type{}
	for xhtml_tbody_type := range stage.Xhtml_tbody_types {
		xhtml_tbody_typeOrdered = append(xhtml_tbody_typeOrdered, xhtml_tbody_type)
	}
	sort.Slice(xhtml_tbody_typeOrdered[:], func(i, j int) bool {
		xhtml_tbody_typei := xhtml_tbody_typeOrdered[i]
		xhtml_tbody_typej := xhtml_tbody_typeOrdered[j]
		xhtml_tbody_typei_order, oki := stage.Xhtml_tbody_typeMap_Staged_Order[xhtml_tbody_typei]
		xhtml_tbody_typej_order, okj := stage.Xhtml_tbody_typeMap_Staged_Order[xhtml_tbody_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_tbody_typei_order < xhtml_tbody_typej_order
	})
	if len(xhtml_tbody_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_tbody_type := range xhtml_tbody_typeOrdered {

		id = generatesIdentifier("Xhtml_tbody_type", idx, xhtml_tbody_type.Name)
		map_Xhtml_tbody_type_Identifiers[xhtml_tbody_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_tbody_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_tbody_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_tbody_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_td_type_Identifiers := make(map[*Xhtml_td_type]string)
	_ = map_Xhtml_td_type_Identifiers

	xhtml_td_typeOrdered := []*Xhtml_td_type{}
	for xhtml_td_type := range stage.Xhtml_td_types {
		xhtml_td_typeOrdered = append(xhtml_td_typeOrdered, xhtml_td_type)
	}
	sort.Slice(xhtml_td_typeOrdered[:], func(i, j int) bool {
		xhtml_td_typei := xhtml_td_typeOrdered[i]
		xhtml_td_typej := xhtml_td_typeOrdered[j]
		xhtml_td_typei_order, oki := stage.Xhtml_td_typeMap_Staged_Order[xhtml_td_typei]
		xhtml_td_typej_order, okj := stage.Xhtml_td_typeMap_Staged_Order[xhtml_td_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_td_typei_order < xhtml_td_typej_order
	})
	if len(xhtml_td_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_td_type := range xhtml_td_typeOrdered {

		id = generatesIdentifier("Xhtml_td_type", idx, xhtml_td_type.Name)
		map_Xhtml_td_type_Identifiers[xhtml_td_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_td_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_td_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_td_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_tfoot_type_Identifiers := make(map[*Xhtml_tfoot_type]string)
	_ = map_Xhtml_tfoot_type_Identifiers

	xhtml_tfoot_typeOrdered := []*Xhtml_tfoot_type{}
	for xhtml_tfoot_type := range stage.Xhtml_tfoot_types {
		xhtml_tfoot_typeOrdered = append(xhtml_tfoot_typeOrdered, xhtml_tfoot_type)
	}
	sort.Slice(xhtml_tfoot_typeOrdered[:], func(i, j int) bool {
		xhtml_tfoot_typei := xhtml_tfoot_typeOrdered[i]
		xhtml_tfoot_typej := xhtml_tfoot_typeOrdered[j]
		xhtml_tfoot_typei_order, oki := stage.Xhtml_tfoot_typeMap_Staged_Order[xhtml_tfoot_typei]
		xhtml_tfoot_typej_order, okj := stage.Xhtml_tfoot_typeMap_Staged_Order[xhtml_tfoot_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_tfoot_typei_order < xhtml_tfoot_typej_order
	})
	if len(xhtml_tfoot_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_tfoot_type := range xhtml_tfoot_typeOrdered {

		id = generatesIdentifier("Xhtml_tfoot_type", idx, xhtml_tfoot_type.Name)
		map_Xhtml_tfoot_type_Identifiers[xhtml_tfoot_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_tfoot_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_tfoot_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_tfoot_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_th_type_Identifiers := make(map[*Xhtml_th_type]string)
	_ = map_Xhtml_th_type_Identifiers

	xhtml_th_typeOrdered := []*Xhtml_th_type{}
	for xhtml_th_type := range stage.Xhtml_th_types {
		xhtml_th_typeOrdered = append(xhtml_th_typeOrdered, xhtml_th_type)
	}
	sort.Slice(xhtml_th_typeOrdered[:], func(i, j int) bool {
		xhtml_th_typei := xhtml_th_typeOrdered[i]
		xhtml_th_typej := xhtml_th_typeOrdered[j]
		xhtml_th_typei_order, oki := stage.Xhtml_th_typeMap_Staged_Order[xhtml_th_typei]
		xhtml_th_typej_order, okj := stage.Xhtml_th_typeMap_Staged_Order[xhtml_th_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_th_typei_order < xhtml_th_typej_order
	})
	if len(xhtml_th_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_th_type := range xhtml_th_typeOrdered {

		id = generatesIdentifier("Xhtml_th_type", idx, xhtml_th_type.Name)
		map_Xhtml_th_type_Identifiers[xhtml_th_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_th_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_th_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_th_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_thead_type_Identifiers := make(map[*Xhtml_thead_type]string)
	_ = map_Xhtml_thead_type_Identifiers

	xhtml_thead_typeOrdered := []*Xhtml_thead_type{}
	for xhtml_thead_type := range stage.Xhtml_thead_types {
		xhtml_thead_typeOrdered = append(xhtml_thead_typeOrdered, xhtml_thead_type)
	}
	sort.Slice(xhtml_thead_typeOrdered[:], func(i, j int) bool {
		xhtml_thead_typei := xhtml_thead_typeOrdered[i]
		xhtml_thead_typej := xhtml_thead_typeOrdered[j]
		xhtml_thead_typei_order, oki := stage.Xhtml_thead_typeMap_Staged_Order[xhtml_thead_typei]
		xhtml_thead_typej_order, okj := stage.Xhtml_thead_typeMap_Staged_Order[xhtml_thead_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_thead_typei_order < xhtml_thead_typej_order
	})
	if len(xhtml_thead_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_thead_type := range xhtml_thead_typeOrdered {

		id = generatesIdentifier("Xhtml_thead_type", idx, xhtml_thead_type.Name)
		map_Xhtml_thead_type_Identifiers[xhtml_thead_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_thead_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_thead_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_thead_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_tr_type_Identifiers := make(map[*Xhtml_tr_type]string)
	_ = map_Xhtml_tr_type_Identifiers

	xhtml_tr_typeOrdered := []*Xhtml_tr_type{}
	for xhtml_tr_type := range stage.Xhtml_tr_types {
		xhtml_tr_typeOrdered = append(xhtml_tr_typeOrdered, xhtml_tr_type)
	}
	sort.Slice(xhtml_tr_typeOrdered[:], func(i, j int) bool {
		xhtml_tr_typei := xhtml_tr_typeOrdered[i]
		xhtml_tr_typej := xhtml_tr_typeOrdered[j]
		xhtml_tr_typei_order, oki := stage.Xhtml_tr_typeMap_Staged_Order[xhtml_tr_typei]
		xhtml_tr_typej_order, okj := stage.Xhtml_tr_typeMap_Staged_Order[xhtml_tr_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_tr_typei_order < xhtml_tr_typej_order
	})
	if len(xhtml_tr_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_tr_type := range xhtml_tr_typeOrdered {

		id = generatesIdentifier("Xhtml_tr_type", idx, xhtml_tr_type.Name)
		map_Xhtml_tr_type_Identifiers[xhtml_tr_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_tr_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_tr_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_tr_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_ul_type_Identifiers := make(map[*Xhtml_ul_type]string)
	_ = map_Xhtml_ul_type_Identifiers

	xhtml_ul_typeOrdered := []*Xhtml_ul_type{}
	for xhtml_ul_type := range stage.Xhtml_ul_types {
		xhtml_ul_typeOrdered = append(xhtml_ul_typeOrdered, xhtml_ul_type)
	}
	sort.Slice(xhtml_ul_typeOrdered[:], func(i, j int) bool {
		xhtml_ul_typei := xhtml_ul_typeOrdered[i]
		xhtml_ul_typej := xhtml_ul_typeOrdered[j]
		xhtml_ul_typei_order, oki := stage.Xhtml_ul_typeMap_Staged_Order[xhtml_ul_typei]
		xhtml_ul_typej_order, okj := stage.Xhtml_ul_typeMap_Staged_Order[xhtml_ul_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_ul_typei_order < xhtml_ul_typej_order
	})
	if len(xhtml_ul_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_ul_type := range xhtml_ul_typeOrdered {

		id = generatesIdentifier("Xhtml_ul_type", idx, xhtml_ul_type.Name)
		map_Xhtml_ul_type_Identifiers[xhtml_ul_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_ul_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_ul_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_ul_type.Name))
		initializerStatements += setValueField

	}

	map_Xhtml_var_type_Identifiers := make(map[*Xhtml_var_type]string)
	_ = map_Xhtml_var_type_Identifiers

	xhtml_var_typeOrdered := []*Xhtml_var_type{}
	for xhtml_var_type := range stage.Xhtml_var_types {
		xhtml_var_typeOrdered = append(xhtml_var_typeOrdered, xhtml_var_type)
	}
	sort.Slice(xhtml_var_typeOrdered[:], func(i, j int) bool {
		xhtml_var_typei := xhtml_var_typeOrdered[i]
		xhtml_var_typej := xhtml_var_typeOrdered[j]
		xhtml_var_typei_order, oki := stage.Xhtml_var_typeMap_Staged_Order[xhtml_var_typei]
		xhtml_var_typej_order, okj := stage.Xhtml_var_typeMap_Staged_Order[xhtml_var_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_var_typei_order < xhtml_var_typej_order
	})
	if len(xhtml_var_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_var_type := range xhtml_var_typeOrdered {

		id = generatesIdentifier("Xhtml_var_type", idx, xhtml_var_type.Name)
		map_Xhtml_var_type_Identifiers[xhtml_var_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xhtml_var_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_var_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_var_type.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(alternative_idOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ALTERNATIVE_ID instances pointers"
	}
	for idx, alternative_id := range alternative_idOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ALTERNATIVE_ID", idx, alternative_id.Name)
		map_ALTERNATIVE_ID_Identifiers[alternative_id] = id

		// Initialisation of values
	}

	if len(attribute_definition_booleanOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_BOOLEAN instances pointers"
	}
	for idx, attribute_definition_boolean := range attribute_definition_booleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_BOOLEAN", idx, attribute_definition_boolean.Name)
		map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[attribute_definition_boolean] = id

		// Initialisation of values
		for _, _alternative_id := range attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_boolean := range attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[_attribute_value_boolean])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_dateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_DATE instances pointers"
	}
	for idx, attribute_definition_date := range attribute_definition_dateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_DATE", idx, attribute_definition_date.Name)
		map_ATTRIBUTE_DEFINITION_DATE_Identifiers[attribute_definition_date] = id

		// Initialisation of values
		for _, _alternative_id := range attribute_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_date := range attribute_definition_date.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_DATE_Identifiers[_attribute_value_date])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_enumerationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_ENUMERATION instances pointers"
	}
	for idx, attribute_definition_enumeration := range attribute_definition_enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_ENUMERATION", idx, attribute_definition_enumeration.Name)
		map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[attribute_definition_enumeration] = id

		// Initialisation of values
		for _, _attribute_value_enumeration := range attribute_definition_enumeration.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[_attribute_value_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _alternative_id := range attribute_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_integerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_INTEGER instances pointers"
	}
	for idx, attribute_definition_integer := range attribute_definition_integerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_INTEGER", idx, attribute_definition_integer.Name)
		map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[attribute_definition_integer] = id

		// Initialisation of values
		for _, _alternative_id := range attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_integer := range attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_INTEGER_Identifiers[_attribute_value_integer])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_realOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_REAL instances pointers"
	}
	for idx, attribute_definition_real := range attribute_definition_realOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_REAL", idx, attribute_definition_real.Name)
		map_ATTRIBUTE_DEFINITION_REAL_Identifiers[attribute_definition_real] = id

		// Initialisation of values
		for _, _alternative_id := range attribute_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_real := range attribute_definition_real.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_REAL_Identifiers[_attribute_value_real])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_stringOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_STRING instances pointers"
	}
	for idx, attribute_definition_string := range attribute_definition_stringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_STRING", idx, attribute_definition_string.Name)
		map_ATTRIBUTE_DEFINITION_STRING_Identifiers[attribute_definition_string] = id

		// Initialisation of values
		for _, _alternative_id := range attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_string := range attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_STRING_Identifiers[_attribute_value_string])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_definition_xhtmlOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_DEFINITION_XHTML instances pointers"
	}
	for idx, attribute_definition_xhtml := range attribute_definition_xhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_XHTML", idx, attribute_definition_xhtml.Name)
		map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[attribute_definition_xhtml] = id

		// Initialisation of values
		for _, _alternative_id := range attribute_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_xhtml := range attribute_definition_xhtml.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_XHTML_Identifiers[_attribute_value_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(attribute_value_booleanOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_BOOLEAN instances pointers"
	}
	for idx, attribute_value_boolean := range attribute_value_booleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_BOOLEAN", idx, attribute_value_boolean.Name)
		map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[attribute_value_boolean] = id

		// Initialisation of values
	}

	if len(attribute_value_dateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_DATE instances pointers"
	}
	for idx, attribute_value_date := range attribute_value_dateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_DATE", idx, attribute_value_date.Name)
		map_ATTRIBUTE_VALUE_DATE_Identifiers[attribute_value_date] = id

		// Initialisation of values
	}

	if len(attribute_value_enumerationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_ENUMERATION instances pointers"
	}
	for idx, attribute_value_enumeration := range attribute_value_enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_ENUMERATION", idx, attribute_value_enumeration.Name)
		map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[attribute_value_enumeration] = id

		// Initialisation of values
	}

	if len(attribute_value_integerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_INTEGER instances pointers"
	}
	for idx, attribute_value_integer := range attribute_value_integerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_INTEGER", idx, attribute_value_integer.Name)
		map_ATTRIBUTE_VALUE_INTEGER_Identifiers[attribute_value_integer] = id

		// Initialisation of values
	}

	if len(attribute_value_realOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_REAL instances pointers"
	}
	for idx, attribute_value_real := range attribute_value_realOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_REAL", idx, attribute_value_real.Name)
		map_ATTRIBUTE_VALUE_REAL_Identifiers[attribute_value_real] = id

		// Initialisation of values
	}

	if len(attribute_value_stringOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_STRING instances pointers"
	}
	for idx, attribute_value_string := range attribute_value_stringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_STRING", idx, attribute_value_string.Name)
		map_ATTRIBUTE_VALUE_STRING_Identifiers[attribute_value_string] = id

		// Initialisation of values
	}

	if len(attribute_value_xhtmlOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ATTRIBUTE_VALUE_XHTML instances pointers"
	}
	for idx, attribute_value_xhtml := range attribute_value_xhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_XHTML", idx, attribute_value_xhtml.Name)
		map_ATTRIBUTE_VALUE_XHTML_Identifiers[attribute_value_xhtml] = id

		// Initialisation of values
		for _, _xhtml_content := range attribute_value_xhtml.THE_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THE_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XHTML_CONTENT_Identifiers[_xhtml_content])
			pointersInitializesStatements += setPointerField
		}

		for _, _xhtml_content := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THE_ORIGINAL_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XHTML_CONTENT_Identifiers[_xhtml_content])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(anytypeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AnyType instances pointers"
	}
	for idx, anytype := range anytypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AnyType", idx, anytype.Name)
		map_AnyType_Identifiers[anytype] = id

		// Initialisation of values
	}

	if len(datatype_definition_booleanOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_BOOLEAN instances pointers"
	}
	for idx, datatype_definition_boolean := range datatype_definition_booleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_BOOLEAN", idx, datatype_definition_boolean.Name)
		map_DATATYPE_DEFINITION_BOOLEAN_Identifiers[datatype_definition_boolean] = id

		// Initialisation of values
		for _, _alternative_id := range datatype_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_dateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_DATE instances pointers"
	}
	for idx, datatype_definition_date := range datatype_definition_dateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_DATE", idx, datatype_definition_date.Name)
		map_DATATYPE_DEFINITION_DATE_Identifiers[datatype_definition_date] = id

		// Initialisation of values
		for _, _alternative_id := range datatype_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_enumerationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_ENUMERATION instances pointers"
	}
	for idx, datatype_definition_enumeration := range datatype_definition_enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_ENUMERATION", idx, datatype_definition_enumeration.Name)
		map_DATATYPE_DEFINITION_ENUMERATION_Identifiers[datatype_definition_enumeration] = id

		// Initialisation of values
		for _, _alternative_id := range datatype_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _enum_value := range datatype_definition_enumeration.SPECIFIED_VALUES.ENUM_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFIED_VALUES.ENUM_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ENUM_VALUE_Identifiers[_enum_value])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_integerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_INTEGER instances pointers"
	}
	for idx, datatype_definition_integer := range datatype_definition_integerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_INTEGER", idx, datatype_definition_integer.Name)
		map_DATATYPE_DEFINITION_INTEGER_Identifiers[datatype_definition_integer] = id

		// Initialisation of values
		for _, _alternative_id := range datatype_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_realOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_REAL instances pointers"
	}
	for idx, datatype_definition_real := range datatype_definition_realOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_REAL", idx, datatype_definition_real.Name)
		map_DATATYPE_DEFINITION_REAL_Identifiers[datatype_definition_real] = id

		// Initialisation of values
		for _, _alternative_id := range datatype_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_stringOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_STRING instances pointers"
	}
	for idx, datatype_definition_string := range datatype_definition_stringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_STRING", idx, datatype_definition_string.Name)
		map_DATATYPE_DEFINITION_STRING_Identifiers[datatype_definition_string] = id

		// Initialisation of values
		for _, _alternative_id := range datatype_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(datatype_definition_xhtmlOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DATATYPE_DEFINITION_XHTML instances pointers"
	}
	for idx, datatype_definition_xhtml := range datatype_definition_xhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_XHTML", idx, datatype_definition_xhtml.Name)
		map_DATATYPE_DEFINITION_XHTML_Identifiers[datatype_definition_xhtml] = id

		// Initialisation of values
		for _, _alternative_id := range datatype_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(embedded_valueOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of EMBEDDED_VALUE instances pointers"
	}
	for idx, embedded_value := range embedded_valueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("EMBEDDED_VALUE", idx, embedded_value.Name)
		map_EMBEDDED_VALUE_Identifiers[embedded_value] = id

		// Initialisation of values
	}

	if len(enum_valueOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ENUM_VALUE instances pointers"
	}
	for idx, enum_value := range enum_valueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ENUM_VALUE", idx, enum_value.Name)
		map_ENUM_VALUE_Identifiers[enum_value] = id

		// Initialisation of values
		for _, _alternative_id := range enum_value.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _embedded_value := range enum_value.PROPERTIES.EMBEDDED_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "PROPERTIES.EMBEDDED_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_EMBEDDED_VALUE_Identifiers[_embedded_value])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(relation_groupOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of RELATION_GROUP instances pointers"
	}
	for idx, relation_group := range relation_groupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RELATION_GROUP", idx, relation_group.Name)
		map_RELATION_GROUP_Identifiers[relation_group] = id

		// Initialisation of values
		for _, _alternative_id := range relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(relation_group_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of RELATION_GROUP_TYPE instances pointers"
	}
	for idx, relation_group_type := range relation_group_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RELATION_GROUP_TYPE", idx, relation_group_type.Name)
		map_RELATION_GROUP_TYPE_Identifiers[relation_group_type] = id

		// Initialisation of values
		for _, _alternative_id := range relation_group_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_boolean := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[_attribute_definition_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_date := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_DATE_Identifiers[_attribute_definition_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_enumeration := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[_attribute_definition_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_integer := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[_attribute_definition_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_real := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_REAL_Identifiers[_attribute_definition_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_string := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_STRING_Identifiers[_attribute_definition_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_xhtml := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[_attribute_definition_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(req_ifOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of REQ_IF instances pointers"
	}
	for idx, req_if := range req_ifOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF", idx, req_if.Name)
		map_REQ_IF_Identifiers[req_if] = id

		// Initialisation of values
		if req_if.THE_HEADER.REQ_IF_HEADER != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THE_HEADER.REQ_IF_HEADER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQ_IF_HEADER_Identifiers[req_if.THE_HEADER.REQ_IF_HEADER])
			pointersInitializesStatements += setPointerField
		}

		if req_if.CORE_CONTENT.REQ_IF_CONTENT != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CORE_CONTENT.REQ_IF_CONTENT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQ_IF_CONTENT_Identifiers[req_if.CORE_CONTENT.REQ_IF_CONTENT])
			pointersInitializesStatements += setPointerField
		}

		for _, _req_if_tool_extension := range req_if.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQ_IF_TOOL_EXTENSION_Identifiers[_req_if_tool_extension])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(req_if_contentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of REQ_IF_CONTENT instances pointers"
	}
	for idx, req_if_content := range req_if_contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF_CONTENT", idx, req_if_content.Name)
		map_REQ_IF_CONTENT_Identifiers[req_if_content] = id

		// Initialisation of values
		for _, _datatype_definition_boolean := range req_if_content.DATATYPES.DATATYPE_DEFINITION_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPES.DATATYPE_DEFINITION_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_BOOLEAN_Identifiers[_datatype_definition_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_date := range req_if_content.DATATYPES.DATATYPE_DEFINITION_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPES.DATATYPE_DEFINITION_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_DATE_Identifiers[_datatype_definition_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_enumeration := range req_if_content.DATATYPES.DATATYPE_DEFINITION_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPES.DATATYPE_DEFINITION_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_ENUMERATION_Identifiers[_datatype_definition_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_integer := range req_if_content.DATATYPES.DATATYPE_DEFINITION_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPES.DATATYPE_DEFINITION_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_INTEGER_Identifiers[_datatype_definition_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_real := range req_if_content.DATATYPES.DATATYPE_DEFINITION_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPES.DATATYPE_DEFINITION_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_REAL_Identifiers[_datatype_definition_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_string := range req_if_content.DATATYPES.DATATYPE_DEFINITION_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPES.DATATYPE_DEFINITION_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_STRING_Identifiers[_datatype_definition_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_xhtml := range req_if_content.DATATYPES.DATATYPE_DEFINITION_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPES.DATATYPE_DEFINITION_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_XHTML_Identifiers[_datatype_definition_xhtml])
			pointersInitializesStatements += setPointerField
		}

		for _, _relation_group_type := range req_if_content.SPEC_TYPES.RELATION_GROUP_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_TYPES.RELATION_GROUP_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RELATION_GROUP_TYPE_Identifiers[_relation_group_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _spec_object_type := range req_if_content.SPEC_TYPES.SPEC_OBJECT_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_TYPES.SPEC_OBJECT_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_OBJECT_TYPE_Identifiers[_spec_object_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _spec_relation_type := range req_if_content.SPEC_TYPES.SPEC_RELATION_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_TYPES.SPEC_RELATION_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_RELATION_TYPE_Identifiers[_spec_relation_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _specification_type := range req_if_content.SPEC_TYPES.SPECIFICATION_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_TYPES.SPECIFICATION_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECIFICATION_TYPE_Identifiers[_specification_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _spec_object := range req_if_content.SPEC_OBJECTS.SPEC_OBJECT {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_OBJECTS.SPEC_OBJECT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_OBJECT_Identifiers[_spec_object])
			pointersInitializesStatements += setPointerField
		}

		for _, _spec_relation := range req_if_content.SPEC_RELATIONS.SPEC_RELATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATIONS.SPEC_RELATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_RELATION_Identifiers[_spec_relation])
			pointersInitializesStatements += setPointerField
		}

		for _, _specification := range req_if_content.SPECIFICATIONS.SPECIFICATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFICATIONS.SPECIFICATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECIFICATION_Identifiers[_specification])
			pointersInitializesStatements += setPointerField
		}

		for _, _relation_group := range req_if_content.SPEC_RELATION_GROUPS.RELATION_GROUP {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATION_GROUPS.RELATION_GROUP")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RELATION_GROUP_Identifiers[_relation_group])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(req_if_headerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of REQ_IF_HEADER instances pointers"
	}
	for idx, req_if_header := range req_if_headerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF_HEADER", idx, req_if_header.Name)
		map_REQ_IF_HEADER_Identifiers[req_if_header] = id

		// Initialisation of values
	}

	if len(req_if_tool_extensionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of REQ_IF_TOOL_EXTENSION instances pointers"
	}
	for idx, req_if_tool_extension := range req_if_tool_extensionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF_TOOL_EXTENSION", idx, req_if_tool_extension.Name)
		map_REQ_IF_TOOL_EXTENSION_Identifiers[req_if_tool_extension] = id

		// Initialisation of values
	}

	if len(specificationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPECIFICATION instances pointers"
	}
	for idx, specification := range specificationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFICATION", idx, specification.Name)
		map_SPECIFICATION_Identifiers[specification] = id

		// Initialisation of values
		for _, _alternative_id := range specification.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_boolean := range specification.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[_attribute_value_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_date := range specification.VALUES.ATTRIBUTE_VALUE_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_DATE_Identifiers[_attribute_value_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_enumeration := range specification.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[_attribute_value_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_integer := range specification.VALUES.ATTRIBUTE_VALUE_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_INTEGER_Identifiers[_attribute_value_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_real := range specification.VALUES.ATTRIBUTE_VALUE_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_REAL_Identifiers[_attribute_value_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_string := range specification.VALUES.ATTRIBUTE_VALUE_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_STRING_Identifiers[_attribute_value_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_xhtml := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_XHTML_Identifiers[_attribute_value_xhtml])
			pointersInitializesStatements += setPointerField
		}

		for _, _spec_hierarchy := range specification.CHILDREN.SPEC_HIERARCHY {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CHILDREN.SPEC_HIERARCHY")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_HIERARCHY_Identifiers[_spec_hierarchy])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(specification_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPECIFICATION_TYPE instances pointers"
	}
	for idx, specification_type := range specification_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFICATION_TYPE", idx, specification_type.Name)
		map_SPECIFICATION_TYPE_Identifiers[specification_type] = id

		// Initialisation of values
		for _, _alternative_id := range specification_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_boolean := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[_attribute_definition_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_date := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_DATE_Identifiers[_attribute_definition_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_enumeration := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[_attribute_definition_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_integer := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[_attribute_definition_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_real := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_REAL_Identifiers[_attribute_definition_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_string := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_STRING_Identifiers[_attribute_definition_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_xhtml := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[_attribute_definition_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spec_hierarchyOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_HIERARCHY instances pointers"
	}
	for idx, spec_hierarchy := range spec_hierarchyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_HIERARCHY", idx, spec_hierarchy.Name)
		map_SPEC_HIERARCHY_Identifiers[spec_hierarchy] = id

		// Initialisation of values
		for _, _alternative_id := range spec_hierarchy.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _spec_hierarchy := range spec_hierarchy.CHILDREN.SPEC_HIERARCHY {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CHILDREN.SPEC_HIERARCHY")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_HIERARCHY_Identifiers[_spec_hierarchy])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spec_objectOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_OBJECT instances pointers"
	}
	for idx, spec_object := range spec_objectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_OBJECT", idx, spec_object.Name)
		map_SPEC_OBJECT_Identifiers[spec_object] = id

		// Initialisation of values
		for _, _alternative_id := range spec_object.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_boolean := range spec_object.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[_attribute_value_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_date := range spec_object.VALUES.ATTRIBUTE_VALUE_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_DATE_Identifiers[_attribute_value_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_enumeration := range spec_object.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[_attribute_value_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_integer := range spec_object.VALUES.ATTRIBUTE_VALUE_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_INTEGER_Identifiers[_attribute_value_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_real := range spec_object.VALUES.ATTRIBUTE_VALUE_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_REAL_Identifiers[_attribute_value_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_string := range spec_object.VALUES.ATTRIBUTE_VALUE_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_STRING_Identifiers[_attribute_value_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_xhtml := range spec_object.VALUES.ATTRIBUTE_VALUE_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_XHTML_Identifiers[_attribute_value_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spec_object_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_OBJECT_TYPE instances pointers"
	}
	for idx, spec_object_type := range spec_object_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_OBJECT_TYPE", idx, spec_object_type.Name)
		map_SPEC_OBJECT_TYPE_Identifiers[spec_object_type] = id

		// Initialisation of values
		for _, _alternative_id := range spec_object_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_boolean := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[_attribute_definition_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_date := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_DATE_Identifiers[_attribute_definition_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_enumeration := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[_attribute_definition_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_integer := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[_attribute_definition_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_real := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_REAL_Identifiers[_attribute_definition_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_string := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_STRING_Identifiers[_attribute_definition_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_xhtml := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[_attribute_definition_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spec_relationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_RELATION instances pointers"
	}
	for idx, spec_relation := range spec_relationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_RELATION", idx, spec_relation.Name)
		map_SPEC_RELATION_Identifiers[spec_relation] = id

		// Initialisation of values
		for _, _alternative_id := range spec_relation.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_boolean := range spec_relation.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[_attribute_value_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_date := range spec_relation.VALUES.ATTRIBUTE_VALUE_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_DATE_Identifiers[_attribute_value_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_enumeration := range spec_relation.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[_attribute_value_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_integer := range spec_relation.VALUES.ATTRIBUTE_VALUE_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_INTEGER_Identifiers[_attribute_value_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_real := range spec_relation.VALUES.ATTRIBUTE_VALUE_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_REAL_Identifiers[_attribute_value_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_string := range spec_relation.VALUES.ATTRIBUTE_VALUE_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_STRING_Identifiers[_attribute_value_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_xhtml := range spec_relation.VALUES.ATTRIBUTE_VALUE_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES.ATTRIBUTE_VALUE_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_XHTML_Identifiers[_attribute_value_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spec_relation_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SPEC_RELATION_TYPE instances pointers"
	}
	for idx, spec_relation_type := range spec_relation_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_RELATION_TYPE", idx, spec_relation_type.Name)
		map_SPEC_RELATION_TYPE_Identifiers[spec_relation_type] = id

		// Initialisation of values
		for _, _alternative_id := range spec_relation_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID.ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_boolean := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[_attribute_definition_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_date := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_DATE_Identifiers[_attribute_definition_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_enumeration := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[_attribute_definition_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_integer := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[_attribute_definition_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_real := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_REAL_Identifiers[_attribute_definition_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_string := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_STRING_Identifiers[_attribute_definition_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_xhtml := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[_attribute_definition_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(xhtml_contentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of XHTML_CONTENT instances pointers"
	}
	for idx, xhtml_content := range xhtml_contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("XHTML_CONTENT", idx, xhtml_content.Name)
		map_XHTML_CONTENT_Identifiers[xhtml_content] = id

		// Initialisation of values
	}

	if len(xhtml_inlpres_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_InlPres_type instances pointers"
	}
	for idx, xhtml_inlpres_type := range xhtml_inlpres_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_InlPres_type", idx, xhtml_inlpres_type.Name)
		map_Xhtml_InlPres_type_Identifiers[xhtml_inlpres_type] = id

		// Initialisation of values
	}

	if len(xhtml_a_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_a_type instances pointers"
	}
	for idx, xhtml_a_type := range xhtml_a_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_a_type", idx, xhtml_a_type.Name)
		map_Xhtml_a_type_Identifiers[xhtml_a_type] = id

		// Initialisation of values
	}

	if len(xhtml_abbr_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_abbr_type instances pointers"
	}
	for idx, xhtml_abbr_type := range xhtml_abbr_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_abbr_type", idx, xhtml_abbr_type.Name)
		map_Xhtml_abbr_type_Identifiers[xhtml_abbr_type] = id

		// Initialisation of values
	}

	if len(xhtml_acronym_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_acronym_type instances pointers"
	}
	for idx, xhtml_acronym_type := range xhtml_acronym_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_acronym_type", idx, xhtml_acronym_type.Name)
		map_Xhtml_acronym_type_Identifiers[xhtml_acronym_type] = id

		// Initialisation of values
	}

	if len(xhtml_address_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_address_type instances pointers"
	}
	for idx, xhtml_address_type := range xhtml_address_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_address_type", idx, xhtml_address_type.Name)
		map_Xhtml_address_type_Identifiers[xhtml_address_type] = id

		// Initialisation of values
	}

	if len(xhtml_blockquote_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_blockquote_type instances pointers"
	}
	for idx, xhtml_blockquote_type := range xhtml_blockquote_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_blockquote_type", idx, xhtml_blockquote_type.Name)
		map_Xhtml_blockquote_type_Identifiers[xhtml_blockquote_type] = id

		// Initialisation of values
	}

	if len(xhtml_br_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_br_type instances pointers"
	}
	for idx, xhtml_br_type := range xhtml_br_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_br_type", idx, xhtml_br_type.Name)
		map_Xhtml_br_type_Identifiers[xhtml_br_type] = id

		// Initialisation of values
	}

	if len(xhtml_caption_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_caption_type instances pointers"
	}
	for idx, xhtml_caption_type := range xhtml_caption_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_caption_type", idx, xhtml_caption_type.Name)
		map_Xhtml_caption_type_Identifiers[xhtml_caption_type] = id

		// Initialisation of values
	}

	if len(xhtml_cite_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_cite_type instances pointers"
	}
	for idx, xhtml_cite_type := range xhtml_cite_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_cite_type", idx, xhtml_cite_type.Name)
		map_Xhtml_cite_type_Identifiers[xhtml_cite_type] = id

		// Initialisation of values
	}

	if len(xhtml_code_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_code_type instances pointers"
	}
	for idx, xhtml_code_type := range xhtml_code_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_code_type", idx, xhtml_code_type.Name)
		map_Xhtml_code_type_Identifiers[xhtml_code_type] = id

		// Initialisation of values
	}

	if len(xhtml_col_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_col_type instances pointers"
	}
	for idx, xhtml_col_type := range xhtml_col_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_col_type", idx, xhtml_col_type.Name)
		map_Xhtml_col_type_Identifiers[xhtml_col_type] = id

		// Initialisation of values
	}

	if len(xhtml_colgroup_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_colgroup_type instances pointers"
	}
	for idx, xhtml_colgroup_type := range xhtml_colgroup_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_colgroup_type", idx, xhtml_colgroup_type.Name)
		map_Xhtml_colgroup_type_Identifiers[xhtml_colgroup_type] = id

		// Initialisation of values
	}

	if len(xhtml_dd_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_dd_type instances pointers"
	}
	for idx, xhtml_dd_type := range xhtml_dd_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_dd_type", idx, xhtml_dd_type.Name)
		map_Xhtml_dd_type_Identifiers[xhtml_dd_type] = id

		// Initialisation of values
	}

	if len(xhtml_dfn_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_dfn_type instances pointers"
	}
	for idx, xhtml_dfn_type := range xhtml_dfn_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_dfn_type", idx, xhtml_dfn_type.Name)
		map_Xhtml_dfn_type_Identifiers[xhtml_dfn_type] = id

		// Initialisation of values
	}

	if len(xhtml_div_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_div_type instances pointers"
	}
	for idx, xhtml_div_type := range xhtml_div_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_div_type", idx, xhtml_div_type.Name)
		map_Xhtml_div_type_Identifiers[xhtml_div_type] = id

		// Initialisation of values
	}

	if len(xhtml_dl_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_dl_type instances pointers"
	}
	for idx, xhtml_dl_type := range xhtml_dl_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_dl_type", idx, xhtml_dl_type.Name)
		map_Xhtml_dl_type_Identifiers[xhtml_dl_type] = id

		// Initialisation of values
	}

	if len(xhtml_dt_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_dt_type instances pointers"
	}
	for idx, xhtml_dt_type := range xhtml_dt_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_dt_type", idx, xhtml_dt_type.Name)
		map_Xhtml_dt_type_Identifiers[xhtml_dt_type] = id

		// Initialisation of values
	}

	if len(xhtml_edit_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_edit_type instances pointers"
	}
	for idx, xhtml_edit_type := range xhtml_edit_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_edit_type", idx, xhtml_edit_type.Name)
		map_Xhtml_edit_type_Identifiers[xhtml_edit_type] = id

		// Initialisation of values
	}

	if len(xhtml_em_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_em_type instances pointers"
	}
	for idx, xhtml_em_type := range xhtml_em_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_em_type", idx, xhtml_em_type.Name)
		map_Xhtml_em_type_Identifiers[xhtml_em_type] = id

		// Initialisation of values
	}

	if len(xhtml_h1_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_h1_type instances pointers"
	}
	for idx, xhtml_h1_type := range xhtml_h1_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_h1_type", idx, xhtml_h1_type.Name)
		map_Xhtml_h1_type_Identifiers[xhtml_h1_type] = id

		// Initialisation of values
	}

	if len(xhtml_h2_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_h2_type instances pointers"
	}
	for idx, xhtml_h2_type := range xhtml_h2_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_h2_type", idx, xhtml_h2_type.Name)
		map_Xhtml_h2_type_Identifiers[xhtml_h2_type] = id

		// Initialisation of values
	}

	if len(xhtml_h3_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_h3_type instances pointers"
	}
	for idx, xhtml_h3_type := range xhtml_h3_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_h3_type", idx, xhtml_h3_type.Name)
		map_Xhtml_h3_type_Identifiers[xhtml_h3_type] = id

		// Initialisation of values
	}

	if len(xhtml_h4_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_h4_type instances pointers"
	}
	for idx, xhtml_h4_type := range xhtml_h4_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_h4_type", idx, xhtml_h4_type.Name)
		map_Xhtml_h4_type_Identifiers[xhtml_h4_type] = id

		// Initialisation of values
	}

	if len(xhtml_h5_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_h5_type instances pointers"
	}
	for idx, xhtml_h5_type := range xhtml_h5_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_h5_type", idx, xhtml_h5_type.Name)
		map_Xhtml_h5_type_Identifiers[xhtml_h5_type] = id

		// Initialisation of values
	}

	if len(xhtml_h6_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_h6_type instances pointers"
	}
	for idx, xhtml_h6_type := range xhtml_h6_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_h6_type", idx, xhtml_h6_type.Name)
		map_Xhtml_h6_type_Identifiers[xhtml_h6_type] = id

		// Initialisation of values
	}

	if len(xhtml_heading_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_heading_type instances pointers"
	}
	for idx, xhtml_heading_type := range xhtml_heading_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_heading_type", idx, xhtml_heading_type.Name)
		map_Xhtml_heading_type_Identifiers[xhtml_heading_type] = id

		// Initialisation of values
	}

	if len(xhtml_hr_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_hr_type instances pointers"
	}
	for idx, xhtml_hr_type := range xhtml_hr_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_hr_type", idx, xhtml_hr_type.Name)
		map_Xhtml_hr_type_Identifiers[xhtml_hr_type] = id

		// Initialisation of values
	}

	if len(xhtml_kbd_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_kbd_type instances pointers"
	}
	for idx, xhtml_kbd_type := range xhtml_kbd_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_kbd_type", idx, xhtml_kbd_type.Name)
		map_Xhtml_kbd_type_Identifiers[xhtml_kbd_type] = id

		// Initialisation of values
	}

	if len(xhtml_li_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_li_type instances pointers"
	}
	for idx, xhtml_li_type := range xhtml_li_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_li_type", idx, xhtml_li_type.Name)
		map_Xhtml_li_type_Identifiers[xhtml_li_type] = id

		// Initialisation of values
	}

	if len(xhtml_object_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_object_type instances pointers"
	}
	for idx, xhtml_object_type := range xhtml_object_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_object_type", idx, xhtml_object_type.Name)
		map_Xhtml_object_type_Identifiers[xhtml_object_type] = id

		// Initialisation of values
	}

	if len(xhtml_ol_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_ol_type instances pointers"
	}
	for idx, xhtml_ol_type := range xhtml_ol_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_ol_type", idx, xhtml_ol_type.Name)
		map_Xhtml_ol_type_Identifiers[xhtml_ol_type] = id

		// Initialisation of values
	}

	if len(xhtml_p_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_p_type instances pointers"
	}
	for idx, xhtml_p_type := range xhtml_p_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_p_type", idx, xhtml_p_type.Name)
		map_Xhtml_p_type_Identifiers[xhtml_p_type] = id

		// Initialisation of values
	}

	if len(xhtml_param_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_param_type instances pointers"
	}
	for idx, xhtml_param_type := range xhtml_param_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_param_type", idx, xhtml_param_type.Name)
		map_Xhtml_param_type_Identifiers[xhtml_param_type] = id

		// Initialisation of values
	}

	if len(xhtml_pre_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_pre_type instances pointers"
	}
	for idx, xhtml_pre_type := range xhtml_pre_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_pre_type", idx, xhtml_pre_type.Name)
		map_Xhtml_pre_type_Identifiers[xhtml_pre_type] = id

		// Initialisation of values
	}

	if len(xhtml_q_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_q_type instances pointers"
	}
	for idx, xhtml_q_type := range xhtml_q_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_q_type", idx, xhtml_q_type.Name)
		map_Xhtml_q_type_Identifiers[xhtml_q_type] = id

		// Initialisation of values
	}

	if len(xhtml_samp_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_samp_type instances pointers"
	}
	for idx, xhtml_samp_type := range xhtml_samp_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_samp_type", idx, xhtml_samp_type.Name)
		map_Xhtml_samp_type_Identifiers[xhtml_samp_type] = id

		// Initialisation of values
	}

	if len(xhtml_span_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_span_type instances pointers"
	}
	for idx, xhtml_span_type := range xhtml_span_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_span_type", idx, xhtml_span_type.Name)
		map_Xhtml_span_type_Identifiers[xhtml_span_type] = id

		// Initialisation of values
	}

	if len(xhtml_strong_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_strong_type instances pointers"
	}
	for idx, xhtml_strong_type := range xhtml_strong_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_strong_type", idx, xhtml_strong_type.Name)
		map_Xhtml_strong_type_Identifiers[xhtml_strong_type] = id

		// Initialisation of values
	}

	if len(xhtml_table_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_table_type instances pointers"
	}
	for idx, xhtml_table_type := range xhtml_table_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_table_type", idx, xhtml_table_type.Name)
		map_Xhtml_table_type_Identifiers[xhtml_table_type] = id

		// Initialisation of values
	}

	if len(xhtml_tbody_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_tbody_type instances pointers"
	}
	for idx, xhtml_tbody_type := range xhtml_tbody_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_tbody_type", idx, xhtml_tbody_type.Name)
		map_Xhtml_tbody_type_Identifiers[xhtml_tbody_type] = id

		// Initialisation of values
	}

	if len(xhtml_td_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_td_type instances pointers"
	}
	for idx, xhtml_td_type := range xhtml_td_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_td_type", idx, xhtml_td_type.Name)
		map_Xhtml_td_type_Identifiers[xhtml_td_type] = id

		// Initialisation of values
	}

	if len(xhtml_tfoot_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_tfoot_type instances pointers"
	}
	for idx, xhtml_tfoot_type := range xhtml_tfoot_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_tfoot_type", idx, xhtml_tfoot_type.Name)
		map_Xhtml_tfoot_type_Identifiers[xhtml_tfoot_type] = id

		// Initialisation of values
	}

	if len(xhtml_th_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_th_type instances pointers"
	}
	for idx, xhtml_th_type := range xhtml_th_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_th_type", idx, xhtml_th_type.Name)
		map_Xhtml_th_type_Identifiers[xhtml_th_type] = id

		// Initialisation of values
	}

	if len(xhtml_thead_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_thead_type instances pointers"
	}
	for idx, xhtml_thead_type := range xhtml_thead_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_thead_type", idx, xhtml_thead_type.Name)
		map_Xhtml_thead_type_Identifiers[xhtml_thead_type] = id

		// Initialisation of values
	}

	if len(xhtml_tr_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_tr_type instances pointers"
	}
	for idx, xhtml_tr_type := range xhtml_tr_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_tr_type", idx, xhtml_tr_type.Name)
		map_Xhtml_tr_type_Identifiers[xhtml_tr_type] = id

		// Initialisation of values
	}

	if len(xhtml_ul_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_ul_type instances pointers"
	}
	for idx, xhtml_ul_type := range xhtml_ul_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_ul_type", idx, xhtml_ul_type.Name)
		map_Xhtml_ul_type_Identifiers[xhtml_ul_type] = id

		// Initialisation of values
	}

	if len(xhtml_var_typeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xhtml_var_type instances pointers"
	}
	for idx, xhtml_var_type := range xhtml_var_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xhtml_var_type", idx, xhtml_var_type.Name)
		map_Xhtml_var_type_Identifiers[xhtml_var_type] = id

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

	fmt.Fprintln(file, res)
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

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
