package query

import sitter "github.com/smacker/go-tree-sitter"

func FindMixinDefinitions(node *sitter.Node) *[]*sitter.Node {
	nodes, err := FindAllNodes(node, MixinDefinitionQ)

	if err != nil {
		return nil
	}

	return nodes
}
