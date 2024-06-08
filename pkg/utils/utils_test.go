package utils_test

import (
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestUriToPath(t *testing.T) {
	testTable := []struct {
		expected string
		input    string
	}{
		// {expected: "c:/tmp/foo.md", input: "file:///c:/tmp/foo.md"},
		{expected: "/src/mixin/logo.pug", input: "file:///src/mixin/logo.pug"},
		{expected: "/Users/vladimirlevin/Repos/WebStorm/kiss-anime/src/page.pug", input: "file:///Users/vladimirlevin/Repos/WebStorm/kiss-anime/src/page.pug"},
	}

	for _, tcase := range testTable {
		result, err := utils.UriToPath(tcase.input)
		assert.NoError(t, err)
		assert.Equal(t, result, tcase.expected)
	}
}
