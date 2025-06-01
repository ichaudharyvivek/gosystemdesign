package director

// here, the various concrete class/object of Car is built
// if we use basic build with .AddBrand().AddColor()... - it becomes messy when we have to build multiple object of same type
// example say if we have to build a sports car, but only brand has to change,
// with primitive way, we will have to chain method multiple times possibly duplicating code
type SportsCarBuilder struct {
	car *Car
}

func NewSportsCarBuilder() *SportsCarBuilder {
	return &SportsCarBuilder{
		car: &Car{}, // Initialize the car struct
	}
}

func (b *SportsCarBuilder) Reset() CarBuilder {
	b.car = &Car{}
	return b
}

func (b *SportsCarBuilder) AddBrand() CarBuilder {
	b.car.Brand = "Ferrari"
	return b
}

func (b *SportsCarBuilder) AddWheels() CarBuilder {
	b.car.Wheels = 4
	return b
}

func (b *SportsCarBuilder) AddColor() CarBuilder {
	b.car.Color = "Red"
	return b
}

func (b *SportsCarBuilder) AddGPS() CarBuilder {
	b.car.GPS = false // Sports cars focus on performance, not navigation
	return b
}

func (b *SportsCarBuilder) AddEngine() CarBuilder {
	b.car.Engine = "V8 Twin Turbo"
	return b
}

func (b *SportsCarBuilder) AddSunRoof() CarBuilder {
	b.car.SunRoof = false // Aerodynamics over comfort
	return b
}

func (b *SportsCarBuilder) GetCar() *Car {
	b.car.Category = "Sports"
	result := b.car
	b.Reset() // Reset for next build
	return result
}
