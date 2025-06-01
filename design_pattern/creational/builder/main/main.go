package main

import (
	b "design-patterns/creational/builder/basic"
	d "design-patterns/creational/builder/director"
	"fmt"
)

func main() {
	car := b.NewCarBuilder().
		SetBrand("Tesla").
		SetColor("Red").
		SetWheels(4).
		AddGPS().
		Build()

	fmt.Printf("Car: %+v\n", car)

	// Director example
	directorExample()
}

func directorExample() {
	sportsCarBuilder := d.NewSportsCarBuilder()
	familyCarBuilder := d.NewFamilyCarBuilder()

	director := d.NewCarDirector(sportsCarBuilder)
	sportsCar := director.MakePerformantSportsCar()

	director = d.NewCarDirector(familyCarBuilder)
	familyCar := director.MakeFamilyCar()

	fmt.Println(sportsCar)
	fmt.Println(familyCar)
}
