package main

import "fmt"

func main() {

	//Map strings to strings
	var mymap map[string]string
	//Build the map
	mymap = make(map[string]string)

	mymap["Bob"] = "Jones"
	mymap["Joe"] = "Puzzo"
	mymap["Connor"] = "Stone"

	fmt.Println("Joe " + mymap["Joe"])

	/*----Shorthand map----*/
	mymap2 := map[string]string{
		"Bob":    "Jones",
		"Joe":    "Puzzo",
		"Connor": "Stone",
	}

	//You need that last comma!!

	/*----Map operations----*/
	//Insert and update
	mymap2["Gregory"] = "Raso"

	//Delete
	delete(mymap2, "Gregory")

	//Check for existance
	//Note the below is an examle of the assignent within an if and the exists is the condition
	if value, exists := mymap2["Joe"]; exists {
		fmt.Println(value + " exists")
	}
}
