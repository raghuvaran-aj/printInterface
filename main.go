package main

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

// func main() {
// 	fmt.Println("Hello Welcome to this printInterface main function")
// 	val := Square{side: 10.0}
// 	fmt.Println("square val:", val.Area())

// 	cval := Circle{radius: 5.0}
// 	fmt.Printf("circle area: %f", cval.Area())
// }
