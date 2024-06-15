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
