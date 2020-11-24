package token

// Type - The token type
type Type string

// Token - holds information of a token
type Token struct {
	Type    Type
	Literal string
}

const (
	// Illegal - Unknown token/character
	Illegal = Type("ILLEGAL")

	// EOF - End of File
	EOF = Type("EOF")

	// Ident - Indentifier, name of variables
	Ident = Type("IDENT")

	// Int - Integer
	Int = Type("INT")
)

// Operators
const (
	Assign = Type("=")
	Plus   = Type("+")
)

// Delimiters
const (
	Comma     = Type(",")
	Semicolon = Type(";")

	LeftParen  = Type("(")
	RightParen = Type(")")
	LeftBrace  = Type("[")
	RightBrace = Type("]")
)

// Keywords
const (
	Function = "FUNCTION"
	Let      = "LET"
)
