package query

import sitter "github.com/smacker/go-tree-sitter"

func HasMixinDefinitionAncestor(node *sitter.Node) bool {
	return FindUpwards(node, MixinDefinitionNode, -1) != nil
}
