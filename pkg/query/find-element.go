package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func FindUpwards(node *sitter.Node, nodeType NodeType, maxDepth int) *sitter.Node {
	if node == nil {
		return nil
	}

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
