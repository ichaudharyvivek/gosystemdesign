// go:build isp

/*
Why this follows ISP:
  - Types only implement interfaces they actually need.
  - Robot doesn’t depend on Eat().

# Interfaces are small and focused, no unnecessary coupling.
# Equivalent java code:
```java

	interface Workable {
	    void work();
	}

	interface Eatable {
	    void eat();
	}

	class Human implements Workable, Eatable {
	    public void work() { System.out.println("Human working"); }
	    public void eat() { System.out.println("Human eating"); }
	}

	class Robot implements Workable {
	    public void work() { System.out.println("Robot working"); }
	}

```
*/
package main

import "fmt"

type Workable interface {
	Work()
}

type Eatable interface {
	Eat()
}

type Human struct{}

func (h Human) Work() {
	fmt.Println("Human working")
}

func (h Human) Eat() {
	fmt.Println("Human eating")
}

type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot working")
}
