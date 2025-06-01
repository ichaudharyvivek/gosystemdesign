package main

import (
	"design-patterns/creational/singleton"
	"fmt"
)

func main() {
	one := singleton.GetInstance()
	two := singleton.GetInstance()

	if one == two {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
}
