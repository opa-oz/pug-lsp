package query

import sitter "github.com/smacker/go-tree-sitter"

// HasCaseAncestor checks if the given node or any of its ancestors
// contains a node of type CaseNode.
//
// It searches upwards from the provided node using the FindUpwards function
// to locate a node of type CaseNode within a maximum depth defined by MaxDepth.
// If such a node is found, HasCaseAncestor returns true; otherwise, it returns false.
//
// Parameters:
//
//	node *sitter.Node - The starting node from which to begin the search.
//
// Returns:
//
//	bool - true if a CaseNode ancestor is found, false otherwise.
func HasCaseAncestor(node *sitter.Node) bool {
	return FindUpwards(node, CaseNode, MaxDepth) != nil
}
