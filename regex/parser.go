package regex

import "fmt"

//Parser 解析类
type Parser struct {
	lex      *Lex
	token    *Token
	stackEnd []TokenType
	end      TokenType
}

//NewParser 生成Parser
func NewParser(lex *Lex) *Parser {
	return &Parser{lex: lex, token: lex.NextToken(), stackEnd: make([]TokenType, 0)}
}

/*
Parse 解析 表达式应以CHAR QMASK LP开头
	expression -> expression(PLUS|START) | (expression) | (CHAR | QMASK)
*/
func (parser *Parser) Parse() (result *Nfa) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	result = parser.expressionWithEnd(EOF)
	return
}

func (parser *Parser) push(tt TokenType) {
	parser.end = tt
	parser.stackEnd = append(parser.stackEnd, tt)
}

func (parser *Parser) pop() {
	cnt := len(parser.stackEnd) - 1
	parser.stackEnd = parser.stackEnd[:cnt]
	if cnt > 0 {
		parser.end = parser.stackEnd[cnt-1]
	}

}

func (parser *Parser) panic(err string) {
	panic(err)
}

func (parser *Parser) consumeToken(tt TokenType) {
	if parser.token.Type != tt {
		parser.panic(fmt.Sprintf("invalid token: %s in: [%s] index: [%d] need: %s",
			parser.token.ToString(), parser.lex.Expression(), parser.lex.Cursor(), TokenDesc[tt]))
	}
	parser.token = parser.lex.NextToken()
}

func (parser *Parser) expressionWithEnd(tt TokenType) *Nfa {
	parser.push(tt)
	defer parser.pop()
	return parser.expression()
}

func (parser *Parser) expression() *Nfa {
	var result *Nfa
	for {
		result = parser.addExpression(result)
		if parser.token.Type == parser.end {
			break
		}
	}
	return result
}

func (parser *Parser) addExpression(head *Nfa) *Nfa {
	nfa := parser.repetExpression()
	if head != nil {
		nfa = transAddNfa(head, nfa)
	}
	return parser.orExpression(nfa)
}

func (parser *Parser) orExpression(nfa *Nfa) *Nfa {
	if parser.token.Type == OR {
		parser.consumeToken(OR)
		nfa = transOrNfa(nfa, parser.expression())
	}
	return nfa
}

func (parser *Parser) repetExpression() *Nfa {
	nfa := parser.baseExpression()
	if parser.token.Type == PLUS {
		parser.consumeToken(PLUS)
		nfa = transPlusNfa(nfa)
	} else if parser.token.Type == STAR {
		parser.consumeToken(STAR)
		nfa = transStarNfa(nfa)
	}
	return nfa
}

func (parser *Parser) baseExpression() *Nfa {
	if parser.token.Type == CHAR {
		nfa := NewNfaWithTrans(parser.token)
		parser.consumeToken(CHAR)
		return nfa
	}
	if parser.token.Type == QMASK {
		nfa := NewNfaWithTrans(parser.token)
		parser.consumeToken(QMASK)
		return nfa
	}
	if parser.token.Type == LP {
		parser.consumeToken(LP)
		nfa := parser.expressionWithEnd(RP)
		parser.consumeToken(RP)
		return nfa
	}
	parser.panic(fmt.Sprintf("invalid token: %s in: [%s] index: [%d] need: %s or %s or %s",
		parser.token.ToString(), parser.lex.Expression(), parser.lex.Cursor(), TokenDesc[LP], TokenDesc[CHAR], TokenDesc[QMASK]))
	return nil
}
