package main

import (
	"fmt"

	"./regex"
)

var pmap map[int]int

func main() {
	lex := regex.NewLex("a*b*")
	parser := regex.NewParser(lex)
	nfa := parser.Parse()
	PrintState(nfa.Init)
}
func PrintState(state *regex.State) {
	pmap = make(map[int]int)
	printState(state)
}

//Print 打印信息
func printState(state *regex.State) {
	if _, have := pmap[state.ID]; have {
		return
	}
	pmap[state.ID] = 1
	fmt.Printf("State :%d Accept:%t \n", state.ID, state.IsAccept)
	for _, link := range state.TransLinks {
		token := "nil"
		if link.Token != nil {
			token = string(link.Token.Value)
		}
		fmt.Printf("trans Token:%s State: %d\n", token, link.EndState.ID)
	}
	for _, link := range state.TransLinks {
		printState(link.EndState)
	}
}
