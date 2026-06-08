package main

import "fmt"

func main() {
	// arrays in Go "literal" and  "separate"
	var ages [5]int
	ages[0] = 1
	ages[1] = 2
	ages[2] = 3
	ages[3] = 4
	ages[4] = 5
	var ages2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(ages)
	fmt.Println(ages2)

	// another way to declare and initialize an array
	colors := [3]string{"red", "green", "blue"}
	fmt.Println(colors)

	// getting the length of an array
	fmt.Println("length of colors array:", len(colors))

	//slice [use arrays under the hood but they are more flexible and easier to work with]
	var scores = []int{90, 80, 70}
	scores[2] = 25
	scores = append(scores, 60)

	fmt.Println(scores, len (scores))

	// slices ranges
	rangeone := colors[1:3] // from index 1 to index 2 (not including index 3)
	rangetwo := colors[2:] // from index 2 to index 4 (not including index 5)
	rangethree := colors[:3] // from index 0 to index 2 (not including index 3)
	fmt.Println(rangeone, rangetwo, rangethree)
}