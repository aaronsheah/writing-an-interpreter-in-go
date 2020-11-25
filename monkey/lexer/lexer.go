package lexer

import (
	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/token"
)

// Lexer - holds the state in the lexing process
// Main point of confusion could be between `position` and `readPosition`. This is mainly used for 'reading-ahead' to see if the token has ended etc.
type Lexer struct {
	input        string
	readPosition int  // current reading position in input (after curreny char)
	position     int  // current position in input (points to current char)
	char         byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func isValidCharacterForIdentifierOrKeywords(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) NextToken() token.Token {
	var outputToken token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		outputToken.Type = token.Assign
		outputToken.Literal = string(l.char)
	case '+':
		outputToken.Type = token.Plus
		outputToken.Literal = string(l.char)

	case ',':
		outputToken.Type = token.Comma
		outputToken.Literal = string(l.char)

	case ';':
		outputToken.Type = token.Semicolon
		outputToken.Literal = string(l.char)

	case '(':
		outputToken.Type = token.LeftParen
		outputToken.Literal = string(l.char)

	case ')':
		outputToken.Type = token.RightParen
		outputToken.Literal = string(l.char)

	case '{':
		outputToken.Type = token.LeftBrace
		outputToken.Literal = string(l.char)

	case '}':
		outputToken.Type = token.RightBrace
		outputToken.Literal = string(l.char)
	case 0: // NUL char
		outputToken.Type = token.EOF
		outputToken.Literal = ""
	default:
		if isValidCharacterForIdentifierOrKeywords(l.char) {
			outputToken.Literal = l.readIdentifier()

			if tokenType, ok := token.KeywordToTokenType[outputToken.Literal]; ok {
				outputToken.Type = tokenType
			} else {
				outputToken.Type = token.Ident
			}

			// return to avoid reading next char
			return outputToken
		} else if isDigit(l.char) {
			outputToken.Type = token.Int
			outputToken.Literal = l.readNumber()

			return outputToken
		} else {
			outputToken.Type = token.Illegal
			outputToken.Literal = string(l.char)
		}
	}

	l.readChar()

	return outputToken
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0 // NUL char
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	originalPosition := l.position
	// iterate through whole identifier
	for isValidCharacterForIdentifierOrKeywords(l.char) {
		l.readChar()
	}
	// grab the identifier
	return l.input[originalPosition:l.position]
}

// Advances the position past all whitespace
func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	originalPosition := l.position
	// iterate through whole identifier
	for isDigit(l.char) {
		l.readChar()

	}
	// grab the identifier
	return l.input[originalPosition:l.position]
}
