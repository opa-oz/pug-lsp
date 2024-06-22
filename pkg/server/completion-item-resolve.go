package server

import (
	"fmt"

	"github.com/opa-oz/pug-lsp/pkg/completion"
	"github.com/opa-oz/pug-lsp/pkg/html"
	"github.com/opa-oz/pug-lsp/pkg/lsp"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

const andAttributesDoc = "Pronounced as “and attributes”, the `&attributes` syntax can be used to explode an object into attributes of an element.\n```pug\ndiv#foo(data-bar=\"foo\")&attributes({'data-foo': 'bar'})\n```"

func (s *Server) CompletionItemResolve(context *glsp.Context, params *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	compDataRaw, ok := params.Data.(string)

	if ok {
		compData := completion.StringToData(compDataRaw)

		if compData.Kind == completion.MixinKind {
			params.Documentation = protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: fmt.Sprintf("```pug\nmixin %s\n```", compData.Def),
			}
		} else if compData.Kind == completion.AttributeKind {
			desc, ok := html.AttrToDesc[compData.Def]
			if ok {
				params.Documentation = protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: desc,
				}
			}
		} else if compData.Kind == completion.AndAttributesKind {
			params.Detail = &compData.Def
			params.Documentation = protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: andAttributesDoc,
			}
		} else if compData.Kind == completion.KeywordKind {
			key := lsp.Keyword(compData.Def)
			desc, ok := lsp.KeywordToDesc[key]

			if ok {
				params.Documentation = protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: desc,
				}
			} else {
				params.Detail = &compData.Def
			}
		} else if compData.Kind == completion.TagKind {
			tag := html.HtmlTag(compData.Def)
			desc, ok := html.TagToDesc[tag]
			if ok {
				params.Documentation = protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: desc,
				}
			} else {
				params.Detail = &compData.Def
			}
		}
	}

	return params, nil
}
