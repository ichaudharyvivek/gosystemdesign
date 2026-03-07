package main

import (
	b "designpatterns/creational/builder/basic"
	d "designpatterns/creational/builder/director"
	f "designpatterns/creational/builder/functional"
	"fmt"
	"time"
)

func main() {
	// Basic example
	car := b.NewCarBuilder().
		SetBrand("Tesla").
		SetColor("Red").
		SetWheels(4).
		AddGPS().
		Build()
	fmt.Printf("Car: %+v\n", car)

	// Director example
	directorExample()

	// Functional options example
	functionalOptionExample()
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

func functionalOptionExample() {
	// With base configurations
	c1 := f.NewHTTPClient()
	c1.Ping()

	// With custom configuration
	c2 := f.NewHTTPClient(f.WithBaseURL("https://google.com"), f.WithRetryCount(1), f.WithTimeout(300*time.Millisecond))
	c2.Ping()
}
