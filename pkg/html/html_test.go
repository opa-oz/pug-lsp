package html_test

import (
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/stretchr/testify/assert"
)

func TestGetAttributes(t *testing.T) {
	assert.Equal(
		t,
		html.GetAttributes(string(html.Li)),
		&[]string{"value"},
	)

	assert.Equal(
		t,
		html.GetAttributes(string(html.Menu)),
		&[]string{"type"},
	)
}

func TestGetInsertText(t *testing.T) {
	testTable := []struct {
		input    string
		expected string
	}{
		{"style", "style={ }"},
		{"checked", "checked"},
		{"disabled", "disabled"},
		{"hidden", "hidden"},
		{"class", "class=[ ]"},
		{"click", "'(click)'"},
	}

	for _, tcase := range testTable {
		assert.Equal(t, html.GetInsertText(tcase.input), &tcase.expected)
	}
}
