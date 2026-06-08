package main
import(
	"fmt"
)

func main(){
	// for loop
	x := 0 

	for x < 5 {
		fmt.Println("value of x:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i:", i)
	}

	// for range loop
	names := []string{"omar", "ahmed", "hassan", "mohamed"}

	for i := 0 ; i < len(names); i++ {
		fmt.Println( names[i])
	}

	for index, value := range names {
		fmt.Printf("index: %d, name: %s\n", index, value)
	}

	// if you don't need the index you can use _ to ignore it
	for _, value := range names {
		fmt.Printf("the value: %v\n", value)
		value = "changed"
	}
	fmt.Println("names after the loop:", names)
}