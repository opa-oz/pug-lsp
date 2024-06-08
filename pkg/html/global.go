package html

var globalAttributes = []string{
	"accesskey",
	"class",
	"contenteditable",
	"data-",
	"dir",
	"draggable",
	"enterkeyhint",
	"hidden",
	// "id", // commented because ids in Pug are done via `#`
	"inert",
	"inputmode",
	"lang",
	"popover",
	"spellcheck",
	"style",
	"tabindex",
	"title",
	"translate",
}

func GlobalAttrs() *[]string {
	return &globalAttributes
}
