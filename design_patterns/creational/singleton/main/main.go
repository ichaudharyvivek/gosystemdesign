package main

import (
	"design_patterns/creational/singleton"
	"fmt"
)

func main() {
	one := singleton.GetInstance("localhost", 8080)
	two := singleton.GetInstance("localhost", 8080)

	if one == two {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
}
