package server

import (
	"fmt"

	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/opa-oz/pug-lsp/pkg/lsp"
	"github.com/opa-oz/pug-lsp/pkg/query"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func ReturnMarkdown(desc string) *protocol.Hover {
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.MarkupKindMarkdown,
			Value: desc,
		},
	}
}

func (s *Server) TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	doc, ok := s.documentStore.Get(params.TextDocument.URI)

	if !ok {
		return nil, nil
	}

	node := doc.GetAtPosition(&params.Position)
	if node == nil {
		return nil, nil
	}

	tagRaw := (*doc.Content)[node.StartByte():node.EndByte()]
	s.logger.Println("TextDocumentHover", node.Type(), tagRaw)

	switch query.NodeType(node.Type()) {
	case query.TagNameNode:
		tag := html.HtmlTag(tagRaw)

		desc, ok := html.TagToDesc[tag]
		if !ok {
			return nil, nil
		}

		return ReturnMarkdown(desc), nil
	case query.MixinNameNode:
		for _, def := range doc.Mixins {
			if def.Name == tagRaw {
				return ReturnMarkdown(fmt.Sprintf("```pug\nmixin %s\n```", def.Definition)), nil
			}
		}

		for _, include := range doc.Includes {
			d, ok := s.documentStore.Get(*include.URI)
			if !ok {
				continue
			}

			for _, def := range d.Mixins {
				if def.Name == tagRaw {
					return ReturnMarkdown(fmt.Sprintf("```pug\nmixin %s\n```", def.Definition)), nil
				}
			}
		}
	case query.KeywordNode:
		keyword := lsp.Keyword(tagRaw)

		desc, ok := lsp.KeywordToDesc[keyword]
		if !ok {
			return nil, nil
		}

		return ReturnMarkdown(desc), nil
	case query.DoctypeNameNode:
		desc, ok := html.DoctypeToDesc[tagRaw]
		if !ok {
			return nil, nil
		}

		return ReturnMarkdown(desc), nil
	case query.AttributeNameNode:
		desc, ok := html.AttrToDesc[tagRaw]
		if !ok {
			return nil, nil
		}

		return ReturnMarkdown(desc), nil
	}

	return nil, nil
}
