// go:build inheritance
package main

import "fmt"

// base object (class in java)
// Has field- Name and method- Speak
type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println(a.Name, "speaks.")
}

// object dog with field- Breed
// go do not have inheritance, but we can still inherit using composition
// here Animal is embedded
type Dog struct {
	Animal
	Breed string
}

// Animal's Speak method is overridden
func (d *Dog) Speak() {
	fmt.Println("Dog of breed", d.Breed, "with name", d.Name, "speaks.")
}

func main() {
	dog := Dog{
		Breed:  "Golden Retriever",
		Animal: Animal{Name: "Tommy"},
	}

	dog.Speak()
}
