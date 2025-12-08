package spectypes

import (
	"fmt"
	"log"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gongreqif/go/icons"
	m "github.com/fullstack-lang/gongreqif/go/models"
)

type SpecTypesTreeStageUpdater struct {
}

func (updater *SpecTypesTreeStageUpdater) UpdateAndCommitSpecTypesTreeStage(stager *m.Stager) {

	stager.GetSpecTypesTreeStage().Reset()
	stage := stager.GetStage()

	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES

	if spectypes == nil {
		return
	}

	rootNode := &tree.Node{
		Name:       "Spec types",
		IsExpanded: true,
		FontStyle:  tree.ITALIC,
	}

	{
		spectypeCategory := &tree.Node{
			Name:       "Spec Object Types",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, spectypeCategory)

		// compute the number of time this spec object type is used
		map_specType_nbInstance := make(map[*m.SPEC_OBJECT_TYPE]int)
		for x := range *m.GetGongstructInstancesSet[m.A_SPEC_OBJECT_TYPE_REF](stager.GetStage()) {

			specObjecType, ok := stager.Map_id_SPEC_OBJECT_TYPE[x.SPEC_OBJECT_TYPE_REF]
			if !ok {
				log.Panicf("Object %s has a x.SPEC_OBJECT_TYPE_REF %s which is not known", x.Name, x.SPEC_OBJECT_TYPE_REF)
			} else {
				map_specType_nbInstance[specObjecType]++
			}
		}

		for _, specObjectType := range spectypes.SPEC_OBJECT_TYPE {
			specObjectTypeRendering := GetSpecObjectTypeRendering(stage, specObjectType)

			specObjectTypeNode := &tree.Node{
				Name: specObjectType.Name + fmt.Sprintf(" (%d/%d)",
					stager.Map_SPEC_OBJECT_TYPE_Spec_nbInstance[specObjectType],
					map_specType_nbInstance[specObjectType]),
				IsExpanded: specObjectTypeRendering.IsNodeExpanded,
				Impl: &SpecObjectTypeNodeProxy{
					stager:         stager,
					specObjectType: specObjectType,
				},
			}
			spectypeCategory.Children = append(spectypeCategory.Children, specObjectTypeNode)

			{
				button := &tree.Button{
					Name: "Show/Unshow identifier",
					Impl: &toggleButtonProxy{
						stager:      stager,
						toggleValue: &specObjectTypeRendering.ShowIdentifier,
					},
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if !specObjectTypeRendering.ShowIdentifier {
					button.ToolTipText = "Show identifier in title"
					button.SVGIcon = icons.SvgIconBadge
				} else {
					button.ToolTipText = "Hide identifier in title"
					button.SVGIcon = icons.SvgIconBadgeOff
				}
				specObjectTypeNode.Buttons = append(specObjectTypeNode.Buttons, button)
			}

			{
				button := &tree.Button{
					Name: "Show/Unshow name",
					Impl: &toggleButtonProxy{
						stager:      stager,
						toggleValue: &specObjectTypeRendering.ShowName,
					},
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if !specObjectTypeRendering.ShowName {
					button.ToolTipText = "Show name in title"
					button.SVGIcon = icons.SvgIconBadge
				} else {
					button.ToolTipText = "Hide name in title"
					button.SVGIcon = icons.SvgIconBadgeOff
				}
				specObjectTypeNode.Buttons = append(specObjectTypeNode.Buttons, button)
			}

			{
				button := &tree.Button{
					Name: "Show/Unshow relation",
					Impl: &toggleButtonProxy{
						stager:      stager,
						toggleValue: &specObjectTypeRendering.ShowRelations,
					},
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if !specObjectTypeRendering.ShowRelations {
					button.ToolTipText = "Show relations in object table"
					button.SVGIcon = icons.SvgIconTable
				} else {
					button.ToolTipText = "Hide relations in object table"
					button.SVGIcon = icons.SvgIconTableOff
				}
				specObjectTypeNode.Buttons = append(specObjectTypeNode.Buttons, button)
			}

			{
				button := &tree.Button{
					Name: "Is Spec Object Type Heading",
					Impl: &toggleButtonProxy{
						stager:      stager,
						toggleValue: &specObjectTypeRendering.IsHeading,
					},
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if !specObjectTypeRendering.IsHeading {
					button.ToolTipText = "Treat as heading"
					button.SVGIcon = icons.SvgIconTitle
				} else {
					button.ToolTipText = "Treat as no heading"
					button.SVGIcon = icons.SvgIconTitleOff
				}
				specObjectTypeNode.Buttons = append(specObjectTypeNode.Buttons, button)
			}

			addAttibutesNodes(
				stager,
				specObjectTypeNode,
				specObjectType.SPEC_ATTRIBUTES)

		}
	}

	{
		spectypeCategory := &tree.Node{
			Name:       "Relation Types",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}
		rootNode.Children = append(rootNode.Children, spectypeCategory)

		// compute the number of time this spec object type is used
		map_specType_nbInstance := make(map[*m.SPEC_RELATION_TYPE]int)
		for x := range *m.GetGongstructInstancesSet[m.A_SPEC_RELATION_TYPE_REF](stager.GetStage()) {

			datatypeDefinition, ok := stager.Map_id_SPEC_RELATION_TYPE[x.SPEC_RELATION_TYPE_REF]
			if !ok {
				log.Panic("x.SPEC_RELATION_TYPE_REF", x.SPEC_RELATION_TYPE_REF,
					"unknown ref")
			} else {
				map_specType_nbInstance[datatypeDefinition]++
			}
		}

		for _, spectype := range spectypes.SPEC_RELATION_TYPE {
			node := &tree.Node{
				Name:       spectype.Name + fmt.Sprintf(" (%d)", map_specType_nbInstance[spectype]),
				IsExpanded: false,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, node)

			addAttibutesNodes(
				stager,
				node,
				spectype.SPEC_ATTRIBUTES)
		}
	}

	{
		spectypeCategory := &tree.Node{
			Name:       "Specification Types",
			IsExpanded: true,
			FontStyle:  tree.ITALIC,
		}

		// compute the number of time this spec object type is used
		map_specType_nbInstance := make(map[*m.SPECIFICATION_TYPE]int)
		for x := range *m.GetGongstructInstancesSet[m.A_SPECIFICATION_TYPE_REF](stager.GetStage()) {

			specificationType, ok := stager.Map_id_SPECIFICATION_TYPE[x.SPECIFICATION_TYPE_REF]
			if !ok {
				log.Panic("x.SPECIFICATION_TYPE_REF: x: ", x.Name, " ,Type Ref: ", x.SPECIFICATION_TYPE_REF,
					", unknown ref")
			} else {
				map_specType_nbInstance[specificationType]++
			}
		}

		for _, spectype := range spectypes.SPECIFICATION_TYPE {
			node := &tree.Node{
				Name: spectype.Name + fmt.Sprintf(" (%d/%d)",
					stager.Map_SPECIFICATION_TYPE_Spec_nbInstance[spectype],
					map_specType_nbInstance[spectype]),
				IsExpanded: false,
			}
			spectypeCategory.Children = append(spectypeCategory.Children, node)

			if spectype.SPEC_ATTRIBUTES == nil {
				continue
			}
			addAttibutesNodes(
				stager,
				node,
				spectype.SPEC_ATTRIBUTES)
		}
	}

	tree.StageBranch(stager.GetSpecTypesTreeStage(),
		&tree.Tree{
			RootNodes: []*tree.Node{
				rootNode,
			},
		},
	)

	stager.GetSpecTypesTreeStage().Commit()
}

// addAttibutesNodes creates category nodes for attribute definitions
// only if there is more than one attribute definition for that category.
func addAttibutesNodes(
	stager *m.Stager,
	nodeSpecType *tree.Node,
	specAttributes *m.A_SPEC_ATTRIBUTES) {

	if specAttributes == nil {
		return
	}

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_XHTML,
		stager.Map_id_ATTRIBUTE_DEFINITION_XHTML,
		stager.Map_ATTRIBUTE_DEFINITION_XHTML_Spec_nbInstance,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_XHTML_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_XHTML,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_STRING,
		stager.Map_id_ATTRIBUTE_DEFINITION_STRING,
		stager.Map_ATTRIBUTE_DEFINITION_STRING_Spec_nbInstance,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_STRING_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_STRING,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_BOOLEAN,
		stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN,
		stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_Spec_nbInstance,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_BOOLEAN,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_INTEGER,
		stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER,
		stager.Map_ATTRIBUTE_DEFINITION_INTEGER_Spec_nbInstance,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_INTEGER_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_INTEGER,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_REAL,
		stager.Map_id_ATTRIBUTE_DEFINITION_REAL,
		stager.Map_ATTRIBUTE_DEFINITION_REAL_Spec_nbInstance,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_REAL_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_REAL,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_DATE,
		stager.Map_id_ATTRIBUTE_DEFINITION_DATE,
		stager.Map_ATTRIBUTE_DEFINITION_DATE_Spec_nbInstance,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_DATE_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_DATE,
		nodeSpecType)

	addAttributeNode(
		stager,
		specAttributes.ATTRIBUTE_DEFINITION_ENUMERATION,
		stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION,
		stager.Map_ATTRIBUTE_DEFINITION_ENUMERATION_Spec_nbInstance,
		*m.GetGongstructInstancesSetFromPointerType[*m.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](stager.GetStage()),
		stager.Map_id_DATATYPE_DEFINITION_ENUMERATION,
		nodeSpecType)
}

type ButtonToggleShowAttributeFieldInTitleProxy[AttrDef m.AttributeDefinition] struct {
	stager              *m.Stager
	attributeDefinition AttrDef
}

func (proxy *ButtonToggleShowAttributeFieldInTitleProxy[AttrDef]) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	switch ad := any(proxy.attributeDefinition).(type) {
	case *m.ATTRIBUTE_DEFINITION_XHTML:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_STRING:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_STRING_ShowInTitle(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_STRING_ShowInTitle(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_BOOLEAN:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_INTEGER:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_DATE:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_DATE_ShowInTitle(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_DATE_ShowInTitle(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_REAL:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_REAL_ShowInTitle(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_REAL_ShowInTitle(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_ENUMERATION:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle(ad, !currentValue)
	default:
		log.Fatalln("Unknown attribute definition type", ad)
	}

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
}

type ButtonToggleShowAttributeFieldInTableProxy[AttrDef m.AttributeDefinition] struct {
	stager              *m.Stager
	attributeDefinition AttrDef
}

func (proxy *ButtonToggleShowAttributeFieldInTableProxy[AttrDef]) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	switch ad := any(proxy.attributeDefinition).(type) {
	case *m.ATTRIBUTE_DEFINITION_XHTML:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_XHTML_ShowInTable(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_XHTML_ShowInTable(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_STRING:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_STRING_ShowInTable(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_STRING_ShowInTable(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_BOOLEAN:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTable(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_INTEGER:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_INTEGER_ShowInTable(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_DATE:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_DATE_ShowInTable(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_DATE_ShowInTable(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_REAL:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_REAL_ShowInTable(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_REAL_ShowInTable(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_ENUMERATION:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTable(ad, !currentValue)
	default:
		log.Fatalln("Unknown attribute definition type", ad)
	}

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
}

type ButtonToggleShowAttributeFieldInSubjectProxy[AttrDef m.AttributeDefinition] struct {
	stager              *m.Stager
	attributeDefinition AttrDef
}

func (proxy *ButtonToggleShowAttributeFieldInSubjectProxy[AttrDef]) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	switch ad := any(proxy.attributeDefinition).(type) {
	case *m.ATTRIBUTE_DEFINITION_XHTML:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_STRING:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_STRING_ShowInSubject(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_STRING_ShowInSubject(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_BOOLEAN:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_INTEGER:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_DATE:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_DATE_ShowInSubject(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_DATE_ShowInSubject(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_REAL:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_REAL_ShowInSubject(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_REAL_ShowInSubject(ad, !currentValue)
	case *m.ATTRIBUTE_DEFINITION_ENUMERATION:
		currentValue := proxy.stager.RenderingConf.Get_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject(ad)
		proxy.stager.RenderingConf.Set_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject(ad, !currentValue)
	default:
		log.Fatalln("Unknown attribute definition type", ad)
	}

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
}
