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

	checkNextChar(t, l, expectedTokens)
}

func TestNextToken_GivenCode(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;

		let add = fn(x, y) {
			x + y;
		};

		let results = add(five, ten);
	`
	l := New(input)

	expectedTokens := []token.Token{
		{Type: token.Let, Literal: "let"},
		{Type: token.Ident, Literal: "five"},
		{Type: token.Assign, Literal: "="},
		{Type: token.Int, Literal: "5"},
		{Type: token.Semicolon, Literal: ";"},

		{Type: token.Let, Literal: "let"},
		{Type: token.Ident, Literal: "ten"},
		{Type: token.Assign, Literal: "="},
		{Type: token.Int, Literal: "10"},
		{Type: token.Semicolon, Literal: ";"},

		{Type: token.Let, Literal: "let"},
		{Type: token.Ident, Literal: "add"},
		{Type: token.Assign, Literal: "="},
		{Type: token.Function, Literal: "fn"},
		{Type: token.LeftParen, Literal: "("},
		{Type: token.Ident, Literal: "x"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Ident, Literal: "y"},
		{Type: token.RightParen, Literal: ")"},
		{Type: token.LeftBrace, Literal: "{"},
		{Type: token.Ident, Literal: "x"},
		{Type: token.Plus, Literal: "+"},
		{Type: token.Ident, Literal: "y"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.RightBrace, Literal: "}"},
		{Type: token.Semicolon, Literal: ";"},

		{Type: token.Let, Literal: "let"},
		{Type: token.Ident, Literal: "results"},
		{Type: token.Assign, Literal: "="},
		{Type: token.Ident, Literal: "add"},
		{Type: token.LeftParen, Literal: "("},
		{Type: token.Ident, Literal: "five"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Ident, Literal: "ten"},
		{Type: token.RightParen, Literal: ")"},
		{Type: token.Semicolon, Literal: ";"},

		{Type: token.EOF, Literal: ""},
	}

	checkNextChar(t, l, expectedTokens)
}

func TestNextToken_SupportBoolean(t *testing.T) {
	testParams := []struct {
		Input          string
		ExpectedTokens []token.Token
	}{
		{
			`false;`,
			[]token.Token{
				{Type: token.False, Literal: "false"},
				{Type: token.Semicolon, Literal: ";"},
			},
		},
		{
			`true;`,
			[]token.Token{
				{Type: token.True, Literal: "true"},
				{Type: token.Semicolon, Literal: ";"},
			},
		},
	}

	for _, testParam := range testParams {
		l := New(testParam.Input)
		checkNextChar(t, l, testParam.ExpectedTokens)
	}
}

func TestNextToken_SupportLogicalOperators(t *testing.T) {
	testParams := []struct {
		Input          string
		ExpectedTokens []token.Token
	}{
		{
			`x < 5;`,
			[]token.Token{
				{Type: token.Ident, Literal: "x"},
				{Type: token.LessThan, Literal: "<"},
				{Type: token.Int, Literal: "5"},
			},
		},
		{
			`x > 5;`,
			[]token.Token{
				{Type: token.Ident, Literal: "x"},
				{Type: token.GreaterThan, Literal: ">"},
				{Type: token.Int, Literal: "5"},
			},
		},
		{
			`let x = !y;`,
			[]token.Token{
				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "x"},
				{Type: token.Assign, Literal: "="},
				{Type: token.Bang, Literal: "!"},
				{Type: token.Ident, Literal: "y"},
			},
		},
	}

	for _, testParam := range testParams {
		l := New(testParam.Input)
		checkNextChar(t, l, testParam.ExpectedTokens)
	}
}

func TestNextToken_SupportMathematicalOperators(t *testing.T) {
	testParams := []struct {
		Input          string
		ExpectedTokens []token.Token
	}{
		{
			`5 + 4 - 1 * 42 / 69;`,
			[]token.Token{
				{Type: token.Int, Literal: "5"},
				{Type: token.Plus, Literal: "+"},
				{Type: token.Int, Literal: "4"},
				{Type: token.Minus, Literal: "-"},
				{Type: token.Int, Literal: "1"},
				{Type: token.Asterisk, Literal: "*"},
				{Type: token.Int, Literal: "42"},
				{Type: token.Slash, Literal: "/"},
				{Type: token.Int, Literal: "69"},
			},
		},
	}

	for _, testParam := range testParams {
		l := New(testParam.Input)
		checkNextChar(t, l, testParam.ExpectedTokens)
	}
}

func TestNextToken_SupportIfElse(t *testing.T) {
	testParams := []struct {
		Input          string
		ExpectedTokens []token.Token
	}{
		{
			`
				let x = 42;
				let theAnswer;
				if (x < 42) {
					theAnswer = true;
				}
				else {
					theAnswer = false;
				}
			`,
			[]token.Token{

				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "x"},
				{Type: token.Assign, Literal: "="},
				{Type: token.Int, Literal: "42"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "theAnswer"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.If, Literal: "if"},
				{Type: token.LeftParen, Literal: "("},
				{Type: token.Ident, Literal: "x"},
				{Type: token.LessThan, Literal: "<"},
				{Type: token.Int, Literal: "42"},
				{Type: token.RightParen, Literal: ")"},
				{Type: token.LeftBrace, Literal: "{"},

				{Type: token.Ident, Literal: "theAnswer"},
				{Type: token.Assign, Literal: "="},
				{Type: token.True, Literal: "true"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.RightBrace, Literal: "}"},

				{Type: token.Else, Literal: "else"},
				{Type: token.LeftBrace, Literal: "{"},

				{Type: token.Ident, Literal: "theAnswer"},
				{Type: token.Assign, Literal: "="},
				{Type: token.False, Literal: "false"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.RightBrace, Literal: "}"},
			},
		},
	}

	for _, testParam := range testParams {
		l := New(testParam.Input)
		checkNextChar(t, l, testParam.ExpectedTokens)
	}
}

func checkNextChar(t *testing.T, l *Lexer, expectedTokens []token.Token) {
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
