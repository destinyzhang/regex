package regex

import (
	"fmt"
	"strings"
)

var (
	stateID = 0
)

//生成状态id
func genStateID() int {
	id := stateID
	stateID++
	return id
}

//TransLink 状态转换
type TransLink struct {
	Token    *Token //对应转换token
	EndState *State //转换目标状态
}

//ToString 打印
func (tl *TransLink) ToString() string {
	str := "nil"
	if tl.Token != nil {
		str = tl.Token.ToString()
	}
	return fmt.Sprintf("trans Token:%s State: %d\n", str, tl.EndState.ID)
}

//EpsilonLink 是否空连接转换
func (tl *TransLink) EpsilonLink() bool {
	return tl.Token == nil
}

//State 状态
type State struct {
	ID         int          //状态id
	IsAccept   bool         //是否可接受状态
	TransLinks []*TransLink //状态转换
}

//NewState 生成状态
func NewState(accept bool) *State {
	return &State{
		ID:         genStateID(),
		IsAccept:   accept,
		TransLinks: make([]*TransLink, 0),
	}
}

//查找状态闭包
func findEpsilonSet(state *State, set *StateSet) {
	for _, tl := range state.TransLinks {
		if tl.EpsilonLink() {
			if set.Push(tl.EndState) {
				findEpsilonSet(tl.EndState, set)
			}
		}
	}
}

//AddTransEpsilonLink 加入空连接转换
func (state *State) AddTransEpsilonLink(endState *State) {
	state.AddTransLink(nil, endState)
}

//AddTransLink 加入转换
func (state *State) AddTransLink(token *Token, endState *State) {
	state.TransLinks = append(state.TransLinks, &TransLink{Token: token, EndState: endState})
}

//FindEpsilonSet 返回闭包集合
func (state *State) FindEpsilonSet() *StateSet {
	epsilonSet := NewStateSet()
	epsilonSet.Push(state)
	findEpsilonSet(state, epsilonSet)
	return epsilonSet
}

//ToString 打印
func (state *State) ToString() string {
	return fmt.Sprintf("State :%d Accept:%t \n", state.ID, state.IsAccept)
}

//StateSet 状态集合
type StateSet struct {
	set []*State
}

//NewStateSet 生成状态集合
func NewStateSet() *StateSet {
	return &StateSet{set: make([]*State, 0)}
}

//Count 状态数量
func (ss *StateSet) Count() int {
	return len(ss.set)
}

//Exists 是否存在状态
func (ss *StateSet) Exists(s *State) bool {
	for _, item := range ss.set {
		if item.ID == s.ID {
			return true
		}
	}
	return false
}

//Push 加入新状态
func (ss *StateSet) Push(s *State) bool {
	if ss.Exists(s) {
		return false
	}
	ss.set = append(ss.set, s)
	return true
}

//IsAccept 是否可接受状态合集
func (ss *StateSet) IsAccept() bool {
	for _, s := range ss.set {
		if s.IsAccept {
			return true
		}
	}
	return false
}

//Megre 合并状态
func (ss *StateSet) Megre(other *StateSet) *StateSet {
	megre := NewStateSet()
	megre.Megre(other)
	megre.Megre(ss)
	return megre
}

func (ss *StateSet) megre(other *StateSet) *StateSet {
	if other != nil {
		for _, s := range other.set {
			ss.Push(s)
		}
	}
	return ss
}

//Equal 是否相等
func (ss *StateSet) Equal(other *StateSet) bool {
	if len(ss.set) != len(other.set) {
		return false
	}
	for _, s := range ss.set {
		if !other.Exists(s) {
			return false
		}
	}
	return true
}

//ToString 打印
func (ss *StateSet) ToString() string {
	strSet := make([]string, 0, len(ss.set))
	for _, s := range ss.set {
		strSet = append(strSet, fmt.Sprintf("%d", s.ID))
	}
	return strings.Join(strSet, ",")
}

//FindNoEpsilonTrans 查询非空连转换
func (ss *StateSet) FindNoEpsilonTrans() []rune {
	runes := make([]rune, 0)
	exist := func(r rune) bool {
		for _, rr := range runes {
			if rr == r {
				return true
			}
		}
		return false
	}
	for _, s := range ss.set {
		for _, tl := range s.TransLinks {
			if !tl.EpsilonLink() {
				if !exist(tl.Token.Value) {
					runes = append(runes, tl.Token.Value)
				}
			}
		}
	}
	return runes
}

//FindTransEpsilonSet 查询对应转换集合
func (ss *StateSet) FindTransEpsilonSet(r rune) *StateSet {
	var ftss *StateSet
	for _, s := range ss.set {
		for _, tl := range s.TransLinks {
			if !tl.EpsilonLink() && tl.Token.Value == r {
				ftss = tl.EndState.FindEpsilonSet().megre(ftss)
			}
		}
	}
	return ftss
}
