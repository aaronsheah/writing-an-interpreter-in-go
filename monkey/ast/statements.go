package ast

import (
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
