package query

import sitter "github.com/smacker/go-tree-sitter"

func HasContentAncestor(node *sitter.Node) bool {
	return FindUpwards(node, ContentNode, -1) != nil
}