package server

import (
	"github.com/opa-oz/pug-lsp/pkg/completion"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *Server) TextDocumentCompletion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completionItems []protocol.CompletionItem
	var completionList protocol.CompletionList

	doc, ok := s.documentStore.Get(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	if params.Context != nil && params.Context.TriggerKind == protocol.CompletionTriggerKindTriggerCharacter {
		switch *params.Context.TriggerCharacter {
		case "(":
			completionItems = *completion.LocalCompletion(doc, completionItems, params, true)
			completionItems = *completion.GlobalCompletion(completionItems)
		case ",":
			completionItems = *completion.LocalCompletion(doc, completionItems, params, false)
			completionItems = *completion.GlobalCompletion(completionItems)
		case "+":
			completionItems = *completion.MixinsCompletion(completionItems)
		}
	} else {
		completionItems = *completion.PlainCompletion(completionItems)
	}

	completionList.IsIncomplete = false
	completionList.Items = completionItems

	return completionList, nil
}
