// go:build ocp

/*
Why this follows OCP:
  - Shape is an interface.
  - You can add new shapes (like Rectangle, Triangle) without changing existing Shape or area() logic.

# No need to modify old code — just extend with new types.
# Closed to modification, open to extension.
# Equivalent java code:
```java

	interface Shape {
	    double area();
	}

	class Circle implements Shape {
	    public double radius;
	    public Circle(double radius) { this.radius = radius; }
	    public double area() { return Math.PI * radius * radius; }
	}

	class Square implements Shape {
	    public double side;
	    public Square(double side) { this.side = side; }
	    public double area() { return side * side; }
	}

```
*/
package main

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}
