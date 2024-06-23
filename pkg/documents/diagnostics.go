package documents

import (
	"fmt"

	"github.com/opa-oz/pug-lsp/pkg/query"
	sitter "github.com/smacker/go-tree-sitter"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

const source = "pug-lsp"

func NewDiagnostic(node *sitter.Node, message string) protocol.Diagnostic {
	severity := protocol.DiagnosticSeverityError
	sourceCp := source
	return protocol.Diagnostic{
		Severity: &severity,
		Source:   &sourceCp,
		Message:  message,
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      node.StartPoint().Row,
				Character: node.StartPoint().Column,
			},
			End: protocol.Position{
				Line:      node.EndPoint().Row,
				Character: node.EndPoint().Column,
			},
		},
	}
}
func NewDiagnosticFromRange(rng *protocol.Range, message string) protocol.Diagnostic {
	severity := protocol.DiagnosticSeverityError
	sourceCp := source
	return protocol.Diagnostic{
		Severity: &severity,
		Source:   &sourceCp,
		Message:  message,
		Range:    *rng,
	}
}

func DiagnostMixins(doc *Document, ds *DocumentStore) *[]protocol.Diagnostic {
	diags := []protocol.Diagnostic{}

	// non-existing mixin detect
	mixinUse, err := query.FindAllNodes(doc.Tree.RootNode(), query.MixinUseQ)
	if err != nil || len(*mixinUse) == 0 {
		return &diags
	}

	mixinsFlat := ds.FlatMixins(doc)
	existingMixins := make(MixinsMap)
	for _, mixin := range *mixinsFlat {
		if mixin == nil {
			continue
		}
		existingMixins[mixin.Name] = mixin
	}

	for _, mixin := range *mixinUse {
		nameNode := query.FindDownwards(mixin, query.MixinNameNode, 2)
		if nameNode == nil {
			continue
		}
		name := (*doc.Content)[nameNode.StartByte():nameNode.EndByte()]

		origin, ok := existingMixins[name]
		if !ok {
			diags = append(diags, NewDiagnostic(nameNode, fmt.Sprintf("Mixin `%s` is not defined", name)))
		} else {
			if len(*origin.Arguments) > 0 {
				attributesRanges, err := query.FindAll(mixin, query.AttributesQ)
				if err != nil {
					diags = append(diags, NewDiagnostic(nameNode, fmt.Sprintf("Arguments list is empty")))
				}
				if len(*origin.Arguments) > len(*attributesRanges) {
					diags = append(diags, NewDiagnostic(nameNode, fmt.Sprintf("Not enough arguments for mixin `%s`\n\twant: `%s`", name, origin.Definition)))
				} else if len(*origin.Arguments) < len(*attributesRanges) {
					diags = append(diags, NewDiagnostic(nameNode, fmt.Sprintf("Extra arguments found for mixin `%s`\n\twant: `%s`", name, origin.Definition)))
				}
			}
		}
	}

	return &diags
}

func DiagnostIncludes(doc *Document) *[]protocol.Diagnostic {
	diags := []protocol.Diagnostic{}

	for _, include := range doc.Includes {
		if !include.IsValid {
			diags = append(diags, NewDiagnosticFromRange(include.Range, fmt.Sprintf("Couldn't locate `%s` in filesystem", *include.Original)))
		}
	}

	return &diags
}
