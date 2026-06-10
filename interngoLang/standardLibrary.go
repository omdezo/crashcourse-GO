package main
import(
	"fmt"
	"sort"
	"strings"
)

func main(){
	// using the strings package from the standard library
	greeting := "hello there dude! "

	fmt.Println(strings.Contains(greeting, "hello")) 
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting)) 
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))
	// the original value it's not changing
	fmt.Println("original string value =", greeting)

	ages := []int{25, 30, 35, 40, 45}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 35)
	fmt.Println("index of 35 in ages slice:", index)

	names := []string{"omar", "ahmed", "hassan", "mohamed"}

	sort.Strings(names)
	fmt.Println(names)
	
	fmt.Println("index of hassan in names slice:", sort.SearchStrings(names, "hassan"))
}