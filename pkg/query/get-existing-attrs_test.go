package query_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestGetExistingAttributes(t *testing.T) {
	content := `
    doctype html
    html
        div
            div
                ul(style="padding: 8px;" class="my class")
                    li
                    li
                    li(disabled)
                    li
    `
	testTree, err := pug.GetParsedTreeFromString(context.TODO(), content)

	assert.NoError(t, err)
	assert.NotNil(t, testTree)

	attributeName := query.FindDownwards(testTree.RootNode(), query.AttributeNode, 19)
	assert.Equal(t, attributeName.Type(), string(query.AttributeNode))
	assert.NotNil(t, attributeName)

	fmt.Println(content[attributeName.StartByte():attributeName.EndByte()])
	fmt.Println(attributeName.Parent().Type(), content[attributeName.Parent().StartByte():attributeName.Parent().EndByte()])

	assert.Equal(t, attributeName.Parent().Type(), string(query.AttributesNode))

	existingAttributes := query.GetExistingAttributes(attributeName, content)
	assert.NotNil(t, existingAttributes)

	keys := make([]string, len(*existingAttributes))
	assert.Equal(t, len(*existingAttributes), 2)

	i := 0
	for key := range *existingAttributes {
		keys[i] = key
		i++
	}
	assert.Equal(t, keys, []string{"style", "class"})
}
