package main 
import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if 0 == i % 2 {
			fmt.Println(i)
		} else {
			fmt.Println(10+i)
		}
	}

	i := 3
	switch i {  // Doesn't require break
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
	}

	var x [5]int // Backwards declaration of arrays
	x[4] = 100
	fmt.Println(x)

	var total float64 = 0 // Ensure correct number types
	for i := 0; i < len(x); i++ {
		total += float64(x[i])
	}
	fmt.Println(total / float64(len(x)))

	// Here we iterate over x[0], x[1], ...
	for _, value := range x { // _ is a nonused variable
		total += float64(value)
	}
	fmt.Println(total / float64(len(x)))

	//y := [5]float64{ 98, 93, 77, 82, 83 } // Alternate

	// Slices: mutable arrays
	s := make([]float64, 5, 10) // Underlying array of 10
	t := x[3:5]
	fmt.Println(s)
	fmt.Println(t)
	t = append(t, 32, 23)
	fmt.Println(t)

}
