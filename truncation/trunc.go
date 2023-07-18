package main

import "fmt"

func main() {

	var x float32
	fmt.Print("Enter a float: ")
	fmt.Scan(&x)
	fmt.Println("The truncated number is: ", int(x))

}
