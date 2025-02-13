package token

import (
	"testing"
)

func TestFileSet_AddFile(t *testing.T) {
	tests := []struct {
		name string
		size int
		want *File
	}{
		{"test.txt", 100, &File{name: "test.txt", base: 0, size: 100}},
		{"test2.txt", 200, &File{name: "test2.txt", base: 100, size: 200}},
		{"test3.txt", 50, &File{name: "test3.txt", base: 300, size: 50}},
		{"test4.txt", 1000, &File{name: "test4.txt", base: 350, size: 1000}},
	}

	s := NewFileSet()

	for _, tt := range tests {
		got := s.AddFile(tt.name, tt.size)
		if got.name != tt.want.name || got.base != tt.want.base || got.size != tt.want.size {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
