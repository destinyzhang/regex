package regex

//TokenType token类型
type TokenType int

const (
	//EOF 结束
	EOF TokenType = iota
	//PLUS +
	PLUS
	//STAR *
	STAR
	//QMASK ?
	QMASK
	//OR |
	OR
	//CHAR 单个字
	CHAR
	//LP (
	LP
	//RP )
	RP
)

var (
	//EOFTOKEN 结束tokey
	EOFTOKEN = Token{Type: EOF}
)

//Token 标识
type Token struct {
	Type  TokenType
	Value rune
}

//NewToken 生成新token
func NewToken(r rune) *Token {
	var tt TokenType
	switch r {
	case '+':
		tt = PLUS
	case '*':
		tt = STAR
	case '?':
		tt = QMASK
	case '|':
		tt = OR
	case '(':
		tt = LP
	case ')':
		tt = RP
	default:
		tt = CHAR
	}
	return &Token{Value: r, Type: tt}
}
