package query

import "github.com/opa-oz/go-todo/todo"

type NodeType string

const (
	TagNode             NodeType = "tag"
	TagNameNode         NodeType = "tag_name"
	AttributeNode       NodeType = "attribute"
	AttributesNode      NodeType = "attributes"
	AttributeNameNode   NodeType = "attribute_name"
	AttributeValueNode  NodeType = "attribute_value"
	ChildrenNode        NodeType = "children"
	MixinNode           NodeType = "mixin_use"
	DoctypeNode         NodeType = "doctype"
	DoctypeNameNode     NodeType = "doctype_name"
	ContentNode         NodeType = "content"
	JSNode              NodeType = "javascript"
	BufferedCodeNode    NodeType = "buffered_code"
	UnBufferedCodeNode  NodeType = "unbuffered_code"
	CaseNode            NodeType = "case"
	MixinDefinitionNode NodeType = "mixin_definition"
	MixinNameNode       NodeType = "mixin_name"
	MixinAttributesNode NodeType = "mixin_attributes"
	KeywordNode         NodeType = "keyword"
	FilenameNode        NodeType = "filename"
	IncludeNode         NodeType = "include"
)

var MaxDepth = todo.Int("Is 5 enough?", 5)

type Query string

const (
	AttributeNamesQ   Query = "(attribute_name) @attr"
	IncludeFilenamesQ Query = "(include (filename) @incl)"
	DoctypeQ          Query = "(doctype) @doc"
	MixinDefinitionQ  Query = "(mixin_definition) @def"
	JSVariablesQ      Query = "(javascript (identifier) @var)"
)
