package completion

import (
	"github.com/opa-oz/pug-lsp/pkg/documents"
	"github.com/opa-oz/pug-lsp/pkg/html"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func DoctypeCompletion(doc *documents.Document, completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	if doc.HasDoctype {
		return &completionItems
	}

	valueKind := protocol.CompletionItemKindSnippet

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
