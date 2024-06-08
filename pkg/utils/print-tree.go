package utils

import (
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
)

func printDepth(t *sitter.Node, depth int, logger Logger) {
	logger.Println(fmt.Sprintf("%"+fmt.Sprint(depth)+"s, (%d:%d)", t.Type(), t.StartPoint().Row, t.StartPoint().Column))
	if t.ChildCount() > 0 {
		for idx := 0; idx < int(t.ChildCount()); idx++ {
			printDepth(t.Child(idx), depth+1, logger)
		}
	}
}

func PrintTree(tree *sitter.Node, logger Logger) {
	printDepth(tree, 0, logger)
}
