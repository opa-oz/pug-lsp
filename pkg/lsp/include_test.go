package lsp_test

import (
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/lsp"
	"github.com/stretchr/testify/assert"
)

func TestNewInclude(t *testing.T) {
	original := "src/mixins/logo"
	path := "/Users/opa-oz/test/main.pug"

	include := lsp.NewInclude(&original, &path)

	expectedPrefix := "mixins"
	expectedUri := "file:///Users/opa-oz/test/main.pug"
	assert.NotNil(t, include)
	assert.Equal(t, &original, include.Original)
	assert.Equal(t, expectedPrefix, *include.Prefix)
	assert.Equal(t, &expectedUri, include.URI)
	assert.Equal(t, path, *include.Path)
}
