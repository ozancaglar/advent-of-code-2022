package main

import (
	"fmt"
	"ozan/utils"
)

func main() {
	scanner := utils.StreamLines("input.txt")
	for scanner.Scan() {
		fmt.Printf("Part One: %d\n", findMarker(scanner.Text(), 4))
		fmt.Printf("Part Two: %d\n", findMarker(scanner.Text(), 14))
	}
}

func findMarker(input string, size int) int {
	var arr []string
	var markerReached int
	for i, c := range input {
		if i <= size-1 {
			arr = append(arr, string(c))
			if len(arr) == size {
				if uniqueLetters(arr) {
					return i + 1
				}
			}
			continue
		}

		arr = keepArraySizeConstant(arr, string(c))
		if uniqueLetters(arr) {
			return i + 1
		}
	}

	return markerReached
}

func keepArraySizeConstant(arr []string, toAppend string) []string {
	arr = arr[1:]
	arr = append(arr, toAppend)
	return arr
}

func uniqueLetters(arr []string) bool {
	m := map[string]int{}
	for _, c := range arr {
		m[string(c)] += 1
	}
	for k := range m {
		if m[k] >= 2 {
			return false
		}
	}
	return true
}
