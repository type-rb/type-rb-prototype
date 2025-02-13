package lexer

import (
	"github.com/type-rb/type-rb-prototype/token"
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `puts("hello, world!")`

	tests := []struct {
		wantType    token.TokenType
		wantLiteral string
		wantPos     int
	}{
		{token.IDENT, "puts", 0},
		{token.LPAREN, "(", 4},
		{token.STRING, "hello, world!", 5},
		{token.RPAREN, ")", 20},
		{token.EOF, "\x00", 21},
	}

	src := strings.NewReader(input)
	l := NewLexer(src)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.wantType {
			t.Fatalf("tests[%d] - tokentype wrong. want=%q, got=%q", i, tt.wantType, tok.Type)
		}

		if tok.Literal != tt.wantLiteral {
			t.Fatalf("tests[%d] - literal wrong. want=%q, got=%q", i, tt.wantLiteral, tok.Literal)
		}

		if int(tok.Pos) != tt.wantPos {
			t.Fatalf("tests[%d] - position wrong. want=%d, got=%d", i, tt.wantPos, tok.Pos)
		}
	}
}
