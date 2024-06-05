package query

import (
	"github.com/opa-oz/pug-lsp/pkg/pug"
	sitter "github.com/smacker/go-tree-sitter"
)

const query = "(include (filename) @incl)"

type StrRange struct {
	StartPos uint32
	EndPos   uint32
}

func FindIncludes(tree *sitter.Tree) (*[]*StrRange, error) {
	var includes []*StrRange

	q, err := sitter.NewQuery([]byte(query), pug.GetLanguage())
	if err != nil {
		return nil, err
	}

	n := tree.RootNode()
	qs := sitter.NewQueryCursor()

	qs.Exec(q, n)

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
