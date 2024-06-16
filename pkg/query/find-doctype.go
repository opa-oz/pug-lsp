package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// FindDoctype searches for the presence of a doctype declaration in the given parse tree.
//
// It uses the FindAll function to search for nodes matching the DoctypeQ query within
// the subtree rooted at the root node of the provided parse tree.
//
// Parameters:
//
//	tree *sitter.Tree - The parse tree in which to search for the doctype declaration.
//
// Returns:
//
//	bool - true if a doctype declaration is found in the parse tree, false otherwise.
//	       Returns false if there's an error during the search.
func FindDoctype(tree *sitter.Tree) bool {
	results, err := FindAll(tree.RootNode(), DoctypeQ)
	if err != nil {
		return false
	}

	return len(*results) > 0
}
