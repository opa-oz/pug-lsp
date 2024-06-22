package completion

import (
	"fmt"
	"strings"

	"github.com/opa-oz/go-todo/todo"
)

type CompletionKind = string

const (
	MixinKind         CompletionKind = "mixin"
	TagKind           CompletionKind = "tag"
	AttributeKind     CompletionKind = "attribute"
	DoctypeKind       CompletionKind = "doctype"
	KeywordKind       CompletionKind = "keyword"
	AndAttributesKind CompletionKind = "andAttributes"
)

type CompletionData struct {
	Kind CompletionKind
	Def  string
}

func DataToString(kind CompletionKind, def string) string {
	return fmt.Sprintf("%s|%s", kind, def)
}

func StringToData(str string) *CompletionData {
	parts := strings.Split(str, "|")
	compData := CompletionData{
		Kind: CompletionKind(parts[0]),
	}

	if len(parts) > 1 {
		compData.Def = parts[1]
	}

	if len(parts) > 2 {
		todo.T("Concat the rest of parts")
	}

	return &compData
}
