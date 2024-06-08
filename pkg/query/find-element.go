package query

import sitter "github.com/smacker/go-tree-sitter"

type NodeType string

const (
	TagNode           NodeType = "tag"
	TagNameNode       NodeType = "tag_name"
	AttributeNode     NodeType = "attribute"
	AttributesNode    NodeType = "attributes"
	AttributeNameNode NodeType = "attribute_name"
	ChildrenNode      NodeType = "children"
)

func FindUpwards(node *sitter.Node, nodeType NodeType, maxDepth int) *sitter.Node {
	iterations := 0
	for {
		if iterations > maxDepth {
			break
		}
		if node.Type() == string(nodeType) {
			return node
		}

		iterations += 1
		node = node.Parent()
	}

	return nil
}
