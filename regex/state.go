package regex

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

//EpsilonLink 是否空连接转换
func (tlink *TransLink) EpsilonLink() bool {
	return tlink.Token == nil
}

//State 状态
type State struct {
	ID         int          //状态id
	IsAccept   bool         //是否可接受状态
	TransLinks []*TransLink //状态转换
	epsilonSet *StateSet    //闭包集合
}

//NewState 生成状态
func NewState() *State {
	return &State{
		ID:         genStateID(),
		IsAccept:   true,
		TransLinks: make([]*TransLink, 0),
		epsilonSet: nil,
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
func (state *State) FindEpsilonSet(new bool) *StateSet {
	//没有生成过
	if state.epsilonSet == nil || new {
		state.epsilonSet = NewStateSet()
		findEpsilonSet(state, state.epsilonSet)
	}
	return state.epsilonSet
}

//StateSet 状态集合
type StateSet struct {
	set []*State
}

//NewStateSet 生成状态集合
func NewStateSet() *StateSet {
	return &StateSet{set: make([]*State, 0)}
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
