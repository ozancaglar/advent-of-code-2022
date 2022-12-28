package main

import (
	"fmt"
	"os"
	"ozan/utils"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Create a scanner to read from the file
	scanner := utils.StreamLines("input.txt")

	// Create a 2d slice to hold the matrix
	var matrix [][]int

	// Read in each line of the matrix
	for scanner.Scan() {
		// Split the line into a slice of strings
		rowStrings := strings.Split(scanner.Text(), "")

		// Convert the slice of strings to a slice of ints
		var rowInts []int
		for _, s := range rowStrings {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			rowInts = append(rowInts, n)
		}

		// Add the row to the matrix
		matrix = append(matrix, rowInts)
	}

	// Initialize the visible slice with the same dimensions as the matrix
	visible := make([][]bool, len(matrix))
	for i := range visible {
		visible[i] = make([]bool, len(matrix[0]))
	}

	// Set all elements in visible to true
	for i := range visible {
		for j := range visible[i] {
			visible[i][j] = false
		}
	}

	// Iterate through each element in the matrix
	for i, row := range matrix {
		for j, x := range row {
			if i == 0 || j == 0 || i == len(matrix)-1 || j == len(matrix[0])-1 {
				// This element is on the edge of the matrix, so it is always visible
				visible[i][j] = true
				continue
			}

			// iterate to see if visible left or right
			for row := range matrix {
				if row != i {
					continue
				}

				treeHeightsToLeft := []int{}
				treeHeightsToRight := []int{}
				for i := range matrix[row] {
					if i == j {
						continue
					}

					// visible from left
					if i < j {
						treeHeightsToLeft = append(treeHeightsToLeft, matrix[row][i])
					}

					if i > j {
						treeHeightsToRight = append(treeHeightsToRight, matrix[row][i])
					}
				}

				sort.Ints(treeHeightsToLeft)
				sort.Ints(treeHeightsToRight)

				if treeHeightsToLeft[len(treeHeightsToLeft)-1] < x || treeHeightsToRight[len(treeHeightsToRight)-1] < x {
					visible[i][j] = true
				}
			}

			treeHeightsAbove := []int{}
			treeHeightsBelow := []int{}
			for row := range matrix {
				if row == i {
					continue
				}

				if row < i {
					treeHeightsAbove = append(treeHeightsAbove, matrix[row][j])
				}

				if row > i {
					treeHeightsBelow = append(treeHeightsBelow, matrix[row][j])
				}
			}
			sort.Ints(treeHeightsAbove)
			sort.Ints(treeHeightsBelow)

			if treeHeightsAbove[len(treeHeightsAbove)-1] < x || treeHeightsBelow[len(treeHeightsBelow)-1] < x {
				visible[i][j] = true
			}
		}
	}

	// Iterate through each element in the matrix
	numberOfVisibleTrees := 0
	for _, row := range visible {
		for j := range row {
			if row[j] == true {
				numberOfVisibleTrees += 1
			}
		}
	}

	// Print the result
	fmt.Println(numberOfVisibleTrees)
}
