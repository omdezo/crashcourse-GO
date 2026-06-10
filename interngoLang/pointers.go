package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"

}

func main() {
	
	name := "Omar"
	
	//updateName(name)
	
	fmt .Println("memory address of name variable:", &name) 

	m := &name 
	fmt.Println("memory address of m variable:", *m)
	fmt.Println(name) // Output: Omar
	updateName(m)
	fmt.Println(name) // Output: wedge
}