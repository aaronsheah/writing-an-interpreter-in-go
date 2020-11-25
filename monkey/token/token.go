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
	Assign           = Type("=")
	Plus             = Type("+")
	Minus            = Type("-")
	Bang             = Type("!")
	Asterisk         = Type("*")
	Slash            = Type("/")
	LessThan         = Type("<")
	GreaterThan      = Type(">")
	LessThanEqual    = Type("<=")
	GreaterThanEqual = Type(">=")
)

// Delimiters
const (
	Comma     = Type(",")
	Semicolon = Type(";")

	LeftParen  = Type("(")
	RightParen = Type(")")
	LeftBrace  = Type("{")
	RightBrace = Type("}")
)

// Keywords
const (
	Function = Type("FUNCTION")
	Let      = Type("LET")
	True     = Type("TRUE")
	False    = Type("FALSE")
)

// KeywordToTokenType - map of keyword string to its corresponding token type
var KeywordToTokenType = map[string]Type{
	"fn":    Function,
	"let":   Let,
	"true":  True,
	"false": False,
}
