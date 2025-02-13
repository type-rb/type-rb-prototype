package matcher

import "github.com/type-rb/type-rb-prototype/token"

type StringMatcher struct {
}

func NewStringMatcher() TokenMatcher {
	return &StringMatcher{}
}

func (m *StringMatcher) Match(ctx LexerContext) (bool, token.Token) {
	if ctx.CurrentChar() != '"' {
		return false, token.Token{}
	}

	pos := ctx.Pos()
	tok := token.Token{Type: token.STRING, Literal: ctx.ReadString(), Pos: pos}
	ctx.ReadChar()
	return true, tok
}
