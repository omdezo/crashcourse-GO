package main
import "fmt"

func main(){

	age := 35
	name := "omar"
	percentage := 94.5

	//printing methods
	fmt.Println("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	fmt.Println("hello ninjas")
	fmt.Println("goodbye ninjas")
	fmt.Println("my age is", age, "and my name is", name)

	// printf (formatted strings) %_ = frormat specifier
	// %v  — default format, Go decides how to print it
	fmt.Printf("name: %v | age: %v | percentage: %v\n", name, age, percentage)

	// %T  — prints the TYPE of the variable
	fmt.Printf("name: %T | age: %T | percentage: %T\n", name, age, percentage)

	// %s  — string as-is, no quotes
	fmt.Printf("name: %s\n", name)

	// %q  — string with double quotes / integer as its unicode character in single quotes
	fmt.Printf("name: %q | age: %q\n", name, age)

	// %d  — integer as decimal number (only works on integers)
	fmt.Printf("age: %d\n", age)

	// %b  — integer as binary
	fmt.Printf("age in binary: %b\n", age)

	// %x  — integer as hexadecimal lowercase
	fmt.Printf("age in hex: %x\n", age)

	// %o  — integer as octal
	fmt.Printf("age in octal: %o\n", age)

	// %f  — float as decimal
	fmt.Printf("percentage: %f\n", percentage)

	// %.2f — float with exactly 2 decimal places
	fmt.Printf("percentage: %.2f\n", percentage)

	// %e  — float in scientific notation
	fmt.Printf("percentage: %e\n", percentage)

	// %t  — boolean, prints true or false (using age here will throw error, only for bool)
	isStudent := true
	fmt.Printf("is student: %t\n", isStudent)

	// %p  — memory address (pointer), use & to get the address of a variable
	fmt.Printf("name address: %p\n", &name)

	// %05d — pad integer with zeros to width 5
	fmt.Printf("age padded: %05d\n", age)

	// %-10s — left align string in a field of width 10
	fmt.Printf("name left aligned: %-10s|\n", name)

	// %10s — right align string in a field of width 10
	fmt.Printf("name right aligned: %10s|\n", name)
	
}