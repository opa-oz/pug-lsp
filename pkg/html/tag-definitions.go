package html

var TagToDesc = map[HtmlTag]string{
	Body:    "The `<body>` element contains all the contents of an HTML document, such as headings, paragraphs, images, hyperlinks, tables, lists, etc.\n\n*Note:* There can only be one `<body>` element in an HTML document.",
	Div:     "The `<div>` tag defines a division or a section in an HTML document.\nThe `<div>` tag is used as a container for HTML elements - which is then styled with CSS or manipulated with JavaScript.\nThe `<div>` tag is easily styled by using the class or id attribute.\nAny sort of content can be put inside the `<div>` tag! \n\n*Note:* By default, browsers always place a line break before and after the `<div>` element.",
	Script:  "The `<script>` tag is used to embed a client-side script (JavaScript).\nThe `<script>` element either contains scripting statements, or it points to an external script file through the `src` attribute.\nCommon uses for JavaScript are image manipulation, form validation, and dynamic changes of content.",
	A:       "The `<a>` tag defines a hyperlink, which is used to link from one page to another.\nThe most important attribute of the `<a>` element is the `href` attribute, which indicates the link's destination.\nBy default, links will appear as follows in all browsers:\n\n- An unvisited link is underlined and blue\n- A visited link is underlined and purple\n- An active link is underlined and red",
	Article: "The `<article>` tag specifies independent, self-contained content.\nAn article should make sense on its own and it should be possible to distribute it independently from the rest of the site.\nPotential sources for the `<article>` element:\n- Forum post\n- Blog post\n- News story\n*Note:* The `<article>` element does not render as anything special in a browser. However, you can use CSS to style the `<article>` element",
}