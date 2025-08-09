// go:build lsp

/*
Why this follows Liskov Substitution Principle (LSP):
  - Sparrow implements FlyingBird and can fly.
  - Ostrich only implements Bird, so it can be used in contexts that don’t require flying.

# All subtypes can safely replace their parent types.
# Equivalent java code:
```java

	interface Bird {}

	interface FlyingBird extends Bird {
	    void fly();
	}

	class Sparrow implements FlyingBird {
	    public void fly() { System.out.println("Sparrow flying"); }
	}

	class Ostrich implements Bird {
	    // No fly method — not expected to fly
	}

```
*/
package main

import "fmt"

// Imagine if Bird interface had fly() method, then we couldn't have Ostrich implement Bird
// That would mean return null or throw exception in Ostrich's fly() implementation => violation of lsp
type Bird interface{}

type FlyingBird interface {
	Bird
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow flying")
}

type Ostrich struct{} // No Fly() method — OK
