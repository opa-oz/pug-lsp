package documents

import (
	"context"

	"github.com/opa-oz/go-todo/todo"
	"github.com/opa-oz/pug-lsp/pkg/lsp"
	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/pkg/errors"
	sitter "github.com/smacker/go-tree-sitter"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Document struct {
	URI      protocol.DocumentUri
	Path     string
	Tree     *sitter.Tree
	Content  *string
	Includes []*lsp.Include
}

// ApplyChanges updates the content of the Document from LSP textDocument/didChange events.
func (d *Document) ApplyChanges(ctx context.Context, changes []interface{}) error {
	for _, change := range changes {
		switch c := change.(type) {
		case protocol.TextDocumentContentChangeEvent:
			originalContent := *d.Content
			startIndex, endIndex := c.Range.IndexesIn(originalContent)
			modified := originalContent[:startIndex] + c.Text + originalContent[endIndex:]
			d.Content = &modified
		case protocol.TextDocumentContentChangeEventWhole:
			d.Content = &c.Text
		}
	}

	newTree, err := pug.UpdateParsedTreeFromString(ctx, d.Tree, *d.Content)
	if err != nil {
		return errors.Wrapf(err, "undable to update content: %s", d.Path)
	}

	todo.T("Applied changes")
	d.Tree = newTree

	return nil
}

func (d *Document) GetAtPosition(position *protocol.Position) *sitter.Node {
	node := d.Tree.RootNode().NamedDescendantForPointRange(
		sitter.Point{
			Row:    position.Line,
			Column: position.Character,
		},
		sitter.Point{
			Row:    position.Line,
			Column: position.Character,
		},
	)

	return node
}

func (d *Document) GetBeforeTrigger(position *protocol.Position) *sitter.Node {
	node := d.Tree.RootNode().NamedDescendantForPointRange(
		sitter.Point{
			Row:    position.Line,
			Column: position.Character - 2,
		},
		sitter.Point{
			Row:    position.Line,
			Column: position.Character - 2,
		},
	)

	return node
}
