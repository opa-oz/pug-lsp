package server

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *Server) initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	s.clientCapabilities = params.Capabilities

	ptrTrue := true

	capabilities := s.handler.CreateServerCapabilities()
	capabilities.CompletionProvider = &protocol.CompletionOptions{
		TriggerCharacters: []string{"+", "("},
		ResolveProvider:   &ptrTrue,
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    s.opts.Name,
			Version: &s.opts.Version,
		},
	}, nil
}
