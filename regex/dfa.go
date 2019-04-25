package regex

//Dfa dfa结构
type Dfa struct {
	SS            *StateSet
	DfaTransLinks []*DfaTransLink
}

//DfaTransLink 状态转换
type DfaTransLink struct {
	Token  rune
	EndDfa *Dfa
}

//NewDfa 生成dfa
func NewDfa(ss *StateSet) *Dfa {
	return &Dfa{SS: ss, DfaTransLinks: make([]*DfaTransLink, 0)}
}

//IsAccept 是否可接受状态
func (dfa *Dfa) IsAccept() bool {
	return dfa.SS.IsAccept()
}

//Equal 是否相等
func (dfa *Dfa) Equal(other *Dfa) bool {
	return dfa.SS.Equal(other.SS)
}

//ExistsTransLink 是否已经存在了转换
func (dfa *Dfa) ExistsTransLink(token rune) bool {
	for _, dtl := range dfa.DfaTransLinks {
		if dtl.Token == token {
			return true
		}
	}
	return false
}

//AddTransLink 加入连接
func (dfa *Dfa) AddTransLink(token rune, other *Dfa) bool {
	if other == nil {
		return false
	}
	if dfa.ExistsTransLink(token) {
		return false
	}
	dfa.DfaTransLinks = append(dfa.DfaTransLinks, &DfaTransLink{Token: token, EndDfa: other})
	return true
}

//FindPathEqualDfa 查询树里面是否有已存在的dfa
func (dfa *Dfa) FindPathEqualDfa(ss *StateSet) *Dfa {
	if dfa.SS.Equal(ss) {
		return dfa
	}
	for _, dtl := range dfa.DfaTransLinks {
		find := dtl.EndDfa.FindPathEqualDfa(ss)
		if find != nil {
			return find
		}
	}
	return nil
}
