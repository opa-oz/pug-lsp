package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type NodeType string

const (
	TagNode            NodeType = "tag"
	TagNameNode        NodeType = "tag_name"
	AttributeNode      NodeType = "attribute"
	AttributesNode     NodeType = "attributes"
	AttributeNameNode  NodeType = "attribute_name"
	AttributeValueNode NodeType = "attribute_value"
	ChildrenNode       NodeType = "children"
	MixinNode          NodeType = "mixin_use"
	DoctypeNode        NodeType = "doctype"
	DoctypeNameNode    NodeType = "doctype_name"
	ContentNodeType    NodeType = "content"
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
		if iterations >= maxDepth {
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
