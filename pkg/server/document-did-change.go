package server

import (
	"context"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *Server) TextDocumentDidChange(ctx *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	doc, ok := s.documentStore.Get(params.TextDocument.URI)

	if !ok {
		return nil
	}
	s.logger.Println("DocDidChange", doc.Path)
	doc.ApplyChanges(context.Background(), params.ContentChanges)
	// todo: refresh diagnostics maybe?
	return nil
}
