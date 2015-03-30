package main

import "fmt"

//Definig types
type Salutation struct {
	Name string
	Last string
}

type Salutations []Salutation

//The renamable interface
type Renamable interface {
	Rename(newName string)
}

func (salutations Salutations) Greet() {

	for _, s := range salutations {
		fmt.Println("Hello: " + s.Name)
	}
}

//Operates on copy
func (sal Salutation) RenameOld(newName string) {
	sal.Name = newName
}

//Operates on actual data
func (sal *Salutation) Rename(newName string) {
	sal.Name = newName
}

//An "object" that impliments the Renamable interface is required
func RenameToFrog(r Renamable) {
	r.Rename("frog")
}

func main() {

	salutations := Salutations{
		{"Joe", "Puzzo"},
		{"Anthony", "Puzzo"},
		{"Bill", "Jones"},
	}

	salutations.Greet()

	RenameToFrog(&salutations[0])
}
