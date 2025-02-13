package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	STRING  = "STRING"
	LPAREN  = "("
	RPAREN  = ")"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Pos     Pos
}
