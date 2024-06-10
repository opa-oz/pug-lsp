package completion

import (
	"github.com/opa-oz/pug-lsp/pkg/documents"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/opa-oz/pug-lsp/pkg/utils"
	sitter "github.com/smacker/go-tree-sitter"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CompletionMetaParams struct {
	Doc           *documents.Document
	Params        *protocol.CompletionParams
	ExistingAttrs *query.ExistingAttributes
	CurrentNode   *sitter.Node
	Logger        *utils.Logger
}
