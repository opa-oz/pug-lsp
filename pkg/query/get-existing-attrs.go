package query

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

type ExistingAttributes = map[string]bool

// GetExistingAttributes retrieves existing attribute names from the given node's
// attributes section within the original content.
//
// It searches upwards from the provided node to find the nearest node of type
// AttributesNode. If found, it extracts attribute names using the AttributeNamesQ
// query and returns a map of attribute names as keys with boolean values indicating
// existence.
//
// Parameters:
//
//	node *sitter.Node - The starting node from which to begin the search.
//	originalContent string - The document's content from which to extract attribute names.
//
// Returns:
//
//	*ExistingAttributes - A pointer to a map[string]bool where keys are attribute names
//	                      found and values are true.
//	                      Returns nil if the node is nil, no AttributesNode is found,
//	                      or if there's an error retrieving attribute ranges.
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
