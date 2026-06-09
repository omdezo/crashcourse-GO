package main
import(
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("saying hi %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("saying bye %v\n", n)
}

func cycleNames(n []string,f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func cycleArea(r float64 ) float64 {
	return math.Pi * r * r
}

func main() {
	
	//sayGreeting("omar")
	//sayGreeting("ahmed")
	//sayBye("ahmed")

	cycleNames([]string{"omar", "ahmed", "hassan"}, sayGreeting)
	cycleNames([]string{"omar", "ahmed", "hassan"}, sayBye)
	a1 := cycleArea(10.5)
	a2 := cycleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("area of circle1 %.2f\n", a1)
	fmt.Printf("area of circle2 %.2f\n", a2)
}

