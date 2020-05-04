package main

import "fmt"

func average(arr [6]int) float64 {
	var sum int
	for _, a := range arr {
		sum += a
	}
	return float64(sum) / float64(len(arr))
}

func max(sl []string) string {
	longestElement := sl[0]
	for i, a := range sl {
		if len(a) > len(longestElement) {
			longestElement = sl[i]
		}
	}
	return longestElement
}

func reverse(sl []int64) []int64 {
	newSl := make([]int64, len(sl), len(sl))
	for i, a := range sl {
		newSl[len(newSl)-1-i] = a
	}
	return newSl
}

func printSorted(m map [int]string) {
	var	keys []int
	var values []string
	for key, value := range m {
		keys = append(keys, key)
		values = append(values, value)
	}
	for i := 0; i < 2; i++ {
		if keys[i+1] < keys[i] {
			tempKey := keys[i+1]
			tempValue := values[i+1]
			keys[i+1] = keys[i]
			values[i+1] = values[i]
			keys[i] = tempKey
			values[i] = tempValue
		}
	}
	fmt.Println(values)
}
func main() {
	//task 1
	var array = [6]int{2, 9, 3, 1, 5, 19}
	fmt.Println(average(array), "\n")
	//task2
	str := []string{"hear", "see", "speak"}
	sequence := []int64{1, 7, 5, 2, 9, 11, 3}
	fmt.Println(max(str))
	fmt.Println(reverse(sequence), "\n")
	//task3
	myMap := map[int]string{5: "KEKW", 2: "LULW", 4: "4Head"}
	printSorted(myMap)
}
