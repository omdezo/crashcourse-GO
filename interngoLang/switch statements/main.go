package main 
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader)(string,error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input),err

}
func CreateBill() Bill{
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("create a new bill name: ",reader)
	b := newBill(name)
	fmt.Println("create the bill - ", b.name)
	return b
}
func promptOptions(b Bill){
	reader := bufio.NewReader(os.Stdin)

	opt,_ := getInput("choose option(a- add item, s- save bill, t - add tip)", reader)
	switch opt {
	case "a":
		name,_ := getInput ("Item name: ", reader)
		price, _ :=getInput("Item price: ",  reader)
		fmt.Println(name , price)
	case "t":
		tip, _ := getInput("Ener tip amount ($): " , reader)

		fmt.Println(tip)
	case "s":
		fmt.Println("you chose s")
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
	
}

func main(){
	mybill := CreateBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}