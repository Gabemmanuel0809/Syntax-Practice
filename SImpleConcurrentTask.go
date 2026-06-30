package main 

import(
  "fmt"
)

func numStore(x int, ch chan int) {
	ch <- x 
}

func add(ch chan int, ch2 chan int, ch3 chan int) {
	val := <- ch
	val2 := <- ch2
	result := val + val2 
	ch3 <- result 
}

func multiply(ch chan int, ch2 chan int, ch3 chan int) {
	val := <- ch
	val2 := <- ch2
	result := val * val2 
	ch3 <- result 
}

func main() {
	ch := make(chan int)
	ch1 := make(chan int)
	res1ch := make(chan int)
	res2ch := make(chan int)
	
	
	go numStore(5, ch)
	go numStore(10, ch1)
	
	go add(ch, ch1, res1ch)
	sum := <-res1ch
	fmt.Println("Operation Result: ", sum)
	
	//re-declare cause the data is gone after goroutine reads data out of an unbuffered channel
	go numStore(5, ch)
	go numStore(10, ch1)
	
	go multiply(ch, ch1, res2ch)
	product := <-res2ch 
	fmt.Println("Operation Result: ", product)
}
