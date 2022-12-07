package day07

import (
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int64
}

type Filesystem struct {
	Name        string
	Parent      *Filesystem
	Files       []*File
	Directories []*Filesystem
}

func (f Filesystem) Size() int64 {
	return 0
}

func NewFilesystem(n string, p *Filesystem) *Filesystem {
	return &Filesystem{
		Name:        n,
		Parent:      p,
		Files:       make([]*File, 0),
		Directories: make([]*Filesystem, 0),
	}
}

func NewFilesystemFromInput(input []string) *Filesystem {

	var root = NewFilesystem("/", nil)
	var cwd = root

scanner:
	for i := 1; i <= len(input)-1; { // ignore the first line since this is cd to root
		ln := input[i]
		cmd := strings.TrimSpace(ln[1:])
		switch cmd {
		case "ls":
			for j := i + 1; j <= len(input)-1; j++ {
				ln := input[j]
				if strings.HasPrefix(ln, "$") {
					i = j
					continue scanner
				}

				if strings.HasPrefix(ln, "dir") {
					cwd.Directories = append(cwd.Directories, NewFilesystem(ln[4:], cwd))
					continue
				}

				fparts := strings.Split(ln, " ")

				fsize, err := strconv.ParseInt(fparts[0], 10, 64)
				if err != nil {
					panic(err)
				}

				cwd.Files = append(cwd.Files, &File{
					Name: fparts[1],
					Size: fsize,
				})
			}
		default: // Assume this is a CD command until proven wrong
			dir := cmd[3:]
			switch dir {
			case "..":
				if cwd.Parent == nil {
					break scanner
				}

				cwd = cwd.Parent
			default:
				for _, cwddir := range cwd.Directories {
					if dir == cwddir.Name {
						cwd = cwddir
						break
					}
				}
			}
		}
		i++
	}

	return root
}
