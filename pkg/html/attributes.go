package html

var allVisible = []string{
	"onblur",
	"onchange",
	"onclick",
	"oncontextmenu",
	"oncopy",
	"oncut",
	"ondblclick",
	"ondrag",
	"ondragend",
	"ondragenter",
	"ondragleave",
	"ondragover",
	"ondragstart",
	"ondrop",
	"onfocus",
	"oninput",
	"oninvalid",
	"onkeydown",
	"onkeypress",
	"onkeyup",
	"onmousedown",
	"onmousemove",
	"onmouseout",
	"onmouseover",
	"onmouseup",
	"onmousewheel",
	"onpaste",
	"onscroll",
	"onselect",
	"onwheel",
}

var tagToAttributes = map[HtmlTag]*[]string{
	Input: {
		"accept",
		"alt",
		"autocomplete",
		"autofocus",
		"checked",
		"dirname",
		"disabled",
		"form",
		"formaction",
		"height",
		"list",
		"max",
		"maxlength",
		"min",
		"multiple",
		"name",
		"onload",
		"onsearch",
		"pattern",
		"placeholder",
		"popovertarget",
		"popovertargetaction",
		"readonly",
		"required",
		"size",
		"src",
		"step",
		"type",
		"value",
		"width",
	},
	Form: {
		"accept-charset",
		"action",
		"autocomplete",
		"enctype",
		"method",
		"name",
		"novalidate",
		"onreset",
		"onsubmit",
		"rel",
		"target",
	},
	Area: {
		"alt",
		"coords",
		"download",
		"href",
		"hreflang",
		"media",
		"rel",
		"shape",
		"target",
	},
	Img: {
		"alt",
		"height",
		"ismap",
		"onabort",
		"onerror",
		"onload",
		"sizes",
		"src",
		"srcset",
		"usemap",
		"width",
	},
	Script: {
		"async",
		"charset",
		"defer",
		"onerror",
		"onload",
		"src",
		"type",
	},
	Button: {
		"autofocus",
		"disabled",
		"form",
		"formaction",
		"name",
		"popovertarget",
		"popovertargetaction",
		"type",
		"value",
	},
	Select: {
		"autofocus",
		"disabled",
		"form",
		"multiple",
		"name",
		"required",
		"size",
	},
	Textarea: {
		"autofocus",
		"cols",
		"dirname",
		"disabled",
		"form",
		"maxlength",
		"name",
		"placeholder",
		"readonly",
		"required",
		"rows",
		"wrap",
	},
	Audio: {
		"autoplay",
		"controls",
		"loop",
		"muted",
		"onabort",
		"oncanplay",
		"oncanplaythrough",
		"ondurationchange",
		"onemptied",
		"onended",
		"onerror",
		"onloadeddata",
		"onloadedmetadata",
		"onloadstart",
		"onpause",
		"onplay",
		"onplaying",
		"onprogress",
		"onratechange",
		"onseeked",
		"onseeking",
		"onstalled",
		"onsuspend",
		"ontimeupdate",
		"onvolumechange",
		"onwaiting",
		"preload",
		"src",
	},
	Video: {
		"autoplay",
		"controls",
		"height",
		"loop",
		"muted",
		"onabort",
		"oncanplay",
		"oncanplaythrough",
		"ondurationchange",
		"onemptied",
		"onended",
		"onerror",
		"onloadeddata",
		"onloadedmetadata",
		"onloadstart",
		"onpause",
		"onplay",
		"onplaying",
		"onprogress",
		"onratechange",
		"onseeked",
		"onseeking",
		"onstalled",
		"onsuspend",
		"ontimeupdate",
		"onvolumechange",
		"onwaiting",
		"poster",
		"preload",
		"src",
		"width",
	},
	Meta: {
		"charset",
		"content",
		"http-equiv",
		"name",
	},
	Blockquote: {
		"cite",
	},
	Del: {
		"cite",
		"datetime",
	},
	Ins: {
		"cite",
		"datetime",
	},
	Q: {
		"cite",
	},
	Td: {
		"colspan",
		"headers",
		"rowspan",
	},
	Th: {
		"colspan",
		"headers",
		"rowspan",
		"scope",
	},
	Object: {
		"data",
		"form",
		"height",
		"name",
		"onabort",
		"oncanplay",
		"onerror",
		"type",
		"usemap",
		"width",
	},
	Time: {
		"datetime",
	},
	Track: {
		"default",
		"kind",
		"label",
		"oncuechange",
		"src",
		"srclang",
	},
	Fieldset: {
		"disabled",
		"form",
		"name",
	},
	OptGroup: {
		"disabled",
		"label",
	},
	Option: {
		"disabled",
		"label",
		"selected",
		"value",
	},
	A: {
		"download",
		"href",
		"hreflang",
		"media",
		"rel",
		"target",
		"type",
	},
	Label: {
		"for",
		"form",
	},
	Output: {
		"for",
		"form",
		"name",
	},
	Meter: {
		"form",
		"high",
		"low",
		"max",
		"min",
		"optimum",
		"value",
	},
	Canvas: {
		"height",
		"width",
	},
	Embed: {
		"height",
		"onabort",
		"oncanplay",
		"onerror",
		"src",
		"type",
		"width",
	},
	IFrame: {
		"height",
		"name",
		"onload",
		"sandbox",
		"src",
		"srcdoc",
		"width",
	},
	Base: {
		"href",
		"target",
	},
	Link: {
		"href",
		"hreflang",
		"media",
		"onload",
		"rel",
		"sizes",
		"type",
	},
	Progress: {
		"max",
		"value",
	},
	Source: {
		"media",
		"sizes",
		"src",
		"srcset",
		"type",
	},
	Style: {
		"media",
		"onerror",
		"onload",
		"type",
	},
	Map: {
		"name",
	},
	Param: {
		"name",
		"value",
	},
	Body: {
		"onafterprint",
		"onbeforeprint",
		"onbeforeunload",
		"onerror",
		"onhashchange",
		"onload",
		"onoffline",
		"ononline",
		"onpagehide",
		"onpageshow",
		"onpopstate",
		"onresize",
		"onstorage",
		"onunload",
	},
	Details: {
		"ontoggle",
		"open",
	},
	Ol: {
		"reversed",
		"start",
	},
	Col: {
		"span",
	},
	ColGroup: {
		"span",
	},
	Menu: {
		"type",
	},
	Li: {
		"value",
	},
	Div:     &allVisible,
	Span:    &allVisible,
	H1:      &allVisible,
	H2:      &allVisible,
	H3:      &allVisible,
	H4:      &allVisible,
	H5:      &allVisible,
	H6:      &allVisible,
	P:       &allVisible,
	Pre:     &allVisible,
	Table:   &allVisible,
	HR:      &allVisible,
	Head:    &allVisible,
	Header:  &allVisible,
	Footer:  &allVisible,
	Article: &allVisible,
	Main:    &allVisible,
	Title:   &allVisible,
	Tr:      &allVisible,
	Html:    &allVisible,
	Ul:      &allVisible,
}

