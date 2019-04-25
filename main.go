package main

import (
	"./regex"
)

func main() {
	lex := regex.NewLex("(a|b|c)*h|l")
	//lex := regex.NewLex("(ab)*")
	parser := regex.NewParser(lex)
	nfa := parser.Parse()
	if nfa != nil {
		//regex.PrintNfa(nfa)
		regex.PrintDfa(nfa.Trans2Dfa())
	}
}
