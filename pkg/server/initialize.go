package server

import (
	"github.com/opa-oz/pug-lsp/pkg/utils"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *Server) initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	s.clientCapabilities = params.Capabilities

	capabilities := s.handler.CreateServerCapabilities()
	capabilities.CompletionProvider = &protocol.CompletionOptions{
		TriggerCharacters: []string{"+", "(", ",", "&"},
		ResolveProvider:   utils.PtrBool(true),
	}

	capabilities.HoverProvider = true

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    s.opts.Name,
			Version: &s.opts.Version,
		},
	}, nil
}
