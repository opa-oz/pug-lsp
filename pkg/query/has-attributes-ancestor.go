package query

import sitter "github.com/smacker/go-tree-sitter"

func HasAttributesAncestor(node *sitter.Node) bool {
	return FindUpwards(node, AttributesNode, maxDepth) != nil
}
