package query

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

type ExistingAttributes = map[string]bool

func GetExistingAttributes(node *sitter.Node, originalContent string) *ExistingAttributes {
	var attributes map[string]bool
	if node == nil {
		return nil
	}

	attributesNode := FindUpwards(node, AttributesNode, 1)

	if attributesNode == nil {
		return nil
	}

	attributesRanges, err := FindAll(attributesNode, AttributeNamesQ)
	if err != nil {
		return nil
	}
	attributes = make(map[string]bool)

	for _, rng := range *attributesRanges {
		attr := strings.Trim(originalContent[rng.StartPos:rng.EndPos], " ")
		attributes[attr] = true
	}

	return &attributes
}
