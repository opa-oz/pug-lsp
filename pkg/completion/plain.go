package completion

import (
	"github.com/opa-oz/pug-lsp/pkg/html"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func PlainCompletion(completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	valueKind := protocol.CompletionItemKindValue
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
