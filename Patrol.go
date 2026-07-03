package main 

import(
  "fmt"
  "sync"
  "time"
)

type Aircraft struct {
   Name string 
   DC int  // Depth Charge
   LowAlt bool
}

func Dive(a *Aircraft, dived chan<- struct{}, wg *sync.WaitGroup) {
   defer wg.Done()
   
   time.Sleep(time.Second)

   if a.LowAlt == true {
	 fmt.Println("Ready for bombing")
   } else {
	a.LowAlt = true 
	fmt.Println("Diving for America......and ugh also for Britain, Soviet, and the world")
   }

   close(dived)
}

func DropLoad(a *Aircraft, dived <-chan struct{}, done chan<- struct{}, wg *sync.WaitGroup) {
   defer wg.Done()
   <-dived 

   time.Sleep(time.Second)

   if a.LowAlt != true {
	 fmt.Println("Too high to drop payload")
   } else if a.DC == 0 {
	 fmt.Println("Out of Depth Charges")
	  fmt.Println("Watch ads to get free 4 depth charges and to support America")
   } else {
	 fmt.Println("Depth Charge released")
	 a.DC -= 2
	 fmt.Println("plink...plank....Kaboom!")
	 fmt.Println("That explosion is very American")
   }

   close(done)
}

func main() {
   var wg sync.WaitGroup
 
   Dived := make(chan struct{})
   Dropped := make(chan struct{})

   var ctl = Aircraft{
	  Name:"PBY Catalina",
	  DC:6,
	  LowAlt:false,
   }

   wg.Add(2)
      go Dive(&ctl, Dived, &wg)
	  go DropLoad(&ctl, Dived, Dropped, &wg)
   wg.Wait()
}
