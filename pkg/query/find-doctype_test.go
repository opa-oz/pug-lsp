package query_test

import (
	"context"
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestFindDoctype(t *testing.T) {
	content := `
    include mixins/logo
    include head
    include footer

    doctype html
    html
        div
    `
	testTree, err := pug.GetParsedTreeFromString(context.TODO(), content)

	assert.NoError(t, err)
	assert.NotNil(t, testTree)

	hasDoctype := query.FindDoctype(testTree)

	assert.True(t, hasDoctype)
}
