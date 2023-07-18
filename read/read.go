package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	var fileName string
	namesSlice := make([]Name, 0)

	// do not leave for loop, until entered file exists
	for {
		fmt.Print("Enter the name of the text file: ")
		fmt.Scan(&fileName)
		if _, err := os.Stat(fileName); err == nil {
			break // text file exists
		} else if errors.Is(err, os.ErrNotExist) {
			// text file does *not* exist
			fmt.Print("Re-enter the name of the text file: ")
			// make fileName an empty string and "collect" entered file name again
			fileName := ""
			fmt.Scan(&fileName)
		}
	}

	readFile, err := os.Open(fileName)
	if err != nil {
		// if file does not exists, prompt user again
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		// read in file line by line
		names := strings.Split(fileScanner.Text(), " ")
		// init empty instance
		n := Name{}
		// set names with a limit of max 20 chars per field
		n.setNames(names)
		// append to slice of Names
		namesSlice = append(namesSlice, n)
	}
	fmt.Println(namesSlice)

}

type Name struct {
	fname string
	lname string
}

func (n *Name) setNames(names []string) {
	var rs []rune
	maxLength := 20 // as per the assignment specification

	n.fname = names[0]
	if utf8.RuneCountInString(names[0]) > maxLength {
		rs = []rune(names[0])
		n.fname = string(rs[:maxLength])
	}

	n.lname = names[1]
	if utf8.RuneCountInString(names[1]) > maxLength {
		rs = []rune(names[1])
		n.lname = string(rs[:maxLength])
	}
}
