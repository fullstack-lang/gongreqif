// generated code - do not edit
package models

import (
	"bufio"
	"embed"
	"errors"
	"go/ast"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError
var _ = dummy_strconv_import
var dummy_time_import time.Time
var _ = dummy_time_import

// swagger:ignore
type GONG__ExpressionType string

const (
	GONG__STRUCT_INSTANCE      GONG__ExpressionType = "STRUCT_INSTANCE"
	GONG__FIELD_OR_CONST_VALUE GONG__ExpressionType = "FIELD_OR_CONST_VALUE"
	GONG__FIELD_VALUE          GONG__ExpressionType = "FIELD_VALUE"
	GONG__ENUM_CAST_INT        GONG__ExpressionType = "ENUM_CAST_INT"
	GONG__ENUM_CAST_STRING     GONG__ExpressionType = "ENUM_CAST_STRING"
	GONG__IDENTIFIER_CONST     GONG__ExpressionType = "IDENTIFIER_CONST"
)

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(stage *Stage, pathToFile string) error {

	ReplaceOldDeclarationsInFile(pathToFile)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	// Read the file content using os.ReadFile
	content, err := os.ReadFile(fileOfInterest)
	if err != nil {
		return errors.New("Unable to read file " + err.Error())
	}

	// Assign the content to stage.contentWhenParsed
	stage.contentWhenParsed = string(content)

	fset := token.NewFileSet()
	// startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
// specified by pathToFile within the provided embed.FS directory and
// stages instances declared in the file using the provided Stage.
//
// Parameters:
//
//	stage:      The staging area to populate.
//	directory:  The embedded filesystem containing the file.
//	pathToFile: The path to the Go source file within the embedded filesystem.
//
// Returns:
//
//	An error if reading or parsing the file fails, or if ParseAstFileFromAst fails.
func ParseAstEmbeddedFile(stage *Stage, directory embed.FS, pathToFile string) error {

	// 1. Read the content from the embedded filesystem.
	//    We don't need filepath.Abs as embed.FS uses relative paths.
	//    We also skip ReplaceOldDeclarationsInFile as embedded files are read-only.
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		// Return a specific error if the file can't be read from the embed.FS
		return errors.New(stage.GetName() + "; Unable to read embedded file " + err.Error())
	}

	// 2. Create a FileSet to manage position information.
	fset := token.NewFileSet()

	// 3. Parse the file content.
	//    Instead of passing a filename for the OS to read, we pass the pathToFile
	//    as the identifier and the actual file content (fileContentBytes) as the source.
	//    parser.ParseComments is included to match the original function's behavior.
	//    The type *ast.File is returned by parser.ParseFile.
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		// Return a specific error if parsing fails
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	// 4. Call the common AST processing logic.
	//    Pass the parsed AST (*ast.File), the FileSet, and the stage.
	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *Stage, inFile *ast.File, fset *token.FileSet) error {
	// if there is a meta package import, it is the third import
	if len(inFile.Imports) > 3 {
		log.Fatalln("Too many imports in file", inFile.Name)
	}
	if len(inFile.Imports) == 3 {
		stage.MetaPackageImportAlias = inFile.Imports[2].Name.Name
		stage.MetaPackageImportPath = inFile.Imports[2].Path.Value
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "_")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.DeclStmt:
						if genDecl, ok := stmt.Decl.(*ast.GenDecl); ok && genDecl.Tok == token.CONST {
							for _, spec := range genDecl.Specs {
								if valueSpec, ok := spec.(*ast.ValueSpec); ok {
									for i, name := range valueSpec.Names {
										if i < len(valueSpec.Values) {
											if basicLit, ok := valueSpec.Values[i].(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
												// Remove quotes from string literal
												value := strings.Trim(basicLit.Value, `"`)

												switch name.Name {
												case "__commitId__":
													if parsedUint, err := strconv.ParseUint(value, 10, 64); err == nil {
														stage.commitId = uint(parsedUint)
														stage.commitIdWhenParsed = stage.commitId
													}
												}
											}
										}
									}
								}
							}
						}
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						// Create an ast.CommentMap from the ast.File's comments.
						// This helps keeping the association between comments
						// and AST nodes.
						// cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						var cmap ast.CommentMap
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName :=
							UnmarshallGongstructStaging(
								stage, &cmap, assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				case *ast.ValueSpec:
					ident := spec.Names[0]
					_ = ident
					if !strings.HasPrefix(ident.Name, "_") {
						continue
					}
					// declaration of a variable without initial value
					if len(spec.Values) == 0 {
						continue
					}
					switch compLit := spec.Values[0].(type) {
					case *ast.CompositeLit:
						var key string
						_ = key
						var value string
						_ = value
						for _, elt := range compLit.Elts {

							// each elt is an expression for struct or for field such as
							// for struct
							//
							//         "dummy.Dummy": &(dummy.Dummy{})
							//
							// or, for field
							//
							//          "dummy.Dummy.Name": (dummy.Dummy{}).Name,
							//
							// first node in the AST is a key value expression
							var ok bool
							var kve *ast.KeyValueExpr
							if kve, ok = elt.(*ast.KeyValueExpr); !ok {
								log.Fatal("Expression should be key value expression" +
									fset.Position(kve.Pos()).String())
							}

							switch bl := kve.Key.(type) {
							case *ast.BasicLit:
								key = bl.Value // "\"dumm.Dummy\"" is the value

								// one remove the ambracing double quotes
								key = strings.TrimPrefix(key, "\"")
								key = strings.TrimSuffix(key, "\"")
							}
							var expressionType GONG__ExpressionType = GONG__STRUCT_INSTANCE
							var docLink GONG__Identifier

							var fieldName string
							var ue *ast.UnaryExpr
							if ue, ok = kve.Value.(*ast.UnaryExpr); !ok {
								expressionType = GONG__FIELD_OR_CONST_VALUE
							}

							var callExpr *ast.CallExpr
							if callExpr, ok = kve.Value.(*ast.CallExpr); ok {

								var se *ast.SelectorExpr
								if se, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector expression" +
										fset.Position(callExpr.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}

								// check the arg type to select wether this is a int or a string enum
								var bl *ast.BasicLit
								if bl, ok = callExpr.Args[0].(*ast.BasicLit); ok {
									switch bl.Kind {
									case token.STRING:
										expressionType = GONG__ENUM_CAST_STRING
									case token.INT:
										expressionType = GONG__ENUM_CAST_INT
									}
								} else {
									log.Fatal("Expression should be a basic lit" +
										fset.Position(se.Pos()).String())
								}

								docLink.Ident = id.Name + "." + se.Sel.Name
								_ = callExpr
							}

							var se2 *ast.SelectorExpr
							switch expressionType {
							case GONG__FIELD_OR_CONST_VALUE:
								if se2, ok = kve.Value.(*ast.SelectorExpr); ok {

									var ident *ast.Ident
									if _, ok = se2.X.(*ast.ParenExpr); ok {
										expressionType = GONG__FIELD_VALUE
										fieldName = se2.Sel.Name
									} else if ident, ok = se2.X.(*ast.Ident); ok {
										expressionType = GONG__IDENTIFIER_CONST
										docLink.Ident = ident.Name + "." + se2.Sel.Name
									} else {
										log.Fatal("Expression should be a selector expression or an ident" +
											fset.Position(kve.Pos()).String())
									}
								} else {

								}
							}

							var pe *ast.ParenExpr
							switch expressionType {
							case GONG__STRUCT_INSTANCE:
								if pe, ok = ue.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							case GONG__FIELD_VALUE:
								if pe, ok = se2.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							}
							switch expressionType {
							case GONG__FIELD_VALUE, GONG__STRUCT_INSTANCE:
								// expect a Composite Litteral with no Element <type>{}
								var cl *ast.CompositeLit
								if cl, ok = pe.X.(*ast.CompositeLit); !ok {
									log.Fatal("Expression should be a composite lit" +
										fset.Position(pe.Pos()).String())
								}

								var se *ast.SelectorExpr
								if se, ok = cl.Type.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector" +
										fset.Position(cl.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}
								docLink.Ident = id.Name + "." + se.Sel.Name
							}

							switch expressionType {
							case GONG__FIELD_VALUE:
								docLink.Ident += "." + fieldName
							}

							// if map_DocLink_Identifier has the same ident, this means
							// that no renaming has occured since the last processing of the
							// file. But it is neccessary to keep it in memory for the
							// marshalling
							if docLink.Ident == key {
								// continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							docLink.Type = expressionType
							stage.Map_DocLink_Renaming[key] = docLink
						}
					}
				}
			}
		}

	}
	return nil
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_ALTERNATIVE_ID = make(map[string]*ALTERNATIVE_ID)
var __gong__map_ATTRIBUTE_DEFINITION_BOOLEAN = make(map[string]*ATTRIBUTE_DEFINITION_BOOLEAN)
var __gong__map_ATTRIBUTE_DEFINITION_DATE = make(map[string]*ATTRIBUTE_DEFINITION_DATE)
var __gong__map_ATTRIBUTE_DEFINITION_ENUMERATION = make(map[string]*ATTRIBUTE_DEFINITION_ENUMERATION)
var __gong__map_ATTRIBUTE_DEFINITION_INTEGER = make(map[string]*ATTRIBUTE_DEFINITION_INTEGER)
var __gong__map_ATTRIBUTE_DEFINITION_REAL = make(map[string]*ATTRIBUTE_DEFINITION_REAL)
var __gong__map_ATTRIBUTE_DEFINITION_STRING = make(map[string]*ATTRIBUTE_DEFINITION_STRING)
var __gong__map_ATTRIBUTE_DEFINITION_XHTML = make(map[string]*ATTRIBUTE_DEFINITION_XHTML)
var __gong__map_ATTRIBUTE_VALUE_BOOLEAN = make(map[string]*ATTRIBUTE_VALUE_BOOLEAN)
var __gong__map_ATTRIBUTE_VALUE_DATE = make(map[string]*ATTRIBUTE_VALUE_DATE)
var __gong__map_ATTRIBUTE_VALUE_ENUMERATION = make(map[string]*ATTRIBUTE_VALUE_ENUMERATION)
var __gong__map_ATTRIBUTE_VALUE_INTEGER = make(map[string]*ATTRIBUTE_VALUE_INTEGER)
var __gong__map_ATTRIBUTE_VALUE_REAL = make(map[string]*ATTRIBUTE_VALUE_REAL)
var __gong__map_ATTRIBUTE_VALUE_STRING = make(map[string]*ATTRIBUTE_VALUE_STRING)
var __gong__map_ATTRIBUTE_VALUE_XHTML = make(map[string]*ATTRIBUTE_VALUE_XHTML)
var __gong__map_A_ALTERNATIVE_ID = make(map[string]*A_ALTERNATIVE_ID)
var __gong__map_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF = make(map[string]*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
var __gong__map_A_ATTRIBUTE_DEFINITION_DATE_REF = make(map[string]*A_ATTRIBUTE_DEFINITION_DATE_REF)
var __gong__map_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF = make(map[string]*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
var __gong__map_A_ATTRIBUTE_DEFINITION_INTEGER_REF = make(map[string]*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
var __gong__map_A_ATTRIBUTE_DEFINITION_REAL_REF = make(map[string]*A_ATTRIBUTE_DEFINITION_REAL_REF)
var __gong__map_A_ATTRIBUTE_DEFINITION_STRING_REF = make(map[string]*A_ATTRIBUTE_DEFINITION_STRING_REF)
var __gong__map_A_ATTRIBUTE_DEFINITION_XHTML_REF = make(map[string]*A_ATTRIBUTE_DEFINITION_XHTML_REF)
var __gong__map_A_ATTRIBUTE_VALUE_BOOLEAN = make(map[string]*A_ATTRIBUTE_VALUE_BOOLEAN)
var __gong__map_A_ATTRIBUTE_VALUE_DATE = make(map[string]*A_ATTRIBUTE_VALUE_DATE)
var __gong__map_A_ATTRIBUTE_VALUE_ENUMERATION = make(map[string]*A_ATTRIBUTE_VALUE_ENUMERATION)
var __gong__map_A_ATTRIBUTE_VALUE_INTEGER = make(map[string]*A_ATTRIBUTE_VALUE_INTEGER)
var __gong__map_A_ATTRIBUTE_VALUE_REAL = make(map[string]*A_ATTRIBUTE_VALUE_REAL)
var __gong__map_A_ATTRIBUTE_VALUE_STRING = make(map[string]*A_ATTRIBUTE_VALUE_STRING)
var __gong__map_A_ATTRIBUTE_VALUE_XHTML = make(map[string]*A_ATTRIBUTE_VALUE_XHTML)
var __gong__map_A_ATTRIBUTE_VALUE_XHTML_1 = make(map[string]*A_ATTRIBUTE_VALUE_XHTML_1)
var __gong__map_A_CHILDREN = make(map[string]*A_CHILDREN)
var __gong__map_A_CORE_CONTENT = make(map[string]*A_CORE_CONTENT)
var __gong__map_A_DATATYPES = make(map[string]*A_DATATYPES)
var __gong__map_A_DATATYPE_DEFINITION_BOOLEAN_REF = make(map[string]*A_DATATYPE_DEFINITION_BOOLEAN_REF)
var __gong__map_A_DATATYPE_DEFINITION_DATE_REF = make(map[string]*A_DATATYPE_DEFINITION_DATE_REF)
var __gong__map_A_DATATYPE_DEFINITION_ENUMERATION_REF = make(map[string]*A_DATATYPE_DEFINITION_ENUMERATION_REF)
var __gong__map_A_DATATYPE_DEFINITION_INTEGER_REF = make(map[string]*A_DATATYPE_DEFINITION_INTEGER_REF)
var __gong__map_A_DATATYPE_DEFINITION_REAL_REF = make(map[string]*A_DATATYPE_DEFINITION_REAL_REF)
var __gong__map_A_DATATYPE_DEFINITION_STRING_REF = make(map[string]*A_DATATYPE_DEFINITION_STRING_REF)
var __gong__map_A_DATATYPE_DEFINITION_XHTML_REF = make(map[string]*A_DATATYPE_DEFINITION_XHTML_REF)
var __gong__map_A_EDITABLE_ATTS = make(map[string]*A_EDITABLE_ATTS)
var __gong__map_A_ENUM_VALUE_REF = make(map[string]*A_ENUM_VALUE_REF)
var __gong__map_A_OBJECT = make(map[string]*A_OBJECT)
var __gong__map_A_PROPERTIES = make(map[string]*A_PROPERTIES)
var __gong__map_A_RELATION_GROUP_TYPE_REF = make(map[string]*A_RELATION_GROUP_TYPE_REF)
var __gong__map_A_SOURCE_1 = make(map[string]*A_SOURCE_1)
var __gong__map_A_SOURCE_SPECIFICATION_1 = make(map[string]*A_SOURCE_SPECIFICATION_1)
var __gong__map_A_SPECIFICATIONS = make(map[string]*A_SPECIFICATIONS)
var __gong__map_A_SPECIFICATION_TYPE_REF = make(map[string]*A_SPECIFICATION_TYPE_REF)
var __gong__map_A_SPECIFIED_VALUES = make(map[string]*A_SPECIFIED_VALUES)
var __gong__map_A_SPEC_ATTRIBUTES = make(map[string]*A_SPEC_ATTRIBUTES)
var __gong__map_A_SPEC_OBJECTS = make(map[string]*A_SPEC_OBJECTS)
var __gong__map_A_SPEC_OBJECT_TYPE_REF = make(map[string]*A_SPEC_OBJECT_TYPE_REF)
var __gong__map_A_SPEC_RELATIONS = make(map[string]*A_SPEC_RELATIONS)
var __gong__map_A_SPEC_RELATION_GROUPS = make(map[string]*A_SPEC_RELATION_GROUPS)
var __gong__map_A_SPEC_RELATION_REF = make(map[string]*A_SPEC_RELATION_REF)
var __gong__map_A_SPEC_RELATION_TYPE_REF = make(map[string]*A_SPEC_RELATION_TYPE_REF)
var __gong__map_A_SPEC_TYPES = make(map[string]*A_SPEC_TYPES)
var __gong__map_A_THE_HEADER = make(map[string]*A_THE_HEADER)
var __gong__map_A_TOOL_EXTENSIONS = make(map[string]*A_TOOL_EXTENSIONS)
var __gong__map_DATATYPE_DEFINITION_BOOLEAN = make(map[string]*DATATYPE_DEFINITION_BOOLEAN)
var __gong__map_DATATYPE_DEFINITION_DATE = make(map[string]*DATATYPE_DEFINITION_DATE)
var __gong__map_DATATYPE_DEFINITION_ENUMERATION = make(map[string]*DATATYPE_DEFINITION_ENUMERATION)
var __gong__map_DATATYPE_DEFINITION_INTEGER = make(map[string]*DATATYPE_DEFINITION_INTEGER)
var __gong__map_DATATYPE_DEFINITION_REAL = make(map[string]*DATATYPE_DEFINITION_REAL)
var __gong__map_DATATYPE_DEFINITION_STRING = make(map[string]*DATATYPE_DEFINITION_STRING)
var __gong__map_DATATYPE_DEFINITION_XHTML = make(map[string]*DATATYPE_DEFINITION_XHTML)
var __gong__map_EMBEDDED_VALUE = make(map[string]*EMBEDDED_VALUE)
var __gong__map_ENUM_VALUE = make(map[string]*ENUM_VALUE)
var __gong__map_EmbeddedJpgImage = make(map[string]*EmbeddedJpgImage)
var __gong__map_EmbeddedPngImage = make(map[string]*EmbeddedPngImage)
var __gong__map_EmbeddedSvgImage = make(map[string]*EmbeddedSvgImage)
var __gong__map_Kill = make(map[string]*Kill)
var __gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry)
var __gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry = make(map[string]*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry)
var __gong__map_Map_SPECIFICATION_Nodes_expandedEntry = make(map[string]*Map_SPECIFICATION_Nodes_expandedEntry)
var __gong__map_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry = make(map[string]*Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry)
var __gong__map_Map_SPEC_OBJECT_TYPE_showIdentifierEntry = make(map[string]*Map_SPEC_OBJECT_TYPE_showIdentifierEntry)
var __gong__map_Map_SPEC_OBJECT_TYPE_showNameEntry = make(map[string]*Map_SPEC_OBJECT_TYPE_showNameEntry)
var __gong__map_RELATION_GROUP = make(map[string]*RELATION_GROUP)
var __gong__map_RELATION_GROUP_TYPE = make(map[string]*RELATION_GROUP_TYPE)
var __gong__map_REQ_IF = make(map[string]*REQ_IF)
var __gong__map_REQ_IF_CONTENT = make(map[string]*REQ_IF_CONTENT)
var __gong__map_REQ_IF_HEADER = make(map[string]*REQ_IF_HEADER)
var __gong__map_REQ_IF_TOOL_EXTENSION = make(map[string]*REQ_IF_TOOL_EXTENSION)
var __gong__map_RenderingConfiguration = make(map[string]*RenderingConfiguration)
var __gong__map_SPECIFICATION = make(map[string]*SPECIFICATION)
var __gong__map_SPECIFICATION_TYPE = make(map[string]*SPECIFICATION_TYPE)
var __gong__map_SPEC_HIERARCHY = make(map[string]*SPEC_HIERARCHY)
var __gong__map_SPEC_OBJECT = make(map[string]*SPEC_OBJECT)
var __gong__map_SPEC_OBJECT_TYPE = make(map[string]*SPEC_OBJECT_TYPE)
var __gong__map_SPEC_RELATION = make(map[string]*SPEC_RELATION)
var __gong__map_SPEC_RELATION_TYPE = make(map[string]*SPEC_RELATION_TYPE)
var __gong__map_StaticWebSite = make(map[string]*StaticWebSite)
var __gong__map_StaticWebSiteChapter = make(map[string]*StaticWebSiteChapter)
var __gong__map_StaticWebSiteGeneratedImage = make(map[string]*StaticWebSiteGeneratedImage)
var __gong__map_StaticWebSiteImage = make(map[string]*StaticWebSiteImage)
var __gong__map_StaticWebSiteParagraph = make(map[string]*StaticWebSiteParagraph)
var __gong__map_XHTML_CONTENT = make(map[string]*XHTML_CONTENT)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	return name, true
}

func lookupSym(recv, name string) bool {
	return recv == ""
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *Stage, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {

	// used for debug purposes
	astCoordinate := "\tAssignStmt: "

	//
	// First parse all comment groups in the assignement
	// if a comment "//gong:ident [DocLink]" is met and is followed by a string assignement.
	// modify the following AST assignement to assigns the DocLink text to the string value
	//

	// get the comment group of the assigStmt
	commentGroups := (*cmap)[assignStmt]
	// get the the prefix
	var hasGongIdentDirective bool
	var commentText string
	var docLinkText string
	for _, commentGroup := range commentGroups {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "//gong:ident") {
				hasGongIdentDirective = true
				commentText = comment.Text
			}
		}
	}
	if hasGongIdentDirective {
		// parser configured to find doclinks
		var docLinkFinder comment.Parser
		docLinkFinder.LookupPackage = lookupPackage
		docLinkFinder.LookupSym = lookupSym
		doc := docLinkFinder.Parse(commentText)

		for _, block := range doc.Content {
			switch paragraph := block.(type) {
			case *comment.Paragraph:
				_ = paragraph
				for _, text := range paragraph.Text {
					switch docLink := text.(type) {
					case *comment.DocLink:
						if docLink.Recv == "" {
							docLinkText = docLink.ImportPath + "." + docLink.Name
						} else {
							docLinkText = docLink.ImportPath + "." + docLink.Recv + "." + docLink.Name
						}

						// we check wether the doc link has been renamed
						// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
						if renamed, ok := (stage.Map_DocLink_Renaming)[docLinkText]; ok {
							docLinkText = renamed.Ident
						}
					}
				}
			}
		}
	}

	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			var basicLit *ast.BasicLit

			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "ALTERNATIVE_ID":
										instanceALTERNATIVE_ID := new(ALTERNATIVE_ID)
										instanceALTERNATIVE_ID.Name = instanceName
										instanceALTERNATIVE_ID.Stage(stage)
										instance = any(instanceALTERNATIVE_ID)
										__gong__map_ALTERNATIVE_ID[identifier] = instanceALTERNATIVE_ID
									case "ATTRIBUTE_DEFINITION_BOOLEAN":
										instanceATTRIBUTE_DEFINITION_BOOLEAN := new(ATTRIBUTE_DEFINITION_BOOLEAN)
										instanceATTRIBUTE_DEFINITION_BOOLEAN.Name = instanceName
										instanceATTRIBUTE_DEFINITION_BOOLEAN.Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_BOOLEAN)
										__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier] = instanceATTRIBUTE_DEFINITION_BOOLEAN
									case "ATTRIBUTE_DEFINITION_DATE":
										instanceATTRIBUTE_DEFINITION_DATE := new(ATTRIBUTE_DEFINITION_DATE)
										instanceATTRIBUTE_DEFINITION_DATE.Name = instanceName
										instanceATTRIBUTE_DEFINITION_DATE.Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_DATE)
										__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier] = instanceATTRIBUTE_DEFINITION_DATE
									case "ATTRIBUTE_DEFINITION_ENUMERATION":
										instanceATTRIBUTE_DEFINITION_ENUMERATION := new(ATTRIBUTE_DEFINITION_ENUMERATION)
										instanceATTRIBUTE_DEFINITION_ENUMERATION.Name = instanceName
										instanceATTRIBUTE_DEFINITION_ENUMERATION.Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_ENUMERATION)
										__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier] = instanceATTRIBUTE_DEFINITION_ENUMERATION
									case "ATTRIBUTE_DEFINITION_INTEGER":
										instanceATTRIBUTE_DEFINITION_INTEGER := new(ATTRIBUTE_DEFINITION_INTEGER)
										instanceATTRIBUTE_DEFINITION_INTEGER.Name = instanceName
										instanceATTRIBUTE_DEFINITION_INTEGER.Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_INTEGER)
										__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier] = instanceATTRIBUTE_DEFINITION_INTEGER
									case "ATTRIBUTE_DEFINITION_REAL":
										instanceATTRIBUTE_DEFINITION_REAL := new(ATTRIBUTE_DEFINITION_REAL)
										instanceATTRIBUTE_DEFINITION_REAL.Name = instanceName
										instanceATTRIBUTE_DEFINITION_REAL.Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_REAL)
										__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier] = instanceATTRIBUTE_DEFINITION_REAL
									case "ATTRIBUTE_DEFINITION_STRING":
										instanceATTRIBUTE_DEFINITION_STRING := new(ATTRIBUTE_DEFINITION_STRING)
										instanceATTRIBUTE_DEFINITION_STRING.Name = instanceName
										instanceATTRIBUTE_DEFINITION_STRING.Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_STRING)
										__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier] = instanceATTRIBUTE_DEFINITION_STRING
									case "ATTRIBUTE_DEFINITION_XHTML":
										instanceATTRIBUTE_DEFINITION_XHTML := new(ATTRIBUTE_DEFINITION_XHTML)
										instanceATTRIBUTE_DEFINITION_XHTML.Name = instanceName
										instanceATTRIBUTE_DEFINITION_XHTML.Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_XHTML)
										__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier] = instanceATTRIBUTE_DEFINITION_XHTML
									case "ATTRIBUTE_VALUE_BOOLEAN":
										instanceATTRIBUTE_VALUE_BOOLEAN := new(ATTRIBUTE_VALUE_BOOLEAN)
										instanceATTRIBUTE_VALUE_BOOLEAN.Name = instanceName
										instanceATTRIBUTE_VALUE_BOOLEAN.Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_BOOLEAN)
										__gong__map_ATTRIBUTE_VALUE_BOOLEAN[identifier] = instanceATTRIBUTE_VALUE_BOOLEAN
									case "ATTRIBUTE_VALUE_DATE":
										instanceATTRIBUTE_VALUE_DATE := new(ATTRIBUTE_VALUE_DATE)
										instanceATTRIBUTE_VALUE_DATE.Name = instanceName
										instanceATTRIBUTE_VALUE_DATE.Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_DATE)
										__gong__map_ATTRIBUTE_VALUE_DATE[identifier] = instanceATTRIBUTE_VALUE_DATE
									case "ATTRIBUTE_VALUE_ENUMERATION":
										instanceATTRIBUTE_VALUE_ENUMERATION := new(ATTRIBUTE_VALUE_ENUMERATION)
										instanceATTRIBUTE_VALUE_ENUMERATION.Name = instanceName
										instanceATTRIBUTE_VALUE_ENUMERATION.Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_ENUMERATION)
										__gong__map_ATTRIBUTE_VALUE_ENUMERATION[identifier] = instanceATTRIBUTE_VALUE_ENUMERATION
									case "ATTRIBUTE_VALUE_INTEGER":
										instanceATTRIBUTE_VALUE_INTEGER := new(ATTRIBUTE_VALUE_INTEGER)
										instanceATTRIBUTE_VALUE_INTEGER.Name = instanceName
										instanceATTRIBUTE_VALUE_INTEGER.Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_INTEGER)
										__gong__map_ATTRIBUTE_VALUE_INTEGER[identifier] = instanceATTRIBUTE_VALUE_INTEGER
									case "ATTRIBUTE_VALUE_REAL":
										instanceATTRIBUTE_VALUE_REAL := new(ATTRIBUTE_VALUE_REAL)
										instanceATTRIBUTE_VALUE_REAL.Name = instanceName
										instanceATTRIBUTE_VALUE_REAL.Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_REAL)
										__gong__map_ATTRIBUTE_VALUE_REAL[identifier] = instanceATTRIBUTE_VALUE_REAL
									case "ATTRIBUTE_VALUE_STRING":
										instanceATTRIBUTE_VALUE_STRING := new(ATTRIBUTE_VALUE_STRING)
										instanceATTRIBUTE_VALUE_STRING.Name = instanceName
										instanceATTRIBUTE_VALUE_STRING.Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_STRING)
										__gong__map_ATTRIBUTE_VALUE_STRING[identifier] = instanceATTRIBUTE_VALUE_STRING
									case "ATTRIBUTE_VALUE_XHTML":
										instanceATTRIBUTE_VALUE_XHTML := new(ATTRIBUTE_VALUE_XHTML)
										instanceATTRIBUTE_VALUE_XHTML.Name = instanceName
										instanceATTRIBUTE_VALUE_XHTML.Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_XHTML)
										__gong__map_ATTRIBUTE_VALUE_XHTML[identifier] = instanceATTRIBUTE_VALUE_XHTML
									case "A_ALTERNATIVE_ID":
										instanceA_ALTERNATIVE_ID := new(A_ALTERNATIVE_ID)
										instanceA_ALTERNATIVE_ID.Name = instanceName
										instanceA_ALTERNATIVE_ID.Stage(stage)
										instance = any(instanceA_ALTERNATIVE_ID)
										__gong__map_A_ALTERNATIVE_ID[identifier] = instanceA_ALTERNATIVE_ID
									case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
										instanceA_ATTRIBUTE_DEFINITION_BOOLEAN_REF := new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
										instanceA_ATTRIBUTE_DEFINITION_BOOLEAN_REF.Name = instanceName
										instanceA_ATTRIBUTE_DEFINITION_BOOLEAN_REF.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
										__gong__map_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF[identifier] = instanceA_ATTRIBUTE_DEFINITION_BOOLEAN_REF
									case "A_ATTRIBUTE_DEFINITION_DATE_REF":
										instanceA_ATTRIBUTE_DEFINITION_DATE_REF := new(A_ATTRIBUTE_DEFINITION_DATE_REF)
										instanceA_ATTRIBUTE_DEFINITION_DATE_REF.Name = instanceName
										instanceA_ATTRIBUTE_DEFINITION_DATE_REF.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_DEFINITION_DATE_REF)
										__gong__map_A_ATTRIBUTE_DEFINITION_DATE_REF[identifier] = instanceA_ATTRIBUTE_DEFINITION_DATE_REF
									case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
										instanceA_ATTRIBUTE_DEFINITION_ENUMERATION_REF := new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
										instanceA_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Name = instanceName
										instanceA_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
										__gong__map_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF[identifier] = instanceA_ATTRIBUTE_DEFINITION_ENUMERATION_REF
									case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
										instanceA_ATTRIBUTE_DEFINITION_INTEGER_REF := new(A_ATTRIBUTE_DEFINITION_INTEGER_REF)
										instanceA_ATTRIBUTE_DEFINITION_INTEGER_REF.Name = instanceName
										instanceA_ATTRIBUTE_DEFINITION_INTEGER_REF.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_DEFINITION_INTEGER_REF)
										__gong__map_A_ATTRIBUTE_DEFINITION_INTEGER_REF[identifier] = instanceA_ATTRIBUTE_DEFINITION_INTEGER_REF
									case "A_ATTRIBUTE_DEFINITION_REAL_REF":
										instanceA_ATTRIBUTE_DEFINITION_REAL_REF := new(A_ATTRIBUTE_DEFINITION_REAL_REF)
										instanceA_ATTRIBUTE_DEFINITION_REAL_REF.Name = instanceName
										instanceA_ATTRIBUTE_DEFINITION_REAL_REF.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_DEFINITION_REAL_REF)
										__gong__map_A_ATTRIBUTE_DEFINITION_REAL_REF[identifier] = instanceA_ATTRIBUTE_DEFINITION_REAL_REF
									case "A_ATTRIBUTE_DEFINITION_STRING_REF":
										instanceA_ATTRIBUTE_DEFINITION_STRING_REF := new(A_ATTRIBUTE_DEFINITION_STRING_REF)
										instanceA_ATTRIBUTE_DEFINITION_STRING_REF.Name = instanceName
										instanceA_ATTRIBUTE_DEFINITION_STRING_REF.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_DEFINITION_STRING_REF)
										__gong__map_A_ATTRIBUTE_DEFINITION_STRING_REF[identifier] = instanceA_ATTRIBUTE_DEFINITION_STRING_REF
									case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
										instanceA_ATTRIBUTE_DEFINITION_XHTML_REF := new(A_ATTRIBUTE_DEFINITION_XHTML_REF)
										instanceA_ATTRIBUTE_DEFINITION_XHTML_REF.Name = instanceName
										instanceA_ATTRIBUTE_DEFINITION_XHTML_REF.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_DEFINITION_XHTML_REF)
										__gong__map_A_ATTRIBUTE_DEFINITION_XHTML_REF[identifier] = instanceA_ATTRIBUTE_DEFINITION_XHTML_REF
									case "A_ATTRIBUTE_VALUE_BOOLEAN":
										instanceA_ATTRIBUTE_VALUE_BOOLEAN := new(A_ATTRIBUTE_VALUE_BOOLEAN)
										instanceA_ATTRIBUTE_VALUE_BOOLEAN.Name = instanceName
										instanceA_ATTRIBUTE_VALUE_BOOLEAN.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_VALUE_BOOLEAN)
										__gong__map_A_ATTRIBUTE_VALUE_BOOLEAN[identifier] = instanceA_ATTRIBUTE_VALUE_BOOLEAN
									case "A_ATTRIBUTE_VALUE_DATE":
										instanceA_ATTRIBUTE_VALUE_DATE := new(A_ATTRIBUTE_VALUE_DATE)
										instanceA_ATTRIBUTE_VALUE_DATE.Name = instanceName
										instanceA_ATTRIBUTE_VALUE_DATE.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_VALUE_DATE)
										__gong__map_A_ATTRIBUTE_VALUE_DATE[identifier] = instanceA_ATTRIBUTE_VALUE_DATE
									case "A_ATTRIBUTE_VALUE_ENUMERATION":
										instanceA_ATTRIBUTE_VALUE_ENUMERATION := new(A_ATTRIBUTE_VALUE_ENUMERATION)
										instanceA_ATTRIBUTE_VALUE_ENUMERATION.Name = instanceName
										instanceA_ATTRIBUTE_VALUE_ENUMERATION.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_VALUE_ENUMERATION)
										__gong__map_A_ATTRIBUTE_VALUE_ENUMERATION[identifier] = instanceA_ATTRIBUTE_VALUE_ENUMERATION
									case "A_ATTRIBUTE_VALUE_INTEGER":
										instanceA_ATTRIBUTE_VALUE_INTEGER := new(A_ATTRIBUTE_VALUE_INTEGER)
										instanceA_ATTRIBUTE_VALUE_INTEGER.Name = instanceName
										instanceA_ATTRIBUTE_VALUE_INTEGER.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_VALUE_INTEGER)
										__gong__map_A_ATTRIBUTE_VALUE_INTEGER[identifier] = instanceA_ATTRIBUTE_VALUE_INTEGER
									case "A_ATTRIBUTE_VALUE_REAL":
										instanceA_ATTRIBUTE_VALUE_REAL := new(A_ATTRIBUTE_VALUE_REAL)
										instanceA_ATTRIBUTE_VALUE_REAL.Name = instanceName
										instanceA_ATTRIBUTE_VALUE_REAL.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_VALUE_REAL)
										__gong__map_A_ATTRIBUTE_VALUE_REAL[identifier] = instanceA_ATTRIBUTE_VALUE_REAL
									case "A_ATTRIBUTE_VALUE_STRING":
										instanceA_ATTRIBUTE_VALUE_STRING := new(A_ATTRIBUTE_VALUE_STRING)
										instanceA_ATTRIBUTE_VALUE_STRING.Name = instanceName
										instanceA_ATTRIBUTE_VALUE_STRING.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_VALUE_STRING)
										__gong__map_A_ATTRIBUTE_VALUE_STRING[identifier] = instanceA_ATTRIBUTE_VALUE_STRING
									case "A_ATTRIBUTE_VALUE_XHTML":
										instanceA_ATTRIBUTE_VALUE_XHTML := new(A_ATTRIBUTE_VALUE_XHTML)
										instanceA_ATTRIBUTE_VALUE_XHTML.Name = instanceName
										instanceA_ATTRIBUTE_VALUE_XHTML.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_VALUE_XHTML)
										__gong__map_A_ATTRIBUTE_VALUE_XHTML[identifier] = instanceA_ATTRIBUTE_VALUE_XHTML
									case "A_ATTRIBUTE_VALUE_XHTML_1":
										instanceA_ATTRIBUTE_VALUE_XHTML_1 := new(A_ATTRIBUTE_VALUE_XHTML_1)
										instanceA_ATTRIBUTE_VALUE_XHTML_1.Name = instanceName
										instanceA_ATTRIBUTE_VALUE_XHTML_1.Stage(stage)
										instance = any(instanceA_ATTRIBUTE_VALUE_XHTML_1)
										__gong__map_A_ATTRIBUTE_VALUE_XHTML_1[identifier] = instanceA_ATTRIBUTE_VALUE_XHTML_1
									case "A_CHILDREN":
										instanceA_CHILDREN := new(A_CHILDREN)
										instanceA_CHILDREN.Name = instanceName
										instanceA_CHILDREN.Stage(stage)
										instance = any(instanceA_CHILDREN)
										__gong__map_A_CHILDREN[identifier] = instanceA_CHILDREN
									case "A_CORE_CONTENT":
										instanceA_CORE_CONTENT := new(A_CORE_CONTENT)
										instanceA_CORE_CONTENT.Name = instanceName
										instanceA_CORE_CONTENT.Stage(stage)
										instance = any(instanceA_CORE_CONTENT)
										__gong__map_A_CORE_CONTENT[identifier] = instanceA_CORE_CONTENT
									case "A_DATATYPES":
										instanceA_DATATYPES := new(A_DATATYPES)
										instanceA_DATATYPES.Name = instanceName
										instanceA_DATATYPES.Stage(stage)
										instance = any(instanceA_DATATYPES)
										__gong__map_A_DATATYPES[identifier] = instanceA_DATATYPES
									case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
										instanceA_DATATYPE_DEFINITION_BOOLEAN_REF := new(A_DATATYPE_DEFINITION_BOOLEAN_REF)
										instanceA_DATATYPE_DEFINITION_BOOLEAN_REF.Name = instanceName
										instanceA_DATATYPE_DEFINITION_BOOLEAN_REF.Stage(stage)
										instance = any(instanceA_DATATYPE_DEFINITION_BOOLEAN_REF)
										__gong__map_A_DATATYPE_DEFINITION_BOOLEAN_REF[identifier] = instanceA_DATATYPE_DEFINITION_BOOLEAN_REF
									case "A_DATATYPE_DEFINITION_DATE_REF":
										instanceA_DATATYPE_DEFINITION_DATE_REF := new(A_DATATYPE_DEFINITION_DATE_REF)
										instanceA_DATATYPE_DEFINITION_DATE_REF.Name = instanceName
										instanceA_DATATYPE_DEFINITION_DATE_REF.Stage(stage)
										instance = any(instanceA_DATATYPE_DEFINITION_DATE_REF)
										__gong__map_A_DATATYPE_DEFINITION_DATE_REF[identifier] = instanceA_DATATYPE_DEFINITION_DATE_REF
									case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
										instanceA_DATATYPE_DEFINITION_ENUMERATION_REF := new(A_DATATYPE_DEFINITION_ENUMERATION_REF)
										instanceA_DATATYPE_DEFINITION_ENUMERATION_REF.Name = instanceName
										instanceA_DATATYPE_DEFINITION_ENUMERATION_REF.Stage(stage)
										instance = any(instanceA_DATATYPE_DEFINITION_ENUMERATION_REF)
										__gong__map_A_DATATYPE_DEFINITION_ENUMERATION_REF[identifier] = instanceA_DATATYPE_DEFINITION_ENUMERATION_REF
									case "A_DATATYPE_DEFINITION_INTEGER_REF":
										instanceA_DATATYPE_DEFINITION_INTEGER_REF := new(A_DATATYPE_DEFINITION_INTEGER_REF)
										instanceA_DATATYPE_DEFINITION_INTEGER_REF.Name = instanceName
										instanceA_DATATYPE_DEFINITION_INTEGER_REF.Stage(stage)
										instance = any(instanceA_DATATYPE_DEFINITION_INTEGER_REF)
										__gong__map_A_DATATYPE_DEFINITION_INTEGER_REF[identifier] = instanceA_DATATYPE_DEFINITION_INTEGER_REF
									case "A_DATATYPE_DEFINITION_REAL_REF":
										instanceA_DATATYPE_DEFINITION_REAL_REF := new(A_DATATYPE_DEFINITION_REAL_REF)
										instanceA_DATATYPE_DEFINITION_REAL_REF.Name = instanceName
										instanceA_DATATYPE_DEFINITION_REAL_REF.Stage(stage)
										instance = any(instanceA_DATATYPE_DEFINITION_REAL_REF)
										__gong__map_A_DATATYPE_DEFINITION_REAL_REF[identifier] = instanceA_DATATYPE_DEFINITION_REAL_REF
									case "A_DATATYPE_DEFINITION_STRING_REF":
										instanceA_DATATYPE_DEFINITION_STRING_REF := new(A_DATATYPE_DEFINITION_STRING_REF)
										instanceA_DATATYPE_DEFINITION_STRING_REF.Name = instanceName
										instanceA_DATATYPE_DEFINITION_STRING_REF.Stage(stage)
										instance = any(instanceA_DATATYPE_DEFINITION_STRING_REF)
										__gong__map_A_DATATYPE_DEFINITION_STRING_REF[identifier] = instanceA_DATATYPE_DEFINITION_STRING_REF
									case "A_DATATYPE_DEFINITION_XHTML_REF":
										instanceA_DATATYPE_DEFINITION_XHTML_REF := new(A_DATATYPE_DEFINITION_XHTML_REF)
										instanceA_DATATYPE_DEFINITION_XHTML_REF.Name = instanceName
										instanceA_DATATYPE_DEFINITION_XHTML_REF.Stage(stage)
										instance = any(instanceA_DATATYPE_DEFINITION_XHTML_REF)
										__gong__map_A_DATATYPE_DEFINITION_XHTML_REF[identifier] = instanceA_DATATYPE_DEFINITION_XHTML_REF
									case "A_EDITABLE_ATTS":
										instanceA_EDITABLE_ATTS := new(A_EDITABLE_ATTS)
										instanceA_EDITABLE_ATTS.Name = instanceName
										instanceA_EDITABLE_ATTS.Stage(stage)
										instance = any(instanceA_EDITABLE_ATTS)
										__gong__map_A_EDITABLE_ATTS[identifier] = instanceA_EDITABLE_ATTS
									case "A_ENUM_VALUE_REF":
										instanceA_ENUM_VALUE_REF := new(A_ENUM_VALUE_REF)
										instanceA_ENUM_VALUE_REF.Name = instanceName
										instanceA_ENUM_VALUE_REF.Stage(stage)
										instance = any(instanceA_ENUM_VALUE_REF)
										__gong__map_A_ENUM_VALUE_REF[identifier] = instanceA_ENUM_VALUE_REF
									case "A_OBJECT":
										instanceA_OBJECT := new(A_OBJECT)
										instanceA_OBJECT.Name = instanceName
										instanceA_OBJECT.Stage(stage)
										instance = any(instanceA_OBJECT)
										__gong__map_A_OBJECT[identifier] = instanceA_OBJECT
									case "A_PROPERTIES":
										instanceA_PROPERTIES := new(A_PROPERTIES)
										instanceA_PROPERTIES.Name = instanceName
										instanceA_PROPERTIES.Stage(stage)
										instance = any(instanceA_PROPERTIES)
										__gong__map_A_PROPERTIES[identifier] = instanceA_PROPERTIES
									case "A_RELATION_GROUP_TYPE_REF":
										instanceA_RELATION_GROUP_TYPE_REF := new(A_RELATION_GROUP_TYPE_REF)
										instanceA_RELATION_GROUP_TYPE_REF.Name = instanceName
										instanceA_RELATION_GROUP_TYPE_REF.Stage(stage)
										instance = any(instanceA_RELATION_GROUP_TYPE_REF)
										__gong__map_A_RELATION_GROUP_TYPE_REF[identifier] = instanceA_RELATION_GROUP_TYPE_REF
									case "A_SOURCE_1":
										instanceA_SOURCE_1 := new(A_SOURCE_1)
										instanceA_SOURCE_1.Name = instanceName
										instanceA_SOURCE_1.Stage(stage)
										instance = any(instanceA_SOURCE_1)
										__gong__map_A_SOURCE_1[identifier] = instanceA_SOURCE_1
									case "A_SOURCE_SPECIFICATION_1":
										instanceA_SOURCE_SPECIFICATION_1 := new(A_SOURCE_SPECIFICATION_1)
										instanceA_SOURCE_SPECIFICATION_1.Name = instanceName
										instanceA_SOURCE_SPECIFICATION_1.Stage(stage)
										instance = any(instanceA_SOURCE_SPECIFICATION_1)
										__gong__map_A_SOURCE_SPECIFICATION_1[identifier] = instanceA_SOURCE_SPECIFICATION_1
									case "A_SPECIFICATIONS":
										instanceA_SPECIFICATIONS := new(A_SPECIFICATIONS)
										instanceA_SPECIFICATIONS.Name = instanceName
										instanceA_SPECIFICATIONS.Stage(stage)
										instance = any(instanceA_SPECIFICATIONS)
										__gong__map_A_SPECIFICATIONS[identifier] = instanceA_SPECIFICATIONS
									case "A_SPECIFICATION_TYPE_REF":
										instanceA_SPECIFICATION_TYPE_REF := new(A_SPECIFICATION_TYPE_REF)
										instanceA_SPECIFICATION_TYPE_REF.Name = instanceName
										instanceA_SPECIFICATION_TYPE_REF.Stage(stage)
										instance = any(instanceA_SPECIFICATION_TYPE_REF)
										__gong__map_A_SPECIFICATION_TYPE_REF[identifier] = instanceA_SPECIFICATION_TYPE_REF
									case "A_SPECIFIED_VALUES":
										instanceA_SPECIFIED_VALUES := new(A_SPECIFIED_VALUES)
										instanceA_SPECIFIED_VALUES.Name = instanceName
										instanceA_SPECIFIED_VALUES.Stage(stage)
										instance = any(instanceA_SPECIFIED_VALUES)
										__gong__map_A_SPECIFIED_VALUES[identifier] = instanceA_SPECIFIED_VALUES
									case "A_SPEC_ATTRIBUTES":
										instanceA_SPEC_ATTRIBUTES := new(A_SPEC_ATTRIBUTES)
										instanceA_SPEC_ATTRIBUTES.Name = instanceName
										instanceA_SPEC_ATTRIBUTES.Stage(stage)
										instance = any(instanceA_SPEC_ATTRIBUTES)
										__gong__map_A_SPEC_ATTRIBUTES[identifier] = instanceA_SPEC_ATTRIBUTES
									case "A_SPEC_OBJECTS":
										instanceA_SPEC_OBJECTS := new(A_SPEC_OBJECTS)
										instanceA_SPEC_OBJECTS.Name = instanceName
										instanceA_SPEC_OBJECTS.Stage(stage)
										instance = any(instanceA_SPEC_OBJECTS)
										__gong__map_A_SPEC_OBJECTS[identifier] = instanceA_SPEC_OBJECTS
									case "A_SPEC_OBJECT_TYPE_REF":
										instanceA_SPEC_OBJECT_TYPE_REF := new(A_SPEC_OBJECT_TYPE_REF)
										instanceA_SPEC_OBJECT_TYPE_REF.Name = instanceName
										instanceA_SPEC_OBJECT_TYPE_REF.Stage(stage)
										instance = any(instanceA_SPEC_OBJECT_TYPE_REF)
										__gong__map_A_SPEC_OBJECT_TYPE_REF[identifier] = instanceA_SPEC_OBJECT_TYPE_REF
									case "A_SPEC_RELATIONS":
										instanceA_SPEC_RELATIONS := new(A_SPEC_RELATIONS)
										instanceA_SPEC_RELATIONS.Name = instanceName
										instanceA_SPEC_RELATIONS.Stage(stage)
										instance = any(instanceA_SPEC_RELATIONS)
										__gong__map_A_SPEC_RELATIONS[identifier] = instanceA_SPEC_RELATIONS
									case "A_SPEC_RELATION_GROUPS":
										instanceA_SPEC_RELATION_GROUPS := new(A_SPEC_RELATION_GROUPS)
										instanceA_SPEC_RELATION_GROUPS.Name = instanceName
										instanceA_SPEC_RELATION_GROUPS.Stage(stage)
										instance = any(instanceA_SPEC_RELATION_GROUPS)
										__gong__map_A_SPEC_RELATION_GROUPS[identifier] = instanceA_SPEC_RELATION_GROUPS
									case "A_SPEC_RELATION_REF":
										instanceA_SPEC_RELATION_REF := new(A_SPEC_RELATION_REF)
										instanceA_SPEC_RELATION_REF.Name = instanceName
										instanceA_SPEC_RELATION_REF.Stage(stage)
										instance = any(instanceA_SPEC_RELATION_REF)
										__gong__map_A_SPEC_RELATION_REF[identifier] = instanceA_SPEC_RELATION_REF
									case "A_SPEC_RELATION_TYPE_REF":
										instanceA_SPEC_RELATION_TYPE_REF := new(A_SPEC_RELATION_TYPE_REF)
										instanceA_SPEC_RELATION_TYPE_REF.Name = instanceName
										instanceA_SPEC_RELATION_TYPE_REF.Stage(stage)
										instance = any(instanceA_SPEC_RELATION_TYPE_REF)
										__gong__map_A_SPEC_RELATION_TYPE_REF[identifier] = instanceA_SPEC_RELATION_TYPE_REF
									case "A_SPEC_TYPES":
										instanceA_SPEC_TYPES := new(A_SPEC_TYPES)
										instanceA_SPEC_TYPES.Name = instanceName
										instanceA_SPEC_TYPES.Stage(stage)
										instance = any(instanceA_SPEC_TYPES)
										__gong__map_A_SPEC_TYPES[identifier] = instanceA_SPEC_TYPES
									case "A_THE_HEADER":
										instanceA_THE_HEADER := new(A_THE_HEADER)
										instanceA_THE_HEADER.Name = instanceName
										instanceA_THE_HEADER.Stage(stage)
										instance = any(instanceA_THE_HEADER)
										__gong__map_A_THE_HEADER[identifier] = instanceA_THE_HEADER
									case "A_TOOL_EXTENSIONS":
										instanceA_TOOL_EXTENSIONS := new(A_TOOL_EXTENSIONS)
										instanceA_TOOL_EXTENSIONS.Name = instanceName
										instanceA_TOOL_EXTENSIONS.Stage(stage)
										instance = any(instanceA_TOOL_EXTENSIONS)
										__gong__map_A_TOOL_EXTENSIONS[identifier] = instanceA_TOOL_EXTENSIONS
									case "DATATYPE_DEFINITION_BOOLEAN":
										instanceDATATYPE_DEFINITION_BOOLEAN := new(DATATYPE_DEFINITION_BOOLEAN)
										instanceDATATYPE_DEFINITION_BOOLEAN.Name = instanceName
										instanceDATATYPE_DEFINITION_BOOLEAN.Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_BOOLEAN)
										__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier] = instanceDATATYPE_DEFINITION_BOOLEAN
									case "DATATYPE_DEFINITION_DATE":
										instanceDATATYPE_DEFINITION_DATE := new(DATATYPE_DEFINITION_DATE)
										instanceDATATYPE_DEFINITION_DATE.Name = instanceName
										instanceDATATYPE_DEFINITION_DATE.Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_DATE)
										__gong__map_DATATYPE_DEFINITION_DATE[identifier] = instanceDATATYPE_DEFINITION_DATE
									case "DATATYPE_DEFINITION_ENUMERATION":
										instanceDATATYPE_DEFINITION_ENUMERATION := new(DATATYPE_DEFINITION_ENUMERATION)
										instanceDATATYPE_DEFINITION_ENUMERATION.Name = instanceName
										instanceDATATYPE_DEFINITION_ENUMERATION.Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_ENUMERATION)
										__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier] = instanceDATATYPE_DEFINITION_ENUMERATION
									case "DATATYPE_DEFINITION_INTEGER":
										instanceDATATYPE_DEFINITION_INTEGER := new(DATATYPE_DEFINITION_INTEGER)
										instanceDATATYPE_DEFINITION_INTEGER.Name = instanceName
										instanceDATATYPE_DEFINITION_INTEGER.Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_INTEGER)
										__gong__map_DATATYPE_DEFINITION_INTEGER[identifier] = instanceDATATYPE_DEFINITION_INTEGER
									case "DATATYPE_DEFINITION_REAL":
										instanceDATATYPE_DEFINITION_REAL := new(DATATYPE_DEFINITION_REAL)
										instanceDATATYPE_DEFINITION_REAL.Name = instanceName
										instanceDATATYPE_DEFINITION_REAL.Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_REAL)
										__gong__map_DATATYPE_DEFINITION_REAL[identifier] = instanceDATATYPE_DEFINITION_REAL
									case "DATATYPE_DEFINITION_STRING":
										instanceDATATYPE_DEFINITION_STRING := new(DATATYPE_DEFINITION_STRING)
										instanceDATATYPE_DEFINITION_STRING.Name = instanceName
										instanceDATATYPE_DEFINITION_STRING.Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_STRING)
										__gong__map_DATATYPE_DEFINITION_STRING[identifier] = instanceDATATYPE_DEFINITION_STRING
									case "DATATYPE_DEFINITION_XHTML":
										instanceDATATYPE_DEFINITION_XHTML := new(DATATYPE_DEFINITION_XHTML)
										instanceDATATYPE_DEFINITION_XHTML.Name = instanceName
										instanceDATATYPE_DEFINITION_XHTML.Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_XHTML)
										__gong__map_DATATYPE_DEFINITION_XHTML[identifier] = instanceDATATYPE_DEFINITION_XHTML
									case "EMBEDDED_VALUE":
										instanceEMBEDDED_VALUE := new(EMBEDDED_VALUE)
										instanceEMBEDDED_VALUE.Name = instanceName
										instanceEMBEDDED_VALUE.Stage(stage)
										instance = any(instanceEMBEDDED_VALUE)
										__gong__map_EMBEDDED_VALUE[identifier] = instanceEMBEDDED_VALUE
									case "ENUM_VALUE":
										instanceENUM_VALUE := new(ENUM_VALUE)
										instanceENUM_VALUE.Name = instanceName
										instanceENUM_VALUE.Stage(stage)
										instance = any(instanceENUM_VALUE)
										__gong__map_ENUM_VALUE[identifier] = instanceENUM_VALUE
									case "EmbeddedJpgImage":
										instanceEmbeddedJpgImage := new(EmbeddedJpgImage)
										instanceEmbeddedJpgImage.Name = instanceName
										instanceEmbeddedJpgImage.Stage(stage)
										instance = any(instanceEmbeddedJpgImage)
										__gong__map_EmbeddedJpgImage[identifier] = instanceEmbeddedJpgImage
									case "EmbeddedPngImage":
										instanceEmbeddedPngImage := new(EmbeddedPngImage)
										instanceEmbeddedPngImage.Name = instanceName
										instanceEmbeddedPngImage.Stage(stage)
										instance = any(instanceEmbeddedPngImage)
										__gong__map_EmbeddedPngImage[identifier] = instanceEmbeddedPngImage
									case "EmbeddedSvgImage":
										instanceEmbeddedSvgImage := new(EmbeddedSvgImage)
										instanceEmbeddedSvgImage.Name = instanceName
										instanceEmbeddedSvgImage.Stage(stage)
										instance = any(instanceEmbeddedSvgImage)
										__gong__map_EmbeddedSvgImage[identifier] = instanceEmbeddedSvgImage
									case "Kill":
										instanceKill := new(Kill)
										instanceKill.Name = instanceName
										instanceKill.Stage(stage)
										instance = any(instanceKill)
										__gong__map_Kill[identifier] = instanceKill
									case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry":
										instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry := new(Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry)
										instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry
									case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry":
										instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry := new(Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry)
										instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry
									case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry":
										instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry := new(Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry)
										instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry
									case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry":
										instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry := new(Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry)
										instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry
									case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry":
										instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry := new(Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry)
										instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry
									case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry":
										instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry := new(Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry)
										instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry
									case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry":
										instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry := new(Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry)
										instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
									case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry":
										instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry := new(Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry)
										instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry
									case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry":
										instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry := new(Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry)
										instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
									case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry":
										instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry := new(Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry)
										instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry
									case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry":
										instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry := new(Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry)
										instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry
									case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry":
										instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry := new(Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry)
										instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry
									case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry":
										instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry := new(Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry)
										instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry
									case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry":
										instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry := new(Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry)
										instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry
									case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry":
										instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry := new(Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry)
										instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry
									case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry":
										instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry := new(Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry)
										instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry
									case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry":
										instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry := new(Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry)
										instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry
									case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry":
										instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry := new(Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry)
										instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry
									case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry":
										instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry := new(Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry)
										instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry
									case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry":
										instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry := new(Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry)
										instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
									case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry":
										instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry := new(Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry)
										instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Name = instanceName
										instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Stage(stage)
										instance = any(instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry)
										__gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry[identifier] = instanceMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
									case "Map_SPECIFICATION_Nodes_expandedEntry":
										instanceMap_SPECIFICATION_Nodes_expandedEntry := new(Map_SPECIFICATION_Nodes_expandedEntry)
										instanceMap_SPECIFICATION_Nodes_expandedEntry.Name = instanceName
										instanceMap_SPECIFICATION_Nodes_expandedEntry.Stage(stage)
										instance = any(instanceMap_SPECIFICATION_Nodes_expandedEntry)
										__gong__map_Map_SPECIFICATION_Nodes_expandedEntry[identifier] = instanceMap_SPECIFICATION_Nodes_expandedEntry
									case "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry":
										instanceMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry := new(Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry)
										instanceMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Name = instanceName
										instanceMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Stage(stage)
										instance = any(instanceMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry)
										__gong__map_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry[identifier] = instanceMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
									case "Map_SPEC_OBJECT_TYPE_showIdentifierEntry":
										instanceMap_SPEC_OBJECT_TYPE_showIdentifierEntry := new(Map_SPEC_OBJECT_TYPE_showIdentifierEntry)
										instanceMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Name = instanceName
										instanceMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Stage(stage)
										instance = any(instanceMap_SPEC_OBJECT_TYPE_showIdentifierEntry)
										__gong__map_Map_SPEC_OBJECT_TYPE_showIdentifierEntry[identifier] = instanceMap_SPEC_OBJECT_TYPE_showIdentifierEntry
									case "Map_SPEC_OBJECT_TYPE_showNameEntry":
										instanceMap_SPEC_OBJECT_TYPE_showNameEntry := new(Map_SPEC_OBJECT_TYPE_showNameEntry)
										instanceMap_SPEC_OBJECT_TYPE_showNameEntry.Name = instanceName
										instanceMap_SPEC_OBJECT_TYPE_showNameEntry.Stage(stage)
										instance = any(instanceMap_SPEC_OBJECT_TYPE_showNameEntry)
										__gong__map_Map_SPEC_OBJECT_TYPE_showNameEntry[identifier] = instanceMap_SPEC_OBJECT_TYPE_showNameEntry
									case "RELATION_GROUP":
										instanceRELATION_GROUP := new(RELATION_GROUP)
										instanceRELATION_GROUP.Name = instanceName
										instanceRELATION_GROUP.Stage(stage)
										instance = any(instanceRELATION_GROUP)
										__gong__map_RELATION_GROUP[identifier] = instanceRELATION_GROUP
									case "RELATION_GROUP_TYPE":
										instanceRELATION_GROUP_TYPE := new(RELATION_GROUP_TYPE)
										instanceRELATION_GROUP_TYPE.Name = instanceName
										instanceRELATION_GROUP_TYPE.Stage(stage)
										instance = any(instanceRELATION_GROUP_TYPE)
										__gong__map_RELATION_GROUP_TYPE[identifier] = instanceRELATION_GROUP_TYPE
									case "REQ_IF":
										instanceREQ_IF := new(REQ_IF)
										instanceREQ_IF.Name = instanceName
										instanceREQ_IF.Stage(stage)
										instance = any(instanceREQ_IF)
										__gong__map_REQ_IF[identifier] = instanceREQ_IF
									case "REQ_IF_CONTENT":
										instanceREQ_IF_CONTENT := new(REQ_IF_CONTENT)
										instanceREQ_IF_CONTENT.Name = instanceName
										instanceREQ_IF_CONTENT.Stage(stage)
										instance = any(instanceREQ_IF_CONTENT)
										__gong__map_REQ_IF_CONTENT[identifier] = instanceREQ_IF_CONTENT
									case "REQ_IF_HEADER":
										instanceREQ_IF_HEADER := new(REQ_IF_HEADER)
										instanceREQ_IF_HEADER.Name = instanceName
										instanceREQ_IF_HEADER.Stage(stage)
										instance = any(instanceREQ_IF_HEADER)
										__gong__map_REQ_IF_HEADER[identifier] = instanceREQ_IF_HEADER
									case "REQ_IF_TOOL_EXTENSION":
										instanceREQ_IF_TOOL_EXTENSION := new(REQ_IF_TOOL_EXTENSION)
										instanceREQ_IF_TOOL_EXTENSION.Name = instanceName
										instanceREQ_IF_TOOL_EXTENSION.Stage(stage)
										instance = any(instanceREQ_IF_TOOL_EXTENSION)
										__gong__map_REQ_IF_TOOL_EXTENSION[identifier] = instanceREQ_IF_TOOL_EXTENSION
									case "RenderingConfiguration":
										instanceRenderingConfiguration := new(RenderingConfiguration)
										instanceRenderingConfiguration.Name = instanceName
										instanceRenderingConfiguration.Stage(stage)
										instance = any(instanceRenderingConfiguration)
										__gong__map_RenderingConfiguration[identifier] = instanceRenderingConfiguration
									case "SPECIFICATION":
										instanceSPECIFICATION := new(SPECIFICATION)
										instanceSPECIFICATION.Name = instanceName
										instanceSPECIFICATION.Stage(stage)
										instance = any(instanceSPECIFICATION)
										__gong__map_SPECIFICATION[identifier] = instanceSPECIFICATION
									case "SPECIFICATION_TYPE":
										instanceSPECIFICATION_TYPE := new(SPECIFICATION_TYPE)
										instanceSPECIFICATION_TYPE.Name = instanceName
										instanceSPECIFICATION_TYPE.Stage(stage)
										instance = any(instanceSPECIFICATION_TYPE)
										__gong__map_SPECIFICATION_TYPE[identifier] = instanceSPECIFICATION_TYPE
									case "SPEC_HIERARCHY":
										instanceSPEC_HIERARCHY := new(SPEC_HIERARCHY)
										instanceSPEC_HIERARCHY.Name = instanceName
										instanceSPEC_HIERARCHY.Stage(stage)
										instance = any(instanceSPEC_HIERARCHY)
										__gong__map_SPEC_HIERARCHY[identifier] = instanceSPEC_HIERARCHY
									case "SPEC_OBJECT":
										instanceSPEC_OBJECT := new(SPEC_OBJECT)
										instanceSPEC_OBJECT.Name = instanceName
										instanceSPEC_OBJECT.Stage(stage)
										instance = any(instanceSPEC_OBJECT)
										__gong__map_SPEC_OBJECT[identifier] = instanceSPEC_OBJECT
									case "SPEC_OBJECT_TYPE":
										instanceSPEC_OBJECT_TYPE := new(SPEC_OBJECT_TYPE)
										instanceSPEC_OBJECT_TYPE.Name = instanceName
										instanceSPEC_OBJECT_TYPE.Stage(stage)
										instance = any(instanceSPEC_OBJECT_TYPE)
										__gong__map_SPEC_OBJECT_TYPE[identifier] = instanceSPEC_OBJECT_TYPE
									case "SPEC_RELATION":
										instanceSPEC_RELATION := new(SPEC_RELATION)
										instanceSPEC_RELATION.Name = instanceName
										instanceSPEC_RELATION.Stage(stage)
										instance = any(instanceSPEC_RELATION)
										__gong__map_SPEC_RELATION[identifier] = instanceSPEC_RELATION
									case "SPEC_RELATION_TYPE":
										instanceSPEC_RELATION_TYPE := new(SPEC_RELATION_TYPE)
										instanceSPEC_RELATION_TYPE.Name = instanceName
										instanceSPEC_RELATION_TYPE.Stage(stage)
										instance = any(instanceSPEC_RELATION_TYPE)
										__gong__map_SPEC_RELATION_TYPE[identifier] = instanceSPEC_RELATION_TYPE
									case "StaticWebSite":
										instanceStaticWebSite := new(StaticWebSite)
										instanceStaticWebSite.Name = instanceName
										instanceStaticWebSite.Stage(stage)
										instance = any(instanceStaticWebSite)
										__gong__map_StaticWebSite[identifier] = instanceStaticWebSite
									case "StaticWebSiteChapter":
										instanceStaticWebSiteChapter := new(StaticWebSiteChapter)
										instanceStaticWebSiteChapter.Name = instanceName
										instanceStaticWebSiteChapter.Stage(stage)
										instance = any(instanceStaticWebSiteChapter)
										__gong__map_StaticWebSiteChapter[identifier] = instanceStaticWebSiteChapter
									case "StaticWebSiteGeneratedImage":
										instanceStaticWebSiteGeneratedImage := new(StaticWebSiteGeneratedImage)
										instanceStaticWebSiteGeneratedImage.Name = instanceName
										instanceStaticWebSiteGeneratedImage.Stage(stage)
										instance = any(instanceStaticWebSiteGeneratedImage)
										__gong__map_StaticWebSiteGeneratedImage[identifier] = instanceStaticWebSiteGeneratedImage
									case "StaticWebSiteImage":
										instanceStaticWebSiteImage := new(StaticWebSiteImage)
										instanceStaticWebSiteImage.Name = instanceName
										instanceStaticWebSiteImage.Stage(stage)
										instance = any(instanceStaticWebSiteImage)
										__gong__map_StaticWebSiteImage[identifier] = instanceStaticWebSiteImage
									case "StaticWebSiteParagraph":
										instanceStaticWebSiteParagraph := new(StaticWebSiteParagraph)
										instanceStaticWebSiteParagraph.Name = instanceName
										instanceStaticWebSiteParagraph.Stage(stage)
										instance = any(instanceStaticWebSiteParagraph)
										__gong__map_StaticWebSiteParagraph[identifier] = instanceStaticWebSiteParagraph
									case "XHTML_CONTENT":
										instanceXHTML_CONTENT := new(XHTML_CONTENT)
										instanceXHTML_CONTENT.Name = instanceName
										instanceXHTML_CONTENT.Stage(stage)
										instance = any(instanceXHTML_CONTENT)
										__gong__map_XHTML_CONTENT[identifier] = instanceXHTML_CONTENT
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Println("gongstructName not found for identifier", identifier)
							break
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "ALTERNATIVE_ID":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_BOOLEAN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_DATE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_ENUMERATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_INTEGER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_REAL":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_STRING":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_XHTML":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_BOOLEAN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_DATE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_ENUMERATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_INTEGER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_REAL":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_STRING":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_XHTML":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ALTERNATIVE_ID":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_DEFINITION_DATE_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_DEFINITION_REAL_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_DEFINITION_STRING_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_VALUE_BOOLEAN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_VALUE_DATE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_VALUE_ENUMERATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_VALUE_INTEGER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_VALUE_REAL":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_VALUE_STRING":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_VALUE_XHTML":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ATTRIBUTE_VALUE_XHTML_1":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_CHILDREN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_CORE_CONTENT":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_DATATYPES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_DATATYPE_DEFINITION_DATE_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_DATATYPE_DEFINITION_INTEGER_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_DATATYPE_DEFINITION_REAL_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_DATATYPE_DEFINITION_STRING_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_DATATYPE_DEFINITION_XHTML_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_EDITABLE_ATTS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_ENUM_VALUE_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_OBJECT":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_PROPERTIES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_RELATION_GROUP_TYPE_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SOURCE_1":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SOURCE_SPECIFICATION_1":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPECIFICATIONS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPECIFICATION_TYPE_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPECIFIED_VALUES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPEC_ATTRIBUTES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPEC_OBJECTS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPEC_OBJECT_TYPE_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPEC_RELATIONS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPEC_RELATION_GROUPS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPEC_RELATION_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPEC_RELATION_TYPE_REF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_SPEC_TYPES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_THE_HEADER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "A_TOOL_EXTENSIONS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_BOOLEAN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_DATE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_ENUMERATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_INTEGER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_REAL":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_STRING":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_XHTML":
							switch fieldName {
							// insertion point for date assign code
							}
						case "EMBEDDED_VALUE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ENUM_VALUE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "EmbeddedJpgImage":
							switch fieldName {
							// insertion point for date assign code
							}
						case "EmbeddedPngImage":
							switch fieldName {
							// insertion point for date assign code
							}
						case "EmbeddedSvgImage":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Kill":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_SPECIFICATION_Nodes_expandedEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_SPEC_OBJECT_TYPE_showIdentifierEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Map_SPEC_OBJECT_TYPE_showNameEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RELATION_GROUP":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RELATION_GROUP_TYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQ_IF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQ_IF_CONTENT":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQ_IF_HEADER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQ_IF_TOOL_EXTENSION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RenderingConfiguration":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECIFICATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECIFICATION_TYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPEC_HIERARCHY":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPEC_OBJECT":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPEC_OBJECT_TYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPEC_RELATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPEC_RELATION_TYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "StaticWebSite":
							switch fieldName {
							// insertion point for date assign code
							}
						case "StaticWebSiteChapter":
							switch fieldName {
							// insertion point for date assign code
							}
						case "StaticWebSiteGeneratedImage":
							switch fieldName {
							// insertion point for date assign code
							}
						case "StaticWebSiteImage":
							switch fieldName {
							// insertion point for date assign code
							}
						case "StaticWebSiteParagraph":
							switch fieldName {
							// insertion point for date assign code
							}
						case "XHTML_CONTENT":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// pick up the first arg
				if len(callExpr.Args) != 1 {
					break
				}
				arg0 := callExpr.Args[0]

				var se *ast.SelectorExpr
				var ok bool
				if se, ok = arg0.(*ast.SelectorExpr); !ok {
					break
				}

				var seXident *ast.Ident
				if seXident = se.X.(*ast.Ident); !ok {
					break
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = "new(" + seXident.Name + "." + se.Sel.Name + ")"
				// following lines are here to avoid warning "unused write to field..."
				_ = basicLit.Kind
				_ = basicLit.Value
				_ = basicLit
			}
			for argNb, arg := range callExpr.Args {
				_ = argNb
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident, *ast.SelectorExpr:
					var ident *ast.Ident
					var ok bool
					_ = ok
					if ident, ok = arg.(*ast.Ident); !ok {
						// log.Println("we are in the case of new(....)")
					}

					var se *ast.SelectorExpr
					if se, ok = arg.(*ast.SelectorExpr); ok {
						if ident, ok = se.X.(*ast.Ident); !ok {
							// log.Println("we are in the case of append(....)")
						}
					}
					_ = ident

					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Println("gongstructName not found for identifier", identifier)
						break
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "ALTERNATIVE_ID":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_BOOLEAN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_DATE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_ENUMERATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_INTEGER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_REAL":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_STRING":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_XHTML":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_BOOLEAN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_DATE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_ENUMERATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_INTEGER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_REAL":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_STRING":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_XHTML":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_ALTERNATIVE_ID":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_ATTRIBUTE_DEFINITION_DATE_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_ATTRIBUTE_DEFINITION_REAL_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_ATTRIBUTE_DEFINITION_STRING_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_ATTRIBUTE_VALUE_BOOLEAN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ATTRIBUTE_VALUE_BOOLEAN":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_BOOLEAN[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_BOOLEAN[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_BOOLEAN = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_BOOLEAN, instanceToAppend)
							}
						}
					case "A_ATTRIBUTE_VALUE_DATE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ATTRIBUTE_VALUE_DATE":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_DATE[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_DATE[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_DATE = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_DATE, instanceToAppend)
							}
						}
					case "A_ATTRIBUTE_VALUE_ENUMERATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ATTRIBUTE_VALUE_ENUMERATION":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_ENUMERATION[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_ENUMERATION[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_ENUMERATION = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_ENUMERATION, instanceToAppend)
							}
						}
					case "A_ATTRIBUTE_VALUE_INTEGER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ATTRIBUTE_VALUE_INTEGER":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_INTEGER[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_INTEGER[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_INTEGER = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_INTEGER, instanceToAppend)
							}
						}
					case "A_ATTRIBUTE_VALUE_REAL":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ATTRIBUTE_VALUE_REAL":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_REAL[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_REAL[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_REAL = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_REAL, instanceToAppend)
							}
						}
					case "A_ATTRIBUTE_VALUE_STRING":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ATTRIBUTE_VALUE_STRING":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_STRING[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_STRING[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_STRING = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_STRING, instanceToAppend)
							}
						}
					case "A_ATTRIBUTE_VALUE_XHTML":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ATTRIBUTE_VALUE_XHTML":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_XHTML[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_XHTML[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_XHTML = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_XHTML, instanceToAppend)
							}
						}
					case "A_ATTRIBUTE_VALUE_XHTML_1":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ATTRIBUTE_VALUE_BOOLEAN":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_BOOLEAN[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_XHTML_1[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_BOOLEAN = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_BOOLEAN, instanceToAppend)
							}
						case "ATTRIBUTE_VALUE_DATE":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_DATE[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_XHTML_1[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_DATE = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_DATE, instanceToAppend)
							}
						case "ATTRIBUTE_VALUE_ENUMERATION":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_ENUMERATION[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_XHTML_1[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_ENUMERATION = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_ENUMERATION, instanceToAppend)
							}
						case "ATTRIBUTE_VALUE_INTEGER":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_INTEGER[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_XHTML_1[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_INTEGER = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_INTEGER, instanceToAppend)
							}
						case "ATTRIBUTE_VALUE_REAL":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_REAL[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_XHTML_1[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_REAL = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_REAL, instanceToAppend)
							}
						case "ATTRIBUTE_VALUE_STRING":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_STRING[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_XHTML_1[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_STRING = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_STRING, instanceToAppend)
							}
						case "ATTRIBUTE_VALUE_XHTML":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_VALUE_XHTML[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_ATTRIBUTE_VALUE_XHTML_1[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_XHTML = append(instanceWhoseFieldIsAppended.ATTRIBUTE_VALUE_XHTML, instanceToAppend)
							}
						}
					case "A_CHILDREN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SPEC_HIERARCHY":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SPEC_HIERARCHY[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_CHILDREN[identifier]
								instanceWhoseFieldIsAppended.SPEC_HIERARCHY = append(instanceWhoseFieldIsAppended.SPEC_HIERARCHY, instanceToAppend)
							}
						}
					case "A_CORE_CONTENT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_DATATYPES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "DATATYPE_DEFINITION_BOOLEAN":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_DATATYPE_DEFINITION_BOOLEAN[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_DATATYPES[identifier]
								instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_BOOLEAN = append(instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_BOOLEAN, instanceToAppend)
							}
						case "DATATYPE_DEFINITION_DATE":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_DATATYPE_DEFINITION_DATE[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_DATATYPES[identifier]
								instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_DATE = append(instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_DATE, instanceToAppend)
							}
						case "DATATYPE_DEFINITION_ENUMERATION":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_DATATYPE_DEFINITION_ENUMERATION[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_DATATYPES[identifier]
								instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_ENUMERATION = append(instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_ENUMERATION, instanceToAppend)
							}
						case "DATATYPE_DEFINITION_INTEGER":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_DATATYPE_DEFINITION_INTEGER[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_DATATYPES[identifier]
								instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_INTEGER = append(instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_INTEGER, instanceToAppend)
							}
						case "DATATYPE_DEFINITION_REAL":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_DATATYPE_DEFINITION_REAL[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_DATATYPES[identifier]
								instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_REAL = append(instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_REAL, instanceToAppend)
							}
						case "DATATYPE_DEFINITION_STRING":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_DATATYPE_DEFINITION_STRING[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_DATATYPES[identifier]
								instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_STRING = append(instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_STRING, instanceToAppend)
							}
						case "DATATYPE_DEFINITION_XHTML":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_DATATYPE_DEFINITION_XHTML[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_DATATYPES[identifier]
								instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_XHTML = append(instanceWhoseFieldIsAppended.DATATYPE_DEFINITION_XHTML, instanceToAppend)
							}
						}
					case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_DATATYPE_DEFINITION_DATE_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_DATATYPE_DEFINITION_INTEGER_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_DATATYPE_DEFINITION_REAL_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_DATATYPE_DEFINITION_STRING_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_DATATYPE_DEFINITION_XHTML_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_EDITABLE_ATTS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_ENUM_VALUE_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_OBJECT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_PROPERTIES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_RELATION_GROUP_TYPE_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_SOURCE_1":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_SOURCE_SPECIFICATION_1":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_SPECIFICATIONS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SPECIFICATION":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SPECIFICATION[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPECIFICATIONS[identifier]
								instanceWhoseFieldIsAppended.SPECIFICATION = append(instanceWhoseFieldIsAppended.SPECIFICATION, instanceToAppend)
							}
						}
					case "A_SPECIFICATION_TYPE_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_SPECIFIED_VALUES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ENUM_VALUE":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ENUM_VALUE[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPECIFIED_VALUES[identifier]
								instanceWhoseFieldIsAppended.ENUM_VALUE = append(instanceWhoseFieldIsAppended.ENUM_VALUE, instanceToAppend)
							}
						}
					case "A_SPEC_ATTRIBUTES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ATTRIBUTE_DEFINITION_BOOLEAN":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_ATTRIBUTES[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_BOOLEAN = append(instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_BOOLEAN, instanceToAppend)
							}
						case "ATTRIBUTE_DEFINITION_DATE":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_DEFINITION_DATE[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_ATTRIBUTES[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_DATE = append(instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_DATE, instanceToAppend)
							}
						case "ATTRIBUTE_DEFINITION_ENUMERATION":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_ATTRIBUTES[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_ENUMERATION = append(instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_ENUMERATION, instanceToAppend)
							}
						case "ATTRIBUTE_DEFINITION_INTEGER":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_ATTRIBUTES[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_INTEGER = append(instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_INTEGER, instanceToAppend)
							}
						case "ATTRIBUTE_DEFINITION_REAL":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_DEFINITION_REAL[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_ATTRIBUTES[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_REAL = append(instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_REAL, instanceToAppend)
							}
						case "ATTRIBUTE_DEFINITION_STRING":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_DEFINITION_STRING[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_ATTRIBUTES[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_STRING = append(instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_STRING, instanceToAppend)
							}
						case "ATTRIBUTE_DEFINITION_XHTML":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ATTRIBUTE_DEFINITION_XHTML[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_ATTRIBUTES[identifier]
								instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_XHTML = append(instanceWhoseFieldIsAppended.ATTRIBUTE_DEFINITION_XHTML, instanceToAppend)
							}
						}
					case "A_SPEC_OBJECTS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SPEC_OBJECT":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SPEC_OBJECT[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_OBJECTS[identifier]
								instanceWhoseFieldIsAppended.SPEC_OBJECT = append(instanceWhoseFieldIsAppended.SPEC_OBJECT, instanceToAppend)
							}
						}
					case "A_SPEC_OBJECT_TYPE_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_SPEC_RELATIONS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SPEC_RELATION":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SPEC_RELATION[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_RELATIONS[identifier]
								instanceWhoseFieldIsAppended.SPEC_RELATION = append(instanceWhoseFieldIsAppended.SPEC_RELATION, instanceToAppend)
							}
						}
					case "A_SPEC_RELATION_GROUPS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "RELATION_GROUP":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_RELATION_GROUP[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_RELATION_GROUPS[identifier]
								instanceWhoseFieldIsAppended.RELATION_GROUP = append(instanceWhoseFieldIsAppended.RELATION_GROUP, instanceToAppend)
							}
						}
					case "A_SPEC_RELATION_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_SPEC_RELATION_TYPE_REF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_SPEC_TYPES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "RELATION_GROUP_TYPE":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_RELATION_GROUP_TYPE[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_TYPES[identifier]
								instanceWhoseFieldIsAppended.RELATION_GROUP_TYPE = append(instanceWhoseFieldIsAppended.RELATION_GROUP_TYPE, instanceToAppend)
							}
						case "SPEC_OBJECT_TYPE":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SPEC_OBJECT_TYPE[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_TYPES[identifier]
								instanceWhoseFieldIsAppended.SPEC_OBJECT_TYPE = append(instanceWhoseFieldIsAppended.SPEC_OBJECT_TYPE, instanceToAppend)
							}
						case "SPEC_RELATION_TYPE":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SPEC_RELATION_TYPE[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_TYPES[identifier]
								instanceWhoseFieldIsAppended.SPEC_RELATION_TYPE = append(instanceWhoseFieldIsAppended.SPEC_RELATION_TYPE, instanceToAppend)
							}
						case "SPECIFICATION_TYPE":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SPECIFICATION_TYPE[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_SPEC_TYPES[identifier]
								instanceWhoseFieldIsAppended.SPECIFICATION_TYPE = append(instanceWhoseFieldIsAppended.SPECIFICATION_TYPE, instanceToAppend)
							}
						}
					case "A_THE_HEADER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "A_TOOL_EXTENSIONS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "REQ_IF_TOOL_EXTENSION":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_REQ_IF_TOOL_EXTENSION[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_A_TOOL_EXTENSIONS[identifier]
								instanceWhoseFieldIsAppended.REQ_IF_TOOL_EXTENSION = append(instanceWhoseFieldIsAppended.REQ_IF_TOOL_EXTENSION, instanceToAppend)
							}
						}
					case "DATATYPE_DEFINITION_BOOLEAN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_DATE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_ENUMERATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_INTEGER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_REAL":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_STRING":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_XHTML":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "EMBEDDED_VALUE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ENUM_VALUE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "EmbeddedJpgImage":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "EmbeddedPngImage":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "EmbeddedSvgImage":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Kill":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_SPECIFICATION_Nodes_expandedEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_SPEC_OBJECT_TYPE_showIdentifierEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Map_SPEC_OBJECT_TYPE_showNameEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RELATION_GROUP":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RELATION_GROUP_TYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQ_IF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQ_IF_CONTENT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQ_IF_HEADER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQ_IF_TOOL_EXTENSION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RenderingConfiguration":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries = append(instanceWhoseFieldIsAppended.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries, instanceToAppend)
							}
						case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries = append(instanceWhoseFieldIsAppended.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries, instanceToAppend)
							}
						case "Map_SPECIFICATION_Nodes_expandedEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_SPECIFICATION_Nodes_expandedEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_SPECIFICATION_Nodes_expandedEntries = append(instanceWhoseFieldIsAppended.Map_SPECIFICATION_Nodes_expandedEntries, instanceToAppend)
							}
						case "Map_SPEC_OBJECT_TYPE_showIdentifierEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_SPEC_OBJECT_TYPE_showIdentifierEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_SPEC_OBJECT_TYPE_showIdentifierEntries = append(instanceWhoseFieldIsAppended.Map_SPEC_OBJECT_TYPE_showIdentifierEntries, instanceToAppend)
							}
						case "Map_SPEC_OBJECT_TYPE_showNameEntries":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Map_SPEC_OBJECT_TYPE_showNameEntry[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RenderingConfiguration[identifier]
								instanceWhoseFieldIsAppended.Map_SPEC_OBJECT_TYPE_showNameEntries = append(instanceWhoseFieldIsAppended.Map_SPEC_OBJECT_TYPE_showNameEntries, instanceToAppend)
							}
						}
					case "SPECIFICATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECIFICATION_TYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPEC_HIERARCHY":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPEC_OBJECT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPEC_OBJECT_TYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPEC_RELATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPEC_RELATION_TYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "StaticWebSite":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Chapters":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_StaticWebSiteChapter[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_StaticWebSite[identifier]
								instanceWhoseFieldIsAppended.Chapters = append(instanceWhoseFieldIsAppended.Chapters, instanceToAppend)
							}
						}
					case "StaticWebSiteChapter":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Paragraphs":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_StaticWebSiteParagraph[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_StaticWebSiteChapter[identifier]
								instanceWhoseFieldIsAppended.Paragraphs = append(instanceWhoseFieldIsAppended.Paragraphs, instanceToAppend)
							}
						}
					case "StaticWebSiteGeneratedImage":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "StaticWebSiteImage":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "StaticWebSiteParagraph":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "XHTML_CONTENT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				}
			}
		case *ast.BasicLit, *ast.UnaryExpr, *ast.CompositeLit:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used
			switch v := expr.(type) {
			case *ast.BasicLit:
				// expression is for instance ... = 18.000
				basicLit = v
			case *ast.UnaryExpr:
				// expression is for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				if bl, ok := v.X.(*ast.BasicLit); ok {
					basicLit = bl
					// Check the operator to set the sign
					if v.Op == token.SUB { // token.SUB is for '-'
						exprSign = -1
					} else if v.Op == token.ADD { // token.ADD is for '+'
						exprSign = 1
					}
				}
			case *ast.CompositeLit:
				var sl *ast.SelectorExpr
				var ident *ast.Ident
				var ok bool

				if sl, ok = v.Type.(*ast.SelectorExpr); !ok {
					break // Exits the switch case
				}

				if ident, ok = sl.X.(*ast.Ident); !ok {
					break // Exits the switch case
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = ident.Name + "." + sl.Sel.Name + "{}"
			}

			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Println("gongstructName not found for identifier", identifier)
				break
			}

			// substitute the RHS part of the assignment if a //gong:ident directive is met
			if hasGongIdentDirective {
				basicLit.Value = "[" + docLinkText + "]"
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "ALTERNATIVE_ID":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ALTERNATIVE_ID[identifier].Name = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ALTERNATIVE_ID[identifier].IDENTIFIER = fielValue
				}
			case "ATTRIBUTE_DEFINITION_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_VALUE_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_BOOLEAN[identifier].Name = fielValue
				}
			case "ATTRIBUTE_VALUE_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_DATE[identifier].Name = fielValue
				case "THE_VALUE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_DATE[identifier].THE_VALUE = fielValue
				}
			case "ATTRIBUTE_VALUE_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_ENUMERATION[identifier].Name = fielValue
				}
			case "ATTRIBUTE_VALUE_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_INTEGER[identifier].Name = fielValue
				case "THE_VALUE":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_VALUE_INTEGER[identifier].THE_VALUE = int(exprSign) * int(fielValue)
				}
			case "ATTRIBUTE_VALUE_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_REAL[identifier].Name = fielValue
				case "THE_VALUE":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_VALUE_REAL[identifier].THE_VALUE = exprSign * fielValue
				}
			case "ATTRIBUTE_VALUE_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_STRING[identifier].Name = fielValue
				case "THE_VALUE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_STRING[identifier].THE_VALUE = fielValue
				}
			case "ATTRIBUTE_VALUE_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].Name = fielValue
				}
			case "A_ALTERNATIVE_ID":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ALTERNATIVE_ID[identifier].Name = fielValue
				}
			case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF[identifier].Name = fielValue
				case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF[identifier].ATTRIBUTE_DEFINITION_BOOLEAN_REF = fielValue
				}
			case "A_ATTRIBUTE_DEFINITION_DATE_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_DATE_REF[identifier].Name = fielValue
				case "ATTRIBUTE_DEFINITION_DATE_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_DATE_REF[identifier].ATTRIBUTE_DEFINITION_DATE_REF = fielValue
				}
			case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF[identifier].Name = fielValue
				case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF[identifier].ATTRIBUTE_DEFINITION_ENUMERATION_REF = fielValue
				}
			case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_INTEGER_REF[identifier].Name = fielValue
				case "ATTRIBUTE_DEFINITION_INTEGER_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_INTEGER_REF[identifier].ATTRIBUTE_DEFINITION_INTEGER_REF = fielValue
				}
			case "A_ATTRIBUTE_DEFINITION_REAL_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_REAL_REF[identifier].Name = fielValue
				case "ATTRIBUTE_DEFINITION_REAL_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_REAL_REF[identifier].ATTRIBUTE_DEFINITION_REAL_REF = fielValue
				}
			case "A_ATTRIBUTE_DEFINITION_STRING_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_STRING_REF[identifier].Name = fielValue
				case "ATTRIBUTE_DEFINITION_STRING_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_STRING_REF[identifier].ATTRIBUTE_DEFINITION_STRING_REF = fielValue
				}
			case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_XHTML_REF[identifier].Name = fielValue
				case "ATTRIBUTE_DEFINITION_XHTML_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_DEFINITION_XHTML_REF[identifier].ATTRIBUTE_DEFINITION_XHTML_REF = fielValue
				}
			case "A_ATTRIBUTE_VALUE_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_VALUE_BOOLEAN[identifier].Name = fielValue
				}
			case "A_ATTRIBUTE_VALUE_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_VALUE_DATE[identifier].Name = fielValue
				}
			case "A_ATTRIBUTE_VALUE_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_VALUE_ENUMERATION[identifier].Name = fielValue
				}
			case "A_ATTRIBUTE_VALUE_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_VALUE_INTEGER[identifier].Name = fielValue
				}
			case "A_ATTRIBUTE_VALUE_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_VALUE_REAL[identifier].Name = fielValue
				}
			case "A_ATTRIBUTE_VALUE_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_VALUE_STRING[identifier].Name = fielValue
				}
			case "A_ATTRIBUTE_VALUE_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_VALUE_XHTML[identifier].Name = fielValue
				}
			case "A_ATTRIBUTE_VALUE_XHTML_1":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ATTRIBUTE_VALUE_XHTML_1[identifier].Name = fielValue
				}
			case "A_CHILDREN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_CHILDREN[identifier].Name = fielValue
				}
			case "A_CORE_CONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_CORE_CONTENT[identifier].Name = fielValue
				}
			case "A_DATATYPES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPES[identifier].Name = fielValue
				}
			case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_BOOLEAN_REF[identifier].Name = fielValue
				case "DATATYPE_DEFINITION_BOOLEAN_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_BOOLEAN_REF[identifier].DATATYPE_DEFINITION_BOOLEAN_REF = fielValue
				}
			case "A_DATATYPE_DEFINITION_DATE_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_DATE_REF[identifier].Name = fielValue
				case "DATATYPE_DEFINITION_DATE_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_DATE_REF[identifier].DATATYPE_DEFINITION_DATE_REF = fielValue
				}
			case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_ENUMERATION_REF[identifier].Name = fielValue
				case "DATATYPE_DEFINITION_ENUMERATION_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_ENUMERATION_REF[identifier].DATATYPE_DEFINITION_ENUMERATION_REF = fielValue
				}
			case "A_DATATYPE_DEFINITION_INTEGER_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_INTEGER_REF[identifier].Name = fielValue
				case "DATATYPE_DEFINITION_INTEGER_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_INTEGER_REF[identifier].DATATYPE_DEFINITION_INTEGER_REF = fielValue
				}
			case "A_DATATYPE_DEFINITION_REAL_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_REAL_REF[identifier].Name = fielValue
				case "DATATYPE_DEFINITION_REAL_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_REAL_REF[identifier].DATATYPE_DEFINITION_REAL_REF = fielValue
				}
			case "A_DATATYPE_DEFINITION_STRING_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_STRING_REF[identifier].Name = fielValue
				case "DATATYPE_DEFINITION_STRING_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_STRING_REF[identifier].DATATYPE_DEFINITION_STRING_REF = fielValue
				}
			case "A_DATATYPE_DEFINITION_XHTML_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_XHTML_REF[identifier].Name = fielValue
				case "DATATYPE_DEFINITION_XHTML_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_DATATYPE_DEFINITION_XHTML_REF[identifier].DATATYPE_DEFINITION_XHTML_REF = fielValue
				}
			case "A_EDITABLE_ATTS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_EDITABLE_ATTS[identifier].Name = fielValue
				case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_EDITABLE_ATTS[identifier].ATTRIBUTE_DEFINITION_BOOLEAN_REF = fielValue
				case "ATTRIBUTE_DEFINITION_DATE_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_EDITABLE_ATTS[identifier].ATTRIBUTE_DEFINITION_DATE_REF = fielValue
				case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_EDITABLE_ATTS[identifier].ATTRIBUTE_DEFINITION_ENUMERATION_REF = fielValue
				case "ATTRIBUTE_DEFINITION_INTEGER_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_EDITABLE_ATTS[identifier].ATTRIBUTE_DEFINITION_INTEGER_REF = fielValue
				case "ATTRIBUTE_DEFINITION_REAL_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_EDITABLE_ATTS[identifier].ATTRIBUTE_DEFINITION_REAL_REF = fielValue
				case "ATTRIBUTE_DEFINITION_STRING_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_EDITABLE_ATTS[identifier].ATTRIBUTE_DEFINITION_STRING_REF = fielValue
				case "ATTRIBUTE_DEFINITION_XHTML_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_EDITABLE_ATTS[identifier].ATTRIBUTE_DEFINITION_XHTML_REF = fielValue
				}
			case "A_ENUM_VALUE_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ENUM_VALUE_REF[identifier].Name = fielValue
				case "ENUM_VALUE_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_ENUM_VALUE_REF[identifier].ENUM_VALUE_REF = fielValue
				}
			case "A_OBJECT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_OBJECT[identifier].Name = fielValue
				case "SPEC_OBJECT_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_OBJECT[identifier].SPEC_OBJECT_REF = fielValue
				}
			case "A_PROPERTIES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_PROPERTIES[identifier].Name = fielValue
				}
			case "A_RELATION_GROUP_TYPE_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_RELATION_GROUP_TYPE_REF[identifier].Name = fielValue
				case "RELATION_GROUP_TYPE_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_RELATION_GROUP_TYPE_REF[identifier].RELATION_GROUP_TYPE_REF = fielValue
				}
			case "A_SOURCE_1":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SOURCE_1[identifier].Name = fielValue
				case "SPEC_OBJECT_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SOURCE_1[identifier].SPEC_OBJECT_REF = fielValue
				}
			case "A_SOURCE_SPECIFICATION_1":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SOURCE_SPECIFICATION_1[identifier].Name = fielValue
				}
			case "A_SPECIFICATIONS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPECIFICATIONS[identifier].Name = fielValue
				}
			case "A_SPECIFICATION_TYPE_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPECIFICATION_TYPE_REF[identifier].Name = fielValue
				case "SPECIFICATION_TYPE_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPECIFICATION_TYPE_REF[identifier].SPECIFICATION_TYPE_REF = fielValue
				}
			case "A_SPECIFIED_VALUES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPECIFIED_VALUES[identifier].Name = fielValue
				}
			case "A_SPEC_ATTRIBUTES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_ATTRIBUTES[identifier].Name = fielValue
				}
			case "A_SPEC_OBJECTS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_OBJECTS[identifier].Name = fielValue
				}
			case "A_SPEC_OBJECT_TYPE_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_OBJECT_TYPE_REF[identifier].Name = fielValue
				case "SPEC_OBJECT_TYPE_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_OBJECT_TYPE_REF[identifier].SPEC_OBJECT_TYPE_REF = fielValue
				}
			case "A_SPEC_RELATIONS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_RELATIONS[identifier].Name = fielValue
				}
			case "A_SPEC_RELATION_GROUPS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_RELATION_GROUPS[identifier].Name = fielValue
				}
			case "A_SPEC_RELATION_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_RELATION_REF[identifier].Name = fielValue
				case "SPEC_RELATION_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_RELATION_REF[identifier].SPEC_RELATION_REF = fielValue
				}
			case "A_SPEC_RELATION_TYPE_REF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_RELATION_TYPE_REF[identifier].Name = fielValue
				case "SPEC_RELATION_TYPE_REF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_RELATION_TYPE_REF[identifier].SPEC_RELATION_TYPE_REF = fielValue
				}
			case "A_SPEC_TYPES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_SPEC_TYPES[identifier].Name = fielValue
				}
			case "A_THE_HEADER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_THE_HEADER[identifier].Name = fielValue
				}
			case "A_TOOL_EXTENSIONS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_A_TOOL_EXTENSIONS[identifier].Name = fielValue
				}
			case "DATATYPE_DEFINITION_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].LONG_NAME = fielValue
				}
			case "DATATYPE_DEFINITION_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].LONG_NAME = fielValue
				}
			case "DATATYPE_DEFINITION_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].LONG_NAME = fielValue
				}
			case "DATATYPE_DEFINITION_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].LONG_NAME = fielValue
				case "MAX":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].MAX = int(exprSign) * int(fielValue)
				case "MIN":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].MIN = int(exprSign) * int(fielValue)
				}
			case "DATATYPE_DEFINITION_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].Name = fielValue
				case "ACCURACY":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].ACCURACY = int(exprSign) * int(fielValue)
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].LONG_NAME = fielValue
				case "MAX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].MAX = exprSign * fielValue
				case "MIN":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].MIN = exprSign * fielValue
				}
			case "DATATYPE_DEFINITION_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].LONG_NAME = fielValue
				case "MAX_LENGTH":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].MAX_LENGTH = int(exprSign) * int(fielValue)
				}
			case "DATATYPE_DEFINITION_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].LONG_NAME = fielValue
				}
			case "EMBEDDED_VALUE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EMBEDDED_VALUE[identifier].Name = fielValue
				case "KEY":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_EMBEDDED_VALUE[identifier].KEY = int(exprSign) * int(fielValue)
				case "OTHER_CONTENT":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EMBEDDED_VALUE[identifier].OTHER_CONTENT = fielValue
				}
			case "ENUM_VALUE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUM_VALUE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUM_VALUE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUM_VALUE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUM_VALUE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUM_VALUE[identifier].LONG_NAME = fielValue
				}
			case "EmbeddedJpgImage":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EmbeddedJpgImage[identifier].Name = fielValue
				case "Base64Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EmbeddedJpgImage[identifier].Base64Content = fielValue
				}
			case "EmbeddedPngImage":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EmbeddedPngImage[identifier].Name = fielValue
				case "Base64Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EmbeddedPngImage[identifier].Base64Content = fielValue
				}
			case "EmbeddedSvgImage":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EmbeddedSvgImage[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EmbeddedSvgImage[identifier].Content = fielValue
				}
			case "Kill":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Kill[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry[identifier].Name = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry[identifier].Name = fielValue
				}
			case "Map_SPECIFICATION_Nodes_expandedEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_SPECIFICATION_Nodes_expandedEntry[identifier].Name = fielValue
				}
			case "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry[identifier].Name = fielValue
				}
			case "Map_SPEC_OBJECT_TYPE_showIdentifierEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_SPEC_OBJECT_TYPE_showIdentifierEntry[identifier].Name = fielValue
				}
			case "Map_SPEC_OBJECT_TYPE_showNameEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Map_SPEC_OBJECT_TYPE_showNameEntry[identifier].Name = fielValue
				}
			case "RELATION_GROUP":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP[identifier].LONG_NAME = fielValue
				}
			case "RELATION_GROUP_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP_TYPE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP_TYPE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP_TYPE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP_TYPE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP_TYPE[identifier].LONG_NAME = fielValue
				}
			case "REQ_IF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF[identifier].Name = fielValue
				case "Lang":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF[identifier].Lang = fielValue
				}
			case "REQ_IF_CONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_CONTENT[identifier].Name = fielValue
				}
			case "REQ_IF_HEADER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].Name = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].IDENTIFIER = fielValue
				case "COMMENT":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].COMMENT = fielValue
				case "CREATION_TIME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].CREATION_TIME = fielValue
				case "REPOSITORY_ID":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].REPOSITORY_ID = fielValue
				case "REQ_IF_TOOL_ID":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].REQ_IF_TOOL_ID = fielValue
				case "REQ_IF_VERSION":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].REQ_IF_VERSION = fielValue
				case "SOURCE_TOOL_ID":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].SOURCE_TOOL_ID = fielValue
				case "TITLE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].TITLE = fielValue
				}
			case "REQ_IF_TOOL_EXTENSION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_TOOL_EXTENSION[identifier].Name = fielValue
				}
			case "RenderingConfiguration":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RenderingConfiguration[identifier].Name = fielValue
				}
			case "SPECIFICATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].LONG_NAME = fielValue
				}
			case "SPECIFICATION_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION_TYPE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION_TYPE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION_TYPE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION_TYPE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION_TYPE[identifier].LONG_NAME = fielValue
				}
			case "SPEC_HIERARCHY":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_HIERARCHY[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_HIERARCHY[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_HIERARCHY[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_HIERARCHY[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_HIERARCHY[identifier].LONG_NAME = fielValue
				}
			case "SPEC_OBJECT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT[identifier].LONG_NAME = fielValue
				}
			case "SPEC_OBJECT_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT_TYPE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT_TYPE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT_TYPE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT_TYPE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT_TYPE[identifier].LONG_NAME = fielValue
				}
			case "SPEC_RELATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION[identifier].LONG_NAME = fielValue
				}
			case "SPEC_RELATION_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION_TYPE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION_TYPE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION_TYPE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION_TYPE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION_TYPE[identifier].LONG_NAME = fielValue
				}
			case "StaticWebSite":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSite[identifier].Name = fielValue
				case "MarkdownContent":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSite[identifier].MarkdownContent = fielValue
				case "InputImagesDir":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSite[identifier].InputImagesDir = fielValue
				case "OutputStaticWebDir":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSite[identifier].OutputStaticWebDir = fielValue
				case "VersionInfo":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSite[identifier].VersionInfo = fielValue
				}
			case "StaticWebSiteChapter":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSiteChapter[identifier].Name = fielValue
				case "MarkdownContent":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSiteChapter[identifier].MarkdownContent = fielValue
				}
			case "StaticWebSiteGeneratedImage":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSiteGeneratedImage[identifier].Name = fielValue
				case "SourceDirectoryPath":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSiteGeneratedImage[identifier].SourceDirectoryPath = fielValue
				case "Width":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_StaticWebSiteGeneratedImage[identifier].Width = int(exprSign) * int(fielValue)
				case "Height":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_StaticWebSiteGeneratedImage[identifier].Height = int(exprSign) * int(fielValue)
				}
			case "StaticWebSiteImage":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSiteImage[identifier].Name = fielValue
				case "SourceDirectoryPath":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSiteImage[identifier].SourceDirectoryPath = fielValue
				case "Width":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_StaticWebSiteImage[identifier].Width = int(exprSign) * int(fielValue)
				case "Height":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_StaticWebSiteImage[identifier].Height = int(exprSign) * int(fielValue)
				}
			case "StaticWebSiteParagraph":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSiteParagraph[identifier].Name = fielValue
				case "LegendMarkdownContent":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_StaticWebSiteParagraph[identifier].LegendMarkdownContent = fielValue
				}
			case "XHTML_CONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_XHTML_CONTENT[identifier].Name = fielValue
				case "EnclosedText":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_XHTML_CONTENT[identifier].EnclosedText = fielValue
				case "PureText":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_XHTML_CONTENT[identifier].PureText = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Println("gongstructName not found for identifier", identifier)
				break
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "ALTERNATIVE_ID":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ATTRIBUTE_DEFINITION_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].IS_EDITABLE = fielValue
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "DEFAULT_VALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].DEFAULT_VALUE = __gong__map_A_ATTRIBUTE_VALUE_BOOLEAN[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].TYPE = __gong__map_A_DATATYPE_DEFINITION_BOOLEAN_REF[targetIdentifier]
				}
			case "ATTRIBUTE_DEFINITION_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].IS_EDITABLE = fielValue
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "DEFAULT_VALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].DEFAULT_VALUE = __gong__map_A_ATTRIBUTE_VALUE_DATE[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].TYPE = __gong__map_A_DATATYPE_DEFINITION_DATE_REF[targetIdentifier]
				}
			case "ATTRIBUTE_DEFINITION_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].IS_EDITABLE = fielValue
				case "MULTI_VALUED":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].MULTI_VALUED = fielValue
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "DEFAULT_VALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].DEFAULT_VALUE = __gong__map_A_ATTRIBUTE_VALUE_ENUMERATION[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].TYPE = __gong__map_A_DATATYPE_DEFINITION_ENUMERATION_REF[targetIdentifier]
				}
			case "ATTRIBUTE_DEFINITION_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].IS_EDITABLE = fielValue
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "DEFAULT_VALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].DEFAULT_VALUE = __gong__map_A_ATTRIBUTE_VALUE_INTEGER[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].TYPE = __gong__map_A_DATATYPE_DEFINITION_INTEGER_REF[targetIdentifier]
				}
			case "ATTRIBUTE_DEFINITION_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].IS_EDITABLE = fielValue
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "DEFAULT_VALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].DEFAULT_VALUE = __gong__map_A_ATTRIBUTE_VALUE_REAL[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].TYPE = __gong__map_A_DATATYPE_DEFINITION_REAL_REF[targetIdentifier]
				}
			case "ATTRIBUTE_DEFINITION_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].IS_EDITABLE = fielValue
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "DEFAULT_VALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].DEFAULT_VALUE = __gong__map_A_ATTRIBUTE_VALUE_STRING[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].TYPE = __gong__map_A_DATATYPE_DEFINITION_STRING_REF[targetIdentifier]
				}
			case "ATTRIBUTE_DEFINITION_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].IS_EDITABLE = fielValue
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "DEFAULT_VALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].DEFAULT_VALUE = __gong__map_A_ATTRIBUTE_VALUE_XHTML[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].TYPE = __gong__map_A_DATATYPE_DEFINITION_XHTML_REF[targetIdentifier]
				}
			case "ATTRIBUTE_VALUE_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "THE_VALUE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_VALUE_BOOLEAN[identifier].THE_VALUE = fielValue
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_VALUE_BOOLEAN[identifier].DEFINITION = __gong__map_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF[targetIdentifier]
				}
			case "ATTRIBUTE_VALUE_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_VALUE_DATE[identifier].DEFINITION = __gong__map_A_ATTRIBUTE_DEFINITION_DATE_REF[targetIdentifier]
				}
			case "ATTRIBUTE_VALUE_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_VALUE_ENUMERATION[identifier].DEFINITION = __gong__map_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF[targetIdentifier]
				case "VALUES":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_VALUE_ENUMERATION[identifier].VALUES = __gong__map_A_ENUM_VALUE_REF[targetIdentifier]
				}
			case "ATTRIBUTE_VALUE_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_VALUE_INTEGER[identifier].DEFINITION = __gong__map_A_ATTRIBUTE_DEFINITION_INTEGER_REF[targetIdentifier]
				}
			case "ATTRIBUTE_VALUE_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_VALUE_REAL[identifier].DEFINITION = __gong__map_A_ATTRIBUTE_DEFINITION_REAL_REF[targetIdentifier]
				}
			case "ATTRIBUTE_VALUE_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_VALUE_STRING[identifier].DEFINITION = __gong__map_A_ATTRIBUTE_DEFINITION_STRING_REF[targetIdentifier]
				}
			case "ATTRIBUTE_VALUE_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_SIMPLIFIED":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].IS_SIMPLIFIED = fielValue
				case "THE_VALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].THE_VALUE = __gong__map_XHTML_CONTENT[targetIdentifier]
				case "THE_ORIGINAL_VALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].THE_ORIGINAL_VALUE = __gong__map_XHTML_CONTENT[targetIdentifier]
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].DEFINITION = __gong__map_A_ATTRIBUTE_DEFINITION_XHTML_REF[targetIdentifier]
				}
			case "A_ALTERNATIVE_ID":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_A_ALTERNATIVE_ID[identifier].ALTERNATIVE_ID = __gong__map_ALTERNATIVE_ID[targetIdentifier]
				}
			case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_DEFINITION_DATE_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_DEFINITION_REAL_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_DEFINITION_STRING_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_VALUE_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_VALUE_DATE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_VALUE_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_VALUE_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_VALUE_REAL":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_VALUE_STRING":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_VALUE_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ATTRIBUTE_VALUE_XHTML_1":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_CHILDREN":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_CORE_CONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "REQ_IF_CONTENT":
					targetIdentifier := ident.Name
					__gong__map_A_CORE_CONTENT[identifier].REQ_IF_CONTENT = __gong__map_REQ_IF_CONTENT[targetIdentifier]
				}
			case "A_DATATYPES":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_DATATYPE_DEFINITION_DATE_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_DATATYPE_DEFINITION_INTEGER_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_DATATYPE_DEFINITION_REAL_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_DATATYPE_DEFINITION_STRING_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_DATATYPE_DEFINITION_XHTML_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_EDITABLE_ATTS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_ENUM_VALUE_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_OBJECT":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_PROPERTIES":
				switch fieldName {
				// insertion point for field dependant code
				case "EMBEDDED_VALUE":
					targetIdentifier := ident.Name
					__gong__map_A_PROPERTIES[identifier].EMBEDDED_VALUE = __gong__map_EMBEDDED_VALUE[targetIdentifier]
				}
			case "A_RELATION_GROUP_TYPE_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SOURCE_1":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SOURCE_SPECIFICATION_1":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPECIFICATIONS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPECIFICATION_TYPE_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPECIFIED_VALUES":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPEC_ATTRIBUTES":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPEC_OBJECTS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPEC_OBJECT_TYPE_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPEC_RELATIONS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPEC_RELATION_GROUPS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPEC_RELATION_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPEC_RELATION_TYPE_REF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_SPEC_TYPES":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "A_THE_HEADER":
				switch fieldName {
				// insertion point for field dependant code
				case "REQ_IF_HEADER":
					targetIdentifier := ident.Name
					__gong__map_A_THE_HEADER[identifier].REQ_IF_HEADER = __gong__map_REQ_IF_HEADER[targetIdentifier]
				}
			case "A_TOOL_EXTENSIONS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "DATATYPE_DEFINITION_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				}
			case "DATATYPE_DEFINITION_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				}
			case "DATATYPE_DEFINITION_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "SPECIFIED_VALUES":
					targetIdentifier := ident.Name
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].SPECIFIED_VALUES = __gong__map_A_SPECIFIED_VALUES[targetIdentifier]
				}
			case "DATATYPE_DEFINITION_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				}
			case "DATATYPE_DEFINITION_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				}
			case "DATATYPE_DEFINITION_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				}
			case "DATATYPE_DEFINITION_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				}
			case "EMBEDDED_VALUE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ENUM_VALUE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_ENUM_VALUE[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "PROPERTIES":
					targetIdentifier := ident.Name
					__gong__map_ENUM_VALUE[identifier].PROPERTIES = __gong__map_A_PROPERTIES[targetIdentifier]
				}
			case "EmbeddedJpgImage":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "EmbeddedPngImage":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "EmbeddedSvgImage":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Kill":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry[identifier].Value = fielValue
				}
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry[identifier].Value = fielValue
				}
			case "Map_SPECIFICATION_Nodes_expandedEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_SPECIFICATION_Nodes_expandedEntry[identifier].Value = fielValue
				}
			case "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry[identifier].Value = fielValue
				}
			case "Map_SPEC_OBJECT_TYPE_showIdentifierEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_SPEC_OBJECT_TYPE_showIdentifierEntry[identifier].Value = fielValue
				}
			case "Map_SPEC_OBJECT_TYPE_showNameEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Map_SPEC_OBJECT_TYPE_showNameEntry[identifier].Value = fielValue
				}
			case "RELATION_GROUP":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_RELATION_GROUP[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "SOURCE_SPECIFICATION":
					targetIdentifier := ident.Name
					__gong__map_RELATION_GROUP[identifier].SOURCE_SPECIFICATION = __gong__map_A_SOURCE_SPECIFICATION_1[targetIdentifier]
				case "SPEC_RELATIONS":
					targetIdentifier := ident.Name
					__gong__map_RELATION_GROUP[identifier].SPEC_RELATIONS = __gong__map_A_SPEC_RELATION_REF[targetIdentifier]
				case "TARGET_SPECIFICATION":
					targetIdentifier := ident.Name
					__gong__map_RELATION_GROUP[identifier].TARGET_SPECIFICATION = __gong__map_A_SOURCE_SPECIFICATION_1[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_RELATION_GROUP[identifier].TYPE = __gong__map_A_RELATION_GROUP_TYPE_REF[targetIdentifier]
				}
			case "RELATION_GROUP_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_RELATION_GROUP_TYPE[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "SPEC_ATTRIBUTES":
					targetIdentifier := ident.Name
					__gong__map_RELATION_GROUP_TYPE[identifier].SPEC_ATTRIBUTES = __gong__map_A_SPEC_ATTRIBUTES[targetIdentifier]
				}
			case "REQ_IF":
				switch fieldName {
				// insertion point for field dependant code
				case "THE_HEADER":
					targetIdentifier := ident.Name
					__gong__map_REQ_IF[identifier].THE_HEADER = __gong__map_A_THE_HEADER[targetIdentifier]
				case "CORE_CONTENT":
					targetIdentifier := ident.Name
					__gong__map_REQ_IF[identifier].CORE_CONTENT = __gong__map_A_CORE_CONTENT[targetIdentifier]
				case "TOOL_EXTENSIONS":
					targetIdentifier := ident.Name
					__gong__map_REQ_IF[identifier].TOOL_EXTENSIONS = __gong__map_A_TOOL_EXTENSIONS[targetIdentifier]
				}
			case "REQ_IF_CONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "DATATYPES":
					targetIdentifier := ident.Name
					__gong__map_REQ_IF_CONTENT[identifier].DATATYPES = __gong__map_A_DATATYPES[targetIdentifier]
				case "SPEC_TYPES":
					targetIdentifier := ident.Name
					__gong__map_REQ_IF_CONTENT[identifier].SPEC_TYPES = __gong__map_A_SPEC_TYPES[targetIdentifier]
				case "SPEC_OBJECTS":
					targetIdentifier := ident.Name
					__gong__map_REQ_IF_CONTENT[identifier].SPEC_OBJECTS = __gong__map_A_SPEC_OBJECTS[targetIdentifier]
				case "SPEC_RELATIONS":
					targetIdentifier := ident.Name
					__gong__map_REQ_IF_CONTENT[identifier].SPEC_RELATIONS = __gong__map_A_SPEC_RELATIONS[targetIdentifier]
				case "SPECIFICATIONS":
					targetIdentifier := ident.Name
					__gong__map_REQ_IF_CONTENT[identifier].SPECIFICATIONS = __gong__map_A_SPECIFICATIONS[targetIdentifier]
				case "SPEC_RELATION_GROUPS":
					targetIdentifier := ident.Name
					__gong__map_REQ_IF_CONTENT[identifier].SPEC_RELATION_GROUPS = __gong__map_A_SPEC_RELATION_GROUPS[targetIdentifier]
				}
			case "REQ_IF_HEADER":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "REQ_IF_TOOL_EXTENSION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "RenderingConfiguration":
				switch fieldName {
				// insertion point for field dependant code
				case "ShowSpecHierachyIdentifiers":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RenderingConfiguration[identifier].ShowSpecHierachyIdentifiers = fielValue
				}
			case "SPECIFICATION":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATION[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "CHILDREN":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATION[identifier].CHILDREN = __gong__map_A_CHILDREN[targetIdentifier]
				case "VALUES":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATION[identifier].VALUES = __gong__map_A_ATTRIBUTE_VALUE_XHTML_1[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATION[identifier].TYPE = __gong__map_A_SPECIFICATION_TYPE_REF[targetIdentifier]
				}
			case "SPECIFICATION_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATION_TYPE[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "SPEC_ATTRIBUTES":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATION_TYPE[identifier].SPEC_ATTRIBUTES = __gong__map_A_SPEC_ATTRIBUTES[targetIdentifier]
				}
			case "SPEC_HIERARCHY":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SPEC_HIERARCHY[identifier].IS_EDITABLE = fielValue
				case "IS_TABLE_INTERNAL":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SPEC_HIERARCHY[identifier].IS_TABLE_INTERNAL = fielValue
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_SPEC_HIERARCHY[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "CHILDREN":
					targetIdentifier := ident.Name
					__gong__map_SPEC_HIERARCHY[identifier].CHILDREN = __gong__map_A_CHILDREN[targetIdentifier]
				case "EDITABLE_ATTS":
					targetIdentifier := ident.Name
					__gong__map_SPEC_HIERARCHY[identifier].EDITABLE_ATTS = __gong__map_A_EDITABLE_ATTS[targetIdentifier]
				case "OBJECT":
					targetIdentifier := ident.Name
					__gong__map_SPEC_HIERARCHY[identifier].OBJECT = __gong__map_A_OBJECT[targetIdentifier]
				}
			case "SPEC_OBJECT":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_SPEC_OBJECT[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "VALUES":
					targetIdentifier := ident.Name
					__gong__map_SPEC_OBJECT[identifier].VALUES = __gong__map_A_ATTRIBUTE_VALUE_XHTML_1[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_SPEC_OBJECT[identifier].TYPE = __gong__map_A_SPEC_OBJECT_TYPE_REF[targetIdentifier]
				}
			case "SPEC_OBJECT_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_SPEC_OBJECT_TYPE[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "SPEC_ATTRIBUTES":
					targetIdentifier := ident.Name
					__gong__map_SPEC_OBJECT_TYPE[identifier].SPEC_ATTRIBUTES = __gong__map_A_SPEC_ATTRIBUTES[targetIdentifier]
				}
			case "SPEC_RELATION":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_SPEC_RELATION[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "VALUES":
					targetIdentifier := ident.Name
					__gong__map_SPEC_RELATION[identifier].VALUES = __gong__map_A_ATTRIBUTE_VALUE_XHTML_1[targetIdentifier]
				case "SOURCE":
					targetIdentifier := ident.Name
					__gong__map_SPEC_RELATION[identifier].SOURCE = __gong__map_A_SOURCE_1[targetIdentifier]
				case "TARGET":
					targetIdentifier := ident.Name
					__gong__map_SPEC_RELATION[identifier].TARGET = __gong__map_A_SOURCE_1[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_SPEC_RELATION[identifier].TYPE = __gong__map_A_SPEC_RELATION_TYPE_REF[targetIdentifier]
				}
			case "SPEC_RELATION_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVE_ID":
					targetIdentifier := ident.Name
					__gong__map_SPEC_RELATION_TYPE[identifier].ALTERNATIVE_ID = __gong__map_A_ALTERNATIVE_ID[targetIdentifier]
				case "SPEC_ATTRIBUTES":
					targetIdentifier := ident.Name
					__gong__map_SPEC_RELATION_TYPE[identifier].SPEC_ATTRIBUTES = __gong__map_A_SPEC_ATTRIBUTES[targetIdentifier]
				}
			case "StaticWebSite":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "StaticWebSiteChapter":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "StaticWebSiteGeneratedImage":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "StaticWebSiteImage":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "StaticWebSiteParagraph":
				switch fieldName {
				// insertion point for field dependant code
				case "Image":
					targetIdentifier := ident.Name
					__gong__map_StaticWebSiteParagraph[identifier].Image = __gong__map_StaticWebSiteImage[targetIdentifier]
				}
			case "XHTML_CONTENT":
				switch fieldName {
				// insertion point for field dependant code
				}
			}
		case *ast.SelectorExpr:
			var basicLit *ast.BasicLit
			var ident *ast.Ident

			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			case *ast.CompositeLit:
				var ok bool
				var sl *ast.SelectorExpr

				if sl, ok = X.Type.(*ast.SelectorExpr); !ok {
					break // Exits the switch case
				}

				if ident, ok = sl.X.(*ast.Ident); !ok {
					break // Exits the switch case
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = ident.Name + "." + sl.Sel.Name + "{}." + selectorExpr.Sel.Name
			}

			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Println("gongstructName not found for identifier", identifier)
					break
				}

				if basicLit == nil {
					// for the meta field written as ref_models.ENUM_VALUE1
					basicLit = new(ast.BasicLit)
					basicLit.Kind = token.STRING // Or another appropriate token.Kind
					basicLit.Value = selectorExpr.X.(*ast.Ident).Name + "." + Sel.Name
					_ = basicLit.Kind
					_ = basicLit.Value
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for selector expr assignments
				case "ALTERNATIVE_ID":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_DEFINITION_BOOLEAN":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_DEFINITION_DATE":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_DEFINITION_ENUMERATION":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_DEFINITION_INTEGER":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_DEFINITION_REAL":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_DEFINITION_STRING":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_DEFINITION_XHTML":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_VALUE_BOOLEAN":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_VALUE_DATE":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_VALUE_ENUMERATION":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_VALUE_INTEGER":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_VALUE_REAL":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_VALUE_STRING":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ATTRIBUTE_VALUE_XHTML":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ALTERNATIVE_ID":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_DEFINITION_DATE_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_DEFINITION_REAL_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_DEFINITION_STRING_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_VALUE_BOOLEAN":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_VALUE_DATE":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_VALUE_ENUMERATION":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_VALUE_INTEGER":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_VALUE_REAL":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_VALUE_STRING":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_VALUE_XHTML":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ATTRIBUTE_VALUE_XHTML_1":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_CHILDREN":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_CORE_CONTENT":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_DATATYPES":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_DATATYPE_DEFINITION_DATE_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_DATATYPE_DEFINITION_INTEGER_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_DATATYPE_DEFINITION_REAL_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_DATATYPE_DEFINITION_STRING_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_DATATYPE_DEFINITION_XHTML_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_EDITABLE_ATTS":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_ENUM_VALUE_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_OBJECT":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_PROPERTIES":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_RELATION_GROUP_TYPE_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SOURCE_1":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SOURCE_SPECIFICATION_1":
					switch fieldName {
					// insertion point for selector expr assign code
					case "SPECIFICATION_REF":
						var val Enum_GLOBAL_REF
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_A_SOURCE_SPECIFICATION_1[identifier].SPECIFICATION_REF = Enum_GLOBAL_REF(val)
					}
				case "A_SPECIFICATIONS":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SPECIFICATION_TYPE_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SPECIFIED_VALUES":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SPEC_ATTRIBUTES":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SPEC_OBJECTS":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SPEC_OBJECT_TYPE_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SPEC_RELATIONS":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SPEC_RELATION_GROUPS":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SPEC_RELATION_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SPEC_RELATION_TYPE_REF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_SPEC_TYPES":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_THE_HEADER":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "A_TOOL_EXTENSIONS":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "DATATYPE_DEFINITION_BOOLEAN":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "DATATYPE_DEFINITION_DATE":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "DATATYPE_DEFINITION_ENUMERATION":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "DATATYPE_DEFINITION_INTEGER":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "DATATYPE_DEFINITION_REAL":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "DATATYPE_DEFINITION_STRING":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "DATATYPE_DEFINITION_XHTML":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "EMBEDDED_VALUE":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ENUM_VALUE":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "EmbeddedJpgImage":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "EmbeddedPngImage":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "EmbeddedSvgImage":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Kill":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_SPECIFICATION_Nodes_expandedEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_SPEC_OBJECT_TYPE_showIdentifierEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Map_SPEC_OBJECT_TYPE_showNameEntry":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "RELATION_GROUP":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "RELATION_GROUP_TYPE":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "REQ_IF":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "REQ_IF_CONTENT":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "REQ_IF_HEADER":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "REQ_IF_TOOL_EXTENSION":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "RenderingConfiguration":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SPECIFICATION":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SPECIFICATION_TYPE":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SPEC_HIERARCHY":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SPEC_OBJECT":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SPEC_OBJECT_TYPE":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SPEC_RELATION":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SPEC_RELATION_TYPE":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "StaticWebSite":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "StaticWebSiteChapter":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "StaticWebSiteGeneratedImage":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "StaticWebSiteImage":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "StaticWebSiteParagraph":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "XHTML_CONTENT":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				}
			}
		}
	}
	return
}

// ReplaceOldDeclarationsInFile replaces specific text in a file at the given path.
func ReplaceOldDeclarationsInFile(pathToFile string) error {
	// Open the file for reading.
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// replacing function with Injection
	pattern := regexp.MustCompile(`\b\w*Injection\b`)
	pattern2 := regexp.MustCompile(`\bmap_DocLink_Identifier_\w*\b`)

	// Temporary slice to hold lines from the file.
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Replace the target text with the desired text.
		line := strings.Replace(scanner.Text(), "var ___dummy__Time_stage time.Time", "var _ time.Time", -1)
		line = pattern.ReplaceAllString(line, "_")
		line = pattern2.ReplaceAllString(line, "_")

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Re-open the file for writing.
	file, err = os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the modified lines back to the file.
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
