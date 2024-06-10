package documents

import (
	"context"
	"os"
	"strings"

	"github.com/opa-oz/go-todo/todo"
	"github.com/opa-oz/pug-lsp/pkg/lsp"
	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/opa-oz/pug-lsp/pkg/utils"
	"github.com/pkg/errors"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type DocumentStore struct {
	documents map[string]*Document
	logger    utils.Logger
	fs        utils.FileStore
}

func NewDocumentStore(logger utils.Logger, fs utils.FileStore) *DocumentStore {
	return &DocumentStore{
		logger:    logger,
		documents: map[string]*Document{},
		fs:        fs,
	}
}

func (ds *DocumentStore) DocumentDidOpen(ctx context.Context, params protocol.DidOpenTextDocumentParams) (*Document, error) {
	langID := params.TextDocument.LanguageID

	if langID != "pug" {
		return nil, nil
	}

	uri := params.TextDocument.URI
	path, err := ds.normalizeFilepath(uri)
	if err != nil {
		return nil, err
	}

	tree, err := pug.GetParsedTreeFromString(ctx, params.TextDocument.Text)

	doc := &Document{
		URI:      uri,
		Path:     path,
		Content:  &params.TextDocument.Text,
		Tree:     tree,
		Includes: make(map[string]*lsp.Include),
	}

	doc.HasDoctype = query.FindDoctype(tree)
	ds.documents[path] = doc

	return doc, nil
}

func (ds *DocumentStore) RefreshIncludes(ctx context.Context, doc *Document) {
	includes, err := query.FindAllIncludes(doc.Tree)
	if err != nil {
		ds.logger.Err(err)
		return
	}

	originalContent := *doc.Content
	for _, strRange := range *includes {
		original := strings.Trim(originalContent[strRange.StartPos:strRange.EndPos], " ")
		uri := ds.partialToUri(original, doc)

		newInclude := lsp.NewInclude(&original, &uri)
		_, ok := doc.Includes[original]

		if !ok {
			doc.Includes[original] = newInclude
			ds.LoadIncludedFile(ctx, newInclude)
		}

	}
}

func (ds *DocumentStore) LoadIncludedFile(ctx context.Context, include *lsp.Include) {
	content, err := os.ReadFile(*include.Path)

	if err != nil {
		ds.logger.Err(err)
		return
	}

	uri := *include.Path
	if !strings.HasPrefix(*include.Path, "file:/") {
		uri = "file://" + *include.Path
	}
	params := protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{
			URI:        uri,
			LanguageID: todo.String("Move to constant", "pug"),
			Version:    1,
			Text:       string(content),
		},
	}

	ds.DocumentDidOpen(ctx, params)
}

func (ds *DocumentStore) partialToUri(original string, doc *Document) string {
	parentParts := strings.Split(doc.Path, "/")
	parentFolder := parentParts[0 : len(parentParts)-1] // end of path is always file, remove it
	includeParts := strings.Split(original, "/")
	includeFilename := includeParts[len(includeParts)-1]

	if !strings.HasSuffix(includeFilename, ".pug") {
		includeFilename = includeFilename + ".pug"
	}

	parentFolder = append(parentFolder, includeParts[0:len(includeParts)-1]...)
	parentFolder = append(parentFolder, includeFilename)

	joined := strings.Join(parentFolder, "/")

	return joined

}

func (ds *DocumentStore) normalizeFilepath(uri string) (string, error) {
	path, err := utils.UriToPath(uri)
	if err != nil {
		return "", errors.Wrapf(err, "unable to parse URI: %s", uri)
	}

	return ds.fs.Canonical(path), nil
}

func (ds *DocumentStore) Get(uri string) (*Document, bool) {
	path, err := ds.normalizeFilepath(uri)
	if err != nil {
		ds.logger.Err(err)
		return nil, false
	}

	d, ok := ds.documents[path]
	return d, ok
}
