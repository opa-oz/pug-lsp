package query_test

import (
	"context"
	"strings"
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestAllIncludes(t *testing.T) {
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

	includes, err := query.FindAllIncludes(testTree)

	assert.NoError(t, err)
	assert.Equal(t, len(*includes), 3)

	var includesFiles []string

	for _, strRange := range *includes {
		original := strings.Trim(content[strRange.StartPos:strRange.EndPos], " ")

		includesFiles = append(includesFiles, original)
	}

	assert.Equal(t, includesFiles, []string{
		"mixins/logo", "head", "footer",
	})
}
