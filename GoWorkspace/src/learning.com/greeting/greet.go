package greeting

import "fmt"

//Definig types
type Salutation struct {
	Name     string
	Greeting string
}

const (
	PI       = 3.14
	Language = "GO"
)

const (
	A = iota
	B = iota
	C = iota
)

//Note index out of bounds will be thrown
func CreateMessage(name string, greeting ...string) (string, string) {
	return greeting[0] + " " + name, "Hey! " + name
}

//Look You can name return values!!!
func createMessageRefactor(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey! " + name
	return
}

type Printer func(string)

func Greet(sal Salutation, do Printer, isFormal bool) {
	//Note how I "throw away one of the return values"
	message, _ := CreateMessage(sal.Name, sal.Greeting)
	if prefix := "Mr "; isFormal {
		do(prefix + message)
	} else {
		do("FUCK")
	}
}

//Functions for do
func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func TestSwitch1(s string) (prefix string) {
	switch s {
	case "Bob":
		prefix = "Mr "
		fallthrough
	case "Joe", "Amy":
		prefix = "Dr "
	default:
		prefix = "Dude "
	}
	return
}

func TestSwitch2(s string) (prefix string) {
	switch {
	case s == "Bob":
		prefix = "Mr "
	case s == "Joe", s == "Amy", len(s) == 10:
		prefix = "Dr "
	default:
		prefix = "Dude "
	}
	return
}

func TypeSwitch(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

//Looping over a "Range
func RangeLoop(aslice []string) {
	//Note: if i is not used you must use an underscore in its place
	for i, s := range aslice {
		fmt.Println("Index: " + i)
		fmt.Println("Data: " + s)
	}
}
