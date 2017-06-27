package lexer

import (
	"testing"

	"github.com/mdaisuke/prattparser/token"
)

func TestNextToken(t *testing.T) {
	input := `(),=+-*/^~!?:`

	tests := []struct {
		expectedType string
		expectedText string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.CARET, "^"},
		{token.TILDA, "~"},
		{token.BANG, "!"},
		{token.RPAREN, "?"},
		{token.COLON, ":"},
	}
}
