package specifications

import (

	// Corrected path
	m "github.com/fullstack-lang/gongreqif/go/models"

	markdown "github.com/fullstack-lang/gong/lib/markdown/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

// UpdateAndCommitSpecificationsMarkdownStage implements models.SpecificationsTreeUpdaterInterface.
func (o *SpecificationsTreeStageUpdater) UpdateAndCommitSpecificationsMarkdownStage(stager *m.Stager) {

	markdownStage := stager.GetMarkdownStage()
	markdownStage.Reset()

	if stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPECIFICATIONS == nil {
		markdownStage.Commit()
		return
	}

	specifications := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPECIFICATIONS.SPECIFICATION
	for _, specification := range specifications {

		if stager.GetSelectedSpecification() == nil {
			continue
		}

		if stager.GetSelectedSpecification() != specification {
			continue
		}

		// --- updated logic to generate and assign markdown content ---

		// 1. Initialize markdown content string
		markDownContent := "# " + specification.Name + "\n\n"

		// 2. A dummy parent node is created because processSpecHierarchy expects a parent
		// to append children to. This node is temporary and will be discarded.
		hierarchyParentNode := &tree.Node{}
		depth := 1 // initial depth for chapters

		// 3. Recursively process spec hierarchies to build the markdown string
		for _, specHierarchy := range specification.CHILDREN.SPEC_HIERARCHY {
			processSpecHierarchy(
				stager,
				specHierarchy,
				hierarchyParentNode,
				depth,
				&markDownContent)
		}
		// --- end of update ---

		content := &markdown.Content{
			Name:    "*" + specification.GetName() + "*",
			Content: markDownContent, // Assign the generated markdown
		}

		markdown.StageBranch(markdownStage, content)
	}

	markdownStage.Commit()
}
