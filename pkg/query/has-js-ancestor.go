package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func HasJSAncestor(node *sitter.Node) bool {
	return FindUpwards(node, JSNode, MaxDepth) != nil || FindUpwards(node, BufferedCodeNode, MaxDepth) != nil || FindUpwards(node, UnBufferedCodeNode, MaxDepth) != nil
}
