package regex

//Regex 正则类
type Regex struct {
	expression string
	dfa        *Dfa
	Debug      bool
}

//NewRegex 生成正则对象
func NewRegex(expression string, debug bool) *Regex {
	regex := &Regex{expression: expression, Debug: debug}
	regex.parse()
	return regex
}

//Parse 分析
func (regex *Regex) Parse(expression string) {
	if regex.expression == expression {
		return
	}
	regex.expression = expression
	regex.parse()
}

func (regex *Regex) parse() {
	nfa := NewParser(NewLex(regex.expression)).Parse()
	if nfa != nil {
		regex.dfa = nfa.Trans2Dfa()
		if regex.Debug {
			PrintNfa(nfa)
			PrintDfa(regex.dfa)
		}
	}
}

//Match 匹配检查
func (regex *Regex) Match(str string) bool {
	if regex.dfa == nil {
		return false
	}
	dfa := regex.dfa
	for _, r := range []rune(str) {
		dfa = dfa.MatchTransLinkDfa(r)
		if dfa == nil {
			return false
		}
	}
	return dfa.IsAccept()
}
