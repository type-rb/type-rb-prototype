package token

type FileSet struct {
	base  Pos
	files []*File
	last  *File
}

func NewFileSet() *FileSet {
	return &FileSet{}
}

func (s *FileSet) AddFile(name string, size int) *File {
	f := &File{name: name, base: s.base, size: size}
	s.files = append(s.files, f)
	s.last = f
	s.base += Pos(size)
	return f
}
