package query

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
	ContentNode        NodeType = "content"
	JSNode             NodeType = "javascript"
)
