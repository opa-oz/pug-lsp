package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// HasJSAncestor checks if the given node or any of its ancestors
// contains specific types of JavaScript-related nodes.
//
// It searches upwards from the provided node using the FindUpwards function
// to locate nodes of types JSNode, BufferedCodeNode, or UnBufferedCodeNode
// within a maximum depth defined by MaxDepth. If any of these nodes are found,
// HasJSAncestor returns true; otherwise, it returns false.
//
// Parameters:
//
//	node *sitter.Node - The starting node from which to begin the search.
//
// Returns:
//
//	bool - true if any JavaScript-related ancestor is found, false otherwise.
func HasJSAncestor(node *sitter.Node) bool {
	return FindUpwards(node, JSNode, MaxDepth) != nil || FindUpwards(node, BufferedCodeNode, MaxDepth) != nil || FindUpwards(node, UnBufferedCodeNode, MaxDepth) != nil
}
