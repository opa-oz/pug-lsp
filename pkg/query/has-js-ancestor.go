package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func HasJSAncestor(node *sitter.Node) bool {
	return FindUpwards(node, JSNode, maxDepth) != nil || FindUpwards(node, BufferedCodeNode, maxDepth) != nil || FindUpwards(node, UnBufferedCodeNode, maxDepth) != nil
}
