/*
The Adapter Design Pattern is a structural design pattern that allows objects with
incompatible interfaces to work together by providing a wrapper (the adapter) that
translates one interface into another.
Diagram:
+--------------------+             +---------------------+
|     Client         |             |    Adaptee          |
|--------------------|             |---------------------|
| Uses Target        |             | incompatible method |
+---------+----------+             +---------+-----------+

	|                                  ^
	|                                  |
	v                                  |

+--------------------+            +---------+----------+
|    <<interface>>   |            |    Adapter         |
|     Target         |<-----------|--------------------|
|--------------------| implements | - adaptee: Adaptee |
| + request()        |            | + request()        |
+--------------------+            +--------------------+
*/
package adapter

import (
	"fmt"
)

// Imagine a situation, where the UI in your app is rendered via XML data
// Now you want to expand your UI with new library that uses JSON data
// But since your app's method signature can't use JSON, we need to create a adapter to transform xml -> json
type XMLData string
type JSONData string

type UI interface {
	Render(xml XMLData)
}

// LegacyUI implements the UI interface to render the UI on screen using XML data
type LegacyUI struct{}

func (lui *LegacyUI) Render(xml XMLData) {
	fmt.Printf("[XML DATA] %s\n", xml)
}

// New modern UI uses only JSON
type ModernUI struct{}

func (mui *ModernUI) RenderJSONData(json JSONData) {
	fmt.Printf("[JSON DATA] %s\n", json)
}

// New Adapter class
// Adaptee:  LegacyUI (the one who is getting adapter)
// Adapter: ModernUI (the one who is adapting older class)
type LegacyToModernAdapter struct {
	Adapter *ModernUI
}

// We kept the interface defination same but the new method adapts to JSON data
// So now, a user can still call x.Render(xml) => but the UI will render modern UI with JSON data
func (a *LegacyToModernAdapter) Render(xml XMLData) {
	// Convert XMLData to JSON
	json := JSONData(xml)

	// Finally call the adapter method
	a.Adapter.RenderJSONData(json)
}
