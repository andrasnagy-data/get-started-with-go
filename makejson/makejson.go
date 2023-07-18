package main

/*
Ask user to enter a name and an address.
Save entered values into a map.
Make a json object from the map, then print the json object.
*/

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	personMap := map[string]string{}
	fmt.Print("Enter a name: ")
	fmt.Scan(&name)
	fmt.Print("Enter an address: ")
	fmt.Scan((&address))
	personMap["name"] = name
	personMap["address"] = address
	data, _ := json.Marshal(personMap)
	fmt.Print(data)
}
