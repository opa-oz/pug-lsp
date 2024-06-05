package documents

import (
	"context"

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
	Content  string
	Includes []*lsp.Include
}

// ApplyChanges updates the content of the Document from LSP textDocument/didChange events.
func (d *Document) ApplyChanges(ctx context.Context, changes []interface{}) error {
	for _, change := range changes {
		switch c := change.(type) {
		case protocol.TextDocumentContentChangeEvent:
			startIndex, endIndex := c.Range.IndexesIn(d.Content)
			d.Content = d.Content[:startIndex] + c.Text + d.Content[endIndex:]
		case protocol.TextDocumentContentChangeEventWhole:
			d.Content = c.Text
		}
	}

	newTree, err := pug.UpdateParsedTreeFromString(ctx, d.Tree, d.Content)
	if err != nil {
		return errors.Wrapf(err, "undable to update content: %s", d.Path)
	}

	d.Tree = newTree

	return nil
}
