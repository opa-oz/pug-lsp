package server

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *Server) initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}
