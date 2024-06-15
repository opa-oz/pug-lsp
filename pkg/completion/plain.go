package completion

import (
	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/opa-oz/pug-lsp/pkg/lsp"
	"github.com/opa-oz/pug-lsp/pkg/query"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func keywords(completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	keywordKind := protocol.CompletionItemKindKeyword
	keytags := lsp.GetKeywords()

	for _, key := range *keytags {
		keyName := string(key)
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      keyName,
			Kind:       &keywordKind,
			Detail:     &keyName,
			InsertText: &keyName,
		})
	}

	return &completionItems
}

func globalTags(completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	valueKind := protocol.CompletionItemKindConstant
	htmlTags := html.GetTags()

	for _, tag := range *htmlTags {
		tagCopy := tag
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      tagCopy,
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
		items := globalTags(completionItems)
		items = keywords(*items)

		return items
	}

	hasAttrsAncestor := query.HasAttributesAncestor(node)
	if hasAttrsAncestor {
		items := LocalCompletion(meta, completionItems)
		items = GlobalCompletion(meta, *items)

		return items
	}

	tag := query.FindUpwards(node, query.TagNode, 1)
	if tag != nil {
		items := globalTags(completionItems)
		items = keywords(*items)

		return items
	}

	return &completionItems
}
