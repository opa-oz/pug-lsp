package completion

import (
	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/opa-oz/pug-lsp/pkg/query"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// DoctypeCompletion generates completion items for document type (doctype) declarations
// based on the provided completion metadata and existing completion items.
//
// If the document (`meta.Doc`) already has a doctype declaration (`meta.Doc.HasDoctype` is true),
// the function returns a pointer to the existing completion items slice without modification.
//
// If the current node (`meta.CurrentNode`) is associated with attributes (checked via
// query.HasAttributesAncestor(meta.CurrentNode)), the function also returns a pointer to
// the existing completion items slice without modification.
//
// Otherwise, the function generates completion items for commonly used doctype declarations
// from `html.Doctypes`, adding them to the `completionItems` slice. Each generated completion
// item includes properties such as `Preselect` (to indicate the first item), `Label` (doctype name),
// `Kind` (set to `protocol.CompletionItemKindSnippet`), `Detail` (doctype name), and `InsertText`
// (text to insert into the document).
//
// Parameters:
//
//	meta *CompletionMetaParams - Metadata parameters for completion, including document information and current node.
//	completionItems []protocol.CompletionItem - Existing list of completion items to which new items will be appended.
//
// Returns:
//
//	*[]protocol.CompletionItem - A pointer to a slice of protocol.CompletionItem containing completion items
//	                             for doctype declarations. Returns nil if no doctype completion items are added.
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
