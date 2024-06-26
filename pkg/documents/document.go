package documents

import (
	"context"

	"github.com/opa-oz/pug-lsp/pkg/lsp"
	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/pkg/errors"
	sitter "github.com/smacker/go-tree-sitter"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type IncludesMap = map[string]*lsp.Include
type MixinsMap = map[string]*query.Mixin

type Document struct {
	URI                      protocol.DocumentUri
	Path                     string
	Tree                     *sitter.Tree
	Content                  *string
	Includes                 IncludesMap
	Mixins                   MixinsMap
	HasDoctype               bool
	NeedToRefreshDiagnostics bool
	NeedToRefreshIncludes    bool
	// JSVariables *[]lsp.JSVariable
}

// ApplyChanges updates the content of the Document from LSP textDocument/didChange events.
//  1. Apply changes
//     1.1 If `TextDocumentContentChangeEvent`
//     Patch original content in place
//     1.2 Else if `TextDocumentContentChangeEvenWhole`
//     Update content completely
//  2. Regenerate sitter.Tree
//     Note: If pass `oldTree` to `UpdateParsedTreeFromString` - it works incorrectly, doing only partial update
//  3. Check if `doctype` present in file
//  4. Find and save mixins definitions
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

	newTree, err := pug.UpdateParsedTreeFromString(ctx, nil, *d.Content)
	if err != nil {
		return errors.Wrapf(err, "undable to update content: %s", d.Path)
	}

	d.Tree = newTree
	d.HasDoctype = query.FindDoctype(newTree)
	// d.JSVariables = query.FindAllJSVariables(newTree.RootNode(), d.Content)
	d.RefreshMixins(ctx)

	return nil
}

func (d *Document) RefreshMixins(ctx context.Context) {
	d.Mixins = make(MixinsMap)
	mixinDefinitions := query.FindMixinDefinitions(d.Tree.RootNode())

	if mixinDefinitions == nil {
		return
	}

	for _, def := range *mixinDefinitions {
		mixin := query.NewMixin(d.URI, def, d.Content)

		if mixin != nil {
			d.Mixins[mixin.Name] = mixin
		}
	}
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
