package server

import (
	"github.com/opa-oz/go-todo/todo"
	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *Server) TextDocumentCompletion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completionItems []protocol.CompletionItem

	doc, ok := s.documentStore.Get(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	s.logger.Println("Trying to completion", doc.Path)
	valueKind := protocol.CompletionItemKindValue

	if params.Context != nil && params.Context.TriggerKind == protocol.CompletionTriggerKindTriggerCharacter {
		switch *params.Context.TriggerCharacter {
		case "(":
			todo.T("Append tag-specific attributes")
			globalAttrs := html.GlobalAttrs()

			for _, attr := range *globalAttrs {
				attrCopy := attr // Create a copy of attr
				attrToInsert := attr + "="
				completionItems = append(completionItems, protocol.CompletionItem{
					Label:      attrCopy + " - [Attr]",
					Kind:       &valueKind,
					Detail:     &attrCopy,
					InsertText: &attrToInsert,
				})
			}
		case "+":
			todo.T("Grab mixins from lined files and offer")
			defaultSymbol := "mixins"
			defaultToInsert := "mixins."

			completionItems = append(completionItems, protocol.CompletionItem{
				Label:      todo.String("Look for mixins", "Default mixin - [Mixin]"),
				Kind:       &valueKind,
				Detail:     &defaultSymbol,
				InsertText: &defaultToInsert,
			})
		}
	} else {
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
	}

	return completionItems, nil
}
