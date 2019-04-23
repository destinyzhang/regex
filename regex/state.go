package regex

//TransLink 状态转换
type TransLink struct {
	EpsilonLink bool   //是否空连转换
	Token       *Token //对应转换token
	EndState    *State //转换目标状态
}

//State 状态
type State struct {
	ID         int          //状态id
	IsAccept   bool         //是否可接受状态
	TransLinks []*TransLink //状态转换
}

//FindEpsilonSet 返回闭包集合
func (state *State) FindEpsilonSet() []*State {
	return nil
}
