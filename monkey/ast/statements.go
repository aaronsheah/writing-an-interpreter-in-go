package ast

import (
	"bytes"

	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/token"
)

type LetStatement struct {
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) Token() token.Token {
	return token.Token{
		Type:    token.Let,
		Literal: "let",
	}
}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token().Literal
}
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())

	if ls.Value != nil {
		out.WriteString(" = ")
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

type ReturnStatement struct {
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) Token() token.Token {
	return token.Token{
		Type:    token.Return,
		Literal: "return",
	}
}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token().Literal
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
