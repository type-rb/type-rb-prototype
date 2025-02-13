package token

import (
	"testing"
)

func TestFile_Position(t *testing.T) {
	src := "line1\nline2\nline3\n"
	file := NewFile("test.txt", 100, src)

	tests := map[string]struct {
		pos  Pos
		want Position
	}{
		"start of file":   {100, Position{Line: 1, Column: 1}},
		"middle of line1": {103, Position{Line: 1, Column: 4}},
		"end of line1":    {105, Position{Line: 1, Column: 6}},
		"start of line2":  {106, Position{Line: 2, Column: 1}},
		"middle of line2": {108, Position{Line: 2, Column: 3}},
		"end of line2":    {111, Position{Line: 2, Column: 6}},
		"start of line3":  {112, Position{Line: 3, Column: 1}},
		"middle of line3": {114, Position{Line: 3, Column: 3}},
		"end of line3":    {117, Position{Line: 3, Column: 6}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			pos := tt.pos
			got := file.Position(pos)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
