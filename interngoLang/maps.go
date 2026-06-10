package main

import "fmt"

func main() {
	// Create a map
	m := make(map[string]float64)

	// Add key-value pairs
	m["apple"] = 4.65
	m["banana"] = 3.22
	// Access values
	fmt.Println(m["apple"]) // Output: 4.65

	// Check if a key exists
	if val, exists := m["banana"]; exists {
		fmt.Println(val) // Output: 3.22
	}

	// Delete a key
	delete(m, "apple")
	fmt.Println(m) // Output: map[banana:3.22]

	for k, v := range m {
		fmt.Println(k, " ", v) // Output: banana 3.22
	
	}

	phoneBook := map[int]string{
		91244525: "Omar",
		91244526: "Ali",
		91244527: "Sara",
	}
	fmt.Println(phoneBook)
	fmt.Println(phoneBook[91244525]) // Output: Omar

	phoneBook[91244528] = "browser"
	fmt.Println(phoneBook)
	 
}