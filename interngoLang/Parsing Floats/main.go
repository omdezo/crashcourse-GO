package main 
import (
	"fmt"
	"bufio"
	"strconv"
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

		p, err := strconv.ParseFloat(price, 64)
		if err != nil { 
			fmt.Println ("The price must be a number ")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("item added ", name , price)
	case "t":
		tip, _ := getInput("Ener tip amount ($): " , reader)
		
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil { 
			fmt.Println ("The tip must be a number ")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - " , tip)
		promptOptions(b)

	case "s":
		fmt.Println("you chose to save the bill ", b)
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