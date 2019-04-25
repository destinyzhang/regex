package regex

import (
	"fmt"
)

var (
	dfaID = 0
)

//Dfa dfa结构
type Dfa struct {
	ID            int
	SS            *StateSet
	DfaTransLinks []*DfaTransLink
}

//DfaTransLink 状态转换
type DfaTransLink struct {
	Token  rune
	EndDfa *Dfa
}

//ToString 打印
func (dtl *DfaTransLink) ToString() string {
	return fmt.Sprintf("trans Token:%s State: %d\n", string(dtl.Token), dtl.EndDfa.ID)
}

//生成状态id
func genDfaID() int {
	id := dfaID
	dfaID++
	return id
}

//NewDfa 生成dfa
func NewDfa(ss *StateSet) *Dfa {
	return &Dfa{SS: ss, ID: genDfaID(), DfaTransLinks: make([]*DfaTransLink, 0)}
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
	return dfa.FindTransLinkDfa(token) != nil
}

//FindTransLinkDfa 查询已经存在的连接
func (dfa *Dfa) FindTransLinkDfa(token rune) *Dfa {
	for _, dtl := range dfa.DfaTransLinks {
		if dtl.Token == token {
			return dtl.EndDfa
		}
	}
	return nil
}

//ToString 打印
func (dfa *Dfa) ToString() string {
	return fmt.Sprintf("Dfa:%d Set:%s Accept:%t \n", dfa.ID, dfa.SS.ToString(), dfa.IsAccept())
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
func (dfa *Dfa) FindPathEqualDfa(ss *StateSet, fmap map[*Dfa]int) *Dfa {
	if _, have := fmap[dfa]; have {
		return nil
	}
	fmap[dfa] = 1
	if dfa.SS.Equal(ss) {
		return dfa
	}
	for _, dtl := range dfa.DfaTransLinks {
		find := dtl.EndDfa.FindPathEqualDfa(ss, fmap)
		if find != nil {
			return find
		}
	}
	return nil
}
