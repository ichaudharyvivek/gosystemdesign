package adapter

import (
	"fmt"
)

/*
The Adapter Design Pattern allows objects with incompatible interfaces to work
together by providing a wrapper (adapter) that translates one interface to another.

Diagram:
+--------------------+             +---------------------+
|     Client         |             |    Adaptee          |
|--------------------|             |---------------------|
| Uses Target        |             | incompatible method |
+---------+----------+             +---------+-----------+
          |                                  ^
          |                                  |
          v                                  |
+--------------------+             +---------------------+
|    <<interface>>   |             |    Adapter          |
|     Target         | <---------- |---------------------|
|--------------------| implements  | - adaptee: Adaptee  |
| + request()        |             | + request()         |
+--------------------+             +---------------------+

Real world scenario:
Your app renders UI using XML data.
You want to use a new ModernUI library that only speaks JSON.
Instead of rewriting your entire app, we wrap ModernUI in an adapter
that accepts XML (familiar interface) and converts it to JSON internally.
*/

type XMLData string
type JSONData string

// Target: the interface your app already knows and uses
type UI interface {
	Render(xml XMLData)
}

// LegacyUI: old implementation of Target, renders XML directly
type LegacyUI struct{}

func (lui *LegacyUI) Render(xml XMLData) {
	fmt.Printf("[XML DATA] %s\n", xml)
}

// Adaptee: new incompatible library that only understands JSON
type ModernUI struct{}

func (mui *ModernUI) RenderJSONData(json JSONData) {
	fmt.Printf("[JSON DATA] %s\n", json)
}

// Adapter: wraps the Adaptee (ModernUI) and implements the Target (UI) interface
// so the client can keep calling Render(xml) without any changes
type ModernToLegacyAdapter struct {
	Adaptee *ModernUI
}

func (a *ModernToLegacyAdapter) Render(xml XMLData) {
	// Convert XML to JSON and delegate to Adaptee
	json := JSONData(xml)
	a.Adaptee.RenderJSONData(json)
}
