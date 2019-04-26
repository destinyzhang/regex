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
	fmt.Printf("%s", state.ToString())
	for _, link := range state.TransLinks {
		fmt.Printf("%s", link.ToString())
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
	fmt.Printf("%s", dfa.ToString())
	for _, link := range dfa.DfaTransLinks {
		fmt.Printf("%s", link.ToString())
		defer printDfa(link.EndDfa, filtMap)
	}
}
