package query_test

import (
	"context"
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/stretchr/testify/assert"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestFindAllOK(t *testing.T) {
	testTree, err := pug.GetParsedTreeFromString(context.TODO(), `
    doctype html
    html
        div
            div
                ul(style="padding: 8px;" class="my class")
                    li
                    li
                    li(disabled)
                    li
    `)

	assert.NoError(t, err)
	assert.NotNil(t, testTree)

	r, err := query.FindAll(testTree.RootNode(), "(tag) @inc")
	assert.Equal(t, len(*r), 8) // html + div*2 + ul + li*4

	r, err = query.FindAll(testTree.RootNode(), "(attribute_name) @inc")
	assert.Equal(t, len(*r), 3) // style + class + disabled

	r, err = query.FindAll(testTree.RootNode(), "(doctype) @inc")
	assert.Equal(t, len(*r), 1)
	exp := query.StrRange{
		StartPos: uint32(5),
		EndPos:   uint32(17),
		Range: &protocol.Range{
			Start: protocol.Position{
				Line:      1,
				Character: 4,
			},
			End: protocol.Position{
				Line:      1,
				Character: 16,
			},
		},
	}
	assert.Equal(t, *r, []*query.StrRange{
		&exp,
	})
}
