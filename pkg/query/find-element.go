package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// FindUpwards searches upwards from the given node to find the nearest ancestor
// node of the specified type within a maximum depth.
//
// It iterates upwards from the provided node using the Parent() method until it
// finds a node whose type matches the specified NodeType or until it reaches the
// root of the tree (node == nil). The search can be limited by a maximum depth,
// specified by maxDepth. If maxDepth is -1, the search continues until the root.
//
// Parameters:
//
//	node *sitter.Node - The starting node from which to begin the upward search.
//	nodeType NodeType - The type of node to search for, represented as a NodeType.
//	maxDepth int - The maximum depth to search upwards. Use -1 to search to the root.
//
// Returns:
//
//	*sitter.Node - A pointer to the nearest ancestor node of the specified type,
//	               or nil if no such node is found within the specified maxDepth
//	               or if the starting node is nil.
func FindUpwards(node *sitter.Node, nodeType NodeType, maxDepth int) *sitter.Node {
	iterations := 0
	for {
		if node == nil {
			break
		}

		if node.Type() == string(nodeType) {
			return node
		}
		if maxDepth != -1 && iterations >= maxDepth {
			break
		}

		iterations += 1
		node = node.Parent()
	}

	return nil
}

// FindDownwards searches downwards from the given node to find the nearest descendant
// node of the specified type within a maximum depth.
//
// It recursively searches through the children of the provided node to find a node
// whose type matches the specified NodeType. The search is limited by the maximum
// depth specified by maxDepth. If maxDepth is less than or equal to 0 or if the
// starting node is nil, the function returns nil immediately.
//
// Parameters:
//
//	node *sitter.Node - The starting node from which to begin the downward search.
//	nodeType NodeType - The type of node to search for, represented as a NodeType.
//	maxDepth int - The maximum depth to search downwards. Use 1 to search only immediate children.
//
// Returns:
//
//	*sitter.Node - A pointer to the nearest descendant node of the specified type,
//	               or nil if no such node is found within the specified maxDepth
//	               or if the starting node is nil.
func FindDownwards(node *sitter.Node, nodeType NodeType, maxDepth int) *sitter.Node {
	if maxDepth <= 0 || node == nil {
		return nil
	}

	if node.Type() == string(nodeType) {
		return node
	}

	childCount := int(node.ChildCount())
	if childCount == 0 {
		return nil
	}

	for i := 0; i < childCount; i++ {
		n := FindDownwards(node.Child(i), nodeType, maxDepth-1)

		if n != nil {
			return n
		}
	}

	return nil
}
