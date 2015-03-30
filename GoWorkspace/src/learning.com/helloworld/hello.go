package main

import "learning.com/greeting"

import "fmt"

func main() {

	//Basic assignment
	var message string = "Hello Go!"
	var a, b, c int = 1, 2, 3
	d, e := true, "test"

	//Pointers
	var greet *string = &message

	var s = greeting.Salutation{"Joe", "Hello"}
	var s2 = greeting.Salutation{Greeting: *greet, Name: "Joe"}

	fmt.Println(message, a, b, c, d, e)
	fmt.Println(s.Name)
	fmt.Println(s.Greeting)

	greeting.Greet(s, greeting.PrintLine, true)
	greeting.Greet(s2, greeting.PrintLine, false)
}