var AttrToDesc = map[string]string{
	"accept":              "Specifies the types of files that the server accepts (only for `type=\"file\"`)",
	"accept-charset":      "Specifies the character encodings that are to be used for the form submission",
	"accesskey":           "Specifies a shortcut key to activate/focus an element",
	"action":              "Specifies where to send the form-data when a form is submitted",
	"align":               "Specifies the alignment according to surrounding elements. Use CSS instead",
	"alt":                 "Specifies an alternate text when the original element fails to display",
	"async":               "Specifies that the script is executed asynchronously (only for external scripts)",
	"autocomplete":        "Specifies whether the `<form>` or the `<input>` element should have autocomplete enabled",
	"autofocus":           "Specifies that the element should automatically get focus when the page loads",
	"autoplay":            "Specifies that the audio/video will start playing as soon as it is ready",
	"bgcolor":             "Specifies the background color of an element. Use CSS instead",
	"border":              "Specifies the width of the border of an element. Use CSS instead",
	"charset":             "Specifies the character encoding",
	"checked":             "Specifies that an `<input>` element should be pre-selected when the page loads (for `type=\"checkbox\"` or `type=\"radio\"`)",
	"cite":                "Specifies a URL which explains the quote/deleted/inserted text",
	"class":               "Specifies one or more class names for an element (refers to a class in a style sheet)",
	"color":               "Specifies the text color of an element. Use CSS instead",
	"cols":                "Specifies the visible width of a text area",
	"colspan":             "Specifies the number of columns a table cell should span",
	"content":             "Gives the value associated with the http-equiv or name attribute",
	"contenteditable":     "Specifies whether the content of an element is editable or not",
	"controls":            "Specifies that audio/video controls should be displayed (such as a play/pause button etc.)",
	"coords":              "Specifies the coordinates of the area",
	"data":                "Specifies the URL of the resource to be used by the object",
	"data-*":              "Used to store custom data private to the page or application",
	"datetime":            "Specifies the date and time",
	"default":             "Specifies that the track is to be enabled if the user's preferences do not indicate that another track would be more appropriate",
	"defer":               "Specifies that the script is executed when the page has finished parsing (only for external scripts)",
	"dir":                 "Specifies the text direction for the content in an element",
	"dirname":             "Specifies that the text direction will be submitted",
	"disabled":            "Specifies that the specified element/group of elements should be disabled",
	"download":            "Specifies that the target will be downloaded when a user clicks on the hyperlink",
	"draggable":           "Specifies whether an element is draggable or not",
	"enctype":             "Specifies how the form-data should be encoded when submitting it to the server (only for `method=\"post\"`)",
	"enterkeyhint":        "Specifies the text of the enter-key on a virtual keyboard",
	"for":                 "Specifies which form element(s) a label/calculation is bound to",
	"form":                "Specifies the name of the form the element belongs to",
	"formaction":          "Specifies where to send the form-data when a form is submitted. Only for `type=\"submit\"`",
	"headers":             "Specifies one or more headers cells a cell is related to",
	"height":              "Specifies the height of the element",
	"hidden":              "Specifies that an element is not yet, or is no longer, relevant",
	"high":                "Specifies the range that is considered to be a high value",
	"href":                "Specifies the URL of the page the link goes to",
	"hreflang":            "Specifies the language of the linked document",
	"http-equiv":          "Provides an HTTP header for the information/value of the content attribute",
	"id":                  "Specifies a unique id for an element",
	"inert":               "Specifies that the browser should ignore this section",
	"inputmode":           "Specifies the mode of a virtual keyboard",
	"ismap":               "Specifies an image as a server-side image map",
	"kind":                "Specifies the kind of text track",
	"label":               "Specifies the title of the text track",
	"lang":                "Specifies the language of the element's content",
	"list":                "Refers to a `<datalist>` element that contains pre-defined options for an `<input>` element",
	"loop":                "Specifies that the audio/video will start over again, every time it is finished",
	"low":                 "Specifies the range that is considered to be a low value",
	"max":                 "Specifies the maximum value",
	"maxlength":           "Specifies the maximum number of characters allowed in an element",
	"media":               "Specifies what media/device the linked document is optimized for",
	"method":              "Specifies the HTTP method to use when sending form-data",
	"min":                 "Specifies a minimum value",
	"multiple":            "Specifies that a user can enter more than one value",
	"muted":               "Specifies that the audio output of the video should be muted",
	"name":                "Specifies the name of the element",
	"novalidate":          "Specifies that the form should not be validated when submitted",
	"onabort":             "Script to be run on abort",
	"onafterprint":        "Script to be run after the document is printed",
	"onbeforeprint":       "Script to be run before the document is printed",
	"onbeforeunload":      "Script to be run when the document is about to be unloaded",
	"onblur":              "Script to be run when the element loses focus",
	"oncanplay":           "Script to be run when a file is ready to start playing (when it has buffered enough to begin)",
	"oncanplaythrough":    "Script to be run when a file can be played all the way to the end without pausing for buffering",
	"onchange":            "Script to be run when the value of the element is changed",
	"onclick":             "Script to be run when the element is being clicked",
	"oncontextmenu":       "Script to be run when a context menu is triggered",
	"oncopy":              "Script to be run when the content of the element is being copied",
	"oncuechange":         "Script to be run when the cue changes in a `<track>` element",
	"oncut":               "Script to be run when the content of the element is being cut",
	"ondblclick":          "Script to be run when the element is being double-clicked",
	"ondrag":              "Script to be run when the element is being dragged",
	"ondragend":           "Script to be run at the end of a drag operation",
	"ondragenter":         "Script to be run when an element has been dragged to a valid drop target",
	"ondragleave":         "Script to be run when an element leaves a valid drop target",
	"ondragover":          "Script to be run when an element is being dragged over a valid drop target",
	"ondragstart":         "Script to be run at the start of a drag operation",
	"ondrop":              "Script to be run when dragged element is being dropped",
	"ondurationchange":    "Script to be run when the length of the media changes",
	"onemptied":           "Script to be run when something bad happens and the file is suddenly unavailable (like unexpectedly disconnects)",
	"onended":             "Script to be run when the media has reach the end (a useful event for messages like \"thanks for listening\")",
	"onerror":             "Script to be run when an error occurs",
	"onfocus":             "Script to be run when the element gets focus",
	"onhashchange":        "Script to be run when there has been changes to the anchor part of the a URL",
	"oninput":             "Script to be run when the element gets user input",
	"oninvalid":           "Script to be run when the element is invalid",
	"onkeydown":           "Script to be run when a user is pressing a key",
	"onkeypress":          "Script to be run when a user presses a key",
	"onkeyup":             "Script to be run when a user releases a key",
	"onload":              "Script to be run when the element is finished loading",
	"onloadeddata":        "Script to be run when media data is loaded",
	"onloadedmetadata":    "Script to be run when meta data (like dimensions and duration) are loaded",
	"onloadstart":         "Script to be run just as the file begins to load before anything is actually loaded",
	"onmousedown":         "Script to be run when a mouse button is pressed down on an element",
	"onmousemove":         "Script to be run as long as the  mouse pointer is moving over an element",
	"onmouseout":          "Script to be run when a mouse pointer moves out of an element",
	"onmouseover":         "Script to be run when a mouse pointer moves over an element",
	"onmouseup":           "Script to be run when a mouse button is released over an element",
	"onmousewheel":        "Script to be run when a mouse wheel is being scrolled over an element",
	"onoffline":           "Script to be run when the browser starts to work offline",
	"ononline":            "Script to be run when the browser starts to work online",
	"onpagehide":          "Script to be run when a user navigates away from a page",
	"onpageshow":          "Script to be run when a user navigates to a page",
	"onpaste":             "Script to be run when the user pastes some content in an element",
	"onpause":             "Script to be run when the media is paused either by the user or programmatically",
	"onplay":              "Script to be run when the media has started playing",
	"onplaying":           "Script to be run when the media has started playing",
	"onpopstate":          "Script to be run when the window's history changes.",
	"onprogress":          "Script to be run when the browser is in the process of getting the media data",
	"onratechange":        "Script to be run each time the playback rate changes (like when a user switches to a slow motion or fast forward mode).",
	"onreset":             "Script to be run when a reset button in a form is clicked.",
	"onresize":            "Script to be run when the browser window is being resized.",
	"onscroll":            "Script to be run when an element's scrollbar is being scrolled",
	"onsearch":            "Script to be run when the user writes something in a search field (for `<input type=\"search\">`)",
	"onseeked":            "Script to be run when the seeking attribute is set to false indicating that seeking has ended",
	"onseeking":           "Script to be run when the seeking attribute is set to true indicating that seeking is active",
	"onselect":            "Script to be run when the element gets selected",
	"onstalled":           "Script to be run when the browser is unable to fetch the media data for whatever reason",
	"onstorage":           "Script to be run when a Web Storage area is updated",
	"onsubmit":            "Script to be run when a form is submitted",
	"onsuspend":           "Script to be run when fetching the media data is stopped before it is completely loaded for whatever reason",
	"ontimeupdate":        "Script to be run when the playing position has changed (like when the user fast forwards to a different point in the media)",
	"ontoggle":            "Script to be run when the user opens or closes the `<details>` element",
	"onunload":            "Script to be run when a page has unloaded (or the browser window has been closed)",
	"onvolumechange":      "Script to be run each time the volume of a video/audio has been changed",
	"onwaiting":           "Script to be run when the media has paused but is expected to resume (like when the media pauses to buffer more data)",
	"onwheel":             "Script to be run when the mouse wheel rolls up or down over an element",
	"open":                "Specifies that the details should be visible (open) to the user",
	"optimum":             "Specifies what value is the optimal value for the gauge",
	"pattern":             "Specifies a regular expression that an `<input>` element's value is checked against",
	"placeholder":         "Specifies a short hint that describes the expected value of the element",
	"popover":             "Specifies a popover element",
	"popovertarget":       "Specifies which popover element to invoked",
	"popovertargetaction": "Specifies what happens to the popover element when the button is clicked",
	"poster":              "Specifies an image to be shown while the video is downloading, or until the user hits the play button",
	"preload":             "Specifies if and how the author thinks the audio/video should be loaded when the page loads",
	"readonly":            "Specifies that the element is read-only",
	"rel":                 "Specifies the relationship between the current document and the linked document",
	"required":            "Specifies that the element must be filled out before submitting the form",
	"reversed":            "Specifies that the list order should be descending (9,8,7...)",
	"rows":                "Specifies the visible number of lines in a text area",
	"rowspan":             "Specifies the number of rows a table cell should span",
	"sandbox":             "Enables an extra set of restrictions for the content in an `<iframe>`",
	"scope":               "Specifies whether a header cell is a header for a column, row, or group of columns or rows",
	"selected":            "Specifies that an option should be pre-selected when the page loads",
	"shape":               "Specifies the shape of the area",
	"size":                "Specifies the width, in characters (for `<input>`) or specifies the number of visible options (for `<select>`)",
	"sizes":               "Specifies the size of the linked resource",
	"span":                "Specifies the number of columns to span",
	"spellcheck":          "Specifies whether the element is to have its spelling and grammar checked or not",
	"src":                 "Specifies the URL of the media file",
	"srcdoc":              "Specifies the HTML content of the page to show in the `<iframe>`",
	"srclang":             "Specifies the language of the track text data (required if `kind=\"subtitles\"`)",
	"srcset":              "Specifies the URL of the image to use in different situations",
	"start":               "Specifies the start value of an ordered list",
	"step":                "Specifies the legal number intervals for an input field",
	"style":               "Specifies an inline CSS style for an element",
	"tabindex":            "Specifies the tabbing order of an element",
	"target":              "Specifies the target for where to open the linked document or where to submit the form",
	"title":               "Specifies extra information about an element",
	"translate":           "Specifies whether the content of an element should be translated or not",
	"type":                "Specifies the type of element",
	"usemap":              "Specifies an image as a client-side image map",
	"value":               "Specifies the value of the element",
	"width":               "Specifies the width of the element",
	"wrap":                "Specifies how the text in a text area is to be wrapped when submitted in a form",
}

func GetAttributes(tagName string) *[]string {
	arr, ok := tagToAttributes[HtmlTag(tagName)]

	if !ok {
		return nil
	}

	return arr
}

func GetTags() *[]HtmlTag {
	keys := make([]HtmlTag, 0, len(tagToAttributes))
	for k := range tagToAttributes {
		htmlK := k
		keys = append(keys, htmlK)
	}

	return &keys
}
