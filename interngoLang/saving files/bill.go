package main
import (
	"fmt"
	"os"
)
type Bill struct {
	
	name string
	items map[string]float64
	tip float64
}
//make a new bill
func newBill(name string) Bill {
	b := Bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
	return b
}
// format the bill 
func (b *Bill) format() string {
	formattedBill := "Bill breakdown:\n"
	var total float64 = 0

	// list items 
	for k,v := range b.items {
		formattedBill += fmt.Sprintf("%-25v ... $%.2f\n", k+ ":", v)
		total += v
	}
	// add tip 
	formattedBill += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)
	//total
	formattedBill += fmt.Sprintf("%-25v ... $%.2f\n", "total:", total + b.tip)
	return formattedBill
}

//update the tip
func (b *Bill) updateTip(tip float64) {
	(*b).tip = tip
}

// add an item to the bill
func (b *Bill) addItem(name string, price float64) {
	b.items[name] = price
}



// save bill 
func (b *Bill) save () {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt",data,0644)
	if err != nil { 
		panic(err)
	}
	fmt.Println("bill was saved to file")
}