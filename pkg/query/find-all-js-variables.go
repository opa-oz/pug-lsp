package query

import (
	"github.com/opa-oz/pug-lsp/pkg/lsp"
	sitter "github.com/smacker/go-tree-sitter"
)

func FindAllJSVariables(node *sitter.Node, originalContent *string) *[]lsp.JSVariable {
	nodes, err := FindAllNodes(node, JSVariablesQ)

	if err != nil {
		return nil
	}

	result := make([]lsp.JSVariable, len(*nodes))

	for _, node := range *nodes {
		val := (*originalContent)[node.StartByte():node.EndByte()]
		result = append(result, lsp.JSVariable(val))
	}

	return &result
}
