package documents_test

import (
	"context"
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/documents"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/stretchr/testify/assert"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestApplyChanges(t *testing.T) {
	orig := `
    doctype html
    html
        body
            div
    `
	newContent := `
    doctype html
    html
        body
            div.my-class#my-div
                p
                    span My awesome text
    `

	doc := documents.Document{}

	changes := []interface{}{
		protocol.TextDocumentContentChangeEventWhole{
			Text: orig,
		},
	}

	err := doc.ApplyChanges(context.TODO(), changes)
	assert.NoError(t, err)
	assert.Equal(t, *doc.Content, orig)
	assert.NotNil(t, doc.Tree)

	changes = []interface{}{
		protocol.TextDocumentContentChangeEventWhole{
			Text: newContent,
		},
	}
	err = doc.ApplyChanges(context.TODO(), changes)
	assert.NoError(t, err)
	assert.Equal(t, *doc.Content, newContent)
	assert.NotNil(t, doc.Tree)
}

func TestPositions(t *testing.T) {
	orig := `
    doctype html
    html
        body
            div.my-class#my-div
                p
                    span My awesome text
    `

	doc := documents.Document{}

	changes := []interface{}{
		protocol.TextDocumentContentChangeEventWhole{
			Text: orig,
		},
	}

	err := doc.ApplyChanges(context.TODO(), changes)
	assert.NoError(t, err)
	assert.Equal(t, *doc.Content, orig)
	assert.NotNil(t, doc.Tree)

	position := protocol.Position{
		Line:      6,
		Character: 25,
	}

	node := doc.GetAtPosition(&position)
	assert.NotNil(t, node)
	assert.Equal(t, node.Type(), string(query.ContentNodeType)) // we found "My awesome text"

	node = doc.GetBeforeTrigger(&position)
	assert.NotNil(t, node)
	assert.Equal(t, node.Type(), string(query.TagNameNode)) // we found tag_name=span before "My awesome text"
}
