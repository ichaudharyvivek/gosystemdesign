/*
+--------------------+             +-------------------+
|     Client         |             |    Adaptee        |
|--------------------|             |-------------------|
| uses Target        |             | incompatible method|
+---------+----------+             +---------+---------+

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

// Client expect to interact with these methods
type Printable interface {
	Print(message string)
}

// Legacy class having methods that might not support new user requirements
type LegacyPrinter struct {
}

func (l *LegacyPrinter) LegacyPrint(message string) {
	fmt.Println("[Legacy]", message)
}

// New Adapter class
// Adaptee: LegacyPrinter (the one who is getting adapter)
// Adapter: PrinterAdapter (the one who is adapting older class)
type PrinterAdapter struct {
	Legacy *LegacyPrinter
}

func (p *PrinterAdapter) Print(message string) {
	p.Legacy.LegacyPrint(fmt.Sprintf("[NewPrinterAdapter] %s", message))
}
