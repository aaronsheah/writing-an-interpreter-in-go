package ast

// Expressions produce values, statements don't

type Node interface {
	TokenLiteral() string
}

// Created separate interfaces for Statement and Expression for type checking
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}
