package regex

//Nfa nfa结构
type Nfa struct {
	Init *State
	End  *State
}

//NewNfa 生成nfa
func NewNfa() *Nfa {
	return &Nfa{Init: NewState(), End: NewState()}
}

//NfaSet nfa集合
type NfaSet struct {
	set []*Nfa
}

//NewNfaSet 生成nfa集合
func NewNfaSet() *NfaSet {
	return &NfaSet{set: make([]*Nfa, 0)}
}

//Push 加入nfa
func (ns *NfaSet) Push(s *Nfa) {
	ns.set = append(ns.set, s)
}

//Pop 弹出nfa
func (ns *NfaSet) Pop() (*Nfa, bool) {
	cnt := len(ns.set)
	if cnt == 0 {
		return nil, false
	}
	cnt--
	nfa := ns.set[cnt]
	ns.set = ns.set[:cnt]
	return nfa, true
}
