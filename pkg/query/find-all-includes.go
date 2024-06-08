package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

const query = "(include (filename) @incl)"

func FindAllIncludes(tree *sitter.Tree) (*[]*StrRange, error) {
	return FindAll(tree.RootNode(), query)
}
