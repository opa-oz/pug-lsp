package completion

import (
	"github.com/opa-oz/pug-lsp/pkg/documents"
	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/opa-oz/pug-lsp/pkg/utils"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func globalTags(completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	valueKind := protocol.CompletionItemKindConstant
	htmlTags := html.GetTags()

	for _, tag := range *htmlTags {
		tagCopy := tag
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      tagCopy + " - [HTML]",
			Kind:       &valueKind,
			Detail:     &tagCopy,
			InsertText: &tagCopy,
		})
	}

	return &completionItems
}

func PlainCompletion(doc *documents.Document, completionItems []protocol.CompletionItem, params *protocol.CompletionParams, logger *utils.Logger) *[]protocol.CompletionItem {
	node := doc.GetAtPosition(&params.Position)
	if node == nil {
		return globalTags(completionItems)
	}

	tag := query.FindUpwards(node, query.TagNode, 1)
	if tag != nil {
		return globalTags(completionItems)
	}

	attributesEl := query.FindUpwards(node, query.AttributesNode, 1)
	if attributesEl != nil {
		existingAttrs := query.GetExistingAttributes(node, *doc.Content)
		items := LocalCompletion(doc, completionItems, params, existingAttrs)
		items = GlobalCompletion(*items, existingAttrs)

		return items
	}

	return &completionItems
}
