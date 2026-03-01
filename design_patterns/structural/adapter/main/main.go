package main

import a "design_patterns/structural/adapter"

func main() {
	newService := &a.ModernUI{}
	adapter := a.ModernToLegacyAdapter{Adaptee: newService}

	adapter.Render("Hello World!")
}
