package lexer

import (
	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/token"
)

type Lexer struct {
}

func New(input string) *Lexer {
	return &Lexer{}
}

func (l *Lexer) NextToken() token.Token {
	return token.Token{}
}
