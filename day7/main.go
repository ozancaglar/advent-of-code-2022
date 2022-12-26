package main

import (
	"errors"
	"fmt"
	"log"
	"ozan/utils"
	"sort"
	"strconv"
	"strings"
)

type dir struct {
	name     string
	files    []file
	parent   *dir
	children *[]dir
	size     int
}

type file struct {
	name string
	size int
}

func newDir(name string, parent *dir) dir {
	return dir{name: name, files: []file{}, parent: parent, children: &[]dir{}, size: 0}
}

func newFile(str string) (file, error) {
	splitString := strings.Split(str, " ")
	name := splitString[1]
	size, err := strconv.Atoi(splitString[0])
	if err != nil {
		return file{}, errors.New("failed to obtain size for file")
	}
	return file{name, size}, nil
}

func main() {
	partOne()
}

func partOne() {
	rootDir, allDirs := buildTree()
	var partOneSum int

	for i := len(allDirs) - 1; i >= 0; i-- {
		if allDirs[i].parent != nil {
			allDirs[i].parent.size += allDirs[i].size
		}
	}

	for _, d := range allDirs {
		if d.size <= 100000 {
			partOneSum += d.size
		}
	}
	fmt.Println(partOneSum)

	fmt.Println(partTwo(rootDir, allDirs))
}

func partTwo(rootDir *dir, allDirs []*dir) int {
	spaceNeeded := 30000000 - (70000000 - rootDir.size)
	var deletableSizes []int
	for _, d := range allDirs {
		if spaceNeeded-d.size < 0 {
			deletableSizes = append(deletableSizes, d.size)
		}
	}
	sort.Ints(deletableSizes)
	return deletableSizes[0]
}

func buildTree() (*dir, []*dir) {
	scanner := utils.StreamLines("input.txt")
	rootDir := newDir("/", nil)
	currDir := &rootDir
	allDirs := []*dir{}
	allDirs = append(allDirs, currDir)

	for scanner.Scan() {
		txt := scanner.Text()
		if strings.Contains(txt, "$ ls") || strings.Contains(txt, "$ cd /") || strings.Contains(txt, "dir") {
			continue
		}

		if txt == "$ cd .." {
			currDir = currDir.parent
			continue
		}

		if strings.Contains(txt, "$ cd") {
			newDirName := strings.Split(txt, " ")[2]
			newDir := newDir(newDirName, currDir)
			*currDir.children = append(*currDir.children, newDir)
			currDir = &newDir
			allDirs = append(allDirs, currDir)

			continue
		}

		f, err := newFile(txt)
		if err != nil {
			log.Fatal(err)
		}

		currDir.files = append(currDir.files, f)
		currDir.size += f.size
	}

	return &rootDir, allDirs
}
