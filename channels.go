package main 
import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))  // sleeping causes it to run
		time.Sleep(time.Millisecond * amt)    // simultaneously with others
	} 
}

func pinger(c chan string) { // attempts to send message to channel
	for i := 0; ; i++ {
		c <- "ping" // send and receive message; send "ping"
	}
}

func printer(c chan string) {  // pinger waits until printer is ready
	for {
		msg := <- c		// receive a message and store it in msg
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger (c chan string) { // use select instead of case to choose
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main() {
	//go f(0) 			// goroutine: execute statements then go to next line
	/*for i := 0; i < 10; i++ {
		go f(i)
	}*/

	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)  // wait for program otherwise it exits */
}

