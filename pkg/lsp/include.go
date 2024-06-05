package lsp

import "strings"

type Include struct {
	Original *string
	Path     *string
	Prefix   *string // This is prefix for suggestions - if you `include mixins/logo`, you should import it as `+mixins.logo`
}

func NewInclude(original, path *string) *Include {
	parts := strings.Split(*original, "/")
	var prefix string

	if len(parts) > 1 {
		prefix = parts[len(parts)-2]
	}

	return &Include{
		Original: original,
		Path:     path,
		Prefix:   &prefix,
	}
}
