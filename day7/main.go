package main

import (
	"errors"
	"fmt"
	"log"
	"ozan/utils"
	"strconv"
	"strings"
)

type dir struct {
	dirName          string
	files            []file
	nestedDirIndexes []int
	dirSize          *int
}

type file struct {
	fileName string
	fileSize int
}

func newDir(dirName string, files []file, dirs []int) dir {
	init := 0
	return dir{dirName: dirName, files: files, nestedDirIndexes: dirs, dirSize: &init}
}

func newFile(str string) (file, error) {
	splitString := strings.Split(str, " ")
	fileName := splitString[1]
	fileSize, err := strconv.Atoi(splitString[0])
	if err != nil {
		return file{}, errors.New("failed to obtain fileSize for file")
	}
	return file{fileName, fileSize}, nil
}

func getDirIndex(dirs []dir, dirName string) int {
	for i, dir := range dirs {
		if dir.dirName == dirName {
			return i
		}
	}

	return -1
}

func main() {
	partOne()
}

func partOne() {
	scanner := utils.StreamLines("input.txt")

	var dirs []dir
	var currDir string

	for scanner.Scan() {
		txt := scanner.Text()
		if strings.Contains(txt, "$ ls") || strings.Contains(txt, "$ cd ..") {
			continue
		}

		if strings.Contains(txt, "$ cd") {
			currDir = strings.Split(txt, " ")[2]
			if getDirIndex(dirs, currDir) == -1 {
				dirs = append(dirs, newDir(currDir, []file{}, nil))
			}
			continue
		}

		i := getDirIndex(dirs, currDir)
		if strings.Contains(txt, "dir") {
			newDirName := strings.Split(txt, " ")[1]
			j := getDirIndex(dirs, newDirName)
			if j == -1 {
				newDir := newDir(newDirName, []file{}, nil)
				dirs = append(dirs, newDir)
				dirs[i].nestedDirIndexes = append(dirs[i].nestedDirIndexes, getDirIndex(dirs, newDirName))
				continue
			}
			dirs[i].nestedDirIndexes = append(dirs[i].nestedDirIndexes, j)
			continue
		}

		f, err := newFile(txt)
		if err != nil {
			log.Fatal(err)
		}

		dirs[i].files = append(dirs[i].files, f)
	}

	for i := len(dirs) - 1; i >= 0; i-- {
		*dirs[i].dirSize = calculateDirSize(dirs[i], dirs)
	}

	for i := len(dirs) - 1; i >= 0; i-- {
		addNestedDirSizes(dirs[i], dirs)
	}

	var partOneSum int

	for i := range dirs {
		if dirSize := *dirs[i].dirSize; dirSize <= 100000 {
			partOneSum += dirSize
		}
	}

	fmt.Println(partOneSum)
}

func calculateDirSize(dir dir, dirs []dir) int {
	var sum int
	for i := range dir.files {
		sum += dir.files[i].fileSize
	}
	return sum
}

func addNestedDirSizes(dir dir, dirs []dir) {
	if len(dir.nestedDirIndexes) == 0 {
		return
	}

	for _, dirIndex := range dir.nestedDirIndexes {
		*dir.dirSize += *dirs[dirIndex].dirSize
	}
}
