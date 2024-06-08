package html

import "fmt"

var customDetails = map[string]string{
	"style":    "style={ }",
	"checked":  "checked",
	"disabled": "disabled",
	"hidden":   "hidden",
	"class":    "class=[ ]",
	"click":    "'(click)'",
}

func GetInsertText(attribute string) *string {
	details, ok := customDetails[attribute]
	if !ok {
		details = fmt.Sprintf("%s=\" \"", attribute)
		return &details
	}

	return &details
}
