package main

import "fmt"

func updateName(x string) string {
	x = "dude"
	fmt.Println("Before function call:", x)
	return x
}
func updateMenu(m map[string]float64) {
	m["burger"] = 6.99
	
}

func main() {
	// Group types -> primitive types(A) (int, float, bool, string, struct, array)
	name := "Omar"
	name = updateName(name)
	fmt.Println("After function call:", name)
	// group types -> reference types (B) (slice, map, channel, pointer, function)
	menu := map[string]float64{
		"burger": 5.99,
		"pizza":  8.99,
	}
	updateMenu(menu)
	fmt.Println("After function call:", menu)

}
