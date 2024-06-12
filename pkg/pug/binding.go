// go:build !plugin
package pug

// #cgo CFLAGS: -std=c11 -fPIC
// #include "tree_sitter/parser.h"
// // NOTE: if your language has an external scanner, add it here.
//TSLanguage *tree_sitter_pug();
import "C"
import (
	"context"
	"unsafe"

	sitter "github.com/smacker/go-tree-sitter"
)

// Get the tree-sitter Language for this grammar.
func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_pug())
	return sitter.NewLanguage(ptr)
}

func NewSitterParser() *sitter.Parser {
	parser := sitter.NewParser()
	parser.SetLanguage(GetLanguage())

	return parser
}

func GetParsedTreeFromString(ctx context.Context, source string) (*sitter.Tree, error) {
	sourceCode := []byte(source)
	parser := NewSitterParser()
	tree, err := parser.ParseCtx(ctx, nil, sourceCode)

	if err != nil {
		return nil, err
	}

	return tree, nil
}

func UpdateParsedTreeFromString(ctx context.Context, oldTree *sitter.Tree, source string) (*sitter.Tree, error) {
	sourceCode := []byte(source)
	parser := NewSitterParser()
	tree, err := parser.ParseCtx(ctx, oldTree, sourceCode)

	if err != nil {
		return nil, err
	}

	return tree, nil
}
