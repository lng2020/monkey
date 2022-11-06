package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT and INT: Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// ASSIGN and PLUS: Operators
	ASSIGN = "="
	PLUS   = "+"

	// COMMA and SEMICOLON: Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// FUNCTION and LET: Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
