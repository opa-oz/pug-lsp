package server

import (
	"context"

	"github.com/opa-oz/go-todo/todo"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *Server) TextDocumentDidChange(ctx *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	doc, ok := s.documentStore.Get(params.TextDocument.URI)

	if !ok {
		return nil
	}
	err := doc.ApplyChanges(context.Background(), params.ContentChanges)
	if err != nil {
		s.logger.Err(err)
		return nil
	}
	todo.T("Refresh diagnostics")

	return nil
}
