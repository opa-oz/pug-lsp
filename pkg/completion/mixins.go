package completion

import (
	"github.com/opa-oz/pug-lsp/pkg/documents"
	"github.com/opa-oz/pug-lsp/pkg/query"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// docMixins generates completion items for mixins defined in a document.
//
// It iterates through the mixins defined in the provided document (`doc`) and creates
// completion items based on each mixin's name and optional arguments. If a mixin's name
// matches the `exclude` string, it skips creating a completion item for that mixin.
//
// Parameters:
//
//	doc *documents.Document - The document containing mixin definitions.
//	completionItems []protocol.CompletionItem - Existing list of completion items to append to.
//	exclude string - The name of a mixin to exclude from completion items.
//
// Returns:
//
//	*[]protocol.CompletionItem - A pointer to a slice of protocol.CompletionItem containing
//	                             completion items for mixins. Returns nil if the document
//	                             or completion items slice is nil.
func docMixins(doc *documents.Document, completionItems []protocol.CompletionItem, exclude string) *[]protocol.CompletionItem {
	functionKind := protocol.CompletionItemKindFunction

	for _, def := range doc.Mixins {
		insert := def.Name
		if insert == exclude {
			continue
		}

		if len(*def.Arguments) > 0 {
			insert += "()"
		}

		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      def.Name,
			Kind:       &functionKind,
			Detail:     &def.Definition,
			InsertText: &insert,
		})
	}

	return &completionItems
}

func MixinsCompletion(meta *CompletionMetaParams, completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	hasMixinDefAncestor := query.HasMixinDefinitionAncestor(meta.CurrentNode)
	var excludeMixin string

	if hasMixinDefAncestor {
		mixinDef := query.FindUpwards(meta.CurrentNode, query.MixinDefinitionNode, -1)
		if mixinDef != nil {
			mixinName := query.FindDownwards(mixinDef, query.MixinNameNode, 2)
			if mixinName != nil {
				excludeMixin = (*meta.Doc.Content)[mixinName.StartByte():mixinName.EndByte()]
			}
		}
	}

	completionItems = *docMixins(meta.Doc, completionItems, excludeMixin)
	for _, include := range meta.Doc.Includes {
		doc, ok := meta.DocumentStore.Get(*include.URI)

		if !ok {
			(*meta.Logger).Println("Something shady with include", include.Path)
			continue
		}

		completionItems = *docMixins(doc, completionItems, excludeMixin)
	}

	return &completionItems
}
