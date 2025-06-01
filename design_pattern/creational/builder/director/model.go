package director

import "fmt"

// Car represents the complex object we want to build
type Car struct {
	Brand    string
	Wheels   int
	Color    string
	GPS      bool
	Engine   string
	SunRoof  bool
	Category string
}

func (c *Car) String() string {
	return fmt.Sprintf("Car{Brand: %s, Wheels: %d, Color: %s, GPS: %t, Engine: %s, SunRoof: %t, Category: %s}",
		c.Brand, c.Wheels, c.Color, c.GPS, c.Engine, c.SunRoof, c.Category)
}

// CarBuilder interface defines the contract for building a Car
type CarBuilder interface {
	AddBrand() CarBuilder
	AddWheels() CarBuilder
	AddColor() CarBuilder
	AddGPS() CarBuilder
	AddEngine() CarBuilder
	AddSunRoof() CarBuilder
	Reset() CarBuilder
	GetCar() *Car
}
