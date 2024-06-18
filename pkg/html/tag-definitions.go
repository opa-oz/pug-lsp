package html

var TagToDesc = map[HtmlTag]string{
	Body:   "The `<body>` element contains all the contents of an HTML document, such as headings, paragraphs, images, hyperlinks, tables, lists, etc.\n\n*Note:* There can only be one `<body>` element in an HTML document.",
	Div:    "The `<div>` tag defines a division or a section in an HTML document.\nThe `<div>` tag is used as a container for HTML elements - which is then styled with CSS or manipulated with JavaScript.\nThe `<div>` tag is easily styled by using the class or id attribute.\nAny sort of content can be put inside the `<div>` tag! \n\n*Note:* By default, browsers always place a line break before and after the `<div>` element.",
	Script: "The `<script>` tag is used to embed a client-side script (JavaScript).\nThe `<script>` element either contains scripting statements, or it points to an external script file through the `src` attribute.\nCommon uses for JavaScript are image manipulation, form validation, and dynamic changes of content.",
}
