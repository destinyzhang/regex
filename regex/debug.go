package regex

import (
	"fmt"
)

//PrintNfa 打印信息
func PrintNfa(nfa *Nfa) {
	fmt.Println("-----------------------nfa info-----------------------")
	PrintState(nfa.Init)
}

//PrintState 打印信息
func PrintState(state *State) {
	printState(state, make(map[int]int))
}

func printState(state *State, filtMap map[int]int) {
	if _, have := filtMap[state.ID]; have {
		return
	}
	filtMap[state.ID] = 1
	fmt.Printf("State :%d Accept:%t \n", state.ID, state.IsAccept)
	for _, link := range state.TransLinks {
		token := "nil"
		if link.Token != nil {
			token = link.Token.ToString()
		}
		fmt.Printf("trans Token:%s State: %d\n", token, link.EndState.ID)
		defer printState(link.EndState, filtMap)
	}
}

//PrintDfa 打印信息
func PrintDfa(dfa *Dfa) {
	fmt.Println("-----------------------dfa info-----------------------")
	printDfa(dfa, make(map[*Dfa]int))
}

func printDfa(dfa *Dfa, filtMap map[*Dfa]int) {
	if _, have := filtMap[dfa]; have {
		return
	}
	filtMap[dfa] = 1
	fmt.Printf("Dfa:%d Set:%s Accept:%t \n", dfa.ID, dfa.SS.ToString(), dfa.IsAccept())
	for _, link := range dfa.DfaTransLinks {
		fmt.Printf("trans Token:%s State: %d\n", string(link.Token), link.EndDfa.ID)
		defer printDfa(link.EndDfa, filtMap)
	}
}
