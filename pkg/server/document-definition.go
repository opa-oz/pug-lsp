package server

import (
	"strings"

	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (s *Server) TextDocumentDefinition(context *glsp.Context, params *protocol.DefinitionParams) (interface{}, error) {
	doc, ok := s.documentStore.Get(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}

	node := doc.GetAtPosition(&params.Position)
	if node == nil {
		return nil, nil
	}

	nodeType := query.NodeType(node.Type())
	if nodeType == query.FilenameNode && (node.Parent() != nil && query.NodeType(node.Parent().Type()) == query.IncludeNode) {
		targetFilename := strings.Trim((*doc.Content)[node.StartByte():node.EndByte()], " ")
		includes := doc.Includes

		for _, include := range includes {
			if *include.Original == targetFilename {
				return protocol.Location{
					URI: *include.URI,
				}, nil
			}
		}
	}

	return nil, nil
}
