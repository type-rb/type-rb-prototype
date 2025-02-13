package token

type File struct {
	name  string
	base  Pos
	size  int
	lines []Pos
}

func NewFile(name string, base Pos, src string) *File {
	f := &File{name: name, base: base, size: len(src)}
	f.setLines(src)
	return f
}

func (f *File) Line(p Pos) Pos {
	return f.Position(p).Line
}

func (f *File) Column(p Pos) Pos {
	return f.Position(p).Column
}

func (f *File) Position(p Pos) Position {
	line := 1
	offset := p - f.base
	i, j := 0, len(f.lines)
	for i < j {
		h := i + (j-i)/2
		if f.lines[h] <= offset {
			line = h + 1
			i = h + 1
		} else {
			j = h
		}
	}

	return Position{
		Line:   Pos(line),
		Column: offset - f.lines[line-1] + 1,
	}
}

func (f *File) setLines(src string) {
	f.lines = []Pos{0}
	for i, c := range src {
		if c == '\n' {
			f.lines = append(f.lines, Pos(i+1))
		}
	}
}
