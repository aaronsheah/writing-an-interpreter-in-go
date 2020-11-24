package lexer

import (
	"testing"

	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"
	l := New(input)

	expectedTokens := []token.Token{
		{Type: token.Assign, Literal: "="},
		{Type: token.Plus, Literal: "+"},
		{Type: token.LeftParen, Literal: "("},
		{Type: token.RightParen, Literal: ")"},
		{Type: token.LeftBrace, Literal: "{"},
		{Type: token.RightBrace, Literal: "}"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Semicolon, Literal: ";"},
	}

	for _, expectedToken := range expectedTokens {
		output := l.NextToken()

		if output.Type != expectedToken.Type {
			t.Fatalf("Wrong token type. expected:%q, got:%q", expectedToken.Type, output.Type)
		}
		if output.Literal != expectedToken.Literal {
			t.Fatalf("Wrong token literal. expected:%q, got:%q", expectedToken.Literal, output.Literal)
		}
	}
}
