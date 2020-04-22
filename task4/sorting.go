package main

import (
	"fmt"
	//"sort"
	"time"
)

type Interface interface {
  Len() int
  Less(i, j int) bool
  Swap(i, j int)
}

type Person struct {
  firstName string
  lastName string
  birthDay time.Time
}

type People []Person

func (pl People) Len() int {
  return len(pl)
}

func (pl People) Less(i, j int) bool {
  return !(pl[i].birthDay.Sub(pl[j].birthDay) < 0)
}

/*func (pl People) Swap(i, j int) {
  if Less(i, j) {
    p := pl[i]
    pl[i] = pl[j]
    pl[j] = p
  }
}*/

func main() {
  person1 := Person{"Sasha", "Kravchenko", time.Date(2000, 1, 4, 0, 0, 0, 0, time.UTC)}
  person2 := Person{"Kirill", "Gasmanov", time.Date(2000, 2, 7, 0, 0, 0, 0, time.UTC)}
  person3 := Person{"Nastya", "Kahovskaya", time.Date(2000, 2, 7, 0, 0, 0, 0, time.UTC)}
  var people = People{person1, person2, person3}
    fmt.Println(people.Len())
    fmt.Println(people.Less(0,1))
    // i dont know how to properly use swap and sort in this task
    // should they be independent from each other?
    // or i just need to use sort with custom comparators to finish this task?
}
