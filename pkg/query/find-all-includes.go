package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func FindAllIncludes(tree *sitter.Tree) (*[]*StrRange, error) {
	return FindAll(tree.RootNode(), IncludeFilenamesQ)
}
