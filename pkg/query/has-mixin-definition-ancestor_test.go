package query_test

import (
	"context"
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestMixinDefinitionHasAncestor(t *testing.T) {
	content := `
    doctype html
    mixin myMixin
        div
            div(style="color: #ffffff;")
                p
    `
	testTree, err := pug.GetParsedTreeFromString(context.TODO(), content)

	assert.NoError(t, err)
	assert.NotNil(t, testTree)

	tag := query.FindDownwards(testTree.RootNode(), query.TagNameNode, 20)

	assert.NotNil(t, tag)
	assert.Equal(t, "div", content[tag.StartByte():tag.EndByte()])
	assert.True(t, query.HasMixinDefinitionAncestor(tag))

	attrs := query.FindDownwards(testTree.RootNode(), query.AttributesNode, 20)
	assert.NotNil(t, attrs)
	assert.True(t, query.HasMixinDefinitionAncestor(attrs))

	doctype := query.FindDownwards(testTree.RootNode(), query.DoctypeNode, 20)
	assert.NotNil(t, doctype)
	assert.False(t, query.HasMixinDefinitionAncestor(doctype))
}
