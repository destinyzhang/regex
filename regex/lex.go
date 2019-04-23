package regex

//Lex 词法分析
type Lex struct {
	cursor int    //当前扫描位置
	runes  []rune //表达式转换成的rune
}

//Expression 返回正则表达式
func (lex *Lex) Expression() string {
	return string(lex.runes)
}

//NextToken 返回下一个token
func (lex *Lex) NextToken() *Token {
	return nil
}
