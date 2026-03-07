package director

type CarDirector struct {
	builder CarBuilder
}

func NewCarDirector(b CarBuilder) *CarDirector {
	return &CarDirector{builder: b}
}

func (d *CarDirector) MakeSportsCar() *Car {
	return d.builder.AddBrand().AddColor().AddWheels().AddGPS().GetCar()
}

// SetBuilder allows changing the builder at runtime
func (d *CarDirector) SetBuilder(builder CarBuilder) {
	d.builder = builder
}

// MakeBasicCar - builds a car with only essential features
func (d *CarDirector) MakeBasicCar() *Car {
	return d.builder.
		AddBrand().
		AddWheels().
		AddColor().
		GetCar()
}

// MakeSportsCar - builds a sports car with performance-focused sequence
func (d *CarDirector) MakePerformantSportsCar() *Car {
	builder := NewSportsCarBuilder()
	return builder.
		AddEngine(). // Engine first for performance
		AddBrand().
		AddWheels().
		AddColor().
		// Skip GPS and SunRoof for pure performance
		GetCar()
}

// MakeFamilyCar - builds a family car with comfort-focused sequence
func (d *CarDirector) MakeFamilyCar() *Car {
	builder := NewFamilyCarBuilder()
	return builder.
		AddBrand().
		AddGPS().     // GPS first for navigation
		AddSunRoof(). // Comfort features
		AddEngine().
		AddWheels().
		AddColor().
		GetCar()
}
