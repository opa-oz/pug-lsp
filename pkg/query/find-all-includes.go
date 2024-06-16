package query

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// FindAllIncludes searches for all nodes in the parse tree that match the criteria
// for include filenames.
//
// It uses the FindAll function to search for nodes that match the IncludeFilenamesQ
// query within the subtree rooted at the root node of the provided parse tree.
//
// Parameters:
//
//	tree *sitter.Tree - The parse tree in which to search for include filenames.
//
// Returns:
//
//	*[]*StrRange - A pointer to a slice of StrRange pointers representing the byte ranges
//	               of nodes that match the include filenames query. Returns nil and an error
//	               if there's any issue during node retrieval.
func FindAllIncludes(tree *sitter.Tree) (*[]*StrRange, error) {
	return FindAll(tree.RootNode(), IncludeFilenamesQ)
}
