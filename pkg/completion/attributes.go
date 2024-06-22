package completion

import (
	"fmt"
	"strings"

	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/opa-oz/pug-lsp/pkg/utils"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GlobalCompletion(meta *CompletionMetaParams, completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	valueKind := protocol.CompletionItemKindProperty
	globalAttrs := html.GlobalAttrs()

	for _, attr := range *globalAttrs {
		if meta.ExistingAttrs != nil {
			// If attribute is already used - skip
			_, ok := (*meta.ExistingAttrs)[attr]
			if ok {
				continue
			}
		}
		attrCopy := attr // Create a copy of attr
		item := protocol.CompletionItem{
			Label:      attrCopy,
			Kind:       &valueKind,
			InsertText: html.GetInsertText(attr),
			Data:       DataToString(AttributeKind, attr),
		}

		completionItems = append(completionItems, item)
	}

	return &completionItems
}

func LocalCompletion(meta *CompletionMetaParams, completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	valueKind := protocol.CompletionItemKindProperty
	eventKind := protocol.CompletionItemKindEvent

	node := meta.CurrentNode
	if node == nil {
		return &completionItems
	}

	mixin := query.FindUpwards(node, query.MixinNode, 1)
	if mixin != nil {
		return &completionItems
	}

	tag := query.FindUpwards(node, query.TagNode, 3)

	if tag == nil {
		return &completionItems
	}
	tagName := query.FindDownwards(tag, query.TagNameNode, 2)

	if tagName == nil {
		return &completionItems
	}

	originalContent := *meta.Doc.Content
	tagNameStr := originalContent[tagName.StartByte():tagName.EndByte()]

	specificAttributes := html.GetAttributes(tagNameStr)

	if specificAttributes != nil {
		for _, attr := range *specificAttributes {
			kind := valueKind
			if strings.HasPrefix(attr, "on") {
				kind = eventKind
			}

			if meta.ExistingAttrs != nil {
				// If attribute is already used - skip
				_, ok := (*meta.ExistingAttrs)[attr]
				if ok {
					continue
				}
			}

			attrCopy := attr // Create a copy of attr
			item := protocol.CompletionItem{
				Label:      attrCopy,
				Kind:       &kind,
				InsertText: html.GetInsertText(attr),
				Data:       DataToString(AttributeKind, attr),
			}
			completionItems = append(completionItems, item)
		}
	}
	return &completionItems
}

func AndCompletion(meta *CompletionMetaParams, completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	snippetKind := protocol.CompletionItemKindSnippet

	node := meta.CurrentNode
	if node == nil {
		return &completionItems
	}

	tag := query.FindUpwards(node, query.TagNode, 2)
	if tag == nil {
		return &completionItems
	}

	attributes := "attributes("
	details := fmt.Sprintf("%s{ })", attributes)
	completionItems = append(completionItems, protocol.CompletionItem{
		Label:      attributes,
		Kind:       &snippetKind,
		InsertText: &details,
		Preselect:  utils.PtrBool(true),
		Data:       DataToString(AndAttributesKind, details),
	})

	return &completionItems
}
