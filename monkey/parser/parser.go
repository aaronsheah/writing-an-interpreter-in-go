package parser

import (
	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/ast"
	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/lexer"
	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/token"
)

type Parser struct {
	lexer   *lexer.Lexer
	current token.Token
	peek    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{lexer: lexer}

	// Populate `current` and `peek`
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.current.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.current.Type {
	case token.Let:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	statement := &ast.LetStatement{}

	// Should have a name (aka Identifier)
	if p.peek.Type != token.Ident {
		return nil
	}

	p.nextToken()
	statement.Name = &ast.Identifier{
		Token: p.current,
		Value: p.current.Literal,
	}

	// Should have a assignment op after name
	if p.peek.Type != token.Assign {
		return nil
	}

	// TODO: skip expression for now until semicolon
	p.nextToken()
	for p.current.Type != token.Semicolon {
		p.nextToken()
	}
	// statement.Value

	return statement
}

// peekTokenTypeIs - checks i the peeked token has the same type as the given type. If it is the same move the 'cursor' and return true
func (p *Parser) peekTokenTypeIs(t token.Type) bool {
	if p.peek.Type == t {
		p.nextToken()
		return true
	}
	return false
}
