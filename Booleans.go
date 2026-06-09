package main

import "fmt"

func main() {
	// IF CONDITIONAL STATEMENTS
	age := 25

	fmt.Println(age<=30)
	fmt.Println(age>=30)
	fmt.Println(age==25)
	fmt.Println(age!=30)

	if age < 30 {
		fmt.Println("you are young")
	} else {
		fmt.Println("you are old")
	}
	names := []string{"omar", "ahmed", "hassan", "mohamed", "sara"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continue at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("break at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v\n", index, value)
	}
}