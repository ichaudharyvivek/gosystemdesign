package main

import a "design-patterns/structural/adapter"

func main() {
	legacy := &a.LegacyPrinter{}
	adapter := &a.PrinterAdapter{Legacy: legacy}

	adapter.Print("Hello World!")
}
