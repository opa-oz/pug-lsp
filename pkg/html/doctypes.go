package html

var Doctypes = []string{
	"doctype html",
	"doctype xml",
	"doctype transitional",
	"doctype strict",
	"doctype frameset",
	"doctype 1.1",
	"doctype basic",
	"doctype mobile",
	"doctype plist",
}

var DoctypeToDesc = map[string]string{
	"html":         "```html\n<!DOCTYPE html>\n```",
	"xml":          "```xml\n<?xml version=\"1.0\" encoding=\"utf-8\" ?>\n```",
	"transitional": "```html\n<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">\n```",
	"strict":       "```html\n<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\">\n```",
	"frameset":     "```html\n<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Frameset//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd\">\n```",
	"1.1":          "```html\n<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\">\n```",
	"basic":        "```html\n<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML Basic 1.1//EN\" \"http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd\">\n```",
	"mobile":       "```html\n<!DOCTYPE html PUBLIC \"-//WAPFORUM//DTD XHTML Mobile 1.2//EN\" \"http://www.openmobilealliance.org/tech/DTD/xhtml-mobile12.dtd\">\n```",
	"plist":        "```html\n<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n```",
}
