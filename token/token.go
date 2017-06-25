package token

const (
	LPAREN   = "("
	RPAREN   = ")"
	COMMA    = ","
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	CARET    = "^"
	TILDA    = "~"
	BANG     = "!"
	QUESTION = "?"
	COLON    = ":"

	EOF  = "EOF"
	NAME = "NAME"
)

type TokenType string

type Token struct {
	Type TokenType
	Text string
}
