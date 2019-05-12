package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/zakuro9715/antex/parser"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: antex expr")
		os.Exit(2)
	}
	input := antlr.NewInputStream(os.Args[1])
	lexer := parser.NewAntexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAntexParser(stream)
	tree := p.Expr()
	fmt.Println(tree.ToStringTree([]string{}, p))
}
