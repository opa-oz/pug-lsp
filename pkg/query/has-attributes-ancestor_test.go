package query_test

import (
	"context"
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestHasAttributesAncestor(t *testing.T) {
	content := `
    doctype html
    html
        body
            div(style={border:"1px solid gray"})
    `
	testTree, err := pug.GetParsedTreeFromString(context.TODO(), content)

	assert.NoError(t, err)
	assert.NotNil(t, testTree)

	jsBlock := query.FindDownwards(testTree.RootNode(), query.JSNode, 20)

	assert.NotNil(t, jsBlock)
	assert.True(t, query.HasAttributesAncestor(jsBlock))

	tag := query.FindDownwards(testTree.RootNode(), query.TagNameNode, 20)

	assert.NotNil(t, tag)
	assert.False(t, query.HasAttributesAncestor(tag))

	attrs := query.FindDownwards(testTree.RootNode(), query.AttributesNode, 20)
	assert.NotNil(t, attrs)
	assert.True(t, query.HasAttributesAncestor(attrs))
}
