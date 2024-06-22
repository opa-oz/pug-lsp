package html

import "fmt"

var customDetails = map[string]string{
	"style":     "style={}",
	"checked":   "checked",
	"disabled":  "disabled",
	"hidden":    "hidden",
	"class":     "class=[]",
	"click":     "'(click)'",
	"draggable": "draggable",
}

// GetInsertText retrieves the insert text for a given attribute from customDetails.
//
// GetInsertText checks if the specified attribute exists in customDetails map.
// If found, it returns a pointer to the insert text associated with the attribute.
// If not found, it creates a default insert text with the attribute name and returns
// a pointer to it.
//
// Example:
//
//	attribute := "href"
//	insertText := GetInsertText(attribute)
//	// insertText will be a pointer to the string "href=\"\"" if 'attribute' is not found in customDetails.
//
// Parameters:
//
//	attribute string: The attribute name for which insert text is requested.
//
// Returns:
//
//	*string: A pointer to the insert text associated with the attribute or a default insert text.
func GetInsertText(attribute string) *string {
	details, ok := customDetails[attribute]
	if !ok {
		details = fmt.Sprintf("%s=\"\"", attribute)
		return &details
	}

	return &details
}
