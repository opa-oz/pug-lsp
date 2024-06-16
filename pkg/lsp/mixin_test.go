package lsp_test

import (
	"context"
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/lsp"
	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestNewMixin(t *testing.T) {
	content := `
    doctype html
    mixin testA(arg1, arg2)
        div
    `
	testTree, err := pug.GetParsedTreeFromString(context.TODO(), content)

	assert.NoError(t, err)
	assert.NotNil(t, testTree)

	doctype := query.FindDownwards(testTree.RootNode(), query.DoctypeNode, 2)

	assert.NotNil(t, doctype)

	mixinDef := doctype.NextSibling()
	assert.NotNil(t, mixinDef)
	assert.Equal(t, mixinDef.Type(), string(query.MixinDefinitionNode))

	mixin := lsp.NewMixin("myfile.pug", mixinDef, &content)

	assert.NotNil(t, mixin)
	assert.Equal(t, mixin.Source, "myfile.pug")
	assert.Equal(t, mixin.Name, "testA")
	assert.Equal(t, mixin.Definition, "testA(arg1, arg2)")

	assert.Equal(t, mixin.Arguments, &([]string{"arg1", "arg2"}))
}

func TestNewMixinWithoutArgs(t *testing.T) {
	content := `
    doctype html
    mixin testA
        div
    `
	testTree, err := pug.GetParsedTreeFromString(context.TODO(), content)

	assert.NoError(t, err)
	assert.NotNil(t, testTree)

	doctype := query.FindDownwards(testTree.RootNode(), query.DoctypeNode, 2)

	assert.NotNil(t, doctype)

	mixinDef := doctype.NextSibling()
	assert.NotNil(t, mixinDef)
	assert.Equal(t, mixinDef.Type(), string(query.MixinDefinitionNode))

	mixin := lsp.NewMixin("myfile.pug", mixinDef, &content)

	assert.NotNil(t, mixin)
	assert.Equal(t, mixin.Source, "myfile.pug")
	assert.Equal(t, mixin.Name, "testA")
	assert.Equal(t, mixin.Definition, "testA")

	var emptyArguments []string
	assert.Equal(t, mixin.Arguments, &emptyArguments)
}
