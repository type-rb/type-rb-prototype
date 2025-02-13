package matcher

import "github.com/type-rb/type-rb-prototype/token"

type IdentifierMatcher struct {
}

func NewIdentifierMatcher() TokenMatcher {
	return &IdentifierMatcher{}
}

func (m *IdentifierMatcher) Match(ctx LexerContext) (bool, token.Token) {
	if !ctx.IsLetter(ctx.CurrentChar()) {
		return false, token.Token{}
	}
	pos := ctx.Pos()
	tok := token.Token{Type: token.IDENT, Literal: ctx.ReadIdentifier(), Pos: pos}
	return true, tok
}
