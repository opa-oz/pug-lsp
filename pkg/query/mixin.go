package query

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Mixin struct {
	Source     string
	Name       string
	Definition string
	Arguments  *[]string
	Range      *protocol.Range
}

func NewMixin(source string, node *sitter.Node, content *string) *Mixin {
	nameNode := FindDownwards(node, MixinNameNode, 2)

	if nameNode == nil {
		return nil
	}

	definition := (*content)[nameNode.StartByte():nameNode.EndByte()]
	mixinAttributes := FindDownwards(node, MixinAttributesNode, 2)
	var arguments []string

	if mixinAttributes != nil {
		attributesRanges, err := FindAll(mixinAttributes, AttributeNamesQ)
		if err == nil {
			for _, rng := range *attributesRanges {
				arguments = append(arguments, strings.Trim((*content)[rng.StartPos:rng.EndPos], ""))
			}

			definition += (*content)[mixinAttributes.StartByte():mixinAttributes.EndByte()]
		}
	}

	return &Mixin{
		Source:     source,
		Name:       (*content)[nameNode.StartByte():nameNode.EndByte()],
		Arguments:  &arguments,
		Definition: definition,
		Range: &protocol.Range{
			Start: protocol.Position{
				Line:      node.StartPoint().Row,
				Character: node.StartPoint().Column,
			},
			End: protocol.Position{
				Line:      node.EndPoint().Row,
				Character: node.EndPoint().Column,
			},
		},
	}
}
