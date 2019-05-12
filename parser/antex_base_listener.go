// Code generated from ./Antex.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Antex

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAntexListener is a complete listener for a parse tree produced by AntexParser.
type BaseAntexListener struct{}

var _ AntexListener = &BaseAntexListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAntexListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAntexListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAntexListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAntexListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseAntexListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseAntexListener) ExitExpr(ctx *ExprContext) {}

// EnterValue is called when production value is entered.
func (s *BaseAntexListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseAntexListener) ExitValue(ctx *ValueContext) {}
