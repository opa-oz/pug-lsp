package query

import (
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
)

func FindDoctype(tree *sitter.Tree) bool {
	q := fmt.Sprintf("(%s) @doc", DoctypeNode)

	results, err := FindAll(tree.RootNode(), q)
	if err != nil {
		return false
	}

	return len(*results) > 0
}
