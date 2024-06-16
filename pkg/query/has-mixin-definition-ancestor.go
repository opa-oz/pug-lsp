package query

import sitter "github.com/smacker/go-tree-sitter"

// HasMixinDefinitionAncestor checks if the given node or any of its ancestors
// contains a mixin definition node.
//
// It searches upwards from the provided node using the FindUpwards function
// to locate a node of type MixinDefinitionNode. If such a node is found,
// HasMixinDefinitionAncestor returns true; otherwise, it returns false.
//
// Parameters:
//
//	node *sitter.Node - The starting node from which to begin the search.
//
// Returns:
//
//	bool - true if a mixin definition ancestor is found, false otherwise.
func HasMixinDefinitionAncestor(node *sitter.Node) bool {
	return FindUpwards(node, MixinDefinitionNode, -1) != nil
}
