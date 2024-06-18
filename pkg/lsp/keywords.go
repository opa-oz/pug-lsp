package lsp

type Keyword string

const (
	IfKeyword           Keyword = "if"
	UnlessKeyword       Keyword = "unless"
	IncludeKeyword      Keyword = "include"
	ForKeyword          Keyword = "for"
	EachKeyword         Keyword = "each"
	MixinKeyword        Keyword = "mixin"
	CaseKeyword         Keyword = "case"
	CaseWhenKeyword     Keyword = "when"
	CaseDefaultKeyword  Keyword = "default"
	BlockKeyword        Keyword = "block"
	ExtendsKeyword      Keyword = "extends"
	BlockAppendKeyword  Keyword = "append"
	BlockPrependKeyword Keyword = "prepend"
	ElseKeyword         Keyword = "else"
	WhileKeyword        Keyword = "while"
)

var baseKeywords = []Keyword{
	IfKeyword, ElseKeyword,
	UnlessKeyword, WhileKeyword, ForKeyword, EachKeyword,
	CaseKeyword,
	BlockKeyword, ExtendsKeyword,
	BlockAppendKeyword, BlockPrependKeyword,
	MixinKeyword, IncludeKeyword,
}

var blockKeywords = []Keyword{
	BlockPrependKeyword, BlockAppendKeyword,
}

var caseKeywords = []Keyword{
	CaseWhenKeyword, CaseDefaultKeyword,
}

func GetKeywords() *[]Keyword {
	return &baseKeywords
}

func GetBlockKeywords() *[]Keyword {
	return &blockKeywords
}

func GetCaseKeywords() *[]Keyword {
	return &caseKeywords
}

var KeywordToDesc = map[Keyword]string{
	IncludeKeyword: "Includes allow you to insert the contents of one Pug file into another.\n```pug\n//- index.pug\ndoctype html\nhtml\n  include includes/head.pug\n  body\n    h1 My Site\n    p Welcome to my super lame site.\n    include includes/foot.pug\n```\n\nIf the path is absolute (e.g., include /root.pug), it is resolved by prepending options.basedir. Otherwise, paths are resolved relative to the current file being compiled.\nIf no file extension is given, .pug is automatically appended to the file name.\n\n[Reference](https://pugjs.org/language/includes.html)",
	WhileKeyword:   "You can use while to create a loop:\n```pug\n- var n = 0;\nul\n  while n < 4\n    li= n++\n```\n\n[Reference](https://pugjs.org/language/iteration.html#while)",
	EachKeyword:    "Pug’s first-class iteration syntax makes it easier to iterate over arrays and objects in a template:\n```pug\nul\n  each val in [1, 2, 3, 4, 5]\n    li= val\n```\n\nYou can also get the index as you iterate:\n```pug\nul\n  each val, index in ['zero', 'one', 'two']\n    li= index + ': ' + val\n```\n\nPug also lets you iterate over the keys in an object:\n```pug\nul\n  each val, key in {1: 'one', 2: 'two', 3: 'three'}\n    li= key + ': ' + val\n```\n\nThe object or array to iterate over is just plain JavaScript. So, it can be a variable, or the result of a function call, or almost anything else.\n```pug\n- var values = [];\nul\n  each val in values.length ? values : ['There are no values']\n    li= val\n```\n\nOne can also add an else block that will be executed if the array or object does not contain values to iterate over.\n\n[Reference](https://pugjs.org/language/iteration.html)",
	IfKeyword:      "Pug’s first-class conditional syntax allows for optional parentheses.\n```pug\n- var user = {description: 'foo bar baz'}\n- var authorised = false\n#user\n  if user.description\n    h2.green Description\n    p.description= user.description\n  else if authorised\n    h2.blue Description\n    p.description.\n      User has no description,\n      why not add one...\n  else\n    h2.red Description\n    p.description User has no description\n```",
	UnlessKeyword:  "Pug provides the conditional unless, which works like a negated if.\n```pug\nunless user.isAnonymous\n  p You're logged in as #{user.name}\n```",
}
