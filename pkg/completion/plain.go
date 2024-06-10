package completion

import (
	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/opa-oz/pug-lsp/pkg/query"
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

func PlainCompletion(meta *CompletionMetaParams, completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	node := meta.CurrentNode
	if node == nil {
		return globalTags(completionItems)
	}

	hasAttrsAncestor := query.HasAttributesAncestor(node)
	if hasAttrsAncestor {
		items := LocalCompletion(meta, completionItems)
		items = GlobalCompletion(meta, *items)

		return items
	}

	tag := query.FindUpwards(node, query.TagNode, 1)
	if tag != nil {
		return globalTags(completionItems)
	}

	return &completionItems
}
