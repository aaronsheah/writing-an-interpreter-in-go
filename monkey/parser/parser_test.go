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
	checkParserErrors(t, p)

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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestLetStatmentError(t *testing.T) {
	// GIVEN
	input := `
	let 123;
	let x 5;
	let = 10;
	`
	// WHEN
	p := New(lexer.New(input))
	p.ParseProgram()

	// THEN
	if len(p.Errors()) != 3 {
		t.Fatalf("Expect 3 errors")
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return true;
		return "test";
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Incorrect number of statements parsed. Got %d statements", len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statment is not a ReturnStatement. got=%T", statement)
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("token literal should be 'return', got %q", returnStatement.TokenLiteral())
		}
	}
}
