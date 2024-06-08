package server

import (
	"github.com/opa-oz/pug-lsp/pkg/completion"
	"github.com/opa-oz/pug-lsp/pkg/query"
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
			completionItems = *completion.LocalCompletion(doc, completionItems, params, nil)
			completionItems = *completion.GlobalCompletion(completionItems, nil)
		case ",":
			existingAttrs := query.GetExistingAttributes(doc.GetAtPosition(&params.Position), *doc.Content)
			completionItems = *completion.LocalCompletion(doc, completionItems, params, existingAttrs)
			completionItems = *completion.GlobalCompletion(completionItems, existingAttrs)
		case " ":
			existingAttrs := query.GetExistingAttributes(doc.GetAtPosition(&params.Position), *doc.Content)
			completionItems = *completion.LocalCompletion(doc, completionItems, params, existingAttrs)
			completionItems = *completion.GlobalCompletion(completionItems, existingAttrs)
		case "+":
			completionItems = *completion.MixinsCompletion(completionItems)
		case "&":
			completionItems = *completion.AndCompletion(doc, completionItems, params)
		}
	} else {
		completionItems = *completion.PlainCompletion(doc, completionItems, params, &s.logger)
	}

	if completionItems == nil {
		return nil, nil
	}

	completionList.IsIncomplete = false
	completionList.Items = completionItems

	return completionList, nil
}
