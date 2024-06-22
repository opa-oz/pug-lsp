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

// DataToString formats the given CompletionKind and string into a formatted string.
//
// DataToString takes a CompletionKind and a string representing a definition,
// and returns a formatted string where the kind and definition are concatenated
// with a '|' separator. The kind parameter specifies the type or category of completion,
// while the def parameter is the definition associated with the completion kind.
//
// Example:
//
//	kind := MixinKind
//	def := "myMixin(param)"
//	result := DataToString(kind, def) // returns "mixin|myMixin(param)"
//
// Parameters:
//
//	kind CompletionKind: The type or category of the completion.
//	def string: The definition associated with the completion kind.
//
// Returns:
//
//	string: A formatted string combining the kind and def parameters with a '|' separator.
func DataToString(kind CompletionKind, def string) string {
	return fmt.Sprintf("%s|%s", kind, def)
}

// StringToData parses the given string into a CompletionData structure.
//
// StringToData takes a string in a specific format (kind|definition) and splits
// it into a CompletionData structure. The first part before '|' is interpreted
// as the CompletionKind, and the second part (if present) is assigned to the Def field.
// If there are additional parts after the second '|', a placeholder message is logged.
//
// Example:
//
//	str := "mixin|myMixin(param)"
//	compData := StringToData(str)
//	// compData will be &CompletionData{ Kind: "mixin", Def: "myMixin(param)" }
//
// Parameters:
//
//	str string: The string to parse, formatted as "kind|definition".
//
// Returns:
//
//	*CompletionData: A pointer to a CompletionData structure containing parsed data.
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
