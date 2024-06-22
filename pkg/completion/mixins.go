package completion

import (
	"github.com/opa-oz/pug-lsp/pkg/query"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// docMixins generates completion items for mixins defined in a document and included docs.
//
// It iterates through the mixins and creates
// completion items based on each mixin's name and optional arguments. If a mixin's name
// matches the `exclude` string, it skips creating a completion item for that mixin.
//
// Parameters:
//
//	mixins *[]*query.Mixin - Flat list opf all mixins.
//	completionItems []protocol.CompletionItem - Existing list of completion items to append to.
//	exclude string - The name of a mixin to exclude from completion items.
//
// Returns:
//
//	*[]protocol.CompletionItem - A pointer to a slice of protocol.CompletionItem containing
//	                             completion items for mixins. Returns nil if the document
//	                             or completion items slice is nil.
func docMixins(mixins *[]*query.Mixin, completionItems []protocol.CompletionItem, exclude string) *[]protocol.CompletionItem {
	functionKind := protocol.CompletionItemKindFunction

	for _, def := range *mixins {
		insert := def.Name
		if insert == exclude {
			continue
		}

		if len(*def.Arguments) > 0 {
			insert += "()"
		}

		item := protocol.CompletionItem{
			Label:      def.Name,
			Kind:       &functionKind,
			InsertText: &insert,
			Data:       DataToString(MixinKind, def.Definition),
		}

		completionItems = append(completionItems, item)
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

	mixins := meta.DocumentStore.FlatMixins(meta.Doc)
	completionItems = *docMixins(mixins, completionItems, excludeMixin)

	return &completionItems
}
