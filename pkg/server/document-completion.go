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
	node := doc.GetAtPosition(&params.Position)
	existingAttrs := query.GetExistingAttributes(node, *doc.Content)
	meta := completion.CompletionMetaParams{
		Doc:           doc,
		Params:        params,
		CurrentNode:   node,
		ExistingAttrs: existingAttrs,
		Logger:        &s.logger,
	}

	if node != nil && query.HasJSAncestor(node) {
		return nil, nil
	}

	if params.Context != nil && params.Context.TriggerKind == protocol.CompletionTriggerKindTriggerCharacter {
		switch *params.Context.TriggerCharacter {
		case "(":
			completionItems = *completion.LocalCompletion(&meta, completionItems)
			completionItems = *completion.GlobalCompletion(&meta, completionItems)
		case ",":
			completionItems = *completion.LocalCompletion(&meta, completionItems)
			completionItems = *completion.GlobalCompletion(&meta, completionItems)
		case "+":
			completionItems = *completion.MixinsCompletion(&meta, completionItems)
		case "&":
			completionItems = *completion.AndCompletion(&meta, completionItems)
		}
	} else {
		completionItems = *completion.DoctypeCompletion(&meta, completionItems)
		completionItems = *completion.PlainCompletion(&meta, completionItems)
	}

	if completionItems == nil {
		return nil, nil
	}

	completionList.IsIncomplete = false
	completionList.Items = completionItems

	return completionList, nil
}
