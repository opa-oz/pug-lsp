package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func FindDoctype(tree *sitter.Tree) bool {
	results, err := FindAll(tree.RootNode(), DoctypeQ)
	if err != nil {
		return false
	}

	return len(*results) > 0
}
