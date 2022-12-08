package types

import (
	"advent-of-code-2022/pkg/util"
)

type File struct {
	Name string
	Size int
}

func NewFile(name, size string) *File {
	return &File{
		Name: name,
		Size: util.MustParseInt(size),
	}
}

type Directory struct {
	Name        string
	Parent      *Directory
	Directories map[string]*Directory
	Files       []*File
	Size        int
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		Name:        name,
		Parent:      parent,
		Directories: make(map[string]*Directory, 0),
		Files:       make([]*File, 0),
		Size:        0,
	}
}

func (d *Directory) AddDirectory(name string, parent *Directory) {
	d.Directories[name] = NewDirectory(name, parent)
}

func (d *Directory) AddFile(name, size string) {
	d.Files = append(d.Files, NewFile(name, size))
}

func (d *Directory) Visit(f func(*Directory)) {
	for _, directory := range d.Directories {
		directory.Visit(f)
	}
	f(d)
}
