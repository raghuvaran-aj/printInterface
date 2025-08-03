package AreaPrint

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	fmt.Print("Print square Area function")
	return s.side * s.side
}
func (c Circle) Area() float64 {
	fmt.Println("print circle area function")
	return 3.14 * c.radius * c.radius
}

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}
