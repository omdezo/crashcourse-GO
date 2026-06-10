package main 
import "fmt"
func main() {
	mybill := newBill("Omar")
	//fmt.Println(mybill)
	mybill.addItem("coffee",3.25)
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}