// déclaration du package pour ce fichier
package main

// importation du package fmt
import (
	"fmt"
	"math"
)

func main() {
	// invocation de la fonction Println déclarée dans le package fmt
	fmt.Println("Hello, World!")
}

type BaseShape struct {
	name      string
	isPolygon bool
}

type Rectangle struct {
	BaseShape
	side1 float64
	side2 float64
}

type Square struct {
	Rectangle
}

type Circle struct {
	BaseShape
	radius float64
}

type Shape interface {
	GetArea() float64
	GetName() string
	IsPolygon() bool
}

func NewBaseShape(name string, shape bool) BaseShape {
	return BaseShape{name, shape}
}

func (bs BaseShape) GetName() string {
	return bs.name
}
func (bs BaseShape) IsPolygon() bool {
	return bs.isPolygon
}

func (rect Rectangle) GetArea() float64 {
	return rect.side1 * rect.side2
}

func (cir Circle) GetArea() float64 {
	return cir.radius * math.Pi
}

func NewRectangle(name string, side1, side2 float64) Rectangle {
	return Rectangle{
		NewBaseShape(name, true),
		side1,
		side2,
	}
}

func NewSquare(name string, side float64) Square {
	return Square{
		NewRectangle(name, side, side),
	}
}

func NewCircle(name string, radius float64) Circle {
	return Circle{
		NewBaseShape(name, false),
		radius,
	}
}

func displayShape(shape Shape) {

}
