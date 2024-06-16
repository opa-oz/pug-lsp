package lsp

import (
	"strings"

	"github.com/opa-oz/pug-lsp/pkg/query"
	sitter "github.com/smacker/go-tree-sitter"
)

type Mixin struct {
	Source     string
	Name       string
	Definition string
	Arguments  *[]string
}

func NewMixin(source string, node *sitter.Node, content *string) *Mixin {
	nameNode := query.FindDownwards(node, query.MixinNameNode, 2)

	if nameNode == nil {
		return nil
	}

	definition := (*content)[nameNode.StartByte():nameNode.EndByte()]
	mixinAttributes := query.FindDownwards(node, query.MixinAttributesNode, 2)
	var arguments []string

	if mixinAttributes != nil {
		attributesRanges, err := query.FindAll(mixinAttributes, query.AttributeNamesQ)
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
	}
}
