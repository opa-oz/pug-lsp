package query_test

import (
	"context"
	"testing"

	"github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestFindUpwardsAndDownwardsOK(t *testing.T) {
	testTree, err := pug.GetParsedTreeFromString(context.TODO(), `
        doctype html
        html
            body
                div(title="My title")
    `)
	// logger := utils.NewStdLogger("pug-lsp", 2)
	// utils.PrintTree(testTree.RootNode(), logger)

	assert.NoError(t, err)

	attribute := query.FindDownwards(testTree.RootNode(), query.AttributeNameNode, 9) // it's exactly 9 steps needed
	assert.NotNil(t, attribute)
	assert.Equal(t, attribute.Type(), string(query.AttributeNameNode))

	bodyOrDiv := query.FindUpwards(attribute, query.TagNode, 8)
	assert.NotNil(t, bodyOrDiv)
	assert.Equal(t, bodyOrDiv.Type(), string(query.TagNode))
}

func TestMaxIterationsOK(t *testing.T) {
	testTree, err := pug.GetParsedTreeFromString(context.TODO(), `
        doctype html
        html
            body
                div(title="My title")
    `)
	assert.NoError(t, err)

	attribute := query.FindDownwards(testTree.RootNode(), query.AttributeNameNode, 2)
	assert.Nil(t, attribute)

	realAttribute := query.FindDownwards(testTree.RootNode(), query.AttributeNameNode, 9) // it's exactly 9 steps needed
	assert.NotNil(t, realAttribute)

	bodyOrDiv := query.FindUpwards(realAttribute, query.TagNode, 1)
	assert.Nil(t, bodyOrDiv)
}

func TestZeroDepth(t *testing.T) {

	testTree, err := pug.GetParsedTreeFromString(context.TODO(), `
    div
    `)
	assert.NoError(t, err)
	assert.Equal(t, testTree.RootNode().ChildCount(), uint32(1))
	firstChild := testTree.RootNode().Child(0)

	assert.Equal(t, firstChild.Type(), string(query.TagNode))
	assert.True(
		t,
		firstChild.Equal(
			query.FindUpwards(firstChild, query.TagNode, 5),
		),
	)

	assert.True(
		t,
		firstChild.Equal(
			query.FindDownwards(firstChild, query.TagNode, 5),
		),
	)
}

func TestNilCases(t *testing.T) {
	assert.Nil(t, query.FindUpwards(nil, query.TagNode, 5))
	assert.Nil(t, query.FindDownwards(nil, query.TagNode, 5))
	assert.Nil(t, query.FindDownwards(nil, query.TagNode, 0))
}
