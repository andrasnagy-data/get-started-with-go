package main

import (
	"fmt"
	"strings"
)

func main() {
	var userString string

	fmt.Print("Enter a string: ")
	fmt.Scan(&userString)

	userStringLower := strings.ToLower(userString)
	if strings.HasPrefix(userStringLower, "i") &&
		strings.Contains(userStringLower, "a") &&
		strings.HasSuffix(userStringLower, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
