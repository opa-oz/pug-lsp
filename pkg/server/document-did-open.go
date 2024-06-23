package server

import (
	"context"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *Server) TextDocumentDidOpen(ctx *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	doc, err := s.documentStore.DocumentDidOpen(context.Background(), *params)

	if err != nil {
		return err
	}
	s.documentStore.RefreshIncludes(context.Background(), doc)
	go func() {
		diags := s.documentStore.RefreshDiagnostics(doc, false)
		if diags != nil && len(*diags) > 0 {
			go ctx.Notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
				URI:         doc.URI,
				Diagnostics: *diags,
			})
		}
	}()

	return nil
}
