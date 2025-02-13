package matcher

import "github.com/type-rb/type-rb-prototype/token"

var singleCharTokens = map[rune]token.TokenType{
	'(': token.LPAREN,
	')': token.RPAREN,
	0:   token.EOF,
}

type SingleCharMatcher struct {
}

func NewSingleCharMatcher() TokenMatcher {
	return &SingleCharMatcher{}
}

func (m *SingleCharMatcher) Match(ctx LexerContext) (bool, token.Token) {
	ch := ctx.CurrentChar()
	if tokType, ok := singleCharTokens[ch]; ok {
		pos := ctx.Pos()
		tok := token.Token{Type: tokType, Literal: string(ch), Pos: pos}
		ctx.ReadChar()
		return true, tok
	}
	return false, token.Token{}
}
