package query

import sitter "github.com/smacker/go-tree-sitter"

func HasJSAncestor(node *sitter.Node) bool {
	return FindUpwards(node, ContentNodeJS, -1) != nil
}
