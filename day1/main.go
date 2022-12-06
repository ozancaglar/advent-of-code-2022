package main

import (
	"fmt"
	"log"
	"ozan/utils"
	"sort"
	"strconv"
)

func main() {
	scanner := utils.StreamLines("input.txt")

	var sum int
	var cals []int

	for scanner.Scan() {
		if scanner.Text() == "" {
			cals = append(cals, sum)
			sum = 0
			continue
		}

		intFromLine, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		sum += intFromLine
	}

	sort.Ints(cals)

	fmt.Println(cals[len(cals)-1])                                         // part one
	fmt.Println(cals[len(cals)-1] + cals[len(cals)-2] + cals[len(cals)-3]) // part two
}
