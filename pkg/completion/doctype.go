package completion

import (
	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/opa-oz/pug-lsp/pkg/query"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func DoctypeCompletion(meta *CompletionMetaParams, completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	if meta.Doc.HasDoctype {
		return &completionItems
	}

	valueKind := protocol.CompletionItemKindSnippet

	if meta.CurrentNode != nil && query.HasAttributesAncestor(meta.CurrentNode) {
		return &completionItems
	}

	for i, doctype := range html.Doctypes {
		preselect := i == 0
		dCopy := doctype
		completionItems = append(completionItems, protocol.CompletionItem{
			Preselect:  &preselect,
			Label:      dCopy,
			Kind:       &valueKind,
			Detail:     &dCopy,
			InsertText: &dCopy,
		})
	}

	return &completionItems
}
