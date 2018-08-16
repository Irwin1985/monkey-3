package ast

import (
	"bytes"
	"strings"

	"github.com/gcoka/monkey/token"
)

// Node is a node of AST.
type Node interface {
	TokenLiteral() string
	String() string
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
	Node
	Statements []Statement
}

// TokenLiteral returns program's token literal.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier is an identifier.
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns identifier token literal.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string {
	return i.Value
}

// ReturnStatement is a return statement.
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns return statement token literal.
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement is an expression statement.
type ExpressionStatement struct {
	Token      token.Token // expression's first token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns expression statement token literal.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// IntegerLiteral is an integer literal.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns integer literal token literal.
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (il *IntegerLiteral) String() string { return il.Token.Literal }

// PrefixExpression is a prefix expression.
type PrefixExpression struct {
	Token    token.Token // prefix token e.g. !, -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns prefix literal token literal.
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression is an infix expression.
type InfixExpression struct {
	Token    token.Token // binary operator e.g. +, *
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

//TokenLiteral returns infix literal token literal.
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean is a boolean literal.
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral returns boolean literal token.
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

func (b *Boolean) String() string { return b.Token.Literal }

// IfExpression is an if expression.
type IfExpression struct {
	Token       token.Token // if token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral returns if statement literal token.
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// BlockStatement is a block statement.
type BlockStatement struct {
	Token      token.Token // {
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

//TokenLiteral returns block statement literal token.
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// FunctionLiteral is a function literal.
type FunctionLiteral struct {
	Token      token.Token // fn token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns function literal token.
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}
