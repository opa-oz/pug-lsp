package lsp

import "strings"

type Include struct {
	Original *string
	Path     *string
	Prefix   *string // This is prefix for suggestions - if you `include mixins/logo`, you should import it as `+mixins.logo`
	URI      *string
}

func NewInclude(original, path *string) *Include {
	parts := strings.Split(*original, "/")
	var prefix string

	if len(parts) > 1 {
		prefix = parts[len(parts)-2]
	}

	uri := *path
	if !strings.HasPrefix(uri, "file:/") {
		uri = "file://" + uri
	}
	return &Include{
		Original: original,
		Path:     path,
		Prefix:   &prefix,
		URI:      &uri,
	}
}
