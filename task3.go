package main

import "fmt"

func average(arr [6]int) float64 {
  var sum int
  for _, a := range arr {
    sum += a
  }
  return float64(sum) / float64(len(arr))
}

func max (sl []string) string {
  min, ind := 0, 0
  for i, a := range sl {
    if(len(a) > min) {
      min = len(a)
      ind = i
    }
  }
  return sl[ind]
}

func reverse (sl []int64) []int64 {
  newSl := make([]int64, len(sl), len(sl))
  for i, a := range sl {
    newSl[len(newSl) - 1 - i] = a
  }
  return newSl
}

func main() {
  //task 1
  var array = [6]int {2, 9, 3, 1, 5, 19}
  fmt.Println(average(array), "\n")
  //task2
  str := []string {"hear", "see", "speak"}
  sequence := []int64 {1, 7, 5, 2, 9, 11, 3}
  fmt.Println(max(str))
  fmt.Println(reverse(sequence))
}
