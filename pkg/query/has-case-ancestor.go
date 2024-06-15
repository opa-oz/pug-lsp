package query

import sitter "github.com/smacker/go-tree-sitter"

func HasCaseAncestor(node *sitter.Node) bool {
	return FindUpwards(node, CaseNode, MaxDepth) != nil
}
