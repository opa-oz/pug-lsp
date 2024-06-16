package query

import sitter "github.com/smacker/go-tree-sitter"

// FindMixinDefinitions finds all nodes in the subtree rooted at the given node
// that match the criteria for mixin definitions.
//
// It uses the FindAllNodes function to search for nodes that match the MixinDefinitionQ
// query within the subtree rooted at the provided node.
//
// Parameters:
//
//	node *sitter.Node - The root node from which to start searching for mixin definitions.
//
// Returns:
//
//	[]*sitter.Node - A slice of pointers to nodes that match the mixin definition query.
//	                  Returns nil if there's an error during the search or if no nodes
//	                  match the query.
func FindMixinDefinitions(node *sitter.Node) *[]*sitter.Node {
	nodes, err := FindAllNodes(node, MixinDefinitionQ)

	if err != nil {
		return nil
	}

	return nodes
}
