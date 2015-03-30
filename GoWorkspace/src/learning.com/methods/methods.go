package main

import "fmt"

//Definig types
type Salutation struct {
	Name string
	Last string
}

type Salutations []Salutation

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

func main() {

	salutations := Salutations{
		{"Joe", "Puzzo"},
		{"Anthony", "Puzzo"},
		{"Bill", "Jones"},
	}

	salutations.Greet()

}
