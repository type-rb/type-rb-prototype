package matcher

import "github.com/type-rb/type-rb-prototype/token"

type LexerContext interface {
	CurrentChar() rune
	ReadChar()
	Pos() token.Pos
	ReadIdentifier() string
	ReadString() string
	IsLetter(rune) bool
}
