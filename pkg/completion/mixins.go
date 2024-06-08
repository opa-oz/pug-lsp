package completion

import (
	"github.com/opa-oz/go-todo/todo"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func MixinsCompletion(completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	valueKind := protocol.CompletionItemKindValue

	todo.T("Grab mixins from lined files and offer")
	defaultSymbol := "mixins"
	defaultToInsert := "mixins."

	completionItems = append(completionItems, protocol.CompletionItem{
		Label:      todo.String("Look for mixins", "Default mixin - [Mixin]"),
		Kind:       &valueKind,
		Detail:     &defaultSymbol,
		InsertText: &defaultToInsert,
	})

	return &completionItems
}
