package token

type Pos int

const NoPos Pos = 0

type Position struct {
	offset int // offset, starting at 0
	Line   Pos // line number, starting at 1
	Column Pos // column number, starting at 1 (byte count)
}
