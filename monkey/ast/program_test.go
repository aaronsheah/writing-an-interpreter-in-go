package ast

import (
	"testing"

	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/token"
)

func TestString(t *testing.T) {

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Name: &Identifier{
					Token: token.Token{
						Type:    token.Ident,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.Ident,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}
	expectedString := "let myVar = anotherVar;"
	if program.String() != expectedString {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
