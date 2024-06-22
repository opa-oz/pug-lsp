package completion

import (
	"github.com/opa-oz/go-todo/todo"
	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/opa-oz/pug-lsp/pkg/lsp"
	"github.com/opa-oz/pug-lsp/pkg/query"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func keywordsCase(completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	keywordKind := protocol.CompletionItemKindKeyword
	snippetKind := protocol.CompletionItemKindSnippet
	keytags := lsp.GetCaseKeywords()

	for _, key := range *keytags {
		keyName := string(key)
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      keyName,
			Kind:       &keywordKind,
			Detail:     &keyName,
			InsertText: &keyName,
		})
	}

	breakKey := todo.String("Move break to snippets", "break")
	insertion := "- break"
	completionItems = append(completionItems, protocol.CompletionItem{
		Label:      breakKey,
		Kind:       &snippetKind,
		Detail:     &insertion,
		InsertText: &insertion,
	})

	return &completionItems
}
func keywords(completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	keywordKind := protocol.CompletionItemKindKeyword
	keytags := lsp.GetKeywords()

	for _, key := range *keytags {
		keyName := string(key)
		item := protocol.CompletionItem{
			Label:      keyName,
			Kind:       &keywordKind,
			InsertText: &keyName,
			Data:       DataToString(KeywordKind, keyName),
		}

		completionItems = append(completionItems, item)
	}

	return &completionItems
}

func globalTags(completionItems []protocol.CompletionItem) *[]protocol.CompletionItem {
	valueKind := protocol.CompletionItemKindConstant
	htmlTags := html.GetTags()

	for _, tag := range *htmlTags {
		tagCopy := string(tag)
		item := protocol.CompletionItem{
			Label:      tagCopy,
			Kind:       &valueKind,
			InsertText: &tagCopy,
			Data:       DataToString(TagKind, tagCopy),
		}

		completionItems = append(completionItems, item)
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
	prevSibling := node.PrevSibling()
	if query.HasCaseAncestor(node) || (prevSibling != nil && prevSibling.Type() == string(query.CaseNode)) {
		return keywordsCase(completionItems)
	}

	// parent := node.Parent()
	// if parent != nil {
	// 	(*meta.Logger).Println("Parent", parent.Type())
	// }
	// if prevSibling != nil {
	// 	(*meta.Logger).Println("Sibling", prevSibling.Type())
	// }

	tag := query.FindUpwards(node, query.TagNode, 1)
	if tag != nil {
		items := globalTags(completionItems)
		items = keywords(*items)

		return items
	}

	return &completionItems
}
