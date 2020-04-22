package main

import "fmt"

type Figure interface {
  area() float64
  perimeter() float64
}

type Square struct {
  a float64
}
type Circle struct {
  a float64
}

func (s Square) area() float64 {
  return s.a * s.a
}
func (s Square) perimeter() float64 {
  return 4 * s.a
}
func (c Circle) area() float64 {
  return 3.14 * c.a * c.a
}
func (c Circle) perimeter() float64 {
  return 2 * 3.14 * c.a
}
func main() {
  var s Figure = Square{2.56}
  var c Figure = Circle{3.78}
  fmt.Println(s.area(), s.perimeter())
  fmt.Println(c.area(), c.perimeter())
}
