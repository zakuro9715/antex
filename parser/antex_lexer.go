// Code generated from ./Antex.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 3, 10, 8, 
	1, 4, 2, 9, 2, 3, 2, 6, 2, 7, 10, 2, 13, 2, 14, 2, 8, 2, 2, 3, 3, 3, 3, 
	2, 3, 3, 2, 50, 59, 2, 10, 2, 3, 3, 2, 2, 2, 3, 6, 3, 2, 2, 2, 5, 7, 9, 
	2, 2, 2, 6, 5, 3, 2, 2, 2, 7, 8, 3, 2, 2, 2, 8, 6, 3, 2, 2, 2, 8, 9, 3, 
	2, 2, 2, 9, 4, 3, 2, 2, 2, 4, 2, 8, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames []string


var lexerSymbolicNames = []string{
	"", "INT",
}

var lexerRuleNames = []string{
	"INT",
}

type AntexLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewAntexLexer(input antlr.CharStream) *AntexLexer {

	l := new(AntexLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Antex.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AntexLexerINT is the AntexLexer token.
const AntexLexerINT = 1


