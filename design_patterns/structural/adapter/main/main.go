package main

import a "design-patterns/structural/adapter"

func main() {
	newService := &a.ModernUI{}
	adapter := a.LegacyToModernAdapter{Adapter: newService}

	adapter.Render("Hello World!")
}
