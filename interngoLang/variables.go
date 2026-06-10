package main
import "fmt"

func main(){
	//strings
	var nameone string = "omar"
	var nametwo = "ahmed"
	var namethree string

	fmt.Println(nameone, nametwo, namethree)

	nameone = "mohamed"
	namethree = "hassan"
	
	fmt.Println(nameone, nametwo, namethree)

	nameFour := "TRAINING"
	fmt.Println(nameFour)

	//ints 
	var ageone uint = 20
	var agetwo = 30
	agethree := 40

	fmt.Println(ageone, agetwo, agethree)

	//bits & memory
	var numone int8 = 25
	var numtwo int8 = -128
	var numthree uint8 = 25
	fmt.Println(numone, numtwo, numthree)

	var scoreone float32 = 25.5
	var scoretwo float64 = 2523235667378732.535
	Scorethree := 1.5
	fmt.Println(scoreone, scoretwo, Scorethree)
	}