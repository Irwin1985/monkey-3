package ast

import "github.com/gcoka/monkey/token"

// Node is a node of AST.
type Node interface {
	TokenLiteral() string
}

// Statement is a statement.
type Statement interface {
	Node
	statementNode()
}

// Expression is an expression.
type Expression interface {
	Node
	expressionNode()
}

// Program is a root node of AST.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns program's token literal.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is a let statement.
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns let statement token literal.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is an identifier.
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns identifier token literal.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
