package main 

import(
	"fmt"
	"sync"
)

type Data struct {
	Username string 
	Password string
}

type Node struct {
	Location string
	Data
}

func sendData(n1 Node, ch chan<- Data) {
     ch <- n1.Data
	 fmt.Println("Sending Data") 
}

func receiveData(ch chan Data, n2 *Node) {
	data := <- ch
	n2.Data = data
	fmt.Println("Data received")
}

func main() {
	var wg sync.WaitGroup
	comms := make(chan Data)
	var n1 = Node {Location:"North America", Data: Data{Username:"Ray", Password:"Ab_C123"}}
	var n2 = Node {Location:"Europe", Data: Data{}}

  wg.Add(2)
	go func() {
		defer wg.Done()
		sendData(n1, comms)
	}()

	go func() {
		defer wg.Done()
		receiveData(comms, &n2)
	}()

	wg.Wait()

	fmt.Println(n2.Data)
}
