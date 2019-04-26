package regex

//Nfa nfa结构
type Nfa struct {
	Init *State
	End  *State
}

//NewNfa 生成nfa
func NewNfa() *Nfa {
	return &Nfa{Init: NewState(false), End: NewState(true)}
}

//NewNfaWithTrans 生成nfa状态
func NewNfaWithTrans(token *Token) *Nfa {
	nfa := NewNfa()
	nfa.Init.AddTransLink(token, nfa.End)
	return nfa
}

func trans2Dfa(cru *Dfa, root *Dfa) {
	for _, r := range cru.SS.FindNoEpsilonTrans() {
		if cru.ExistsTransLink(r) {
			continue
		}
		ftss := cru.SS.FindTransEpsilonSet(r)
		if ftss == nil {
			continue
		}
		fdfa := root.FindPathEqualDfa(ftss, make(map[int]int))
		if fdfa == nil {
			fdfa = NewDfa(ftss)
		}
		cru.AddTransLink(r, fdfa)
		trans2Dfa(fdfa, root)
	}
}

//Trans2Dfa 转换到dfa
func (nfa *Nfa) Trans2Dfa() *Dfa {
	dfa := NewDfa(nfa.Init.FindEpsilonSet())
	trans2Dfa(dfa, dfa)
	return dfa
}

//ab 格式
func transAddNfa(nfa1 *Nfa, nfa2 *Nfa) *Nfa {
	nfa1.End.IsAccept = false
	nfa1.End.AddTransEpsilonLink(nfa2.Init)
	nfa1.End = nfa2.End
	return nfa1
}

//ac 格式
func transOrNfa(nfa1 *Nfa, nfa2 *Nfa) *Nfa {
	orNfa := NewNfa()
	orNfa.Init.AddTransEpsilonLink(nfa1.Init)
	orNfa.Init.AddTransEpsilonLink(nfa2.Init)

	nfa1.End.IsAccept = false
	nfa2.End.IsAccept = false

	nfa1.End.AddTransEpsilonLink(orNfa.End)
	nfa2.End.AddTransEpsilonLink(orNfa.End)
	return orNfa
}

//a* 格式
func transStarNfa(nfa *Nfa) *Nfa {
	eNfa := NewNfa()
	eNfa.Init.AddTransEpsilonLink(eNfa.End)
	eNfa.Init.AddTransEpsilonLink(nfa.Init)

	nfa.End.IsAccept = false
	nfa.End.AddTransEpsilonLink(eNfa.End)
	nfa.End.AddTransEpsilonLink(nfa.Init)
	return eNfa
}

//a+ 格式
func transPlusNfa(nfa *Nfa) *Nfa {
	nfa.End.AddTransEpsilonLink(nfa.Init)
	return nfa
}
