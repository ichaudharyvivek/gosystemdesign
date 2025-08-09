package builder

// The complex object we want to build
type Car struct {
	Brand  string
	Wheels int
	Color  string
	GPS    bool
}

// Helper class that will initialize the complex object sequentially
type CarBuilder struct {
	car *Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{car: &Car{}}
}

func (b *CarBuilder) SetBrand(brand string) *CarBuilder {
	b.car.Brand = brand
	return b
}

func (b *CarBuilder) SetWheels(wheels int) *CarBuilder {
	b.car.Wheels = wheels
	return b
}

func (b *CarBuilder) SetColor(color string) *CarBuilder {
	b.car.Color = color
	return b
}

func (b *CarBuilder) AddGPS() *CarBuilder {
	b.car.GPS = true
	return b
}

// Returns the built concrete object
func (b *CarBuilder) Build() *Car {
	return b.car
}
