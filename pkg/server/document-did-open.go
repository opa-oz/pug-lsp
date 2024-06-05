package server

import (
	"context"

	"github.com/opa-oz/go-todo/todo"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *Server) TextDocumentDidOpen(ctx *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	doc, err := s.documentStore.DocumentDidOpen(context.Background(), *params, ctx.Notify)

	if err != nil {
		return err
	}

	todo.T("Refresh diagnostics", doc.Path)

	s.documentStore.RefreshIncludes(context.Background(), doc)

	return nil
}
