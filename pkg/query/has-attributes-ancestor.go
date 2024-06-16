package query

import sitter "github.com/smacker/go-tree-sitter"

// HasAttributesAncestor checks if the given node or any of its ancestors
// contains a node of type AttributesNode.
//
// It searches upwards from the provided node using the FindUpwards function
// to locate a node of type AttributesNode within a maximum depth defined by MaxDepth.
// If such a node is found, HasAttributesAncestor returns true; otherwise, it returns false.
//
// Parameters:
//
//	node *sitter.Node - The starting node from which to begin the search.
//
// Returns:
//
//	bool - true if an AttributesNode ancestor is found, false otherwise.
func HasAttributesAncestor(node *sitter.Node) bool {
	return FindUpwards(node, AttributesNode, MaxDepth) != nil
}
