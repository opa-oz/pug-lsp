package query

import sitter "github.com/smacker/go-tree-sitter"

// HasContentAncestor checks if the given node or any of its ancestors
// contains a node of type ContentNode.
//
// It searches upwards from the provided node using the FindUpwards function
// to locate a node of type ContentNode within a maximum depth defined by MaxDepth.
// If such a node is found, HasContentAncestor returns true; otherwise, it returns false.
//
// Parameters:
//
//	node *sitter.Node - The starting node from which to begin the search.
//
// Returns:
//
//	bool - true if a ContentNode ancestor is found, false otherwise.
func HasContentAncestor(node *sitter.Node) bool {
	return FindUpwards(node, ContentNode, MaxDepth) != nil
}
