package query

import (
	"github.com/opa-oz/pug-lsp/pkg/pug"
	sitter "github.com/smacker/go-tree-sitter"
)

type StrRange struct {
	StartPos uint32
	EndPos   uint32
}

func FindAll(node *sitter.Node, query string) (*[]*StrRange, error) {
	var includes []*StrRange

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
			includes = append(includes, &StrRange{StartPos: c.Node.StartByte(), EndPos: c.Node.EndByte()})
		}
	}

	return &includes, nil
}
