package regex

//Lex 词法分析
type Lex struct {
	cursor int    //当前扫描位置
	runes  []rune //表达式转换成的rune
}

//NewLex 生成词法
func NewLex(s string) *Lex {
	return &Lex{cursor: 0, runes: []rune(s)}
}

//Expression 返回正则式
func (lex *Lex) Expression() string {
	return string(lex.runes)
}

//NextToken 返回下一个token
func (lex *Lex) NextToken() *Token {
	if lex.cursor >= len(lex.runes) {
		return &EOFTOKEN
	}
	r := lex.runes[lex.cursor]
	lex.cursor++
	return NewToken(r)
}
