package regex

//Parser 解析类
type Parser struct {
	lex    *Lex
	token  *Token
	nfaSet *NfaSet
}

//NewParser 生成Parser
func NewParser(lex *Lex) *Parser {
	return &Parser{lex: lex, token: lex.NextToken(), nfaSet: NewNfaSet()}
}

//Parse 解析
func (parser *Parser) Parse() *Nfa {
	/*
		expression -> expression(PLUS|START) | (expression) | (CHAR | QMASK)
	*/
	//表达式应以CHAR QMASK LP开头
	for {
		nfa := parser.baseToken()
		if nfa == nil {
			break
		}
		//正常完成判断后面是否有运算符号
		if parser.token.Type == PLUS {
			parser.consumeToken()
			nfa = transPlusNfa(nfa)
		} else if parser.token.Type == STAR {
			parser.consumeToken()
			nfa = transStarNfa(nfa)
		}
		parser.nfaSet.Push(nfa)
	}
	result, _ := parser.nfaSet.Pop()
	if result != nil {
		for {
			if nfa, suc := parser.nfaSet.Pop(); suc {
				result = transAddNfa(result, nfa)
			} else {
				break
			}
		}
	}
	return result
}

func transAddNfa(nfa1 *Nfa, nfa2 *Nfa) *Nfa {
	nfa1.End.IsAccept = false
	nfa1.End.AddTransEpsilonLink(nfa2.Init)
	return nfa1
}

func transOrNfa(nfa1 *Nfa, nfa2 *Nfa) *Nfa {
	orNfa := NewNfa()
	orNfa.Init.IsAccept = false
	orNfa.Init.AddTransEpsilonLink(nfa1.Init)
	orNfa.Init.AddTransEpsilonLink(nfa2.Init)

	nfa1.End.IsAccept = false
	nfa2.End.IsAccept = false

	nfa1.End.AddTransEpsilonLink(orNfa.End)
	nfa2.End.AddTransEpsilonLink(orNfa.End)
	return orNfa
}

func transStarNfa(nfa *Nfa) *Nfa {
	//a* == 空|a+
	nfa = transPlusNfa(nfa)
	//空
	eNfa := NewNfa()
	eNfa.Init.IsAccept = false
	eNfa.Init.AddTransEpsilonLink(eNfa.End)
	return transOrNfa(eNfa, nfa)
}

func transPlusNfa(nfa *Nfa) *Nfa {
	nfa.End.AddTransEpsilonLink(nfa.Init)
	return nfa
}

func (parser *Parser) consumeToken() {
	parser.token = parser.lex.NextToken()
}

//解析最基本的token
func (parser *Parser) baseToken() *Nfa {
	if parser.token.Type == CHAR || parser.token.Type == QMASK {
		nfa := NewNfa()
		nfa.Init.IsAccept = false
		nfa.End.IsAccept = true
		nfa.Init.AddTransLink(parser.token, nfa.End)
		parser.consumeToken()
		return nfa
	}
	return nil
}
