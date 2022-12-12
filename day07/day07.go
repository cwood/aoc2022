package day07

import (
	"sort"
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

	size int64
}

func (f Filesystem) Size() int64 {
	if f.size != 0 {
		return f.size
	}

	for _, fi := range f.Files {
		f.size += fi.Size
	}

	for _, d := range f.Directories {
		f.size += d.Size()
	}

	return f.size
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

func computeSize(prefix string, fs *Filesystem, m map[string]int64) {
	for _, dir := range fs.Directories {
		m[prefix+"/"+dir.Name] = dir.Size()
		computeSize(prefix+"/"+dir.Name, dir, m)
	}
}

func calculateSizes(fs *Filesystem) map[string]int64 {
	var dirs = make(map[string]int64)
	computeSize("", fs, dirs)
	dirs["/"] = fs.Size()
	return dirs
}

// DirectoriesNeariest will take a filesystem and try to find all directories nearest that value and
// sum them and return that size
func DirectoriesNeariest(fs *Filesystem, nearest int64) int64 {
	dirs := calculateSizes(fs)

	var total int64 = 0

	for dir, size := range dirs {
		if size > nearest {
			delete(dirs, dir)
			continue
		}

		total += size
	}

	return total
}

func FreeUpSpace(fs *Filesystem, totalSize int64, atleast int64) int64 {
	dirs := calculateSizes(fs)

	var sizes = make([]int64, 0)

	used := totalSize - dirs["/"]
	minToFree := atleast - used

	for _, size := range dirs {
		if size >= minToFree {
			sizes = append(sizes, size)
		}
	}

	sort.Slice(sizes, func(x int, y int) bool {
		return sizes[x] > sizes[y]
	})

	return sizes[len(sizes)-1]
}
