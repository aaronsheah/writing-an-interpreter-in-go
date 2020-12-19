package parser

import (
	"testing"

	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/ast"
	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 5;
		let z = 5;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Incorrect number of statements parsed. Got %d statements", len(program.Statements))
	}

	expectedNames := []string{
		"x",
		"y",
		"z",
	}

	for i, expectedName := range expectedNames {
		statements := program.Statements[i]
		testLetStatement(t, statements, expectedName)
	}
}

func testLetStatement(t *testing.T, statement ast.Statement, expectedName string) {
	// Check that tokent literal is 'let'
	if statement.TokenLiteral() != "let" {
		t.Fatalf("token literal should have been 'let', instead got=%q", statement.TokenLiteral())
	}

	// Check that the given statement is LetStatement type
	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Fatalf("s not *ast.LetStatement. got=%T", statement)
	}

	// Check that the identifier of the let statement matches the expected name
	if letStatement.Name.Value != expectedName {
		t.Fatalf("expectedName not '%s', got=%s", expectedName, letStatement.Name.Value)
	}
}
