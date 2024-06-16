package query_test

import (
	"context"
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestFindMixinDefinitions(t *testing.T) {
	content := `
        mixin testA
            div

        mixin testB(arg1, arg2)
            p.paragraph
    `
	testTree, err := pug.GetParsedTreeFromString(context.TODO(), content)

	assert.NoError(t, err)
	assert.NotNil(t, testTree)

	defs := query.FindMixinDefinitions(testTree.RootNode())

	assert.NotNil(t, defs)
	assert.Equal(t, len(*defs), 2)
}
