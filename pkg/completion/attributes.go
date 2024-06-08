package completion

import (
	"fmt"

	"github.com/opa-oz/go-todo/todo"
	"github.com/opa-oz/pug-lsp/pkg/documents"
	"github.com/opa-oz/pug-lsp/pkg/html"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GlobalCompletion(completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	valueKind := protocol.CompletionItemKindValue
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

	return &completionItems
}

func LocalCompletion(doc *documents.Document, completionItems []protocol.CompletionItem, params *protocol.CompletionParams, firstTime bool) *[]protocol.CompletionItem {
	valueKind := protocol.CompletionItemKindValue
	node := doc.GetAtPosition(&params.Position)

	if node != nil {
		node = node.Parent()
		nodeType := node.Type()
		todo.String(fmt.Sprintf("nodeType='%s'", nodeType))

		if nodeType == "tag" {
			node = node.Child(0)
			nodeType = node.Type()

			todo.String(fmt.Sprintf("nodeType2='%s', %s, %d, %d", node.Type(), node.String(), node.StartByte(), node.EndByte()))
		}

		originalContent := *doc.Content

		if nodeType == "tag_name" {
			tagName := originalContent[node.StartByte():node.EndByte()]
			todo.String(fmt.Sprintf("tagname='%s'", tagName))

			specificAttributes := html.GetAttributes(tagName)
			if specificAttributes != nil {
				for _, attr := range *specificAttributes {
					attrCopy := attr // Create a copy of attr
					attrToInsert := attr + "="
					completionItems = append(completionItems, protocol.CompletionItem{
						Label:      attrCopy + " - [S-Attr]",
						Kind:       &valueKind,
						Detail:     &attrCopy,
						InsertText: &attrToInsert,
					})
				}
			}
		}
	}
	return &completionItems
}
