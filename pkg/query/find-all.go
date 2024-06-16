package query

import (
	"github.com/opa-oz/pug-lsp/pkg/pug"
	sitter "github.com/smacker/go-tree-sitter"
)

type StrRange struct {
	StartPos uint32
	EndPos   uint32
}

func FindAll(node *sitter.Node, query Query) (*[]*StrRange, error) {
	var nodes []*StrRange

	ns, err := FindAllNodes(node, query)

	if err != nil {
		return nil, err
	}

	for _, n := range *ns {
		nodes = append(nodes, &StrRange{StartPos: n.StartByte(), EndPos: n.EndByte()})
	}

	return &nodes, nil
}

func FindAllNodes(node *sitter.Node, query Query) (*[]*sitter.Node, error) {
	var nodes []*sitter.Node

	q, err := sitter.NewQuery([]byte(query), pug.GetLanguage())
	if err != nil {
		return nil, err
	}

	qs := sitter.NewQueryCursor()
	qs.Exec(q, node)

	for {
		m, ok := qs.NextMatch()
		if !ok {
			break
		}

		for _, c := range m.Captures {
			nodes = append(nodes, c.Node)
		}
	}

	return &nodes, nil
}
