package server

import (
	"strings"

	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/opa-oz/pug-lsp/pkg/utils"
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

	if nodeType == query.MixinNameNode && (node.Parent() != nil && query.NodeType(node.Parent().Type()) == query.MixinNode) {
		mixins := s.documentStore.FlatMixins(doc)
		targetMixin := strings.Trim((*doc.Content)[node.StartByte():node.EndByte()], " ")

		for _, mixin := range *mixins {
			if mixin.Name == targetMixin {
				if utils.PtrIsTrue(s.clientCapabilities.TextDocument.Definition.LinkSupport) {
					return protocol.LocationLink{
						TargetURI:            mixin.Source,
						OriginSelectionRange: mixin.Range,
					}, nil
				} else {
					return protocol.Location{
						URI: *&mixin.Source,
					}, nil
				}
			}
		}
	}

	return nil, nil
}
