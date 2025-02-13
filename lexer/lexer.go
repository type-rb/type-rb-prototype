package lexer

import (
	"github.com/type-rb/type-rb-prototype/lexer/matcher"
	"github.com/type-rb/type-rb-prototype/token"
	"io"
	"strings"
)

type Lexer struct {
	src      io.RuneReader
	ch       rune
	pos      token.Pos
	readPos  token.Pos
	matchers []matcher.TokenMatcher
}

func NewLexer(src io.RuneReader) *Lexer {
	l := &Lexer{src: src}
	l.matchers = []matcher.TokenMatcher{
		matcher.NewSingleCharMatcher(),
		matcher.NewStringMatcher(),
		matcher.NewIdentifierMatcher(),
	}
	l.ReadChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	for _, m := range l.matchers {
		if ok, tok := m.Match(l); ok {
			return tok
		}
	}

	tok := token.Token{Type: token.ILLEGAL, Literal: string(l.ch), Pos: l.pos}
	l.ReadChar()
	return tok
}

func (l *Lexer) CurrentChar() rune {
	return l.ch
}

func (l *Lexer) ReadChar() {
	var size int
	l.ch, size, _ = l.src.ReadRune()
	l.pos = l.readPos
	l.readPos = l.pos + token.Pos(size)
}

func (l *Lexer) Pos() token.Pos {
	return l.pos
}

func (l *Lexer) ReadIdentifier() string {
	var ident strings.Builder
	for l.IsLetter(l.ch) {
		ident.WriteRune(l.ch)
		l.ReadChar()
	}
	return ident.String()
}

func (l *Lexer) ReadString() string {
	var str strings.Builder
	l.ReadChar()
	for l.ch != '"' && l.ch != 0 {
		str.WriteRune(l.ch)
		l.ReadChar()
	}
	return str.String()
}

func (l *Lexer) IsLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
