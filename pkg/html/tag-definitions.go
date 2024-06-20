package html

const headerD = "The `<h1>` to `<h6>` tags are used to define HTML headings.\n`<h1>` defines the most important heading. `<h6>` defines the least important heading."
const bodyD = "The `<body>` element contains all the contents of an HTML document, such as headings, paragraphs, images, hyperlinks, tables, lists, etc.\n\n*Note:* There can only be one `<body>` element in an HTML document."
const divD = "The `<div>` tag defines a division or a section in an HTML document.\nThe `<div>` tag is used as a container for HTML elements - which is then styled with CSS or manipulated with JavaScript.\nThe `<div>` tag is easily styled by using the class or id attribute.\nAny sort of content can be put inside the `<div>` tag! \n\n*Note:* By default, browsers always place a line break before and after the `<div>` element."
const scriptD = "The `<script>` tag is used to embed a client-side script (JavaScript).\nThe `<script>` element either contains scripting statements, or it points to an external script file through the `src` attribute.\nCommon uses for JavaScript are image manipulation, form validation, and dynamic changes of content."
const aD = "The `<a>` tag defines a hyperlink, which is used to link from one page to another.\nThe most important attribute of the `<a>` element is the `href` attribute, which indicates the link's destination.\nBy default, links will appear as follows in all browsers:\n\n- An unvisited link is underlined and blue\n- A visited link is underlined and purple\n- An active link is underlined and red"
const articleD = "The `<article>` tag specifies independent, self-contained content.\nAn article should make sense on its own and it should be possible to distribute it independently from the rest of the site.\nPotential sources for the `<article>` element:\n- Forum post\n- Blog post\n- News story\n*Note:* The `<article>` element does not render as anything special in a browser. However, you can use CSS to style the `<article>` element"
const baseD = "The `<base>` tag specifies the base URL and/or target for all relative URLs in a document.\nThe `<base>` tag must have either an href or a target attribute present, or both.\nThere can only be one single `<base>` element in a document, and it must be inside the `<head>` element."
const buttonD = "The `<button>` tag defines a clickable button.\nInside a `<button>` element you can put text (and tags like `<i>`, `<b>`, `<strong>`, `<br>`, `<img>`, etc.). That is not possible with a button created with the `<input>` element!\n*Tip:* Always specify the `type` attribute for a `<button>` element, to tell browsers what type of button it is."
const footerD = "The `<footer>` tag defines a footer for a document or section.\nA `<footer>` element typically contains:\n- authorship information\n- copyright information\n- contact information\n- sitemap\n- back to top links\n- related documents\n\nYou can have several `<footer>` elements in one document.\n\n*Tip:* Contact information inside a <footer> element should go inside an `<address>` tag."
const formD = "The `<form>` tag is used to create an HTML form for user input.\nThe `<form>` element can contain one or more of the following form elements:\n- `<input>`\n- `<textarea>`\n- `<button>`\n- `<select>`\n- `<option>`\n- `<optgroup>`\n- `<fieldset>`\n- `<label>`\n- `<output>`"
const imgD = "The `<img>` tag is used to embed an image in an HTML page.\nImages are not technically inserted into a web page; images are linked to web pages. The `<img>` tag creates a holding space for the referenced image.\nThe `<img>` tag has two required attributes:\n- `src` - Specifies the path to the image\n- `alt` - Specifies an alternate text for the image, if the image for some reason cannot be displayed\nNote: Also, always specify the width and height of an image. If width and height are not specified, the page might flicker while the image loads.\n*Tip:* To link an image to another document, simply nest the `<img>` tag inside an `<a>` tag"

var TagToDesc = map[HtmlTag]string{
	H1:      headerD,
	H2:      headerD,
	H3:      headerD,
	H4:      headerD,
	H6:      headerD,
	H5:      headerD,
	Body:    bodyD,
	Div:     divD,
	Script:  scriptD,
	A:       aD,
	Article: articleD,
	Base:    baseD,
	Button:  buttonD,
	Footer:  footerD,
	Form:    formD,
	Img:     imgD,
}
