package director

// FamilyCarBuilder builds family cars with different configurations
type FamilyCarBuilder struct {
	car *Car
}

func NewFamilyCarBuilder() *FamilyCarBuilder {
	return &FamilyCarBuilder{
		car: &Car{}, // Initialize the car struct
	}
}

func (b *FamilyCarBuilder) Reset() CarBuilder {
	b.car = &Car{}
	return b
}

func (b *FamilyCarBuilder) AddBrand() CarBuilder {
	b.car.Brand = "Toyota"
	return b
}

func (b *FamilyCarBuilder) AddWheels() CarBuilder {
	b.car.Wheels = 4
	return b
}

func (b *FamilyCarBuilder) AddColor() CarBuilder {
	b.car.Color = "Silver"
	return b
}

func (b *FamilyCarBuilder) AddGPS() CarBuilder {
	b.car.GPS = true // Family cars need navigation
	return b
}

func (b *FamilyCarBuilder) AddEngine() CarBuilder {
	b.car.Engine = "2.0L Hybrid"
	return b
}

func (b *FamilyCarBuilder) AddSunRoof() CarBuilder {
	b.car.SunRoof = true // Comfort for family
	return b
}

func (b *FamilyCarBuilder) GetCar() *Car {
	b.car.Category = "Family"
	result := b.car
	b.Reset() // Reset for next build
	return result
}
