package main

import (
	"fmt"
	"sort"
)

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	return Point{x: s.start.x + int(s.a), y: s.start.y - int(s.a)}
}

func (s Square) Perimeter() uint {
	return 4 * s.a
}

func (s Square) Area() uint {
	return s.a * s.a
}

func median(i []int) float64 {
	sort.Ints(i)
	if len(i)%2 == 0 {
		return (float64(i[len(i)/2-1]) + float64(i[len(i)/2])) / 2
	}
	return float64(i[((len(i)+1)/2)-1])
}

func main() {
	// task 1
	var arrayOne = []int{1, 8, 3, 10, 5}
	var arrayTwo = []int{5, 1, 9, 6, 7, 2}
	fmt.Println(median(arrayOne))
	fmt.Println(median(arrayTwo))
	fmt.Println()
	// task 2
	s := Square{start: Point{1, 1}, a: 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
