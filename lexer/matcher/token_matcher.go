package matcher

import "github.com/type-rb/type-rb-prototype/token"

type TokenMatcher interface {
	Match(ctx LexerContext) (bool, token.Token)
}
