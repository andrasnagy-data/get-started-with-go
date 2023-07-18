package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slc := make([]int, 3)
	for {
		var tempIntStr string
		// ask user for an integer
		fmt.Print("Enter an integer or 'X' to exit the program: ")
		// "save" input into tempIntStr
		fmt.Scanln(&tempIntStr)
		// check if entered value is "X"
		if tempIntStr == "X" {
			break
		} else {
			// convert string to integer
			integer, _ := strconv.Atoi(tempIntStr)
			// append integer to slice
			slc = append(slc, integer)
			// sort slice
			sort.Slice(slc, func(i int, j int) bool {
				return slc[i] < slc[j]
			})
			// print slice
			fmt.Print(slc)
		}
	}
}
