package main 
import "fmt"

// func name([name type]*) (return, *)
// can pass in multiple values and get multiple return parameters
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// We can take 0 or more arguments, and also pass a slice
func add(args ... float64) int {
	total := 0.0
	for _, v := range args {
		total += float64(v)
	}
	return int(total)
}

// Returns the next even integer, as a function call
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return 
	}
}

func main() {
	x := make(map[string]int) // x is a map of strings to ints
	x["key"] = 10
	fmt.Println(x["key"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"

	// name -> result of lookup, ok -> if successful
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok) // only prints if successful
	}

	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))
	fmt.Println(add(xs...))
}

