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
	err := doc.ApplyChanges(context.Background(), params.ContentChanges)
	if err != nil {
		s.logger.Err(err)
		return nil
	}

	if !doc.NeedToRefreshIncludes {
		doc.NeedToRefreshIncludes = true

		s.documentStore.RefreshIncludes(context.Background(), doc, false)
	}

	if !doc.NeedToRefreshDiagnostics {
		doc.NeedToRefreshDiagnostics = true
		go func() {
			diags := s.documentStore.RefreshDiagnostics(doc, true)
			go ctx.Notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
				URI:         doc.URI,
				Diagnostics: *diags,
			})
		}()
	}

	return nil
}
