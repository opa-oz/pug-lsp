package query

import (
	"github.com/opa-oz/pug-lsp/pkg/pug"
	sitter "github.com/smacker/go-tree-sitter"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type StrRange struct {
	StartPos uint32 // Starting position (inclusive)
	EndPos   uint32 // Ending position (exclusive)
	Range    *protocol.Range
}

// FindAll searches for nodes matching the specified query within the subtree rooted at the given node.
//
// It uses the FindAllNodes function to retrieve nodes that match the provided query,
// and converts their byte offsets into StrRange objects indicating the start and end positions.
//
// Parameters:
//
//	node *sitter.Node - The root node from which to begin searching.
//	query Query - The query to match nodes against.
//
// Returns:
//
//	*[]*StrRange - A pointer to a slice of StrRange pointers representing the byte ranges
//	               of nodes that match the query. Returns nil and an error if there's
//	               any issue during node retrieval.
func FindAll(node *sitter.Node, query Query) (*[]*StrRange, error) {
	var nodes []*StrRange

	ns, err := FindAllNodes(node, query)

	if err != nil {
		return nil, err
	}

	for _, n := range *ns {
		nodes = append(nodes, &StrRange{Range: &protocol.Range{
			Start: protocol.Position{
				Line:      n.StartPoint().Row,
				Character: n.StartPoint().Column,
			},
			End: protocol.Position{
				Line:      n.EndPoint().Row,
				Character: n.EndPoint().Column,
			},
		}, StartPos: n.StartByte(), EndPos: n.EndByte()})
	}

	return &nodes, nil
}

// FindAllNodes finds all nodes matching the specified query within the subtree rooted at the given node.
//
// It compiles the query using sitter.NewQuery, initializes a query cursor, and executes the query
// on the node using qs.Exec. It then iterates through each match in the query cursor and collects
// the corresponding nodes into a slice.
//
// Parameters:
//
//	node *sitter.Node - The root node from which to begin searching.
//	query Query - The query to match nodes against.
//
// Returns:
//
//	*[]*sitter.Node - A pointer to a slice of pointers to nodes that match the query.
//	                  Returns nil and an error if there's any issue during node retrieval.
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
