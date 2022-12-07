package main

import (
	"aoc/utils"
	"fmt"
	"strings"
	"time"
)

const (
	maxFileSystemSize = 70_000_000
	updateSize        = 30_000_000
)

type Folder struct {
	name      string
	parent    *Folder
	files     []File
	folders   map[string]*Folder
	sizeCache *int
}

func (f *Folder) Size() int {
	if f.sizeCache != nil {
		return *f.sizeCache
	}
	total := 0
	for _, file := range f.files {
		total += file.size
	}
	for _, folder := range f.folders {
		total += folder.Size()
	}
	f.sizeCache = &total
	return total
}

type File struct {
	name string
	size int
}

func main() {
	defer utils.MeasureTime(time.Now())
	rootFolder := buildFS("input.txt")
	needToFreeUpSpace := updateSize - (maxFileSystemSize - rootFolder.Size())
	part1, part2 := walkDirectories(*rootFolder, needToFreeUpSpace, 0)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func buildFS(input string) *Folder {
	inputs := utils.ReadAndSplit(input, "$ ")
	current := &Folder{
		name:    "root",
		parent:  nil,
		files:   make([]File, 0),
		folders: make(map[string]*Folder),
	}

	rootFolder := current

	for _, line := range inputs[1:] {
		lines := strings.Split(line, "\n")
		switch {
		case lines[0] == "ls":
			parseLS(current, lines[1:])
		case lines[0] == "cd /":
			current = rootFolder
		case lines[0] == "cd ..":
			current = current.parent
		default: // cd into a dir
			folder := strings.TrimPrefix(lines[0], "cd ")
			current = current.folders[folder]
		}
	}

	return rootFolder
}

func parseLS(f *Folder, lines []string) {
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		parts := strings.Split(l, " ")

		if strings.HasPrefix(parts[0], "dir") { // parse folder
			f.folders[parts[1]] = &Folder{
				name:    parts[1],
				parent:  f,
				files:   make([]File, 0),
				folders: make(map[string]*Folder),
			}
			continue
		}
		// parse file
		f.files = append(f.files, File{
			name: parts[1],
			size: utils.MustParseInt(parts[0]),
		})
	}
}

func walkDirectories(root Folder, needToUpdate int, depth int) (int, int) {
	part1 := 0
	part2 := maxFileSystemSize
	size := root.Size()
	if size <= 100_000 {
		part1 += size
	}
	if size >= needToUpdate {
		part2 = size
	}

	for _, f := range root.folders {
		p1, p2 := walkDirectories(*f, needToUpdate, depth+3)
		part1 += p1
		part2 = utils.Min(part2, p2)
	}
	return part1, part2
}
