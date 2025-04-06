package lexer

import (
	"github.com/type-rb/type-rb-prototype/lexer/matcher"
	"github.com/type-rb/type-rb-prototype/token"
	"io"
	"strings"
)

type Lexer interface {
	NextToken() token.Token
}

type lexer struct {
	src      io.RuneReader
	ch       rune
	pos      token.Pos
	readPos  token.Pos
	matchers []matcher.TokenMatcher
}

func NewLexer(src io.RuneReader) Lexer {
	l := &lexer{src: src}
	l.matchers = []matcher.TokenMatcher{
		matcher.NewSingleCharMatcher(),
		matcher.NewStringMatcher(),
		matcher.NewIdentifierMatcher(),
	}
	l.ReadChar()
	return l
}

func (l *lexer) NextToken() token.Token {
	for _, m := range l.matchers {
		if ok, tok := m.Match(l); ok {
			return tok
		}
	}

	tok := token.Token{Type: token.ILLEGAL, Literal: string(l.ch), Pos: l.pos}
	l.ReadChar()
	return tok
}

func (l *lexer) CurrentChar() rune {
	return l.ch
}

func (l *lexer) ReadChar() {
	var size int
	l.ch, size, _ = l.src.ReadRune()
	l.pos = l.readPos
	l.readPos = l.pos + token.Pos(size)
}

func (l *lexer) Pos() token.Pos {
	return l.pos
}

func (l *lexer) ReadIdentifier() string {
	var ident strings.Builder
	for l.IsLetter(l.ch) {
		ident.WriteRune(l.ch)
		l.ReadChar()
	}
	return ident.String()
}

func (l *lexer) ReadString() string {
	var str strings.Builder
	l.ReadChar()
	for l.ch != '"' && l.ch != 0 {
		str.WriteRune(l.ch)
		l.ReadChar()
	}
	return str.String()
}

func (l *lexer) IsLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
