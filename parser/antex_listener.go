// Code generated from ./Antex.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Antex

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AntexListener is a complete listener for a parse tree produced by AntexParser.
type AntexListener interface {
	antlr.ParseTreeListener

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
