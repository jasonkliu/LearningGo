package main 
import "fmt"
import "math"

type Circle struct {
	x, y, r float64
}

type Android struct {
	Person  		// don't have to declare a name
	Model string 
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	/*
	defer func() { // run after function completes, even if panic
		str := recover() // stop the panic and returns panic value
		fmt.Println(str)
	}()
	panic("PANIC")
	*/

	x := 5
	zero(&x)	// Find address of variable x
	fmt.Println(x)

	y := new(int) // like malloc
	zero(y)
	fmt.Println(*y)

	c := Circle{0, 0, 5}
	//c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(c.area())

	a := new(Android)
	a.Talk() 	 // also a.Person.Talk()
}

func zero(xPtr *int) {
	*xPtr = 0
}

func (c Circle) area() float64 { // automatically passes pointer
	return math.Pi * c.r * c.r   // function called by c.area()
}

